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

package cygnus

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestNamemappingsList(t *testing.T) {
	c := setupTest([]string{"namemappings", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[]}}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsListPretty(t *testing.T) {
	c := setupTest([]string{"namemappings", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": []\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"namemappings", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestNamemappingsListErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"namemappings", "list", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestNamemappingsListErrorPretty(t *testing.T) {
	c := setupTest([]string{"namemappings", "list", "--host", "cygnus", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[]}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := namemappingsList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestNamemappingsCreate(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[{\"originalService\":\"^(.*)\",\"newService\":\"null\",\"servicePathMappings\":[{\"originalServicePath\":\"/myservicepath1\",\"newServicePath\":\"/new_myservicepath1\",\"entityMappings\":[{\"originalEntityId\":\"myentityid1\",\"originalEntityType\":\"myentitytype1\",\"newEntityId\":\"new_myentityid1\",\"newEntityType\":\"new_myentitytype1\",\"attributeMappings\":[{\"originalAttributeName\":\"myattributename1\",\"originalAttributeType\":\"myattributetype1\",\"newAttributeName\":\"new_myattributename1\",\"newAttributeType\":\"new_myattributetype1\"}]}]}]}]}}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsCreatePretty(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": [\n      {\n        \"originalService\": \"^(.*)\",\n        \"newService\": \"null\",\n        \"servicePathMappings\": [\n          {\n            \"originalServicePath\": \"/myservicepath1\",\n            \"newServicePath\": \"/new_myservicepath1\",\n            \"entityMappings\": [\n              {\n                \"originalEntityId\": \"myentityid1\",\n                \"originalEntityType\": \"myentitytype1\",\n                \"newEntityId\": \"new_myentityid1\",\n                \"newEntityType\": \"new_myentitytype1\",\n                \"attributeMappings\": [\n                  {\n                    \"originalAttributeName\": \"myattributename1\",\n                    \"originalAttributeType\": \"myattributetype1\",\n                    \"newAttributeName\": \"new_myattributename1\",\n                    \"newAttributeType\": \"new_myattributetype1\"\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ]\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsCreateErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"namemappings", "create", "--host", "cygnus", "--data", "@"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestNamemappingsCreateErrorHTTP(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestNamemappingsCreateErrorStatusCode(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "create", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestNamemappingsCreateErrorPretty(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "create", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := namemappingsCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestNamemappingsUpdate(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "update", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[{\"originalService\":\"^(.*)\",\"newService\":\"null\",\"servicePathMappings\":[{\"originalServicePath\":\"/myservicepath1\",\"newServicePath\":\"/new_myservicepath1\",\"entityMappings\":[{\"originalEntityId\":\"myentityid1\",\"originalEntityType\":\"myentitytype1\",\"newEntityId\":\"new_myentityid1\",\"newEntityType\":\"new_myentitytype1\",\"attributeMappings\":[{\"originalAttributeName\":\"myattributename1\",\"originalAttributeType\":\"myattributetype1\",\"newAttributeName\":\"new_myattributename1\",\"newAttributeType\":\"new_myattributetype1\"}]}]}]}]}}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsUpdatePretty(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "update", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": [\n      {\n        \"originalService\": \"^(.*)\",\n        \"newService\": \"null\",\n        \"servicePathMappings\": [\n          {\n            \"originalServicePath\": \"/myservicepath1\",\n            \"newServicePath\": \"/new_myservicepath1\",\n            \"entityMappings\": [\n              {\n                \"originalEntityId\": \"myentityid1\",\n                \"originalEntityType\": \"myentitytype1\",\n                \"newEntityId\": \"new_myentityid1\",\n                \"newEntityType\": \"new_myentitytype1\",\n                \"attributeMappings\": [\n                  {\n                    \"originalAttributeName\": \"myattributename1\",\n                    \"originalAttributeType\": \"myattributetype1\",\n                    \"newAttributeName\": \"new_myattributename1\",\n                    \"newAttributeType\": \"new_myattributetype1\"\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ]\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"namemappings", "update", "--host", "cygnus", "--data", "@"})

	err := namemappingsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestNamemappingsUpdateErrorHTTP(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "update", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestNamemappingsUpdateErrorStatusCode(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "update", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestNamemappingsUpdateErrorPretty(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "update", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := namemappingsUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestNamemappingsDelete(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "delete", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"success\":\"true\",\"result\":{\"serviceMappings\":[{\"originalService\":\"^(.*)\",\"newService\":\"null\",\"servicePathMappings\":[{\"originalServicePath\":\"/myservicepath1\",\"newServicePath\":\"/new_myservicepath1\",\"entityMappings\":[{\"originalEntityId\":\"myentityid1\",\"originalEntityType\":\"myentitytype1\",\"newEntityId\":\"new_myentityid1\",\"newEntityType\":\"new_myentitytype1\",\"attributeMappings\":[{\"originalAttributeName\":\"myattributename1\",\"originalAttributeType\":\"myattributetype1\",\"newAttributeName\":\"new_myattributename1\",\"newAttributeType\":\"new_myattributetype1\"}]}]}]}]}}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsDeletePretty(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "delete", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsDelete(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"success\": \"true\",\n  \"result\": {\n    \"serviceMappings\": [\n      {\n        \"originalService\": \"^(.*)\",\n        \"newService\": \"null\",\n        \"servicePathMappings\": [\n          {\n            \"originalServicePath\": \"/myservicepath1\",\n            \"newServicePath\": \"/new_myservicepath1\",\n            \"entityMappings\": [\n              {\n                \"originalEntityId\": \"myentityid1\",\n                \"originalEntityType\": \"myentitytype1\",\n                \"newEntityId\": \"new_myentityid1\",\n                \"newEntityType\": \"new_myentitytype1\",\n                \"attributeMappings\": [\n                  {\n                    \"originalAttributeName\": \"myattributename1\",\n                    \"originalAttributeType\": \"myattributetype1\",\n                    \"newAttributeName\": \"new_myattributename1\",\n                    \"newAttributeType\": \"new_myattributetype1\"\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ]\n  }\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestNamemappingsDeleteErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"namemappings", "delete", "--host", "cygnus", "--data", "@"})

	err := namemappingsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestNamemappingsDeleteErrorHTTP(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "delete", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)
	reqRes.Err = errors.New("error")

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestNamemappingsDeleteErrorStatusCode(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "delete", "--host", "cygnus", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := namemappingsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestNamemappingsDeleteErrorPretty(t *testing.T) {
	data := `{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`
	c := setupTest([]string{"namemappings", "delete", "--host", "cygnus", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/namemappings"
	reqRes.ReqData = []byte(`{"serviceMappings":[{"servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}`)
	reqRes.ResBody = []byte(`{"success":"true","result":{"serviceMappings":[{"originalService":"^(.*)","newService":"null","servicePathMappings":[{"originalServicePath":"/myservicepath1","newServicePath":"/new_myservicepath1","entityMappings":[{"originalEntityId":"myentityid1","originalEntityType":"myentitytype1","newEntityId":"new_myentityid1","newEntityType":"new_myentitytype1","attributeMappings":[{"originalAttributeName":"myattributename1","originalAttributeType":"myattributetype1","newAttributeName":"new_myattributename1","newAttributeType":"new_myattributetype1"}]}]}]}]}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := namemappingsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
