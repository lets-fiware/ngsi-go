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

package iotagent

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestIdasServicesList(t *testing.T) {
	c := setupTest([]string{"services", "list", "--host", "iota"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"count\":1,\"services\":[{\"commands\":[],\"lazy\":[],\"attributes\":[],\"_id\":\"601e25597d7b3d691be82d23\",\"resource\":\"/iot/d\",\"apikey\":\"apikey\",\"service\":\"openiot\",\"subservice\":\"/\",\"__v\":0,\"static_attributes\":[],\"internal_attributes\":[],\"entity_type\":\"Event\"}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasServicesListPretty(t *testing.T) {
	c := setupTest([]string{"services", "list", "--host", "iota", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"count\": 1,\n  \"services\": [\n    {\n      \"commands\": [],\n      \"lazy\": [],\n      \"attributes\": [],\n      \"_id\": \"601e25597d7b3d691be82d23\",\n      \"resource\": \"/iot/d\",\n      \"apikey\": \"apikey\",\n      \"service\": \"openiot\",\n      \"subservice\": \"/\",\n      \"__v\": 0,\n      \"static_attributes\": [],\n      \"internal_attributes\": [],\n      \"entity_type\": \"Event\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasServicesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"services", "list", "--host", "iota"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/service"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasServicesListErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"services", "list", "--host", "iota"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasServicesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"services", "list", "--host", "iota", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := idasServicesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasServicesCreateData(t *testing.T) {
	data := `{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`
	c := setupTest([]string{"services", "create", "--host", "iota", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasServicesCreateParam(t *testing.T) {
	c := setupTest([]string{"services", "create", "--host", "iota", "--apikey", "apikey", "--cbroker", "http://orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event","resource":"/iot/d"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasServicesCreateErrorData(t *testing.T) {
	c := setupTest([]string{"services", "create", "--host", "iota", "--data", "@"})

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestIdasServicesCreateErrorCbroker(t *testing.T) {
	c := setupTest([]string{"services", "create", "--host", "iota", "--apikey", "apikey", "--cbroker", "orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event","resource":"/iot/d"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "specify url or broker alias to --cbroker", ngsiErr.Message)
	}
}

func TestIdasServicesCreateErrorMarshal(t *testing.T) {
	c := setupTest([]string{"services", "create", "--host", "iota", "--apikey", "apikey", "--cbroker", "http://orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event","resource":"/iot/d"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasServicesCreateErrorParam(t *testing.T) {
	c := setupTest([]string{"services", "create", "--host", "iota", "--apikey", "apikey", "--data", "@"})

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "apikey, type and resource are needed", ngsiErr.Message)
	}
}

func TestIdasServicesCreateErrorHTTP(t *testing.T) {
	data := `{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`
	c := setupTest([]string{"services", "create", "--host", "iota", "apikey", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasServicesCreateErrorHTTPStatus(t *testing.T) {
	data := `{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`
	c := setupTest([]string{"services", "create", "--host", "iota", "apikey", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasServicesUpdateData(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--apikey", "apikey", "--resource", "/iot/d", "--data", `{"type":"Event"}`})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"type":"Event"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasServicesUpdateParam(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--apikey", "apikey", "--cbroker", "http://orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasServicesUpdateErrorData(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--apikey", "apikey", "--resource", "/iot/d", "--data", "@"})

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestIdasServicesUpdateErrorCbroker(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--apikey", "apikey", "--cbroker", "orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "specify url or broker alias to --cbroker", ngsiErr.Message)
	}
}
func TestIdasServicesUpdateErrorMarshal(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--apikey", "apikey", "--cbroker", "http://orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasServicesUpdateErrorParam(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--resource", "/iot/d"})

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "configuration group field not found", ngsiErr.Message)
	}
}

func TestIdasServicesUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--apikey", "apikey", "--cbroker", "http://orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/service"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasServicesUpdateErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"services", "update", "--host", "iota", "--apikey", "apikey", "--cbroker", "http://orion:1026", "--type", "Event", "--resource", "/iot/d", "--token", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasServicesDelete(t *testing.T) {
	c := setupTest([]string{"services", "delete", "--host", "iota", "--apikey", "apikey", "--resource", "/iot/d"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasServicesDeleteDataNoAPIKey(t *testing.T) {
	c := setupTest([]string{"services", "delete", "--host", "iota", "--resource", "/iot/d"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasServicesDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"services", "delete", "--host", "iota", "--apikey", "apikey", "--resource", "/iot/d"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/service"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasServicesDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"services", "delete", "--host", "iota", "--apikey", "apikey", "--resource", "/iot/d"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasServicesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
