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

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return CygnusApp
}

var CygnusApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "cygnus command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&adminCmd,
		&NamemappingsCmd,
		&GroupingrulesCmd,
	},
}

var adminCmd = ngsicli.Command{
	Name:     "admin",
	Usage:    "admin command for FIWARE Orion, Cygnus, Perseo, Scorpio",
	Category: "CONVENIENCE",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		&LoggersCmd,
		&AppendersCmd,
	},
}

var AppendersCmd = ngsicli.Command{
	Name:     "appenders",
	Usage:    "manage appenders for Cygnus",
	Category: "sub-command",
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list appenders",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusAppendersTransientFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appendersList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get appender",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusAppendersNameRFlag,
				cygnusAppendersTransientFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appendersGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create appender",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusAppendersNameFlag,
				cygnusAppendersRDataFlag,
				cygnusAppendersTransientFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appendersCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update appender",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusAppendersNameRFlag,
				cygnusAppendersRDataFlag,
				cygnusAppendersTransientFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appendersUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete appender",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusAppendersNameRFlag,
				cygnusAppendersTransientFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appendersDelete(c, ngsi, client)
			},
		},
	},
}

var LoggersCmd = ngsicli.Command{
	Name:     "loggers",
	Usage:    "manage loggers for Cygnus",
	Category: "sub-command",
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list loggers",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusLoggersTransientFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return loggersList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get logger",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusLoggersNameRFlag,
				cygnusLoggersTransientFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return loggersGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create logger",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusLoggersDataRFlag,
				cygnusLoggersTransientFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return loggersCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update logger",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusLoggersDataRFlag,
				cygnusLoggersTransientFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return loggersUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete logger",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusLoggersNameRFlag,
				cygnusLoggersTransientFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return loggersDelete(c, ngsi, client)
			},
		},
	},
}

var NamemappingsCmd = ngsicli.Command{
	Name:     "namemappings",
	Usage:    "manage namemappings for Cygnus",
	Category: "PERSISTING CONTEXT DATA",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list namemappings",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return namemappingsList(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create namemapping",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusNamemappingsDataRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return namemappingsCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update namemapping",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusNamemappingsDataRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return namemappingsUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete namemapping",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusNamemappingsDataRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return namemappingsDelete(c, ngsi, client)
			},
		},
	},
}

var GroupingrulesCmd = ngsicli.Command{
	Name:     "groupingrules",
	Usage:    "manage groupingrules for Cygnus",
	Category: "PERSISTING CONTEXT DATA",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list groupingrules",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return groupingrulesList(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create groupingrule",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusGroupingrulesDataRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return groupingrulesCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update groupingrule",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusGroupingrulesIDRFlag,
				cygnusGroupingrulesDataRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"id", "data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return groupingrulesUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete groupingrule",
			ServerList: []string{"cygnus"},
			Flags: []ngsicli.Flag{
				cygnusGroupingrulesIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return groupingrulesDelete(c, ngsi, client)
			},
		},
	},
}
