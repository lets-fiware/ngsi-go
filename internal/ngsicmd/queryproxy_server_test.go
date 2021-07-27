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
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestQueryProxy(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--qhost=0.0.0.0", "--port=1028"})

	err := queryProxyServer(c)

	assert.NoError(t, err)
}

func TestQueryProxyOptions(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--qhost=0.0.0.0", "--port=1028", "-replaceURL=/v3/entities"})

	err := queryProxyServer(c)

	assert.NoError(t, err)
}

func TestQueryProxyHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--qhost=0.0.0.0", "--port=1028", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := queryProxyServer(c)

	assert.NoError(t, err)
}

func TestQueryProxyError(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := receiver(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestQueryProxyErrorServerType(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=queryproxy", "--qhost=0.0.0.0", "--port=1028", "--https"})

	err := queryProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by queryproxy", ngsiErr.Message)
	}
}

func TestQueryProxyErrorKey(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--qhost=0.0.0.0", "--port=1028", "--https"})

	err := queryProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestQueryProxyErrorCert(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--qhost=0.0.0.0", "--port=1028", "--https", "--key=test.key"})

	err := queryProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestQueryProxyErrorHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--port=1028", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := queryProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestQueryProxyErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--qhost=0.0.0.0", "--port=1028"})

	err := queryProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestQueryProxyRootHandler(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,replaceURL,qhost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	queryProxyGlobal = &queryProxyParam{ngsi: ngsi, http: mockHTTP, verbose: true, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://queryProxy/", nil)
	got := httptest.NewRecorder()

	queryProxyRootHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHealthHandler(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	u, _ := url.Parse("http://orion:1026")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mockHTTP,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	req := httptest.NewRequest(http.MethodGet, "http://queryProxy/", nil)
	got := httptest.NewRecorder()

	queryProxyHealthHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHealthHandlerError(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	u, _ := url.Parse("http://orion:1026")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mockHTTP,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/", nil)
	got := httptest.NewRecorder()

	queryProxyHealthHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderPost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	client.Server.IdmType = ""

	resHeader := http.Header{}
	resHeader["Content-Type"] = []string{"application/json"}
	query := "options=keyValues&type=Deivce"
	reqRes := MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = &query
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = resHeader
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	req.Header.Set("Authorization", "Bearer 23d7500c85d2f05ffb102e1b7165e325d75f4290")

	got := httptest.NewRecorder()

	queryProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderPostIDM(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	client.Server.IdmType = "basic"
	client.Server.Username = "fiware"
	client.Server.Password = "1234"

	resHeader := http.Header{}
	resHeader["Content-Type"] = []string{"application/json"}
	query := "options=keyValues&type=Deivce"
	reqRes := MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = &query
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = resHeader
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	queryProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderErrorMethod(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	u, _ := url.Parse("http://orion:1026")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mockHTTP,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	req := httptest.NewRequest(http.MethodGet, "http://queryProxy/", nil)
	got := httptest.NewRecorder()

	queryProxyHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderPostErrorURL(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	resHeader := http.Header{}
	resHeader["Content-Type"] = []string{"application/json"}
	query := "options=keyValues&type=Deivce"
	reqRes := MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = &query
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = resHeader
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u := &url.URL{Scheme: ":"}
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	queryProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"queryProxyHandler003 parse \"::\": missing protocol scheme\"}", got.Body.String())
}

func TestQueryProxyHanderPostErrorIDM(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	client.Server.IdmType = "unknown"

	resHeader := http.Header{}
	resHeader["Content-Type"] = []string{"application/json"}
	query := "options=keyValues&type=Deivce"
	reqRes := MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = &query
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = resHeader
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	queryProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"queryProxyHandler004 unknown idm type: unknown\"}", got.Body.String())
}

func TestQueryProxyHanderPostErrorParam(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	client.Server.IdmType = ""

	resHeader := http.Header{}
	resHeader["Content-Type"] = []string{"application/json"}
	query := "options=keyValues&type=Deivce"
	reqRes := MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = &query
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = resHeader
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	queryProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokeProxyRequestToken003 Content-Type error\"}", got.Body.String())
}

func TestQueryProxyHanderPostErrorHTTP(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	client.Server.IdmType = ""

	resHeader := http.Header{}
	resHeader["Content-Type"] = []string{"application/json"}
	query := "options=keyValues&type=Deivce"
	reqRes := MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Err = errors.New("http error")
	reqRes.RawQuery = &query
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = resHeader
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	req.Header.Set("Authorization", "Bearer 23d7500c85d2f05ffb102e1b7165e325d75f4290")

	got := httptest.NewRecorder()

	queryProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"queryProxyHandler006 http error\"}", got.Body.String())
}

func TestQueryProxyResposeError(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	queryProxyGlobal = &queryProxyParam{
		failure: 1,
		gLock:   &sync.Mutex{},
	}
	got := httptest.NewRecorder()

	queryProxyResposeError(ngsi, got, http.StatusBadRequest, errors.New("test"))

	assert.Equal(t, int64(2), queryProxyGlobal.failure)
	assert.Equal(t, "{\"error\":\"test\"}", got.Body.String())
}

func TestQueryProxySetQueryParam(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	u := &url.URL{}

	err := queryProxySetQueryParam(ngsi, req, u)

	assert.NoError(t, err)
}

func TestQueryProxySetQueryParamErrorContentTypeMissing(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	u := &url.URL{}

	err := queryProxySetQueryParam(ngsi, req, u)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing Content-Type", ngsiErr.Message)
	}
}

func TestQueryProxySetQueryParamErrorParseForm(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = nil
	req.Form = nil
	u := &url.URL{}

	err := queryProxySetQueryParam(ngsi, req, u)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing form body", ngsiErr.Message)
	}
}

func TestQueryProxySetQueryParamErrorContentType(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	u := &url.URL{}

	err := queryProxySetQueryParam(ngsi, req, u)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Content-Type error", ngsiErr.Message)
	}
}

func TestQueryProxyGetStat(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	b := queryProxyGetStat()

	stat := &queryProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "OK")
}

func TestQueryProxyGetStatError(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	queryProxyGlobal = &queryProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	setJSONEncodeErr(ngsi, 0)

	b := queryProxyGetStat()

	stat := &queryProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "NG")
}

func TestQueryProxyHealthCmd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"queryproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=queryproxy"})

	err := queryProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"ngsi-go\":\"queryproxy\",\"version\":\"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)\",\"health\":\"OK\",\"orion\":\"http://orion:1026/v2/entities\",\"verbose\":true,\"uptime\":\"0 d, 2 h, 49 m, 1 s\",\"timesent\":3,\"success\":3,\"failure\":0}"
		assert.Equal(t, expected, actual)
	}
}

func TestQueryProxyHealthCmdPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"queryproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=queryproxy", "--pretty"})

	err := queryProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"ngsi-go\": \"queryproxy\",\n  \"version\": \"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)\",\n  \"health\": \"OK\",\n  \"orion\": \"http://orion:1026/v2/entities\",\n  \"verbose\": true,\n  \"uptime\": \"0 d, 2 h, 49 m, 1 s\",\n  \"timesent\": 3,\n  \"success\": 3,\n  \"failure\": 0\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQueryProxyHealthCmdErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := queryProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestQueryProxyHealthCmdErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})

	err := queryProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestQueryProxyHealthCmdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=queryproxy"})

	err := queryProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestQueryProxyHealthCmdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=queryproxy"})

	err := queryProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestQueryProxyHealthCmdIotaErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"queryproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=queryproxy", "--pretty"})

	setJSONIndentError(ngsi)

	err := queryProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
