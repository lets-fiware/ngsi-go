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
AUTHORS OR COPYRIGHT HOv2ERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package ngsicmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestOpUpdateArrayData(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", testOpUpdateArrayData})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"

	helper.SetClientHTTP(c, reqRes)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	assert.NoError(t, err)
}

func TestOpUpdateArrayDataOver100(t *testing.T) {
	testData := "["
	for i := 0; i < 105; i++ {
		testData = testData + fmt.Sprintf("{\"id\":\"urn:ngsi-ld:Product:%d\",\"type\":\"Product\"},", i)
	}
	testData = testData[:len(testData)-1] + "]"

	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", testData})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNoContent
	reqRes1.Path = "/v2/op/update"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.Path = "/v2/op/update"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	assert.NoError(t, err)
}

func TestOpUpdateLineData(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", testOpUpdateLineData})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"

	helper.SetClientHTTP(c, reqRes)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	assert.NoError(t, err)
}

func TestOpUpdateErrorReadAll(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "@"})

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestOpUpdateErrorData(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "@test"})

	c.Ngsi.FileReader = &helper.MockFileLib{IoReader: bufio.NewReader(bytes.NewReader([]byte("")))}

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "EOF", ngsiErr.Message)
	}
}

func TestOpUpdateErrorJSONDelim(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "1"})

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "data not JSON", ngsiErr.Message)
	}
}

func TestOpUpdateErrorJSONDelim2(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "{{"})

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '{' looking for beginning of object key string (2)", ngsiErr.Message)
	}
}

func TestOpUpdateErrorDecode(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "[" + testOpUpdateArrayData})

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal array into Go value of type map[string]interface {}", ngsiErr.Message)
	}
}
func TestOpUpdateArrayErrorHTTP(t *testing.T) {
	testData := "["
	for i := 0; i < 105; i++ {
		testData = testData + fmt.Sprintf("{\"id\":\"urn:ngsi-ld:Product:%d\",\"type\":\"Product\"},", i)
	}
	testData = testData[:len(testData)-1] + "]"

	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", testData})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNoContent
	reqRes1.Err = errors.New("error")
	reqRes1.Path = "/v2/op/update"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.Path = "/v2/op/update"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOpUpdateArrayErrorStatusCode(t *testing.T) {
	testData := "["
	for i := 0; i < 105; i++ {
		testData = testData + fmt.Sprintf("{\"id\":\"urn:ngsi-ld:Product:%d\",\"type\":\"Product\"},", i)
	}
	testData = testData[:len(testData)-1] + "]"

	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", testData})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.Path = "/v2/op/update"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.Path = "/v2/op/update"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestOpUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/op/update"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestOpUpdateErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/op/update"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestOpUpdateArrayDataError(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", testOpUpdateArrayDataError})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"

	helper.SetClientHTTP(c, reqRes)

	err := opUpdate(c, c.Ngsi, c.Client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "EOF", ngsiErr.Message)
	}
}

var testOpUpdateArrayData = `[
	{
	  "id":"urn:ngsi-ld:Product:001",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Brandy"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1299}
	},
	{
	  "id":"urn:ngsi-ld:Product:002",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Port"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1199}
	},
	{
	  "id":"urn:ngsi-ld:Product:003",
	  "type":"Product",
	  "offerPrice":{"type":"Integer", "value": 59}
	}
  ]`

var testOpUpdateArrayDataError = `[
	{
	  "id":"urn:ngsi-ld:Product:001",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Brandy"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1299}
	},
	{
	  "id":"urn:ngsi-ld:Product:002",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Port"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1199}
	},
	{
	  "id":"urn:ngsi-ld:Product:003",
	  "type":"Product",
	  "offerPrice":{"type":"Integer", "value": 59}
	}
  `

var testOpUpdateLineData = `{"id":"urn:ngsi-ld:Product:001","type":"Product","name":{"type":"Text","value":"Brandy"},"size":{"type":"Text","value":"M"},"price":{"type":"Integer","value":1299}}
{"id":"urn:ngsi-ld:Product:002","type":"Product","name":{"type":"Text","value":"Port"},"size":{"type":"Text","value":"M"},"price":{"type":"Integer","value":1199}}
{"id":"urn:ngsi-ld:Product:003","type":"Product","offerPrice":{"type":"Integer","value":59}}`
