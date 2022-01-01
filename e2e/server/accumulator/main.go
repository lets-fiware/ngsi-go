/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"sync"
	"syscall"
	"time"
)

const (
	urlRoot   = "/"
	urlHealth = "/health"
	urlKill   = "/kill"
	urlAcc    = "/acc"
	urlClear  = "/clear/"
	urlDump   = "/dump/"
	urlGet    = "/get/"
)

var allNotifications sync.Map

var (
	gHost = flag.String("host", "0.0.0.0", "host")
	gPort = flag.String("port", "8000", "port")
)

func main() {
	os.Exit(accumulator())
}

func accumulator() int {
	const funcName = "accumulator"

	printMsg(funcName, 1, "Start accumulator")

	flag.Parse()

	allNotifications = sync.Map{}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	m := http.NewServeMux()

	m.HandleFunc(urlRoot, http.HandlerFunc(rootHandler))
	m.HandleFunc(urlAcc, http.HandlerFunc(accHandler))
	m.HandleFunc(urlClear, http.HandlerFunc(clearHandler))
	m.HandleFunc(urlDump, http.HandlerFunc(dumpHandler))
	m.HandleFunc(urlGet, http.HandlerFunc(getHandler))
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

func accHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "accHandler"

	var err error

	printMsg(funcName, 1, r.URL.Path)

	status := http.StatusNoContent

	switch r.Method {
	default:
		printMsg(funcName, 2, "Method not allowed.")
		status = http.StatusMethodNotAllowed
	case http.MethodPost:
		body := r.Body
		defer func() { setNewError(funcName, 3, body.Close(), &err) }()
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, body)
		if err != nil {
			printMsg(funcName, 3, err.Error())
			status = http.StatusInternalServerError
			break
		}
		b := buf.Bytes()

		var j map[string]interface{}
		err = json.Unmarshal(b, &j)
		if err != nil {
			printMsg(funcName, 3, err.Error())
		} else {
			if id, ok := j["subscriptionId"].(string); ok {
				subs := [][]byte{}
				if v, ok := allNotifications.LoadAndDelete(id); ok {
					subs = v.([][]byte)
				}
				allNotifications.Store(id, subs)
				status = http.StatusNoContent
			} else {
				printMsg(funcName, 4, "subscription id not found")
			}
		}
	}
	w.WriteHeader(status)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "clearHandler"

	printMsg(funcName, 1, r.URL.Path)

	id := r.URL.Path[len(urlClear):]

	if id == "" {
		allNotifications = sync.Map{}
		printMsg(funcName, 2, "all subs clear")
	} else {
		if _, ok := allNotifications.LoadAndDelete(id); ok {
			printMsg(funcName, 3, id+"deleted")
		} else {
			printMsg(funcName, 4, id+" not found")
		}
	}
}

func dumpHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "dumpHandler"

	printMsg(funcName, 1, r.URL.Path)

	switch r.Method {
	default:
		printMsg(funcName, 2, "Method not allowed.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodGet:
		id := r.URL.Path[len(urlDump):]
		payload := getNotifications(id, false)
		if payload == "" {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(payload))
			if err != nil {
				printMsg(funcName, 3, err.Error())
			}
		}
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "getHandler"

	printMsg(funcName, 1, r.URL.Path)

	switch r.Method {
	default:
		printMsg(funcName, 2, "Method not allowed.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodGet:
		id := r.URL.Path[len(urlGet):]
		payload := getNotifications(id, true)
		if payload == "" {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(payload))
			if err != nil {
				printMsg(funcName, 3, err.Error())
			}
		}
	}
}

func getNotifications(id string, clear bool) string {
	const funcName = "getNotifications"

	var payload = ""

	if id == "" {
		allNotifications.Range(func(key interface{}, value interface{}) bool {
			n := value.([][]byte)
			for _, s := range n {
				b := new(bytes.Buffer)
				if err := json.Indent(b, s, "", "  "); err != nil {
					printMsg(funcName, 1, err.Error())
				} else {
					payload = payload + b.String() + "\n"
				}
			}
			return true
		})
		if clear {
			allNotifications = sync.Map{}
		}
	} else {
		if v, ok := allNotifications.Load(id); ok {
			n := v.([][]byte)
			for _, s := range n {
				b := new(bytes.Buffer)
				if err := json.Indent(b, s, "", "  "); err != nil {
					printMsg(funcName, 2, err.Error())
				} else {
					payload = payload + b.String() + "\n"
				}
			}
			if clear {
				allNotifications.LoadAndDelete(id)
			}
		}
	}

	return payload
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
