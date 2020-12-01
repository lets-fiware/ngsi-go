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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestSubscriptionssubscriptionsListLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"6"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListLdCount(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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

func TestSubscriptionssubscriptionsListLdPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(subscriptionLdData)
	reqRes1.Path = "/ngsi-ld/v1/subscriptions"
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"106"}}
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(subscriptionLdData)
	reqRes2.Path = "/ngsi-ld/v1/subscriptions"
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
		expected := "3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n3ea2e78f675f2d199d3025ff\n5f64060ef6752d199d302600\n1f32db4bf6752d199d302601\n3978fabd87752d199d302602\n9f6c254ac4a6068bb276774e\n4f6c2576c4a6068bb276774f\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListLdCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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

func TestSubscriptionssubscriptionsListLdJson(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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

func TestSubscriptionssubscriptionsListLdJsonCount0(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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

func TestSubscriptionssubscriptionsListLdVerbose(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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
		expected := "3ea2e78f675f2d199d3025ff ngsi source subscription\n5f64060ef6752d199d302600 ngsi source subscription\n1f32db4bf6752d199d302601 ngsi source subscription\n3978fabd87752d199d302602 ngsi source subscription\n9f6c254ac4a6068bb276774e ngsi source subscription\n4f6c2576c4a6068bb276774f FIWARE\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListLdErrorHTTP(t *testing.T) {
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
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListLdErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListLdErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsListLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListLdErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsListLdErrorJSONMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	setupFlagBool(set, "json")
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(subscriptionLdData)
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
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
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsGetLd(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
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
		expected := "{}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionssubscriptionsGetLdSafeString(t *testing.T) {
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

func TestSubscriptionssubscriptionsGetLdErrorHTTP(t *testing.T) {
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

func TestSubscriptionssubscriptionsGetLdErrorHTTPStatus(t *testing.T) {
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

func TestSubscriptionssubscriptionsGetLdErrorSafeString(t *testing.T) {
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

func TestSubscriptionsCreateLdErrorReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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

	err := subscriptionsCreateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
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
		assert.Equal(t, 2, ngsiErr.ErrNo)
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
		assert.Equal(t, 3, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	c := cli.NewContext(app, set, nil)
	client, _ := newClient(ngsi, c, false)

	err := subscriptionsUpdateLd(c, ngsi, client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "not yet implemented", ngsiErr.Message)
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
		expected := "{\"description\":\"description\",\"entities\":[{\"type\":\"Template\"}],\"notification\":{\"attributes\":[\"attribute\"],\"endpoint\":{\"accept\":\"application/ld+json\",\"uri\":\"http://template\"},\"format\":\"normalized\"},\"type\":\"Subscription\",\"watchedAttributes\":[\"watchedAttribute\"]}"
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
		expected := "{\"description\":\"description\",\"entities\":[{\"type\":\"Template\"}],\"notification\":{\"attributes\":[\"attribute\"],\"endpoint\":{\"accept\":\"application/ld+json\",\"uri\":\"http://template\"},\"format\":\"keyValues\"},\"type\":\"Subscription\",\"watchedAttributes\":[\"watchedAttribute\"]}"
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
		expected := "{\"@Context\":\"https://schema.lab.fiware.org/ld/context\",\"description\":\"test\",\"entities\":[{\"type\":\"Device\"}],\"notification\":{\"attributes\":[\"abc\",\"xyz\"],\"endpoint\":{\"accept\":\"application/ld+json\",\"uri\":\"http://ngsiproxy\"},\"format\":\"normalized\"},\"q\":\"abc\",\"type\":\"Subscription\",\"watchedAttributes\":[\"abc\",\"xyz\"]}"
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

func TestSubscriptionsTemplateLdErrorLink(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc"})

	err := subscriptionsTemplateLd(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
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
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

var subscriptionLdData = `[
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

var testDataLdRespose = "[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-01T01:24:01.00Z\",\"id\":\"3ea2e78f675f2d199d3025ff\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastFailure\":\"2020-09-01T07:42:07.00Z\",\"lastFailureReason\":\"Timeout was reached\",\"lastNotification\":\"2020-09-01T07:43:00.00Z\",\"lastSuccess\":\"2020-09-01T07:43:04.00Z\",\"lastSuccessCode\":204,\"onlyChangedAttrs\":false,\"timesSent\":3406},\"status\":\"expired\",\"subject\":{\"condition\":{\"attrs\":[\"observed\",\"location\"]},\"entities\":[{\"idPattern\":\".*\"}]}},{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-16T03:57:49.00Z\",\"id\":\"5f64060ef6752d199d302600\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-16T03:40:27.00Z\",\"lastSuccess\":\"2020-09-16T03:40:28.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":27},\"status\":\"expired\",\"subject\":{\"condition\":{\"attrs\":[\"dateRetrieved\"]},\"entities\":[{\"idPattern\":\".*\",\"type\":\"WeatherObserved\"}]}},{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-16T04:03:05.00Z\",\"id\":\"1f32db4bf6752d199d302601\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastFailure\":\"2020-09-16T03:40:06.00Z\",\"lastFailureReason\":\"Timeout was reached\",\"lastNotification\":\"2020-09-16T04:03:00.00Z\",\"lastSuccess\":\"2020-09-16T04:03:04.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":3408},\"status\":\"expired\",\"subject\":{\"condition\":{\"attrs\":[\"observed\",\"location\"]},\"entities\":[{\"idPattern\":\".*\"}]}},{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-16T04:03:07.00Z\",\"id\":\"3978fabd87752d199d302602\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-16T04:00:13.00Z\",\"lastSuccess\":\"2020-09-16T04:00:13.00Z\",\"lastSuccessCode\":204,\"onlyChangedAttrs\":false,\"timesSent\":10},\"status\":\"expired\",\"subject\":{\"condition\":{\"attrs\":[\"dateRetrieved\"]},\"entities\":[{\"idPattern\":\".*\",\"type\":\"WeatherObserved\"}]}},{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}},{\"description\":\"FIWARE\",\"expires\":\"2020-09-24T07:49:56.00Z\",\"id\":\"4f6c2576c4a6068bb276774f\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:40:26.00Z\",\"lastSuccess\":\"2020-09-24T07:40:26.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":278},\"status\":\"active\",\"subject\":{\"condition\":{\"attrs\":[\"dateRetrieved\"]},\"entities\":[{\"idPattern\":\".*\",\"type\":\"WeatherObserved\"}]}}]\n"
