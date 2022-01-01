/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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

package ngsilib

import (
	"fmt"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestIsTenantString(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "fiware", expected: true},
		{arg: "open_iot", expected: true},
		{arg: "open@iot", expected: false},
		{arg: "FIWARE", expected: false},
	}

	for _, c := range cases {
		actual := isTenantString(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsScopeString(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "", expected: true},
		{arg: "/", expected: true},
		{arg: "/fiware", expected: true},
		{arg: "/FIWARE", expected: true},
		{arg: "/FIWARE_orion", expected: true},
		{arg: "/fiware/orion", expected: true},
		{arg: "/fiware/orion,/keyrock", expected: true},
		{arg: "/fiware/orion,/keyrock, /abc, /def, /xyz/abc", expected: true},
		{arg: "/#", expected: true},
		{arg: "/*", expected: true},
		{arg: "/FIWARE@orion", expected: false},
		{arg: "FIWARE", expected: false},
	}

	for _, c := range cases {
		actual := isScopeString(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsHTTP(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "http://orion", expected: true},
		{arg: "https://orion", expected: true},
		{arg: "http:/orion", expected: false},
		{arg: "https:/orion", expected: false},
		{arg: "orion", expected: false},
	}

	for _, c := range cases {
		actual := IsHTTP(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsIPAddress(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "192.168.1.1", expected: true},
		{arg: "192.168.1.1:1026", expected: true},
		{arg: "orion", expected: false},
		{arg: "orion:1026", expected: false},
	}

	for _, c := range cases {
		actual := isIPAddress(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsLocalHost(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "localhost", expected: true},
		{arg: "localhost:1026", expected: true},
		{arg: "192.168.1.1", expected: false},
		{arg: "192.168.1.1:1026", expected: false},
	}

	for _, c := range cases {
		actual := isLocalHost(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestContains(t *testing.T) {
	list := []string{"abc", "def", "xyz", "123"}

	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "abc", expected: true},
		{arg: "123", expected: true},
		{arg: "orion", expected: false},
		{arg: "576", expected: false},
		{arg: "", expected: false},
	}

	for _, c := range cases {
		actual := Contains(list, c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsExpirationDate(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "10years", expected: true},
		{arg: "1year", expected: true},
		{arg: "65months", expected: true},
		{arg: "1month", expected: true},
		{arg: "365days", expected: true},
		{arg: "1day", expected: true},
		{arg: "123hours", expected: true},
		{arg: "1hour", expected: true},
		{arg: "orion", expected: false},
		{arg: "576", expected: false},
		{arg: "", expected: false},
	}

	for _, c := range cases {
		actual := isExpirationDate(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestGetExpirationDate(t *testing.T) {
	ngsi := NewNGSI()

	cases := []struct {
		value    string
		expected string
	}{
		{value: "1year", expected: "2007-01-02T15:04:05.000Z"},
		{value: "5years", expected: "2011-01-02T15:04:05.000Z"},
		{value: "1month", expected: "2006-02-02T15:04:05.000Z"},
		{value: "5months", expected: "2006-06-02T15:04:05.000Z"},
		{value: "1month", expected: "2006-02-02T15:04:05.000Z"},
		{value: "5months", expected: "2006-06-02T15:04:05.000Z"},
		{value: "1day", expected: "2006-01-03T15:04:05.000Z"},
		{value: "5days", expected: "2006-01-07T15:04:05.000Z"},
		{value: "1hour", expected: "2006-01-02T16:04:05.000Z"},
		{value: "5hours", expected: "2006-01-02T20:04:05.000Z"},
		{value: "1minute", expected: "2006-01-02T15:05:05.000Z"},
		{value: "-1minute", expected: "2006-01-02T15:03:05.000Z"},
		{value: "5minutes", expected: "2006-01-02T15:09:05.000Z"},
	}

	for _, c := range cases {
		ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

		date, err := GetExpirationDate(c.value)

		if assert.NoError(t, err) {
			assert.Equal(t, c.expected, date)
		}
	}
}

func TestGetExpirationDateError(t *testing.T) {
	_ = NewNGSI()

	_, err := GetExpirationDate("test")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error test", ngsiErr.Message)
	}
}

func TestIsOrionDateTime(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{value: "2022-09-24T12:07:54.035Z", expected: true},
		{value: "2022-09-24T12:07:54.035+09:00", expected: true},
		{value: "2022-09-24", expected: true},
	}

	for _, c := range cases {
		actual := IsOrionDateTime(c.value)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsNameString(t *testing.T) {
	cases := []struct {
		name string
		rc   bool
	}{
		{name: "a", rc: true},
		{name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", rc: true},
		{name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", rc: false},
		{name: "a-________---0", rc: true},
		{name: "0123", rc: false},
		{name: "user@fware", rc: true},
		{name: "user@fware.org", rc: true},
		{name: "localhost:1026", rc: true},
		{name: "", rc: false},
		{name: "0_", rc: false},
		{name: "_", rc: false},
		{name: "-", rc: false},
		{name: "@", rc: false},
	}

	for _, c := range cases {
		if b := IsNameString(c.name); b != c.rc {
			t.Error(fmt.Printf("error \"%s\" is %v", c.name, b))
		}
	}
}

func TestIsNgsiV2(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{value: "ngsi-v2", expected: true},
		{value: "ngsiv2", expected: true},
		{value: "v2", expected: true},
		{value: "ld", expected: false},
	}

	for _, c := range cases {
		actual := IsNgsiV2(c.value)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsNgsiLd(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{value: "ngsi-ld", expected: true},
		{value: "ld", expected: true},
		{value: "v2", expected: false},
	}

	for _, c := range cases {
		actual := IsNgsiLd(c.value)
		assert.Equal(t, c.expected, actual)
	}
}
