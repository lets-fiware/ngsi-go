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

package ngsicmd

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestTroeList(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--type", "Sensor"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:sensor100\",\"type\":\"Sensor\",\"temperature\":[{\"type\":\"Property\",\"value\":25,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-03-01T00:00:01Z\"},{\"type\":\"Property\",\"value\":21,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-02-01T00:00:01Z\"}],\"@context\":[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestTroeListPretty(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--type", "Sensor", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"urn:ngsi-ld:sensor100\",\n  \"type\": \"Sensor\",\n  \"temperature\": [\n    {\n      \"type\": \"Property\",\n      \"value\": 25,\n      \"instanceId\": \"REGEX(.*)\",\n      \"observedAt\": \"2017-03-01T00:00:01Z\"\n    },\n    {\n      \"type\": \"Property\",\n      \"value\": 21,\n      \"instanceId\": \"REGEX(.*)\",\n      \"observedAt\": \"2017-02-01T00:00:01Z\"\n    }\n  ],\n  \"@context\": [\n    \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTroeListErrorTemporalQuery(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--fromDate", "1"})

	err := troeList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}

func TestTroeListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--type", "Sensor"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := troeList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTroeListErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--type", "Sensor"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTroeListErrorPretty(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--type", "Sensor", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := troeList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTroeCreate(t *testing.T) {
	data := `{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`
	c := setupTest([]string{"create", "tentity", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeCreateErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"create", "tentity", "--host", "orion-ld", "--data", "@"})

	err := troeCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestTroeCreateErrorSafeString(t *testing.T) {
	data := `{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}`
	c := setupTest([]string{"create", "tentity", "--host", "orion-ld", "--data", data, "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error: -01T00:00:01Z\"}", ngsiErr.Message)
	}
}

func TestTroeCreateErrorContext(t *testing.T) {
	data := `{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`
	c := setupTest([]string{"create", "tentity", "--host", "orion-ld", "--data", data, "--context", "ctx"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ctx not found", ngsiErr.Message)
	}
}

func TestTroeCreateErrorHTTP(t *testing.T) {
	data := `{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`
	c := setupTest([]string{"create", "tentity", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := troeCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTroeCreateErrorHTTPStatus(t *testing.T) {
	data := `{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`
	c := setupTest([]string{"create", "tentity", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTroeRead(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:sensor100\",\"type\":\"Sensor\",\"temperature\":{\"type\":\"Property\",\"value\":20,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-01-01T00:00:01Z\"},\"@context\":[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTroeReadAcceptJSON(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--acceptJson"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:sensor100\",\"type\":\"Sensor\",\"temperature\":{\"type\":\"Property\",\"value\":20,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-01-01T00:00:01Z\"},\"@context\":[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTroeReadPretty(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"urn:ngsi-ld:sensor100\",\n  \"type\": \"Sensor\",\n  \"temperature\": {\n    \"type\": \"Property\",\n    \"value\": 20,\n    \"instanceId\": \"REGEX(.*)\",\n    \"observedAt\": \"2017-01-01T00:00:01Z\"\n  },\n  \"@context\": [\n    \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTroeReadErrorTemporalQuery(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--fromDate", "1"})

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}

func TestTroeReadErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTroeReadErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error:  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTroeReadErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]`)

	helper.SetClientHTTP(c, reqRes)

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error: ontext.jsonld\"]", ngsiErr.Message)
	}
}

func TestTroeReadErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := troeRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTroeDelete(t *testing.T) {
	c := setupTest([]string{"delete", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"

	helper.SetClientHTTP(c, reqRes)

	err := troeDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := troeDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTroeDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"delete", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTroeAttrsAppend(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"append", "tattrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrsAppend(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeAttrsAppendErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"append", "tattrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--data", "@"})

	err := troeAttrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestTroeAttrsAppendErrorSafeString(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}`
	c := setupTest([]string{"append", "tattrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--data", data, "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error: -01T00:00:01Z\"}", ngsiErr.Message)
	}
}

func TestTroeAttrsAppendErrorContext(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"append", "tattrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--data", data, "--context", "ctx"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ctx not found", ngsiErr.Message)
	}
}

func TestTroeAttrsAppendErrorHTTP(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"append", "tattrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTroeAttrsAppendErrorHTTPStatus(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"append", "tattrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTroeAttrDelete(t *testing.T) {
	c := setupTest([]string{"delete", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteDeleteAll(t *testing.T) {
	c := setupTest([]string{"delete", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--deleteAll"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteDatasetId(t *testing.T) {
	c := setupTest([]string{"delete", "tattr", "--host", "orion-ld", "--datasetId", "datasetid001", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteInstanceID(t *testing.T) {
	c := setupTest([]string{"delete", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0"

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteErrorInstanceId(t *testing.T) {
	c := setupTest([]string{"delete", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "attr", "--datasetId", "dataset001", "--instanceId", "instance001"})

	err := troeAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specify --deleteALl and/or --datasetId with --instanceId", ngsiErr.Message)
	}
}

func TestTroeAttrDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTroeAttrDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"delete", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTroeAttrUpdate(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"update", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "instance001", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTroeAttrsUpdateErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"update", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "instance001", "--data", "@"})

	err := troeAttrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestTroeAttrsUpdateErrorSafeString(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}`
	c := setupTest([]string{"update", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "instance001", "--data", data, "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error: -01T00:00:01Z\"}", ngsiErr.Message)
	}
}

func TestTroeAttrsUpdateErrorContext(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"update", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "instance001", "--data", data, "--context", "ctx"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ctx not found", ngsiErr.Message)
	}
}

func TestTroeAttrsUpdateErrorHTTP(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"update", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "instance001", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTroeAttrsUpdateErrorHTTPStatus(t *testing.T) {
	data := `{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`
	c := setupTest([]string{"update", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "instance001", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := troeAttrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestBuildTemporalQueryBetween(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--fromDate", "2016-09-13T00:00:00.000Z", "--toDate", "2017-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "between", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("timeAt"))
	assert.Equal(t, "2017-09-13T00:00:00.000Z", v.Get("endTimeAt"))
}

func TestBuildTemporalQueryAfter(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--fromDate", "2016-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "after", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("timeAt"))
}

func TestBuildTemporalQueryBefore(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--toDate", "2016-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "before", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("timeAt"))
}

func TestBuildTemporalQueryBetweenETSI10(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--etsi10", "--fromDate", "2016-09-13T00:00:00.000Z", "--toDate", "2017-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "between", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("time"))
	assert.Equal(t, "2017-09-13T00:00:00.000Z", v.Get("endTime"))
}

func TestBuildTemporalQueryError(t *testing.T) {
	c := setupTest([]string{"list", "tentities", "--host", "orion-ld", "--etsi10", "--fromDate", "1"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}
