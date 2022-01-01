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
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestBatchCreateLd(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion-ld", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "create")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchUpdateLd(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion-ld", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "update")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchUsertLd(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "upsert")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchDeleteLd(t *testing.T) {
	c := setupTest([]string{"delete", "entities", "--host", "orion-ld", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "delete")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchCreateV2(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "create")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchUpdateV2(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "update")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchAppendV2(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "upsert")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}
func TestBatchReplaceV2(t *testing.T) {
	c := setupTest([]string{"replace", "entities", "--host", "orion", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "replace")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchDeleteV2(t *testing.T) {
	c := setupTest([]string{"delete", "entities", "--host", "orion", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "delete")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchErrorModeV2(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "get")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: get", ngsiErr.Message)
	}
}

func TestBatchErrorModeLD(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion-ld", "--data", "@"})

	err := batch(c, c.Ngsi, c.Client, "get")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: get", ngsiErr.Message)
	}
}

func TestBatchCreate(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion-ld", "--data", testData})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResBody = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/create"
	helper.SetClientHTTP(c, reqRes)

	err := batchCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBatchCreateContext(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion-ld", "--data", testData, "--context", "[\"http://context\"]"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`[{"@context":["http://context"],"category":{"type":"Property","value":"sensor"},"id":"urn:ngsi-ld:TemperatureSensor:002","temperature":{"type":"Property","unitCode":"CEL","value":21},"type":"TemperatureSensor"},{"@context":["http://context"],"category":{"type":"Property","value":"sensor"},"id":"urn:ngsi-ld:TemperatureSensor:003","temperature":{"type":"Property","unitCode":"CEL","value":27},"type":"TemperatureSensor"}]`)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/create"
	helper.SetClientHTTP(c, reqRes)

	err := batchCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBatchCreateErrorReadAll(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion-ld", "--data", "@"})

	err := batchCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchCreateErrorContext(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion-ld", "--data", testData, "--context", "[\"http://context\""})

	err := batchCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestBatchCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := batchCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestBatchCreateErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"create", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := batchCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestBatchUpdate(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion-ld", "--data", testData})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/update"
	helper.SetClientHTTP(c, reqRes)

	err := batchUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBatchUpdateContext(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion-ld", "--data", testData, "--context", "[\"http://context\"]"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`[{"@context":["http://context"],"category":{"type":"Property","value":"sensor"},"id":"urn:ngsi-ld:TemperatureSensor:002","temperature":{"type":"Property","unitCode":"CEL","value":21},"type":"TemperatureSensor"},{"@context":["http://context"],"category":{"type":"Property","value":"sensor"},"id":"urn:ngsi-ld:TemperatureSensor:003","temperature":{"type":"Property","unitCode":"CEL","value":27},"type":"TemperatureSensor"}]`)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/update"
	helper.SetClientHTTP(c, reqRes)

	err := batchUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBatchUpdateErrorReadAll(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion-ld", "--data", "@"})

	err := batchUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchUpdateErrorContext(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion-ld", "--data", testData, "--context", "[\"http://context\""})

	err := batchUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestBatchUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entityOperations/update"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := batchUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestBatchUpdateErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"update", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/update"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := batchUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestBatchUpsertNoContent(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", testData})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/upsert"
	helper.SetClientHTTP(c, reqRes)

	err := batchUpsert(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBatchUpsertCreateted(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", testData})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(testData)
	reqRes.ResBody = []byte(`["urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/upsert"
	helper.SetClientHTTP(c, reqRes)

	err := batchUpsert(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\"urn:ngsi-ld:TemperatureSensor:002\",\"urn:ngsi-ld:TemperatureSensor:003\"]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestBatchUpsertContext(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", testData, "--context", "[\"http://context\"]"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`[{"@context":["http://context"],"category":{"type":"Property","value":"sensor"},"id":"urn:ngsi-ld:TemperatureSensor:002","temperature":{"type":"Property","unitCode":"CEL","value":21},"type":"TemperatureSensor"},{"@context":["http://context"],"category":{"type":"Property","value":"sensor"},"id":"urn:ngsi-ld:TemperatureSensor:003","temperature":{"type":"Property","unitCode":"CEL","value":27},"type":"TemperatureSensor"}]`)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/upsert"
	helper.SetClientHTTP(c, reqRes)

	err := batchUpsert(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBatchUpsertErrorReadAll(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", "@"})

	err := batchUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchUpsertErrorContext(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", testData, "--context", "[\"http://context\""})

	err := batchUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestBatchUpsertErrorHTTP(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entityOperations/upsert"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := batchUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestBatchUpsertErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/upsert"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := batchUpsert(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestBatchDelete(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", testData})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte(testData)
	reqRes.Path = "/ngsi-ld/v1/entityOperations/delete"
	helper.SetClientHTTP(c, reqRes)

	err := batchDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBatchDeleteErrorReadAll(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", "@"})

	err := batchDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestBatchDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entityOperations/delete"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := batchDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestBatchDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"upsert", "entities", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entityOperations/delete"
	reqRes.ResBody = []byte("error")
	helper.SetClientHTTP(c, reqRes)

	err := batchDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

var testData = `[
    {
      "id": "urn:ngsi-ld:TemperatureSensor:002",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 21,
            "unitCode": "CEL"
      }
    },
    {
      "id": "urn:ngsi-ld:TemperatureSensor:003",
      "type": "TemperatureSensor",
      "category": {
            "type": "Property",
            "value": "sensor"
      },
      "temperature": {
            "type": "Property",
            "value": 27,
            "unitCode": "CEL"
      }
    }
]`
