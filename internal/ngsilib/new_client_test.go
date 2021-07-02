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
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	ngsi.Updated = true
	ngsi.PreviousArgs.Host = "http://localhost:1026"
	_, err := ngsi.NewClient("http://localhost:1026", flags, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "", ngsi.PreviousArgs.Host)
	}
}

func TestNewClientHTTP2(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	ngsi.Updated = true
	ngsi.PreviousArgs.Host = "localhost:1026"
	_, err := ngsi.NewClient("localhost:1026", flags, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "", ngsi.PreviousArgs.Host)
	}
}

func TestNewClientURL(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("http://orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientIPAdress(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("192.168.1.1?options=keyValues", flags, false)

	assert.NoError(t, err)
}

func TestNewClientTenatScope(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/"}
	ngsi.serverList["orion"] = broker

	ngsi.PreviousArgs = &Settings{Tenant: "test", Scope: "/test"}
	tenant := "fiware"
	scope := "/iot"
	flags := &CmdFlags{Tenant: &tenant, Scope: &scope}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientAPIPath(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", APIPath: "/,/orion"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientNgsiTypeV2(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", NgsiType: "ld"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientToken(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/"}
	ngsi.serverList["orion"] = broker

	ngsi.PreviousArgs = &Settings{Token: "b8ab85c5e7f8708b91dde91979729287b1dbd6d2"}
	token := "e08ff73ae501d19225152e426ea74d0c4fe458c2"
	flags := &CmdFlags{Token: &token}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientIdmType(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	ngsi.CacheFile = &MockIoLib{filename: &fileName}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token": "ad5252cd520cnaddacdc5d2e63899f0cdcf946f3", "expires_in": 3599, "refresh_token": "03e33a311e03317b390956729bcac2794b695670", "scope": [ "bearer" ], "token_type": "Bearer" }`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", IdmType: CTokenproxy, Username: "fiware", Password: "1234"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientSafeString(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	ngsi.CacheFile = &MockIoLib{filename: &fileName}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", SafeString: "on"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientSafeStringCmdFlag(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	ngsi.CacheFile = &MockIoLib{filename: &fileName}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", SafeString: "on"}
	ngsi.serverList["orion"] = broker

	safeString := "on"
	flags := &CmdFlags{SafeString: &safeString}

	_, err := ngsi.NewClient("orion", flags, false)

	assert.NoError(t, err)
}

func TestNewClientErrorURL(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("http://orion\n", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "illegal url: http://orion\n", ngsiErr.Message)
	}
}

func TestNewClientErrorHost(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: ""}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found", ngsiErr.Message)
	}
}

func TestNewClientErrorHost2NotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "orion-ld"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion-ld not found", ngsiErr.Message)
	}
}

func TestNewClientErrorHost2Empty(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "orion-ld"}
	ngsi.serverList["orion"] = broker
	broker2 := &Server{ServerHost: ""}
	ngsi.serverList["orion-ld"] = broker2

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error: orion-ld", ngsiErr.Message)
	}
}

func TestNewClientErrorHostNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("192.168.1", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error host: 192.168.1", ngsiErr.Message)
	}
}

func TestNewClientErrorURLParse(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion\n"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "illegal url: orion, http://orion\n", ngsiErr.Message)
	}
}

func TestNewClientErrorAPIPath(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", APIPath: "/"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "apiPath error: /", ngsiErr.Message)
	}
}

func TestNewClientErrorIdmType(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", IdmType: CKeyrock}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "username is required", ngsiErr.Message)
	}
}

func TestNewClientErrorSafeString(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	ngsi.CacheFile = &MockIoLib{filename: &fileName}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", SafeString: "enable"}
	ngsi.serverList["orion"] = broker

	flags := &CmdFlags{}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: enable", ngsiErr.Message)
	}
}

func TestNewClientErrorSafeStringCmdFlag(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	ngsi.CacheFile = &MockIoLib{filename: &fileName}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	InitServerList()

	broker := &Server{ServerHost: "http://orion/"}
	ngsi.serverList["orion"] = broker

	safeString := "enable"
	flags := &CmdFlags{SafeString: &safeString}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: enable", ngsiErr.Message)
	}
}

func TestNewClientInitHeader(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	ngsi.CacheFile = &MockIoLib{filename: &fileName}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	InitServerList()

	broker := &Server{ServerHost: "http://orion/", SafeString: "on"}
	ngsi.serverList["orion"] = broker

	tenant := "FIWARE"
	flags := &CmdFlags{Tenant: &tenant}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
		assert.Equal(t, "error FIWARE Service: FIWARE", ngsiErr.Message)
	}
}

func TestNewClientErrorSaveConfig(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	InitServerList()

	broker := &Server{ServerHost: "http://orion/"}
	ngsi.serverList["orion"] = broker

	tenant := "fiware"
	scope := "/iot"
	flags := &CmdFlags{Tenant: &tenant, Scope: &scope}

	_, err := ngsi.NewClient("orion", flags, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 12, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestSetTenantAndScope(t *testing.T) {
	client := &Client{Server: &Server{Tenant: "iot", Scope: "/device"}}

	setTenantAndScope(client, nil, nil)

	assert.Equal(t, "iot", client.Tenant)
	assert.Equal(t, "/device", client.Scope)
}

func TestSetTenantAndScopeTenant(t *testing.T) {
	client := &Client{Server: &Server{Tenant: "iot", Scope: "/device"}}

	tenant := "FIWARE"
	setTenantAndScope(client, &tenant, nil)

	assert.Equal(t, "FIWARE", client.Tenant)
	assert.Equal(t, "/device", client.Scope)
}

func TestSetTenantAndScopeScope(t *testing.T) {
	client := &Client{Server: &Server{Tenant: "iot", Scope: "/device"}}

	scope := "/iotagent"
	setTenantAndScope(client, nil, &scope)

	assert.Equal(t, "iot", client.Tenant)
	assert.Equal(t, "/iotagent", client.Scope)
}

func TestSetTenantAndScopeTeantAndScope(t *testing.T) {
	client := &Client{Server: &Server{Tenant: "iot", Scope: "/device"}}

	tenant := ""
	scope := ""
	setTenantAndScope(client, &tenant, &scope)

	assert.Equal(t, "", client.Tenant)
	assert.Equal(t, "", client.Scope)
}

func TestParseURLHost(t *testing.T) {
	host, path, query := parseURL("orion")

	assert.Equal(t, "orion", host)
	assert.Equal(t, "", path)
	assert.Equal(t, "", query)
}

func TestParseURLHostPath(t *testing.T) {
	host, path, query := parseURL("orion/version")

	assert.Equal(t, "orion", host)
	assert.Equal(t, "/version", path)
	assert.Equal(t, "", query)
}

func TestParseURLHostPathQuery(t *testing.T) {
	host, path, query := parseURL("orion/v2/entities?options=keyValues")

	assert.Equal(t, "orion", host)
	assert.Equal(t, "/v2/entities", path)
	assert.Equal(t, "options=keyValues", query)
}
