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
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTokenTokenproxy(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	broker := &Server{}
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmTokenProxy{}

	actual, err := idm.requestToken(ngsi, client, broker, "")

	if assert.NoError(t, err) {
		assert.Equal(t, CTokenproxy, actual.Type)
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual.Oauth.AccessToken)
	}
}

func TestRequestTokenTokenproxyErrorUser(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	broker := &Server{}
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmTokenProxy{}

	_, err := idm.requestToken(ngsi, client, broker, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestRequestTokenTokenproxyErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	broker := &Server{}
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmTokenProxy{}

	_, err := idm.requestToken(ngsi, client, broker, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenTokenproxyErrorHTTPStatus(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`bad request`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	broker := &Server{}
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmTokenProxy{}

	_, err := idm.requestToken(ngsi, client, broker, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestRequestTokenTokenproxyErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	broker := &Server{}
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmTokenProxy{}

	_, err := idm.requestToken(ngsi, client, broker, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		actual := "json: cannot unmarshal string into Go value of type ngsilib.OauthToken Field: (14) \"access_token\": \"ad5252cd520c"
		assert.Equal(t, actual, ngsiErr.Message)
	}
}

func TestGetAuthHeaderTokenproxy(t *testing.T) {
	idm := &idmTokenProxy{}

	key, value := idm.getAuthHeader("9e7067026d0aac494e8fedf66b1f585e79f52935")

	assert.Equal(t, "Authorization", key)
	assert.Equal(t, "Bearer 9e7067026d0aac494e8fedf66b1f585e79f52935", value)
}
