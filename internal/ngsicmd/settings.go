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
	"strings"

	"github.com/urfave/cli/v2"
)

func settingsList(c *cli.Context) error {
	const funcName = "settingsList"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	d := ngsi.GetPreviousArgs()

	all := c.Bool("all")

	printItem(ngsi.StdWriter, "Host", d.Host, all)
	printItem(ngsi.StdWriter, "FIWARE-Service", d.Tenant, all)
	printItem(ngsi.StdWriter, "FIWARE-ServicePath", d.Scope, all)
	printItem(ngsi.StdWriter, "Token", d.Token, all)
	printItem(ngsi.StdWriter, "Syslog", d.Syslog, all)
	printItem(ngsi.StdWriter, "Stderr", d.Stderr, all)
	printItem(ngsi.StdWriter, "LogFile", d.Logfile, all)
	printItem(ngsi.StdWriter, "LogLevel", d.Loglevel, all)

	return nil
}

func settingsDelete(c *cli.Context) error {
	const funcName = "settingsDelete"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if !c.IsSet("items") {
		return &ngsiCmdError{funcName, 2, "Required itmes not found", nil}
	}

	items := c.String("items")

	d := ngsi.GetPreviousArgs()

	for _, item := range strings.Split(items, ",") {
		item = strings.ToLower(item)
		switch item {
		default:
			return &ngsiCmdError{funcName, 3, item + " not found", nil}
		case "host":
			d.Host = ""
		case "service", "fiware-service", "tenant":
			d.Tenant = ""
		case "path", "fiware-servicepath", "scope":
			d.Scope = ""
		case "token":
			d.Token = ""
		case "syslog":
			d.Syslog = ""
		case "stderr":
			d.Stderr = ""
		case "logfile", "loglevel":
			d.Logfile = ""
			d.Loglevel = ""
		}
	}

	if d.Host == "" {
		d.Tenant = ""
		d.Scope = ""
	}

	err = ngsi.SavePreviousArgs()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	return nil
}

func settingsClear(c *cli.Context) error {
	const funcName = "settingsClear"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	d := ngsi.GetPreviousArgs()

	d.Host = ""
	d.Tenant = ""
	d.Scope = ""
	d.Token = ""
	d.Syslog = ""
	d.Stderr = ""
	d.Logfile = ""
	d.Loglevel = ""

	err = ngsi.SavePreviousArgs()
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return nil
}

func printItem(w io.Writer, k, v string, f bool) {
	if f || v != "" {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}
