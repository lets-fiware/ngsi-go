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
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func registrationsList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	if client.IsNgsiV2() {
		return registrationsListV2(c, ngsi, client)
	}
	return registrationsListLd(c, ngsi, client)
}

func registrationsGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	if client.IsNgsiV2() {
		return registrationsGetV2(c, ngsi, client)
	}
	return registrationsGetLd(c, ngsi, client)
}

func registrationsCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	if client.IsNgsiV2() {
		return registrationsCreateV2(c, ngsi, client)
	}
	return registrationsCreateLd(c, ngsi, client)
}

func registrationsDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	if client.IsNgsiV2() {
		return registrationsDeleteV2(c, ngsi, client)
	}
	return registrationsDeleteLd(c, ngsi, client)
}

func registrationsTemplate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsTemplate"

	if !c.IsSet("ngsiType") {
		return ngsierr.New(funcName, 1, "missing ngsiType", nil)
	}

	t := strings.ToLower(c.String("ngsiType"))

	if ngsilib.IsNgsiV2(t) {
		return registrationsTemplateV2(c, ngsi)
	}

	if ngsilib.IsNgsiLd(t) {
		return registrationsTemplateLd(c, ngsi)
	}

	return ngsierr.New(funcName, 2, "ngsiType error: "+t, nil)
}

func registrationsCount(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsCount"

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
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	count, err := client.ResultsCount(res)
	if err != nil {
		return ngsierr.New(funcName, 3, "ResultsCount error", nil)
	}

	fmt.Fprintln(ngsi.StdWriter, count)

	return nil
}
