/*
MIT License

Removeright (c) 2020 Kazuhito Suda

This file is part of NGSI Go

https://github.com/lets-fiware/ngsi-go

Permission is hereby granted, free of charge, to any person obtaining a remove
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, remove, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above removeright notice and this permission notice shall be included in all
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

func TestRemove(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"
	mock := NewMockHTTP()
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	err := remove(c)

	assert.NoError(t, err)
}

func TestRemovePage(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"
	mock := NewMockHTTP()
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	err := remove(c)

	assert.NoError(t, err)
}

func TestRemoveCountZero(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	err := remove(c)

	assert.NoError(t, err)
}

func TestRemoveErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorNewClient(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion"})
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorIsNgsiLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion-ld", "https://orion-ld", "ld")

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorRunFlag(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "run remove with --run option", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entitie"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorResultCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"\": invalid syntax", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"
	mock := NewMockHTTP()
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorOpUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupAddBroker(t, ngsi, "orion", "https://orion", "v2")

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"
	mock := NewMockHTTP()
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})
	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
