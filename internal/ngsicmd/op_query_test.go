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

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestOpQuery(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "Sensor001\nSensor002\nSensor003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestOpQuerylines(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	setupFlagBool(set, "lines")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--lines", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"Sensor001\",\"type\":\"Sensor\"}\n{\"id\":\"Sensor002\",\"type\":\"Sensor\"}\n{\"id\":\"Sensor003\",\"type\":\"Sensor\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryValues(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	setupFlagBool(set, "values")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--values", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "Sensor001\nSensor002\nSensor003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryLinesValues(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	setupFlagBool(set, "lines,values")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--lines", "--values", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[10.148599472]\n[14.627960669]\n[-2.461631059]\n[-15.999248065]\n[-4.553473866]\n[1.147149609]\n[1.003624237]\n[11.747977585]\n[-4.264932072]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v2/op/query"
	reqRes1.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"103"}}
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/v2/op/query"
	reqRes2.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"103"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "Sensor001\nSensor002\nSensor003\nSensor001\nSensor002\nSensor003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	set.Bool("verbose", false, "dock")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\": \"Sensor001\",\"type\":\"Sensor\"},{\"id\": \"Sensor002\",\"type\":\"Sensor\"},{\"id\": \"Sensor003\",\"type\":\"Sensor\"}]"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryCount(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	set.Bool("count", false, "doc")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--count", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOpQueryInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "LD")

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data="})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorResultsCountCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	set.Bool("count", false, "doc")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--count", "--data={}"})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryVerboseErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data,safeString")
	set.Bool("verbose", false, "dock")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type:"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--verbose", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'S' after object key (27) sor001\",\"type:\"Sensor\"},{\"id\":", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorLinesValues(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	setupFlagBool(set, "lines,values")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--lines", "--values", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorLinesValues2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	setupFlagBool(set, "lines,values")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--lines", "--values", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 12, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorLines(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	setupFlagBool(set, "lines")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--lines", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{Jsonlib: j, DecodeErr: errors.New("json error")}
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 13, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorLines2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	setupFlagBool(set, "lines")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--lines", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 14, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpQueryErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,data")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/op/query"
	reqRes.ResBody = []byte(`[{"id": "Sensor001","type":"Sensor"},{"id": "Sensor002","type":"Sensor"},{"id": "Sensor003","type":"Sensor"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={\"entities\":[{\"idPattern\":\".*\",\"type\":\"Sensor\"}]}"})
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	err := opQuery(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 15, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
