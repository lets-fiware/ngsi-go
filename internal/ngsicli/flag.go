/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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
	"fmt"
	"strconv"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

type Flag interface {
	FlagName() string
	FlagNameList() string
	FlagAliases() []string
	FlagUsage() string
	FlagHidden() bool
	Check(name, alias string) bool
	IsSet() bool
	IsRequired() bool
	IsInitClient() bool
	Copy(reset bool) Flag
	SetValue(v interface{}) error
	AllowEmpty() bool
}

type StringFlag struct {
	Name         string
	Aliases      []string
	Usage        string
	Value        string
	Hidden       bool
	Set          bool
	Required     bool
	InitClient   bool
	SkipGetToken bool
	ValueEmpty   bool
	PreviousArgs bool
	SkipRefHost  bool
	Choices      []string
}

func (f *StringFlag) FlagName() string {
	return f.Name
}

func (f *StringFlag) FlagNameList() string {
	return flagName(f.Name, f.Aliases, f.Usage, " VALUE")
}

func (f *StringFlag) FlagAliases() []string {
	return f.Aliases
}

func (f *StringFlag) FlagUsage() string {
	u := strings.Replace(f.Usage, "`", "", -1)
	if f.Required {
		u += " (required)"
	}
	return u
}

func (f *StringFlag) FlagHidden() bool {
	return f.Hidden
}

func (f *StringFlag) IsSet() bool {
	return f.Set
}

func (f *StringFlag) IsRequired() bool {
	return f.Required
}

func (f *StringFlag) IsInitClient() bool {
	return f.InitClient
}

func (f *StringFlag) Check(name, alias string) bool {
	if name != "" && f.Name == name {
		return true
	}
	if alias != "" {
		for _, a := range f.Aliases {
			if alias == a {
				return true
			}
		}
	}
	return false
}

func (f *StringFlag) Copy(reset bool) Flag {
	nf := *f
	nf.Set = !reset

	aliases := make([]string, len(f.Aliases))
	_ = copy(aliases, f.Aliases)
	nf.Aliases = aliases

	return &nf
}

func (f *StringFlag) SetValue(v interface{}) error {
	const funcName = "StringFlagSetValue"
	var ok bool
	f.Value, ok = v.(string)
	if ok {
		f.Set = true
		return nil
	} else {
		return ngsierr.New(funcName, 1, "type error", nil)

	}
}

func (f *StringFlag) AllowEmpty() bool {
	return f.ValueEmpty
}

type BoolFlag struct {
	Name     string
	Aliases  []string
	Usage    string
	Value    bool
	Hidden   bool
	Set      bool
	Required bool
}

func (f *BoolFlag) FlagName() string {
	return f.Name
}

func (f *BoolFlag) FlagNameList() string {
	return flagName(f.Name, f.Aliases, f.Usage, "")
}

func (f *BoolFlag) FlagAliases() []string {
	return f.Aliases
}

func (f *BoolFlag) FlagUsage() string {
	u := strings.Replace(f.Usage, "`", "", -1)
	u = fmt.Sprintf("%s (default: %t)", u, f.Value)
	if f.Required {
		u += " (required)"
	}
	return u
}

func (f *BoolFlag) FlagHidden() bool {
	return f.Hidden
}

func (f *BoolFlag) IsSet() bool {
	return f.Set
}

func (f *BoolFlag) IsRequired() bool {
	return f.Required
}

func (f *BoolFlag) IsInitClient() bool {
	return false
}

func (f *BoolFlag) Check(name, alias string) bool {
	if name != "" && f.Name == name {
		return true
	}
	if alias != "" {
		for _, a := range f.Aliases {
			if alias == a {
				return true
			}
		}
	}
	return false
}

func (f *BoolFlag) Copy(reset bool) Flag {
	nf := *f
	nf.Set = !reset

	aliases := make([]string, len(f.Aliases))
	_ = copy(aliases, f.Aliases)
	nf.Aliases = aliases

	return &nf
}

func (f *BoolFlag) SetValue(v interface{}) error {
	const funcName = "BoolFlagSetValue"
	var b bool

	switch val := v.(type) {
	default:
		return ngsierr.New(funcName, 1, "type error", nil)
	case bool:
		b = val
	case string:
		switch strings.ToLower(val) {
		default:
			return ngsierr.New(funcName, 2, val+" is not boolean value", nil)
		case "true", "on":
			b = true
		case "false", "off":
			b = false
		}
	}
	f.Value = b
	f.Set = true
	return nil
}

func (f *BoolFlag) AllowEmpty() bool {
	return false
}

type Int64Flag struct {
	Name     string
	Aliases  []string
	Usage    string
	Value    int64
	Hidden   bool
	Set      bool
	Required bool
}

func (f *Int64Flag) FlagName() string {
	return f.Name
}

func (f *Int64Flag) FlagNameList() string {
	return flagName(f.Name, f.Aliases, f.Usage, " VALUE")
}

func (f *Int64Flag) FlagAliases() []string {
	return f.Aliases
}

func (f *Int64Flag) FlagUsage() string {
	u := strings.Replace(f.Usage, "`", "", -1)
	if f.Required {
		u += " (required)"
	}
	return u
}

func (f *Int64Flag) FlagHidden() bool {
	return f.Hidden
}

func (f *Int64Flag) IsSet() bool {
	return f.Set
}

func (f *Int64Flag) IsRequired() bool {
	return f.Required
}

func (f *Int64Flag) IsInitClient() bool {
	return false
}

func (f *Int64Flag) Check(name, alias string) bool {
	if name != "" && f.Name == name {
		return true
	}
	if alias != "" {
		for _, a := range f.Aliases {
			if alias == a {
				return true
			}
		}
	}
	return false
}

// Copy Flag
func (f *Int64Flag) Copy(reset bool) Flag {
	nf := *f
	nf.Set = !reset

	aliases := make([]string, len(f.Aliases))
	_ = copy(aliases, f.Aliases)
	nf.Aliases = aliases

	return &nf
}

// Set Valut to Flag
func (f *Int64Flag) SetValue(v interface{}) error {
	const funcName = "Int64FlagSetValue"
	switch val := v.(type) {
	default:
		return ngsierr.New(funcName, 1, "type error", nil)
	case int:
		f.Value = int64(val)
	case int64:
		f.Value = val
	case string:
		var err error
		f.Value, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return ngsierr.New(funcName, 2, val+" is not number", nil)
		}
	}
	f.Set = true
	return nil
}

func (f *Int64Flag) AllowEmpty() bool {
	return false
}

// FlagName for print
func flagName(name string, aliases []string, usage string, arg string) string {

	s := strings.Index(usage, "`")
	e := strings.LastIndex(usage, "`")
	if s != -1 && e != -1 && s < e {
		arg = " " + usage[s+1:e]
	}

	n := "--" + name + arg

	for _, alias := range aliases {
		n += ", -" + alias + arg
	}

	return n
}

// remove flag from flags
func removeFlag(flags []Flag, name string) []Flag {
	if flags == nil {
		return flags
	}
	newFlags := []Flag{}
	for _, f := range flags {
		if f.FlagName() != name {
			newFlags = append(newFlags, f)
		}
	}
	return newFlags
}
