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

func TestPerseoRulesList(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo"})

	err := perseoRulesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "blood_rule_update\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesListPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "raw,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--raw", "--pretty"})

	err := perseoRulesList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"error\": null,\n  \"data\": [\n    {\n      \"_id\": \"6024cb208e2bfc0012c77488\",\n      \"name\": \"blood_rule_update\",\n      \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n      \"action\": {\n        \"type\": \"update\",\n        \"parameters\": {\n          \"attributes\": [\n            {\n              \"name\": \"abnormal\",\n              \"value\": \"true\",\n              \"type\": \"boolean\"\n            }\n          ]\n        }\n      },\n      \"subservice\": \"/\",\n      \"service\": \"unknownt\"\n    }\n  ],\n  \"count\": 1\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := perseoRulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesListErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := perseoRulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesListErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo"})

	err := perseoRulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestPerseoRulesListErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo"})

	err := perseoRulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoRulesListErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "raw,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--raw", "--pretty"})

	setJSONIndentError(ngsi)

	err := perseoRulesList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoRulesGet(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update"})

	err := perseoRulesGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"error\":null,\"data\":{\"_id\":\"6024cb208e2bfc0012c77488\",\"name\":\"blood_rule_update\",\"text\":\"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"action\":{\"type\":\"update\",\"parameters\":{\"attributes\":[{\"name\":\"abnormal\",\"value\":\"true\",\"type\":\"boolean\"}]}},\"subservice\":\"/\",\"service\":\"unknownt\"}}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesGetPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update", "--pretty"})

	err := perseoRulesGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"error\": null,\n  \"data\": {\n    \"_id\": \"6024cb208e2bfc0012c77488\",\n    \"name\": \"blood_rule_update\",\n    \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n    \"action\": {\n      \"type\": \"update\",\n      \"parameters\": {\n        \"attributes\": [\n          {\n            \"name\": \"abnormal\",\n            \"value\": \"true\",\n            \"type\": \"boolean\"\n          }\n        ]\n      }\n    },\n    \"subservice\": \"/\",\n    \"service\": \"unknownt\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := perseoRulesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesGetErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := perseoRulesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesGetErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/device"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo"})

	err := perseoRulesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "rule name not found", ngsiErr.Message)
	}
}

func TestPerseoRulesGetErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update"})

	err := perseoRulesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestPerseoRulesGetErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update"})

	err := perseoRulesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoRulesGetErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"error":null,"data":{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update", "--pretty"})

	setJSONIndentError(ngsi)

	err := perseoRulesGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoRulesCreate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", `--data={"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`})

	err := perseoRulesCreate(c)

	assert.NoError(t, err)
}

func TestPerseoRulesCreateVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[false,[{"code":200,"body":{"name":"ctxt$unknownt$","timeLastStateChange":1613024032325,"text":"create context ctxt$unknownt$ partition by service from iotEvent(service=\"unknownt\" and subservice=\"/\")","state":"STARTED"}},{"code":200,"body":{"name":"blood_rule_update@unknownt/","timeLastStateChange":1613028692624,"text":"context ctxt$unknownt$ select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","state":"STARTED"}},null],{"n":1,"ok":1},null]}`)
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "verbose")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--verbose", `--data={"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`})

	err := perseoRulesCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"error\":null,\"data\":[false,[{\"code\":200,\"body\":{\"name\":\"ctxt$unknownt$\",\"timeLastStateChange\":1613024032325,\"text\":\"create context ctxt$unknownt$ partition by service from iotEvent(service=\\\"unknownt\\\" and subservice=\\\"/\\\")\",\"state\":\"STARTED\"}},{\"code\":200,\"body\":{\"name\":\"blood_rule_update@unknownt/\",\"timeLastStateChange\":1613028692624,\"text\":\"context ctxt$unknownt$ select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"state\":\"STARTED\"}},null],{\"n\":1,\"ok\":1},null]}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesCreatePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[false,[{"code":200,"body":{"name":"ctxt$unknownt$","timeLastStateChange":1613024032325,"text":"create context ctxt$unknownt$ partition by service from iotEvent(service=\"unknownt\" and subservice=\"/\")","state":"STARTED"}},{"code":200,"body":{"name":"blood_rule_update@unknownt/","timeLastStateChange":1613028692624,"text":"context ctxt$unknownt$ select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","state":"STARTED"}},null],{"n":1,"ok":1},null]}`)
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--pretty", `--data={"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`})

	err := perseoRulesCreate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"error\": null,\n  \"data\": [\n    false,\n    [\n      {\n        \"code\": 200,\n        \"body\": {\n          \"name\": \"ctxt$unknownt$\",\n          \"timeLastStateChange\": 1613024032325,\n          \"text\": \"create context ctxt$unknownt$ partition by service from iotEvent(service=\\\"unknownt\\\" and subservice=\\\"/\\\")\",\n          \"state\": \"STARTED\"\n        }\n      },\n      {\n        \"code\": 200,\n        \"body\": {\n          \"name\": \"blood_rule_update@unknownt/\",\n          \"timeLastStateChange\": 1613028692624,\n          \"text\": \"context ctxt$unknownt$ select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n          \"state\": \"STARTED\"\n        }\n      },\n      null\n    ],\n    {\n      \"n\": 1,\n      \"ok\": 1\n    },\n    null\n  ]\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoRulesCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := perseoRulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesCreateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := perseoRulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesCreateErrorData(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/rules"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo"})

	err := perseoRulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "--data not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesCreateErrorDataEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/rules"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", `--data=`})

	err := perseoRulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesCreateErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rule"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", `--data={"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`})
	err := perseoRulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestPerseoRulesCreateErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules"
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", `--data={"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`})

	err := perseoRulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoRulesCreateErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules"
	reqRes.ResBody = []byte(`{"error":null,"data":[false,[{"code":200,"body":{"name":"ctxt$unknownt$","timeLastStateChange":1613024032325,"text":"create context ctxt$unknownt$ partition by service from iotEvent(service=\"unknownt\" and subservice=\"/\")","state":"STARTED"}},{"code":200,"body":{"name":"blood_rule_update@unknownt/","timeLastStateChange":1613028692624,"text":"context ctxt$unknownt$ select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","state":"STARTED"}},null],{"n":1,"ok":1},null]}`)
	reqRes.ReqData = []byte(`{"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`)

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,data")
	setupFlagBool(set, "pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--pretty", `--data={"name":"blood_rule_update","text":"select *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}}}`})

	setJSONIndentError(ngsi)

	err := perseoRulesCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoRulesDelete(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update"})

	err := perseoRulesDelete(c)

	assert.NoError(t, err)
}

func TestPerseoRulesDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := perseoRulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesDeleteErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := perseoRulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesDeleteErrorName(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/rules/blood_rule_update"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo"})

	err := perseoRulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "rule name not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPerseoRulesDeleteErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/rules"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update"})

	err := perseoRulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestPerseoRulesDeleteErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/rules/blood_rule_update"
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,name")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=perseo", "--name=blood_rule_update"})

	err := perseoRulesDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeNoParam(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagBool(set, "raw,verboser, pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "blood_rule_update\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeRaw(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagBool(set, "raw,verboser, pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--raw"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"error\":null,\"data\":[{\"_id\":\"6024cb208e2bfc0012c77488\",\"name\":\"blood_rule_update\",\"text\":\"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"action\":{\"type\":\"update\",\"parameters\":{\"attributes\":[{\"name\":\"abnormal\",\"value\":\"true\",\"type\":\"boolean\"}]}},\"subservice\":\"/\",\"service\":\"unknownt\"}],\"count\":1}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeRawPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagBool(set, "raw,verboser,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--raw", "--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"error\": null,\n  \"data\": [\n    {\n      \"_id\": \"6024cb208e2bfc0012c77488\",\n      \"name\": \"blood_rule_update\",\n      \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n      \"action\": {\n        \"type\": \"update\",\n        \"parameters\": {\n          \"attributes\": [\n            {\n              \"name\": \"abnormal\",\n              \"value\": \"true\",\n              \"type\": \"boolean\"\n            }\n          ]\n        }\n      },\n      \"subservice\": \"/\",\n      \"service\": \"unknownt\"\n    }\n  ],\n  \"count\": 1\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeCount(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagBool(set, "raw,verbose,count,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--count"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "1\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagBool(set, "raw,verbose,count,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--verbose"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"_id\":\"6024cb208e2bfc0012c77488\",\"name\":\"blood_rule_update\",\"text\":\"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\"action\":{\"parameters\":{\"attributes\":[{\"name\":\"abnormal\",\"type\":\"boolean\",\"value\":\"true\"}]},\"type\":\"update\"},\"subservice\":\"/\",\"service\":\"unknownt\"}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposePretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagBool(set, "raw,verbose,count,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"_id\": \"6024cb208e2bfc0012c77488\",\n    \"name\": \"blood_rule_update\",\n    \"text\": \"select \\\"blood_rule_update\\\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\\\"BloodMeter\\\")]\",\n    \"action\": {\n      \"parameters\": {\n        \"attributes\": [\n          {\n            \"name\": \"abnormal\",\n            \"type\": \"boolean\",\n            \"value\": \"true\"\n          }\n        ]\n      },\n      \"type\": \"update\"\n    },\n    \"subservice\": \"/\",\n    \"service\": \"unknownt\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestPerseoPrintResposeErrorRawPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "raw,verboser,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--raw", "--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	setJSONIndentError(ngsi)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeErrorNoParam(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "raw,verboser, pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	setJSONDecodeErr(ngsi, 0)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeErrorVerbose(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "raw,verbose,count,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--verbose"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	setJSONEncodeErr(ngsi, 0)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestPerseoPrintResposeErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "raw,verbose,count,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--pretty"})

	b := []byte(`{"error":null,"data":[{"_id":"6024cb208e2bfc0012c77488","name":"blood_rule_update","text":"select \"blood_rule_update\" as ruleName, *, *, ev.BloodPressure? as Pressure, ev.id? as Meter from pattern [every ev=iotEvent(cast(cast(BloodPressure?,String),float)>1.5 and type=\"BloodMeter\")]","action":{"type":"update","parameters":{"attributes":[{"name":"abnormal","value":"true","type":"boolean"}]}},"subservice":"/","service":"unknownt"}],"count":1}`)

	setJSONIndentError(ngsi)

	err := perseoPrintRespose(c, ngsi, b)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
