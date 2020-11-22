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

func TestRegistrationsListV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2CountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2Page(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"101"}}
	reqRes1.Path = "/v2/registrations"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"101"}}
	reqRes2.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2Verbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"inactive"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "verbose")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5fb9f88ca723657d763c631f Weather Context Source 2040-01-01T14:00:00.000Z\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2Localtime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"inactive"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "verbose,localTime")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--localTime"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5fb9f88ca723657d763c631f Weather Context Source 2040-01-01T23:00:00.000+0900\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2JSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "json")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--json"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T14:00:00.000Z\",\"status\":\"active\"}]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type []ngsicmd.registrationResposeV2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "json")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--json"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsGetV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T14:00:00.000Z\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual, client)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2LocalTime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")
	setupFlagBool(set, "localTime")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3", "--localTime"})
	err := registrationsGetV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T23:00:00.000+0900\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual, client)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2SafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,id,safeString")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5fb9f88ca723657d763c631f","description":"%25Weather Context Source","dataProvided":{"entities":[{"id":"urn:ngsi-ld:WeatherObserved:sensor003","type":"WeatherObserved"}],"attrs":["temperature","relativeHumidity","atmosphericPressure"]},"provider":{"http":{"url":"http://192.168.1.3/v1/weatherObserved"},"supportedForwardingMode":"all","legacyForwarding":true},"expires":"2040-01-01T14:00:00.000Z","status":"active"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	err := registrationsGetV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5fb9f88ca723657d763c631f\",\"description\":\"%Weather Context Source\",\"dataProvided\":{\"entities\":[{\"id\":\"urn:ngsi-ld:WeatherObserved:sensor003\",\"type\":\"WeatherObserved\"}],\"attrs\":[\"temperature\",\"relativeHumidity\",\"atmosphericPressure\"]},\"provider\":{\"http\":{\"url\":\"http://192.168.1.3/v1/weatherObserved\"},\"supportedForwardingMode\":\"all\",\"legacyForwarding\":true},\"expires\":\"2040-01-01T14:00:00.000Z\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual, client)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetErrorV2StatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2ErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,id,safeString")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	err := registrationsGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2ErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,id,safeString")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	err := registrationsGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/registrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2ErrorSetValule(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--data=aaa"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'a' looking for beginning of value (1) aaa", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2ErrorJSONMarshalEncode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsDeleteV2(c, ngsi, client)

	assert.NoError(t, err)
}

func TestRegistrationsDeleteV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsDeleteV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsDeleteV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsTemplateV2(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateV2Args(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "description,providedId,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--providedId=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	err := registrationsTemplateV2(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"description\":\"test\",\"dataProvided\":{\"entities\":[{\"id\":\"device001\",\"type\":\"Device\"}],\"attrs\":[\"abc\",\"xyz\"]},\"provider\":{\"http\":{\"url\":\"http://provider\"}}}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSetRegistrationsValuleV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "idPattern,forwardingModeFlag,provider,expires,status")
	setupFlagBool(set, "legacy")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=.*", "--legacy", "--provider=http://provider", "--forwardingModeFlag=all", "--expires=2040-01-01T14:00:00.000Z", "--status=active"})

	var r registrationQueryV2
	err := setRegistrationsValuleV2(c, ngsi, &r)

	assert.NoError(t, err)
}

func TestSetRegistrationsValuleV2Expires(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "idPattern,forwardingModeFlag,provider,expires,status")
	setupFlagBool(set, "legacy")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=.*", "--provider=http://provider", "--expires=1day"})

	var r registrationQueryV2
	err := setRegistrationsValuleV2(c, ngsi, &r)

	assert.NoError(t, err)
}

func TestRegistrationsTemplateV2ErrorProvider(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=provider"})
	err := registrationsTemplateV2(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "provider url error: provider", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateV2ErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := registrationsTemplateV2(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetRegistrationsValuleV2ErrorForwardingModeFlag(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "idPattern,forwardingModeFlag,provider,expires,status")
	setupFlagBool(set, "legacy")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--forwardingModeFlag=on", "--idPattern=.*", "--provider=http://provider"})

	var r registrationQueryV2
	err := setRegistrationsValuleV2(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "unknown mode: on", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetRegistrationsValuleV2ErrorExpires(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "idPattern,forwardingModeFlag,provider,expires,status")
	setupFlagBool(set, "legacy")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--expires=1", "--idPattern=.*", "--legacy", "--provider=http://provider"})

	var r registrationQueryV2
	err := setRegistrationsValuleV2(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetRegistrationsValuleV2ErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "idPattern,forwardingModeFlag,provider,expires,status")
	setupFlagBool(set, "legacy")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--status=on", "--idPattern=.*", "--legacy", "--provider=http://provider", "--forwardingModeFlag=all", "--expires=2040-01-01T14:00:00.000Z"})

	var r registrationQueryV2
	err := setRegistrationsValuleV2(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unknown status: on", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetRegistrationsValuleV2ErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data="})

	var r registrationQueryV2
	err := setRegistrationsValuleV2(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
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
