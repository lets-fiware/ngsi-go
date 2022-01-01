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

package convenience

import (
	"github.com/lets-fiware/ngsi-go/internal/cygnus"
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return ConvenienceApp
}

var ConvenienceApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "Convenience command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&AdminCmd,
		&ApisCmd,
		&DebugCmd,
		&CopyCmd,
		&DocumentsCmd,
		&HealthCmd,
		&QueryProxyCmd,
		&ReceiverCmd,
		&RegProxyCmd,
		&RemoveCmd,
		&TokenProxyCmd,
		&VersionCmd,
	},
}

var AdminCmd = ngsicli.Command{
	Name:     "admin",
	Usage:    "admin command for FIWARE Orion, Cygnus, Perseo, Scorpio",
	Category: "CONVENIENCE",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "log",
			Usage:      "admin log",
			ServerList: []string{"brokerv2", "cygnus", "perseo"},
			Flags: []ngsicli.Flag{
				levelFlag,
				loggingFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return adminLog(c, ngsi, client)
			},
		},
		{
			Name:       "trace",
			Usage:      "admin trace",
			ServerList: []string{"brokerv2"},
			Flags: []ngsicli.Flag{
				levelFlag,
				setFlag,
				deleteFlag,
				loggingFlag,
			},
			OptionFlags: &ngsicli.ValidationFlag{Mode: ngsicli.NandCondition, Flags: []string{"set", "delete"}},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return adminTrace(c, ngsi, client)
			},
		},
		{
			Name:       "semaphore",
			Usage:      "print semaphore",
			ServerList: []string{"brokerv2"},
			Flags: []ngsicli.Flag{
				loggingFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return adminSemaphore(c, ngsi, client)
			},
		},
		{
			Name:       "metrics",
			Usage:      "manage metrics",
			ServerList: []string{"brokerv2", "perseo", "cygnus"},
			Flags: []ngsicli.Flag{
				deleteFlag,
				resetFlag,
				loggingFlag,
				ngsicli.PrettyFlag,
			},
			OptionFlags: &ngsicli.ValidationFlag{Mode: ngsicli.NandCondition, Flags: []string{"reset", "delete"}},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return adminMetrics(c, ngsi, client)
			},
		},
		{
			Name:       "statistics",
			Usage:      "print statistics",
			ServerList: []string{"brokerv2", "cygnus"},
			Flags: []ngsicli.Flag{
				deleteFlag,
				loggingFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return adminStatistics(c, ngsi, client)
			},
		},
		{
			Name:       "cacheStatistics",
			Usage:      "print cache statistics",
			ServerList: []string{"brokerv2"},
			Flags: []ngsicli.Flag{
				deleteFlag,
				loggingFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return adminCacheStatistics(c, ngsi, client)
			},
		},
		&cygnus.AppendersCmd,
		&cygnus.LoggersCmd,
		&ScorpioCmd,
	},
}

var ApisCmd = ngsicli.Command{
	Name:       "apis",
	Usage:      "print endpoints of API",
	Category:   "CONVENIENCE",
	ServerList: []string{"brokerv2", "brokerld", "quantumleap"},
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.PrettyFlag,
	},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return apis(c, ngsi, client)
	},
}

var CopyCmd = ngsicli.Command{
	Name:       "cp",
	Usage:      "copy entities",
	Category:   "CONVENIENCE",
	ServerList: []string{"brokerv2", "brokerld"},
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		destinationFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
		linkFlag,
		typeRFlag,
		token2Flag,
		tenant2Flag,
		scope2Flag,
		context2Flag,
		ngsiV1Flag,
		skipForwardingFlag,
		ngsicli.RunFlag,
	},
	RequiredFlags: []string{"type"},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return copy(c, ngsi, client)
	},
}

var DebugCmd = ngsicli.Command{
	Name:     "debug",
	Category: "CONVENIENCE",
	Usage:    "test",
	Hidden:   true,
	Flags: []ngsicli.Flag{
		&ngsicli.StringFlag{
			Name:         "host",
			Usage:        "broker or server host `VALUE`",
			Aliases:      []string{"h"},
			Required:     true,
			InitClient:   true,
			SkipGetToken: true,
			PreviousArgs: true,
		},
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	RequiredFlags: []string{"host"},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return debugCommand(c, ngsi, client)
	},
}

var DocumentsCmd = ngsicli.Command{
	Name:     "man",
	Usage:    "print urls of document",
	Category: "CONVENIENCE",
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return documents(c, ngsi, client)
	},
}

var HealthCmd = ngsicli.Command{
	Name:       "health",
	Usage:      "print health status",
	Category:   "CONVENIENCE",
	ServerList: []string{"quantumleap", "brokerv2", "brokerld", "regproxy", "tokenproxy", "queryproxy"},
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return healthCheck(c, ngsi, client)
	},
}

var QueryProxyCmd = ngsicli.Command{
	Name:     "queryproxy",
	Category: "CONVENIENCE",
	Usage:    "query proxy",
	Subcommands: []*ngsicli.Command{
		{
			Name:       "server",
			Usage:      "start up queryproxy server",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				queryProxyHostFlag,
				getPorxyReplaceURLFlag,
				queryProxyGHostFlag,
				queryProxyPortFlag,
				queryProxyHTTPSFlag,
				queryProxyKeyFlag,
				queryProxyCertFlag,
				ngsicli.VerboseFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return queryProxyServer(c, ngsi, client)
			},
		},
		{
			Name:       "health",
			Usage:      "sanity check for queryproxy server",
			ServerList: []string{"queryproxy"},
			Flags: []ngsicli.Flag{
				queryProxyQueryProxyHostFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return queryProxyHealthCmd(c, ngsi, client)
			},
		},
	},
}

var ReceiverCmd = ngsicli.Command{
	Name:     "receiver",
	Category: "CONVENIENCE",
	Usage:    "notification receiver",
	Flags: []ngsicli.Flag{
		receiverHostFlag,
		receiverPortFlag,
		receiverURLFlag,
		ngsicli.PrettyFlag,
		receiverHTTPSFlag,
		receiverKeyFlag,
		receiverCertFlag,
		ngsicli.VerboseFlag,
		headerFlag,
	},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return receiver(c, ngsi, client)
	},
}

var RegProxyCmd = ngsicli.Command{
	Name:     "regproxy",
	Category: "CONVENIENCE",
	Usage:    "registration proxy",
	Subcommands: []*ngsicli.Command{
		{
			Name:  "server",
			Usage: "start up regproxy server",
			Flags: []ngsicli.Flag{
				regProxyHostFlag,
				regProxyRhostFlag,
				regProxyPortFlag,
				regProxyURLFlag,
				regProxyReplaceTenantFlag,
				regProxyReplaceScopeFlag,
				regProxyAddScopeFlag,
				regProxyReplaceURLFlag,
				regProxyHTTPSFlag,
				regProxyKeyFlag,
				regProxyCertFlag,
				ngsicli.VerboseFlag,
			},
			ServerList: []string{"brokerv2"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return regProxyServer(c, ngsi, client)
			},
		},
		{
			Name:       "health",
			Usage:      "sanity check for regproxy server",
			ServerList: []string{"regproxy"},
			Flags: []ngsicli.Flag{
				regProxyRegProxyHostFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return regProxyHealthCmd(c, ngsi, client)
			},
		},
		{
			Name:       "config",
			Usage:      "change configuration for regproxy server",
			ServerList: []string{"regproxy"},
			Flags: []ngsicli.Flag{
				regProxyRegProxyHostFlag,
				regProxyVerboseFlag,
				regProxyReplaceTenantFlag,
				regProxyReplaceScopeFlag,
				regProxyAddScopeFlag,
				regProxyReplaceURLFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return regProxyConfigCmd(c, ngsi, client)
			},
		},
	},
}

var RemoveCmd = ngsicli.Command{
	Name:       "rm",
	Usage:      "remove entities",
	Category:   "CONVENIENCE",
	ServerList: []string{"brokerv2", "brokerld"},
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
		typeRFlag,
		linkFlag,
		ngsiV1Flag,
		skipForwardingFlag,
		ngsicli.RunFlag,
	},
	RequiredFlags: []string{"type"},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return remove(c, ngsi, client)
	},
}

var ScorpioCmd = ngsicli.Command{
	Name:     "scorpio",
	Usage:    "information command for Scorpio broker",
	Category: "sub-command",
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "List of information paths",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return scorpioCommand(c, c.Ngsi, c.Client, "")
			},
		},
		{
			Name:       "types",
			Usage:      "print types",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return scorpioCommand(c, c.Ngsi, c.Client, "types")
			},
		},
		{
			Name:       "localtypes",
			Usage:      "print local types",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return scorpioCommand(c, c.Ngsi, c.Client, "localtypes")
			},
		},
		{
			Name:       "stats",
			Usage:      "print stats",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return scorpioCommand(c, c.Ngsi, c.Client, "stats")
			},
		},
		{
			Name:       "health",
			Usage:      "print health",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return scorpioCommand(c, c.Ngsi, c.Client, "health")
			},
		},
	},
}

var TokenProxyCmd = ngsicli.Command{
	Name:     "tokenproxy",
	Category: "CONVENIENCE",
	Usage:    "token proxy",
	Subcommands: []*ngsicli.Command{
		{
			Name:  "server",
			Usage: "start up regproxy server",
			Flags: []ngsicli.Flag{
				tokenProxyHostFlag,
				tokenProxyPortFlag,
				tokenProxyHTTPSFlag,
				tokenProxyKeyFlag,
				tokenProxyCertFlag,
				tokenProxyIdmHostTenantFlag,
				tokenProxyClientIDFlag,
				tokenProxyClientSecretFlag,
				ngsicli.VerboseFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return tokenProxyServer(c, ngsi, client)
			},
		},
		{
			Name:       "health",
			Usage:      "sanity check for regproxy server",
			ServerList: []string{"tokenproxy"},
			Flags: []ngsicli.Flag{
				tokenProxyTokenProxyHostFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return tokenProxyHealthCmd(c, ngsi, client)
			},
		},
	},
}

var VersionCmd = ngsicli.Command{
	Name:     "version",
	Category: "CONVENIENCE",
	Usage:    "print the version",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.PrettyFlag,
	},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return cbVersion(c, ngsi, client)
	},
}
