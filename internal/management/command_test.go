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

package management

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
		{args: []string{"broker", "add", "--host", "orion"}, rc: 1},
		{args: []string{"broker", "delete", "--host", "orion"}, rc: 0},
		{args: []string{"broker", "get", "--host", "orion"}, rc: 0},
		{args: []string{"broker", "list"}, rc: 0},
		{args: []string{"broker", "update", "--host", "comet"}, rc: 0},
		{args: []string{"server", "add", "--host", "comet"}, rc: 1},
		{args: []string{"server", "delete", "--host", "comet"}, rc: 0},
		{args: []string{"server", "get", "--host", "comet"}, rc: 0},
		{args: []string{"server", "list"}, rc: 0},
		{args: []string{"server", "update", "--host", "orion"}, rc: 0},
		{args: []string{"context", "add", "--name", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"context", "delete", "--name", "abc"}, rc: 1},
		{args: []string{"context", "list"}, rc: 0},
		{args: []string{"context", "update", "--name", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"context", "server", "--data", "@"}, rc: 1},
		{args: []string{"settings", "list"}, rc: 0},
		{args: []string{"settings", "clear"}, rc: 0},
		{args: []string{"settings", "delete"}, rc: 1},
		{args: []string{"settings", "previousArgs"}, rc: 1},
		{args: []string{"token", "--host", "orion"}, rc: 1},
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
