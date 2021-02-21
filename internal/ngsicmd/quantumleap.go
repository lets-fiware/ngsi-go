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

// https://github.com/smartsdk/ngsi-timeseries-api/blob/master/specification/quantumleap.yml

/*
func qlEntitiesRead(c *cli.Context) error {
	const funcName = "qlEntitiesRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return qlEntitiesReadMain(c, ngsi, client)
}
*/

func qlEntitiesReadMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "qlEntitiesReadMain"

	path := "/v2/entities"
	client.SetPath(path)

	v := parseOptions(c, []string{"type", "hLimit", "hOffset"}, nil)
	for _, p := range []string{"fromDate", "toDate"} {
		if c.IsSet(p) {
			dt, err := getDateTime(c.String(p))
			if err != nil {
				return &ngsiCmdError{funcName, 1, err.Error(), err}
			}
			v.Set(p, dt)
		}
	}
	client.SetQuery(v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

/*
func qlAttrRead(c *cli.Context) error {
	const funcName = "qlAttrRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return qlAttrReadMain(c, ngsi, client)
}
*/

func qlAttrReadMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "qlAttrReadMain"

	if !c.IsSet("attrName") {
		return &ngsiCmdError{funcName, 1, "missing attrName", nil}
	}

	if c.IsSet("sameType") && c.IsSet("nTypes") {
		return &ngsiCmdError{funcName, 2, "sameType and nTypes are incompatible", nil}
	}

	if (c.IsSet("georel") != c.IsSet("geometry")) != c.IsSet("coords") {
		return &ngsiCmdError{funcName, 3, "georel, geometry and coords are needed", nil}
	}

	path := ""
	param := []string{"aggrMethod", "aggrPeriod", "lastN", "hLimit", "hOffset", "lastN", "georel", "geometry", "coords"}

	if c.Bool("nTypes") { // History of an attribute of N entities of N types.
		path = fmt.Sprintf("/v2/attrs/%s", c.String("attrName"))
		param = append([]string{"id", "type", "aggrScope"}, param...)

	} else if c.Bool("sameType") { // History of an attribute of N entities of the same type.
		if !c.IsSet("type") {
			return &ngsiCmdError{funcName, 4, "missing type", nil}
		}
		path = fmt.Sprintf("/v2/types/%s/attrs/%s", c.String("type"), c.String("attrName"))
		param = append([]string{"id", "aggrScope"}, param...)

	} else { // History of an attribute of a given entity instance.
		if !c.IsSet("id") {
			return &ngsiCmdError{funcName, 5, "missing id", nil}
		}
		path = fmt.Sprintf("/v2/entities/%s/attrs/%s", c.String("id"), c.String("attrName"))
		param = append([]string{"type"}, param...)
	}

	v := parseOptions(c, param, nil)
	for _, p := range []string{"fromDate", "toDate"} {
		if c.IsSet(p) {
			dt, err := getDateTime(c.String(p))
			if err != nil {
				return &ngsiCmdError{funcName, 6, err.Error(), err}
			}
			v.Set(p, dt)
		}
	}
	if c.IsSet("hLimit") {
		v.Set("limit", v.Get("hLimit"))
		v.Del("hLimit")
	}
	if c.IsSet("hOffset") {
		v.Set("offset", v.Get("hOffset"))
		v.Del("hOffset")
	}
	client.SetQuery(v)

	if c.Bool("value") {
		path += "/value"
	}
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 9, err.Error(), err}
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 10, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

/*
func qlAttrsRead(c *cli.Context) error {
	const funcName = "qlAttrsRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return qlAttrsReadMain(c, ngsi, client)
}
*/

func qlAttrsReadMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "qlAttrsReadMain"

	if (c.IsSet("georel") != c.IsSet("geometry")) != c.IsSet("coords") {
		return &ngsiCmdError{funcName, 1, "georel, geometry and coords are needed", nil}
	}

	path := ""
	param := []string{"attrs", "aggrMethod", "aggrPeriod", "lastN", "hLimit", "hOffset", "lastN", "georel", "geometry", "coords"}

	if c.Bool("nTypes") {
		path = "/v2/attrs"
		param = append([]string{"type", "id", "aggrScope"}, param...)

	} else if c.Bool("sameType") { // History of an attribute of N entities of the same type.
		if !c.IsSet("type") {
			return &ngsiCmdError{funcName, 2, "missing type", nil}
		}
		path = fmt.Sprintf("/v2/types/%s", c.String("type"))
		param = append([]string{"id", "aggrScope"}, param...)

	} else { // History of an attribute of a given entity instance.
		if !c.IsSet("id") {
			return &ngsiCmdError{funcName, 3, "missing id", nil}
		}
		path = fmt.Sprintf("/v2/entities/%s", c.String("id"))
		param = append([]string{"type"}, param...)
	}
	for _, p := range param {
		if c.IsSet(p) {
			param = append(param, p)
		}
	}

	v := parseOptions(c, param, []string{"keyValues"})
	for _, p := range []string{"fromDate", "toDate"} {
		if c.IsSet(p) {
			dt, err := getDateTime(c.String(p))
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
			v.Set(p, dt)
		}
	}
	if c.IsSet("hLimit") {
		v.Set("limit", v.Get("hLimit"))
		v.Del("hLimit")
	}
	if c.IsSet("hOffset") {
		v.Set("offset", v.Get("hOffset"))
		v.Del("hOffset")
	}
	client.SetQuery(v)

	if c.Bool("value") {
		path += "/value"
	}
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

/*
func qlEntityDelete(c *cli.Context) error {
	const funcName = "qlEntityDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return qlEntityDeleteMain(c, ngsi, client)
}
*/

func qlEntityDeleteMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "qlEntityDeleteMain"

	if !c.IsSet("id") {
		return &ngsiCmdError{funcName, 1, "missing id", nil}
	}

	v := parseOptions(c, []string{"type"}, nil)
	for _, p := range []string{"fromDate", "toDate"} {
		if c.IsSet(p) {
			dt, err := getDateTime(c.String(p))
			if err != nil {
				return &ngsiCmdError{funcName, 2, err.Error(), err}
			}
			v.Set(p, dt)
		}
	}
	client.SetQuery(v)

	client.SetPath(fmt.Sprintf("/v2/entities/%s", c.String("id")))

	if !c.IsSet("run") {
		fmt.Fprintf(ngsi.StdWriter, "historical data of the entity <%s> will be deleted. run delete with -run option\n", c.String("id"))
		return nil
	}

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

// Delete historical data of all entities of a certain type.

/*
func qlEntitiesDelete(c *cli.Context) error {
	const funcName = "qlEntitiesDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return qlEntitiesDeleteMain(c, ngsi, client)
}
*/

func qlEntitiesDeleteMain(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "qlEntitiesDeleteMain"

	if !c.IsSet("type") {
		return &ngsiCmdError{funcName, 1, "missing type", nil}
	}

	v := url.Values{}
	for _, p := range []string{"fromDate", "toDate"} {
		if c.IsSet(p) {
			dt, err := getDateTime(c.String(p))
			if err != nil {
				return &ngsiCmdError{funcName, 2, err.Error(), err}
			}
			v.Set(p, dt)
		}
	}
	if c.Bool("dropTable") {
		v.Set("dropTable", "true")
	}
	client.SetQuery(&v)

	client.SetPath(fmt.Sprintf("/v2/types/%s", c.String("type")))

	if !c.IsSet("run") {
		fmt.Fprintf(ngsi.StdWriter, "historical data of all entities of the type <%s> will be deleted. run delete with -run option\n", c.String("type"))
		return nil
	}

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}
