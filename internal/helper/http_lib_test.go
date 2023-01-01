/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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

package helper

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestHTTPNewMockHTTP(t *testing.T) {
	h := NewMockHTTP()

	assert.NotEqual(t, (*MockHTTP)(nil), h)
}

func TestNewHttpHeader(t *testing.T) {
	h := NewHttpHeader("options", "keyValues")

	assert.Equal(t, "keyValues", h.Get("options"))
}

func TestHTTPSetClientHTTP(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	SetClientHTTP(c, MockHTTPReqRes{})

	assert.NotEqual(t, (*MockHTTP)(nil), c.Client.HTTP)
}

func TestHTTPAddReqRes(t *testing.T) {
	ngsi := &ngsilib.NGSI{HTTP: &MockHTTP{}}

	AddReqRes(ngsi, MockHTTPReqRes{})

	assert.NotEqual(t, (*http.Request)(nil), ngsi.HTTP)
}

func TestHTTPRequestGet(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{Path: "/v2", RawQuery: raw}
	header := map[string]string{}

	res, body, err := c.Client.HTTP.Request(http.MethodGet, u, header, nil)

	if assert.NoError(t, err) {
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, "{\"entities_url\":\"/v2/entities\",\"types_url\":\"/v2/types\",\"subscriptions_url\":\"/v2/subscriptions\",\"registrations_url\":\"/v2/registrations\"}", string(body))
	}
}

func TestHTTPRequestPOSTByte(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{Path: "/v2", RawQuery: raw}
	header := map[string]string{}

	res, body, err := c.Client.HTTP.Request(http.MethodPost, u, header, []byte("fiware"))

	if assert.NoError(t, err) {
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, "{\"entities_url\":\"/v2/entities\",\"types_url\":\"/v2/types\",\"subscriptions_url\":\"/v2/subscriptions\",\"registrations_url\":\"/v2/registrations\"}", string(body))
	}
}

func TestHTTPRequestPOSTString(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{Path: "/v2", RawQuery: raw}
	header := map[string]string{}

	res, body, err := c.Client.HTTP.Request(http.MethodPost, u, header, "fiware")

	if assert.NoError(t, err) {
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, "{\"entities_url\":\"/v2/entities\",\"types_url\":\"/v2/types\",\"subscriptions_url\":\"/v2/subscriptions\",\"registrations_url\":\"/v2/registrations\"}", string(body))
	}
}

func TestHTTPRequestError(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{Err: errors.New("HTTP error")}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{Path: "/v2", RawQuery: raw}
	header := map[string]string{}

	_, _, err := c.Client.HTTP.Request(http.MethodGet, u, header, nil)

	if assert.Error(t, err) {
		assert.Equal(t, "HTTP error", err.Error())
	}
}

func TestHTTPRequestErrorReqRes0(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{HTTP: &MockHTTP{}}}

	u := &url.URL{Path: "/v2"}
	header := map[string]string{}

	_, _, err := c.Client.HTTP.Request(http.MethodGet, u, header, nil)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "ReqRes length is 0", ngsiErr.Message)
	}
}

func TestHTTPRequestErrorPOSTDataType(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{Path: "/v2", RawQuery: raw}
	header := map[string]string{}

	_, _, err := c.Client.HTTP.Request(http.MethodPost, u, header, 1)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Unsupported type", ngsiErr.Message)
	}
}

func TestHTTPRequestErrorPOSTData(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{ReqData: []byte("orion")}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{Path: "/v2", RawQuery: raw}
	header := map[string]string{}

	_, _, err := c.Client.HTTP.Request(http.MethodPost, u, header, "fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "body data error", ngsiErr.Message)
	}
}

func TestHTTPRequestErrorURL(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{}
	header := map[string]string{}

	_, _, err := c.Client.HTTP.Request(http.MethodPost, u, header, "fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestHTTPRequestErrorRawQuery(t *testing.T) {
	c := &ngsicli.Context{Client: &ngsilib.Client{}}

	reqRes := MockHTTPReqRes{ReqData: []byte("fiware")}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2"
	reqRes.ResHeader = make(http.Header)
	raw := "raw"
	reqRes.RawQuery = &raw

	SetClientHTTP(c, reqRes)

	u := &url.URL{Path: "/v2", RawQuery: "keyValue"}
	header := map[string]string{}

	_, _, err := c.Client.HTTP.Request(http.MethodPost, u, header, "fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "raw query error: keyValue", ngsiErr.Message)
	}
}
