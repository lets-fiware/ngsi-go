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

package cygnus

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestAppendersList(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","appenders":[{"name":"DAILY","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"},{"name":"LOGFILE","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"true"},{"name":"console","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appendersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"appenders\":[{\"name\":\"DAILY\",\"layout\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\"active\":\"false\"},{\"name\":\"LOGFILE\",\"layout\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\"active\":\"true\"},{\"name\":\"console\",\"layout\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\"active\":\"false\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersListPretty(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","appenders":[{"name":"DAILY","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"},{"name":"LOGFILE","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"true"},{"name":"console","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appendersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"appenders\": [\n    {\n      \"name\": \"DAILY\",\n      \"layout\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\n      \"active\": \"false\"\n    },\n    {\n      \"name\": \"LOGFILE\",\n      \"layout\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\n      \"active\": \"true\"\n    },\n    {\n      \"name\": \"console\",\n      \"layout\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\n      \"active\": \"false\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appendersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestAppendersListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appendersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppendersListErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","appenders":[{"name":"DAILY","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"},{"name":"LOGFILE","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"true"},{"name":"console","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appendersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppendersGet(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "get", "--host", "cygnus", "--name", "test"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"appender\":{\"name\":\"test\",\"class\":\"\"},\"pattern\":{\"layout\":\"\",\"ConversionPattern\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersGetTransient(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "get", "--host", "cygnus", "--name", "test", "--transient"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=true"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"appender\":{\"name\":\"test\",\"class\":\"\"},\"pattern\":{\"layout\":\"\",\"ConversionPattern\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersGetPretty(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "get", "--host", "cygnus", "--name", "test", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"appender\": {\n    \"name\": \"test\",\n    \"class\": \"\"\n  },\n  \"pattern\": {\n    \"layout\": \"\",\n    \"ConversionPattern\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\"\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "get", "--host", "cygnus", "--name", "test"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appendersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestAppendersGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "get", "--host", "cygnus", "--name", "test"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appendersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppendersGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "get", "--host", "cygnus", "--name", "test", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appendersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppendersCreate(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appendersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":\"Appender 'test' posted\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersCreatePretty(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appendersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": \"Appender 'test' posted\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersCreateErrorData(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "create", "--host", "cygnus", "--data", "@"})

	err := appendersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestAppendersCreateErrorHTTP(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appendersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppendersCreateErrorStatusCode(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appendersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppendersCreateErrorPretty(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appendersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppendersUpdate(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "update", "--host", "cygnus", "--name", "test", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":\"Appender 'test' put\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersUpdatePretty(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "update", "--host", "cygnus", "--name", "test", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": \"Appender 'test' put\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersErrorData(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "update", "--host", "cygnus", "--name", "test", "--data", "@"})

	err := appendersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestAppendersUpdateErrorHTTP(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "update", "--host", "cygnus", "--name", "test", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appendersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppendersUpdateErrorStatusCode(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "update", "--host", "cygnus", "--name", "test", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppendersUpdateErrorPretty(t *testing.T) {
	data := `{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`
	c := setupTest([]string{"admin", "appenders", "update", "--host", "cygnus", "--name", "test", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appendersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppendersDelete(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "delete", "--host", "cygnus", "--name", "test"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":\" Appender 'test' removed successfully\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersDeleteTransient(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "delete", "--host", "cygnus", "--name", "test", "--transient"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=true"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":\" Appender 'test' removed successfully\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersDeletePretty(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "delete", "--host", "cygnus", "--name", "test", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": \" Appender 'test' removed successfully\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "delete", "--host", "cygnus", "--name", "test", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appendersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppendersDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "delete", "--host", "cygnus", "--name", "test", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	err := appendersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppendersDeleteErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "appenders", "delete", "--host", "cygnus", "--name", "test", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appendersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
