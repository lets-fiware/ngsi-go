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

func TestOrganizationsList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := organizationsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose"})

	err := organizationsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"organizations\":[{\"role\":\"owner\",\"Organization\":{\"id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"name\":\"Testorganization2\",\"description\":\"description2\",\"image\":\"default\",\"website\":null}},{\"role\":\"owner\",\"Organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"name\":\"Testorganization\",\"description\":\"description\",\"image\":\"default\",\"website\":null}}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty"})

	err := organizationsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"organizations\": [\n    {\n      \"role\": \"owner\",\n      \"Organization\": {\n        \"id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n        \"name\": \"Testorganization2\",\n        \"description\": \"description2\",\n        \"image\": \"default\",\n        \"website\": null\n      }\n    },\n    {\n      \"role\": \"owner\",\n      \"Organization\": {\n        \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n        \"name\": \"Testorganization\",\n        \"description\": \"description\",\n        \"image\": \"default\",\n        \"website\": null\n      }\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsListNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/organizations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := organizationsList(c)

	assert.NoError(t, err)
}

func TestOrganizationsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := organizationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := organizationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := organizationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestOrganizationsListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := organizationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty"})

	setJSONIndentError(ngsi)

	err := organizationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsListErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	setJSONDecodeErr(ngsi, 1)
	err := organizationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"name\":\"Testorganization\",\"description\":\"description\",\"website\":null,\"image\":\"default\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"organization\": {\n    \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"name\": \"Testorganization\",\n    \"description\": \"description\",\n    \"website\": null,\n    \"image\": \"default\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := organizationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := organizationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsGetErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := organizationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	setJSONIndentError(ngsi)

	err := organizationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3e20722f-d420-422d-89ba-3ae87bc1c0cd\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsCreateVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"image\":\"default\",\"name\":\"Testorganization\",\"description\":\"description\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"organization\": {\n    \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"image\": \"default\",\n    \"name\": \"Testorganization\",\n    \"description\": \"description\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsCreateErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty"})

	setJSONEncodeErr(ngsi, 2)

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization name", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsCreateErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	setJSONEncodeErr(ngsi, 2)

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	setJSONIndentError(ngsi)

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsCreateErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name,description,website")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	setJSONDecodeErr(ngsi, 1)

	err := organizationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,name,description,website")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"image\":\"default\",\"name\":\"Testorganization\",\"description\":\"description\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"organization\": {\n    \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"image\": \"default\",\n    \"name\": \"Testorganization\",\n    \"description\": \"description\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := organizationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := organizationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := organizationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsUpdateErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	setJSONEncodeErr(ngsi, 2)

	err := organizationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,name,description,website")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	err := organizationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,name,description,website")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name=Testorganization", "--description=description", "--website=http://keyrock"})

	setJSONIndentError(ngsi)

	err := organizationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsDelete(c)

	assert.NoError(t, err)
}

func TestOrganizationsDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := organizationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := organizationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsDeleteErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := organizationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrganizationsDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := organizationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
