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

func TestApplicationsList(t *testing.T) {
	c := setupTest([]string{"applications", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\nfd7fe349-f7da-4c27-b404-74da17641025\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsListVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "list", "--host", "keyrock", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"applications\":[{\"id\":\"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\"name\":\"Test_application2\",\"description\":\"Description\",\"image\":\"default\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost/login\",\"grant_type\":\"client_credentials,password,implicit,authorization_code,refresh_token\",\"response_type\":\"code,token\",\"token_types\":\"bearer,jwt,permanent\",\"client_type\":null},{\"id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"name\":\"Test_application1\",\"description\":\"description\",\"image\":\"default\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost/login\",\"grant_type\":\"password,authorization_code,implicit\",\"response_type\":\"code,token\",\"token_types\":\"bearer\",\"client_type\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsListPretty(t *testing.T) {
	c := setupTest([]string{"applications", "list", "--host", "keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"applications\": [\n    {\n      \"id\": \"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\n      \"name\": \"Test_application2\",\n      \"description\": \"Description\",\n      \"image\": \"default\",\n      \"url\": \"http://localhost\",\n      \"redirect_uri\": \"http://localhost/login\",\n      \"grant_type\": \"client_credentials,password,implicit,authorization_code,refresh_token\",\n      \"response_type\": \"code,token\",\n      \"token_types\": \"bearer,jwt,permanent\",\n      \"client_type\": null\n    },\n    {\n      \"id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n      \"name\": \"Test_application1\",\n      \"description\": \"description\",\n      \"image\": \"default\",\n      \"url\": \"http://localhost\",\n      \"redirect_uri\": \"http://localhost/login\",\n      \"grant_type\": \"password,authorization_code,implicit\",\n      \"response_type\": \"code,token\",\n      \"token_types\": \"bearer\",\n      \"client_type\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := applicationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestApplicationsListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestApplicationsListErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "list", "--host", "keyrock", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := applicationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsListErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "list", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := applicationsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsGet(t *testing.T) {
	c := setupTest([]string{"applications", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"application\":{\"id\":\"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\"name\":\"Test_application2\",\"description\":\"Description\",\"secret\":\"61f5def7-bcf9-45b1-9c69-d0887e403737\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost/login\",\"image\":\"default\",\"grant_type\":\"client_credentials,password,implicit,authorization_code,refresh_token\",\"response_type\":\"code,token\",\"token_types\":\"bearer,jwt,permanent\",\"jwt_secret\":\"3f1164da20d50c62\",\"client_type\":null,\"scope\":null,\"extra\":null}}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsGetPretty(t *testing.T) {
	c := setupTest([]string{"applications", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"application\": {\n    \"id\": \"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\n    \"name\": \"Test_application2\",\n    \"description\": \"Description\",\n    \"secret\": \"61f5def7-bcf9-45b1-9c69-d0887e403737\",\n    \"url\": \"http://localhost\",\n    \"redirect_uri\": \"http://localhost/login\",\n    \"image\": \"default\",\n    \"grant_type\": \"client_credentials,password,implicit,authorization_code,refresh_token\",\n    \"response_type\": \"code,token\",\n    \"token_types\": \"bearer,jwt,permanent\",\n    \"jwt_secret\": \"3f1164da20d50c62\",\n    \"client_type\": null,\n    \"scope\": null,\n    \"extra\": null\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := applicationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestApplicationsGetErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestApplicationsGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := applicationsGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsCreate(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "fd7fe349-f7da-4c27-b404-74da17641025\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsCreateVerbose(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{}", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"application\":{\"id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"secret\":\"9dc463cf-8318-4f65-bc02-778424fdfd77\",\"image\":\"default\",\"name\":\"Test_application1\",\"description\":\"description\",\"redirect_uri\":\"http://localhost/login\",\"url\":\"http://localhost\",\"grant_type\":\"password,authorization_code,implicit\",\"token_types\":\"jwt,permanent\",\"jwt_secret\":\"3f1164da20d50c62\",\"response_type\":\"code,token\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsCreatePretty(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{}", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"application\": {\n    \"id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n    \"secret\": \"9dc463cf-8318-4f65-bc02-778424fdfd77\",\n    \"image\": \"default\",\n    \"name\": \"Test_application1\",\n    \"description\": \"description\",\n    \"redirect_uri\": \"http://localhost/login\",\n    \"url\": \"http://localhost\",\n    \"grant_type\": \"password,authorization_code,implicit\",\n    \"token_types\": \"jwt,permanent\",\n    \"jwt_secret\": \"3f1164da20d50c62\",\n    \"response_type\": \"code,token\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsCreateErrorMakeBody(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "@"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestApplicationsCreateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestApplicationsCreateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestApplicationsCreateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{}", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsCreateErrorID(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := applicationsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsUpdate(t *testing.T) {
	c := setupTest([]string{"applications", "update", "--host", "keyrock", "--aid", "fd7fe349-f7da-4c27-b404-74da17641025", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"application\":{\"id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"secret\":\"9dc463cf-8318-4f65-bc02-778424fdfd77\",\"image\":\"default\",\"name\":\"Test_application1\",\"description\":\"description\",\"redirect_uri\":\"http://localhost/login\",\"url\":\"http://localhost\",\"grant_type\":\"password,authorization_code,implicit\",\"token_types\":\"jwt,permanent\",\"jwt_secret\":\"3f1164da20d50c62\",\"response_type\":\"code,token\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsUpdatePretty(t *testing.T) {
	c := setupTest([]string{"applications", "update", "--host", "keyrock", "--aid", "fd7fe349-f7da-4c27-b404-74da17641025", "--data", "{}", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"application\": {\n    \"id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n    \"secret\": \"9dc463cf-8318-4f65-bc02-778424fdfd77\",\n    \"image\": \"default\",\n    \"name\": \"Test_application1\",\n    \"description\": \"description\",\n    \"redirect_uri\": \"http://localhost/login\",\n    \"url\": \"http://localhost\",\n    \"grant_type\": \"password,authorization_code,implicit\",\n    \"token_types\": \"jwt,permanent\",\n    \"jwt_secret\": \"3f1164da20d50c62\",\n    \"response_type\": \"code,token\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsErrorMakeBody(t *testing.T) {
	c := setupTest([]string{"applications", "update", "--host", "keyrock", "--aid", "fd7fe349-f7da-4c27-b404-74da17641025", "--data", "@"})

	err := applicationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestApplicationsUpdateErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "update", "--host", "keyrock", "--aid", "fd7fe349-f7da-4c27-b404-74da17641025", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := applicationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestApplicationsUpdateErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "update", "--host", "keyrock", "--aid", "fd7fe349-f7da-4c27-b404-74da17641025", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestApplicationsUpdateErrorPretty(t *testing.T) {
	c := setupTest([]string{"applications", "update", "--host", "keyrock", "--aid", "fd7fe349-f7da-4c27-b404-74da17641025", "--data", "{}", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := applicationsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsDelete(t *testing.T) {
	c := setupTest([]string{"applications", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"

	helper.SetClientHTTP(c, reqRes)

	err := applicationsDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestApplicationsDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"applications", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := applicationsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestApplicationsDeleteErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"applications", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := applicationsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestMakeAppBodyData(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "{\"application\":{}}"})

	actual, err := makeAppBody(c, c.Ngsi, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{}}", string(actual))
	}
}

func TestMakeAppBodyParam1(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--name", "app1", "--description", "application1", "--redirectUri", "http://ruri", "--redirectSignOutUri", "http://suri", "--url", "http://url"})

	actual, err := makeAppBody(c, c.Ngsi, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{\"name\":\"app1\",\"description\":\"application1\",\"url\":\"http://url\",\"redirect_uri\":\"http://ruri\",\"redirect_sign_out_uri\":\"http://suri\",\"grant_type\":[\"client_credentials\",\"password\",\"implicit\",\"authorization_code\",\"refresh_token\"]}}", string(actual))
	}
}

func TestMakeAppBodyParam2(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--name", "app", "--grantType", "abc,def", "--tokenTypes", "123,456", "--responseType", "123,xyz", "--clientType", "987,654"})

	actual, err := makeAppBody(c, c.Ngsi, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{\"name\":\"app\",\"description\":\"app\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost\",\"redirect_sign_out_uri\":\"\",\"grant_type\":[\"abc\",\"def\"],\"response_type\":[\"123\",\"xyz\"],\"token_types\":[\"123\",\"456\"],\"client_type\":[\"987\",\"654\"]}}", string(actual))
	}
}

func TestMakeAppBodyScope(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--name", "app", "--scope", "jwt"})

	actual, err := makeAppBody(c, c.Ngsi, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{\"name\":\"app\",\"description\":\"app\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost\",\"redirect_sign_out_uri\":\"\",\"grant_type\":[\"client_credentials\",\"password\",\"implicit\",\"authorization_code\",\"refresh_token\"],\"scope\":[\"jwt\"]}}", string(actual))
	}
}
func TestMakeAppBodyOpenID(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--name", "app", "--openid"})

	actual, err := makeAppBody(c, c.Ngsi, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{\"name\":\"app\",\"description\":\"app\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost\",\"redirect_sign_out_uri\":\"\",\"grant_type\":[\"client_credentials\",\"password\",\"implicit\",\"authorization_code\",\"refresh_token\"],\"token_types\":[\"jwt\"],\"scope\":\"openid\"}}", string(actual))
	}
}

func TestMakeAppBodyErrorData(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--data", "@"})

	_, err := makeAppBody(c, c.Ngsi, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestMakeAppBodyErrorParam(t *testing.T) {
	c := setupTest([]string{"applications", "create", "--host", "keyrock", "--name", "app1", "--description", "application1", "--redirectUri", "http://ruril", "--url", "http://url"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, err := makeAppBody(c, c.Ngsi, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
