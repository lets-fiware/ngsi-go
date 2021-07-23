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

func TestServersList(t *testing.T) {
	_, set, app, buf := setupTest()

	c := cli.NewContext(app, set, nil)

	err := serverList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "comet cygnus geoproxy iota keyrock perseo perseo-core ql regproxy tokenproxy wirecloud\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestServersListHost(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=ql"})
	c := cli.NewContext(app, set, nil)

	err := serverList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "serverType quantumleap\nserverHost https://quantumleap\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestServersListJSON(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json"})

	err := serverList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"comet\":{\"serverType\":\"comet\",\"serverHost\":\"https://comet\"},\"cygnus\":{\"serverType\":\"cygnus\",\"serverHost\":\"https://cygnus\"},\"geoproxy\":{\"serverType\":\"geoproxy\",\"serverHost\":\"https://geoproxy\"},\"iota\":{\"serverType\":\"iota\",\"serverHost\":\"https://iota\"},\"keyrock\":{\"serverType\":\"keyrock\",\"serverHost\":\"https://keyrock\"},\"perseo\":{\"serverType\":\"perseo\",\"serverHost\":\"https://perseo\"},\"perseo-core\":{\"serverType\":\"perseo-core\",\"serverHost\":\"https://perseo-core\"},\"ql\":{\"serverType\":\"quantumleap\",\"serverHost\":\"https://quantumleap\"},\"regproxy\":{\"serverType\":\"regproxy\",\"serverHost\":\"https://regproxy\"},\"tokenproxy\":{\"serverType\":\"tokenproxy\",\"serverHost\":\"https://tokenproxy\"},\"wirecloud\":{\"serverType\":\"wirecloud\",\"serverHost\":\"https://wirecloud\"}}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestServersListJSONPretty(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json", "--pretty"})

	err := serverList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"comet\": {\n    \"serverType\": \"comet\",\n    \"serverHost\": \"https://comet\"\n  },\n  \"cygnus\": {\n    \"serverType\": \"cygnus\",\n    \"serverHost\": \"https://cygnus\"\n  },\n  \"geoproxy\": {\n    \"serverType\": \"geoproxy\",\n    \"serverHost\": \"https://geoproxy\"\n  },\n  \"iota\": {\n    \"serverType\": \"iota\",\n    \"serverHost\": \"https://iota\"\n  },\n  \"keyrock\": {\n    \"serverType\": \"keyrock\",\n    \"serverHost\": \"https://keyrock\"\n  },\n  \"perseo\": {\n    \"serverType\": \"perseo\",\n    \"serverHost\": \"https://perseo\"\n  },\n  \"perseo-core\": {\n    \"serverType\": \"perseo-core\",\n    \"serverHost\": \"https://perseo-core\"\n  },\n  \"ql\": {\n    \"serverType\": \"quantumleap\",\n    \"serverHost\": \"https://quantumleap\"\n  },\n  \"regproxy\": {\n    \"serverType\": \"regproxy\",\n    \"serverHost\": \"https://regproxy\"\n  },\n  \"tokenproxy\": {\n    \"serverType\": \"tokenproxy\",\n    \"serverHost\": \"https://tokenproxy\"\n  },\n  \"wirecloud\": {\n    \"serverType\": \"wirecloud\",\n    \"serverHost\": \"https://wirecloud\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestServersListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := serverList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersListErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=sth"})
	c := cli.NewContext(app, set, nil)

	err := serverList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersListErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "json")
	setJSONEncodeErr(ngsi, 2)

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json"})

	err := serverList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersListErrorJSONPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--json", "--pretty"})

	setJSONIndentError(ngsi)

	err := serverList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersGet(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet"})

	err := serverGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "serverType comet\nserverHost https://comet\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestServersGetJSON(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--json"})

	err := serverGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\"serverType\":\"comet\",\"serverHost\":\"https://comet\"}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestServersGetJSONPretty(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--json", "--pretty"})

	err := serverGet(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "{\n  \"serverType\": \"comet\",\n  \"serverHost\": \"https://comet\"\n}\n"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestServersGetErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := serverGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersGetErrorHostNotFound(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := serverGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersGetErrorServerListErrorJSON(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=sth", "--json"})

	err := serverGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersGetErrorJSONPretty(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "json,pretty")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--json", "--pretty"})

	setJSONIndentError(ngsi)

	err := serverGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersGetErrorServerList(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=sth"})

	err := serverGet(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersAdd(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,serverHost,serverType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=sth", "--serverHost=http://sth", "--serverType=comet"})

	err := serverAdd(c)

	if assert.NoError(t, err) {
		list := ngsi.AllServersList()
		assert.Equal(t, "http://sth", (*list)["sth"].ServerHost)
		assert.Equal(t, "comet", (*list)["sth"].ServerType)
	}
}

func TestServersAddKeyrock(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,serverHost,serverType,username,password")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=idm", "--serverHost=http://keyrock", "--serverType=keyrock", "--username=fiware", "--password=1234"})

	err := serverAdd(c)

	if assert.NoError(t, err) {
		list := ngsi.AllServersList()
		assert.Equal(t, "http://keyrock", (*list)["idm"].ServerHost)
		assert.Equal(t, "keyrock", (*list)["idm"].ServerType)
	}
}

func TestServersAddTenant(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,ngsiType,serverHost,serverType,service")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=sth", "--serverHost=http://sth", "--serverType=comet", "--service=Foo"})

	err := serverAdd(c)

	if assert.NoError(t, err) {
		list := ngsi.AllServersList()
		v := (*list)["sth"]
		assert.Equal(t, "http://sth", (*list)["sth"].ServerHost)
		assert.Equal(t, "comet", (*list)["sth"].ServerType)
		assert.Equal(t, "foo", v.Tenant)
	}
}

func TestServersAddErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersAddErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersAddErrorNameString(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=@ql"})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "name error @ql", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersAddErrorAlreadyExists(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet"})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "comet already exists", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestServersAddErrorServerHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,serverType,serverHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion-v2"})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "serverHost is missing", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersAddErrorServerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,serverType,serverHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=sth", "--serverHost=http://sth"})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "serverType is missing", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersAddErrorUnknownServerType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,serverType,serverHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=fiware", "--serverHost=http://fiware", "--serverType=fiware"})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "serverType error: fiware (Comet, Cygnus, Iota, Keyrock, Perseo, QuantumLeap, WireCloud, Geoproxy, Regproxy, Tokenproxy)", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersAddErrorAdd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,serverType,serverHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=sth", "--serverType=comet", "--serverHost=fiware"})

	err := serverAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "host error: fiware", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersUpdate(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,serverHost")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--serverHost=http://localhost"})

	err := serverUpdate(c)

	if assert.NoError(t, err) {
		list := ngsi.AllServersList()
		assert.Equal(t, "http://localhost", (*list)["comet"].ServerHost)
	}
}

func TestServersUpdateService(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,service")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet", "--service=Foo"})

	err := serverUpdate(c)

	if assert.NoError(t, err) {
		list := ngsi.AllServersList()
		assert.Equal(t, "foo", (*list)["comet"].Tenant)
	}
}

func TestServersUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := serverUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersUpdateErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := serverUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersUpdateErrorNotFound(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=fiware"})

	err := serverUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersUpdateErrorCreateBroker(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,idmType")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--idmType=fiware"})

	err := serverUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unknown idm type: fiware", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersDelete(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql"})

	err := serverDelete(c)

	assert.NoError(t, err)
}

func TestServersDeleteNgsiType(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--items=idmType"})

	err := serverDelete(c)

	assert.NoError(t, err)
}

func TestServersDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := serverDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersDeleteErrorHost(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host="})

	err := serverDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "required host not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersDeleteErrorAlreadyExists(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=sth"})

	err := serverDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "sth not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersDeleteNoItem(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=ql", "--items=noitem"})

	err := serverDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "delete error - noitem", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersDeleteErrorUpdateBroker(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "host,items")
	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("open error")}
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--items=idmType"})

	err := serverDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestServersDeleteErrorReference(t *testing.T) {
	ngsi, set, app, _ := setupTest()

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

	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=comet"})

	err := serverDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "comet is referenced", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPrintServerInfo(t *testing.T) {
	ngsi, _, _, buf := setupTest()

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

	printServerInfo(ngsi, &broker, false)

	actual := buf.String()
	expected := "serverType comet\nserverHost http://sth\nFIWARE-Service openiot\nFIWARE-ServicePath /iot\nContext http://conetxt\nSafeString on\nIdmType keyrock\nIdmHost http://keyrock\nUsername fiware\nPassword ****\nClientID ********\nClientSecret ************\nXAuthToken false\nToken token\nAPIPath /path\n"
	assert.Equal(t, expected, actual)
}

func TestPrintServerInfoClearText(t *testing.T) {
	ngsi, _, _, buf := setupTest()

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

	printServerInfo(ngsi, &broker, true)

	actual := buf.String()
	expected := "serverType comet\nserverHost http://sth\nFIWARE-Service openiot\nFIWARE-ServicePath /iot\nContext http://conetxt\nSafeString on\nIdmType keyrock\nIdmHost http://keyrock\nUsername fiware\nPassword 1234\nClientID clientid\nClientSecret clientsecret\nXAuthToken false\nToken token\nAPIPath /path\n"
	assert.Equal(t, expected, actual)
}

func TestPrintServerInfoError(t *testing.T) {
	ngsi, _, _, buf := setupTest()

	broker := ngsilib.Server{
		ServerType: "broker",
	}

	printServerInfo(ngsi, &broker, false)

	actual := buf.String()
	expected := "server type error\n"
	assert.Equal(t, expected, actual)
}
