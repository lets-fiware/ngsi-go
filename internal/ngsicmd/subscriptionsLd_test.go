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
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestSubscriptionsListLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdCount(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagBool(set, "count")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--count"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "6\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(subscriptionLdData)
	reqRes1.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"106"}}
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(subscriptionLdData)
	reqRes2.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes2.ResHeader = http.Header{"Ngsild-Results-Count": []string{"106"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdStatus(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=active"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdQuery(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "query")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--query=FIWARE*"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdJson(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	set.Bool("json", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--json"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := testDataLdRespose
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdJsonPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	setupFlagBool(set, "json,pretty")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--json", "--pretty"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "[\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\",\n    \"type\": \"Subscription\",\n    \"description\": \"FIWARE\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"q\": \"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"normalized\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store002\",\n        \"accept\": \"application/ld+json\"\n      }\n    },\n    \"expires\": \"2021-12-10T11:16:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\",\n    \"type\": \"Subscription\",\n    \"description\": \"Notify me of low stock in Store 001\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"keyValues\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store001\",\n        \"accept\": \"application/json\"\n      }\n    },\n    \"expires\": \"2020-12-09T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\",\n    \"type\": \"Subscription\",\n    \"description\": \"LD Notify me of low stock in Store 003\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"q\": \"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"normalized\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store003\",\n        \"accept\": \"application/ld+json\"\n      }\n    },\n    \"expires\": \"2021-12-10T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\",\n    \"type\": \"Subscription\",\n    \"description\": \"Notify me of low stock in Store 004\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"keyValues\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store004\",\n        \"accept\": \"application/json\"\n      }\n    },\n    \"expires\": \"2020-12-09T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\",\n    \"type\": \"Subscription\",\n    \"description\": \"LD Notify me of low stock in Store 005\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"q\": \"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"normalized\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store005\",\n        \"accept\": \"application/ld+json\"\n      }\n    },\n    \"expires\": \"2021-12-09T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\",\n    \"type\": \"Subscription\",\n    \"description\": \"Notify me of low stock in Store 006\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"keyValues\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store006\",\n        \"accept\": \"application/json\"\n      }\n    },\n    \"expires\": \"2020-12-09T11:06:29.693Z\",\n    \"status\": \"active\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdJsonCount0(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	set.Bool("json", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--json"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := ""
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	set.Bool("verbose", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--verbose"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e  2021-12-10T11:16:29.693Z FIWARE\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f  2020-12-09T11:06:29.693Z Notify me of low stock in Store 001\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940  2021-12-10T11:06:29.693Z LD Notify me of low stock in Store 003\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941  2020-12-09T11:06:29.693Z Notify me of low stock in Store 004\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942  2021-12-09T11:06:29.693Z LD Notify me of low stock in Store 005\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943 active 2020-12-09T11:06:29.693Z Notify me of low stock in Store 006\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdLocaltime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagBool(set, "verbose,localTime")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--verbose", "--localTime"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e  2021-12-10T20:16:29.693+0900 FIWARE\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f  2020-12-09T20:06:29.693+0900 Notify me of low stock in Store 001\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940  2021-12-10T20:06:29.693+0900 LD Notify me of low stock in Store 003\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941  2020-12-09T20:06:29.693+0900 Notify me of low stock in Store 004\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942  2021-12-09T20:06:29.693+0900 LD Notify me of low stock in Store 005\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943 active 2020-12-09T20:06:29.693+0900 Notify me of low stock in Store 006\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ld/subscription"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--status=err"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: err (active, paused, expired)", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ld/subscription"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	setupFlagBool(set, "json")
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json"})
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListLdErrorJsonPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query")
	setupFlagBool(set, "json,pretty")

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{IndentErr: errors.New("json error"), Jsonlib: j}

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--json", "--pretty"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}

}

func TestSubscriptionsListLdErrorItems(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "status,query,items")
	setupFlagBool(set, "verbose")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--verbose", "--items=id"})

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, "error: id in --items", ngsiErr.Message)
		assert.Equal(t, 8, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff"})
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"3ea2e78f675f2d199d3025ff\",\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-01T01:24:01.00Z\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	setupFlagBool(set, "pretty")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff", "--pretty"})
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"id\": \"3ea2e78f675f2d199d3025ff\",\n  \"description\": \"ngsi source subscription\",\n  \"expires\": \"2020-09-01T01:24:01.00Z\"\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdLocalTime(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	setupFlagBool(set, "localTime")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff", "--localTime"})
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"id\":\"3ea2e78f675f2d199d3025ff\",\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-01T10:24:01.00+0900\"}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdSafeString(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,safeString")
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ld/subscription"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff"})
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdErrorSafeString(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	setupFlagString(set, "id,safeString")
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "id,safeString")
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff", "--safeString=on"})
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetLdErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	setupFlagBool(set, "pretty")

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{IndentErr: errors.New("json error"), Jsonlib: j}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--id=3ea2e78f675f2d199d3025ff", "--pretty"})
	client, _ := newClient(ngsi, c, false)
	err := subscriptionGetLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--data={}"})

	err := subscriptionsCreateLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "5f0a44789dd803416ccbf15c\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateLdErrorSetSubscriptionValuesLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--data="})

	err := subscriptionsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateLdErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--data={}"})

	err := subscriptionsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--data={}"})

	err := subscriptionsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateLdErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "data")
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--data={}"})

	err := subscriptionsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateLd(c, ngsi, client)

	assert.NoError(t, err)
}

func TestSubscriptionsUpdateLdErrorSetSubscriptionValuesLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data,id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--data="})
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateLdErrorJSONMarshalEncode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v2/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateLdErrorStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id,throttling,expires")
	set.Bool("get", false, "doc")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-05T00:58:26.929Z"})

	err := subscriptionsUpdateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionsDeleteLd(c, ngsi, client)

	assert.NoError(t, err)
}

func TestSubscriptionsDeleteLdErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ld/subscription/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionsDeleteLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteLdErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "id")
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)
	_ = set.Parse([]string{"--id=5f0a44789dd803416ccbf15c"})

	err := subscriptionsDeleteLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	c := cli.NewContext(app, set, nil)

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"type\":\"Subscription\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdKeyValues(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupFlagBool(set, "keyValues")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--keyValues"})
	err := subscriptionsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"type\":\"Subscription\",\"notification\":{\"format\":\"keyValues\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdArgs(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	setupFlagString(set, "type,uri,query,link,wAttrs,nAttrs,description")
	setupFlagBool(set, "keyValues")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--type=Device", "--uri=http://ngsiproxy", "--query=abc"})
	_ = set.Parse([]string{"--link=ld", "--wAttrs=abc,xyz", "--nAttrs=abc,xyz"})
	_ = set.Parse([]string{"--description=test"})

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"type\":\"Subscription\",\"description\":\"test\",\"entities\":[{\"type\":\"Device\"}],\"watchedAttributes\":[\"abc\",\"xyz\"],\"q\":\"abc\",\"notification\":{\"attributes\":[\"abc\",\"xyz\"],\"endpoint\":{\"uri\":\"http://ngsiproxy\"}},\"@context\":\"https://schema.lab.fiware.org/ld/context\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdArgsPretty(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	setupFlagString(set, "type,uri,query,link,wAttrs,nAttrs,description")
	setupFlagBool(set, "keyValues,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--type=Device", "--uri=http://ngsiproxy", "--query=abc"})
	_ = set.Parse([]string{"--link=ld", "--wAttrs=abc,xyz", "--nAttrs=abc,xyz"})
	_ = set.Parse([]string{"--description=test", "--pretty"})

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"type\": \"Subscription\",\n  \"description\": \"test\",\n  \"entities\": [\n    {\n      \"type\": \"Device\"\n    }\n  ],\n  \"watchedAttributes\": [\n    \"abc\",\n    \"xyz\"\n  ],\n  \"q\": \"abc\",\n  \"notification\": {\n    \"attributes\": [\n      \"abc\",\n      \"xyz\"\n    ],\n    \"endpoint\": {\n      \"uri\": \"http://ngsiproxy\"\n    }\n  },\n  \"@context\": \"https://schema.lab.fiware.org/ld/context\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdErrorUri(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "uri")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--uri=ngsiproxy"})

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "notification url error: ngsiproxy", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateLdErrorSetSubscriptionValuesLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc"})

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateLdErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	c := cli.NewContext(app, set, nil)

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateLdArgsErrorPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	setupFlagString(set, "type,uri,query,link,wAttrs,nAttrs,description")
	setupFlagBool(set, "keyValues,pretty")

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{IndentErr: errors.New("json error"), Jsonlib: j}

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--type=Device", "--uri=http://ngsiproxy", "--query=abc"})
	_ = set.Parse([]string{"--link=ld", "--wAttrs=abc,xyz", "--nAttrs=abc,xyz"})
	_ = set.Parse([]string{"--description=test", "--pretty"})

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLd1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagString(set, "subscriptionId,name,entityId,idPattern,type,wAttrs,query")
	setupFlagBool(set, "active")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--subscriptionId=subsId", "--name=subsName", "--entityId==device001", "--idPattern=.*", "--type=device", "--wAttrs=temperature", "--query=subs*", "--active"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"id\":\"subsId\",\"type\":\"Subscription\",\"name\":\"subsName\",\"entities\":[{\"id\":\"=device001\",\"idPattern\":\".*\",\"type\":\"device\"}],\"watchedAttributes\":[\"temperature\"],\"q\":\"subs*\",\"isActive\":true}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagString(set, "geometry,coords,georel,geoproperty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--geometry=Point", "--coords=[0, 100]", "--georel=near", "--geoproperty=geo"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"geoQ\":{\"geometry\":\"Point\",\"coordinates\":[0,100],\"georel\":\"near\",\"geoproperty\":\"geo\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagInt64(set, "timeInterval")
	setupFlagString(set, "csf")
	setupFlagBool(set, "inactive")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--timeInterval=1", "--csf=abc", "--inactive"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"timeInterval\":1,\"csf\":\"abc\",\"isActive\":false}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdAccept1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagString(set, "accept")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--accept=json"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"notification\":{\"endpoint\":{\"uri\":\"\",\"accept\":\"application/json\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdAccept2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagString(set, "accept")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--accept=ld+json"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"notification\":{\"endpoint\":{\"uri\":\"\",\"accept\":\"application/ld+json\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdAccept3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagString(set, "accept")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--accept=ld"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"notification\":{\"endpoint\":{\"uri\":\"\",\"accept\":\"application/ld+json\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd4(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagString(set, "timeRel,timeAt,endTimeAt,timeProperty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--timeRel=before", "--timeAt=2020-09-24T07:49:56.00Z", "--endTimeAt=2020-09-24T07:49:56.00Z", "--timeProperty=timeProp"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"temporalQ\":{\"timerel\":\"before\",\"timeAt\":\"2020-09-24T07:49:56.00Z\",\"endTimeAt\":\"2020-09-24T07:49:56.00Z\",\"timeproperty\":\"timeProp\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd5(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagInt64(set, "throttling")
	setupFlagString(set, "expires,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-01T01:10:00.00Z", "--link=http://context"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"expires\":\"2020-10-01T01:10:00.00Z\",\"throttling\":1,\"@context\":\"http://context\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd6(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	var s subscriptionLd

	setupFlagInt64(set, "throttling")
	setupFlagString(set, "expires,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--throttling=1", "--expires=2020-10-01T01:10:00.000Z", "--link=http://context"})
	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"expires\":\"2020-10-01T01:10:00.000Z\",\"throttling\":1,\"@context\":\"http://context\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data="})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLdErrorJSONUnMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={}"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLdErrorCoords(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "coords")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--coords=1,100"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "coords: not JSON Array:1,100", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLdErrorActive(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "active,inactive")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--active", "--inactive"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specify both active and inactive options", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLdErrorUri(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "uri")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--uri=ngsiproxy"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "notification url error: ngsiproxy", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLdErrorAccept(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "accept")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--accept=xml"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "unknown param: xml", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestSetSubscriptionValuesLdErrorExpires(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "expires")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--expires=day"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error day", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLdErrorTimeRel(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "timeRel")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--timeRel=current"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "unknown param: current", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSetSubscriptionValuesLdErrorLink(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=context"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "context not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestToLocaltimeLd(t *testing.T) {
	var s subscriptionLd

	s.Expires = "2020-10-01T01:10:00.000Z"
	toLocaltimeLd(&s)

	assert.Equal(t, "2020-10-01T10:10:00.000+0900", s.Expires)
}

func TestToLocaltimeLd2(t *testing.T) {
	var s subscriptionLd
	s.Notification = new(notificationParamsLd)

	s.Expires = "2020-10-01T01:10:00.000Z"
	s.Notification.LastNotification = "2020-10-02T01:10:00.00Z"
	s.Notification.LastSuccess = "2020-10-03T01:10:00.00Z"
	s.Notification.LastFailure = "2020-10-04T01:10:00.000Z"

	toLocaltimeLd(&s)

	assert.Equal(t, "2020-10-01T10:10:00.000+0900", s.Expires)
	assert.Equal(t, "2020-10-02T10:10:00.00+0900", s.Notification.LastNotification)
	assert.Equal(t, "2020-10-03T10:10:00.00+0900", s.Notification.LastSuccess)
	assert.Equal(t, "2020-10-04T10:10:00.000+0900", s.Notification.LastFailure)
}

func TestCheckItemLd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--items=description,timessent,lastnotification,lastsuccess"})

	actual, err := checkItemsLd(c)

	if assert.NoError(t, err) {
		expected := []string{"id", "description", "timessent", "lastnotification", "lastsuccess"}
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCheckItemLdError(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--items=id"})

	_, err := checkItemsLd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: id in --items", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSprintItemsLd(t *testing.T) {
	var s subscriptionLd
	s.Notification = new(notificationParamsLd)
	s.Notification.Endpoint = new(endpointLd)

	var time int64

	s.Status = "active"
	s.Notification.TimesSent = &time
	s.Notification.Status = "active"
	s.Notification.Endpoint.URI = "http://ngsiproxy"

	items := []string{"id", "description", "timessent", "lastnotification", "lastsuccess", "notificationstatus", "uri", "expires", "status"}

	actual := sprintItemsLd(&s, items)

	assert.Equal(t, "  0   active http://ngsiproxy  active", actual)
}

var subscriptionLdData = `[
	{
		"id": "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e",
		"type": "Subscription",
		"description": "FIWARE",
		"entities": [
		  {
			"type": "Shelf"
		  }
		],
		"watchedAttributes": [
		  "numberOfItems"
		],
		"q": "https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002",
		"notification": {
		  "attributes": [
			"numberOfItems",
			"stocks",
			"locatedIn"
		  ],
		  "format": "normalized",
		  "endpoint": {
			"uri": "http://tutorial:3000/subscription/low-stock-store002",
			"accept": "application/ld+json"
		  }
		},
		"expires": "2021-12-10T11:16:29.693Z"
	  },
	  {
		"id": "urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f",
		"type": "Subscription",
		"description": "Notify me of low stock in Store 001",
		"entities": [
		  {
			"type": "Shelf"
		  }
		],
		"watchedAttributes": [
		  "numberOfItems"
		],
		"notification": {
		  "attributes": [
			"numberOfItems",
			"stocks",
			"locatedIn"
		  ],
		  "format": "keyValues",
		  "endpoint": {
			"uri": "http://tutorial:3000/subscription/low-stock-store001",
			"accept": "application/json"
		  }
		},
		"expires": "2020-12-09T11:06:29.693Z"
	  },
	  {
		"id": "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940",
		"type": "Subscription",
		"description": "LD Notify me of low stock in Store 003",
		"entities": [
		  {
			"type": "Shelf"
		  }
		],
		"watchedAttributes": [
		  "numberOfItems"
		],
		"q": "https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002",
		"notification": {
		  "attributes": [
			"numberOfItems",
			"stocks",
			"locatedIn"
		  ],
		  "format": "normalized",
		  "endpoint": {
			"uri": "http://tutorial:3000/subscription/low-stock-store003",
			"accept": "application/ld+json"
		  }
		},
		"expires": "2021-12-10T11:06:29.693Z"
	  },
	  {
		"id": "urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941",
		"type": "Subscription",
		"description": "Notify me of low stock in Store 004",
		"entities": [
		  {
			"type": "Shelf"
		  }
		],
		"watchedAttributes": [
		  "numberOfItems"
		],
		"notification": {
		  "attributes": [
			"numberOfItems",
			"stocks",
			"locatedIn"
		  ],
		  "format": "keyValues",
		  "endpoint": {
			"uri": "http://tutorial:3000/subscription/low-stock-store004",
			"accept": "application/json"
		  }
		},
		"expires": "2020-12-09T11:06:29.693Z"
	  },
	  {
		"id": "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942",
		"type": "Subscription",
		"description": "LD Notify me of low stock in Store 005",
		"entities": [
		  {
			"type": "Shelf"
		  }
		],
		"watchedAttributes": [
		  "numberOfItems"
		],
		"q": "https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002",
		"notification": {
		  "attributes": [
			"numberOfItems",
			"stocks",
			"locatedIn"
		  ],
		  "format": "normalized",
		  "endpoint": {
			"uri": "http://tutorial:3000/subscription/low-stock-store005",
			"accept": "application/ld+json"
		  }
		},
		"expires": "2021-12-09T11:06:29.693Z"
	  },
	  {
		"id": "urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943",
		"type": "Subscription",
		"description": "Notify me of low stock in Store 006",
		"entities": [
		  {
			"type": "Shelf"
		  }
		],
		"watchedAttributes": [
		  "numberOfItems"
		],
		"notification": {
		  "attributes": [
			"numberOfItems",
			"stocks",
			"locatedIn"
		  ],
		  "format": "keyValues",
		  "endpoint": {
			"uri": "http://tutorial:3000/subscription/low-stock-store006",
			"accept": "application/json"
		  }
		},
		"expires": "2020-12-09T11:06:29.693Z",
		"status": "active"
	  }
  ]`

var testDataLdRespose = "[{\"id\":\"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\",\"type\":\"Subscription\",\"description\":\"FIWARE\",\"entities\":[{\"type\":\"Shelf\"}],\"watchedAttributes\":[\"numberOfItems\"],\"q\":\"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\"notification\":{\"attributes\":[\"numberOfItems\",\"stocks\",\"locatedIn\"],\"format\":\"normalized\",\"endpoint\":{\"uri\":\"http://tutorial:3000/subscription/low-stock-store002\",\"accept\":\"application/ld+json\"}},\"expires\":\"2021-12-10T11:16:29.693Z\"},{\"id\":\"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\",\"type\":\"Subscription\",\"description\":\"Notify me of low stock in Store 001\",\"entities\":[{\"type\":\"Shelf\"}],\"watchedAttributes\":[\"numberOfItems\"],\"notification\":{\"attributes\":[\"numberOfItems\",\"stocks\",\"locatedIn\"],\"format\":\"keyValues\",\"endpoint\":{\"uri\":\"http://tutorial:3000/subscription/low-stock-store001\",\"accept\":\"application/json\"}},\"expires\":\"2020-12-09T11:06:29.693Z\"},{\"id\":\"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\",\"type\":\"Subscription\",\"description\":\"LD Notify me of low stock in Store 003\",\"entities\":[{\"type\":\"Shelf\"}],\"watchedAttributes\":[\"numberOfItems\"],\"q\":\"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\"notification\":{\"attributes\":[\"numberOfItems\",\"stocks\",\"locatedIn\"],\"format\":\"normalized\",\"endpoint\":{\"uri\":\"http://tutorial:3000/subscription/low-stock-store003\",\"accept\":\"application/ld+json\"}},\"expires\":\"2021-12-10T11:06:29.693Z\"},{\"id\":\"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\",\"type\":\"Subscription\",\"description\":\"Notify me of low stock in Store 004\",\"entities\":[{\"type\":\"Shelf\"}],\"watchedAttributes\":[\"numberOfItems\"],\"notification\":{\"attributes\":[\"numberOfItems\",\"stocks\",\"locatedIn\"],\"format\":\"keyValues\",\"endpoint\":{\"uri\":\"http://tutorial:3000/subscription/low-stock-store004\",\"accept\":\"application/json\"}},\"expires\":\"2020-12-09T11:06:29.693Z\"},{\"id\":\"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\",\"type\":\"Subscription\",\"description\":\"LD Notify me of low stock in Store 005\",\"entities\":[{\"type\":\"Shelf\"}],\"watchedAttributes\":[\"numberOfItems\"],\"q\":\"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\"notification\":{\"attributes\":[\"numberOfItems\",\"stocks\",\"locatedIn\"],\"format\":\"normalized\",\"endpoint\":{\"uri\":\"http://tutorial:3000/subscription/low-stock-store005\",\"accept\":\"application/ld+json\"}},\"expires\":\"2021-12-09T11:06:29.693Z\"},{\"id\":\"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\",\"type\":\"Subscription\",\"description\":\"Notify me of low stock in Store 006\",\"entities\":[{\"type\":\"Shelf\"}],\"watchedAttributes\":[\"numberOfItems\"],\"notification\":{\"attributes\":[\"numberOfItems\",\"stocks\",\"locatedIn\"],\"format\":\"keyValues\",\"endpoint\":{\"uri\":\"http://tutorial:3000/subscription/low-stock-store006\",\"accept\":\"application/json\"}},\"expires\":\"2020-12-09T11:06:29.693Z\",\"status\":\"active\"}]\n"
