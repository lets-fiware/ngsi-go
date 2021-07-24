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
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNGSICommand(t *testing.T) {
	cases := []struct {
		args []string
		rc   int
	}{
		//{args: []string{}, rc: 0},
		{args: []string{"admin", "log"}, rc: 1},
		{args: []string{"admin", "trace"}, rc: 1},
		{args: []string{"admin", "semaphore"}, rc: 1},
		{args: []string{"admin", "metrics"}, rc: 1},
		{args: []string{"admin", "statistics"}, rc: 1},
		{args: []string{"admin", "cacheStatistics"}, rc: 1},
		{args: []string{"cp", "--type", "abc", "--host2", "abc"}, rc: 1},
		{args: []string{"wc", "entities"}, rc: 1},
		{args: []string{"wc", "subscriptions"}, rc: 1},
		{args: []string{"wc", "registrations"}, rc: 1},
		{args: []string{"wc", "types"}, rc: 1},
		{args: []string{"ls"}, rc: 1},
		{args: []string{"rm", "--type", "abc"}, rc: 1},
		{args: []string{"receiver"}, rc: 1},
		{args: []string{"regproxy", "server"}, rc: 1},
		{args: []string{"regproxy", "health"}, rc: 1},
		{args: []string{"regproxy", "config"}, rc: 1},
		{args: []string{"tokenproxy", "server"}, rc: 1},
		{args: []string{"tokenproxy", "health"}, rc: 1},
		{args: []string{"queryproxy", "server"}, rc: 1},
		{args: []string{"queryproxy", "health"}, rc: 1},
		{args: []string{"template", "registration"}, rc: 1},
		{args: []string{"template", "subscription", "--url", "abc"}, rc: 1},
		{args: []string{"version"}, rc: 1},
		{args: []string{"man"}, rc: 1},
		{args: []string{"apis"}, rc: 1},
		{args: []string{"health"}, rc: 1},
		{args: []string{"broker", "add"}, rc: 1},
		{args: []string{"broker", "delete"}, rc: 1},
		{args: []string{"broker", "get"}, rc: 1},
		{args: []string{"broker", "list"}, rc: 1},
		{args: []string{"broker", "update"}, rc: 1},
		{args: []string{"server", "add"}, rc: 1},
		{args: []string{"server", "delete"}, rc: 1},
		{args: []string{"server", "get"}, rc: 1},
		{args: []string{"server", "list"}, rc: 1},
		{args: []string{"server", "update"}, rc: 1},
		{args: []string{"context", "add", "--name", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"context", "delete", "--name", "abc"}, rc: 1},
		{args: []string{"context", "list"}, rc: 1},
		{args: []string{"context", "update", "--name", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"context", "server"}, rc: 1},
		{args: []string{"settings", "list"}, rc: 1},
		{args: []string{"settings", "clear"}, rc: 1},
		{args: []string{"settings", "delete"}, rc: 1},
		{args: []string{"settings", "previousArgs"}, rc: 1},
		{args: []string{"append", "attrs", "--id", "abc"}, rc: 1},
		{args: []string{"create", "entities"}, rc: 1},
		{args: []string{"create", "entity"}, rc: 1},
		{args: []string{"create", "registration"}, rc: 1},
		{args: []string{"create", "subscription", "--url", "abc"}, rc: 1},
		{args: []string{"delete", "entities"}, rc: 1},
		{args: []string{"delete", "entity", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "attr", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"delete", "registration", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "subscription", "--id", "abc"}, rc: 1},
		{args: []string{"get", "attr", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"get", "attrs", "--id", "abc"}, rc: 1},
		{args: []string{"get", "entity", "--id", "abc"}, rc: 1},
		{args: []string{"get", "entities"}, rc: 1},
		{args: []string{"get", "type", "--type", "abc"}, rc: 1},
		{args: []string{"get", "registration", "--id", "abc"}, rc: 1},
		{args: []string{"get", "subscription", "--id", "abc"}, rc: 1},
		{args: []string{"list", "entities"}, rc: 1},
		{args: []string{"list", "registrations"}, rc: 1},
		{args: []string{"list", "subscriptions"}, rc: 1},
		{args: []string{"list", "types"}, rc: 1},
		{args: []string{"replace", "entities"}, rc: 1},
		{args: []string{"replace", "attrs", "--id", "abc"}, rc: 1},
		{args: []string{"update", "entities"}, rc: 1},
		{args: []string{"update", "attrs", "--id", "abc"}, rc: 1},
		{args: []string{"update", "attr", "--id", "abc", "--attr", "abc"}, rc: 1},
		{args: []string{"update", "subscription", "--id", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"upsert", "entity"}, rc: 1},
		{args: []string{"upsert", "entities"}, rc: 1},
		{args: []string{"token"}, rc: 1},
		{args: []string{"debug"}, rc: 1},
		{args: []string{"hget", "attr"}, rc: 1},
		{args: []string{"hget", "attrs"}, rc: 1},
		{args: []string{"hget", "entities"}, rc: 1},
		{args: []string{"hdelete", "attr"}, rc: 1},
		{args: []string{"hdelete", "entity"}, rc: 1},
		{args: []string{"hdelete", "entities"}, rc: 1},
		{args: []string{"services", "list"}, rc: 1},
		{args: []string{"services", "create"}, rc: 1},
		{args: []string{"services", "update"}, rc: 1},
		{args: []string{"services", "delete"}, rc: 1},
		{args: []string{"devices", "list"}, rc: 1},
		{args: []string{"devices", "create"}, rc: 1},
		{args: []string{"devices", "get"}, rc: 1},
		{args: []string{"devices", "update"}, rc: 1},
		{args: []string{"devices", "delete"}, rc: 1},
		{args: []string{"rules", "list"}, rc: 1},
		{args: []string{"rules", "create"}, rc: 1},
		{args: []string{"rules", "get"}, rc: 1},
		{args: []string{"rules", "delete"}, rc: 1},
		{args: []string{"applications", "list"}, rc: 1},
		{args: []string{"applications", "get"}, rc: 1},
		{args: []string{"applications", "create"}, rc: 1},
		{args: []string{"applications", "update"}, rc: 1},
		{args: []string{"applications", "delete"}, rc: 1},
		{args: []string{"applications", "roles", "list"}, rc: 1},
		{args: []string{"applications", "roles", "get"}, rc: 1},
		{args: []string{"applications", "roles", "create"}, rc: 1},
		{args: []string{"applications", "roles", "update"}, rc: 1},
		{args: []string{"applications", "roles", "delete"}, rc: 1},
		{args: []string{"applications", "roles", "permissions"}, rc: 1},
		{args: []string{"applications", "roles", "assign"}, rc: 1},
		{args: []string{"applications", "roles", "unassign"}, rc: 1},
		{args: []string{"applications", "permissions", "list"}, rc: 1},
		{args: []string{"applications", "permissions", "get"}, rc: 1},
		{args: []string{"applications", "permissions", "create"}, rc: 1},
		{args: []string{"applications", "permissions", "update"}, rc: 1},
		{args: []string{"applications", "permissions", "delete"}, rc: 1},
		{args: []string{"applications", "pep", "list"}, rc: 1},
		{args: []string{"applications", "pep", "create"}, rc: 1},
		{args: []string{"applications", "pep", "reset"}, rc: 1},
		{args: []string{"applications", "pep", "delete"}, rc: 1},
		{args: []string{"applications", "iota", "list"}, rc: 1},
		{args: []string{"applications", "iota", "get"}, rc: 1},
		{args: []string{"applications", "iota", "create"}, rc: 1},
		{args: []string{"applications", "iota", "reset"}, rc: 1},
		{args: []string{"applications", "iota", "delete"}, rc: 1},
		{args: []string{"applications", "users", "list"}, rc: 1},
		{args: []string{"applications", "users", "get"}, rc: 1},
		{args: []string{"applications", "users", "assign"}, rc: 1},
		{args: []string{"applications", "users", "unassign"}, rc: 1},
		{args: []string{"applications", "organizations", "list"}, rc: 1},
		{args: []string{"applications", "organizations", "get"}, rc: 1},
		{args: []string{"applications", "organizations", "assign"}, rc: 1},
		{args: []string{"applications", "organizations", "unassign"}, rc: 1},
		{args: []string{"applications", "trusted", "list"}, rc: 1},
		{args: []string{"applications", "trusted", "add"}, rc: 1},
		{args: []string{"applications", "trusted", "delete"}, rc: 1},
		{args: []string{"organizations", "list"}, rc: 1},
		{args: []string{"organizations", "get"}, rc: 1},
		{args: []string{"organizations", "create"}, rc: 1},
		{args: []string{"organizations", "update"}, rc: 1},
		{args: []string{"organizations", "delete"}, rc: 1},
		{args: []string{"organizations", "users", "list"}, rc: 1},
		{args: []string{"organizations", "users", "get"}, rc: 1},
		{args: []string{"organizations", "users", "add"}, rc: 1},
		{args: []string{"organizations", "users", "remove"}, rc: 1},
		{args: []string{"users", "list"}, rc: 1},
		{args: []string{"users", "get"}, rc: 1},
		{args: []string{"users", "create"}, rc: 1},
		{args: []string{"users", "update"}, rc: 1},
		{args: []string{"users", "delete"}, rc: 1},
		{args: []string{"providers"}, rc: 1},
		{args: []string{"admin", "appenders", "list"}, rc: 1},
		{args: []string{"admin", "appenders", "get"}, rc: 1},
		{args: []string{"admin", "appenders", "create"}, rc: 1},
		{args: []string{"admin", "appenders", "update"}, rc: 1},
		{args: []string{"admin", "appenders", "delete"}, rc: 1},
		{args: []string{"admin", "loggers", "list"}, rc: 1},
		{args: []string{"admin", "loggers", "get"}, rc: 1},
		{args: []string{"admin", "loggers", "create"}, rc: 1},
		{args: []string{"admin", "loggers", "update"}, rc: 1},
		{args: []string{"admin", "loggers", "delete"}, rc: 1},
		{args: []string{"namemappings", "list"}, rc: 1},
		{args: []string{"namemappings", "create"}, rc: 1},
		{args: []string{"namemappings", "update"}, rc: 1},
		{args: []string{"namemappings", "delete"}, rc: 1},
		{args: []string{"groupingrules", "list"}, rc: 1},
		{args: []string{"groupingrules", "create"}, rc: 1},
		{args: []string{"groupingrules", "update"}, rc: 1},
		{args: []string{"groupingrules", "delete"}, rc: 1},
		{args: []string{"admin", "scorpio", "list"}, rc: 1},
		{args: []string{"admin", "scorpio", "types"}, rc: 1},
		{args: []string{"admin", "scorpio", "localtypes"}, rc: 1},
		{args: []string{"admin", "scorpio", "stats"}, rc: 1},
		{args: []string{"admin", "scorpio", "health"}, rc: 1},
		{args: []string{"list", "tentities"}, rc: 1},
		{args: []string{"create", "tentity"}, rc: 1},
		{args: []string{"get", "tentity"}, rc: 1},
		{args: []string{"delete", "tentity"}, rc: 1},
		{args: []string{"append", "tattrs"}, rc: 1},
		{args: []string{"update", "tattr"}, rc: 1},
		{args: []string{"delete", "tattr"}, rc: 1},
		{args: []string{"preferences", "get"}, rc: 1},
		{args: []string{"macs", "list"}, rc: 1},
		{args: []string{"macs", "get"}, rc: 1},
		{args: []string{"macs", "download"}, rc: 1},
		{args: []string{"macs", "install"}, rc: 1},
		{args: []string{"macs", "uninstall"}, rc: 1},
		{args: []string{"workspaces", "list"}, rc: 1},
		{args: []string{"workspaces", "get"}, rc: 1},
		{args: []string{"tabs", "list"}, rc: 1},
		{args: []string{"tabs", "get"}, rc: 1},
	}

	for _, c := range cases {
		setupTest()

		syslog := []string{"ngsi", "--stderr", "no"}
		args := append(syslog, c.args...)
		in := new(bytes.Buffer)
		out := new(bytes.Buffer)
		err := new(bytes.Buffer)

		if Run(args, in, out, err) != c.rc {
			t.Error(fmt.Printf("*** %s *** is wrong (%d)\n", strings.Join(c.args, " "), c.rc))
		}
	}
}

func TestInitCmdNormal(t *testing.T) {
	setupTest()

	args := []string{"ngsi", "man"}
	in := new(bytes.Buffer)
	out := new(bytes.Buffer)
	err := new(bytes.Buffer)
	if Run(args, in, out, err) != 0 {
		t.Error(fmt.Printf("*** %s *** \n", args[0]))
	}
}

func TestNGSIMessage(t *testing.T) {

	e := errors.New("error message")
	s := message(e)

	assert.Equal(t, "error message", s)
}

func TestNGSIMessageLibError(t *testing.T) {

	e := &ngsilib.LibError{ErrNo: 1, Message: "error message"}
	s := message(e)

	assert.Equal(t, "001 error message", s)
}

func TestNGSIMessageCmdError(t *testing.T) {

	e := &ngsiCmdError{ErrNo: 1, Message: "error message"}
	s := message(e)

	assert.Equal(t, "001 error message", s)
}

func TestIsSetsORTrue(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--count"})
	actual := isSetOR(c, []string{"host"})
	expected := true

	assert.Equal(t, expected, actual)
}

func TestIsSetsORFalse(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")

	c := cli.NewContext(app, set, nil)

	actual := isSetOR(c, []string{"host"})
	expected := false

	assert.Equal(t, expected, actual)
}

func TestIsSetsANDTrue(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=orion", "--count"})
	actual := isSetAND(c, []string{"host", "count"})
	expected := true

	assert.Equal(t, expected, actual)
}

func TestIsSetsANDFalse(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host,type")
	setupFlagBool(set, "count")

	c := cli.NewContext(app, set, nil)

	actual := isSetAND(c, []string{"host"})
	expected := false

	assert.Equal(t, expected, actual)
}
