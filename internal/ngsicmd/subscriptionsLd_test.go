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
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestSubscriptionsListLd(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdCount(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--count"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "6\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdPage(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(subscriptionLdData)
	reqRes1.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"106"}}

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(subscriptionLdData)
	reqRes2.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes2.ResHeader = http.Header{"Ngsild-Results-Count": []string{"106"}}

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdStatus(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--status", "active"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdQuery(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--query", "FIWARE"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdCountZero(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdJson(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := testDataLdRespose
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdJsonPretty(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "[\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e\",\n    \"type\": \"Subscription\",\n    \"description\": \"FIWARE\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"q\": \"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"normalized\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store002\",\n        \"accept\": \"application/ld+json\"\n      }\n    },\n    \"expires\": \"2021-12-10T11:16:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f\",\n    \"type\": \"Subscription\",\n    \"description\": \"Notify me of low stock in Store 001\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"keyValues\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store001\",\n        \"accept\": \"application/json\"\n      }\n    },\n    \"expires\": \"2020-12-09T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940\",\n    \"type\": \"Subscription\",\n    \"description\": \"LD Notify me of low stock in Store 003\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"q\": \"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"normalized\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store003\",\n        \"accept\": \"application/ld+json\"\n      }\n    },\n    \"expires\": \"2021-12-10T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941\",\n    \"type\": \"Subscription\",\n    \"description\": \"Notify me of low stock in Store 004\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"keyValues\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store004\",\n        \"accept\": \"application/json\"\n      }\n    },\n    \"expires\": \"2020-12-09T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942\",\n    \"type\": \"Subscription\",\n    \"description\": \"LD Notify me of low stock in Store 005\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"q\": \"https://fiware.github.io/tutorials.Step-by-Step/schema/numberOfItems<10;https://fiware.github.io/tutorials.Step-by-Step/schema/locatedIn==urn:ngsi-ld:Building:store002\",\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"normalized\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store005\",\n        \"accept\": \"application/ld+json\"\n      }\n    },\n    \"expires\": \"2021-12-09T11:06:29.693Z\"\n  },\n  {\n    \"id\": \"urn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943\",\n    \"type\": \"Subscription\",\n    \"description\": \"Notify me of low stock in Store 006\",\n    \"entities\": [\n      {\n        \"type\": \"Shelf\"\n      }\n    ],\n    \"watchedAttributes\": [\n      \"numberOfItems\"\n    ],\n    \"notification\": {\n      \"attributes\": [\n        \"numberOfItems\",\n        \"stocks\",\n        \"locatedIn\"\n      ],\n      \"format\": \"keyValues\",\n      \"endpoint\": {\n        \"uri\": \"http://tutorial:3000/subscription/low-stock-store006\",\n        \"accept\": \"application/json\"\n      }\n    },\n    \"expires\": \"2020-12-09T11:06:29.693Z\",\n    \"status\": \"active\"\n  }\n]\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdJsonCount0(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := ""
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdVerbose(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--verbose"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e  2021-12-10T11:16:29.693Z FIWARE\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f  2020-12-09T11:06:29.693Z Notify me of low stock in Store 001\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940  2021-12-10T11:06:29.693Z LD Notify me of low stock in Store 003\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941  2020-12-09T11:06:29.693Z Notify me of low stock in Store 004\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942  2021-12-09T11:06:29.693Z LD Notify me of low stock in Store 005\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943 active 2020-12-09T11:06:29.693Z Notify me of low stock in Store 006\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdLocaltime(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--verbose", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "urn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c93e  2021-12-10T20:16:29.693+0900 FIWARE\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c93f  2020-12-09T20:06:29.693+0900 Notify me of low stock in Store 001\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c940  2021-12-10T20:06:29.693+0900 LD Notify me of low stock in Store 003\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c941  2020-12-09T20:06:29.693+0900 Notify me of low stock in Store 004\nurn:ngsi-ld:Subscription:5fcf5b65f6c9a661e958c942  2021-12-09T20:06:29.693+0900 LD Notify me of low stock in Store 005\nurn:ngsi-ld:Subscription:5fcf5e3ff6c9a661e958c943 active 2020-12-09T20:06:29.693+0900 Notify me of low stock in Store 006\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsListLdErrorStatus(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--status", "err"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ld/subscription"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: err (active, paused, expired)", ngsiErr.Message)
	}
}

func TestSubscriptionsListLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ld/subscription"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsListLdErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestSubscriptionsListLdErrorResultsCount(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	}
}

func TestSubscriptionsListLdErrorJSONUnmarshal(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	}
}

func TestSubscriptionsListLdErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--json"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, "json error", ngsiErr.Message)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	}
}

func TestSubscriptionsListLdErrorJsonPretty(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--json", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}

}

func TestSubscriptionsListLdErrorItems(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--verbose", "--items", "id"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsListLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, "error: id in --items", ngsiErr.Message)
		assert.Equal(t, 8, ngsiErr.ErrNo)
	}
}

func TestSubscriptionsGetLd(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"3ea2e78f675f2d199d3025ff\",\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-01T01:24:01.00Z\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsGetLdPretty(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"id\": \"3ea2e78f675f2d199d3025ff\",\n  \"description\": \"ngsi source subscription\",\n  \"expires\": \"2020-09-01T01:24:01.00Z\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsGetLdLocalTime(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff", "--localTime"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"id\":\"3ea2e78f675f2d199d3025ff\",\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-01T10:24:01.00+0900\"}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsGetLdSafeString(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsGetLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ld/subscription"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsGetLdErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	}
}

func TestSubscriptionsGetLdErrorSafeString(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsGetLdErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff", "--safeString", "on"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsGetLdErrorPretty(t *testing.T) {
	c := setupTest([]string{"get", "subscription", "--host", "orion-ld", "--id", "3ea2e78f675f2d199d3025ff", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"id": "3ea2e78f675f2d199d3025ff", "description": "ngsi source subscription", "expires": "2020-09-01T01:24:01.00Z"}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/3ea2e78f675f2d199d3025ff"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := subscriptionGetLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateLd(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	reqRes.ResHeader = http.Header{"Location": []string{"/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"}}

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateLd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "5f0a44789dd803416ccbf15c\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsCreateLdErrorSetSubscriptionValuesLd(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion-ld", "--data", "@"})

	err := subscriptionsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateLdErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion-ld", "--data", "{}"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusCreated
	reqRes.Path = "/ngsi-ld/subscriptions"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsCreateLdErrorStatus(t *testing.T) {
	c := setupTest([]string{"create", "subscription", "--host", "orion-ld", "--data", "{}"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsCreateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateLd(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateLd(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSubscriptionsUpdateLdErrorSetSubscriptionValuesLd(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c", "--data", "@"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateLdErrorJSONMarshalEncode(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsUpdateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v2/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsUpdateLdErrorStatus(t *testing.T) {
	c := setupTest([]string{"update", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c", "--throttling", "1", "--expires", "2020-10-05T00:58:26.929Z"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte(`{"type":"Subscription","expires":"2020-10-05T00:58:26.929Z","throttling":1}`)
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsUpdateLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " error 5f0a44789dd803416ccbf15c", ngsiErr.Message)
	}
}

func TestSubscriptionsDeleteLd(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDeleteLd(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSubscriptionsDeleteLdErrorHTTP(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ld/subscription/5f0a44789dd803416ccbf15c"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDeleteLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestSubscriptionsDeleteLdErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"delete", "subscription", "--host", "orion-ld", "--id", "5f0a44789dd803416ccbf15c"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/subscriptions/5f0a44789dd803416ccbf15c"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := subscriptionsDeleteLd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error 5f0a44789dd803416ccbf15c", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateLd(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld"})

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"Subscription\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdKeyValues(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--keyValues"})

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"Subscription\",\"notification\":{\"format\":\"keyValues\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdArgs(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--type", "Device", "--uri", "http://ngsiproxy", "--query", "abc", "--wAttrs", "abc,xyz", "--nAttrs", "abc,xyz", "--description", "test"})

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"type\":\"Subscription\",\"description\":\"test\",\"entities\":[{\"type\":\"Device\"}],\"watchedAttributes\":[\"abc\",\"xyz\"],\"q\":\"abc\",\"notification\":{\"attributes\":[\"abc\",\"xyz\"],\"endpoint\":{\"uri\":\"http://ngsiproxy\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdArgsPretty(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--type", "Device", "--uri", "http://ngsiproxy", "--query", "abc", "--wAttrs", "abc,xyz", "--nAttrs", "abc,xyz", "--description", "test", "--pretty"})

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"type\": \"Subscription\",\n  \"description\": \"test\",\n  \"entities\": [\n    {\n      \"type\": \"Device\"\n    }\n  ],\n  \"watchedAttributes\": [\n    \"abc\",\n    \"xyz\"\n  ],\n  \"q\": \"abc\",\n  \"notification\": {\n    \"attributes\": [\n      \"abc\",\n      \"xyz\"\n    ],\n    \"endpoint\": {\n      \"uri\": \"http://ngsiproxy\"\n    }\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestSubscriptionsTemplateLdErrorUri(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--uri", "ngsiproxy"})

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "notification url error: ngsiproxy", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateLdErrorSetSubscriptionValuesLd(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--context", "abc"})

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateLdErrorJSONMarshal(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--context", "{}"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSubscriptionsTemplateLdArgsErrorPretty(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--type", "Device", "--uri", "http://ngsiproxy", "--query", "abc", "--link", "ld", "--wAttrs", "abc,xyz", "--nAttrs", "abc,xyz", "--description", "test", "--pretty"})

	helper.SetJSONIndentError(c.Ngsi)

	err := subscriptionsTemplateLd(c, c.Ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLd1(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--subscriptionId", "subsId", "--name", "subsName", "--entityId", "device001", "--idPattern", ".*", "--type", "device", "--wAttrs", "temperature", "--query", "subs*", "--active"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"id\":\"subsId\",\"type\":\"Subscription\",\"name\":\"subsName\",\"entities\":[{\"id\":\"device001\",\"idPattern\":\".*\",\"type\":\"device\"}],\"watchedAttributes\":[\"temperature\"],\"q\":\"subs*\",\"isActive\":true}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--geometry", "Point", "--coords", "[0, 100]", "--georel", "near", "--geoproperty", "geo"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"geoQ\":{\"geometry\":\"Point\",\"coordinates\":[0,100],\"georel\":\"near\",\"geoproperty\":\"geo\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd3(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--timeInterval", "1", "--csf", "abc", "--inactive"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"timeInterval\":1,\"csf\":\"abc\",\"isActive\":false}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdAccept1(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--accept", "json"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"notification\":{\"endpoint\":{\"uri\":\"\",\"accept\":\"application/json\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdAccept2(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--accept", "ld+json"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"notification\":{\"endpoint\":{\"uri\":\"\",\"accept\":\"application/ld+json\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdAccept3(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--accept", "ld"})
	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"notification\":{\"endpoint\":{\"uri\":\"\",\"accept\":\"application/ld+json\"}}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd4(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--timeRel", "before", "--timeAt", "2020-09-24T07:49:56.00Z", "--endTimeAt", "2020-09-24T07:49:56.00Z", "--timeProperty", "timeProp"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"temporalQ\":{\"timerel\":\"before\",\"timeAt\":\"2020-09-24T07:49:56.00Z\",\"endTimeAt\":\"2020-09-24T07:49:56.00Z\",\"timeproperty\":\"timeProp\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd5(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--throttling", "1", "--expires", "2020-10-01T01:10:00.00Z"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"expires\":\"2020-10-01T01:10:00.00Z\",\"throttling\":1}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLd6(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--throttling", "1", "--expires", "2020-10-01T01:10:00.000Z"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"expires\":\"2020-10-01T01:10:00.000Z\",\"throttling\":1}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdContext(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--context", "[\"http://context\"]"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	b, _ := json.Marshal(s)
	actual := string(b)

	if assert.NoError(t, err) {
		expected := "{\"type\":\"Subscription\",\"@context\":[\"http://context\"]}"
		assert.Equal(t, expected, actual)
	}
}

func TestSetSubscriptionValuesLdErrorReadAll(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--data", "@"})

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLdErrorJSONUnMarshal(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--data", "{}"})

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLdErrorCoords(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--coords", "1,100"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "coords: not JSON Array:1,100", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLdErrorActive(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--active", "--inactive"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "cannot specify both active and inactive options", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLdErrorUri(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--uri", "ngsiproxy"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "notification url error: ngsiproxy", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLdErrorAccept(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--accept", "xml"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "unknown param: xml", ngsiErr.Message)
	}
}
func TestSetSubscriptionValuesLdErrorExpires(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--expires", "day"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error day", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLdErrorTimeRel(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--timeRel", "current"})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "unknown param: current", ngsiErr.Message)
	}
}

func TestSetSubscriptionValuesLdErrorContext(t *testing.T) {
	c := setupTest([]string{"template", "subscription", "--ngsiType", "ld", "--context", "[\"http://context\""})

	var s subscriptionLd

	err := setSubscriptionValuesLd(c, c.Ngsi, &s, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
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
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--items", "description,timessent,lastnotification,lastsuccess"})

	actual, err := checkItemsLd(c)

	if assert.NoError(t, err) {
		expected := []string{"id", "description", "timessent", "lastnotification", "lastsuccess"}
		assert.Equal(t, expected, actual)
	}
}

func TestCheckItemLdError(t *testing.T) {
	c := setupTest([]string{"list", "subscriptions", "--host", "orion-ld", "--items", "id"})

	_, err := checkItemsLd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error: id in --items", ngsiErr.Message)
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
