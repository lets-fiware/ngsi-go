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

package convenience

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestAdminLog(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"level":"DEBUG"}`)
	reqRes.Path = "/admin/log"
	helper.SetClientHTTP(c, reqRes)

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"level":"DEBUG"}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminLogPretty(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion", "--logging", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"level":"DEBUG"}`)
	reqRes.Path = "/admin/log"
	helper.SetClientHTTP(c, reqRes)

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"level\": \"DEBUG\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminLogErrorHTTPLevelError(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion", "--level", "off"})

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "log level error: off (none, fatal, error, warn, info, debug)", ngsiErr.Message)
	}
}

func TestAdminLogErrorHTTPLevel(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion", "--level", "none"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminLogErrorStatusCodeLevel(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion", "--level", "none"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	helper.SetClientHTTP(c, reqRes)

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminLogErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminLogErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/log"
	helper.SetClientHTTP(c, reqRes)

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminLogErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "log", "--host", "orion", "--logging", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"level":"DEBUG"}`)
	reqRes.Path = "/admin/log"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := adminLog(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
func TestAdminTrace(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tracelevels":"empty"}`)
	reqRes.Path = "/log/trace"
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"tracelevels":"empty"}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminTraceSet(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion", "--set", "--logging", "--level", "t1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tracelevels":"poorly formatted trace level string"}`)
	reqRes.Path = "/log/trace/t1"
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"tracelevels":"poorly formatted trace level string"}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminTraceErrorLevel(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion", "--set"})

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing level", ngsiErr.Message)
	}
}

func TestAdminTraceErrorSetHTTPLevel(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion", "--set", "--level", "t1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminTraceErrorSetStatusCodeLevel(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion", "--set", "--level", "t1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace/t1"
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminTraceErrorDeleteHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion", "--delete", "--level", "t1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminTraceErrorDeleteStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion", "--delete", "--level", "t1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace/t1"
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminTraceErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminTraceErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "trace", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/log/trace"
	helper.SetClientHTTP(c, reqRes)

	err := adminTrace(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminSemaphore(t *testing.T) {
	c := setupTest([]string{"admin", "semaphore", "--host", "orion", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`)
	reqRes.Path = "/admin/sem"
	helper.SetClientHTTP(c, reqRes)

	err := adminSemaphore(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminSemaphorePretty(t *testing.T) {
	c := setupTest([]string{"admin", "semaphore", "--host", "orion", "--logging", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`)
	reqRes.Path = "/admin/sem"
	helper.SetClientHTTP(c, reqRes)

	err := adminSemaphore(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"dbConnectionPool\": {\n    \"status\": \"free\"\n  },\n  \"dbConnection\": {\n    \"status\": \"free\"\n  },\n  \"request\": {\n    \"status\": \"free\"\n  },\n  \"subCache\": {\n    \"status\": \"free\"\n  },\n  \"transaction\": {\n    \"status\": \"free\"\n  },\n  \"timeStat\": {\n    \"status\": \"free\"\n  },\n  \"logMsg\": {\n    \"status\": \"free\"\n  },\n  \"alarmMgr\": {\n    \"status\": \"free\"\n  },\n  \"metricsMgr\": {\n    \"status\": \"free\"\n  },\n  \"connectionContext\": {\n    \"status\": \"free\"\n  },\n  \"connectionEndpoints\": {\n    \"status\": \"free\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminSemaphoreErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "semaphore", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/sem"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminSemaphore(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminSemaphoreErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "semaphore", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/sem"
	helper.SetClientHTTP(c, reqRes)

	err := adminSemaphore(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminSemaphoreErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "semaphore", "--host", "orion", "--logging", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"dbConnectionPool":{"status":"free"},"dbConnection":{"status":"free"},"request":{"status":"free"},"subCache":{"status":"free"},"transaction":{"status":"free"},"timeStat":{"status":"free"},"logMsg":{"status":"free"},"alarmMgr":{"status":"free"},"metricsMgr":{"status":"free"},"connectionContext":{"status":"free"},"connectionEndpoints":{"status":"free"}}`)
	reqRes.Path = "/admin/sem"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := adminSemaphore(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAdminMetrics(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion", "--reset", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`)
	reqRes.Path = "/admin/metrics"
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminMetricsCygnus(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "cygnus", "--reset", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"services":{},"sum": {"subservs":{},"sum":{}}}`)
	reqRes.Path = "/v1/admin/metrics"
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"services":{},"sum": {"subservs":{},"sum":{}}}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminMetricsPretty(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion", "--reset", "--logging", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`)
	reqRes.Path = "/admin/metrics"
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"services\": {\n    \"default-service\": {\n      \"subservs\": {\n        \"root-subserv\": {\n          \"incomingTransactionResponseSize\": 1103,\n          \"serviceTime\": 0.000524,\n          \"incomingTransactions\": 7\n        }\n      },\n      \"sum\": {\n        \"incomingTransactionResponseSize\": 1103,\n        \"serviceTime\": 0.000524,\n        \"incomingTransactions\": 7\n      }\n    }\n  },\n  \"sum\": {\n    \"subservs\": {\n      \"root-subserv\": {\n        \"incomingTransactionResponseSize\": 1103,\n        \"serviceTime\": 0.000524,\n        \"incomingTransactions\": 7\n      }\n    },\n    \"sum\": {\n      \"incomingTransactionResponseSize\": 1103,\n      \"serviceTime\": 0.000524,\n      \"incomingTransactions\": 7\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminMetricsDelete(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte("")
	reqRes.Path = "/admin/metrics"
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestAdminMetricsErrorDeleteHTTPLevel(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminMetricsErrorDeleteStatusCodeLevel(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminMetricsErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminMetricsErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/admin/metrics"
	helper.SetClientHTTP(c, reqRes)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminMetricsErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "metrics", "--host", "orion", "--reset", "--logging", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"services":{"default-service":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}},"sum":{"subservs":{"root-subserv":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}},"sum":{"incomingTransactionResponseSize":1103,"serviceTime":0.000524,"incomingTransactions":7}}}`)
	reqRes.Path = "/admin/metrics"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := adminMetrics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAdminStatistics(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`)
	reqRes.Path = "/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminStatisticsCygnus(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "cygnus", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"success":"true","stats":{"sources":[{"name":"http-source-mongo","status":"START","setup_time":"2021-02-22T00:52:39.98Z","num_received_events":0,"num_processed_events":0}],"channels":[{"name":"mongo-channel","status":"START","setup_time":"2021-02-22T00:52:39.439Z","num_events":0,"num_puts_ok":0,"num_puts_failed":0,"num_takes_ok":0,"num_takes_failed":136},{"name":"sth-channel","status":"START","setup_time":"2021-02-22T00:52:39.436Z","num_events":0,"num_puts_ok":0,"num_puts_failed":0,"num_takes_ok":0,"num_takes_failed":136}],"sinks":[{"name":"mongo-sink","status":"START","setup_time":"2021-02-22T00:52:39.182Z","num_processed_events":0,"num_persisted_events":0},{"name":"sth-sink","status":"START","setup_time":"2021-02-22T00:52:39.196Z","num_processed_events":0,"num_persisted_events":0}]}}`)
	reqRes.Path = "/v1/stats"
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"success":"true","stats":{"sources":[{"name":"http-source-mongo","status":"START","setup_time":"2021-02-22T00:52:39.98Z","num_received_events":0,"num_processed_events":0}],"channels":[{"name":"mongo-channel","status":"START","setup_time":"2021-02-22T00:52:39.439Z","num_events":0,"num_puts_ok":0,"num_puts_failed":0,"num_takes_ok":0,"num_takes_failed":136},{"name":"sth-channel","status":"START","setup_time":"2021-02-22T00:52:39.436Z","num_events":0,"num_puts_ok":0,"num_puts_failed":0,"num_takes_ok":0,"num_takes_failed":136}],"sinks":[{"name":"mongo-sink","status":"START","setup_time":"2021-02-22T00:52:39.182Z","num_processed_events":0,"num_persisted_events":0},{"name":"sth-sink","status":"START","setup_time":"2021-02-22T00:52:39.196Z","num_processed_events":0,"num_persisted_events":0}]}}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminStatisticsCygnusDelete(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "cygnus", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"success":"true","stats":{"sources":[{"name":"http-source-mongo","status":"START","setup_time":"2021-02-22T00:52:39.98Z","num_received_events":0,"num_processed_events":0}],"channels":[{"name":"mongo-channel","status":"START","setup_time":"2021-02-22T00:52:39.439Z","num_events":0,"num_puts_ok":0,"num_puts_failed":0,"num_takes_ok":0,"num_takes_failed":136},{"name":"sth-channel","status":"START","setup_time":"2021-02-22T00:52:39.436Z","num_events":0,"num_puts_ok":0,"num_puts_failed":0,"num_takes_ok":0,"num_takes_failed":136}],"sinks":[{"name":"mongo-sink","status":"START","setup_time":"2021-02-22T00:52:39.182Z","num_processed_events":0,"num_persisted_events":0},{"name":"sth-sink","status":"START","setup_time":"2021-02-22T00:52:39.196Z","num_processed_events":0,"num_persisted_events":0}]}}`)
	reqRes.Path = "/v1/stats"
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAdminStatisticsPretty(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion", "--logging", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`)
	reqRes.Path = "/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"uptime_in_secs\": 152275,\n  \"measuring_interval_in_secs\": 152275\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminStatisticsDelete(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("")
	reqRes.Path = "/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestAdminStatisticsErrorDeleteHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminStatisticsErrorDeleteStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminStatisticsErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminStatisticsErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminStatisticsErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "statistics", "--host", "orion", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"uptime_in_secs":152275,"measuring_interval_in_secs":152275}`)
	reqRes.Path = "/statistics"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := adminStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAdminCacheStatistics(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`)
	reqRes.Path = "/cache/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminCacheStatisticsLogging(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion", "--logging"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`)
	reqRes.Path = "/cache/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`
		assert.Equal(t, expected, actual)
	}
}

func TestAdminCacheStatisticsPretty(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`)
	reqRes.Path = "/cache/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"ids\": \"\",\n  \"refresh\": 1949,\n  \"inserts\": 0,\n  \"removes\": 0,\n  \"updates\": 0,\n  \"items\": 0\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAdminCacheStatisticsDelete(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("")
	reqRes.Path = "/cache/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestAdminCacheStatisticsErrorDeleteHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminCacheStatisticsErrorDeleteStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion", "--delete"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminCacheStatisticsErrorHTTP(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAdminCacheStatisticsErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/cache/statistics"
	helper.SetClientHTTP(c, reqRes)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestAdminCacheStatisticsErrorPretty(t *testing.T) {
	c := setupTest([]string{"admin", "cacheStatistics", "--host", "orion", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"ids":"","refresh":1949,"inserts":0,"removes":0,"updates":0,"items":0}`)
	reqRes.Path = "/cache/statistics"
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := adminCacheStatistics(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
