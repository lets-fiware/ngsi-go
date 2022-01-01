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
		{args: []string{"create", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"create", "entity", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"create", "registration", "--host", "orion"}, rc: 1},
		{args: []string{"create", "subscription", "--host", "orion", "--url", "abc"}, rc: 1},
		{args: []string{"create", "ldContext", "--host", "orion-ld", "--data", "@"}, rc: 1},
		{args: []string{"delete", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"delete", "entity", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "attr", "--host", "orion", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"delete", "registration", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "subscription", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "ldContext", "--host", "orion-ld", "--id", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"}, rc: 1},
		{args: []string{"get", "attr", "--host", "orion", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"get", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "entity", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "entities", "--host", "orion"}, rc: 1},
		{args: []string{"get", "type", "--host", "orion", "--type", "abc"}, rc: 1},
		{args: []string{"get", "registration", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "subscription", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "ldContext", "--host", "orion-ld", "--id", "2fa4dbc4-ece8-11eb-a645-0242c0a8a010"}, rc: 1},
		{args: []string{"list", "entities", "--host", "orion"}, rc: 1},
		{args: []string{"list", "registrations", "--host", "orion"}, rc: 1},
		{args: []string{"list", "subscriptions", "--host", "orion"}, rc: 1},
		{args: []string{"list", "types", "--host", "orion"}, rc: 1},
		{args: []string{"list", "attributes", "--host", "orion"}, rc: 1},
		{args: []string{"list", "ldContexts", "--host", "orion-ld"}, rc: 1},
		{args: []string{"append", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"replace", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"replace", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"update", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"update", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"update", "attr", "--host", "orion", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"update", "subscription", "--host", "orion", "--id", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"upsert", "entity", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"upsert", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"ls", "--host", "orion"}, rc: 1},
		{args: []string{"wc", "entities", "--host", "orion"}, rc: 1},
		{args: []string{"wc", "subscriptions", "--host", "orion"}, rc: 1},
		{args: []string{"wc", "registrations", "--host", "orion"}, rc: 1},
		{args: []string{"wc", "types", "--host", "orion"}, rc: 1},
		{args: []string{"template", "registration", "--ngsiType", "v2"}, rc: 0},
		{args: []string{"template", "subscription", "--ngsiType", "ld"}, rc: 0},
		{args: []string{"list", "tentities", "--host", "orion-ld"}, rc: 1},
		{args: []string{"create", "tentity", "--host", "orion-ld", "--data", "@"}, rc: 1},
		{args: []string{"get", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"}, rc: 1},
		{args: []string{"delete", "tentity", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100"}, rc: 1},
		{args: []string{"append", "tattrs", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--data", "@"}, rc: 1},
		{args: []string{"update", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature", "--instanceId", "instance001", "--data", "@"}, rc: 1},
		{args: []string{"delete", "tattr", "--host", "orion-ld", "--id", "urn:ngsi-ld:sensor100", "--attr", "temperature"}, rc: 1},
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
