/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestSubscriptionssubscriptionsListV2(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Count(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "6\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2CountZero(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Page(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(subscriptionData)
	reqRes1.Path = "/v2/subscriptions"
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"106"}}

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(subscriptionData)
	reqRes2.Path = "/v2/subscriptions"
	reqRes2.ResHeader = http.Header{"Fiware-Total-Count": []string{"106"}}

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Status(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "active"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Query(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--query", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Json(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "inactive", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2JsonPretty(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "inactive", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"9f6c254ac4a6068bb276774e\",\n    \"description\": \"ngsi source subscription\",\n    \"subject\": {\n      \"entities\": [\n        {\n          \"idPattern\": \".*\"\n        }\n      ],\n      \"condition\": {\n        \"attrs\": [\n          \"dateObserved\"\n        ]\n      }\n    },\n    \"notification\": {\n      \"timesSent\": 28,\n      \"lastNotification\": \"2020-09-24T07:30:02.00Z\",\n      \"lastSuccess\": \"2020-09-24T07:30:02.00Z\",\n      \"lastSuccessCode\": 404,\n      \"onlyChangedAttrs\": false,\n      \"http\": {\n        \"url\": \"https://ngsiproxy\"\n      },\n      \"attrsFormat\": \"keyValues\"\n    },\n    \"expires\": \"2020-09-24T07:49:13.00Z\",\n    \"status\": \"inactive\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2JsonCount0(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "inactive", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Verbose(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "inactive", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "9f6c254ac4a6068bb276774e inactive 2020-09-24T07:49:13.00Z ngsi source subscription\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Localtime(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "inactive", "--verbose", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "9f6c254ac4a6068bb276774e inactive 2020-09-24T16:49:13.00+0900 ngsi source subscription\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2Items(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "inactive", "--verbose", "--localTime", "--items", "status,expires"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "9f6c254ac4a6068bb276774e inactive 2020-09-24T16:49:13.00+0900\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorStatus(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--host", "orion", "--status", "abc", "--verbose", "--localTime", "--items", "status,expires"})

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: abc (active, inactive, oneshot, expired, failed)", ngsiErr.Message)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscription"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorRessultsCount(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorMarshal(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorJSONPretty(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionssubscriptionsListV2ErrorHTTPItems(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--status", "inactive", "--verbose", "--localTime", "--items", "status,expires,error"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionData)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
	}
}

func TestSubscriptionsGetV2(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"4f6c2576c4a6068bb276774f\",\"description\":\"FIWARE\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"WeatherObserved\"}],\"condition\":{\"attrs\":[\"dateRetrieved\"]}},\"notification\":{\"timesSent\":278,\"lastNotification\":\"2020-09-24T07:40:26.00Z\",\"lastSuccess\":\"2020-09-24T07:40:26.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:56.00Z\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsGetV2Raw(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "615d5b66f19d2a10c44e264c", "--raw"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/615d5b66f19d2a10c44e264c"
	reqRes.ResBody = []byte(`{"id":"615d5b66f19d2a10c44e264c","description":"TestNotification","status":"active","subject":{"entities":[{"idPattern":"Alert.*","type":"Alert"}],"condition":{"attrs":[]}},"notification":{"attrs":[],"onlyChangedAttrs":false,"attrsFormat":"keyValues","httpCustom":{"url":"http://dev/null","headers":{"fiware-shared-key":"test"}}}}`)
	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"615d5b66f19d2a10c44e264c\",\"description\":\"TestNotification\",\"status\":\"active\",\"subject\":{\"entities\":[{\"idPattern\":\"Alert.*\",\"type\":\"Alert\"}],\"condition\":{\"attrs\":[]}},\"notification\":{\"attrs\":[],\"onlyChangedAttrs\":false,\"attrsFormat\":\"keyValues\",\"httpCustom\":{\"url\":\"http://dev/null\",\"headers\":{\"fiware-shared-key\":\"test\"}}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsGetV2Pretty(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"4f6c2576c4a6068bb276774f\",\n  \"description\": \"FIWARE\",\n  \"subject\": {\n    \"entities\": [\n      {\n        \"idPattern\": \".*\",\n        \"type\": \"WeatherObserved\"\n      }\n    ],\n    \"condition\": {\n      \"attrs\": [\n        \"dateRetrieved\"\n      ]\n    }\n  },\n  \"notification\": {\n    \"timesSent\": 278,\n    \"lastNotification\": \"2020-09-24T07:40:26.00Z\",\n    \"lastSuccess\": \"2020-09-24T07:40:26.00Z\",\n    \"lastSuccessCode\": 404,\n    \"onlyChangedAttrs\": false,\n    \"http\": {\n      \"url\": \"https://ngsiproxy\"\n    },\n    \"attrsFormat\": \"keyValues\"\n  },\n  \"expires\": \"2020-09-24T07:49:56.00Z\",\n  \"status\": \"active\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsGetV2LocalTime(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"4f6c2576c4a6068bb276774f\",\"description\":\"FIWARE\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"WeatherObserved\"}],\"condition\":{\"attrs\":[\"dateRetrieved\"]}},\"notification\":{\"timesSent\":278,\"lastNotification\":\"2020-09-24T16:40:26.00+0900\",\"lastSuccess\":\"2020-09-24T16:40:26.00+0900\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T16:49:56.00+0900\",\"status\":\"active\"}"
		assert.Equal(t, expected, actual)
	}
}
func TestSubscriptionsGetV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "  5f0a44789dd803416ccbf15c", ngsiErr.Message)
	}
}

func TestSubscriptionsGetV2ErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsGetV2ErrorMarshal(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsGetV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte(`{"id":"4f6c2576c4a6068bb276774f","description":"FIWARE","subject":{"entities":[{"idPattern":".*","type":"WeatherObserved"}],"condition":{"attrs":["dateRetrieved"]}},"notification":{"timesSent":278,"lastNotification":"2020-09-24T07:40:26.00Z","lastSuccess":"2020-09-24T07:40:26.00Z","lastSuccessCode":404,"onlyChangedAttrs":false,"http":{"url":"https://ngsiproxy"},"attrsFormat":"keyValues"},"expires":"2020-09-24T07:49:56.00Z","status":"active"}`)

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := subscriptionGetV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateV2(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--entityId", "abc", "--uri", "http://ngsiproxy", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/subscriptions/5f0a44789dd803416ccbf15c"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f0a44789dd803416ccbf15c\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsCreateV2OnlyURI(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--uri", "http://ngsiproxy"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"idPattern":".*"}]},"notification":{"http":{"url":"http://ngsiproxy"}}}`)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/subscriptions/5f0a44789dd803416ccbf15c"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f0a44789dd803416ccbf15c\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsCreateV2DataRaw(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--raw", "--data", `{"description":"TestNotification","subject":{"entities":[{"idPattern":"Alert.*","type":"Alert"}],"condition":{"attrs":[]}},"notification":{"httpCustom":{"url":"http://dev/null","headers":{"fiware-shared-key":"test"}},"attrsFormat":"keyValues"}}`})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"description":"TestNotification","subject":{"entities":[{"idPattern":"Alert.*","type":"Alert"}],"condition":{"attrs":[]}},"notification":{"httpCustom":{"url":"http://dev/null","headers":{"fiware-shared-key":"test"}},"attrsFormat":"keyValues"}}`)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/subscriptions/5f0a44789dd803416ccbf15c"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f0a44789dd803416ccbf15c\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsCreateV2ErrorDataRaw(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--raw", "--data", "@"})

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateV2ErrorsetSubscriptionValuesV2(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z", "--data", "@"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/v2/subscriptions"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateV2ErrorMarshal(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--entityId", "abc", "--url", "http://ngsiproxy", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/v2/subscriptions/5f0a44789dd803416ccbf15c"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--entityId", "abc", "--url", "http://ngsiproxy", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscription"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateV2ErrorStatus(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion", "--entityId", "abc", "--url", "http://ngsiproxy", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte(`{"subject":{"entities":[{"id":"abc"}]},"notification":{"http":{"url":"http://ngsiproxy"}},"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateV2(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSubscriptionsUpdateV2DataRaw(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--raw", "--data", `{"description":"TestNotification","subject":{"entities":[{"idPattern":"Alert.*","type":"Alert"}],"condition":{"attrs":[]}},"notification":{"httpCustom":{"url":"http://dev/null","headers":{"fiware-shared-key":"test"}},"attrsFormat":"keyValues"}}`})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"description":"TestNotification","subject":{"entities":[{"idPattern":"Alert.*","type":"Alert"}],"condition":{"attrs":[]}},"notification":{"httpCustom":{"url":"http://dev/null","headers":{"fiware-shared-key":"test"}},"attrsFormat":"keyValues"}}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	assert.NoError(t, err)

}

func TestSubscriptionsUpdateV2ErrorDataRaw(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--raw", "--data", "@"})

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateV2wAttr(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--wAttrs", "Temp"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"subject":{"condition":{"attrs":["Temp"]}}}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSubscriptionsUpdateV2ErrotsetSubscriptionValuesV2(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscription/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error 2", ngsiErr.Message)
	}
}
func TestSubscriptionsUpdateV2Marshal(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
func TestSubscriptionsUpdateV2ErrotHTTP(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscription/5f0a44789dd803416ccbf15c"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateV2ErrorStatus(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte(`{"throttling":1,"expires":"2020-10-05T00:58:26.929Z"}`)
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " error 5f0a44789dd803416ccbf15c", ngsiErr.Message)
	}
}

func TestSubscriptionsDeleteV2(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDeleteV2(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSubscriptionsDeleteV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscription/5f0a44789dd803416ccbf15c"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDeleteV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsDeleteV2ErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDeleteV2(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error 5f0a44789dd803416ccbf15c", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateV2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--data", "{}", "--entityId", "abc", "--url", "http://ngsiproxy"})

	err := subscriptionsTemplateV2(c, c.Ngsi)

	assert.NoError(t, err)
}

func TestSubscriptionsTemplateV2Pretty(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--pretty", "--data", "{}", "--entityId", "abc", "--url", "http://ngsiproxy"})

	err := subscriptionsTemplateV2(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"subject\": {\n    \"entities\": [\n      {\n        \"id\": \"abc\"\n      }\n    ]\n  },\n  \"notification\": {\n    \"http\": {\n      \"url\": \"http://ngsiproxy\"\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateV2Error(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--url", "http://ngsiproxy", "--data", "@"})

	err := subscriptionsTemplateV2(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateV2ErrorMarshal(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--entityId", "abc", "--url", "http://ngsiproxy"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsTemplateV2(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateV2ErrorPretty(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--pretty", "--data", "{}", "--entityId", "abc", "--url", "http://ngsiproxy"})

	helper.SetJSONIndentError(c.Ngsi)

	err := subscriptionsTemplateV2(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Data(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--data", "{}", "--entityId", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2getAttributes(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--entityId", "abc", "--url", "http://ngsiproxy", "--host", "orion", "--type", "abc", "--get"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}")

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2IdPattern(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--data", "{}", "--idPattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2TypePattern(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--data", "{}", "--typePattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2wAttrs1(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--wAttrs", "abc,def,xyz", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2wAttrs2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--wAttrs", "abc,def,xyz", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2query1(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--query", "abc", "--mq", "def", "--georel", "123", "--geometry", "456", "--coords", "789", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2query2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--query", "abc", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2query3(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--query", "abc", "--mq", "def", "--georel", "123", "--geometry", "456", "--coords", "789", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)
	sub.Subject.Condition = new(subscriptionConditionV2)
	sub.Subject.Condition.Expression = new(subscriptionExpressionV2)

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2url(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}
	sub.Notification = new(subscriptionNotificationV2)
	sub.Notification.HTTP = new(subscriptionHTTPV2)

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2Metadata(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--metadata", "abc"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2headers(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy", "--headers", `{"abc":"123","xyz":"456"}`})

	sub := subscriptionV2{}
	sub.Notification = new(subscriptionNotificationV2)
	sub.Notification.HTTP = new(subscriptionHTTPV2)

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2qs(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy", "--qs", `{"abc":"123","xyz":"456"}`})

	sub := subscriptionV2{}
	sub.Notification = new(subscriptionNotificationV2)
	sub.Notification.HTTPCustom = new(subscriptionHTTPCustomV2)

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2method(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy", "--method", "post", "--payload", "abc", "--nAttrs", "abc,xyz", "--metadata", "abc,xyz"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2exceptAttrs(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy", "--exceptAttrs", "abc,xyz", "--attrsFormat", "abc", "--throttling", "1", "--expires", "1day", "--status", "oneshot"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2expires(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--idPattern", "abc", "--typePattern", "abc", "--url", "http://ngsiproxy", "--exceptAttrs", "abc,xyz", "--attrsFormat", "abc", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z", "--status", "active"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	assert.NoError(t, err)
}

func TestSetSubscriptionValuesV2Error1(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--data", "@", "--entityId", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--data", "{id}", "--entityId", "abc", "--url", "http://ngsiproxy"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'i' looking for beginning of object key string (2) {id}", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error3(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--entityId", "abc", "--url", "http://ngsiproxy", "--host", "orion", "--type", "abc", "--get", "--link", "abc"})

	sub := subscriptionV2{}

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error5(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--entityId", "abc", "--url", "http://ngsiproxy", "--type", "123", "--typePattern", "xyz"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "type or typePattern", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error6(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--headers", "{id}", "--entityId", "abc", "--url", "http://ngsiproxy", "--type", "123"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "err{id}", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error7(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--qs", "{id}", "--entityId", "abc", "--url", "http://ngsiproxy", "--type", "123"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "err{id}", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error8(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--exceptAttrs", "abc", "--nAttrs", "abc", "--entityId", "abc", "--url", "http://ngsiproxy", "--type", "123"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error exceptAttrs or nAttrs", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error9(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--expires", "1", "--entityId", "abc", "--url", "http://ngsiproxy", "--type", "123"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesV2Error10(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--status", "error", "--entityId", "abc", "--url", "http://ngsiproxy", "--type", "123"})

	sub := subscriptionV2{}

	err := setSubscriptionValuesV2(c, c.Ngsi, &sub, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "error: error (active, inactive, oneshot)", ngsiErr.Message)
	}
}

func TestGtAttributesV2Ok1(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--host", "orion", "--type", "abc", "--get"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}")

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	sub := subscriptionV2{}

	err := getAttributesV2(c, c.Ngsi, &sub)

	assert.NoError(t, err)
}

func TestGtAttributesV2Ok2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--host", "orion", "--type", "abc", "--get"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}")

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)
	sub.Subject.Entities = append(sub.Subject.Entities, *new(subscriptionEntityV2))

	err := getAttributesV2(c, c.Ngsi, &sub)

	assert.NoError(t, err)
}

func TestGtAttributesV2NoError(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2"})

	sub := subscriptionV2{}

	err := getAttributesV2(c, c.Ngsi, &sub)

	assert.NoError(t, err)
}

func TestGtAttributesV2ErrorNewClient(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "abc", "--ngsiType", "v2", "--type", "abc", "--get"})

	sub := subscriptionV2{}

	err := getAttributesV2(c, c.Ngsi, &sub)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error host: abc", ngsiErr.Message)
	}
}

func TestGtAttributesV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--type", "abc", "--get"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Path = "/v2/types"
	reqRes.Err = errors.New("http error")

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	sub := subscriptionV2{}

	err := getAttributesV2(c, c.Ngsi, &sub)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestGtAttributesV2ErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--host", "orion", "--type", "abc", "--get"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("{\"attrs\":{\"CO\":{\"types\":[\"Number\"]},\"CO_Level\":{\"types\":[\"Text\"]},\"NO\":{\"types\":[\"Number\"]},\"NO2\":{\"types\":[\"Number\"]},\"NOx\":{\"types\":[\"Number\"]},\"SO2\":{\"types\":[\"Number\"]},\"address\":{\"types\":[\"StructuredValue\"]},\"airQualityIndex\":{\"types\":[\"Number\"]},\"airQualityLevel\":{\"types\":[\"Text\"]},\"dateObserved\":{\"types\":[\"DateTime\",\"Text\"]},\"location\":{\"types\":[\"StructuredValue\",\"geo:json\"]},\"precipitation\":{\"types\":[\"Number\"]},\"refPointOfInterest\":{\"types\":[\"Text\"]},\"relativeHumidity\":{\"types\":[\"Number\"]},\"reliability\":{\"types\":[\"Number\"]},\"source\":{\"types\":[\"Text\",\"URL\"]},\"temperature\":{\"types\":[\"Number\"]},\"windDirection\":{\"types\":[\"Number\"]},\"windSpeed\":{\"types\":[\"Number\"]}},\"count\":18}")

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	sub := subscriptionV2{}
	sub.Subject = new(subscriptionSubjectV2)
	sub.Subject.Condition = new(subscriptionConditionV2)
	sub.Subject.Entities = append(sub.Subject.Entities, *new(subscriptionEntityV2))

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := getAttributesV2(c, c.Ngsi, &sub)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestGtAttributesV2ErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--host", "orion", "--ngsiType", "v2", "--type", "abc", "--get"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/types/abc"
	reqRes.ResBody = []byte("error")

	mock := helper.NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	c.Ngsi.HTTP = mock

	sub := subscriptionV2{}

	err := getAttributesV2(c, c.Ngsi, &sub)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
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
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--items", "description,timessent,lastnotification,lastsuccess,lastsuccesscode,url,expires,status"})

	_, err := checkItems(c)

	assert.NoError(t, err)
}

func TestCheckItemsError(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion", "--items", "abc"})

	_, err := checkItems(c)

	if assert.Error(t, err) {
		actual := err.Error()
		expected := "error: abc in --items"
		assert.Equal(t, expected, actual)
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
	timesSent := int64(10)

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
