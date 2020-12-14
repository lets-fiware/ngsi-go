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
	"errors"
	"flag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestSubscriptionssubscriptionsListV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Count(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagBool(set, "count")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--count"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "6\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2CountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Page(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(subscriptionData)
	reqRes1.Path = "/v2/subscriptions"
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"106"}}
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(subscriptionData)
	reqRes2.Path = "/v2/subscriptions"
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"106"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Status(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=active"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Query(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--query=FIWARE*"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Json(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	set.Bool("json", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--json"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2JsonCount0(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	set.Bool("json", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--json"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Verbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	set.Bool("verbose", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--verbose"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "9f6c254ac4a6068bb276774e inactive 2020-09-24T07:49:13.00Z ngsi source subscription\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Localtime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	setupFlagBool(set, "verbose,localTime")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--verbose", "--localTime"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "9f6c254ac4a6068bb276774e inactive 2020-09-24T16:49:13.00+0900 ngsi source subscription\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2Items(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query,items")
	setupFlagBool(set, "verbose,localTime")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--verbose", "--localTime", "--items=status,expires"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "9f6c254ac4a6068bb276774e inactive 2020-09-24T16:49:13.00+0900\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2ErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query,items")
	setupFlagBool(set, "verbose,localTime")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=abc", "--verbose", "--localTime", "--items=status,expires"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: abc (active, inactive, oneshot, expired, failed)", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscription"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query,items")
	setupFlagBool(set, "verbose,localTime")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--verbose", "--localTime", "--items=status,expires"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2ErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query,items")
	setupFlagBool(set, "verbose,localTime")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--verbose", "--localTime", "--items=status,expires"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2ErrorRessultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2ErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2ErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagBool(set, "json")
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json"})
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListV2ErrorHTTPItems(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query,items")
	setupFlagBool(set, "verbose,localTime")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=inactive", "--verbose", "--localTime", "--items=status,expires,error"})

	err := subscriptionsListV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	setupFlagString(set, "id")
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionGetV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"4f6c2576c4a6068bb276774f\",\"description\":\"FIWARE\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"WeatherObserved\"}],\"condition\":{\"attrs\":[\"dateRetrieved\"]}},\"notification\":{\"timesSent\":278,\"lastNotification\":\"2020-09-24T07:40:26.00Z\",\"lastSuccess\":\"2020-09-24T07:40:26.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:56.00Z\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetV2LocalTime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	setupFlagString(set, "id")
	setupFlagBool(set, "localTime")
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c", "--localTime"})

	err := subscriptionGetV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"4f6c2576c4a6068bb276774f\",\"description\":\"FIWARE\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"WeatherObserved\"}],\"condition\":{\"attrs\":[\"dateRetrieved\"]}},\"notification\":{\"timesSent\":278,\"lastNotification\":\"2020-09-24T16:40:26.00+0900\",\"lastSuccess\":\"2020-09-24T16:40:26.00+0900\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T16:49:56.00+0900\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
func TestSubscriptionsGetV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	setupFlagString(set, "id")
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "  5f0a44789dd803416ccbf15c", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetV2ErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	reqRes := MockHTTPReqRes{}
	setupFlagString(set, "id")
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetV2ErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	reqRes := MockHTTPReqRes{}
	setupFlagString(set, "id")
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionGetV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/subscriptions/5f0a44789dd803416ccbf15c"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires,entityId,uri")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--entityId=abc", "--uri=http://ngsiproxy"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsCreateV2(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f0a44789dd803416ccbf15c\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateV2ErrorsetSubscriptionValuesV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v2/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires,data")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--data=", "--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateV2ErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/subscriptions/5f0a44789dd803416ccbf15c"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires,entityId,url")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--entityId=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscription"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires,entityId,url")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--entityId=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateV2ErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires,entityId,url")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--entityId=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsCreateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateV2(c, ngsi, client)

	assert.NoError(t, err)
}

func TestSubscriptionsUpdateV2wAttr(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"subject":{"condition":{"attrs":["Temp"]}}}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,wAttrs")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--wAttrs=Temp"})

	err := subscriptionsUpdateV2(c, ngsi, client)

	assert.NoError(t, err)
}

func TestSubscriptionsUpdateV2ErrotsetSubscriptionValuesV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscription/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2"})

	err := subscriptionsUpdateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
func TestSubscriptionsUpdateV2Marshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestSubscriptionsUpdateV2ErrotHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscription/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateV2ErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionsDeleteV2(c, ngsi, client)

	assert.NoError(t, err)
}

func TestSubscriptionsDeleteV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscription/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionsDeleteV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionsDeleteV2(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,entityId,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--entityId=abc", "--url=http://ngsiproxy"})

	err := subscriptionsTemplateV2(c, ngsi)

	assert.NoError(t, err)
}

func TestSubscriptionsTemplateV2Error(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data,entityId,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=", "--url=http://ngsiproxy"})

	err := subscriptionsTemplateV2(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateV2ErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	setupFlagString(set, "data,entityId,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--entityId=abc", "--url=http://ngsiproxy"})

	err := subscriptionsTemplateV2(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Data(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "data,entityId,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--entityId=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2getAttributes(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	sub := subscriptionV2{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data,entityId,url,host,type,link")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--entityId=abc", "--url=http://ngsiproxy", "--host=orion", "--type=abc", "--get"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2IdPattern(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "data,entityId,idPattern,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--idPattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2TypePattern(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "typePattern,data,entityId,idPattern,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}", "--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2wAttrs1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,wAttrs")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--wAttrs=abc,def,xyz", "--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2wAttrs2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,wAttrs")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--wAttrs=abc,def,xyz", "--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2query1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,query,mq,georel,geometry,coords")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--query=abc", "--mq=def", "--georel=123", "--geometry=456", "--coords=789", "--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2query2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,query,mq,georel,geometry,coords")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--query=abc", "--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2query3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)
	sub.Subject.Condition = new(subscriptionConditionV2)
	sub.Subject.Condition.Expression = new(subscriptionExpressionV2)

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,query,mq,georel,geometry,coords")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--query=abc", "--mq=def", "--georel=123", "--geometry=456", "--coords=789", "--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2url(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}
	sub.Notification = new(subscriptionNotificationV2)
	sub.Notification.HTTP = new(subscriptionHTTPV2)

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,query,mq,georel,geometry,coords")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2headers(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}
	sub.Notification = new(subscriptionNotificationV2)
	sub.Notification.HTTP = new(subscriptionHTTPV2)

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,headers,qs,method,payload")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{`--headers={"abc":"123","xyz":"456"}`})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2qs(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}
	sub.Notification = new(subscriptionNotificationV2)
	sub.Notification.HTTPCustom = new(subscriptionHTTPCustomV2)

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,headers,qs,method,payload")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{`--qs={"abc":"123","xyz":"456"}`})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2method(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,method,payload,nAttrs,metadata")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{"--method=post", "--payload=abc", "--nAttrs=abc,xyz", "--metadata=abc,xyz"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2exceptAttrs(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,exceptAttrs,attrsFormat,throttling,expires,status")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{"--exceptAttrs=abc,xyz", "--attrsFormat=abc", "--throttling=1", "--expires=1day", "--status=oneshot"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2expires(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "typePattern,data,entityId,idPattern,url,exceptAttrs,attrsFormat,throttling,expires,status")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--idPattern=abc", "--typePattern=abc", "--url=http://ngsiproxy"})
	_ = set.Parse([]string{"--exceptAttrs=abc,xyz", "--attrsFormat=abc", "--throttling=1", "--expires=2020-10-05T00:58:26.929Z", "--status=active"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2Error1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "data,entityId,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=", "--entityId=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "data,entityId,url")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={id}", "--entityId=abc", "--url=http://ngsiproxy"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data,entityId,url,host,type,get,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--entityId=abc", "--url=http://ngsiproxy", "--host=orion", "--type=abc", "--get", "--link=abc"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error5(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "entityId,url,type,typePattern")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--entityId=abc", "--url=http://ngsiproxy", "--type=123", "--typePattern=xyz"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error6(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "entityId,url,type,typePattern,headers")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--headers={id}", "--entityId=abc", "--url=http://ngsiproxy", "--type=123"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error7(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "entityId,url,type,typePattern,headers,qs")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--qs={id}", "--entityId=abc", "--url=http://ngsiproxy", "--type=123"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error8(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "entityId,url,type,typePattern,headers,qs,expires,exceptAttrs,nAttrs")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--exceptAttrs=abc", "--nAttrs=abc", "--entityId=abc", "--url=http://ngsiproxy", "--type=123"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error9(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "entityId,url,type,typePattern,headers,qs,expires")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--expires=1", "--entityId=abc", "--url=http://ngsiproxy", "--type=123"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesV2Error10(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	sub := subscriptionV2{}

	setupFlagString(set, "entityId,url,type,typePattern,headers,qs,expires,status")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--status=error", "--entityId=abc", "--url=http://ngsiproxy", "--type=123"})

	err := setSubscriptionValuesV2(c, ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestGtAttributesV2Ok1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,link")
	set.Bool("get", false, "")
	sub := subscriptionV2{}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=abc", "--get"})
	err := getAttributesV2(c, ngsi, &sub)

	assert.NoError(t, err)
}

func TestGtAttributesV2Ok2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,link")
	set.Bool("get", false, "")
	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)
	sub.Subject.Entities = append(sub.Subject.Entities, *new(subscriptionEntityV2))

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=abc", "--get"})
	err := getAttributesV2(c, ngsi, &sub)

	assert.NoError(t, err)
}

func TestGtAttributesV2NoError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,type")
	sub := subscriptionV2{}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := getAttributesV2(c, ngsi, &sub)

	assert.NoError(t, err)
}

func TestGtAttributesV2ErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,type,link")
	set.Bool("get", false, "")
	sub := subscriptionV2{}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=abc", "--get", "--link=abc"})
	err := getAttributesV2(c, ngsi, &sub)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGtAttributesV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Path = "/v2/types"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,link")
	set.Bool("get", false, "")
	sub := subscriptionV2{}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=abc", "--get"})
	err := getAttributesV2(c, ngsi, &sub)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGtAttributesV2ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/abc"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type,link")
	set.Bool("get", false, "")
	sub := subscriptionV2{}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=abc", "--get"})
	err := getAttributesV2(c, ngsi, &sub)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestToLocaltime(t *testing.T) {
	sub := subscriptionResposeV2{}
	sub.Expires = "2020-10-01T00:00:00.00Z"
	sub.Notification.LastNotification = "2020-10-01T01:10:00.00Z"
	sub.Notification.LastSuccess = "2020-10-01T02:12:00.00Z"
	sub.Notification.LastFailure = "2020-10-01T03:13:00.00Z"

	toLocaltime(&sub)

	assert.Equal(t, "2020-10-01T09:00:00.00+0900", sub.Expires)
	assert.Equal(t, "2020-10-01T10:10:00.00+0900", sub.Notification.LastNotification)
	assert.Equal(t, "2020-10-01T11:12:00.00+0900", sub.Notification.LastSuccess)
	assert.Equal(t, "2020-10-01T12:13:00.00+0900", sub.Notification.LastFailure)
}

func TestGetLocalTime1(t *testing.T) {
	actual := getLocalTime("")
	expected := ""
	assert.Equal(t, expected, actual)
}

func TestGetLocalTime2(t *testing.T) {
	actual := getLocalTime("2020-10-01T00:00:00.00Z")
	expected := "2020-10-01T09:00:00.00+0900"
	assert.Equal(t, expected, actual)
}

func TestGetLocalTime3(t *testing.T) {
	actual := getLocalTime("2020-10-01T00:00:00.000Z")
	expected := "2020-10-01T09:00:00.000+0900"
	assert.Equal(t, expected, actual)
}

func TestGetLocalTime4(t *testing.T) {
	actual := getLocalTime("2020-10-01T00:00:00")
	expected := "2020-10-01T00:00:00"
	assert.Equal(t, expected, actual)
}

func TestCheckItems(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	setupFlagString(set, "items")
	app := cli.NewApp()
	_ = set.Parse([]string{"--items=description,timessent,lastnotification,lastsuccess,lastsuccesscode,url,expires,status"})
	c := cli.NewContext(app, set, nil)
	_, err := checkItems(c)

	assert.NoError(t, err)
}

func TestCheckItemsError(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	setupFlagString(set, "items")
	app := cli.NewApp()
	_ = set.Parse([]string{"--items=abc"})
	c := cli.NewContext(app, set, nil)
	_, err := checkItems(c)

	if assert.Error(t, err) {
		actual := err.Error()
		expected := "error: abc in --items"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSprintItems1(t *testing.T) {
	sub := subscriptionResposeV2{}
	http := subscriptionHTTPV2{}
	code := 404

	sub.ID = "6f6c2576c4a6068bb2767743"
	sub.Description = "test subscription"
	sub.Notification.LastNotification = "2020-10-01T00:00:00.00Z"
	sub.Notification.LastSuccess = "2020-10-02T01:00:00.00Z"
	sub.Notification.LastSuccessCode = &code
	sub.Notification.HTTP = &http
	sub.Notification.HTTP.URL = "https://ngsiproxy"
	sub.Expires = "2021-12-31T01:00:00.00Z"
	sub.Status = "expired"

	items := []string{"id", "description", "timessent", "lastnotification", "lastsuccess", "lastsuccesscode", "url", "expires", "status"}

	actual := sprintItems(&sub, items)
	expected := "6f6c2576c4a6068bb2767743 test subscription - 2020-10-01T00:00:00.00Z 2020-10-02T01:00:00.00Z 404 https://ngsiproxy 2021-12-31T01:00:00.00Z expired"
	assert.Equal(t, expected, actual)
}

func TestSprintItems2(t *testing.T) {
	sub := subscriptionResposeV2{}
	http := subscriptionHTTPCustomV2{}
	code := 404
	var timesSent int64
	timesSent = 10

	sub.ID = "6f6c2576c4a6068bb2767743"
	sub.Description = "test subscription"
	sub.Notification.TimesSent = &timesSent
	sub.Notification.LastNotification = "2020-10-01T00:00:00.00Z"
	sub.Notification.LastSuccess = "2020-10-02T01:00:00.00Z"
	sub.Notification.LastSuccessCode = &code
	sub.Notification.HTTPCustom = &http
	sub.Notification.HTTPCustom.URL = "https://ngsiproxy"
	sub.Expires = "2021-12-31T01:00:00.00Z"
	sub.Status = "expired"

	items := []string{"id", "description", "timessent", "lastnotification", "lastsuccess", "lastsuccesscode", "url", "expires", "status"}

	// assert.NoError(t, err)
	actual := sprintItems(&sub, items)
	expected := "6f6c2576c4a6068bb2767743 test subscription 10 2020-10-01T00:00:00.00Z 2020-10-02T01:00:00.00Z 404 https://ngsiproxy 2021-12-31T01:00:00.00Z expired"
	assert.Equal(t, expected, actual)
}

var subscriptionData = `[
	{
	  "id": "3ea2e78f675f2d199d3025ff",
	  "description": "ngsi source subscription",
	  "subject": {
		"entities": [
		  {
			"idPattern": ".*"
		  }
		],
		"condition": {
		  "attrs": [
			"observed",
			"location"
		  ]
		}
	  },
	  "notification": {
		"timesSent": 3406,
		"lastNotification": "2020-09-01T07:43:00.00Z",
		"lastSuccess": "2020-09-01T07:43:04.00Z",
		"lastSuccessCode": 204,
		"lastFailure": "2020-09-01T07:42:07.00Z",
		"lastFailureReason": "Timeout was reached",
		"onlyChangedAttrs": false,
		"http": {
		  "url": "https://ngsiproxy"
		},
		"attrsFormat": "keyValues"
	  },
	  "expires": "2020-09-01T01:24:01.00Z",
	  "status": "expired"
	},
	{
	  "id": "5f64060ef6752d199d302600",
	  "description": "ngsi source subscription",
	  "subject": {
		"entities": [
		  {
			"idPattern": ".*",
			"type": "WeatherObserved"
		  }
		],
		"condition": {
		  "attrs": [
			"dateRetrieved"
		  ]
		}
	  },
	  "notification": {
		"timesSent": 27,
		"lastNotification": "2020-09-16T03:40:27.00Z",
		"lastSuccess": "2020-09-16T03:40:28.00Z",
		"lastSuccessCode": 404,
		"onlyChangedAttrs": false,
		"http": {
		  "url": "https://ngsiproxy"
		},
		"attrsFormat": "keyValues"
	  },
	  "expires": "2020-09-16T03:57:49.00Z",
	  "status": "expired"
	},
	{
	  "id": "1f32db4bf6752d199d302601",
	  "description": "ngsi source subscription",
	  "subject": {
		"entities": [
		  {
			"idPattern": ".*"
		  }
		],
		"condition": {
		  "attrs": [
			"observed",
			"location"
		  ]
		}
	  },
	  "notification": {
		"timesSent": 3408,
		"lastNotification": "2020-09-16T04:03:00.00Z",
		"lastSuccess": "2020-09-16T04:03:04.00Z",
		"lastSuccessCode": 404,
		"lastFailure": "2020-09-16T03:40:06.00Z",
		"lastFailureReason": "Timeout was reached",
		"onlyChangedAttrs": false,
		"http": {
		  "url": "https://ngsiproxy"
		},
		"attrsFormat": "keyValues"
	  },
	  "expires": "2020-09-16T04:03:05.00Z",
	  "status": "expired"
	},
	{
	  "id": "3978fabd87752d199d302602",
	  "description": "ngsi source subscription",
	  "subject": {
		"entities": [
		  {
			"idPattern": ".*",
			"type": "WeatherObserved"
		  }
		],
		"condition": {
		  "attrs": [
			"dateRetrieved"
		  ]
		}
	  },
	  "notification": {
		"timesSent": 10,
		"lastNotification": "2020-09-16T04:00:13.00Z",
		"lastSuccess": "2020-09-16T04:00:13.00Z",
		"lastSuccessCode": 204,
		"onlyChangedAttrs": false,
		"http": {
		  "url": "https://ngsiproxy"
		},
		"attrsFormat": "keyValues"
	  },
	  "expires": "2020-09-16T04:03:07.00Z",
	  "status": "expired"
	},
	{
	  "id": "9f6c254ac4a6068bb276774e",
	  "description": "ngsi source subscription",
	  "subject": {
		"entities": [
		  {
			"idPattern": ".*"
		  }
		],
		"condition": {
		  "attrs": [
			"dateObserved"
		  ]
		}
	  },
	  "notification": {
		"timesSent": 28,
		"lastNotification": "2020-09-24T07:30:02.00Z",
		"lastSuccess": "2020-09-24T07:30:02.00Z",
		"lastSuccessCode": 404,
		"onlyChangedAttrs": false,
		"http": {
		  "url": "https://ngsiproxy"
		},
		"attrsFormat": "keyValues"
	  },
	  "expires": "2020-09-24T07:49:13.00Z",
	  "status": "inactive"
	},
	{
	  "id": "4f6c2576c4a6068bb276774f",
	  "description": "FIWARE",
	  "subject": {
		"entities": [
		  {
			"idPattern": ".*",
			"type": "WeatherObserved"
		  }
		],
		"condition": {
		  "attrs": [
			"dateRetrieved"
		  ]
		}
	  },
	  "notification": {
		"timesSent": 278,
		"lastNotification": "2020-09-24T07:40:26.00Z",
		"lastSuccess": "2020-09-24T07:40:26.00Z",
		"lastSuccessCode": 404,
		"onlyChangedAttrs": false,
		"http": {
		  "url": "https://ngsiproxy"
		},
		"attrsFormat": "keyValues"
	  },
	  "expires": "2020-09-24T07:49:56.00Z",
	  "status": "active"
	}
  ]`
