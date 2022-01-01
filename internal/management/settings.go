/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"fmt"
	"io"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func settingsList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	d := ngsi.GetPreviousArgs()

	all := c.Bool("all")

	if !d.UsePreviousArgs {
		fmt.Fprintln(ngsi.StdWriter, "PreviousArgs off")
	}
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

func settingsDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "settingsDelete"

	if !c.IsSet("items") {
		return ngsierr.New(funcName, 1, "Required itmes not found", nil)
	}

	items := c.String("items")

	d := ngsi.GetPreviousArgs()

	for _, item := range strings.Split(items, ",") {
		item = strings.ToLower(item)
		switch item {
		default:
			return ngsierr.New(funcName, 2, item+" not found", nil)
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

	err := ngsi.SavePreviousArgs()
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	return nil
}

func settingsClear(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "settingsClear"

	d := ngsi.GetPreviousArgs()

	d.Host = ""
	d.Tenant = ""
	d.Scope = ""
	d.Token = ""
	d.Syslog = ""
	d.Stderr = ""
	d.Logfile = ""
	d.Loglevel = ""

	err := ngsi.SavePreviousArgs()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	return nil
}

func settingsPreviousArgs(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "settingsPreviousArgs"

	on := c.Bool("on")
	off := c.Bool("off")

	if on == off {
		return ngsierr.New(funcName, 1, "specify either on or off", nil)
	}

	d := ngsi.GetPreviousArgs()

	d.UsePreviousArgs = on
	d.Host = ""
	d.Tenant = ""
	d.Scope = ""
	d.Token = ""
	d.Syslog = ""
	d.Stderr = ""
	d.Logfile = ""
	d.Loglevel = ""

	err := ngsi.SavePreviousArgs()
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	return nil
}

func printItem(w io.Writer, k, v string, f bool) {
	if f || v != "" {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}
