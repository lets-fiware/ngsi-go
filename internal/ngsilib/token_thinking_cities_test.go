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

func TestRequestTokenErrorThinkingCities(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResBody = []byte(`{"token":{"domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"methods":["password"],"roles":[],"expires_at":"2021-04-16T11:30:47.000000Z","catalog":[],"extras":{"password_creation_time":"2021-04-16T08:29:01Z","last_login_attempt_time":"2021-04-16T08:29:05.000000","pwd_user_in_blacklist":false,"password_expiration_time":"2022-04-16T08:29:01Z"},"user":{"password_expires_at":"2022-04-16T08:29:00.000000","domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"id":"80e292b7dae445e7af66c284162ff049","name":"usertest"},"audit_ids":["6kJ9zBFCQaKRa7aCFc6bpw"],"issued_at":"2021-04-16T08:30:47.000000Z"}}`)
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: CThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}
	idm := &idmThinkingCities{}
	tokenInfo := &TokenInfo{}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CThinkingCities, actual.Type)
		expected := "gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"
		assert.Equal(t, expected, actual.Token)
		assert.Equal(t, "2021-04-16 11:30:47", actual.Expires.Format("2006-01-02 15:04:05"))
	}
}

func TestRequestTokenErrorThinkingCitiesErrorUser(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResBody = []byte(`{"token":{"domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"methods":["password"],"roles":[],"expires_at":"2021-04-16T11:30:47.000000Z","catalog":[],"extras":{"password_creation_time":"2021-04-16T08:29:01Z","last_login_attempt_time":"2021-04-16T08:29:05.000000","pwd_user_in_blacklist":false,"password_expiration_time":"2022-04-16T08:29:01Z"},"user":{"password_expires_at":"2022-04-16T08:29:00.000000","domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"id":"80e292b7dae445e7af66c284162ff049","name":"usertest"},"audit_ids":["6kJ9zBFCQaKRa7aCFc6bpw"],"issued_at":"2021-04-16T08:30:47.000000Z"}}`)
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: CThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Tenant: "smartcity", Scope: "/madrid"}}
	idm := &idmThinkingCities{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestRequestTokenErrorThinkingCitiesErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResBody = []byte(`{"token":{"domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"methods":["password"],"roles":[],"expires_at":"2021-04-16T11:30:47.000000Z","catalog":[],"extras":{"password_creation_time":"2021-04-16T08:29:01Z","last_login_attempt_time":"2021-04-16T08:29:05.000000","pwd_user_in_blacklist":false,"password_expiration_time":"2022-04-16T08:29:01Z"},"user":{"password_expires_at":"2022-04-16T08:29:00.000000","domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"id":"80e292b7dae445e7af66c284162ff049","name":"usertest"},"audit_ids":["6kJ9zBFCQaKRa7aCFc6bpw"],"issued_at":"2021-04-16T08:30:47.000000Z"}}`)
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"}}
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: CThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}
	idm := &idmThinkingCities{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenErrorThinkingCitiesErrorStatus(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`bad request`)
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: CThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}
	idm := &idmThinkingCities{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestRequestTokenThinkingCitiesErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResBody = []byte(`{"token":{"domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"methods":["password"],"roles":[],"expires_at":"2021-04-16T11:30:47.000000Z","catalog":[],"extras":{"password_creation_time":"2021-04-16T08:29:01Z","last_login_attempt_time":"2021-04-16T08:29:05.000000","pwd_user_in_blacklist":false,"password_expiration_time":"2022-04-16T08:29:01Z"},"user":{"password_expires_at":"2022-04-16T08:29:00.000000","domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"id":"80e292b7dae445e7af66c284162ff049","name":"usertest"},"audit_ids":["6kJ9zBFCQaKRa7aCFc6bpw"],"issued_at":"2021-04-16T08:30:47.000000Z"}`)
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: CThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}
	idm := &idmThinkingCities{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestGetAuthHeaderThinkingCities(t *testing.T) {
	idm := &idmThinkingCities{}

	key, value := idm.getAuthHeader("9e7067026d0aac494e8fedf66b1f585e79f52935")

	assert.Equal(t, "X-Auth-Token", key)
	assert.Equal(t, "9e7067026d0aac494e8fedf66b1f585e79f52935", value)
}
