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

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestTypesListV2(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes)

	err := typesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "AEDFacilities\nAirQualityObserved\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLD(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--link", "etsi"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "TemperatureSensor\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLDEmpty(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--link", "etsi"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte("{\n\"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n\"id\": \"urn:ngsi-ld:EntityTypeList:b6c79274-78c4-11eb-a948-0242ac12000f\",\n\"type\": \"EntityTypeList\",\n\"typeList\": []\n}")

	helper.SetClientHTTP(c, reqRes)

	err := typesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLDEmptyPretty(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--link", "etsi", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte("{\n\"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n\"id\": \"urn:ngsi-ld:EntityTypeList:b6c79274-78c4-11eb-a948-0242ac12000f\",\n\"type\": \"EntityTypeList\",\n\"typeList\": []\n}")

	helper.SetClientHTTP(c, reqRes)

	err := typesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n  \"id\": \"urn:ngsi-ld:EntityTypeList:b6c79274-78c4-11eb-a948-0242ac12000f\",\n  \"type\": \"EntityTypeList\",\n  \"typeList\": []\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListV2V2(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "AEDFacilities\nAirQualityObserved\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListV2CountZero(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListV2CountPage(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"12"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "AEDFacilities\nAirQualityObserved\nAEDFacilities\nAirQualityObserved\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListV2JSON(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\"AEDFacilities\",\"AirQualityObserved\"]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListV2Pretty(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  \"AEDFacilities\",\n  \"AirQualityObserved\"\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/type"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTypesListV2ErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestTypesListV2ErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestTypesListV2ErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}

	helper.SetClientHTTP(c, reqRes)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "EOF", ngsiErr.Message)
	}
}

func TestTypesListV2ErrorJSON(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTypesListV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"2"}}
	reqRes.ResBody = []byte(`["AEDFacilities","AirQualityObserved"]`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := typesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTypesListLDLD(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:419047fa-4ae6-11eb-b8c1-0242ac140003","type":"EntityTypeList","typeList":["https://uri.fiware.org/ns/data-models#TemperatureSensor"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "https://uri.fiware.org/ns/data-models#TemperatureSensor\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLDLink(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--link", "etsi"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "TemperatureSensor\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLDPretty(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--link", "etsi", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"@context\": \"http://context/ngsi-context.jsonld\",\n  \"id\": \"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003\",\n  \"type\": \"EntityTypeList\",\n  \"typeList\": [\n    \"TemperatureSensor\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLDJSON(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"@context\":\"http://context/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003\",\"type\":\"EntityTypeList\",\"typeList\":[\"TemperatureSensor\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLDDetails(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--details"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.RawQuery = helper.StrPtr("details=true")
	reqRes.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://uri.fiware.org/ns/data-models#TemperatureSensor","type":"EntityType","typeName":"https://uri.fiware.org/ns/data-models#TemperatureSensor","attributeNames":["https://uri.fiware.org/ns/data-models#category","https://w3id.org/saref#temperature"]}]`)

	helper.SetClientHTTP(c, reqRes)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\"type\":\"EntityType\",\"typeName\":\"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\"attributeNames\":[\"https://uri.fiware.org/ns/data-models#category\",\"https://w3id.org/saref#temperature\"]}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesListLDErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/type"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTypesListLDErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestTypesListLDErrorPretty(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTypesListLDErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"list", "types", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types"
	reqRes.ResBody = []byte(`{"@context":"http://context/ngsi-context.jsonld","id":"urn:ngsi-ld:EntityTypeList:b4d7fa50-4ae6-11eb-9f5b-0242ac140003","type":"EntityTypeList","typeList":["TemperatureSensor"]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := typesListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTypesGetV2(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion", "--type", "AirQualityObserved"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesGetLD(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion-ld", "--type", "TemperatureSensor"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types/TemperatureSensor"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://uri.fiware.org/ns/data-models#TemperatureSensor","type":"EntityTypeInformation","typeName":"https://uri.fiware.org/ns/data-models#TemperatureSensor","entityCount":3,"attributeDetails":[{"id":"https://uri.fiware.org/ns/data-models#category","type":"Attribute","attributeName":"https://uri.fiware.org/ns/data-models#category","attributeTypes":["Property"]},{"id":"https://w3id.org/saref#temperature","type":"Attribute","attributeName":"https://w3id.org/saref#temperature","attributeTypes":["Property"]}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\"type\":\"EntityTypeInformation\",\"typeName\":\"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\"entityCount\":3,\"attributeDetails\":[{\"id\":\"https://uri.fiware.org/ns/data-models#category\",\"type\":\"Attribute\",\"attributeName\":\"https://uri.fiware.org/ns/data-models#category\",\"attributeTypes\":[\"Property\"]},{\"id\":\"https://w3id.org/saref#temperature\",\"type\":\"Attribute\",\"attributeName\":\"https://w3id.org/saref#temperature\",\"attributeTypes\":[\"Property\"]}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesGetArg(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion", "AirQualityObserved"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesGetErrorArg(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion"})

	err := typeGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing entity type", ngsiErr.Message)
	}
}

func TestTypeGetV2(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion", "--type", "AirQualityObserved"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGetV2(c, c.Ngsi, c.Client, "AirQualityObserved")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypeGetV2Pretty(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion", "--type", "AirQualityObserved", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGetV2(c, c.Ngsi, c.Client, "AirQualityObserved")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"attrs\": {\n    \"CO\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"CO_Level\": {\n      \"types\": [\n        \"Text\"\n      ]\n    },\n    \"NO\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"NO2\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"NOx\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"SO2\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"address\": {\n      \"types\": [\n        \"StructuredValue\"\n      ]\n    },\n    \"airQualityIndex\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"airQualityLevel\": {\n      \"types\": [\n        \"Text\"\n      ]\n    },\n    \"dateObserved\": {\n      \"types\": [\n        \"DateTime\",\n        \"Text\"\n      ]\n    },\n    \"location\": {\n      \"types\": [\n        \"StructuredValue\",\n        \"geo:json\"\n      ]\n    },\n    \"precipitation\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"refPointOfInterest\": {\n      \"types\": [\n        \"Text\"\n      ]\n    },\n    \"relativeHumidity\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"reliability\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"source\": {\n      \"types\": [\n        \"Text\",\n        \"URL\"\n      ]\n    },\n    \"temperature\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"windDirection\": {\n      \"types\": [\n        \"Number\"\n      ]\n    },\n    \"windSpeed\": {\n      \"types\": [\n        \"Number\"\n      ]\n    }\n  },\n  \"count\": 18\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesGetV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion", "--type", "AirQualityObserved"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/error"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := typeGetV2(c, c.Ngsi, c.Client, "AirQualityObserved")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTypesGetV2ErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion", "--type", "AirQualityObserved"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGetV2(c, c.Ngsi, c.Client, "AirQualityObserved")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestTypesGetV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion", "--type", "AirQualityObserved", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/AirQualityObserved"
	reqRes.ResBody = []byte(`{"attrs":{"CO":{"types":["Number"]},"CO_Level":{"types":["Text"]},"NO":{"types":["Number"]},"NO2":{"types":["Number"]},"NOx":{"types":["Number"]},"SO2":{"types":["Number"]},"address":{"types":["StructuredValue"]},"airQualityIndex":{"types":["Number"]},"airQualityLevel":{"types":["Text"]},"dateObserved":{"types":["DateTime","Text"]},"location":{"types":["StructuredValue","geo:json"]},"precipitation":{"types":["Number"]},"refPointOfInterest":{"types":["Text"]},"relativeHumidity":{"types":["Number"]},"reliability":{"types":["Number"]},"source":{"types":["Text","URL"]},"temperature":{"types":["Number"]},"windDirection":{"types":["Number"]},"windSpeed":{"types":["Number"]}},"count":18}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := typeGetV2(c, c.Ngsi, c.Client, "AirQualityObserved")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTypeGetLD(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion-ld", "--type", "TemperatureSensor"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types/TemperatureSensor"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://uri.fiware.org/ns/data-models#TemperatureSensor","type":"EntityTypeInformation","typeName":"https://uri.fiware.org/ns/data-models#TemperatureSensor","entityCount":3,"attributeDetails":[{"id":"https://uri.fiware.org/ns/data-models#category","type":"Attribute","attributeName":"https://uri.fiware.org/ns/data-models#category","attributeTypes":["Property"]},{"id":"https://w3id.org/saref#temperature","type":"Attribute","attributeName":"https://w3id.org/saref#temperature","attributeTypes":["Property"]}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGetLd(c, c.Ngsi, c.Client, "TemperatureSensor")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\"type\":\"EntityTypeInformation\",\"typeName\":\"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\"entityCount\":3,\"attributeDetails\":[{\"id\":\"https://uri.fiware.org/ns/data-models#category\",\"type\":\"Attribute\",\"attributeName\":\"https://uri.fiware.org/ns/data-models#category\",\"attributeTypes\":[\"Property\"]},{\"id\":\"https://w3id.org/saref#temperature\",\"type\":\"Attribute\",\"attributeName\":\"https://w3id.org/saref#temperature\",\"attributeTypes\":[\"Property\"]}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypeGetLDPretty(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion-ld", "--type", "TemperatureSensor", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types/TemperatureSensor"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://uri.fiware.org/ns/data-models#TemperatureSensor","type":"EntityTypeInformation","typeName":"https://uri.fiware.org/ns/data-models#TemperatureSensor","entityCount":3,"attributeDetails":[{"id":"https://uri.fiware.org/ns/data-models#category","type":"Attribute","attributeName":"https://uri.fiware.org/ns/data-models#category","attributeTypes":["Property"]},{"id":"https://w3id.org/saref#temperature","type":"Attribute","attributeName":"https://w3id.org/saref#temperature","attributeTypes":["Property"]}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := typeGetLd(c, c.Ngsi, c.Client, "TemperatureSensor")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n  \"id\": \"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\n  \"type\": \"EntityTypeInformation\",\n  \"typeName\": \"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\n  \"entityCount\": 3,\n  \"attributeDetails\": [\n    {\n      \"id\": \"https://uri.fiware.org/ns/data-models#category\",\n      \"type\": \"Attribute\",\n      \"attributeName\": \"https://uri.fiware.org/ns/data-models#category\",\n      \"attributeTypes\": [\n        \"Property\"\n      ]\n    },\n    {\n      \"id\": \"https://w3id.org/saref#temperature\",\n      \"type\": \"Attribute\",\n      \"attributeName\": \"https://w3id.org/saref#temperature\",\n      \"attributeTypes\": [\n        \"Property\"\n      ]\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypeGetLDErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion-ld", "--type", "TemperatureSensor"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v2/types/TemperatureSensor"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := typeGetLd(c, c.Ngsi, c.Client, "TemperatureSensor")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTypeGetLDErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion-ld", "--type", "TemperatureSensor"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/types/TemperatureSensor"
	reqRes.ResBody = []byte(`bad request`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := typeGetLd(c, c.Ngsi, c.Client, "TemperatureSensor")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestTypeGetLDErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "type", "--host", "orion-ld", "--type", "TemperatureSensor", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/types/TemperatureSensor"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://uri.fiware.org/ns/data-models#TemperatureSensor","type":"EntityTypeInformation","typeName":"https://uri.fiware.org/ns/data-models#TemperatureSensor","entityCount":3,"attributeDetails":[{"id":"https://uri.fiware.org/ns/data-models#category","type":"Attribute","attributeName":"https://uri.fiware.org/ns/data-models#category","attributeTypes":["Property"]},{"id":"https://w3id.org/saref#temperature","type":"Attribute","attributeName":"https://w3id.org/saref#temperature","attributeTypes":["Property"]}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := typeGetLd(c, c.Ngsi, c.Client, "TemperatureSensor")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTypesCountV2(t *testing.T) {
	c := setupTest([]string{"wc", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"10"}}

	helper.SetClientHTTP(c, reqRes)

	err := typesCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "10\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTypesCountErrorOnlyV2(t *testing.T) {
	c := setupTest([]string{"wc", "types", "--host", "orion-ld"})

	err := typesCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Only available on NGSIv2", ngsiErr.Message)
	}
}

func TestTypesCountErrorHTTP(t *testing.T) {
	c := setupTest([]string{"wc", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/error"
	reqRes.ResBody = []byte(`error`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := typesCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestTypesCountErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"wc", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := typesCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  error", ngsiErr.Message)
	}
}

func TestTypesErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"wc", "types", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types"

	helper.SetClientHTTP(c, reqRes)

	err := typesCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}
