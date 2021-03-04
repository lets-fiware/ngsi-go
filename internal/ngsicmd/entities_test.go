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

func TestEntitiesListLDMain(t *testing.T) {
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

func TestEntitiesListV2Main(t *testing.T) {
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

func TestEntitiesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
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

func TestEntitiesListErrorLDParam(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,typePattern")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--typePattern=Thing.*"})

	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specfiy typePattern, mq, metadata, value or uniq", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListErrorV2Param(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "acceptJson")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--acceptJson"})

	err := entitiesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specfiy link or acceptJson", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "9\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListCountV2AttrNone(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	rawQuery := "attrs=__NONE&limit=1&options=count"
	reqRes.RawQuery = &rawQuery
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--count"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "9\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2AttrNone(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	rawQuery := "attrs=__NONE&limit=100&offset=0&options=count"
	reqRes.RawQuery = &rawQuery
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "acceptJson")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--acceptJson"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2Page(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\nairqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2Verbose(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"airqualityobserved_0\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.727447926,\"metadata\":{}}},{\"id\":\"airqualityobserved_1\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":19.012560208,\"metadata\":{}}},{\"id\":\"airqualityobserved_2\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-3.196384014,\"metadata\":{}}},{\"id\":\"airqualityobserved_3\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":7.992932652,\"metadata\":{}}},{\"id\":\"airqualityobserved_4\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-6.620346091,\"metadata\":{}}},{\"id\":\"airqualityobserved_5\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-16.634766746,\"metadata\":{}}},{\"id\":\"airqualityobserved_6\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":20.263618173,\"metadata\":{}}},{\"id\":\"airqualityobserved_7\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":14.285382467,\"metadata\":{}}},{\"id\":\"airqualityobserved_8\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.998595286,\"metadata\":{}}}]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2VerbosePretty(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"id\": \"airqualityobserved_0\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.727447926,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_1\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 19.012560208,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_2\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -3.196384014,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_3\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 7.992932652,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_4\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -6.620346091,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_5\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -16.634766746,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_6\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 20.263618173,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_7\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 14.285382467,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_8\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.998595286,\n      \"metadata\": {}\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2VerboseLines(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"airqualityobserved_0\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":6.727447926},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_1\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":19.012560208},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_2\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-3.196384014},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_3\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":7.992932652},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_4\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-6.620346091},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_5\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-16.634766746},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_6\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":20.263618173},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_7\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":14.285382467},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_8\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":6.998595286},\"type\":\"AirQualityObserved\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2Values(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2ValuesLines(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[10.148599472]\n[14.627960669]\n[-2.461631059]\n[-15.999248065]\n[-4.553473866]\n[1.147149609]\n[1.003624237]\n[11.747977585]\n[-4.264932072]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2ErrorHTTP(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2ErrorHTTPStatus(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2ErrorResultsCount1(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2ErrorResultsCount2(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2ResultsCount3(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	assert.NoError(t, err)
}

func TestEntitiesListV2ErrorVerboseSafeString(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character ':' after array element (5) [\"id\":\"airqualityobs", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListV2ErrorVerboseLinesValues(t *testing.T) {
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
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

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "acceptJson")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--acceptJson"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:TemperatureSensor:001\nurn:ngsi-ld:TemperatureSensor:002\nurn:ngsi-ld:TemperatureSensor:003\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"102"}}
	reqRes1.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entities"
	reqRes2.ResHeader = http.Header{"Ngsild-Results-Count": []string{"102"}}
	reqRes2.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:TemperatureSensor:001\nurn:ngsi-ld:TemperatureSensor:002\nurn:ngsi-ld:TemperatureSensor:003\nurn:ngsi-ld:TemperatureSensor:001\nurn:ngsi-ld:TemperatureSensor:002\nurn:ngsi-ld:TemperatureSensor:003\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--verbose", "--attrs=temperature"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"}},{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":26,\"unitCode\":\"CEL\"}},{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:003\",\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":27,\"unitCode\":\"CEL\"}}]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDVerbosePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--verbose", "--attrs=temperature", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:001\",\n    \"type\": \"TemperatureSensor\",\n    \"temperature\": {\n      \"type\": \"Property\",\n      \"value\": 25,\n      \"unitCode\": \"CEL\"\n    }\n  },\n  {\n    \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:002\",\n    \"type\": \"TemperatureSensor\",\n    \"temperature\": {\n      \"type\": \"Property\",\n      \"value\": 26,\n      \"unitCode\": \"CEL\"\n    }\n  },\n  {\n    \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:003\",\n    \"type\": \"TemperatureSensor\",\n    \"temperature\": {\n      \"type\": \"Property\",\n      \"value\": 27,\n      \"unitCode\": \"CEL\"\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDVerboseLines(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,lines,keyValues")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--verbose", "--lines", "--attrs=temperature", "--keyValues"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"temperature\":{\"type\":\"Property\",\"unitCode\":\"CEL\",\"value\":25},\"type\":\"TemperatureSensor\"}\n{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"temperature\":{\"type\":\"Property\",\"unitCode\":\"CEL\",\"value\":26},\"type\":\"TemperatureSensor\"}\n{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:003\",\"temperature\":{\"type\":\"Property\",\"unitCode\":\"CEL\",\"value\":27},\"type\":\"TemperatureSensor\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Device"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Device"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDErrorResultsCount1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Device", "--count"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDErrorResultsCount2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Device"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDResultsCount3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Device"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	assert.NoError(t, err)
}

func TestEntitiesListLDErrorVerboseSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,attrs,safeString")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--verbose", "--attrs=temperature"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '@' (2) [{@context\":\"http", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesListLDErrorEntitiesPrint(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setJSONDecodeErr(ngsi, 1)

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type,attrs")
	setupFlagBool(set, "verbose,lines")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--verbose", "--lines", "--attrs=temperature"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = entitiesListLD(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrint(t *testing.T) {
	ngsi, _, _, buffer := setupTest()

	pretty := false
	lines := false
	values := false
	verbose := false

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.NoError(t, err) {
		actual := buffer.String()
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintLinesValues(t *testing.T) {
	ngsi, _, _, buffer := setupTest()

	pretty := false
	lines := true
	values := true
	verbose := false

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.NoError(t, err) {
		actual := buffer.String()
		expected := "[10.148599472]\n[14.627960669]\n[-2.461631059]\n[-15.999248065]\n[-4.553473866]\n[1.147149609]\n[1.003624237]\n[11.747977585]\n[-4.264932072]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintVerbosePretty(t *testing.T) {
	ngsi, _, _, buffer := setupTest()

	pretty := true
	lines := false
	values := false
	verbose := true

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	buf.bufferClose()

	if assert.NoError(t, err) {
		actual := buffer.String()
		expected := "[\n  {\n    \"id\": \"airqualityobserved_0\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.727447926,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_1\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 19.012560208,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_2\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -3.196384014,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_3\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 7.992932652,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_4\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -6.620346091,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_5\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -16.634766746,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_6\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 20.263618173,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_7\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 14.285382467,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_8\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.998595286,\n      \"metadata\": {}\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintVerbose(t *testing.T) {
	ngsi, _, _, buffer := setupTest()

	pretty := false
	lines := false
	values := false
	verbose := true

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	buf.bufferClose()

	if assert.NoError(t, err) {
		actual := buffer.String()
		expected := "[{\"id\":\"airqualityobserved_0\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.727447926,\"metadata\":{}}},{\"id\":\"airqualityobserved_1\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":19.012560208,\"metadata\":{}}},{\"id\":\"airqualityobserved_2\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-3.196384014,\"metadata\":{}}},{\"id\":\"airqualityobserved_3\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":7.992932652,\"metadata\":{}}},{\"id\":\"airqualityobserved_4\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-6.620346091,\"metadata\":{}}},{\"id\":\"airqualityobserved_5\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-16.634766746,\"metadata\":{}}},{\"id\":\"airqualityobserved_6\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":20.263618173,\"metadata\":{}}},{\"id\":\"airqualityobserved_7\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":14.285382467,\"metadata\":{}}},{\"id\":\"airqualityobserved_8\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.998595286,\"metadata\":{}}}]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintErrorVerboseLinesValuesDecode(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	pretty := false
	lines := true
	values := true
	verbose := false

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	setJSONDecodeErr(ngsi, 0)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintErrorVerboseLinesValuesEncode(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	pretty := false
	lines := true
	values := true
	verbose := false

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	setJSONEncodeErr(ngsi, 0)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintErrorVerboseLinesDecode(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	pretty := false
	lines := true
	values := false
	verbose := false

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	setJSONDecodeErr(ngsi, 0)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintErrorVerboseLinesEncode(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	pretty := false
	lines := true
	values := false
	verbose := false

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	setJSONEncodeErr(ngsi, 0)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintErrorVerbosePretty(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	pretty := true
	lines := false
	values := false
	verbose := true

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	setJSONIndentError(ngsi)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntitiesPrintErrorUnmarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	pretty := false
	lines := false
	values := false
	verbose := false

	buf := jsonBuffer{}
	buf.bufferOpen(ngsi.StdWriter)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	setJSONDecodeErr(ngsi, 0)

	err := entitiesPrint(ngsi, body, &buf, pretty, lines, values, verbose)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
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
		assert.Equal(t, "required host not found", ngsiErr.Message)
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
