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
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestRegistrationsListLd(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListLdLocalTime(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld", "--verbose", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"expires": "2020-09-01T01:24:01.00Z", "id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3 sensor source 2020-09-01T10:24:01.00+0900\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListLdCountZero(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListLdCountZeroPretty(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListLdCountPage(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"101"}}
	reqRes1.Path = "/ngsi-ld/v1/csourceRegistrations/"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes2.ResHeader = http.Header{"Ngsild-Results-Count": []string{"101"}}
	reqRes2.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3\n5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListLdVerbose(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3 sensor source \n"
		assert.Equal(t, expected, actual)
	}
}
func TestRegistrationsListLdJSON(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\"}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListLdJSONPretty(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"5f5dcb551e715bc7f1ad79e3\",\n    \"type\": \"ContextSourceRegistration\",\n    \"description\": \"sensor source\",\n    \"endpoint\": \"http://raspi\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsListLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsListLdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestRegistrationsListLdErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestRegistrationsListLdErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type []ngsicmd.cSourceRegistration Field: (1) {}", ngsiErr.Message)
	}
}

func TestRegistrationsListLdErrorJSON(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsListLdErrorJSONPretty(t *testing.T) {
	c := setupTest([]string{"list", "registrations", "--host", "orion-ld", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}]`)
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := registrationsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsLdGet(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsLdGetPretty(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"5f5dcb551e715bc7f1ad79e3\",\n  \"type\": \"ContextSourceRegistration\",\n  \"description\": \"sensor source\",\n  \"endpoint\": \"http://raspi\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsLdGetLocalTime(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"expires": "2020-09-01T01:24:01.00Z", "id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"expires\":\"2020-09-01T10:24:01.00+0900\",\"endpoint\":\"http://raspi\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsLdGetSafeString(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"5f5dcb551e715bc7f1ad79e3\",\"type\":\"ContextSourceRegistration\",\"description\":\"sensor source\",\"endpoint\":\"http://raspi\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsGetLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsGetLdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "5f5dcb551e715bc7f1ad79e3  error", ngsiErr.Message)
	}
}

func TestRegistrationsLdGetErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsLdGetErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsLdGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"5f5dcb551e715bc7f1ad79e3","description":"sensor source","endpoint":"http://raspi","information":[{"entities":[{"id":"urn:ngsi-ld:Device:device001","type":"Device"}],"properties":["temperature","pressure","humidity"]}],"type":"ContextSourceRegistration"}`)
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := registrationsGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateLd(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"}}
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreateLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f5dcb551e715bc7f1ad79e3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsCreateLdErrorSetRegistrationsValuleLd(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion-ld", "--data", "@"})

	err := registrationsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateLdErrorJSONMarshalEncode(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion-ld", "--data", "{}"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsCreateLdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"create", "registration", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	}
}

func TestRegistrationsDeleteLd(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDeleteLd(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRegistrationsDeleteLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDeleteLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRegistrationsDeleteLdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"delete", "registration", "--host", "orion-ld", "--id", "5f5dcb551e715bc7f1ad79e3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations/5f5dcb551e715bc7f1ad79e3"
	reqRes.ResBody = []byte(`error`)

	helper.SetClientHTTP(c, reqRes)

	err := registrationsDeleteLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "5f5dcb551e715bc7f1ad79e3  error", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateLd(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "{}"})

	err := registrationsTemplateLd(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsTemplateLdArgs(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	err := registrationsTemplateLd(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"description\":\"test\",\"registrationInfo\":[{\"entities\":[{\"id\":\"device001\",\"type\":\"Device\"}]}],\"endpoint\":\"http://provider\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsTemplateLdArgsPretty(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--pretty", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	err := registrationsTemplateLd(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"description\": \"test\",\n  \"registrationInfo\": [\n    {\n      \"entities\": [\n        {\n          \"id\": \"device001\",\n          \"type\": \"Device\"\n        }\n      ]\n    }\n  ],\n  \"endpoint\": \"http://provider\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRegistrationsTemplateLdErrorProvider(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "provider"})

	err := registrationsTemplateLd(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "provider url error: provider", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateLdErrorMarshal(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := registrationsTemplateLd(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRegistrationsTemplateLdErrorArgsPretty(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--pretty", "--description", "test", "--providedId", "device001", "--type", "Device", "--attrs", "abc,xyz", "--provider", "http://provider"})

	helper.SetJSONIndentError(c.Ngsi)

	err := registrationsTemplateLd(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLd1(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "{}", "--description", "reg", "--provider", "http://csource", "--expires", "2020-12-01T19:17:35.000Z"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"description\":\"reg\",\"expires\":\"2020-12-01T19:17:35.000Z\",\"endpoint\":\"http://csource\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLd2(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--type", "device", "--providedId", "device001", "--idPattern", ".*"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"registrationInfo\":[{\"entities\":[{\"id\":\"device001\",\"idPattern\":\".*\",\"type\":\"device\"}]}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLd3(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--properties", "speed", "--relationships", "urn:ngsi-ld:car"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"registrationInfo\":[{\"entities\":[{}],\"properties\":[\"speed\"],\"relationships\":[\"urn:ngsi-ld:car\"]}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLdContext(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "{}", "--context", "[\"http://context\"]"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.NoError(t, err) {
		b, _ := json.Marshal(r)
		actual := string(b)
		expected := "{\"@context\":[\"http://context\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetRegistrationsValuleLdErrorReadAll(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "@", "--description", "reg", "--provider", "http://csource", "--expires", "day"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "{}", "--description", "reg", "--provider", "http://csource", "--expires", "day"})

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorExpires(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "{}", "--description", "reg", "--provider", "http://csource", "--expires", "day"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error day", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorProvider(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "{}", "--description", "reg", "--provider", "csource", "--expires", "1day"})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "provider url error: csource", ngsiErr.Message)
	}
}

func TestSetRegistrationsValuleLdErrorContext(t *testing.T) {
	c := setupTest([]string{"template", "registration", "--ngsiType", "ld", "--data", "{}", "--context", "[\"http://context\""})

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, c.Ngsi, &r)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}
