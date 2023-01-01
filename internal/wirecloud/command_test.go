/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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

package wirecloud

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
		{args: []string{"preferences", "get", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "list", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "get", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "download", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "install", "--host", "wirecloud", "--file", "NGSIGO_ngsigo-widget_1.0.0.wgt"}, rc: 1},
		{args: []string{"macs", "install", "--host", "wirecloud", "NGSIGO_ngsigo-widget_1.0.0.wgt"}, rc: 1},
		{args: []string{"macs", "uninstall", "--host", "wirecloud"}, rc: 1},
		{args: []string{"workspaces", "list", "--host", "wirecloud"}, rc: 1},
		{args: []string{"workspaces", "get", "--host", "wirecloud", "-wid", "1"}, rc: 1},
		{args: []string{"workspaces", "get", "--host", "wirecloud", "1"}, rc: 1},
		{args: []string{"tabs", "list", "--host", "wirecloud", "--wid", "33"}, rc: 1},
		{args: []string{"tabs", "list", "--host", "wirecloud", "33"}, rc: 1},
		{args: []string{"tabs", "get", "--host", "wirecloud", "--wid", "33", "--tid", "36"}, rc: 1},
		{args: []string{"tabs", "get", "--host", "wirecloud", "33", "36"}, rc: 1},
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
