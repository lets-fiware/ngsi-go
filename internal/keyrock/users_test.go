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

func TestUsersList(t *testing.T) {
	c := setupTest([]string{"users", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\nadmin\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersListVerbose(t *testing.T) {
	c := setupTest([]string{"users", "list", "--host", "keyrock", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"users\":[{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"username\":\"alice\",\"email\":\"alice@test.com\",\"enabled\":true,\"gravatar\":false,\"date_password\":\"2018-03-20T09:31:07.000Z\",\"description\":null,\"website\":null},{\"id\":\"admin\",\"username\":\"admin\",\"email\":\"admin@test.com\",\"enabled\":true,\"gravatar\":false,\"date_password\":\"2018-03-20T08:40:14.000Z\",\"description\":null,\"website\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersListPretty(t *testing.T) {
	c := setupTest([]string{"users", "list", "--host", "keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"users\": [\n    {\n      \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n      \"username\": \"alice\",\n      \"email\": \"alice@test.com\",\n      \"enabled\": true,\n      \"gravatar\": false,\n      \"date_password\": \"2018-03-20T09:31:07.000Z\",\n      \"description\": null,\n      \"website\": null\n    },\n    {\n      \"id\": \"admin\",\n      \"username\": \"admin\",\n      \"email\": \"admin@test.com\",\n      \"enabled\": true,\n      \"gravatar\": false,\n      \"date_password\": \"2018-03-20T08:40:14.000Z\",\n      \"description\": null,\n      \"website\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"users", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := usersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestUsersListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"users", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestUsersListErrorPretty(t *testing.T) {
	c := setupTest([]string{"users", "list", "--host", "keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := usersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersListErrorID(t *testing.T) {
	c := setupTest([]string{"users", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users"
	reqRes.ResBody = []byte(`{"users":[{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null},{"id":"admin","username":"admin","email":"admin@test.com","enabled":true,"gravatar":false,"date_password":"2018-03-20T08:40:14.000Z","description":null,"website":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := usersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersGet(t *testing.T) {
	c := setupTest([]string{"users", "get", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"user\":{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"username\":\"alice\",\"email\":\"alice@test.com\",\"enabled\":true,\"admin\":false,\"image\":\"default\",\"gravatar\":false,\"date_password\":\"2018-03-20T09:31:07.000Z\",\"description\":null,\"website\":null}}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersGetPretty(t *testing.T) {
	c := setupTest([]string{"users", "get", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"user\": {\n    \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"username\": \"alice\",\n    \"email\": \"alice@test.com\",\n    \"enabled\": true,\n    \"admin\": false,\n    \"image\": \"default\",\n    \"gravatar\": false,\n    \"date_password\": \"2018-03-20T09:31:07.000Z\",\n    \"description\": null,\n    \"website\": null\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"users", "get", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := usersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestUsersGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"users", "get", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestUsersGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"users", "get", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","username":"alice","email":"alice@test.com","enabled":true,"admin":false,"image":"default","gravatar":false,"date_password":"2018-03-20T09:31:07.000Z","description":null,"website":null}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := usersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersCreate(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "2d6f5391-6130-48d8-a9d0-01f20699a7eb\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersCreateVerbose(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"user\":{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"image\":\"default\",\"gravatar\":false,\"enabled\":true,\"admin\":false,\"username\":\"alice\",\"email\":\"alice@test.com\",\"date_password\":\"2018-03-20T09:31:07.104Z\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersCreatePretty(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"user\": {\n    \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"image\": \"default\",\n    \"gravatar\": false,\n    \"enabled\": true,\n    \"admin\": false,\n    \"username\": \"alice\",\n    \"email\": \"alice@test.com\",\n    \"date_password\": \"2018-03-20T09:31:07.104Z\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersCreateErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd", "--pretty"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestUsersCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestUsersCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersCreateErrorID(t *testing.T) {
	c := setupTest([]string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/users"
	reqRes.ReqData = []byte(`{"user":{"username":"alice","email":"alice@test.com","password":"passw0rd"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := usersCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersUpdate(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"user\":{\"id\":\"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\"image\":\"default\",\"gravatar\":false,\"enabled\":true,\"admin\":false,\"username\":\"alice\",\"email\":\"alice@test.com\",\"date_password\":\"2018-03-20T09:31:07.104Z\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersUpdatePretty(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"user\": {\n    \"id\": \"2d6f5391-6130-48d8-a9d0-01f20699a7eb\",\n    \"image\": \"default\",\n    \"gravatar\": false,\n    \"enabled\": true,\n    \"admin\": false,\n    \"username\": \"alice\",\n    \"email\": \"alice@test.com\",\n    \"date_password\": \"2018-03-20T09:31:07.104Z\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestUsersUpdateErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := usersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := usersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestUsersUpdateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestUsersUpdateErrorPretty(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ReqData = []byte(`{"user":{"username":"alice"}}`)
	reqRes.ResBody = []byte(`{"user":{"id":"2d6f5391-6130-48d8-a9d0-01f20699a7eb","image":"default","gravatar":false,"enabled":true,"admin":false,"username":"alice","email":"alice@test.com","date_password":"2018-03-20T09:31:07.104Z"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := usersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestUsersDelete(t *testing.T) {
	c := setupTest([]string{"users", "delete", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"

	helper.SetClientHTTP(c, reqRes)

	err := usersDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestUsersDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"users", "delete", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := usersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestUsersDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"users", "delete", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/users/2d6f5391-6130-48d8-a9d0-01f20699a7eb"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := usersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestSetUsersParam(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice", "--email", "alice@test.com", "--password", "1234", "--description", "Description", "--website", "http://keyrock", "--gravatar", "on", "--extra", "Extra"})

	actual, err := setUsersParam(c)

	if assert.NoError(t, err) {
		expected := "{\"user\":{\"username\":\"alice\",\"email\":\"alice@test.com\",\"password\":\"1234\",\"gravatar\":true,\"description\":\"Description\",\"website\":\"http://keyrock\",\"extra\":\"Extra\"}}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestSetUsersParamErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--username", "alice", "--email", "alice@test.com", "--password", "1234", "--description", "Description", "--website", "http://keyrock", "--gravatar", "on", "--extra", "Extra"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, err := setUsersParam(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
