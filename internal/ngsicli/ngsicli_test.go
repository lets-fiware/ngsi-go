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
	"errors"
	"strings"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestRun(t *testing.T) {
	setInitNGSI(nil)

	f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return nil
	}
	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
		Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			return f(c, ngsi, client)
		},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "version", "--host", "orion"}

	err := r.Run(args)

	assert.NoError(t, err)

}

func TestRunSavePreviousArgs(t *testing.T) {
	ngsi := setInitNGSI(nil)
	ngsi.Updated = true

	f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return nil
	}
	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
		Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			ngsi.PreviousArgs.UsePreviousArgs = true
			return f(c, ngsi, client)
		},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "version", "--host", "orion"}

	err := r.Run(args)

	assert.NoError(t, err)
}

func TestRunHelp(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(nil)
	ngsi.StdWriter = buf

	f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return nil
	}
	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
		Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			return f(c, ngsi, client)
		},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi"}

	err := r.Run(args)

	if assert.NoError(t, err) {
		expected := "NAME:\n   ngsi - \n\nUSAGE:\n   ngsi [global options] command [options] [arguments...]\n\nVERSION:\n   ngsi version \n\nCOMMANDS:\n   help, h  Shows a list of commands or help for one command\n   :\n     fiware   \n     version  \n\nGLOBAL OPTIONS:\n   --help         show help (default: false)\n   --version, -v  print the version (default: false)\n\nPREVIOUS ARGS:\n   off\n   (To enable it, run 'ngsi settings previousArgs --on')\n"
		assert.Equal(t, expected, buf.String())
	}

}

func TestRunHelp2(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(nil)
	ngsi.StdWriter = buf

	f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return nil
	}
	create := &Command{
		Name:  "create",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
		Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			return f(c, ngsi, client)
		},
		Subcommands: []*Command{
			{Name: "entity", Subcommands: []*Command{}},
		},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			create,
		},
	}
	args := []string{"ngsi", "create", "entity"}

	err := r.Run(args)

	if assert.NoError(t, err) {
		expected := "NAME:\n   ngsi create entity - \n\nUSAGE:\n   ngsi [global options] create entity [options] [command] [arguments...]\n\nCATEGORY:\n   \n\nCOMMANDS:\n%!(EXTRA string=help, h, string=Shows a list of commands or help for one command)\nOPTIONS:\n   --host VALUE  \n   --help        show help (default: false)\n\nGLOBAL OPTIONS:\n\nPREVIOUS ARGS:\n   off\n   (To enable it, run 'ngsi settings previousArgs --on')\n"
		assert.Equal(t, expected, buf.String())
	}

}

func TestRunErrorBashCompletion(t *testing.T) {
	setInitNGSI(nil)

	cases := []struct {
		args []string
	}{
		{args: []string{"ngsi", "creat", "--generate-bash-completion"}},
		{args: []string{"ngsi", "--stderr", "--generate-bash-completion"}},
	}

	for _, c := range cases {
		f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			return nil
		}
		version := &Command{
			Name:  "version",
			Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
			Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return f(c, ngsi, client)
			},
		}
		r := &App{
			Commands: []*Command{
				{Name: "fiware"},
				version,
			},
		}

		err := r.Run(c.args)

		assert.NoError(t, err)
	}
}

func TestRunError(t *testing.T) {
	setInitNGSI(nil)

	f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return nil
	}
	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
		Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			return f(c, ngsi, client)
		},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "list", "--host", "orion", "entity"}

	err := r.Run(args)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "list not found", ngsiErr.Message)
	}
}

func TestRunSavePreviousArgsError(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(buf)
	ngsi.Updated = true
	ngsi.ConfigFile = &MockIoLib{Trunc: []error{nil, errors.New("trunc error")}}

	f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return nil
	}
	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
		Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			ngsi.PreviousArgs.UsePreviousArgs = true
			return f(c, ngsi, client)
		},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "version", "--host", "orion"}

	err := r.Run(args)

	if assert.NoError(t, err) {
		assert.Equal(t, "Run002 trunc error", buf.String())
	}
}

func TestParse(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "version", "--host", "orion"}

	_, _, err := r.Parse(args)

	assert.NoError(t, err)
}

func TestParseError(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{}

	_, _, err := r.Parse(args)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "command name error", ngsiErr.Message)
	}
}

func TestNgsiRun(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "version", "--host", "orion"}

	_, _, err := ngsiRun(r, args)

	assert.NoError(t, err)
}

func TestNgsiRunCmdHelp(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(nil)
	ngsi.StdWriter = buf

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "version", "--help"}

	cmd, c, err := ngsiRun(r, args)

	if assert.NoError(t, err) {
		assert.NotEqual(t, nil, cmd)
		assert.NotEqual(t, nil, c)
		expected := "NAME:\n   ngsi version - \n\nUSAGE:\n   ngsi [global options] version [options] [arguments...]\n\nCATEGORY:\n   \n\nOPTIONS:\n   --host VALUE  \n   --help        show help (default: true)\n\nGLOBAL OPTIONS:\n\nPREVIOUS ARGS:\n   off\n   (To enable it, run 'ngsi settings previousArgs --on')\n"
		assert.Equal(t, expected, buf.String())
	}
}

func TestNgsiRunPrintVersion(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(nil)
	ngsi.StdWriter = buf

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "--version"}

	cmd, c, err := ngsiRun(r, args)

	if assert.NoError(t, err) {
		assert.NotEqual(t, nil, cmd)
		assert.NotEqual(t, nil, c)
		expected := "ngsi version \n"
		assert.Equal(t, expected, buf.String())
	}
}

func TestNgsiRunPrintVersionShortHand(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(nil)
	ngsi.StdWriter = buf

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "-v"}

	cmd, c, err := ngsiRun(r, args)

	if assert.NoError(t, err) {
		assert.NotEqual(t, nil, cmd)
		assert.NotEqual(t, nil, c)
		expected := "ngsi version \n"
		assert.Equal(t, expected, buf.String())
	}
}

func TestNgsiRunPrintHelp(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(nil)
	ngsi.StdWriter = buf

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi"}

	cmd, c, err := ngsiRun(r, args)

	if assert.NoError(t, err) {
		assert.NotEqual(t, nil, cmd)
		assert.NotEqual(t, nil, c)
		expected := "NAME:\n   ngsi - \n\nUSAGE:\n   ngsi [global options] command [options] [arguments...]\n\nVERSION:\n   ngsi version \n\nCOMMANDS:\n   help, h  Shows a list of commands or help for one command\n   :\n     fiware   \n     version  \n\nGLOBAL OPTIONS:\n   --help         show help (default: false)\n   --version, -v  print the version (default: false)\n\nPREVIOUS ARGS:\n   off\n   (To enable it, run 'ngsi settings previousArgs --on')\n"
		assert.Equal(t, expected, buf.String())
	}
}

func TestNgsiRunErrorCmdName(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}

	_, _, err := ngsiRun(r, nil)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "command name error", ngsiErr.Message)
	}
}

func TestNgsiRunErrorGlobalFlags(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "--host"}

	_, _, err := ngsiRun(r, args)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unknown flag: --host", ngsiErr.Message)
	}
}

func TestNgsiRunErrorInitCmd(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Flags: []Flag{
			&StringFlag{Name: "stderr"},
		},
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "--stderr", "unknown", "version", "--host", "orion"}

	_, _, err := ngsiRun(r, args)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "stderr logLevel error", ngsiErr.Message)
	}
}

func TestNgsiRunErrorRunCmd(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}},
	}
	r := &App{
		Flags: []Flag{
			&StringFlag{Name: "stderr"},
		},
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "--stderr", "err", "unknown", "--host", "orion"}

	_, _, err := ngsiRun(r, args)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unknown not found", ngsiErr.Message)
	}
}

func TestNgsiRunErrorCreateNewClient(t *testing.T) {
	setInitNGSI(nil)

	version := &Command{
		Name:  "version",
		Flags: []Flag{&StringFlag{Name: "host", Value: "orion", InitClient: true}},
	}
	r := &App{
		Flags: []Flag{
			&StringFlag{Name: "stderr"},
		},
		Commands: []*Command{
			{Name: "fiware"},
			version,
		},
	}
	args := []string{"ngsi", "--stderr", "err", "version", "--host", "orion-ld"}

	_, _, err := ngsiRun(r, args)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "error host: orion-ld", ngsiErr.Message)
	}
}

func TestRunCmd(t *testing.T) {
	buf := &bytes.Buffer{}
	ngsi := setInitNGSI(nil)
	ngsi.StdWriter = buf

	cmd := &Command{Name: "ngsi", Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}}}
	c := &Context{
		App: &App{
			Commands: []*Command{
				{Name: "fiware"},
				cmd,
			},
		},
		Ngsi: ngsi,
	}
	token := newToken([]string{"--host", "orion"})

	actual, err := runCmd(c, token, "ngsi")

	if assert.NoError(t, err) {
		assert.Equal(t, cmd, actual)
	}
}

func TestRunCmdHelp(t *testing.T) {
	buf := &bytes.Buffer{}
	cmd := &Command{Name: "ngsi", Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}}}
	c := &Context{
		App: &App{
			Commands: []*Command{
				{Name: "fiware"},
				cmd,
			},
		},
		Ngsi: &ngsilib.NGSI{
			StdWriter: buf,
		},
	}
	token := newToken([]string{"--help"})

	actual, err := runCmd(c, token, "ngsi")

	if assert.NoError(t, err) {
		assert.Equal(t, (*Command)(nil), actual)
		expected := "NAME:\n   ngsi ngsi - \n\nUSAGE:\n    [global options] ngsi [options] [arguments...]\n\nCATEGORY:\n   \n\nOPTIONS:\n   --host VALUE  \n   --help        show help (default: true)\n\n"
		assert.Equal(t, expected, buf.String())
	}
}

func TestRunCmdErrorUnkownFlag(t *testing.T) {
	setInitNGSI(nil)

	cmd := &Command{Name: "ngsi", Flags: []Flag{&StringFlag{Name: "host", Value: "orion"}}}
	c := &Context{
		App: &App{
			Commands: []*Command{
				{Name: "fiware"},
				cmd,
			},
		},
	}
	token := newToken([]string{"--host2", "orion"})

	_, err := runCmd(c, token, "ngsi")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown flag: --host2", ngsiErr.Message)
	}
}

func TestRunCmdErrorCmdNotFound(t *testing.T) {
	setInitNGSI(nil)

	c := &Context{App: &App{}}

	_, err := runCmd(c, nil, "ngsi")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "ngsi not found", ngsiErr.Message)
	}
}

func TestCreateNewClientNoHost(t *testing.T) {
	c := &Context{Flags: []Flag{}}

	err := createNewClient(c)

	assert.NoError(t, err)
}

func TestCreateNewClientHost(t *testing.T) {
	buf := &bytes.Buffer{}
	host := &StringFlag{Name: "host", Value: "orion", Set: true, Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true}
	c := &Context{
		Flags: []Flag{
			host,
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Ngsi: &ngsilib.NGSI{
			Stderr:       buf,
			PreviousArgs: &ngsilib.Settings{},
			ServerList:   ngsilib.ServerList{"orion": &ngsilib.Server{ServerHost: "http://orion"}},
		},
	}

	err := createNewClient(c)

	assert.NoError(t, err)
}

func TestCreateNewClientHostPreviousHost(t *testing.T) {
	buf := &bytes.Buffer{}
	host := &StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true}
	c := &Context{
		Flags: []Flag{
			host,
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Ngsi: &ngsilib.NGSI{
			Stderr:       buf,
			PreviousArgs: &ngsilib.Settings{Host: "orion"},
			ServerList:   ngsilib.ServerList{"orion": &ngsilib.Server{ServerHost: "http://orion"}},
		},
	}

	err := createNewClient(c)

	assert.NoError(t, err)
}

func TestCreateNewClientHost2(t *testing.T) {
	buf := &bytes.Buffer{}
	host2 := &StringFlag{Name: "host2", Value: "orion", Set: true, Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: false, InitClient: true}
	c := &Context{
		Flags: []Flag{
			host2,
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Ngsi: &ngsilib.NGSI{
			Stderr:       buf,
			PreviousArgs: &ngsilib.Settings{},
			ServerList:   ngsilib.ServerList{"orion": &ngsilib.Server{ServerHost: "http://orion"}},
		},
	}

	err := createNewClient(c)

	assert.NoError(t, err)
}

func TestCreateNewClientHostNotFound(t *testing.T) {
	buf := &bytes.Buffer{}
	host2 := &StringFlag{Name: "host", Value: "orion-ld", Set: true, Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: false, InitClient: true}
	c := &Context{
		Flags: []Flag{
			host2,
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Ngsi: &ngsilib.NGSI{
			Stderr:       buf,
			PreviousArgs: &ngsilib.Settings{},
			ServerList:   ngsilib.ServerList{"orion": &ngsilib.Server{ServerHost: "http://orion"}},
		},
	}

	err := createNewClient(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error host: orion-ld", ngsiErr.Message)
	}
}

func TestCreateNewClientHost2NotFound(t *testing.T) {
	buf := &bytes.Buffer{}
	host2 := &StringFlag{Name: "host2", Value: "orion2", Set: true, Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: false, InitClient: true}
	c := &Context{
		Flags: []Flag{
			host2,
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Ngsi: &ngsilib.NGSI{
			Stderr:       buf,
			PreviousArgs: &ngsilib.Settings{},
			ServerList:   ngsilib.ServerList{"orion": &ngsilib.Server{ServerHost: "http://orion"}},
		},
	}

	err := createNewClient(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: orion2 (destination)", ngsiErr.Message)
	}
}

func TestParseCmdFlag(t *testing.T) {
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
	}
	token := newToken([]string{"--host", "orion"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.NoError(t, err) {
		assert.NotEqual(t, (*Command)(nil), acutal)
	}
}

func TestParseCmdFlagHelp(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{StdWriter: buf}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Subcommands: []*Command{},
	}
	token := newToken([]string{""})
	_ = token.Next()

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.NoError(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		assert.Equal(t, true, c.HelpCommand)
	}
}

func TestParseCmdFlagHelpSub(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
	}
	token := newToken([]string{"--help"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.NoError(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		assert.Equal(t, true, c.HelpCommand)
	}
}

func TestParseCmdFlagSub(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{StdWriter: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Subcommands: []*Command{
			{Name: "list", Usage: "list applications", Flags: []Flag{&StringFlag{Name: "-aid"}}},
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
	}
	token := newToken([]string{"list", "--host", "keyrock"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.NoError(t, err) {
		assert.NotEqual(t, (*Command)(nil), acutal)
		assert.Equal(t, false, c.HelpCommand)
	}
}

func TestParseCmdFlagArg(t *testing.T) {
	buf := &bytes.Buffer{}
	cmdArgs := cmdArgs{}
	c := &Context{Ngsi: &ngsilib.NGSI{StdWriter: buf, PreviousArgs: &ngsilib.Settings{}}, Arg: &cmdArgs}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
	}
	token := newToken([]string{"--host", "keyrock", "put"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.NoError(t, err) {
		assert.NotEqual(t, (*Command)(nil), acutal)
		assert.Equal(t, 1, cmdArgs.Len())
		assert.Equal(t, "put", cmdArgs.Get(0))
		assert.Equal(t, false, c.HelpCommand)
	}
}

func TestParseCmdFlagErrFlag(t *testing.T) {
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
	}
	token := newToken([]string{"--host"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.Error(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "value missing", ngsiErr.Message)
	}
}

func TestParseCmdFlagSubErrorSubcmdNotFound(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
		Subcommands: []*Command{
			{Name: "list", Usage: "list applications", Flags: []Flag{&StringFlag{Name: "-aid"}}},
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
	}
	token := newToken([]string{"put"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.Error(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "put not found", ngsiErr.Message)
	}
}

func TestParseCmdFlagErrorRequired(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
	}
	token := newToken([]string{"--oAuthToken", "1234567890"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.Error(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "missing required options", ngsiErr.Message)
		expected := "required002 --host not found\n"
		assert.Equal(t, expected, buf.String())
	}
}

func TestParseCmdFlagEmpty(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "token", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
	}
	token := newToken([]string{"--host", "orion", "--path", ""})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.Error(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "option values are empty", ngsiErr.Message)
		assert.Equal(t, "checkEmpty001 --path: value is empty\n", buf.String())
	}
}

func TestParseCmdFlagValidation(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "token", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
		},
		OptionFlags: &ValidationFlag{Mode: NonCondition},
	}
	token := newToken([]string{"--host", "orion"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.Error(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "validation mode error", err.Error())
	}
}

func TestParseCmdFlagCheckChoices(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "token", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "serverType", Value: "", Choices: []string{"keyrock", "cygnus", "comet"}, Set: true},
		},
	}
	token := newToken([]string{"--host", "orion", "--serverType", "wirecloud"})

	acutal, err := parseCmdFlag(c, token, cmds)

	if assert.Error(t, err) {
		assert.Equal(t, (*Command)(nil), acutal)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "checkChoices001 specify one of keyrock, cygnus and comet to --serverType\n", buf.String())
	}
}

func TestRequired(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "host", Value: "orion", Set: true, Required: true},
		&BoolFlag{Name: "pretty"},
		&StringFlag{Name: "data", Value: "{}", Set: true},
	}
	buf := &bytes.Buffer{}
	c := &Context{Flags: f, Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}

	actual := required(f, c)

	assert.Equal(t, false, actual)
	expected := ""
	assert.Equal(t, expected, buf.String())
}

func TestRequiredHostNotFound(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "host", Required: true},
		&BoolFlag{Name: "pretty"},
		&StringFlag{Name: "data"},
	}
	buf := &bytes.Buffer{}
	c := &Context{Flags: f, Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}

	actual := required(f, c)

	assert.Equal(t, true, actual)
	expected := "required002 --host not found\n"
	assert.Equal(t, expected, buf.String())
}

func TestRequiredDataEmpty(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "host", Value: "orion", Set: true, Required: true},
		&BoolFlag{Name: "pretty"},
		&StringFlag{Name: "data", Value: "", Set: true},
	}
	buf := &bytes.Buffer{}
	c := &Context{Flags: f, Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}

	actual := required(f, c)

	assert.Equal(t, true, actual)
	expected := "required003 data is empty\n"
	assert.Equal(t, expected, buf.String())
}

func TestRequiredRequiredFlagsSet(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "id", Value: "123", Set: true},
		&StringFlag{Name: "attr", Value: "speed", Set: true},
	}
	buf := &bytes.Buffer{}
	c := &Context{
		Flags:         f,
		Ngsi:          &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}},
		RequiredFlags: []string{"id", "attr"},
		Arg:           &cmdArgs{},
	}

	actual := required(f, c)

	assert.Equal(t, false, actual)
	assert.Equal(t, "123", f[0].(*StringFlag).Value)
	assert.Equal(t, "speed", f[1].(*StringFlag).Value)
}

func TestRequiredRequiredFlagsSetArg(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "id"},
		&StringFlag{Name: "attr"},
	}
	buf := &bytes.Buffer{}
	c := &Context{
		Flags:         f,
		Ngsi:          &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}},
		RequiredFlags: []string{"id", "attr"},
		Arg:           &cmdArgs{"123", "speed"},
	}

	actual := required(f, c)

	assert.Equal(t, false, actual)
	assert.Equal(t, "123", f[0].(*StringFlag).Value)
	assert.Equal(t, "speed", f[1].(*StringFlag).Value)
}

func TestRequiredRequiredFlagsError(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "id"},
		&Int64Flag{Name: "value"},
	}
	buf := &bytes.Buffer{}
	c := &Context{
		Flags:         f,
		Ngsi:          &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}},
		RequiredFlags: []string{"id", "value"},
		Arg:           &cmdArgs{"123", "speed"},
	}

	actual := required(f, c)

	assert.Equal(t, true, actual)
	expected := "--value: speed is not number\n"
	assert.Equal(t, expected, buf.String())
}

func TestRequiredRequiredFlagsErrorName(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "id"},
		&Int64Flag{Name: "value"},
	}
	buf := &bytes.Buffer{}
	c := &Context{
		Flags:         f,
		Ngsi:          &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}},
		RequiredFlags: []string{"id", "attr"},
		Arg:           &cmdArgs{"123", "speed"},
	}

	actual := required(f, c)

	assert.Equal(t, true, actual)
	expected := "required001 --attr: not found\n"
	assert.Equal(t, expected, buf.String())
}

func TestRequiredRequiredFlagsNoSet(t *testing.T) {
	f := []Flag{
		&StringFlag{Name: "id", Value: "123", Set: false, Required: true},
		&StringFlag{Name: "attr", Value: "speed", Set: false, Required: true},
	}
	buf := &bytes.Buffer{}
	c := &Context{
		Flags:         f,
		Ngsi:          &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}},
		RequiredFlags: []string{"id", "attr"},
		Arg:           &cmdArgs{},
	}

	actual := required(f, c)

	assert.Equal(t, true, actual)
	expected := "required002 --id not found\nrequired002 --attr not found\n"
	assert.Equal(t, expected, buf.String())
}

func TestCheckEmpty(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`"},
		},
	}

	acutal := checkEmpty(cmds.Flags, c)

	assert.Equal(t, false, acutal)
}

func TestCheckEmptyError(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "host", Usage: "broker or server host `VALUE`", Aliases: []string{"h"}, Required: true, InitClient: true},
			&StringFlag{Name: "oAuthToken", Usage: "oauth token `VALUE`"},
			&StringFlag{Name: "service", Aliases: []string{"s"}, Usage: "FIWARE Service `VALUE`"},
			&StringFlag{Name: "path", Aliases: []string{"p"}, Usage: "FIWARE ServicePath `VALUE`", Value: "", Set: true},
		},
	}

	acutal := checkEmpty(cmds.Flags, c)

	assert.Equal(t, true, acutal)
	assert.Equal(t, "checkEmpty001 --path: value is empty\n", buf.String())
}

func TestSetPrevArgsUnusePreviousArgs(t *testing.T) {
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: &ngsilib.Settings{UsePreviousArgs: false}}}

	setPrevArgs(c)
}

func TestCheckChoices(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "serverType", Value: "cygnus", Choices: []string{"keyrock", "cygnus", "comet"}, Set: true},
		},
	}

	acutal := checkChoices(cmds.Flags, c)

	assert.Equal(t, false, acutal)
	assert.Equal(t, "", buf.String())
}

func TestCheckChoicesErrorOneOfThem(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "serverType", Value: "", Choices: []string{"keyrock", "cygnus", "comet"}, Set: true},
		},
	}

	acutal := checkChoices(cmds.Flags, c)

	assert.Equal(t, true, acutal)
	assert.Equal(t, "checkChoices001 specify one of keyrock, cygnus and comet to --serverType\n", buf.String())
}

func TestCheckChoicesErrorEither(t *testing.T) {
	buf := &bytes.Buffer{}
	c := &Context{Ngsi: &ngsilib.NGSI{Stderr: buf, PreviousArgs: &ngsilib.Settings{}}}
	cmds := &Command{
		Flags: []Flag{
			&StringFlag{Name: "safeString", Value: "", Choices: []string{"off", "on"}, Set: true},
		},
	}

	acutal := checkChoices(cmds.Flags, c)

	assert.Equal(t, true, acutal)
	assert.Equal(t, "checkChoices001 specify either off or on to --safeString\n", buf.String())
}

func TestSetPrevArgsNoHost(t *testing.T) {
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: &ngsilib.Settings{UsePreviousArgs: true}}}

	setPrevArgs(c)
}

func TestSetPrevArgsPriousArgs(t *testing.T) {
	server := make(ngsilib.ServerList)
	server["orion"] = &ngsilib.Server{ServerHost: "http://oiron:1026", Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Value: "orion", Set: true, PreviousArgs: false},
			&StringFlag{Name: "oAuthToken"},
			&StringFlag{Name: "service"},
			&StringFlag{Name: "path"},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: &ngsilib.Settings{UsePreviousArgs: true},
			ServerList:   server,
		},
	}

	setPrevArgs(c)
}

func TestSetPrevArgsServerValue(t *testing.T) {
	server := make(ngsilib.ServerList)
	server["orion"] = &ngsilib.Server{ServerHost: "http://oiron:1026", Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Value: "orion", Set: true, PreviousArgs: true},
			&StringFlag{Name: "oAuthToken"},
			&StringFlag{Name: "service"},
			&StringFlag{Name: "path"},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: &ngsilib.Settings{UsePreviousArgs: true},
			ServerList:   server,
		},
	}

	setPrevArgs(c)

	assert.Equal(t, "orion", c.Flags[0].(*StringFlag).Value)
	assert.Equal(t, "", c.Flags[1].(*StringFlag).Value)
	assert.Equal(t, "openiot", c.Flags[2].(*StringFlag).Value)
	assert.Equal(t, "/", c.Flags[3].(*StringFlag).Value)
}

func TestSetPrevArgsFlagValue(t *testing.T) {
	server := make(ngsilib.ServerList)
	server["orion"] = &ngsilib.Server{Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Value: "orion", Set: true},
			&StringFlag{Name: "oAuthToken"},
			&StringFlag{Name: "service", Value: "fiware", Set: true},
			&StringFlag{Name: "path", Value: "/fiware", Set: true},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: &ngsilib.Settings{UsePreviousArgs: true},
			ServerList:   server,
		},
	}

	setPrevArgs(c)

	assert.Equal(t, "fiware", c.Flags[2].(*StringFlag).Value)
	assert.Equal(t, "/fiware", c.Flags[3].(*StringFlag).Value)
}

func TestSetPrevArgsHostEqualPrevArg(t *testing.T) {
	server := make(ngsilib.ServerList)
	server["orion"] = &ngsilib.Server{Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Value: "orion", Set: true, PreviousArgs: true},
			&StringFlag{Name: "oAuthToken"},
			&StringFlag{Name: "service", Value: "fiware", Set: true},
			&StringFlag{Name: "path", Value: "/fiware", Set: true},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: &ngsilib.Settings{UsePreviousArgs: true, Host: "orion"},
			ServerList:   server,
		},
	}

	setPrevArgs(c)

	assert.Equal(t, "orion", c.Flags[0].(*StringFlag).Value)
}

func TestSetPrevArgsHostNotEqualPrevArg(t *testing.T) {
	server := make(ngsilib.ServerList)
	server["orion"] = &ngsilib.Server{Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Value: "orion", Set: true, PreviousArgs: true},
			&StringFlag{Name: "oAuthToken"},
			&StringFlag{Name: "service", Value: "fiware", Set: true},
			&StringFlag{Name: "path", Value: "/fiware", Set: true},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: &ngsilib.Settings{UsePreviousArgs: true, Host: "orion-ld"},
			ServerList:   server,
		},
	}

	setPrevArgs(c)

	assert.Equal(t, "orion", c.Ngsi.PreviousArgs.Host)
}

func TestGetPrevFlags(t *testing.T) {
	c := &Context{Ngsi: &ngsilib.NGSI{PreviousArgs: &ngsilib.Settings{UsePreviousArgs: true, Host: "orion-ld"}}}

	actual := getPrevFlags(c)

	assert.NotEqual(t, (*prevFlags)(nil), actual)
}

func TestCopyPrevArgsToFlags(t *testing.T) {
	p := &ngsilib.Settings{UsePreviousArgs: true, Host: "orion-ld", Token: "token", Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Value: "orion", Set: true, PreviousArgs: true},
			&StringFlag{Name: "oAuthToken"},
			&StringFlag{Name: "service", Value: "fiware", Set: true},
			&StringFlag{Name: "path"},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: p,
		},
	}

	copyPrevArgsToFlags(c)

	assert.Equal(t, "orion", c.Flags[0].(*StringFlag).Value)
	assert.Equal(t, "token", c.Flags[1].(*StringFlag).Value)
	assert.Equal(t, "fiware", c.Flags[2].(*StringFlag).Value)
	assert.Equal(t, "/", c.Flags[3].(*StringFlag).Value)
}

func TestCopyPrevArgsToFlagsHost(t *testing.T) {
	p := &ngsilib.Settings{UsePreviousArgs: true, Host: "orion-ld", Token: "token", Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Set: false, PreviousArgs: true},
			&StringFlag{Name: "oAuthToken"},
			&StringFlag{Name: "service", Value: "fiware", Set: true},
			&StringFlag{Name: "path"},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: p,
		},
	}

	copyPrevArgsToFlags(c)

	assert.Equal(t, "orion-ld", c.Flags[0].(*StringFlag).Value)
	assert.Equal(t, "token", c.Flags[1].(*StringFlag).Value)
	assert.Equal(t, "fiware", c.Flags[2].(*StringFlag).Value)
	assert.Equal(t, "/", c.Flags[3].(*StringFlag).Value)
}

func TestCopyFlagsToPrevArgs(t *testing.T) {
	p := &ngsilib.Settings{UsePreviousArgs: true, Host: "orion-ld", Token: "token", Tenant: "openiot", Scope: "/"}
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "host", Value: "orion", Set: true, PreviousArgs: true},
			&StringFlag{Name: "token"},
			&StringFlag{Name: "service", Value: "fiware", Set: true},
			&StringFlag{Name: "path", Value: "/fiware", Set: true},
		},
		Ngsi: &ngsilib.NGSI{
			PreviousArgs: p,
		},
	}

	copyFlagsToPrevArgs(c)

	assert.Equal(t, true, c.Ngsi.Updated)
	assert.Equal(t, "orion", p.Host)
	assert.Equal(t, "", p.Token)
	assert.Equal(t, "fiware", p.Tenant)
	assert.Equal(t, "/fiware", p.Scope)
}

func TestSearchSubCommnad(t *testing.T) {
	cmds := []*Command{
		{Name: "create"},
		{Name: "get"},
		{Name: "update"},
		{Name: "delete"},
	}
	actual := searchSubCommnad(cmds, "update")

	assert.NotEqual(t, nil, actual)
}

func TestSearchSubCommnadNotFound(t *testing.T) {
	cmds := []*Command{
		{Name: "create"},
		{Name: "get"},
		{Name: "update"},
		{Name: "delete"},
	}
	actual := searchSubCommnad(cmds, "list")

	assert.Equal(t, (*Command)(nil), actual)
}

func TestParseGlobalFlag(t *testing.T) {
	c := &Context{App: &App{Flags: []Flag{
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
	}}}

	token := newToken([]string{"--syslog", "info"})

	err := parseGlobalFlag(c, token)

	assert.NoError(t, err)

}

func TestParseGlobalFlagNoFlag(t *testing.T) {
	c := &Context{App: &App{Flags: []Flag{
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
	}}}

	token := newToken([]string{""})

	err := parseGlobalFlag(c, token)

	assert.NoError(t, err)

}

func TestParseGlobalFlagError(t *testing.T) {
	c := &Context{App: &App{Flags: []Flag{
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
	}}}

	token := newToken([]string{"--syslog"})

	err := parseGlobalFlag(c, token)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "value missing", ngsiErr.Message)
	}
}

func TestIsOption(t *testing.T) {
	cases := []struct {
		args     string
		name     string
		alias    string
		expected bool
	}{
		{args: "--host", name: "host", alias: "", expected: true},
		{args: "-h", name: "", alias: "h", expected: true},
		{args: "host", name: "host", alias: "", expected: false},
		{args: "h", name: "h", alias: "", expected: false},
		{args: "--", name: "--", alias: "", expected: false},
		{args: "-", name: "-", alias: "", expected: false},
	}

	for _, c := range cases {
		name, alias, err := isOption(c.args)

		assert.Equal(t, c.name, name)
		assert.Equal(t, c.alias, alias)
		assert.Equal(t, c.expected, err)
	}

}

func TestParseOptString(t *testing.T) {
	f := []Flag{&StringFlag{Name: "host"}}
	token := newToken([]string{"orion"})

	err := parseOpt(f, token, "host", "")

	assert.NoError(t, err)
}

func TestParseOptInt64(t *testing.T) {
	f := []Flag{&Int64Flag{Name: "int64"}}
	token := newToken([]string{"123"})

	err := parseOpt(f, token, "int64", "")

	assert.NoError(t, err)
}

func TestParseOptBool(t *testing.T) {
	f := []Flag{&BoolFlag{Name: "bool"}}

	cases := []struct {
		value string
	}{
		{value: "true"},
		{value: "True"},
		{value: "TRUE"},
		{value: "false"},
		{value: "False"},
		{value: "FALSE"},
		{value: "on"},
		{value: "ON"},
		{value: "off"},
		{value: "OFF"},
		{value: "1"},
	}

	for _, c := range cases {
		token := newToken([]string{c.value})
		err := parseOpt(f, token, "bool", "")

		assert.NoError(t, err)
	}
}

func TestParseOptBoolNul(t *testing.T) {
	f := []Flag{&BoolFlag{Name: "bool"}}

	token := newToken([]string{""})
	_ = token.Next()

	err := parseOpt(f, token, "bool", "")

	assert.NoError(t, err)
}

func TestParseOptBoolErrorStringValueMissing(t *testing.T) {
	f := []Flag{&StringFlag{Name: "string"}}

	token := newToken([]string{""})
	_ = token.Next()

	err := parseOpt(f, token, "string", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "value missing", ngsiErr.Message)
	}
}

func TestParseOptBoolErrorInt64ValueMissing(t *testing.T) {
	f := []Flag{&Int64Flag{Name: "int64"}}

	token := newToken([]string{""})
	_ = token.Next()

	err := parseOpt(f, token, "int64", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "value missing", ngsiErr.Message)
	}
}

func TestParseOptBoolErrorInt64Value(t *testing.T) {
	f := []Flag{&Int64Flag{Name: "int64"}}

	token := newToken([]string{"abc"})

	err := parseOpt(f, token, "int64", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "abc is not number", ngsiErr.Message)
	}
}

func TestParseOptBoolErrorUnknown(t *testing.T) {
	f := []Flag{&StringFlag{Name: "string"}}

	token := newToken([]string{""})
	_ = token.Next()

	err := parseOpt(f, token, "unknown", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unknown flag: --unknown", ngsiErr.Message)
	}
}

func TestParseOptBoolErrorUnknownAliase(t *testing.T) {
	f := []Flag{&StringFlag{Name: "string"}}

	token := newToken([]string{""})
	_ = token.Next()

	err := parseOpt(f, token, "", "u")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "unknown flag: -u", ngsiErr.Message)
	}
}

func TestHintCmdList(t *testing.T) {
	cases := []struct {
		cmds     []*Command
		expected string
	}{
		{cmds: nil, expected: ""},
		{cmds: []*Command{{Name: "create"}, {Name: "delete"}}, expected: "create\ndelete\n"},
		{cmds: []*Command{{Name: "create"}, {Name: "update"}, {Name: "delete"}}, expected: "create\nupdate\ndelete\n"},
	}

	for _, c := range cases {
		actual := hintCmdList(c.cmds)

		assert.Equal(t, c.expected, actual)
	}
}

func TestHintFlagList(t *testing.T) {
	cases := []struct {
		arg      string
		expected string
	}{
		{arg: "-s", expected: ""},
		{arg: "", expected: "--service\n--path\n-s\n-p\n"},
		{arg: "--host", expected: ""},
		{arg: "--", expected: "--service\n--path\n"},
		{arg: "--service", expected: ""},
		{arg: "-h", expected: ""},
		{arg: "-s", expected: ""},
		{arg: "-", expected: "--service\n--path\n-s\n-p\n"},
	}

	for _, c := range cases {
		flags := []Flag{
			&StringFlag{Name: "host", Aliases: []string{"h"}, Value: "orion", Set: true, PreviousArgs: true},
			&StringFlag{Name: "oAuthToken", Hidden: true},
			&StringFlag{Name: "service", Aliases: []string{"s"}},
			&StringFlag{Name: "path", Aliases: []string{"p"}},
		}

		actual := hintFlagList(flags, c.arg)

		assert.Equal(t, c.expected, actual)
	}

}

func TestRunBashCompletion(t *testing.T) {
	cases := []struct {
		args     []string
		expected string
	}{
		{args: []string{"ngsi"}, expected: "admin\ncreate\ncp\nlist\nget\nupdate\ndelete\n"},
		{args: []string{"ngsi", "-"}, expected: "--batch\n--stderr\n--insecureSkipVerify\n--margin\n--help\n--version\n-B\n-v\n"},
		{args: []string{"ngsi", "--"}, expected: "--batch\n--stderr\n--insecureSkipVerify\n--margin\n--help\n--version\n"},
		{args: []string{"ngsi", "--std"}, expected: "--stderr\n"},
		{args: []string{"ngsi", "--stderr"}, expected: ""},
		{args: []string{"ngsi", "--stderr", "err"}, expected: "admin\ncreate\ncp\nlist\nget\nupdate\ndelete\n"},
		{args: []string{"ngsi", "--stderr", "err", "create"}, expected: "entity\nentities\nsubscription\nregistration\ntentity\nldContext\n"},
		{args: []string{"ngsi", "--stderr", "err", "create", "-"}, expected: "--host\n--help\n-h\n"},
		{args: []string{"ngsi", "--stderr", "err", "create", "-a"}, expected: ""},
		{args: []string{"ngsi", "--stderr", "err", "create", "--"}, expected: "--host\n--help\n"},
		{args: []string{"ngsi", "--stderr", "err", "create", "--h"}, expected: "--host\n--help\n"},
		{args: []string{"ngsi", "--stderr", "err", "create", "entity"}, expected: ""},
		{args: []string{"ngsi", "--stderr", "err", "create", "entity", "-"}, expected: "--host\n--id\n--type\n--help\n-h\n-i\n-t\n"},
		{args: []string{"ngsi", "--stderr", "err", "create", "entity", "-a"}, expected: ""},
		{args: []string{"ngsi", "--stderr", "err", "create", "entity", "--"}, expected: "--host\n--id\n--type\n--help\n"},
		{args: []string{"ngsi", "--stderr", "err", "create", "entity", "--h"}, expected: "--host\n--help\n"},
		{args: []string{"ngsi", "--stderr", "err", "create", "--host", "orion"}, expected: "entity\nentities\nsubscription\nregistration\ntentity\nldContext\n"},
		{args: []string{"ngsi", "--stderr", "err", "admin"}, expected: "cacheStatistics\nlog\nmetrics\nsemaphore\nstatistics\ntrace\nscorpio\n"},
		{args: []string{"ngsi", "--stderr", "err", "admin", "scorpio"}, expected: "health\nlist\nlocaltypes\nstats\ntypes\n"},

		{args: []string{"ngsi", "reate"}, expected: ""},
	}

	for _, c := range cases {
		buf := &bytes.Buffer{}

		ngsi := setInitNGSI(nil)
		ngsi.StdWriter = buf

		f := func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
			return nil
		}
		admin := &Command{
			Name: "admin",
			Subcommands: []*Command{
				{Name: "cacheStatistics"}, {Name: "log"}, {Name: "metrics"}, {Name: "semaphore"}, {Name: "statistics"}, {Name: "trace"},
				{
					Name: "scorpio",
					Subcommands: []*Command{
						{Name: "health"}, {Name: "list"}, {Name: "localtypes"}, {Name: "stats"}, {Name: "types"},
					},
				},
			},
		}
		create := &Command{
			Name:  "create",
			Flags: []Flag{&StringFlag{Name: "host", Aliases: []string{"h"}}},
			Action: func(c *Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return f(c, ngsi, client)
			},
			Subcommands: []*Command{
				{Name: "entity",
					Flags: []Flag{
						&StringFlag{Name: "id", Aliases: []string{"i"}},
						&StringFlag{Name: "type", Aliases: []string{"t"}},
					},
				},
				{Name: "entities"},
				{Name: "subscription"},
				{Name: "registration"},
				{Name: "tentity"},
				{Name: "ldContext"},
			},
		}
		r := &App{
			Commands: []*Command{
				admin,
				create,
				{Name: "cp"},
				{Name: "list"},
				{Name: "get"},
				{Name: "update"},
				{Name: "delete"},
			},
			Flags: []Flag{
				&StringFlag{Name: "batch", Aliases: []string{"B"}},
				&StringFlag{Name: "stderr"},
				&BoolFlag{Name: "insecureSkipVerify"},
				&Int64Flag{Name: "margin"},
			},
		}

		args := append(c.args, "--generate-bash-completion")

		err := r.Run(args)

		if assert.NoError(t, err) {
			if !assert.Equal(t, c.expected, buf.String()) {
				t.Error(strings.Join(args, " "))
			}
		}
	}
}

func setInitNGSI(buf *bytes.Buffer) *ngsilib.NGSI {
	ngsilib.Reset()
	ngsi := ngsilib.NewNGSI()
	if buf != nil {
		ngsi.Stderr = buf
	}

	ngsi.ConfigFile = &MockIoLib{}
	filename := ""
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.FileReader = &MockFileLib{}
	ngsi.CacheFile = &MockIoLib{}
	ngsi.CacheFile.SetFileName(&filename)
	ngsi.PreviousArgs = &ngsilib.Settings{}

	return ngsi
}
