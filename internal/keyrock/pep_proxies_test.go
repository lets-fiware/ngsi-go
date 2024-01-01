/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

func TestPepProxiesList(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"pep_proxy\":{\"id\":\"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a\",\"oauth_client_id\":\"fd7fe349-f7da-4c27-b404-74da17641025\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPepProxiesListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"pep_proxy\": {\n    \"id\": \"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a\",\n    \"oauth_client_id\": \"fd7fe349-f7da-4c27-b404-74da17641025\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPepProxiesListNotFound(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "pep proxy not found\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPepProxiesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestPepProxiesListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPepProxiesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := pepProxiesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPepProxiesCreate(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","password":"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"pep_proxy\":{\"id\":\"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a\",\"password\":\"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPepProxiesCreatePretty(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","password":"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"pep_proxy\": {\n    \"id\": \"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a\",\n    \"password\": \"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPepProxiesCreateErrorRun(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := pepProxiesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "run create with --run option", ngsiErr.Message)
	}
}

func TestPepProxiesCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/role"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","password":"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestPepProxiesCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPepProxiesCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","password":"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := pepProxiesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPepProxiesReset(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","password":"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesReset(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"pep_proxy\":{\"id\":\"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a\",\"password\":\"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestPepProxiesResetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","password":"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesReset(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"pep_proxy\": {\n    \"id\": \"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a\",\n    \"password\": \"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPepProxiesResetErrorRun(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := pepProxiesReset(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "run reset with --run option", ngsiErr.Message)
	}
}

func TestPepProxiesResetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"permission":{"id":"33fd15c0-e919-47b0-9e05-5f47999f6d91","name":"permission1","description":null,"is_internal":false,"action":"GET","resource":"login","xml":null,"oauth_client_id":"fd7fe349-f7da-4c27-b404-74da17641025"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesReset(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestPepProxiesResetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesReset(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPepProxiesResetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"pep_proxy":{"id":"pep_proxy_2d19d297-e555-4e3a-a18d-22deda37036a","password":"pep_proxy_950468f7-4198-42e1-8f41-4b1f8c5bc1f2"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := pepProxiesReset(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
func TestPepProxiesDelete(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestPepProxiesDeleteErrorRun(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "run delete with --run option", ngsiErr.Message)
	}
}

func TestPepProxiesDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestPepProxiesDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "pep", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/pep_proxies"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := pepProxiesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
