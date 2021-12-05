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

package ngsicli

import (
	"bytes"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestPrintVersion(t *testing.T) {
	buffer := &bytes.Buffer{}

	c := &Context{Ngsi: &ngsilib.NGSI{StdWriter: buffer}, App: &App{Version: "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab"}}
	token := newToken([]string{"--version"})

	actual := printVersion(c, token)

	assert.Equal(t, true, actual)
	expected := "ngsi version 0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab\n"
	assert.Equal(t, expected, buffer.String())
}

func TestPrintVersionV(t *testing.T) {
	buffer := &bytes.Buffer{}

	c := &Context{Ngsi: &ngsilib.NGSI{StdWriter: buffer}, App: &App{Version: "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab"}}
	token := newToken([]string{"-v"})

	actual := printVersion(c, token)

	assert.Equal(t, true, actual)
	expected := "ngsi version 0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab\n"
	assert.Equal(t, expected, buffer.String())
}

func TestPrintSerial(t *testing.T) {
	buffer := &bytes.Buffer{}

	c := &Context{Ngsi: &ngsilib.NGSI{StdWriter: buffer}, App: &App{Version: "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab"}}
	token := newToken([]string{"--serial"})

	actual := printVersion(c, token)

	assert.Equal(t, true, actual)
	expected := "00900"
	assert.Equal(t, expected, buffer.String())
}

func TestPrintVersionFalse(t *testing.T) {
	buffer := &bytes.Buffer{}

	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buffer}, App: &App{Version: "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab"}}
	token := newToken([]string{"ngsi"})

	actual := printVersion(c, token)

	assert.Equal(t, false, actual)
}

func TestPrintHelp(t *testing.T) {
	buffer := &bytes.Buffer{}

	c := &Context{
		Ngsi: &ngsilib.NGSI{StdWriter: buffer},
		App:  &App{Version: "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab", Usage: "command-line tool for FIWARE NGSI and NGSI-LD"},
	}

	printHelp(c)

	expected := "NAME:\n   ngsi - command-line tool for FIWARE NGSI and NGSI-LD\n\nUSAGE:\n    [global options] command [options] [arguments...]\n\nVERSION:\n   ngsi version 0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab\n\nCOMMANDS:\n   help, h  Shows a list of commands or help for one command\n\n"
	assert.Equal(t, expected, buffer.String())
}

func TestPrintCommandHelp(t *testing.T) {
	buffer := &bytes.Buffer{}

	c := &Context{
		Ngsi: &ngsilib.NGSI{StdWriter: buffer},
		Commands: []*Command{
			{
				Name:     "applications",
				Usage:    "manage applications for Keyrock",
				Category: "Keyrock",
				Flags: []Flag{
					&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
				},
				Subcommands: []*Command{
					{Name: "list", Usage: "list applications"},
					{Name: "get", Usage: "get application"},
					{Name: "create", Usage: "create application"},
					{Name: "update", Usage: "update application"},
					{Name: "delete", Usage: "delete application"},
					{
						Name:     "roles",
						Usage:    "manage roles for Keyrock",
						Category: "sub-command",
						Subcommands: []*Command{
							{Name: "list", Usage: "list roles"},
							{Name: "get", Usage: "get a role"},
							{Name: "create", Usage: "create a role"},
							{Name: "update", Usage: "update a role"},
							{Name: "delete", Usage: "delete a role"},
							{Name: "permissions", Usage: "list permissions associated to a role"},
							{Name: "assign", Usage: "Assign a permission to a role"},
							{Name: "unassign", Usage: "delete a permission from a role"},
						},
					},
				},
			},
		},
		GlobalFlags: []Flag{
			&StringFlag{Name: "syslog", Usage: "specify logging `LEVEL` (off, err, info, debug)"},
			&StringFlag{Name: "stderr", Usage: "specify logging `LEVEL` (off, err, info, debug)"},
			&StringFlag{Name: "config", Usage: "specify configuration `FILE`"},
			&StringFlag{Name: "cache", Usage: "specify cache `FILE`"},
			&BoolFlag{Name: "help", Usage: "show help"},
			&Int64Flag{Name: "margin", Usage: "I/O time out (second)", Hidden: true, Value: 180},
			&Int64Flag{Name: "timeout", Usage: "I/O time out (second)", Hidden: true, Value: 60},
			&Int64Flag{Name: "maxCount", Usage: "maxCount", Hidden: true, Value: 100},
			&BoolFlag{Name: "batch", Aliases: []string{"B"}, Usage: "don't use previous args (batch)"},
			&StringFlag{Name: "cmdName", Hidden: true},
			&BoolFlag{Name: "insecureSkipVerify", Usage: "TLS/SSL skip certificate verification"},
		},
		App: &App{
			Version: "0.9.0 (git_hash:bfd1ec240a8a8421929e2923f8fb5d3f6cab18ab",
			Usage:   "command-line tool for FIWARE NGSI and NGSI-LD",
		},
	}

	printCommandHelp(c)

	expected := "NAME:\n   ngsi applications - manage applications for Keyrock\n\nUSAGE:\n    [global options] applications [options] [command] [arguments...]\n\nCATEGORY:\n   Keyrock\n\nCOMMANDS:\n   list     list applications\n   get      get application\n   create   create application\n   update   update application\n   delete   delete application\n   help, h  Shows a list of commands or help for one command\n   sub-command:\n     roles  manage roles for Keyrock\n\nGLOBAL OPTIONS:\n   --syslog LEVEL        specify logging LEVEL (off, err, info, debug)\n   --stderr LEVEL        specify logging LEVEL (off, err, info, debug)\n   --config FILE         specify configuration FILE\n   --cache FILE          specify cache FILE\n   --help                show help (default: false)\n   --batch, -B           don't use previous args (batch) (default: false)\n   --insecureSkipVerify  TLS/SSL skip certificate verification (default: false)\n\n"
	assert.Equal(t, expected, buffer.String())
}

func TestCommandList(t *testing.T) {
	cmd := []*Command{
		{Name: "list", Usage: "list applications"},
		{Name: "get", Usage: "get application"},
		{Name: "create", Usage: "create application"},
		{Name: "update", Usage: "update application"},
		{Name: "delete", Usage: "delete application"},
	}

	c := &Context{App: &App{Commands: cmd}}

	actual := commandList(c)
	expected := "COMMANDS:\n   help, h  Shows a list of commands or help for one command\n   :\n     list    list applications\n     get     get application\n     create  create application\n     update  update application\n     delete  delete application\n\n"

	assert.Equal(t, expected, actual)
}
func TestSubCommandList(t *testing.T) {
	cmd := &Command{
		Subcommands: []*Command{
			{Name: "list", Usage: "list roles"},
			{Name: "get", Usage: "get a role"},
			{Name: "create", Usage: "create a role"},
			{Name: "update", Usage: "update a role"},
			{Name: "delete", Usage: "delete a role"},
			{Name: "permissions", Usage: "list permissions associated to a role"},
			{Name: "assign", Usage: "Assign a permission to a role"},
			{Name: "unassign", Usage: "delete a permission from a role"},
		},
	}

	actual := subCommandList(cmd)
	expected := "COMMANDS:\n   list         list roles\n   get          get a role\n   create       create a role\n   update       update a role\n   delete       delete a role\n   permissions  list permissions associated to a role\n   assign       Assign a permission to a role\n   unassign     delete a permission from a role\n   help, h      Shows a list of commands or help for one command\n\n"

	assert.Equal(t, expected, actual)
}

func TestSubCommandListEmpty(t *testing.T) {
	cmd := &Command{}

	actual := subCommandList(cmd)
	expected := ""

	assert.Equal(t, expected, actual)
}

func TestCommandFlags(t *testing.T) {
	hostRFlag := &StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true}
	tokenFlag := &StringFlag{Name: "token", Usage: "oauth token `VALUE`"}
	tenantFlag := &StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"}
	scopeFlag := &StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"}
	f := []Flag{
		hostRFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	}

	s := commandFlags("OPTIONS", f)

	expected := "OPTIONS:\n   --host VALUE, -h VALUE     broker or server host VALUE (required)\n   --token VALUE              oauth token VALUE\n   --service VALUE, -s VALUE  FIWARE Service VALUE\n   --path VALUE, -p VALUE     FIWARE ServicePath VALUE\n\n"
	assert.Equal(t, expected, s)

}

func TestCommandFlagsEmpty(t *testing.T) {
	actual := commandFlags("OPTIONS", nil)

	assert.Equal(t, "", actual)
}

func TestPreviousArg(t *testing.T) {
	cases := []struct {
		name     string
		msg      string
		len      int
		expected string
	}{
		{name: "test", msg: "message", len: 5, expected: "   test   message\n"},
		{name: "test", msg: "message", len: 10, expected: "   test        message\n"},
		{name: "test", msg: "message", len: 11, expected: "   test         message\n"},
		{name: "test", msg: "", len: 11, expected: ""},
	}

	for _, c := range cases {
		actual := previousArg(c.name, c.msg, c.len)
		assert.Equal(t, c.expected, actual)
	}
}

func TestPreviousArgsPreviousArgsOff(t *testing.T) {
	p := &ngsilib.Settings{UsePreviousArgs: false}
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: p}}

	actual := previousArgs(c)
	expected := "PREVIOUS ARGS:\n   off\n   (To enable it, run 'ngsi settings previousArgs --on')\n"

	assert.Equal(t, expected, actual)
}

func TestPreviousArgsPreviousArgsOn(t *testing.T) {
	p := &ngsilib.Settings{UsePreviousArgs: true, Host: "orion", Tenant: "iot", Scope: "/smartcity"}
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: p}}

	actual := previousArgs(c)
	expected := "PREVIOUS ARGS:\n   Host                orion\n   FIWARE-Service      iot\n   FIWARE-ServicePath  /smartcity\n   (To clear args, run 'ngsi settings clear')\n"

	assert.Equal(t, expected, actual)
}

func TestPreviousArgsPreviousArgsOnEmpty(t *testing.T) {
	p := &ngsilib.Settings{UsePreviousArgs: true}
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: p}}

	actual := previousArgs(c)
	expected := "PREVIOUS ARGS:\n   None\n"

	assert.Equal(t, expected, actual)
}

func TestPreviousArgsPreviousArgsNull(t *testing.T) {
	c := &Context{Ngsi: &ngsilib.NGSI{}}

	actual := previousArgs(c)
	expected := ""

	assert.Equal(t, expected, actual)
}

func TestPreviousArgsNGSINull(t *testing.T) {
	c := &Context{}

	actual := previousArgs(c)

	assert.Equal(t, "", actual)
}

func TestPreviousArgsMaxLen(t *testing.T) {
	p := &ngsilib.Settings{Host: "orion"}
	max := previousArgsMaxLen(p)
	assert.Equal(t, 4, max)

	p = &ngsilib.Settings{Host: "orion", Tenant: "iot"}
	max = previousArgsMaxLen(p)
	assert.Equal(t, 14, max)

	p = &ngsilib.Settings{Host: "orion", Tenant: "iot", Scope: "/smartcity"}
	max = previousArgsMaxLen(p)
	assert.Equal(t, 18, max)
}

func TestPreviousArgsLen(t *testing.T) {
	cases := []struct {
		name     string
		value    string
		expected int
	}{
		{name: "host", value: "orion", expected: 4},
		{name: "host", value: "", expected: 0},
		{name: "", value: "orion", expected: 0},
		{name: "", value: "", expected: 0},
	}

	for _, c := range cases {
		actual := previousArgsLen(c.name, c.value)
		assert.Equal(t, c.expected, actual)
	}
}

func TestMaxInt(t *testing.T) {

	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 10, b: 50, expected: 50},
		{a: 50, b: 10, expected: 50},
		{a: 0, b: 100, expected: 100},
		{a: 100, b: 0, expected: 100},
		{a: -1, b: 100, expected: 100},
		{a: 100, b: -1, expected: 100},
		{a: -10, b: -20, expected: -10},
		{a: -20, b: -10, expected: -10},
	}

	for _, c := range cases {
		actual := maxInt(c.a, c.b)
		assert.Equal(t, c.expected, actual)
	}

}

func TestCategoriesLen(t *testing.T) {
	categories := categories{"abc", "def", "aaa", "MANAGEMENT", "zzz"}

	actual := categories.Len()
	expected := 5

	assert.Equal(t, expected, actual)
}

func TestCategoriesLess(t *testing.T) {
	categories := categories{"abc", "def", "aaa", "MANAGEMENT", "zzz"}

	cases := []struct {
		i        int
		j        int
		expected bool
	}{
		{i: 0, j: 1, expected: true},
		{i: 1, j: 2, expected: false},
		{i: 2, j: 3, expected: true},
		{i: 3, j: 4, expected: false},
	}

	for _, c := range cases {
		actual := categories.Less(c.i, c.j)
		assert.Equal(t, c.expected, actual)
	}
}
func TestCategoriesSwap(t *testing.T) {
	categories := categories{"abc", "def", "aaa", "MANAGEMENT", "zzz"}

	categories.Swap(0, 4)

	assert.Equal(t, "zzz", categories[0])
	assert.Equal(t, "abc", categories[4])
}
