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

func appendersList(c *cli.Context) error {
	const funcName = "appendersList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"cygnus"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	v := cygnusAdminSetParam(c)
	client.SetQuery(v)

	client.SetPath("/v1/admin/log/appenders")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		fmt.Fprint(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func appendersGet(c *cli.Context) error {
	const funcName = "appendersGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"cygnus"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("name") {
		return &ngsiCmdError{funcName, 3, "specify appender name", err}
	}

	v := cygnusAdminSetParam(c)
	client.SetQuery(v)

	client.SetPath("/v1/admin/log/appenders")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprint(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func appendersCreate(c *cli.Context) error {
	const funcName = "appendersCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"cygnus"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("data") {
		return &ngsiCmdError{funcName, 3, "specify data", nil}
	}

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	v := cygnusAdminSetParam(c)
	client.SetQuery(v)

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/v1/admin/log/appenders")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		fmt.Fprint(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func appendersUpdate(c *cli.Context) error {
	const funcName = "appendersUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"cygnus"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("name") {
		return &ngsiCmdError{funcName, 3, "specify name", nil}
	}

	if !c.IsSet("data") {
		return &ngsiCmdError{funcName, 4, "specify data", nil}
	}

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}

	v := cygnusAdminSetParam(c)
	client.SetQuery(v)

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/v1/admin/log/appenders")

	res, body, err := client.HTTPPut(b)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		fmt.Fprint(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func appendersDelete(c *cli.Context) error {
	const funcName = "appendersDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"cygnus"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("name") {
		return &ngsiCmdError{funcName, 3, "specify name", nil}
	}

	v := cygnusAdminSetParam(c)
	client.SetQuery(v)

	client.SetPath("/v1/admin/log/appenders")

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprint(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}
