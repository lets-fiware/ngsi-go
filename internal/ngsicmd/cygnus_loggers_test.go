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

func TestLoggersList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","loggers":[{"name":"org.mongodb","level":"WARN"},{"name":"org.apache.http","level":"WARN"},{"name":"org.apache.hadoop","level":"WARN"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := loggersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"loggers\":[{\"name\":\"org.mongodb\",\"level\":\"WARN\"},{\"name\":\"org.apache.http\",\"level\":\"WARN\"},{\"name\":\"org.apache.hadoop\",\"level\":\"WARN\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","loggers":[{"name":"org.mongodb","level":"WARN"},{"name":"org.apache.http","level":"WARN"},{"name":"org.apache.hadoop","level":"WARN"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	err := loggersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"loggers\": [\n    {\n      \"name\": \"org.mongodb\",\n      \"level\": \"WARN\"\n    },\n    {\n      \"name\": \"org.apache.http\",\n      \"level\": \"WARN\"\n    },\n    {\n      \"name\": \"org.apache.hadoop\",\n      \"level\": \"WARN\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := loggersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := loggersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := loggersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestLoggersListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := loggersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","loggers":[{"name":"org.mongodb","level":"WARN"},{"name":"org.apache.http","level":"WARN"},{"name":"org.apache.hadoop","level":"WARN"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	setJSONIndentError(ngsi)

	err := loggersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLoggersGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"WARN"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb"})

	err := loggersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"logger\":\"[{\"name\":\"org.mongodb\",\"level\":\"WARN\"}]\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersGetTransient(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"WARN"}]"}`)
	name := "name=org.mongodb&transient=true"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "transient")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb", "--transient"})

	err := loggersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"logger\":\"[{\"name\":\"org.mongodb\",\"level\":\"WARN\"}]\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"WARN"}]}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--name=org.mongodb"})

	err := loggersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"logger\": [\n    {\n      \"name\": \"org.mongodb\",\n      \"level\": \"WARN\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := loggersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := loggersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersGetErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := loggersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify logger name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"WARN"}]}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb"})

	err := loggersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestLoggersGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb"})

	err := loggersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"WARN"}]}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--name=org.mongodb"})

	setJSONIndentError(ngsi)

	err := loggersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLoggersCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"logger":{"name":"org.mongodb","level":"WARN"}}`})

	err := loggersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":\"Logger 'org.mongodb' put\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"logger":{"name":"org.mongodb","level":"WARN"}}`})

	err := loggersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": \"Logger 'org.mongodb' put\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := loggersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := loggersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	setJSONEncodeErr(ngsi, 2)

	err := loggersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersCreateErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--data="})

	err := loggersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"logger":{"name":"org.mongodb","level":"WARN"}}`})

	err := loggersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"logger":{"name":"org.mongodb","level":"WARN"}}`})

	err := loggersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"logger":{"name":"org.mongodb","level":"WARN"}}`})

	setJSONIndentError(ngsi)

	err := loggersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb", `--data={"logger":{"name":"org.mongodb","level":"INFO"}}`})

	err := loggersUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"logger\":\"[{\"name\":\"org.mongodb\",\"level\":\"INFO\"}]\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"INFO"}]}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--name=org.mongodb", `--data={"logger":{"name":"org.mongodb","level":"INFO"}}`})

	err := loggersUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"logger\": [\n    {\n      \"name\": \"org.mongodb\",\n      \"level\": \"INFO\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=name", "--data="})

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb", `--data={"logger":{"name":"org.mongodb","level":"INFO"}}`})

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb", `--data={"logger":{"name":"org.mongodb","level":"INFO"}}`})

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--name=org.mongodb", `--data={"logger":{"name":"org.mongodb","level":"INFO"}}`})

	setJSONIndentError(ngsi)

	err := loggersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersDelete(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed succesfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb"})

	err := loggersDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":\" Logger 'org.mongodb' removed succesfully\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeleteTransient(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed succesfully"}`)
	name := "name=org.mongodb&transient=true"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "transient")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb", "--transient"})

	err := loggersDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":\" Logger 'org.mongodb' removed succesfully\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeletePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed succesfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb", "--pretty"})

	err := loggersDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": \" Logger 'org.mongodb' removed succesfully\"\n}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := loggersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := loggersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeleteErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed succesfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := loggersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed succesfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb"})

	err := loggersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb"})

	err := loggersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLoggersDeleteErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed succesfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=org.mongodb", "--pretty"})

	setJSONIndentError(ngsi)

	err := loggersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCygnusAdminSetParam(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "name")
	setupFlagBool(set, "transient")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=org.mongodb", "--transient"})

	v := cygnusAdminSetParam(c)

	assert.Equal(t, "org.mongodb", v.Get("name"))
	assert.Equal(t, "true", v.Get("transient"))
}

func TestCygnusAdminSetParam2(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "name")
	setupFlagBool(set, "transient")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=org.mongodb"})

	v := cygnusAdminSetParam(c)

	assert.Equal(t, "org.mongodb", v.Get("name"))
	assert.Equal(t, "false", v.Get("transient"))
}
