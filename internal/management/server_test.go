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

package management

import (
	"errors"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestServersList(t *testing.T) {
	c := setupTest([]string{"server", "list"})

	err := serverList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "comet cygnus iota keyrock perseo perseo-core ql queryproxy regproxy tokenproxy wirecloud\n"
		assert.Equal(t, expected, actual)
	}
}

func TestServersListHost(t *testing.T) {
	c := setupTest([]string{"server", "list", "--host", "ql"})

	err := serverList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "serverType quantumleap\nserverHost https://quantumleap\n"
		assert.Equal(t, expected, actual)
	}
}

func TestServersListJSON(t *testing.T) {
	c := setupTest([]string{"server", "list", "--json"})

	err := serverList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"comet\":{\"serverType\":\"comet\",\"serverHost\":\"https://comet\"},\"cygnus\":{\"serverType\":\"cygnus\",\"serverHost\":\"https://cygnus\"},\"iota\":{\"serverType\":\"iota\",\"serverHost\":\"https://iota\"},\"keyrock\":{\"serverType\":\"keyrock\",\"serverHost\":\"https://keyrock\"},\"perseo\":{\"serverType\":\"perseo\",\"serverHost\":\"https://perseo\"},\"perseo-core\":{\"serverType\":\"perseo-core\",\"serverHost\":\"https://perseo-core\"},\"ql\":{\"serverType\":\"quantumleap\",\"serverHost\":\"https://quantumleap\"},\"queryproxy\":{\"serverType\":\"queryproxy\",\"serverHost\":\"https://queryproxy\"},\"regproxy\":{\"serverType\":\"regproxy\",\"serverHost\":\"https://regproxy\"},\"tokenproxy\":{\"serverType\":\"tokenproxy\",\"serverHost\":\"https://tokenproxy\"},\"wirecloud\":{\"serverType\":\"wirecloud\",\"serverHost\":\"https://wirecloud\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestServersListJSONPretty(t *testing.T) {
	c := setupTest([]string{"server", "list", "--json", "--pretty"})

	err := serverList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"comet\": {\n    \"serverType\": \"comet\",\n    \"serverHost\": \"https://comet\"\n  },\n  \"cygnus\": {\n    \"serverType\": \"cygnus\",\n    \"serverHost\": \"https://cygnus\"\n  },\n  \"iota\": {\n    \"serverType\": \"iota\",\n    \"serverHost\": \"https://iota\"\n  },\n  \"keyrock\": {\n    \"serverType\": \"keyrock\",\n    \"serverHost\": \"https://keyrock\"\n  },\n  \"perseo\": {\n    \"serverType\": \"perseo\",\n    \"serverHost\": \"https://perseo\"\n  },\n  \"perseo-core\": {\n    \"serverType\": \"perseo-core\",\n    \"serverHost\": \"https://perseo-core\"\n  },\n  \"ql\": {\n    \"serverType\": \"quantumleap\",\n    \"serverHost\": \"https://quantumleap\"\n  },\n  \"queryproxy\": {\n    \"serverType\": \"queryproxy\",\n    \"serverHost\": \"https://queryproxy\"\n  },\n  \"regproxy\": {\n    \"serverType\": \"regproxy\",\n    \"serverHost\": \"https://regproxy\"\n  },\n  \"tokenproxy\": {\n    \"serverType\": \"tokenproxy\",\n    \"serverHost\": \"https://tokenproxy\"\n  },\n  \"wirecloud\": {\n    \"serverType\": \"wirecloud\",\n    \"serverHost\": \"https://wirecloud\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestServersListErrorHost(t *testing.T) {
	c := setupTest([]string{"server", "list", "--host", "sth"})

	err := serverList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	}
}

func TestServersListErrorJSON(t *testing.T) {
	c := setupTest([]string{"server", "list", "--json"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := serverList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestServersListErrorJSONPretty(t *testing.T) {
	c := setupTest([]string{"server", "list", "--json", "--pretty"})

	helper.SetJSONIndentError(c.Ngsi)

	err := serverList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestServersGet(t *testing.T) {
	c := setupTest([]string{"server", "get", "--host", "comet"})

	err := serverGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "serverType comet\nserverHost https://comet\n"
		assert.Equal(t, expected, actual)
	}
}

func TestServersGetJSON(t *testing.T) {
	c := setupTest([]string{"server", "get", "--host", "comet", "--json"})

	err := serverGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"serverType\":\"comet\",\"serverHost\":\"https://comet\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestServersGetJSONPretty(t *testing.T) {
	c := setupTest([]string{"server", "get", "--host", "comet", "--json", "--pretty"})

	err := serverGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"serverType\": \"comet\",\n  \"serverHost\": \"https://comet\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestServersGetErrorServerListErrorJSON(t *testing.T) {
	c := setupTest([]string{"server", "get", "--host", "sth", "--json"})

	err := serverGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	}
}

func TestServersGetErrorJSONPretty(t *testing.T) {
	c := setupTest([]string{"server", "get", "--host", "comet", "--json", "--pretty"})

	helper.SetJSONIndentError(c.Ngsi)

	err := serverGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestServersGetErrorServerList(t *testing.T) {
	c := setupTest([]string{"server", "get", "--host", "sth"})

	err := serverGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	}
}

func TestServersAdd(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "sth", "--serverHost", "http://sth", "--serverType", "comet"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		assert.Equal(t, "http://sth", (*list)["sth"].ServerHost)
		assert.Equal(t, "comet", (*list)["sth"].ServerType)
	}
}

func TestServersAddKeyrock(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "idm", "--serverHost", "http://keyrock", "--serverType", "keyrock", "--username", "fiware", "--password", "1234"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		assert.Equal(t, "http://keyrock", (*list)["idm"].ServerHost)
		assert.Equal(t, "keyrock", (*list)["idm"].ServerType)
	}
}

func TestServersAddTenant(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "sth", "--serverHost", "http://sth", "--serverType", "comet", "--service", "Foo"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["sth"]
		assert.Equal(t, "http://sth", (*list)["sth"].ServerHost)
		assert.Equal(t, "comet", (*list)["sth"].ServerType)
		assert.Equal(t, "foo", v.Tenant)
	}
}

func TestServersAddErrorNameString(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "@comet", "--serverHost", "http://comet", "--serverType", "comet", "--service", "Foo"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "name error @comet", ngsiErr.Message)
	}
}

func TestServersAddErrorAlreadyExists(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "comet", "--serverHost", "http://comet", "--serverType", "comet", "--service", "Foo"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "comet already exists", ngsiErr.Message)
	}
}
func TestServersAddErrorServerHost(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "server", "http://comet", "--serverType", "comet", "--service", "Foo"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "serverHost is missing", ngsiErr.Message)
	}
}

func TestServersAddErrorServerType(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "server", "--serverHost", "http://comet", "--service", "Foo"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "serverType is missing", ngsiErr.Message)
	}
}

func TestServersAddErrorUnknownServerType(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "fiware", "--serverHost", "http://fiware", "--serverType", "fiware"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "serverType error: fiware (Comet, Cygnus, Iota, Keyrock, Perseo, QuantumLeap, WireCloud, Queryproxy, Regproxy, Tokenproxy)", ngsiErr.Message)
	}
}

func TestServersAddErrorAdd(t *testing.T) {
	c := setupTest([]string{"server", "add", "--host", "sth", "--serverType", "comet", "--serverHost", "fiware"})

	err := serverAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "host error: fiware", ngsiErr.Message)
	}
}

func TestServersUpdate(t *testing.T) {
	c := setupTest([]string{"server", "update", "--host", "comet", "--serverHost", "http://localhost"})

	err := serverUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		assert.Equal(t, "http://localhost", (*list)["comet"].ServerHost)
	}
}

func TestServersUpdateService(t *testing.T) {
	c := setupTest([]string{"server", "update", "--host", "comet", "--service", "Foo"})

	err := serverUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		assert.Equal(t, "foo", (*list)["comet"].Tenant)
	}
}

func TestServersUpdateErrorNotFound(t *testing.T) {
	c := setupTest([]string{"server", "update", "--host", "fiware", "--service", "Foo"})

	err := serverUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestServersUpdateErrorUpdateBroker(t *testing.T) {
	c := setupTest([]string{"server", "update", "--host", "ql", "--idmType", "fiware"})

	err := serverUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: fiware", ngsiErr.Message)
	}
}

func TestServersDelete(t *testing.T) {
	c := setupTest([]string{"server", "delete", "--host", "ql"})

	err := serverDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestServersDeleteNgsiType(t *testing.T) {
	c := setupTest([]string{"server", "delete", "--host", "ql", "--items", "idmType"})

	err := serverDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestServersDeleteErrorAlreadyExists(t *testing.T) {
	c := setupTest([]string{"server", "delete", "--host", "sth"})

	err := serverDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	}
}

func TestServersDeleteNoItem(t *testing.T) {
	c := setupTest([]string{"server", "delete", "--host", "ql", "--items", "noitem"})

	err := serverDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "delete error - noitem", ngsiErr.Message)
	}
}

func TestServersDeleteErrorUpdateServer(t *testing.T) {
	c := setupTest([]string{"server", "delete", "--host", "comet", "--items", "idmType"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("open error"), Filename: helper.StrPtr("ngsi-config.json")}

	err := serverDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestServersDeleteErrorReference(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"comet": {
				  "serverHost": "https://comet",
				  "serverType": "comet"
			},
			"comet2": {
				  "serverHost": "comet",
				  "serverType": "comet"
			}
		}
	}`
	c := setupTestWithConfig([]string{"server", "delete", "--host", "comet"}, conf)

	c.Ngsi.FileReader = &helper.MockFileLib{ReadFileData: []byte(conf)}

	err := serverDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "comet is referenced", ngsiErr.Message)
	}
}

func TestServersDeleteErrorDeleteServer(t *testing.T) {
	c := setupTest([]string{"server", "delete", "--host", "comet"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("open error"), Filename: helper.StrPtr("ngsi-config.json")}

	err := serverDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestPrintServerInfo(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	broker := ngsilib.Server{
		ServerType:   "comet",
		ServerHost:   "http://sth",
		Context:      "http://conetxt",
		SafeString:   "on",
		Tenant:       "openiot",
		Scope:        "/iot",
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

	printServerInfo(c.Ngsi, &broker, false)

	actual := helper.GetStdoutString(c)
	expected := "serverType comet\nserverHost http://sth\nFIWARE-Service openiot\nFIWARE-ServicePath /iot\nContext http://conetxt\nSafeString on\nIdmType keyrock\nIdmHost http://keyrock\nUsername fiware\nPassword ****\nClientID ********\nClientSecret ************\nXAuthToken false\nToken token\nAPIPath /path\n"
	assert.Equal(t, expected, actual)
}

func TestPrintServerInfoClearText(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	broker := ngsilib.Server{
		ServerType:   "comet",
		ServerHost:   "http://sth",
		Context:      "http://conetxt",
		SafeString:   "on",
		Tenant:       "openiot",
		Scope:        "/iot",
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

	printServerInfo(c.Ngsi, &broker, true)

	actual := helper.GetStdoutString(c)
	expected := "serverType comet\nserverHost http://sth\nFIWARE-Service openiot\nFIWARE-ServicePath /iot\nContext http://conetxt\nSafeString on\nIdmType keyrock\nIdmHost http://keyrock\nUsername fiware\nPassword 1234\nClientID clientid\nClientSecret clientsecret\nXAuthToken false\nToken token\nAPIPath /path\n"
	assert.Equal(t, expected, actual)
}

func TestPrintServerInfoError(t *testing.T) {
	c := helper.SetupTestInitCmd(nil)

	broker := ngsilib.Server{
		ServerType: "broker",
	}

	printServerInfo(c.Ngsi, &broker, false)

	actual := helper.GetStdoutString(c)
	expected := "server type error\n"
	assert.Equal(t, expected, actual)
}
