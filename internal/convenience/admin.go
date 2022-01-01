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

package convenience

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func adminLog(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "adminLog"

	path := "/admin/log"

	if c.IsSet("level") {
		level := strings.ToLower(c.String("level"))

		if !ngsilib.Contains([]string{"none", "fatal", "error", "warn", "info", "debug"}, level) {
			return ngsierr.New(funcName, 1, "log level error: "+level+" (none, fatal, error, warn, info, debug)", nil)
		}

		client.SetPath(path)

		v := url.Values{}
		v.Set("level", strings.ToUpper(c.String("level")))
		client.SetQuery(&v)

		res, body, err := client.HTTPPut("")
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 6, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}

		if c.IsSet("logging") {
			ngsi.Logging(ngsilib.LogInfo, string(body))
		}
	}

	return nil
}

func adminTrace(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "adminTrace"

	path := "/log/trace"

	if c.IsSet("set") {
		if !c.IsSet("level") {
			return ngsierr.New(funcName, 1, "missing level", nil)
		}

		path = path + "/" + c.String("level")
		client.SetPath(path)

		res, body, err := client.HTTPPut("")
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
		fmt.Fprint(ngsi.StdWriter, string(body))
		if c.IsSet("logging") {
			ngsi.Logging(ngsilib.LogInfo, string(body))
		}
	} else if c.IsSet("delete") {

		if c.IsSet("level") {
			path = path + "/" + c.String("level")
		}
		client.SetPath(path)

		res, body, err := client.HTTPDelete(nil)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return ngsierr.New(funcName, 6, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
		fmt.Fprint(ngsi.StdWriter, string(body))
		if c.IsSet("logging") {
			ngsi.Logging(ngsilib.LogInfo, string(body))
		}
	}

	return nil
}

func adminSemaphore(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "adminSemaphore"

	client.SetPath("/admin/sem")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
	} else {
		fmt.Fprint(ngsi.StdWriter, string(body))
	}

	if c.IsSet("logging") {
		ngsi.Logging(ngsilib.LogInfo, string(body))
	}

	return nil
}

func adminMetrics(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "adminMetrics"

	path := "/admin/metrics"
	if client.Server.ServerType == "cygnus" {
		path = "/v1/admin/metrics"
	}

	if c.IsSet("delete") {
		client.SetPath(path)

		res, body, err := client.HTTPDelete(nil)
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		if res.StatusCode != http.StatusNoContent {
			return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
	} else {
		client.SetPath(path)

		if c.IsSet("reset") {
			v := url.Values{}
			v.Set("reset", "true")
			client.SetQuery(&v)
		}

		res, body, err := client.HTTPGet()
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}

		if c.IsSet("logging") {
			ngsi.Logging(ngsilib.LogInfo, string(body))
		}
	}

	return nil
}

func adminStatistics(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "adminStatistics"

	path := "/statistics"
	if client.Server.ServerType == "cygnus" {
		path = "/v1/stats"
	}

	if c.IsSet("delete") {
		client.SetPath(path)
		var res *http.Response
		var body []byte
		var err error

		if client.Server.ServerType == "cygnus" {
			res, body, err = client.HTTPPut("")
		} else {
			res, body, err = client.HTTPDelete(nil)
		}
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}

		if c.IsSet("logging") {
			ngsi.Logging(ngsilib.LogInfo, string(body))
		}
	}

	return nil
}

func adminCacheStatistics(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "adminCacheStatistics"

	path := "/cache/statistics"

	if c.IsSet("delete") {
		client.SetPath(path)

		res, body, err := client.HTTPDelete(nil)
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}

		if c.IsSet("logging") {
			ngsi.Logging(ngsilib.LogInfo, string(body))
		}
	}

	return nil
}
