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
	"archive/zip"
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestWcResourcesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "NGSIGO/ngsigo-mashup/1.0.0\nNGSIGO/ngsigo-operator/1.0.0\nNGSIGO/ngsigo-widget/1.0.0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListVender(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--vender=FIWARE", "--name=NGSIGO", "--version=2.0.0"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListName(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--name=NGSIGO", "--version=2.0.0"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListVersion(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--version=2.0.0"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--json"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"NGSIGO/ngsigo-widget/1.0.0\":{\"type\":\"widget\",\"vendor\":\"LetsFIWARE\",\"version\":\"1.0.0\"},\"NGSIGO/ngsigo-operator/1.0.0\":{\"type\":\"operator\",\"vendor\":\"LetsFIWARE\",\"version\":\"1.0.0\"},\"NGSIGO/ngsigo-mashup/1.0.0\":{\"type\":\"mashup\",\"vendor\":\"LetsFIWARE\",\"version\":\"1.0.0\"}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListJSON2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--widget", "--json"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"NGSIGO/ngsigo-widget/1.0.0\":{\"authors\":null,\"changelog\":\"\",\"contributors\":null,\"default_lang\":\"\",\"description\":\"\",\"doc\":\"\",\"email\":\"\",\"homepage\":\"\",\"image\":\"\",\"issuetracker\":\"\",\"js_files\":null,\"license\":\"\",\"licenseurl\":\"\",\"longdescription\":\"\",\"name\":\"\",\"preferences\":null,\"properties\":null,\"requirements\":null,\"smartphoneimage\":\"\",\"title\":\"\",\"type\":\"widget\",\"vendor\":\"LetsFIWARE\",\"version\":\"1.0.0\",\"wiring\":{\"inputs\":null,\"outputs\":null}}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"NGSIGO/ngsigo-widget/1.0.0\": {\n    \"type\": \"widget\",\n    \"vendor\": \"LetsFIWARE\",\n    \"version\": \"1.0.0\"\n  },\n  \"NGSIGO/ngsigo-operator/1.0.0\": {\n    \"type\": \"operator\",\n    \"vendor\": \"LetsFIWARE\",\n    \"version\": \"1.0.0\"\n  },\n  \"NGSIGO/ngsigo-mashup/1.0.0\": {\n    \"type\": \"mashup\",\n    \"vendor\": \"LetsFIWARE\",\n    \"version\": \"1.0.0\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListPretty2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--widget", "--pretty"})

	err := wireCloudResourcesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"NGSIGO/ngsigo-widget/1.0.0\": {\n    \"authors\": null,\n    \"changelog\": \"\",\n    \"contributors\": null,\n    \"default_lang\": \"\",\n    \"description\": \"\",\n    \"doc\": \"\",\n    \"email\": \"\",\n    \"homepage\": \"\",\n    \"image\": \"\",\n    \"issuetracker\": \"\",\n    \"js_files\": null,\n    \"license\": \"\",\n    \"licenseurl\": \"\",\n    \"longdescription\": \"\",\n    \"name\": \"\",\n    \"preferences\": null,\n    \"properties\": null,\n    \"requirements\": null,\n    \"smartphoneimage\": \"\",\n    \"title\": \"\",\n    \"type\": \"widget\",\n    \"vendor\": \"LetsFIWARE\",\n    \"version\": \"1.0.0\",\n    \"wiring\": {\n      \"inputs\": null,\n      \"outputs\": null\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourcesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourcesListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc"})

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourcesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourcesListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourcesListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty"})

	setJSONIndentError(ngsi)

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcResourcesListErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	setJSONDecodeErr(ngsi, 1)

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcResourcesListErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--widget", "--json"})

	setJSONEncodeErr(ngsi, 2)

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcResourcesListErrorPretty2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,widget,operator,mashup")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--widget", "--pretty"})

	setJSONIndentError(ngsi)

	err := wireCloudResourcesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcResourceGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--vender=NGSIGO", "--name=ngsigo-widget", "--version=1.0.0"})

	err := wireCloudResourceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"authors\":null,\"changelog\":\"\",\"contributors\":null,\"default_lang\":\"\",\"description\":\"\",\"doc\":\"\",\"email\":\"\",\"homepage\":\"\",\"image\":\"\",\"issuetracker\":\"\",\"js_files\":null,\"license\":\"\",\"licenseurl\":\"\",\"longdescription\":\"\",\"name\":\"NGSIGO\",\"preferences\":null,\"properties\":null,\"requirements\":null,\"smartphoneimage\":\"\",\"title\":\"\",\"type\":\"widget\",\"vendor\":\"LetsFIWARE\",\"version\":\"1.0.0\",\"wiring\":{\"inputs\":null,\"outputs\":null}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourceGetArg(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"authors\":null,\"changelog\":\"\",\"contributors\":null,\"default_lang\":\"\",\"description\":\"\",\"doc\":\"\",\"email\":\"\",\"homepage\":\"\",\"image\":\"\",\"issuetracker\":\"\",\"js_files\":null,\"license\":\"\",\"licenseurl\":\"\",\"longdescription\":\"\",\"name\":\"NGSIGO\",\"preferences\":null,\"properties\":null,\"requirements\":null,\"smartphoneimage\":\"\",\"title\":\"\",\"type\":\"widget\",\"vendor\":\"LetsFIWARE\",\"version\":\"1.0.0\",\"wiring\":{\"inputs\":null,\"outputs\":null}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourceGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"authors\": null,\n  \"changelog\": \"\",\n  \"contributors\": null,\n  \"default_lang\": \"\",\n  \"description\": \"\",\n  \"doc\": \"\",\n  \"email\": \"\",\n  \"homepage\": \"\",\n  \"image\": \"\",\n  \"issuetracker\": \"\",\n  \"js_files\": null,\n  \"license\": \"\",\n  \"licenseurl\": \"\",\n  \"longdescription\": \"\",\n  \"name\": \"NGSIGO\",\n  \"preferences\": null,\n  \"properties\": null,\n  \"requirements\": null,\n  \"smartphoneimage\": \"\",\n  \"title\": \"\",\n  \"type\": \"widget\",\n  \"vendor\": \"LetsFIWARE\",\n  \"version\": \"1.0.0\",\n  \"wiring\": {\n    \"inputs\": null,\n    \"outputs\": null\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourceGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceGetErrorArg(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "argument error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourceGetErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourceGetErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	setJSONDecodeErr(ngsi, 1)

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourceGetErrorNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/2.0.0"})

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "NGSIGO/ngsigo-widget/2.0.0 not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceGetErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	setJSONEncodeErr(ngsi, 2)

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcResourceGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty", "NGSIGO/ngsigo-widget/1.0.0"})

	setJSONIndentError(ngsi)

	err := wireCloudResourceGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcResourceDownload(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--vender=NGSIGO", "--name=ngsigo-widget", "--version=1.0.0"})

	err := wireCloudResourceDownload(c)

	assert.NoError(t, err)
}

func TestWcResourceDownloadArg(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceDownload(c)

	assert.NoError(t, err)
}

func TestWcResourceDownloadErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudResourceDownload(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceDownloadErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceDownload(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceDownloadErrorArg(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudResourceDownload(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "argument error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceDownloadErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceDownload(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourceDownloadErrorNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceDownload(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "NGSIGO/ngsigo-widget/1.0.0 not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceDownloadErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceDownload(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourceDownloadErrorWriteFile(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(``)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{WriteFileErr: errors.New("write file error")}

	setupFlagString(set, "host,vender,name,version")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceDownload(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "write file error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstall(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--file=NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	assert.NoError(t, err)
}

func TestWcResourceInstallJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.ResBody = []byte(`{"authors": [{"email": "wirecloud@letsfiware.jp", "name": "Let's FIWARE"}], "changelog": "doc/changelog.md", "contributors": [], "default_lang": "en", "description": "ol-ext poi", "doc": "https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/doc/userguide.md", "email": "kazuhito@fisuda.jp", "homepage": "https://github.com/lets-fiware/ol3-bubble-map-operator", "image": "", "issuetracker": "https://github.com/lets-fiware/ol3-bubble-map-operator/issues", "js_files": ["https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/js/main.js"], "license": "MIT", "licenseurl": "", "longdescription": "DESCRIPTION.md", "name": "test-widget", "preferences": [{"default": "radius", "description": "Name of attribute that specifies the radius", "label": "Radius attribute", "multiuser": false, "name": "radiusAttr", "readonly": false, "required": true, "secure": false, "type": "text", "value": null}, {"default": "", "description": "Name of attribute that specifies the text", "label": "Text attribute", "multiuser": false, "name": "textAttr", "readonly": false, "required": false, "secure": false, "type": "text", "value": null}], "properties": [], "requirements": [], "smartphoneimage": "", "title": "ol-ext poi", "type": "operator", "vendor": "NGSIGO", "version": "0.1.0", "wiring": {"inputs": [{"actionlabel": "", "description": "Received entities will be transform into PoIs", "friendcode": "entity", "label": "Entities", "name": "entityInput", "type": "text"}], "outputs": [{"description": "Transformed Points of Interests from the received entities", "friendcode": "poi", "label": "PoIs", "name": "poiOutput", "type": "text"}]}}`)
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--json", "--file=NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"authors\": [{\"email\": \"wirecloud@letsfiware.jp\", \"name\": \"Let's FIWARE\"}], \"changelog\": \"doc/changelog.md\", \"contributors\": [], \"default_lang\": \"en\", \"description\": \"ol-ext poi\", \"doc\": \"https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/doc/userguide.md\", \"email\": \"kazuhito@fisuda.jp\", \"homepage\": \"https://github.com/lets-fiware/ol3-bubble-map-operator\", \"image\": \"\", \"issuetracker\": \"https://github.com/lets-fiware/ol3-bubble-map-operator/issues\", \"js_files\": [\"https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/js/main.js\"], \"license\": \"MIT\", \"licenseurl\": \"\", \"longdescription\": \"DESCRIPTION.md\", \"name\": \"test-widget\", \"preferences\": [{\"default\": \"radius\", \"description\": \"Name of attribute that specifies the radius\", \"label\": \"Radius attribute\", \"multiuser\": false, \"name\": \"radiusAttr\", \"readonly\": false, \"required\": true, \"secure\": false, \"type\": \"text\", \"value\": null}, {\"default\": \"\", \"description\": \"Name of attribute that specifies the text\", \"label\": \"Text attribute\", \"multiuser\": false, \"name\": \"textAttr\", \"readonly\": false, \"required\": false, \"secure\": false, \"type\": \"text\", \"value\": null}], \"properties\": [], \"requirements\": [], \"smartphoneimage\": \"\", \"title\": \"ol-ext poi\", \"type\": \"operator\", \"vendor\": \"NGSIGO\", \"version\": \"0.1.0\", \"wiring\": {\"inputs\": [{\"actionlabel\": \"\", \"description\": \"Received entities will be transform into PoIs\", \"friendcode\": \"entity\", \"label\": \"Entities\", \"name\": \"entityInput\", \"type\": \"text\"}], \"outputs\": [{\"description\": \"Transformed Points of Interests from the received entities\", \"friendcode\": \"poi\", \"label\": \"PoIs\", \"name\": \"poiOutput\", \"type\": \"text\"}]}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourceInstallPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.ResBody = []byte(`{"authors": [{"email": "wirecloud@letsfiware.jp", "name": "Let's FIWARE"}], "changelog": "doc/changelog.md", "contributors": [], "default_lang": "en", "description": "ol-ext poi", "doc": "https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/doc/userguide.md", "email": "kazuhito@fisuda.jp", "homepage": "https://github.com/lets-fiware/ol3-bubble-map-operator", "image": "", "issuetracker": "https://github.com/lets-fiware/ol3-bubble-map-operator/issues", "js_files": ["https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/js/main.js"], "license": "MIT", "licenseurl": "", "longdescription": "DESCRIPTION.md", "name": "test-widget", "preferences": [{"default": "radius", "description": "Name of attribute that specifies the radius", "label": "Radius attribute", "multiuser": false, "name": "radiusAttr", "readonly": false, "required": true, "secure": false, "type": "text", "value": null}, {"default": "", "description": "Name of attribute that specifies the text", "label": "Text attribute", "multiuser": false, "name": "textAttr", "readonly": false, "required": false, "secure": false, "type": "text", "value": null}], "properties": [], "requirements": [], "smartphoneimage": "", "title": "ol-ext poi", "type": "operator", "vendor": "NGSIGO", "version": "0.1.0", "wiring": {"inputs": [{"actionlabel": "", "description": "Received entities will be transform into PoIs", "friendcode": "entity", "label": "Entities", "name": "entityInput", "type": "text"}], "outputs": [{"description": "Transformed Points of Interests from the received entities", "friendcode": "poi", "label": "PoIs", "name": "poiOutput", "type": "text"}]}}`)
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty", "--file=NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"authors\": [\n    {\n      \"email\": \"wirecloud@letsfiware.jp\",\n      \"name\": \"Let's FIWARE\"\n    }\n  ],\n  \"changelog\": \"doc/changelog.md\",\n  \"contributors\": [],\n  \"default_lang\": \"en\",\n  \"description\": \"ol-ext poi\",\n  \"doc\": \"https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/doc/userguide.md\",\n  \"email\": \"kazuhito@fisuda.jp\",\n  \"homepage\": \"https://github.com/lets-fiware/ol3-bubble-map-operator\",\n  \"image\": \"\",\n  \"issuetracker\": \"https://github.com/lets-fiware/ol3-bubble-map-operator/issues\",\n  \"js_files\": [\n    \"https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/js/main.js\"\n  ],\n  \"license\": \"MIT\",\n  \"licenseurl\": \"\",\n  \"longdescription\": \"DESCRIPTION.md\",\n  \"name\": \"test-widget\",\n  \"preferences\": [\n    {\n      \"default\": \"radius\",\n      \"description\": \"Name of attribute that specifies the radius\",\n      \"label\": \"Radius attribute\",\n      \"multiuser\": false,\n      \"name\": \"radiusAttr\",\n      \"readonly\": false,\n      \"required\": true,\n      \"secure\": false,\n      \"type\": \"text\",\n      \"value\": null\n    },\n    {\n      \"default\": \"\",\n      \"description\": \"Name of attribute that specifies the text\",\n      \"label\": \"Text attribute\",\n      \"multiuser\": false,\n      \"name\": \"textAttr\",\n      \"readonly\": false,\n      \"required\": false,\n      \"secure\": false,\n      \"type\": \"text\",\n      \"value\": null\n    }\n  ],\n  \"properties\": [],\n  \"requirements\": [],\n  \"smartphoneimage\": \"\",\n  \"title\": \"ol-ext poi\",\n  \"type\": \"operator\",\n  \"vendor\": \"NGSIGO\",\n  \"version\": \"0.1.0\",\n  \"wiring\": {\n    \"inputs\": [\n      {\n        \"actionlabel\": \"\",\n        \"description\": \"Received entities will be transform into PoIs\",\n        \"friendcode\": \"entity\",\n        \"label\": \"Entities\",\n        \"name\": \"entityInput\",\n        \"type\": \"text\"\n      }\n    ],\n    \"outputs\": [\n      {\n        \"description\": \"Transformed Points of Interests from the received entities\",\n        \"friendcode\": \"poi\",\n        \"label\": \"PoIs\",\n        \"name\": \"poiOutput\",\n        \"type\": \"text\"\n      }\n    ]\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourceInstallArg(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	assert.NoError(t, err)
}

func TestWcResourceInstallArgPublic(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--public", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	assert.NoError(t, err)
}

func TestWcResourceInstallErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"NGSIGO/ngsigo-widget/1.0.0":{"type":"widget","vendor":"LetsFIWARE","name":"NGSIGO","version":"1.0.0"},"NGSIGO/ngsigo-operator/1.0.0":{"type":"operator","vendor":"LetsFIWARE","version":"1.0.0"},"NGSIGO/ngsigo-mashup/1.0.0":{"type":"mashup","vendor":"LetsFIWARE","version":"1.0.0"}}`)
	reqRes.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorArg(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "argument error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorFilePathAbs(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.FilePath = &MockFilePathLib{PathAbsErr: errors.New("file path abs error")}
	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "file path abs error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorReadFile(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileErr: errors.New("read file error")}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "read file error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorGetMacName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.ZipLib = &MockZipLib{Zip: errors.New("newreader error")}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "newreader error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorExistsMac(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorOverWrite(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--overwrite", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorAlreadyExists(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "NGSIGO/ngsi-go-widget/1.0.0 already exists", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorMultiPart(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resource"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}
	ngsi.MultiPart = &MockMultiPart{CreatePartErr: errors.New("createpart error")}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "createpart error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resource"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.Path = "/api/resources"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 12, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceInstallErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusNotFound
	reqRes1.Path = "/api/resource/NGSIGO/ngsi-go-widget/1.0.0"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/api/resources"
	reqRes2.ResBody = []byte(`{"authors": [{"email": "wirecloud@letsfiware.jp", "name": "Let's FIWARE"}], "changelog": "doc/changelog.md", "contributors": [], "default_lang": "en", "description": "ol-ext poi", "doc": "https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/doc/userguide.md", "email": "kazuhito@fisuda.jp", "homepage": "https://github.com/lets-fiware/ol3-bubble-map-operator", "image": "", "issuetracker": "https://github.com/lets-fiware/ol3-bubble-map-operator/issues", "js_files": ["https://mashup.lab.e-suda.info/showcase/media/NGSIGO/test-widget/0.1.0/js/main.js"], "license": "MIT", "licenseurl": "", "longdescription": "DESCRIPTION.md", "name": "test-widget", "preferences": [{"default": "radius", "description": "Name of attribute that specifies the radius", "label": "Radius attribute", "multiuser": false, "name": "radiusAttr", "readonly": false, "required": true, "secure": false, "type": "text", "value": null}, {"default": "", "description": "Name of attribute that specifies the text", "label": "Text attribute", "multiuser": false, "name": "textAttr", "readonly": false, "required": false, "secure": false, "type": "text", "value": null}], "properties": [], "requirements": [], "smartphoneimage": "", "title": "ol-ext poi", "type": "operator", "vendor": "NGSIGO", "version": "0.1.0", "wiring": {"inputs": [{"actionlabel": "", "description": "Received entities will be transform into PoIs", "friendcode": "entity", "label": "Entities", "name": "entityInput", "type": "text"}], "outputs": [{"description": "Transformed Points of Interests from the received entities", "friendcode": "poi", "label": "PoIs", "name": "poiOutput", "type": "text"}]}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	ngsi.Ioutil = &MockIoutilLib{ReadFileData: widgetZipFile}

	setupFlagString(set, "host,file")
	setupFlagBool(set, "json,pretty,public,overwrite")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--pretty", "NGSIGO_ngsigo-widget_1.0.0.wgt"})

	setJSONIndentError(ngsi)

	err := wireCloudResourceInstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 13, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestMakeMultipart(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	var body bytes.Buffer
	mw := &MockMultiPartLib{Mw: multipart.NewWriter(&body)}

	_, err := makeMultipart(ngsi, mw, "ngsigo.wgt", []byte(""))

	assert.NoError(t, err)
}

func TestMakeMultipartErrorCreatePart(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	var body bytes.Buffer
	mw := &MockMultiPartLib{Mw: multipart.NewWriter(&body), CreatePartErr: errors.New("createpart error")}

	_, err := makeMultipart(ngsi, mw, "ngsigo.wgt", []byte(""))

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "createpart error", ngsiErr.Message)
	}
}

func TestMakeMultipartErrorCopy(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	var body bytes.Buffer
	mw := &MockMultiPartLib{Mw: multipart.NewWriter(&body)}
	ngsi.Ioutil = &MockIoutilLib{CopyErr: errors.New("io.Copy error")}

	_, err := makeMultipart(ngsi, mw, "ngsigo.wgt", []byte(""))

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "io.Copy error", ngsiErr.Message)
	}
}

func TestMakeMultipartErrorClose(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	var body bytes.Buffer
	mw := &MockMultiPartLib{Mw: multipart.NewWriter(&body), CloseErr: errors.New("close error")}

	_, err := makeMultipart(ngsi, mw, "ngsigo.wgt", []byte(""))

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "close error", ngsiErr.Message)
	}
}

func TestWcResourceUninstall(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--vender=NGSIGO", "--name=ngsigo-widget", "--version=1.0.0", "--run"})

	err := wireCloudResourceUninstall(c)

	assert.NoError(t, err)
}

func TestWcResourceUninstallJSON(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"affectedVersions": ["0.1.0"]}`)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--vender=NGSIGO", "--name=ngsigo-widget", "--version=1.0.0", "--json", "--run"})

	err := wireCloudResourceUninstall(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"affectedVersions\": [\"0.1.0\"]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourceUninstallPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"affectedVersions": ["0.1.0", "0.2.0"]}`)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--vender=NGSIGO", "--name=ngsigo-widget", "--pretty", "--run"})

	err := wireCloudResourceUninstall(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"affectedVersions\": [\n    \"0.1.0\",\n    \"0.2.0\"\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestWcResourceUninstallArg(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--run", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceUninstall(c)

	assert.NoError(t, err)
}

func TestWcResourceUninstallErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := wireCloudResourceUninstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceUninstallErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wc", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceUninstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: wc", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceUninstallErrorArg(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud"})

	err := wireCloudResourceUninstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "argument error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceUninstallErrorRun(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceUninstall(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "NGSIGO/ngsigo-widget/1.0.0 will be uninstalled. run uninstall with --run option\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestWcResourceUninstallErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--run", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceUninstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourceUninstallErrorNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--run", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceUninstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "NGSIGO/ngsigo-widget/1.0.0 not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcResourceUninstallErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--run", "NGSIGO/ngsigo-widget/1.0.0"})

	err := wireCloudResourceUninstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcResourceUninstallErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"affectedVersions": ["0.1.0"]}`)
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,vender,name,version")
	setupFlagBool(set, "json,pretty,run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=wirecloud", "--run", "--pretty", "NGSIGO/ngsigo-widget/1.0.0"})

	setJSONIndentError(ngsi)

	err := wireCloudResourceUninstall(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestWcUninstallMac(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=wirecloud"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	err = uninstallMac(ngsi, client, "NGSIGO/ngsigo-widget/1.0.0")

	assert.NoError(t, err)
}

func TestWcUninstallMacErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=wirecloud"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	err = uninstallMac(ngsi, client, "NGSIGO/ngsigo-widget/1.0.0")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestWcUninstallMacErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=wirecloud"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	err = uninstallMac(ngsi, client, "NGSIGO/ngsigo-widget/1.0.0")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

var widgetZipFile = []byte{
	0x50, 0x4b, 0x03, 0x04, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0x2a, 0x39, 0xde, 0x52, 0x0e, 0x23,
	0x99, 0xb2, 0x8c, 0x00, 0x00, 0x00, 0xa4, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x1c, 0x00, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x78, 0x6d, 0x6c, 0x55, 0x54, 0x09, 0x00, 0x03, 0x10, 0x9a, 0xdb,
	0x60, 0x10, 0x9a, 0xdb, 0x60, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xe8, 0x03, 0x00, 0x00, 0x04,
	0xe8, 0x03, 0x00, 0x00, 0x3d, 0x8e, 0xb1, 0x0e, 0xc2, 0x20, 0x14, 0x00, 0xf7, 0x7e, 0x05, 0x79,
	0x0b, 0x53, 0xa1, 0xdd, 0x8c, 0x29, 0xed, 0x66, 0xe3, 0xa2, 0x83, 0xfa, 0x01, 0x0d, 0x3c, 0x91,
	0xa4, 0x3c, 0x08, 0x50, 0xeb, 0xe7, 0x4b, 0x62, 0xe2, 0x7e, 0xb9, 0xbb, 0x61, 0xfa, 0xf8, 0x95,
	0xbd, 0x31, 0x65, 0x17, 0x48, 0xf1, 0x5e, 0x74, 0x9c, 0x21, 0xe9, 0x60, 0x1c, 0x59, 0xc5, 0x1f,
	0xf7, 0x53, 0x7b, 0xe0, 0xd3, 0xd8, 0x0c, 0xbb, 0x33, 0x16, 0x0b, 0xab, 0x30, 0x65, 0x05, 0xaf,
	0x52, 0xe2, 0x51, 0xca, 0xdd, 0x25, 0xd4, 0x6b, 0xd8, 0x8c, 0xd0, 0x81, 0x76, 0x2c, 0xe2, 0xe9,
	0xc4, 0x16, 0xbd, 0xc0, 0x2c, 0x29, 0x4b, 0xbf, 0x68, 0x83, 0x59, 0x27, 0x17, 0x4b, 0x75, 0xcb,
	0x1e, 0x6a, 0x86, 0x4c, 0x48, 0x0a, 0x2e, 0xf3, 0xed, 0x3c, 0x5f, 0x81, 0xd1, 0xe2, 0x51, 0x01,
	0xd9, 0xec, 0x5a, 0x1b, 0xda, 0x5f, 0x02, 0xfe, 0x33, 0x50, 0x67, 0x44, 0x07, 0x63, 0xf3, 0x05,
	0x50, 0x4b, 0x01, 0x02, 0x1e, 0x03, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00, 0x2a, 0x39, 0xde, 0x52,
	0x0e, 0x23, 0x99, 0xb2, 0x8c, 0x00, 0x00, 0x00, 0xa4, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x18, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa4, 0x81, 0x00, 0x00, 0x00, 0x00, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x78, 0x6d, 0x6c, 0x55, 0x54, 0x05, 0x00, 0x03, 0x10, 0x9a, 0xdb,
	0x60, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xe8, 0x03, 0x00, 0x00, 0x04, 0xe8, 0x03, 0x00, 0x00,
	0x50, 0x4b, 0x05, 0x06, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x50, 0x00, 0x00, 0x00,
	0xd0, 0x00, 0x00, 0x00, 0x00, 0x00,
}

func TestGetMacName(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	mashup, name, err := getMacName(ngsi, widgetZipFile)

	if assert.NoError(t, err) {
		assert.Equal(t, mashup, "widget")
		assert.Equal(t, name, "NGSIGO/ngsi-go-widget/1.0.0")
	}
}

func TestGetMacNameErrorNewReader(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	ngsi.ZipLib = &MockZipLib{Zip: errors.New("newreader error")}

	_, _, err := getMacName(ngsi, nil)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "newreader error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetMacNameErrorNotFound(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	zf := zip.File{}
	zf.FileHeader.Name = "config.json"
	zr := zip.Reader{File: []*zip.File{&zf}}
	ngsi.ZipLib = &MockZipLib{ZipReader: &zr}

	_, _, err := getMacName(ngsi, nil)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "config.xml not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetFromConfigXML(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	data := []byte(`<widget xmlns="http://wirecloud.conwet.fi.upm.es/ns/macdescription/1" vendor="NGSIGO" name="ngsi-go-widget" version="1.0.0">`)
	ngsi.Ioutil = &MockIoutilLib{ReadFullData: data}

	mashup, name, err := getFromConfigXML(ngsi, nil, uint32(len(data)))

	if assert.NoError(t, err) {
		assert.Equal(t, mashup, "widget")
		assert.Equal(t, name, "NGSIGO/ngsi-go-widget/1.0.0")
	} else {
		t.FailNow()
	}
}

func TestGetFromConfigXMLErrorReadFull(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	ngsi.Ioutil = &MockIoutilLib{ReadFullErr: errors.New("readfull error")}

	_, _, err := getFromConfigXML(ngsi, nil, 10)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "readfull error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetFromConfigXMLErrorConfigXML(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	data := []byte(`<widget xmlns="http://wirecloud.conwet.fi.upm.es/ns/macdescription/1" vendor="NGSIGO" name="ngsi-go-widget" version="1.0.0"`)
	ngsi.Ioutil = &MockIoutilLib{ReadFullData: data}

	_, _, err := getFromConfigXML(ngsi, nil, uint32(len(data)))

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "config.xml error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetFromConfigXMLErrorConfigXML3(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	data := []byte(`<?xml version='1.0' encoding='UTF-8'?>`)
	ngsi.Ioutil = &MockIoutilLib{ReadFullData: data}

	_, _, err := getFromConfigXML(ngsi, nil, uint32(len(data)))

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "config.xml error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestWcExistsMacFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=wirecloud"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	actual, err := existsMac(ngsi, client, "NGSIGO/ngsigo-widget/1.0.0")

	if assert.NoError(t, err) {
		assert.Equal(t, true, actual)
	}
}

func TestWcExistsMacNotFound(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=wirecloud"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	actual, err := existsMac(ngsi, client, "NGSIGO/ngsigo-widget/1.0.0")

	if assert.NoError(t, err) {
		assert.Equal(t, false, actual)
	}
}

func TestWcExistsMacErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNotFound
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=wirecloud"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	actual, err := existsMac(ngsi, client, "NGSIGO/ngsigo-widget/1.0.0")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, false, actual)
	} else {
		t.FailNow()
	}
}

func TestWcExistsMacErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/api/resource/NGSIGO/ngsigo-widget/1.0.0"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=wirecloud"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, nil)
	assert.NoError(t, err)

	actual, err := existsMac(ngsi, client, "NGSIGO/ngsigo-widget/1.0.0")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, false, actual)
	} else {
		t.FailNow()
	}
}
