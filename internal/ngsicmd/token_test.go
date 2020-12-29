/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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

func TestVersionTokenCommand(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

func TestVersionTokenCommandJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

func TestVersionTokenCommandJSONPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

func TestVersionTokenCommandJSONExpiresZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

func TestVersionTokenCommandExpires(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

func TestVersionTokenCommandExpiresZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

// initCmd() Error: no host
func TestTokenCommandErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
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

func TestVersionTokenCommandErrorHostNotFound(t *testing.T) {
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

func TestVersionTokenCommandErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

func TestVersionTokenCommandErrorJSONPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"brokers": {
			"orion": {
				"brokerHost": "http://orion",
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

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{IndentErr: errors.New("json error"), Jsonlib: j}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--verbose", "--pretty"})
	err := tokenCommand(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
