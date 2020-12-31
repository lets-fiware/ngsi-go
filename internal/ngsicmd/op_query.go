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

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func opQuery(c *cli.Context) error {
	const funcName = "opQuery"

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
	limit := 100

	verbose := c.IsSet("verbose")
	if c.Bool("pretty") {
		verbose = true
	}
	lines := c.Bool("lines")

	buf := jsonBuffer{}
	if verbose {
		buf.bufferOpen(ngsi.StdWriter)
	}

	for {
		client.SetPath("/op/query")

		var args = []string{"orderBy"}
		var opts = []string{"count", "keyValues", "values", "unique"}
		v := parseOptions(c, args, opts)
		if c.IsSet("count") {
			v.Set("limit", "1")
			v.Set("options", "count")
		} else {
			options := v.Get("options")
			if options == "" {
				options = "count"
			} else {
				options = options + ",count"
			}
			v.Set("options", options)
			v.Set("limit", fmt.Sprintf("%d", limit))
			v.Set("offset", fmt.Sprintf("%d", page*limit))
		}
		client.SetQuery(v)

		client.SetHeader("Content-Type", "application/json")

		b, err := readAll(c, ngsi)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}

		res, body, err := client.HTTPPost(b)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		if c.IsSet("count") {
			count, err := client.ResultsCount(res)
			if err != nil {
				return &ngsiCmdError{funcName, 8, "ResultsCount error", nil}
			}
			fmt.Fprintln(ngsi.StdWriter, count)
			return nil
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 9, "ResultsCount error", err}
		}
		if count == 0 {
			break
		}

		if client.IsSafeString() {
			body, err = ngsilib.JSONSafeStringDecode(body)
			if err != nil {
				return &ngsiCmdError{funcName, 10, err.Error(), err}
			}
		}

		if lines {
			if c.IsSet("values") {
				var values [][]interface{}
				err = ngsilib.JSONUnmarshal(body, &values)
				if err != nil {
					return &ngsiCmdError{funcName, 11, err.Error(), err}
				}
				for _, e := range values {
					b, err := ngsilib.JSONMarshal(&e)
					if err != nil {
						return &ngsiCmdError{funcName, 12, err.Error(), err}
					}
					fmt.Fprintln(ngsi.StdWriter, string(b))
				}
			} else {
				var entities entitiesRespose
				err = ngsilib.JSONUnmarshal(body, &entities)
				if err != nil {
					return &ngsiCmdError{funcName, 13, err.Error(), err}
				}
				for _, e := range entities {
					b, err := ngsilib.JSONMarshal(&e)
					if err != nil {
						return &ngsiCmdError{funcName, 14, err.Error(), err}
					}
					fmt.Fprintln(ngsi.StdWriter, string(b))
				}
			}
		} else if verbose {
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
				if err != nil {
					return &ngsiCmdError{funcName, 15, err.Error(), err}
				}
				buf.bufferWrite(newBuf.Bytes())
			} else {
				buf.bufferWrite(body)
			}
		} else {
			var entities entitiesRespose
			err = ngsilib.JSONUnmarshalDecode(body, &entities, client.IsSafeString())
			if err != nil {
				return &ngsiCmdError{funcName, 16, err.Error(), err}
			}
			for _, e := range entities {
				fmt.Fprintln(ngsi.StdWriter, e["id"])
			}
		}

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	if verbose {
		buf.bufferClose()
	}
	return nil
}
