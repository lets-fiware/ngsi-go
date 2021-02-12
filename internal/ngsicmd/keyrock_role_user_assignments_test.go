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

func TestAppsUsersList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_user_assignments\":[{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"provider\"},{\"user_id\":\"admin\",\"role_id\":\"purchaser\"},{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"},{\"user_id\":\"admin\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"}]} "
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_user_assignments\": [\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"provider\"\n    },\n    {\n      \"user_id\": \"admin\",\n      \"role_id\": \"purchaser\"\n    },\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    },\n    {\n      \"user_id\": \"admin\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    }\n  ]\n} \n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/appsUsers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersListErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	reqRes.Err = errors.New("error")
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	setJSONIndentError(ngsi)

	err := appsUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersListErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"admin","role_id":"purchaser"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"},{"user_id":"admin","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]} `)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	setJSONDecodeErr(ngsi, 1)

	err := appsUsersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := appsUsersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "provider\nee2ec16f-694b-447f-b61a-e293b6fe5f7b\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersGetVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := appsUsersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_user_assignments\":[{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"provider\"},{\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"role_id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := appsUsersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_user_assignments\": [\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"provider\"\n    },\n    {\n      \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"role_id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/appsUsers"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersGetErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersGetErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	setJSONIndentError(ngsi)

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersGetErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles"
	reqRes.ResBody = []byte(`{"role_user_assignments":[{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"provider"},{"user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","role_id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b"}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	setJSONDecodeErr(ngsi, 1)

	err := appsUsersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestAppsUsersAssign(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	err := appsUsersAssign(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"role_user_assignments\":{\"role_id\":\"purchaser\",\"user_id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersAssignPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	err := appsUsersAssign(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"role_user_assignments\": {\n    \"role_id\": \"purchaser\",\n    \"user_id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestAppsUsersAssignErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsUsersAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersAssignErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsUsersAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersAssignErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsUsersAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersAssignErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersAssignErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := appsUsersAssign(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersAssignErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	err := appsUsersAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersAssignErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	err := appsUsersAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersAssignErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ReqData = []byte("")
	reqRes.ResBody = []byte(`{"role_user_assignments":{"role_id":"purchaser","user_id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	setJSONIndentError(ngsi)

	err := appsUsersAssign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersUnassign(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	err := appsUsersUnassign(c)

	assert.NoError(t, err)
}

func TestAppsUsersUnassignErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := appsUsersUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersUnassignErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := appsUsersUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersUnassignErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := appsUsersUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersUnassignErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := appsUsersUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersUnassignErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := appsUsersUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify role id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersUnassignErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	err := appsUsersUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAppsUsersUnassignErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb/roles/purchaser"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,uid,rid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid=purchaser"})

	err := appsUsersUnassign(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
