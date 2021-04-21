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

package ngsilib

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitTokenMgr(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	filename := "cache-file"

	err := ngsi.InitTokenMgr(&filename)

	assert.NoError(t, err)
}

func TestInitTokenMgrFileNil(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}

	err := ngsi.InitTokenMgr(nil)
	assert.NoError(t, err)
}

func TestInitTokenMgrNoFileName(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	filename := ""

	err := ngsi.InitTokenMgr(&filename)
	assert.NoError(t, err)
}

func TestInitTokenMgrErrorgetConfigDir(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{HomeDirErr: errors.New("error homedir")}
	ngsi.LogWriter = &bytes.Buffer{}

	err := ngsi.InitTokenMgr(nil)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error homedir", ngsiErr.Message)
	}
}

func TestInitTokenMgrErrorPathAbs(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{PathAbsErr: errors.New("path abs error")}
	ngsi.LogWriter = &bytes.Buffer{}
	filename := "cache-file"

	err := ngsi.InitTokenMgr(&filename)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "path abs error cache-file", ngsiErr.Message)
	}
}

func TestInitTokenMgrErrorInitTokenList(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{OpenErr: errors.New("open error")}
	ngsi.LogWriter = &bytes.Buffer{}
	filename := "cache-file"

	err := ngsi.InitTokenMgr(&filename)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "open error cache-file", ngsiErr.Message)
	}
}

func TestInitTokenList(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{}
	filename := "cache-file"
	io.SetFileName(&filename)

	err := initTokenList(io)

	assert.NoError(t, err)
}

func TestInitTokenListNoExistsFile(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{StatErr: errors.New("stat error")}
	filename := "cache-file"
	io.SetFileName(&filename)

	err := initTokenList(io)

	assert.NoError(t, err)
}

func TestInitTokenListNoFile(t *testing.T) {
	testNgsiLibInit()
	io := &ioLib{}
	filename := ""
	io.SetFileName(&filename)

	err := initTokenList(io)

	assert.NoError(t, err)
}

func TestTokenList(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.tokenList["token1"] = TokenInfo{}

	actual := ngsi.TokenList()
	expected := "token1"
	assert.Equal(t, expected, actual)
}

func TestInitTokenListErrorOpen(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{OpenErr: errors.New("open error")}
	filename := "cache-file"
	io.SetFileName(&filename)

	err := initTokenList(io)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

/*
func TestInitTokenListErrorClose(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{CloseErr: errors.New("close error")}
	filename := "cache-file"
	io.SetFileName(&filename)

	err := initTokenList(io)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "close error", ngsiErr.Message)
	}
}
*/

func TestInitTokenListErrorDecode(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{DecodeErr: errors.New("decode error")}
	filename := "cache-file"
	io.SetFileName(&filename)

	err := initTokenList(io)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "decode error", ngsiErr.Message)
	}
}

func TestTokenInfo(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware"}}

	_, err := ngsi.TokenInfo(client)

	assert.NoError(t, err)
}

func TestTokenInfoNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware"}}

	_, err := ngsi.TokenInfo(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "not found", ngsiErr.Message)
	}
}

func TestNgsiGetToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cTokenproxy, Username: "fiware", Password: "1234"}}

	actual, err := ngsi.GetToken(client)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestNgsiGetTokenKeyrock(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	token := "1234"
	ngsi.tokenList["9e7067026d0aac494e8fedf66b1f585e79f52935"] = TokenInfo{KeyrockToken: &token, Expires: 9613169598}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: cKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}

	actual, err := ngsi.GetToken(client)

	if assert.NoError(t, err) {
		expected := "1234"
		assert.Equal(t, expected, actual)
	}
}

func TestNgsiGetTokenExpires(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.TimeLib = &MockTimeLib{unixTime: 0}
	token := OauthToken{AccessToken: "123456"}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Expires: 3600, OauthToken: &token}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware"}}

	actual, err := ngsi.GetToken(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "123456", actual)
	}
}

func TestNgsiGetTokenThinkingCities(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.TimeLib = &MockTimeLib{unixTime: 0}
	token := "123456"
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Expires: 3600, KeystoneToken: &token}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware", IdmType: cThinkingCities}}

	actual, err := ngsi.GetToken(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "123456", actual)
	}
}

func TestNgsiGetTokenErrorTokenList(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	ngsi.tokenList["9e7067026d0aac494e8fedf66b1f585e79f52935"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: cKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}

	_, err := ngsi.GetToken(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "token list error", ngsiErr.Message)
	}
}

func TestNgsiGetTokenNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/"}}

	_, err := ngsi.GetToken(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "username is required", ngsiErr.Message)
	}
}

func TestGetToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cPasswordCredentials, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	actual, err := getToken(ngsi, client)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestGetTokenKeyrockIDM(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: cKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}

	actual, err := getToken(ngsi, client)

	if assert.NoError(t, err) {
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestNgsiGetTokenThinkingCitiesIDM(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResBody = []byte(`{"token":{"domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"methods":["password"],"roles":[],"expires_at":"2021-04-16T11:30:47.000000Z","catalog":[],"extras":{"password_creation_time":"2021-04-16T08:29:01Z","last_login_attempt_time":"2021-04-16T08:29:05.000000","pwd_user_in_blacklist":false,"password_expiration_time":"2022-04-16T08:29:01Z"},"user":{"password_expires_at":"2022-04-16T08:29:00.000000","domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"id":"80e292b7dae445e7af66c284162ff049","name":"usertest"},"audit_ids":["6kJ9zBFCQaKRa7aCFc6bpw"],"issued_at":"2021-04-16T08:30:47.000000Z"}}`)
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: cThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}

	actual, err := getToken(ngsi, client)

	if assert.NoError(t, err) {
		expected := "gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"
		assert.Equal(t, expected, actual)
	}
}

func TestGetTokenExpires(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.TimeLib = &MockTimeLib{unixTime: 0}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{Expires: 1000}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	actual, err := getToken(ngsi, client)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestGetTokenPasswordCredentials(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cPasswordCredentials, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	actual, err := getToken(ngsi, client)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestGetTokenKeyrocktokenprovider(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrocktokenprovider, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	assert.NoError(t, err)
}

func TestGetTokenTokenproxy(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cTokenproxy, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	actual, err := getToken(ngsi, client)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestGetTokenKeyrock(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	actual, err := getToken(ngsi, client)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestGetTokenErrorUsername(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "username is required", ngsiErr.Message)
	}
}

func TestGetTokenErrorPassword(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestGetTokenErrorIdmType(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: "fiware", IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: fiware", ngsiErr.Message)
	}
}

func TestGetTokenErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := "cache-file"
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Err = errors.New("http error")
	reqRes.Res.StatusCode = http.StatusBadRequest
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestGetTokenErrorHTTPStatus(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := "cache-file"
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestGetTokenErrorJSONUnmarshal6(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.JSONConverter = &MockJSONLib{DecodeErr: errors.New("decode error")}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "decode error", ngsiErr.Message)
	}
}

func TestGetTokenErrorJSONUnmarshal7(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.JSONConverter = &MockJSONLib{DecodeErr: errors.New("decode error")}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrocktokenprovider, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "decode error", ngsiErr.Message)
	}
}

func TestGetTokenErrorKeyrockIDM(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: cKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal string into Go value of type ngsilib.KeyrockToken Field: (8) \"tokens\":{\"9e7067026d0a", ngsiErr.Message)
	}
}

func TestNgsiGetTokenErrorThinkingCitiesIDM(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ResBody = []byte(`{"token":{"domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"methods":["password"],"roles":[],"expires_at":"2021-04-16T11:30:47.000000Z","catalog":[],"extras":{"password_creation_time":"2021-04-16T08:29:01Z","last_login_attempt_time":"2021-04-16T08:29:05.000000","pwd_user_in_blacklist":false,"password_expiration_time":"2022-04-16T08:29:01Z"},"user":{"password_expires_at":"2022-04-16T08:29:00.000000","domain":{"id":"9f60e700f04544379932d59a17985cff","name":"smartcity"},"id":"80e292b7dae445e7af66c284162ff049","name":"usertest"},"audit_ids":["6kJ9zBFCQaKRa7aCFc6bpw"],"issued_at":"2021-04-16T08:30:47.000000Z"}`)
	reqRes.ResHeader = http.Header{"X-Subject-Token": []string{"gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: cThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestGetTokenErrorSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := "cache-file"
	ngsi.CacheFile = &MockIoLib{filename: &filename, EncodeErr: errors.New("encode error")}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "encode error", ngsiErr.Message)
	}
}

func TestSaveToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	tokens := make(map[string]interface{})

	err := saveToken("cache-file", tokens)
	assert.NoError(t, err)
}

func TestSaveTokenNoFileName(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	tokens := make(map[string]interface{})

	err := saveToken("", tokens)
	assert.NoError(t, err)
}

func TestSaveTokenErrorOpenFile(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{OpenErr: errors.New("open error")}
	ngsi.LogWriter = &bytes.Buffer{}
	tokens := make(map[string]interface{})

	err := saveToken("cache-file", tokens)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "open error cache-file", ngsiErr.Message)
	}
}

func TestSaveTokenErrorTruncate(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{TruncateErr: errors.New("truncate error")}
	ngsi.LogWriter = &bytes.Buffer{}
	tokens := make(map[string]interface{})

	err := saveToken("cache-file", tokens)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "truncate error", ngsiErr.Message)
	}
}

func TestSaveTokenErrorEncode(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{EncodeErr: errors.New("encode error")}
	ngsi.LogWriter = &bytes.Buffer{}
	tokens := make(map[string]interface{})

	err := saveToken("cache-file", tokens)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "encode error", ngsiErr.Message)
	}
}

func TestGetHash(t *testing.T) {
	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware"}}

	actual := getHash(client)
	expected := "583a5c111b603ff8925585f48503e343403115f9"

	assert.Equal(t, expected, actual)

}

func TestGetHashThinkingCities(t *testing.T) {
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: cThinkingCities, Username: "fiware", Tenant: "smartcity", Scope: "/madrid"}}

	actual := getHash(client)
	expected := "a50ef2c09c126141f16967c62ef78a4c031bdb6f"

	assert.Equal(t, expected, actual)

}

func TestGetUserName(t *testing.T) {
	client := &Client{Server: &Server{Username: "fiware"}}

	actual, err := getUserName(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", actual)
	}

}

func TestGetPassword(t *testing.T) {
	client := &Client{Server: &Server{Password: "12345"}}

	actual, err := getPassword(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "12345", actual)
	}

}
