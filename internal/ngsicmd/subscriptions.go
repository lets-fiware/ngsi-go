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
	"strings"

	"github.com/urfave/cli/v2"
)

func subscriptionsList(c *cli.Context) error {
	const funcName = "subscriptionsList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return subscriptionsListV2(c, ngsi, client)
	}
	return subscriptionsListLd(c, ngsi, client)
}

func subscriptionGet(c *cli.Context) error {
	const funcName = "subscriptionGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return subscriptionGetV2(c, ngsi, client)
	}
	return subscriptionGetLd(c, ngsi, client)
}

func subscriptionsCreate(c *cli.Context) error {
	const funcName = "subscriptionsCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return subscriptionsCreateV2(c, ngsi, client)
	}
	return subscriptionsCreateLd(c, ngsi, client)
}

func subscriptionsUpdate(c *cli.Context) error {
	const funcName = "subscriptionsUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	if client.IsNgsiV2() {
		return subscriptionsUpdateV2(c, ngsi, client)
	}
	return subscriptionsUpdateLd(c, ngsi, client)
}

func subscriptionsDelete(c *cli.Context) error {
	const funcName = "subscriptionsDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return subscriptionsDeleteV2(c, ngsi, client)
	}
	return subscriptionsDeleteLd(c, ngsi, client)
}

func subscriptionsTemplate(c *cli.Context) error {
	const funcName = "subscriptionsTemplate"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if !c.IsSet("ngsiType") {
		return &ngsiCmdError{funcName, 2, "Required ngsiType not found", err}
	}

	t := strings.ToLower(c.String("ngsiType"))
	if t == "v2" || t == "ngsiv2" || t == "ngsi-v2" {
		return subscriptionsTemplateV2(c, ngsi)
	}

	if t == "ld" || t == "ngsi-ld" {
		return subscriptionsTemplateLd(c, ngsi)
	}

	return &ngsiCmdError{funcName, 3, "ngsiType error " + t, err}
}

func subscriptionsCount(c *cli.Context) error {
	const funcName = "subscriptionsCount"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	client.SetPath("/subscriptions")
	v := url.Values{}

	if client.IsNgsiLd() {
		v.Set("limit", "0")
		v.Set("count", "true")
	} else {
		v.Set("limit", "1")
		v.Set("options", "count")
	}
	client.SetQuery(&v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	count, err := client.ResultsCount(res)
	if err != nil {
		return &ngsiCmdError{funcName, 5, "ResultsCount error", nil}
	}

	fmt.Fprintln(ngsi.StdWriter, count)

	return nil
}
