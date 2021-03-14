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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestQlEntitiesReadMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3", "--fromDate=1day"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"Event001\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"},{\"id\":\"Event002\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"}]"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesReadMainSafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"Event001\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"},{\"id\":\"Event002\",\"index\":[\"2016-11-13T00:11:22\"],\"type\":\"Event\"}]"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesReadMainPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,safeString")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"id\": \"Event001\",\n    \"index\": [\n      \"2016-11-13T00:11:22\"\n    ],\n    \"type\": \"Event\"\n  },\n  {\n    \"id\": \"Event002\",\n    \"index\": [\n      \"2016-11-13T00:11:22\"\n    ],\n    \"type\": \"Event\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesReadMainErrorDate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3", "--fromDate=123"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error: \"type\":\"Event\"}", ngsiErr.Message)
	}
}

func TestQlEntitiesReadMainErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte(`[{"id":"Event001","index":["2016-11-13T00:11:22"],"type":"Event"},{"id":"Event002","index":["2016-11-13T00:11:22"],"type":"Event"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,hLimit,hOffset,safeString")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--hLimit=3", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = qlEntitiesReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestQlAttrReadMainLastN(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1", "--lastN=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainLimitOffset(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1", "--hLimit=10", "--hOffset=1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainValue(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/attrs/A1/value"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	setupFlagBool(set, "value")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1", "--hLimit=10", "--hOffset=1", "--value"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainNtypes(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	setupFlagBool(set, "value,nTypes")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1", "--nTypes"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainSameType(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate")
	setupFlagBool(set, "value,nTypes,sameType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--attr=A1", "--sameType", "--fromDate=1day"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91.0,92.0,93.0]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainSafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--attr=A1", "--sameType", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attrName\":\"A1\",\"entityId\":\"device001\",\"index\":[\"2016-09-13T03:01:00.000+00:00\",\"2016-09-13T03:03:00.000+00:00\",\"2016-09-13T03:05:00.000+00:00\"],\"values\":[91,92,93]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--pretty", "--attr=A1", "--sameType"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"attrName\": \"A1\",\n  \"entityId\": \"device001\",\n  \"index\": [\n    \"2016-09-13T03:01:00.000+00:00\",\n    \"2016-09-13T03:03:00.000+00:00\",\n    \"2016-09-13T03:05:00.000+00:00\"\n  ],\n  \"values\": [\n    91.0,\n    92.0,\n    93.0\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrReadMainErrorAttrName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString,georel")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing attr", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorTypes(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString,georel")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--sameType", "--nTypes", "--attr=A1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "sameType and nTypes are incompatible", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorGeo(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString,georel")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--sameType", "--attr=A1", "--georel=line"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "georel, geometry and coords are needed", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorType(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--sameType", "--attr=A1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--attr=A1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorDate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--pretty", "--attr=A1", "--sameType", "--fromDate=123"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--pretty", "--attr=A1", "--sameType"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/device/attrs/A1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--pretty", "--attr=A1", "--sameType"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--attr=A1", "--sameType", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error: ues\":[91,92,93]", ngsiErr.Message)
	}
}

func TestQlAttrReadMainErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device/attrs/A1"
	reqRes.ResBody = []byte(`{"attrName":"A1","entityId":"device001","index":["2016-09-13T03:01:00.000+00:00","2016-09-13T03:03:00.000+00:00","2016-09-13T03:05:00.000+00:00"],"values":[91.0,92.0,93.0]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "value,nTypes,sameType,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--pretty", "--attr=A1", "--sameType"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = qlAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestQlAttrsReadMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--lastN=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainSameType(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	setupFlagBool(set, "sameType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--sameType", "--attr=A1,A2", "--lastN=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainNtypes(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/attrs"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	setupFlagBool(set, "nTypes")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--nTypes", "--lastN=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainDate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--lastN=3", "--fromDate=1day"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainLimitOffset(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--hLimit=10", "--hOffset=1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainValue(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001/value"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate")
	setupFlagBool(set, "value")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--value"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1.0,2.0,3.0]},{\"attrName\":\"A2\",\"values\":[2.0,3.0,4.0,]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainSafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attributes\":[{\"attrName\":\"A1\",\"values\":[1,2,3]},{\"attrName\":\"A2\",\"values\":[2,3,4]}],\"entityId\":\"device001\",\"index\":[\"2016-09-13T00:01:00.000+00:00\",\"2016-09-13T00:03:00.000+00:00\",\"2016-09-13T00:05:00.000+00:00\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--safeString=on", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"attributes\": [\n    {\n      \"attrName\": \"A1\",\n      \"values\": [\n        1,\n        2,\n        3\n      ]\n    },\n    {\n      \"attrName\": \"A2\",\n      \"values\": [\n        2,\n        3,\n        4\n      ]\n    }\n  ],\n  \"entityId\": \"device001\",\n  \"index\": [\n    \"2016-09-13T00:01:00.000+00:00\",\n    \"2016-09-13T00:03:00.000+00:00\",\n    \"2016-09-13T00:05:00.000+00:00\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlAttrsReadMainErrorGeo(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,georel")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--lastN=3", "--georel=line"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "georel, geometry and coords are needed", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorType(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN")
	setupFlagBool(set, "sameType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--sameType", "--attr=A1,A2", "--lastN=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--attr=A1,A2", "--lastN=3", "--fromDate=123"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorDate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--lastN=3", "--fromDate=123"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0,]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,georel")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--lastN=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/device001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,georel")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--lastN=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error: :00.000+00:00\"]", ngsiErr.Message)
	}
}

func TestQlAttrsReadMainErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	reqRes.ResBody = []byte(`{"attributes":[{"attrName":"A1","values":[1.0,2.0,3.0]},{"attrName":"A2","values":[2.0,3.0,4.0]}],"entityId":"device001","index":["2016-09-13T00:01:00.000+00:00","2016-09-13T00:03:00.000+00:00","2016-09-13T00:05:00.000+00:00"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,hOffset,lastN,fromDate,safeString")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--attr=A1,A2", "--safeString=on", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = qlAttrsReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/device001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntityDeleteMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "historical data of the entity <device001> will be deleted. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntityDeleteMainWithRun(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/device001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--fromDate=2days", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntityDeleteMain(c, ngsi, client)

	assert.NoError(t, err)
}

func TestQlEntityDeleteMainErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/device001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--fromDate=2days", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMainErrorDate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/device001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--fromDate=123", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/entities/device001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestQlEntityDeleteMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/device001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--id=device001", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesDeleteMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "historical data of all entities of the type <device> will be deleted. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesDeleteMainWithDropTable(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/device"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type")
	setupFlagBool(set, "dropTable")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--dropTable"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesDeleteMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "historical data of all entities of the type <device> will be deleted. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestQlEntitiesDeleteMainWithRun(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/types/device"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--fromDate=1day", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesDeleteMain(c, ngsi, client)

	assert.NoError(t, err)
}

func TestQlEntitiesDeleteMainErrorType(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/types/device"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMainErrorDate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/types/device"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate,toDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--fromDate=123", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error 123", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/types/device"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate,toDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestQlEntitiesDeleteMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/device"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,fromDate,toDate")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--type=device", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	assert.NoError(t, err)

	err = qlEntitiesDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}
