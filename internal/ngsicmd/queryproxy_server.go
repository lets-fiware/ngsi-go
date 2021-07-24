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
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type queryProxyParam struct {
	ngsi    *ngsilib.NGSI
	url     *url.URL
	client  *ngsilib.Client
	http    ngsilib.HTTPRequest
	verbose bool
	mutex   *sync.Mutex
	gLock   *sync.Mutex

	startTime time.Time
	timeSent  int64
	success   int64
	failure   int64
}

type queryProxyStat struct {
	NgsiGo   string `json:"ngsi-go"`
	Version  string `json:"version"`
	Health   string `json:"health"`
	Orion    string `json:"orion"`
	Verbose  bool   `json:"verbose"`
	Uptime   string `json:"uptime"`
	Timesent int64  `json:"timesent"`
	Success  int64  `json:"success"`
	Failure  int64  `json:"failure"`
}

var queryProxyGlobal *queryProxyParam

func queryProxyServer(c *cli.Context) error {
	const funcName = "queryProxy"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClientSkipGetToken(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	host := c.String("ghost")
	port := c.String("port")
	addr := host + ":" + port
	proxyUrl := addr

	if c.Bool("https") {
		if !c.IsSet("key") {
			return &ngsiCmdError{funcName, 3, "no key file provided", nil}
		}
		if !c.IsSet("cert") {
			return &ngsiCmdError{funcName, 4, "no cert file provided", nil}
		}
		proxyUrl = "https://" + proxyUrl
	} else {
		proxyUrl = "http://" + proxyUrl
	}

	u, _ := url.Parse(client.URL.String())
	u.Path = path.Join(u.Path, "/v2/entities")

	queryProxyGlobal = &queryProxyParam{
		ngsi:      ngsi,
		url:       u,
		client:    client,
		http:      ngsi.HTTP,
		verbose:   c.Bool("verbose"),
		mutex:     &sync.Mutex{},
		gLock:     &sync.Mutex{},
		startTime: time.Now(),
	}

	proxyPath := "/v2/ex/entities"
	if c.IsSet("replaceURL") {
		proxyPath = c.String("replaceURL")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(queryProxyRootHandler))
	mux.HandleFunc(proxyPath, http.HandlerFunc(queryProxyHandler))
	mux.HandleFunc("/health", http.HandlerFunc(queryProxyHealthHandler))

	ngsi.Logging(ngsilib.LogErr, "Start geo proxy: "+proxyUrl+"\n")

	if c.Bool("https") {
		err = gNetLib.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), mux)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	} else {
		err = gNetLib.ListenAndServe(addr, mux)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	}

	return nil
}

func queryProxyRootHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "queryProxyRootHandler"

	ngsi := queryProxyGlobal.ngsi
	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, r.URL.Path)+"\n")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s not found"}`, r.URL.Path)))
}

func queryProxyHealthHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "queryProxyHealthHandler"

	ngsi := queryProxyGlobal.ngsi

	if r.Method != http.MethodGet {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(queryProxyGetStat())
	}
}

func queryProxyHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "queryProxyHandler"

	status := http.StatusBadRequest
	ngsi := queryProxyGlobal.ngsi
	client := queryProxyGlobal.client
	verbose := queryProxyGlobal.verbose

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 2, r.URL.Path)+"\n")

	u, err := url.Parse(queryProxyGlobal.url.String())
	if err != nil {
		queryProxyResposeError(ngsi, w, status, &ngsiCmdError{funcName, 3, err.Error(), err})
		return
	}

	tenant := ""
	scope := ""

	headers := map[string]string{}
	for k := range r.Header {
		switch strings.ToLower(k) {
		case "content-length", "user-agent", "content-type":
			continue
		case "fiware-service", "ngsild-tenant":
			tenant = r.Header.Get(k)
			headers[k] = tenant
			continue
		case "fiware-servicepath":
			scope = r.Header.Get(k)
			headers[k] = scope
			continue
		}
		headers[k] = r.Header.Get(k)
	}
	if client.Server.IdmType != "" {
		queryProxyGlobal.mutex.Lock()
		key, token, err := ngsi.GetAuthHeader(client)
		queryProxyGlobal.mutex.Unlock()
		if err != nil {
			queryProxyResposeError(ngsi, w, status, &ngsiCmdError{funcName, 4, err.Error(), err})
			return
		}
		headers[key] = token
	}

	err = queryProxySetQueryParam(ngsi, r, u)
	if err != nil {
		queryProxyResposeError(ngsi, w, status, err)
		return
	}

	res, resBody, err := queryProxyGlobal.http.Request("GET", u, headers, nil)
	if err == nil {
		queryProxyGlobal.gLock.Lock()
		queryProxyGlobal.timeSent += 1
		queryProxyGlobal.success += 1
		queryProxyGlobal.gLock.Unlock()

		for k := range res.Header {
			if strings.ToLower(k) != "content-length" {
				w.Header().Set(k, res.Header.Get(k))
			}
		}
		w.WriteHeader(res.StatusCode)
		_, _ = w.Write(resBody)
		if verbose {
			ngsi.Logging(ngsilib.LogInfo, sprintMsg(funcName, 5, fmt.Sprintf("%d %s\n", res.StatusCode, string(resBody))))
		}
		return
	} else {
		queryProxyResposeError(ngsi, w, status, &ngsiCmdError{funcName, 6, err.Error(), err})
	}
}

func queryProxyResposeError(ngsi *ngsilib.NGSI, w http.ResponseWriter, status int, err error) {
	queryProxyGlobal.gLock.Lock()
	queryProxyGlobal.timeSent += 1
	queryProxyGlobal.failure += 1
	queryProxyGlobal.gLock.Unlock()

	msg := message(err)
	ngsi.Logging(ngsilib.LogErr, msg+"\n")

	body := []byte(fmt.Sprintf(`{"error":"%s"}`, msg))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func queryProxySetQueryParam(ngsi *ngsilib.NGSI, r *http.Request, u *url.URL) error {
	const funcName = "tokeProxyRequestToken"

	ctype := r.Header["Content-Type"]
	if ctype == nil || (ctype != nil && len(ctype) == 0) {
		return &ngsiCmdError{funcName, 1, "missing Content-Type", nil}
	}

	q := u.Query()

	if ctype[0] == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
		for k, v := range r.Form {
			q.Set(k, v[0])
		}
	} else {
		return &ngsiCmdError{funcName, 3, "Content-Type error", nil}
	}

	u.RawQuery = q.Encode()

	return nil
}

func queryProxyGetStat() []byte {
	uptime := time.Now().Unix() - queryProxyGlobal.startTime.Unix()

	queryProxyGlobal.gLock.Lock()
	stat := queryProxyStat{
		NgsiGo:   "queryproxy",
		Version:  Version,
		Health:   "OK",
		Orion:    queryProxyGlobal.url.String(),
		Verbose:  queryProxyGlobal.verbose,
		Uptime:   humanizeUptime(uptime),
		Timesent: queryProxyGlobal.timeSent,
		Success:  queryProxyGlobal.success,
		Failure:  queryProxyGlobal.failure,
	}
	queryProxyGlobal.gLock.Unlock()

	b, err := ngsilib.JSONMarshal(stat)
	if err != nil {
		return []byte(`{"ngsi-go":"queryproxy","health":"NG"}`)
	}

	return b
}

func queryProxyHealthCmd(c *cli.Context) error {
	const funcName = "queryProxyHealth"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"queryproxy"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	client.SetPath("/health")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))
	return nil
}
