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

func registrationsList(c *cli.Context) error {
	const funcName = "registratinsList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return registrationsListV2(c, ngsi, client)
	}
	return registrationsListLd(c, ngsi, client)
}

func registrationsGet(c *cli.Context) error {
	const funcName = "registratinsList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return registrationsGetV2(c, ngsi, client)
	}
	return registrationsGetLd(c, ngsi, client)
}

func registrationsCreate(c *cli.Context) error {
	const funcName = "registrationsCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return registrationsCreateV2(c, ngsi, client)
	}
	return registrationsCreateLd(c, ngsi, client)
}

func registrationsDelete(c *cli.Context) error {
	const funcName = "registrationsDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiV2() {
		return registrationsDeleteV2(c, ngsi, client)
	}
	return registrationsDeleteLd(c, ngsi, client)
}

func registrationsTemplate(c *cli.Context) error {
	const funcName = "registrationsTemplate"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	t := strings.ToLower(c.String("ngsiType"))
	if t == "v2" || t == "ngsiv2" || t == "ngsi-v2" {
		return registrationsTemplateV2(c, ngsi)
	}

	return registrationsTemplateLd(c, ngsi)
}

func registrationsCount(c *cli.Context) error {
	const funcName = "registratinsList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	v := url.Values{}

	if client.IsNgsiLd() {
		client.SetPath("/csourceRegistrations")
		v.Set("limit", "0")
		v.Set("count", "true")
	} else {
		client.SetPath("/registrations")
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
