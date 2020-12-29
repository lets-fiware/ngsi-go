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

func TestCopyErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
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
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorParse2(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link2,destination")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link2=abc", "--destination=orion-ld"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorNewClientDestination(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link,destination")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--destination=orion-v2"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error host: orion-v2 (destination)", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorIsNgsiLdSrc(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,destination")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--destination=orion-ld"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorIsNgsiLdDest(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,destination")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--destination=orion-ld"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorRunFlag(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "run copy with --run option", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusNoContent
	reqRes.Path = "/v2/entitie"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest", "--run"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorHTTPStatus(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest", "--run"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorResultCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest", "--run"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "strconv.Atoi: parsing \"\": invalid syntax", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyResultCountZero(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"0"}}
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest", "--run"})
	err := copy(c)

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
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte("{}")
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/entities"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest", "--run"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type ngsicmd.entitiesRespose Field: (1) {}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyErrorOpUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	setupFlagBool(set, "run")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"1"}}
	reqRes.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Err = errors.New("opupdate error")
	reqRes2.Res.StatusCode = http.StatusBadRequest
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest", "--run"})
	err := copy(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 11, ngsiErr.ErrNo)
		assert.Equal(t, "opupdate error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestCopyPage(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"brokers": {
			"orion-src": {
				"brokerHost": "https://orion-src",
				"ngsiType": "v2"
			},
			"orion-dest": {
				"brokerHost": "https://orion-dest",
				"ngsiType": "v2"
			}
		}
	}`
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,destination")
	setupFlagBool(set, "run")
	reqRes1 := MockHTTPReqRes{}
	reqRes1.Res.StatusCode = http.StatusOK
	reqRes1.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes1.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes1.Path = "/v2/entities"
	reqRes2 := MockHTTPReqRes{}
	reqRes2.Res.StatusCode = http.StatusOK
	reqRes3 := MockHTTPReqRes{}
	reqRes3.Res.StatusCode = http.StatusOK
	reqRes3.ResBody = []byte(`[{"id":"device001"}]`)
	reqRes3.ResHeader = http.Header{"Fiware-Total-Count": []string{"150"}}
	reqRes3.Path = "/v2/entities"
	reqRes4 := MockHTTPReqRes{}
	reqRes4.Res.StatusCode = http.StatusOK
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes1)
	mock.ReqRes = append(mock.ReqRes, reqRes2)
	mock.ReqRes = append(mock.ReqRes, reqRes3)
	mock.ReqRes = append(mock.ReqRes, reqRes4)
	ngsi.HTTP = mock
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-src", "--destination=orion-dest", "--run"})
	err := copy(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "2\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
