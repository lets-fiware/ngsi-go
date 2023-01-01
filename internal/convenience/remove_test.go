/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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

package convenience

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestRemoveV2(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := remove(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRemoveV2SkipForwarding(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--skipForwarding", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	q := "attrs=__NONE&limit=100&options=count%2CskipForwarding&type=Thing"
	reqRes1.RawQuery = &q
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := remove(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRemoveV1(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run", "--ngsiV1"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"200","reasonPhrase":"OK"}}]}`)
	reqRes2.Path = "/v1/updateContext"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := remove(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRemoveV2AttrNone(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	rawQuery := "attrs=__NONE&limit=100&options=count&type=Thing"
	reqRes1.RawQuery = &rawQuery

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := remove(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRemoveLD(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes3.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := remove(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestRemoveErrorV2Link(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--link", "ld"})

	err := remove(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "can't specify --link option on NGSIv2", ngsiErr.Message)
	}
}

func TestRemoveErrorRemove(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes.Path = "/entities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := remove(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRemoveV2TestRun(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "191 entities will be removed. run remove with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRemoveV2Page(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveV2CountZero(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveV2ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entitie"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRemoveV2ErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	}
}

func TestRemoveV2ErrorResultCount(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestRemoveV2ErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRemoveV2ErrorOpUpdate(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.Path = "/v2/op/update"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}
func TestRemoveV2ErrorOpUpdateStatus(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ReqData = []byte("{\"actionType\":\"delete\",\"entities\":[{\"description\":\"ngsi source subscription\",\"expires\":\"2020-09-24T07:49:13.00Z\",\"id\":\"9f6c254ac4a6068bb276774e\",\"notification\":{\"attrsFormat\":\"keyValues\",\"http\":{\"url\":\"https://ngsiproxy\"},\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"timesSent\":28},\"status\":\"inactive\",\"subject\":{\"condition\":{\"attrs\":[\"dateObserved\"]},\"entities\":[{\"idPattern\":\".*\"}]}}]}")
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes2.Path = "/v2/op/update"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[{\"id\":\"9f6c254ac4a6068bb276774e\",\"description\":\"ngsi source subscription\",\"subject\":{\"entities\":[{\"idPattern\":\".*\"}],\"condition\":{\"attrs\":[\"dateObserved\"]}},\"notification\":{\"timesSent\":28,\"lastNotification\":\"2020-09-24T07:30:02.00Z\",\"lastSuccess\":\"2020-09-24T07:30:02.00Z\",\"lastSuccessCode\":404,\"onlyChangedAttrs\":false,\"http\":{\"url\":\"https://ngsiproxy\"},\"attrsFormat\":\"keyValues\"},\"expires\":\"2020-09-24T07:49:13.00Z\",\"status\":\"inactive\"}]\n")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := removeV2(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}
func TestRemoveLDTestRun(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "191 entities will be removed. run remove with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRemoveLDPage(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes3.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveLDCountZero(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveLDErrorHTTP(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRemoveLDErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/ngsi-ld/v1/entities"
	reqRes.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestRemoveLDErrorResultCount(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestRemoveLDErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRemoveLDErrorMarshal(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRemoveLDErrorHTTP2(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Err = errors.New("error")
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestRemoveLDErrorStatus2(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion-ld", "--type", "Thing", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"TemperatureSensor"},{"id":"urn:ngsi-ld:TemperatureSensor:003","type":"TemperatureSensor"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"191"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ReqData = []byte(`["urn:ngsi-ld:TemperatureSensor:001","urn:ngsi-ld:TemperatureSensor:002","urn:ngsi-ld:TemperatureSensor:003"]`)
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/delete"
	reqRes2.ResBody = []byte("error")

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := removeLD(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " error", ngsiErr.Message)
	}
}

func TestRemoveV1TestRun(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3 entities will be removed. run remove with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestRemoveV1Page(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"200","reasonPhrase":"OK"}}]}`)
	reqRes2.Path = "/v1/updateContext"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveV1CountZero(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	assert.NoError(t, err)
}

func TestRemoveV1ErrorHTTP(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entitie"
	reqRes.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorHTTPStatus(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
	}
}

func TestRemoveV1ErrorResultCount(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ResultsCount error", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorUnmarshal(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"200","reasonPhrase":"OK"}}]}`)
	reqRes2.Path = "/v1/updateContext"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorMarshal(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"200","reasonPhrase":"OK"}}]}`)
	reqRes2.Path = "/v1/updateContext"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte("[]")
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes3.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1, reqRes2, reqRes3)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorHTTPV1(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"200","reasonPhrase":"OK"}}]}`)
	reqRes2.Path = "/v1/updateContext"
	reqRes2.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorHTTPStatusV1(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes2.Path = "/v1/updateContext"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorUnmarshalv1Res(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"200","reasonPhrase":"OK"}}]}`)
	reqRes2.Path = "/v1/updateContext"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	helper.SetJSONDecodeErr(c.Ngsi, 1)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorErrorCode(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"400","reasonPhrase":"Bad Request"}}], "errorCode":{"code":"400","reasonPhrase":"Bad Request"}}`)
	reqRes2.Path = "/v1/updateContext"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "error 400 Bad Request ", ngsiErr.Message)
	}
}

func TestRemoveV1ErrorContextResError(t *testing.T) {
	c := setupTest([]string{"rm", "--host", "orion", "--type", "Thing", "--ngsiV1", "--run"})

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"type":"AEDFacilities","id":"AEDFacilities.1"},{"type":"AEDFacilities","id":"AEDFacilities.2"},{"type":"AEDFacilities","id":"AEDFacilities.3"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"191"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ReqData = []byte(`{"contextElements":[{"id":"AEDFacilities.1","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.2","isPattern":"false","type":"AEDFacilities"},{"id":"AEDFacilities.3","isPattern":"false","type":"AEDFacilities"}],"updateAction":"DELETE"}`)
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.1"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.2"},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"AEDFacilities","isPattern":"false","id":"AEDFacilities.3"},"statusCode":{"code":"400","reasonPhrase":"Bad Request"}}]}`)
	reqRes2.Path = "/v1/updateContext"

	helper.SetClientHTTP(c, reqRes1, reqRes2)

	err := removeV1(c, c.Ngsi, c.Client, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "error 400 Bad Request", ngsiErr.Message)
	}
}
