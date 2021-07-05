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
AUTHORS OR COPYRIGHT HOv2ERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package ngsicmd

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestSettingsList(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)

	err := settingsList(c)

	if assert.NoError(t, err) {
		assert.Equal(t, "", buf.String())
	} else {
		t.FailNow()
	}
}

func TestSettingsListPreviousArgsOff(t *testing.T) {
	ngsi, set, app, buf := setupTest()

	conf := `{
		"version": "1",
		"servers": {
		},
		"settings": {
			"usePreviousArgs": false,
			"syslog": "",
			"stderr": "",
			"logfile": "",
			"loglevel": "",
			"cachefile": "",
			"host": "",
			"tenant": "",
			"scope": "",
			"token": ""
		}
	}`

	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(conf)}
	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)

	err := settingsList(c)

	if assert.NoError(t, err) {
		assert.Equal(t, "PreviousArgs off\n", buf.String())
	} else {
		t.FailNow()
	}
}

func TestSettingsListAll(t *testing.T) {
	_, set, app, buf := setupTest()

	setupFlagString(set, "host")
	set.Bool("all", false, "doc")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--all"})

	err := settingsList(c)

	if assert.NoError(t, err) {
		assert.Equal(t, "Host: \nFIWARE-Service: \nFIWARE-ServicePath: \nToken: \nSyslog: \nStderr: \nLogFile: \nLogLevel: \n", buf.String())
	} else {
		t.FailNow()
	}
}

func TestSettingsListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := settingsList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsDelete(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--items=host,service,path,token,syslog,stderr,logfile,loglevel"})
	err := settingsDelete(c)

	assert.NoError(t, err)
}

func TestSettingsDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := settingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsDeleteErrorItems(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,syslog")
	c := cli.NewContext(app, set, nil)
	err := settingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "Required itmes not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsDeleteErrorItem(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=abc", "--items=item"})
	err := settingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "item not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsDeleteErrorSave(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("save error")}
	ngsi.PreviousArgs.UsePreviousArgs = true
	setupFlagString(set, "host,items")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--items=host,service,path,token,syslog,stderr,logfile,loglevel"})
	err := settingsDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "save error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsClear(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	err := settingsClear(c)

	assert.NoError(t, err)
}

func TestSettingsClearErrInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := settingsClear(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsClearErrorSave(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("save error")}
	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)

	err := settingsClear(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "save error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsPreviousArgs(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "on,off")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--on"})

	err := settingsPreviousArgs(c)

	assert.NoError(t, err)
}

func TestSettingsPreviousArgsErrInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,syslog")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})

	err := settingsPreviousArgs(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsPreviousArgsErrParam(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "on,off")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--on", "--off"})

	err := settingsPreviousArgs(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "specify either on or off", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestSettingsPreviousArgsErrParamOnOff(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	setupFlagBool(set, "on,off")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--on", "--off"})

	err := settingsPreviousArgs(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "specify either on or off", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestSettingsPreviousArgsErrorSave(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	ngsi.ConfigFile = &MockIoLib{OpenErr: errors.New("save error")}
	setupFlagString(set, "host")
	setupFlagBool(set, "on,off")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--on"})

	err := settingsPreviousArgs(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "save error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestPrintItem1(t *testing.T) {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "")

	printItem(buf, "", "", false)

	assert.Equal(t, "", buf.String())
}

func TestPrintItem2(t *testing.T) {
	buf := &bytes.Buffer{}
	printItem(buf, "host", "orion", false)

	assert.Equal(t, "host: orion\n", buf.String())
}

func TestPrintItem3(t *testing.T) {
	buf := &bytes.Buffer{}
	printItem(buf, "host", "", true)

	assert.Equal(t, "host: \n", buf.String())
}
