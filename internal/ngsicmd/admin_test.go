/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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

func TestAdminLog(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"level":"DEBUG"}`)
	reqRes.Path = "/admin/log"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging"})
	err := adminLog(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"level":"DEBUG"}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminLogPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"level":"DEBUG"}`)
	reqRes.Path = "/admin/log"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})
	err := adminLog(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"level\": \"DEBUG\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminLogErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAdminLogErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/admin/log"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminLogErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/admin/log"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminLogErrorHTTPLevelError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--level=off"})
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "log level error: off (none, fatal, error, warn, info, debug)", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminLogErrorHTTPLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--level=none"})
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminLogErrorStatusCodeLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--level=none"})
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminLogErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminLogErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminLogErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"level":"DEBUG"}`)
	reqRes.Path = "/admin/log"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})

	setJSONIndentError(ngsi)

	err := adminLog(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAdminTrace(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tracelevels":"empty"}`)
	reqRes.Path = "/log/trace"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging"})
	err := adminTrace(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"tracelevels":"empty"}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminTraceSet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tracelevels":"poorly formatted trace level string"}`)
	reqRes.Path = "/log/trace/t1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")
	setupFlagBool(set, "set,logging")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--set", "--level=t1", "--logging"})
	err := adminTrace(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"tracelevels":"poorly formatted trace level string"}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminTraceErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAdminTraceErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/log/trace"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/log/trace"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorParam(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/log/trace"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "set,delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--set", "--delete"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify either --set or --delete", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/log/trace"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "set")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--set"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "missing level", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorSetHTTPLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")
	setupFlagBool(set, "set")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--set", "--level=t1"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorSetStatusCodeLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace/t1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")
	setupFlagBool(set, "set")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--set", "--level=t1"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorDeleteHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete", "--level=t1"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorDeleteStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace/t1"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete", "--level=t1"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminTraceErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,level")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminTrace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetrics(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`)
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging,reset")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--reset", "--logging"})
	err := adminMetrics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminMetricsPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`)
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging,reset,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--reset", "--logging", "--pretty"})
	err := adminMetrics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"services\": {\n    \"default-service\": {\n      \"subservs\": {\n        \"root-subserv\": {\n          \"incomingTransactionResponseSize\": 1103,\n          \"serviceTime\": 0.000524,\n          \"incomingTransactions\": 7\n        }\n      },\n      \"sum\": {\n        \"incomingTransactionResponseSize\": 1103,\n        \"serviceTime\": 0.000524,\n        \"incomingTransactions\": 7\n      }\n    }\n  },\n  \"sum\": {\n    \"subservs\": {\n      \"root-subserv\": {\n        \"incomingTransactionResponseSize\": 1103,\n        \"serviceTime\": 0.000524,\n        \"incomingTransactions\": 7\n      }\n    },\n    \"sum\": {\n      \"incomingTransactionResponseSize\": 1103,\n      \"serviceTime\": 0.000524,\n      \"incomingTransactions\": 7\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminMetricsDelete(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte("")
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete,logging")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminMetrics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestAdminMetricsErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAdminMetricsErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetricsErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetricsErrorParam(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "reset,delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--reset", "--delete"})
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify either --reset or --delete", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetricsErrorDeleteHTTPLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetricsErrorDeleteStatusCodeLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetricsErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetricsErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminMetricsErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`)
	reqRes.Path = "/admin/metrics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "logging,reset,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--reset", "--logging", "--pretty"})

	setJSONIndentError(ngsi)

	err := adminMetrics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminSemaphore(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`)
	reqRes.Path = "/admin/sem"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging"})
	err := adminSemaphore(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminSemaphorePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`)
	reqRes.Path = "/admin/sem"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})
	err := adminSemaphore(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"dbConnectionPool\": {\n    \"status\": \"free\"\n  },\n  \"dbConnection\": {\n    \"status\": \"free\"\n  },\n  \"request\": {\n    \"status\": \"free\"\n  },\n  \"subCache\": {\n    \"status\": \"free\"\n  },\n  \"transaction\": {\n    \"status\": \"free\"\n  },\n  \"timeStat\": {\n    \"status\": \"free\"\n  },\n  \"logMsg\": {\n    \"status\": \"free\"\n  },\n  \"alarmMgr\": {\n    \"status\": \"free\"\n  },\n  \"metricsMgr\": {\n    \"status\": \"free\"\n  },\n  \"connectionContext\": {\n    \"status\": \"free\"\n  },\n  \"connectionEndpoints\": {\n    \"status\": \"free\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminSemaphoreErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := adminSemaphore(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAdminSemaphoreErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/admin/sem"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := adminSemaphore(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminSemaphoreErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/admin/sem"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := adminSemaphore(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminSemaphoreErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/sem"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminSemaphore(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminSemaphoreErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/sem"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminSemaphore(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminSemaphoreErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`)
	reqRes.Path = "/admin/sem"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})

	setJSONIndentError(ngsi)

	err := adminSemaphore(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminStatistics(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`)
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging"})
	err := adminStatistics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminStatisticsPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`)
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})
	err := adminStatistics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"uptime_in_secs\": 152275,\n  \"measuring_interval_in_secs\": 152275\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminStatisticsDelete(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("")
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminStatistics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestAdminStatisticsErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAdminStatisticsErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminStatisticsErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminStatisticsErrorDeleteHTTPLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminStatisticsErrorDeleteStatusCodeLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminStatisticsErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminStatisticsErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminStatisticsErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`)
	reqRes.Path = "/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})

	setJSONIndentError(ngsi)

	err := adminStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminCacheStatistics(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`)
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging"})
	err := adminCacheStatistics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := `{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminCacheStatisticsPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`)
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})
	err := adminCacheStatistics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"ids\": \"\",\n  \"refresh\": 1949,\n  \"inserts\": 0,\n  \"removes\": 0,\n  \"updates\": 0,\n  \"items\": 0\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminCacheStatisticsDelete(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("")
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminCacheStatistics(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestAdminCacheStatisticsErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAdminCacheStatisticsErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminCacheStatisticsErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminCacheStatisticsErrorDeleteHTTPLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminCacheStatisticsErrorDeleteStatusCodeLevel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "delete")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--delete"})
	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminCacheStatisticsErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminCacheStatisticsErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAdminCacheStatisticsErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`)
	reqRes.Path = "/cache/statistics"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "logging,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--logging", "--pretty"})

	setJSONIndentError(ngsi)

	err := adminCacheStatistics(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
