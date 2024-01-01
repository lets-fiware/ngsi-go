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
)

func TestNewContext(t *testing.T) {
	actual := NewContext(&App{})

	assert.NotEqual(t, (*Context)(nil), actual)

}

func TestString(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Value: "ok"},
			&Int64Flag{Name: "int64", Value: 999},
			&BoolFlag{Name: "bool", Value: true},
		},
	}

	assert.Equal(t, "ok", c.String("string"))
	assert.Equal(t, "999", c.String("int64"))
	assert.Equal(t, "true", c.String("bool"))
	assert.Equal(t, "", c.String("unknown"))
}

func TestBool(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Value: "ok"},
			&Int64Flag{Name: "int64", Value: 999},
			&BoolFlag{Name: "bool", Value: true},
		},
	}

	assert.Equal(t, false, c.Bool("string"))
	assert.Equal(t, false, c.Bool("int64"))
	assert.Equal(t, true, c.Bool("bool"))
	assert.Equal(t, false, c.Bool("unknown"))
}

func TestInt64(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Value: "ok"},
			&Int64Flag{Name: "int64", Value: 999},
			&BoolFlag{Name: "bool", Value: true},
		},
	}

	assert.Equal(t, int64(0), c.Int64("string"))
	assert.Equal(t, int64(999), c.Int64("int64"))
	assert.Equal(t, int64(0), c.Int64("bool"))
	assert.Equal(t, int64(0), c.Int64("unknown"))
}

func TestIsSet(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "on-string", Set: true},
			&StringFlag{Name: "off-string", Set: false},
			&BoolFlag{Name: "on-bool", Set: true},
			&BoolFlag{Name: "off-bool", Set: false},
			&Int64Flag{Name: "on-int64", Set: true},
			&Int64Flag{Name: "off-int64", Set: false},
		},
	}

	assert.Equal(t, true, c.IsSet("on-string"))
	assert.Equal(t, false, c.IsSet("off-string"))
	assert.Equal(t, true, c.IsSet("on-bool"))
	assert.Equal(t, false, c.IsSet("off-bool"))
	assert.Equal(t, true, c.IsSet("on-int64"))
	assert.Equal(t, false, c.IsSet("off-int64"))
	assert.Equal(t, false, c.IsSet("unknown"))
}

func TestFlagNames(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Value: "ok"},
			&Int64Flag{Name: "int64", Value: 999},
			&BoolFlag{Name: "bool", Value: true},
		},
	}

	actual := c.FlagNames()

	assert.Equal(t, []string{"string", "int64", "bool"}, actual)
}

func TestHasFlag(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Value: "ok"},
			&Int64Flag{Name: "int64", Value: 999},
			&BoolFlag{Name: "bool", Value: true},
		},
	}

	assert.Equal(t, true, c.HasFlag("string"))
	assert.Equal(t, true, c.HasFlag("int64"))
	assert.Equal(t, true, c.HasFlag("bool"))
	assert.Equal(t, false, c.HasFlag("unknown"))
}

func TestGetFlag(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Value: "ok"},
			&Int64Flag{Name: "int64", Value: 999},
			&BoolFlag{Name: "bool", Value: true},
		},
	}

	assert.Equal(t, "ok", c.GetFlag("string").(*StringFlag).Value)
	assert.Equal(t, int64(999), c.GetFlag("int64").(*Int64Flag).Value)
	assert.Equal(t, true, c.GetFlag("bool").(*BoolFlag).Value)
	assert.Equal(t, nil, c.GetFlag("unknown"))
}

func TestIsRequired(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Required: true},
			&Int64Flag{Name: "int64", Required: true},
			&BoolFlag{Name: "bool", Required: true},
		},
	}

	assert.Equal(t, true, c.IsRequired("string"))
	assert.Equal(t, true, c.IsRequired("int64"))
	assert.Equal(t, true, c.IsRequired("bool"))
	assert.Equal(t, false, c.IsRequired("unknown"))
}

func TestArgs(t *testing.T) {
	cmdArgs := &cmdArgs{"orion", "iota", "wirecloud", "cygnus", "comet"}
	c := &Context{Arg: cmdArgs}

	acutal := c.Args()

	assert.Equal(t, "orion", acutal.Get(0))
	assert.Equal(t, "iota", acutal.Get(1))
	assert.Equal(t, "wirecloud", acutal.Get(2))
	assert.Equal(t, "cygnus", acutal.Get(3))
	assert.Equal(t, "comet", acutal.Get(4))
	assert.Equal(t, 5, acutal.Len())
}

func TestIsSetORTrue(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Set: true},
			&StringFlag{Name: "string2", Set: false},
			&Int64Flag{Name: "int64", Set: false},
			&BoolFlag{Name: "bool", Set: true},
		},
	}

	actual := c.IsSetOR([]string{"string", "int64", "bool"})
	expected := true
	assert.Equal(t, expected, actual)
}

func TestIsSetORFalse(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Set: true},
			&StringFlag{Name: "string2", Set: false},
			&Int64Flag{Name: "int64", Set: false},
			&BoolFlag{Name: "bool", Set: true},
		},
	}

	actual := c.IsSetOR([]string{"string2", "int64"})
	expected := false
	assert.Equal(t, expected, actual)
}

func TestIsSetANDTrue(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Set: true},
			&StringFlag{Name: "string2", Set: false},
			&Int64Flag{Name: "int64", Set: false},
			&BoolFlag{Name: "bool", Set: true},
		},
	}

	actual := c.IsSetAND([]string{"string", "bool"})
	expected := true

	assert.Equal(t, expected, actual)
}

func TestIsSetANDFalse(t *testing.T) {
	c := &Context{
		Flags: []Flag{
			&StringFlag{Name: "string", Set: true},
			&StringFlag{Name: "string2", Set: false},
			&Int64Flag{Name: "int64", Set: false},
			&BoolFlag{Name: "bool", Set: true},
		},
	}

	actual := c.IsSetAND([]string{"string", "string2", "int64", "bool"})
	expected := false

	assert.Equal(t, expected, actual)
}
