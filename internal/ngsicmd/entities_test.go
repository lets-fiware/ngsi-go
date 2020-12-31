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

func TestEntitiesListCountV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--count"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "9\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListCountLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--count"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")
	setupFlagBool(set, "acceptJson")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--acceptJson"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v2/entities"
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"102"}}
	reqRes1.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/v2/entities"
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"102"}}
	reqRes2.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\nairqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"airqualityobserved_0\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.727447926,\"metadata\":{}}},{\"id\":\"airqualityobserved_1\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":19.012560208,\"metadata\":{}}},{\"id\":\"airqualityobserved_2\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-3.196384014,\"metadata\":{}}},{\"id\":\"airqualityobserved_3\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":7.992932652,\"metadata\":{}}},{\"id\":\"airqualityobserved_4\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-6.620346091,\"metadata\":{}}},{\"id\":\"airqualityobserved_5\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-16.634766746,\"metadata\":{}}},{\"id\":\"airqualityobserved_6\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":20.263618173,\"metadata\":{}}},{\"id\":\"airqualityobserved_7\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":14.285382467,\"metadata\":{}}},{\"id\":\"airqualityobserved_8\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.998595286,\"metadata\":{}}}]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListVerbosePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--attrs=temperature", "--pretty"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"id\": \"airqualityobserved_0\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.727447926,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_1\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 19.012560208,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_2\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -3.196384014,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_3\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 7.992932652,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_4\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -6.620346091,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_5\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -16.634766746,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_6\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 20.263618173,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_7\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 14.285382467,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_8\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.998595286,\n      \"metadata\": {}\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListVerboseLines(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,lines")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--lines", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"airqualityobserved_0\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":6.727447926},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_1\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":19.012560208},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_2\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-3.196384014},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_3\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":7.992932652},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_4\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-6.620346091},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_5\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-16.634766746},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_6\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":20.263618173},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_7\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":14.285382467},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_8\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":6.998595286},\"type\":\"AirQualityObserved\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListValues(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,values")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--values", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListValuesLines(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,values,lines")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--values", "--lines", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[10.148599472]\n[14.627960669]\n[-2.461631059]\n[-15.999248065]\n[-4.553473866]\n[1.147149609]\n[1.003624237]\n[11.747977585]\n[-4.264932072]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-List": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-List": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorResultsCount1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device", "--count"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorResultsCount2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorVerboseSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`["id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs,safeString")
	setupFlagBool(set, "verbose")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--verbose", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character ':' after array element (5) [\"id\":\"airqualityobs", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorVerboseLinesValues(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setJSONDecodeErr(ngsi, 1)
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,lines,values")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--values", "--verbose", "--lines", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorVerboseLinesValues2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setJSONEncodeErr(ngsi, 2)
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,lines,values")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--values", "--verbose", "--lines", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorVerboseLines(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setJSONDecodeErr(ngsi, 1)
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,lines")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--lines", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorVerboseLines2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setJSONEncodeErr(ngsi, 2)
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,lines")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--lines", "--attrs=temperature"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorVerbosePretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--attrs=temperature", "--pretty"})

	setJSONIndentError(ngsi)

	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 12, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setJSONDecodeErr(ngsi, 1)
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 13, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorResultsCount3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesList(c)

	assert.NoError(t, err)
}

func TestEntitiesCountV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"10"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := entitiesCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "10\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountV2Type(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"10"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "10\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"15"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := entitiesCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "15\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountLDType(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Device"})
	err := entitiesCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := entitiesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := entitiesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestEntitiesCountErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entitiesCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
