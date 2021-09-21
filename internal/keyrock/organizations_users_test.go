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

func TestOrgUsersList(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersListVerbose(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"organization_users\":[{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role\":\"owner\"},{\"user_id\":\"admin\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role\":\"member\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersListPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"organization_users\": [\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role\": \"owner\"\n    },\n    {\n      \"user_id\": \"admin\",\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role\": \"member\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrgUsersListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	helper.SetClientHTTP(c, reqRes)

	err := orgUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestOrgUsersListErrorPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := orgUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrgUsersListErrorID(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := orgUsersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrgUsersGet(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_user":{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"organization_user\":{\"user_id\":\"admin\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role\":\"member\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersGetPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_user":{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"organization_user\": {\n    \"user_id\": \"admin\",\n    \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"role\": \"member\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrgUsersGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestOrgUsersGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := orgUsersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrgUsersCreate(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "add", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"user_organization_assignments\":{\"role\":\"owner\",\"user_id\":\"admin\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersCreatePretty(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "add", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"user_organization_assignments\": {\n    \"role\": \"owner\",\n    \"user_id\": \"admin\",\n    \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "add", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrgUsersCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "add", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestOrgUsersCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "add", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := orgUsersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrgUsersDelete(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "remove", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestOrgUsersDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "remove", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestOrgUsersDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"organizations", "users", "remove", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := orgUsersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
