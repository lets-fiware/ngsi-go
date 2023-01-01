/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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
	"net/url"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func entityCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entityCreate"

	if client.IsNgsiLd() {
		for _, name := range []string{"keyValues", "upsert"} {
			if c.IsSet(name) {
				return ngsierr.New(funcName, 1, fmt.Sprintf("--%s only available on NGSIv2", name), nil)
			}
		}
	}

	client.SetPath("/entities")

	var opts = []string{"keyValues", "upsert"}
	v := ngsicli.ParseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetContentType()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = ngsi.InsertAtContext(b, c.String("context"))
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 5, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func entityRead(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entityRead"

	id := c.String("id")
	client.SetPath("/entities/" + id)

	args := []string{"type", "attrs"}
	var opts = []string{"keyValues", "values", "unique", "sysAttrs"}
	v := ngsicli.ParseOptions(c, args, opts)
	client.SetQuery(v)

	if c.Bool("acceptJson") {
		client.SetAcceptJSON()
	} else if c.Bool("acceptGeoJson") {
		client.SetAcceptGeoJSON()
	}

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error: %s %s", res.Status, string(body)), nil)
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))
	return nil
}

func entityUpsert(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entityUpsert"

	client.SetPath("/entities")

	v := url.Values{}
	options := "upsert"
	if c.IsSet("keyValues") {
		options = options + ",keyValues"
	}
	v.Set("options", options)
	client.SetQuery(&v)

	client.SetContentType()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func entityDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entityDelete"

	id := c.String("id")
	client.SetPath("/entities/" + id)

	args := []string{"type"}
	v := ngsicli.ParseOptions(c, args, nil)
	client.SetQuery(v)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}
