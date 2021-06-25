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
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestWcWorkspacesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	tt := "2016/10/28 00:00:00"
	ngsi.TimeLib = &MockTimeLib{format: &tt}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudWorkspacesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "1 ws1 ws1 2016/10/28 00:00:00\n1 ws2 ws2 2016/10/28 00:00:00\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspacesListJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	tt := "2016/10/28 00:00:00"
	ngsi.TimeLib = &MockTimeLib{format: &tt}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--json"})

	err := wireCloudWorkspacesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\": \"1\", \"name\": \"ws1\", \"title\": \"ws1\", \"lastmodified\": 1624744059084},{\"id\": \"1\", \"name\": \"ws2\", \"title\": \"ws2\", \"lastmodified\": 1624744059084}]"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspacesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	tt := "2016/10/28 00:00:00"
	ngsi.TimeLib = &MockTimeLib{format: &tt}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty"})

	err := wireCloudWorkspacesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"id\": \"1\",\n    \"name\": \"ws1\",\n    \"title\": \"ws1\",\n    \"lastmodified\": 1624744059084\n  },\n  {\n    \"id\": \"1\",\n    \"name\": \"ws2\",\n    \"title\": \"ws2\",\n    \"lastmodified\": 1624744059084\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspacesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudWorkspacesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspacesListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc"})

	err := wireCloudWorkspacesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspacesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudWorkspacesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspacesListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudWorkspacesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspacesListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty"})

	setJSONIndentError(ngsi)

	err := wireCloudWorkspacesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcWorkspacesListErrorList(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	setJSONDecodeErr(ngsi, 1)

	err := wireCloudWorkspacesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcWorkspaceGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	tt := "2016/10/28 00:00:00"
	ngsi.TimeLib = &MockTimeLib{format: &tt}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "1 ws1 ws1 2016/10/28 00:00:00\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetArg(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	tt := "2016/10/28 00:00:00"
	ngsi.TimeLib = &MockTimeLib{format: &tt}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "1"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "1 ws1 ws1 2016/10/28 00:00:00\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	tt := "2016/10/28 00:00:00"
	ngsi.TimeLib = &MockTimeLib{format: &tt}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--json"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"1\",\"name\":\"ws1\",\"title\":\"ws1\",\"lastmodified\":1624744059084,\"users\":[{\"fullname\":\"\",\"username\":\"admin\",\"organization\":false,\"accesslevel\":\"owner\"}],\"tabs\":[{\"id\":\"17\",\"name\":\"tab\",\"title\":\"Tab\",\"visible\":true,\"preferences\":{}}],\"wiring\":{\"visualdescription\":{\"components\":{\"operator\":{},\"widget\":{\"1\":{\"name\":\"NGSIGO/ngsigo-widget/1.0.0\",\"position\":{\"x\":799,\"y\":139}}}}}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	tt := "2016/10/28 00:00:00"
	ngsi.TimeLib = &MockTimeLib{format: &tt}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--pretty"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"id\": \"1\",\n  \"name\": \"ws1\",\n  \"title\": \"ws1\",\n  \"lastmodified\": 1624744059084,\n  \"users\": [\n    {\n      \"fullname\": \"\",\n      \"username\": \"admin\",\n      \"organization\": false,\n      \"accesslevel\": \"owner\"\n    }\n  ],\n  \"tabs\": [\n    {\n      \"id\": \"17\",\n      \"name\": \"tab\",\n      \"title\": \"Tab\",\n      \"visible\": true,\n      \"preferences\": {}\n    }\n  ],\n  \"wiring\": {\n    \"visualdescription\": {\n      \"components\": {\n        \"operator\": {},\n        \"widget\": {\n          \"1\": {\n            \"name\": \"NGSIGO/ngsigo-widget/1.0.0\",\n            \"position\": {\n              \"x\": 799,\n              \"y\": 139\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetTabs(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--tabs"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "17 tab Tab\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetWidgets(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--widgets"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "NGSIGO/ngsigo-widget/1.0.0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetOperators(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{"1":{"name":"NGSIGO/ngsigo-operator/1.0.0","position":{"x":799,"y":139}}},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--operators"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "NGSIGO/ngsigo-operator/1.0.0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetUsers(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{"1":{"name":"NGSIGO/ngsigo-operator/1.0.0","position":{"x":799,"y":139}}},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--users"})

	err := wireCloudWorkspaceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "admin owner\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspaceGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc", "--wid=1", "--pretty"})

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspaceGetErrorWid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "workspace id required", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspaceGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1"})

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspaceGetErrorNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1"})

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "workspace not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspaceGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1"})

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcWorkspaceGetErrorList(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--tabs"})

	setJSONDecodeErr(ngsi, 1)

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
func TestWcWorkspaceGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty,tabs,widgets,operators,users")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=1", "--pretty"})

	setJSONIndentError(ngsi)

	err := wireCloudWorkspaceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcSortWCWorkspaceInfos(t *testing.T) {

	var wss = wireCloudWorkspaceInfos{
		{ID: "1"},
		{ID: "100"},
		{ID: "10"},
	}
	sort.Sort(wireCloudWorkspaceInfos(wss))
}

func TestWcSortWCTabs(t *testing.T) {

	var tabs = wireCloudTabs{
		{ID: "1"},
		{ID: "100"},
		{ID: "10"},
	}
	sort.Sort(wireCloudTabs(tabs))
}
