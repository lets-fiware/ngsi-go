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

package cygnus

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

// Cygnus namemappings

var (
	cygnusNamemappingsDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Aliases:  []string{"d"},
		Usage:    "name mapping data",
		Required: true,
	}
)

// Cygnus appenders
var (
	cygnusAppendersNameFlag = &ngsicli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "appender name",
	}
	cygnusAppendersNameRFlag = &ngsicli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "appender name",
		Required: true,
	}
	cygnusAppendersRDataFlag = &ngsicli.StringFlag{
		Name:     "data",
		Aliases:  []string{"d"},
		Usage:    "appender information",
		Required: true,
	}
	cygnusAppendersTransientFlag = &ngsicli.BoolFlag{
		Name:    "transient",
		Aliases: []string{"t"},
		Usage:   "true, retrieving from memory, or false, retrieving from file",
		Value:   false,
	}
)

// Cygnus groupingrules

var (
	cygnusGroupingrulesIDRFlag = &ngsicli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "grouping rule id",
		Required: true,
	}
	cygnusGroupingrulesDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Aliases:  []string{"d"},
		Usage:    "grouping rule data",
		Required: true,
	}
)

// Cygnus loggers
var (
	cygnusLoggersNameRFlag = &ngsicli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "logger name",
		Required: true,
	}
	cygnusLoggersDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Aliases:  []string{"d"},
		Usage:    "logger information",
		Required: true,
	}
	cygnusLoggersTransientFlag = &ngsicli.BoolFlag{
		Name:    "transient",
		Aliases: []string{"t"},
		Usage:   "true, retrieving from memory, or false, retrieving from file",
		Value:   false,
	}
)
