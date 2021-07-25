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

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestJsonldContextsList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	rawQuery := "details=true"
	reqRes.RawQuery = &rawQuery
	reqRes.ResBody = []byte(`[{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"fd564040-ece7-11eb-8e4a-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:08.133Z","hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","attributes":"https://uri.etsi.org/ngsi-ld/attributes","properties":"https://uri.etsi.org/ngsi-ld/properties"}},{"url":"http://atcontext:8000/ngsi-context.jsonld","id":"08d25d00-ece8-11eb-8d65-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:28.945Z","lastUse":"2021-07-25T01:31:35.335Z","lookups":187,"hash-table":{"familyName":"https://schema.org/familyName","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","providedBy":"https://uri.fiware.org/ns/data-models#providedBy","irrSection":"https://w3id.org/saref#irrSection","multimedia":"https://w3id.org/saref#multimedia"}},{"url":"http://atcontext:8000/test-context.jsonld","id":"0c6484d4-ece8-11eb-a312-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:34.936Z","lastUse":"2021-07-25T01:31:34.997Z","lookups":6,"hash-table":{"letsfiware":"https://context.lab.letsfiware.jp/dataset#","temperature":"https://w3id.org/saref#temperature","id":"@id","ｎａｍｅ":"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ","name":"https://context.lab.letsfiware.jp/dataset#name"}},{"url":"https://fiware.github.io/data-models/context.jsonld","id":"30abb6fa-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"hash-table":{"roadClosed":"https://uri.fiware.org/ns/data-models#roadClosed","copyMachineOrService":"https://uri.fiware.org/ns/data-models#copyMachineOrService","carSharing":"https://uri.fiware.org/ns/data-models#carSharing","areaServed":"https://schema.org/areaServed","anyVehicle":"https://uri.fiware.org/ns/data-models#anyVehicle"}},{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","id":"31443434-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T02:51:34.930Z","lookups":2,"hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","MultiPoint":"https://purl.org/geojson/vocab#MultiPoint","EntityType":"https://uri.etsi.org/ngsi-ld/EntityType"}},{"url":"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010","id":"2fa4dbc4-ece8-11eb-a645-0242c0a8a010","type":"array","origin":"Inline","createdAt":"1970-01-01T00:00:00.000Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"URLs":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := jsonldContextsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "fd564040-ece7-11eb-8e4a-0242c0a8a010 https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\n08d25d00-ece8-11eb-8d65-0242c0a8a010 http://atcontext:8000/ngsi-context.jsonld\n0c6484d4-ece8-11eb-a312-0242c0a8a010 http://atcontext:8000/test-context.jsonld\n30abb6fa-ece8-11eb-a645-0242c0a8a010 https://fiware.github.io/data-models/context.jsonld\n31443434-ece8-11eb-a645-0242c0a8a010 https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\n2fa4dbc4-ece8-11eb-a645-0242c0a8a010 http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte(`["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","http://atcontext:8000/ngsi-context.jsonld","http://atcontext:8000/test-context.jsonld","https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json"})

	err := jsonldContextsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"http://atcontext:8000/ngsi-context.jsonld\",\"http://atcontext:8000/test-context.jsonld\",\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\"]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListDetails(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	rawQuery := "details=true"
	reqRes.RawQuery = &rawQuery
	reqRes.ResBody = []byte(`[{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"fd564040-ece7-11eb-8e4a-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:08.133Z","hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","attributes":"https://uri.etsi.org/ngsi-ld/attributes","properties":"https://uri.etsi.org/ngsi-ld/properties"}},{"url":"http://atcontext:8000/ngsi-context.jsonld","id":"08d25d00-ece8-11eb-8d65-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:28.945Z","lastUse":"2021-07-25T01:31:35.335Z","lookups":187,"hash-table":{"familyName":"https://schema.org/familyName","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","providedBy":"https://uri.fiware.org/ns/data-models#providedBy","irrSection":"https://w3id.org/saref#irrSection","multimedia":"https://w3id.org/saref#multimedia"}},{"url":"http://atcontext:8000/test-context.jsonld","id":"0c6484d4-ece8-11eb-a312-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:34.936Z","lastUse":"2021-07-25T01:31:34.997Z","lookups":6,"hash-table":{"letsfiware":"https://context.lab.letsfiware.jp/dataset#","temperature":"https://w3id.org/saref#temperature","id":"@id","ｎａｍｅ":"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ","name":"https://context.lab.letsfiware.jp/dataset#name"}},{"url":"https://fiware.github.io/data-models/context.jsonld","id":"30abb6fa-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"hash-table":{"roadClosed":"https://uri.fiware.org/ns/data-models#roadClosed","copyMachineOrService":"https://uri.fiware.org/ns/data-models#copyMachineOrService","carSharing":"https://uri.fiware.org/ns/data-models#carSharing","areaServed":"https://schema.org/areaServed","anyVehicle":"https://uri.fiware.org/ns/data-models#anyVehicle"}},{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","id":"31443434-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T02:51:34.930Z","lookups":2,"hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","MultiPoint":"https://purl.org/geojson/vocab#MultiPoint","EntityType":"https://uri.etsi.org/ngsi-ld/EntityType"}},{"url":"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010","id":"2fa4dbc4-ece8-11eb-a645-0242c0a8a010","type":"array","origin":"Inline","createdAt":"1970-01-01T00:00:00.000Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"URLs":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--details"})

	err := jsonldContextsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"url\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"fd564040-ece7-11eb-8e4a-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:31:08.133Z\",\"hash-table\":{\"instanceId\":\"https://uri.etsi.org/ngsi-ld/instanceId\",\"notifiedAt\":\"https://uri.etsi.org/ngsi-ld/notifiedAt\",\"observedAt\":\"https://uri.etsi.org/ngsi-ld/observedAt\",\"attributes\":\"https://uri.etsi.org/ngsi-ld/attributes\",\"properties\":\"https://uri.etsi.org/ngsi-ld/properties\"}},{\"url\":\"http://atcontext:8000/ngsi-context.jsonld\",\"id\":\"08d25d00-ece8-11eb-8d65-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:31:28.945Z\",\"lastUse\":\"2021-07-25T01:31:35.335Z\",\"lookups\":187,\"hash-table\":{\"familyName\":\"https://schema.org/familyName\",\"observedAt\":\"https://uri.etsi.org/ngsi-ld/observedAt\",\"providedBy\":\"https://uri.fiware.org/ns/data-models#providedBy\",\"irrSection\":\"https://w3id.org/saref#irrSection\",\"multimedia\":\"https://w3id.org/saref#multimedia\"}},{\"url\":\"http://atcontext:8000/test-context.jsonld\",\"id\":\"0c6484d4-ece8-11eb-a312-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:31:34.936Z\",\"lastUse\":\"2021-07-25T01:31:34.997Z\",\"lookups\":6,\"hash-table\":{\"letsfiware\":\"https://context.lab.letsfiware.jp/dataset#\",\"temperature\":\"https://w3id.org/saref#temperature\",\"id\":\"@id\",\"ｎａｍｅ\":\"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ\",\"name\":\"https://context.lab.letsfiware.jp/dataset#name\"}},{\"url\":\"https://fiware.github.io/data-models/context.jsonld\",\"id\":\"30abb6fa-ece8-11eb-a645-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:32:34.083Z\",\"lastUse\":\"2021-07-25T01:32:36.855Z\",\"lookups\":1,\"hash-table\":{\"roadClosed\":\"https://uri.fiware.org/ns/data-models#roadClosed\",\"copyMachineOrService\":\"https://uri.fiware.org/ns/data-models#copyMachineOrService\",\"carSharing\":\"https://uri.fiware.org/ns/data-models#carSharing\",\"areaServed\":\"https://schema.org/areaServed\",\"anyVehicle\":\"https://uri.fiware.org/ns/data-models#anyVehicle\"}},{\"url\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"id\":\"31443434-ece8-11eb-a645-0242c0a8a010\",\"type\":\"hash-table\",\"origin\":\"Downloaded\",\"createdAt\":\"2021-07-25T01:32:34.083Z\",\"lastUse\":\"2021-07-25T02:51:34.930Z\",\"lookups\":2,\"hash-table\":{\"instanceId\":\"https://uri.etsi.org/ngsi-ld/instanceId\",\"notifiedAt\":\"https://uri.etsi.org/ngsi-ld/notifiedAt\",\"observedAt\":\"https://uri.etsi.org/ngsi-ld/observedAt\",\"MultiPoint\":\"https://purl.org/geojson/vocab#MultiPoint\",\"EntityType\":\"https://uri.etsi.org/ngsi-ld/EntityType\"}},{\"url\":\"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\",\"id\":\"2fa4dbc4-ece8-11eb-a645-0242c0a8a010\",\"type\":\"array\",\"origin\":\"Inline\",\"createdAt\":\"1970-01-01T00:00:00.000Z\",\"lastUse\":\"2021-07-25T01:32:36.855Z\",\"lookups\":1,\"URLs\":[\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte(`["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","http://atcontext:8000/ngsi-context.jsonld","http://atcontext:8000/test-context.jsonld","https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json", "--pretty"})

	err := jsonldContextsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n  \"http://atcontext:8000/ngsi-context.jsonld\",\n  \"http://atcontext:8000/test-context.jsonld\",\n  \"https://fiware.github.io/data-models/context.jsonld\",\n  \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\n  \"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010\"\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := jsonldContextsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextsListErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})

	err := jsonldContextsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextsListErrorOrion(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := jsonldContextsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextsListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.Err = errors.New("http get error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := jsonldContextsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextsListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := jsonldContextsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextsListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ResBody = []byte(`["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","http://atcontext:8000/ngsi-context.jsonld","http://atcontext:8000/test-context.jsonld","https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json", "--pretty"})

	setJSONIndentError(ngsi)

	err := jsonldContextsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestJsonldContextsListErrorJSONUnmarsal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	rawQuery := "details=true"
	reqRes.RawQuery = &rawQuery
	reqRes.ResBody = []byte(`[{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"fd564040-ece7-11eb-8e4a-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:08.133Z","hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","attributes":"https://uri.etsi.org/ngsi-ld/attributes","properties":"https://uri.etsi.org/ngsi-ld/properties"}},{"url":"http://atcontext:8000/ngsi-context.jsonld","id":"08d25d00-ece8-11eb-8d65-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:28.945Z","lastUse":"2021-07-25T01:31:35.335Z","lookups":187,"hash-table":{"familyName":"https://schema.org/familyName","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","providedBy":"https://uri.fiware.org/ns/data-models#providedBy","irrSection":"https://w3id.org/saref#irrSection","multimedia":"https://w3id.org/saref#multimedia"}},{"url":"http://atcontext:8000/test-context.jsonld","id":"0c6484d4-ece8-11eb-a312-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:31:34.936Z","lastUse":"2021-07-25T01:31:34.997Z","lookups":6,"hash-table":{"letsfiware":"https://context.lab.letsfiware.jp/dataset#","temperature":"https://w3id.org/saref#temperature","id":"@id","ｎａｍｅ":"https://context.lab.letsfiware.jp/dataset#ｎａｍｅ","name":"https://context.lab.letsfiware.jp/dataset#name"}},{"url":"https://fiware.github.io/data-models/context.jsonld","id":"30abb6fa-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"hash-table":{"roadClosed":"https://uri.fiware.org/ns/data-models#roadClosed","copyMachineOrService":"https://uri.fiware.org/ns/data-models#copyMachineOrService","carSharing":"https://uri.fiware.org/ns/data-models#carSharing","areaServed":"https://schema.org/areaServed","anyVehicle":"https://uri.fiware.org/ns/data-models#anyVehicle"}},{"url":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld","id":"31443434-ece8-11eb-a645-0242c0a8a010","type":"hash-table","origin":"Downloaded","createdAt":"2021-07-25T01:32:34.083Z","lastUse":"2021-07-25T02:51:34.930Z","lookups":2,"hash-table":{"instanceId":"https://uri.etsi.org/ngsi-ld/instanceId","notifiedAt":"https://uri.etsi.org/ngsi-ld/notifiedAt","observedAt":"https://uri.etsi.org/ngsi-ld/observedAt","MultiPoint":"https://purl.org/geojson/vocab#MultiPoint","EntityType":"https://uri.etsi.org/ngsi-ld/EntityType"}},{"url":"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010","id":"2fa4dbc4-ece8-11eb-a645-0242c0a8a010","type":"array","origin":"Inline","createdAt":"1970-01-01T00:00:00.000Z","lastUse":"2021-07-25T01:32:36.855Z","lookups":1,"URLs":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link")
	setupFlagBool(set, "json,details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	setJSONDecodeErr(ngsi, 1)

	err := jsonldContextsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestJsonldContextGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"@context\":[\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextGetArg(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"@context\":[\"https://fiware.github.io/data-models/context.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--pretty", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"@context\": [\n    \"https://fiware.github.io/data-models/context.jsonld\",\n    \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestJsonldContextGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := jsonldContextGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextGetErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})

	err := jsonldContextGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextGetErrorOrion(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := jsonldContextGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextGetErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := jsonldContextGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing jsonldContext id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.Err = errors.New("http get error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte(`{"@context":["https://fiware.github.io/data-models/context.jsonld","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--pretty", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	setJSONIndentError(ngsi)

	err := jsonldContextGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestJsonldContextCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.ResHeader = http.Header{"Location": []string{"http://58dac41cd926:1026/ngsi-ld/v1/jsonldContexts/b1bd90a2-ed23-11eb-8a1f-0242c0a8a010"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", `--data=["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`})

	err := jsonldContextCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "b1bd90a2-ed23-11eb-8a1f-0242c0a8a010\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateLocationEmpty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.ResHeader = http.Header{"Location": []string{""}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", `--data=["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`})

	err := jsonldContextCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := jsonldContextCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})

	err := jsonldContextCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateErrorOrion(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := jsonldContextCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := jsonldContextCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing jsonldContext data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateErrorData(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data="})

	err := jsonldContextCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.Err = errors.New("http get error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", `--data=["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`})
	err := jsonldContextCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts"
	reqRes.ReqData = []byte((`["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`))
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", `--data=["https://fiware.github.io/tutorials.Step-by-Step/tutorials-context.jsonld","https://fiware.github.io/data-models/context.jsonld"]`})

	err := jsonldContextCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextDelete(c)

	assert.NoError(t, err)
}

func TestJsonldContextDeleteID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextDelete(c)

	assert.NoError(t, err)
}

func TestJsonldContextDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := jsonldContextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextDeleteErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})

	err := jsonldContextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextDeleteErrorOrion(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := jsonldContextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextDeleteErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := jsonldContextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "missing jsonldContext id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.Err = errors.New("http get error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJsonldContextDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/jsonldContexts/2fa4dbc4-ece8-11eb-a645-0242c0a8a010"
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=2fa4dbc4-ece8-11eb-a645-0242c0a8a010"})

	err := jsonldContextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
