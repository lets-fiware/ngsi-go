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
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestContextList(t *testing.T) {
	ngsilib.Reset()

	_, set, app, buf := setupTest()
	setupFlagString(set, "name")

	c := cli.NewContext(app, set, nil)
	err := contextList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "etsi https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\nld https://schema.lab.fiware.org/ld/context\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListName(t *testing.T) {
	ngsilib.Reset()

	_, set, app, buf := setupTest()
	setupFlagString(set, "name")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi"})
	err := contextList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListErrorInitCmd(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextListErrorName(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestContextAdd(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware", "--url=http://fiware"})
	err := contextAdd(c)

	if assert.NoError(t, err) {
		_, set, app, _ := setupTest()
		setupFlagString(set, "name")
		c := cli.NewContext(app, set, nil)
		_ = set.Parse([]string{"--name=fiware"})
		_ = contextDelete(c)
	}
}

func TestContextAddErrorInitCmd(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextAddErrorName(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "name not found", ngsiErr.Message)
	}
}

func TestContextAddErrorNameString(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=@fiware"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "name error @fiware", ngsiErr.Message)
	}
}

func TestContextAddErrorUrl(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url not found", ngsiErr.Message)
	}
}

func TestContextAddErrorUrlError(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware", "--url=abc"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "abc is not url", ngsiErr.Message)
	}
}

func TestContextUpdate(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi", "--url=http://fiware"})
	err := contextUpdate(c)

	assert.NoError(t, err)
}

func TestContextUpdateErrorInitCmd(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextUpdateErrorName(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "name not found", ngsiErr.Message)
	}
}

func TestContextUpdateErrorUrl(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url not found", ngsiErr.Message)
	}
}

func TestContextUpdateErrorUrlError(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi", "--url=abc"})
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "abc is not url", ngsiErr.Message)
	}
}

func TestContextDelete(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi", "--url=http://fiware"})
	err := contextDelete(c)

	assert.NoError(t, err)
}

func TestContextDeleteErrorInitCmd(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextDeleteErrorName(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	err := contextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "name not found", ngsiErr.Message)
	}
}

func TestContextDeleteErrorUrl(t *testing.T) {
	ngsilib.Reset()

	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}
