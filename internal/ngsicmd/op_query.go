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

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func opQuery(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "opQuery"

	if client.IsNgsiLd() {
		return ngsierr.New(funcName, 0, "Only available on NGSIv2", nil)
	}

	page := 0
	count := 0
	limit := 100

	verbose := c.IsSet("verbose")
	if c.Bool("pretty") {
		verbose = true
	}
	lines := c.Bool("lines")

	buf := ngsilib.NewJsonBuffer()
	if verbose {
		buf.BufferOpen(ngsi.StdWriter, false, false)
	}

	for {
		client.SetPath("/op/query")

		var args = []string{"orderBy"}
		var opts = []string{"count", "keyValues", "values", "unique"}
		v := ngsicli.ParseOptions(c, args, opts)
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

		b, err := ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}

		res, body, err := client.HTTPPost(b)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
		}

		if c.IsSet("count") {
			count, err := client.ResultsCount(res)
			if err != nil {
				return ngsierr.New(funcName, 4, "ResultsCount error", nil)
			}
			fmt.Fprintln(ngsi.StdWriter, count)
			return nil
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return ngsierr.New(funcName, 5, "ResultsCount error", err)
		}
		if count == 0 {
			break
		}

		if client.IsSafeString() {
			body, err = ngsilib.JSONSafeStringDecode(body)
			if err != nil {
				return ngsierr.New(funcName, 6, err.Error(), err)
			}
		}

		if lines {
			if c.IsSet("values") {
				var values [][]interface{}
				err = ngsilib.JSONUnmarshal(body, &values)
				if err != nil {
					return ngsierr.New(funcName, 7, err.Error(), err)
				}
				for _, e := range values {
					b, err := ngsilib.JSONMarshal(&e)
					if err != nil {
						return ngsierr.New(funcName, 8, err.Error(), err)
					}
					fmt.Fprintln(ngsi.StdWriter, string(b))
				}
			} else {
				var entities ngsilib.EntitiesRespose
				err = ngsilib.JSONUnmarshal(body, &entities)
				if err != nil {
					return ngsierr.New(funcName, 9, err.Error(), err)
				}
				for _, e := range entities {
					b, err := ngsilib.JSONMarshal(&e)
					if err != nil {
						return ngsierr.New(funcName, 10, err.Error(), err)
					}
					fmt.Fprintln(ngsi.StdWriter, string(b))
				}
			}
		} else if verbose {
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
				if err != nil {
					return ngsierr.New(funcName, 11, err.Error(), err)
				}
				buf.BufferWrite(newBuf.Bytes())
			} else {
				buf.BufferWrite(body)
			}
		} else {
			var entities ngsilib.EntitiesRespose
			err = ngsilib.JSONUnmarshalDecode(body, &entities, client.IsSafeString())
			if err != nil {
				return ngsierr.New(funcName, 12, err.Error(), err)
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
		buf.BufferClose()
	}
	return nil
}
