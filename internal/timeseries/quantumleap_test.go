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

package timeseries

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestQlEntitiesReadMain(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3", "--fromDate", "1day"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"Event001\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"},{\"id\":\"Event002\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"}]"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesReadMainSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"Event001\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"},{\"id\":\"Event002\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"}]"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesReadMainPretty(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"Event001\",\n    \"index\": [\n      \"2016-11-13T00:11:22\"\n    ],\n    \"type\": \"Event\"\n  },\n  {\n    \"id\": \"Event002\",\n    \"index\": [\n      \"2016-11-13T00:11:22\"\n    ],\n    \"type\": \"Event\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesReadMainErrorDate(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3", "--fromDate", "123"})

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error: \"type\":\"Event\"}", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorPretty(t *testing.T) {
	c := setupTest([]string{"hget", "entities", "--host", "ql", "--hLimit", "3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := qlEntitiesRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestQlAttrReadMainLastN(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--id", "device001", "--attr", "A1", "--lastN", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainLimitOffset(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--id", "device001", "--attr", "A1", "--hLimit", "10", "--hOffset", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainValue(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--id", "device001", "--attr", "A1", "--hLimit", "10", "--hOffset", "1", "--value"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/attrs/A1/value"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainNtypes(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--id", "device001", "--attr", "A1", "--nTypes"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainSameType(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--attr", "A1", "--sameType", "--fromDate", "1day"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--attr", "A1", "--sameType", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91,92,93]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainPretty(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--pretty", "--attr", "A1", "--sameType"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"attrName\": \"A1\",\n  \"entityId\": \"device001\",\n  \"index\": [\n    \"2016-09-13T03:01:00.000+00:00\",\n    \"2016-09-13T03:03:00.000+00:00\",\n    \"2016-09-13T03:05:00.000+00:00\"\n  ],\n  \"values\": [\n    91.0,\n    92.0,\n    93.0\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainErrorAttrName(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql"})

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing attr", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorTypes(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--sameType", "--nTypes", "--attr", "A1"})

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "sameType and nTypes are incompatible", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorGeo(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--sameType", "--attr", "A1", "--georel", "line"})

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "georel, geometry and coords are needed", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorType(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--sameType", "--attr", "A1"})

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorID(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--attr", "A1"})

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorDate(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--pretty", "--attr", "A1", "--sameType", "--fromDate", "123"})

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--pretty", "--attr", "A1", "--sameType"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--pretty", "--attr", "A1", "--sameType"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--attr", "A1", "--sameType", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error: ues\":[91,92,93]", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorPretty(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql", "--type", "device", "--pretty", "--attr", "A1", "--sameType"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := qlAttrReadMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestQlAttrsReadMain(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--lastN", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainSameType(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--type", "device", "--sameType", "--attrs", "A1,A2", "--lastN", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainNtypes(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--type", "device", "--nTypes", "--lastN", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/attrs"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainDate(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--lastN", "3", "--fromDate", "1day"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainLimitOffset(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--hLimit", "10", "--hOffset", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainValue(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--value"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/value"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1,2,3]},{\"attrName\":\"A2\",\"values\":[2,3,4]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainPretty(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--safeString", "on", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"attributes\": [\n    {\n      \"attrName\": \"A1\",\n      \"values\": [\n        1,\n        2,\n        3\n      ]\n    },\n    {\n      \"attrName\": \"A2\",\n      \"values\": [\n        2,\n        3,\n        4\n      ]\n    }\n  ],\n  \"entityId\": \"device001\",\n  \"index\": [\n    \"2016-09-13T00:01:00.000+00:00\",\n    \"2016-09-13T00:03:00.000+00:00\",\n    \"2016-09-13T00:05:00.000+00:00\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainErrorGeo(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--lastN", "3", "--georel", "line"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "georel, geometry and coords are needed", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorType(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--sameType", "--attrs", "A1,A2", "--lastN", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorID(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--attrs", "sA1,A2", "--lastN", "3", "--fromDate", "123"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorDate(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--lastN", "3", "--fromDate", "123"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--lastN", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--lastN", "3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/device001"

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorSafeString(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]`)

	helper.SetClientHTTP(c, reqRes)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error: :00.000+00:00\"]", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorPretty(t *testing.T) {
	c := setupTest([]string{"hget", "attrs", "--host", "ql", "--id", "device001", "--attrs", "A1,A2", "--safeString", "on", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := qlAttrsRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMain(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "ql", "--id", "device001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"

	helper.SetClientHTTP(c, reqRes)

	err := qlEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "historical data of the entity <device001> will be deleted. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntityDeleteMainWithRun(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "ql", "--id", "device001", "--fromDate", "2days", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/device001"

	helper.SetClientHTTP(c, reqRes)

	err := qlEntityDeleteMain(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestQlEntityDeleteMainErrorID(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "ql", "--fromDate", "2days", "--run"})

	err := qlEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMainErrorDate(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "ql", "--id", "device001", "--fromDate", "123", "--run"})

	err := qlEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "ql", "--id", "device001", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/entities/device001"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := qlEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "ql", "--id", "device001", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/device001"

	helper.SetClientHTTP(c, reqRes)

	err := qlEntityDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMain(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql", "--type", "device"})

	err := qlEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "historical data of all entities of the type <device> will be deleted. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesDeleteMainWithDropTable(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql", "--type", "device", "--dropTable"})

	err := qlEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "historical data of all entities of the type <device> will be deleted. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesDeleteMainWithRun(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql", "--type", "device", "--fromDate", "1day", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/types/device"

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesDeleteMain(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestQlEntitiesDeleteMainErrorType(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql", "--run"})

	err := qlEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMainErrorDate(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql", "--type", "device", "--fromDate", "123", "--run"})

	err := qlEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMainErrorHTTP(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql", "--type", "device", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/types/device"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMainErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql", "--type", "device", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/device"

	helper.SetClientHTTP(c, reqRes)

	err := qlEntitiesDeleteMain(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}
