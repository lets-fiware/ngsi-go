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

package timeseries

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestCometAttrReadMain(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"StructuredValue\",\"value\":[{\"recvTime\":\"2016-09-13T00:00:00.000Z\",\"attrType\":\"Number\",\"attrValue\":1},{\"recvTime\":\"2016-09-13T00:01:00.000Z\",\"attrType\":\"Number\",\"attrValue\":2},{\"recvTime\":\"2016-09-13T00:02:00.000Z\",\"attrType\":\"Number\",\"attrValue\":3}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainDate(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3", "--fromDate", "1day", "--toDate", "2days"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"StructuredValue\",\"value\":[{\"recvTime\":\"2016-09-13T00:00:00.000Z\",\"attrType\":\"Number\",\"attrValue\":1},{\"recvTime\":\"2016-09-13T00:01:00.000Z\",\"attrType\":\"Number\",\"attrValue\":2},{\"recvTime\":\"2016-09-13T00:02:00.000Z\",\"attrType\":\"Number\",\"attrValue\":3}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"StructuredValue\",\"value\":[{\"recvTime\":\"2016-09-13T00:00:00.000Z\",\"attrType\":\"Number\",\"attrValue\":1},{\"recvTime\":\"2016-09-13T00:01:00.000Z\",\"attrType\":\"Number\",\"attrValue\":2},{\"recvTime\":\"2016-09-13T00:02:00.000Z\",\"attrType\":\"Number\",\"attrValue\":3}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainPretty(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"type\": \"StructuredValue\",\n  \"value\": [\n    {\n      \"recvTime\": \"2016-09-13T00:00:00.000Z\",\n      \"attrType\": \"Number\",\n      \"attrValue\": 1\n    },\n    {\n      \"recvTime\": \"2016-09-13T00:01:00.000Z\",\n      \"attrType\": \"Number\",\n      \"attrValue\": 2\n    },\n    {\n      \"recvTime\": \"2016-09-13T00:02:00.000Z\",\n      \"attrType\": \"Number\",\n      \"attrValue\": 3\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainErrorNoType(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestCometAttrReadMainNoWayToConsumeData(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no way to consume data", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorDate(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3", "--fromDate", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/entities/device001/attrs/A1"

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}}`)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '}' after array element (256) ,\"attrValue\":3}}", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorPretty(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet", "--id", "device001", "--type", "device", "--attr", "A1", "--hLimit", "3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := cometAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestCometEntitiesDeleteMain(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "comet"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "all the data associated to certain service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometEntitiesDeleteMainWithRun(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "comet", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometEntitiesDeleteMain(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestCometEntitiesDeleteMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "comet", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := cometEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestCometEntitiesDeleteMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "comet", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/contextEntities"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := cometEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMain(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "comet", "--type", "device", "--id", "device001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "all the data associated to entity <device001>, service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometEntityDeleteMainWithRun(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "comet", "--type", "device", "--id", "device001", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometEntityDeleteMain(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestCometEntityDeleteMainErrorID(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "comet", "--type", "device"})

	err := cometEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMainErrorType(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "comet", "--id", "device001"})

	err := cometEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "comet", "--type", "device", "--id", "device001", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := cometEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "comet", "--type", "device", "--id", "device001", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := cometEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestCometAttrDelete(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--type", "device", "--id", "device001", "--attr", "A1"})

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "all the data associated to attribute <A1> of entity <device001>, service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrDeleteMain(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--type", "device", "--id", "device001", "--attr", "A1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "all the data associated to attribute <A1> of entity <device001>, service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrDeleteMainWithRun(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--type", "device", "--id", "device001", "--attr", "A1", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestCometAttrDeleteMainErrorType(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--id", "device001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorID(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--type", "device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorAttrName(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--type", "device", "--id", "device001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "missing attr", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--type", "device", "--id", "device001", "--attr", "A1", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hdelete", "attr", "--host", "comet", "--type", "device", "--id", "device001", "--attr", "A1", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := cometAttrDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}
