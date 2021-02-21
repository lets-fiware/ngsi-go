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

func TestGroupingrulesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true","grouping_rules": []}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := groupingrulesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"grouping_rules\": []}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true","grouping_rules": []}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	err := groupingrulesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"grouping_rules\": []\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := groupingrulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := groupingrulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := groupingrulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestGroupingrulesListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := groupingrulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true","grouping_rules": []}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	setJSONIndentError(ngsi)

	err := groupingrulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestGroupingrulesCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"success":"true"}`
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := groupingrulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := groupingrulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	setJSONEncodeErr(ngsi, 2)

	err := groupingrulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesCreateErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--data="})

	err := groupingrulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	setJSONIndentError(ngsi)

	err := groupingrulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1", "--pretty", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1"})

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1", "--data="})

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--id=1", `--data={"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`})

	setJSONIndentError(ngsi)

	err := groupingrulesUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDelete(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1"})

	err := groupingrulesDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"success":"true"}`
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDeletePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1", "--pretty"})

	err := groupingrulesDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\"\n}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := groupingrulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := groupingrulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDeleteErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := groupingrulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1"})

	err := groupingrulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1"})

	err := groupingrulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGroupingrulesDeleteErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--id=1", "--pretty"})

	setJSONIndentError(ngsi)

	err := groupingrulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
