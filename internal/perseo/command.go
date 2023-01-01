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

package perseo

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return PerseoApp
}

var PerseoApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "Perseo command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&RulesCmd,
	},
}

// PERSEO FE Rules
var RulesCmd = ngsicli.Command{
	Name:     "rules",
	Usage:    "rules command for PERSEO",
	Category: "Context-Aware CEP",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list all plain rules",
			ServerList: []string{"perseo"},
			Flags: []ngsicli.Flag{
				perseoRulesLimitFlag,
				perseoRulesOffsetFlag,
				perseoRulesCount,
				perseoRulesRaw,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return perseoRulesList(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create a plain rule",
			ServerList: []string{"perseo"},
			Flags: []ngsicli.Flag{
				perseoRulesDataRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return perseoRulesCreate(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get a plain rule",
			ServerList: []string{"perseo"},
			Flags: []ngsicli.Flag{
				perseoRulesNameRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return perseoRulesGet(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete a plain rule",
			ServerList: []string{"perseo"},
			Flags: []ngsicli.Flag{
				perseoRulesNameRFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return perseoRulesDelete(c, ngsi, client)
			},
		},
	},
}
