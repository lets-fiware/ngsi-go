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

func TestIdasServicesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasServicesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"count\":1,\"services\":[{\"commands\":[],\"lazy\":[],\"attributes\":[],\"_id\":\"601e25597d7b3d691be82d23\",\"resource\":\"/iot/d\",\"apikey\":\"apikey\",\"service\":\"openiot\",\"subservice\":\"/\",\"__v\":0,\"static_attributes\":[],\"internal_attributes\":[],\"entity_type\":\"Event\"}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasServicesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--pretty"})

	err := idasServicesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"count\": 1,\n  \"services\": [\n    {\n      \"commands\": [],\n      \"lazy\": [],\n      \"attributes\": [],\n      \"_id\": \"601e25597d7b3d691be82d23\",\n      \"resource\": \"/iot/d\",\n      \"apikey\": \"apikey\",\n      \"service\": \"openiot\",\n      \"subservice\": \"/\",\n      \"__v\": 0,\n      \"static_attributes\": [],\n      \"internal_attributes\": [],\n      \"entity_type\": \"Event\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasServicesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasServicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesListErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasServicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/service"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasServicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasServicesListErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasServicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasServicesListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"count":1,"services":[{"commands":[],"lazy":[],"attributes":[],"_id":"601e25597d7b3d691be82d23","resource":"/iot/d","apikey":"apikey","service":"openiot","subservice":"/","__v":0,"static_attributes":[],"internal_attributes":[],"entity_type":"Event"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--pretty"})

	setJSONIndentError(ngsi)

	err := idasServicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasServicesCreateData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data={"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`})

	err := idasServicesCreate(c)

	assert.NoError(t, err)
}

func TestIdasServicesCreateParam(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event","resource":"/iot/d"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=http://orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	err := idasServicesCreate(c)

	assert.NoError(t, err)
}

func TestIdasServicesCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesCreateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data=`})

	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesCreateErrorCbroker(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event","resource":"/iot/d"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify url or broker alias to --cbroker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesCreateErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event","resource":"/iot/d"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=http://orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	setJSONEncodeErr(ngsi, 2)

	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesCreateErrorParam(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,data,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--data=", "--apikey=apikey"})

	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "apikey, type and resource are needed", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/service"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data={"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`})
	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasServicesCreateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data={"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`})

	err := idasServicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasServicesUpdateData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--resource=/iot/d", `--data={"type":"Event"}`})

	err := idasServicesUpdate(c)

	assert.NoError(t, err)
}

func TestIdasServicesUpdateParam(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=http://orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	err := idasServicesUpdate(c)

	assert.NoError(t, err)
}

func TestIdasServicesUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesUpdateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesUpdateErrorResource(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--data="})

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "resource not fuond", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesUpdateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--resource=/iot/d", "--data="})

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesUpdateErrorCbroker(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify url or broker alias to --cbroker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestIdasServicesUpdateErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=http://orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	setJSONEncodeErr(ngsi, 2)

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesUpdateErrorParam(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,data,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--resource=/iot/d"})

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "configuration group field not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/service"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=http://orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasServicesUpdateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ReqData = []byte(`{"token":"FIWARE","cbroker":"http://orion:1026","entity_type":"Event"}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,apikey,cbroker,type,resource,token")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--cbroker=http://orion:1026", "--type=Event", "--resource=/iot/d", "--token=FIWARE"})

	err := idasServicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasServicesDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--resource=/iot/d"})

	err := idasServicesDelete(c)

	assert.NoError(t, err)
}

func TestIdasServicesDeleteDataNoAPIKey(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--resource=/iot/d"})

	err := idasServicesDelete(c)

	assert.NoError(t, err)
}

func TestIdasServicesDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasServicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesDeleteErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasServicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesDeleteErrorResource(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/services"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey"})

	err := idasServicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "resource not fuond", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasServicesDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/service"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--resource=/iot/d"})

	err := idasServicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasServicesDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/services"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,resource,apikey")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--apikey=apikey", "--resource=/iot/d"})

	err := idasServicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasDevicesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"count\":1,\"devices\":[{\"device_id\":\"sensor001\",\"service\":\"openiot\",\"service_path\":\"/\",\"entity_name\":\"urn:ngsi-ld:WeatherObserved:sensor001\",\"entity_type\":\"Sensor\",\"transport\":\"HTTP\",\"attributes\":[{\"object_id\":\"d\",\"name\":\"dateObserved\",\"type\":\"DateTime\"},{\"object_id\":\"t\",\"name\":\"temperature\",\"type\":\"Number\"},{\"object_id\":\"h\",\"name\":\"relativeHumidity\",\"type\":\"Number\"},{\"object_id\":\"p\",\"name\":\"atmosphericPressure\",\"type\":\"Number\"}],\"lazy\":[],\"commands\":[],\"static_attributes\":[{\"name\":\"location\",\"type\":\"geo:json\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.7671,35.68117]}}],\"explicitAttrs\":false}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--pretty"})

	err := idasDevicesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"count\": 1,\n  \"devices\": [\n    {\n      \"device_id\": \"sensor001\",\n      \"service\": \"openiot\",\n      \"service_path\": \"/\",\n      \"entity_name\": \"urn:ngsi-ld:WeatherObserved:sensor001\",\n      \"entity_type\": \"Sensor\",\n      \"transport\": \"HTTP\",\n      \"attributes\": [\n        {\n          \"object_id\": \"d\",\n          \"name\": \"dateObserved\",\n          \"type\": \"DateTime\"\n        },\n        {\n          \"object_id\": \"t\",\n          \"name\": \"temperature\",\n          \"type\": \"Number\"\n        },\n        {\n          \"object_id\": \"h\",\n          \"name\": \"relativeHumidity\",\n          \"type\": \"Number\"\n        },\n        {\n          \"object_id\": \"p\",\n          \"name\": \"atmosphericPressure\",\n          \"type\": \"Number\"\n        }\n      ],\n      \"lazy\": [],\n      \"commands\": [],\n      \"static_attributes\": [\n        {\n          \"name\": \"location\",\n          \"type\": \"geo:json\",\n          \"value\": {\n            \"type\": \"Point\",\n            \"coordinates\": [\n              139.7671,\n              35.68117\n            ]\n          }\n        }\n      ],\n      \"explicitAttrs\": false\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasDevicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesListErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasDevicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesListErrorDetailed(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,detailed")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--detailed=true"})

	err := idasDevicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify either on or off to --detailed", ngsiErr.Message)
	}
}

func TestIdasDevicesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasDevicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasDevicesListErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasDevicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--pretty"})

	setJSONIndentError(ngsi)

	err := idasDevicesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasDevicesGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001"})

	err := idasDevicesGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"device_id\":\"sensor001\",\"service\":\"openiot\",\"service_path\":\"/\",\"entity_name\":\"urn:ngsi-ld:WeatherObserved:sensor001\",\"entity_type\":\"Sensor\",\"transport\":\"HTTP\",\"attributes\":[{\"object_id\":\"d\",\"name\":\"dateObserved\",\"type\":\"DateTime\"},{\"object_id\":\"t\",\"name\":\"temperature\",\"type\":\"Number\"},{\"object_id\":\"h\",\"name\":\"relativeHumidity\",\"type\":\"Number\"},{\"object_id\":\"p\",\"name\":\"atmosphericPressure\",\"type\":\"Number\"}],\"lazy\":[],\"commands\":[],\"static_attributes\":[{\"name\":\"location\",\"type\":\"geo:json\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.7671,35.68117]}}],\"explicitAttrs\":false}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001", "--pretty"})

	err := idasDevicesGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"device_id\": \"sensor001\",\n  \"service\": \"openiot\",\n  \"service_path\": \"/\",\n  \"entity_name\": \"urn:ngsi-ld:WeatherObserved:sensor001\",\n  \"entity_type\": \"Sensor\",\n  \"transport\": \"HTTP\",\n  \"attributes\": [\n    {\n      \"object_id\": \"d\",\n      \"name\": \"dateObserved\",\n      \"type\": \"DateTime\"\n    },\n    {\n      \"object_id\": \"t\",\n      \"name\": \"temperature\",\n      \"type\": \"Number\"\n    },\n    {\n      \"object_id\": \"h\",\n      \"name\": \"relativeHumidity\",\n      \"type\": \"Number\"\n    },\n    {\n      \"object_id\": \"p\",\n      \"name\": \"atmosphericPressure\",\n      \"type\": \"Number\"\n    }\n  ],\n  \"lazy\": [],\n  \"commands\": [],\n  \"static_attributes\": [\n    {\n      \"name\": \"location\",\n      \"type\": \"geo:json\",\n      \"value\": {\n        \"type\": \"Point\",\n        \"coordinates\": [\n          139.7671,\n          35.68117\n        ]\n      }\n    }\n  ],\n  \"explicitAttrs\": false\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasDevicesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesGetErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasDevicesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesGetErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasDevicesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "device id not found", ngsiErr.Message)
	}
}

func TestIdasDevicesGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001"})

	err := idasDevicesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasDevicesGetErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001"})

	err := idasDevicesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001", "--pretty"})

	setJSONIndentError(ngsi)

	err := idasDevicesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasDevicesCreate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data={"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`})

	err := idasDevicesCreate(c)

	assert.NoError(t, err)
}

func TestIdasDevicesCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasDevicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesCreateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasDevicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasDevicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "--data not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesCreateErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data=`})

	err := idasDevicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data={"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`})
	err := idasDevicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasDevicesCreateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data={"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`})

	err := idasDevicesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001", `--data={"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`})

	err := idasDevicesUpdate(c)

	assert.NoError(t, err)
}

func TestIdasDevicesUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasDevicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesUpdateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasDevicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesUpdateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001"})

	err := idasDevicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "--data not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesUpdateErrorNoData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001", "--data="})

	err := idasDevicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesUpdateErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", `--data={"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`})

	err := idasDevicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "device id not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001", `--data={"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`})

	err := idasDevicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasDevicesUpdateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001", `--data={"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`})

	err := idasDevicesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices/sensor001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001"})

	err := idasDevicesDelete(c)

	assert.NoError(t, err)
}

func TestIdasDevicesDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := idasDevicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesDeleteErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := idasDevicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesDeleteErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices/sensor001"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota"})

	err := idasDevicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "device id not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIdasDevicesDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001"})

	err := idasDevicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestIdasDevicesDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=iota", "--id=sensor001"})

	err := idasDevicesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGetCbroker(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	ngsi, err := initCmd(c, "", false)
	assert.NoError(t, err)

	cases := []struct {
		arg      string
		expected string
	}{
		{arg: "http://orion:1026", expected: "http://orion:1026"},
		{arg: "orion", expected: "https://orion"},
		{arg: "orion-ld", expected: "https://orion-ld"},
		{arg: "orion-alias", expected: "https://orion-ld"},
	}

	for _, c := range cases {
		actual, err := getCbroker(ngsi, c.arg)

		if assert.NoError(t, err) {
			assert.Equal(t, c.expected, actual)
		} else {
			ngsiErr := err.(*ngsiCmdError)
			assert.Equal(t, 1, ngsiErr.ErrNo)
			assert.Equal(t, c.expected, ngsiErr.Message)
		}
	}
}

func TestGetCbrokerError(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	ngsi, err := initCmd(c, "", false)
	assert.NoError(t, err)

	cases := []struct {
		arg      string
		expected string
	}{
		{arg: "orion:1026", expected: "specify url or broker alias to --cbroker"},
		{arg: "http:/orion:1026", expected: "specify url or broker alias to --cbroker"},
	}

	for _, c := range cases {
		actual, err := getCbroker(ngsi, c.arg)

		if assert.Error(t, err) {
			ngsiErr := err.(*ngsiCmdError)
			assert.Equal(t, 1, ngsiErr.ErrNo)
			assert.Equal(t, "", actual)
			assert.Equal(t, c.expected, ngsiErr.Message)
		}
	}
}
