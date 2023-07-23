/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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
	"errors"
	"flag"
	"fmt"
	"io"
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
	urlRoot     = "/"
	urlHealth   = "/health"
	urlKill     = "/kill"
	urlRegister = "/register"
	urlQuery    = "/v2/op/query"
	urlUpdate   = "/v2/op/update"
)

type entityDef map[string]interface{}
type entitiesDef map[string]entityDef      // key is entity id
type allEntitiesDef map[string]entitiesDef // key is entity type

var mutex = &sync.Mutex{}

type entity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type queryDef struct {
	Entities []entity `json:"entities"`
	Attrs    []string `json:"attrs"`
}

var allEntities allEntitiesDef

var (
	gHost       = flag.String("host", "0.0.0.0", "host")
	gPort       = flag.String("port", "8000", "port")
	gConfigFile = flag.String("config", "", "entities file")
)

func main() {
	os.Exit(csourceServer())
}

func csourceServer() int {
	const funcName = "csourceServer"

	printMsg(funcName, 1, "Start csource server")

	flag.Parse()

	allEntities = make(map[string]entitiesDef)

	if err := loadEntitites(); err != nil {
		printMsg(funcName, 2, err.Error())
		return 1
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	m := http.NewServeMux()

	m.HandleFunc(urlRoot, http.HandlerFunc(rootHandler))
	m.HandleFunc(urlQuery, http.HandlerFunc(queryHandler))
	m.HandleFunc(urlUpdate, http.HandlerFunc(updateHandler))
	m.HandleFunc(urlRegister, http.HandlerFunc(registerHandler))
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

	if *gConfigFile == "" {
		return nil
	}

	printMsg(funcName, 1, *gConfigFile)

	b, err := os.ReadFile(*gConfigFile)
	if err != nil {
		printMsg(funcName, 2, err.Error())
		return err
	}

	return storeEntities(b)
}

func storeEntities(b []byte) error {
	const funcName = "storeEntities"

	var entities []entityDef

	err := json.Unmarshal(b, &entities)
	if err != nil {
		printMsg(funcName, 1, err.Error())
		return err
	}

	for _, e := range entities {
		t, ok := e["type"].(string)
		if !ok {
			s := sprintMsg(funcName, 2, "entity type error")
			printJSONIndent(s, b)
			return errors.New(s)
		}
		if t == "" {
			s := sprintMsg(funcName, 3, "entity type is empty")
			printJSONIndent(s, b)
			return errors.New(s)
		}
		i, ok := e["id"].(string)
		if !ok {
			s := sprintMsg(funcName, 4, "entity id error")
			printJSONIndent(s, b)
			return errors.New(s)
		}
		if i == "" {
			s := sprintMsg(funcName, 5, "entity id is empty")
			printJSONIndent(s, b)
			return errors.New(s)
		}
		mutex.Lock()
		if allEntities[t] == nil {
			allEntities[t] = make(entitiesDef)
		}
		allEntities[t][i] = e
		mutex.Unlock()
	}

	return nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "registerhandler"

	var err error

	printMsg(funcName, 1, r.URL.Path)

	if r.Method != http.MethodPost {
		printMsg(funcName, 2, "Method not allowed.")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body := r.Body
	defer func() { setNewError(funcName, 3, body.Close(), &err) }()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, body)
	if err != nil {
		printMsg(funcName, 4, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = storeEntities(buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
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
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
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
	_, err = r.Body.Read(body)
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
	t := query.Entities[0].Type

	newEntities := []entityDef{}

	mutex.Lock()
	if entities, ok := allEntities[t]; ok {
		if e, ok := entities[id]; ok {
			newEntities = append(newEntities, e)
		}
	}
	mutex.Unlock()

	if len(newEntities) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		b, err := json.Marshal(newEntities)
		if err != nil {
			printMsg(funcName, 7, err.Error())
		}
		printJSONIndent("Respose:", b)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(b)
		if err != nil {
			printMsg(funcName, 8, err.Error())
		}
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
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
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
	_, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		printMsg(funcName, 5, "Body read error:"+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	printJSONIndent("Request:", body)

	w.WriteHeader(http.StatusBadRequest)
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
	fmt.Println(sprintMsg(funcName, no, msg+"\n"))
}

func sprintMsg(funcName string, no int, msg string) string {
	return fmt.Sprintf("%s%03d %s", funcName, no, msg)
}

func setNewError(funcName string, num int, newErr error, err *error) {
	if *err == nil && newErr != nil {
		*err = errors.New(sprintMsg(funcName, num, newErr.Error()))
	}
}
