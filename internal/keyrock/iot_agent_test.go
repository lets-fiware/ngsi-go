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

package keyrock

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestIotAgentsList(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agents":[{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"},{"id":"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d"},{"id":"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89"}]}	`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7\niot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d\niot_sensor_f95fc041-7102-4be7-8948-37fb571afa89\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsListVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agents":[{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"},{"id":"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d"},{"id":"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89"}]}	`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"iot_agents\":[{\"id\":\"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7\"},{\"id\":\"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d\"},{\"id\":\"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89\"}]}\t"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agents":[{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"},{"id":"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d"},{"id":"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89"}]}	`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"iot_agents\": [\n    {\n      \"id\": \"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7\"\n    },\n    {\n      \"id\": \"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d\"\n    },\n    {\n      \"id\": \"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89\"\n    }\n  ]\n}\t\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsListNotFound(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "iot agents not found\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agents":[{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"},{"id":"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d"},{"id":"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89"}]}	`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestIotAgentsListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIotAgentsListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agents":[{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"},{"id":"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d"},{"id":"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89"}]}	`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIotAgentsListErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agents":[{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"},{"id":"iot_sensor_74df472c-0fd1-4b8a-8f13-c6848307bc7d"},{"id":"iot_sensor_f95fc041-7102-4be7-8948-37fb571afa89"}]}	`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := iotAgentsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIotAgentssGet(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"iot_agent\":{\"id\":\"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentssGetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"iot_agent\": {\n    \"id\": \"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestIotAgentsGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIotAgentsGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := iotAgentsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIotAgentsCreate(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","password":"iot_sensor_14569214-60d7-44d1-856c-ec8d8967aba2"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"iot_agent\":{\"id\":\"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7\",\"password\":\"iot_sensor_14569214-60d7-44d1-856c-ec8d8967aba2\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsCreatePretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","password":"iot_sensor_14569214-60d7-44d1-856c-ec8d8967aba2"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"iot_agent\": {\n    \"id\": \"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7\",\n    \"password\": \"iot_sensor_14569214-60d7-44d1-856c-ec8d8967aba2\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agent"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","password":"iot_sensor_14569214-60d7-44d1-856c-ec8d8967aba2"}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIotAgentsCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIotAgentsCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents"
	reqRes.ResBody = []byte(`{"iot_agent":{"id":"iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7","password":"iot_sensor_14569214-60d7-44d1-856c-ec8d8967aba2"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := iotAgentsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIotAgentsReset(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"
	reqRes.ResBody = []byte(`{"new_password": "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsReset(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"new_password\": \"iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsResetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"
	reqRes.ResBody = []byte(`{"new_password": "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsReset(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"new_password\": \"iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIotAgentsResetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"
	reqRes.ResBody = []byte(`{"new_password": "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsReset(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestIotAgentsResetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsReset(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIotAgentsResetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"
	reqRes.ResBody = []byte(`{"new_password": "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := iotAgentsReset(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
func TestIotAgentsDelete(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIotAgentsDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestIotAgentsDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "iota", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/iot_agents/iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := iotAgentsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
