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

package ngsilib

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestRequestTokenKeyrockIDM(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"3be06d20-b231-4430-8c38-ab5b12a8fad1"}}
	reqRes.ResBody = []byte(`{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: CKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CKeyrockIDM, actual.Type)
		assert.Equal(t, "3be06d20-b231-4430-8c38-ab5b12a8fad1", actual.Token)
		assert.Equal(t, "2021-02-12 22:56:03", actual.Expires.Format("2006-01-02 15:04:05"))
	}
}

func TestRequestTokenKeyrockIDMErrorUser(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: CKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrockIDMErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: CKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrockIDMErrorHTTPStatus(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: CKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrockIDMErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: CKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal string into Go value of type ngsilib.KeyrockIDMToken Field: (8) \"tokens\":{\"9e7067026d0a", ngsiErr.Message)
	}
}

func TestRevokeKeyrockIDM(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{Token: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	assert.NoError(t, err)
}

func TestRevokeKeyrockIDMErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{Token: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRevokeKeyrockIDMErrorHTTPStatus(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{Token: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestGetAuthHeaderKeyrockIDM(t *testing.T) {
	idm := &idmKeyrockIDM{}

	key, value := idm.getAuthHeader("9e7067026d0aac494e8fedf66b1f585e79f52935")

	assert.Equal(t, "X-Auth-Token", key)
	assert.Equal(t, "9e7067026d0aac494e8fedf66b1f585e79f52935", value)
}

func TestGetTokenInfoKeyrockIDM(t *testing.T) {
	testNgsiLibInit()

	idm := &idmKeyrockIDM{}
	tokenInfo := &TokenInfo{
		KeyrockIDM: &KeyrockIDMToken{},
	}
	tokenInfo.KeyrockIDM.Token.Methods = []string{"password"}
	tokenInfo.KeyrockIDM.Token.ExpiresAt = "2021-02-12T22:56:03.410Z"
	tokenInfo.KeyrockIDM.IdmAuthorizationConfig.Level = "basic"
	tokenInfo.KeyrockIDM.IdmAuthorizationConfig.Authzforce = false

	_, err := idm.getTokenInfo(tokenInfo)

	assert.NoError(t, err)
}

func TestCheckIdmParamsKeyrockIDM(t *testing.T) {
	idm := &idmKeyrockIDM{}
	idmParams := &IdmParams{
		IdmHost:  "https://idm.letsfiware.jp",
		Username: "admin@letsfiware.jp",
		Password: "1234",
	}

	err := idm.checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsKeyrockIDMError(t *testing.T) {
	idm := &idmKeyrockIDM{}
	idmParams := &IdmParams{
		IdmHost:  "https://idm.letsfiware.jp",
		Username: "admin@letsfiware.jp",
	}

	err := idm.checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "username and password are needed", ngsiErr.Message)
	}
}
