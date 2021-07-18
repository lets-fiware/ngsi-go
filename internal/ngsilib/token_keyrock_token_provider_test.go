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

func TestRequestTokenKeyrocktokenprovider(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`ad5252cd520cnaddacdc5d2e63899f0cdcf946f3`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrocktokenprovider, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockTokenProvider{}
	tokenInfo := &TokenInfo{}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CKeyrocktokenprovider, actual.Type)
		assert.Equal(t, "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", actual.TokenProvider.AccessToken)
	}
}

func TestRequestTokenKeyrocktokenproviderErrorUser(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrocktokenprovider, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockTokenProvider{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrocktokenproviderErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrocktokenprovider, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockTokenProvider{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrocktokenproviderErrorHTTPStatus(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrocktokenprovider, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockTokenProvider{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestRevokeTokenKeyrocktokenprovider(t *testing.T) {
	ngsi := testNgsiLibInit()

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockTokenProvider{}
	tokenInfo := &TokenInfo{}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	assert.NoError(t, err)
}

func TestGetAuthHeaderKeyrocktokenprovider(t *testing.T) {
	idm := &idmKeyrockTokenProvider{}

	key, value := idm.getAuthHeader("9e7067026d0aac494e8fedf66b1f585e79f52935")

	assert.Equal(t, "Authorization", key)
	assert.Equal(t, "Bearer 9e7067026d0aac494e8fedf66b1f585e79f52935", value)
}

func TestGetTokenInfoKeyrocktokenprovider(t *testing.T) {
	testNgsiLibInit()

	idm := &idmKeyrockTokenProvider{}
	tokenInfo := &TokenInfo{
		TokenProvider: &TokenProvider{
			AccessToken: "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3",
		},
	}

	actual, err := idm.getTokenInfo(tokenInfo)

	if assert.NoError(t, err) {
		expected := "{\"access_token\":\"ad5252cd520cnaddacdc5d2e63899f0cdcf946f3\"}"

		assert.Equal(t, expected, string(actual))
	}
}

func TestGetTokenInfoKeyrocktokenproviderError(t *testing.T) {
	testNgsiLibInit()

	idm := &idmKeyrockTokenProvider{}
	tokenInfo := &TokenInfo{
		TokenProvider: &TokenProvider{
			AccessToken: "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3",
		},
	}

	gNGSI.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error")}

	_, err := idm.getTokenInfo(tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestCheckIdmParamsKeyrocktokenprovider(t *testing.T) {
	idm := &idmKeyrockTokenProvider{}
	idmParams := &IdmParams{
		IdmHost:  "https://tokenprovider",
		Username: "keyrock001@letsfiware.jp",
		Password: "1234",
	}

	err := idm.checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsKeyrocktokenproviderError(t *testing.T) {
	idm := &idmKeyrockTokenProvider{}
	idmParams := &IdmParams{
		IdmHost:  "https://tokenprovider",
		Username: "keyrock001@letsfiware.jp",
	}

	err := idm.checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "idmHost, username and password are needed", ngsiErr.Message)
	}
}
