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

func TestRegistrationsListV2(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2CountZero(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2CountZeroPretty(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2Page(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"101"}}
	reqRes1.Path = "/v2/registrations"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"101"}}
	reqRes2.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3\n5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2Verbose(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"inactive"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5fb9f88ca723657d763c631f Weather Context Source 2040-01-01T14:00:00.000Z\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2Localtime(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion", "--verbose", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"inactive"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5fb9f88ca723657d763c631f Weather Context Source 2040-01-01T23:00:00.000+0900\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2JSON(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T14:00:00.000Z\",\"status\":\"active\"}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2JSONPretty(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"5fb9f88ca723657d763c631f\",\n    \"description\": \"Weather Context Source\",\n    \"dataProvided\": {\n      \"entities\": [\n        {\n          \"id\": \"urn:ngsi-ld:WeatherObserved:sensor003\",\n          \"type\": \"WeatherObserved\"\n        }\n      ],\n      \"attrs\": [\n        \"temperature\",\n        \"relativeHumidity\",\n        \"atmosphericPressure\"\n      ]\n    },\n    \"provider\": {\n      \"http\": {\n        \"url\": \"http://192.168.1.3/v1/weatherObserved\"\n      },\n      \"supportedForwardingMode\": \"all\",\n      \"legacyForwarding\": true\n    },\n    \"expires\": \"2040-01-01T14:00:00.000Z\",\n    \"status\": \"active\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsListV2ErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestRegistrationsListV2ErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestRegistrationsListV2ErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type []ngsicmd.registrationResposeV2 Field: (1) {}", ngsiErr.Message)
	}
}

func TestRegistrationsListV2ErrorJSON(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsListV2ErrorJSONPretty(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := registrationsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsGetV2(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T14:00:00.000Z\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsGetV2Pretty(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"5fb9f88ca723657d763c631f\",\n  \"description\": \"Weather Context Source\",\n  \"dataProvided\": {\n    \"entities\": [\n      {\n        \"id\": \"urn:ngsi-ld:WeatherObserved:sensor003\",\n        \"type\": \"WeatherObserved\"\n      }\n    ],\n    \"attrs\": [\n      \"temperature\",\n      \"relativeHumidity\",\n      \"atmosphericPressure\"\n    ]\n  },\n  \"provider\": {\n    \"http\": {\n      \"url\": \"http://192.168.1.3/v1/weatherObserved\"\n    },\n    \"supportedForwardingMode\": \"all\",\n    \"legacyForwarding\": true\n  },\n  \"expires\": \"2040-01-01T14:00:00.000Z\",\n  \"status\": \"active\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsGetV2LocalTime(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T23:00:00.000+0900\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsGetV2SafeString(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"%25Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"%Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T14:00:00.000Z\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsGetV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsGetErrorV2StatusCode(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "5f5dcb551e715bc7f1ad79e3  error", ngsiErr.Message)
	}
}

func TestRegistrationsGetV2ErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsGetV2ErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsGetV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := registrationsGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateV2(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/registrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/v2/registrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreateV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsCreateV2ErrorSetValule(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion", "--data", "aaa"})

	err := registrationsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'a' looking for beginning of value (1) aaa", ngsiErr.Message)
	}
}

func TestRegistrationsCreateV2ErrorJSONMarshalEncode(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion", "--data", "{}"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateV2ErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestRegistrationsDeleteV2(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDeleteV2(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRegistrationsDeleteV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDeleteV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsDeleteV2ErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDeleteV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "5f5dcb551e715bc7f1ad79e3  error", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateV2(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2"})

	err := registrationsTemplateV2(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsTemplateV2Args(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	err := registrationsTemplateV2(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"description\":\"test\",\"dataProvided\":{\"entities\":[{\"id\":\"device001\",\"type\":\"Device\"}],\"attrs\":[\"abc\",\"xyz\"]},\"provider\":{\"http\":{\"url\":\"http://provider\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsTemplateV2ArgsPretty(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--pretty", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	err := registrationsTemplateV2(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"description\": \"test\",\n  \"dataProvided\": {\n    \"entities\": [\n      {\n        \"id\": \"device001\",\n        \"type\": \"Device\"\n      }\n    ],\n    \"attrs\": [\n      \"abc\",\n      \"xyz\"\n    ]\n  },\n  \"provider\": {\n    \"http\": {\n      \"url\": \"http://provider\"\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleV2(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--idPattern", ".*", "--legacy", "--provider", "http://provider", "--forwardingMode", "all", "--expires", "2040-01-01T14:00:00.000Z", "--status", "active"})

	var r registrationQueryV2

	err := setRegistrationsValuleV2(c, c.Ngsi, &r)

	assert.NoError(t, err)
}

func TestSetRegistrationsValuleV2Expires(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--idPattern", ".*", "--provider", "http://provider", "--expires", "1day"})

	var r registrationQueryV2

	err := setRegistrationsValuleV2(c, c.Ngsi, &r)

	assert.NoError(t, err)
}

func TestRegistrationsTemplateV2ErrorProvider(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--description", "test", "--type", "Device", "--attrs", "abc,xyz", "--provider", "provider"})

	err := registrationsTemplateV2(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "provider url error: provider", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateV2ErrorMarshal(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--description", "test", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsTemplateV2(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateV2ErrorArgsPretty(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--pretty", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	helper.SetJSONIndentError(c.Ngsi)

	err := registrationsTemplateV2(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleV2ErrorForwardingMode(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--forwardingMode", "on", "--idPattern", ".*", "--provider", "http://provider"})

	var r registrationQueryV2

	err := setRegistrationsValuleV2(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "unknown mode: on", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleV2ErrorExpires(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--expires", "1", "--idPattern", ".*", "--legacy", "--provider", "http://provider"})

	var r registrationQueryV2

	err := setRegistrationsValuleV2(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleV2ErrorStatus(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--status", "on", "--idPattern", ".*", "--legacy", "--provider", "http://provider", "--forwardingMode", "all", "--expires", "2040-01-01T14:00:00.000Z"})

	var r registrationQueryV2

	err := setRegistrationsValuleV2(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unknown status: on", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleV2ErrorReadAll(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "v2", "--data", "@"})

	var r registrationQueryV2

	err := setRegistrationsValuleV2(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestToLocaltimeRegistrationNil(t *testing.T) {
	var r registrationResposeV2

	toLocaltimeRegistration(&r)

	assert.Equal(t, (*registrationForwardingInformationV2)(nil), r.ForwardingInformation)
}

func TestToLocaltimeRegistrationNil2(t *testing.T) {
	var r registrationResposeV2
	r.ForwardingInformation = new(registrationForwardingInformationV2)

	toLocaltimeRegistration(&r)

	assert.Equal(t, "", r.ForwardingInformation.LastForwarding)
	assert.Equal(t, "", r.ForwardingInformation.LastSuccess)
	assert.Equal(t, "", r.ForwardingInformation.LastFailure)
	assert.Equal(t, "", r.Expires)
}

func TestToLocaltimeRegistration(t *testing.T) {
	var r registrationResposeV2
	r.ForwardingInformation = new(registrationForwardingInformationV2)

	r.ForwardingInformation.LastForwarding = "2040-01-01T14:00:00.000Z"
	r.ForwardingInformation.LastSuccess = "2040-01-02T14:00:00.000Z"
	r.ForwardingInformation.LastFailure = "2040-01-03T14:00:00.000Z"
	r.Expires = "2040-01-04T14:00:00.000Z"

	toLocaltimeRegistration(&r)

	assert.Equal(t, "2040-01-01T23:00:00.000+0900", r.ForwardingInformation.LastForwarding)
	assert.Equal(t, "2040-01-02T23:00:00.000+0900", r.ForwardingInformation.LastSuccess)
	assert.Equal(t, "2040-01-03T23:00:00.000+0900", r.ForwardingInformation.LastFailure)
	assert.Equal(t, "2040-01-04T23:00:00.000+0900", r.Expires)
}
