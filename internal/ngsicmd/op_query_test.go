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
AUTHORS OR COPYRIGHT HOv2ERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package ngsicmd

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestOpQuery(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "Sensor001\nSensor002\nSensor003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryCountZero(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestOpQuerylines(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--lines"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"Sensor001\",\"type\":\"Sensor\"}\n{\"id\":\"Sensor002\",\"type\":\"Sensor\"}\n{\"id\":\"Sensor003\",\"type\":\"Sensor\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryValues(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--values"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "Sensor001\nSensor002\nSensor003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryLinesValues(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--lines", "--values"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[10.148599472]\n[14.627960669]\n[-2.461631059]\n[-15.999248065]\n[-4.553473866]\n[1.147149609]\n[1.003624237]\n[11.747977585]\n[-4.264932072]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryPage(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v2/op/query"
	reqRes1.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"103"}}

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/v2/op/query"
	reqRes2.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"103"}}

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "Sensor001\nSensor002\nSensor003\nSensor001\nSensor002\nSensor003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryVerbose(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\": \"Sensor001\",\"type\":\"Sensor\"},{\"id\": \"Sensor002\",\"type\":\"Sensor\"},{\"id\": \"Sensor003\",\"type\":\"Sensor\"}]"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryVerbosePretty(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--verbose", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"Sensor001\",\n    \"type\": \"Sensor\"\n  },\n  {\n    \"id\": \"Sensor002\",\n    \"type\": \"Sensor\"\n  },\n  {\n    \"id\": \"Sensor003\",\n    \"type\": \"Sensor\"\n  }\n]"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryCount(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--verbose", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryErrorOnlyV2(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion-ld"})

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 0, ngsiErr.ErrNo)
		assert.Equal(t, "Only available on NGSIv2", ngsiErr.Message)
	}
}

func TestOpQueryErrorReadAll(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "@"})

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestOpQueryErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestOpQueryErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	}
}

func TestOpQueryErrorResultsCountCount(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestOpQueryErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestOpQueryVerboseErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type:"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'S' after object key (27) sor001\",\"type:\"Sensor\"},{\"id\":", ngsiErr.Message)
	}
}

func TestOpQueryErrorLinesValues(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--lines", "--values"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOpQueryErrorLinesValues2(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--lines", "--values"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOpQueryErrorLines(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--lines"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOpQueryErrorLines2(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--lines"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOpQueryErrorVerbosePretty(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}", "--verbose", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOpQueryErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"get", "entities", "--host", "orion", "--data", "{\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := opQuery(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 12, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
