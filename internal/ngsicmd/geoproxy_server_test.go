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

func TestGeoProxy(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ghost=0.0.0.0", "--port=1028"})

	err := geoProxyServer(c)

	assert.NoError(t, err)
}

func TestGeoProxyOptions(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ghost=0.0.0.0", "--port=1028", "-replaceURL=/v3/entities"})

	err := geoProxyServer(c)

	assert.NoError(t, err)
}

func TestGeoProxyHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ghost=0.0.0.0", "--port=1028", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := geoProxyServer(c)

	assert.NoError(t, err)
}

func TestGeoProxyError(t *testing.T) {
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

func TestGeoProxyErrorServerType(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=geoproxy", "--ghost=0.0.0.0", "--port=1028", "--https"})

	err := geoProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by geoproxy", ngsiErr.Message)
	}
}

func TestGeoProxyErrorKey(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ghost=0.0.0.0", "--port=1028", "--https"})

	err := geoProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestGeoProxyErrorCert(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ghost=0.0.0.0", "--port=1028", "--https", "--key=test.key"})

	err := geoProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestGeoProxyErrorHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--port=1028", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := geoProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestGeoProxyErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ghost=0.0.0.0", "--port=1028"})

	err := geoProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestGeoProxyRootHandler(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,replaceURL,ghost,port,key,cert")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	geoProxyGlobal = &geoProxyParam{ngsi: ngsi, http: mockHTTP, verbose: true, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://geoProxy/", nil)
	got := httptest.NewRecorder()

	geoProxyRootHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestGeoProxyHealthHandler(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mockHTTP,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	req := httptest.NewRequest(http.MethodGet, "http://geoProxy/", nil)
	got := httptest.NewRecorder()

	geoProxyHealthHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestGeoProxyHealthHandlerError(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mockHTTP,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/", nil)
	got := httptest.NewRecorder()

	geoProxyHealthHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestGeoProxyHanderPost(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	req.Header.Set("Authorization", "Bearer 23d7500c85d2f05ffb102e1b7165e325d75f4290")

	got := httptest.NewRecorder()

	geoProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestGeoProxyHanderPostIDM(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	geoProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestGeoProxyHanderErrorMethod(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mockHTTP,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	req := httptest.NewRequest(http.MethodGet, "http://geoProxy/", nil)
	got := httptest.NewRecorder()

	geoProxyHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestGeoProxyHanderPostErrorURL(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	geoProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"geoProxyHandler003 parse \"::\": missing protocol scheme\"}", got.Body.String())
}

func TestGeoProxyHanderPostErrorIDM(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	geoProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"geoProxyHandler004 unknown idm type: unknown\"}", got.Body.String())
}

func TestGeoProxyHanderPostErrorParam(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	geoProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokeProxyRequestToken003 Content-Type error\"}", got.Body.String())
}

func TestGeoProxyHanderPostErrorHTTP(t *testing.T) {
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
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     u,
		client:  client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	req.Header.Set("Authorization", "Bearer 23d7500c85d2f05ffb102e1b7165e325d75f4290")

	got := httptest.NewRecorder()

	geoProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"geoProxyHandler006 http error\"}", got.Body.String())
}

func TestGeoProxyResposeError(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	geoProxyGlobal = &geoProxyParam{
		failure: 1,
		gLock:   &sync.Mutex{},
	}
	got := httptest.NewRecorder()

	geoProxyResposeError(ngsi, got, http.StatusBadRequest, errors.New("test"))

	assert.Equal(t, int64(2), geoProxyGlobal.failure)
	assert.Equal(t, "{\"error\":\"test\"}", got.Body.String())
}

func TestGeoProxySetQueryParam(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	u := &url.URL{}

	err := geoProxySetQueryParam(ngsi, req, u)

	assert.NoError(t, err)
}

func TestGeoProxySetQueryParamErrorContentTypeMissing(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	u := &url.URL{}

	err := geoProxySetQueryParam(ngsi, req, u)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing Content-Type", ngsiErr.Message)
	}
}

func TestGeoProxySetQueryParamErrorParseForm(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = nil
	req.Form = nil
	u := &url.URL{}

	err := geoProxySetQueryParam(ngsi, req, u)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing form body", ngsiErr.Message)
	}
}

func TestGeoProxySetQueryParamErrorContentType(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://geoProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	u := &url.URL{}

	err := geoProxySetQueryParam(ngsi, req, u)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Content-Type error", ngsiErr.Message)
	}
}

func TestGeoProxyGetStat(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	b := geoProxyGetStat()

	stat := &geoProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "OK")
}

func TestGeoProxyGetStatError(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	uHost, _ := url.Parse("http://orion:1026/v2/ex/entities")
	geoProxyGlobal = &geoProxyParam{
		ngsi:    ngsi,
		url:     uHost,
		verbose: true,
		mutex:   &sync.Mutex{},
		gLock:   &sync.Mutex{},
	}

	setJSONEncodeErr(ngsi, 0)

	b := geoProxyGetStat()

	stat := &geoProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "NG")
}

func TestGeoProxyHealthCmd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"geoproxy","version":"0.8.4-next (git_hash:445dfc6166004baf512cad612df05fe137ce5e61)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=geoproxy"})

	err := geoProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"ngsi-go\":\"geoproxy\",\"version\":\"0.8.4-next (git_hash:445dfc6166004baf512cad612df05fe137ce5e61)\",\"health\":\"OK\",\"orion\":\"http://orion:1026/v2/entities\",\"verbose\":true,\"uptime\":\"0 d, 2 h, 49 m, 1 s\",\"timesent\":3,\"success\":3,\"failure\":0}"
		assert.Equal(t, expected, actual)
	}
}

func TestGeoProxyHealthCmdPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"geoproxy","version":"0.8.4-next (git_hash:445dfc6166004baf512cad612df05fe137ce5e61)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=geoproxy", "--pretty"})

	err := geoProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"ngsi-go\": \"geoproxy\",\n  \"version\": \"0.8.4-next (git_hash:445dfc6166004baf512cad612df05fe137ce5e61)\",\n  \"health\": \"OK\",\n  \"orion\": \"http://orion:1026/v2/entities\",\n  \"verbose\": true,\n  \"uptime\": \"0 d, 2 h, 49 m, 1 s\",\n  \"timesent\": 3,\n  \"success\": 3,\n  \"failure\": 0\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestGeoProxyHealthCmdErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := geoProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGeoProxyHealthCmdErrorNewClient(t *testing.T) {
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

	err := geoProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGeoProxyHealthCmdErrorHTTP(t *testing.T) {
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
	_ = set.Parse([]string{"--host=geoproxy"})

	err := geoProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGeoProxyHealthCmdErrorStatusCode(t *testing.T) {
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
	_ = set.Parse([]string{"--host=geoproxy"})

	err := geoProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGeoProxyHealthCmdIotaErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"geoproxy","version":"0.8.4-next (git_hash:445dfc6166004baf512cad612df05fe137ce5e61)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=geoproxy", "--pretty"})

	setJSONIndentError(ngsi)

	err := geoProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
