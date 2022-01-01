/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestHealthCheckV2(t *testing.T) {
	c := setupTest([]string{"health", "--host", "ql"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{\n"status": "pass"\n}\n`)
	reqRes.Path = "/health"

	helper.SetClientHTTP(c, reqRes)

	err := healthCheck(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{\n"status": "pass"\n}\n`
		assert.Equal(t, expected, actual)
	}
}

func TestHealthCheckScorpio(t *testing.T) {
	c := setupTest([]string{"health", "--host", "scorpio"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{ "Status of Registrymanager": "Up and running", "Status of Entitymanager": "Up and running", "Status of Subscriptionmanager": "Not running", "Status of Storagemanager": "Up and running", "Status of Querymanager": "Up and running", "Status of Historymanager": "Up and running"}`)
	reqRes.Path = "/scorpio/v1/info/health"

	helper.SetClientHTTP(c, reqRes)

	err := healthCheck(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{ \"Status of Registrymanager\": \"Up and running\", \"Status of Entitymanager\": \"Up and running\", \"Status of Subscriptionmanager\": \"Not running\", \"Status of Storagemanager\": \"Up and running\", \"Status of Querymanager\": \"Up and running\", \"Status of Historymanager\": \"Up and running\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestHealthCheckErrorBrokerType(t *testing.T) {
	c := setupTest([]string{"health", "--host", "orion-ld"})

	err := healthCheck(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "brokerType error", ngsiErr.Message)
	}
}

func TestHealthCheckErrorHTTP(t *testing.T) {
	c := setupTest([]string{"health", "--host", "ql"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := healthCheck(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestHealthCheckErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"health", "--host", "ql"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/health"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := healthCheck(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}
