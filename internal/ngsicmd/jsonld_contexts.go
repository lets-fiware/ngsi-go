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
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type jsonldContexts []struct {
	URL       string            `json:"url"`
	ID        string            `json:"id"`
	Type      string            `json:"type"`
	Origin    string            `json:"origin"`
	CreatedAt time.Time         `json:"createdAt"`
	LastUse   time.Time         `json:"lastUse"`
	Lookups   int               `json:"lookups"`
	HashTable map[string]string `json:"hash-table"`
	URLs      []string          `json:"URLs"`
}

func jsonldContextsList(c *cli.Context) error {
	const funcName = "jsonldContextsList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "only available on NGSI-LD", nil}
	}

	path := "/jsonldContexts"

	if c.IsSet("details") || !c.IsSet("json") {
		v := &url.Values{}
		v.Set("details", "true")
		client.SetQuery(v)
	}

	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if isSetOR(c, []string{"json", "details"}) {
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
	} else {
		var contexts jsonldContexts
		err := ngsilib.JSONUnmarshal(body, &contexts)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}

		for _, context := range contexts {
			fmt.Fprintf(ngsi.StdWriter, "%s %s\n", context.ID, context.URL)
		}
	}

	return nil
}

func jsonldContextGet(c *cli.Context) error {
	const funcName = "jsonldContextGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "only available on NGSI-LD", nil}
	}

	id := ""
	if c.IsSet("id") && c.Args().Len() == 0 {
		id = c.String("id")
	} else if !c.IsSet("id") && c.Args().Len() == 1 {
		id = c.Args().Get(0)
	} else {
		return &ngsiCmdError{funcName, 4, "missing jsonldContext id", nil}
	}

	path := "/jsonldContexts/" + id
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
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

func jsonldContextCreate(c *cli.Context) error {
	const funcName = "jsonldContextCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "only available on NGSI-LD", nil}
	}

	if !c.IsSet("data") {
		return &ngsiCmdError{funcName, 4, "missing jsonldContext data", nil}
	}

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}

	path := "/jsonldContexts"
	client.SetPath(path)
	client.SetContentJSON()

	res, Body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(Body)), nil}
	}

	location := res.Header.Get("Location")
	p := "/ngsi-ld/v1/jsonldContexts/"
	pos := strings.Index(location, p)
	if pos >= 0 {
		location = location[pos+len(p):]
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is created, FIWARE-Service: %s", res.Header.Get("Location"), c.String("service")))

	fmt.Fprintln(ngsi.StdWriter, location)

	return nil
}

func jsonldContextDelete(c *cli.Context) error {
	const funcName = "jsonldContextDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.NgsiType != "ld" {
		return &ngsiCmdError{funcName, 3, "only available on NGSI-LD", nil}
	}

	id := ""
	if c.IsSet("id") && c.Args().Len() == 0 {
		id = c.String("id")
	} else if !c.IsSet("id") && c.Args().Len() == 1 {
		id = c.Args().Get(0)
	} else {
		return &ngsiCmdError{funcName, 4, "missing jsonldContext id", nil}
	}

	path := "/jsonldContexts/" + id
	client.SetPath(path)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}
