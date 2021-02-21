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
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNamemappingsList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := namemappingsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[]}}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	err := namemappingsList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": []\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := namemappingsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := namemappingsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := namemappingsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestNamemappingsListErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := namemappingsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[]}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty"})

	setJSONIndentError(ngsi)

	err := namemappingsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestNamemappingsCreate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[{\"originalService\":\"^(.*)\",\"newService\":\"null\",\"servicePathMappings\":[{\"originalServicePath\":\"/myservicepath1\",\"newServicePath\":\"/new_myservicepath1\",\"entityMappings\":[{\"originalEntityId\":\"myentityid1\",\"originalEntityType\":\"myentitytype1\",\"newEntityId\":\"new_myentityid1\",\"newEntityType\":\"new_myentitytype1\",\"attributeMappings\":[{\"originalAttributeName\":\"myattributename1\",\"originalAttributeType\":\"myattributetype1\",\"newAttributeName\":\"new_myattributename1\",\"newAttributeType\":\"new_myattributetype1\"}]}]}]}]}}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": [\n      {\n        \"originalService\": \"^(.*)\",\n        \"newService\": \"null\",\n        \"servicePathMappings\": [\n          {\n            \"originalServicePath\": \"/myservicepath1\",\n            \"newServicePath\": \"/new_myservicepath1\",\n            \"entityMappings\": [\n              {\n                \"originalEntityId\": \"myentityid1\",\n                \"originalEntityType\": \"myentitytype1\",\n                \"newEntityId\": \"new_myentityid1\",\n                \"newEntityType\": \"new_myentitytype1\",\n                \"attributeMappings\": [\n                  {\n                    \"originalAttributeName\": \"myattributename1\",\n                    \"originalAttributeType\": \"myattributetype1\",\n                    \"newAttributeName\": \"new_myattributename1\",\n                    \"newAttributeType\": \"new_myattributetype1\"\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ]\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := namemappingsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := namemappingsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	setJSONEncodeErr(ngsi, 2)

	err := namemappingsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsCreateErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--data="})

	err := namemappingsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsCreateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	setJSONIndentError(ngsi)

	err := namemappingsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsUpdate(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[{\"originalService\":\"^(.*)\",\"newService\":\"null\",\"servicePathMappings\":[{\"originalServicePath\":\"/myservicepath1\",\"newServicePath\":\"/new_myservicepath1\",\"entityMappings\":[{\"originalEntityId\":\"myentityid1\",\"originalEntityType\":\"myentitytype1\",\"newEntityId\":\"new_myentityid1\",\"newEntityType\":\"new_myentitytype1\",\"attributeMappings\":[{\"originalAttributeName\":\"myattributename1\",\"originalAttributeType\":\"myattributetype1\",\"newAttributeName\":\"new_myattributename1\",\"newAttributeType\":\"new_myattributetype1\"}]}]}]}]}}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsUpdatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsUpdate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": [\n      {\n        \"originalService\": \"^(.*)\",\n        \"newService\": \"null\",\n        \"servicePathMappings\": [\n          {\n            \"originalServicePath\": \"/myservicepath1\",\n            \"newServicePath\": \"/new_myservicepath1\",\n            \"entityMappings\": [\n              {\n                \"originalEntityId\": \"myentityid1\",\n                \"originalEntityType\": \"myentitytype1\",\n                \"newEntityId\": \"new_myentityid1\",\n                \"newEntityType\": \"new_myentitytype1\",\n                \"attributeMappings\": [\n                  {\n                    \"originalAttributeName\": \"myattributename1\",\n                    \"originalAttributeType\": \"myattributetype1\",\n                    \"newAttributeName\": \"new_myattributename1\",\n                    \"newAttributeType\": \"new_myattributetype1\"\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ]\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := namemappingsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := namemappingsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := namemappingsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--data="})

	err := namemappingsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsUpdateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsUpdateErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsUpdateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	setJSONIndentError(ngsi)

	err := namemappingsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDelete(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[{\"originalService\":\"^(.*)\",\"newService\":\"null\",\"servicePathMappings\":[{\"originalServicePath\":\"/myservicepath1\",\"newServicePath\":\"/new_myservicepath1\",\"entityMappings\":[{\"originalEntityId\":\"myentityid1\",\"originalEntityType\":\"myentitytype1\",\"newEntityId\":\"new_myentityid1\",\"newEntityType\":\"new_myentitytype1\",\"attributeMappings\":[{\"originalAttributeName\":\"myattributename1\",\"originalAttributeType\":\"myattributetype1\",\"newAttributeName\":\"new_myattributename1\",\"newAttributeType\":\"new_myattributetype1\"}]}]}]}]}}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeletePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsDelete(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": [\n      {\n        \"originalService\": \"^(.*)\",\n        \"newService\": \"null\",\n        \"servicePathMappings\": [\n          {\n            \"originalServicePath\": \"/myservicepath1\",\n            \"newServicePath\": \"/new_myservicepath1\",\n            \"entityMappings\": [\n              {\n                \"originalEntityId\": \"myentityid1\",\n                \"originalEntityType\": \"myentitytype1\",\n                \"newEntityId\": \"new_myentityid1\",\n                \"newEntityType\": \"new_myentitytype1\",\n                \"attributeMappings\": [\n                  {\n                    \"originalAttributeName\": \"myattributename1\",\n                    \"originalAttributeType\": \"myattributetype1\",\n                    \"newAttributeName\": \"new_myattributename1\",\n                    \"newAttributeType\": \"new_myattributetype1\"\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ]\n  }\n}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := namemappingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeleteErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := namemappingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeleteErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus"})

	err := namemappingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify data", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeleteErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--data="})

	err := namemappingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	reqRes.Err = errors.New("error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeleteErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	err := namemappingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestNamemappingsDeleteErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	mock := NewMockHTTP()
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=cygnus", "--pretty", `--data={"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`})

	setJSONIndentError(ngsi)

	err := namemappingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
