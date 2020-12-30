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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestAttrReadV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"

	AddReqRes(ngsi, reqRes)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "89\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAttrReadV2Pretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues,pretty")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":"fiware"}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"

	AddReqRes(ngsi, reqRes)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--pretty", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"name\": \"fiware\"\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAttrReadV2SafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":"%25"}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"

	AddReqRes(ngsi, reqRes)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--safeString=on"})
	err := attrRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"name\":\"%\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAttrReadLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "89\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAttrReadErrorInitCmd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrReadErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,link")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrReadErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("99")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrReadErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte("99")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestAttrReadV2ErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"

	AddReqRes(ngsi, reqRes)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--safeString=on"})
	err := attrRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error: {\"name\"", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAttrReadV2ErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"name":"fiware"}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	AddReqRes(ngsi, reqRes)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--pretty", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})

	setJSONIndentError(ngsi)

	err := attrRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrUpdateV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data=89"})
	err := attrUpdate(c)

	assert.NoError(t, err)
}

func TestAttrUpdateV2JSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("{\"value\":89}")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data={\"value\":89}"})
	err := attrUpdate(c)

	assert.NoError(t, err)
}

func TestAttrUpdateLD(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("99")
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data=99"})
	err := attrUpdate(c)

	assert.NoError(t, err)
}

func TestAttrUpdateLDJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("{\"value\":99}")
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data={\"value\":99}"})
	err := attrUpdate(c)

	assert.NoError(t, err)
}

func TestAttrUpdateLDJSONContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,context")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"value":89}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data={\"value\":89}", "--context=[\"http://context\"]"})
	err := attrUpdate(c)

	assert.NoError(t, err)
}

func TestAttrUpdateErrorInitCmd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.ReqData = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data=89"})
	err := attrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,link")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.ReqData = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data=89"})
	err := attrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrUpdateErrorReadALl(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.ReqData = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrUpdateLDJSONContextError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,context")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"value":89}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data={\"value\":89}", "--context=[\"http://context\""})
	err := attrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestAttrUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte("99")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data=89"})
	err := attrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "body data error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrUpdateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price/value"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price", "--data=89"})
	err := attrUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestAttrDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrDelete(c)

	assert.NoError(t, err)
}

func TestAttrDeleteErrorInitCmd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte("89")
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:001/attrs/price"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:001", "--attrName=price"})
	err := attrDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
