/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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

func TestEntityCreateV2(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"create", "entity", "--host", "orion", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityCreateV2SafeString(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"create", "entity", "--host", "orion", "--data", data, "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityCreateLd(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityCreateLdContext(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--data", data, "--context", "[\"http://context\"]"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ReqData = []byte(`{"@context":["http://context"],"id":"urn:ngsi-ld:Product:010","name":{"type":"Text","value":"Lemonade"},"price":{"type":"Integer","value":99},"size":{"type":"Text","value":"S"},"type":"Product"}`)

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityCreateErrorLdKeyValues(t *testing.T) {
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--keyValues", "--data", "{}"})

	err := entityCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "--keyValues only available on NGSIv2", ngsiErr.Message)
	}
}

func TestEntityCreateErrorLdUpsert(t *testing.T) {
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--upsert", "--data", "@"})

	err := entityCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "--upsert only available on NGSIv2", ngsiErr.Message)
	}
}
func TestEntityCreateErrorReadAll(t *testing.T) {
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--data", "@"})

	err := entityCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestEntityCreateErrorSafeString(t *testing.T) {
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--data", "{", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error: {", ngsiErr.Message)
	}
}

func TestEntityCreateLdErrorContext(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--data", data, "--context", "[\"http://context\""})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ReqData = []byte(`{"@context":["http://context"],"id":"urn:ngsi-ld:Product:010","name":{"type":"Text","value":"Lemonade"},"price":{"type":"Integer","value":99},"size":{"type":"Text","value":"S"},"type":"Product"}`)

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestEntityCreateErrorHTTP(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestEntityCreateErrorHTTPStatus(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"create", "entity", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	}
}

func TestEntityReadV2(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntityReadV2Pretty(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"urn:ngsi-ld:Product:010\",\n  \"type\": \"Product\",\n  \"name\": {\n    \"type\": \"Text\",\n    \"value\": \"Lemonade\"\n  },\n  \"size\": {\n    \"type\": \"Text\",\n    \"value\": \"S\"\n  },\n  \"price\": {\n    \"type\": \"Integer\",\n    \"value\": 99\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntityReadV2SafeString(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name%25":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name%\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntityReadLd(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntityReadLdJSON(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010", "--acceptJson"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntityReadLdGeoJSON(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010", "--acceptGeoJson"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}},"@context":"http://atcontext:8000/ngsi-context.jsonld","geometry":{"type":"Point","coordinates":[139.76,35.68]}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}},\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"geometry\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntityReadErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestEntityReadErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error:  error", ngsiErr.Message)
	}
}

func TestEntityReadV2ErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type:"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'P' after object key (39) ct:010\",\"type:\"Product\",\"name\"", ngsiErr.Message)
	}
}

func TestEntityReadV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := entityRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntityUpsertV2(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"upsert", "entity", "--host", "orion", "--keyValues", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityUpsert(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityUpsertV2SafeString(t *testing.T) {
	data := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"
	c := setupTest([]string{"upsert", "entity", "--host", "orion", "--data", data, "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityUpsert(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityUpsertErrorReadAll(t *testing.T) {
	c := setupTest([]string{"upsert", "entity", "--host", "orion", "--data", "@"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestEntityUpsertErrorSafeString(t *testing.T) {
	c := setupTest([]string{"upsert", "entity", "--host", "orion", "--data", "{,}", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character ',' (1) {,}", ngsiErr.Message)
	}
}

func TestEntityUpsertErrorHTTP(t *testing.T) {
	c := setupTest([]string{"upsert", "entity", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := entityUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestEntityUpsertErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"upsert", "entity", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entityUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	}
}

func TestEntityDeleteV2(t *testing.T) {
	c := setupTest([]string{"delete", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityDeleteLd(t *testing.T) {
	c := setupTest([]string{"delete", "entity", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntityDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := entityDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestEntityDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"delete", "entity", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"

	helper.SetClientHTTP(c, reqRes)

	err := entityDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	}
}
