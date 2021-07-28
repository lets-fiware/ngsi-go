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

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestAttrReadV2(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "89\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttrReadV2Pretty(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":"fiware"}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"name\": \"fiware\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttrReadV2SafeString(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":"%25"}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"name\":\"%\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttrReadLD(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	helper.SetClientHTTP(c, reqRes)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "89\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttrReadErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestAttrReadErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte("error")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestAttrReadV2ErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error: {\"name\"", ngsiErr.Message)
	}
}
func TestAttrReadV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":"fiware"}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := attrRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAttrUpdateV2Int(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "89"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2Float(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "123.45"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("123.45")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2Null(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "null"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("null")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2BoolTrue(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "true"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("true")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2BoolFalse(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "false"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("false")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2StringNull(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "\"null\""})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`"null"`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2StringEmpty(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "\"\""})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`""`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2StringTrue(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "\"true\""})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`"true"`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2StringFalse(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "\"false\""})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`"false"`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2String(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`"FIWARE"`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2StringWithSpace(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "\"Open APIs\""})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`"Open APIs"`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2StringSafeString(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "<>", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`"%3C%3E"`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2StringSafeString2(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "\"\"", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`""`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2JSON(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "{\"value\":89}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("{\"value\":89}")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateV2JSONSafeString(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "{\"value\":\"<>\"}", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("{\"value\":\"%3C%3E\"}")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateLD(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "99"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("99")
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateLDJSON(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "{\"value\":99}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("{\"value\":99}")
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateLDJSONContext(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "{\"value\":89}", "--context", "[\"http://context\"]"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"value":89}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrUpdateErrorReadALl(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:001", "--attr", "price"})

	err := attrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestAttrUpdateLDJSONContextError(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion-ld", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "{\"value\":89}", "--context", "[\"http://context\""})

	err := attrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}
func TestAttrUpdateV2ErrorJSONSafeString(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "{\"value\":\"<>}", "--safeString", "on"})

	err := attrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestAttrUpdateV2ErrorLength(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "\""})

	err := attrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data length error", ngsiErr.Message)
	}
}

func TestAttrUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "89"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Err = errors.New("http error")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestAttrUpdateErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"update", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price", "--data", "89"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte("89")
	reqRes.ResBody = []byte("http error")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	helper.SetClientHTTP(c, reqRes)

	err := attrUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " http error", ngsiErr.Message)
	}
}

func TestAttrDelete(t *testing.T) {
	c := setupTest([]string{"delete", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	helper.SetClientHTTP(c, reqRes)

	err := attrDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAttrDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Err = errors.New("http error")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	helper.SetClientHTTP(c, reqRes)

	err := attrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestAttrDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"delete", "attr", "--host", "orion", "--id", "urn:ngsi-ld:Product:001", "--attr", "price"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte("error")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	helper.SetClientHTTP(c, reqRes)

	err := attrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}
