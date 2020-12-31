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
		ngsiErr := err.(*NgsiLibError)
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
		ngsiErr := err.(*NgsiLibError)
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
		ngsiErr := err.(*NgsiLibError)
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

func TestTokenInfo(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", Username: "fiware"}}

	_, err := ngsi.TokenInfo(client)

	assert.NoError(t, err)
}

func TestTokenInfoNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", Username: "fiware"}}

	_, err := ngsi.TokenInfo(client)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
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
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cTokenproxy, Username: "fiware", Password: "1234"}}

	_, err := ngsi.GetToken(client)

	assert.NoError(t, err)
}

func TestNgsiGetTokenExpires(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.TimeLib = &MockTimeLib{unixTime: 0}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}
	ngsi.tokenList["583a5c111b603ff8925585f48503e343403115f9"] = TokenInfo{Expires: 3600}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", Username: "fiware"}}

	_, err := ngsi.GetToken(client)

	assert.NoError(t, err)
}

func TestNgsiGetTokenNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/"}}

	_, err := ngsi.GetToken(client)

	assert.Error(t, err)
}

func TestGetToken(t *testing.T) {
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cPasswordCredentials, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	assert.NoError(t, err)
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
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{Expires: 1000}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	assert.NoError(t, err)
}

func TestGetTokenPasswordCredentials(t *testing.T) {
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cPasswordCredentials, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	assert.NoError(t, err)
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cKeyrocktokenprovider, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

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
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cTokenproxy, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	assert.NoError(t, err)
}

func TestGetTokenKeyrock(t *testing.T) {
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	assert.NoError(t, err)
}

func TestGetTokenErrorUsername(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.tokenList = tokenInfoList{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.LogWriter = &bytes.Buffer{}
	ngsi.HTTP = &MockHTTP{}
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", Username: "fiware"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: "fiware", IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
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

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
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
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", IdmType: cKeyrock, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}

	_, err := getToken(ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
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
		ngsiErr := err.(*NgsiLibError)
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
		ngsiErr := err.(*NgsiLibError)
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
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "encode error", ngsiErr.Message)
	}
}

func TestGetHash(t *testing.T) {
	client := &Client{Broker: &Broker{BrokerHost: "http://orion/", Username: "fiware"}}

	actual := getHash(client)
	expected := "583a5c111b603ff8925585f48503e343403115f9"

	assert.Equal(t, expected, actual)

}

func TestGetUserName(t *testing.T) {
	client := &Client{Broker: &Broker{Username: "fiware"}}

	actual, err := getUserName(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", actual)
	}

}

func TestGetPassword(t *testing.T) {
	client := &Client{Broker: &Broker{Password: "12345"}}

	actual, err := getPassword(client)

	if assert.NoError(t, err) {
		assert.Equal(t, "12345", actual)
	}

}
