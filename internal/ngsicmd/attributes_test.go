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

func TestAttributesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/attributes"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:EntityAttributeList:52890704-ecdf-11eb-9c78-0242c0a8800d","type":"EntityAttributeList","attributeList":["https://uri.fiware.org/ns/data-models#category","https://w3id.org/saref#temperature"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,attrib")
	setupFlagBool(set, "details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := attributesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"urn:ngsi-ld:EntityAttributeList:52890704-ecdf-11eb-9c78-0242c0a8800d\",\"type\":\"EntityAttributeList\",\"attributeList\":[\"https://uri.fiware.org/ns/data-models#category\",\"https://w3id.org/saref#temperature\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttributesListDetails(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/attributes"
	rawQuery := "details=true"
	reqRes.RawQuery = &rawQuery
	reqRes.ResBody = []byte(`[{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://uri.fiware.org/ns/data-models#category","type":"Attribute","attributeCount":1,"attributeTypes":["Property"],"typeNames":["https://uri.fiware.org/ns/data-models#TemperatureSensor"]},{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://w3id.org/saref#temperature","type":"Attribute","attributeCount":1,"attributeTypes":["Property"],"typeNames":["https://uri.fiware.org/ns/data-models#TemperatureSensor"]}]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,attrib")
	setupFlagBool(set, "details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--details"})

	err := attributesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"https://uri.fiware.org/ns/data-models#category\",\"type\":\"Attribute\",\"attributeCount\":1,\"attributeTypes\":[\"Property\"],\"typeNames\":[\"https://uri.fiware.org/ns/data-models#TemperatureSensor\"]},{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"https://w3id.org/saref#temperature\",\"type\":\"Attribute\",\"attributeCount\":1,\"attributeTypes\":[\"Property\"],\"typeNames\":[\"https://uri.fiware.org/ns/data-models#TemperatureSensor\"]}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttributesListAttrId(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/attributes/https%3A%2F%2Fw3id.org%2Fsaref%23temperature"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"https://w3id.org/saref#temperature","type":"Attribute","attributeCount":1,"attributeTypes":["Property"],"typeNames":["https://uri.fiware.org/ns/data-models#TemperatureSensor"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,attr")
	setupFlagBool(set, "details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--attr=https://w3id.org/saref#temperature"})

	err := attributesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\"id\":\"https://w3id.org/saref#temperature\",\"type\":\"Attribute\",\"attributeCount\":1,\"attributeTypes\":[\"Property\"],\"typeNames\":[\"https://uri.fiware.org/ns/data-models#TemperatureSensor\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttributesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/attributes"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:EntityAttributeList:52890704-ecdf-11eb-9c78-0242c0a8800d","type":"EntityAttributeList","attributeList":["https://uri.fiware.org/ns/data-models#category","https://w3id.org/saref#temperature"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,attrib")
	setupFlagBool(set, "details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--pretty"})

	err := attributesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"@context\": \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\",\n  \"id\": \"urn:ngsi-ld:EntityAttributeList:52890704-ecdf-11eb-9c78-0242c0a8800d\",\n  \"type\": \"EntityAttributeList\",\n  \"attributeList\": [\n    \"https://uri.fiware.org/ns/data-models#category\",\n    \"https://w3id.org/saref#temperature\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAttributesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := attributesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttributesListErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})

	err := attributesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttributesListErrorOrion(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := attributesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttributesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/attributes"
	reqRes.Err = errors.New("http get error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,attrib")
	setupFlagBool(set, "details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := attributesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		expected := "http get error"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttributesListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/attributes"
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,attrib")
	setupFlagBool(set, "details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	err := attributesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		expected := " bad request"
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttributesListIotaErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/attributes"
	reqRes.ResBody = []byte(`{"@context":"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld","id":"urn:ngsi-ld:EntityAttributeList:52890704-ecdf-11eb-9c78-0242c0a8800d","type":"EntityAttributeList","attributeList":["https://uri.fiware.org/ns/data-models#category","https://w3id.org/saref#temperature"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,link,attrib")
	setupFlagBool(set, "details,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--pretty"})

	setJSONIndentError(ngsi)

	err := attributesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
