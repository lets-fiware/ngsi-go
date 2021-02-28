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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestTypesListV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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

func TestTypesListLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=etsi"})

	err := typesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "TemperatureSensor\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDEmpty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte("{\n\"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n\"id\": \"urn:ngsi-ld:EntityTypeList:b6c79274-78c4-11eb-a948-0242ac12000f\",\n\"type\": \"EntityTypeList\",\n\"typeList\": []\n}")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=etsi"})

	err := typesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDEmptyPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte("{\n\"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n\"id\": \"urn:ngsi-ld:EntityTypeList:b6c79274-78c4-11eb-a948-0242ac12000f\",\n\"type\": \"EntityTypeList\",\n\"typeList\": []\n}")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=etsi", "--pretty"})

	err := typesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n  \"id\": \"urn:ngsi-ld:EntityTypeList:b6c79274-78c4-11eb-a948-0242ac12000f\",\n  \"type\": \"EntityTypeList\",\n  \"typeList\": []\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := typesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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

//
// ngsi list types
//
func TestTypesListV2V2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

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

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--json"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\"AEDFacilities\",\"AirQualityObserved\"]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2Pretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  \"AEDFacilities\",\n  \"AirQualityObserved\"\n]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/type"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2ErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2ErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2ErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setJSONEncodeErr(ngsi, 2)

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--json"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListV2ErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = typesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:419047fa-4ae6-11eb-b8c1-0242ac140003","type":"EntityTypeList","typeList":["https://uri.fiware.org/ns/data-models#TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "https://uri.fiware.org/ns/data-models#TemperatureSensor\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDLink(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=etsi"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "TemperatureSensor\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=etsi", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"@context\": \"http://context/ngsi-context.jsonld\",\n  \"id\": \"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003\",\n  \"type\": \"EntityTypeList\",\n  \"typeList\": [\n    \"TemperatureSensor\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"@context\":\"http://context/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003\",\"type\":\"EntityTypeList\",\"typeList\":[\"TemperatureSensor\"]}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/type"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = typesListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=etsi", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = typesListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesListLDErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	setJSONDecodeErr(ngsi, 0)

	err = typesListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
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

func TestTypesGetV2Pretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved", "--pretty"})
	err := typeGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"attrs\": {\n    \"CO\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"CO_Level\": {\n      \"types\": [\n        \"Text\"\n      ]\n    },\n    \"NO\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"NO2\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"NOx\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"SO2\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"address\": {\n      \"types\": [\n        \"StructuredValue\"\n      ]\n    },\n    \"airQualityIndex\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"airQualityLevel\": {\n      \"types\": [\n        \"Text\"\n      ]\n    },\n    \"dateObserved\": {\n      \"types\": [\n        \"DateTime\",\n        \"Text\"\n      ]\n    },\n    \"location\": {\n      \"types\": [\n        \"StructuredValue\",\n        \"geo:json\"\n      ]\n    },\n    \"precipitation\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"refPointOfInterest\": {\n      \"types\": [\n        \"Text\"\n      ]\n    },\n    \"relativeHumidity\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"reliability\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"source\": {\n      \"types\": [\n        \"Text\",\n        \"URL\"\n      ]\n    },\n    \"temperature\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"windDirection\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"windSpeed\": {\n      \"types\": [\n        \"Number\"\n      ]\n    }\n  },\n  \"count\": 18\n}\n"
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
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTypesCountErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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

func TestTypesGetV2ErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved", "--pretty"})

	setJSONIndentError(ngsi)

	err := typeGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
