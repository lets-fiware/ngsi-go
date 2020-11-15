/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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

	"github.com/stretchr/testify/assert"
)

func TestNGSICommand(t *testing.T) {
	cases := []struct {
		args []string
		rc   int
	}{
		//{args: []string{}, rc: 0},
		{args: []string{"cp", "--type", "abc", "--destination", "abc"}, rc: 1},
		{args: []string{"wc", "entities"}, rc: 1},
		{args: []string{"wc", "subscriptions"}, rc: 1},
		{args: []string{"wc", "registrations"}, rc: 1},
		{args: []string{"wc", "types"}, rc: 1},
		{args: []string{"ls"}, rc: 1},
		{args: []string{"rm", "--type", "abc"}, rc: 1},
		{args: []string{"template", "registration"}, rc: 1},
		{args: []string{"template", "subscription", "--url", "abc"}, rc: 1},
		{args: []string{"version"}, rc: 1},
		{args: []string{"man"}, rc: 1},
		{args: []string{"broker", "add"}, rc: 1},
		{args: []string{"broker", "delete"}, rc: 1},
		{args: []string{"broker", "get"}, rc: 1},
		{args: []string{"broker", "list"}, rc: 1},
		{args: []string{"broker", "update"}, rc: 1},
		{args: []string{"context", "add", "--name", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"context", "delete", "--name", "abc"}, rc: 1},
		{args: []string{"context", "list"}, rc: 1},
		{args: []string{"context", "update", "--name", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"settings", "list"}, rc: 1},
		{args: []string{"settings", "clear"}, rc: 1},
		{args: []string{"settings", "delete"}, rc: 1},
		{args: []string{"append", "attrs", "--id", "abc"}, rc: 1},
		{args: []string{"create", "entities"}, rc: 1},
		{args: []string{"create", "entity"}, rc: 1},
		{args: []string{"create", "registration"}, rc: 1},
		{args: []string{"create", "subscription", "--url", "abc"}, rc: 1},
		{args: []string{"delete", "entities"}, rc: 1},
		{args: []string{"delete", "entity", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "attr", "--id", "abc", "--attrName", "abc"}, rc: 1},
		{args: []string{"delete", "registration", "--id", "abc"}, rc: 1},
		{args: []string{"delete", "subscription", "--id", "abc"}, rc: 1},
		{args: []string{"get", "attr", "--id", "abc", "--attrName", "abc"}, rc: 1},
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
		{args: []string{"update", "attr", "--id", "abc", "--attrName", "abc"}, rc: 1},
		{args: []string{"update", "subscription", "--id", "abc", "--url", "abc"}, rc: 1},
		{args: []string{"upsert", "entities"}, rc: 1},
		{args: []string{"token"}, rc: 1},
		{args: []string{"debug"}, rc: 1},
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
