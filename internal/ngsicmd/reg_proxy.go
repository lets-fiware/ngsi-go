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
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type regProxyParam struct {
	ngsi     *ngsilib.NGSI
	client   *ngsilib.Client
	http     ngsilib.HTTPRequest
	verbose  bool
	bearer   bool
	tenant   *string
	scope    *string
	addScope *string
	url      *string
	replace  bool
	mutex    *sync.Mutex
	gLock    *sync.Mutex

	startTime time.Time
	timeSent  int64
	success   int64
	failure   int64
}

type regProxyReplace struct {
	Verbose *bool   `json:"verbose,omitempty"`
	Service *string `json:"service,omitempty"`
	Path    *string `json:"path,omitempty"`
	AddPath *string `json:"add_path,omitempty"`
	URL     *string `json:"url,omitempty"`
}

type regProxyStat struct {
	NgsiGo   string           `json:"ngsi-go"`
	Version  string           `json:"version"`
	Health   string           `json:"health"`
	Csource  string           `json:"csource"`
	Verbose  bool             `json:"verbose"`
	Uptime   string           `json:"uptime"`
	Timesent int64            `json:"timesent"`
	Success  int64            `json:"success"`
	Failure  int64            `json:"failure"`
	Replace  *regProxyReplace `json:"replace,omitempty"`
}

var regProxyGlobal *regProxyParam

func regProxyServer(c *cli.Context) error {
	const funcName = "regProxy"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClientSkipGetToken(ngsi, c, false, []string{"broker", "csource"})
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
		ngsi:      ngsi,
		client:    client,
		http:      ngsi.HTTP,
		verbose:   c.Bool("verbose"),
		mutex:     &sync.Mutex{},
		gLock:     &sync.Mutex{},
		startTime: time.Now(),
	}

	if c.IsSet("replaceService") {
		tenant := c.String("replaceService")
		regProxyGlobal.tenant = &tenant
		regProxyGlobal.replace = true
	}
	if c.IsSet("replacePath") {
		scope := c.String("replacePath")
		regProxyGlobal.scope = &scope
		regProxyGlobal.replace = true
	}
	if c.IsSet("addPath") {
		addScope := c.String("addPath")
		regProxyGlobal.addScope = &addScope
		regProxyGlobal.replace = true
	}
	if c.IsSet("replaceURL") {
		repalceURL := c.String("replaceURL")
		regProxyGlobal.url = &repalceURL
		regProxyGlobal.replace = true
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(regProxyRootHandler))
	mux.HandleFunc("/v1/", http.HandlerFunc(regProxyHandler))
	mux.HandleFunc("/v2/", http.HandlerFunc(regProxyHandler))
	mux.HandleFunc("/health", http.HandlerFunc(regProxyHealthHandler))
	mux.HandleFunc("/config", http.HandlerFunc(regProxyConfigHandler))

	ngsi.Logging(ngsilib.LogInfo, "Start registration proxy: "+url+"\n")

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

func regProxyRootHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyRootHandler"

	ngsi := regProxyGlobal.ngsi
	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, r.URL.Path)+"\n")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s not found"}`, r.URL.Path)))
}

func regProxyHealthHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyHealthHandler"

	ngsi := regProxyGlobal.ngsi

	if r.Method != http.MethodGet {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(regProxyGetStat(regProxyGlobal.client.Server.ServerHost))
	}
}

func regProxyConfigHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyConfigHandler"

	ngsi := regProxyGlobal.ngsi

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.Header().Set("Content-Type", "Application/json")
		b := getRequestBody(r.Body)
		status, body := regProxyConfig(ngsi, b)
		w.WriteHeader(status)
		_, _ = w.Write(body)
	}
}

func regProxyHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyHandler"

	status := http.StatusBadRequest
	ngsi := regProxyGlobal.ngsi
	verbose := regProxyGlobal.verbose
	client := regProxyGlobal.client
	host := client.Server.ServerHost

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, r.URL.Path)+"\n")

	b := getRequestBody(r.Body)

	origTenant := ""
	origScope := ""
	tenant := ""
	scope := ""
	headers := map[string]string{}

	for k := range r.Header {
		switch strings.ToLower(k) {
		case "content-length", "user-agent":
			continue
		case "fiware-service", "ngsild-tenant":
			tenant = r.Header.Get(k)
			origTenant = tenant
			if regProxyGlobal.tenant != nil {
				tenant = *regProxyGlobal.tenant
			}
			headers[k] = tenant
			continue
		case "fiware-servicepath":
			scope = r.Header.Get(k)
			origScope = scope
			if regProxyGlobal.scope != nil {
				scope = *regProxyGlobal.scope
			}
			headers[k] = scope
			continue
		}
		headers[k] = r.Header.Get(k)
	}

	if regProxyGlobal.addScope != nil {
		scope = path.Join(*regProxyGlobal.addScope, scope)
	}

	uPath := r.URL.Path
	if regProxyGlobal.url != nil {
		uPath = *regProxyGlobal.url
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("Path:%s, Tenant: %s, Scope: %s\n", r.URL.Path, origTenant, origScope))
	if regProxyGlobal.replace {
		ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("Path:%s, Tenant: %s, Scope: %s\n", uPath, tenant, scope))
	}

	u, err := url.Parse(host)
	if err != nil {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 4, err.Error()))
		regProxyFailureUp()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u.Path = path.Join(u.Path, uPath)

	regProxyGlobal.mutex.Lock()
	key, token, err := ngsi.GetAuthHeader(client)
	regProxyGlobal.mutex.Unlock()

	if err != nil {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 5, err.Error()))
		regProxyFailureUp()
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	headers[key] = token

	if verbose {
		ngsi.Logging(ngsilib.LogInfo, string(b)+"\n")
	}

	res, resBody, err := regProxyGlobal.http.Request("POST", u, headers, b)
	if err == nil {
		regProxyGlobal.gLock.Lock()
		regProxyGlobal.timeSent += 1
		regProxyGlobal.success += 1
		regProxyGlobal.gLock.Unlock()

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
	} else {
		regProxyFailureUp()
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 6, err.Error()))

	w.WriteHeader(status)
}

func getRequestBody(body io.ReadCloser) []byte {
	defer func() { _ = body.Close() }()
	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, body)
	return buf.Bytes()
}

func regProxyFailureUp() {
	regProxyGlobal.gLock.Lock()
	regProxyGlobal.timeSent += 1
	regProxyGlobal.failure += 1
	regProxyGlobal.gLock.Unlock()
}

func regProxyGetStat(host string) []byte {
	uptime := time.Now().Unix() - regProxyGlobal.startTime.Unix()

	regProxyGlobal.gLock.Lock()
	stat := regProxyStat{
		NgsiGo:   "regproxy",
		Version:  Version,
		Health:   "OK",
		Csource:  host,
		Verbose:  regProxyGlobal.verbose,
		Uptime:   humanizeUptime(uptime),
		Timesent: regProxyGlobal.timeSent,
		Success:  regProxyGlobal.success,
		Failure:  regProxyGlobal.failure,
	}

	if regProxyGlobal.replace {
		stat.Replace = &regProxyReplace{
			Service: regProxyGlobal.tenant,
			Path:    regProxyGlobal.scope,
			AddPath: regProxyGlobal.addScope,
			URL:     regProxyGlobal.url,
		}
	}
	regProxyGlobal.gLock.Unlock()

	b, err := ngsilib.JSONMarshal(stat)
	if err != nil {
		return []byte(`{"ngsi-go":"regproxy","health":"NG"}`)
	}

	return b
}

func regProxyConfig(ngsi *ngsilib.NGSI, body []byte) (int, []byte) {
	const funcName = "regProxyConfig"

	req := &regProxyReplace{}

	err := ngsilib.JSONUnmarshal(body, &req)
	if err != nil {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, err.Error()+"\n"))
		return http.StatusBadRequest, []byte(`{"error":"` + err.Error() + `"}`)
	}

	if req.Verbose != nil {
		regProxyGlobal.verbose = *req.Verbose
	}
	if req.Service != nil {
		if *req.Service == "" {
			regProxyGlobal.tenant = nil
		} else {
			regProxyGlobal.tenant = req.Service
		}
	}
	if req.Path != nil {
		if *req.Path == "" {
			regProxyGlobal.scope = nil
		} else {
			regProxyGlobal.scope = req.Path
		}
	}
	if req.AddPath != nil {
		if *req.AddPath == "" {
			regProxyGlobal.addScope = nil
		} else {
			regProxyGlobal.addScope = req.AddPath
		}
	}
	if req.URL != nil {
		if *req.URL == "" {
			regProxyGlobal.url = nil
		} else {
			regProxyGlobal.url = req.URL
		}
	}

	if regProxyGlobal.tenant == nil &&
		regProxyGlobal.scope == nil &&
		regProxyGlobal.addScope == nil &&
		regProxyGlobal.url == nil {
		regProxyGlobal.replace = false
	} else {
		regProxyGlobal.replace = true
	}

	req.Verbose = &regProxyGlobal.verbose
	req.Service = regProxyGlobal.tenant
	req.Path = regProxyGlobal.scope
	req.AddPath = regProxyGlobal.addScope
	req.URL = regProxyGlobal.url

	b, err := ngsilib.JSONMarshal(req)
	if err != nil {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 2, err.Error()+"\n"))
		return http.StatusBadRequest, []byte(`{"error":"` + err.Error() + `"}`)
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 3, string(b)+"\n"))

	return http.StatusOK, b
}

func regProxyHealthCmd(c *cli.Context) error {
	const funcName = "regProxyHealth"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"regproxy"})
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

func regProxyConfigCmd(c *cli.Context) error {
	const funcName = "regProxyConfig"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"regproxy"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	req := &regProxyReplace{}

	if c.IsSet("verbose") {
		verbose := strings.ToLower(c.String("verbose"))
		switch verbose {
		default:
			return &ngsiCmdError{funcName, 3, "error: set on or off to --verbose option", err}
		case "on", "true":
			v := true
			req.Verbose = &v
		case "off", "false":
			v := false
			req.Verbose = &v
		}
	}
	if c.IsSet("replaceService") {
		service := c.String("replaceService")
		req.Service = &service
	}
	if c.IsSet("replacePath") {
		path := c.String("replacePath")
		req.Path = &path
	}
	if c.IsSet("addPath") {
		addPath := c.String("addPath")
		req.AddPath = &addPath
	}
	if c.IsSet("replaceURL") {
		url := c.String("replaceURL")
		req.URL = &url
	}

	client.SetPath("/config")

	b, err := ngsilib.JSONMarshal(req)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))
	return nil
}
