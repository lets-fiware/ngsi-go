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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewClient(t *testing.T) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	_ = set.Parse([]string{"--host=orion"})
	c := cli.NewContext(app, set, nil)

	ngsi, err := initCmd(c, "", true)
	assert.NoError(t, err)

	_, err = newClient(ngsi, c, false, []string{"broker"})

	assert.NoError(t, err)
}

func TestNewClientError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "link")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--link=fiware"})

	_, err := newClient(ngsi, c, false, []string{"broker"})

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
