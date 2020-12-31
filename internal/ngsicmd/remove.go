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
	"net/url"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func remove(c *cli.Context) error {
	const funcName = "remove"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		if c.IsSet("link") {
			return &ngsiCmdError{funcName, 3, "can't specify --link option on NGSIv2", nil}
		}
		return removeV2(c, ngsi, client)
	}
	return removeLD(c, ngsi, client)
}

func removeV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "removeV2"

	entityType := c.String("type")

	limit := 100
	total := 0
	for {
		client.SetPath("/entities")

		v := url.Values{}
		v.Set("type", entityType)
		v.Set("options", "count")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("attrs", "id")
		client.SetQuery(&v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		count, err := client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 3, "ResultsCount error", nil}
		}

		if !c.IsSet("run") {
			fmt.Fprintf(ngsi.StdWriter, fmt.Sprintf("%d entities will be removed. run remove with --run option\n", count))
			return nil
		}

		if count == 0 {
			break
		}
		if count >= limit {
			total += limit
		} else {
			total += count
		}

		var entities entitiesRespose
		err = ngsilib.JSONUnmarshalDecode(body, &entities, false)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}

		_, _, err = client.OpUpdate(&entities, "delete", false, false)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		client.RemoveHeader("Content-Type")
	}

	fmt.Fprintf(ngsi.StdWriter, "%d\n", total)

	return nil
}

func removeLD(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "removeLD"

	entityType := c.String("type")

	limit := 100
	total := 0
	for {
		// get count
		client.SetPath("/entities")

		v := url.Values{}
		v.Set("type", entityType)
		v.Set("count", "true")
		v.Set("limit", fmt.Sprintf("%d", limit))
		client.SetQuery(&v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		count, err := client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 3, "ResultsCount error", nil}
		}

		if !c.IsSet("run") {
			fmt.Fprintf(ngsi.StdWriter, fmt.Sprintf("%d entities will be removed. run remove with --run option\n", count))
			return nil
		}

		if count == 0 {
			break
		}
		if count >= limit {
			total += limit
		} else {
			total += count
		}

		var entities entitiesRespose
		err = ngsilib.JSONUnmarshalDecode(body, &entities, false)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}

		data := []string{}
		for _, e := range entities {
			data = append(data, e["id"].(string))
		}
		b, err := ngsilib.JSONMarshal(data)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}

		client.SetPath("/entityOperations/delete")
		v = url.Values{}
		client.SetQuery(&v)
		client.SetContentType()

		res, body, err = client.HTTPPost(b)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		client.RemoveHeader("Content-Type")
	}

	fmt.Fprintf(ngsi.StdWriter, "%d\n", total)

	return nil
}
