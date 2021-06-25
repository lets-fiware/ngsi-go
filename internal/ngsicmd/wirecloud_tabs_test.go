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

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestWcTabsList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33"})

	err := wireCloudTabsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "36 tab tab\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListArg(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "33"})

	err := wireCloudTabsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "36 tab tab\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--json"})

	err := wireCloudTabsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"36\",\"name\":\"tab\",\"title\":\"tab\",\"visible\":true,\"preferences\":{\"public\":{\"inherit\":false,\"value\":\"\"},\"requireauth\":{\"inherit\":true,\"value\":\"false\"},\"sharelist\":{\"inherit\":true,\"value\":\"[]\"},\"initiallayout\":{\"inherit\":true,\"value\":\"Fixed\"},\"baselayout\":{\"inherit\":true,\"value\":\"{\\\"type\\\":\\\"columnlayout\\\",\\\"smart\\\":\\\"false\\\",\\\"columns\\\":20,\\\"cellheight\\\":12,\\\"horizontalmargin\\\":4,\\\"verticalmargin\\\":3}\"}},\"iwidgets\":[]}]"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--pretty"})

	err := wireCloudTabsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"id\": \"36\",\n    \"name\": \"tab\",\n    \"title\": \"tab\",\n    \"visible\": true,\n    \"preferences\": {\n      \"public\": {\n        \"inherit\": false,\n        \"value\": \"\"\n      },\n      \"requireauth\": {\n        \"inherit\": true,\n        \"value\": \"false\"\n      },\n      \"sharelist\": {\n        \"inherit\": true,\n        \"value\": \"[]\"\n      },\n      \"initiallayout\": {\n        \"inherit\": true,\n        \"value\": \"Fixed\"\n      },\n      \"baselayout\": {\n        \"inherit\": true,\n        \"value\": \"{\\\"type\\\":\\\"columnlayout\\\",\\\"smart\\\":\\\"false\\\",\\\"columns\\\":20,\\\"cellheight\\\":12,\\\"horizontalmargin\\\":4,\\\"verticalmargin\\\":3}\"\n      }\n    },\n    \"iwidgets\": []\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc"})

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListErrorWid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "workspace id required", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33"})

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListErrorNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33"})

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "workspace not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33"})

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33"})

	setJSONDecodeErr(ngsi, 1)

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--pretty"})

	setJSONEncodeErr(ngsi, 2)

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcTabsListIotaErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--pretty"})

	setJSONIndentError(ngsi)

	err := wireCloudTabsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcTabGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--tid=36"})

	err := wireCloudTabGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\": \"36\", \"iwidgets\": [], \"name\": \"tab\", \"preferences\": {\"baselayout\": {\"inherit\": true, \"value\": \"{\\\"type\\\": \\\"columnlayout\\\", \\\"smart\\\": \\\"false\\\", \\\"columns\\\": 20, \\\"cellheight\\\": 12, \\\"horizontalmargin\\\": 4, \\\"verticalmargin\\\": 3}\"}, \"initiallayout\": {\"inherit\": true, \"value\": \"Fixed\"}, \"requireauth\": {\"inherit\": true, \"value\": \"false\"}, \"sharelist\": {\"inherit\": true, \"value\": \"[]\"}}, \"title\": \"tab\", \"visible\": true}"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabGetArgs(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "33", "36"})

	err := wireCloudTabGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\": \"36\", \"iwidgets\": [], \"name\": \"tab\", \"preferences\": {\"baselayout\": {\"inherit\": true, \"value\": \"{\\\"type\\\": \\\"columnlayout\\\", \\\"smart\\\": \\\"false\\\", \\\"columns\\\": 20, \\\"cellheight\\\": 12, \\\"horizontalmargin\\\": 4, \\\"verticalmargin\\\": 3}\"}, \"initiallayout\": {\"inherit\": true, \"value\": \"Fixed\"}, \"requireauth\": {\"inherit\": true, \"value\": \"false\"}, \"sharelist\": {\"inherit\": true, \"value\": \"[]\"}}, \"title\": \"tab\", \"visible\": true}"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--tid=36", "--pretty"})

	err := wireCloudTabGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"id\": \"36\",\n  \"iwidgets\": [],\n  \"name\": \"tab\",\n  \"preferences\": {\n    \"baselayout\": {\n      \"inherit\": true,\n      \"value\": \"{\\\"type\\\": \\\"columnlayout\\\", \\\"smart\\\": \\\"false\\\", \\\"columns\\\": 20, \\\"cellheight\\\": 12, \\\"horizontalmargin\\\": 4, \\\"verticalmargin\\\": 3}\"\n    },\n    \"initiallayout\": {\n      \"inherit\": true,\n      \"value\": \"Fixed\"\n    },\n    \"requireauth\": {\n      \"inherit\": true,\n      \"value\": \"false\"\n    },\n    \"sharelist\": {\n      \"inherit\": true,\n      \"value\": \"[]\"\n    }\n  },\n  \"title\": \"tab\",\n  \"visible\": true\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc"})

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabGetErrorWid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "workspace id required", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabGetErrortid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33"})

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "tab id required", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--tid=36"})

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcTabGetErrorNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--tid=36"})

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "workspace or tab not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcTabGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--tid=36"})

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcTabGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,wid,tid")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--wid=33", "--tid=36", "--pretty"})

	setJSONIndentError(ngsi)

	err := wireCloudTabGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
