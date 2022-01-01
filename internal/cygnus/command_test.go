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

package cygnus

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestNewNgsiApp(t *testing.T) {
	actual := NewNgsiApp()

	assert.NotEqual(t, nil, actual)
}

func TestNGSICommand(t *testing.T) {
	cases := []struct {
		args  []string
		flags []ngsicli.Flag
		rc    int
	}{
		{args: []string{"namemappings", "list", "--host", "cygnus"}, rc: 1},
		{args: []string{"namemappings", "create", "--host", "cygnus", "--data", "@"}, rc: 1},
		{args: []string{"namemappings", "update", "--host", "cygnus", "--data", "@"}, rc: 1},
		{args: []string{"namemappings", "delete", "--host", "cygnus", "--data", "@"}, rc: 1},
		{args: []string{"groupingrules", "list", "--host", "cygnus"}, rc: 1},
		{args: []string{"groupingrules", "create", "--host", "cygnus", "--data", "@"}, rc: 1},
		{args: []string{"groupingrules", "update", "--host", "cygnus", "--id", "1", "--data", "@"}, rc: 1},
		{args: []string{"groupingrules", "delete", "--host", "cygnus", "--id", "1"}, rc: 1},
		{args: []string{"admin", "appenders", "list", "--host", "cygnus"}, rc: 1},
		{args: []string{"admin", "appenders", "get", "--host", "cygnus", "--name", "test"}, rc: 1},
		{args: []string{"admin", "appenders", "create", "--host", "cygnus", "--data", "@"}, rc: 1},
		{args: []string{"admin", "appenders", "update", "--host", "cygnus", "--name", "test", "--data", "@"}, rc: 1},
		{args: []string{"admin", "appenders", "delete", "--host", "cygnus", "--name", "test"}, rc: 1},
		{args: []string{"admin", "loggers", "list", "--host", "cygnus"}, rc: 1},
		{args: []string{"admin", "loggers", "get", "--host", "cygnus", "--name", "org.mongodb"}, rc: 1},
		{args: []string{"admin", "loggers", "create", "--host", "cygnus", "--data", "@"}, rc: 1},
		{args: []string{"admin", "loggers", "update", "--host", "cygnus", "--data", "@"}, rc: 1},
		{args: []string{"admin", "loggers", "delete", "--host", "cygnus", "--name", "org.mongodb"}, rc: 1},
	}

	for _, c := range cases {
		ngsi := helper.SetupTestInitNGSI()

		ngsi.HTTP = &helper.MockHTTP{ReqRes: []helper.MockHTTPReqRes{{StatusCode: http.StatusBadRequest}}}
		syslog := []string{"ngsi", "--stderr", "off"}
		args := append(syslog, c.args...)

		app := NewNgsiApp()

		err := app.Run(args)
		rc := 0
		if err != nil {
			if err.(*ngsierr.NgsiError).Message == "missing required flags" {
				fmt.Println(strings.Join(args, " "))
				os.Exit(1)
			}
			rc = 1
		}

		if rc != c.rc {
			fmt.Printf("*** %s *** rc expected:%d, actual:%d)\n", strings.Join(c.args, " "), c.rc, rc)
		}
	}
}
