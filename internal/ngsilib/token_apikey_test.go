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
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestRequestTokenApikey(t *testing.T) {
	ngsi := testNgsiLibInit()

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CApikey, HeaderName: "apikey", HeaderValue: "1234"}}
	idm := &idmApikey{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	assert.NoError(t, err)
}

func TestRevokeTokenApikey(t *testing.T) {
	ngsi := testNgsiLibInit()

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CApikey, HeaderName: "apikey", HeaderValue: "1234"}}
	idm := &idmApikey{}
	tokenInfo := &TokenInfo{}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	assert.NoError(t, err)
}

func TestGetAuthHeaderApikey(t *testing.T) {
	idm := &idmApikey{}

	key, value := idm.getAuthHeader("b7308719683033900d37384e723c1660")

	assert.Equal(t, "", key)
	assert.Equal(t, "", value)
}

func TestGetTokenInfoApikey(t *testing.T) {
	idm := &idmApikey{}
	tokenInfo := &TokenInfo{}

	_, err := idm.getTokenInfo(tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "no information available", ngsiErr.Message)
	}
}

func TestCheckIdmParamsApikey(t *testing.T) {
	idm := &idmApikey{}
	idmParams := &IdmParams{HeaderName: "apikey", HeaderValue: "1234"}

	err := idm.checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsApikeyEnv(t *testing.T) {
	idm := &idmApikey{}
	idmParams := &IdmParams{HeaderName: "apikey", HeaderEnvValue: "TOKEN"}

	err := idm.checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsApikeyError(t *testing.T) {
	idm := &idmApikey{}
	idmParams := &IdmParams{}

	err := idm.checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "headerName and either headerValue or headerEnvValue", ngsiErr.Message)
	}
}

func TestGetApikeyHeader(t *testing.T) {
	client := &Client{Server: &Server{HeaderName: "apikey", HeaderValue: "1234"}}

	key, value := GetApikeyHeader(client)

	assert.Equal(t, "apikey", key)
	assert.Equal(t, "1234", value)
}

func TestGetApikeyHeaderEnv(t *testing.T) {
	client := &Client{Server: &Server{HeaderName: "apikey", HeaderEnvValue: "NGSO_UNKNOWN_1234"}}

	key, value := GetApikeyHeader(client)

	assert.Equal(t, "apikey", key)
	assert.Equal(t, "", value)
}

func TestGetApikeyHeaderBlank(t *testing.T) {
	client := &Client{Server: &Server{HeaderName: "apikey", HeaderValue: ""}}

	key, value := GetApikeyHeader(client)

	assert.Equal(t, "apikey", key)
	assert.Equal(t, "", value)
}
