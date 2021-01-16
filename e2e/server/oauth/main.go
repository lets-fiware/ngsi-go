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
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	urlRoot             = "/"
	urlHealth           = "/health"
	urlKill             = "/kill"
	urlAdmin            = "/admin/"
	keyrock             = "/keyrock"
	passwordCredentials = "/PasswordCredentials"
	tokenProvider       = "/tokenprovider"
	tokenProxy          = "/tokenproxy"
)

type tokenInfo struct {
	GrantType    string `json:"grantType"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Token        interface{}
}

var gTokens sync.Map

var (
	gHost        = flag.String("host", "0.0.0.0", "host")
	gPort        = flag.String("port", "8000", "port")
	gCconfigFile = flag.String("config", "", "config file")
)

func main() {
	os.Exit(oauthServer())
}

func oauthServer() int {
	const funcName = "oauthServer"

	printMsg(funcName, 1, "Start oauth server")

	flag.Parse()

	if err := readTokens(*gCconfigFile); err != nil {
		printMsg(funcName, 2, err.Error())
		return 1
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	m := http.NewServeMux()

	m.HandleFunc(urlRoot, http.HandlerFunc(oauthHandler))
	m.HandleFunc(urlAdmin, http.HandlerFunc(adminHandler))
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

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "oauthHandler"

	var err error

	printMsg(funcName, 1, r.URL.Path)

	if r.Method != http.MethodPost {
		printMsg(funcName, 2, "Method not allowed.")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	status := http.StatusNotFound

	body := r.Body
	defer func() { setNewError(funcName, 3, body.Close(), &err) }()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, body)
	if err != nil {
		printMsg(funcName, 4, err.Error())
	}

	switch r.URL.Path {
	case keyrock, passwordCredentials, tokenProvider:
		params := parseParam(buf.String())
		var t []byte
		t, err = getToken(params["username"])
		if err != nil {
			printMsg(funcName, 5, err.Error())
			w.WriteHeader(status)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(t)
		}
	case tokenProxy:
		var param struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err = json.Unmarshal(buf.Bytes(), &param)
		if err != nil {
			printMsg(funcName, 6, err.Error())
			w.WriteHeader(status)
		}
		var t []byte
		t, err = getToken(param.Username)
		if err != nil {
			printMsg(funcName, 7, err.Error())
			w.WriteHeader(status)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(t)
		}
	default:
		fmt.Println("url not found")
		w.WriteHeader(status)
	}
}

func parseParam(d string) map[string]string {
	params := make(map[string]string)
	for _, p := range strings.Split(d, "&") {
		q := strings.Split(p, "=")
		if len(q) > 1 {
			params[q[0]] = q[1]
		}
	}
	return params
}

func readTokens(fileName string) error {
	const funcName = "readTokens"

	if fileName == "" {
		return nil
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		printMsg(funcName, 1, err.Error())
		return err
	}

	var tokens map[string]tokenInfo

	err = json.Unmarshal(b, &tokens)
	if err != nil {
		printMsg(funcName, 2, err.Error())
		return err
	}

	for k, v := range tokens {
		gTokens.Store(k, v)
	}

	return nil
}

func getToken(name string) ([]byte, error) {
	const funcName = "getToken"

	if t, ok := gTokens.Load(name); ok {
		info := t.(tokenInfo)
		b, err := json.Marshal(&info.Token)
		if err != nil {
			printMsg(funcName, 1, err.Error())
			return nil, err
		}
		return b, nil
	}
	return nil, errors.New("token not found: " + name)
}

func getTokenInfo(id string) ([]byte, error) {
	const funcName = "getTokenInfo"

	var b []byte
	var err error

	if id == "" {
		var tokens map[string]interface{}
		b, err = json.Marshal(&tokens)
		if err != nil {
			printMsg(funcName, 1, err.Error())
			return nil, err
		}
		for k, v := range tokens {
			gTokens.Store(k, v)
		}

	} else {
		if t, ok := gTokens.Load(id); ok {
			b, err = json.Marshal(&t)
			if err != nil {
				printMsg(funcName, 2, err.Error())
				return nil, err
			}
		} else {
			return nil, errors.New(id + " not found.")
		}
	}
	buf := new(bytes.Buffer)
	if err := json.Indent(buf, b, "", "  "); err != nil {
		printMsg(funcName, 3, err.Error())
	}
	payload := buf.String() + "\n"

	return []byte(payload), nil
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "adminHandler"

	printMsg(funcName, 1, r.URL.Path)

	id := r.URL.Path[len(urlAdmin):]

	switch r.Method {
	default:
		printMsg(funcName, 2, "Method not allowed.")
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodPost:
		if id != "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body := r.Body
		defer func() { _ = body.Close() }()
		buf := new(bytes.Buffer)
		_, err := io.Copy(buf, body)
		if err != nil {
			printMsg(funcName, 3, err.Error())
		}

		var t tokenInfo
		err = json.Unmarshal(buf.Bytes(), &t)
		if err != nil {
			printMsg(funcName, 4, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if t.Username == "" {
			printMsg(funcName, 5, "username == \"\"")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		gTokens.Store(t.Username, t)
		w.WriteHeader(http.StatusNoContent)
	case http.MethodGet:
		b, err := getTokenInfo(id)
		if err != nil {
			printMsg(funcName, 6, err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	case http.MethodDelete:
		if id == "" {
			gTokens = sync.Map{}
		} else {
			if _, ok := gTokens.Load(id); ok {
				gTokens.Delete(id)
			} else {
				printMsg(funcName, 7, id+" not found")
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func printMsg(funcName string, no int, msg string) {
	fmt.Println(sprintMsg(funcName, no, msg))
}

func sprintMsg(funcName string, no int, msg string) string {
	return fmt.Sprintf("%s%03d %s", funcName, no, msg)
}

func setNewError(funcName string, num int, newErr error, err *error) {
	if *err == nil && newErr != nil {
		*err = errors.New(sprintMsg(funcName, num, newErr.Error()))
	}
}
