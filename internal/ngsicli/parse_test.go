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

package ngsicli

import (
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestParseFlagsNoset(t *testing.T) {
	c := setupTestInitCmd()

	flag, err := parseFlags(c.Ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, (*string)(nil), flag.Token)
		assert.Equal(t, (*string)(nil), flag.Tenant)
		assert.Equal(t, (*string)(nil), flag.Scope)
		assert.Equal(t, (*string)(nil), flag.Link)
		assert.Equal(t, (*string)(nil), flag.SafeString)
		assert.Equal(t, false, flag.XAuthToken)
	}
}

func TestParseFlagsSet(t *testing.T) {
	c := setupTestInitCmd()

	token := tokenFlag.Copy(true)
	err := token.SetValue("oAuthToken")
	assert.NoError(t, err)

	service := tenantFlag.Copy(true)
	err = service.SetValue("service")
	assert.NoError(t, err)

	scope := scopeFlag.Copy(true)
	err = scope.SetValue("path")
	assert.NoError(t, err)

	link := linkFlag.Copy(true)
	err = link.SetValue("ld")
	assert.NoError(t, err)

	safeString := safeStringFlag.Copy(true)
	err = safeString.SetValue("on")
	assert.NoError(t, err)

	xAuthToken := xAuthTokenFlag.Copy(true)
	err = xAuthToken.SetValue(true)
	assert.NoError(t, err)

	c.Flags = []Flag{token, service, scope, link, safeString, xAuthToken}

	flag, err := parseFlags(c.Ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, "oAuthToken", *flag.Token)
		assert.Equal(t, "service", *flag.Tenant)
		assert.Equal(t, "path", *flag.Scope)
		assert.Equal(t, "https://schema.lab.fiware.org/ld/context", *flag.Link)
		assert.Equal(t, "on", *flag.SafeString)
		assert.Equal(t, true, flag.XAuthToken)
	}
}

func TestParseFlagsError(t *testing.T) {
	c := setupTestInitCmd()

	link := linkFlag.Copy(true)
	err := link.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{link}

	_, err = parseFlags(c.Ngsi, c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestParseFlags2Noset(t *testing.T) {
	c := setupTestInitCmd()

	flag, err := parseFlags2(c.Ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, (*string)(nil), flag.Token)
		assert.Equal(t, (*string)(nil), flag.Tenant)
		assert.Equal(t, (*string)(nil), flag.Scope)
		assert.Equal(t, (*string)(nil), flag.Link)
		assert.Equal(t, (*string)(nil), flag.SafeString)
		assert.Equal(t, false, flag.XAuthToken)
	}
}

func TestParseFlags2Set(t *testing.T) {
	c := setupTestInitCmd()

	token := token2Flag.Copy(true)
	err := token.SetValue("token")
	assert.NoError(t, err)

	service := tenant2Flag.Copy(true)
	err = service.SetValue("service")
	assert.NoError(t, err)

	scope := scope2Flag.Copy(true)
	err = scope.SetValue("path")
	assert.NoError(t, err)

	link := link2Flag.Copy(true)
	err = link.SetValue("etsi")
	assert.NoError(t, err)

	c.Flags = []Flag{token, service, scope, link}

	flag, err := parseFlags2(c.Ngsi, c)

	if assert.NoError(t, err) {
		assert.Equal(t, "token", *flag.Token)
		assert.Equal(t, "service", *flag.Tenant)
		assert.Equal(t, "path", *flag.Scope)
		assert.Equal(t, "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld", *flag.Link)
	}
}

func TestParseFlags2Error(t *testing.T) {
	c := setupTestInitCmd()

	link := link2Flag.Copy(true)
	err := link.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{link}

	_, err = parseFlags2(c.Ngsi, c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestParseOptions(t *testing.T) {
	c := setupTestInitCmd()

	query := queryFlag.Copy(true)
	err := query.SetValue("id=1")
	assert.NoError(t, err)

	device := servicesDeviceFlag.Copy(true)
	err = device.SetValue(true)
	assert.NoError(t, err)

	details := &BoolFlag{Name: "details", Value: true, Set: true}

	key := &StringFlag{Name: "key", Value: "value", Set: true}

	c.Flags = []Flag{query, device, details, key}

	u := ParseOptions(c, []string{"query", "device", "details", "key"}, nil)

	assert.Equal(t, "id=1", u.Get("q"))
	assert.Equal(t, "true", u.Get("device"))
	assert.Equal(t, "true", u.Get("details"))
	assert.Equal(t, "value", u.Get("key"))
}

func TestParseOptionsDetailsFalse(t *testing.T) {
	c := setupTestInitCmd()

	details := &BoolFlag{Name: "details", Value: false, Set: true}

	c.Flags = []Flag{details}

	u := ParseOptions(c, []string{"details"}, nil)

	assert.Equal(t, "false", u.Get("details"))
}

func TestParseOptionsLimit(t *testing.T) {
	c := setupTestInitCmd()

	limit := &Int64Flag{Name: "limit", Value: 1, Set: true}
	offset := &Int64Flag{Name: "offset", Value: 2, Set: true}
	hLimit := &Int64Flag{Name: "hLimit", Value: 3, Set: true}
	hOffset := &Int64Flag{Name: "hOffset", Value: 4, Set: true}
	lastN := &Int64Flag{Name: "lastN", Value: 5, Set: true}

	c.Flags = []Flag{limit, offset, hLimit, hOffset, lastN}

	u := ParseOptions(c, []string{"limit", "offset", "hLimit", "hOffset", "lastN"}, nil)

	assert.Equal(t, "1", u.Get("limit"))
	assert.Equal(t, "2", u.Get("offset"))
	assert.Equal(t, "3", u.Get("hLimit"))
	assert.Equal(t, "4", u.Get("hOffset"))
	assert.Equal(t, "5", u.Get("lastN"))
}

func TestParseOptionsOptions(t *testing.T) {
	c := setupTestInitCmd()

	keyValues := &BoolFlag{Name: "keyValues", Value: true, Set: true}
	values := &BoolFlag{Name: "values", Value: true, Set: true}

	c.Flags = []Flag{keyValues, values}

	u := ParseOptions(c, nil, []string{"keyValues", "values"})

	assert.Equal(t, "keyValues,values", u.Get("options"))
}

func TestParseOptionsOptionsKeyValues(t *testing.T) {
	c := setupTestInitCmd()

	keyValues := &BoolFlag{Name: "keyValues", Value: true, Set: true}
	values := &BoolFlag{Name: "values", Value: false, Set: true}

	c.Flags = []Flag{keyValues, values}

	u := ParseOptions(c, nil, []string{"keyValues", "values"})

	assert.Equal(t, "keyValues", u.Get("options"))
}

func TestParseOptionsOptionsValues(t *testing.T) {
	c := setupTestInitCmd()

	keyValues := &BoolFlag{Name: "keyValues", Value: false, Set: true}
	values := &BoolFlag{Name: "values", Value: true, Set: true}

	c.Flags = []Flag{keyValues, values}

	u := ParseOptions(c, nil, []string{"keyValues", "values"})

	assert.Equal(t, "values", u.Get("options"))
}

func TestParseOptionsOptionsOff(t *testing.T) {
	c := setupTestInitCmd()

	keyValues := &BoolFlag{Name: "keyValues", Value: true, Set: false}
	values := &BoolFlag{Name: "values", Value: true, Set: false}

	c.Flags = []Flag{keyValues, values}

	u := ParseOptions(c, nil, []string{"keyValues", "values"})

	assert.Equal(t, "", u.Get("options"))
}
