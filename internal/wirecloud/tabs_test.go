/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestWcTabsList(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "36 tab tab\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListArg(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "33"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "36 tab tab\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListJSON(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"36\",\"name\":\"tab\",\"title\":\"tab\",\"visible\":true,\"preferences\":{\"public\":{\"inherit\":false,\"value\":\"\"},\"requireauth\":{\"inherit\":true,\"value\":\"false\"},\"sharelist\":{\"inherit\":true,\"value\":\"[]\"},\"initiallayout\":{\"inherit\":true,\"value\":\"Fixed\"},\"baselayout\":{\"inherit\":true,\"value\":\"{\\\"type\\\":\\\"columnlayout\\\",\\\"smart\\\":\\\"false\\\",\\\"columns\\\":20,\\\"cellheight\\\":12,\\\"horizontalmargin\\\":4,\\\"verticalmargin\\\":3}\"}},\"iwidgets\":[]}]"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListPretty(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"36\",\n    \"name\": \"tab\",\n    \"title\": \"tab\",\n    \"visible\": true,\n    \"preferences\": {\n      \"public\": {\n        \"inherit\": false,\n        \"value\": \"\"\n      },\n      \"requireauth\": {\n        \"inherit\": true,\n        \"value\": \"false\"\n      },\n      \"sharelist\": {\n        \"inherit\": true,\n        \"value\": \"[]\"\n      },\n      \"initiallayout\": {\n        \"inherit\": true,\n        \"value\": \"Fixed\"\n      },\n      \"baselayout\": {\n        \"inherit\": true,\n        \"value\": \"{\\\"type\\\":\\\"columnlayout\\\",\\\"smart\\\":\\\"false\\\",\\\"columns\\\":20,\\\"cellheight\\\":12,\\\"horizontalmargin\\\":4,\\\"verticalmargin\\\":3}\"\n      }\n    },\n    \"iwidgets\": []\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabsListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestWcTabsListErrorNotFound(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "workspace not found", ngsiErr.Message)
	}
}

func TestWcTabsListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/workspace/33"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestWcTabsListErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcTabsListErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcTabsListIotaErrorPretty(t *testing.T) {
	c := setupTest([]string{"tabs", "list", "--host", "wirecloud", "--wid", "33", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tabs":[{"id":"36","name":"tab","title":"tab","visible":true,"preferences":{"public":{"inherit":false,"value":""},"requireauth":{"inherit":true,"value":"false"},"sharelist":{"inherit":true,"value":"[]"},"initiallayout":{"inherit":true,"value":"Fixed"},"baselayout":{"inherit":true,"value":"{\"type\":\"columnlayout\",\"smart\":\"false\",\"columns\":20,\"cellheight\":12,\"horizontalmargin\":4,\"verticalmargin\":3}"}},"iwidgets":[]}]}`)
	reqRes.Path = "/api/workspace/33"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := wireCloudTabsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcTabGet(t *testing.T) {
	c := setupTest([]string{"tabs", "get", "--host", "wirecloud", "--wid", "33", "--tid", "36"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\": \"36\", \"iwidgets\": [], \"name\": \"tab\", \"preferences\": {\"baselayout\": {\"inherit\": true, \"value\": \"{\\\"type\\\": \\\"columnlayout\\\", \\\"smart\\\": \\\"false\\\", \\\"columns\\\": 20, \\\"cellheight\\\": 12, \\\"horizontalmargin\\\": 4, \\\"verticalmargin\\\": 3}\"}, \"initiallayout\": {\"inherit\": true, \"value\": \"Fixed\"}, \"requireauth\": {\"inherit\": true, \"value\": \"false\"}, \"sharelist\": {\"inherit\": true, \"value\": \"[]\"}}, \"title\": \"tab\", \"visible\": true}"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabGetArgs(t *testing.T) {
	c := setupTest([]string{"tabs", "get", "--host", "wirecloud", "33", "36"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\": \"36\", \"iwidgets\": [], \"name\": \"tab\", \"preferences\": {\"baselayout\": {\"inherit\": true, \"value\": \"{\\\"type\\\": \\\"columnlayout\\\", \\\"smart\\\": \\\"false\\\", \\\"columns\\\": 20, \\\"cellheight\\\": 12, \\\"horizontalmargin\\\": 4, \\\"verticalmargin\\\": 3}\"}, \"initiallayout\": {\"inherit\": true, \"value\": \"Fixed\"}, \"requireauth\": {\"inherit\": true, \"value\": \"false\"}, \"sharelist\": {\"inherit\": true, \"value\": \"[]\"}}, \"title\": \"tab\", \"visible\": true}"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabGetPretty(t *testing.T) {
	c := setupTest([]string{"tabs", "get", "--host", "wirecloud", "--wid", "33", "--tid", "36", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"36\",\n  \"iwidgets\": [],\n  \"name\": \"tab\",\n  \"preferences\": {\n    \"baselayout\": {\n      \"inherit\": true,\n      \"value\": \"{\\\"type\\\": \\\"columnlayout\\\", \\\"smart\\\": \\\"false\\\", \\\"columns\\\": 20, \\\"cellheight\\\": 12, \\\"horizontalmargin\\\": 4, \\\"verticalmargin\\\": 3}\"\n    },\n    \"initiallayout\": {\n      \"inherit\": true,\n      \"value\": \"Fixed\"\n    },\n    \"requireauth\": {\n      \"inherit\": true,\n      \"value\": \"false\"\n    },\n    \"sharelist\": {\n      \"inherit\": true,\n      \"value\": \"[]\"\n    }\n  },\n  \"title\": \"tab\",\n  \"visible\": true\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcTabGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"tabs", "get", "--host", "wirecloud", "--wid", "33", "--tid", "36"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestWcTabGetErrorNotFound(t *testing.T) {
	c := setupTest([]string{"tabs", "get", "--host", "wirecloud", "--wid", "33", "--tid", "36"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "workspace or tab not found", ngsiErr.Message)
	}
}

func TestWcTabGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"tabs", "get", "--host", "wirecloud", "--wid", "33", "--tid", "36"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/workspace/33/tab/36"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := wireCloudTabGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestWcTabGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"tabs", "get", "--host", "wirecloud", "--wid", "33", "--tid", "36", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "36", "iwidgets": [], "name": "tab", "preferences": {"baselayout": {"inherit": true, "value": "{\"type\": \"columnlayout\", \"smart\": \"false\", \"columns\": 20, \"cellheight\": 12, \"horizontalmargin\": 4, \"verticalmargin\": 3}"}, "initiallayout": {"inherit": true, "value": "Fixed"}, "requireauth": {"inherit": true, "value": "false"}, "sharelist": {"inherit": true, "value": "[]"}}, "title": "tab", "visible": true}`)
	reqRes.Path = "/api/workspace/33/tab/36"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := wireCloudTabGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
