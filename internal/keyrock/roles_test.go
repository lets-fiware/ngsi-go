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

func TestRolesList(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"id":"purchaser","name":"Purchaser"},{"id":"provider","name":"Provider"},{"id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b","name":"role2"},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "purchaser\nprovider\nee2ec16f-694b-447f-b61a-e293b6fe5f7b\n33fd15c0-e919-47b0-9e05-5f47999f6d91\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesListVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"id":"purchaser","name":"Purchaser"},{"id":"provider","name":"Provider"},{"id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b","name":"role2"},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"roles\":[{\"id\":\"purchaser\",\"name\":\"Purchaser\"},{\"id\":\"provider\",\"name\":\"Provider\"},{\"id\":\"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\",\"name\":\"role2\"},{\"id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"name\":\"role1\"}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"id":"purchaser","name":"Purchaser"},{"id":"provider","name":"Provider"},{"id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b","name":"role2"},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"roles\": [\n    {\n      \"id\": \"purchaser\",\n      \"name\": \"Purchaser\"\n    },\n    {\n      \"id\": \"provider\",\n      \"name\": \"Provider\"\n    },\n    {\n      \"id\": \"ee2ec16f-694b-447f-b61a-e293b6fe5f7b\",\n      \"name\": \"role2\"\n    },\n    {\n      \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n      \"name\": \"role1\"\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"id":"purchaser","name":"Purchaser"},{"id":"provider","name":"Provider"},{"id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b","name":"role2"},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1"}]}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := rolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestRolesListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestRolesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"id":"purchaser","name":"Purchaser"},{"id":"provider","name":"Provider"},{"id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b","name":"role2"},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := rolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesListErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"id":"purchaser","name":"Purchaser"},{"id":"provider","name":"Provider"},{"id":"ee2ec16f-694b-447f-b61a-e293b6fe5f7b","name":"role2"},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1"}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := rolesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesGet(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1","is_internal":false,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role\":{\"id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"name\":\"role1\",\"is_internal\":false,\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesGetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1","is_internal":false,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role\": {\n    \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n    \"name\": \"role1\",\n    \"is_internal\": false,\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1","is_internal":false,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := rolesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestRolesGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1","is_internal":false,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestRolesGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"role1","is_internal":false,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := rolesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesCreate(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes1.ResBody = []byte(`{"roles":[]}`)

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes2.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes2.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "33fd15c0-e919-47b0-9e05-5f47999f6d91\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesCreateExist(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"name":"role1","id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "33fd15c0-e919-47b0-9e05-5f47999f6d91\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesCreateData(t *testing.T) {
	data := `{"role":{"name":"role1"}}`
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "33fd15c0-e919-47b0-9e05-5f47999f6d91\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesCreateVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1", "--verbose"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes1.ResBody = []byte(`{"roles":[]}`)

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes2.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes2.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"role\":{\"id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"is_internal\":false,\"name\":\"role1\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesCreatePretty(t *testing.T) {
	data := `{"role":{"name":"role1"}}`
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role\": {\n    \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n    \"is_internal\": false,\n    \"name\": \"role1\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesCreateErrorData(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", "@"})

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestRolesCreateErrorGetRoleByName(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestRolesCreateErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"name":"role1","id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)

	helper.SetJSONEncodeErr(c.Ngsi, 0)
	helper.SetClientHTTP(c, reqRes)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesCreateErrorPrintRole(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"name":"role1","id":"33fd15c0-e919-47b0-9e05-5f47999f6d91"}]}`)

	helper.SetJSONIndentError(c.Ngsi)
	helper.SetClientHTTP(c, reqRes)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesCreateErrorName(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[]}`)

	helper.SetClientHTTP(c, reqRes)
	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1", "--verbose"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes1.ResBody = []byte(`{"roles":[]}`)

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/role"
	reqRes2.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes2.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRolesCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1", "--verbose"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes1.ResBody = []byte(`{"roles":[]}`)

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes2.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestRolesCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1", "--verbose", "--pretty"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes1.ResBody = []byte(`{"roles":[]}`)

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes2.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes2.ResBody = []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	helper.SetJSONIndentError(c.Ngsi)

	err := rolesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPrintRole(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1", "--verbose", "--pretty"})

	body := []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	err := printRole(c, c.Ngsi, body)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"role\": {\n    \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n    \"is_internal\": false,\n    \"name\": \"role1\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPrintRoleErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1", "--verbose", "--pretty"})

	body := []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetJSONIndentError(c.Ngsi)

	err := printRole(c, c.Ngsi, body)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPrintRoleErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	body := []byte(`{"role":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","is_internal":false,"name":"role1","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := printRole(c, c.Ngsi, body)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesUpdate(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission","xml":"xmlrule"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"values_updated\":{\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\",\"xml\":\"xmlrule\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesUpdateData(t *testing.T) {
	data := `{"role":{"name":"role1"}}`
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission","xml":"xmlrule"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"values_updated\":{\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\",\"xml\":\"xmlrule\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesUpdatePretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", `{"role":{"name":"role1"}}`, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission","xml":"xmlrule"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"values_updated\": {\n    \"name\": \"newnamepermission\",\n    \"description\": \"newdescriptionpermission\",\n    \"xml\": \"xmlrule\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRolesUpdateErrorData(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", "@"})

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestRolesUpdateErrorName(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--name", "role1"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", `{"role":{"name":"role1"}}`})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/role/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission","xml":"xmlrule"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestRolesUpdateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", `{"role":{"name":"role1"}}`})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestRolesUpdateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", `{"role":{"name":"role1"}}`, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ReqData = []byte(`{"role":{"name":"role1"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission","xml":"xmlrule"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := rolesUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRolesDelete(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"

	helper.SetClientHTTP(c, reqRes)

	err := rolesDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRolesDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := rolesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestRolesDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := rolesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGetRoleByName(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"roles":[{"id":"1c5d64b2-f023-455b-8b45-35452745961a","name":"role1"}]}`)

	helper.SetClientHTTP(c, reqRes)

	r, err := getRoleByName(c.Ngsi, c.Client, "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "role1")

	if assert.NoError(t, err) {
		actual := r.ID
		expected := "1c5d64b2-f023-455b-8b45-35452745961a"
		assert.Equal(t, expected, actual)
	}
}

func TestGetRoleByNameNotFound(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{}`)

	helper.SetClientHTTP(c, reqRes)

	actual, err := getRoleByName(c.Ngsi, c.Client, "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "role1")

	if assert.NoError(t, err) {
		assert.Equal(t, (*keyrockRoleItmes)(nil), actual)
	}
}

func TestGetRoleByNameErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	_, err := getRoleByName(c.Ngsi, c.Client, "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "role1")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestGetRoleByNameErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	_, err := getRoleByName(c.Ngsi, c.Client, "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "role1")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestGetRoleByNameErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "role1"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/roles"
	reqRes.ResBody = []byte(`{}`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)
	helper.SetClientHTTP(c, reqRes)

	_, err := getRoleByName(c.Ngsi, c.Client, "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "role1")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
