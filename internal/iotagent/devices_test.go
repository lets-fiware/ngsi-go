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

package iotagent

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestIdasDevicesList(t *testing.T) {
	c := setupTest([]string{"devices", "list", "--host", "iota"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"count\":1,\"devices\":[{\"device_id\":\"sensor001\",\"service\":\"openiot\",\"service_path\":\"/\",\"entity_name\":\"urn:ngsi-ld:WeatherObserved:sensor001\",\"entity_type\":\"Sensor\",\"transport\":\"HTTP\",\"attributes\":[{\"object_id\":\"d\",\"name\":\"dateObserved\",\"type\":\"DateTime\"},{\"object_id\":\"t\",\"name\":\"temperature\",\"type\":\"Number\"},{\"object_id\":\"h\",\"name\":\"relativeHumidity\",\"type\":\"Number\"},{\"object_id\":\"p\",\"name\":\"atmosphericPressure\",\"type\":\"Number\"}],\"lazy\":[],\"commands\":[],\"static_attributes\":[{\"name\":\"location\",\"type\":\"geo:json\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.7671,35.68117]}}],\"explicitAttrs\":false}]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesListPretty(t *testing.T) {
	c := setupTest([]string{"devices", "list", "--host", "iota", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"count\": 1,\n  \"devices\": [\n    {\n      \"device_id\": \"sensor001\",\n      \"service\": \"openiot\",\n      \"service_path\": \"/\",\n      \"entity_name\": \"urn:ngsi-ld:WeatherObserved:sensor001\",\n      \"entity_type\": \"Sensor\",\n      \"transport\": \"HTTP\",\n      \"attributes\": [\n        {\n          \"object_id\": \"d\",\n          \"name\": \"dateObserved\",\n          \"type\": \"DateTime\"\n        },\n        {\n          \"object_id\": \"t\",\n          \"name\": \"temperature\",\n          \"type\": \"Number\"\n        },\n        {\n          \"object_id\": \"h\",\n          \"name\": \"relativeHumidity\",\n          \"type\": \"Number\"\n        },\n        {\n          \"object_id\": \"p\",\n          \"name\": \"atmosphericPressure\",\n          \"type\": \"Number\"\n        }\n      ],\n      \"lazy\": [],\n      \"commands\": [],\n      \"static_attributes\": [\n        {\n          \"name\": \"location\",\n          \"type\": \"geo:json\",\n          \"value\": {\n            \"type\": \"Point\",\n            \"coordinates\": [\n              139.7671,\n              35.68117\n            ]\n          }\n        }\n      ],\n      \"explicitAttrs\": false\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"devices", "list", "--host", "iota"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasDevicesListErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"devices", "list", "--host", "iota"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"devices", "list", "--host", "iota", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"count":1,"devices":[{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := idasDevicesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasDevicesGet(t *testing.T) {
	c := setupTest([]string{"devices", "get", "--host", "iota", "--id", "sensor001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"device_id\":\"sensor001\",\"service\":\"openiot\",\"service_path\":\"/\",\"entity_name\":\"urn:ngsi-ld:WeatherObserved:sensor001\",\"entity_type\":\"Sensor\",\"transport\":\"HTTP\",\"attributes\":[{\"object_id\":\"d\",\"name\":\"dateObserved\",\"type\":\"DateTime\"},{\"object_id\":\"t\",\"name\":\"temperature\",\"type\":\"Number\"},{\"object_id\":\"h\",\"name\":\"relativeHumidity\",\"type\":\"Number\"},{\"object_id\":\"p\",\"name\":\"atmosphericPressure\",\"type\":\"Number\"}],\"lazy\":[],\"commands\":[],\"static_attributes\":[{\"name\":\"location\",\"type\":\"geo:json\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.7671,35.68117]}}],\"explicitAttrs\":false}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesGetPretty(t *testing.T) {
	c := setupTest([]string{"devices", "get", "--host", "iota", "--id", "sensor001", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"device_id\": \"sensor001\",\n  \"service\": \"openiot\",\n  \"service_path\": \"/\",\n  \"entity_name\": \"urn:ngsi-ld:WeatherObserved:sensor001\",\n  \"entity_type\": \"Sensor\",\n  \"transport\": \"HTTP\",\n  \"attributes\": [\n    {\n      \"object_id\": \"d\",\n      \"name\": \"dateObserved\",\n      \"type\": \"DateTime\"\n    },\n    {\n      \"object_id\": \"t\",\n      \"name\": \"temperature\",\n      \"type\": \"Number\"\n    },\n    {\n      \"object_id\": \"h\",\n      \"name\": \"relativeHumidity\",\n      \"type\": \"Number\"\n    },\n    {\n      \"object_id\": \"p\",\n      \"name\": \"atmosphericPressure\",\n      \"type\": \"Number\"\n    }\n  ],\n  \"lazy\": [],\n  \"commands\": [],\n  \"static_attributes\": [\n    {\n      \"name\": \"location\",\n      \"type\": \"geo:json\",\n      \"value\": {\n        \"type\": \"Point\",\n        \"coordinates\": [\n          139.7671,\n          35.68117\n        ]\n      }\n    }\n  ],\n  \"explicitAttrs\": false\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestIdasDevicesGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"devices", "get", "--host", "iota", "--id", "sensor001", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasDevicesGetErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"devices", "get", "--host", "iota", "--id", "sensor001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"devices", "get", "--host", "iota", "--id", "sensor001", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"device_id":"sensor001","service":"openiot","service_path":"/","entity_name":"urn:ngsi-ld:WeatherObserved:sensor001","entity_type":"Sensor","transport":"HTTP","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"lazy":[],"commands":[],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}],"explicitAttrs":false}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := idasDevicesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestIdasDevicesCreate(t *testing.T) {
	data := `{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`
	c := setupTest([]string{"devices", "create", "--host", "iota", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasDevicesCreateErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"devices", "create", "--host", "iota", "--data", "@"})

	err := idasDevicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestIdasDevicesCreateErrorHTTP(t *testing.T) {
	data := `{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`
	c := setupTest([]string{"devices", "create", "--host", "iota", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasDevicesCreateErrorHTTPStatus(t *testing.T) {
	data := `{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`
	c := setupTest([]string{"devices", "create", "--host", "iota", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"devices":[{"device_id":"${device1}","entity_name":"urn:ngsi-ld:WeatherObserved:${device1}","entity_type":"Sensor","timezone":"Asia/Tokyo","attributes":[{"object_id":"d","name":"dateObserved","type":"DateTime"},{"object_id":"t","name":"temperature","type":"Number"},{"object_id":"h","name":"relativeHumidity","type":"Number"},{"object_id":"p","name":"atmosphericPressure","type":"Number"}],"static_attributes":[{"name":"location","type":"geo:json","value":{"type":"Point","coordinates":[139.7671,35.68117]}}]}]`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesUpdate(t *testing.T) {
	data := `{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`
	c := setupTest([]string{"devices", "update", "--host", "iota", "--id", "sensor001", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasDevicesUpdateErrorNoData(t *testing.T) {
	c := setupTest([]string{"devices", "update", "--host", "iota", "--id", "sensor001", "--data", "@"})

	err := idasDevicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestIdasDevicesUpdateErrorHTTP(t *testing.T) {
	data := `{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`
	c := setupTest([]string{"devices", "update", "--host", "iota", "--id", "sensor001", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasDevicesUpdateErrorHTTPStatus(t *testing.T) {
	data := `{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`
	c := setupTest([]string{"devices", "update", "--host", "iota", "--id", "sensor001", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ReqData = []byte(`{"entity_name": "urn:ngsi-ld:WeatherObserved:sensor333"}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestIdasDevicesDelete(t *testing.T) {
	c := setupTest([]string{"devices", "delete", "--host", "iota", "--id", "sensor001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices/sensor001"

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestIdasDevicesDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"devices", "delete", "--host", "iota", "--id", "sensor001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/iot/devices"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestIdasDevicesDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"devices", "delete", "--host", "iota", "--id", "sensor001"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/devices/sensor001"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := idasDevicesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGetCbroker(t *testing.T) {
	data := `{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`
	cc := setupTest([]string{"services", "create", "--host", "iota", "--data", data})

	cases := []struct {
		arg      string
		expected string
	}{
		{arg: "http://orion:1026", expected: "http://orion:1026"},
		{arg: "orion", expected: "https://orion"},
		{arg: "orion-ld", expected: "https://orion-ld"},
		{arg: "orion-alias", expected: "https://orion-ld"},
	}

	for _, c := range cases {
		actual, err := getCbroker(cc.Ngsi, c.arg)

		if assert.NoError(t, err) {
			assert.Equal(t, c.expected, actual)
		} else {
			ngsiErr := err.(*ngsierr.NgsiError)
			assert.Equal(t, 1, ngsiErr.ErrNo)
			assert.Equal(t, c.expected, ngsiErr.Message)
		}
	}
}

func TestGetCbrokerError(t *testing.T) {
	data := `{"services":[{"apikey":"apikey","cbroker":"http://orion:1026","entity_type":"Thing","resource":"/iot/d"}]}`
	cc := setupTest([]string{"services", "create", "--host", "iota", "--data", data})

	cases := []struct {
		arg      string
		expected string
	}{
		{arg: "orion:1026", expected: "specify url or broker alias to --cbroker"},
		{arg: "http:/orion:1026", expected: "specify url or broker alias to --cbroker"},
	}

	for _, c := range cases {
		actual, err := getCbroker(cc.Ngsi, c.arg)

		if assert.Error(t, err) {
			ngsiErr := err.(*ngsierr.NgsiError)
			assert.Equal(t, 1, ngsiErr.ErrNo)
			assert.Equal(t, "", actual)
			assert.Equal(t, c.expected, ngsiErr.Message)
		}
	}
}
