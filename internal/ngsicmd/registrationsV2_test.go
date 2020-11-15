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

func TestRegistrationsListV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2CountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2Page(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"101"}}
	reqRes1.Path = "/v2/registrations"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"101"}}
	reqRes2.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2Verbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "verbose")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3 sensor source\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
func TestRegistrationsListV2JSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "json")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--json"})
	err := registrationsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"description\":\"sensor source\",\"endpoint\":\"http://raspi\",\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"information\":[{\"entities\":[{\"id\":\"urn:ngsi-ld:Device:device001\",\"type\":\"Device\"}],\"properties\":[\"temperature\",\"pressure\",\"humidity\"]}],\"type\":\"ContextSourceRegistration\"}]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type []map[string]interface {}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListV2ErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "json")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--json"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := registrationsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsGetV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\",\"information\":[{\"entities\":[{\"id\":\"urn:ngsi-ld:Device:device001\",\"type\":\"Device\"}],\"properties\":[\"temperature\",\"pressure\",\"humidity\"]}],\"type\":\"ContextSourceRegistration\"}\n"
		assert.Equal(t, expected, actual, client)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2SafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,id,safeString")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	err := registrationsGetV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"description\":\"sensor source\",\"endpoint\":\"http://raspi\",\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"information\":[{\"entities\":[{\"id\":\"urn:ngsi-ld:Device:device001\",\"type\":\"Device\"}],\"properties\":[\"temperature\",\"pressure\",\"humidity\"]}],\"type\":\"ContextSourceRegistration\"}\n"
		assert.Equal(t, expected, actual, client)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetErrorV2StatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetV2ErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")
	setupFlagString(set, "host,id,safeString")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	err := registrationsGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/registrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2ErrorReadALl(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/registrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := registrationsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsDeleteV2(c, ngsi, client)

	assert.NoError(t, err)
}

func TestRegistrationsDeleteV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsDeleteV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--host=orion", "--id=5f5dcb551e715bc7f1ad79e3"})
	err := registrationsDeleteV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsTemplateV2(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"description\":\"Registration template\",\"dataProvided\":{\"entities\":[{\"id\":\"\",\"type\":\"Room\"}],\"attrs\":[\"attr\"]},\"provider\":{\"http\":{\"url\":\"http://localhost:1234\"}}}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateV2Args(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	err := registrationsTemplateV2(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"description\":\"test\",\"dataProvided\":{\"entities\":[{\"id\":\"device001\",\"type\":\"Device\"}],\"attrs\":[\"abc\",\"xyz\"]},\"provider\":{\"http\":{\"url\":\"http://provider\"}}}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateV2ErrorProvider(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=provider"})
	err := registrationsTemplateV2(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "provider url error: provider", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateV2ErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "description,id,type,attrs,provider")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--description=test", "--id=device001", "--type=Device", "--attrs=abc,xyz", "--provider=http://provider"})
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	err := registrationsTemplateV2(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
