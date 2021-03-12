/*
MIT License

Removeright (c) 2020-2021 Kazuhito Suda

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

func TestRemoveV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Thing", "--run"})

	err := remove(c)

	assert.NoError(t, err)
}

func TestRemoveV2AttrNone(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	rawQuery := "attrs=__NONE&limit=100&options=count&type=Thing"
	reqRes1.RawQuery = &rawQuery
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=Thing", "--run"})

	err := remove(c)

	assert.NoError(t, err)
}

func TestRemoveLD(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes3.Path = "/ngsi-ld/v1/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--type=Thing", "--run"})

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
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

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
func TestRemoveErrorType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=ld", "--host=orion"})

	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify entity type", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorTypeEmpty(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=ld", "--type=", "--host=orion"})

	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "no entity type", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorV2Link(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=ld", "--type=Thing", "--host=orion"})

	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "can't specify --link option on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveErrorRemove(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--type=Thing", "--host=orion"})

	err := remove(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveV2TestRun(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "191 entities will be removed. run remove with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRemoveV2Page(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveV2CountZero(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveV2ErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entitie"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveV2ErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRemoveV2ErrorResultCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveV2ErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock

	setJSONDecodeErr(ngsi, 1)

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveV2ErrorOpUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeV2(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveLDTestRun(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "191 entities will be removed. run remove with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRemoveLDPage(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes3.Path = "/ngsi-ld/v1/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveLDCountZero(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes.Path = "/ngsi-ld/v1/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveLDErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entitie"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveLDErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRemoveLDErrorResultCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveLDErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setJSONDecodeErr(ngsi, 1)

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveLDErrorMarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setJSONEncodeErr(ngsi, 2)

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveLDErrorHTTP2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRemoveLDErrorStatus2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	client, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)

	err = removeLD(c, ngsi, client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " ", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
