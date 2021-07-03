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

	"github.com/stretchr/testify/assert"
)

func TestRequestTokenBasic(t *testing.T) {
	ngsi := testNgsiLibInit()

	broker := &Server{}
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmBasic{}

	actual, err := idm.requestToken(ngsi, client, broker, "")

	if assert.NoError(t, err) {
		assert.Equal(t, "basic", actual.Type)
		assert.Equal(t, "Zml3YXJlOjEyMzQ=", actual.Token)
	}
}

func TestRequestTokenBasicErrorUser(t *testing.T) {
	ngsi := testNgsiLibInit()

	broker := &Server{}
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmBasic{}

	_, err := idm.requestToken(ngsi, client, broker, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestGetAuthHeaderBasic(t *testing.T) {
	idm := &idmBasic{}

	key, value := idm.getAuthHeader("b7308719683033900d37384e723c1660")

	assert.Equal(t, "Authorization", key)
	assert.Equal(t, "Basic b7308719683033900d37384e723c1660", value)
}
