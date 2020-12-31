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
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddContex(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")

	assert.NoError(t, err)
}

func TestAddContexJSON(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", `["https://fiware.org/"]`)

	assert.NoError(t, err)
}

func TestAddContexErrorAlreadyExists(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	err = ngsi.AddContext("fiware", "https://fiware.org/")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware already exists", ngsiErr.Message)
	}
}

func TestAddContexErrorNotUrl(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "fiware.org")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "fiware.org is neither url nor json", ngsiErr.Message)
	}
}

func TestAddContexErrorSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	err := ngsi.AddContext("fiware", "https://fiware.org/")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestUpdateContex(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	err = ngsi.UpdateContext("fiware", "http://fiware.org")
	assert.NoError(t, err)
}

func TestUpdateContexErrorNotUrl(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	err = ngsi.UpdateContext("fiware", "fiware.org")
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware.org is not url", ngsiErr.Message)
	}
}

func TestUpdateContexErrorSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.Error(t, err)

	err = ngsi.UpdateContext("fiware", "http://fiware.org")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestUpdateContexErrorNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	err = ngsi.UpdateContext("core", "http://fiware.org")
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "core not found", ngsiErr.Message)
	}
}

func TestDeleteContex(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	err = ngsi.DeleteContext("fiware")
	assert.NoError(t, err)
}

func TestDeleteContexErrorReferenced(t *testing.T) {
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

	err = ngsi.DeleteContext("fiware")
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware is referenced", ngsiErr.Message)
	}
}

func TestDeleteContexErrorSave(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	fileName = "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	err = ngsi.DeleteContext("fiware")
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestDeleteContexErrorNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	fileName = "config"
	ngsi.ConfigFile = &MockIoLib{filename: &fileName, OpenErr: errors.New("open error")}

	err = ngsi.DeleteContext("core")
	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "core not found", ngsiErr.Message)
	}
}

func TestGetContext(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	actual, err := ngsi.GetContext("fiware")
	expected := "https://fiware.org/"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

}

func TestGetContextHttpValue(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}

	context := "https://fiware.org/"

	actual, err := ngsi.GetContext(context)
	expected := context

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

}
func TestAddContext(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")

	assert.NoError(t, err)
}

func TestGetContextErrorOtherType(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	ngsi.contextList["fiware"] = 1

	actual, err := ngsi.GetContext("fiware")
	expected := ""

	if assert.Error(t, err) {
		assert.Equal(t, expected, actual)
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware neither url nor json", ngsiErr.Message)
	}
}

func TestGetContextErrorString(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	ngsi.contextList["fiware"] = "context"

	actual, err := ngsi.GetContext("fiware")
	expected := ""

	if assert.Error(t, err) {
		assert.Equal(t, expected, actual)
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "fiware is not url", ngsiErr.Message)
	}
}

func TestGetContextErrorJSON(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	var v []interface{}
	json.Unmarshal([]byte(`[]`), &v)
	ngsi.contextList["fiware"] = v

	actual, err := ngsi.GetContext("fiware")
	expected := ""

	if assert.Error(t, err) {
		assert.Equal(t, expected, actual)
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestGetContextErrorNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	actual, err := ngsi.GetContext("core")
	expected := ""

	if assert.Error(t, err) {
		assert.Equal(t, expected, actual)
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "core not found", ngsiErr.Message)
	}
}

func TestGetContextHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	actual, err := ngsi.GetContextHTTP("fiware")
	expected := "https://fiware.org/"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestGetContextHTTPErrorNoFound(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	actual, err := ngsi.GetContextHTTP("core")
	expected := ""

	if assert.Error(t, err) {
		assert.Equal(t, expected, actual)
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "core not found", ngsiErr.Message)
	}
}

func TestGetContextHTTPErrorJSON(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	array := []string{"https://fiware.org/"}
	b, _ := json.Marshal(array)
	err := ngsi.AddContext("fiware", string(b))
	assert.NoError(t, err)

	actual, err := ngsi.GetContextHTTP("fiware")
	expected := ""

	if assert.Error(t, err) {
		assert.Equal(t, expected, actual)
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "fiware is not url", ngsiErr.Message)
	}
}

func TestGetContextList(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.contextList = ContextsInfo{}
	fileName := ""
	ngsi.ConfigFile = &MockIoLib{filename: &fileName}

	err := ngsi.AddContext("fiware", "https://fiware.org/")
	assert.NoError(t, err)

	err = ngsi.AddContext("core", "http://fiware.org/")
	assert.NoError(t, err)

	info := ngsi.GetContextList()

	assert.Equal(t, "https://fiware.org/", info["fiware"])
	assert.Equal(t, "http://fiware.org/", info["core"])
}
