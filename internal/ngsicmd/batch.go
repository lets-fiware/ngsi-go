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
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func batch(c *cli.Context, mode string) error {
	const funcName = "batch"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		switch mode {
		case "create":
			return batchCreate(c, ngsi, client)
		case "update":
			return batchUpdate(c, ngsi, client)
		case "upsert":
			return batchUpsert(c, ngsi, client)
		case "delete":
			return batchDelete(c, ngsi, client)
		}
	} else {
		switch mode {
		case "create":
			return opUpdate(c, ngsi, client, "append_strict")
		case "update":
			return opUpdate(c, ngsi, client, "update")
		case "upsert":
			return opUpdate(c, ngsi, client, "append")
		case "replace":
			return opUpdate(c, ngsi, client, "replace")
		case "delete":
			return opUpdate(c, ngsi, client, "delete")
		}
	}
	return &ngsiCmdError{funcName, 3, "error: " + mode, nil}
}

func batchCreate(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchCreate"

	client.SetPath("/entityOperations/create")

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func batchUpdate(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchUpdate"

	client.SetPath("/entityOperations/update")

	var opts = []string{"noOverwrite", "replace"}
	v := parseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}
func batchUpsert(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchUpsert"

	client.SetPath("/entityOperations/upsert")

	var opts = []string{"replace", "update"}
	v := parseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func batchDelete(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchDelete"

	client.SetPath("/entityOperations/delete")

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}
