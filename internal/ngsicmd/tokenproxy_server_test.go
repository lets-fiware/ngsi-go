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

func TestTokenProxy(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=0.0.0.0", "--port=1028"})

	err := tokenProxyServer(c)

	assert.NoError(t, err)
}

func TestTokenProxyOptions(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=0.0.0.0", "--port=1028",
		"--idmHost=http://keyrock:3000", "--clientId=a1a6048b-df1d-4d4f-9a08-5cf836041d14", "--clientSecret=e4cc0147-e38f-4211-b8ad-8ae5e6a107f9"})

	err := tokenProxyServer(c)

	assert.NoError(t, err)
}

func TestTokenProxyHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{}

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=0.0.0.0", "--port=1028", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := tokenProxyServer(c)

	assert.NoError(t, err)
}

func TestTokenProxyError(t *testing.T) {
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
func TestTokenProxyErrorKey(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=0.0.0.0", "--port=1028", "--https"})

	err := tokenProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestTokenProxyErrorCert(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=0.0.0.0", "--port=1028", "--https", "--key=test.key"})

	err := tokenProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestTokenProxyErrorIdmHost(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idmHost=:", "--host=0.0.0.0", "--port=1028", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := tokenProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "parse \":\": missing protocol scheme", ngsiErr.Message)
	}
}

func TestTokenProxyErrorHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--port=1028", "--https", "--key=test.key", "--cert=test.cert", "--verbose"})

	err := tokenProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestTokenProxyErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	gNetLib = &MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=0.0.0.0", "--port=1028"})

	err := tokenProxyServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestTokenProxyRootHandler(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,port,key,cert,idmHost,clientId,clientSecret")
	setupFlagBool(set, "verbose,https")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi, http: mockHTTP, verbose: true, gLock: &sync.Mutex{}}

	req := httptest.NewRequest(http.MethodGet, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	tokenProxyRootHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHealthHandler(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		http:    mockHTTP,
		verbose: true,
		gLock:   &sync.Mutex{},
		idmHost: u,
	}

	req := httptest.NewRequest(http.MethodGet, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	tokenProxyHealthHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHealthHandlerError(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		http:    mockHTTP,
		verbose: true,
		gLock:   &sync.Mutex{},
		idmHost: u,
	}

	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	tokenProxyHealthHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderErrorMethod(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	mockHTTP := NewMockHTTP()
	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		http:      mockHTTP,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	req := httptest.NewRequest(http.MethodGet, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	tokenProxyHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderPostToken(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`grant_type=password&username=admin@test.com&password=1234`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		http:      mock,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	tokenProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderPostRevoke(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`token=9721b640cafb39882cf9a71d2249760134c0073d&token_type_hint=refresh_token`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		http:      mock,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"token":"9721b640cafb39882cf9a71d2249760134c0073d"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	tokenProxyHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderErrorToken(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`grant_type=password&username=admin@test.com&password=1234`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		http:      mock,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	tokenProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokeProxyRequestToken006 parameter error\"}", got.Body.String())
}

func TestTokenProxyHanderErrorRevoke(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`token=9721b640cafb39882cf9a71d2249760134c0073d&token_type_hint=refresh_token`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		http:      mock,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	tokenProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokenProxyRevokeToken006 parameter error\"}", got.Body.String())
}

func TestTokenProxyHanderErrorHTTP(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,stderr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--stderr=info"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`grant_type=password&username=admin@test.com&password=1234`)
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		http:      mock,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	tokenProxyHandler(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokenProxyHandler004 http error\"}", got.Body.String())
}

func TestTokenProxyResposeError(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	tokenProxyGlobal = &tokenProxyParam{
		failure: 1,
		gLock:   &sync.Mutex{},
	}
	got := httptest.NewRecorder()

	tokenProxyResposeError(ngsi, got, http.StatusBadRequest, errors.New("test"))

	assert.Equal(t, int64(2), tokenProxyGlobal.failure)
	assert.Equal(t, "{\"error\":\"test\"}", got.Body.String())
}

func TestTokenProxyRequestTokenJSON(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   false,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234", string(actual))
	}
}

func TestTokenProxyRequestTokenJSONVerbose(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234", string(actual))
	}
}

func TestTokenProxyRequestTokenRefreshToken(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=refresh_token&refresh_token=2981fed8a6810c8a6131eb445f029dcb14a4eff3", string(actual))
	}
}

func TestTokenProxyRequestTokenURLEncorded(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`username=admin@test.com&password=1234&scope=openid`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	actual, err := tokenProxyRequestToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234&scope=openid", string(actual))
	}
}

func TestTokenProxyRequestTokenErrorContentType(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`username=admin@test.com&password=1234`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)

	_, err := tokenProxyRequestToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing Content-Type", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorUnknownContentType(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`username=admin@test.com&password=1234`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "unknown")

	_, err := tokenProxyRequestToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Content-Type error", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorJSONUnmarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	setJSONDecodeErr(ngsi, 0)

	_, err := tokenProxyRequestToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorURLEncoded(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", nil)
	req.Body = nil
	req.Form = nil
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRequestToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing form body", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorUnknownParam(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`user=admin@test.com`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err := tokenProxyRequestToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: user", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorParam(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(``)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err := tokenProxyRequestToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "parameter error", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenJSON(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   false,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRevokeToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "token=2981fed8a6810c8a6131eb445f029dcb14a4eff3&token_type_hint=refresh_token", string(actual))
	}
}

func TestTokenProxyRevokeTokenJSONVerbose(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRevokeToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "token=2981fed8a6810c8a6131eb445f029dcb14a4eff3&token_type_hint=refresh_token", string(actual))
	}
}

func TestTokenProxyRevokeTokenURLEncorded(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`token=2981fed8a6810c8a6131eb445f029dcb14a4eff3`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	actual, err := tokenProxyRevokeToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "token=2981fed8a6810c8a6131eb445f029dcb14a4eff3&token_type_hint=refresh_token", string(actual))
	}
}

func TestTokenProxyRevokeTokenErrorContentType(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`token=2981fed8a6810c8a6131eb445f029dcb14a4eff3`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)

	_, err := tokenProxyRevokeToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing Content-Type", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorUnknownContentType(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`token=2981fed8a6810c8a6131eb445f029dcb14a4eff3`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "unknown")

	_, err := tokenProxyRevokeToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Content-Type error", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorJSONUnmarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")

	setJSONDecodeErr(ngsi, 0)

	_, err := tokenProxyRevokeToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorURLEncoded(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", nil)
	req.Body = nil
	req.Form = nil
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRevokeToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing form body", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorUnkownParam(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`user=admin@test.com`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err := tokenProxyRevokeToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: user", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorParam(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`token_type_hint=password`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	_, err := tokenProxyRevokeToken(ngsi, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "parameter error", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenURLEncoded(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(ngsi, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234", string(actual))
	}
}

func TestTokenProxyGetStat(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	b := tokenProxyGetStat()

	stat := &tokenProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "OK")
}

func TestTokenProxyGetStatError(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	u, _ := url.Parse("http://keyrock:3000")
	tokenProxyGlobal = &tokenProxyParam{ngsi: ngsi,
		verbose:   true,
		gLock:     &sync.Mutex{},
		idmHost:   u,
		RevokeURL: u,
	}

	setJSONEncodeErr(ngsi, 0)

	b := tokenProxyGetStat()

	stat := &tokenProxyStat{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "NG")
}

func TestTokenProxyHealthCmd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"tokenproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","idm":"http://keyrock:3000/oauth2/token","clientId":"a1a6048b-df1d-4d4f-9a08-5cf836041d14","clientSecret":"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9","verbose":true,"uptime":"0 d, 1 h, 55 m, 39 s","timesent":3,"success":1,"revoke":1,"failure":1}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=tokenproxy"})

	err := tokenProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"ngsi-go\":\"tokenproxy\",\"version\":\"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)\",\"health\":\"OK\",\"idm\":\"http://keyrock:3000/oauth2/token\",\"clientId\":\"a1a6048b-df1d-4d4f-9a08-5cf836041d14\",\"clientSecret\":\"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9\",\"verbose\":true,\"uptime\":\"0 d, 1 h, 55 m, 39 s\",\"timesent\":3,\"success\":1,\"revoke\":1,\"failure\":1}"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenProxyHealthCmdPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"tokenproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","idm":"http://keyrock:3000/oauth2/token","clientId":"a1a6048b-df1d-4d4f-9a08-5cf836041d14","clientSecret":"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9","verbose":true,"uptime":"0 d, 1 h, 55 m, 39 s","timesent":3,"success":1,"revoke":1,"failure":1}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=tokenproxy", "--pretty"})

	err := tokenProxyHealthCmd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"ngsi-go\": \"tokenproxy\",\n  \"version\": \"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)\",\n  \"health\": \"OK\",\n  \"idm\": \"http://keyrock:3000/oauth2/token\",\n  \"clientId\": \"a1a6048b-df1d-4d4f-9a08-5cf836041d14\",\n  \"clientSecret\": \"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9\",\n  \"verbose\": true,\n  \"uptime\": \"0 d, 1 h, 55 m, 39 s\",\n  \"timesent\": 3,\n  \"success\": 1,\n  \"revoke\": 1,\n  \"failure\": 1\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenProxyHealthCmdErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := tokenProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTokenProxyHealthCmdErrorNewClient(t *testing.T) {
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

	err := tokenProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTokenProxyHealthCmdErrorHTTP(t *testing.T) {
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
	_ = set.Parse([]string{"--host=tokenproxy"})

	err := tokenProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTokenProxyHealthCmdErrorStatusCode(t *testing.T) {
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
	_ = set.Parse([]string{"--host=tokenproxy"})

	err := tokenProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTokenProxyHealthCmdIotaErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"tokenproxy","version":"0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab)","health":"OK","idm":"http://keyrock:3000/oauth2/token","clientId":"a1a6048b-df1d-4d4f-9a08-5cf836041d14","clientSecret":"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9","verbose":true,"uptime":"0 d, 1 h, 55 m, 39 s","timesent":3,"success":1,"revoke":1,"failure":1}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=tokenproxy", "--pretty"})

	setJSONIndentError(ngsi)

	err := tokenProxyHealthCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
