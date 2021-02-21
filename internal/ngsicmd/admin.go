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

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"

	"github.com/urfave/cli/v2"
)

func adminLog(c *cli.Context) error {
	const funcName = "adminLog"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker", "cygnus", "perseo"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.ServerType == "broker" && client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "only available on NGSIv2", err}
	}

	path := "/admin/log"

	if c.IsSet("level") {
		level := strings.ToLower(c.String("level"))

		if !ngsilib.Contains([]string{"none", "fatal", "error", "warn", "info", "debug"}, level) {
			return &ngsiCmdError{funcName, 4, "log level error: " + level + " (none, fatal, error, warn, info, debug)", nil}
		}

		client.SetPath(path)

		v := url.Values{}
		v.Set("level", strings.ToUpper(c.String("level")))
		client.SetQuery(&v)

		res, body, err := client.HTTPPut("")
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 8, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 9, err.Error(), err}
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

func adminTrace(c *cli.Context) error {
	const funcName = "adminTrace"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "only available on NGSIv2", err}
	}

	if isSetAND(c, []string{"set", "delete"}) {
		return &ngsiCmdError{funcName, 4, "specify either --set or --delete", nil}
	}

	path := "/log/trace"

	if c.IsSet("set") {
		if !c.IsSet("level") {
			return &ngsiCmdError{funcName, 5, "missing level", err}
		}

		path = path + "/" + c.String("level")
		client.SetPath(path)

		res, body, err := client.HTTPPut("")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
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
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 9, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 10, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 11, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
		fmt.Fprint(ngsi.StdWriter, string(body))
		if c.IsSet("logging") {
			ngsi.Logging(ngsilib.LogInfo, string(body))
		}
	}

	return nil
}

func adminMetrics(c *cli.Context) error {
	const funcName = "adminMetrics"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker", "perseo", "cygnus"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() && client.Server.ServerType == "broker" {
		return &ngsiCmdError{funcName, 3, "only available on NGSIv2", err}
	}

	if isSetAND(c, []string{"reset", "delete"}) {
		return &ngsiCmdError{funcName, 4, "specify either --reset or --delete", nil}
	}

	path := "/admin/metrics"
	if client.Server.ServerType == "cygnus" {
		path = "/v1/admin/metrics"
	}

	if c.IsSet("delete") {
		client.SetPath(path)

		res, body, err := client.HTTPDelete(nil)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		if res.StatusCode != http.StatusNoContent {
			return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
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
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 8, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 9, err.Error(), err}
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

func adminSemaphore(c *cli.Context) error {
	const funcName = "adminSemaphore"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "only available on NGSIv2", err}
	}

	client.SetPath("/admin/sem")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
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

func adminStatistics(c *cli.Context) error {
	const funcName = "adminStatistics"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker", "cygnus"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.ServerType == "broker" && client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "only available on NGSIv2", err}
	}

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
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 8, err.Error(), err}
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

func adminCacheStatistics(c *cli.Context) error {
	const funcName = "adminCacheStatistics"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.IsNgsiLd() {
		return &ngsiCmdError{funcName, 3, "only available on NGSIv2", err}
	}

	path := "/cache/statistics"

	if c.IsSet("delete") {
		client.SetPath(path)

		res, body, err := client.HTTPDelete(nil)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
	} else {
		client.SetPath(path)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 8, err.Error(), err}
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
