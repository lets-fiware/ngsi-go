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

func TestPermissionsList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := permissionsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "6\n5\n4\n3\n2\n33fd15c0-e919-47b0-9e05-5f47999f6d91\n1\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := permissionsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"permissions\":[{\"id\":\"6\",\"name\":\"Getandassignonlypublicownedroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"5\",\"name\":\"Getandassignallpublicapplicationroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"4\",\"name\":\"Manageauthorizations\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"3\",\"name\":\"Manageroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"2\",\"name\":\"Managetheapplication\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null},{\"id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"name\":\"permission1\",\"description\":null,\"action\":\"GET\",\"resource\":\"login\",\"xml\":null},{\"id\":\"1\",\"name\":\"Getandassignallinternalapplicationroles\",\"description\":null,\"action\":null,\"resource\":null,\"xml\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := permissionsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"permissions\": [\n    {\n      \"id\": \"6\",\n      \"name\": \"Getandassignonlypublicownedroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"5\",\n      \"name\": \"Getandassignallpublicapplicationroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"4\",\n      \"name\": \"Manageauthorizations\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"3\",\n      \"name\": \"Manageroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"2\",\n      \"name\": \"Managetheapplication\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    },\n    {\n      \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n      \"name\": \"permission1\",\n      \"description\": null,\n      \"action\": \"GET\",\n      \"resource\": \"login\",\n      \"xml\": null\n    },\n    {\n      \"id\": \"1\",\n      \"name\": \"Getandassignallinternalapplicationroles\",\n      \"description\": null,\n      \"action\": null,\n      \"resource\": null,\n      \"xml\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := permissionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/permissions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := permissionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsListErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := permissionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := permissionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := permissionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	setJSONIndentError(ngsi)

	err := permissionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsListErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ResBody = []byte(`{"permissions":[{"id":"6","name":"Getandassignonlypublicownedroles","description":null,"action":null,"resource":null,"xml":null},{"id":"5","name":"Getandassignallpublicapplicationroles","description":null,"action":null,"resource":null,"xml":null},{"id":"4","name":"Manageauthorizations","description":null,"action":null,"resource":null,"xml":null},{"id":"3","name":"Manageroles","description":null,"action":null,"resource":null,"xml":null},{"id":"2","name":"Managetheapplication","description":null,"action":null,"resource":null,"xml":null},{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"action":"GET","resource":"login","xml":null},{"id":"1","name":"Getandassignallinternalapplicationroles","description":null,"action":null,"resource":null,"xml":null}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	setJSONDecodeErr(ngsi, 1)
	err := permissionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := permissionsGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"permission\":{\"id\":\"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\"name\":\"permission1\",\"description\":null,\"is_internal\":false,\"action\":\"GET\",\"resource\":\"login\",\"xml\":null,\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := permissionsGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"permission\": {\n    \"id\": \"33fd15c0-e919-47b0-9e05-5f47999f6d91\",\n    \"name\": \"permission1\",\n    \"description\": null,\n    \"is_internal\": false,\n    \"action\": \"GET\",\n    \"resource\": \"login\",\n    \"xml\": null,\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := permissionsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/permissions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := permissionsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsGetErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := permissionsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsGetErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := permissionsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify permission id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := permissionsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	err := permissionsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/33fd15c0-e919-47b0-9e05-5f47999f6d91"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=33fd15c0-e919-47b0-9e05-5f47999f6d91"})

	setJSONIndentError(ngsi)

	err := permissionsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPermissionsCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name=permission1", "--action=GET", "--resource=login"})

	err := permissionsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreateData(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", `--data={"permission":{"name":"permission1","action":"GET","resource":"login"}}`})

	err := permissionsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreateVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name=permission1", "--action=GET", "--resource=login"})

	err := permissionsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"permission\":{\"id\":\"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\"is_internal\":false,\"name\":\"permission1\",\"action\":\"GET\",\"resource\":\"login\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name=permission1", "--action=GET", "--resource=login"})

	err := permissionsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"permission\": {\n    \"id\": \"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3\",\n    \"is_internal\": false,\n    \"name\": \"permission1\",\n    \"action\": \"GET\",\n    \"resource\": \"login\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/permissions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsCreateErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--name=permission1", "--action=GET", "--resource=login"})

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", `--data=`})

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/role"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name=permission1", "--action=GET", "--resource=login"})

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name=permission1", "--action=GET", "--resource=login"})

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name=permission1", "--action=GET", "--resource=login"})

	setJSONIndentError(ngsi)

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsCreateErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions"
	reqRes.ReqData = []byte(`{"permission":{"name":"permission1","action":"GET","resource":"login"}}`)
	reqRes.ResBody = []byte(`{"permission":{"id":"15ca810b-27d1-44a1-8491-a3fb4b6bc6f3","is_internal":false,"name":"permission1","action":"GET","resource":"login","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,name,action,resource")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--name=permission1", "--action=GET", "--resource=login"})

	setJSONDecodeErr(ngsi, 1)

	err := permissionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,name,description")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", "--name=newnamepermission", "--description=newdescriptionpermission"})

	err := permissionsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"values_updated\":{\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsUpdateData(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", `--data={"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`})

	err := permissionsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"values_updated\":{\"name\":\"newnamepermission\",\"description\":\"newdescriptionpermission\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", `--data={"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`})

	err := permissionsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"values_updated\": {\n    \"name\": \"newnamepermission\",\n    \"description\": \"newdescriptionpermission\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPermissionsUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/permissions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdateErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", `--data={}`})

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdateErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", `--data={}`})

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify permission id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", `--data=`})

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/role/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"values_updated":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", `--data={"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`})

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", `--data={"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`})

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ReqData = []byte(`{"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3", `--data={"permission":{"name":"newnamepermission","description":"newdescriptionpermission"}}`})

	setJSONIndentError(ngsi)

	err := permissionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"})

	err := permissionsDelete(c)

	assert.NoError(t, err)
}

func TestPermissionsDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := permissionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/permissions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := permissionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsDeleteErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := permissionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsDeleteErrorRid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := permissionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify permission id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"})

	err := permissionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPermissionsDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/permissions/15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,pid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid=15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"})

	err := permissionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetPermissionsData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={\"permission\":{\"name\":\"abc\",\"description\":\"xyz\",\"action\":\"GET\",\"resource\":\"login\"}}"})

	actual, err := makePermissionBody(c, ngsi)

	if assert.NoError(t, err) {
		expected := "{\"permission\":{\"name\":\"abc\",\"description\":\"xyz\",\"action\":\"GET\",\"resource\":\"login\"}}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestSetPermissionsParam(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "name,description,action,resource")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=abc", "--description=xyz", "--action=GET", "--resource=login"})

	actual, err := makePermissionBody(c, ngsi)

	if assert.NoError(t, err) {
		expected := "{\"permission\":{\"name\":\"abc\",\"description\":\"xyz\",\"action\":\"GET\",\"resource\":\"login\"}}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestSetPermissionsParamErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data="})

	setJSONEncodeErr(ngsi, 0)

	_, err := makePermissionBody(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetPermissionsParamErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "name,description,action,resource")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=abc", "--description=xyz", "--action=GET", "--resource=login"})

	setJSONEncodeErr(ngsi, 0)

	_, err := makePermissionBody(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
