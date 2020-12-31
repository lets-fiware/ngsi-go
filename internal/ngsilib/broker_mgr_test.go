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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBroker(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)

	if assert.NoError(t, err) {
		actual := ngsi.brokerList["orion"]
		expected := &Broker{BrokerHost: "http://orion"}
		assert.Equal(t, expected, actual)
	}
}

func TestCreateBrokerErrorParam(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["host"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "brokerHost not found", ngsiErr.Message)
	}
}

func TestCreateBrokerErrorSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestUpdateBroker(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	param = make(map[string]string)
	param["brokerHost"] = "https://orion-ld"
	err = ngsi.UpdateBroker("orion", param)

	if assert.NoError(t, err) {
		actual := ngsi.brokerList["orion"]
		expected := &Broker{BrokerHost: "https://orion-ld"}
		assert.Equal(t, expected, actual)
	}
}

func TestUpdateBrokerErrorParam(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	param = make(map[string]string)
	param["host"] = "https://orion-ld"
	err = ngsi.UpdateBroker("orion", param)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "host not found", ngsiErr.Message)
	}
}

func TestUpdateBrokerErrorCheckParam(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	param = make(map[string]string)
	param["brokerHost"] = "orion-ld"
	err = ngsi.UpdateBroker("orion", param)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "brokerHost error: orion-ld", ngsiErr.Message)
	}
}

func TestUpdateBrokerErrorSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	fileName = "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	param = make(map[string]string)
	param["brokerHost"] = "https://orion-ld"
	err = ngsi.UpdateBroker("orion", param)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestUpdateBrokerErrorNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	param = make(map[string]string)
	param["brokerHost"] = "orion-ld"
	err = ngsi.UpdateBroker("orion-ld", param)

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "orion-ld not found", ngsiErr.Message)
	}
}

func TestDeleteBroker(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	param = make(map[string]string)
	param["brokerHost"] = "https://orion-ld"
	err = ngsi.DeleteBroker("orion")

	assert.NoError(t, err)
}

func TestDeleteBrokerErrorSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	fileName = "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	param = make(map[string]string)
	param["brokerHost"] = "https://orion-ld"
	err = ngsi.DeleteBroker("orion")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestDeleteBrokerErrorNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	InitBrokerList()

	param := make(map[string]string)
	param["brokerHost"] = "http://orion"

	err := ngsi.CreateBroker("orion", param)
	assert.NoError(t, err)

	param = make(map[string]string)
	param["brokerHost"] = "orion-ld"
	err = ngsi.DeleteBroker("orion-ld")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "orion-ld not found", ngsiErr.Message)
	}
}
