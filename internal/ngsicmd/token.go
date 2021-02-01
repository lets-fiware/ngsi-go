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

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func tokenCommand(c *cli.Context) error {
	const funcName = "tokenCommand"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := ngsi.Host

	client, err := newClient(ngsi, c, false, nil)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	token, err := ngsi.TokenInfo(client)
	if err != nil {
		return &ngsiCmdError{funcName, 3, host + " has no token", err}
	}

	time := token.Expires - ngsi.TimeLib.NowUnix()
	if time < 0 {
		time = 0
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		token.Token.ExpiresIn = time
		b, err := ngsilib.JSONMarshal(token.Token)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(b))
		}
	} else if c.Bool("expires") {
		fmt.Fprintf(ngsi.StdWriter, "%d\n", time)
	} else {
		fmt.Fprintln(ngsi.StdWriter, client.Token)
	}

	return nil
}
