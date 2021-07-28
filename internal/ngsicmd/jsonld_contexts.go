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

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
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

func jsonldContextsList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "jsonldContextsList"

	client.SetPath("/jsonldContexts")

	if c.IsSet("details") || !c.IsSet("json") {
		v := &url.Values{}
		v.Set("details", "true")
		client.SetQuery(v)
	}

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	if c.IsSetOR([]string{"json", "details"}) {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			return nil
		}
		fmt.Fprintln(ngsi.StdWriter, string(body))
	} else {
		var contexts jsonldContexts
		err := ngsilib.JSONUnmarshal(body, &contexts)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}

		for _, context := range contexts {
			fmt.Fprintf(ngsi.StdWriter, "%s %s\n", context.ID, context.URL)
		}
	}

	return nil
}

func jsonldContextGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "jsonldContextGet"

	client.SetPath("/jsonldContexts/" + c.String("id"))

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}
	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func jsonldContextCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "jsonldContextCreate"

	client.SetPath("/jsonldContexts")
	client.SetContentJSON()

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	res, Body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated {
		return ngsierr.New(funcName, 3, fmt.Sprintf("%s %s", res.Status, string(Body)), nil)
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

func jsonldContextDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "jsonldContextDelete"

	client.SetPath("/jsonldContexts/" + c.String("id"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}
