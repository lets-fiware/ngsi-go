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

func TestAppendersList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","appenders":[{"name":"DAILY","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"},{"name":"LOGFILE","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"true"},{"name":"console","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := appendersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"appenders\":[{\"name\":\"DAILY\",\"layout\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\"active\":\"false\"},{\"name\":\"LOGFILE\",\"layout\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\"active\":\"true\"},{\"name\":\"console\",\"layout\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\"active\":\"false\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","appenders":[{"name":"DAILY","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"},{"name":"LOGFILE","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"true"},{"name":"console","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	err := appendersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"appenders\": [\n    {\n      \"name\": \"DAILY\",\n      \"layout\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\n      \"active\": \"false\"\n    },\n    {\n      \"name\": \"LOGFILE\",\n      \"layout\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\n      \"active\": \"true\"\n    },\n    {\n      \"name\": \"console\",\n      \"layout\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\",\n      \"active\": \"false\"\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appendersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appendersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := appendersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAppendersListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := appendersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","appenders":[{"name":"DAILY","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"},{"name":"LOGFILE","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"true"},{"name":"console","layout":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n","active":"false"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	setJSONIndentError(ngsi)

	err := appendersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppendersGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := appendersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"appender\":{\"name\":\"test\",\"class\":\"\"},\"pattern\":{\"layout\":\"\",\"ConversionPattern\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersGetTransient(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=true"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "transient")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", "--transient"})

	err := appendersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"appender\":{\"name\":\"test\",\"class\":\"\"},\"pattern\":{\"layout\":\"\",\"ConversionPattern\":\"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--name=test"})

	err := appendersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"appender\": {\n    \"name\": \"test\",\n    \"class\": \"\"\n  },\n  \"pattern\": {\n    \"layout\": \"\",\n    \"ConversionPattern\": \"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n\"\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appendersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appendersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersGetErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := appendersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify appender name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := appendersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAppendersGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := appendersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--name=test"})

	setJSONIndentError(ngsi)

	err := appendersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppendersCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":\"Appender 'test' posted\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": \"Appender 'test' posted\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appendersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appendersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	setJSONEncodeErr(ngsi, 2)

	err := appendersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersCreateErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--data="})

	err := appendersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' posted"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	setJSONIndentError(ngsi)

	err := appendersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":\"Appender 'test' put\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", "-pretty", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": \"Appender 'test' put\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppendersUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=name", "--data="})

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ReqData = []byte(`{"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`)
	reqRes.ResBody = []byte(`{"success":"true","result":"Appender 'test' put"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", "--name=test", `--data={"appender":{"name":"test","class":""},"pattern":{"layout":"","ConversionPattern":"time=%d{yyyy-MM-dd}T%d{HH:mm:ss.SSS}Z|lvl=%p|corr=%X{correlatorId}|trans=%X{transactionId}|srv=%X{service}|subsrv=%X{subservice}|comp=%X{agent}|op=%M|msg=%C[%L]:%m%n"}}`})

	setJSONIndentError(ngsi)

	err := appendersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersDelete(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := appendersDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":\" Appender 'test' removed successfully\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeleteTransient(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=true"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "transient")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", "--transient"})

	err := appendersDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":\" Appender 'test' removed successfully\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeletePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", "--pretty"})

	err := appendersDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": \" Appender 'test' removed successfully\"\n}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appendersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appendersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeleteErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := appendersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := appendersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test"})

	err := appendersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppendersDeleteErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/admin/log/appenders"
	reqRes.ResBody = []byte(`{"success":"true","result":" Appender 'test' removed successfully"}`)
	name := "name=test&transient=false"
	reqRes.RawQuery = &name
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--name=test", "--pretty"})

	setJSONIndentError(ngsi)

	err := appendersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
