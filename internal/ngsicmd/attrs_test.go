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

func TestAttrsReadV2(t *testing.T) {
	c := setupTest([]string{"get", "attrs", "--host", "orion", "--id", "airqualityobserved1", "--attrs", "CO"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"CO\":{\"type\":\"Number\",\"value\":400.463869544,\"metadata\":{}}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttrsReadV2Pretty(t *testing.T) {
	c := setupTest([]string{"get", "attrs", "--host", "orion", "--id", "airqualityobserved1", "--attrs", "CO", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"CO\": {\n    \"type\": \"Number\",\n    \"value\": 400.463869544,\n    \"metadata\": {}\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttrsReadV2SafeString(t *testing.T) {
	c := setupTest([]string{"get", "attrs", "--host", "orion", "--id", "airqualityobserved1", "--attrs", "CO", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"CO\":{\"type\":\"Number\",\"value\":400.463869544,\"metadata\":{}}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttrsReadErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "attrs", "--host", "orion", "--id", "airqualityobserved1", "--attrs", "CO"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestAttrsReadErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"get", "attrs", "--host", "orion", "--id", "airqualityobserved1", "--attrs", "CO"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestAttrsReadV2ErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "attrs", "--host", "orion", "--id", "airqualityobserved1", "--attrs", "CO", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value"400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		expected := "invalid character '4' after object key (30) Number\",\"value\"400.463869544,\""
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestAttrsReadV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "attrs", "--host", "orion", "--id", "airqualityobserved1", "--attrs", "CO", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := attrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAttrsAppendV2(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsAppend(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsAppendV2SafeString(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsAppend(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsAppendLD(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}", "--context", "[\"http://context\"]"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"specialOffer":{"value":true}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsAppend(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsAppendErrorData(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	err := attrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestAttrsAppendLDErrorContext(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}", "--context", "[\"http://context\""})

	err := attrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestAttrsAppendV2SafeStringError(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value: true}", "--safeString", "on"})

	err := attrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestAttrsAppendErrorHTTP(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestAttrsAppendErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"append", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsAppend(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestAttrsUpdateV2(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsUpdateV2SafeString(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsUpdateLD(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}", "--context", "[\"http://context\"]"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"specialOffer":{"value":true}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsUpdateErrorData(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestAttrsUpdateLDErrorContext(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}", "--context", "[\"http://context\""})

	err := attrsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestAttrsUpdateV2SafeStringError(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}", "--safeString", "on"})

	err := attrsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error: :{\"value\":true}", ngsiErr.Message)
	}
}

func TestAttrsUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attr"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestAttrsUpdateErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestAttrsReplaceV2(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsReplace(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsReplaceV2SafeString(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	helper.SetClientHTTP(c, reqRes)

	err := attrsReplace(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrsReplaceErrorData(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010"})

	err := attrsReplace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestAttrsReplaceV2SafeStringError(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}", "--safeString", "on"})

	err := attrsReplace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error: :{\"value\":true}", ngsiErr.Message)
	}
}

func TestAttrsReplaceErrorHTTP(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attr"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsReplace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestAttrsReplaceErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"update", "attrs", "--host", "orion", "--id", "urn:ngsi-ld:Product:010", "--data", "{\"specialOffer\":{\"value\": true}}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := attrsReplace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}
