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
		assert.Equal(t, 6, ngsiErr.ErrNo)
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
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
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
