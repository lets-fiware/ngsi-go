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

func TestParseFlagsNoset(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)

	flag, err := parseFlags(ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, (*string)(nil), flag.Token)
		assert.Equal(t, (*string)(nil), flag.Tenant)
		assert.Equal(t, (*string)(nil), flag.Scope)
		assert.Equal(t, (*string)(nil), flag.Link)
		assert.Equal(t, (*string)(nil), flag.SafeString)
		assert.Equal(t, false, flag.XAuthToken)
	} else {
		t.FailNow()
	}
}

func TestParseFlagsSet(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "token,service,path,link,safeString")
	setupFlagBool(set, "xAuthToken")

	_ = set.Parse([]string{"--token=token", "--service=service", "--path=path", "--link=ld"})
	_ = set.Parse([]string{"--safeString=on", "--xAuthToken"})

	flag, err := parseFlags(ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, "token", *flag.Token)
		assert.Equal(t, "service", *flag.Tenant)
		assert.Equal(t, "path", *flag.Scope)
		assert.Equal(t, "https://schema.lab.fiware.org/ld/context", *flag.Link)
		assert.Equal(t, "on", *flag.SafeString)
		assert.Equal(t, true, flag.XAuthToken)
	} else {
		t.FailNow()
	}
}

func TestParseFlagsError(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "token,service,path,link,safeString")
	setupFlagBool(set, "xAuthToken")

	_ = set.Parse([]string{"--link=fiware"})

	_, err := parseFlags(ngsi, c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestParseFlags2Noset(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	// setupFlagString(set, "data")

	flag, err := parseFlags2(ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, (*string)(nil), flag.Token)
		assert.Equal(t, (*string)(nil), flag.Tenant)
		assert.Equal(t, (*string)(nil), flag.Scope)
		assert.Equal(t, (*string)(nil), flag.Link)
		assert.Equal(t, (*string)(nil), flag.SafeString)
		assert.Equal(t, false, flag.XAuthToken)
	} else {
		t.FailNow()
	}
}

func TestParseFlags2Set(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "token2,service2,path2,link2,safeString2")
	setupFlagBool(set, "xAuthToken2")

	_ = set.Parse([]string{"--token2=token", "--service2=service", "--path2=path", "--link2=ld"})
	_ = set.Parse([]string{"--safeString2=on", "--xAuthToken2"})

	flag, err := parseFlags2(ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, "token", *flag.Token)
		assert.Equal(t, "service", *flag.Tenant)
		assert.Equal(t, "path", *flag.Scope)
		assert.Equal(t, "https://schema.lab.fiware.org/ld/context", *flag.Link)
		assert.Equal(t, "on", *flag.SafeString)
		assert.Equal(t, true, flag.XAuthToken)
	} else {
		t.FailNow()
	}
}

func TestParseFlags2Error(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "token2,service2,path2,link2,safeString2")
	setupFlagBool(set, "xAuthToken2")

	_ = set.Parse([]string{"--link2=fiware"})

	_, err := parseFlags2(ngsi, c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestParseOptions(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "limit,offset,link,query")
	setupFlagBool(set, "keyValues,values,device")

	_ = set.Parse([]string{"--limit=100", "--offset=-1", "--link=etsi", "--query=query"})
	_ = set.Parse([]string{"--keyValues", "--values", "--device"})

	args := []string{"limit", "offset", "link", "query", "device"}
	opts := []string{"keyValues", "values"}

	values := parseOptions(c, args, opts)

	assert.Equal(t, "100", values.Get("limit"))
	assert.Equal(t, "etsi", values.Get("link"))
	assert.Equal(t, "query", values.Get("q"))
	assert.Equal(t, "keyValues,values", values.Get("options"))
	assert.Equal(t, "true", values.Get("device"))
}

func TestParseOptions2(t *testing.T) {
	_, set, app, _ := setupTest()

	c := cli.NewContext(app, set, nil)
	setupFlagString(set, "limit,offset,link,query")
	setupFlagBool(set, "keyValues,values")

	_ = set.Parse([]string{"--limit=100", "--offset=-1", "--link=etsi", "--query=query"})
	_ = set.Parse([]string{"--keyValues"})

	args := []string{"limit", "offset", "link", "query"}
	opts := []string{"keyValues", "values"}

	values := parseOptions(c, args, opts)

	assert.Equal(t, "100", values.Get("limit"))
	assert.Equal(t, "etsi", values.Get("link"))
	assert.Equal(t, "query", values.Get("q"))
	assert.Equal(t, "keyValues", values.Get("options"))
}
