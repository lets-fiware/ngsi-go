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

func TestTrustedAppList(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d",]}`)

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"trusted_applications\":[\"8692ec57-8514-4ef6-a347-3d1ac6409f79\",\"78b4763f-139a-4820-a42b-3e265fb9d56e\",\"462ee067-f10a-4c9c-aefe-079038830043\",\"6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d\",]}"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d"]}`)

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"trusted_applications\": [\n    \"8692ec57-8514-4ef6-a347-3d1ac6409f79\",\n    \"78b4763f-139a-4820-a42b-3e265fb9d56e\",\n    \"462ee067-f10a-4c9c-aefe-079038830043\",\n    \"6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppListNotFound(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "Trusted applications nof found\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d"]}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestTrustedAppListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTrustedAppListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d"]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := trustedAppList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTrustedAppAdd(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "add", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"oauth_client_id\": \"3e34cbf6-3579-4e54-878a-0a406962ac36\",\"trusted_oauth_client_id\": \"0118ccb7-756e-42f9-8a19-5b4e83ca8c46\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppAddPretty(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "add", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"oauth_client_id\": \"3e34cbf6-3579-4e54-878a-0a406962ac36\",\n  \"trusted_oauth_client_id\": \"0118ccb7-756e-42f9-8a19-5b4e83ca8c46\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppAddErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "add", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestTrustedAppAddErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "add", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestTrustedAppAddErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "add", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := trustedAppAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTrustedAppDelete(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTrustedAppDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestTrustedAppDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "trusted", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := trustedAppDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
