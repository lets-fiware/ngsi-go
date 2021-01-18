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
	"strconv"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"

	"github.com/urfave/cli/v2"
)

func attrRead(c *cli.Context) error {
	const funcName = "attrRead"
	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	attrName := c.String("attrName")
	path := fmt.Sprintf("/entities/%s/attrs/%s", id, attrName)
	if client.IsNgsiV2() {
		path = path + "/value"
	}
	client.SetPath(path)

	args := []string{"type"}
	v := parseOptions(c, args, nil)
	client.SetQuery(v)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func attrUpdate(c *cli.Context) error {
	const funcName = "attrUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	attrName := c.String("attrName")
	path := fmt.Sprintf("/entities/%s/attrs/%s", id, attrName)
	if client.IsNgsiV2() {
		path = path + "/value"
	}
	client.SetPath(path)

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	if ngsilib.IsJSON(b) {
		if client.IsNgsiLd() && c.IsSet("context") {
			b, err = insertAtContext(ngsi, b, c.String("context"))
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
		}
		if client.IsSafeString() {
			b, err = ngsilib.JSONSafeStringEncode(b)
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
		}
		client.SetContentType()
	} else {
		s := string(b)
		if _, err := strconv.ParseFloat(s, 64); err != nil {
			if !ngsilib.Contains([]string{"null", "true", "false"}, s) {
				if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
					if len(s) < 2 {
						return &ngsiCmdError{funcName, 6, "data length error", nil}
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
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func attrDelete(c *cli.Context) error {
	const funcName = "attrDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	id := c.String("id")
	attrName := c.String("attrName")
	path := fmt.Sprintf("/entities/%s/attrs/%s", id, attrName)
	client.SetPath(path)

	args := []string{"type"}
	v := parseOptions(c, args, nil)
	client.SetQuery(v)

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}
