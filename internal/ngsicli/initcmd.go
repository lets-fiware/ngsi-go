/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func InitCmd(c *Context) (*ngsilib.NGSI, error) {
	const funcName = "InitCmd"

	ngsi := ngsilib.NewNGSI()

	if c.IsSet("batch") {
		b := c.Bool("batch")
		ngsi.BatchFlag = &b
	}

	if c.IsSet("configDir") && (c.IsSet("config") || c.IsSet("cache")) {
		return nil, ngsierr.New(funcName, 1, "configDir cannot be specified with config or cache at the same time", nil)
	}
	var file *string
	if c.IsSet("config") {
		s := c.String("config")
		file = &s
	}
	if c.IsSet("configDir") {
		s := c.String("configDir")
		ngsi.ConfigDir = &s
	}

	err := ngsi.InitConfig(file)
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}

	prevArgs := ngsi.GetPreviousArgs()

	err = initStdErrOption(ngsi, c, prevArgs)
	if err != nil {
		return nil, ngsierr.New(funcName, 3, err.Error(), err)
	}

	err = initSyslogOption(ngsi, c, prevArgs)
	if err != nil {
		return nil, ngsierr.New(funcName, 4, err.Error(), err)
	}

	initHiddenOptions(ngsi, c)

	ngsi.InsecureSkipVerify = c.Bool("insecureSkipVerify")

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s (git_hash:%s)\n", Version, Revision))

	tokenCache := initCacheFileOption(ngsi, c, prevArgs)
	err = ngsi.InitTokenMgr(tokenCache)
	if err != nil {
		return nil, ngsierr.New(funcName, 5, err.Error(), err)
	}

	return ngsi, nil
}

func initStdErrOption(ngsi *ngsilib.NGSI, c *Context, prevArgs *ngsilib.Settings) error {
	const funcName = "initStrErrOption"

	s := "err"
	if prevArgs.Stderr != "" {
		s = prevArgs.Stderr
	}
	if c.IsSet("stderr") {
		s = c.String("stderr")
		prevArgs.Stderr = s
		ngsi.Updated = true
	}
	level, err := ngsilib.LogLevel(s)
	if err != nil {
		return ngsierr.New(funcName, 1, "stderr logLevel error", err)
	}
	ngsi.LogWriter = &ngsilib.LogWriter{Writer: os.Stderr, LogLevel: level}

	return nil
}

func initSyslogOption(ngsi *ngsilib.NGSI, c *Context, prevArgs *ngsilib.Settings) error {
	const funcName = "initSyslogOption"

	s := "off"
	if prevArgs.Syslog != "" {
		s = prevArgs.Syslog
	}

	if c.IsSet("syslog") {
		s = c.String("syslog")
		prevArgs.Syslog = s
		ngsi.Updated = true
	}

	level, err := ngsilib.LogLevel(s)
	if err != nil {
		return ngsierr.New(funcName, 1, "syslog logLevel error", err)
	}

	if level != ngsilib.LogOff && ngsi.OsType != "windows" {
		syslog, err := ngsi.SyslogLib.New()
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		syslogWriter := ngsilib.LogWriter{Writer: syslog, LogLevel: level}
		ngsi.LogWriter = io.MultiWriter(ngsi.LogWriter, &syslogWriter)
	}

	return nil
}

func initHiddenOptions(ngsi *ngsilib.NGSI, c *Context) {
	if c.IsSet("margin") {
		margin := c.Int64("margin")
		if margin > 600 || margin < 10 {
			margin = 180
		}
		ngsi.Margin = margin
	}

	if c.IsSet("timeout") {
		timeout := c.Int64("timeout")
		if timeout > 600 || timeout < 10 {
			timeout = 60
		}
		ngsi.Timeout = time.Duration(timeout) * time.Second
	}

	if c.IsSet("maxCount") {
		maxsize := c.Int64("maxCount")
		if maxsize > 3000 || maxsize < 1 {
			maxsize = 100
		}
		ngsi.Maxsize = maxsize
	}
}

func initCacheFileOption(ngsi *ngsilib.NGSI, c *Context, prevArgs *ngsilib.Settings) *string {
	var cache *string

	if prevArgs.CacheFile != "" {
		cache = &prevArgs.CacheFile
	}

	if c.IsSet("cache") {
		prevArgs.CacheFile = c.String("cache")
		cache = &prevArgs.CacheFile
		ngsi.Updated = true
	}

	return cache
}
