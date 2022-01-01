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

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func entitiesList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entitiesList"

	if client.IsNgsiLd() {
		if c.IsSetOR([]string{"typePattern", "mq", "metadata", "value", "uniq"}) {
			return ngsierr.New(funcName, 1, "cannot specfiy typePattern, mq, metadata, value or uniq", nil)
		}
		return entitiesListLD(c, ngsi, client)
	}
	if c.IsSetOR([]string{"link", "acceptJson", "acceptGeoJson"}) {
		return ngsierr.New(funcName, 2, "cannot specfiy link acceptJson or acceptGeoJson", nil)
	}
	return entitiesListV2(c, ngsi, client)
}

func entitiesListV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entitiesList"

	attrs := "__NONE"

	if c.IsSet("attrs") {
		attrs = c.String("attrs")
	}

	page := 0
	count := 0
	limit := 100

	verbose := c.IsSet("verbose")
	if c.Bool("pretty") || c.Bool("keyValues") || c.IsSetOR([]string{"attrs", "metadata", "orderBy"}) {
		verbose = true
	}
	values := c.IsSet("values")
	if values {
		verbose = true
	}
	lines := c.Bool("lines")

	buf := ngsilib.NewJsonBuffer()
	if verbose {
		buf.BufferOpen(ngsi.StdWriter, false, false)
		attrs = ""
	}

	for {
		client.SetPath("/entities")

		args := []string{"id", "type", "idPattern", "typePattern", "query", "mq", "georel", "geometry", "coords", "attrs", "metadata", "orderBy"}
		opts := []string{"keyValues", "values", "unique", "skipForwarding"}
		v := ngsicli.ParseOptions(c, args, opts)

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
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
		}

		if c.IsSet("count") {
			count, err := client.ResultsCount(res)
			if err != nil {
				return ngsierr.New(funcName, 3, "ResultsCount error", nil)
			}
			fmt.Fprintln(ngsi.StdWriter, count)
			break
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return ngsierr.New(funcName, 4, "ResultsCount error", err)
		}
		if count == 0 {
			break
		}

		if client.IsSafeString() {
			body, err = ngsilib.JSONSafeStringDecode(body)
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
		}

		err = entitiesPrint(ngsi, body, buf, c.Bool("pretty"), lines, values, verbose, false)
		if err != nil {
			return ngsierr.New(funcName, 6, err.Error(), err)
		}

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}
	if verbose {
		buf.BufferClose()
	}
	return nil
}

func entitiesListLD(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entitiesListLD"

	idPattern := ""
	if !c.IsSetOR([]string{"type", "idPattern", "query", "attrs", "georel"}) {
		idPattern = ".*"
	}

	page := 0
	count := 0
	limit := 100

	verbose := c.IsSet("verbose")
	if c.IsSetOR([]string{"pretty", "keyValues", "acceptGeoJson", "attrs", "orderBy"}) {
		verbose = true
	}
	lines := c.Bool("lines")

	buf := ngsilib.NewJsonBuffer()
	if verbose {
		buf.BufferOpen(ngsi.StdWriter, c.Bool("acceptGeoJson"), c.Bool("pretty"))
	}

	for {
		client.SetPath("/entities")

		args := []string{"id", "type", "idPattern", "query", "georel", "geometry", "coords", "attrs", "orderBy"}
		opts := []string{"keyValues"}
		v := ngsicli.ParseOptions(c, args, opts)

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
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
		}

		if c.IsSet("count") {
			count, err := client.ResultsCount(res)
			if err != nil {
				return ngsierr.New(funcName, 3, "ResultsCount error", nil)
			}
			fmt.Fprintln(ngsi.StdWriter, count)
			break
		}

		count, err = client.ResultsCount(res)
		if err != nil {
			return ngsierr.New(funcName, 4, "ResultsCount error", err)
		}
		if count == 0 {
			break
		}

		if client.IsSafeString() {
			body, err = ngsilib.JSONSafeStringDecode(body)
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
		}

		err = entitiesPrint(ngsi, body, buf, c.Bool("pretty"), lines, false, verbose, c.Bool("acceptGeoJson"))
		if err != nil {
			return ngsierr.New(funcName, 6, err.Error(), err)
		}

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}
	if verbose {
		buf.BufferClose()
	}
	return nil
}

func entitiesPrint(ngsi *ngsilib.NGSI, body []byte, buf *ngsilib.JsonBuffer, pretty, lines, values, verbose, geoJSON bool) error {
	const funcName = "entitiesPrint"
	const geoJSONFeatures = `{"type":"FeatureCollection","features":`
	var err error

	if geoJSON {
		if bytes.HasPrefix(body, []byte(geoJSONFeatures)) && bytes.HasSuffix(body, []byte(`}`)) {
			body = body[len(geoJSONFeatures) : len(body)-1]
		} else {
			return ngsierr.New(funcName, 1, "geojson error: "+string(body), err)
		}
	}

	if lines {
		if values {
			var values [][]interface{}
			err = ngsilib.JSONUnmarshal(body, &values)
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error(), err)
			}
			for _, e := range values {
				b, err := ngsilib.JSONMarshal(&e)
				if err != nil {
					return ngsierr.New(funcName, 3, err.Error(), err)
				}
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		} else {
			var entities ngsilib.EntitiesRespose
			err = ngsilib.JSONUnmarshal(body, &entities)
			if err != nil {
				return ngsierr.New(funcName, 4, err.Error(), err)
			}
			for _, e := range entities {
				b, err := ngsilib.JSONMarshal(&e)
				if err != nil {
					return ngsierr.New(funcName, 5, err.Error(), err)
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
				return ngsierr.New(funcName, 6, err.Error(), err)
			}
			buf.BufferWrite(newBuf.Bytes())
		} else {
			buf.BufferWrite(body)
		}
	} else {
		var entities ngsilib.EntitiesRespose
		err = ngsilib.JSONUnmarshal(body, &entities)
		if err != nil {
			return ngsierr.New(funcName, 7, err.Error(), err)
		}
		for _, e := range entities {
			fmt.Fprintln(ngsi.StdWriter, e["id"])
		}
	}
	return nil
}

func entitiesCount(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "entitiesCount"

	client.SetPath("/entities")

	args := []string{"idPattern", "typePattern", "query", "mq", "georel", "geometry", "coords"}
	v := ngsicli.ParseOptions(c, args, nil)

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
