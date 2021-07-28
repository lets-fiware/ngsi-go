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

package helper

import (
	"bytes"
	"errors"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
)

var testApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "wirecloud command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&testCmd,
	},
}

var testCmd = ngsicli.Command{
	Name:     "preferences",
	Usage:    "manage preferences for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		&ngsicli.StringFlag{Name: "host2"},
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "get",
			Usage:      "get preferences",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return nil
			},
		},
	},
}

func TestSetupTestInitNGSI(t *testing.T) {
	ngsi := SetupTestInitNGSI()

	assert.NotEqual(t, (*ngsilib.NGSI)(nil), ngsi)
}

func TestSetupTestInitCmd(t *testing.T) {
	c := SetupTestInitCmd(nil)

	assert.NotEqual(t, (*ngsicli.Context)(nil), c)
}

func callBackForInitCmd(ngsi *ngsilib.NGSI) {
	ngsi.ConfigFile = &MockIoLib{HomeDir: errors.New("open error")}
}

func TestSetupTestInitCmdError(t *testing.T) {
	c := SetupTestInitCmd(callBackForInitCmd)

	assert.Equal(t, (*ngsicli.Context)(nil), c)
}

func TestSetupTest(t *testing.T) {
	args := []string{"preferences", "get", "--host", "wirecloud", "--host2", "wirecloud"}

	c := SetupTest(testApp, args)

	assert.NotEqual(t, (*ngsicli.Context)(nil), c)
}

func TestSetupTestWithConfig(t *testing.T) {
	args := []string{"preferences", "get", "--host", "wirecloud", "--host2", "wirecloud"}

	c := SetupTestWithConfig(testApp, args, "")

	assert.NotEqual(t, (*ngsicli.Context)(nil), c)
}

func TestSetupTestWithConfigAndCache(t *testing.T) {
	args := []string{"preferences", "get", "--host", "wirecloud", "--host2", "wirecloud"}

	c := SetupTestWithConfigAndCache(testApp, args, "", "")

	assert.NotEqual(t, (*ngsicli.Context)(nil), c)
}

func TestSetupTestWithConfigAndCacheCache(t *testing.T) {
	args := []string{"preferences", "get", "--host", "wirecloud"}

	c := SetupTestWithConfigAndCache(testApp, args, "", "cache")

	assert.NotEqual(t, (*ngsicli.Context)(nil), c)
}

func TestSetupTestWithConfigAndCacheError(t *testing.T) {
	args := []string{"preferences", "--host", "wirecloud", "--pretty"}

	c := SetupTestWithConfigAndCache(testApp, args, "", "")

	assert.Equal(t, (*ngsicli.Context)(nil), c)
}

func TestGetStdoutString(t *testing.T) {
	buf := bytes.NewBuffer([]byte("fiware"))
	c := &ngsicli.Context{Ngsi: &ngsilib.NGSI{StdWriter: buf}}

	s := GetStdoutString(c)

	assert.Equal(t, "fiware", s)

}

func TestStrPtr(t *testing.T) {
	s := StrPtr("fiware")

	assert.Equal(t, "fiware", *s)
}

func TestUrlParse(t *testing.T) {
	u := UrlParse("http://orion:1026/version")

	assert.Equal(t, "orion:1026", u.Host)
	assert.Equal(t, "/version", u.Path)
}
