/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

package convenience

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

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type regProxyConfigParam struct {
	verbose  bool
	bearer   bool
	tenant   *string
	scope    *string
	addScope *string
	url      *string
	replace  bool
}

type regProxyReplace struct {
	Verbose *bool   `json:"verbose,omitempty"`
	Service *string `json:"service,omitempty"`
	Path    *string `json:"path,omitempty"`
	AddPath *string `json:"add_path,omitempty"`
	URL     *string `json:"url,omitempty"`
}

type regProxyStat struct {
	mutex *sync.Mutex

	startTime time.Time
	timeSent  int64
	success   int64
	failure   int64
}

type regProxyStatInfo struct {
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

func regProxyServer(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "regProxy"

	host := c.String("rhost")
	port := c.String("port")
	addr := host + ":" + port

	path := c.String("url")
	url := addr + path

	if c.Bool("https") {
		if !c.IsSet("key") {
			return ngsierr.New(funcName, 1, "no key file provided", nil)
		}
		if !c.IsSet("cert") {
			return ngsierr.New(funcName, 2, "no cert file provided", nil)
		}
		url = "https://" + url
	} else {
		url = "http://" + url
	}

	config := &regProxyConfigParam{
		verbose: c.Bool("verbose"),
	}
	if c.IsSet("replaceService") {
		tenant := c.String("replaceService")
		config.tenant = &tenant
		config.replace = true
	}
	if c.IsSet("replacePath") {
		scope := c.String("replacePath")
		config.scope = &scope
		config.replace = true
	}
	if c.IsSet("addPath") {
		addScope := c.String("addPath")
		config.addScope = &addScope
		config.replace = true
	}
	if c.IsSet("replaceURL") {
		repalceURL := c.String("replaceURL")
		config.url = &repalceURL
		config.replace = true
	}

	stat := &regProxyStat{
		mutex:     &sync.Mutex{},
		startTime: time.Now(),
	}
	regProxyHandler := &regProxyHandler{
		ngsi:   ngsi,
		client: client,
		http:   ngsi.HTTP,
		mutex:  &sync.Mutex{},
		config: config,
		stat:   stat,
	}

	mux := http.NewServeMux()
	mux.Handle("/", &regProxyRootHandler{ngsi: ngsi})
	mux.Handle("/v1/", regProxyHandler)
	mux.Handle("/v2/", regProxyHandler)
	mux.Handle("/health", &regProxyHealthHandler{ngsi: ngsi, host: client.Server.ServerHost, stat: stat, config: config})
	mux.Handle("/config", &regProxyConfigHandler{ngsi: ngsi, config: config})

	ngsi.Logging(ngsilib.LogInfo, "Start registration proxy: "+url+"\n")

	if c.Bool("https") {
		err := ngsi.NetLib.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), mux)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	} else {
		err := ngsi.NetLib.ListenAndServe(addr, mux)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
	}

	return nil
}

type regProxyRootHandler struct {
	ngsi *ngsilib.NGSI
}

func (h *regProxyRootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyRootHandler"

	ngsi := h.ngsi
	ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, r.URL.Path)+"\n")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s not found"}`, r.URL.Path)))
}

type regProxyHealthHandler struct {
	ngsi   *ngsilib.NGSI
	host   string
	stat   *regProxyStat
	config *regProxyConfigParam
}

func (h *regProxyHealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyHealthHandler"

	ngsi := h.ngsi

	if r.Method != http.MethodGet {
		ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(regProxyGetStat(h))
	}
}

type regProxyConfigHandler struct {
	ngsi   *ngsilib.NGSI
	config *regProxyConfigParam
}

func (h *regProxyConfigHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyConfigHandler"

	ngsi := h.ngsi

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.Header().Set("Content-Type", "Application/json")
		b := getRequestBody(r.Body)
		status, body := regProxyConfig(h, b)
		w.WriteHeader(status)
		_, _ = w.Write(body)
	}
}

type regProxyHandler struct {
	ngsi   *ngsilib.NGSI
	client *ngsilib.Client
	http   ngsilib.HTTPRequest
	mutex  *sync.Mutex
	config *regProxyConfigParam
	stat   *regProxyStat
}

func (h *regProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const funcName = "regProxyHandler"

	status := http.StatusBadRequest
	ngsi := h.ngsi
	verbose := h.config.verbose
	client := h.client
	host := client.Server.ServerHost

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, r.URL.Path)+"\n")

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
			if h.config.tenant != nil {
				tenant = *h.config.tenant
			}
			headers[k] = tenant
			continue
		case "fiware-servicepath":
			scope = r.Header.Get(k)
			origScope = scope
			if h.config.scope != nil {
				scope = *h.config.scope
			}
			headers[k] = scope
			continue
		}
		headers[k] = r.Header.Get(k)
	}

	if h.config.addScope != nil {
		scope = path.Join(*h.config.addScope, scope)
	}

	uPath := r.URL.Path
	if h.config.url != nil {
		uPath = *h.config.url
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("Path:%s, Tenant: %s, Scope: %s\n", r.URL.Path, origTenant, origScope))
	if h.config.replace {
		ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("Path:%s, Tenant: %s, Scope: %s\n", uPath, tenant, scope))
	}

	u, err := url.Parse(host)
	if err != nil {
		ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 4, err.Error()))
		regProxyFailureUp(h)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u.Path = path.Join(u.Path, uPath)

	h.mutex.Lock()
	key, token, err := ngsi.GetAuthHeader(client)
	h.mutex.Unlock()

	if err != nil {
		ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 5, err.Error()))
		regProxyFailureUp(h)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	headers[key] = token

	if verbose {
		ngsi.Logging(ngsilib.LogInfo, string(b)+"\n")
	}

	res, resBody, err := h.http.Request("POST", u, headers, b)
	if err == nil {
		h.stat.mutex.Lock()
		h.stat.timeSent += 1
		h.stat.success += 1
		h.stat.mutex.Unlock()

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
		regProxyFailureUp(h)
	}

	ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 6, err.Error()))

	w.WriteHeader(status)
}

func getRequestBody(body io.ReadCloser) []byte {
	defer func() { _ = body.Close() }()
	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, body)
	return buf.Bytes()
}

func regProxyFailureUp(h *regProxyHandler) {
	h.stat.mutex.Lock()
	h.stat.timeSent += 1
	h.stat.failure += 1
	h.stat.mutex.Unlock()
}

func regProxyGetStat(h *regProxyHealthHandler) []byte {
	uptime := time.Now().Unix() - h.stat.startTime.Unix()

	h.stat.mutex.Lock()
	stat := regProxyStatInfo{
		NgsiGo:   "regproxy",
		Version:  ngsicli.Version,
		Health:   "OK",
		Csource:  h.host,
		Verbose:  h.config.verbose,
		Uptime:   ngsilib.HumanizeUptime(uptime),
		Timesent: h.stat.timeSent,
		Success:  h.stat.success,
		Failure:  h.stat.failure,
	}
	h.stat.mutex.Unlock()

	if h.config.replace {
		stat.Replace = &regProxyReplace{
			Service: h.config.tenant,
			Path:    h.config.scope,
			AddPath: h.config.addScope,
			URL:     h.config.url,
		}
	}

	b, err := ngsilib.JSONMarshal(stat)
	if err != nil {
		return []byte(`{"ngsi-go":"regproxy","health":"NG"}`)
	}

	return b
}

func regProxyConfig(h *regProxyConfigHandler, body []byte) (int, []byte) {
	const funcName = "regProxyConfig"

	req := &regProxyReplace{}

	err := ngsilib.JSONUnmarshal(body, &req)
	if err != nil {
		h.ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, err.Error()+"\n"))
		return http.StatusBadRequest, []byte(`{"error":"` + err.Error() + `"}`)
	}

	if req.Verbose != nil {
		h.config.verbose = *req.Verbose
	}
	if req.Service != nil {
		if *req.Service == "" {
			h.config.tenant = nil
		} else {
			h.config.tenant = req.Service
		}
	}
	if req.Path != nil {
		if *req.Path == "" {
			h.config.scope = nil
		} else {
			h.config.scope = req.Path
		}
	}
	if req.AddPath != nil {
		if *req.AddPath == "" {
			h.config.addScope = nil
		} else {
			h.config.addScope = req.AddPath
		}
	}
	if req.URL != nil {
		if *req.URL == "" {
			h.config.url = nil
		} else {
			h.config.url = req.URL
		}
	}

	if h.config.tenant == nil &&
		h.config.scope == nil &&
		h.config.addScope == nil &&
		h.config.url == nil {
		h.config.replace = false
	} else {
		h.config.replace = true
	}

	req.Verbose = &h.config.verbose
	req.Service = h.config.tenant
	req.Path = h.config.scope
	req.AddPath = h.config.addScope
	req.URL = h.config.url

	b, err := ngsilib.JSONMarshal(req)
	if err != nil {
		h.ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 2, err.Error()+"\n"))
		return http.StatusBadRequest, []byte(`{"error":"` + err.Error() + `"}`)
	}

	h.ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 3, string(b)+"\n"))

	return http.StatusOK, b
}

func regProxyHealthCmd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "regProxyHealth"

	client.SetPath("/health")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))
	return nil
}

func regProxyConfigCmd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "regProxyConfig"

	req := &regProxyReplace{}

	if c.IsSet("verbose") {
		verbose := strings.ToLower(c.String("verbose"))
		switch verbose {
		default:
			return ngsierr.New(funcName, 1, "error: set on or off to --verbose option", nil)
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
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))
	return nil
}
