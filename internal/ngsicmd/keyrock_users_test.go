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

func TestUsersList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := usersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose"})

	err := usersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"users\":[{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"username\":\"alice\",\"email\":\"alice@test.com\",\"enabled\":true,\"gravatar\":false,\"date_password\":\"2018-03-20T09:31:07.000Z\",\"description\":null,\"website\":null},{\"id\":\"admin\",\"username\":\"admin\",\"email\":\"admin@test.com\",\"enabled\":true,\"gravatar\":false,\"date_password\":\"2018-03-20T08:40:14.000Z\",\"description\":null,\"website\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty"})

	err := usersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"users\": [\n    {\n      \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"username\": \"alice\",\n      \"email\": \"alice@test.com\",\n      \"enabled\": true,\n      \"gravatar\": false,\n      \"date_password\": \"2018-03-20T09:31:07.000Z\",\n      \"description\": null,\n      \"website\": null\n    },\n    {\n      \"id\": \"admin\",\n      \"username\": \"admin\",\n      \"email\": \"admin@test.com\",\n      \"enabled\": true,\n      \"gravatar\": false,\n      \"date_password\": \"2018-03-20T08:40:14.000Z\",\n      \"description\": null,\n      \"website\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := usersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := usersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := usersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestUsersListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := usersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty"})

	setJSONIndentError(ngsi)

	err := usersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersListErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	setJSONDecodeErr(ngsi, 1)
	err := usersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"user\":{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"username\":\"alice\",\"email\":\"alice@test.com\",\"enabled\":true,\"admin\":false,\"image\":\"default\",\"gravatar\":false,\"date_password\":\"2018-03-20T09:31:07.000Z\",\"description\":null,\"website\":null}}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"user\": {\n    \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"username\": \"alice\",\n    \"email\": \"alice@test.com\",\n    \"enabled\": true,\n    \"admin\": false,\n    \"image\": \"default\",\n    \"gravatar\": false,\n    \"date_password\": \"2018-03-20T09:31:07.000Z\",\n    \"description\": null,\n    \"website\": null\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := usersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := usersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersGetErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,oid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := usersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	setJSONIndentError(ngsi)

	err := usersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	err := usersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersCreateVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	err := usersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"user\":{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"image\":\"default\",\"gravatar\":false,\"enabled\":true,\"admin\":false,\"username\":\"alice\",\"email\":\"alice@test.com\",\"date_password\":\"2018-03-20T09:31:07.104Z\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	err := usersCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"user\": {\n    \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"image\": \"default\",\n    \"gravatar\": false,\n    \"enabled\": true,\n    \"admin\": false,\n    \"username\": \"alice\",\n    \"email\": \"alice@test.com\",\n    \"date_password\": \"2018-03-20T09:31:07.104Z\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersCreateErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--email=alice@test.com", "--password=passw0rd"})

	setJSONEncodeErr(ngsi, 2)

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify username, email and password", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersCreateErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	setJSONEncodeErr(ngsi, 2)

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	setJSONIndentError(ngsi)

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersCreateErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,username,password,email")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--username=alice", "--email=alice@test.com", "--password=passw0rd"})

	setJSONDecodeErr(ngsi, 1)

	err := usersCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid,username")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username=alice"})

	err := usersUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"user\":{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"image\":\"default\",\"gravatar\":false,\"enabled\":true,\"admin\":false,\"username\":\"alice\",\"email\":\"alice@test.com\",\"date_password\":\"2018-03-20T09:31:07.104Z\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid,username")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username=alice"})

	err := usersUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"user\": {\n    \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"image\": \"default\",\n    \"gravatar\": false,\n    \"enabled\": true,\n    \"admin\": false,\n    \"username\": \"alice\",\n    \"email\": \"alice@test.com\",\n    \"date_password\": \"2018-03-20T09:31:07.104Z\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := usersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := usersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := usersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersUpdateErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	setJSONEncodeErr(ngsi, 2)

	err := usersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid,username")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username=alice"})

	err := usersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid,username")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username=alice"})

	setJSONIndentError(ngsi)

	err := usersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersDelete(c)

	assert.NoError(t, err)
}

func TestUsersDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := usersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := usersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersDeleteErrorUid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := usersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestUsersDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,uid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--uid=2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	err := usersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetUsersParam(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,username,password,email,description,website,gravatar,extra")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--username=alice", "--email=alice@test.com", "--password=1234", "--description=Description", "--website=http://keyrock", "--gravatar=on", "--extra=Extra"})

	actual, err := setUsersParam(c)

	if assert.NoError(t, err) {
		expected := "{\"user\":{\"username\":\"alice\",\"email\":\"alice@test.com\",\"password\":\"1234\",\"gravatar\":true,\"description\":\"Description\",\"website\":\"http://keyrock\",\"extra\":\"Extra\"}}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestSetUsersParamErrorGravatar(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,username,password,email,description,website,gravatar,extra")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--gravatar=aaaa"})

	setJSONEncodeErr(ngsi, 0)

	_, err := setUsersParam(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "specify either true or false to --gravatar", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetUsersParamErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,username,password,email,description,website,gravatar,extra")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--username=alice", "--email=alice@test.com", "--password=1234", "--description=Description", "--website=http://keyrock", "--gravatar=on", "--extra=Extra"})

	setJSONEncodeErr(ngsi, 0)

	_, err := setUsersParam(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetBoolOn(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "gravatar")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--gravatar=on"})

	actual, err := getBool(c, "gravatar")

	if assert.NoError(t, err) {
		assert.Equal(t, true, actual)
	}
}

func TestGetBoolTrue(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "gravatar")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--gravatar=true"})

	actual, err := getBool(c, "gravatar")

	if assert.NoError(t, err) {
		assert.Equal(t, true, actual)
	}
}

func TestGetBoolOff(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "gravatar")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--gravatar=off"})

	actual, err := getBool(c, "gravatar")

	if assert.NoError(t, err) {
		assert.Equal(t, false, actual)
	}
}

func TestGetBoolFalse(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "gravatar")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--gravatar=false"})

	actual, err := getBool(c, "gravatar")

	if assert.NoError(t, err) {
		assert.Equal(t, false, actual)
	}
}

func TestGetBoolError(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "gravatar")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--gravatar=aaaaa"})

	_, err := getBool(c, "gravatar")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "specify either true or false to --gravatar", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
