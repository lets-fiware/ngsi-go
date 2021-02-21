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

func TestAppsOrgsRolesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n33cf4d3c-8dfb-4bed-bf37-7647f45528ec\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n3e20722f-d420-422d-89ba-3ae87bc1c0cd\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_organization_assignments\":[{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_organization\":\"owner\",\"role_id\":\"purchaser\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_organization\":\"owner\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_organization\":\"member\",\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\"},{\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role_organization\":\"owner\",\"role_id\":\"provider\"},{\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role_organization\":\"owner\",\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\"},{\"organization_id\":\"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\"role_organization\":\"member\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_organization_assignments\": [\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"purchaser\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_organization\": \"member\",\n      \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\"\n    },\n    {\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"provider\"\n    },\n    {\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role_organization\": \"owner\",\n      \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\"\n    },\n    {\n      \"organization_id\": \"3e20722f-d420-422d-89ba-3ae87bc1c0cd\",\n      \"role_organization\": \"member\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsOrgsRolesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/appsOrgsRoles"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsOrgsRolesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesListErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsOrgsRolesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	reqRes.Err = errors.New("error")
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	setJSONIndentError(ngsi)

	err := appsOrgsRolesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesListErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"owner","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_organization":"member","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"provider"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"owner","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"},{"organization_id":"3e20722f-d420-422d-89ba-3ae87bc1c0cd","role_organization":"member","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	setJSONDecodeErr(ngsi, 1)

	err := appsOrgsRolesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	err := appsOrgsRolesGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_organization_assignments\":[{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_id\":\"purchaser\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"},{\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"role_id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	err := appsOrgsRolesGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_organization_assignments\": [\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_id\": \"purchaser\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    },\n    {\n      \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n      \"role_id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsOrgsRolesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/appsOrgsRoles"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsOrgsRolesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesGetErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsOrgsRolesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesGetErrorOid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	err := appsOrgsRolesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	err := appsOrgsRolesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles"
	reqRes.ResBody = []byte(`{"role_organization_assignments":[{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"purchaser"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","role_id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	setJSONIndentError(ngsi)

	err := appsOrgsRolesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsOrgsRolesAssign(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	err := appsOrgsRolesAssign(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_organization_assignments\":{\"role_id\":\"provider\",\"organization_id\":\"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"role_organization\":\"owner\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesAssignPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	err := appsOrgsRolesAssign(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_organization_assignments\": {\n    \"role_id\": \"provider\",\n    \"organization_id\": \"33cf4d3c-8dfb-4bed-bf37-7647f45528ec\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n    \"role_organization\": \"owner\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsOrgsRolesAssignErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsOrgsRolesAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsOrgsRolesAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsOrgsRolesAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorOid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	err := appsOrgsRolesAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorOrid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	reqRes.Err = errors.New("error")
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider"})

	err := appsOrgsRolesAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_organization_assignments":{"role_id":"provider","organization_id":"33cf4d3c-8dfb-4bed-bf37-7647f45528ec","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025","role_organization":"owner"}}`)
	reqRes.Err = errors.New("error")
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	err := appsOrgsRolesAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	err := appsOrgsRolesAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesAssignErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	setJSONIndentError(ngsi)

	err := appsOrgsRolesAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassign(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	err := appsOrgsRolesUnassign(c)

	assert.NoError(t, err)
}

func TestAppsOrgsRolesUnassignErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassignErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassignErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassignErrorOid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassignErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec"})

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassignErrorOrid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider"})

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "specify organization role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassignErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsOrgsRolesUnassignErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/organizations/33cf4d3c-8dfb-4bed-bf37-7647f45528ec/roles/provider/organization_roles/owner"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,oid,rid,orid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid=33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid=provider", "--orid=owner"})

	err := appsOrgsRolesUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
