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
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	urlRoot   = "/"
	urlHealth = "/health"
	urlKill   = "/kill"
)

var gContext sync.Map

var (
	gHost = flag.String("host", "0.0.0.0", "host")
	gPort = flag.String("port", "8000", "port")
	gDir  = flag.String("dir", "", "context directory")
)

func main() {
	os.Exit(contextServer())
}

func contextServer() int {
	const funcName = "contextServer"

	printMsg(funcName, 1, "Start context server")

	flag.Parse()

	gContext = sync.Map{}

	err := loadContext()
	if err != nil {
		printMsg(funcName, 2, err.Error())
		return 1
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	m := http.NewServeMux()
	m.HandleFunc(urlRoot, http.HandlerFunc(rootHandler))
	m.HandleFunc(urlHealth, http.HandlerFunc(healthHandler))
	m.HandleFunc(urlKill, func(w http.ResponseWriter, r *http.Request) {
		const funcName = "killHandler"
		printMsg(funcName, 3, r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
		sig <- syscall.SIGINT
	})

	addr := *gHost + ":" + *gPort
	printMsg(funcName, 4, addr)
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
			printMsg(funcName, 5, err.Error())
		}
	}()

	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		printMsg(funcName, 6, err.Error())
	}

	return 0
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "rootHandler"

	printMsg(funcName, 1, r.URL.Path)

	switch r.Method {
	default:
		fmt.Println("Method not allowed.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodGet:
		if v, ok := gContext.Load(r.URL.Path[1:]); ok {
			w.Header().Set("Content-Type", "application/ld+json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(v.([]byte))
			if err != nil {
				printMsg(funcName, 1, err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPost:
		body := r.Body
		defer body.Close()
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		status := storeContext(r.URL.Path[1:], buf.Bytes())
		w.WriteHeader(status)
	case http.MethodDelete:
		status := deleteContext(r.URL.Path[1:])
		w.WriteHeader(status)
	}
}

func storeContext(name string, b []byte) int {
	const funcName = "storeContext"

	if json.Valid(b) == false {
		printMsg(funcName, 1, "json error: "+string(b))
		return http.StatusBadRequest
	}
	gContext.Store(name, b)

	return http.StatusCreated
}

func deleteContext(name string) int {
	const funcName = "deleteContext"

	if _, ok := gContext.Load(name); ok {
		gContext.Delete(name)
	} else {
		printMsg(funcName, 1, name+" not found")
		return http.StatusBadRequest
	}

	return http.StatusNoContent
}

func loadContext() error {
	const funcName = "loadContext"

	if *gDir == "" {
		return nil
	}

	printMsg(funcName, 1, "dir: "+*gDir)

	files, err := ioutil.ReadDir(*gDir)
	if err != nil {
		printMsg(funcName, 2, "ReadDir: "+err.Error())
		return err
	}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".jsonld") {
			fname := f.Name()
			fmt.Println("load: " + fname)
			b, err := ioutil.ReadFile(filepath.Join(*gDir, fname))
			if err != nil {
				printMsg(funcName, 3, "ReadFile: "+err.Error())
				return err
			}
			if json.Valid(b) == false {
				printMsg(funcName, 4, "json error")
				return err
			}
			gContext.Store(fname, b)
		}
	}
	return nil
}

func jsonIndent(b []byte) ([]byte, error) {
	const funName = "jsonIndent"

	buf := new(bytes.Buffer)
	if err := json.Indent(buf, b, "", "  "); err != nil {
		printMsg(funName, 1, "json.Indent: "+err.Error())
		return nil, err
	}
	return buf.Bytes(), nil
}

func printMsg(funcName string, no int, msg string) {
	fmt.Printf(sprintMsg(funcName, no, msg+"\n"))
}

func sprintMsg(funcName string, no int, msg string) string {
	return fmt.Sprintf("%s%03d %s", funcName, no, msg)
}
