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

func TestStringFlag(t *testing.T) {
	f := StringFlag{
		Name:         "test",
		Aliases:      []string{"a", "b", "c"},
		Value:        "value",
		Usage:        "usage",
		Hidden:       true,
		Set:          true,
		Required:     true,
		InitClient:   true,
		SkipGetToken: true,
		ValueEmpty:   true,
	}

	assert.Equal(t, "test", f.FlagName())
	assert.Equal(t, "--test VALUE, -a VALUE, -b VALUE, -c VALUE", f.FlagNameList())
	assert.Equal(t, []string{"a", "b", "c"}, f.FlagAliases())
	assert.Equal(t, "usage (required)", f.FlagUsage())
	assert.Equal(t, true, f.FlagHidden())
	assert.Equal(t, true, f.IsSet())
	assert.Equal(t, true, f.IsRequired())
	assert.Equal(t, true, f.IsInitClient())
	assert.Equal(t, true, f.Check("test", ""))
	assert.Equal(t, true, f.Check("", "a"))
	assert.Equal(t, false, f.Check("", "t"))
}

func TestStringFlagCopy(t *testing.T) {
	f := StringFlag{
		Name:  "test",
		Value: "value",
		Set:   true,
	}

	ff := f.Copy(false)
	assert.Equal(t, "value", ff.(*StringFlag).Value)
	assert.Equal(t, true, ff.IsSet())

	ff = f.Copy(true)
	assert.Equal(t, "value", ff.(*StringFlag).Value)
	assert.Equal(t, false, ff.IsSet())
}

func TestStringFlagSetValue(t *testing.T) {
	f := StringFlag{
		Name:  "test",
		Value: "",
		Set:   false,
	}

	assert.Equal(t, "", f.Value)
	assert.Equal(t, false, f.Set)

	err := f.SetValue("value")

	if assert.NoError(t, err) {
		assert.Equal(t, "value", f.Value)
		assert.Equal(t, true, f.Set)
	}
}

func TestStringFlagSetValueError(t *testing.T) {
	f := StringFlag{
		Name:  "test",
		Value: "",
		Set:   false,
	}

	assert.Equal(t, "", f.Value)
	assert.Equal(t, false, f.Set)

	err := f.SetValue(1)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "type error", ngsiErr.Message)
	}
}

func TestStringFlagAllowEmptyFalse(t *testing.T) {
	cases := []struct {
		f        StringFlag
		expected bool
	}{
		{StringFlag{Name: "test", Value: "", Set: false, ValueEmpty: false}, false},
		{StringFlag{Name: "test", Value: "", Set: false, ValueEmpty: true}, true},
	}

	for _, c := range cases {
		actual := c.f.AllowEmpty()
		assert.Equal(t, c.expected, actual)
	}
}

func TestBoolFlag(t *testing.T) {
	f := BoolFlag{
		Name:     "test",
		Aliases:  []string{"a", "b", "c"},
		Value:    true,
		Usage:    "usage",
		Hidden:   true,
		Set:      true,
		Required: true,
	}

	assert.Equal(t, "test", f.FlagName())
	assert.Equal(t, "--test, -a, -b, -c", f.FlagNameList())
	assert.Equal(t, []string{"a", "b", "c"}, f.FlagAliases())
	assert.Equal(t, "usage (default: true) (required)", f.FlagUsage())
	assert.Equal(t, true, f.FlagHidden())
	assert.Equal(t, true, f.IsSet())
	assert.Equal(t, true, f.IsRequired())
	assert.Equal(t, false, f.IsInitClient())
	assert.Equal(t, true, f.Check("test", ""))
	assert.Equal(t, true, f.Check("", "a"))
	assert.Equal(t, false, f.Check("", "t"))
}
func TestBoolFlagCopy(t *testing.T) {
	f := BoolFlag{
		Name:  "test",
		Value: true,
		Set:   true,
	}

	ff := f.Copy(false)
	assert.Equal(t, true, ff.(*BoolFlag).Value)
	assert.Equal(t, true, ff.IsSet())

	ff = f.Copy(true)
	assert.Equal(t, true, ff.(*BoolFlag).Value)
	assert.Equal(t, false, ff.IsSet())
}

func TestBoolFlagSetValue(t *testing.T) {
	cases := []struct {
		InitValue bool
		Value     interface{}
		Expect    bool
	}{
		{InitValue: false, Value: true, Expect: true},
		{InitValue: true, Value: false, Expect: false},
		{InitValue: false, Value: "true", Expect: true},
		{InitValue: true, Value: "false", Expect: false},
		{InitValue: false, Value: "True", Expect: true},
		{InitValue: true, Value: "False", Expect: false},
		{InitValue: false, Value: "tRUE", Expect: true},
		{InitValue: true, Value: "fALSE", Expect: false},
	}

	for _, c := range cases {
		f := BoolFlag{
			Name:  "test",
			Value: c.InitValue,
			Set:   false,
		}

		assert.Equal(t, c.InitValue, f.Value)
		assert.Equal(t, false, f.Set)

		err := f.SetValue(c.Value)

		if assert.NoError(t, err) {
			assert.Equal(t, c.Expect, f.Value)
			assert.Equal(t, true, f.Set)
		}
	}
}

func TestBoolFlagSetValueErrorInt(t *testing.T) {
	f := BoolFlag{
		Name:  "test",
		Value: true,
		Set:   false,
	}

	assert.Equal(t, true, f.Value)
	assert.Equal(t, false, f.Set)

	err := f.SetValue(1)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "type error", ngsiErr.Message)
	}
}

func TestBoolFlagSetValueErrorString(t *testing.T) {
	f := BoolFlag{
		Name:  "test",
		Value: true,
		Set:   false,
	}

	assert.Equal(t, true, f.Value)
	assert.Equal(t, false, f.Set)

	err := f.SetValue("enable")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "enable is not boolean value", ngsiErr.Message)
	}
}

func TestBoolFlagAllowEmpty(t *testing.T) {
	f := BoolFlag{Name: "test", Value: true, Set: false}

	actual := f.AllowEmpty()
	expected := false

	assert.Equal(t, expected, actual)
}

func TestInt64Flag(t *testing.T) {
	f := Int64Flag{
		Name:     "test",
		Aliases:  []string{"a", "b", "c"},
		Value:    1,
		Usage:    "usage",
		Hidden:   true,
		Set:      true,
		Required: true,
	}

	assert.Equal(t, "test", f.FlagName())
	assert.Equal(t, "--test VALUE, -a VALUE, -b VALUE, -c VALUE", f.FlagNameList())
	assert.Equal(t, []string{"a", "b", "c"}, f.FlagAliases())
	assert.Equal(t, "usage (required)", f.FlagUsage())
	assert.Equal(t, true, f.FlagHidden())
	assert.Equal(t, true, f.IsSet())
	assert.Equal(t, true, f.IsRequired())
	assert.Equal(t, false, f.IsInitClient())
	assert.Equal(t, true, f.Check("test", ""))
	assert.Equal(t, true, f.Check("", "a"))
	assert.Equal(t, false, f.Check("", "t"))
}
func TestInt64FlagCopy(t *testing.T) {
	f := Int64Flag{
		Name:  "test",
		Value: 1,
		Set:   true,
	}

	ff := f.Copy(false)
	assert.Equal(t, int64(1), ff.(*Int64Flag).Value)
	assert.Equal(t, true, ff.IsSet())

	ff = f.Copy(true)
	assert.Equal(t, int64(1), ff.(*Int64Flag).Value)
	assert.Equal(t, false, ff.IsSet())
}

func TestInt64FlagSetValue(t *testing.T) {
	cases := []struct {
		InitValue int64
		Value     interface{}
		Expect    int64
	}{
		{InitValue: 1, Value: 9, Expect: 9},
		{InitValue: 1, Value: int64(100), Expect: 100},
		{InitValue: 1, Value: "1000", Expect: 1000},
	}

	for _, c := range cases {
		f := Int64Flag{
			Name:  "test",
			Value: c.InitValue,
			Set:   false,
		}

		assert.Equal(t, c.InitValue, f.Value)
		assert.Equal(t, false, f.Set)

		err := f.SetValue(c.Value)

		if assert.NoError(t, err) {
			assert.Equal(t, c.Expect, f.Value)
			assert.Equal(t, true, f.Set)
		}
	}
}

func TestInt64FlagSetValueErrorBool(t *testing.T) {
	f := Int64Flag{
		Name:  "test",
		Value: 1,
		Set:   false,
	}

	assert.Equal(t, int64(1), f.Value)
	assert.Equal(t, false, f.Set)

	err := f.SetValue(true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "type error", ngsiErr.Message)
	}
}

func TestInt64FlagSetValueErrorString(t *testing.T) {
	f := Int64Flag{
		Name:  "test",
		Value: 1,
		Set:   false,
	}

	assert.Equal(t, int64(1), f.Value)
	assert.Equal(t, false, f.Set)

	err := f.SetValue("abc")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "abc is not number", ngsiErr.Message)
	}
}

func TestInt64FlagAllowEmpty(t *testing.T) {
	f := Int64Flag{Name: "test", Value: 1, Set: false}

	actual := f.AllowEmpty()
	expected := false

	assert.Equal(t, expected, actual)
}

func TestFlagName(t *testing.T) {
	cases := []struct {
		name     string
		aliases  []string
		usage    string
		expected string
	}{
		{name: "test", aliases: []string{"a", "b", "c"}, usage: "usage", expected: "--test VALUE, -a VALUE, -b VALUE, -c VALUE"},
		{name: "test", aliases: []string{"a", "b", "c"}, usage: "usage `FLAG`", expected: "--test FLAG, -a FLAG, -b FLAG, -c FLAG"},
		{name: "test", aliases: []string{"a", "b", "c"}, usage: "usage `FLAG", expected: "--test VALUE, -a VALUE, -b VALUE, -c VALUE"},
	}

	for _, c := range cases {
		actual := flagName(c.name, c.aliases, c.usage, " VALUE")

		assert.Equal(t, c.expected, actual)
	}
}

func TestRemoveFlag(t *testing.T) {
	cases := []struct {
		name     string
		expected int
	}{
		{name: "test", expected: 3},
		{name: "test1", expected: 2},
		{name: "test2", expected: 2},
		{name: "test3", expected: 2},
	}

	for _, c := range cases {
		flags := []Flag{
			&StringFlag{Name: "test1", Value: "value", Set: true},
			&StringFlag{Name: "test2", Value: "value", Set: true},
			&StringFlag{Name: "test3", Value: "value", Set: true},
		}
		f := removeFlag(flags, c.name)
		actual := len(f)
		assert.Equal(t, c.expected, actual)
	}
}

func TestRemoveFlagNil(t *testing.T) {
	actual := removeFlag(nil, "test")
	expected := []Flag(nil)
	assert.Equal(t, expected, actual)
}
