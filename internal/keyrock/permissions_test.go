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

func TestPermissionsList(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "6\n5\n4\n3\n2\n33fd15c0-e919-47b0-9e05-5f47999f6d91\n1\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsListVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"permissions\":[{\"id\":\"6\",\"name\":\"Getandassignonlypublicownedroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"5\",\"name\":\"Getandassignallpublicapplicationroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"4\",\"name\":\"Manageauthorizations\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"3\",\"name\":\"Manageroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"2\",\"name\":\"Managetheapplication\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"name\":\"permission1\",\"description\":null,\"action\":\"GET\",\"resource\":\"login\",\"xml\":null},{\"id\":\"1\",\"name\":\"Getandassignallinternalapplicationroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"permissions\": [\n    {\n      \"id\": \"6\",\n      \"name\": \"Getandassignonlypublicownedroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"5\",\n      \"name\": \"Getandassignallpublicapplicationroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"4\",\n      \"name\": \"Manageauthorizations\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"3\",\n      \"name\": \"Manageroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"2\",\n      \"name\": \"Managetheapplication\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n      \"name\": \"permission1\",\n      \"description\": null,\n      \"action\": \"GET\",\n      \"resource\": \"login\",\n      \"xml\": null\n    },\n    {\n      \"id\": \"1\",\n      \"name\": \"Getandassignallinternalapplicationroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := permissionsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestPermissionsListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPermissionsListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := permissionsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsListErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := permissionsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsGet(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"permission\":{\"id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"name\":\"permission1\",\"description\":null,\"is_internal\":false,\"action\":\"GET\",\"resource\":\"login\",\"xml\":null,\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsGetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"permission\": {\n    \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n    \"name\": \"permission1\",\n    \"description\": null,\n    \"is_internal\": false,\n    \"action\": \"GET\",\n    \"resource\": \"login\",\n    \"xml\": null,\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := permissionsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestPermissionsGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPermissionsGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := permissionsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsCreate(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "permission1", "--action", "GET", "--resource", "login"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreateData(t *testing.T) {
	data := `{"permission":{"name":"permission1","action":"GET","resource":"login"}}`
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreateVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "permission1", "--action", "GET", "--resource", "login", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"permission\":{\"id\":\"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\"is_internal\":false,\"name\":\"permission1\",\"action\":\"GET\",\"resource\":\"login\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreatePretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "permission1", "--action", "GET", "--resource", "login", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"permission\": {\n    \"id\": \"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\n    \"is_internal\": false,\n    \"name\": \"permission1\",\n    \"action\": \"GET\",\n    \"resource\": \"login\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreateErrorData(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", "@"})

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestPermissionsCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "permission1", "--action", "GET", "--resource", "login"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestPermissionsCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "permission1", "--action", "GET", "--resource", "login"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPermissionsCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "permission1", "--action", "GET", "--resource", "login", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsCreateErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "permission1", "--action", "GET", "--resource", "login"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := permissionsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsUpdate(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--name", "newnamepermission", "--description", "newdescriptionpermission"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"values_updated\":{\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsUpdateData(t *testing.T) {
	data := `{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"values_updated\":{\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsUpdatePretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--name", "newnamepermission", "--description", "newdescriptionpermission", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"values_updated\": {\n    \"name\": \"newnamepermission\",\n    \"description\": \"newdescriptionpermission\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsUpdateErrorData(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--data", "@"})

	err := permissionsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestPermissionsUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--name", "newnamepermission", "--description", "newdescriptionpermission"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/role/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := permissionsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestPermissionsUpdateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--name", "newnamepermission", "--description", "newdescriptionpermission"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPermissionsUpdateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--name", "newnamepermission", "--description", "newdescriptionpermission", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := permissionsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsDelete(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"

	helper.SetClientHTTP(c, reqRes)

	err := permissionsDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestPermissionsDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := permissionsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestPermissionsDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := permissionsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestSetPermissionsData(t *testing.T) {
	data := "{\"permission\":{\"name\":\"abc\",\"description\":\"xyz\",\"action\":\"GET\",\"resource\":\"login\"}}"
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", data})

	actual, err := makePermissionBody(c, c.Ngsi)

	if assert.NoError(t, err) {
		expected := "{\"permission\":{\"name\":\"abc\",\"description\":\"xyz\",\"action\":\"GET\",\"resource\":\"login\"}}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestSetPermissionsParam(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "abc", "--description", "xyz", "--action", "GET", "--resource", "login"})

	actual, err := makePermissionBody(c, c.Ngsi)

	if assert.NoError(t, err) {
		expected := "{\"permission\":{\"name\":\"abc\",\"description\":\"xyz\",\"action\":\"GET\",\"resource\":\"login\"}}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestSetPermissionsParamErrorData(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", "@"})

	_, err := makePermissionBody(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSetPermissionsParamErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name", "abc", "--description", "xyz", "--action", "GET", "--resource", "login"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, err := makePermissionBody(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
