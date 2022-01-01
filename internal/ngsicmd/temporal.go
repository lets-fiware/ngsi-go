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

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func troeList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "troeList"

	args := []string{"id", "type", "idPattern", "attrs", "query", "csf", "georel", "geometry", "coords", "geoProperty", "timeProperty", "lastN"}
	opts := []string{"sysattrs", "temporalValues"}
	v := ngsicli.ParseOptions(c, args, opts)

	err := buildTemporalQuery(c, v)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	client.SetQuery(v)

	client.SetPath("/temporal/entities/")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func troeCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "troeCreate"

	client.SetPath("/temporal/entities/")

	client.SetContentType()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = ngsi.InsertAtContext(b, c.String("context"))
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func troeRead(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "troeRead"

	client.SetPath("/temporal/entities/" + c.String("id"))

	args := []string{"attrs"}
	var opts = []string{"sysAttrs", "temporalValues"}
	v := ngsicli.ParseOptions(c, args, opts)

	err := buildTemporalQuery(c, v)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	client.SetQuery(v)

	if c.Bool("acceptJson") {
		client.SetAcceptJSON()
	}

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 3, fmt.Sprintf("error: %s %s", res.Status, string(body)), nil)
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func troeDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "troeDelete"

	client.SetPath("/temporal/entities/" + c.String("id"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func troeAttrsAppend(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "troeAttrsAppend"

	client.SetPath("/temporal/entities/" + c.String("id") + "/attrs/")

	client.SetContentType()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = ngsi.InsertAtContext(b, c.String("context"))
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func troeAttrDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "troeAttrDelete"

	if c.IsSet("instanceId") {
		if c.IsSetOR([]string{"deleteAll", "datasetId"}) {
			return ngsierr.New(funcName, 1, "cannot specify --deleteALl and/or --datasetId with --instanceId", nil)
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
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func troeAttrUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "troeAttrUpdate"

	client.SetPath(fmt.Sprintf("/temporal/entities/%s/attrs/%s/%s", c.String("id"), c.String("attr"), c.String("instanceId")))

	client.SetContentType()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if client.IsSafeString() {
		b, err = ngsilib.JSONSafeStringEncode(b)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	if client.IsNgsiLd() && c.IsSet("context") {
		b, err = ngsi.InsertAtContext(b, c.String("context"))
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func buildTemporalQuery(c *ngsicli.Context, v *url.Values) error {
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
			dt, err := ngsilib.GetDateTime(c.String(p))
			if err != nil {
				return ngsierr.New(funcName, 1, err.Error(), err)
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
