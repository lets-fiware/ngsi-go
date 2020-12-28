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
	"bytes"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func attrsRead(c *cli.Context) error {
	const funcName = "attrsRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	client.SetPath("/entities/" + id + "/attrs")

	args := []string{"type", "attrs", "metadata"}
	opts := []string{"keyValues", "values", "unique"}
	v := parseOptions(c, args, opts)
	client.SetQuery(v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, string(newBuf.Bytes()))
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func attrsAppend(c *cli.Context) error {
	const funcName = "attrsAppend"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	client.SetPath("/entities/" + id + "/attrs")

	args := []string{"type"}
	opts := []string{"append", "keyValues"}
	v := parseOptions(c, args, opts)
	client.SetQuery(v)

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func attrsUpdate(c *cli.Context) error {
	const funcName = "attrsUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	client.SetPath("/entities/" + id + "/attrs")

	args := []string{"type"}
	opts := []string{"append", "keyValues"}
	v := parseOptions(c, args, opts)
	client.SetQuery(v)

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func attrsReplace(c *cli.Context) error {
	const funcName = "attrsReplace"
	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	client.SetPath("/entities/" + id + "/attrs")

	args := []string{"type"}
	opts := []string{"append", "keyValues"}
	v := parseOptions(c, args, opts)
	client.SetQuery(v)

	client.SetHeader("Content-Type", "application/json")

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}
