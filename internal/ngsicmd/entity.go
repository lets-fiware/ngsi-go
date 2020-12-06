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
	"net/http"
	"net/url"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func entityCreate(c *cli.Context) error {
	const funcName = "entityCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	client.SetPath("/entities")

	var opts = []string{"keyValues", "upsert"}
	v := parseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetContentType()

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
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func entityRead(c *cli.Context) error {
	const funcName = "entityRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	client.SetPath("/entities/" + id)

	args := []string{"type", "attrs"}
	var opts = []string{"keyValues", "values", "unique", "sysAttrs"}
	v := parseOptions(c, args, opts)
	client.SetQuery(v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("error: %s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	}
	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}
func entityUpsert(c *cli.Context) error {
	const funcName = "entityUpsert"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	client.SetPath("/entities")

	v := url.Values{}
	options := "upsert"
	if c.IsSet("keyValues") {
		options = options + ",keyValues"
	}
	v.Set("options", options)
	client.SetQuery(&v)

	client.SetContentType()

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
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func entityDelete(c *cli.Context) error {

	const funcName = "entityDelete"
	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	client.SetPath("/entities/" + id)

	args := []string{"type"}
	v := parseOptions(c, args, nil)
	client.SetQuery(v)

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}
