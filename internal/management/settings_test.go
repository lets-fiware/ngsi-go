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
AUTHORS OR COPYRIGHT HOv2ERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package management

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestSettingsList(t *testing.T) {
	c := setupTest([]string{"settings", "list"})

	err := settingsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		assert.Equal(t, "", actual)
	}
}

func TestSettingsListPreviousArgsOff(t *testing.T) {
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
	c := setupTestWithConfig([]string{"settings", "list"}, conf)

	c.Ngsi.FileReader = &helper.MockFileLib{ReadFileData: []byte(conf)}

	err := settingsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		assert.Equal(t, "PreviousArgs off\n", actual)
	}
}

func TestSettingsListAll(t *testing.T) {
	c := setupTest([]string{"settings", "list", "--all"})

	err := settingsList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "Host: \nFIWARE-Service: \nFIWARE-ServicePath: \nToken: \nSyslog: \nStderr: \nLogFile: \nLogLevel: \n"
		assert.Equal(t, expected, actual)
	}
}

func TestSettingsDelete(t *testing.T) {
	c := setupTest([]string{"settings", "delete", "--items", "host,service,path,token,syslog,stderr,logfile,loglevel"})

	err := settingsDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSettingsDeleteErrorItems(t *testing.T) {
	c := setupTest([]string{"settings", "delete"})

	err := settingsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "Required itmes not found", ngsiErr.Message)
	}
}

func TestSettingsDeleteErrorItem(t *testing.T) {
	c := setupTest([]string{"settings", "delete", "--items", "item"})

	err := settingsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "item not found", ngsiErr.Message)
	}
}

func TestSettingsDeleteErrorSave(t *testing.T) {
	c := setupTest([]string{"settings", "delete", "--items", "host,service,path,token,syslog,stderr,logfile,loglevel"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("save error"), Filename: helper.StrPtr("ngsi-config.json")}
	c.Ngsi.PreviousArgs.UsePreviousArgs = true

	err := settingsDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "save error", ngsiErr.Message)
	}
}

func TestSettingsClear(t *testing.T) {
	c := setupTest([]string{"settings", "clear"})

	err := settingsClear(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSettingsClearErrorSave(t *testing.T) {
	c := setupTest([]string{"settings", "clear"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("save error"), Filename: helper.StrPtr("ngsi-config.json")}

	err := settingsClear(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "save error", ngsiErr.Message)
	}
}

func TestSettingsPreviousArgs(t *testing.T) {
	c := setupTest([]string{"settings", "previousArgs", "--on"})

	err := settingsPreviousArgs(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestSettingsPreviousArgsErrParam(t *testing.T) {
	c := setupTest([]string{"settings", "previousArgs"})

	err := settingsPreviousArgs(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "specify either on or off", ngsiErr.Message)
	}
}

func TestSettingsPreviousArgsErrParamOnOff(t *testing.T) {
	c := setupTest([]string{"settings", "previousArgs", "--on", "--off"})

	err := settingsPreviousArgs(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "specify either on or off", ngsiErr.Message)
	}
}
func TestSettingsPreviousArgsErrorSave(t *testing.T) {
	c := setupTest([]string{"settings", "previousArgs", "--on"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("save error"), Filename: helper.StrPtr("ngsi-config.json")}

	err := settingsPreviousArgs(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "save error", ngsiErr.Message)
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
