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

func copy(c *cli.Context) error {
	const funcName = "copy"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	source, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	entityType := c.String("type")

	flags, err := parseFlags2(ngsi, c)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	destination, err := ngsi.NewClient(ngsi.Destination, flags, false)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error() + " (destination)", err}
	}

	if source.IsNgsiLd() || destination.IsNgsiLd() {
		return &ngsiCmdError{funcName, 5, "only available on NGSIv2", err}
	}

	if !c.IsSet("run") {
		return &ngsiCmdError{funcName, 6, "run copy with --run option", err}
	}

	page := 0
	count := 0
	limit := 100
	total := 0
	for {
		source.SetPath("/entities")

		v := url.Values{}
		v.Set("type", entityType)
		v.Set("options", "count")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		source.SetQuery(&v)

		res, body, err := source.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 8, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}
		count, err = source.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 9, err.Error(), err}
		}
		if count == 0 {
			break
		}

		var entities entitiesRespose
		err = ngsilib.JSONUnmarshal(body, &entities)
		if err != nil {
			return &ngsiCmdError{funcName, 10, err.Error(), err}
		}

		res, _, err = destination.OpUpdate(&entities, "append", false, false)
		if err != nil {
			return &ngsiCmdError{funcName, 11, err.Error(), err}
		}

		total += len(entities)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	fmt.Fprintln(ngsi.StdWriter, total)

	return nil
}
