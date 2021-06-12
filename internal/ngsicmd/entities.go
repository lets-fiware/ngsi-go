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

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func entitiesList(c *cli.Context) error {
	const funcName = "entitiesList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		if isSetOR(c, []string{"typePattern", "mq", "metadata", "value", "uniq"}) {
			return &ngsiCmdError{funcName, 3, "cannot specfiy typePattern, mq, metadata, value or uniq", nil}
		}
		return entitiesListLD(c, ngsi, client)
	}
	if isSetOR(c, []string{"link", "acceptJson", "acceptGeoJson"}) {
		return &ngsiCmdError{funcName, 4, "cannot specfiy link acceptJson or acceptGeoJson", nil}
	}
	return entitiesListV2(c, ngsi, client)
}

func entitiesListV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entitiesList"

	attrs := "__NONE"

	if c.IsSet("attrs") {
		attrs = c.String("attrs")
	}

	page := 0
	count := 0
	limit := 100

	verbose := c.IsSet("verbose")
	if c.Bool("pretty") || c.Bool("keyValues") || isSetOR(c, []string{"attrs", "metadata", "orderBy"}) {
		verbose = true
	}
	values := c.IsSet("values")
	if values {
		verbose = true
	}
	lines := c.Bool("lines")

	buf := jsonBuffer{}
	if verbose {
		buf.bufferOpen(ngsi.StdWriter, false, false)
		attrs = ""
	}

	for {
		client.SetPath("/entities")

		args := []string{"id", "type", "idPattern", "typePattern", "query", "mq", "georel", "geometry", "coords", "attrs", "metadata", "orderBy"}
		opts := []string{"keyValues", "values", "unique", "skipForwarding"}
		v := parseOptions(c, args, opts)

		if attrs != "" {
			v.Set("attrs", attrs)
		}
		if c.IsSet("count") {
			v.Set("limit", "1")
			v.Set("options", "count")
		} else {
			options := v.Get("options")
			if options == "" {
				options = "count"
			} else {
				options = options + ",count"
			}
			v.Set("options", options)
			v.Set("limit", fmt.Sprintf("%d", limit))
			v.Set("offset", fmt.Sprintf("%d", page*limit))
		}
		client.SetQuery(v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		if c.IsSet("count") {
			count, err := client.ResultsCount(res)
			if err != nil {
				return &ngsiCmdError{funcName, 3, "ResultsCount error", nil}
			}
			fmt.Fprintln(ngsi.StdWriter, count)
			break
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 4, "ResultsCount error", err}
		}
		if count == 0 {
			break
		}

		if client.IsSafeString() {
			body, err = ngsilib.JSONSafeStringDecode(body)
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
		}

		err = entitiesPrint(ngsi, body, &buf, c.Bool("pretty"), lines, values, verbose, false)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}
	if verbose {
		buf.bufferClose()
	}
	return nil
}

func entitiesListLD(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entitiesListLD"

	idPattern := ""
	if !isSetOR(c, []string{"type", "idPattern", "query", "attrs", "georel"}) {
		idPattern = ".*"
	}

	page := 0
	count := 0
	limit := 100

	verbose := c.IsSet("verbose")
	if isSetOR(c, []string{"pretty", "keyValues", "acceptGeoJson", "attrs", "orderBy"}) {
		verbose = true
	}
	lines := c.Bool("lines")

	buf := jsonBuffer{}
	if verbose {
		buf.bufferOpen(ngsi.StdWriter, c.Bool("acceptGeoJson"), c.Bool("pretty"))
	}

	for {
		client.SetPath("/entities")

		args := []string{"id", "type", "idPattern", "query", "georel", "geometry", "coords", "attrs", "orderBy"}
		opts := []string{"keyValues"}
		v := parseOptions(c, args, opts)

		if idPattern != "" {
			v.Set("idPattern", idPattern)
		}
		if c.IsSet("count") {
			v.Set("limit", "0")
			v.Set("count", "true")
		} else {
			options := v.Get("options")
			if options == "" {
				options = "count"
			} else {
				options = options + ",count"
			}
			v.Set("options", options)
			v.Set("limit", fmt.Sprintf("%d", limit))
			v.Set("offset", fmt.Sprintf("%d", page*limit))
		}
		client.SetQuery(v)

		if c.Bool("acceptJson") {
			client.SetAcceptJSON()
		} else if c.Bool("acceptGeoJson") {
			client.SetAcceptGeoJSON()
		}

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		if c.IsSet("count") {
			count, err := client.ResultsCount(res)
			if err != nil {
				return &ngsiCmdError{funcName, 3, "ResultsCount error", nil}
			}
			fmt.Fprintln(ngsi.StdWriter, count)
			break
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 4, "ResultsCount error", err}
		}
		if count == 0 {
			break
		}

		if client.IsSafeString() {
			body, err = ngsilib.JSONSafeStringDecode(body)
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
		}

		err = entitiesPrint(ngsi, body, &buf, c.Bool("pretty"), lines, false, verbose, c.Bool("acceptGeoJson"))
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}
	if verbose {
		buf.bufferClose()
	}
	return nil
}

func entitiesPrint(ngsi *ngsilib.NGSI, body []byte, buf *jsonBuffer, pretty, lines, values, verbose, geoJSON bool) error {
	const funcName = "entitiesPrint"
	const geoJSONFeatures = `{"type":"FeatureCollection","features":`
	var err error

	if geoJSON {
		if bytes.HasPrefix(body, []byte(geoJSONFeatures)) && bytes.HasSuffix(body, []byte(`}`)) {
			body = body[len(geoJSONFeatures) : len(body)-1]
		} else {
			return &ngsiCmdError{funcName, 1, "geojson error: " + string(body), err}
		}
	}

	if lines {
		if values {
			var values [][]interface{}
			err = ngsilib.JSONUnmarshal(body, &values)
			if err != nil {
				return &ngsiCmdError{funcName, 2, err.Error(), err}
			}
			for _, e := range values {
				b, err := ngsilib.JSONMarshal(&e)
				if err != nil {
					return &ngsiCmdError{funcName, 3, err.Error(), err}
				}
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		} else {
			var entities entitiesRespose
			err = ngsilib.JSONUnmarshal(body, &entities)
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
			for _, e := range entities {
				b, err := ngsilib.JSONMarshal(&e)
				if err != nil {
					return &ngsiCmdError{funcName, 5, err.Error(), err}
				}
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		}
	} else if verbose {
		if pretty {
			newBuf := new(bytes.Buffer)
			indent := ""
			if geoJSON {
				indent = "  "
			}
			err := ngsi.JSONConverter.Indent(newBuf, body, indent, "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 6, err.Error(), err}
			}
			buf.bufferWrite(newBuf.Bytes())
		} else {
			buf.bufferWrite(body)
		}
	} else {
		var entities entitiesRespose
		err = ngsilib.JSONUnmarshal(body, &entities)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		for _, e := range entities {
			fmt.Fprintln(ngsi.StdWriter, e["id"])
		}
	}
	return nil
}

func entitiesCount(c *cli.Context) error {
	const funcName = "entitiesCount"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	client.SetPath("/entities")

	args := []string{"idPattern", "typePattern", "query", "mq", "georel", "geometry", "coords"}
	v := parseOptions(c, args, nil)

	if c.IsSet("type") {
		v.Set("type", c.String("type"))
	}

	if client.IsNgsiLd() {
		v.Set("limit", "0")
		v.Set("count", "true")
	} else {
		v.Set("limit", "1")
		if c.Bool("skipForwarding") {
			v.Set("options", "count,skipForwarding")
		} else {
			v.Set("options", "count")
		}
		v.Set("attrs", "__NONE")
	}
	client.SetQuery(v)

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
