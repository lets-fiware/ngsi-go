/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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

func TestAppsOrgsRolesList(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesListVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_organization_assignments\":[{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_organization\":\"owner\",\"role_id\":\"purchaser\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_organization\":\"owner\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_organization\":\"member\",\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\"},{\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role_organization\":\"owner\",\"role_id\":\"provider\"},{\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role_organization\":\"owner\",\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\"},{\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role_organization\":\"member\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_organization_assignments\": [\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"purchaser\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_organization\": \"member\",\n      \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\"\n    },\n    {\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"provider\"\n    },\n    {\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\"\n    },\n    {\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role_organization\": \"member\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsOrgsRolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesListErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := appsOrgsRolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesGet(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_organization_assignments\":[{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_id\":\"purchaser\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesGetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_organization_assignments\": [\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_id\": \"purchaser\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsOrgsRolesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesAssign(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesAssign(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_organization_assignments\":{\"role_id\":\"provider\",\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"role_organization\":\"owner\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesAssignPretty(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesAssign(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_organization_assignments\": {\n    \"role_id\": \"provider\",\n    \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n    \"role_organization\": \"owner\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesAssignErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesAssignErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesAssignErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsOrgsRolesAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesUnassign(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesUnassign(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAppsOrgsRolesUnassignErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesUnassign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesUnassignErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "organizations", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsOrgsRolesUnassign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
