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

package convenience

import (
	"bytes"
	"fmt"
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

type queryProxyStatInfo struct {
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

type queryProxyStat struct {
	mutex *sync.Mutex

	startTime time.Time
	timeSent  int64
	success   int64
	failure   int64
}

func queryProxyServer(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "queryProxy"

	host := c.String("ghost")
	port := c.String("port")
	addr := host + ":" + port
	proxyUrl := addr

	if c.Bool("https") {
		if !c.IsSet("key") {
			return ngsierr.New(funcName, 1, "no key file provided", nil)
		}
		if !c.IsSet("cert") {
			return ngsierr.New(funcName, 2, "no cert file provided", nil)
		}
		proxyUrl = "https://" + proxyUrl
	} else {
		proxyUrl = "http://" + proxyUrl
	}

	u, _ := url.Parse(client.URL.String())
	u.Path = path.Join(u.Path, "/v2/entities")

	proxyPath := "/v2/ex/entities"
	if c.IsSet("replaceURL") {
		proxyPath = c.String("replaceURL")
	}

	stat := &queryProxyStat{mutex: &sync.Mutex{}, startTime: time.Now()}

	mux := http.NewServeMux()
	mux.Handle("/", &queryProxyRootHandler{ngsi: ngsi})
	mux.Handle("/health", &queryProxyHealthHandler{ngsi: ngsi, stat: stat, broker: u.String(), verbose: c.Bool("verbose")})
	mux.Handle(proxyPath, &queryProxyHandler{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    ngsi.HTTP,
		verbose: c.Bool("verbose"),
		mutex:   &sync.Mutex{},
		stat:    stat,
	})

	ngsi.Logging(ngsilib.LogErr, "Start geo proxy: "+proxyUrl+"\n")

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

type queryProxyRootHandler struct {
	ngsi *ngsilib.NGSI
}

func (h *queryProxyRootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const funcName = "queryProxyRootHandler"

	h.ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, r.URL.Path)+"\n")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s not found"}`, r.URL.Path)))
}

type queryProxyHealthHandler struct {
	ngsi *ngsilib.NGSI
	stat *queryProxyStat

	broker  string
	verbose bool
}

func (h *queryProxyHealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const funcName = "queryProxyHealthHandler"

	if r.Method != http.MethodGet {
		h.ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(queryProxyGetStat(h))
	}
}

type queryProxyHandler struct {
	ngsi    *ngsilib.NGSI
	url     *url.URL
	client  *ngsilib.Client
	http    ngsilib.HTTPRequest
	verbose bool
	mutex   *sync.Mutex
	stat    *queryProxyStat
}

func (h *queryProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const funcName = "queryProxyHandler"

	status := http.StatusBadRequest
	ngsi := h.ngsi
	client := h.client
	verbose := h.verbose

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ngsi.Logging(ngsilib.LogErr, ngsierr.SprintMsg(funcName, 2, r.URL.Path)+"\n")

	u, err := url.Parse(h.url.String())
	if err != nil {
		queryProxyResposeError(h, w, status, ngsierr.New(funcName, 3, err.Error(), err))
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
		h.mutex.Lock()
		key, token, err := ngsi.GetAuthHeader(client)
		h.mutex.Unlock()
		if err != nil {
			queryProxyResposeError(h, w, status, ngsierr.New(funcName, 4, err.Error(), err))
			return
		}
		headers[key] = token
	}

	err = queryProxySetQueryParam(ngsi, r, u)
	if err != nil {
		queryProxyResposeError(h, w, status, err)
		return
	}

	res, resBody, err := h.http.Request("GET", u, headers, nil)
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
			ngsi.Logging(ngsilib.LogInfo, ngsierr.SprintMsg(funcName, 5, fmt.Sprintf("%d %s\n", res.StatusCode, string(resBody))))
		}
		return
	} else {
		queryProxyResposeError(h, w, status, ngsierr.New(funcName, 6, err.Error(), err))
	}
}

func queryProxyResposeError(h *queryProxyHandler, w http.ResponseWriter, status int, err error) {
	h.stat.mutex.Lock()
	h.stat.timeSent += 1
	h.stat.failure += 1
	h.stat.mutex.Unlock()

	msg := ngsierr.Message(err)
	h.ngsi.Logging(ngsilib.LogErr, msg+"\n")

	body := []byte(fmt.Sprintf(`{"error":"%s"}`, msg))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func queryProxySetQueryParam(ngsi *ngsilib.NGSI, r *http.Request, u *url.URL) error {
	const funcName = "tokeProxyRequestToken"

	ctype := r.Header["Content-Type"]
	if ctype == nil || (ctype != nil && len(ctype) == 0) {
		return ngsierr.New(funcName, 1, "missing Content-Type", nil)
	}

	q := u.Query()

	if ctype[0] == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		for k, v := range r.Form {
			q.Set(k, v[0])
		}
	} else {
		return ngsierr.New(funcName, 3, "Content-Type error", nil)
	}

	u.RawQuery = q.Encode()

	return nil
}

func queryProxyGetStat(h *queryProxyHealthHandler) []byte {
	uptime := time.Now().Unix() - h.stat.startTime.Unix()

	h.stat.mutex.Lock()
	stat := queryProxyStatInfo{
		NgsiGo:   "queryproxy",
		Version:  ngsicli.Version,
		Health:   "OK",
		Orion:    h.broker,
		Verbose:  h.verbose,
		Uptime:   ngsilib.HumanizeUptime(uptime),
		Timesent: h.stat.timeSent,
		Success:  h.stat.success,
		Failure:  h.stat.failure,
	}
	h.stat.mutex.Unlock()

	b, err := ngsilib.JSONMarshal(stat)
	if err != nil {
		return []byte(`{"ngsi-go":"queryproxy","health":"NG"}`)
	}

	return b
}

func queryProxyHealthCmd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "queryProxyHealth"

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
