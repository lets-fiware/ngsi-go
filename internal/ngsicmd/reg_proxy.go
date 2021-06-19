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

package ngsicmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type regProxyParam struct {
	ngsi    *ngsilib.NGSI
	client  *ngsilib.Client
	http    ngsilib.HTTPRequest
	verbose bool
	bearer  bool
	mutex   *sync.Mutex
}

var regProxyGlobal *regProxyParam

func regProxy(c *cli.Context) error {
	const funcName = "regProxy"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker", "csource"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	host := c.String("rhost")
	port := c.String("port")
	addr := host + ":" + port

	path := c.String("url")
	url := addr + path

	if c.Bool("https") {
		if !c.IsSet("key") {
			return &ngsiCmdError{funcName, 3, "no key file provided", nil}
		}
		if !c.IsSet("cert") {
			return &ngsiCmdError{funcName, 4, "no cert file provided", nil}
		}
		url = "https://" + url
	} else {
		url = "http://" + url
	}

	regProxyGlobal = &regProxyParam{
		ngsi:    ngsi,
		client:  client,
		http:    ngsi.HTTP,
		verbose: c.Bool("verbose"),
		bearer:  !(client.Server.IdmType == "thinkingcities"),
		mutex:   &sync.Mutex{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc(path, http.HandlerFunc(regProxyHandler))

	ngsi.Logging(ngsilib.LogInfo, "Start registration proxy: "+url+"\n")

	if c.Bool("https") {
		err = gNetLib.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), mux)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), nil}
		}
	} else {
		err = gNetLib.ListenAndServe(addr, mux)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), nil}
		}
	}

	return nil
}

func regProxyHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyHandler"

	status := http.StatusBadRequest
	ngsi := regProxyGlobal.ngsi
	verbose := regProxyGlobal.verbose
	client := regProxyGlobal.client
	host := client.Server.ServerHost

	switch r.Method {
	default:
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		status = http.StatusMethodNotAllowed

	case http.MethodGet:
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 2, r.URL.Path))
		if r.URL.Path == "/health" {
			w.Header().Set("Content-Type", "Application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(fmt.Sprintf(`{"ngsi-go":{"version":"%s","csource":"%s"}}`, Version, host)))
			return
		} else {
			status = http.StatusBadRequest
		}
	case http.MethodPost:
		body := r.Body
		defer func() { _ = body.Close() }()
		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, body)

		tenant := ""
		scope := ""
		headers := map[string]string{}
		for k := range r.Header {
			switch strings.ToLower(k) {
			case "content-length", "user-agent":
				continue
			case "fiware-service", "ngsild-tenant":
				tenant = r.Header.Get(k)
			case "fiware-servicepath":
				scope = r.Header.Get(k)
			}
			headers[k] = r.Header.Get(k)
		}

		ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("Path:%s, Tenant: %s, Scope: %s\n", r.URL.Path, tenant, scope))

		u, err := url.Parse(host)
		if err != nil {
			ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 3, err.Error()))
			break
		}
		u.Path = path.Join(u.Path, r.URL.Path)

		regProxyGlobal.mutex.Lock()
		token, err := ngsi.GetToken(client)
		regProxyGlobal.mutex.Unlock()
		if err != nil {
			ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 4, err.Error()))
			break
		}
		if regProxyGlobal.bearer {
			headers["Authorization"] = "Bearer " + token
		} else {
			headers["X-Auth-Token"] = token
		}

		b := buf.Bytes()
		if verbose {
			ngsi.Logging(ngsilib.LogInfo, string(b)+"\n")
		}

		res, resBody, err := regProxyGlobal.http.Request("POST", u, headers, b)
		if err == nil {
			for k := range res.Header {
				if strings.ToLower(k) != "content-length" {
					w.Header().Set(k, res.Header.Get(k))
				}
			}
			w.WriteHeader(res.StatusCode)
			_, _ = w.Write(resBody)
			if verbose {
				ngsi.Logging(ngsilib.LogInfo, string(resBody)+"\n")
			}
			return
		}

		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 5, err.Error()))
	}
	w.WriteHeader(status)
}
