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
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

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

	err := regProxy(c)

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

	err := regProxy(c)

	assert.NoError(t, err)
}

func TestRegProxyErrorHost(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := regProxy(c)

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

	err := regProxy(c)

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

	err := regProxy(c)

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

	err := regProxy(c)

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

	err := regProxy(c)

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

	err := regProxy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestRegProxyHanderGetHealth(t *testing.T) {
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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/health", nil)
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusOK

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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	regProxyHandler(got, req)

	expected := http.StatusBadRequest

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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}}

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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}}

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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: false, bearer: false, mutex: &sync.Mutex{}}

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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}}

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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}}

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
	regProxyGlobal = &regProxyParam{ngsi: ngsi, client: client, http: mockHTTP, verbose: true, bearer: true, mutex: &sync.Mutex{}}

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
