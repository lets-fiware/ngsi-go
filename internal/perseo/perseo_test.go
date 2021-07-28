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

package perseo

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestPerseoRulesList(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "blood_rule_update\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesListPretty(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"_id\": \"6024cb208e2bfc0012c77488\",\n    \"name\": \"blood_rule_update\",\n    \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n    \"action\": {\n      \"parameters\": {\n        \"attributes\": [\n          {\n            \"name\": \"abnormal\",\n            \"type\": \"boolean\",\n            \"value\": \"true\"\n          }\n        ]\n      },\n      \"type\": \"update\"\n    },\n    \"subservice\": \"/\",\n    \"service\": \"unknownt\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesListPrettyRaw(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--raw", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"error\": null,\n  \"data\": [\n    {\n      \"_id\": \"6024cb208e2bfc0012c77488\",\n      \"name\": \"blood_rule_update\",\n      \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n      \"action\": {\n        \"type\": \"update\",\n        \"parameters\": {\n          \"attributes\": [\n            {\n              \"name\": \"abnormal\",\n              \"value\": \"true\",\n              \"type\": \"boolean\"\n            }\n          ]\n        }\n      },\n      \"subservice\": \"/\",\n      \"service\": \"unknownt\"\n    }\n  ],\n  \"count\": 1\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesListErrorHTTP(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/iot/device"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestPerseoRulesListErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoRulesListErrorPretty(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := perseoRulesList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoRulesGet(t *testing.T) {
	c := setupTest([]string{"rules", "get", "--host", "perseo", "--name", "blood_rule_update"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"error\":null,\"data\":{\"_id\":\"6024cb208e2bfc0012c77488\",\"name\":\"blood_rule_update\",\"text\":\"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"action\":{\"type\":\"update\",\"parameters\":{\"attributes\":[{\"name\":\"abnormal\",\"value\":\"true\",\"type\":\"boolean\"}]}},\"subservice\":\"/\",\"service\":\"unknownt\"}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesGetPretty(t *testing.T) {
	c := setupTest([]string{"rules", "get", "--host", "perseo", "--name", "blood_rule_update", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"error\": null,\n  \"data\": {\n    \"_id\": \"6024cb208e2bfc0012c77488\",\n    \"name\": \"blood_rule_update\",\n    \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n    \"action\": {\n      \"type\": \"update\",\n      \"parameters\": {\n        \"attributes\": [\n          {\n            \"name\": \"abnormal\",\n            \"value\": \"true\",\n            \"type\": \"boolean\"\n          }\n        ]\n      }\n    },\n    \"subservice\": \"/\",\n    \"service\": \"unknownt\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesGetErrorHTTP(t *testing.T) {
	c := setupTest([]string{"rules", "get", "--host", "perseo", "--name", "blood_rule_update"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestPerseoRulesGetErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"rules", "get", "--host", "perseo", "--name", "blood_rule_update"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoRulesGetErrorPretty(t *testing.T) {
	c := setupTest([]string{"rules", "get", "--host", "perseo", "--name", "blood_rule_update", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := perseoRulesGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoRulesCreate(t *testing.T) {
	data := `{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`
	c := setupTest([]string{"rules", "create", "--host", "perseo", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesCreate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestPerseoRulesCreateVerbose(t *testing.T) {
	data := `{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`
	c := setupTest([]string{"rules", "create", "--host", "perseo", "--data", data, "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[false,[{"code":200,"body":{"name":"ctxt$unknownt$","timeLastStateChange":1613024032325,"text":"create context ctxt$unknownt$ partition by service from iotEvent(service=\"unknownt\" and subservice=\"/\")","state":"STARTED"}},{"code":200,"body":{"name":"blood_rule_update@unknownt/","timeLastStateChange":1613028692624,"text":"context ctxt$unknownt$ select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","state":"STARTED"}},null],{"n":1,"ok":1},null]}`)
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"error\":null,\"data\":[false,[{\"code\":200,\"body\":{\"name\":\"ctxt$unknownt$\",\"timeLastStateChange\":1613024032325,\"text\":\"create context ctxt$unknownt$ partition by service from iotEvent(service=\\\"unknownt\\\" and subservice=\\\"/\\\")\",\"state\":\"STARTED\"}},{\"code\":200,\"body\":{\"name\":\"blood_rule_update@unknownt/\",\"timeLastStateChange\":1613028692624,\"text\":\"context ctxt$unknownt$ select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"state\":\"STARTED\"}},null],{\"n\":1,\"ok\":1},null]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesCreatePretty(t *testing.T) {
	data := `{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`
	c := setupTest([]string{"rules", "create", "--host", "perseo", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[false,[{"code":200,"body":{"name":"ctxt$unknownt$","timeLastStateChange":1613024032325,"text":"create context ctxt$unknownt$ partition by service from iotEvent(service=\"unknownt\" and subservice=\"/\")","state":"STARTED"}},{"code":200,"body":{"name":"blood_rule_update@unknownt/","timeLastStateChange":1613028692624,"text":"context ctxt$unknownt$ select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","state":"STARTED"}},null],{"n":1,"ok":1},null]}`)
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesCreate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"error\": null,\n  \"data\": [\n    false,\n    [\n      {\n        \"code\": 200,\n        \"body\": {\n          \"name\": \"ctxt$unknownt$\",\n          \"timeLastStateChange\": 1613024032325,\n          \"text\": \"create context ctxt$unknownt$ partition by service from iotEvent(service=\\\"unknownt\\\" and subservice=\\\"/\\\")\",\n          \"state\": \"STARTED\"\n        }\n      },\n      {\n        \"code\": 200,\n        \"body\": {\n          \"name\": \"blood_rule_update@unknownt/\",\n          \"timeLastStateChange\": 1613028692624,\n          \"text\": \"context ctxt$unknownt$ select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n          \"state\": \"STARTED\"\n        }\n      },\n      null\n    ],\n    {\n      \"n\": 1,\n      \"ok\": 1\n    },\n    null\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesCreateErrorDataEmpty(t *testing.T) {
	c := setupTest([]string{"rules", "create", "--host", "perseo", "--data", "@"})

	err := perseoRulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestPerseoRulesCreateErrorHTTP(t *testing.T) {
	data := `{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`
	c := setupTest([]string{"rules", "create", "--host", "perseo", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rule"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestPerseoRulesCreateErrorHTTPStatus(t *testing.T) {
	data := `{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`
	c := setupTest([]string{"rules", "create", "--host", "perseo", "--data", data})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoRulesCreateErrorPretty(t *testing.T) {
	data := `{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`
	c := setupTest([]string{"rules", "create", "--host", "perseo", "--data", data, "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[false,[{"code":200,"body":{"name":"ctxt$unknownt$","timeLastStateChange":1613024032325,"text":"create context ctxt$unknownt$ partition by service from iotEvent(service=\"unknownt\" and subservice=\"/\")","state":"STARTED"}},{"code":200,"body":{"name":"blood_rule_update@unknownt/","timeLastStateChange":1613028692624,"text":"context ctxt$unknownt$ select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","state":"STARTED"}},null],{"n":1,"ok":1},null]}`)
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := perseoRulesCreate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoRulesDelete(t *testing.T) {
	c := setupTest([]string{"rules", "delete", "--host", "perseo", "--name", "blood_rule_update"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestPerseoRulesDeleteErrorHTTP(t *testing.T) {
	c := setupTest([]string{"rules", "delete", "--host", "perseo", "--name", "blood_rule_update"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestPerseoRulesDeleteErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"rules", "delete", "--host", "perseo", "--name", "blood_rule_update"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	helper.SetClientHTTP(c, reqRes)

	err := perseoRulesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeNoParam(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "blood_rule_update\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeRaw(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--raw"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"error\":null,\"data\":[{\"_id\":\"6024cb208e2bfc0012c77488\",\"name\":\"blood_rule_update\",\"text\":\"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"action\":{\"type\":\"update\",\"parameters\":{\"attributes\":[{\"name\":\"abnormal\",\"value\":\"true\",\"type\":\"boolean\"}]}},\"subservice\":\"/\",\"service\":\"unknownt\"}],\"count\":1}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeRawPretty(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--raw", "--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"error\": null,\n  \"data\": [\n    {\n      \"_id\": \"6024cb208e2bfc0012c77488\",\n      \"name\": \"blood_rule_update\",\n      \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n      \"action\": {\n        \"type\": \"update\",\n        \"parameters\": {\n          \"attributes\": [\n            {\n              \"name\": \"abnormal\",\n              \"value\": \"true\",\n              \"type\": \"boolean\"\n            }\n          ]\n        }\n      },\n      \"subservice\": \"/\",\n      \"service\": \"unknownt\"\n    }\n  ],\n  \"count\": 1\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeCount(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--count"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeVerbose(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--verbose"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"_id\":\"6024cb208e2bfc0012c77488\",\"name\":\"blood_rule_update\",\"text\":\"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"action\":{\"parameters\":{\"attributes\":[{\"name\":\"abnormal\",\"type\":\"boolean\",\"value\":\"true\"}]},\"type\":\"update\"},\"subservice\":\"/\",\"service\":\"unknownt\"}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposePretty(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"_id\": \"6024cb208e2bfc0012c77488\",\n    \"name\": \"blood_rule_update\",\n    \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n    \"action\": {\n      \"parameters\": {\n        \"attributes\": [\n          {\n            \"name\": \"abnormal\",\n            \"type\": \"boolean\",\n            \"value\": \"true\"\n          }\n        ]\n      },\n      \"type\": \"update\"\n    },\n    \"subservice\": \"/\",\n    \"service\": \"unknownt\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeErrorRawPretty(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--raw", "--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetJSONIndentError(c.Ngsi)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeErrorNoParam(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeErrorVerbose(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--verbose"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeErrorPretty(t *testing.T) {
	c := setupTest([]string{"rules", "list", "--host", "perseo", "--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	helper.SetJSONIndentError(c.Ngsi)

	err := perseoPrintRespose(c, c.Ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
