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

package timeseries

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return TimeseriesApp
}

var TimeseriesApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "Time series command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&HGetCmd,
		&HDeleteCmd,
	},
}

var HGetCmd = ngsicli.Command{
	Name:     "hget",
	Usage:    "get historical raw and aggregated time series context information",
	Category: "TIME SERIES",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attr",
			Usage:      "history of an attribute",
			ServerList: []string{"comet", "quantumleap"},
			Flags: []ngsicli.Flag{
				typeFlag,
				idFlag,
				attrFlag,
				sameTypeFlag,
				nTypesFlag,
				aggrMethodFlag,
				aggrPeriodFlag,
				fromDateFlag,
				toDateFlag,
				lastNFlag,
				hLimitFlag,
				hOffsetFlag,
				georelFlag,
				geometryFlag,
				coordsFlag,
				valueFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return tsAttrRead(c, ngsi, client)
			},
		},
		{
			Name:       "attrs",
			Usage:      "history of attributes",
			ServerList: []string{"quantumleap"},
			Flags: []ngsicli.Flag{
				typeFlag,
				idFlag,
				attrsFlag,
				sameTypeFlag,
				nTypesFlag,
				aggrMethodFlag,
				aggrPeriodFlag,
				fromDateFlag,
				toDateFlag,
				lastNFlag,
				hLimitFlag,
				hOffsetFlag,
				georelFlag,
				geometryFlag,
				coordsFlag,
				valueFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return qlAttrsRead(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "list of all the entity id",
			ServerList: []string{"quantumleap"},
			Flags: []ngsicli.Flag{
				typeFlag,
				fromDateFlag,
				toDateFlag,
				hLimitFlag,
				hOffsetFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return qlEntitiesRead(c, ngsi, client)
			},
		},
	},
}

var HDeleteCmd = ngsicli.Command{
	Name:     "hdelete",
	Usage:    "delete historical raw and aggregated time series context information",
	Category: "TIME SERIES",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attr",
			Usage:      "delete all the data associated to certain attribute of certain entity",
			ServerList: []string{"comet"},
			Flags: []ngsicli.Flag{
				idFlag,
				typeFlag,
				attrFlag,
				ngsicli.RunFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return cometAttrDelete(c, ngsi, client)
			},
		},
		{
			Name:       "entity",
			Usage:      "delete historical data of a certain entity",
			ServerList: []string{"comet", "quantumleap"},
			Flags: []ngsicli.Flag{
				idFlag,
				typeFlag,
				fromDateFlag,
				toDateFlag,
				ngsicli.RunFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return tsEntityDelete(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "delete historical data of all entities of a certain type",
			ServerList: []string{"comet", "quantumleap"},
			Flags: []ngsicli.Flag{
				idFlag,
				typeFlag,
				dropTableFlag,
				fromDateFlag,
				toDateFlag,
				ngsicli.RunFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return tsEntitiesDelete(c, ngsi, client)
			},
		},
	},
}
