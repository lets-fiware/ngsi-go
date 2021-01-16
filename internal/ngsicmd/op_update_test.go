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
AUTHORS OR COPYRIGHT HOv2ERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package ngsicmd

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestOpUpdateArrayData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	ngsi.Host = "orion"
	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testOpUpdateArrayData})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	assert.NoError(t, err)
}

func TestOpUpdateArrayDataOver100(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	testData := "["
	for i := 0; i < 105; i++ {
		testData = testData + fmt.Sprintf("{\"id\":\"urn:ngsi-ld:Product:%d\",\"type\":\"Product\"},", i)
	}
	testData = testData[:len(testData)-1] + "]"
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNoContent
	reqRes1.Path = "/v2/op/update"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testData})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")
	assert.NoError(t, err)
}

func TestOpUpdateLineData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testOpUpdateLineData})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	assert.NoError(t, err)
}

func TestOpUpdateErrorReadAll(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data="})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

/*
func TestOpUpdateErrorClose(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testOpUpdateLineData})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	ngsi.FileReader = &MockFileLib{CloseError: errors.New("close error")}

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "close error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
*/

func TestOpUpdateErrorToken(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=}"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '}' looking for beginning of value", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpUpdateErrorJSONDelim(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is not JSON", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpUpdateErrorJSONDelim2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={{"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '{' looking for beginning of object key string (2)", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpUpdateErrorDecode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=[" + testOpUpdateArrayData})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal array into Go value of type map[string]interface {}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestOpUpdateArrayErrorHTTP2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	testData := "["
	for i := 0; i < 105; i++ {
		testData = testData + fmt.Sprintf("{\"id\":\"urn:ngsi-ld:Product:%d\",\"type\":\"Product\"},", i)
	}
	testData = testData[:len(testData)-1] + "]"
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNoContent
	reqRes1.Err = errors.New("error")
	reqRes1.Path = "/v2/op/update"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testData})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpUpdateArrayErrorHTTP2StatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	testData := "["
	for i := 0; i < 105; i++ {
		testData = testData + fmt.Sprintf("{\"id\":\"urn:ngsi-ld:Product:%d\",\"type\":\"Product\"},", i)
	}
	testData = testData[:len(testData)-1] + "]"
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.Path = "/v2/op/update"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testData})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestOpUpdateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestOpUpdateArrayDataError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/op/update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data=" + testOpUpdateArrayDataError})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false)
	assert.NoError(t, err)

	err = opUpdate(c, ngsi, client, "append_strict")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 12, ngsiErr.ErrNo)
		assert.Equal(t, "EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

var testOpUpdateArrayData = `[
	{
	  "id":"urn:ngsi-ld:Product:001",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Brandy"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1299}
	},
	{
	  "id":"urn:ngsi-ld:Product:002",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Port"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1199}
	},
	{
	  "id":"urn:ngsi-ld:Product:003",
	  "type":"Product",
	  "offerPrice":{"type":"Integer", "value": 59}
	}
  ]`

var testOpUpdateArrayDataError = `[
	{
	  "id":"urn:ngsi-ld:Product:001",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Brandy"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1299}
	},
	{
	  "id":"urn:ngsi-ld:Product:002",
	  "type":"Product",
	  "name":{"type":"Text", "value":"Port"},
	  "size":{"type":"Text", "value": "M"},
	  "price":{"type":"Integer", "value": 1199}
	},
	{
	  "id":"urn:ngsi-ld:Product:003",
	  "type":"Product",
	  "offerPrice":{"type":"Integer", "value": 59}
	}
  `

var testOpUpdateLineData = `{"id":"urn:ngsi-ld:Product:001","type":"Product","name":{"type":"Text","value":"Brandy"},"size":{"type":"Text","value":"M"},"price":{"type":"Integer","value":1299}}
{"id":"urn:ngsi-ld:Product:002","type":"Product","name":{"type":"Text","value":"Port"},"size":{"type":"Text","value":"M"},"price":{"type":"Integer","value":1199}}
{"id":"urn:ngsi-ld:Product:003","type":"Product","offerPrice":{"type":"Integer","value":59}}`
