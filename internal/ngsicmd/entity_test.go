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

func TestEntityCreateV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"})
	err := entityCreate(c)

	assert.NoError(t, err)
}

func TestEntityCreateV2SafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data,safeString")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"})
	err := entityCreate(c)

	assert.NoError(t, err)
}

func TestEntityCreateLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"})
	err := entityCreate(c)

	assert.NoError(t, err)
}

func TestEntityCreateLdContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ReqData = []byte(`{"@context":["http://context"],"id":"urn:ngsi-ld:Product:010","name":{"type":"Text","value":"Lemonade"},"price":{"type":"Integer","value":99},"size":{"type":"Text","value":"S"},"type":"Product"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data,context")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}", "--context=[\"http://context\"]"})
	err := entityCreate(c)

	assert.NoError(t, err)
}

func TestEntityCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityCreateErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityCreateErrorLdKeyValues(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--keyValues"})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "--keyValues only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityCreateErrorLdUpsert(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")
	setupFlagBool(set, "upsert")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--upsert"})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "--upsert only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestEntityCreateErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityCreateErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data,safeString")
	setJSONDecodeErr(ngsi, 1)

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={", "--safeString=on"})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error: {", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityCreateLdErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ReqData = []byte(`{"@context":["http://context"],"id":"urn:ngsi-ld:Product:010","name":{"type":"Text","value":"Lemonade"},"price":{"type":"Integer","value":99},"size":{"type":"Text","value":"S"},"type":"Product"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data,context")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}", "--context=[\"http://context\""})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityCreateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := entityCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestEntityReadV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010"})
	err := entityRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntityReadV2Pretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--pretty"})
	err := entityRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"id\": \"urn:ngsi-ld:Product:010\",\n  \"type\": \"Product\",\n  \"name\": {\n    \"type\": \"Text\",\n    \"value\": \"Lemonade\"\n  },\n  \"size\": {\n    \"type\": \"Text\",\n    \"value\": \"S\"\n  },\n  \"price\": {\n    \"type\": \"Integer\",\n    \"value\": 99\n  }\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntityReadV2SafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name%25":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id,safeString")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--safeString=on"})
	err := entityRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name%\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntityReadLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")
	setupFlagBool(set, "acceptJson")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:010", "--acceptJson"})
	err := entityRead(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestEntityReadErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := entityRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityReadErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := entityRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityReadErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := entityRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityReadErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010"})
	err := entityRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestEntityReadV2ErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type:"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id,safeString")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--safeString=on"})
	err := entityRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'P' after object key (39) ct:010\",\"type:\"Product\",\"name\"", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityReadV2ErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id":"urn:ngsi-ld:Product:010","type":"Product","name":{"type":"Text","value":"Lemonade"},"size":{"type":"Text","value":"S"},"price":{"type":"Integer","value":99}}`)
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,id")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010", "--pretty"})

	setJSONIndentError(ngsi)

	err := entityRead(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityUpsertV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")
	setupFlagBool(set, "keyValues")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--keyValues", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"})
	err := entityUpsert(c)

	assert.NoError(t, err)
}

func TestEntityUpsertV2SafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data,safeString")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--safeString=on", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"})
	err := entityUpsert(c)

	assert.NoError(t, err)
}

func TestEntityUpsertErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := entityUpsert(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityUpsertErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := entityUpsert(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityUpsertErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data={\"id\":\"urn:ngsi-ld:Product:010\",\"type\":\"Product\",\"name\":{\"type\":\"Text\",\"value\":\"Lemonade\"},\"size\":{\"type\":\"Text\",\"value\":\"S\"},\"price\":{\"type\":\"Integer\",\"value\":99}}"})
	err := entityUpsert(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityUpsertErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := entityUpsert(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityUpsertErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data,safeString")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={,}", "--safeString=on"})
	err := entityUpsert(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character ',' (1) {,}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityUpsertErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := entityUpsert(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityUpsertErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--data={}"})
	err := entityUpsert(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestEntityDeleteV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010"})
	err := entityDelete(c)

	assert.NoError(t, err)
}

func TestEntityDeleteLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/entities/urn:ngsi-ld:Product:010"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--id=urn:ngsi-ld:Product:010"})
	err := entityDelete(c)

	assert.NoError(t, err)
}

func TestEntityDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := entityDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityDeleteErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := entityDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Device"})
	err := entityDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestEntityDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities/urn:ngsi-ld:Product:010"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"8"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,id")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--id=urn:ngsi-ld:Product:010"})
	err := entityDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
