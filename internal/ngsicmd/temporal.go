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

func troeList(c *cli.Context) error {
	const funcName = "troeList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "ngsiType error", nil}
	}

	args := []string{"id", "type", "idPattern", "attrs", "query", "csf", "georel", "geometry", "coords", "geoProperty", "timeProperty", "lastN"}
	opts := []string{"sysattrs", "temporalValues"}
	v := parseOptions(c, args, opts)

	err = buildTemporalQuery(c, v)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	client.SetQuery(v)

	client.SetPath("/temporal/entities/")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
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

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func troeCreate(c *cli.Context) error {
	const funcName = "troeCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "ngsiType error", nil}
	}

	client.SetPath("/temporal/entities/")

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func troeRead(c *cli.Context) error {
	const funcName = "troeRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "ngsiType error", nil}
	}

	if !c.IsSet("id") {
		return &ngsiCmdError{funcName, 4, "specify temporal entity id", err}
	}

	client.SetPath("/temporal/entities/" + c.String("id"))

	args := []string{"attrs"}
	var opts = []string{"sysAttrs", "temporalValues"}
	v := parseOptions(c, args, opts)

	err = buildTemporalQuery(c, v)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}

	client.SetQuery(v)

	if c.Bool("acceptJson") {
		client.SetAcceptJSON()
	}

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("error: %s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 9, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func troeDelete(c *cli.Context) error {
	const funcName = "troeDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "ngsiType error", nil}
	}

	if !c.IsSet("id") {
		return &ngsiCmdError{funcName, 4, "specify temporal entity id", err}
	}

	client.SetPath("/temporal/entities/" + c.String("id"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func troeAttrsAppend(c *cli.Context) error {
	const funcName = "troeAttrsAppend"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "ngsiType error", nil}
	}

	if !c.IsSet("id") {
		return &ngsiCmdError{funcName, 4, "specify temporal entity id", err}
	}

	client.SetPath("/temporal/entities/" + c.String("id") + "/attrs/")

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 8, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 9, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func troeAttrDelete(c *cli.Context) error {
	const funcName = "troeAttrDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "ngsiType error", nil}
	}

	params := []struct {
		arg string
		msg string
	}{
		{"id", "specify temporal entity id"},
		{"attr", "specify attribute name"},
	}
	for _, param := range params {
		if !c.IsSet(param.arg) {
			return &ngsiCmdError{funcName, 4, param.msg, err}
		}
	}

	if c.IsSet("instanceId") {
		if isSetOR(c, []string{"deleteAll", "datasetId"}) {
			return &ngsiCmdError{funcName, 5, "cannot specify --deleteALl and/or --datasetId with --instanceId", nil}
		}
		client.SetPath(fmt.Sprintf("/temporal/entities/%s/attrs/%s/%s", c.String("id"), c.String("attr"), c.String("instanceId")))
	} else {
		v := url.Values{}
		if c.Bool("deleteAll") {
			v.Set("deleteAll", "true")
		}
		if c.IsSet("datasetId") {
			v.Set("datasetId", c.String("datasetId"))
		}
		client.SetQuery(&v)
		client.SetPath(fmt.Sprintf("/temporal/entities/%s/attrs/%s", c.String("id"), c.String("attr")))
	}

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func troeAttrUpdate(c *cli.Context) error {
	const funcName = "troeAttrUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "ngsiType error", nil}
	}

	params := []struct {
		arg string
		msg string
	}{
		{"id", "specify temporal entity id"},
		{"attr", "specify attribute name"},
		{"instanceId", "specify instance id"},
	}
	for _, param := range params {
		if !c.IsSet(param.arg) {
			return &ngsiCmdError{funcName, 4, param.msg, err}
		}
	}

	client.SetPath(fmt.Sprintf("/temporal/entities/%s/attrs/%s/%s", c.String("id"), c.String("attr"), c.String("instanceId")))

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = insertAtContext(ngsi, b, c.String("context"))
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
	}

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return &ngsiCmdError{funcName, 8, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 9, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func buildTemporalQuery(c *cli.Context, v *url.Values) error {
	const funcName = "buildTemporalQuery"

	timeAt := "timeAt"
	endTimeAt := "endTimeAt"

	if c.Bool("etsi10") {
		timeAt = "time"
		endTimeAt = "endTime"
	}

	fromDate := ""
	toDate := ""

	for _, p := range []string{"fromDate", "toDate"} {
		if c.IsSet(p) {
			dt, err := getDateTime(c.String(p))
			if err != nil {
				return &ngsiCmdError{funcName, 1, err.Error(), err}
			}
			if p == "fromDate" {
				fromDate = dt
			} else {
				toDate = dt
			}
		}
	}

	if fromDate != "" && toDate != "" {
		v.Set("timerel", "between")
		v.Set(timeAt, fromDate)
		v.Set(endTimeAt, toDate)
	} else if fromDate != "" {
		v.Set("timerel", "after")
		v.Set(timeAt, fromDate)
	} else if toDate != "" {
		v.Set("timerel", "before")
		v.Set(timeAt, toDate)
	}

	return nil
}
