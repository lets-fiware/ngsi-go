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

package ngsicmd

import (
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestRegistrationsListErrorV2(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsListErrorLd(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsGetErrorV2(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "57458eb60962ef754e7c0998"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsGetErrorLd(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "57458eb60962ef754e7c0998"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateErrorV2(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateErrorLd(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsDeleteErrorV2(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion", "--id", "57458eb60962ef754e7c0998"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsDeleteErrorLd(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion-ld", "--id", "57458eb60962ef754e7c0998"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateNgsiV2(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2"})

	err := registrationsTemplate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRegistrationsTemplateNgsiLd(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld"})

	err := registrationsTemplate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRegistrationsTemplateErrorNgsiTypeMissing(t *testing.T) {
	c := setupTest([]string{"template", "registration"})

	err := registrationsTemplate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing ngsiType", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateErrorNgsiTypeError(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v1"})

	err := registrationsTemplate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error: v1", ngsiErr.Message)
	}
}

func TestRegistrationsCountV2(t *testing.T) {
	c := setupTest([]string{"wc", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/registrations"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"12"}}

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "12\n"
		assert.Equal(t, expected, actual)
	}
}
func TestRegistrationsCountLD(t *testing.T) {
	c := setupTest([]string{"wc", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"21"}}

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "21\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsCountErrorHTTP(t *testing.T) {
	c := setupTest([]string{"wc", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, ngsiErr.Message, "url error")
	}
}

func TestRegistrationsCountErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"wc", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	}
}

func TestRegistrationsCountErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"wc", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	}
}
