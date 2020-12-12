/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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

package ngsicmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestReceiver(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url")
	setupFlagBool(set, "verbose,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--verbose", "--port=aaaa", "--url=/"})
	err := receiver(c)

	assert.NoError(t, err)
}

func TestReceiverHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--https", "--key=test.key", "--cert=test.cert", "--verbose", "--port=aaaa", "--url=/test"})
	err := receiver(c)

	assert.NoError(t, err)
}

func TestReceiverError(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := receiver(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestReceiverErrorKey(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--https", "--port=aaaa", "--url=/"})
	err := receiver(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestReceiverErrorCert(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,key,cert")
	setupFlagBool(set, "verbose,pretty,https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--https", "--key=a", "--port=aaaa", "--url=/"})
	err := receiver(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestReceiverHander(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	pretty := false
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	receiverGlobal = &receiverParam{ngsi: ngsi, pretty: pretty}

	reqBody := bytes.NewBufferString(`{"subscriptionId":"5fd412e8ecb082767349b975","data":[{"id":"device001","type":"device","temperature":{"type":"Number","value":25,"metadata":{}}}]}`)
	req := httptest.NewRequest(http.MethodPost, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	receiverHandler(got, req)

	expected := http.StatusNoContent

	assert.Equal(t, expected, got.Code)
}

func TestReceiverHanderPretty(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	pretty := true
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	receiverGlobal = &receiverParam{ngsi: ngsi, pretty: pretty}

	reqBody := bytes.NewBufferString(`{"subscriptionId":"5fd412e8ecb082767349b975","data":[{"id":"device001","type":"device","temperature":{"type":"Number","value":25,"metadata":{}}}]}`)
	req := httptest.NewRequest(http.MethodPost, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	receiverHandler(got, req)

	expected := http.StatusNoContent

	assert.Equal(t, expected, got.Code)
}

func TestReceiverHanderErrorPretty(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	pretty := true
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	receiverGlobal = &receiverParam{ngsi: ngsi, pretty: pretty}

	reqBody := bytes.NewBufferString(`{"subscriptionId`)
	req := httptest.NewRequest(http.MethodPost, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	receiverHandler(got, req)

	expected := http.StatusNoContent

	assert.Equal(t, expected, got.Code)
}

func TestReceiverHanderErrorMethodNotAllowed(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	pretty := false
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf
	receiverGlobal = &receiverParam{ngsi: ngsi, pretty: pretty}

	reqBody := bytes.NewBufferString(`{"subscriptionId":"5fd412e8ecb082767349b975","data":[{"id":"device001","type":"device","temperature":{"type":"Number","value":25,"metadata":{}}}]}`)
	req := httptest.NewRequest(http.MethodGet, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	receiverHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}
