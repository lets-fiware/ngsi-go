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

// initCmd() Error: no host
func TestSubscriptionsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := subscriptionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := subscriptionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsListErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := subscriptionsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := subscriptionGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionGetErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := subscriptionGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsGetErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := subscriptionGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := subscriptionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := subscriptionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCreateErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data="})
	err := subscriptionsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := subscriptionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := subscriptionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsUpdateErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := subscriptionsUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := subscriptionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := subscriptionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsDeleteErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := subscriptionsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateNgsiTypeV2(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "ngsiType")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--ngsiType=v2"})
	err := subscriptionsTemplate(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateNgsiTypeLd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "ngsiType")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--ngsiType=ld"})
	err := subscriptionsTemplate(c)

	assert.NoError(t, err)
}

func TestSubscriptionsTemplateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := subscriptionsTemplate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateErrorNgsiTypeNotFound(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := subscriptionsTemplate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Required ngsiType not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsTemplateErrorNgsiType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "ngsiType")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--ngsiType=v1"})
	err := subscriptionsTemplate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error v1", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCountErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := subscriptionsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCountErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--link=abc"})
	err := subscriptionsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCountErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2//subscription"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, ngsiErr.Message, "url error")
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCountErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := subscriptionsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCountErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCountV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/subscriptions"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"12"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := subscriptionsCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "12\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestSubscriptionsCountLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/subscriptions"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"21"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := subscriptionsCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "21\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
