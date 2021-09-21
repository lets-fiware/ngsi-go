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

package ngsicmd

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestSubscriptionsListErrorV2(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsListErrorLd(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsGetErrorV2(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "0000"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsGetErrorLd(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "0000"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateErrorV2(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateErrorLd(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateErrorV2(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "3ea2e78f675f2d199d3025ff"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateErrorLd(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsDeleteErrorV2(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion", "--id", "3ea2e78f675f2d199d3025ff"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsDeleteErrorLd(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateNgsiTypeV2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--data", "{}"})

	err := subscriptionsTemplate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateNgsiTypeLd(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "ld", "--data", "{}"})

	err := subscriptionsTemplate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"Subscription\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateErrorNgsiTypeMissing(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion"})

	err := subscriptionsTemplate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required ngsiType not found", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateErrorNgsiType(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v1", "--data", "{}"})

	err := subscriptionsTemplate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error v1", ngsiErr.Message)
	}
}

func TestSubscriptionsCountErrorHTTP(t *testing.T) {
	c := setupTest([]string{"wc", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsCountErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"wc", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestSubscriptionsCountErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"wc", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestSubscriptionsCountV2(t *testing.T) {
	c := setupTest([]string{"wc", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"12"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "12\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsCountLD(t *testing.T) {
	c := setupTest([]string{"wc", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"21"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "21\n"
		assert.Equal(t, expected, actual)
	}
}
