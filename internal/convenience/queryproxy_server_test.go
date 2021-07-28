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
	"net/url"
	"sync"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestQueryProxy(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--qhost", "0.0.0.0", "--port", "1028"})

	err := queryProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestQueryProxyReplaceURL(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--qhost", "0.0.0.0", "--port", "1028", "--replaceURL", "http://replaceURL"})

	err := queryProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestQueryProxyOptions(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--qhost", "0.0.0.0", "--port", "1028", "--replaceURL", "/v3/entities"})

	err := queryProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestQueryProxyHTTPS(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--qhost", "0.0.0.0", "--port", "1028", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose"})

	err := queryProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestQueryProxyErrorKey(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--qhost", "0.0.0.0", "--port", "1028", "--https"})

	err := queryProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestQueryProxyErrorCert(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--qhost", "0.0.0.0", "--port", "1028", "--https", "--key", "test.key"})

	err := queryProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestQueryProxyErrorHTTPS(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--port", "1028", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	err := queryProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestQueryProxyErrorHTTP(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion", "--qhost", "0.0.0.0", "--port", "1028"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	err := queryProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestQueryProxyRootHandler(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	h := &queryProxyRootHandler{ngsi: c.Ngsi}

	req := httptest.NewRequest(http.MethodGet, "http://queryProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHealthHandler(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	h := &queryProxyHealthHandler{
		ngsi:    c.Ngsi,
		broker:  helper.UrlParse("http://orion:1026").String(),
		verbose: true,
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	req := httptest.NewRequest(http.MethodGet, "http://queryProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHealthHandlerError(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	h := &queryProxyHealthHandler{
		ngsi:    c.Ngsi,
		broker:  helper.UrlParse("http://orion:1026").String(),
		verbose: true,
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderPost(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	c.Client.Server.IdmType = ""

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = helper.StrPtr("options=keyValues&type=Deivce")
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = http.Header{"Content-Type": []string{"application/json"}}

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &queryProxyHandler{
		ngsi:    c.Ngsi,
		url:     helper.UrlParse("http://orion:1026/v2/ex/entities"),
		client:  c.Client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	req.Header.Set("Authorization", "Bearer 23d7500c85d2f05ffb102e1b7165e325d75f4290")

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderPostIDM(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	c.Client.Server.IdmType = "basic"
	c.Client.Server.Username = "fiware"
	c.Client.Server.Password = "1234"

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = helper.StrPtr("options=keyValues&type=Deivce")
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = http.Header{"Content-Type": []string{"application/json"}}

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &queryProxyHandler{
		ngsi:    c.Ngsi,
		url:     helper.UrlParse("http://orion:1026/v2/ex/entities"),
		client:  c.Client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderErrorMethod(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	h := &queryProxyHandler{
		ngsi:    c.Ngsi,
		url:     helper.UrlParse("http://orion:1026"),
		client:  c.Client,
		http:    helper.NewMockHTTP(),
		verbose: true,
		mutex:   &sync.Mutex{},
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	req := httptest.NewRequest(http.MethodGet, "http://queryProxy/", nil)

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestQueryProxyHanderPostErrorURL(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = helper.StrPtr("options=keyValues&type=Deivce")
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = http.Header{"Content-Type": []string{"application/json"}}

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &queryProxyHandler{
		ngsi:    c.Ngsi,
		url:     &url.URL{Scheme: ":"},
		client:  c.Client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"queryProxyHandler003 parse \"::\": missing protocol scheme\"}", got.Body.String())
}

func TestQueryProxyHanderPostErrorIDM(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	c.Client.Server.IdmType = "unknown"

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = helper.StrPtr("options=keyValues&type=Deivce")
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = http.Header{"Content-Type": []string{"application/json"}}

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &queryProxyHandler{
		ngsi:    c.Ngsi,
		url:     helper.UrlParse("http://orion:1026/v2/ex/entities"),
		client:  c.Client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"queryProxyHandler004 unknown idm type: unknown\"}", got.Body.String())
}

func TestQueryProxyHanderPostErrorParam(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	c.Client.Server.IdmType = ""

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.RawQuery = helper.StrPtr("options=keyValues&type=Deivce")
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = http.Header{"Content-Type": []string{"application/json"}}

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &queryProxyHandler{
		ngsi:    c.Ngsi,
		url:     helper.UrlParse("http://orion:1026/v2/ex/entities"),
		client:  c.Client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	reqBody := bytes.NewBufferString(`options`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/json")

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokeProxyRequestToken003 Content-Type error\"}", got.Body.String())
}

func TestQueryProxyHanderPostErrorHTTP(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	c.Client.Server.IdmType = ""

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Path = "/v2/ex/entities"
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Err = errors.New("http error")
	reqRes.RawQuery = helper.StrPtr("options=keyValues&type=Deivce")
	reqRes.ResBody = []byte("")
	reqRes.ResHeader = http.Header{"Content-Type": []string{"application/json"}}

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &queryProxyHandler{
		ngsi:    c.Ngsi,
		url:     helper.UrlParse("http://orion:1026/v2/ex/entities"),
		client:  c.Client,
		http:    mock,
		verbose: true,
		mutex:   &sync.Mutex{},
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")
	req.Header.Set("Authorization", "Bearer 23d7500c85d2f05ffb102e1b7165e325d75f4290")

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"queryProxyHandler006 http error\"}", got.Body.String())
}

func TestQueryProxyResposeError(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	h := &queryProxyHandler{
		ngsi: c.Ngsi,
		stat: &queryProxyStat{mutex: &sync.Mutex{}, failure: 1},
	}
	got := httptest.NewRecorder()

	queryProxyResposeError(h, got, http.StatusBadRequest, errors.New("test"))

	assert.Equal(t, int64(2), h.stat.failure)
	assert.Equal(t, "{\"error\":\"test\"}", got.Body.String())
}

func TestQueryProxySetQueryParam(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")

	u := &url.URL{}

	err := queryProxySetQueryParam(c.Ngsi, req, u)

	if assert.NoError(t, err) {
		assert.Equal(t, "options=keyValues&type=Deivce", u.RawQuery)
	}

}

func TestQueryProxySetQueryParamErrorContentTypeMissing(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)

	err := queryProxySetQueryParam(c.Ngsi, req, &url.URL{})

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing Content-Type", ngsiErr.Message)
	}
}

func TestQueryProxySetQueryParamErrorParseForm(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	reqBody := bytes.NewBufferString(`options=keyValues&type=Deivce`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = nil
	req.Form = nil

	err := queryProxySetQueryParam(c.Ngsi, req, &url.URL{})

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing form body", ngsiErr.Message)
	}
}

func TestQueryProxySetQueryParamErrorContentType(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://queryProxy/v2/ex/entities", reqBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FIWARE-Service", "fiware")
	req.Header.Set("FIWARE-ServicePath", "/iot")

	err := queryProxySetQueryParam(c.Ngsi, req, &url.URL{})

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Content-Type error", ngsiErr.Message)
	}
}

func TestQueryProxyGetStat(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	h := &queryProxyHealthHandler{
		ngsi:    c.Ngsi,
		broker:  helper.UrlParse("http://orion:1026/v2/ex/entities").String(),
		verbose: true,
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	b := queryProxyGetStat(h)

	stat := &queryProxyStatInfo{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "OK")
}

func TestQueryProxyGetStatError(t *testing.T) {
	c := setupTest([]string{"queryproxy", "server", "--host", "orion"})

	h := &queryProxyHealthHandler{
		ngsi:    c.Ngsi,
		broker:  helper.UrlParse("http://orion:1026/v2/ex/entities").String(),
		verbose: true,
		stat:    &queryProxyStat{mutex: &sync.Mutex{}},
	}

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	b := queryProxyGetStat(h)

	stat := &queryProxyStatInfo{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "NG")
}

func TestQueryProxyHealthCmd(t *testing.T) {
	c := setupTest([]string{"queryproxy", "health", "--host", "queryproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"queryproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)

	helper.SetClientHTTP(c, reqRes)

	err := queryProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"ngsi-go\":\"queryproxy\",\"version\":\"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)\",\"health\":\"OK\",\"orion\":\"http://orion:1026/v2/entities\",\"verbose\":true,\"uptime\":\"0 d, 2 h, 49 m, 1 s\",\"timesent\":3,\"success\":3,\"failure\":0}"
		assert.Equal(t, expected, actual)
	}
}

func TestQueryProxyHealthCmdPretty(t *testing.T) {
	c := setupTest([]string{"queryproxy", "health", "--host", "queryproxy", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"queryproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)

	helper.SetClientHTTP(c, reqRes)

	err := queryProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"ngsi-go\": \"queryproxy\",\n  \"version\": \"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)\",\n  \"health\": \"OK\",\n  \"orion\": \"http://orion:1026/v2/entities\",\n  \"verbose\": true,\n  \"uptime\": \"0 d, 2 h, 49 m, 1 s\",\n  \"timesent\": 3,\n  \"success\": 3,\n  \"failure\": 0\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQueryProxyHealthCmdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"queryproxy", "health", "--host", "queryproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := queryProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestQueryProxyHealthCmdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"queryproxy", "health", "--host", "queryproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := queryProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestQueryProxyHealthCmdIotaErrorPretty(t *testing.T) {
	c := setupTest([]string{"queryproxy", "health", "--host", "queryproxy", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"queryproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","orion":"http://orion:1026/v2/entities","verbose":true,"uptime":"0 d, 2 h, 49 m, 1 s","timesent":3,"success":3,"failure":0}`)
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := queryProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
