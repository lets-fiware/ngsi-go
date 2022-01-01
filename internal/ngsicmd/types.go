/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"strconv"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func typesList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	if client.IsNgsiV2() {
		return typesListV2(c, ngsi, client)
	}
	return typesListLd(c, ngsi, client)
}

func typesListV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "typesListV2"

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
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return ngsierr.New(funcName, 3, "ResultsCount error", err)
		}
		if count == 0 {
			break
		}

		var t []string
		err = ngsilib.JSONUnmarshal(body, &t)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		types = append(types, t...)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	if c.IsSetOR([]string{"json", "pretty"}) {
		b, err := ngsilib.JSONMarshal(types)
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 6, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(b))
		}
	} else {
		for _, e := range types {
			fmt.Fprintln(ngsi.StdWriter, e)
		}
	}

	return nil
}

// 5.2.24 EntityTypeList
type entityTypeList struct {
	AtContext interface{} `json:"@context,omitempty"`
	ID        string      `json:"id,omitempty"`
	Type      string      `json:"type,omitempty"`
	TypeList  []string    `json:"typeList,omitempty"`
}

func typesListLd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "typesListLd"

	client.SetPath("/types")

	if c.IsSet("details") {
		v := &url.Values{}
		v.Set("details", "true")
		client.SetQuery(v)
	}

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.IsSetOR([]string{"json", "details", "pretty"}) {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(body))
		}
	} else {
		var list entityTypeList
		err := ngsilib.JSONUnmarshal(body, &list)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		for _, e := range list.TypeList {
			fmt.Fprintln(ngsi.StdWriter, e)
		}
	}

	return nil
}

func typeGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "typesGet"

	t := ""
	if c.IsSet("type") && c.Args().Len() == 0 {
		t = c.String("type")
	} else if !c.IsSet("type") && c.Args().Len() == 1 {
		t = c.Args().Get(0)
	} else {
		return ngsierr.New(funcName, 1, "missing entity type", nil)
	}

	if client.IsNgsiV2() {
		return typeGetV2(c, ngsi, client, t)
	}
	return typeGetLd(c, ngsi, client, t)
}

func typeGetV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client, t string) error {
	const funcName = "typeGetV2"

	client.SetPath("/types/" + t)

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
	} else {
		fmt.Fprintln(ngsi.StdWriter, string(body))
	}

	return nil
}

func typeGetLd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client, t string) error {
	const funcName = "typeGetLd"

	client.SetPath("/types/" + url.QueryEscape(t))

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
	} else {
		fmt.Fprintln(ngsi.StdWriter, string(body))
	}

	return nil
}

func typesCount(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "typesCount"

	if client.IsNgsiLd() {
		return ngsierr.New(funcName, 1, "Only available on NGSIv2", nil)
	}

	client.SetPath("/types")

	v := url.Values{}
	v.Set("options", "values,count")
	client.SetQuery(&v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	count, err := client.ResultsCount(res)
	if err != nil {
		return ngsierr.New(funcName, 4, "ResultsCount error", nil)
	}

	fmt.Fprintln(ngsi.StdWriter, strconv.Itoa(count))

	return nil
}
