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

package management

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
)

func TestTokenCommand(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "c312d32a36a8a1df219a807a79323bb31941f462\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandRevoke(t *testing.T) {
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

	c := setupTestWithConfig([]string{"token", "--host", "orion", "--revoke"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion has no token", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandJSON(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--verbose"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"access_token\":\"c312d32a36a8a1df219a807a79323bb31941f462\",\"expires_in\":1156,\"refresh_token\":\"7cb75b47782195839ecbc7c7457f18abed853fe1\",\"scope\":[\"bearer\"],\"token_type\":\"Bearer\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandJSONPretty(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--pretty"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"access_token\": \"c312d32a36a8a1df219a807a79323bb31941f462\",\n  \"expires_in\": 1156,\n  \"refresh_token\": \"7cb75b47782195839ecbc7c7457f18abed853fe1\",\n  \"scope\": [\n    \"bearer\"\n  ],\n  \"token_type\": \"Bearer\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandJSONExpiresZero(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--verbose"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":0,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock
	c.Ngsi.TimeLib = &helper.MockTimeLib{UnixTime: 1200}

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"access_token\":\"c312d32a36a8a1df219a807a79323bb31941f462\",\"expires_in\":0,\"refresh_token\":\"7cb75b47782195839ecbc7c7457f18abed853fe1\",\"scope\":[\"bearer\"],\"token_type\":\"Bearer\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandExpires(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--expires"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock
	c.Ngsi.TimeLib = &helper.MockTimeLib{UnixTime: 1156}

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1156\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandExpiresZero(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--expires"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":-1,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock
	c.Ngsi.TimeLib = &helper.MockTimeLib{UnixTime: 2312}

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandKeyrock(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"idm": {
				"serverHost": "http://localhost:3000/",
				"serverType": "keyrock",
				"idmType": "idm",
				"idmHost": "http://keyrock/v1/auth/tokens",
				"username": "admin@letsfiware.jp",
				"password": "1234"
			}
		}
	}`
	c := setupTestWithConfig([]string{"token", "--host", "idm", "--verbose"}, conf)

	tokens := `{"version":"1", "tokens":{"b3193239b2d60a2dac06044845650c3470fdb1d5":{"type":"idm","expires":"2121-07-03T00:43:44.000Z","keyrock":{"token":{"methods":["password"],"expires_at":"2121-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`
	c.Ngsi.CacheFile = &helper.MockIoLib{Tokens: &tokens, Filename: helper.StrPtr("ngsi-token-cache.json")}

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"token":{"methods":["password"],"expires_at":"2121-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"token":"81868db8-d45c-4675-b68c-68860ba6b561"}`)
	reqRes1.Path = "/v1/auth/tokens"
	reqRes1.ResHeader = helper.NewHttpHeader("X-Subject-Token", "81868db8-d45c-4675-b68c-68860ba6b561")

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"access_token":"b832280e-a05c-4224-bb75-92c1df96a9b2","expires":"2031-09-04T23:05:05.000Z","valid":true,"User":{"scope":[],"id":"admin","username":"admin","email":"admin@test.com","date_password":"2019-01-01T02:06:40.000Z","enabled":true,"admin":true}}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	c.Ngsi.HTTP = mock

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"access_token\":\"b832280e-a05c-4224-bb75-92c1df96a9b2\",\"expires\":\"2031-09-04T23:05:05.000Z\",\"valid\":true,\"User\":{\"scope\":[],\"id\":\"admin\",\"username\":\"admin\",\"email\":\"admin@test.com\",\"date_password\":\"2019-01-01T02:06:40.000Z\",\"enabled\":true,\"admin\":true}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestTokenCommandErrorNewClient(t *testing.T) {
	c := setupTest([]string{"token", "--host", "unknown"})

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error host: unknown", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorHostNotFound(t *testing.T) {
	c := setupTest([]string{"token", "--host", "orion"})

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "orion has no token", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorKeyrock(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"idm": {
				"serverHost": "http://idm",
				"serverType": "keyrock",
				"idmType": "idm",
				"idmHost": "http://keyrock/v1/auth/tokens",
				"username": "testuser",
				"password": "1234"
			}
		}
	}`
	c := setupTestWithConfig([]string{"token", "--host", "idm", "--verbose"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"token":{"methods":["password"],"expires_at":"2121-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"token":"81868db8-d45c-4675-b68c-68860ba6b561"}`)
	reqRes.Path = "/v1/auth/tokens"

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "token is empty", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandBasic(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--verbose"}, conf)

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "no information available", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestTokenCommandErrorJSONPretty(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--verbose", "--pretty"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"c312d32a36a8a1df219a807a79323bb31941f462","expires_in":1156,"refresh_token":"7cb75b47782195839ecbc7c7457f18abed853fe1","scope":["bearer"],"token_type":"Bearer"}`)

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	helper.SetJSONIndentError(c.Ngsi)

	err := tokenCommand(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRevokeTokenCommand(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://localhost:3000/",
				"serverType": "broker",
				"ngsiType": "v2",
				"idmType": "basic",
				"username": "admin",
				"password": "1234"
			}
		}
	}`

	tokens := `{
		"version": "1",
		"tokens": {
			"8bc7807f8d65c12b9265801dfa8a6392b5236013": {
			"type": "tokenproxy",
			"token": "69e45641cf8e29d647929a3514e27874c031fbae",
			"refresh_token": "1a8346b8df2881c8b3407b0f39c80d1374204b93",
			"expires": "2031-01-01T00:00:08+09:00",
			"tokenproxy": {
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
	c := setupTestWithConfigAndCache([]string{"token", "--host", "orion", "--revoke"}, conf, tokens)

	err := revokeTokenCommand(c, c.Ngsi)

	assert.NoError(t, err)
}

func TestRevokeTokenCommandErrorOptions(t *testing.T) {
	c := setupTest([]string{"token", "--host", "orion", "--revoke", "--pretty"})

	err := revokeTokenCommand(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "only --revoke can be specified", ngsiErr.Message)
	}
}

func TestRevokeTokenCommandErrorRevoke(t *testing.T) {
	c := setupTest([]string{"token", "--host", "unknown", "--revoke"})

	err := revokeTokenCommand(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: unknown", ngsiErr.Message)
	}
}

func TestRevokeTokenCommandErrorTokenInfo(t *testing.T) {
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
	c := setupTestWithConfig([]string{"token", "--host", "orion", "--revoke"}, conf)

	err := revokeTokenCommand(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion has no token", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestRevokeTokenCommandErrorRevokeToken(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				"serverHost": "http://localhost:3000/",
				"serverType": "broker",
				"ngsiType": "v2",
				"idmType": "basic",
				"username": "admin",
				"password": "1234"
			}
		}
	}`

	tokens := `{
		"version": "1",
		"tokens": {
			"8bc7807f8d65c12b9265801dfa8a6392b5236013": {
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

	c := setupTestWithConfigAndCache([]string{"token", "--host", "orion", "--revoke"}, conf, tokens)

	c.Ngsi.CacheFile = &helper.MockIoLib{OpenErr: errors.New("save error"), Filename: helper.StrPtr("ngsi-config.json")}

	err := revokeTokenCommand(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "save error ngsi-config.json", ngsiErr.Message)
		assert.Error(t, err)
	}
}

func TestGetKeyrockUserInfoError(t *testing.T) {
	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"b832280e-a05c-4224-bb75-92c1df96a9b2","expires":"2031-09-04T23:05:05.000Z","valid":true,"User":{"scope":[],"id":"admin","username":"admin","email":"admin@test.com","date_password":"2019-01-01T02:06:40.000Z","enabled":true,"admin":true}}`)
	reqRes.Path = "/v1/auth/tokens"

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	var u url.URL
	client := &ngsilib.Client{HTTP: mock, Headers: map[string]string{}, URL: &u, Server: &ngsilib.Server{ServerType: "keyrock"}}

	actual, err := getKeyrockUserInfo(client, "1234")

	if assert.NoError(t, err) {
		expected := "{\"access_token\":\"b832280e-a05c-4224-bb75-92c1df96a9b2\",\"expires\":\"2031-09-04T23:05:05.000Z\",\"valid\":true,\"User\":{\"scope\":[],\"id\":\"admin\",\"username\":\"admin\",\"email\":\"admin@test.com\",\"date_password\":\"2019-01-01T02:06:40.000Z\",\"enabled\":true,\"admin\":true}}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestGetKeyrockUserInfoErrorToken(t *testing.T) {
	client := &ngsilib.Client{}
	_, err := getKeyrockUserInfo(client, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "token is empty", ngsiErr.Message)
	}
}

func TestGetKeyrockUserInfoErrorHTTP(t *testing.T) {
	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{}`)
	reqRes.Path = "/v1/auth/token"

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	var u url.URL
	client := &ngsilib.Client{HTTP: mock, Headers: map[string]string{}, URL: &u, Server: &ngsilib.Server{ServerType: "keyrock"}}

	_, err := getKeyrockUserInfo(client, "1234")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestGetKeyrockUserInfoErrorStatus(t *testing.T) {
	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Res.Status = "400"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes.Path = "/v1/auth/tokens"

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)

	var u url.URL
	client := &ngsilib.Client{HTTP: mock, Headers: map[string]string{}, URL: &u, Server: &ngsilib.Server{ServerType: "keyrock"}}

	_, err := getKeyrockUserInfo(client, "1234")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error 400 {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
