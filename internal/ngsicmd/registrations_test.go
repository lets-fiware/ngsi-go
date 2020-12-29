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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestRegistrationsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := registrationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsListErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := registrationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsGetErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := registrationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCreateErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld", "--data="})
	err := registrationsCreate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := registrationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteErrorV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsDeleteErrorLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,link")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := registrationsTemplate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateNgsiV2(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--ngsiType=v2"})
	err := registrationsTemplate(c)

	assert.NoError(t, err)
}

func TestRegistrationsTemplateNgsiLd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--ngsiType=ld"})
	err := registrationsTemplate(c)

	assert.NoError(t, err)
}

func TestRegistrationsTemplateErrorNgsiTypeMissing(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	err := registrationsTemplate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "missing ngsiType", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsTemplateErrorNgsiTypeError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType")
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--ngsiType=v1"})
	err := registrationsTemplate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType error: v1", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCountV2(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/registrations"
	reqRes.ResHeader = http.Header{"Fiware-Total-Count": []string{"12"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "12\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
func TestRegistrationsCountLD(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/v1/csourceRegistrations"
	reqRes.ResHeader = http.Header{"Ngsild-Results-Count": []string{"21"}}
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-ld"})
	err := registrationsCount(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "21\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCountErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	err := registrationsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCountErrorNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,link")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--link=abc"})
	err := registrationsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCountErrorHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, ngsiErr.Message, "url error")
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCountErrorStatusCode(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/v2/registrations"
	reqRes.ResBody = []byte(`error`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host,type")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--type=AirQualityObserved"})
	err := registrationsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}

func TestRegistrationsCountErrorResultsCount(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v2/registrations"
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	setupFlagString(set, "host")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})
	err := registrationsCount(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
	} else {
		t.FailNow()
	}
}
