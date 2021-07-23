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

type geoProxyParam struct {
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

type geoProxyStat struct {
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

var geoProxyGlobal *geoProxyParam

func geoProxyServer(c *cli.Context) error {
	const funcName = "geoProxy"

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

	geoProxyGlobal = &geoProxyParam{
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
	mux.HandleFunc("/", http.HandlerFunc(geoProxyRootHandler))
	mux.HandleFunc(proxyPath, http.HandlerFunc(geoProxyHandler))
	mux.HandleFunc("/health", http.HandlerFunc(geoProxyHealthHandler))

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

func geoProxyRootHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "geoProxyRootHandler"

	ngsi := geoProxyGlobal.ngsi
	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, r.URL.Path)+"\n")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s not found"}`, r.URL.Path)))
}

func geoProxyHealthHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "geoProxyHealthHandler"

	ngsi := geoProxyGlobal.ngsi

	if r.Method != http.MethodGet {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(geoProxyGetStat())
	}
}

func geoProxyHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "geoProxyHandler"

	status := http.StatusBadRequest
	ngsi := geoProxyGlobal.ngsi
	client := geoProxyGlobal.client
	verbose := geoProxyGlobal.verbose

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 2, r.URL.Path)+"\n")

	u, err := url.Parse(geoProxyGlobal.url.String())
	if err != nil {
		geoProxyResposeError(ngsi, w, status, &ngsiCmdError{funcName, 3, err.Error(), err})
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
		geoProxyGlobal.mutex.Lock()
		key, token, err := ngsi.GetAuthHeader(client)
		geoProxyGlobal.mutex.Unlock()
		if err != nil {
			geoProxyResposeError(ngsi, w, status, &ngsiCmdError{funcName, 4, err.Error(), err})
			return
		}
		headers[key] = token
	}

	err = geoProxySetQueryParam(ngsi, r, u)
	if err != nil {
		geoProxyResposeError(ngsi, w, status, err)
		return
	}

	res, resBody, err := geoProxyGlobal.http.Request("GET", u, headers, nil)
	if err == nil {
		geoProxyGlobal.gLock.Lock()
		geoProxyGlobal.timeSent += 1
		geoProxyGlobal.success += 1
		geoProxyGlobal.gLock.Unlock()

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
		geoProxyResposeError(ngsi, w, status, &ngsiCmdError{funcName, 6, err.Error(), err})
	}
}

func geoProxyResposeError(ngsi *ngsilib.NGSI, w http.ResponseWriter, status int, err error) {
	geoProxyGlobal.gLock.Lock()
	geoProxyGlobal.timeSent += 1
	geoProxyGlobal.failure += 1
	geoProxyGlobal.gLock.Unlock()

	msg := message(err)
	ngsi.Logging(ngsilib.LogErr, msg+"\n")

	body := []byte(fmt.Sprintf(`{"error":"%s"}`, msg))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func geoProxySetQueryParam(ngsi *ngsilib.NGSI, r *http.Request, u *url.URL) error {
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

func geoProxyGetStat() []byte {
	uptime := time.Now().Unix() - geoProxyGlobal.startTime.Unix()

	geoProxyGlobal.gLock.Lock()
	stat := geoProxyStat{
		NgsiGo:   "geoproxy",
		Version:  Version,
		Health:   "OK",
		Orion:    geoProxyGlobal.url.String(),
		Verbose:  geoProxyGlobal.verbose,
		Uptime:   humanizeUptime(uptime),
		Timesent: geoProxyGlobal.timeSent,
		Success:  geoProxyGlobal.success,
		Failure:  geoProxyGlobal.failure,
	}
	geoProxyGlobal.gLock.Unlock()

	b, err := ngsilib.JSONMarshal(stat)
	if err != nil {
		return []byte(`{"ngsi-go":"geoproxy","health":"NG"}`)
	}

	return b
}

func geoProxyHealthCmd(c *cli.Context) error {
	const funcName = "geoProxyHealth"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"geoproxy"})
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
