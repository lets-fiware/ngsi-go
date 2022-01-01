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

package wirecloud

import (
	"errors"
	"net/http"
	"sort"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestWcWorkspacesList(t *testing.T) {
	c := setupTest([]string{"workspaces", "list", "--host", "wirecloud"})

	c.Ngsi.TimeLib = &helper.MockTimeLib{TimeFormat: helper.StrPtr("2016/10/28 00:00:00")}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspacesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1 ws1 ws1 2016/10/28 00:00:00\n1 ws2 ws2 2016/10/28 00:00:00\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspacesListJSON(t *testing.T) {
	c := setupTest([]string{"workspaces", "list", "--host", "wirecloud", "--json"})

	c.Ngsi.TimeLib = &helper.MockTimeLib{TimeFormat: helper.StrPtr("2016/10/28 00:00:00")}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspacesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\": \"1\", \"name\": \"ws1\", \"title\": \"ws1\", \"lastmodified\": 1624744059084},{\"id\": \"1\", \"name\": \"ws2\", \"title\": \"ws2\", \"lastmodified\": 1624744059084}]"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspacesListPretty(t *testing.T) {
	c := setupTest([]string{"workspaces", "list", "--host", "wirecloud", "--pretty"})

	c.Ngsi.TimeLib = &helper.MockTimeLib{TimeFormat: helper.StrPtr("2016/10/28 00:00:00")}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspacesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"1\",\n    \"name\": \"ws1\",\n    \"title\": \"ws1\",\n    \"lastmodified\": 1624744059084\n  },\n  {\n    \"id\": \"1\",\n    \"name\": \"ws2\",\n    \"title\": \"ws2\",\n    \"lastmodified\": 1624744059084\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspacesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"workspaces", "list", "--host", "wirecloud"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspacesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestWcWorkspacesListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"workspaces", "list", "--host", "wirecloud"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/workspaces"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspacesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestWcWorkspacesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"workspaces", "list", "--host", "wirecloud", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := wireCloudWorkspacesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcWorkspacesListErrorList(t *testing.T) {
	c := setupTest([]string{"workspaces", "list", "--host", "wirecloud"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id": "1", "name": "ws1", "title": "ws1", "lastmodified": 1624744059084},{"id": "1", "name": "ws2", "title": "ws2", "lastmodified": 1624744059084}]`)
	reqRes.Path = "/api/workspaces"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := wireCloudWorkspacesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcWorkspaceGet(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1"})

	c.Ngsi.TimeLib = &helper.MockTimeLib{TimeFormat: helper.StrPtr("2016/10/28 00:00:00")}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1 ws1 ws1 2016/10/28 00:00:00\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetArg(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "1"})

	c.Ngsi.TimeLib = &helper.MockTimeLib{TimeFormat: helper.StrPtr("2016/10/28 00:00:00")}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1 ws1 ws1 2016/10/28 00:00:00\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetJSON(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--json"})

	c.Ngsi.TimeLib = &helper.MockTimeLib{TimeFormat: helper.StrPtr("2016/10/28 00:00:00")}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"1\",\"name\":\"ws1\",\"title\":\"ws1\",\"lastmodified\":1624744059084,\"users\":[{\"fullname\":\"\",\"username\":\"admin\",\"organization\":false,\"accesslevel\":\"owner\"}],\"tabs\":[{\"id\":\"17\",\"name\":\"tab\",\"title\":\"Tab\",\"visible\":true,\"preferences\":{}}],\"wiring\":{\"visualdescription\":{\"components\":{\"operator\":{},\"widget\":{\"1\":{\"name\":\"NGSIGO/ngsigo-widget/1.0.0\",\"position\":{\"x\":799,\"y\":139}}}}}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetPretty(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--pretty"})

	c.Ngsi.TimeLib = &helper.MockTimeLib{TimeFormat: helper.StrPtr("2016/10/28 00:00:00")}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"1\",\n  \"name\": \"ws1\",\n  \"title\": \"ws1\",\n  \"lastmodified\": 1624744059084,\n  \"users\": [\n    {\n      \"fullname\": \"\",\n      \"username\": \"admin\",\n      \"organization\": false,\n      \"accesslevel\": \"owner\"\n    }\n  ],\n  \"tabs\": [\n    {\n      \"id\": \"17\",\n      \"name\": \"tab\",\n      \"title\": \"Tab\",\n      \"visible\": true,\n      \"preferences\": {}\n    }\n  ],\n  \"wiring\": {\n    \"visualdescription\": {\n      \"components\": {\n        \"operator\": {},\n        \"widget\": {\n          \"1\": {\n            \"name\": \"NGSIGO/ngsigo-widget/1.0.0\",\n            \"position\": {\n              \"x\": 799,\n              \"y\": 139\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetTabs(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--tabs"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "17 tab Tab\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetWidgets(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--widgets"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "NGSIGO/ngsigo-widget/1.0.0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetOperators(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--operators"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{"1":{"name":"NGSIGO/ngsigo-operator/1.0.0","position":{"x":799,"y":139}}},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "NGSIGO/ngsigo-operator/1.0.0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetUsers(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--users"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{"1":{"name":"NGSIGO/ngsigo-operator/1.0.0","position":{"x":799,"y":139}}},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "admin owner\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcWorkspaceGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestWcWorkspaceGetErrorNotFound(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "workspace not found", ngsiErr.Message)
	}
}

func TestWcWorkspaceGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/workspace/1"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestWcWorkspaceGetErrorList(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--tabs"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
func TestWcWorkspaceGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"workspaces", "get", "--host", "wirecloud", "--wid", "1", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"1","name":"ws1","title":"ws1","lastmodified":1624744059084,"users":[{"fullname":"","username":"admin","organization":false,"accesslevel":"owner"}],"tabs":[{"id":"17","name":"tab","title":"Tab","visible":true,"preferences":{}}],"wiring":{"visualdescription":{"components":{"operator":{},"widget":{"1":{"name":"NGSIGO/ngsigo-widget/1.0.0","position":{"x":799,"y":139}}}}}}}`)
	reqRes.Path = "/api/workspace/1"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := wireCloudWorkspaceGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
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
