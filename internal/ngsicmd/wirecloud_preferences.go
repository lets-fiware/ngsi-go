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
	"bytes"
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func wireCloudPreferencesGet(c *cli.Context) error {
	const funcName = "workspacesList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, nil)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.ServerType == "keyrock" {
		return &ngsiCmdError{funcName, 3, "not supported by keyrock", nil}
	}

	switch client.Server.ServerType {
	default:
		client.SetPath("/version")
	case "cygnus":
		client.SetPath("/v1/version")
	case "iota":
		client.SetPath("/iot/about")
	case "perseo-core":
		client.SetPath("/perseo-core/version")
	case "wirecloud":
		client.SetPath("/api/features")
	}

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") && client.Server.ServerType != "perseo-core" {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}
