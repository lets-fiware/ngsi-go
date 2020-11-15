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
	"strconv"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func typesList(c *cli.Context) error {
	const funcName = "typesList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "Only available on NGSIv2", nil}
	}

	page := 0
	count := 0
	limit := 10

	var types []string

	for {
		client.SetPath("/types")

		v := url.Values{}
		v.Set("options", "values,count")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		client.SetQuery(&v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 6, "ResultsCount error", err}
		}
		if count == 0 {
			break
		}

		var t []string
		err = ngsilib.JSONUnmarshal(body, &t)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		types = append(types, t...)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	if c.IsSet("json") {
		b, err := ngsilib.JSONMarshal(types)
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, string(b))
	} else {
		for _, e := range types {
			fmt.Fprintln(ngsi.StdWriter, e)
		}
	}

	return nil
}

func typeGet(c *cli.Context) error {
	const funcName = "typeGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "Only available on NGSIv2", nil}
	}

	client.SetPath("/types/" + c.String("type"))

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func typesCount(c *cli.Context) error {
	const funcName = "typesCount"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "Only available on NGSIv2", nil}
	}

	client.SetPath("/types")

	v := url.Values{}
	v.Set("options", "values,count")
	client.SetQuery(&v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	count, err := client.ResultsCount(res)
	if err != nil {
		return &ngsiCmdError{funcName, 6, "ResultsCount error", nil}
	}

	fmt.Fprintln(ngsi.StdWriter, strconv.Itoa(count))

	return nil
}
