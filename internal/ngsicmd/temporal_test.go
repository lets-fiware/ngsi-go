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
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestTroeList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Sensor"})

	err := troeList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"urn:ngsi-ld:sensor100\",\"type\":\"Sensor\",\"temperature\":[{\"type\":\"Property\",\"value\":25,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-03-01T00:00:01Z\"},{\"type\":\"Property\",\"value\":21,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-02-01T00:00:01Z\"}],\"@context\":[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"]}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTroeListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Sensor", "--pretty"})

	err := troeList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"id\": \"urn:ngsi-ld:sensor100\",\n  \"type\": \"Sensor\",\n  \"temperature\": [\n    {\n      \"type\": \"Property\",\n      \"value\": 25,\n      \"instanceId\": \"REGEX(.*)\",\n      \"observedAt\": \"2017-03-01T00:00:01Z\"\n    },\n    {\n      \"type\": \"Property\",\n      \"value\": 21,\n      \"instanceId\": \"REGEX(.*)\",\n      \"observedAt\": \"2017-02-01T00:00:01Z\"\n    }\n  ],\n  \"@context\": [\n    \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTroeListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := troeList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeListErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	err := troeList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeListErrorBrokerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := troeList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeListErrorTemporalQuery(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,fromDate")
	_ = set.Parse([]string{"--host=orion-ld", "--fromDate=1"})
	c := cli.NewContext(app, set, nil)

	err := troeList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Sensor"})

	err := troeList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeListErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Sensor"})

	err := troeList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":[{"type":"Property","value":25,"instanceId":"REGEX(.*)","observedAt":"2017-03-01T00:00:01Z"},{"type":"Property","value":21,"instanceId":"REGEX(.*)","observedAt":"2017-02-01T00:00:01Z"}],"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Sensor", "--pretty"})

	setJSONIndentError(ngsi)

	err := troeList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", `--data={"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`})

	err := troeCreate(c)

	assert.NoError(t, err)
}

func TestTroeCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreateErrorBrokerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreateErrorDataEmpty(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,data")
	_ = set.Parse([]string{"--host=orion-ld", "--data="})
	c := cli.NewContext(app, set, nil)

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreateErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", `--data={"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}`})

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error: -01T00:00:01Z\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreateErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,context")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--context=ctx", `--data={"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`})

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ctx not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", `--data={"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`})

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeCreateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.ReqData = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", `--data={"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":21,"observedAt":"2017-02-01T00:00:01Z"}}`})

	err := troeCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeRead(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})

	err := troeRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"urn:ngsi-ld:sensor100\",\"type\":\"Sensor\",\"temperature\":{\"type\":\"Property\",\"value\":20,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-01-01T00:00:01Z\"},\"@context\":[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTroeReadAcceptJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "acceptJson")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--acceptJson", "--id=urn:ngsi-ld:sensor100"})

	err := troeRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"urn:ngsi-ld:sensor100\",\"type\":\"Sensor\",\"temperature\":{\"type\":\"Property\",\"value\":20,\"instanceId\":\"REGEX(.*)\",\"observedAt\":\"2017-01-01T00:00:01Z\"},\"@context\":[\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"]}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTroeReadPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--pretty", "--id=urn:ngsi-ld:sensor100"})

	err := troeRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"id\": \"urn:ngsi-ld:sensor100\",\n  \"type\": \"Sensor\",\n  \"temperature\": {\n    \"type\": \"Property\",\n    \"value\": 20,\n    \"instanceId\": \"REGEX(.*)\",\n    \"observedAt\": \"2017-01-01T00:00:01Z\"\n  },\n  \"@context\": [\n    \"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorBrokerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion-ld"})
	c := cli.NewContext(app, set, nil)

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify temporal entity id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorTemporalQuery(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id,fromDate")
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--fromDate=1"})
	c := cli.NewContext(app, set, nil)

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error:  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--id=urn:ngsi-ld:sensor100"})

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error: ontext.jsonld\"]", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeReadErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:sensor100","type":"Sensor","temperature":{"type":"Property","value":20,"instanceId":"REGEX(.*)","observedAt":"2017-01-01T00:00:01Z"},"@context":["https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--pretty", "--id=urn:ngsi-ld:sensor100"})

	setJSONIndentError(ngsi)

	err := troeRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})

	err := troeDelete(c)

	assert.NoError(t, err)
}

func TestTroeDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := troeDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeDeleteErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	err := troeDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeDeleteErrorBrokerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := troeDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeDeleteErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion-ld"})
	c := cli.NewContext(app, set, nil)

	err := troeDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify temporal entity id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})

	err := troeDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})

	err := troeDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppend(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrsAppend(c)

	assert.NoError(t, err)
}

func TestTroeAttrsAppendErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorBrokerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion-ld"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify temporal entity id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorDataEmpty(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id,data")
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--data="})
	c := cli.NewContext(app, set, nil)

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--id=urn:ngsi-ld:sensor100", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}`})

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error: -01T00:00:01Z\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data,context")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--context=ctx", "--id=urn:ngsi-ld:sensor100", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "ctx not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsAppendErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/"
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,attrName")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature"})

	err := troeAttrDelete(c)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteDeleteAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,attrName")
	setupFlagBool(set, "deleteAll")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--deleteAll", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature"})

	err := troeAttrDelete(c)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteDatasetId(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,attrName,datasetId")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--datasetId=datasetid001", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature"})

	err := troeAttrDelete(c)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteInstanceID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,attrName,instanceId")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature", "--instanceId=urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0"})

	err := troeAttrDelete(c)

	assert.NoError(t, err)
}

func TestTroeAttrDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDeleteErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDeleteErrorBrokerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDeleteErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion-ld"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify temporal entity id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDeleteErrorAttrName(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id")
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify attribute name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDeleteErrorInstanceId(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id,attrName,datasetId,instanceId")
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=attr", "--datasetId=dataset001", "--instanceId=instance001"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specify --deleteALl and/or --datasetId with --instanceId", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,attrName")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature"})

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,attrName")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature"})

	err := troeAttrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data,attrName,instanceId")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature", "--instanceId=instance001", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrUpdate(c)

	assert.NoError(t, err)
}

func TestTroeAttrsUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorBrokerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorID(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion-ld"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify temporal entity id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorAttrName(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id")
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify attribute name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorInstanceId(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id,attrName")
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify instance id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorDataEmpty(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,id,data,attrName,instanceId")
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature", "--instanceId=instance001"})
	c := cli.NewContext(app, set, nil)

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data,attrName,instanceId,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature", "--instanceId=instance001", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}`})
	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error: -01T00:00:01Z\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data,attrName,instanceId,context")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--context=ctx", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature", "--instanceId=instance001", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "ctx not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001/"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data,attrName,instanceId")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature", "--instanceId=instance001", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTroeAttrsUpdateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/temporal/entities/urn:ngsi-ld:sensor100/attrs/temperature/instance001"
	reqRes.ReqData = []byte(`{"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,data,attrName,instanceId")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:sensor100", "--attrName=temperature", "--instanceId=instance001", `--data={"temperature":{"type":"Property","value":20,"instanceId":"urn:ngsi-ld:1d293e44-01e2-4527-9a31-9cbdae761fe0","observedAt":"2017-01-01T00:00:01Z"}}`})

	err := troeAttrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBuildTemporalQueryBetween(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,fromDate,toDate")
	setupFlagBool(set, "etsi10")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	_ = set.Parse([]string{"--fromDate=2016-09-13T00:00:00.000Z", "--toDate=2017-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "between", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("timeAt"))
	assert.Equal(t, "2017-09-13T00:00:00.000Z", v.Get("endTimeAt"))
}

func TestBuildTemporalQueryAfter(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,fromDate,toDate")
	setupFlagBool(set, "etsi10")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	_ = set.Parse([]string{"--fromDate=2016-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "after", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("timeAt"))
}

func TestBuildTemporalQueryBefore(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,fromDate,toDate")
	setupFlagBool(set, "etsi10")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	_ = set.Parse([]string{"--toDate=2016-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "before", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("timeAt"))
}

func TestBuildTemporalQueryBetweenETSI10(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,fromDate,toDate")
	setupFlagBool(set, "etsi10")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	_ = set.Parse([]string{"--etsi10", "--fromDate=2016-09-13T00:00:00.000Z", "--toDate=2017-09-13T00:00:00.000Z"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	assert.NoError(t, err)
	assert.Equal(t, "between", v.Get("timerel"))
	assert.Equal(t, "2016-09-13T00:00:00.000Z", v.Get("time"))
	assert.Equal(t, "2017-09-13T00:00:00.000Z", v.Get("endTime"))
}

func TestBuildTemporalQueryError(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,fromDate,toDate")
	setupFlagBool(set, "etsi10")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	_ = set.Parse([]string{"--etsi10", "--fromDate=1"})

	v := url.Values{}

	err := buildTemporalQuery(c, &v)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}
