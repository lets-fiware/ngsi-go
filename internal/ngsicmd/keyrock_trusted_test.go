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

func TestTrustedAppList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d",]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := trustedAppList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"trusted_applications\":[\"8692ec57-8514-4ef6-a347-3d1ac6409f79\",\"78b4763f-139a-4820-a42b-3e265fb9d56e\",\"462ee067-f10a-4c9c-aefe-079038830043\",\"6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d\",]}"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d"]}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pretty"})

	err := trustedAppList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"trusted_applications\": [\n    \"8692ec57-8514-4ef6-a347-3d1ac6409f79\",\n    \"78b4763f-139a-4820-a42b-3e265fb9d56e\",\n    \"462ee067-f10a-4c9c-aefe-079038830043\",\n    \"6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppListNotFound(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := trustedAppList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "Trusted applications nof found\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := trustedAppList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/trustedApp"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := trustedAppList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppListErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d"]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := trustedAppList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d"]}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := trustedAppList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := trustedAppList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications"
	reqRes.ResBody = []byte(`{"trusted_applications":["8692ec57-8514-4ef6-a347-3d1ac6409f79","78b4763f-139a-4820-a42b-3e265fb9d56e","462ee067-f10a-4c9c-aefe-079038830043","6781fd6c-9dd3-46d7-bdcc-4ca2af1ae42d"]}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	setJSONIndentError(ngsi)

	err := trustedAppList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestTrustedAppAdd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := trustedAppAdd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"oauth_client_id\": \"3e34cbf6-3579-4e54-878a-0a406962ac36\",\"trusted_oauth_client_id\": \"0118ccb7-756e-42f9-8a19-5b4e83ca8c46\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppAddPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := trustedAppAdd(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"oauth_client_id\": \"3e34cbf6-3579-4e54-878a-0a406962ac36\",\n  \"trusted_oauth_client_id\": \"0118ccb7-756e-42f9-8a19-5b4e83ca8c46\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTrustedAppAddErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := trustedAppAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppAddErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := trustedAppAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppAddErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/trustedApp"
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := trustedAppAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppAddErrorTid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/trustedApp"
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := trustedAppAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify trusted application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppAddErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := trustedAppAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppAddErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := trustedAppAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppAddErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"oauth_client_id": "3e34cbf6-3579-4e54-878a-0a406962ac36","trusted_oauth_client_id": "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--pretty", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	setJSONIndentError(ngsi)

	err := trustedAppAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := trustedAppDelete(c)

	assert.NoError(t, err)
}

func TestTrustedAppDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := trustedAppDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := trustedAppDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppDeleteErrorAid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock"})

	err := trustedAppDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppDeleteErrorTid(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"})

	err := trustedAppDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "specify trusted application id", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := trustedAppDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestTrustedAppDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/applications/0fbfa58c-e5b6-41c3-b748-ab29f1567a9c/trusted_applications/0118ccb7-756e-42f9-8a19-5b4e83ca8c46"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,aid,tid")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=keyrock", "--aid=0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid=0118ccb7-756e-42f9-8a19-5b4e83ca8c46"})

	err := trustedAppDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
