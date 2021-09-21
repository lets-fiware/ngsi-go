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

package helper

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

// MockHTTPReqRes is ...
type MockHTTP struct {
	index  int
	ReqRes []MockHTTPReqRes
}

type MockHTTPReqRes struct {
	Res        http.Response
	ResBody    []byte
	ResHeader  http.Header
	Err        error
	StatusCode int
	ReqData    []byte
	Path       string
	RawQuery   *string
}

func NewMockHTTP() *MockHTTP {
	m := MockHTTP{}
	return &m
}

func NewHttpHeader(key, value string) http.Header {
	header := make(http.Header)
	header.Set(key, value)
	return header
}

func SetClientHTTP(c *ngsicli.Context, reqRes ...MockHTTPReqRes) {
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes...)
	c.Client.HTTP = mock
}

func AddReqRes(ngsi *ngsilib.NGSI, r MockHTTPReqRes) {
	h, _ := ngsi.HTTP.(*MockHTTP)
	h.ReqRes = append(h.ReqRes, r)
}

// Request is ...
func (h *MockHTTP) Request(method string, url *url.URL, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	const funcName = "Request"

	if len(h.ReqRes) == 0 {
		return nil, nil, ngsierr.New(funcName, 1, "ReqRes length is 0", nil)
	}
	r := h.ReqRes[h.index]
	h.index++

	if r.Err != nil {
		return nil, nil, r.Err
	}
	var data []byte
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		switch body := body.(type) {
		case []byte:
			data = body
		case string:
			data = []byte(body)
		default:
			return nil, nil, ngsierr.New(funcName, 2, "Unsupported type", nil)
		}
	}
	if data != nil && r.ReqData != nil {
		if !reflect.DeepEqual(r.ReqData, data) {
			fmt.Printf("r.ReqData: %s\n", string(r.ReqData))
			fmt.Printf("Data:      %s\n", string(data))
			return nil, nil, ngsierr.New(funcName, 3, "body data error", nil)
		}
	}
	if r.Path != "" && r.Path != url.Path {
		return nil, nil, ngsierr.New(funcName, 4, "url error", nil)
	}
	if r.RawQuery != nil {
		if *r.RawQuery != url.RawQuery {
			return nil, nil, ngsierr.New(funcName, 5, "raw query error: "+url.RawQuery, nil)
		}
	}
	if r.ResHeader != nil {
		r.Res.Header = r.ResHeader
	}
	return &r.Res, r.ResBody, r.Err
}
