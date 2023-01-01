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

package cygnus

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestLoggersList(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","loggers":[{"name":"org.mongodb","level":"WARN"},{"name":"org.apache.http","level":"WARN"},{"name":"org.apache.hadoop","level":"WARN"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := loggersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"loggers\":[{\"name\":\"org.mongodb\",\"level\":\"WARN\"},{\"name\":\"org.apache.http\",\"level\":\"WARN\"},{\"name\":\"org.apache.hadoop\",\"level\":\"WARN\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersListPretty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","loggers":[{"name":"org.mongodb","level":"WARN"},{"name":"org.apache.http","level":"WARN"},{"name":"org.apache.hadoop","level":"WARN"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := loggersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"loggers\": [\n    {\n      \"name\": \"org.mongodb\",\n      \"level\": \"WARN\"\n    },\n    {\n      \"name\": \"org.apache.http\",\n      \"level\": \"WARN\"\n    },\n    {\n      \"name\": \"org.apache.hadoop\",\n      \"level\": \"WARN\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := loggersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestLoggersListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := loggersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestLoggersListErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","loggers":[{"name":"org.mongodb","level":"WARN"},{"name":"org.apache.http","level":"WARN"},{"name":"org.apache.hadoop","level":"WARN"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := loggersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLoggersGet(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"WARN"}]"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"logger\":\"[{\"name\":\"org.mongodb\",\"level\":\"WARN\"}]\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersGetTransient(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb", "--transient"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"WARN"}]"}`)
	name := "name=org.mongodb&transient=true"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"logger\":\"[{\"name\":\"org.mongodb\",\"level\":\"WARN\"}]\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersGetPretty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"WARN"}]}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"logger\": [\n    {\n      \"name\": \"org.mongodb\",\n      \"level\": \"WARN\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"WARN"}]}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := loggersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestLoggersGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := loggersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestLoggersGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"WARN"}]}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := loggersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLoggersCreate(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"WARN"}}`
	c := setupTest([]string{"admin", "loggers", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)

	helper.SetClientHTTP(c, reqRes)

	err := loggersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":\"Logger 'org.mongodb' put\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersCreatePretty(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"WARN"}}`
	c := setupTest([]string{"admin", "loggers", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)

	helper.SetClientHTTP(c, reqRes)

	err := loggersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": \"Logger 'org.mongodb' put\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersCreateErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "create", "--host", "cygnus", "--data", "@"})

	err := loggersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestLoggersCreateErrorHTTP(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"WARN"}}`
	c := setupTest([]string{"admin", "loggers", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := loggersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestLoggersCreateErrorStatusCode(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"WARN"}}`
	c := setupTest([]string{"admin", "loggers", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := loggersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestLoggersCreateErrorPretty(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"WARN"}}`
	c := setupTest([]string{"admin", "loggers", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"WARN"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Logger 'org.mongodb' put"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := loggersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLoggersUpdate(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"INFO"}}`
	c := setupTest([]string{"admin", "loggers", "update", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"logger\":\"[{\"name\":\"org.mongodb\",\"level\":\"INFO\"}]\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersUpdatePretty(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"INFO"}}`
	c := setupTest([]string{"admin", "loggers", "update", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":[{"name":"org.mongodb","level":"INFO"}]}`)
	name := "transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"logger\": [\n    {\n      \"name\": \"org.mongodb\",\n      \"level\": \"INFO\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "update", "--host", "cygnus", "--data", "@"})

	err := loggersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestLoggersUpdateErrorHTTP(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"INFO"}}`
	c := setupTest([]string{"admin", "loggers", "update", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := loggersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestLoggersUpdateErrorStatusCode(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"INFO"}}`
	c := setupTest([]string{"admin", "loggers", "update", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestLoggersUpdateErrorPretty(t *testing.T) {
	data := `{"logger":{"name":"org.mongodb","level":"INFO"}}`
	c := setupTest([]string{"admin", "loggers", "update", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ReqData = []byte(`{"logger":{"name":"org.mongodb","level":"INFO"}}`)
	reqRes.ResBody = []byte(`{"success":"true","logger":"[{"name":"org.mongodb","level":"INFO"}]"}`)
	name := "transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := loggersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLoggersDelete(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "delete", "--host", "cygnus", "--name", "org.mongodb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed successfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":\" Logger 'org.mongodb' removed successfully\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersDeleteTransient(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "delete", "--host", "cygnus", "--name", "org.mongodb", "--transient"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed successfully"}`)
	name := "name=org.mongodb&transient=true"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":\" Logger 'org.mongodb' removed successfully\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersDeletePretty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "delete", "--host", "cygnus", "--name", "org.mongodb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed successfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": \" Logger 'org.mongodb' removed successfully\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestLoggersDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "delete", "--host", "cygnus", "--name", "org.mongodb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed successfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := loggersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestLoggersDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "delete", "--host", "cygnus", "--name", "org.mongodb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := loggersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestLoggersDeleteErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "delete", "--host", "cygnus", "--name", "org.mongodb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/loggers"
	reqRes.ResBody = []byte(`{"success":"true","result":" Logger 'org.mongodb' removed successfully"}`)
	name := "name=org.mongodb&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := loggersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestCygnusAdminSetParam(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb", "--transient"})

	v := cygnusAdminSetParam(c)

	assert.Equal(t, "org.mongodb", v.Get("name"))
	assert.Equal(t, "true", v.Get("transient"))
}

func TestCygnusAdminSetParam2(t *testing.T) {
	c := setupTest([]string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb"})

	v := cygnusAdminSetParam(c)

	assert.Equal(t, "org.mongodb", v.Get("name"))
	assert.Equal(t, "false", v.Get("transient"))
}
