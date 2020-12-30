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
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestRegistrationsListLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdLocalTime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"expires": "2020-09-01T01:24:01.00Z", "id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "verbose,localTime")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--verbose", "--localTime"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)
	err = registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3 sensor source 2020-09-01T10:24:01.00+0900\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)
	err = registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdCountPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"101"}}
	reqRes1.Path = "/ngsi-ld/v1/csourceRegistrations/"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes2.ResHeader = http.Header{"Ngsild-Results-Count": []string{"101"}}
	reqRes2.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--verbose"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3 sensor source \n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
func TestRegistrationsListLdJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\"}]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdJSONPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"id\": \"5f5dcb551e715bc7f1ad79e3\",\n    \"type\": \"ContextSourceRegistration\",\n    \"description\": \"sensor source\",\n    \"endpoint\": \"http://raspi\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=AirQualityObserved"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type []ngsicmd.cSourceRegistration Field: (1) {}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)
	setJSONEncodeErr(ngsi, 0)

	err = registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorJSONPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--json", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"id\": \"5f5dcb551e715bc7f1ad79e3\",\n  \"type\": \"ContextSourceRegistration\",\n  \"description\": \"sensor source\",\n  \"endpoint\": \"http://raspi\"\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetLocalTime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"expires": "2020-09-01T01:24:01.00Z", "id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "localTime")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--localTime", "--id=5f5dcb551e715bc7f1ad79e3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"expires\":\"2020-09-01T10:24:01.00+0900\",\"endpoint\":\"http://raspi\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetSafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,safeString")
	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--id=5f5dcb551e715bc7f1ad79e3"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetLdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setJSONDecodeErr(ngsi, 1)

	setupFlagString(set, "host,id,safeString")
	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--id=5f5dcb551e715bc7f1ad79e3"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setJSONEncodeErr(ngsi, 2)
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3", "--pretty"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	setJSONIndentError(ngsi)

	err = registrationsGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={}"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsCreateLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLdErrorSetRegistrationsValuleLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/registrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data="})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLdErrorJSONMarshalEncode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setJSONEncodeErr(ngsi, 2)
	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={}"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={}"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={}"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsDeleteLd(c, ngsi, client)

	assert.NoError(t, err)
}

func TestRegistrationsDeleteLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsDeleteLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteLdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = registrationsDeleteLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLdArgs(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "description,providedId,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--providedId=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	err := registrationsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"description\":\"test\",\"registrationInfo\":[{\"entities\":[{\"id\":\"device001\",\"type\":\"Device\"}]}],\"endpoint\":\"http://provider\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLdArgsPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "description,providedId,type,attrs,provider")
	setupFlagBool(set, "pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--pretty", "--description=test", "--providedId=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	err := registrationsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"description\": \"test\",\n  \"registrationInfo\": [\n    {\n      \"entities\": [\n        {\n          \"id\": \"device001\",\n          \"type\": \"Device\"\n        }\n      ]\n    }\n  ],\n  \"endpoint\": \"http://provider\"\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLdErrorProvider(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,providedId,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--providedId=device001", "--type=Device", "--attrs=abc,xyz", "--provider=provider"})
	err := registrationsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "provider url error: provider", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLdErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,providedId,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--providedId=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	setJSONEncodeErr(ngsi, 0)

	err := registrationsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLdErrorArgsPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,providedId,type,attrs,provider")
	setupFlagBool(set, "pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--pretty", "--description=test", "--providedId=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})

	setJSONIndentError(ngsi)

	err := registrationsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetRegistrationsValuleLd1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,description,provider,expires")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--description=reg", "--provider=http://csource", "--expires=2020-12-01T19:17:35.000Z"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"description\":\"reg\",\"expires\":\"2020-12-01T19:17:35.000Z\",\"endpoint\":\"http://csource\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLd2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "type,providedId,idPattern")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--type=device", "--providedId=device001", "--idPattern=.*"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"registrationInfo\":[{\"entities\":[{\"id\":\"device001\",\"idPattern\":\".*\",\"type\":\"device\"}]}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLd3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "properties,relationships")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--properties=speed", "--relationships=urn:ngsi-ld:car"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"registrationInfo\":[{\"entities\":[{}],\"properties\":[\"speed\"],\"relationships\":[\"urn:ngsi-ld:car\"]}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLdContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,context")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--context=[\"http://context\"]"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"@context\":[\"http://context\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLdErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,description,provider,expires")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=", "--description=reg", "--provider=http://csource", "--expires=day"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setJSONDecodeErr(ngsi, 0)

	setupFlagString(set, "data,description,provider,expires")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--description=reg", "--provider=http://csource", "--expires=day"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorProvider(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,description,provider,expires")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--description=reg", "--provider=csource", "--expires=1day"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "provider url error: csource", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,context")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--context=[\"http://context\""})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorExpires(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,description,provider,expires")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--description=reg", "--provider=http://csource", "--expires=day"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error day", ngsiErr.Message)
	}
}
