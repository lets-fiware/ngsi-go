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

package ngsilib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAllParams(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	err = ngsi.checkAllParams(host)

	assert.NoError(t, err)
}

func TestCheckAllParamsErrorBrokerHost(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.BrokerHost = ""
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "brokerHost not found", ngsiErr.Message)
	}
}

func TestCheckAllParamsErrorBrokerHostNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.BrokerHost = "orion-ld"
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "brokerHost error: orion-ld", ngsiErr.Message)
	}
}

func TestCheckAllParamsErrorNgsiType(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.NgsiType = "v1"
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "v1 not found", ngsiErr.Message)
	}
}

func TestCheckAllParamsErrorAPIPath(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.APIPath = "/"
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "apiPath error: /", ngsiErr.Message)
	}
}

func TestCheckAllParamsErrorIdmParams(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.IdmType = "unknown"
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "idmType error: unknown", ngsiErr.Message)
	}
}

func TestCheckAllParamsErrorTenant(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.Tenant = "FIWARE"
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error FIWARE Service: FIWARE", ngsiErr.Message)
	}
}

func TestCheckAllParamsErrorScope(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.Scope = "Scope"
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error FIWARE ServicePath: Scope", ngsiErr.Message)
	}
}

func TestCheckAllParamsErrorSafeString(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	host := ngsi.brokerList["orion"]
	host.SafeString = "none"
	err = ngsi.checkAllParams(host)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "unkown parameter: none", ngsiErr.Message)
	}
}

func TestGetAPIPath(t *testing.T) {
	b, a, err := getAPIPath("/,/api")

	if assert.NoError(t, err) {
		assert.Equal(t, "/", b)
		assert.Equal(t, "/api", a)
	}
}

func TestGetAPIPathErrorIndex(t *testing.T) {
	_, _, err := getAPIPath("/")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "apiPath error: /", ngsiErr.Message)
	}
}

func TestGetAPIPathErrorBeforePath(t *testing.T) {
	_, _, err := getAPIPath("path,path")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "apiPath error: path", ngsiErr.Message)
	}
}

func TestGetAPIPathErrorAfterPath(t *testing.T) {
	_, _, err := getAPIPath("/,path")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "must start with '/': path", ngsiErr.Message)
	}
}

func TestGetAPIPathErrorAfterPathTail(t *testing.T) {
	_, _, err := getAPIPath("/,/path/")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "trailing '/' is not required: /path/", ngsiErr.Message)
	}
}

func TestCheckIdmParams(t *testing.T) {
	err := checkIdmParams(cKeyrock, "https://keyrock/oauth2/token", "keyrock001@fiware", "0123456789", "00000000-1111-2222-3333-444444444444", "55555555-6666-7777-8888-999999999999")

	assert.NoError(t, err)
}

func TestCheckIdmParamsNoIdm(t *testing.T) {
	err := checkIdmParams("", "", "", "", "", "")

	assert.NoError(t, err)
}

func TestCheckIdmParamsErrorNoIdmType(t *testing.T) {
	err := checkIdmParams("", "https://keyrock/oauth2/token", "keyrock001@fiware", "0123456789", "00000000-1111-2222-3333-444444444444", "55555555-6666-7777-8888-999999999999")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required idmType not found", ngsiErr.Message)
	}
}

func TestCheckIdmParamsErrorIdmType(t *testing.T) {
	err := checkIdmParams("fiware", "https://keyrock/oauth2/token", "keyrock001@fiware", "0123456789", "00000000-1111-2222-3333-444444444444", "55555555-6666-7777-8888-999999999999")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "idmType error: fiware", ngsiErr.Message)
	}
}

func TestCheckIdmParamsErrorNoIdmHost(t *testing.T) {
	err := checkIdmParams(cKeyrock, "", "keyrock001@fiware", "0123456789", "00000000-1111-2222-3333-444444444444", "55555555-6666-7777-8888-999999999999")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "required idmHost not found", ngsiErr.Message)
	}
}

func TestCheckIdmParamsErrorIdmHost(t *testing.T) {
	err := checkIdmParams(cKeyrock, "fiware", "keyrock001@fiware", "0123456789", "00000000-1111-2222-3333-444444444444", "55555555-6666-7777-8888-999999999999")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "idmHost error: fiware", ngsiErr.Message)
	}
}

func TestCheckIdmParamsErrorClientId(t *testing.T) {
	err := checkIdmParams(cKeyrock, "http://keyrock", "keyrock001@fiware", "0123456789", "", "55555554-6666-7777-8888-999999999999")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "clientID and clientSecret are needed", ngsiErr.Message)
	}
}

func TestCheckIdmParamsErrorUserPassword(t *testing.T) {
	err := checkIdmParams(cKeyrock, "http://keyrock", "", "0123456789", "00000000-1111-2222-3333-444444444444", "55555555-6666-7777-8888-999999999999")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "username is needed", ngsiErr.Message)
	}
}

func TestExistsBrokerHostTrue(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	b := ngsi.ExistsBrokerHost("orion")
	assert.Equal(t, true, b)
}

func TestExistsBrokerHostFalse(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	b := ngsi.ExistsBrokerHost("orion-ld")
	assert.Equal(t, false, b)
}

func TestServerInfoArgs(t *testing.T) {
	ngsi := testNgsiLibInit()
	args := ngsi.ServerInfoArgs()

	assert.Equal(t, brokerArgs, args)
}

func TestCopyBrokerInfo(t *testing.T) {
	broker := Broker{}
	param := make(map[string]string)
	param[cBrokerHost] = "orion"
	param[cNgsiType] = "v2"
	param[cAPIPath] = "/,/orion"
	param[cIdmType] = cKeyrock
	param[cIdmHost] = "https://keyrock"
	param[cToken] = "00000000000000000"
	param[cUsername] = "fiware"
	param[cPassword] = "123"
	param[cClientID] = "111111111111"
	param[cClientSecret] = "222222222222"
	param[cContext] = "http://context"
	param[cFiwareService] = "iot"
	param[cFiwareServicePath] = "/iot"
	param[cSafeString] = "off"
	param[cXAuthToken] = "on"
	setBrokerParam(&broker, param)

	broker2 := Broker{}

	copyBrokerInfo(&broker, &broker2)

	assert.Equal(t, broker, broker2)
}

func TestCopyBrokerInfo2(t *testing.T) {
	broker := Broker{}
	param := make(map[string]string)
	param[cBrokerHost] = "orion"
	param[cNgsiType] = "v2"
	param[cAPIPath] = "/,/orion"
	param[cIdmType] = cKeyrock
	param[cIdmHost] = "https://keyrock"
	param[cToken] = "00000000000000000"
	param[cUsername] = "fiware"
	param[cPassword] = "123"
	param[cClientID] = "111111111111"
	param[cClientSecret] = "222222222222"
	param[cContext] = "http://context"
	param[cFiwareService] = "iot"
	param[cFiwareServicePath] = "/iot"
	param[cSafeString] = "off"
	param[cXAuthToken] = "on"
	setBrokerParam(&broker, param)

	broker2 := Broker{}

	copyBrokerInfo(&broker2, &broker)

	expected := Broker{}

	assert.Equal(t, expected, broker2)
}

func TestSetBrokerParam(t *testing.T) {
	broker := Broker{}
	param := make(map[string]string)
	param[cBrokerHost] = "orion"
	param[cNgsiType] = "v2"
	param[cAPIPath] = "/,/orion"
	param[cIdmType] = cKeyrock
	param[cIdmHost] = "https://keyrock"
	param[cToken] = "00000000000000000"
	param[cUsername] = "fiware"
	param[cPassword] = "123"
	param[cClientID] = "111111111111"
	param[cClientSecret] = "222222222222"
	param[cContext] = "http://context"
	param[cFiwareService] = "iot"
	param[cFiwareServicePath] = "/iot"
	param[cSafeString] = "off"
	param[cXAuthToken] = "on"
	err := setBrokerParam(&broker, param)

	assert.NoError(t, err)
}

func TestSetBrokerParamError(t *testing.T) {
	broker := Broker{}
	param := make(map[string]string)
	param["fiware"] = "orion"
	err := setBrokerParam(&broker, param)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestDeleteItem(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	param["safeString"] = "on"
	err := ngsi.CreateBroker("orion", param)

	if assert.NoError(t, err) {
		actual := ngsi.brokerList["orion"].SafeString
		expected := "on"
		assert.Equal(t, expected, actual)
	}

	err = ngsi.DeleteItem("orion", "safeString")

	if assert.NoError(t, err) {
		actual := ngsi.brokerList["orion"].SafeString
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestDeleteItemErrorHostNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	param["safeString"] = "on"
	err := ngsi.CreateBroker("orion", param)

	err = ngsi.DeleteItem("orion-ld", "safeString")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "orion-ld not found", ngsiErr.Message)
	}
}

func TestDeleteItemErrorItem(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	param["safeString"] = "on"
	err := ngsi.CreateBroker("orion", param)

	err = ngsi.DeleteItem("orion", "SafeString")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "SafeString not found", ngsiErr.Message)
	}
}

func TestIsHostReferenced(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)

	err = ngsi.IsHostReferenced("orion")

	assert.NoError(t, err)
}

func TestIsHostReferencedError(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"
	err := ngsi.CreateBroker("orion", param)

	param = make(map[string]string)
	param["brokerHost"] = "orion"
	err = ngsi.CreateBroker("fiware", param)

	err = ngsi.IsHostReferenced("orion")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "orion is referenced in fiware", ngsiErr.Message)
	}
}

func TestIsContextReferenced(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	InitBrokerList()

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	fileName = ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	orion := Broker{BrokerHost: "http://orion/", Context: "fiware"}
	ngsi.brokerList["orion"] = &orion

	err = ngsi.IsContextReferenced("orion")

	assert.NoError(t, err)
}

func TestIsContextReferencedError(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	InitBrokerList()

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	fileName = ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	orion := Broker{BrokerHost: "http://orion/", Context: "fiware"}
	ngsi.brokerList["orion"] = &orion

	err = ngsi.IsContextReferenced("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware is referenced in orion", ngsiErr.Message)
	}
}

func TestIsIdmTypeTrue(t *testing.T) {
	b := isIdmType(cKeyrock)
	assert.Equal(t, b, true)
}

func TestIsIdmTypeFalse(t *testing.T) {
	b := isIdmType("orion")
	assert.Equal(t, b, false)
}

func TestSafeStringTrue(t *testing.T) {
	info := Broker{SafeString: "on"}

	b, err := info.safeString()

	if assert.NoError(t, err) {
		assert.Equal(t, b, true)
	}
}

func TestSafeStringFalse(t *testing.T) {
	info := Broker{SafeString: "off"}

	b, err := info.safeString()

	if assert.NoError(t, err) {
		assert.Equal(t, b, false)
	}
}

func TestSafeStringError(t *testing.T) {
	info := Broker{SafeString: "error"}

	b, err := info.safeString()

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unkown parameter: error", ngsiErr.Message)
		assert.Equal(t, b, false)
	}
}

func TestXAuthTokenTrue(t *testing.T) {
	info := Broker{XAuthToken: "on"}

	b, err := info.xAuthToken()

	if assert.NoError(t, err) {
		assert.Equal(t, b, true)
	}
}

func TestXAuthTokenFalse(t *testing.T) {
	info := Broker{XAuthToken: "off"}

	b, err := info.xAuthToken()

	if assert.NoError(t, err) {
		assert.Equal(t, b, false)
	}
}

func TestXAuthTokenError(t *testing.T) {
	info := Broker{XAuthToken: "error"}

	b, err := info.xAuthToken()

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unkown parameter: error", ngsiErr.Message)
		assert.Equal(t, b, false)
	}
}
