/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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

func TestAppsRolePermList(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "permissions", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_permission_assignments\":[{\"id\":\"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\"is_internal\":false,\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\",\"action\":null,\"resource\":null,\"xml\":\"xmlrule\"},{\"id\":\"3\",\"is_internal\":true,\"name\":\"Manageroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"1\",\"is_internal\":true,\"name\":\"Getandassignallinternalapplicationroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermListNotFound(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "permissions", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "Assignments not found\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "permissions", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_permission_assignments\": [\n    {\n      \"id\": \"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\n      \"is_internal\": false,\n      \"name\": \"newnamepermission\",\n      \"description\": \"newdescriptionpermission\",\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": \"xmlrule\"\n    },\n    {\n      \"id\": \"3\",\n      \"is_internal\": true,\n      \"name\": \"Manageroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"1\",\n      \"is_internal\": true,\n      \"name\": \"Getandassignallinternalapplicationroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "permissions", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsRolePermListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "permissions", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsRolePermListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "permissions", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsRolePermList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsRolePermAssign(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermAssign(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_permission_assignments\":{\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"permission_id\":\"4\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermAssignPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermAssign(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_permission_assignments\": {\n    \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n    \"permission_id\": \"4\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermAssignErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsRolePermAssignErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsRolePermAssignErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsRolePermAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsRolePermDelete(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAppsRolePermDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsRolePermDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsRolePermDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
