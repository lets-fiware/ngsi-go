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

func remove(c *cli.Context) error {
	const funcName = "remove"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	entityType := c.String("type")

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "only available on NGSIv2", err}
	}

	if !c.IsSet("run") {
		return &ngsiCmdError{funcName, 4, "run remove with --run option", err}
	}

	limit := 100
	total := 0
	for {
		// get count
		client.SetPath("/entities")

		v := url.Values{}
		v.Set("type", entityType)
		v.Set("options", "count")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("attrs", "id")
		client.SetQuery(&v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		count, err := strconv.Atoi(res.Header.Get("fiware-total-count"))
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
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
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}

		_, _, err = client.OpUpdate(&entities, "delete", false, false)
		if err != nil {
			return &ngsiCmdError{funcName, 9, err.Error(), err}
		}
	}

	fmt.Fprintf(ngsi.StdWriter, "%d", total)

	return nil
}
