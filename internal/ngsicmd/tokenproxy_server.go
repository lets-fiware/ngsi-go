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
	"encoding/base64"
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

type tokenProxyParam struct {
	ngsi          *ngsilib.NGSI
	http          ngsilib.HTTPRequest
	verbose       bool
	idmHost       *url.URL
	RevokeURL     *url.URL
	clientID      string
	clientSecret  string
	Authorization string
	gLock         *sync.Mutex

	startTime time.Time
	timeSent  int64
	success   int64
	revoke    int64
	failure   int64
}

type tokenProxyStat struct {
	NgsiGo       string `json:"ngsi-go"`
	Version      string `json:"version"`
	Health       string `json:"health"`
	Idm          string `json:"idm"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Verbose      bool   `json:"verbose"`
	Uptime       string `json:"uptime"`
	Timesent     int64  `json:"timesent"`
	Success      int64  `json:"success"`
	Revoke       int64  `json:"revoke"`
	Failure      int64  `json:"failure"`
}

type tokenProxyRequest struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Token    *string `json:"token,omitempty"`
	Scope    *string `json:"scope,omitempty"`
}

type tokenProxyRevoke struct {
	Token         *string `json:"token,omitempty"`
	TokenTypeHint *string `json:"token_type_hint,omitempty"`
}

var tokenProxyGlobal *tokenProxyParam

func tokenProxyServer(c *cli.Context) error {
	const funcName = "tokenProxy"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")
	port := c.String("port")
	addr := host + ":" + port
	proxyUrl := addr

	if c.Bool("https") {
		if !c.IsSet("key") {
			return &ngsiCmdError{funcName, 2, "no key file provided", nil}
		}
		if !c.IsSet("cert") {
			return &ngsiCmdError{funcName, 3, "no cert file provided", nil}
		}
		proxyUrl = "https://" + proxyUrl
	} else {
		proxyUrl = "http://" + proxyUrl
	}

	u, err := url.Parse(c.String("idmHost"))
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if !strings.HasSuffix(u.Path, "/oauth2/token") {
		u.Path = path.Join(u.Path, "/oauth2/token")
	}
	revokeURL, _ := url.Parse(c.String("idmHost"))
	revokeURL.Path = u.Path[:len(u.Path)-6] + "/revoke"

	clientID := c.String("clientId")
	clientSecret := c.String("clientSecret")

	tokenProxyGlobal = &tokenProxyParam{
		ngsi:          ngsi,
		http:          ngsi.HTTP,
		idmHost:       u,
		RevokeURL:     revokeURL,
		clientID:      clientID,
		clientSecret:  clientSecret,
		Authorization: "Basic " + base64.URLEncoding.EncodeToString([]byte(clientID+":"+clientSecret)),
		verbose:       c.Bool("verbose"),
		gLock:         &sync.Mutex{},
		startTime:     time.Now(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(tokenProxyRootHandler))
	mux.HandleFunc("/token", http.HandlerFunc(tokenProxyHandler))
	mux.HandleFunc("/revoke", http.HandlerFunc(tokenProxyHandler))
	mux.HandleFunc("/health", http.HandlerFunc(tokenProxyHealthHandler))

	ngsi.Logging(ngsilib.LogErr, "Start token proxy: "+proxyUrl+"\n")

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

func tokenProxyRootHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "tokenProxyRootHandler"

	ngsi := tokenProxyGlobal.ngsi
	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, r.URL.Path)+"\n")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s not found"}`, r.URL.Path)))
}

func tokenProxyHealthHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "tokenProxyHealthHandler"

	ngsi := tokenProxyGlobal.ngsi

	if r.Method != http.MethodGet {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(tokenProxyGetStat())
	}
}

func tokenProxyHandler(w http.ResponseWriter, r *http.Request) {
	const funcName = "tokenProxyHandler"

	status := http.StatusBadRequest
	ngsi := tokenProxyGlobal.ngsi
	verbose := tokenProxyGlobal.verbose

	if r.Method != http.MethodPost {
		ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 1, "Method not allowed"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 2, r.URL.Path)+"\n")

	var body []byte
	var err error
	var u *url.URL

	revoke := r.URL.Path == "/revoke"
	if revoke {
		body, err = tokenProxyRevokeToken(ngsi, r)
		if err != nil {
			tokenProxyResposeError(ngsi, w, status, err)
			return
		}
		u = tokenProxyGlobal.RevokeURL
	} else {
		body, err = tokenProxyRequestToken(ngsi, r)
		if err != nil {
			tokenProxyResposeError(ngsi, w, status, err)
			return
		}
		u = tokenProxyGlobal.idmHost
	}

	headers := map[string]string{}
	headers["Authorization"] = tokenProxyGlobal.Authorization
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	res, resBody, err := tokenProxyGlobal.http.Request("POST", u, headers, body)
	if err == nil {
		tokenProxyGlobal.gLock.Lock()
		tokenProxyGlobal.timeSent += 1
		if revoke {
			tokenProxyGlobal.revoke += 1
		} else {
			tokenProxyGlobal.success += 1
		}
		tokenProxyGlobal.gLock.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.StatusCode)
		_, _ = w.Write(resBody)
		if verbose {
			ngsi.Logging(ngsilib.LogInfo, sprintMsg(funcName, 3, fmt.Sprintf("%d %s\n", res.StatusCode, string(resBody))))
		}
		return
	} else {
		tokenProxyResposeError(ngsi, w, status, &ngsiCmdError{funcName, 4, err.Error(), err})
	}
}

func tokenProxyResposeError(ngsi *ngsilib.NGSI, w http.ResponseWriter, status int, err error) {
	tokenProxyGlobal.gLock.Lock()
	tokenProxyGlobal.timeSent += 1
	tokenProxyGlobal.failure += 1
	tokenProxyGlobal.gLock.Unlock()

	msg := message(err)
	ngsi.Logging(ngsilib.LogErr, msg+"\n")

	body := []byte(fmt.Sprintf(`{"error":"%s"}`, msg))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func tokenProxyRequestToken(ngsi *ngsilib.NGSI, r *http.Request) ([]byte, error) {
	const funcName = "tokeProxyRequestToken"

	ctype := r.Header["Content-Type"]
	if ctype == nil || (ctype != nil && len(ctype) == 0) {
		return nil, &ngsiCmdError{funcName, 1, "missing Content-Type", nil}
	}
	req := &tokenProxyRequest{}
	var body, msg string
	verbose := tokenProxyGlobal.verbose

	switch ctype[0] {
	default:
		return nil, &ngsiCmdError{funcName, 2, "Content-Type error", nil}
	case "application/json":
		b := getRequestBody(r.Body)
		err := ngsilib.JSONUnmarshal(b, req)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 3, err.Error(), err}
		}
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			return nil, &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		for k, v := range r.Form {
			switch k {
			default:
				return nil, &ngsiCmdError{funcName, 5, "unknown parameter: " + k, nil}
			case "username":
				req.Username = &v[0]
			case "password":
				req.Password = &v[0]
			case "scope":
				req.Scope = &v[0]
			}
		}
	}

	if req.Username != nil && req.Password != nil && req.Token == nil {
		body = fmt.Sprintf("grant_type=password&username=%s&password=%s", *req.Username, *req.Password)
		if verbose {
			msg = body
		} else {
			msg = fmt.Sprintf("grant_type=password&username=%s&password=*****", *req.Username)
		}
	} else if req.Username == nil && req.Password == nil && req.Token != nil {
		body = fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", *req.Token)
		msg = body
	} else {
		return nil, &ngsiCmdError{funcName, 6, "parameter error", nil}
	}
	if req.Scope != nil {
		body += "&scope=" + *req.Scope
		msg += "&scope=" + *req.Scope
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 7, msg+"\n"))

	return []byte(body), nil
}

func tokenProxyRevokeToken(ngsi *ngsilib.NGSI, r *http.Request) ([]byte, error) {
	const funcName = "tokenProxyRevokeToken"

	ctype := r.Header["Content-Type"]
	if ctype == nil || (ctype != nil && len(ctype) == 0) {
		return nil, &ngsiCmdError{funcName, 1, "missing Content-Type", nil}
	}
	req := &tokenProxyRevoke{}
	var msg string
	var body []byte

	switch ctype[0] {
	default:
		return nil, &ngsiCmdError{funcName, 2, "Content-Type error", nil}
	case "application/json":
		b := getRequestBody(r.Body)
		err := ngsilib.JSONUnmarshal(b, req)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 3, err.Error(), err}
		}
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			return nil, &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		for k, v := range r.Form {
			switch k {
			default:
				return nil, &ngsiCmdError{funcName, 5, "unknown parameter: " + k, nil}
			case "token":
				req.Token = &v[0]
			case "token_type_hint":
				req.TokenTypeHint = &v[0]
			}
		}
	}

	if req.Token != nil {
		if req.TokenTypeHint == nil {
			hint := "refresh_token"
			req.TokenTypeHint = &hint
		}
		msg = fmt.Sprintf("token=%s&token_type_hint=%s", *req.Token, *req.TokenTypeHint)
		body = []byte(msg)
	} else {
		return nil, &ngsiCmdError{funcName, 6, "parameter error", nil}
	}

	ngsi.Logging(ngsilib.LogErr, sprintMsg(funcName, 7, msg+"\n"))

	return body, nil
}

func tokenProxyGetStat() []byte {
	uptime := time.Now().Unix() - tokenProxyGlobal.startTime.Unix()

	tokenProxyGlobal.gLock.Lock()
	stat := tokenProxyStat{
		NgsiGo:       "tokenproxy",
		Version:      Version,
		Health:       "OK",
		Idm:          tokenProxyGlobal.idmHost.String(),
		ClientID:     tokenProxyGlobal.clientID,
		ClientSecret: tokenProxyGlobal.clientSecret,
		Verbose:      tokenProxyGlobal.verbose,
		Uptime:       humanizeUptime(uptime),
		Timesent:     tokenProxyGlobal.timeSent,
		Success:      tokenProxyGlobal.success,
		Revoke:       tokenProxyGlobal.revoke,
		Failure:      tokenProxyGlobal.failure,
	}
	tokenProxyGlobal.gLock.Unlock()

	b, err := ngsilib.JSONMarshal(stat)
	if err != nil {
		return []byte(`{"ngsi-go":"tokenproxy","health":"NG"}`)
	}

	return b
}

func tokenProxyHealthCmd(c *cli.Context) error {
	const funcName = "tokenProxyHealth"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"tokenproxy"})
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
