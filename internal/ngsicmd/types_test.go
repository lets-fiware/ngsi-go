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

//
// ngsi list types
//
func TestTypesListV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "AEDFacilities\nAirQualityObserved\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2CountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2CountPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"12"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "AEDFacilities\nAirQualityObserved\nAEDFacilities\nAirQualityObserved\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2JSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	set.Bool("json", false, "doc")
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--json"})
	err := typesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\"AEDFacilities\",\"AirQualityObserved\"]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

// initCmd() Error: no host
func TestTypesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/Path"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/type"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2ErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	set.Bool("json", false, "doc")
	setupFlagString(set, "host")
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--json"})
	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

///
// ngsi get type
//
// ngsi get --host orion type --type AirQualityObserved
func TestTypesGetV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := typeGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := typeGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc", "--type=AirQualityObserved"})
	err := typeGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}

}

func TestTypesGetErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=AirQualityObserved"})
	err := typeGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/error"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := typeGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := typeGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}

}

///
// ngsi wc types
//
// ngsi wc --host fisudalab types
func TestTypesCountV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"10"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "10\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesCountErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := typesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesCountErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := typesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesCountErrorOnlyV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := typesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesCountErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/error"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesCountErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := typesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestTypesErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := typesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
