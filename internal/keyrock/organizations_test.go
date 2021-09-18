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

package keyrock

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestOrganizationsList(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsListVerbose(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"organizations\":[{\"role\":\"owner\",\"Organization\":{\"id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"name\":\"Testorganization2\",\"description\":\"description2\",\"image\":\"default\",\"website\":null}},{\"role\":\"owner\",\"Organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"name\":\"Testorganization\",\"description\":\"description\",\"image\":\"default\",\"website\":null}}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsListPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"organizations\": [\n    {\n      \"role\": \"owner\",\n      \"Organization\": {\n        \"id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n        \"name\": \"Testorganization2\",\n        \"description\": \"description2\",\n        \"image\": \"default\",\n        \"website\": null\n      }\n    },\n    {\n      \"role\": \"owner\",\n      \"Organization\": {\n        \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n        \"name\": \"Testorganization\",\n        \"description\": \"description\",\n        \"image\": \"default\",\n        \"website\": null\n      }\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsListNotFound(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/organizations"

	helper.SetClientHTTP(c, reqRes)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "Organizations not found\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestOrganizationsListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestOrganizationsListErrorPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsListErrorID(t *testing.T) {
	c := setupTest([]string{"organizations", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations"
	reqRes.ResBody = []byte(`{"organizations":[{"role":"owner","Organization":{"id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","name":"Testorganization2","description":"description2","image":"default","website":null}},{"role":"owner","Organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","image":"default","website":null}}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := organizationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsGet(t *testing.T) {
	c := setupTest([]string{"organizations", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"name\":\"Testorganization\",\"description\":\"description\",\"website\":null,\"image\":\"default\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsGetPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"organization\": {\n    \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"name\": \"Testorganization\",\n    \"description\": \"description\",\n    \"website\": null,\n    \"image\": \"default\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := organizationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrganizationsGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestOrganizationsGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","name":"Testorganization","description":"description","website":null,"image":"default"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := organizationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsCreate(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3e20722f-d420-422d-89ba-3ae87bc1c0cd\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsCreateVerbose(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"image\":\"default\",\"name\":\"Testorganization\",\"description\":\"description\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsCreatePretty(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"organization\": {\n    \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"image\": \"default\",\n    \"name\": \"Testorganization\",\n    \"description\": \"description\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsCreateErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrganizationsCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestOrganizationsCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsCreateErrorID(t *testing.T) {
	c := setupTest([]string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := organizationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsUpdate(t *testing.T) {
	c := setupTest([]string{"organizations", "update", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"organization\":{\"id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"image\":\"default\",\"name\":\"Testorganization\",\"description\":\"description\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsUpdatePretty(t *testing.T) {
	c := setupTest([]string{"organizations", "update", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"organization\": {\n    \"id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"image\": \"default\",\n    \"name\": \"Testorganization\",\n    \"description\": \"description\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrganizationsUpdateErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"organizations", "update", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := organizationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "update", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := organizationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrganizationsUpdateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "update", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestOrganizationsUpdateErrorPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "update", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--name", "Testorganization", "--description", "description", "--website", "http://keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ReqData = []byte(`{"organization":{"name":"Testorganization","description":"description","website":"http://keyrock"}}`)
	reqRes.ResBody = []byte(`{"organization":{"id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","image":"default","name":"Testorganization","description":"description"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := organizationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrganizationsDelete(t *testing.T) {
	c := setupTest([]string{"organizations", "delete", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"

	helper.SetClientHTTP(c, reqRes)

	err := organizationsDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestOrganizationsDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "delete", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := organizationsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrganizationsDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "delete", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := organizationsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
