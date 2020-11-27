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

	"github.com/urfave/cli/v2"
)

var debugCmd = cli.Command{
	Name:     "debug",
	Category: "CONVENIENCE",
	Usage:    "test",
	Hidden:   true,
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Action: func(c *cli.Context) error {
		return debugCommand(c)
	},
}

func debugCommand(c *cli.Context) error {
	const funcName = "debugCommand"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	fmt.Fprintf(ngsi.StdWriter, "config file: %s\n", *ngsi.ConfigFile.FileName())
	fmt.Fprintf(ngsi.StdWriter, "cache file: %s\n", *ngsi.CacheFile.FileName())

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	fmt.Fprintf(ngsi.StdWriter, "Host: %s\n", ngsi.Host)
	fmt.Fprintf(ngsi.StdWriter, "Tenant: %s\n", client.Tenant)
	fmt.Fprintf(ngsi.StdWriter, "Scope: %s\n", client.Scope)

	return nil
}
