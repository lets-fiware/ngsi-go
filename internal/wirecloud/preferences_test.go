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

package wirecloud

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestWcPreferences(t *testing.T) {
	c := setupTest([]string{"preferences", "get", "--host", "wirecloud"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"language": {"inherit": false, "value": "default"}}`)
	reqRes.Path = "/api/preferences/platform"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudPreferencesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"language": {"inherit": false, "value": "default"}}`
		assert.Equal(t, expected, actual)
	}
}

func TestWcPreferencesPretty(t *testing.T) {
	c := setupTest([]string{"preferences", "get", "--host", "wirecloud", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"language": {"inherit": false, "value": "default"}}`)
	reqRes.Path = "/api/preferences/platform"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudPreferencesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"language\": {\n    \"inherit\": false,\n    \"value\": \"default\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcPreferencesErrorHTTP(t *testing.T) {
	c := setupTest([]string{"preferences", "get", "--host", "wirecloud"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"language": {"inherit": false, "value": "default"}}`)
	reqRes.Path = "/api/preferences/platform"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudPreferencesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestWcPreferencesErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"preferences", "get", "--host", "wirecloud"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`error`)
	reqRes.Path = "/api/preferences/platform"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudPreferencesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestWcPreferencesErrorPretty(t *testing.T) {
	c := setupTest([]string{"preferences", "get", "--host", "wirecloud", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"language": {"inherit": false, "value": "default"}}`)
	reqRes.Path = "/api/preferences/platform"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := wireCloudPreferencesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
