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

func TestRequestTokenKong(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte("grant_type=client_credentials&client_id=0000&client_secret=1111")
	reqRes.ResBody = []byte(`{"expires_in":7200,"access_token":"G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO","token_type":"bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CKong, actual.Type)
		expected := "G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO"
		assert.Equal(t, expected, actual.Kong.AccessToken)
	}
}

func TestRequestTokenKongErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte("grant_type=client_credentials&client_id=0000&client_secret=1111")
	reqRes.ResBody = []byte(`{"expires_in":7200,"access_token":"G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO","token_type":"bearer"}`)
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenKongErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte("grant_type=client_credentials&client_id=0000&client_secret=1111")
	reqRes.ResBody = []byte(`"expires_in":7200,"access_token":"G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO","token_type":"bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		actual := "json: cannot unmarshal string into Go value of type ngsilib.KongToken Field: (12) \"expires_in\":7200,\"access_t"
		assert.Equal(t, actual, ngsiErr.Message)
	}
}

func TestRequestTokenKongErrorHTTPStatus(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestRequestTokenKongErrorHTTPStatusUnauthorized(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  Unauthorized", ngsiErr.Message)
	}
}

func TestRevokeTokenKong(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/oauth2_tokens/G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{Token: "G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	assert.NoError(t, err)
}

func TestRevokeTokenKongErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/oauth2_tokens/G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO"
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{Token: "G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRevokeTokenKongErrorHTTPStatus(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/oauth2_tokens/G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO"
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKong, IdmHost: "http://kong-service,http://kong-idm", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKong{}
	tokenInfo := &TokenInfo{Token: "G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestGetAuthHeaderKong(t *testing.T) {
	idm := &idmKong{}

	key, value := idm.getAuthHeader("G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO")

	assert.Equal(t, "Authorization", key)
	assert.Equal(t, "Bearer G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO", value)
}

func TestGetKongHostService(t *testing.T) {
	idmHost := "http://kong-service,http://kong-idm"

	actual := getKongHost(idmHost, cKongService)

	expected := "http://kong-service"

	assert.Equal(t, expected, actual)
}

func TestGetKongHostIdm(t *testing.T) {
	idmHost := "http://kong-service,http://kong-idm"

	actual := getKongHost(idmHost, cKongIdm)

	expected := "http://kong-idm"

	assert.Equal(t, expected, actual)
}

func TestGetKongHostError(t *testing.T) {
	idmHost := "http://kong-service"

	actual := getKongHost(idmHost, cKongIdm)

	expected := "http://kong-error/"

	assert.Equal(t, expected, actual)
}

func TestGetTokenInfoKong(t *testing.T) {
	testNgsiLibInit()

	idm := &idmKong{}
	tokenInfo := &TokenInfo{
		Kong: &KongToken{
			ExpiresIn:   7200,
			AccessToken: "G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO",
			TokenType:   "bearer",
		},
	}

	actual, err := idm.getTokenInfo(tokenInfo)

	if assert.NoError(t, err) {
		expected := "{\"expires_in\":7200,\"access_token\":\"G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO\",\"token_type\":\"bearer\"}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestGetTokenInfoKongError(t *testing.T) {
	testNgsiLibInit()

	idm := &idmKong{}
	tokenInfo := &TokenInfo{
		Kong: &KongToken{
			ExpiresIn:   7200,
			AccessToken: "G1y60yGbFE8OKXH6VHEOO1LGP0A5qyeO",
			TokenType:   "bearer",
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

func TestCheckIdmParamsKong(t *testing.T) {
	idm := &idmKong{}
	idmParams := &IdmParams{
		IdmHost:      "https://localhost:8443/ngsi/oauth2/token,http://localhost:8001/",
		ClientID:     "orion",
		ClientSecret: "1234",
	}

	err := idm.checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsKongError(t *testing.T) {
	idm := &idmKong{}
	idmParams := &IdmParams{
		IdmHost:      "https://localhost:8443/ngsi/oauth2/token,http://localhost:8001/",
		Username:     "fiware",
		Password:     "1234",
		ClientID:     "orion",
		ClientSecret: "1234",
	}

	err := idm.checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "idmHost, clientID and clientSecret are needed", ngsiErr.Message)
	}
}
