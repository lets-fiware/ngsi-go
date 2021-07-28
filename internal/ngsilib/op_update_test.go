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

package ngsilib

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestOpUpdate(t *testing.T) {
	testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	client := &Client{HTTP: mock, URL: &url.URL{}, Headers: map[string]string{}, Server: &Server{ServerType: "broker"}}

	entities := `[]`
	actionType := "update"
	keyValues := true
	safeString := false

	_, _, err := client.OpUpdate(entities, actionType, keyValues, safeString)

	assert.NoError(t, err)
}

func TestOpUpdateKeyValues(t *testing.T) {
	testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	client := &Client{HTTP: mock, URL: &url.URL{}, Headers: map[string]string{}, Server: &Server{ServerType: "broker"}}

	entities := `[]`
	actionType := "update"
	keyValues := true
	safeString := false

	_, _, err := client.OpUpdate(entities, actionType, keyValues, safeString)

	assert.NoError(t, err)
}

func TestOpUpdateErrorJSON(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	client := &Client{HTTP: mock, URL: &url.URL{}, Headers: map[string]string{}, Server: &Server{ServerType: "broker"}}
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}

	entities := `[]`
	actionType := "update"
	keyValues := true
	safeString := false

	_, _, err := client.OpUpdate(entities, actionType, keyValues, safeString)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshal error", ngsiErr.Message)
	}
}
