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

package convenience

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestReceiver(t *testing.T) {
	c := setupTest([]string{"receiver", "--verbose", "--port", "8000", "--url", "/"})

	err := receiver(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestReceiverHTTPS(t *testing.T) {
	c := setupTest([]string{"receiver", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose", "--port", "8000", "--url", "/test"})

	err := receiver(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestReceiverErrorKey(t *testing.T) {
	c := setupTest([]string{"receiver", "--https", "--port", "8000", "--url", "/"})

	err := receiver(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestReceiverErrorCert(t *testing.T) {
	c := setupTest([]string{"receiver", "--https", "--key", "test.key", "--port", "8000", "--url", "/"})

	err := receiver(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestReceiverErrorHTTPS(t *testing.T) {
	c := setupTest([]string{"receiver", "--https", "--key", "test.key", "--cert", "test.cert", "--verbose", "--port", "8000", "--url", "/test"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	err := receiver(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServeTLS error", ngsiErr.Message)
	}
}

func TestReceiverErrorHTTP(t *testing.T) {
	c := setupTest([]string{"receiver", "--verbose", "--port", "8000", "--url", "/"})

	c.Ngsi.NetLib = &helper.MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	err := receiver(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ListenAndServe error", ngsiErr.Message)
	}
}

func TestReceiverHander(t *testing.T) {
	c := setupTest([]string{"receiver"})

	h := &receiverHandler{ngsi: c.Ngsi, pretty: false}

	reqBody := bytes.NewBufferString(`{"subscriptionId":"5fd412e8ecb082767349b975","data":[{"id":"device001","type":"device","temperature":{"type":"Number","value":25,"metadata":{}}}]}`)
	req := httptest.NewRequest(http.MethodPost, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusNoContent

	assert.Equal(t, expected, got.Code)

	expected2 := "{\"subscriptionId\":\"5fd412e8ecb082767349b975\",\"data\":[{\"id\":\"device001\",\"type\":\"device\",\"temperature\":{\"type\":\"Number\",\"value\":25,\"metadata\":{}}}]}\n"
	assert.Equal(t, expected2, helper.GetStdoutString(c))
}

func TestReceiverHanderPretty(t *testing.T) {
	c := setupTest([]string{"receiver"})

	h := &receiverHandler{ngsi: c.Ngsi, pretty: true}

	reqBody := bytes.NewBufferString(`{"subscriptionId":"5fd412e8ecb082767349b975","data":[{"id":"device001","type":"device","temperature":{"type":"Number","value":25,"metadata":{}}}]}`)
	req := httptest.NewRequest(http.MethodPost, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusNoContent
	assert.Equal(t, expected, got.Code)

	expected2 := "{\n  \"subscriptionId\": \"5fd412e8ecb082767349b975\",\n  \"data\": [\n    {\n      \"id\": \"device001\",\n      \"type\": \"device\",\n      \"temperature\": {\n        \"type\": \"Number\",\n        \"value\": 25,\n        \"metadata\": {}\n      }\n    }\n  ]\n}\n"
	assert.Equal(t, expected2, helper.GetStdoutString(c))
}

func TestReceiverHanderErrorPretty(t *testing.T) {
	c := setupTest([]string{"receiver"})

	h := &receiverHandler{ngsi: c.Ngsi, pretty: true}

	reqBody := bytes.NewBufferString(`{"subscriptionId`)
	req := httptest.NewRequest(http.MethodPost, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusNoContent
	assert.Equal(t, expected, got.Code)

	expected2 := "{\"subscriptionId\n"
	assert.Equal(t, expected2, helper.GetStdoutString(c))

}

func TestReceiverHanderErrorMethodNotAllowed(t *testing.T) {
	c := setupTest([]string{"receiver"})

	h := &receiverHandler{ngsi: c.Ngsi, pretty: false}

	reqBody := bytes.NewBufferString(`{"subscriptionId":"5fd412e8ecb082767349b975","data":[{"id":"device001","type":"device","temperature":{"type":"Number","value":25,"metadata":{}}}]}`)
	req := httptest.NewRequest(http.MethodGet, "http://receiver/", reqBody)

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed
	assert.Equal(t, expected, got.Code)

	expected2 := "Method not allowed.\n"
	assert.Equal(t, expected2, c.Ngsi.Stderr.(*bytes.Buffer).String())
}

func TestReceiverHanderHeader(t *testing.T) {
	c := setupTest([]string{"receiver"})

	h := &receiverHandler{ngsi: c.Ngsi, header: true}

	reqBody := bytes.NewBufferString(`{"subscriptionId":"5fd412e8ecb082767349b975","data":[{"id":"device001","type":"device","temperature":{"type":"Number","value":25,"metadata":{}}}]}`)
	req := httptest.NewRequest(http.MethodPost, "http://receiver/", reqBody)
	req.Header.Set("Fiware-Service", "openiot")

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusNoContent

	assert.Equal(t, expected, got.Code)

	expected2 := "Fiware-Service: openiot\n\n{\"subscriptionId\":\"5fd412e8ecb082767349b975\",\"data\":[{\"id\":\"device001\",\"type\":\"device\",\"temperature\":{\"type\":\"Number\",\"value\":25,\"metadata\":{}}}]}\n"
	assert.Equal(t, expected2, helper.GetStdoutString(c))
}
