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

func TestRequestTokenKeyrock(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CKeyrock, actual.Type)
		assert.Equal(t, "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", actual.Keyrock.AccessToken)
	}
}

func TestRequestTokenKeyrockRefresh(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{RefreshToken: "refresh"}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CKeyrock, actual.Type)
		assert.Equal(t, "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", actual.Keyrock.AccessToken)
	}
}

func TestRequestTokenKeyrockErrorUser(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrockErrorHTTP(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrockErrorUnmarshal(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		actual := "json: cannot unmarshal string into Go value of type ngsilib.KeyrockToken Field: (14) \"access_token\": \"ad5252cd520c"
		assert.Equal(t, actual, ngsiErr.Message)
	}
}

func TestRequestTokenKeyrockErrorHTTPStatus(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestRequestTokenKeyrockErrorHTTPStatusUnauthorized(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusUnauthorized
	reqRes.ResBody = []byte(`Unauthorized`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  Unauthorized", ngsiErr.Message)
	}
}

func TestRevokeTokenKeyrock(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte("token=1a8346b8df2881c8b3407b0f39c80d1374204b93&token_type_hint=refresh_token")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{RefreshToken: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	assert.NoError(t, err)
}

func TestRevokeTokenKeyrockErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte("token=1a8346b8df2881c8b3407b0f39c80d1374204b93&token_type_hint=refresh_token")
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{RefreshToken: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRevokeTokenKeyrockErrorHTTPStatus(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte("token=1a8346b8df2881c8b3407b0f39c80d1374204b93&token_type_hint=refresh_token")
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{RefreshToken: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestGetAuthHeaderKeyrock(t *testing.T) {
	idm := &idmKeyrock{}

	key, value := idm.getAuthHeader("9e7067026d0aac494e8fedf66b1f585e79f52935")

	assert.Equal(t, "Authorization", key)
	assert.Equal(t, "Bearer 9e7067026d0aac494e8fedf66b1f585e79f52935", value)
}

func TestGetTokenInfoKeyrock(t *testing.T) {
	testNgsiLibInit()

	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{
		Keyrock: &KeyrockToken{
			AccessToken:  "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3",
			ExpiresIn:    3599,
			RefreshToken: "03e33a311e03317b390956729bcac2794b695670",
			Scope:        []string{"bearer"},
			TokenType:    "Bearer",
		},
	}

	actual, err := idm.getTokenInfo(tokenInfo)

	if assert.NoError(t, err) {
		expected := "{\"access_token\":\"ad5252cd520cnaddacdc5d2e63899f0cdcf946f3\",\"expires_in\":3599,\"refresh_token\":\"03e33a311e03317b390956729bcac2794b695670\",\"scope\":[\"bearer\"],\"token_type\":\"Bearer\"}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestGetTokenInfoKeyrockError(t *testing.T) {
	testNgsiLibInit()

	idm := &idmKeyrock{}
	tokenInfo := &TokenInfo{
		Keyrock: &KeyrockToken{
			AccessToken:  "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3",
			ExpiresIn:    3599,
			RefreshToken: "03e33a311e03317b390956729bcac2794b695670",
			Scope:        []string{"bearer"},
			TokenType:    "Bearer",
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

func TestCheckIdmParamsKeyrock(t *testing.T) {
	idm := &idmKeyrock{}
	idmParams := &IdmParams{
		IdmHost:      "https://keyrock/oauth2/token",
		Username:     "keyrock001@letsfiware.jp",
		Password:     "1234",
		ClientID:     "00000000-1111-2222-3333-444444444444",
		ClientSecret: "55555555-6666-7777-8888-999999999999",
	}

	err := idm.checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsKeyrockError(t *testing.T) {
	idm := &idmKeyrock{}
	idmParams := &IdmParams{
		IdmHost:  "https://keyrock/oauth2/token",
		Username: "keyrock001@letsfiware.jp",
		Password: "1234",
		ClientID: "00000000-1111-2222-3333-444444444444",
	}

	err := idm.checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "idmHost, username, password, clientID and clientSecret are needed", ngsiErr.Message)
	}
}
