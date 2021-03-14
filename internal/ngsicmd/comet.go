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
	"net/url"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

// https://github.com/telefonicaid/fiware-sth-comet/blob/master/apiary.apib

/*
func cometAttrRead(c *cli.Context) error {
	const funcName = "cometAttrRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"comet"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return cometAttrReadMain(c, ngsi, client)
}
*/

func cometAttrReadMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "cometAttrReadMain"

	for _, v := range []string{"type", "id", "attr"} {
		if !c.IsSet(v) {
			return &ngsiCmdError{funcName, 1, "missing " + v, nil}
		}
	}

	param := []string{}
	for _, p := range []string{"hLimit", "hOffset", "lastN", "aggrMethod", "aggrPeriod"} {
		if c.IsSet(p) {
			param = append(param, p)
		}
	}

	if len(param) == 0 {
		return &ngsiCmdError{funcName, 2, "no way to consume data", nil}
	}

	v := url.Values{}
	param = append(param, "type")
	for _, p := range param {
		v.Set(p, c.String(p))
	}
	for _, p := range []string{"fromDate", "toDate"} {
		if c.IsSet(p) {
			dt, err := getDateTime(c.String(p))
			if err != nil {
				return &ngsiCmdError{funcName, 3, err.Error(), err}
			}
			if p == "fromDate" {
				v.Set("dateFrom", dt)
			} else {
				v.Set("dateTo", dt)
			}
		}
	}
	client.SetQuery(&v)

	path := fmt.Sprintf("/STH/v2/entities/%s/attrs/%s", c.String("id"), c.String("attr"))
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

/*
func cometEntitiesDelete(c *cli.Context) error {
	const funcName = "cometEntitiesDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"comet"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return cometEntitiesDeleteMain(c, ngsi, client)
}
*/

func cometEntitiesDeleteMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "cometEntitiesDeleteMain"

	if !c.IsSet("run") {
		fmt.Fprintf(ngsi.StdWriter, "all the data associated to certain service and service path wiil be removed. run delete with -run option\n")
		return nil
	}

	client.SetPath("/STH/v1/contextEntities")

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

/*
func cometEntityDelete(c *cli.Context) error {
	const funcName = "cometEntityDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"comet"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return cometEntityDeleteMain(c, ngsi, client)
}
*/

func cometEntityDeleteMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "cometEntityDeleteMain"

	path := ""
	if c.IsSet("type") && c.IsSet("id") {
		path = fmt.Sprintf("/STH/v1/contextEntities/type/%s/id/%s", c.String("type"), c.String("id"))
	} else {
		if c.IsSet("type") && !c.IsSet("id") {
			return &ngsiCmdError{funcName, 1, "missing id", nil}
		}
		return &ngsiCmdError{funcName, 2, "missing type", nil}
	}

	if !c.IsSet("run") {
		fmt.Fprintf(ngsi.StdWriter, "all the data associated to entity <%s>, service and service path wiil be removed. run delete with -run option\n", c.String("id"))
		return nil
	}

	client.SetPath(path)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func cometAttrDelete(c *cli.Context) error {
	const funcName = "cometAttrDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"comet"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return cometAttrDeleteMain(c, ngsi, client)
}

func cometAttrDeleteMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "cometAttrDeleteMain"

	path := ""
	if c.IsSet("type") && c.IsSet("id") && c.IsSet("attr") {
		path = fmt.Sprintf("/STH/v1/contextEntities/type/%s/id/%s/attributes/%s", c.String("type"), c.String("id"), c.String("attr"))
	} else {
		if !c.IsSet("type") && c.IsSet("id") {
			return &ngsiCmdError{funcName, 1, "missing type", nil}
		} else if c.IsSet("type") && !c.IsSet("id") {
			return &ngsiCmdError{funcName, 2, "missing id", nil}
		}
		return &ngsiCmdError{funcName, 3, "missing attr", nil}
	}

	if !c.IsSet("run") {
		fmt.Fprintf(ngsi.StdWriter, "all the data associated to attribute <%s> of entity <%s>, service and service path wiil be removed. run delete with -run option\n", c.String("attr"), c.String("id"))
		return nil
	}

	client.SetPath(path)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}
