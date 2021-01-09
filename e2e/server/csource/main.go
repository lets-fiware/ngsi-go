/*
MIT License

Copyright (c) 2020-2021 Kazuhito Suda

This file is part of NGSI Go

https://github.com/lets-fiware/ngsi-go

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	urlRoot   = "/"
	urlHealth = "/health"
	urlKill   = "/kill"
	urlQuery  = "/v2/op/query"
	urlUpdate = "/v2/op/update"
)

type entityDef map[string]interface{}
type entitiesDef []entityDef

var mutex = &sync.Mutex{}

type entity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type queryDef struct {
	Entities []entity `json:"entities"`
	Attrs    []string `json:"attrs"`
}

var entities entitiesDef

var (
	gHost       = flag.String("host", "0.0.0.0", "host")
	gPort       = flag.String("port", "8000", "port")
	gConfigFile = flag.String("config", "entities.json", "entities file")
)

func main() {
	os.Exit(csourceServer())
}

func csourceServer() int {
	const funcName = "csourceServer"

	printMsg(funcName, 1, "Start csource server")

	flag.Parse()

	if err := loadEntitites(); err != nil {
		printMsg(funcName, 2, err.Error())
		return 1
	}
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	m := http.NewServeMux()

	m.HandleFunc(urlRoot, http.HandlerFunc(rootHandler))
	m.HandleFunc(urlQuery, http.HandlerFunc(queryHandler))
	m.HandleFunc(urlUpdate, http.HandlerFunc(updateHandler))
	m.HandleFunc(urlHealth, http.HandlerFunc(healthHandler))
	m.HandleFunc(urlKill, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
		sig <- syscall.SIGINT
	})

	addr := *gHost + ":" + *gPort
	println(addr)
	server := &http.Server{
		Addr:              addr,
		Handler:           m,
		ReadHeaderTimeout: 30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			printMsg(funcName, 2, err.Error())
		}
	}()

	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		printMsg(funcName, 3, err.Error())
	}

	return 0
}

func loadEntitites() error {
	const funcName = "loadEntitites"

	printMsg(funcName, 1, *gConfigFile)

	b, err := ioutil.ReadFile(*gConfigFile)
	if err != nil {
		printMsg(funcName, 2, err.Error())
		return err
	}

	mutex.Lock()
	err = json.Unmarshal(b, &entities)
	mutex.Unlock()
	if err != nil {
		printMsg(funcName, 3, err.Error())
		return err
	}

	return nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "rootHandler"

	printMsg(funcName, 1, r.URL.Path)

	w.WriteHeader(http.StatusNotFound)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "healthHandler"

	printMsg(funcName, 1, r.URL.Path)

	if r.Method != http.MethodGet {
		printMsg(funcName, 2, "Method not allowed.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "queryHandler"

	printMsg(funcName, 1, r.URL.Path)

	if r.Method != http.MethodPost {
		printMsg(funcName, 2, "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") == false {
		printMsg(funcName, 3, "Content-Type error: "+r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		printMsg(funcName, 4, "Content-Length error:"+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		printMsg(funcName, 5, "Body read error:"+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	printJSONIndent("Request:", body)

	var query queryDef
	if err := json.Unmarshal(body, &query); err != nil {
		printMsg(funcName, 6, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := query.Entities[0].ID
	newEntities := new(entitiesDef)

	mutex.Lock()
	for _, e := range entities {
		if v, ok := e["id"]; ok {
			if i, ok := v.(string); ok {
				if id == i {
					*newEntities = append(*newEntities, e)
				}
			}
		}
	}
	mutex.Unlock()

	if len(*newEntities) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		b, err := json.Marshal(newEntities)
		if err != nil {
			printMsg(funcName, 7, err.Error())
		}
		printJSONIndent("Respose:", b)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "updateHandler"

	printMsg(funcName, 1, r.URL.Path)

	if r.Method != http.MethodPost {
		printMsg(funcName, 2, "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") == false {
		printMsg(funcName, 3, "Content-Type error: "+r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		printMsg(funcName, 4, "Content-Length error:"+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		printMsg(funcName, 5, "Body read error:"+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	printJSONIndent("Request:", body)
}

func printJSONIndent(s string, b []byte) {
	if s != "" {
		fmt.Println(s)
	}
	buf := new(bytes.Buffer)
	if err := json.Indent(buf, b, "", "  "); err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println(buf.String())
}

func printMsg(funcName string, no int, msg string) {
	fmt.Printf(sprintMsg(funcName, no, msg+"\n"))
}

func sprintMsg(funcName string, no int, msg string) string {
	return fmt.Sprintf("%s%03d %s", funcName, no, msg)
}
