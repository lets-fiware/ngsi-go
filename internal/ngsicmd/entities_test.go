/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestEntitiesListLDMain(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "8\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2Main(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	helper.SetClientHTTP(c, reqRes)

	err := entitiesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "9\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2MainSkipForwarding(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--count", "--skipForwarding"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	q := "attrs=__NONE&limit=1&options=count"
	reqRes.RawQuery = &q
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "9\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListErrorLDParam(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--typePattern", "Thing.*"})

	err := entitiesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specfiy typePattern, mq, metadata, value or uniq", ngsiErr.Message)
	}
}

func TestEntitiesListErrorV2Param(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--acceptJson"})

	err := entitiesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specfiy link acceptJson or acceptGeoJson", ngsiErr.Message)
	}
}

func TestEntitiesListCountV2(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "9\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListCountV2AttrNone(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	rawQuery := "attrs=__NONE&limit=1&options=count"
	reqRes.RawQuery = &rawQuery

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "9\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2AttrNone(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)
	rawQuery := "attrs=__NONE&limit=100&offset=0&options=count"
	reqRes.RawQuery = &rawQuery

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2Page(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v2/entities"
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"102"}}
	reqRes1.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/v2/entities"
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"102"}}
	reqRes2.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\nairqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2Verbose(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--verbose", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"airqualityobserved_0\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.727447926,\"metadata\":{}}},{\"id\":\"airqualityobserved_1\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":19.012560208,\"metadata\":{}}},{\"id\":\"airqualityobserved_2\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-3.196384014,\"metadata\":{}}},{\"id\":\"airqualityobserved_3\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":7.992932652,\"metadata\":{}}},{\"id\":\"airqualityobserved_4\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-6.620346091,\"metadata\":{}}},{\"id\":\"airqualityobserved_5\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-16.634766746,\"metadata\":{}}},{\"id\":\"airqualityobserved_6\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":20.263618173,\"metadata\":{}}},{\"id\":\"airqualityobserved_7\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":14.285382467,\"metadata\":{}}},{\"id\":\"airqualityobserved_8\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.998595286,\"metadata\":{}}}]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2VerbosePretty(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--verbose", "--attrs", "temperature", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"airqualityobserved_0\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.727447926,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_1\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 19.012560208,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_2\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -3.196384014,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_3\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 7.992932652,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_4\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -6.620346091,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_5\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -16.634766746,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_6\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 20.263618173,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_7\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 14.285382467,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_8\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.998595286,\n      \"metadata\": {}\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2VerboseLines(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--verbose", "--lines", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"airqualityobserved_0\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":6.727447926},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_1\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":19.012560208},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_2\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-3.196384014},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_3\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":7.992932652},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_4\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-6.620346091},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_5\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":-16.634766746},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_6\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":20.263618173},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_7\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":14.285382467},\"type\":\"AirQualityObserved\"}\n{\"id\":\"airqualityobserved_8\",\"temperature\":{\"metadata\":{},\"type\":\"Number\",\"value\":6.998595286},\"type\":\"AirQualityObserved\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2Values(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--values", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2ValuesLines(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--values", "--lines", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[10.148599472]\n[14.627960669]\n[-2.461631059]\n[-15.999248065]\n[-4.553473866]\n[1.147149609]\n[1.003624237]\n[11.747977585]\n[-4.264932072]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListV2ResultsCount(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntitiesListV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-List": []string{"8"}}
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestEntitiesListV2ErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-List": []string{"8"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	}
}

func TestEntitiesListV2ErrorResultsCount1(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--type", "Device", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestEntitiesListV2ErrorResultsCount2(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestEntitiesListV2ErrorVerboseSafeString(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--safeString", "on", "--verbose", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`["id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character ':' after array element (5) [\"id\":\"airqualityobs", ngsiErr.Message)
	}
}

func TestEntitiesListV2ErrorVerboseLinesValues(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--values", "--verbose", "--lines", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := entitiesListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesListCountLD(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "8\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLD(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:TemperatureSensor:001\nurn:ngsi-ld:TemperatureSensor:002\nurn:ngsi-ld:TemperatureSensor:003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDAcceptJson(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--pretty", "--acceptJson"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"9"}}
	reqRes.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:001\",\n    \"type\": \"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\n    \"https://uri.fiware.org/ns/data-models#category\": {\n      \"type\": \"Property\",\n      \"value\": \"sensor\"\n    },\n    \"https://w3id.org/saref#temperature\": {\n      \"type\": \"Property\",\n      \"value\": 25,\n      \"unitCode\": \"CEL\"\n    }\n  },\n  {\n    \"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:002\",\n    \"type\": \"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\n    \"https://uri.fiware.org/ns/data-models#category\": {\n      \"type\": \"Property\",\n      \"value\": \"sensor\"\n    },\n    \"https://w3id.org/saref#temperature\": {\n      \"type\": \"Property\",\n      \"value\": 26,\n      \"unitCode\": \"CEL\"\n    }\n  },\n  {\n    \"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:003\",\n    \"type\": \"https://uri.fiware.org/ns/data-models#TemperatureSensor\",\n    \"https://uri.fiware.org/ns/data-models#category\": {\n      \"type\": \"Property\",\n      \"value\": \"sensor\"\n    },\n    \"https://w3id.org/saref#temperature\": {\n      \"type\": \"Property\",\n      \"value\": 27,\n      \"unitCode\": \"CEL\"\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDGeoJSON(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--acceptGeoJson"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`{"type":"FeatureCollection","features":[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}},"@context":"http://atcontext:8000/ngsi-context.jsonld","geometry":{"type":"Point","coordinates":[139.76,35.68]}},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[135.75,34.98]}}},"@context":"http://atcontext:8000/ngsi-context.jsonld","geometry":{"type":"Point","coordinates":[135.75,34.98]}},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[135.49,34.7]}}},"@context":"http://atcontext:8000/ngsi-context.jsonld","geometry":{"type":"Point","coordinates":[135.49,34.7]}}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"FeatureCollection\",\"features\":[{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}},\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"geometry\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}},{\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":26,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[135.75,34.98]}}},\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"geometry\":{\"type\":\"Point\",\"coordinates\":[135.75,34.98]}},{\"id\":\"urn:ngsi-ld:TemperatureSensor:003\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":27,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[135.49,34.7]}}},\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"geometry\":{\"type\":\"Point\",\"coordinates\":[135.49,34.7]}}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDGeoJSONPretty(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--acceptGeoJson", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`{"type":"FeatureCollection","features":[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}},"@context":"http://atcontext:8000/ngsi-context.jsonld","geometry":{"type":"Point","coordinates":[139.76,35.68]}},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[135.75,34.98]}}},"@context":"http://atcontext:8000/ngsi-context.jsonld","geometry":{"type":"Point","coordinates":[135.75,34.98]}},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[135.49,34.7]}}},"@context":"http://atcontext:8000/ngsi-context.jsonld","geometry":{"type":"Point","coordinates":[135.49,34.7]}}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": [\n    {\n      \"id\": \"urn:ngsi-ld:TemperatureSensor:001\",\n      \"type\": \"Feature\",\n      \"properties\": {\n        \"type\": \"TemperatureSensor\",\n        \"temperature\": {\n          \"type\": \"Property\",\n          \"value\": 25,\n          \"unitCode\": \"CEL\"\n        },\n        \"location\": {\n          \"type\": \"GeoProperty\",\n          \"value\": {\n            \"type\": \"Point\",\n            \"coordinates\": [\n              139.76,\n              35.68\n            ]\n          }\n        }\n      },\n      \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n      \"geometry\": {\n        \"type\": \"Point\",\n        \"coordinates\": [\n          139.76,\n          35.68\n        ]\n      }\n    },\n    {\n      \"id\": \"urn:ngsi-ld:TemperatureSensor:002\",\n      \"type\": \"Feature\",\n      \"properties\": {\n        \"type\": \"TemperatureSensor\",\n        \"temperature\": {\n          \"type\": \"Property\",\n          \"value\": 26,\n          \"unitCode\": \"CEL\"\n        },\n        \"location\": {\n          \"type\": \"GeoProperty\",\n          \"value\": {\n            \"type\": \"Point\",\n            \"coordinates\": [\n              135.75,\n              34.98\n            ]\n          }\n        }\n      },\n      \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n      \"geometry\": {\n        \"type\": \"Point\",\n        \"coordinates\": [\n          135.75,\n          34.98\n        ]\n      }\n    },\n    {\n      \"id\": \"urn:ngsi-ld:TemperatureSensor:003\",\n      \"type\": \"Feature\",\n      \"properties\": {\n        \"type\": \"TemperatureSensor\",\n        \"temperature\": {\n          \"type\": \"Property\",\n          \"value\": 27,\n          \"unitCode\": \"CEL\"\n        },\n        \"location\": {\n          \"type\": \"GeoProperty\",\n          \"value\": {\n            \"type\": \"Point\",\n            \"coordinates\": [\n              135.49,\n              34.7\n            ]\n          }\n        }\n      },\n      \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n      \"geometry\": {\n        \"type\": \"Point\",\n        \"coordinates\": [\n          135.49,\n          34.7\n        ]\n      }\n    }\n  ]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDPage(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"102"}}
	reqRes1.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entities"
	reqRes2.ResHeader = http.Header{"Ngsild-Results-Count": []string{"102"}}
	reqRes2.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"https://uri.fiware.org/ns/data-models#TemperatureSensor","https://uri.fiware.org/ns/data-models#category":{"type":"Property","value":"sensor"},"https://w3id.org/saref#temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:TemperatureSensor:001\nurn:ngsi-ld:TemperatureSensor:002\nurn:ngsi-ld:TemperatureSensor:003\nurn:ngsi-ld:TemperatureSensor:001\nurn:ngsi-ld:TemperatureSensor:002\nurn:ngsi-ld:TemperatureSensor:003\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDVerbose(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--verbose", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"}},{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":26,\"unitCode\":\"CEL\"}},{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:003\",\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":27,\"unitCode\":\"CEL\"}}]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDVerbosePretty(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--verbose", "--attrs", "temperature", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:001\",\n    \"type\": \"TemperatureSensor\",\n    \"temperature\": {\n      \"type\": \"Property\",\n      \"value\": 25,\n      \"unitCode\": \"CEL\"\n    }\n  },\n  {\n    \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:002\",\n    \"type\": \"TemperatureSensor\",\n    \"temperature\": {\n      \"type\": \"Property\",\n      \"value\": 26,\n      \"unitCode\": \"CEL\"\n    }\n  },\n  {\n    \"@context\": \"http://atcontext:8000/ngsi-context.jsonld\",\n    \"id\": \"urn:ngsi-ld:TemperatureSensor:003\",\n    \"type\": \"TemperatureSensor\",\n    \"temperature\": {\n      \"type\": \"Property\",\n      \"value\": 27,\n      \"unitCode\": \"CEL\"\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDVerboseLines(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--verbose", "--lines", "--attrs", "temperature", "--keyValues"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"temperature\":{\"type\":\"Property\",\"unitCode\":\"CEL\",\"value\":25},\"type\":\"TemperatureSensor\"}\n{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"temperature\":{\"type\":\"Property\",\"unitCode\":\"CEL\",\"value\":26},\"type\":\"TemperatureSensor\"}\n{\"@context\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"urn:ngsi-ld:TemperatureSensor:003\",\"temperature\":{\"type\":\"Property\",\"unitCode\":\"CEL\",\"value\":27},\"type\":\"TemperatureSensor\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesListLDResultsCount(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestEntitiesListLDErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestEntitiesListLDErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	}
}

func TestEntitiesListLDErrorResultsCount1(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--type", "Device", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestEntitiesListLDErrorResultsCount2(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--type", "Device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestEntitiesListLDErrorVerboseSafeString(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--safeString", "on", "--verbose", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '@' (2) [{@context\":\"http", ngsiErr.Message)
	}
}

func TestEntitiesListLDErrorEntitiesPrint(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld", "--verbose", "--lines", "--attrs", "temperature"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes.ResBody = []byte(`[{@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor","temperature":{"type":"Property","value":26,"unitCode":"CEL"}},{"@context":"http://atcontext:8000/ngsi-context.jsonld","id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor","temperature":{"type":"Property","value":27,"unitCode":"CEL"}}]`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := entitiesListLD(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesPrint(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := false
	values := false
	verbose := false

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "airqualityobserved_0\nairqualityobserved_1\nairqualityobserved_2\nairqualityobserved_3\nairqualityobserved_4\nairqualityobserved_5\nairqualityobserved_6\nairqualityobserved_7\nairqualityobserved_8\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesPrintLinesValues(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := true
	values := true
	verbose := false

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[10.148599472]\n[14.627960669]\n[-2.461631059]\n[-15.999248065]\n[-4.553473866]\n[1.147149609]\n[1.003624237]\n[11.747977585]\n[-4.264932072]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesPrintVerbosePretty(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := true
	lines := false
	values := false
	verbose := true

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	buf.BufferClose()

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"airqualityobserved_0\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.727447926,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_1\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 19.012560208,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_2\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -3.196384014,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_3\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 7.992932652,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_4\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -6.620346091,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_5\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": -16.634766746,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_6\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 20.263618173,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_7\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 14.285382467,\n      \"metadata\": {}\n    }\n  },\n  {\n    \"id\": \"airqualityobserved_8\",\n    \"type\": \"AirQualityObserved\",\n    \"temperature\": {\n      \"type\": \"Number\",\n      \"value\": 6.998595286,\n      \"metadata\": {}\n    }\n  }\n]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesPrintVerbose(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := false
	values := false
	verbose := true

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	buf.BufferClose()

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"airqualityobserved_0\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.727447926,\"metadata\":{}}},{\"id\":\"airqualityobserved_1\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":19.012560208,\"metadata\":{}}},{\"id\":\"airqualityobserved_2\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-3.196384014,\"metadata\":{}}},{\"id\":\"airqualityobserved_3\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":7.992932652,\"metadata\":{}}},{\"id\":\"airqualityobserved_4\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-6.620346091,\"metadata\":{}}},{\"id\":\"airqualityobserved_5\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":-16.634766746,\"metadata\":{}}},{\"id\":\"airqualityobserved_6\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":20.263618173,\"metadata\":{}}},{\"id\":\"airqualityobserved_7\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":14.285382467,\"metadata\":{}}},{\"id\":\"airqualityobserved_8\",\"type\":\"AirQualityObserved\",\"temperature\":{\"type\":\"Number\",\"value\":6.998595286,\"metadata\":{}}}]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesPrintGeoJSON(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := false
	values := false
	verbose := true

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`{"type":"FeatureCollection","features":[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}]}`)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	buf.BufferClose()

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\"type\":\"FeatureCollection\",\"features\":[{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}]]"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesPrintErrorGeoJSON(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := true
	values := true
	verbose := false
	geoJSON := true

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`{}`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, geoJSON)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "geojson error: {}", ngsiErr.Message)
	}
}

func TestEntitiesPrintErrorVerboseLinesValuesDecode(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := true
	values := true
	verbose := false

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesPrintErrorVerboseLinesValuesEncode(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := true
	values := true
	verbose := false

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[[10.148599472],[14.627960669],[-2.461631059],[-15.999248065],[-4.553473866],[1.147149609],[1.003624237],[11.747977585],[-4.264932072]]`)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesPrintErrorVerboseLinesDecode(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := true
	values := false
	verbose := false

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesPrintErrorVerboseLinesEncode(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := true
	values := false
	verbose := false

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesPrintErrorVerbosePretty(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := true
	lines := false
	values := false
	verbose := true

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetJSONIndentError(c.Ngsi)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesPrintErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	pretty := false
	lines := false
	values := false
	verbose := false

	buf := ngsilib.NewJsonBuffer()
	buf.BufferOpen(c.Ngsi.StdWriter, false, false)

	body := []byte(`[{"id":"airqualityobserved_0","type":"AirQualityObserved","temperature":{"type":"Number","value":6.727447926,"metadata":{}}},{"id":"airqualityobserved_1","type":"AirQualityObserved","temperature":{"type":"Number","value":19.012560208,"metadata":{}}},{"id":"airqualityobserved_2","type":"AirQualityObserved","temperature":{"type":"Number","value":-3.196384014,"metadata":{}}},{"id":"airqualityobserved_3","type":"AirQualityObserved","temperature":{"type":"Number","value":7.992932652,"metadata":{}}},{"id":"airqualityobserved_4","type":"AirQualityObserved","temperature":{"type":"Number","value":-6.620346091,"metadata":{}}},{"id":"airqualityobserved_5","type":"AirQualityObserved","temperature":{"type":"Number","value":-16.634766746,"metadata":{}}},{"id":"airqualityobserved_6","type":"AirQualityObserved","temperature":{"type":"Number","value":20.263618173,"metadata":{}}},{"id":"airqualityobserved_7","type":"AirQualityObserved","temperature":{"type":"Number","value":14.285382467,"metadata":{}}},{"id":"airqualityobserved_8","type":"AirQualityObserved","temperature":{"type":"Number","value":6.998595286,"metadata":{}}}]`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := entitiesPrint(c.Ngsi, body, buf, pretty, lines, values, verbose, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestEntitiesCountV2(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"10"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "10\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesCountV2SkipForwarding(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--skipForwarding"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.RawQuery = helper.StrPtr("attrs=__NONE&limit=1&options=count%2CskipForwarding")
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"10"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "10\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesCountV2Type(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion", "--type", "device"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"10"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "10\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesCountLD(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"15"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "15\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesCountLDType(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "8\n"
		assert.Equal(t, expected, actual)
	}
}

func TestEntitiesCountErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestEntitiesCountErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestEntitiesCountErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"list", "entities", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := entitiesCount(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}
