/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

package ngsimain

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
)

func TestNGSICommand(t *testing.T) {
	cases := []struct {
		args  []string
		flags []ngsicli.Flag
		rc    int
	}{
		//{args: []string{}, rc: 0},
		{args: []string{"admin", "log", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "trace", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "semaphore", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "metrics", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "statistics", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "cacheStatistics", "--host", "orion"}, rc: 1},
		{args: []string{"cp", "--type", "abc", "--host", "orion", "--host2", "orion-ld"}, rc: 1},
		{args: []string{"wc", "entities", "--host", "orion"}, rc: 1},
		{args: []string{"wc", "subscriptions", "--host", "orion"}, rc: 1},
		{args: []string{"wc", "registrations", "--host", "orion"}, rc: 1},
		{args: []string{"wc", "types", "--host", "orion"}, rc: 1},
		{args: []string{"ls", "--host", "orion"}, rc: 1},
		{args: []string{"rm", "--host", "orion", "--type", "abc"}, rc: 1},
		{args: []string{"receiver", "--host", "orion", "--https"}, rc: 1},
		{args: []string{"regproxy", "server", "--host", "orion", "--https"}, rc: 1},
		{args: []string{"regproxy", "health", "--host", "regproxy"}, rc: 1},
		{args: []string{"regproxy", "config", "--host", "regproxy"}, rc: 1},
		{args: []string{"tokenproxy", "server", "--https"}, rc: 1},
		{args: []string{"tokenproxy", "health", "--host", "tokenproxy"}, rc: 1},
		{args: []string{"queryproxy", "server", "--host", "orion", "--https"}, rc: 1},
		{args: []string{"queryproxy", "health", "--host", "queryproxy"}, rc: 1},
		{args: []string{"template", "registration", "--ngsiType", "v2"}, rc: 0},
		{args: []string{"template", "subscription", "--ngsiType", "ld"}, rc: 0},
		{args: []string{"version", "--host", "orion"}, rc: 1},
		{args: []string{"man"}, rc: 0},
		{args: []string{"apis", "--host", "orion"}, rc: 1},
		{args: []string{"health", "--host", "orion"}, rc: 1},
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
		{args: []string{"context", "server"}, rc: 1},
		{args: []string{"settings", "list"}, rc: 0},
		{args: []string{"settings", "clear"}, rc: 0},
		{args: []string{"settings", "delete"}, rc: 1},
		{args: []string{"settings", "previousArgs"}, rc: 1},
		{args: []string{"append", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"create", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"create", "entity", "--host", "orion"}, rc: 1},
		{args: []string{"create", "registration", "--host", "orion"}, rc: 1},
		{args: []string{"create", "subscription", "--host", "orion", "--url", "abc"}, rc: 1},
		{args: []string{"create", "ldContext", "--host", "orion-ld"}, rc: 1},
		{args: []string{"delete", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"delete", "entity", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "attr", "--host", "orion", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"delete", "registration", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "subscription", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "ldContext", "--host", "orion-ld"}, rc: 1},
		{args: []string{"get", "attr", "--host", "orion", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"get", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "entity", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "entities", "--host", "orion"}, rc: 1},
		{args: []string{"get", "type", "--host", "orion", "--type", "abc"}, rc: 1},
		{args: []string{"get", "registration", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "subscription", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"get", "ldContext", "--host", "orion-ld"}, rc: 1},
		{args: []string{"list", "entities", "--host", "orion"}, rc: 1},
		{args: []string{"list", "registrations", "--host", "orion"}, rc: 1},
		{args: []string{"list", "subscriptions", "--host", "orion"}, rc: 1},
		{args: []string{"list", "types", "--host", "orion"}, rc: 1},
		{args: []string{"list", "attributes", "--host", "orion"}, rc: 1},
		{args: []string{"list", "ldContexts", "--host", "orion-ld"}, rc: 1},
		{args: []string{"replace", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"replace", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"update", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"update", "attrs", "--host", "orion", "--id", "abc"}, rc: 1},
		{args: []string{"update", "attr", "--host", "orion", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"update", "subscription", "--host", "orion", "--id", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"upsert", "entity", "--host", "orion"}, rc: 1},
		{args: []string{"upsert", "entities", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"token", "--host", "orion"}, rc: 1},
		{args: []string{"debug", "--host", "orion"}, rc: 0},
		{args: []string{"hget", "attr", "--host", "comet"}, rc: 1},
		{args: []string{"hget", "attrs", "--host", "ql"}, rc: 1},
		{args: []string{"hget", "entities", "--host", "ql"}, rc: 1},
		{args: []string{"hdelete", "attr", "--host", "comet"}, rc: 1},
		{args: []string{"hdelete", "entity", "--host", "comet"}, rc: 1},
		{args: []string{"hdelete", "entities", "--host", "comet"}, rc: 0},
		{args: []string{"services", "list", "--host", "iota"}, rc: 1},
		{args: []string{"services", "create", "--host", "iota"}, rc: 1},
		{args: []string{"services", "update", "--host", "iota"}, rc: 1},
		{args: []string{"services", "delete", "--host", "iota"}, rc: 1},
		{args: []string{"devices", "list", "--host", "iota"}, rc: 1},
		{args: []string{"devices", "create", "--host", "iota", "--data", "@"}, rc: 1},
		{args: []string{"devices", "get", "--host", "iota"}, rc: 1},
		{args: []string{"devices", "update", "--host", "iota", "--data", "@"}, rc: 1},
		{args: []string{"devices", "delete", "--host", "iota"}, rc: 1},
		{args: []string{"rules", "list", "--host", "perseo"}, rc: 1},
		{args: []string{"rules", "create", "--host", "perseo", "--data", "@"}, rc: 1},
		{args: []string{"rules", "get", "--host", "perseo"}, rc: 1},
		{args: []string{"rules", "delete", "--host", "perseo"}, rc: 1},
		{args: []string{"applications", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "update", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "update", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "permissions", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "assign", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "roles", "unassign", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "permissions", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "permissions", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "permissions", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "permissions", "update", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "permissions", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "pep", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "pep", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "pep", "reset", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "pep", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "iota", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "iota", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "iota", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "iota", "reset", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "iota", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "users", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "users", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "users", "assign", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "users", "unassign", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "organizations", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "organizations", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "organizations", "assign", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "organizations", "unassign", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "trusted", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "trusted", "add", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "trusted", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "update", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "users", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "users", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "users", "add", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "users", "remove", "--host", "keyrock"}, rc: 1},
		{args: []string{"users", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"users", "get", "--host", "keyrock"}, rc: 1},
		{args: []string{"users", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"users", "update", "--host", "keyrock"}, rc: 1},
		{args: []string{"users", "delete", "--host", "keyrock"}, rc: 1},
		{args: []string{"providers", "--host", "keyrock"}, rc: 1},
		{args: []string{"admin", "appenders", "list", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "appenders", "get", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "appenders", "create", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"admin", "appenders", "update", "--host", "orion", "--data", "@"}, rc: 1},
		{args: []string{"admin", "appenders", "delete", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "loggers", "list", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "loggers", "get", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "loggers", "create", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "loggers", "update", "--host", "orion"}, rc: 1},
		{args: []string{"admin", "loggers", "delete", "--host", "orion"}, rc: 1},
		{args: []string{"namemappings", "list", "--host", "cygnus"}, rc: 1},
		{args: []string{"namemappings", "create", "--host", "cygnus"}, rc: 1},
		{args: []string{"namemappings", "update", "--host", "cygnus"}, rc: 1},
		{args: []string{"namemappings", "delete", "--host", "cygnus"}, rc: 1},
		{args: []string{"groupingrules", "list", "--host", "cygnus"}, rc: 1},
		{args: []string{"groupingrules", "create", "--host", "cygnus"}, rc: 1},
		{args: []string{"groupingrules", "update", "--host", "cygnus"}, rc: 1},
		{args: []string{"groupingrules", "delete", "--host", "cygnus"}, rc: 1},
		{args: []string{"admin", "scorpio", "list", "--host", "scorpio"}, rc: 1},
		{args: []string{"admin", "scorpio", "types", "--host", "scorpio"}, rc: 1},
		{args: []string{"admin", "scorpio", "localtypes", "--host", "scorpio"}, rc: 1},
		{args: []string{"admin", "scorpio", "stats", "--host", "scorpio"}, rc: 1},
		{args: []string{"admin", "scorpio", "health", "--host", "scorpio"}, rc: 1},
		{args: []string{"list", "tentities", "--host", "orion-ld"}, rc: 1},
		{args: []string{"create", "tentity", "--host", "orion-ld"}, rc: 1},
		{args: []string{"get", "tentity", "--host", "orion-ld"}, rc: 1},
		{args: []string{"delete", "tentity", "--host", "orion-ld"}, rc: 1},
		{args: []string{"append", "tattrs", "--host", "orion-ld"}, rc: 1},
		{args: []string{"update", "tattr", "--host", "orion-ld"}, rc: 1},
		{args: []string{"delete", "tattr", "--host", "orion-ld"}, rc: 1},
		{args: []string{"preferences", "get", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "list", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "get", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "download", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "install", "--host", "wirecloud"}, rc: 1},
		{args: []string{"macs", "uninstall", "--host", "wirecloud"}, rc: 1},
		{args: []string{"workspaces", "list", "--host", "wirecloud"}, rc: 1},
		{args: []string{"workspaces", "get", "--host", "wirecloud"}, rc: 1},
		{args: []string{"tabs", "list", "--host", "wirecloud"}, rc: 1},
		{args: []string{"tabs", "get", "--host", "wirecloud"}, rc: 1},
	}

	for _, c := range cases {
		ngsi := helper.SetupTestInitNGSI()

		ngsi.HTTP = &helper.MockHTTP{ReqRes: []helper.MockHTTPReqRes{{StatusCode: http.StatusBadRequest}}}
		syslog := []string{"ngsi", "--stderr", "off"}
		args := append(syslog, c.args...)
		in := new(bytes.Buffer)
		out := new(bytes.Buffer)
		err := new(bytes.Buffer)

		if rc := Run(args, in, out, err); rc != c.rc {
			fmt.Printf("*** %s *** rc expected:%d, actual:%d)\n", strings.Join(c.args, " "), c.rc, rc)
		}
	}
}

func TestInitCmdNormal(t *testing.T) {
	_ = helper.SetupTestInitNGSI()

	args := []string{"ngsi", "man"}
	in := new(bytes.Buffer)
	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	if Run(args, in, out, err) != 0 {
		t.Error(fmt.Printf("*** %s *** \n", args[0]))
	}
}
