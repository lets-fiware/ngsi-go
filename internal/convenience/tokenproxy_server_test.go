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

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestTokenProxy(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "0.0.0.0", "--port", "1028"})

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTokenProxyOptions(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "0.0.0.0", "--port", "1028", "--idmHost", "http://keyrock:3000", "--clientId", "a1a6048b-df1d-4d4f-9a08-5cf836041d14", "--clientSecret", "e4cc0147-e38f-4211-b8ad-8ae5e6a107f9"})

	buf := new(bytes.Buffer)
	c.Ngsi.Stderr = buf

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTokenProxyHTTPS(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "0.0.0.0", "--port", "1028", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose"})

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTokenProxyErrorKey(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "0.0.0.0", "--port", "1028", "--https"})

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestTokenProxyErrorCert(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "0.0.0.0", "--port", "1028", "--https", "--key", "test.key"})

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestTokenProxyErrorIdmHost(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--idmHost", ":", "--host", "0.0.0.0", "--port", "1028", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "parse \":\": missing protocol scheme", ngsiErr.Message)
	}
}

func TestTokenProxyErrorHTTPS(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion", "--port", "1028", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestTokenProxyErrorHTTP(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "0.0.0.0", "--port", "1028"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	err := tokenProxyServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestTokenProxyRootHandler(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyRootHandler{ngsi: c.Ngsi}

	req := httptest.NewRequest(http.MethodGet, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHealthHandler(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHealthHandler{
		ngsi: c.Ngsi,
		config: &tokenProxyConfig{
			verbose: true,
			idmHost: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHealthHandlerError(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHealthHandler{
		ngsi: c.Ngsi,
		config: &tokenProxyConfig{
			verbose: true,
			idmHost: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderErrorMethod(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "http://tokenProxy/", nil)
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderPostToken(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`grant_type=password&username=admin@test.com&password=1234`)
	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: mock,
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderPostRevoke(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`token=9721b640cafb39882cf9a71d2249760134c0073d&token_type_hint=refresh_token`)
	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: mock,
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"token":"9721b640cafb39882cf9a71d2249760134c0073d"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestTokenProxyHanderErrorToken(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`grant_type=password&username=admin@test.com&password=1234`)
	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: mock,
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokeProxyRequestToken006 parameter error\"}", got.Body.String())
}

func TestTokenProxyHanderErrorRevoke(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`token=9721b640cafb39882cf9a71d2249760134c0073d&token_type_hint=refresh_token`)
	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: mock,
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokenProxyRevokeToken006 parameter error\"}", got.Body.String())
}

func TestTokenProxyHanderErrorHTTP(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte(`grant_type=password&username=admin@test.com&password=1234`)
	reqRes.Err = errors.New("http error")

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: mock,
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")
	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusBadRequest

	assert.Equal(t, expected, got.Code)
	assert.Equal(t, "{\"error\":\"tokenProxyHandler004 http error\"}", got.Body.String())
}

func TestTokenProxyResposeError(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex:   &sync.Mutex{},
			failure: 1,
		},
	}

	got := httptest.NewRecorder()

	tokenProxyResposeError(h, got, http.StatusBadRequest, errors.New("test"))

	assert.Equal(t, int64(2), h.stat.failure)
	assert.Equal(t, "{\"error\":\"test\"}", got.Body.String())
}

func TestTokenProxyRequestTokenJSON(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   false,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234", string(actual))
	}
}

func TestTokenProxyRequestTokenJSONVerbose(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234", string(actual))
	}
}

func TestTokenProxyRequestTokenRefreshToken(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=refresh_token&refresh_token=2981fed8a6810c8a6131eb445f029dcb14a4eff3", string(actual))
	}
}

func TestTokenProxyRequestTokenURLEncorded(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`username=admin@test.com&password=1234&scope=openid`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	actual, err := tokenProxyRequestToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234&scope=openid", string(actual))
	}
}

func TestTokenProxyRequestTokenErrorContentType(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`username=admin@test.com&password=1234`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)

	_, err := tokenProxyRequestToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing Content-Type", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorUnknownContentType(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`username=admin@test.com&password=1234`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "unknown")

	_, err := tokenProxyRequestToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Content-Type error", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	_, err := tokenProxyRequestToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorURLEncoded(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", nil)
	req.Body = nil
	req.Form = nil
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRequestToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing form body", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorUnknownParam(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`user=admin@test.com`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRequestToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: user", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenErrorParam(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(``)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRequestToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "parameter error", ngsiErr.Message)
	}
}
func TestTokenProxyRevokeTokenJSON(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRevokeToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "token=2981fed8a6810c8a6131eb445f029dcb14a4eff3&token_type_hint=refresh_token", string(actual))
	}
}

func TestTokenProxyRevokeTokenJSONVerbose(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRevokeToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "token=2981fed8a6810c8a6131eb445f029dcb14a4eff3&token_type_hint=refresh_token", string(actual))
	}
}

func TestTokenProxyRevokeTokenURLEncorded(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`token=2981fed8a6810c8a6131eb445f029dcb14a4eff3`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	actual, err := tokenProxyRevokeToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "token=2981fed8a6810c8a6131eb445f029dcb14a4eff3&token_type_hint=refresh_token", string(actual))
	}
}

func TestTokenProxyRevokeTokenErrorContentType(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`token=2981fed8a6810c8a6131eb445f029dcb14a4eff3`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)

	_, err := tokenProxyRevokeToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing Content-Type", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorUnknownContentType(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`token=2981fed8a6810c8a6131eb445f029dcb14a4eff3`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "unknown")

	_, err := tokenProxyRevokeToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Content-Type error", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"token":"2981fed8a6810c8a6131eb445f029dcb14a4eff3"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/json")

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	_, err := tokenProxyRevokeToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorURLEncoded(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", nil)
	req.Body = nil
	req.Form = nil
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRevokeToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing form body", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorUnkownParam(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`user=admin@test.com`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRevokeToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: user", ngsiErr.Message)
	}
}

func TestTokenProxyRevokeTokenErrorParam(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`token_type_hint=password`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/revoke", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err := tokenProxyRevokeToken(h, req)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "parameter error", ngsiErr.Message)
	}
}

func TestTokenProxyRequestTokenURLEncoded(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHandler{
		ngsi: c.Ngsi,
		http: helper.NewMockHTTP(),
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	reqBody := bytes.NewBufferString(`{"username":"admin@test.com","password":"1234"}`)
	req := httptest.NewRequest(http.MethodPost, "http://tokenProxy/token", reqBody)
	req.Header.Set("Content-Type", "application/json")

	actual, err := tokenProxyRequestToken(h, req)

	if assert.NoError(t, err) {
		assert.Equal(t, "grant_type=password&username=admin@test.com&password=1234", string(actual))
	}
}

func TestTokenProxyGetStat(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHealthHandler{
		ngsi: c.Ngsi,
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	b := tokenProxyGetStat(h)

	stat := &tokenProxyStatInfo{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "OK")
}

func TestTokenProxyGetStatError(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "server", "--host", "orion"})

	h := &tokenProxyHealthHandler{
		ngsi: c.Ngsi,
		config: &tokenProxyConfig{
			verbose:   true,
			idmHost:   helper.UrlParse("http://keyrock:3000"),
			RevokeURL: helper.UrlParse("http://keyrock:3000"),
		},
		stat: &tokenProxyStat{
			mutex: &sync.Mutex{},
		},
	}

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	b := tokenProxyGetStat(h)

	stat := &tokenProxyStatInfo{}
	_ = json.Unmarshal(b, &stat)

	assert.Equal(t, stat.Health, "NG")
}
func TestTokenProxyHealthCmd(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "health", "--host", "tokenproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"tokenproxy","version":"0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)","health":"OK","idm":"http://keyrock:3000/oauth2/token","clientId":"a1a6048b-df1d-4d4f-9a08-5cf836041d14","clientSecret":"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9","verbose":true,"uptime":"0 d, 1 h, 55 m, 39 s","timesent":3,"success":1,"revoke":1,"failure":1}`)

	helper.SetClientHTTP(c, reqRes)

	err := tokenProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"ngsi-go\":\"tokenproxy\",\"version\":\"0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)\",\"health\":\"OK\",\"idm\":\"http://keyrock:3000/oauth2/token\",\"clientId\":\"a1a6048b-df1d-4d4f-9a08-5cf836041d14\",\"clientSecret\":\"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9\",\"verbose\":true,\"uptime\":\"0 d, 1 h, 55 m, 39 s\",\"timesent\":3,\"success\":1,\"revoke\":1,\"failure\":1}"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenProxyHealthCmdPretty(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "health", "--host", "tokenproxy", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"tokenproxy","version":"0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)","health":"OK","idm":"http://keyrock:3000/oauth2/token","clientId":"a1a6048b-df1d-4d4f-9a08-5cf836041d14","clientSecret":"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9","verbose":true,"uptime":"0 d, 1 h, 55 m, 39 s","timesent":3,"success":1,"revoke":1,"failure":1}`)

	helper.SetClientHTTP(c, reqRes)

	err := tokenProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"ngsi-go\": \"tokenproxy\",\n  \"version\": \"0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)\",\n  \"health\": \"OK\",\n  \"idm\": \"http://keyrock:3000/oauth2/token\",\n  \"clientId\": \"a1a6048b-df1d-4d4f-9a08-5cf836041d14\",\n  \"clientSecret\": \"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9\",\n  \"verbose\": true,\n  \"uptime\": \"0 d, 1 h, 55 m, 39 s\",\n  \"timesent\": 3,\n  \"success\": 1,\n  \"revoke\": 1,\n  \"failure\": 1\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenProxyHealthCmdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "health", "--host", "tokenproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := tokenProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestTokenProxyHealthCmdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "health", "--host", "tokenproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := tokenProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestTokenProxyHealthCmdIotaErrorPretty(t *testing.T) {
	c := setupTest([]string{"tokenproxy", "health", "--host", "tokenproxy", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/health"
	reqRes.ResBody = []byte(`{"ngsi-go":"tokenproxy","version":"0.10.0 (git_hash:8385af6dff05e842ef3786a231a4bdfe0880b4bf)","health":"OK","idm":"http://keyrock:3000/oauth2/token","clientId":"a1a6048b-df1d-4d4f-9a08-5cf836041d14","clientSecret":"e4cc0147-e38f-4211-b8ad-8ae5e6a107f9","verbose":true,"uptime":"0 d, 1 h, 55 m, 39 s","timesent":3,"success":1,"revoke":1,"failure":1}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := tokenProxyHealthCmd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
