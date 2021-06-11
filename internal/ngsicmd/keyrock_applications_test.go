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

func TestApplicationsList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := applicationsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\nfd7fe349-f7da-4c27-b404-74da17641025\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsListVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--verbose"})

	err := applicationsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"applications\":[{\"id\":\"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\"name\":\"Test_application2\",\"description\":\"Description\",\"image\":\"default\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost/login\",\"grant_type\":\"client_credentials,password,implicit,authorization_code,refresh_token\",\"response_type\":\"code,token\",\"token_types\":\"bearer,jwt,permanent\",\"client_type\":null},{\"id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"name\":\"Test_application1\",\"description\":\"description\",\"image\":\"default\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost/login\",\"grant_type\":\"password,authorization_code,implicit\",\"response_type\":\"code,token\",\"token_types\":\"bearer\",\"client_type\":null}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty"})

	err := applicationsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"applications\": [\n    {\n      \"id\": \"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\n      \"name\": \"Test_application2\",\n      \"description\": \"Description\",\n      \"image\": \"default\",\n      \"url\": \"http://localhost\",\n      \"redirect_uri\": \"http://localhost/login\",\n      \"grant_type\": \"client_credentials,password,implicit,authorization_code,refresh_token\",\n      \"response_type\": \"code,token\",\n      \"token_types\": \"bearer,jwt,permanent\",\n      \"client_type\": null\n    },\n    {\n      \"id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n      \"name\": \"Test_application1\",\n      \"description\": \"description\",\n      \"image\": \"default\",\n      \"url\": \"http://localhost\",\n      \"redirect_uri\": \"http://localhost/login\",\n      \"grant_type\": \"password,authorization_code,implicit\",\n      \"response_type\": \"code,token\",\n      \"token_types\": \"bearer\",\n      \"client_type\": null\n    }\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := applicationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := applicationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := applicationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := applicationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty"})

	setJSONIndentError(ngsi)

	err := applicationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsListErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	reqRes.ResBody = []byte(`{"applications":[{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","client_type":null},{"id":"fd7fe349-f7da-4c27-b404-74da17641025","name":"Test_application1","description":"description","image":"default","url":"http://localhost","redirect_uri":"http://localhost/login","grant_type":"password,authorization_code,implicit","response_type":"code,token","token_types":"bearer","client_type":null}]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	setJSONDecodeErr(ngsi, 1)
	err := applicationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := applicationsGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"application\":{\"id\":\"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\"name\":\"Test_application2\",\"description\":\"Description\",\"secret\":\"61f5def7-bcf9-45b1-9c69-d0887e403737\",\"url\":\"http://localhost\",\"redirect_uri\":\"http://localhost/login\",\"image\":\"default\",\"grant_type\":\"client_credentials,password,implicit,authorization_code,refresh_token\",\"response_type\":\"code,token\",\"token_types\":\"bearer,jwt,permanent\",\"jwt_secret\":\"3f1164da20d50c62\",\"client_type\":null,\"scope\":null,\"extra\":null}}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := applicationsGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"application\": {\n    \"id\": \"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c\",\n    \"name\": \"Test_application2\",\n    \"description\": \"Description\",\n    \"secret\": \"61f5def7-bcf9-45b1-9c69-d0887e403737\",\n    \"url\": \"http://localhost\",\n    \"redirect_uri\": \"http://localhost/login\",\n    \"image\": \"default\",\n    \"grant_type\": \"client_credentials,password,implicit,authorization_code,refresh_token\",\n    \"response_type\": \"code,token\",\n    \"token_types\": \"bearer,jwt,permanent\",\n    \"jwt_secret\": \"3f1164da20d50c62\",\n    \"client_type\": null,\n    \"scope\": null,\n    \"extra\": null\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := applicationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := applicationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsGetErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := applicationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := applicationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := applicationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	setJSONIndentError(ngsi)

	err := applicationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestApplicationsCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}"})

	err := applicationsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "fd7fe349-f7da-4c27-b404-74da17641025\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsCreateVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}", "--verbose"})

	err := applicationsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"application\":{\"id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"secret\":\"9dc463cf-8318-4f65-bc02-778424fdfd77\",\"image\":\"default\",\"name\":\"Test_application1\",\"description\":\"description\",\"redirect_uri\":\"http://localhost/login\",\"url\":\"http://localhost\",\"grant_type\":\"password,authorization_code,implicit\",\"token_types\":\"jwt,permanent\",\"jwt_secret\":\"3f1164da20d50c62\",\"response_type\":\"code,token\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}", "--pretty"})

	err := applicationsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"application\": {\n    \"id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n    \"secret\": \"9dc463cf-8318-4f65-bc02-778424fdfd77\",\n    \"image\": \"default\",\n    \"name\": \"Test_application1\",\n    \"description\": \"description\",\n    \"redirect_uri\": \"http://localhost/login\",\n    \"url\": \"http://localhost\",\n    \"grant_type\": \"password,authorization_code,implicit\",\n    \"token_types\": \"jwt,permanent\",\n    \"jwt_secret\": \"3f1164da20d50c62\",\n    \"response_type\": \"code,token\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := applicationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := applicationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsCreateErrorMakeBody(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data=", "--pretty"})

	err := applicationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}"})

	err := applicationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}"})

	err := applicationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}", "--pretty"})

	setJSONIndentError(ngsi)

	err := applicationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsCreateErrorID(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}"})

	setJSONDecodeErr(ngsi, 1)

	err := applicationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=fd7fe349-f7da-4c27-b404-74da17641025", "--data={}"})

	err := applicationsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"application\":{\"id\":\"fd7fe349-f7da-4c27-b404-74da17641025\",\"secret\":\"9dc463cf-8318-4f65-bc02-778424fdfd77\",\"image\":\"default\",\"name\":\"Test_application1\",\"description\":\"description\",\"redirect_uri\":\"http://localhost/login\",\"url\":\"http://localhost\",\"grant_type\":\"password,authorization_code,implicit\",\"token_types\":\"jwt,permanent\",\"jwt_secret\":\"3f1164da20d50c62\",\"response_type\":\"code,token\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=fd7fe349-f7da-4c27-b404-74da17641025", "--data={}"})

	err := applicationsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"application\": {\n    \"id\": \"fd7fe349-f7da-4c27-b404-74da17641025\",\n    \"secret\": \"9dc463cf-8318-4f65-bc02-778424fdfd77\",\n    \"image\": \"default\",\n    \"name\": \"Test_application1\",\n    \"description\": \"description\",\n    \"redirect_uri\": \"http://localhost/login\",\n    \"url\": \"http://localhost\",\n    \"grant_type\": \"password,authorization_code,implicit\",\n    \"token_types\": \"jwt,permanent\",\n    \"jwt_secret\": \"3f1164da20d50c62\",\n    \"response_type\": \"code,token\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestApplicationsUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := applicationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := applicationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data=", "--pretty"})

	err := applicationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsErrorMakeBody(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=fd7fe349-f7da-4c27-b404-74da17641025", "--data=", "--pretty"})

	err := applicationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=fd7fe349-f7da-4c27-b404-74da17641025"})

	err := applicationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}", "--aid=fd7fe349-f7da-4c27-b404-74da17641025"})

	err := applicationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/fd7fe349-f7da-4c27-b404-74da17641025"
	reqRes.ReqData = []byte(`{}`)
	reqRes.ResBody = []byte(`{"application":{"id":"fd7fe349-f7da-4c27-b404-74da17641025","secret":"9dc463cf-8318-4f65-bc02-778424fdfd77","image":"default","name":"Test_application1","description":"description","redirect_uri":"http://localhost/login","url":"http://localhost","grant_type":"password,authorization_code,implicit","token_types":"jwt,permanent","jwt_secret":"3f1164da20d50c62","response_type":"code,token"}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--data={}", "--pretty", "--aid=fd7fe349-f7da-4c27-b404-74da17641025"})

	setJSONIndentError(ngsi)

	err := applicationsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := applicationsDelete(c)

	assert.NoError(t, err)
}

func TestApplicationsDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := applicationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := applicationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsDeleteErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"application":{"id":"0fbfa58c-e5b6-41c3-b748-ab29f1567a9c","name":"Test_application2","description":"Description","secret":"61f5def7-bcf9-45b1-9c69-d0887e403737","url":"http://localhost","redirect_uri":"http://localhost/login","image":"default","grant_type":"client_credentials,password,implicit,authorization_code,refresh_token","response_type":"code,token","token_types":"bearer,jwt,permanent","jwt_secret":"3f1164da20d50c62","client_type":null,"scope":null,"extra":null}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := applicationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify user id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := applicationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestApplicationsDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := applicationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestMakeAppBodyData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={\"application\":{}}"})

	actual, err := makeAppBody(c, ngsi)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{}}", string(actual))
	}
}

func TestMakeAppBodyParam1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "name,description,redirectUri,redirectSignOutUri,url,grantType,tokenTypes,resposeType,clientType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=app1", "--description=application1", "--redirectUri=http://ruri", "--redirectSignOutUri=http://suri", "--url=http://url"})

	actual, err := makeAppBody(c, ngsi)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{\"name\":\"app1\",\"description\":\"application1\",\"url\":\"http://url\",\"redirect_uri\":\"http://ruri\",\"redirect_sign_out_uri\":\"http://suri\"}}", string(actual))
	}
}

func TestMakeAppBodyParam2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "name,description,redirectUri,url,grantType,tokenTypes,responseType,clientType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--grantType=abc,def", "--tokenTypes=123,456", "--responseType=123,xyz", "--clientType=987,654"})

	actual, err := makeAppBody(c, ngsi)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"application\":{\"grant_type\":[\"abc\",\"def\"],\"response_type\":[\"123\",\"xyz\"],\"token_types\":[\"123\",\"456\"],\"client_type\":[\"987\",\"654\"]}}", string(actual))
	}
}

func TestMakeAppBodyErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data="})

	_, err := makeAppBody(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestMakeAppBodyErrorParam(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "name,description,redirectUri,url,grantType,tokenTypes,resposeType,clientType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=app1", "--description=application1", "--redirectUri=http://ruril", "--url=http://url"})

	setJSONEncodeErr(ngsi, 0)
	_, err := makeAppBody(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
