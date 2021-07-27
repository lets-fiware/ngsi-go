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
	"sync"
	"testing"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestRegProxy(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,rhost,port,url")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--rhost=0.0.0.0", "--port=1028", "--url=/"})

	err := regProxyServer(c)

	assert.NoError(t, err)
}

func TestRegProxyOptions(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,rhost,port,url,replaceService,replacePath,replaceURL,addPath")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--rhost=0.0.0.0", "--port=1028", "--url=/",
		"--replaceService=fiware", "--replacePath=/orion", "--replaceURL=/v3/queryConxt", "--addPath=/federation"})

	err := regProxyServer(c)

	assert.NoError(t, err)
}

func TestRegProxyHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,rhost,port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--rhost=0.0.0.0", "--port=1028", "--url=/", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := regProxyServer(c)

	assert.NoError(t, err)
}

func TestRegProxyErrorHost(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := regProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	}
}

func TestRegProxyErrorHostType(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "host,rhost,port,url,key,cert")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := regProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	}
}

func TestRegProxyErrorKey(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,rhost,port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--rhost=0.0.0.0", "--port=1028", "--url=/", "--https"})

	err := regProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestRegProxyErrorCert(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,rhost,port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--rhost=0.0.0.0", "--port=1028", "--url=/", "--https", "--key=test.key"})

	err := regProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestRegProxyErrorHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	setupFlagString(set, "host,rhost,port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--rhost=0.0.0.0", "--port=1028", "--url=/", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := regProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestRegProxyErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	setupFlagString(set, "host,rhost,port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--rhost=0.0.0.0", "--port=1028", "--url=/"})

	err := regProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestRegProxyRootHandler(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	regProxyRootHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHealthHandler(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	regProxyHealthHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHealthHandlerError(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodPost, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	regProxyHealthHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyConfigHandler(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	b := bytes.NewReader([]byte(`{"verbose":false}`))
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/", b)
	got := httptest.NewRecorder()

	regProxyConfigHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyConfigHandlerError(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	regProxyConfigHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorGet(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorMethod(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodDelete, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderPost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	client.Server.IdmType = "tokenproxy"
	client.Server.IdmHost = "/token"
	client.Server.Username = "testuser"
	client.Server.Password = "1234"

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	h := http.Header{}
	h["Context-Length"] = []string{"0"}
	reqRes2.ResHeader = h
	mockHTTP := NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	req.Header.Set("Authorization", "Basic 1234")
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderPostOptions(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	client.Server.IdmType = "tokenproxy"
	client.Server.IdmHost = "/token"
	client.Server.Username = "testuser"
	client.Server.Password = "1234"

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	h := http.Header{}
	h["Context-Length"] = []string{"0"}
	reqRes2.ResHeader = h
	mockHTTP := NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)
	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"
	regProxyGlobal = &regProxyParam{
		ngsi:     ngsi,
		client:   client,
		http:     mockHTTP,
		verbose:  true,
		bearer:   true,
		tenant:   &tenant,
		scope:    &scope,
		addScope: &addScope,
		replace:  true,
		url:      &url,
		mutex:    &sync.Mutex{},
		gLock:    &sync.Mutex{},
	}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderPostXAuth(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	client.Server.IdmType = "tokenproxy"
	client.Server.IdmHost = "/token"
	client.Server.Username = "testuser"
	client.Server.Password = "1234"

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	mockHTTP := NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: false, bearer: false, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	client.Server.ServerHost = ":"
	client.Server.IdmType = "tokenproxy"
	client.Server.IdmHost = "/token"
	client.Server.Username = "testuser"
	client.Server.Password = "1234"

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	mockHTTP := NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorToken(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	client.Server.IdmType = "tokenproxy"
	client.Server.IdmHost = "/token"

	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	mockHTTP := NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderBadRequest(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	client.Server.IdmType = "tokenproxy"
	client.Server.IdmHost = "/token"
	client.Server.Username = "testuser"
	client.Server.Password = "1234"

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/"
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	mockHTTP := NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestGetRequestBody(t *testing.T) {
	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)

	b := getRequestBody(req.Body)

	assert.Equal(t, "{\"entities\":[{\"id\":\"urn:ngsi-ld:Device:uDr8vgsJ0Xbe\",\"type\":\"Device\"}],\"attrs\":[\"temperature\"]}", string(b))
}

func TestRegProxyFailureUp(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	regProxyGlobal = &regProxyParam{ngsi: ngsi, verbose: true, bearer: true, mutex: &sync.Mutex{}, gLock: &sync.Mutex{}}

	regProxyFailureUp()

	assert.Equal(t, int64(1), regProxyGlobal.failure)
}

func TestRegProxyGetStat(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"
	regProxyGlobal = &regProxyParam{
		ngsi:      ngsi,
		verbose:   true,
		bearer:    true,
		tenant:    &tenant,
		scope:     &scope,
		addScope:  &addScope,
		replace:   true,
		url:       &url,
		startTime: time.Now(),
		mutex:     &sync.Mutex{},
		gLock:     &sync.Mutex{},
	}

	b := regProxyGetStat("http://orion")

	stat := &regProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "OK")
}

func TestRegProxyGetStatError(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"
	regProxyGlobal = &regProxyParam{
		ngsi:      ngsi,
		verbose:   true,
		bearer:    true,
		tenant:    &tenant,
		scope:     &scope,
		addScope:  &addScope,
		replace:   true,
		url:       &url,
		startTime: time.Now(),
		mutex:     &sync.Mutex{},
		gLock:     &sync.Mutex{},
	}

	setJSONEncodeErr(ngsi, 0)

	b := regProxyGetStat("http://orion")

	stat := &regProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "NG")
}

func TestRegProxyConfig(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"
	regProxyGlobal = &regProxyParam{
		ngsi:      ngsi,
		verbose:   true,
		bearer:    true,
		tenant:    &tenant,
		scope:     &scope,
		addScope:  &addScope,
		replace:   true,
		url:       &url,
		startTime: time.Now(),
		mutex:     &sync.Mutex{},
	}

	req := &regProxyReplace{}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(ngsi, b)

	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "{\"verbose\":true,\"service\":\"fiware\",\"path\":\"/orion\",\"add_path\":\"/federation\",\"url\":\"/v3/queryContext\"}", string(body))
}

func TestRegProxyConfigValue(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	regProxyGlobal = &regProxyParam{
		ngsi:      ngsi,
		verbose:   true,
		bearer:    true,
		replace:   false,
		startTime: time.Now(),
		mutex:     &sync.Mutex{},
	}

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"

	verbose := false
	req := &regProxyReplace{
		Verbose: &verbose,
		Service: &tenant,
		Path:    &scope,
		AddPath: &addScope,
		URL:     &url,
	}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(ngsi, b)

	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "{\"verbose\":false,\"service\":\"fiware\",\"path\":\"/orion\",\"add_path\":\"/federation\",\"url\":\"/v3/queryContext\"}", string(body))
	assert.Equal(t, true, regProxyGlobal.replace)
}

func TestRegProxyConfigEmpty(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"
	regProxyGlobal = &regProxyParam{
		ngsi:      ngsi,
		verbose:   true,
		bearer:    true,
		tenant:    &tenant,
		scope:     &scope,
		addScope:  &addScope,
		replace:   true,
		url:       &url,
		startTime: time.Now(),
		mutex:     &sync.Mutex{},
	}

	empty := ""
	req := &regProxyReplace{
		Service: &empty,
		Path:    &empty,
		AddPath: &empty,
		URL:     &empty,
	}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(ngsi, b)

	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "{\"verbose\":true}", string(body))
	assert.Equal(t, false, regProxyGlobal.replace)
}

func TestRegProxyConfigErrorUnmarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"
	regProxyGlobal = &regProxyParam{
		ngsi:      ngsi,
		verbose:   true,
		bearer:    true,
		tenant:    &tenant,
		scope:     &scope,
		addScope:  &addScope,
		replace:   true,
		url:       &url,
		startTime: time.Now(),
		mutex:     &sync.Mutex{},
	}

	setJSONDecodeErr(ngsi, 0)

	req := &regProxyReplace{}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(ngsi, b)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "{\"error\":\"json error\"}", string(body))
}

func TestRegProxyConfigErrorMarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"
	regProxyGlobal = &regProxyParam{
		ngsi:      ngsi,
		verbose:   true,
		bearer:    true,
		tenant:    &tenant,
		scope:     &scope,
		addScope:  &addScope,
		replace:   true,
		url:       &url,
		startTime: time.Now(),
		mutex:     &sync.Mutex{},
	}

	setJSONEncodeErr(ngsi, 0)

	req := &regProxyReplace{}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(ngsi, b)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "{\"error\":\"json error\"}", string(body))
}

func TestRegProxyHealthCmd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go": "regproxy", "version": "0.8.4-next (git_hash:7392ed9962f42c6eca1f894465b6f7450d65958a)", "health": "OK", "csource": "https://orion.letsfiware.jp", "verbose": false, "uptime": "0 d, 0 h, 51 m, 51 s", "timesent": 0, "success": 0, "failure": 0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy"})

	err := regProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"ngsi-go\": \"regproxy\", \"version\": \"0.8.4-next (git_hash:7392ed9962f42c6eca1f894465b6f7450d65958a)\", \"health\": \"OK\", \"csource\": \"https://orion.letsfiware.jp\", \"verbose\": false, \"uptime\": \"0 d, 0 h, 51 m, 51 s\", \"timesent\": 0, \"success\": 0, \"failure\": 0}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyHealthCmdPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go": "regproxy", "version": "0.8.4-next (git_hash:7392ed9962f42c6eca1f894465b6f7450d65958a)", "health": "OK", "csource": "https://orion.letsfiware.jp", "verbose": false, "uptime": "0 d, 0 h, 51 m, 51 s", "timesent": 0, "success": 0, "failure": 0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--pretty"})

	err := regProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"ngsi-go\": \"regproxy\",\n  \"version\": \"0.8.4-next (git_hash:7392ed9962f42c6eca1f894465b6f7450d65958a)\",\n  \"health\": \"OK\",\n  \"csource\": \"https://orion.letsfiware.jp\",\n  \"verbose\": false,\n  \"uptime\": \"0 d, 0 h, 51 m, 51 s\",\n  \"timesent\": 0,\n  \"success\": 0,\n  \"failure\": 0\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyHealthCmdErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := regProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyHealthCmdErrorNewClient(t *testing.T) {
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

	err := regProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyHealthCmdErrorHTTP(t *testing.T) {
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
	_ = set.Parse([]string{"--host=regproxy"})

	err := regProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyHealthCmdErrorStatusCode(t *testing.T) {
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
	_ = set.Parse([]string{"--host=regproxy"})

	err := regProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyHealthCmdIotaErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go": "regproxy", "version": "0.8.4-next (git_hash:7392ed9962f42c6eca1f894465b6f7450d65958a)", "health": "OK", "csource": "https://orion.letsfiware.jp", "verbose": false, "uptime": "0 d, 0 h, 51 m, 51 s", "timesent": 0, "success": 0, "failure": 0}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--pretty"})

	setJSONIndentError(ngsi)

	err := regProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegProxyConfigCmd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=on", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext"})

	err := regProxyConfigCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyConfigCmdVerboseOff(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":false,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":false,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=off", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext"})

	err := regProxyConfigCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"verbose":false,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`
		assert.Equal(t, expected, actual)
	}
}
func TestRegProxyConfigCmdPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=on", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext", "--pretty"})

	err := regProxyConfigCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"verbose\": true,\n  \"service\": \"federatin\",\n  \"path\": \"/fiware\",\n  \"add_path\": \"/orion\",\n  \"url\": \"/v3/queryContext\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyConfigCmdErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := regProxyConfigCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyConfigCmdErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})

	err := regProxyConfigCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyConfigCmdErrorVerbose(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=unknown", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext"})

	err := regProxyConfigCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error: set on or off to --verbose option", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyConfigCmdErrorJSONMarshl(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=on", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext"})

	setJSONEncodeErr(ngsi, 2)

	err := regProxyConfigCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyConfigCmdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=on", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext"})

	err := regProxyConfigCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyConfigCmdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=on", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext"})

	err := regProxyConfigCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegProxyConfigCmdIotaErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,verbose,replaceService,replacePath,addPath,replaceURL")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=regproxy", "--verbose=on", "--replaceService=federatin", "--replacePath=/fiware", "--addPath=/orion", "--replaceURL=/v3/queryContext", "--pretty"})

	setJSONIndentError(ngsi)

	err := regProxyConfigCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
