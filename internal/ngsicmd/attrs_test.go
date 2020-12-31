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

package ngsicmd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestAttrsReadV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=airqualityobserved1", "--attrName=CO"})
	err := attrsRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"CO\":{\"type\":\"Number\",\"value\":400.463869544,\"metadata\":{}}}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadV2Pretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues,pretty")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=airqualityobserved1", "--attrName=CO", "--pretty"})
	err := attrsRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"CO\": {\n    \"type\": \"Number\",\n    \"value\": 400.463869544,\n    \"metadata\": {}\n  }\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadV2SafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=airqualityobserved1", "--attrName=CO"})
	err := attrsRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"CO\":{\"type\":\"Number\",\"value\":400.463869544,\"metadata\":{}}}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadErrorInitCmd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=airqualityobserved1", "--attrName=CO"})
	err := attrsRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,link")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion", "--id=airqualityobserved1", "--attrName=CO"})
	err := attrsRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities/airqualityobserved1/attr"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=airqualityobserved1", "--attrName=CO"})
	err := attrsRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=airqualityobserved1", "--attrName=CO"})
	err := attrsRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadV2ErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value"400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=airqualityobserved1", "--attrName=CO"})
	err := attrsRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		expected := "invalid character '4' after object key (30) Number\",\"value\"400.463869544,\""
		assert.Equal(t, expected, ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReadV2ErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"CO":{"type":"Number","value":400.463869544,"metadata":{}}}`)
	reqRes.Path = "/v2/entities/airqualityobserved1/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=airqualityobserved1", "--attrName=CO", "--pretty"})

	setJSONIndentError(ngsi)

	err := attrsRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsAppendV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsAppend(c)

	assert.NoError(t, err)
}

func TestAttrsAppendV2SafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsAppend(c)

	assert.NoError(t, err)
}

func TestAttrsAppendLD(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,context")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"specialOffer":{"value":true}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}", "--context=[\"http://context\"]"})
	err := attrsAppend(c)

	assert.NoError(t, err)
}

func TestAttrsAppendErrorInitCmd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsAppendErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,link")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsAppendErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010"})
	err := attrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsAppendLDErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,context")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"specialOffer":{"value":true}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}", "--context=[\"http://context\""})
	err := attrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsAppendV2SafeStringError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value: true}"})
	err := attrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsAppendErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attr"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsAppendErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsAppend(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestAttrsUpdateV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsUpdate(c)

	assert.NoError(t, err)
}

func TestAttrsUpdateV2SafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsUpdate(c)

	assert.NoError(t, err)
}

func TestAttrsUpdateLD(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,context")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"specialOffer":{"value":true}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}", "--context=[\"http://context\"]"})
	err := attrsUpdate(c)

	assert.NoError(t, err)
}

func TestAttrsUpdateErrorInitCmd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,link")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsUpdateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010"})
	err := attrsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsUpdateLDErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,context")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"@context":["http://context"],"specialOffer":{"value":true}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}", "--context=[\"http://context\""})
	err := attrsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsUpdateV2SafeStringError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}"})
	err := attrsUpdate(c)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error: :{\"value\":true}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attr"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsUpdateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestAttrsReplaceV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsReplace(c)

	assert.NoError(t, err)
}

func TestAttrsReplaceV2SafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsReplace(c)

	assert.NoError(t, err)
}

func TestAttrsReplaceErrorInitCmd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsReplace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReplaceErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,link")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsReplace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReplaceErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010"})
	err := attrsReplace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReplaceV2SafeStringError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data,safeString")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}"})
	err := attrsReplace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error: :{\"value\":true}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReplaceErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attr"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data={\"specialOffer\":{\"value\": true}}"})
	err := attrsReplace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestAttrsReplaceErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,id,type,attrName,data")
	setupFlagBool(set, "append,keyValues")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010/attrs"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--data=\"specialOffer\":{\"value\": true}"})
	err := attrsReplace(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
