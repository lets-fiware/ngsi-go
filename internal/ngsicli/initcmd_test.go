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

package ngsicli

import (
	"errors"
	"testing"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestInitCmd(t *testing.T) {
	_ = setupTestInitNGSI()

	c := &Context{}

	_, err := InitCmd(c)

	assert.NoError(t, err)
}

func TestInitCmdBatchFlagTrue(t *testing.T) {
	_ = setupTestInitNGSI()

	f := batchFlag.Copy(true)
	err := f.SetValue(true)
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	ngsi, err := InitCmd(c)

	if assert.NoError(t, err) {
		assert.Equal(t, true, *ngsi.BatchFlag)
	}
}

func TestInitCmdBatchFlagFalse(t *testing.T) {
	_ = setupTestInitNGSI()

	f := batchFlag.Copy(true)
	err := f.SetValue(false)
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	ngsi, err := InitCmd(c)

	if assert.NoError(t, err) {
		assert.Equal(t, false, *ngsi.BatchFlag)
	}
}

func TestInitCmdConfig(t *testing.T) {
	_ = setupTestInitNGSI()

	f := configFlag.Copy(true)
	err := f.SetValue("config.json")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	ngsi, err := InitCmd(c)

	if assert.NoError(t, err) {
		assert.Equal(t, "config.json", *ngsi.ConfigFile.FileName())
	}

}

func TestInitCmdinsecureSkipVerifyFlagFalse(t *testing.T) {
	_ = setupTestInitNGSI()

	f := insecureSkipVerifyFlag.Copy(true)
	err := f.SetValue(false)
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	ngsi, err := InitCmd(c)

	if assert.NoError(t, err) {
		assert.Equal(t, false, ngsi.InsecureSkipVerify)
	}
}

func TestInitCmdinsecureSkipVerifyFlagTrue(t *testing.T) {
	_ = setupTestInitNGSI()

	f := insecureSkipVerifyFlag.Copy(true)
	err := f.SetValue(true)
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	ngsi, err := InitCmd(c)

	if assert.NoError(t, err) {
		assert.Equal(t, true, ngsi.InsecureSkipVerify)
	}
}

func TestInitCmdErrorInitConfig(t *testing.T) {
	ngsi := setupTestInitNGSI()

	ngsi.ConfigFile = &MockIoLib{HomeDir: errors.New("InitConfig error")}
	c := &Context{}

	_, err := InitCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "InitConfig error", ngsiErr.Message)
	}
}

func TestInitCmdErrorInitStdErrOption(t *testing.T) {
	_ = setupTestInitNGSI()

	f := stderrFlag.Copy(true)
	err := f.SetValue("unknown")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	_, err = InitCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "stderr logLevel error", ngsiErr.Message)
	}
}

func TestInitCmdErrorInitSyslogOption(t *testing.T) {
	_ = setupTestInitNGSI()

	f := syslogFlag.Copy(true)
	err := f.SetValue("unknown")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	_, err = InitCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestInitCmdErrorInitTokenMgr(t *testing.T) {
	ngsi := setupTestInitNGSI()

	ngsi.CacheFile = &MockIoLib{HomeDir: errors.New("InitTOkenMgr error")}

	c := &Context{}

	_, err := InitCmd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "InitTOkenMgr error", ngsiErr.Message)
	}
}

func TestInitStdErrOption(t *testing.T) {
	ngsi := setupTestInitNGSI()

	f := stderrFlag.Copy(true)
	err := f.SetValue("err")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	prevArgs := &ngsilib.Settings{}

	err = initStdErrOption(ngsi, c, prevArgs)

	if assert.NoError(t, err) {
		assert.Equal(t, "err", prevArgs.Stderr)
	}
}

func TestInitStdErrOptionPrevArgs(t *testing.T) {
	ngsi := setupTestInitNGSI()

	c := &Context{}

	prevArgs := &ngsilib.Settings{Stderr: "debug"}

	err := initStdErrOption(ngsi, c, prevArgs)

	assert.NoError(t, err)
}

func TestInitStdErrOptionStdErr(t *testing.T) {
	ngsi := setupTestInitNGSI()

	f := stderrFlag.Copy(true)
	err := f.SetValue("debug")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	prevArgs := &ngsilib.Settings{}

	err = initStdErrOption(ngsi, c, prevArgs)

	if assert.NoError(t, err) {
		assert.Equal(t, "debug", prevArgs.Stderr)
		assert.Equal(t, true, ngsi.Updated)
	}
}

func TestInitStdErrOptionErrorLogLevel(t *testing.T) {
	ngsi := setupTestInitNGSI()

	f := stderrFlag.Copy(true)
	err := f.SetValue("unknown")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	prevArgs := &ngsilib.Settings{}

	err = initStdErrOption(ngsi, c, prevArgs)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "stderr logLevel error", ngsiErr.Message)
	}
}
func TestInitSyslogOption(t *testing.T) {
	ngsi := setupTestInitNGSI()

	c := &Context{}

	prevArgs := &ngsilib.Settings{}

	err := initSyslogOption(ngsi, c, prevArgs)

	assert.NoError(t, err)
}

func TestInitSyslogOptionPrevArgs(t *testing.T) {
	ngsi := setupTestInitNGSI()

	c := &Context{}

	prevArgs := &ngsilib.Settings{Syslog: "info"}

	err := initSyslogOption(ngsi, c, prevArgs)

	assert.NoError(t, err)
}

func TestInitSyslogOptionSyslogFlag(t *testing.T) {
	ngsi := setupTestInitNGSI()

	f := syslogFlag.Copy(true)
	err := f.SetValue("debug")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	prevArgs := &ngsilib.Settings{}

	err = initSyslogOption(ngsi, c, prevArgs)

	if assert.NoError(t, err) {
		assert.Equal(t, "debug", prevArgs.Syslog)
		assert.Equal(t, true, ngsi.Updated)
	}
}

func TestInitSyslogOptionError(t *testing.T) {
	ngsi := setupTestInitNGSI()

	f := syslogFlag.Copy(true)
	err := f.SetValue("unknown")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	prevArgs := &ngsilib.Settings{}

	err = initSyslogOption(ngsi, c, prevArgs)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestInitSyslogOptionErrorSyslogLibNew(t *testing.T) {
	ngsi := setupTestInitNGSI()

	f := syslogFlag.Copy(true)
	err := f.SetValue("debug")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	prevArgs := &ngsilib.Settings{}

	ngsi.SyslogLib = &MockSyslogLib{Err: errors.New("SyslogLib")}

	err = initSyslogOption(ngsi, c, prevArgs)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "SyslogLib", ngsiErr.Message)
	}
}

func TestInitHiddenOptionsMargin(t *testing.T) {
	ngsi := setupTestInitNGSI()

	cases := []struct {
		expected int64
		arg      int64
	}{
		{expected: 180, arg: -1},
		{expected: 180, arg: 0},
		{expected: 180, arg: 9},
		{expected: 10, arg: 10},
		{expected: 100, arg: 100},
		{expected: 600, arg: 600},
		{expected: 180, arg: 601},
	}

	f := marginFlag.Copy(true)
	for _, cc := range cases {
		err := f.SetValue(cc.arg)
		assert.NoError(t, err)

		c := &Context{Flags: []Flag{f}}

		initHiddenOptions(ngsi, c)

		assert.Equal(t, cc.expected, ngsi.Margin)
	}
}

func TestInitHiddenOptionsTimeout(t *testing.T) {
	ngsi := setupTestInitNGSI()

	cases := []struct {
		expected int64
		arg      int64
	}{
		{expected: 60, arg: -1},
		{expected: 60, arg: 0},
		{expected: 60, arg: 9},
		{expected: 10, arg: 10},
		{expected: 100, arg: 100},
		{expected: 600, arg: 600},
		{expected: 60, arg: 601},
	}

	f := timeOutFlag.Copy(true)
	for _, cc := range cases {
		err := f.SetValue(cc.arg)
		assert.NoError(t, err)

		c := &Context{Flags: []Flag{f}}

		initHiddenOptions(ngsi, c)

		assert.Equal(t, time.Duration(cc.expected)*time.Second, ngsi.Timeout)
	}
}

func TestInitHiddenOptionsMaxCount(t *testing.T) {
	ngsi := setupTestInitNGSI()

	cases := []struct {
		expected int64
		arg      int64
	}{
		{expected: 100, arg: -1},
		{expected: 100, arg: 0},
		{expected: 1, arg: 1},
		{expected: 10, arg: 10},
		{expected: 100, arg: 100},
		{expected: 3000, arg: 3000},
		{expected: 100, arg: 3001},
	}

	f := maxCountFlag.Copy(true)
	for _, cc := range cases {
		err := f.SetValue(cc.arg)
		assert.NoError(t, err)

		c := &Context{Flags: []Flag{f}}

		initHiddenOptions(ngsi, c)

		assert.Equal(t, cc.expected, ngsi.Maxsize)
	}
}

func TestInitCacheFileOption(t *testing.T) {
	ngsi := setupTestInitNGSI()

	c := &Context{}

	prevArgs := &ngsilib.Settings{}

	actual := initCacheFileOption(ngsi, c, prevArgs)

	assert.Equal(t, (*string)(nil), actual)
}

func TestInitCacheFileOptionPrevArgs(t *testing.T) {
	ngsi := setupTestInitNGSI()

	c := &Context{}

	prevArgs := &ngsilib.Settings{CacheFile: "ngsi-cache.json"}

	actual := initCacheFileOption(ngsi, c, prevArgs)

	assert.Equal(t, "ngsi-cache.json", *actual)
}

func TestInitCacheFileOptionCacheFlag(t *testing.T) {
	ngsi := setupTestInitNGSI()

	f := cacheFlag.Copy(true)
	err := f.SetValue("ngsi-cache.json")
	assert.NoError(t, err)

	c := &Context{Flags: []Flag{f}}

	prevArgs := &ngsilib.Settings{}

	actual := initCacheFileOption(ngsi, c, prevArgs)

	assert.Equal(t, "ngsi-cache.json", *actual)
}
