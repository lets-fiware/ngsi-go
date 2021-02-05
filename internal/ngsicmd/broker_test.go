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
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestBrokersList(t *testing.T) {
	_, set, app, buf := setupTest()

	c := cli.NewContext(app, set, nil)

	err := brokersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "orion orion-alias orion-ld\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBrokersListHost(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	err := brokersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "brokerHost https://orion\nngsiType v2\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBrokersListJSON(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json"})

	err := brokersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"orion\":{\"serverType\":\"broker\",\"serverHost\":\"https://orion\",\"ngsiType\":\"v2\"},\"orion-alias\":{\"serverType\":\"broker\",\"serverHost\":\"orion-ld\",\"ngsiType\":\"ld\"},\"orion-ld\":{\"serverType\":\"broker\",\"serverHost\":\"https://orion-ld\",\"ngsiType\":\"ld\"}}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBrokersListJSONPretty(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json", "--pretty"})

	err := brokersList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"orion\": {\n    \"serverType\": \"broker\",\n    \"serverHost\": \"https://orion\",\n    \"ngsiType\": \"v2\"\n  },\n  \"orion-alias\": {\n    \"serverType\": \"broker\",\n    \"serverHost\": \"orion-ld\",\n    \"ngsiType\": \"ld\"\n  },\n  \"orion-ld\": {\n    \"serverType\": \"broker\",\n    \"serverHost\": \"https://orion-ld\",\n    \"ngsiType\": \"ld\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBrokersListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := brokersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersListErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion-v2"})
	c := cli.NewContext(app, set, nil)

	err := brokersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "orion-v2 not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersListErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "json")
	setJSONEncodeErr(ngsi, 2)

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json"})

	err := brokersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersListErrorJSONPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json", "--pretty"})

	setJSONIndentError(ngsi)

	err := brokersList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersGet(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := brokersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "brokerHost https://orion\nngsiType v2\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBrokersGetJSON(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--json"})

	err := brokersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"serverType\":\"broker\",\"serverHost\":\"https://orion\",\"ngsiType\":\"v2\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBrokersGetJSONPretty(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--json", "--pretty"})

	err := brokersGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"serverType\": \"broker\",\n  \"serverHost\": \"https://orion\",\n  \"ngsiType\": \"v2\"\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBrokersGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := brokersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersGetErrorHostNotFound(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := brokersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersGetErrorBrokerListErrorJSON(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-v2", "--json"})

	err := brokersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion-v2 not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersGetErrorJSONPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--json", "--pretty"})

	setJSONIndentError(ngsi)

	err := brokersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersGetErrorBrokerList(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-v2"})

	err := brokersGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "orion-v2 not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersAdd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-v2", "--brokerHost=http://orion", "--ngsiType=v2"})

	err := brokersAdd(c)

	assert.NoError(t, err)
}

func TestBrokersAddTenant(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost,service")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-v2", "--brokerHost=http://orion", "--ngsiType=v2", "--service=Foo"})

	err := brokersAdd(c)

	if assert.NoError(t, err) {
		list := ngsi.AllServersList()
		v := (*list)["orion-v2"]
		assert.Equal(t, "foo", v.Tenant)
	}
}

func TestBrokersAddLD(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orionld", "--brokerHost=http://orion", "--ngsiType=ld"})

	err := brokersAdd(c)

	assert.NoError(t, err)
}

func TestBrokersAddLDSafeString(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orionld", "--brokerHost=http://orion", "--ngsiType=ld", "--safeString=on"})

	err := brokersAdd(c)

	assert.NoError(t, err)
}

func TestBrokersAddErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersAddErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersAddErrorNameString(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=@orion"})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "name error @orion", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersAddErrorAlreadyExists(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "orion already exists", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestBrokersAddErrorBrokerHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-v2"})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "brokerHost is missing", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersAddErrorNgsiType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orionv2", "--brokerHost=http://orion"})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType is missing", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersAddErrorNgsiType2(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orionv2", "--brokerHost=orion2", "--ngsiType=v2"})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "can't specify ngsiType", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersAddErrorCreateBroker(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,brokerHost,safeString")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orionld", "--ngsiType=ld", "--brokerHost=http://orion-ld", "--safeString=123"})

	err := brokersAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: 123", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersUpdate(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ngsiType=ld"})

	err := brokersUpdate(c)

	assert.NoError(t, err)
}

func TestBrokersUpdateService(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,service")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ngsiType=ld", "--service=Foo"})

	err := brokersUpdate(c)

	assert.NoError(t, err)
}

func TestBrokersUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := brokersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersUpdateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := brokersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersUpdateErrorAlreadyExists(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orionld"})

	err := brokersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orionld not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersUpdateErrorCreateBroker(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ngsiType=v1"})

	err := brokersUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "v1 not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersDelete(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--ngsiType=ld"})

	err := brokersDelete(c)

	assert.NoError(t, err)
}

func TestBrokersDeleteNgsiType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--items=ngsiType"})

	err := brokersDelete(c)

	assert.NoError(t, err)
}

func TestBrokersDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := brokersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersDeleteErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := brokersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersDeleteErrorAlreadyExists(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orionld"})

	err := brokersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orionld not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersDeleteNoItem(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--items=noitem"})

	err := brokersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "delete error - noitem", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersDeleteErrorUpdateBroker(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,items")
	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("open error")}
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--items=ngsiType"})

	err := brokersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestBrokersDeleteErrorReference(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				  "serverHost": "https://orion",
				  "ngsiType": "v2"
			},
			"orion-ld": {
				  "serverHost": "orion"
			}
		}
	}`

	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host,ngsiType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion"})

	err := brokersDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "orion is referenced", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPrintBrokerInfoV2(t *testing.T) {
	ngsi, _, _, buf := setupTest()

	broker := ngsilib.Server{
		ServerHost:   "http://orion",
		NgsiType:     "v2",
		Tenant:       "openiot",
		Scope:        "/iot",
		Context:      "http://context",
		SafeString:   "on",
		IdmType:      "keyrock",
		IdmHost:      "http://keyrock",
		Username:     "fiware",
		Password:     "1234",
		ClientID:     "clientid",
		ClientSecret: "clientsecret",
		XAuthToken:   "false",
		Token:        "token",
		APIPath:      "/path",
	}

	printBrokerInfo(ngsi, &broker)

	actual := buf.String()
	expected := "brokerHost http://orion\nngsiType v2\nFIWARE-Serivce openiot\nFIWARE-SerivcePath /iot\nContext http://context\nSafeString on\nIdmType keyrock\nIdmHost http://keyrock\nUsername fiware\nPassword 1234\nClientID clientid\nClientSecret clientsecret\nXAuthToken false\nToken token\nAPIPath /path\n"
	assert.Equal(t, expected, actual)
}

func TestPrintBrokerInfoLD(t *testing.T) {
	ngsi, _, _, buf := setupTest()

	broker := ngsilib.Server{
		ServerHost: "http://orion-ld",
		NgsiType:   "ld",
		Tenant:     "openiot",
		Scope:      "/iot",
	}

	printBrokerInfo(ngsi, &broker)

	actual := buf.String()
	expected := "brokerHost http://orion-ld\nngsiType ld\nTenant openiot\nScope /iot\n"
	assert.Equal(t, expected, actual)
}
