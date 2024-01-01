/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

func TestAPIsV2(t *testing.T) {
	c := setupTest([]string{"apis", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`)
	reqRes.Path = "/v2/"
	helper.SetClientHTTP(c, reqRes)

	err := apis(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"entities_url":"/v2/entities","types_url":"/v2/types","subscriptions_url":"/v2/subscriptions","registrations_url":"/v2/registrations"}`
		assert.Equal(t, expected, actual)
	}
}

func TestAPIsQuantumleap(t *testing.T) {
	c := setupTest([]string{"apis", "--host", "ql"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{\n"attributes_url": "/v2/attrs",\n"entities_url": "/v2/entities",\n"notify_url": "/v2/notify",\n"subscriptions_url": "/v2/subscriptions",\n"types_url": "/v2/types"\n}\n`)
	reqRes.Path = "/v2"
	helper.SetClientHTTP(c, reqRes)

	err := apis(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{\n"attributes_url": "/v2/attrs",\n"entities_url": "/v2/entities",\n"notify_url": "/v2/notify",\n"subscriptions_url": "/v2/subscriptions",\n"types_url": "/v2/types"\n}\n`
		assert.Equal(t, expected, actual)
	}
}

func TestAPIsQuantumleapPretty(t *testing.T) {
	c := setupTest([]string{"apis", "--host", "ql", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"attributes_url": "/v2/attrs","entities_url": "/v2/entities","notify_url": "/v2/notify","subscriptions_url": "/v2/subscriptions","types_url": "/v2/types"}`)
	reqRes.Path = "/v2"
	helper.SetClientHTTP(c, reqRes)

	err := apis(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"attributes_url\": \"/v2/attrs\",\n  \"entities_url\": \"/v2/entities\",\n  \"notify_url\": \"/v2/notify\",\n  \"subscriptions_url\": \"/v2/subscriptions\",\n  \"types_url\": \"/v2/types\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAPIsErrorHTTP(t *testing.T) {
	c := setupTest([]string{"apis", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/version"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := apis(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAPIsErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"apis", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/"
	helper.SetClientHTTP(c, reqRes)

	err := apis(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAPIsQuantumleapErrorPretty(t *testing.T) {
	c := setupTest([]string{"apis", "--host", "ql", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"attributes_url": "/v2/attrs","entities_url": "/v2/entities","notify_url": "/v2/notify","subscriptions_url": "/v2/subscriptions","types_url": "/v2/types"}`)
	reqRes.Path = "/v2"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := apis(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
