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

package ngsicli

import (
	"fmt"
	"strconv"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type Context struct {
	App            *App
	Ngsi           *ngsilib.NGSI
	Client         *ngsilib.Client
	Client2        *ngsilib.Client
	Flags          []Flag
	GlobalFlags    []Flag
	RequiredFlags  []string
	Arg            *cmdArgs
	Commands       []*Command
	ServerList     []string
	CommandName    string
	HelpCommand    bool
	Bashcompletion bool
}

func NewContext(r *App) *Context {
	return &Context{
		App:           r,
		Ngsi:          ngsilib.NewNGSI(),
		Flags:         []Flag{},
		GlobalFlags:   []Flag{},
		RequiredFlags: []string{},
		Arg:           &cmdArgs{},
	}
}

func (c *Context) String(flag string) string {
	for _, ff := range c.Flags {
		if ff.FlagName() == flag {
			switch f := ff.(type) {
			case *StringFlag:
				return f.Value
			case *Int64Flag:
				return strconv.FormatInt(f.Value, 10)
			case *BoolFlag:
				return fmt.Sprintf("%t", f.Value)
			}
		}
	}
	return ""
}

func (c *Context) Bool(flag string) bool {
	for _, ff := range c.Flags {
		if ff.FlagName() == flag {
			f, ok := ff.(*BoolFlag)
			if ok {
				return f.Value
			}
		}
	}
	return false
}

func (c *Context) Int64(flag string) int64 {
	for _, ff := range c.Flags {
		if ff.FlagName() == flag {
			f, ok := ff.(*Int64Flag)
			if ok {
				return f.Value
			}
		}
	}
	return 0
}

func (c *Context) IsSet(flag string) bool {
	for _, ff := range c.Flags {
		if ff.FlagName() == flag {
			return ff.IsSet()
		}
	}
	return false
}

func (c *Context) FlagNames() []string {
	names := []string{}

	for _, flag := range c.Flags {
		names = append(names, flag.FlagName())
	}
	return names
}

func (c *Context) HasFlag(name string) bool {
	for _, flag := range c.Flags {
		if flag.FlagName() == name {
			return true
		}
	}
	return false
}

func (c *Context) GetFlag(name string) Flag {
	for _, flag := range c.Flags {
		if flag.FlagName() == name {
			return flag
		}
	}
	return nil
}

func (c *Context) GetStringFlag(name string) *StringFlag {
	for _, flag := range c.Flags {
		if flag.FlagName() == name {
			if f, ok := flag.(*StringFlag); ok {
				return f
			}
		}
	}
	return nil
}

func (c *Context) IsRequired(name string) bool {
	for _, flag := range c.Flags {
		if flag.FlagName() == name {
			return flag.IsRequired()
		}
	}
	return false
}

func (c *Context) Args() Args {
	return c.Arg
}

func (c *Context) IsSetOR(params []string) bool {
	for _, param := range params {
		if c.IsSet(param) {
			return true
		}
	}
	return false
}

func (c *Context) IsSetAND(params []string) bool {
	for _, param := range params {
		if !c.IsSet(param) {
			return false
		}
	}
	return true
}
