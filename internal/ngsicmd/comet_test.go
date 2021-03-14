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

func TestCometAttrReadMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"type\":\"StructuredValue\",\"value\":[{\"recvTime\":\"2016-09-13T00:00:00.000Z\",\"attrType\":\"Number\",\"attrValue\":1},{\"recvTime\":\"2016-09-13T00:01:00.000Z\",\"attrType\":\"Number\",\"attrValue\":2},{\"recvTime\":\"2016-09-13T00:02:00.000Z\",\"attrType\":\"Number\",\"attrValue\":3}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainDate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,fromDate,toDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3", "--fromDate=1day", "--toDate=2days"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"type\":\"StructuredValue\",\"value\":[{\"recvTime\":\"2016-09-13T00:00:00.000Z\",\"attrType\":\"Number\",\"attrValue\":1},{\"recvTime\":\"2016-09-13T00:01:00.000Z\",\"attrType\":\"Number\",\"attrValue\":2},{\"recvTime\":\"2016-09-13T00:02:00.000Z\",\"attrType\":\"Number\",\"attrValue\":3}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainSafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"type\":\"StructuredValue\",\"value\":[{\"recvTime\":\"2016-09-13T00:00:00.000Z\",\"attrType\":\"Number\",\"attrValue\":1},{\"recvTime\":\"2016-09-13T00:01:00.000Z\",\"attrType\":\"Number\",\"attrValue\":2},{\"recvTime\":\"2016-09-13T00:02:00.000Z\",\"attrType\":\"Number\",\"attrValue\":3}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"type\": \"StructuredValue\",\n  \"value\": [\n    {\n      \"recvTime\": \"2016-09-13T00:00:00.000Z\",\n      \"attrType\": \"Number\",\n      \"attrValue\": 1\n    },\n    {\n      \"recvTime\": \"2016-09-13T00:01:00.000Z\",\n      \"attrType\": \"Number\",\n      \"attrValue\": 2\n    },\n    {\n      \"recvTime\": \"2016-09-13T00:02:00.000Z\",\n      \"attrType\": \"Number\",\n      \"attrValue\": 3\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrReadMainErrorNoType(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestCometAttrReadMainNoWayToConsumeData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "no way to consume data", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorDate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,fromDate")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3", "--fromDate=1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/entities/device001/attrs/A1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3", "--safeString=on"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '}' after array element (256) ,\"attrValue\":3}}", ngsiErr.Message)
	}
}

func TestCometAttrReadMainErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v2/entities/device001/attrs/A1"
	reqRes.ResBody = []byte(`{"type":"StructuredValue","value":[{"recvTime":"2016-09-13T00:00:00.000Z","attrType":"Number","attrValue":1},{"recvTime":"2016-09-13T00:01:00.000Z","attrType":"Number","attrValue":2},{"recvTime":"2016-09-13T00:02:00.000Z","attrType":"Number","attrValue":3}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attr,hLimit")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001", "--type=device", "--attr=A1", "--hLimit=3", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = cometAttrReadMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestCometEntitiesDeleteMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntitiesDeleteMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "all the data associated to certain service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometEntitiesDeleteMainWithRun(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntitiesDeleteMain(c, ngsi, client)

	assert.NoError(t, err)
}

func TestCometEntitiesDeleteMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/contextEntities"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntitiesDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestCometEntitiesDeleteMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/contextEntities"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntitiesDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntityDeleteMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "all the data associated to entity <device001>, service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometEntityDeleteMainWithRun(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntityDeleteMain(c, ngsi, client)

	assert.NoError(t, err)
}

func TestCometEntityDeleteMainErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMainErrorType(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestCometEntityDeleteMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometEntityDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}

func TestCometAttrDelete(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host,type,id,attr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--attr=A1"})

	err := cometAttrDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "all the data associated to attribute <A1> of entity <device001>, service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := cometAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	}
}

func TestCometAttrDeleteErrorNotSupported(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := cometAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMain(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id,attr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--attr=A1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrDeleteMain(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "all the data associated to attribute <A1> of entity <device001>, service and service path wiil be removed. run delete with -run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCometAttrDeleteMainWithRun(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id,attr")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--attr=A1", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrDeleteMain(c, ngsi, client)

	assert.NoError(t, err)
}

func TestCometAttrDeleteMainErrorType(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id,attr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--id=device001"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id,attr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorAttrName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id,attr")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "missing attr", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/STH/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id,attr")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--attr=A1", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestCometAttrDeleteMainErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/STH/v1/contextEntities/type/device/id/device001/attributes/A1"
	reqRes.ResBody = []byte(``)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,id,attr")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--type=device", "--id=device001", "--attr=A1", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"comet"})
	assert.NoError(t, err)

	err = cometAttrDeleteMain(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	}
}
