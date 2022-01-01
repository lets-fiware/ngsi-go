/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"time"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
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
		ngsiErr := err.(*ngsierr.NgsiError)
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
		ngsiErr := err.(*ngsierr.NgsiError)
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
		ngsiErr := err.(*ngsierr.NgsiError)
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

func TestInitTokenListVersion1(t *testing.T) {
	testNgsiLibInit()
	tokens := `{"version":"1", "tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"type":"idm","expires":"2121-07-03T00:43:44.000Z","keyrock":{"token":{"methods":["password"],"expires_at":"2121-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`
	io := &MockIoLib{Data: &tokens}
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
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
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
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "not found", ngsiErr.Message)
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, Username: "fiware", Password: "1234"}}

	actual, err := ngsi.GetToken(client)

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
	reqRes.ResBody = []byte(`{"tokens":{"9e7067026d0aac494e8fedf66b1f585e79f52935":{"expires":1613170563,"keyrock":{"token":{"methods":["password"],"expires_at":"2021-02-12T22:56:03.410Z"},"idm_authorization_config":{"level":"basic","authzforce":false}},"keyrock_token":"81868db8-d45c-4675-b68c-68860ba6b561"}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	token := "1234"
	ngsi.tokenList["b3193239b2d60a2dac06044845650c3470fdb1d5"] = TokenInfo{Token: token, Expires: time.Unix(9613169598, 0)}

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: CKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}

	actual, err := ngsi.GetToken(client)

	if assert.NoError(t, err) {
		expected := "1234"
		assert.Equal(t, expected, actual)
	}
}

func TestGetTokenExpires(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.TimeLib = &MockTimeLib{unixTime: 0}
	token := "123456"
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Token: token, Expires: time.Unix(3600, 0)}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware"}}

	actual, err := ngsi.GetToken(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "123456", actual)
	}
}

func TestGetTokenThinkingCities(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.TimeLib = &MockTimeLib{unixTime: 0}
	token := "123456"
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Token: token, Expires: time.Unix(3600, 0)}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware", IdmType: CThinkingCities}}

	actual, err := ngsi.GetToken(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "123456", actual)
	}
}

func TestGetTokenErrorRequestToken(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: "unknown", IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}

	_, err := ngsi.GetToken(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: unknown", ngsiErr.Message)
	}
}

func TestGetAuthHeader(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, Username: "fiware", Password: "1234"}}

	key, value, err := ngsi.GetAuthHeader(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "Authorization", key)
		assert.Equal(t, "Bearer ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", value)
	}
}

func TestGetAuthHeaderErrorGetToken(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: "unknown", Username: "fiware", Password: "1234"}}

	_, _, err := ngsi.GetAuthHeader(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: unknown", ngsiErr.Message)
	}
}

func TestGetAuthHeaderErrorIdmType(t *testing.T) {
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
	ngsi.tokenList["d40fc58b9c5fc05c623f7e046c5b447954901c02"] = TokenInfo{Token: "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", Expires: time.Unix(9999999999, 0)}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: "unknown", Username: "fiware", Password: "1234"}}

	_, _, err := ngsi.GetAuthHeader(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: unknown", ngsiErr.Message)
	}
}

func TestRequestToken(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CPasswordCredentials, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	tokenInfo := &TokenInfo{}

	actual, err := requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestRequestTokenMainKeyrockIDM(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://localhost:3000/", ServerType: "keyrock", IdmType: CKeyrockIDM, IdmHost: "http://localhost:3000/", Username: "admin@letsfiware.jp", Password: "1234"}}
	tokenInfo := &TokenInfo{}

	actual, err := requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestRequestTokenThinkingCitiesIDM(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion:1026/", IdmType: CThinkingCities, IdmHost: "http://localhost:5001/v3/auth/tokens", Username: "usertest", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}
	tokenInfo := &TokenInfo{}

	actual, err := requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		expected := "gAAAAABgeojDoWDHy9r4Lq1sNRbss2ncweTzmQ5jBpefFI5eYFh6fA3DyzQM8mjzoiGqrUH6JNWl4Sk1XVVMwTf18eFJ7FluEkPklrM_AFSGXv1IO0j_Dy-UQxNUAEYyxqT8Ny3O2TNC78MOKkt2UoR3oOg4HBcjkf6iCsVFwPhW9BGjC37LWdk"
		assert.Equal(t, expected, actual)
	}
}

func TestRequestTokenExpires(t *testing.T) {
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
	ngsi.tokenList["token1"] = TokenInfo{Expires: time.Unix(1000, 0)}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	tokenInfo := &TokenInfo{}

	actual, err := requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		expected := "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3"
		assert.Equal(t, expected, actual)
	}
}

func TestRequestTokenErrorIdmType(t *testing.T) {
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
	tokenInfo := &TokenInfo{}

	_, err := requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: fiware", ngsiErr.Message)
	}
}

func TestRequestTokenErrorRequestToken(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	tokenInfo := &TokenInfo{}

	_, err := requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenErrorSave(t *testing.T) {
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

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	tokenInfo := &TokenInfo{}

	_, err := requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "encode error", ngsiErr.Message)
	}
}

func TestAppendToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	hash := "token3"
	tokenInfo := &TokenInfo{}

	actual := appendToken(ngsi, hash, tokenInfo)

	_, ok := (*actual)["token3"]

	assert.Equal(t, 1, len(*actual))
	assert.Equal(t, true, ok)
}

func TestSaveToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	tokenInfo := &tokenInfoList{}

	err := saveToken("cache-file", tokenInfo)
	assert.NoError(t, err)
}

func TestSaveTokenNoFileName(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	tokenInfo := &tokenInfoList{}

	err := saveToken("", tokenInfo)
	assert.NoError(t, err)
}

func TestSaveTokenErrorOpenFile(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{OpenErr: errors.New("open error")}
	ngsi.LogWriter = &bytes.Buffer{}
	tokenInfo := &tokenInfoList{}

	err := saveToken("cache-file", tokenInfo)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "open error cache-file", ngsiErr.Message)
	}
}

func TestSaveTokenErrorTruncate(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{TruncateErr: errors.New("truncate error")}
	ngsi.LogWriter = &bytes.Buffer{}
	tokenInfo := &tokenInfoList{}

	err := saveToken("cache-file", tokenInfo)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "truncate error", ngsiErr.Message)
	}
}

func TestSaveTokenErrorEncode(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.CacheFile = &MockIoLib{EncodeErr: errors.New("encode error")}
	ngsi.LogWriter = &bytes.Buffer{}
	tokenInfo := &tokenInfoList{}

	err := saveToken("cache-file", tokenInfo)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "encode error", ngsiErr.Message)
	}
}

func TestGetHash(t *testing.T) {
	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware", Password: "1234"}}

	actual := getHash(client)
	expected := "d40fc58b9c5fc05c623f7e046c5b447954901c02"

	assert.Equal(t, expected, actual)

}

func TestGetHashThinkingCities(t *testing.T) {
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CThinkingCities, Username: "fiware", Password: "1234", Tenant: "smartcity", Scope: "/madrid"}}

	actual := getHash(client)
	expected := "40b60171102fa5b505ef1929cf5f23267ba050d6"

	assert.Equal(t, expected, actual)

}

func TestGetHashKeycloak(t *testing.T) {
	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, Username: "fiware", Password: "1234", ClientID: "orion", ClientSecret: "1234"}}

	actual := getHash(client)
	expected := "0565444d5249dc3ceb130af73a8700dbf068485c"

	assert.Equal(t, expected, actual)

}

func TestGetUserName(t *testing.T) {
	client := &Client{Server: &Server{Username: "fiware"}}

	actual, err := getUserName(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", actual)
	}

}

func TestGetUserNameError(t *testing.T) {
	client := &Client{Server: &Server{}}

	_, err := getUserName(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "username is required", ngsiErr.Message)
	}

}

func TestGetPassword(t *testing.T) {
	client := &Client{Server: &Server{Password: "12345"}}

	actual, err := getPassword(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "12345", actual)
	}

}

func TestGetPasswordError(t *testing.T) {
	client := &Client{Server: &Server{}}

	_, err := getPassword(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}

}

func TestGetUserNamePassword(t *testing.T) {
	client := &Client{Server: &Server{Username: "fiware", Password: "12345"}}

	user, password, err := getUserNamePassword(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", user)
		assert.Equal(t, "12345", password)
	}

}

func TestGetUserNamePasswordErrorUser(t *testing.T) {
	client := &Client{Server: &Server{Password: "12345"}}

	_, _, err := getUserNamePassword(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "username is required", ngsiErr.Message)
	}

}

func TestGetUserNamePasswordErrorPassword(t *testing.T) {
	client := &Client{Server: &Server{Username: "fiware"}}

	_, _, err := getUserNamePassword(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestRevodeToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Token: "token", Expires: time.Unix(3600, 0)}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware", IdmType: CBasic}}

	err := ngsi.RevokeToken(client)

	assert.NoError(t, err)
}

func TestRevodeTokenErrorIdm(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Token: "token", Expires: time.Unix(3600, 0)}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware", IdmType: "unknown"}}

	err := ngsi.RevokeToken(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: unknown", ngsiErr.Message)
	}
}

func TestRevodeTokenErrorRevokeToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := "file"
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Token: "token", Expires: time.Unix(3600, 0)}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Err = errors.New("reovke token error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware", IdmType: "keyrock"}}

	err := ngsi.RevokeToken(client)

	assert.NoError(t, err)
}

func TestRevodeTokenErrorSaveToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	filename := "file"
	ngsi.CacheFile = &MockIoLib{filename: &filename, OpenErr: errors.New("open error")}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Token: "token", Expires: time.Unix(3600, 0)}

	client := &Client{Server: &Server{ServerHost: "http://orion/", Username: "fiware", IdmType: CBasic}}

	err := ngsi.RevokeToken(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "open error file", ngsiErr.Message)
	}
}

func TestGetTokenInfo(t *testing.T) {
	ngsi := testNgsiLibInit()

	tokenInfo := &TokenInfo{
		Type: CKeyrock,
		Keyrock: &KeyrockToken{
			AccessToken:  "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3",
			ExpiresIn:    3599,
			RefreshToken: "03e33a311e03317b390956729bcac2794b695670",
			Scope:        []string{"bearer"},
			TokenType:    "Bearer",
		},
	}

	actual, err := ngsi.GetTokenInfo(tokenInfo)

	if assert.NoError(t, err) {
		expected := "{\"access_token\":\"ad5252cd520cnaddacdc5d2e63899f0cdcf946f3\",\"expires_in\":3599,\"refresh_token\":\"03e33a311e03317b390956729bcac2794b695670\",\"scope\":[\"bearer\"],\"token_type\":\"Bearer\"}"
		assert.Equal(t, expected, string(actual))
	}
}

func TestGetTokenInfoErrorUnkownType(t *testing.T) {
	ngsi := testNgsiLibInit()

	tokenInfo := &TokenInfo{
		Type: "unknown",
	}

	_, err := ngsi.GetTokenInfo(tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: unknown", ngsiErr.Message)
	}
}

func TestGetTokenInfoErrorJSON(t *testing.T) {
	ngsi := testNgsiLibInit()

	tokenInfo := &TokenInfo{
		Type: CKeyrock,
		Keyrock: &KeyrockToken{
			AccessToken:  "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3",
			ExpiresIn:    3599,
			RefreshToken: "03e33a311e03317b390956729bcac2794b695670",
			Scope:        []string{"bearer"},
			TokenType:    "Bearer",
		},
	}

	gNGSI.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}}

	_, err := ngsi.GetTokenInfo(tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestCheckIdmParams(t *testing.T) {
	idmParams := &IdmParams{
		IdmType:      CKeyrock,
		IdmHost:      "https://keyrock/oauth2/token",
		Username:     "keyrock001@letsfiware.jp",
		Password:     "1234",
		ClientID:     "00000000-1111-2222-3333-444444444444",
		ClientSecret: "55555555-6666-7777-8888-999999999999",
	}

	err := checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsNoIDM(t *testing.T) {
	idmParams := &IdmParams{}

	err := checkIdmParams(idmParams)

	assert.NoError(t, err)
}

func TestCheckIdmParamsErrorNoIDM(t *testing.T) {
	idmParams := &IdmParams{
		Username: "fiware",
	}

	err := checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required idmType not found", ngsiErr.Message)
	}
}

func TestCheckIdmParamsErrorUnkownIDM(t *testing.T) {
	idmParams := &IdmParams{
		IdmType: "unknown",
	}

	err := checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: unknown", ngsiErr.Message)
	}
}

func TestCheckIdmParamsErrorParam(t *testing.T) {
	idmParams := &IdmParams{
		IdmType:  CKeyrock,
		IdmHost:  "https://keyrock/oauth2/token",
		Username: "keyrock001@letsfiware.jp",
		Password: "1234",
		ClientID: "00000000-1111-2222-3333-444444444444",
	}

	err := checkIdmParams(idmParams)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "idmHost, username, password, clientID and clientSecret are needed", ngsiErr.Message)
	}
}
