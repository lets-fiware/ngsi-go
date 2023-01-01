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
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestCopyV2V2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`

	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes1)

	err := copy(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyV1V1(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--ngsiV1"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"
	helper.SetClientHTTP(c, reqRes1)

	err := copy(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyLDLD(t *testing.T) {

	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	helper.SetClientHTTP(c, reqRes1)

	err := copy(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyV2LD(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes1)

	err := copy(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyErrorSavePreviousArgs(t *testing.T) {
	c := setupTest([]string{"cp", "--host", "orion", "--host2", "orion-ld", "--type", "Thing"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("open error"), Filename: helper.StrPtr("config")}

	err := copy(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestCopyErrorSrcDestSame(t *testing.T) {
	c := setupTest([]string{"cp", "--host", "orion", "--host2", "orion", "--type", "Thing"})

	err := copy(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "source and destination are same", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrroLDV2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes1)

	err := copy(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "cannot copy entities from NGSI-LD to NGSI v2", ngsiErr.Message)
	}
}

func TestCopyErrorCopy(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes1)

	err := copy(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestV2V2CopyPage(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes3.Path = "/v2/entities"

	reqRes4 := helper.MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusNoContent

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes3)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes4)
	c.Client2.HTTP = mockDest

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "2\n"
		assert.Equal(t, expected, actual)
	}
}

func TestV2V2CopyPageSkipForwarding(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--skipForwarding", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes1.Path = "/v2/entities"
	q := "limit=100&offset=0&options=count%2CskipForwarding&type=Thing"
	reqRes1.RawQuery = &q

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes3.Path = "/v2/entities"

	reqRes4 := helper.MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusNoContent

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes3)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes4)
	c.Client2.HTTP = mockDest

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "2\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyV2V2ErrorRunFlag(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes)

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyErrorHTTP(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes.Path = "/v2/entities"
	reqRes.Err = errors.New("http error")
	helper.SetClientHTTP(c, reqRes)

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestCopyErrorHTTPStatus(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes)

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestCopyErrorResultCount(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{""}}
	reqRes.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes)

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"\": invalid syntax", ngsiErr.Message)
	}
}

func TestCopyResultCountZero(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes)

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "0\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorJSONUnmarshal(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/entities"
	helper.SetClientHTTP(c, reqRes)

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type ngsilib.EntitiesRespose Field: (1) {}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorOpUpdate(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Err = errors.New("opupdate error")
	reqRes2.Res.StatusCode = http.StatusBadRequest

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "opupdate error", ngsiErr.Message)
	}
}

func TestCopyErrorOpUpdateStatus(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV2V2(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestMakeV1Entities(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)

	b, c, err := makeV1Entities(entities, "APPEND")

	if assert.NoError(t, err) {
		expected := "{\"contextElements\":[{\"attributes\":[{\"name\":\"abc\",\"type\":\"Text\",\"value\":\"001\"}],\"id\":\"thing001\",\"isPattern\":\"false\",\"type\":\"Thing\"},{\"attributes\":[{\"name\":\"abc\",\"type\":\"Text\",\"value\":\"002\"}],\"id\":\"thing002\",\"isPattern\":\"false\",\"type\":\"Thing\"},{\"attributes\":[{\"name\":\"abc\",\"type\":\"Text\",\"value\":\"003\"}],\"id\":\"thing002\",\"isPattern\":\"false\",\"type\":\"Thing\"}],\"updateAction\":\"APPEND\"}"
		assert.Equal(t, expected, string(b))
		assert.Equal(t, 3, c)
	}
}

func TestMakeV1EntitiesErrorUnmarshal(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestMakeV1EntitiesErrorStatus(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"400","reasonPhrase":"OK","details":"Count: 3"}}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "{400 OK Count: 3} OK", ngsiErr.Message)
	}
}

func TestMakeV1EntitiesErrorCount(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count:3"}}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "count error", ngsiErr.Message)
	}
}

func TestMakeV1EntitiesErrorAtoi(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: abc"}}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"abc\": invalid syntax", ngsiErr.Message)
	}
}

func TestMakeV1EntitiesErrorMarshal(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestCopyV1V1Page(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 300"}}`)
	reqRes1.Path = "/v1/queryContext"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes3.Path = "/v1/queryContext"

	reqRes4 := helper.MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusOK
	reqRes4.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes4.Path = "/v1/updateContext"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes3)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes4)
	c.Client2.HTTP = mockDest

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "6\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1PageZero(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 0"}}`)
	reqRes1.Path = "/v1/queryContext"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyV1V1ErrorHTTP1(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v2/queryContext"
	reqRes1.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes1)

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestCopyV1V1ErrorHTTPStatus1(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes1.Path = "/v1/queryContext"

	helper.SetClientHTTP(c, reqRes1)

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestCopyV1V1ErrorMkaeV1Entities(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: abc"}}`)
	reqRes1.Path = "/v1/queryContext"

	helper.SetClientHTTP(c, reqRes1)

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"abc\": invalid syntax", ngsiErr.Message)
	}
}

func TestCopyV1V1ErrorHTTP2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"
	reqRes2.Err = errors.New("http error")

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestCopyV1V1ErrorHTTPStatus2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes2.Path = "/v1/updateContext"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestCopyV1V1ErrorUnmarshal(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}`)
	reqRes2.Path = "/v1/updateContext"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestCopyV1V1ErrorStatusCode(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--ngsiV1", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"400","reasonPhrase":"Bad Request"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV1V1(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error 400 Bad Request", ngsiErr.Message)
	}
}

func TestLDLDCopyPage(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"200"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes3.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes3.Path = "/ngsi-ld/v1/entities"

	reqRes4 := helper.MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusCreated
	reqRes4.Path = "/ngsi-ld/v1/entityOperations/create"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes3)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes4)
	c.Client2.HTTP = mockDest

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "201\n"
		assert.Equal(t, expected, actual)
	}
}

func TestLDLDCopyPageCountZero(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}

	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestLDLDCopyPageContext(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestLDLDCopyErrorHTTP1(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes1.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes1)

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestLDLDCopyErrorHTTPStatus1(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestLDLDCopyErrorCount(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.Path = "/ngsi-ld/v1/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "results count error", ngsiErr.Message)
	}
}

func TestLDLDCopyErrorContext(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--context2", "abc", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	}
}

func TestLDLDCopyErrorHTTP2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes2.Err = errors.New("http error")

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestLDLDCopyErrorHTTPStatus2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "ld"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyLDLD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestV2LDCopyPage(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"200"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes3.Path = "/v2/entities"

	reqRes4 := helper.MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusCreated
	reqRes4.Path = "/ngsi-ld/v1/entityOperations/create"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes3)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes4)
	c.Client2.HTTP = mockDest

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "201\n"
		assert.Equal(t, expected, actual)
	}
}

func TestV2LDCopyPageSkipForwarding(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--skipForwarding", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"200"}}
	q := "limit=100&offset=0&options=count%2CskipForwarding&type=Thing"
	reqRes1.RawQuery = &q
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"

	reqRes3 := helper.MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes3.Path = "/v2/entities"

	reqRes4 := helper.MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusCreated
	reqRes4.Path = "/ngsi-ld/v1/entityOperations/create"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes3)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes4)
	c.Client2.HTTP = mockDest

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "201\n"
		assert.Equal(t, expected, actual)
	}
}

func TestV2LDCopyPageCountZero(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes1.Path = "/v2/entities"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "0\n"
		assert.Equal(t, expected, actual)
	}
}

func TestV2LDCopyPageContext(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--context2", "ld", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "3\n"
		assert.Equal(t, expected, actual)
	}
}

func TestV2LDCopyErrorHTTP1(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"
	reqRes1.Err = errors.New("http error")

	helper.SetClientHTTP(c, reqRes1)

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestV2LDCopyErrorHTTPStatus1(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestV2LDCopyErrorCount(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "results count error", ngsiErr.Message)
	}
}

func TestV2LDCopyErrornormalized2LD(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T","location":{"type":"geo:point","value":1}}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal number into Go value of type string Field: (1) 1", ngsiErr.Message)
	}
}

func TestV2LDCopyErrorContext(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--context2", "abc", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	helper.SetClientHTTP(c, reqRes1)

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	}
}

func TestV2LDCopyErrorHTTP2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes2.Err = errors.New("http error")

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestV2LDCopyErrorHTTPStatus2(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion-src": {
				"serverHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"serverHost": "https://orion-dest",
				"ngsiType": "ld"
			}
		}
	}`
	c := setupTestWithConfig([]string{"cp", "--host", "orion-src", "--host2", "orion-dest", "--type", "Thing", "--run"}, conf)

	reqRes1 := helper.MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001","type":"T"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes1.Path = "/v2/entities"

	reqRes2 := helper.MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)

	mockSource := helper.NewMockHTTP()
	mockSource.ReqRes = append(mockSource.ReqRes, reqRes1)
	c.Client.HTTP = mockSource

	mockDest := helper.NewMockHTTP()
	mockDest.ReqRes = append(mockDest.ReqRes, reqRes2)
	c.Client2.HTTP = mockDest

	err := copyV2LD(c, c.Ngsi, c.Client, c.Client2, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	}
}

func TestNormalized2LD(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	body := []byte(`[{"type":"T","id":"E","name":{"type":"Text","value":"FIWARE"}}]`)

	b, err := normalized2LD(body)

	if assert.NoError(t, err) {
		expected := `[{"id":"urn:ngsi-ld:T:E","name":{"type":"Property","value":"FIWARE"},"type":"T"}]`
		assert.Equal(t, expected, string(b))
	}
}

func TestNormalized2LDErrorJSONUnmarshal(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	body := []byte(`[{"type":"T","id":"E","name":{"type":"Text","value":"FIWARE"}}]`)

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	_, err := normalized2LD(body)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestNormalized2LDErrorNormalized2LDEntity(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	body := []byte(`[{"type":"T","id":"E","location":{"type":"geo:point","value":1}}]`)

	_, err := normalized2LD(body)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal number into Go value of type string Field: (1) 1", ngsiErr.Message)
	}
}

func TestNormalized2LDErrorJSONMarshal(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	body := []byte(`[{"type":"T","id":"E","name":{"type":"Text","value":"FIWARE"}}]`)

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, err := normalized2LD(body)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestNormalized2LDEntity(t *testing.T) {
	var v2 ngsilib.NgsiEntity

	v2Entity := `{"type":"T", "id": "E", "name": {"type":"Text", "value": "FIWARE"}}`
	e := json.Unmarshal([]byte(v2Entity), &v2)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	ldEntity, err := normalized2LDEntity(v2)

	actual, e := json.Marshal(ldEntity)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	if assert.NoError(t, err) {
		expected := `{"id":"urn:ngsi-ld:T:E","name":{"type":"Property","value":"FIWARE"},"type":"T"}`
		assert.Equal(t, expected, string(actual))
	}
}

func TestNormalized2LDEntityErrorId(t *testing.T) {
	var v2 ngsilib.NgsiEntity

	v2Entity := `{"type":"T", "id": 123, "name": {"type":"Text", "value": "FIWARE"}}`
	e := json.Unmarshal([]byte(v2Entity), &v2)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := normalized2LDEntity(v2)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "id not string", ngsiErr.Message)
	}
}

func TestNormalized2LDEntityErrorIdType(t *testing.T) {
	var v2 ngsilib.NgsiEntity

	v2Entity := `{"id": "E", "type": 123}`
	e := json.Unmarshal([]byte(v2Entity), &v2)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := normalized2LDEntity(v2)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "type not string", ngsiErr.Message)
	}
}

func TestNormalized2LDEntityErrorType(t *testing.T) {
	var v2 ngsilib.NgsiEntity

	v2Entity := `{"type": 123, "name": {"type":"Text", "value": "FIWARE"}}`
	e := json.Unmarshal([]byte(v2Entity), &v2)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := normalized2LDEntity(v2)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "type not string", ngsiErr.Message)
	}
}

func TestNormalized2LDEntityErrorAttrType(t *testing.T) {
	var v2 ngsilib.NgsiEntity

	v2Entity := `{"name": ""}`
	e := json.Unmarshal([]byte(v2Entity), &v2)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := normalized2LDEntity(v2)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "name: attribute error", ngsiErr.Message)
	}
}

func TestNormalized2LDEntityErrorAttrValue(t *testing.T) {
	var v2 ngsilib.NgsiEntity

	v2Entity := `{"name": {"type":1,"value": 1}}`
	e := json.Unmarshal([]byte(v2Entity), &v2)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := normalized2LDEntity(v2)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "type not string", ngsiErr.Message)
	}
}

func TestLdAttribute(t *testing.T) {
	cases := []struct {
		key      string
		attr     string
		expected string
	}{
		{
			key:      "speed",
			attr:     `{"type": "Number", "value": 123, "metadata": {"timestamp": {"type": "DateTime", "value": "2017-06-17T07:21:24.238Z"}}}`,
			expected: `{"observedAt":"2017-06-17T07:21:24.238Z","type":"Property","value":123}`,
		},
		{
			key:      "speed",
			attr:     `{"type": "Number", "value": 123, "metadata": {"accuracy": {"type": "Number", "value": 0.8}}}`,
			expected: `{"accuracy":{"type":"Property","value":0.8},"type":"Property","value":123}`,
		},
		{
			key:      "speed",
			attr:     `{"type": "Number", "value": 123, "metadata": {"unitCode": {"type": "Text", "value": "GP"}}}`,
			expected: `{"type":"Property","unitCode":"GP","value":123}`,
		},
		{
			key:      "refEntity",
			attr:     `{"type": "Relationship", "value": "urn:NGSI-LD:Device:device001"}`,
			expected: `{"object":"urn:NGSI-LD:Device:device001","type":"Relationship"}`,
		},
		{
			key:      "observedAt",
			attr:     `{"type": "DateTime", "value": "2017-06-17T07:21:24.238Z"}`,
			expected: `{"type":"Property","value":{"@type":"DateTime","@value":"2017-06-17T07:21:24.238Z"}}`,
		},
		{
			key:      "location",
			attr:     `{"type": "geo:point", "value": "35.1,135.1"}`,
			expected: `{"type":"GeoProperty","value":{"type":"Point","coordinates":[135.1,35.1]}}`,
		},
		{
			key:      "location",
			attr:     `{"type": "geo:line", "value": ["35.1,135.1", "35.2,135.2"]}`,
			expected: `{"type":"GeoProperty","value":{"type":"LineString","coordinates":[[135.1,35.1],[135.2,35.2]]}}`,
		},
		{
			key:      "location",
			attr:     `{"type": "geo:json", "value": {"type": "Point", "coordinates": [-3.703,40.417]}}`,
			expected: `{"type":"GeoProperty","value":{"coordinates":[-3.703,40.417],"type":"Point"}}`,
		},
		{
			key:      "name",
			attr:     `{"type": "Text", "value": "FIWARE"}`,
			expected: `{"type":"Property","value":"FIWARE"}`,
		},
		{
			key:      "count",
			attr:     `{"type": "Number", "value": 123}`,
			expected: `{"type":"Property","value":123}`,
		},
	}

	_ = helper.SetupTestInitCmd(nil)

	for _, c := range cases {
		var attr map[string]interface{}
		e := json.Unmarshal([]byte(c.attr), &attr)
		if !assert.NoError(t, e) {
			t.FailNow()
		}

		attr, err := ldAttribute(c.key, attr)

		actual, e := json.Marshal(attr)
		assert.NoError(t, e)

		if assert.NoError(t, err) {
			assert.Equal(t, c.expected, string(actual))
		}
	}
}

func TestLdAttributeErrorType(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type": 1}`)

	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := ldAttribute("", attr)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "type not string", ngsiErr.Message)
	}
}

func TestLdAttributeErrorRelationship(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"Relationship","value": 1}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Relationship: value not string", ngsiErr.Message)
	}
}

func TestLdAttributeErrorDateTime(t *testing.T) {
	_ = helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"DateTime","value": 1}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "DateTime: value not string", ngsiErr.Message)
	}
}

func TestLdAttributeErrorGeoEncodeGeoPoint(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"geo:point","value":["35.1, 135.1", "35.2, 135.2"]}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLdAttributeErrorGeoDecodeGeoPoint(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"geo:point","value":["35.1, 135.1", "35.2, 135.2"]}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLdAttributeErrorGeoEncodeGeoXXX(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"geo:line","value":["35.1, 135.1", "35.2, 135.2"]}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLdAttributeErrorGeoDecodeGeoXXX(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"geo:line","value":["35.1, 135.1", "35.2, 135.2"]}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLdAttributeErrorMetadataEncode(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"Text","value":"", "metadata": {"accuracy": {"type": "Number", "value": 0.8}}}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestLdAttributeErrorMetadataDecode(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	attribute := []byte(`{"type":"Text","value":"", "metadata": {"accuracy": {"type": "Number", "value": 0.8}}}`)
	var attr map[string]interface{}
	e := json.Unmarshal(attribute, &attr)
	if !assert.NoError(t, e) {
		t.FailNow()
	}

	helper.SetJSONDecodeErr(c.Ngsi, 0)

	_, err := ldAttribute("", attr)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestToGeoJSON(t *testing.T) {
	cases := []struct {
		geoType       string
		location      []string
		expectedType  string
		expectedValue string
	}{
		{
			geoType:       "geo:point",
			location:      []string{"35.1, 135.1"},
			expectedType:  "Point",
			expectedValue: "[135.1,35.1]",
		},
		{
			geoType:       "geo:line",
			location:      []string{"35.1, 135.1", "35.2, 135.2"},
			expectedType:  "LineString",
			expectedValue: "[[135.1,35.1],[135.2,35.2]]",
		},
		{
			geoType:       "geo:polygon",
			location:      []string{"35.1, 135.1", "35.2, 135.2", "35.3, 135.3", "35.1, 135.1"},
			expectedType:  "Polygon",
			expectedValue: "[[[135.1,35.1],[135.2,35.2],[135.3,35.3],[135.1,35.1]]]",
		},
		{
			geoType:       "geo:box",
			location:      []string{"35.1, 135.1", "35.2, 135.2"},
			expectedType:  "Polygon",
			expectedValue: "[[[135.1,35.1],[135.2,35.1],[135.2,35.2],[135.1,35.2],[135.1,35.1]]]",
		},
	}
	for _, c := range cases {
		actual, err := toGeoJSON(c.geoType, c.location)
		coords, e := json.Marshal(actual.Coordinates)
		if !assert.NoError(t, e) {
			t.FailNow()
		}

		if assert.NoError(t, err) {
			assert.Equal(t, c.expectedType, actual.Type)
			assert.Equal(t, c.expectedValue, string(coords))
		}
	}
}

func TestToGeoJSONErrorCoord1(t *testing.T) {
	_, err := toGeoJSON("geo:point", []string{"abc, 135.1"})

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.ParseFloat: parsing \"abc\": invalid syntax", ngsiErr.Message)
	}
}

func TestToGeoJSONErrorCoord2(t *testing.T) {
	_, err := toGeoJSON("geo:point", []string{"35.1, abc"})

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.ParseFloat: parsing \"abc\": invalid syntax", ngsiErr.Message)
	}
}

func TestToGeoJSONErrorGeoType(t *testing.T) {
	_, err := toGeoJSON("geo:unknown", []string{"35.1, 135.1"})

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "unknown type: geo:unknown", ngsiErr.Message)
	}
}

func TestLdRelationship(t *testing.T) {
	cases := []struct {
		attrName string
		entityId string
		expected string
	}{
		{attrName: "", entityId: "http://letsfiware.jp/ns/data-models#fiware", expected: "http://letsfiware.jp/ns/data-models#fiware"},
		{attrName: "", entityId: "https://letsfiware.jp/ns/data-models#fiware", expected: "https://letsfiware.jp/ns/data-models#fiware"},
		{attrName: "", entityId: "urn:ngsi-ld:Device:001", expected: "urn:ngsi-ld:Device:001"},
		{attrName: "refName", entityId: "Device001", expected: "urn:ngsi-ld:Name:Device001"},
		{attrName: "Device", entityId: "device001", expected: "urn:ngsi-ld:Device:device001"},
	}
	for _, c := range cases {
		actual := ldRelationship(c.attrName, c.entityId)
		assert.Equal(t, c.expected, actual)
	}
}

func TestNgsildUri(t *testing.T) {
	cases := []struct {
		typePart string
		idPart   string
		expected string
	}{
		{typePart: "Device", idPart: "device:001", expected: "urn:ngsi-ld:device:001"},
		{typePart: "Device", idPart: "device001", expected: "urn:ngsi-ld:Device:device001"},
	}
	for _, c := range cases {
		actual := ngsildUri(c.typePart, c.idPart)
		assert.Equal(t, c.expected, actual)
	}
}

func TestLdId(t *testing.T) {
	cases := []struct {
		entityId   string
		entityType string
		expected   string
	}{
		{entityId: "http://letsfiware.jp/ns/data-models#fiware", entityType: "", expected: "http://letsfiware.jp/ns/data-models#fiware"},
		{entityId: "https://letsfiware.jp/ns/data-models#fiware", entityType: "", expected: "https://letsfiware.jp/ns/data-models#fiware"},
		{entityId: "urn:ngsi-ld:Device:001", entityType: "", expected: "urn:ngsi-ld:Device:001"},
		{entityId: "Device001", entityType: "Device", expected: "urn:ngsi-ld:Device:Device001"},
		{entityId: "device001", entityType: "Device", expected: "urn:ngsi-ld:Device:device001"},
	}

	for _, c := range cases {
		actual := ldId(c.entityId, c.entityType)
		assert.Equal(t, c.expected, actual)
	}
}

func TestNormalizeDate(t *testing.T) {
	cases := []struct {
		arg      string
		expected string
	}{
		{arg: "2014-10-01T00:00:00.00Z", expected: "2014-10-01T00:00:00.00Z"},
		{arg: "2014-10-01T00:00:00.00", expected: "2014-10-01T00:00:00.00Z"},
	}

	for _, c := range cases {
		actual := normalizeDate(c.arg)
		assert.Equal(t, "DateTime", actual.Type)
		assert.Equal(t, c.expected, actual.Value)
	}
}
