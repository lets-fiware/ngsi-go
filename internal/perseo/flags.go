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

package perseo

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

// PERSEO FE
var (
	perseoRulesNameRFlag = &ngsicli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "rule name",
		Required: true,
	}
	perseoRulesDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Aliases:  []string{"d"},
		Usage:    "rule data",
		Required: true,
	}
	perseoRulesLimitFlag = &ngsicli.Int64Flag{
		Name:  "limit",
		Usage: "maximum number of rules",
	}
	perseoRulesOffsetFlag = &ngsicli.Int64Flag{
		Name:  "offset",
		Usage: "offset to skip a given number of rules at the beginning",
	}
	perseoRulesRaw = &ngsicli.BoolFlag{
		Name:  "raw",
		Usage: "print raw data",
	}
	perseoRulesCount = &ngsicli.BoolFlag{
		Name:  "count",
		Usage: "print number of rules",
	}
)
