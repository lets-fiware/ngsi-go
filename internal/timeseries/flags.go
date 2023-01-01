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

package timeseries

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

var (
	typeFlag = &ngsicli.StringFlag{
		Name:    "type",
		Aliases: []string{"t"},
		Usage:   "Entity Type",
	}
	idFlag = &ngsicli.StringFlag{
		Name:    "id",
		Aliases: []string{"i"},
		Usage:   "id",
		Value:   "",
	}
	attrFlag = &ngsicli.StringFlag{
		Name:  "attr",
		Usage: "attribute name",
	}
	georelFlag = &ngsicli.StringFlag{
		Name:  "georel",
		Usage: "georel",
	}
	geometryFlag = &ngsicli.StringFlag{
		Name:  "geometry",
		Usage: "geometry",
	}
	coordsFlag = &ngsicli.StringFlag{
		Name:  "coords",
		Usage: "coords",
	}
	attrsFlag = &ngsicli.StringFlag{
		Name:  "attrs",
		Usage: "attributes",
	}
)

// TIME SERIES
var (
	hLimitFlag = &ngsicli.Int64Flag{
		Name:  "hLimit",
		Usage: "maximum number of data entries to retrieve",
	}
	hOffsetFlag = &ngsicli.Int64Flag{
		Name:  "hOffset",
		Usage: "offset to be applied to data entries to be retrieved",
	}
	lastNFlag = &ngsicli.Int64Flag{
		Name:  "lastN",
		Usage: "number of data entries to retrieve since the final date backwards",
	}
	aggrMethodFlag = &ngsicli.StringFlag{
		Name:  "aggrMethod",
		Usage: "aggregation method (max, min, sum, sum, occur)",
	}
	aggrPeriodFlag = &ngsicli.StringFlag{
		Name:  "aggrPeriod",
		Usage: "aggregation period or resolution of the aggregated data to be retrieved",
	}
	fromDateFlag = &ngsicli.StringFlag{
		Name:  "fromDate",
		Usage: "starting date from which data should be retrieved",
	}
	toDateFlag = &ngsicli.StringFlag{
		Name:  "toDate",
		Usage: "final date until which data should be retrieved",
	}
)

// TIME SERIES (quantumleap)
var (
	dropTableFlag = &ngsicli.BoolFlag{
		Name:  "dropTable",
		Usage: "drop the table storing an entity type",
	}
	sameTypeFlag = &ngsicli.BoolFlag{
		Name:  "sameType",
		Usage: "same type",
	}
	nTypesFlag = &ngsicli.BoolFlag{
		Name:  "nTypes",
		Usage: "nTypes",
	}
	valueFlag = &ngsicli.BoolFlag{
		Name:  "value",
		Usage: "values only",
	}
)
