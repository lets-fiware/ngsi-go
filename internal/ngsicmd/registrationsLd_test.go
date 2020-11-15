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

func TestRegistrationsListLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsListLd(c, ngsi, client)

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

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"101"}}
	reqRes1.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes2.ResHeader = http.Header{"Ngsild-Results-Count": []string{"101"}}
	reqRes2.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsListLd(c, ngsi, client)

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

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "verbose")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--verbose"})
	err := registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3 sensor source\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
func TestRegistrationsListLdJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "json")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--json"})
	err := registrationsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"description\":\"sensor source\",\"endpoint\":\"http://raspi\",\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"information\":[{\"entities\":[{\"id\":\"urn:ngsi-ld:Device:device001\",\"type\":\"Device\"}],\"properties\":[\"temperature\",\"pressure\",\"humidity\"]}],\"type\":\"ContextSourceRegistration\"}]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsListLd(c, ngsi, client)

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

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--type=AirQualityObserved"})
	err := registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type []map[string]interface {}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListLdErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "json")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--json"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := registrationsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\",\"information\":[{\"entities\":[{\"id\":\"urn:ngsi-ld:Device:device001\",\"type\":\"Device\"}],\"properties\":[\"temperature\",\"pressure\",\"humidity\"]}],\"type\":\"ContextSourceRegistration\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetSafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")
	setupFlagString(set, "host,id,safeString")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--id=5f5dcb551e715bc7f1ad79e3"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	err := registrationsGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"description\":\"sensor source\",\"endpoint\":\"http://raspi\",\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"information\":[{\"entities\":[{\"id\":\"urn:ngsi-ld:Device:device001\",\"type\":\"Device\"}],\"properties\":[\"temperature\",\"pressure\",\"humidity\"]}],\"type\":\"ContextSourceRegistration\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsGetLd(c, ngsi, client)

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

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsLdGetErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")
	setupFlagString(set, "host,id,safeString")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	_ = set.Parse([]string{"--host=orion-ld", "--safeString=on", "--id=5f5dcb551e715bc7f1ad79e3"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := registrationsGetLd(c, ngsi, client)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--data={}"})
	err := registrationsCreateLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLdErrorReadALl(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/registrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--data={}"})
	err := registrationsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateLdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--data={}"})
	err := registrationsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsDeleteLd(c, ngsi, client)

	assert.NoError(t, err)
}

func TestRegistrationsDeleteLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/csourceRegistrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsDeleteLd(c, ngsi, client)

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

	setupAddBroker(t, ngsi, "orion-ld", "https://orion", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion-ld", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsDeleteLd(c, ngsi, client)

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
		expected := "{\"description\":\"registration template\",\"endpoint\":\"http://registration\",\"information\":[{\"entities\":[{\"id\":\"urn:ngsi-ld:Registration:001\",\"type\":\"Registration\"}],\"properties\":[\"attr\"]}],\"type\":\"ContextSourceRegistration\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLdArgs(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	err := registrationsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"description\":\"test\",\"endpoint\":\"http://provider\",\"information\":[{\"entities\":[{\"id\":\"device001\",\"type\":\"Device\"}],\"properties\":[\"abc\",\"xyz\"]}],\"type\":\"ContextSourceRegistration\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateLdErrorProvider(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=provider"})
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

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := registrationsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
