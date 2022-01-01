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
	"strconv"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func attrRead(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "attrRead"

	id := c.String("id")
	attr := c.String("attr")
	path := fmt.Sprintf("/entities/%s/attrs/%s", id, attr)
	if client.IsNgsiV2() {
		path = path + "/value"
	}
	client.SetPath(path)

	args := []string{"type"}
	v := ngsicli.ParseOptions(c, args, nil)
	client.SetQuery(v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
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

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func attrUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "attrUpdate"

	id := c.String("id")
	attr := c.String("attr")
	path := fmt.Sprintf("/entities/%s/attrs/%s", id, attr)
	if client.IsNgsiV2() {
		path = path + "/value"
	}
	client.SetPath(path)

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if ngsilib.IsJSON(b) {
		if client.IsNgsiLd() && c.IsSet("context") {
			b, err = ngsi.InsertAtContext(b, c.String("context"))
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error(), err)
			}
		}
		if client.IsSafeString() {
			b, err = ngsilib.JSONSafeStringEncode(b)
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
		}
		client.SetContentType()
	} else {
		s := string(b)
		if _, err := strconv.ParseFloat(s, 64); err != nil {
			if !ngsilib.Contains([]string{"null", "true", "false"}, s) {
				if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
					if len(s) < 2 {
						return ngsierr.New(funcName, 4, "data length error", nil)
					}
					b = []byte(s[1 : len(s)-1])
				}
				if client.IsSafeString() {
					b, _ = ngsilib.JSONSafeStringEncode(b)
				}
				b = []byte(`"` + string(b) + `"`)
			}
		}
		client.SetHeader("Content-Type", "text/plain")
	}

	var res *http.Response
	var body []byte

	if client.IsNgsiLd() {
		res, body, err = client.HTTPPatch(b)
	} else {
		res, body, err = client.HTTPPut(b)
	}
	if err != nil {
		return ngsierr.New(funcName, 5, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func attrDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "attrDelete"

	id := c.String("id")
	attr := c.String("attr")
	path := fmt.Sprintf("/entities/%s/attrs/%s", id, attr)
	client.SetPath(path)

	args := []string{"type"}
	v := ngsicli.ParseOptions(c, args, nil)
	client.SetQuery(v)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}
