/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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

func TestGroupingrulesList(t *testing.T) {
	c := setupTest([]string{"groupingrules", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true","grouping_rules": []}`)

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"grouping_rules\": []}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesListPretty(t *testing.T) {
	c := setupTest([]string{"groupingrules", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true","grouping_rules": []}`)

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"grouping_rules\": []\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"groupingrules", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}
func TestGroupingrulesListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"groupingrules", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGroupingrulesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"groupingrules", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true","grouping_rules": []}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := groupingrulesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestGroupingrulesCreate(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"success":"true"}`
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesCreatePretty(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesCreateErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"groupingrules", "create", "--host", "cygnus", "--data", "@"})

	err := groupingrulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestGroupingrulesCreateErrorHTTP(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestGroupingrulesCreateErrorStatusCode(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGroupingrulesCreateErrorPretty(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := groupingrulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	}
}

func TestGroupingrulesUpdate(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "update", "--host", "cygnus", "--data", data, "--id", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesUpdatePretty(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "update", "--host", "cygnus", "--data", data, "--id", "1", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"groupingrules", "update", "--host", "cygnus", "--id", "1", "--data", "@"})

	err := groupingrulesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestGroupingrulesUpdateErrorHTTP(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "update", "--host", "cygnus", "--data", data, "--id", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestGroupingrulesUpdateErrorStatusCode(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "update", "--host", "cygnus", "--data", data, "--id", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGroupingrulesUpdateErrorPretty(t *testing.T) {
	data := `{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`
	c := setupTest([]string{"groupingrules", "update", "--host", "cygnus", "--data", data, "--id", "1", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ReqData = []byte(`{"regex":"Room","destination":"allrooms","fiware_service_path":"/rooms","fields":["entityType"]}`)
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := groupingrulesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestGroupingrulesDelete(t *testing.T) {
	c := setupTest([]string{"groupingrules", "delete", "--host", "cygnus", "--id", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"success":"true"}`
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesDeletePretty(t *testing.T) {
	c := setupTest([]string{"groupingrules", "delete", "--host", "cygnus", "--id", "1", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\"\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestGroupingrulesDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"groupingrules", "delete", "--host", "cygnus", "--id", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestGroupingrulesDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"groupingrules", "delete", "--host", "cygnus", "--id", "1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	err := groupingrulesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGroupingrulesDeleteErrorPretty(t *testing.T) {
	c := setupTest([]string{"groupingrules", "delete", "--host", "cygnus", "--id", "1", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/groupingrules"
	reqRes.ResBody = []byte(`{"success":"true"}`)
	id := "id=1"
	reqRes.RawQuery = &id

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := groupingrulesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
