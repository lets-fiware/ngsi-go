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

func TestOrgUsersList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"organization_users\":[{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role\":\"owner\"},{\"user_id\":\"admin\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role\":\"member\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"organization_users\": [\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role\": \"owner\"\n    },\n    {\n      \"user_id\": \"admin\",\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role\": \"member\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := orgUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := orgUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersListErrorOid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := orgUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	setJSONIndentError(ngsi)

	err := orgUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrgUsersListErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	setJSONDecodeErr(ngsi, 1)
	err := orgUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrgUsersGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_user":{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin"})

	err := orgUsersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"organization_user\":{\"user_id\":\"admin\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role\":\"member\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_user":{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin"})

	err := orgUsersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"organization_user\": {\n    \"user_id\": \"admin\",\n    \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n    \"role\": \"member\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := orgUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := orgUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersGetErrorOid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := orgUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersGetErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin"})

	err := orgUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin"})

	err := orgUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles"
	reqRes.ResBody = []byte(`{"organization_users":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"owner"},{"user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role":"member"}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin"})

	setJSONIndentError(ngsi)

	err := orgUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestOrgUsersCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	err := orgUsersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"user_organization_assignments\":{\"role\":\"owner\",\"user_id\":\"admin\",\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	err := orgUsersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"user_organization_assignments\": {\n    \"role\": \"owner\",\n    \"user_id\": \"admin\",\n    \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestOrgUsersCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersCreateErrorNewClient(t *testing.T) {
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

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersCreateErrorOid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersCreateErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersCreateErrorOrid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin"})

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ReqData = []byte(``)
	reqRes.ResBody = []byte(`{"user_organization_assignments":{"role":"owner","user_id":"admin","organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	setJSONIndentError(ngsi)

	err := orgUsersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	err := orgUsersDelete(c)

	assert.NoError(t, err)
}

func TestOrgUsersDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := orgUsersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := orgUsersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersDeleteErrorOid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := orgUsersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersDeleteErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd"})

	err := orgUsersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersDeleteErrorOrid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin"})

	err := orgUsersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	err := orgUsersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOrgUsersDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/organizations/3e20722f-d420-422d-89ba-3ae87bc1c0cd/users/admin/organization_roles/owner"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid,uid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--oid=3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid=admin", "--orid=owner"})

	err := orgUsersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
