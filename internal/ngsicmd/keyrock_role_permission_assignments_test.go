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

func TestAppsRolePermList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := appsRolePermList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_permission_assignments\":[{\"id\":\"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\"is_internal\":false,\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\",\"action\":null,\"resource\":null,\"xml\":\"xmlrule\"},{\"id\":\"3\",\"is_internal\":true,\"name\":\"Manageroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"1\",\"is_internal\":true,\"name\":\"Getandassignallinternalapplicationroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermListNotFound(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := appsRolePermList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "Assignments not found\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := appsRolePermList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_permission_assignments\": [\n    {\n      \"id\": \"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\n      \"is_internal\": false,\n      \"name\": \"newnamepermission\",\n      \"description\": \"newdescriptionpermission\",\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": \"xmlrule\"\n    },\n    {\n      \"id\": \"3\",\n      \"is_internal\": true,\n      \"name\": \"Manageroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"1\",\n      \"is_internal\": true,\n      \"name\": \"Getandassignallinternalapplicationroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsRolePermList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/appsRolePerm"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsRolePermList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermListErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsRolePermList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermListErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsRolePermList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	reqRes.Err = errors.New("error")
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := appsRolePermList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := appsRolePermList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions"
	reqRes.ResBody = []byte(`{"role_permission_assignments":[{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"newnamepermission","description":"newdescriptionpermission","action":null,"resource":null,"xml":"xmlrule"},{"id":"3","is_internal":true,"name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"1","is_internal":true,"name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	setJSONIndentError(ngsi)

	err := appsRolePermList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsRolePermAssign(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := appsRolePermAssign(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_permission_assignments\":{\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"permission_id\":\"4\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermAssignPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := appsRolePermAssign(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_permission_assignments\": {\n    \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n    \"permission_id\": \"4\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsRolePermAssignErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsRolePermAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermAssignErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsRolePermAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermAssignErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsRolePermAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermAssignErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsRolePermAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermAssignErrorPid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := appsRolePermAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify permission id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermAssignErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := appsRolePermAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermAssignErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := appsRolePermAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermAssignErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"role_permission_assignments":{"role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","permission_id":"4"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	setJSONIndentError(ngsi)

	err := appsRolePermAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := appsRolePermDelete(c)

	assert.NoError(t, err)
}

func TestAppsRolePermDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsRolePermDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsRolePermDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermDeleteErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsRolePermDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermDeleteErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsRolePermDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermDeleteErrorPid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := appsRolePermDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify permission id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := appsRolePermDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsRolePermDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91/permissions/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid=33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := appsRolePermDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
