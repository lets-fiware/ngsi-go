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

package management

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return ManagementApp
}

var ManagementApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "keyrock command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&BrokersCmd,
		&ServerCmd,
		&ContextCmd,
		&SettingsCmd,
		&TokenCmd,
	},
}

var BrokersCmd = ngsicli.Command{
	Name:     "broker",
	Usage:    "manage config for broker",
	Category: "MANAGEMENT",
	Subcommands: []*ngsicli.Command{
		{
			Name:  "list",
			Usage: "list brokers",
			Flags: []ngsicli.Flag{
				hostBrokerFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
				clearTextFlag,
				singleLineFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return brokersList(c, ngsi, client)
			},
		},
		{
			Name:  "get",
			Usage: "get broker",
			Flags: []ngsicli.Flag{
				hostBrokerRPFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
				clearTextFlag,
			},
			RequiredFlags: []string{"host"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return brokersGet(c, ngsi, client)
			},
		},
		{
			Name:  "add",
			Usage: "add broker",
			Flags: []ngsicli.Flag{
				hostBrokerRFlag,
				brokerHostFlag,
				ngsicli.NgsiTypeFlag,
				brokerTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				headerNameFlag,
				headerValueFlag,
				headerEnvValueFlag,
				tokenScopeFlag,
				tokenFlag,
				ngsicli.TenantFlag,
				ngsicli.ScopeFlag,
				ngsicli.SafeStringFlag,
				ngsicli.XAuthTokenFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return brokersAdd(c, ngsi, client)
			},
		},
		{
			Name:  "update",
			Usage: "update broker",
			Flags: []ngsicli.Flag{
				hostBrokerRFlag,
				brokerHostFlag,
				ngsiTypeFlag,
				brokerTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				headerNameFlag,
				headerValueFlag,
				headerEnvValueFlag,
				tokenScopeFlag,
				tokenFlag,
				ngsicli.TenantFlag,
				ngsicli.ScopeFlag,
				ngsicli.SafeStringFlag,
				ngsicli.XAuthTokenFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return brokersUpdate(c, ngsi, client)
			},
		},
		{
			Name:  "delete",
			Usage: "delete broker",
			Flags: []ngsicli.Flag{
				hostBrokerRFlag,
				itemsFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return brokersDelete(c, ngsi, client)
			},
		},
	},
}

var ServerCmd = ngsicli.Command{
	Name:     "server",
	Usage:    "manage config for server",
	Category: "MANAGEMENT",
	Subcommands: []*ngsicli.Command{
		{
			Name:  "list",
			Usage: "list servers",
			Flags: []ngsicli.Flag{
				hostServerFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
				allServersFlag,
				clearTextFlag,
				singleLineFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return serverList(c, ngsi, client)
			},
		},
		{
			Name:  "get",
			Usage: "get server",
			Flags: []ngsicli.Flag{
				hostServerRPFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
				clearTextFlag,
			},
			RequiredFlags: []string{"host"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return serverGet(c, ngsi, client)
			},
		},
		{
			Name:  "add",
			Usage: "add server",
			Flags: []ngsicli.Flag{
				hostServerRFlag,
				serverHostFlag,
				serverTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				headerNameFlag,
				headerValueFlag,
				headerEnvValueFlag,
				tokenScopeFlag,
				tokenFlag,
				ngsicli.TenantFlag,
				ngsicli.ScopeFlag,
				ngsicli.SafeStringFlag,
				ngsicli.XAuthTokenFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return serverAdd(c, ngsi, client)
			},
		},
		{
			Name:  "update",
			Usage: "update server",
			Flags: []ngsicli.Flag{
				hostServerRFlag,
				serverHostFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				headerNameFlag,
				headerValueFlag,
				headerEnvValueFlag,
				tokenScopeFlag,
				tokenFlag,
				ngsicli.TenantFlag,
				ngsicli.ScopeFlag,
				ngsicli.SafeStringFlag,
				ngsicli.XAuthTokenFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return serverUpdate(c, ngsi, client)
			},
		},
		{
			Name:  "delete",
			Usage: "delete server",
			Flags: []ngsicli.Flag{
				hostServerRFlag,
				itemsHFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return serverDelete(c, ngsi, client)
			},
		},
	},
}

var ContextCmd = ngsicli.Command{
	Name:     "context",
	Usage:    "manage @context",
	Category: "MANAGEMENT",
	Subcommands: []*ngsicli.Command{
		{
			Name:  "list",
			Usage: "List @context",
			Flags: []ngsicli.Flag{
				contextNameFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return contextList(c, ngsi, client)
			},
		},
		{
			Name:  "add",
			Usage: "Add @context",
			Flags: []ngsicli.Flag{
				contextNameRFlag,
				contextUrlFlag,
				contextJsonFlag,
			},
			OptionFlags: &ngsicli.ValidationFlag{Mode: ngsicli.XnorCondition, Flags: []string{"url", "json"}},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return contextAdd(c, ngsi, client)
			},
		},
		{
			Name:  "update",
			Usage: "Update @context",
			Flags: []ngsicli.Flag{
				contextNameRFlag,
				contextUrlRFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return contextUpdate(c, ngsi, client)
			},
		},
		{
			Name:  "delete",
			Usage: "Delete @context",
			Flags: []ngsicli.Flag{
				contextNameRFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return contextDelete(c, ngsi, client)
			},
		},
		{
			Name:  "server",
			Usage: "serve @context",
			Flags: []ngsicli.Flag{
				contextNameFlag,
				contextDataFlag,
				contextServerHostFlag,
				contextServerPortFlag,
				contextServerURLFlag,
				contextServerHTTPSFlag,
				contextServerKeyFlag,
				contextServerCertFlag,
			},
			OptionFlags: &ngsicli.ValidationFlag{Mode: ngsicli.XnorCondition, Flags: []string{"name", "data"}},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return contextServer(c, ngsi, client)
			},
		},
	},
}

var SettingsCmd = ngsicli.Command{
	Name:     "settings",
	Category: "MANAGEMENT",
	Usage:    "manage settings",
	Subcommands: []*ngsicli.Command{
		{
			Name:  "list",
			Usage: "List settings",
			Flags: []ngsicli.Flag{
				allFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return settingsList(c, ngsi, client)
			},
		},
		{
			Name:  "delete",
			Usage: "Delete setting",
			Flags: []ngsicli.Flag{
				itemsFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return settingsDelete(c, ngsi, client)
			},
		},
		{
			Name:  "clear",
			Usage: "Clear settings",
			Flags: []ngsicli.Flag{},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return settingsClear(c, ngsi, client)
			},
		},
		{
			Name:  "previousArgs",
			Usage: "Set PreviousArgs mode",
			Flags: []ngsicli.Flag{
				offFlag,
				onFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return settingsPreviousArgs(c, ngsi, client)
			},
		},
	},
}

var TokenCmd = ngsicli.Command{
	Name:  "token",
	Usage: "manage token",
	Flags: []ngsicli.Flag{
		ngsicli.HostFlag,
		ngsicli.VerboseFlag,
		ngsicli.PrettyFlag,
		expiresFlag,
		revokeFlag,
	},
	Category: "MANAGEMENT",
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return tokenCommand(c, ngsi, client)
	},
}
