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

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestAppsUsersList(t *testing.T) {
	c := setupTest([]string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersListVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_user_assignments\":[{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"provider\"},{\"user_id\":\"admin\",\"role_id\":\"purchaser\"},{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"},{\"user_id\":\"admin\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"}]} "
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_user_assignments\": [\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"provider\"\n    },\n    {\n      \"user_id\": \"admin\",\n      \"role_id\": \"purchaser\"\n    },\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    },\n    {\n      \"user_id\": \"admin\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    }\n  ]\n} \n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsUsersListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsUsersListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersListErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := appsUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersGet(t *testing.T) {
	c := setupTest([]string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "provider\nee2ec16f-694b-447f-b61a-e293b6fe5f7b\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersGetVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_user_assignments\":[{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"provider\"},{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersGetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_user_assignments\": [\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"provider\"\n    },\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsUsersGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsUsersGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsUsersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersGetErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := appsUsersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersAssign(t *testing.T) {
	c := setupTest([]string{"applications", "users", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersAssign(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role_user_assignments\":{\"role_id\":\"purchaser\",\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersAssignPretty(t *testing.T) {
	c := setupTest([]string{"applications", "users", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersAssign(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role_user_assignments\": {\n    \"role_id\": \"purchaser\",\n    \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersAssignErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "users", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsUsersAssignErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "users", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestAppsUsersAssignErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "users", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := appsUsersAssign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersUnassign(t *testing.T) {
	c := setupTest([]string{"applications", "users", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersUnassign(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestAppsUsersUnassignErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "users", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersUnassign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestAppsUsersUnassignErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "users", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := appsUsersUnassign(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
