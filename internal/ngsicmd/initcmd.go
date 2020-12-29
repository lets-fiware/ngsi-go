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
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"

	"github.com/urfave/cli/v2"
)

func initCmd(c *cli.Context, cmdName string, requiredHost bool) (*ngsilib.NGSI, error) {
	const funcName = "initCmd"

	ngsi := ngsilib.NewNGSI()

	if c.IsSet("batch") {
		b := c.Bool("batch")
		ngsi.BatchFlag = &b
	}

	var file *string
	if c.IsSet("config") {
		s := c.String("config")
		file = &s
	}

	err := ngsi.InitConfig(file)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	d := ngsi.GetPreviousArgs()

	s := "err"
	if d.Stderr != "" {
		s = d.Stderr
	}
	if c.IsSet("stderr") {
		s = c.String("stderr")
		d.Stderr = s
		ngsi.Updated = true
	}
	level, err := ngsilib.LogLevel(s)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 2, "stderr logLevel error", err}
	}
	ngsi.LogWriter = &ngsilib.LogWriter{Writer: os.Stderr, LogLevel: level}

	s = "off"
	if d.Syslog != "" {
		s = d.Syslog
	}
	if c.IsSet("syslog") {
		s = c.String("syslog")
		d.Syslog = s
		ngsi.Updated = true
	}
	level, err = ngsilib.LogLevel(s)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 3, "syslog logLevel error", err}
	}
	if level != ngsilib.LogOff && ngsi.OsType != "windows" {
		syslog, err := ngsi.SyslogLib.New()
		if err != nil {
			return nil, &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		syslogWriter := ngsilib.LogWriter{Writer: syslog, LogLevel: level}
		ngsi.LogWriter = io.MultiWriter(ngsi.LogWriter, &syslogWriter)
	}

	if c.IsSet("margin") {
		margin := c.Int64("margin")
		if margin > 600 || margin < 10 {
			margin = 180
		}
		ngsi.Margin = margin
	}

	if c.IsSet("timeout") {
		timeout := c.Int("timeout")
		if timeout > 600 || timeout < 10 {
			timeout = 60
		}
		ngsi.Timeout = time.Duration(timeout) * time.Second
	}

	if c.IsSet("maxCount") {
		maxsize := c.Int("maxCount")
		if maxsize > 3000 || maxsize < 1 {
			maxsize = 100
		}
		ngsi.Maxsize = maxsize
	}

	c.App.Writer = ngsi.StdWriter
	c.App.ErrWriter = ngsi.LogWriter

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s (git_hash:%s)\n", Version, Revision))
	ngsi.Logging(ngsilib.LogInfo, cmdName+"\n")

	ngsi.Host = c.String("host")
	if ngsi.Host == "" {
		ngsi.Host = d.Host
	}
	if d.Host != ngsi.Host {
		d.Host = ngsi.Host
		d.Tenant = ""
		d.Scope = ""
		ngsi.Updated = true
	}

	ngsi.Destination = c.String("destination")

	if requiredHost && ngsi.Host == "" {
		return nil, &ngsiCmdError{funcName, 5, "Required host not found", err}
	}

	var cacheFile *string
	if d.CacheFile != "" {
		cacheFile = &d.CacheFile
	}
	if c.IsSet("cacheFile") {
		d.CacheFile = c.String("cacheFile")
		cacheFile = &d.CacheFile
		ngsi.Updated = true
	}

	err = ngsi.InitTokenMgr(cacheFile)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 6, err.Error(), err}
	}

	return ngsi, nil
}
