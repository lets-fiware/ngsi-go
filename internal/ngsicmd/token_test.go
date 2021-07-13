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

func TestTokenCommand(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "c312d32a36a8a1df219a807a79323bb31941f462\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandRevoke(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "revoke")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--revoke"})

	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion has no token", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"access_token\":\"c312d32a36a8a1df219a807a79323bb31941f462\",\"expires_in\":1156,\"refresh_token\":\"7cb75b47782195839ecbc7c7457f18abed853fe1\",\"scope\":[\"bearer\"],\"token_type\":\"Bearer\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandJSONPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "verbose,pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--pretty"})
	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"access_token\": \"c312d32a36a8a1df219a807a79323bb31941f462\",\n  \"expires_in\": 1156,\n  \"refresh_token\": \"7cb75b47782195839ecbc7c7457f18abed853fe1\",\n  \"scope\": [\n    \"bearer\"\n  ],\n  \"token_type\": \"Bearer\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandJSONExpiresZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":0,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.TimeLib = &MockTimeLib{unixTime: 1200}
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"access_token\":\"c312d32a36a8a1df219a807a79323bb31941f462\",\"expires_in\":0,\"refresh_token\":\"7cb75b47782195839ecbc7c7457f18abed853fe1\",\"scope\":[\"bearer\"],\"token_type\":\"Bearer\"}\n"
		assert.Equal(t, expected, actual)
	}

}

func TestTokenCommandExpires(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.TimeLib = &MockTimeLib{unixTime: 1156}
	setupFlagString(set, "host")
	set.Bool("expires", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--expires"})
	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "1156\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandExpiresZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":-1,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.TimeLib = &MockTimeLib{unixTime: 2312}
	setupFlagString(set, "host")
	set.Bool("expires", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--expires"})
	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandKeyrock(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"idm": {
				"serverHost": "http://localhost:3000/",
				"serverType": "keyrock",
				"idmType": "idm",
				"idmHost": "http://idm",
				"username": "admin@letsfiware.jp",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	tokens := `{"version":"1", "tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"type":"idm","expires":"2121-07-03T00:43:44.000Z","keyrock":{"token":{"methods":["password"],"expires_at":"2121-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`
	ngsi.CacheFile = &MockIoLib{Tokens: &tokens}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"access_token":"7921ac63-57bc-4063-8d92-816b6b4b118a","expires":"2121-07-03T00:43:44.000Z","valid":true,"User":{"scope":[],"id":"admin","username":"admin","email":"admin@letsfilware.jp","date_password":"2019-11-09T02:06:40.000Z","enabled":true,"admin":true}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{}`)
	reqRes2.Path = "/v1/auth/token"
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=idm", "--verbose"})

	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"access_token\":\"7921ac63-57bc-4063-8d92-816b6b4b118a\",\"expires\":\"2121-07-03T00:43:44.000Z\",\"valid\":true,\"User\":{\"scope\":[],\"id\":\"admin\",\"username\":\"admin\",\"email\":\"admin@letsfilware.jp\",\"date_password\":\"2019-11-09T02:06:40.000Z\",\"enabled\":true,\"admin\":true}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandBasic(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "basic",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})

	err := tokenCommand(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "no information available\n"
		assert.Equal(t, expected, actual)
	}
}

// initCmd() Error: no host
func TestTokenCommandErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorHostNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":0,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion has no token", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorOAuthJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	setJSONEncodeErr(ngsi, 0)

	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorThinkingCitesJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "thinkingCities",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"token":{"domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"methods":["password"],"roles":[],"expires_at":"2021-04-16T11:30:47.000000Z","catalog":[],"extras":{"password_creation_time":"2021-04-16T08:29:01Z","last_login_attempt_time":"2021-04-16T08:29:05.000000","pwd_user_in_blacklist":false,"password_expiration_time":"2022-04-16T08:29:01Z"},"user":{"password_expires_at":"2022-04-16T08:29:00.000000","domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"id":"80e292b7dae445e7af66c284162ff049","name":"usertest"},"audit_ids":["6kJ9zBFCQaKRa7aCFc6bpw"],"issued_at":"2021-04-16T08:30:47.000000Z"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	setJSONEncodeErr(ngsi, 1)

	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorKeycloakJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "keycloak",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234",
				"clientId": "11111111-2222-3333-4444-555555555555",
				"clientSecret": "66666666-7777-8888-9999-000000000000"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA","expires_in":300,"refresh_expires_in":1800,"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJjOTUzMTYwYy0yMGEyLTRiNDQtYmI5NC01ZDgyNjQ2ODZmMzUifQ.eyJleHAiOjE2MjU5NTkzMzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiYzYzZjhhYmEtYWIzOC00OWVkLWExYjYtMjNhZjcyYzQ1MmI4IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXV0aC9yZWFsbXMvZml3YXJlX3NlcnZpY2UiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoibmdzaV9hcGkiLCJzZXNzaW9uX3N0YXRlIjoiMGZkZGZhZDctODA1Yi00MzcxLTkwMDktZmIxOTI3YWZkYjAyIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSJ9.rJ7sMlxKZ-IA0AoqKznZLZnNAK7xsKHoPIRc2owIsw0","token_type":"Bearer","id_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwiYXV0aF90aW1lIjowLCJqdGkiOiI4YzA4YWJiYy05YjI1LTRmMDgtOWMwYi1iMDRmNmEwYWJjNmMiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0L2F1dGgvcmVhbG1zL2Zpd2FyZV9zZXJ2aWNlIiwiYXVkIjoibmdzaV9hcGkiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJJRCIsImF6cCI6Im5nc2lfYXBpIiwic2Vzc2lvbl9zdGF0ZSI6IjBmZGRmYWQ3LTgwNWItNDM3MS05MDA5LWZiMTkyN2FmZGIwMiIsImF0X2hhc2giOiJfalh0eDQxbVlobWRrZGsyc1h5Y0FnIiwiYWNyIjoiMSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.JYjRzJckKsYhVChFw1Y2SGHJgY_PXhjTZRJMDzls7Xj5J7nJ5IiglHY1cSI-A1wzcJHttJYvGwKRK9QziVZfQh2LQoQSChnhHYY1Uq0VMrWfuOgofCQSqOmoXJiu7VwFRWVbg6RNVxlgT_z1SyncXQrCmjtfUBUKsBSSAZxOucVgDmHptT_JcNqKPeeo8-7PelDtx4PZDkf4Qf_77qHFPy0cwXio57UFAGtJAsztxd6nwZ6Q0QQY7XxCQLLALIsJeYJzfB2b58YwTdnpmSHG6oMrWp_Ie-P8cYkHgmNmI_Q1KIYuWYwA6NRqL26rC5CwN7irxn2sgEwShBNeqwhJlA","not-before-policy":0,"session_state":"0fddfad7-805b-4371-9009-fb1927afdb02","scope":"openid email profile"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	setJSONEncodeErr(ngsi, 0)

	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorWSO2JSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "wso2",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234",
				"clientId": "11111111-2222-3333-4444-555555555555",
				"clientSecret": "66666666-7777-8888-9999-000000000000"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"scope":"default","token_type":"Bearer","expires_in":3600,"refresh_token":"a7d6bae2b1d36c041787e9c9e2d6cbf8","access_token":"cba95432f1f8227f5bc6cf4a20633cb3"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose"})
	setJSONEncodeErr(ngsi, 0)

	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorKeyrock(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"idm": {
				"serverHost": "http://idm",
				"serverType": "keyrock",
				"idmType": "idm",
				"idmHost": "http://idm",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	set.Bool("verbose", false, "doc")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=idm", "--verbose"})

	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "token is empty", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorJSONPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "verbose,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--pretty"})

	setJSONIndentError(ngsi)

	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRevokeTokenCommand(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://localhost:3000/",
				"serverType": "broker",
				"ngsiType": "v2",
				"idmType": "basic",
				"idmHost": "http://orion",
				"username": "admin",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	file := "cache"
	tokens := `{
		"version": "1",
		"tokens": {
		  "67c97b389b0b38551db075a6e5d0ba998d5bcb2a": {
			"type": "tokenproxy",
			"token": "69e45641cf8e29d647929a3514e27874c031fbae",
			"refresh_token": "1a8346b8df2881c8b3407b0f39c80d1374204b93",
			"expires": "2031-01-01T00:00:08+09:00",
			"Oauth": {
			  "access_token": "69e45641cf8e29d647929a3514e27874c031fbae",
			  "expires_in": 3599,
			  "refresh_token": "1a8346b8df2881c8b3407b0f39c80d1374204b93",
			  "scope": [
				"bearer"
			  ],
			  "token_type": "Bearer"
			}
		  }
		}
	}`
	ngsi.CacheFile = &MockIoLib{filename: &file, Tokens: &tokens}

	setupFlagString(set, "host")
	setupFlagBool(set, "revoke")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--revoke"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	err = revokeTokenCommand(c, ngsi)

	assert.NoError(t, err)
}

func TestRevokeTokenCommandErrorOptions(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}
	ngsi.CacheFile = &MockIoLib{Trunc: []error{nil, errors.New("encode error")}}

	setupFlagString(set, "host")
	setupFlagBool(set, "revoke,verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--revoke", "--verbose"})

	err := revokeTokenCommand(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "only --revoke can be specified", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRevokeTokenCommandErrorRevoke(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host")
	setupFlagBool(set, "revoke")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--revoke"})

	err := revokeTokenCommand(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRevokeTokenCommandErrorTokenInfo(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://orion",
				"ngsiType": "v2",
				"idmType": "tokenproxy",
				"idmHost": "/token",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host")
	setupFlagBool(set, "revoke")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--revoke"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	err = revokeTokenCommand(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion has no token", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestRevokeTokenCommandErrorRevokeToken(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://localhost:3000/",
				"serverType": "broker",
				"ngsiType": "v2",
				"idmType": "basic",
				"idmHost": "http://orion",
				"username": "admin",
				"password": "1234"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}
	file := "cache"
	tokens := `{
		"version": "1",
		"tokens": {
		  "67c97b389b0b38551db075a6e5d0ba998d5bcb2a": {
			"type": "tokenproxy",
			"token": "69e45641cf8e29d647929a3514e27874c031fbae",
			"refresh_token": "1a8346b8df2881c8b3407b0f39c80d1374204b93",
			"expires": "2031-01-01T00:00:08+09:00",
			"Oauth": {
			  "access_token": "69e45641cf8e29d647929a3514e27874c031fbae",
			  "expires_in": 3599,
			  "refresh_token": "1a8346b8df2881c8b3407b0f39c80d1374204b93",
			  "scope": [
				"bearer"
			  ],
			  "token_type": "Bearer"
			}
		  }
		}
	}`
	ngsi.CacheFile = &MockIoLib{filename: &file, Tokens: &tokens, Trunc: []error{errors.New("encode error")}}

	setupFlagString(set, "host")
	setupFlagBool(set, "revoke")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--revoke"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	err = revokeTokenCommand(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "encode error", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestGetKeyrockUserInfoError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.Path = "/v1/auth/tokens"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	actual, err := getKeyrockUserInfo(client, "1234")

	if assert.NoError(t, err) {
		expected := []byte("{}")
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestGetKeyrockUserInfoErrorToken(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.Path = "/v1/auth/token"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	_, err = getKeyrockUserInfo(client, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "token is empty", ngsiErr.Message)
	}
}

func TestGetKeyrockUserInfoErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.Path = "/v1/auth/token"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	_, err = getKeyrockUserInfo(client, "1234")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestGetKeyrockUserInfoErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Res.Status = "400"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes.Path = "/v1/auth/tokens"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=keyrock"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	_, err = getKeyrockUserInfo(client, "1234")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error 400 {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
