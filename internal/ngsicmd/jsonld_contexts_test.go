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
)

func TestJsonldContextsList(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	rawQuery := "details=true"
	reqRes.RawQuery = &rawQuery
	reqRes.ResBody = []byte(`[{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"fd564040-ece7-11eb-8e4a-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:08.133Z","hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","attributes":"https://uri.etsi.org/ngsi-ld/attributes","properties":"https://uri.etsi.org/ngsi-ld/properties"}},{"url":"http://atcontext:8000/ngsi-context.jsonld","id":"08d25d00-ece8-11eb-8d65-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:28.945Z","lastUse":"2021-07-25T01:31:35.335Z","lookups":187,"hash-table":{"familyName":"https://schema.org/familyName","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","providedBy":"https://uri.fiware.org/ns/data-models#providedBy","irrSection":"https://w3id.org/saref#irrSection","multimedia":"https://w3id.org/saref#multimedia"}},{"url":"http://atcontext:8000/test-context.jsonld","id":"0c6484d4-ece8-11eb-a312-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:34.936Z","lastUse":"2021-07-25T01:31:34.997Z","lookups":6,"hash-table":{"letsfiware":"https://context.lab.letsfiware.jp/dataset#","temperature":"https://w3id.org/saref#temperature","id":"@id","ｎａｍｅ":"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ","name":"https://context.lab.letsfiware.jp/dataset#name"}},{"url":"https://fiware.github.io/data-models/context.jsonld","id":"30abb6fa-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"hash-table":{"roadClosed":"https://uri.fiware.org/ns/data-models#roadClosed","copyMachineOrService":"https://uri.fiware.org/ns/data-models#copyMachineOrService","carSharing":"https://uri.fiware.org/ns/data-models#carSharing","areaServed":"https://schema.org/areaServed","anyVehicle":"https://uri.fiware.org/ns/data-models#anyVehicle"}},{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","id":"31443434-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T02:51:34.930Z","lookups":2,"hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","MultiPoint":"https://purl.org/geojson/vocab#MultiPoint","EntityType":"https://uri.etsi.org/ngsi-ld/EntityType"}},{"url":"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010","id":"2fa4dbc4-ece8-11eb-a645-0242c0a8a010","type":"array","origin":"Inline","createdAt":"1970-01-01T00:00:00.000Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"URLs":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}]`)

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "fd564040-ece7-11eb-8e4a-0242c0a8a010 https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\n08d25d00-ece8-11eb-8d65-0242c0a8a010 http://atcontext:8000/ngsi-context.jsonld\n0c6484d4-ece8-11eb-a312-0242c0a8a010 http://atcontext:8000/test-context.jsonld\n30abb6fa-ece8-11eb-a645-0242c0a8a010 https://fiware.github.io/data-models/context.jsonld\n31443434-ece8-11eb-a645-0242c0a8a010 https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\n2fa4dbc4-ece8-11eb-a645-0242c0a8a010 http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListJSON(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte(`["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","http://atcontext:8000/ngsi-context.jsonld","http://atcontext:8000/test-context.jsonld","https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"]`)

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"http://atcontext:8000/ngsi-context.jsonld\",\"http://atcontext:8000/test-context.jsonld\",\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\"]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListDetails(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld", "--details"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	rawQuery := "details=true"
	reqRes.RawQuery = &rawQuery
	reqRes.ResBody = []byte(`[{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"fd564040-ece7-11eb-8e4a-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:08.133Z","hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","attributes":"https://uri.etsi.org/ngsi-ld/attributes","properties":"https://uri.etsi.org/ngsi-ld/properties"}},{"url":"http://atcontext:8000/ngsi-context.jsonld","id":"08d25d00-ece8-11eb-8d65-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:28.945Z","lastUse":"2021-07-25T01:31:35.335Z","lookups":187,"hash-table":{"familyName":"https://schema.org/familyName","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","providedBy":"https://uri.fiware.org/ns/data-models#providedBy","irrSection":"https://w3id.org/saref#irrSection","multimedia":"https://w3id.org/saref#multimedia"}},{"url":"http://atcontext:8000/test-context.jsonld","id":"0c6484d4-ece8-11eb-a312-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:34.936Z","lastUse":"2021-07-25T01:31:34.997Z","lookups":6,"hash-table":{"letsfiware":"https://context.lab.letsfiware.jp/dataset#","temperature":"https://w3id.org/saref#temperature","id":"@id","ｎａｍｅ":"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ","name":"https://context.lab.letsfiware.jp/dataset#name"}},{"url":"https://fiware.github.io/data-models/context.jsonld","id":"30abb6fa-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"hash-table":{"roadClosed":"https://uri.fiware.org/ns/data-models#roadClosed","copyMachineOrService":"https://uri.fiware.org/ns/data-models#copyMachineOrService","carSharing":"https://uri.fiware.org/ns/data-models#carSharing","areaServed":"https://schema.org/areaServed","anyVehicle":"https://uri.fiware.org/ns/data-models#anyVehicle"}},{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","id":"31443434-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T02:51:34.930Z","lookups":2,"hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","MultiPoint":"https://purl.org/geojson/vocab#MultiPoint","EntityType":"https://uri.etsi.org/ngsi-ld/EntityType"}},{"url":"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010","id":"2fa4dbc4-ece8-11eb-a645-0242c0a8a010","type":"array","origin":"Inline","createdAt":"1970-01-01T00:00:00.000Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"URLs":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}]`)

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"url\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"fd564040-ece7-11eb-8e4a-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:31:08.133Z\",\"hash-table\":{\"instanceId\":\"https://uri.etsi.org/ngsi-ld/instanceId\",\"notifiedAt\":\"https://uri.etsi.org/ngsi-ld/notifiedAt\",\"observedAt\":\"https://uri.etsi.org/ngsi-ld/observedAt\",\"attributes\":\"https://uri.etsi.org/ngsi-ld/attributes\",\"properties\":\"https://uri.etsi.org/ngsi-ld/properties\"}},{\"url\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"08d25d00-ece8-11eb-8d65-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:31:28.945Z\",\"lastUse\":\"2021-07-25T01:31:35.335Z\",\"lookups\":187,\"hash-table\":{\"familyName\":\"https://schema.org/familyName\",\"observedAt\":\"https://uri.etsi.org/ngsi-ld/observedAt\",\"providedBy\":\"https://uri.fiware.org/ns/data-models#providedBy\",\"irrSection\":\"https://w3id.org/saref#irrSection\",\"multimedia\":\"https://w3id.org/saref#multimedia\"}},{\"url\":\"http://atcontext:8000/test-context.jsonld\",\"id\":\"0c6484d4-ece8-11eb-a312-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:31:34.936Z\",\"lastUse\":\"2021-07-25T01:31:34.997Z\",\"lookups\":6,\"hash-table\":{\"letsfiware\":\"https://context.lab.letsfiware.jp/dataset#\",\"temperature\":\"https://w3id.org/saref#temperature\",\"id\":\"@id\",\"ｎａｍｅ\":\"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ\",\"name\":\"https://context.lab.letsfiware.jp/dataset#name\"}},{\"url\":\"https://fiware.github.io/data-models/context.jsonld\",\"id\":\"30abb6fa-ece8-11eb-a645-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:32:34.083Z\",\"lastUse\":\"2021-07-25T01:32:36.855Z\",\"lookups\":1,\"hash-table\":{\"roadClosed\":\"https://uri.fiware.org/ns/data-models#roadClosed\",\"copyMachineOrService\":\"https://uri.fiware.org/ns/data-models#copyMachineOrService\",\"carSharing\":\"https://uri.fiware.org/ns/data-models#carSharing\",\"areaServed\":\"https://schema.org/areaServed\",\"anyVehicle\":\"https://uri.fiware.org/ns/data-models#anyVehicle\"}},{\"url\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"id\":\"31443434-ece8-11eb-a645-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:32:34.083Z\",\"lastUse\":\"2021-07-25T02:51:34.930Z\",\"lookups\":2,\"hash-table\":{\"instanceId\":\"https://uri.etsi.org/ngsi-ld/instanceId\",\"notifiedAt\":\"https://uri.etsi.org/ngsi-ld/notifiedAt\",\"observedAt\":\"https://uri.etsi.org/ngsi-ld/observedAt\",\"MultiPoint\":\"https://purl.org/geojson/vocab#MultiPoint\",\"EntityType\":\"https://uri.etsi.org/ngsi-ld/EntityType\"}},{\"url\":\"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\",\"id\":\"2fa4dbc4-ece8-11eb-a645-0242c0a8a010\",\"type\":\"array\",\"origin\":\"Inline\",\"createdAt\":\"1970-01-01T00:00:00.000Z\",\"lastUse\":\"2021-07-25T01:32:36.855Z\",\"lookups\":1,\"URLs\":[\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListPretty(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte(`["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","http://atcontext:8000/ngsi-context.jsonld","http://atcontext:8000/test-context.jsonld","https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"]`)

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n  \"http://atcontext:8000/ngsi-context.jsonld\",\n  \"http://atcontext:8000/test-context.jsonld\",\n  \"https://fiware.github.io/data-models/context.jsonld\",\n  \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\n  \"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\"\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.Err = errors.New("http get error")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestJsonldContextsListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte("bad request")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestJsonldContextsListErrorPretty(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte(`["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","http://atcontext:8000/ngsi-context.jsonld","http://atcontext:8000/test-context.jsonld","https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"]`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestJsonldContextsListErrorJSONUnmarsal(t *testing.T) {
	c := setupTest([]string{"list", "ldContexts", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	rawQuery := "details=true"
	reqRes.RawQuery = &rawQuery
	reqRes.ResBody = []byte(`[{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"fd564040-ece7-11eb-8e4a-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:08.133Z","hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","attributes":"https://uri.etsi.org/ngsi-ld/attributes","properties":"https://uri.etsi.org/ngsi-ld/properties"}},{"url":"http://atcontext:8000/ngsi-context.jsonld","id":"08d25d00-ece8-11eb-8d65-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:28.945Z","lastUse":"2021-07-25T01:31:35.335Z","lookups":187,"hash-table":{"familyName":"https://schema.org/familyName","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","providedBy":"https://uri.fiware.org/ns/data-models#providedBy","irrSection":"https://w3id.org/saref#irrSection","multimedia":"https://w3id.org/saref#multimedia"}},{"url":"http://atcontext:8000/test-context.jsonld","id":"0c6484d4-ece8-11eb-a312-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:34.936Z","lastUse":"2021-07-25T01:31:34.997Z","lookups":6,"hash-table":{"letsfiware":"https://context.lab.letsfiware.jp/dataset#","temperature":"https://w3id.org/saref#temperature","id":"@id","ｎａｍｅ":"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ","name":"https://context.lab.letsfiware.jp/dataset#name"}},{"url":"https://fiware.github.io/data-models/context.jsonld","id":"30abb6fa-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"hash-table":{"roadClosed":"https://uri.fiware.org/ns/data-models#roadClosed","copyMachineOrService":"https://uri.fiware.org/ns/data-models#copyMachineOrService","carSharing":"https://uri.fiware.org/ns/data-models#carSharing","areaServed":"https://schema.org/areaServed","anyVehicle":"https://uri.fiware.org/ns/data-models#anyVehicle"}},{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","id":"31443434-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T02:51:34.930Z","lookups":2,"hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","MultiPoint":"https://purl.org/geojson/vocab#MultiPoint","EntityType":"https://uri.etsi.org/ngsi-ld/EntityType"}},{"url":"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010","id":"2fa4dbc4-ece8-11eb-a645-0242c0a8a010","type":"array","origin":"Inline","createdAt":"1970-01-01T00:00:00.000Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"URLs":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}]`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := jsonldContextsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestJsonldContextGet(t *testing.T) {
	c := setupTest([]string{"get", "ldContext", "--host", "orion-ld", "--id", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"@context\":[\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextGetArg(t *testing.T) {
	c := setupTest([]string{"get", "ldContext", "--host", "orion-ld", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"@context\":[\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextGetPretty(t *testing.T) {
	c := setupTest([]string{"get", "ldContext", "--host", "orion-ld", "--pretty", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"@context\": [\n    \"https://fiware.github.io/data-models/context.jsonld\",\n    \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "ldContext", "--host", "orion-ld", "--id", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.Err = errors.New("http get error")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestJsonldContextGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"get", "ldContext", "--host", "orion-ld", "--id", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte("bad request")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestJsonldContextGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "ldContext", "--host", "orion-ld", "--id", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := jsonldContextGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestJsonldContextCreate(t *testing.T) {
	data := `["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`
	c := setupTest([]string{"create", "ldContext", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.ResHeader = http.Header{"Location": []string{"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/b1bd90a2-ed23-11eb-8a1f-0242c0a8a010"}}

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "b1bd90a2-ed23-11eb-8a1f-0242c0a8a010\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextCreateLocationEmpty(t *testing.T) {
	data := `["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`
	c := setupTest([]string{"create", "ldContext", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.ResHeader = http.Header{"Location": []string{""}}

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextCreateErrorData(t *testing.T) {
	c := setupTest([]string{"create", "ldContext", "--host", "orion-ld", "--data", "@"})

	err := jsonldContextCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestJsonldContextCreateErrorHTTP(t *testing.T) {
	data := `["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`
	c := setupTest([]string{"create", "ldContext", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.Err = errors.New("http get error")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestJsonldContextCreateErrorStatusCode(t *testing.T) {
	data := `["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`
	c := setupTest([]string{"create", "ldContext", "--host", "orion-ld", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.ResBody = []byte("bad request")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestJsonldContextDelete(t *testing.T) {
	c := setupTest([]string{"delete", "ldContext", "--host", "orion-ld", "--id", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestJsonldContextDeleteID(t *testing.T) {
	c := setupTest([]string{"delete", "ldContext", "--host", "orion-ld", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestJsonldContextDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "ldContext", "--host", "orion-ld", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.Err = errors.New("http get error")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}

func TestJsonldContextDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"delete", "ldContext", "--host", "orion-ld", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte("bad request")

	helper.SetClientHTTP(c, reqRes)

	err := jsonldContextDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	}
}
