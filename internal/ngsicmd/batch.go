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
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func batch(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client, mode string) error {
	const funcName = "batch"

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
	return ngsierr.New(funcName, 1, "error: "+mode, nil)
}

func batchCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchCreate"

	client.SetPath("/entityOperations/create")

	client.SetContentType()
	client.SetAcceptJSON()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = ngsi.InsertAtContext(b, c.String("context"))
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated {
		return ngsierr.New(funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func batchUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchUpdate"

	client.SetPath("/entityOperations/update")

	var opts = []string{"noOverwrite", "replace"}
	v := ngsicli.ParseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetContentType()
	client.SetAcceptJSON()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = ngsi.InsertAtContext(b, c.String("context"))
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}
func batchUpsert(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchUpsert"

	client.SetPath("/entityOperations/upsert")

	var opts = []string{"replace", "update"}
	v := ngsicli.ParseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetContentType()
	client.SetAcceptJSON()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = ngsi.InsertAtContext(b, c.String("context"))
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}
	if res.StatusCode == http.StatusCreated {
		fmt.Fprintln(ngsi.StdWriter, string(body))
		return nil
	} else if res.StatusCode == http.StatusNoContent {
		return nil
	}

	return ngsierr.New(funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
}

func batchDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "batchDelete"

	client.SetPath("/entityOperations/delete")

	client.SetContentType()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}
