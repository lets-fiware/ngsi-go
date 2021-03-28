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
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestCopyV2V2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--type=Thing"})

	err := copy(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "1 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,type")
	setupFlagBool(set, "ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--type=Thing", "--ngsiV1"})

	err := copy(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyLDLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--type=Thing"})

	err := copy(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "1 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorV2LD(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--type=Thing"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "not yet implemented", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorSavePreviousArgs(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("open error")}
	filename := "config"
	ngsi.ConfigFile.SetFileName(&filename)

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=abc", "--host=orion"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorParse2(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link2,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link2=abc", "--host2=orion-ld"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorNewClientDestination(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--host2=orion-v2"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error host: orion-v2 (destination)", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorNotBroker(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--host2=keyrock"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "destination not broker", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorSrcDestSame(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--host2=orion"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "source and destination are same", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrroLDV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--type=Thing"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "cannot copy entites from NGSI-LD to NGSI v2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrroEntityType(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "specify entity type", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorCopy(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,type")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--type=Thing"})

	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestV2V2CopyPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusNoContent
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes3.Path = "/v2/entities"
	reqRes4 := MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusNoContent
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	mock.ReqRes = append(mock.ReqRes, reqRes4)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "2\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyV2V2ErrorRunFlag(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3 entities will be copied. run copy with --run option\n"
		assert.Equal(t, expected, actual)
	}
}

func TestCopyErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes.Path = "/v2/entitie"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"3"}}
	reqRes.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorResultCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{""}}
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"\": invalid syntax", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyResultCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "0\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorJSONUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type ngsicmd.entitiesRespose Field: (1) {}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorOpUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Err = errors.New("opupdate error")
	reqRes2.Res.StatusCode = http.StatusBadRequest
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "opupdate error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorOpUpdateStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV2V2(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestMakeV1Entities(t *testing.T) {
	setupTest()
	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)

	b, c, err := makeV1Entities(entities, "APPEND")

	if assert.NoError(t, err) {
		expected := "{\"contextElements\":[{\"attributes\":[{\"name\":\"abc\",\"type\":\"Text\",\"value\":\"001\"}],\"id\":\"thing001\",\"isPattern\":\"false\",\"type\":\"Thing\"},{\"attributes\":[{\"name\":\"abc\",\"type\":\"Text\",\"value\":\"002\"}],\"id\":\"thing002\",\"isPattern\":\"false\",\"type\":\"Thing\"},{\"attributes\":[{\"name\":\"abc\",\"type\":\"Text\",\"value\":\"003\"}],\"id\":\"thing002\",\"isPattern\":\"false\",\"type\":\"Thing\"}],\"updateAction\":\"APPEND\"}"
		assert.Equal(t, expected, string(b))
		assert.Equal(t, 3, c)
	}
}

func TestMakeV1EntitiesErrorUnmarshal(t *testing.T) {
	setupTest()

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestMakeV1EntitiesErrorStatus(t *testing.T) {
	setupTest()

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"400","reasonPhrase":"OK","details":"Count: 3"}}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "{400 OK Count: 3} OK", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestMakeV1EntitiesErrorCount(t *testing.T) {
	setupTest()

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count:3"}}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "count error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestMakeV1EntitiesErrorAtoi(t *testing.T) {
	setupTest()

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: abc"}}`)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"abc\": invalid syntax", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestMakeV1EntitiesErrorMarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	entities := []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)

	setJSONEncodeErr(ngsi, 0)

	_, _, err := makeV1Entities(entities, "APPEND")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV2LDMain(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	err := copyV2LD(nil, ngsi, nil, nil, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "not yet implemented", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1Page(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 300"}}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes3.Path = "/v1/queryContext"
	reqRes4 := MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusOK
	reqRes4.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes4.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	mock.ReqRes = append(mock.ReqRes, reqRes4)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "6\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1PageZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 0"}}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "0\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1ErrorHTTP1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v2/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1ErrorHTTPStatus1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1ErrorMkaeV1Entities(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: abc"}}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"abc\": invalid syntax", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1ErrorHTTP2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v2/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1ErrorHTTPStatus2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes2.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1ErrorUnmarshal(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}`)
	reqRes2.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyV1V1ErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes1.Path = "/v1/queryContext"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.ResBody = []byte(`{"contextResponses":[{"contextElement":{"type":"Thing","isPattern":"false","id":"thing001","attributes":[{"name":"abc","type":"Text","value":"001"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"002"}]},"statusCode":{"code":"200","reasonPhrase":"OK"}},{"contextElement":{"type":"Thing","isPattern":"false","id":"thing002","attributes":[{"name":"abc","type":"Text","value":"003"}]},"statusCode":{"code":"400","reasonPhrase":"Bad Request"}}],"errorCode":{"code":"200","reasonPhrase":"OK","details":"Count: 3"}}`)
	reqRes2.Path = "/v1/updateContext"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run,ngsiV1")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--ngsiV1"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyV1V1(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "error 400 Bad Request", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"200"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes3.ResHeader = http.Header{"Ngsild-Results-Count": []string{"1"}}
	reqRes3.Path = "/ngsi-ld/v1/entities"
	reqRes4 := MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusCreated
	reqRes4.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	mock.ReqRes = append(mock.ReqRes, reqRes4)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "201\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyPageCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"0"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "0\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyPageContext(t *testing.T) {
	ngsi, set, app, buf := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusCreated
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,context2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--context2=ld"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "3\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyErrorHTTP1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyErrorHTTPStatus1(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusBadRequest
	reqRes1.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyErrorCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "results count error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyErrorContext(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2,context2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run", "--context2=abc"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyErrorHTTP2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes2.Path = "/ngsi-ld/v2/entityOperations/create"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestLDLDCopyErrorHTTPStatus2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Ngsild-Results-Count": []string{"3"}}
	reqRes1.Path = "/ngsi-ld/v1/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusBadRequest
	reqRes2.Path = "/ngsi-ld/v1/entityOperations/create"
	reqRes2.ResBody = []byte(`{"code":"400","reasonPhrase":"Bad Request"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock

	setupFlagString(set, "host,host2")
	setupFlagBool(set, "run")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--host2=orion-dest", "--run"})

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)
	source, err := newClient(ngsi, c, false, []string{"broker"})
	assert.NoError(t, err)
	flags, err := parseFlags2(ngsi, c)
	assert.NoError(t, err)
	dest, err := ngsi.NewClient(ngsi.Destination, flags, false)
	assert.NoError(t, err)

	err = copyLDLD(c, ngsi, source, dest, "Thing")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, " {\"code\":\"400\",\"reasonPhrase\":\"Bad Request\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
