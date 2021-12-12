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
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestRegProxy(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion", "--rhost", "0.0.0.0", "--port", "1028", "--url", "/"})

	err := regProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRegProxyOptions(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion", "--rhost", "0.0.0.0", "--port", "1028", "--url", "/", "--replaceService", "fiware", "--replacePath", "/orion", "--replaceURL", "/v3/queryConxt", "--addPath", "/federation"})

	err := regProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRegProxyHTTPS(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion", "--rhost", "0.0.0.0", "--port", "1028", "--url", "/", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose"})

	err := regProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRegProxyErrorKey(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion", "--rhost", "0.0.0.0", "--port", "1028", "--url", "/", "--https"})

	err := regProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestRegProxyErrorCert(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion", "--rhost", "0.0.0.0", "--port", "1028", "--url", "/", "--https", "--key", "test.key"})

	err := regProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestRegProxyErrorHTTPS(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion", "--rhost", "0.0.0.0", "--port", "1028", "--url", "/", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	err := regProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestRegProxyErrorHTTP(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion", "--rhost", "0.0.0.0", "--port", "1028", "--url", "/"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	err := regProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestRegProxyRootHandler(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	h := &regProxyRootHandler{ngsi: c.Ngsi}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHealthHandler(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	h := &regProxyHealthHandler{
		ngsi:   c.Ngsi,
		host:   "orion",
		config: &regProxyConfigParam{},
		stat:   &regProxyStat{mutex: &sync.Mutex{}},
	}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHealthHandlerError(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	h := &regProxyHealthHandler{
		ngsi:   c.Ngsi,
		host:   "orion",
		config: &regProxyConfigParam{},
		stat:   &regProxyStat{mutex: &sync.Mutex{}},
	}

	req := httptest.NewRequest(http.MethodPost, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyConfigHandler(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	h := &regProxyConfigHandler{
		ngsi: c.Ngsi,
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
	}

	b := bytes.NewReader([]byte(`{"verbose":false}`))
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/", b)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyConfigHandlerError(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	h := &regProxyConfigHandler{
		ngsi: c.Ngsi,
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorGet(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   helper.NewMockHTTP(),
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorMethod(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   helper.NewMockHTTP(),
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	req := httptest.NewRequest(http.MethodDelete, "http://regProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderPost(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	c.Client.Server.IdmType = "tokenproxy"
	c.Client.Server.IdmHost = "/token"
	c.Client.Server.Username = "testuser"
	c.Client.Server.Password = "1234"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	reqRes2.ResHeader = http.Header{"Context-Length": []string{"0"}}

	mockHTTP := helper.NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   mockHTTP,
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	req.Header.Set("Authorization", "Basic 1234")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderPostOptions(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	c.Client.Server.IdmType = "tokenproxy"
	c.Client.Server.IdmHost = "/token"
	c.Client.Server.Username = "testuser"
	c.Client.Server.Password = "1234"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)
	reqRes2.ResHeader = http.Header{"Context-Length": []string{"0"}}

	mockHTTP := helper.NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   mockHTTP,
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose:  true,
			bearer:   true,
			tenant:   helper.StrPtr("fiware"),
			scope:    helper.StrPtr("/orion"),
			addScope: helper.StrPtr("/federation"),
			replace:  true,
			url:      helper.StrPtr("/v3/queryContext"),
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderPostXAuth(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	c.Ngsi.HTTP = mock

	c.Client.Server.IdmType = "tokenproxy"
	c.Client.Server.IdmHost = "/token"
	c.Client.Server.Username = "testuser"
	c.Client.Server.Password = "1234"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)

	mockHTTP := helper.NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   mockHTTP,
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorHost(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	c.Client.Server.ServerHost = ":"
	c.Client.Server.IdmType = "tokenproxy"
	c.Client.Server.IdmHost = "/token"
	c.Client.Server.Username = "testuser"
	c.Client.Server.Password = "1234"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)

	mockHTTP := helper.NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   mockHTTP,
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderErrorToken(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	c.Ngsi.HTTP = mock

	c.Client.Server.IdmType = "tokenproxy"
	c.Client.Server.IdmHost = "/token"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)

	mockHTTP := helper.NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   mockHTTP,
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestRegProxyHanderBadRequest(t *testing.T) {
	c := setupTest([]string{"regproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	c.Ngsi.HTTP = mock

	c.Client.Server.IdmType = "tokenproxy"
	c.Client.Server.IdmHost = "/token"
	c.Client.Server.Username = "testuser"
	c.Client.Server.Password = "1234"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/"
	reqRes2.ResBody = []byte(`[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device","temperature":{"type":"Number","value":25.47,"metadata":{"TimeInstant":{"type":"DateTime","value":"2020-07-12T05:00:52.00Z"}}}}]`)

	mockHTTP := helper.NewMockHTTP()
	mockHTTP.ReqRes = append(mockHTTP.ReqRes, reqRes2)

	h := &regProxyHandler{
		ngsi:   c.Ngsi,
		client: c.Client,
		http:   mockHTTP,
		mutex:  &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"entities":[{"id":"urn:ngsi-ld:Device:uDr8vgsJ0Xbe","type":"Device"}],"attrs":["temperature"]}`)
	req := httptest.NewRequest(http.MethodPost, "http://regProxy/v2/op/query", reqBody)
	req.Header.Set("User-Agent", "NGSI Go")
	req.Header.Set("Fiware-Service", "iot")
	req.Header.Set("Fiware-ServicePath", "/")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

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
	h := &regProxyHandler{
		mutex: &sync.Mutex{},
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
		},
		stat: &regProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	regProxyFailureUp(h)

	assert.Equal(t, int64(1), h.stat.failure)
}

func TestRegProxyGetStat(t *testing.T) {
	ngsi := helper.SetupTestInitNGSI()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"

	h := &regProxyHealthHandler{
		ngsi: ngsi,
		host: "http://orion",
		config: &regProxyConfigParam{
			verbose:  true,
			bearer:   true,
			tenant:   &tenant,
			scope:    &scope,
			addScope: &addScope,
			replace:  true,
			url:      &url,
		},
		stat: &regProxyStat{
			mutex:     &sync.Mutex{},
			startTime: time.Now(),
		},
	}

	b := regProxyGetStat(h)

	stat := &regProxyStatInfo{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "OK")
}

func TestRegProxyGetStatError(t *testing.T) {
	ngsi := helper.SetupTestInitNGSI()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"

	h := &regProxyHealthHandler{
		ngsi: ngsi,
		host: "http://orion",
		config: &regProxyConfigParam{
			verbose:  true,
			bearer:   true,
			tenant:   &tenant,
			scope:    &scope,
			addScope: &addScope,
			replace:  true,
			url:      &url,
		},
		stat: &regProxyStat{
			mutex:     &sync.Mutex{},
			startTime: time.Now(),
		},
	}
	helper.SetJSONEncodeErr(ngsi, 0)

	b := regProxyGetStat(h)

	stat := &regProxyStatInfo{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "NG")
}

func TestRegProxyConfig(t *testing.T) {
	ngsi := helper.SetupTestInitNGSI()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"

	h := &regProxyConfigHandler{
		ngsi: ngsi,
		config: &regProxyConfigParam{
			verbose:  true,
			bearer:   true,
			tenant:   &tenant,
			scope:    &scope,
			addScope: &addScope,
			replace:  true,
			url:      &url,
		},
	}

	req := &regProxyReplace{}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(h, b)

	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "{\"verbose\":true,\"service\":\"fiware\",\"path\":\"/orion\",\"add_path\":\"/federation\",\"url\":\"/v3/queryContext\"}", string(body))
}

func TestRegProxyConfigValue(t *testing.T) {
	ngsi := helper.SetupTestInitNGSI()

	h := &regProxyConfigHandler{
		ngsi: ngsi,
		config: &regProxyConfigParam{
			verbose: true,
			bearer:  true,
			replace: false,
		},
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

	status, body := regProxyConfig(h, b)

	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "{\"verbose\":false,\"service\":\"fiware\",\"path\":\"/orion\",\"add_path\":\"/federation\",\"url\":\"/v3/queryContext\"}", string(body))
	assert.Equal(t, true, h.config.replace)
}

func TestRegProxyConfigEmpty(t *testing.T) {
	ngsi := helper.SetupTestInitNGSI()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"

	h := &regProxyConfigHandler{
		ngsi: ngsi,
		config: &regProxyConfigParam{
			verbose:  true,
			bearer:   true,
			tenant:   &tenant,
			scope:    &scope,
			addScope: &addScope,
			replace:  true,
			url:      &url,
		},
	}

	empty := ""
	req := &regProxyReplace{
		Service: &empty,
		Path:    &empty,
		AddPath: &empty,
		URL:     &empty,
	}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(h, b)

	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "{\"verbose\":true}", string(body))
	assert.Equal(t, false, h.config.replace)
}

func TestRegProxyConfigErrorUnmarshal(t *testing.T) {
	ngsi := helper.SetupTestInitNGSI()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"

	h := &regProxyConfigHandler{
		ngsi: ngsi,
		config: &regProxyConfigParam{
			verbose:  true,
			bearer:   true,
			tenant:   &tenant,
			scope:    &scope,
			addScope: &addScope,
			replace:  true,
			url:      &url,
		},
	}

	helper.SetJSONDecodeErr(ngsi, 0)

	req := &regProxyReplace{}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(h, b)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "{\"error\":\"json error\"}", string(body))
}

func TestRegProxyConfigErrorMarshal(t *testing.T) {
	ngsi := helper.SetupTestInitNGSI()

	tenant := "fiware"
	scope := "/orion"
	addScope := "/federation"
	url := "/v3/queryContext"

	h := &regProxyConfigHandler{
		ngsi: ngsi,
		config: &regProxyConfigParam{
			verbose:  true,
			bearer:   true,
			tenant:   &tenant,
			scope:    &scope,
			addScope: &addScope,
			replace:  true,
			url:      &url,
		},
	}

	helper.SetJSONEncodeErr(ngsi, 0)

	req := &regProxyReplace{}

	b, _ := json.Marshal(req)

	status, body := regProxyConfig(h, b)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "{\"error\":\"json error\"}", string(body))
}

func TestRegProxyHealthCmd(t *testing.T) {
	c := setupTest([]string{"regproxy", "health", "--host", "regproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go": "regproxy", "version": "0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)", "health": "OK", "csource": "https://orion.letsfiware.jp", "verbose": false, "uptime": "0 d, 0 h, 51 m, 51 s", "timesent": 0, "success": 0, "failure": 0}`)
	helper.SetClientHTTP(c, reqRes)

	err := regProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"ngsi-go\": \"regproxy\", \"version\": \"0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)\", \"health\": \"OK\", \"csource\": \"https://orion.letsfiware.jp\", \"verbose\": false, \"uptime\": \"0 d, 0 h, 51 m, 51 s\", \"timesent\": 0, \"success\": 0, \"failure\": 0}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyHealthCmdPretty(t *testing.T) {
	c := setupTest([]string{"regproxy", "health", "--host", "regproxy", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go": "regproxy", "version": "0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)", "health": "OK", "csource": "https://orion.letsfiware.jp", "verbose": false, "uptime": "0 d, 0 h, 51 m, 51 s", "timesent": 0, "success": 0, "failure": 0}`)

	helper.SetClientHTTP(c, reqRes)

	err := regProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"ngsi-go\": \"regproxy\",\n  \"version\": \"0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)\",\n  \"health\": \"OK\",\n  \"csource\": \"https://orion.letsfiware.jp\",\n  \"verbose\": false,\n  \"uptime\": \"0 d, 0 h, 51 m, 51 s\",\n  \"timesent\": 0,\n  \"success\": 0,\n  \"failure\": 0\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyHealthCmdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"regproxy", "health", "--host", "regproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := regProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestRegProxyHealthCmdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"regproxy", "health", "--host", "regproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := regProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestRegProxyHealthCmdErrorPretty(t *testing.T) {
	c := setupTest([]string{"regproxy", "health", "--host", "regproxy", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go": "regproxy", "version": "0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)", "health": "OK", "csource": "https://orion.letsfiware.jp", "verbose": false, "uptime": "0 d, 0 h, 51 m, 51 s", "timesent": 0, "success": 0, "failure": 0}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := regProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegProxyConfigCmd(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "on", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)

	helper.SetClientHTTP(c, reqRes)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyConfigCmdVerboseOff(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "off", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":false,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":false,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)

	helper.SetClientHTTP(c, reqRes)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"verbose":false,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyConfigCmdPretty(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "on", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)

	helper.SetClientHTTP(c, reqRes)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"verbose\": true,\n  \"service\": \"federatin\",\n  \"path\": \"/fiware\",\n  \"add_path\": \"/orion\",\n  \"url\": \"/v3/queryContext\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegProxyConfigCmdErrorVerbose(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "unknown", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: set on or off to --verbose option", ngsiErr.Message)
	}
}

func TestRegProxyConfigCmdErrorJSONMarshl(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "on", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegProxyConfigCmdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "on", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestRegProxyConfigCmdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "on", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestRegProxyConfigCmdErrorPretty(t *testing.T) {
	c := setupTest([]string{"regproxy", "config", "--host", "regproxy", "--verbose", "on", "--replaceService", "federatin", "--replacePath", "/fiware", "--addPath", "/orion", "--replaceURL", "/v3/queryContext", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/config"
	reqRes.ReqData = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)
	reqRes.ResBody = []byte(`{"verbose":true,"service":"federatin","path":"/fiware","add_path":"/orion","url":"/v3/queryContext"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := regProxyConfigCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
