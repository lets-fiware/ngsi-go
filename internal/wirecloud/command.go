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

package wirecloud

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return WirecloudApp
}

var WirecloudApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "wirecloud command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&PreferencesCmd,
		&ResourcesCmd,
		&TabsCmd,
		&WorkspacesCmd,
	},
}

var PreferencesCmd = ngsicli.Command{
	Name:     "preferences",
	Usage:    "manage preferences for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "get",
			Usage:      "get preferences",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudPreferencesGet(c, ngsi, client)
			},
		},
	},
}

var ResourcesCmd = ngsicli.Command{
	Name:     "macs",
	Usage:    "manage mashable application components for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list mashable application components",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudWidgetFlag,
				wireCloudOperatorFlag,
				wireCloudMashupFlag,
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudResourcesList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get mashable application component",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudResourceGet(c, ngsi, client)
			},
		},
		{
			Name:       "download",
			Usage:      "download mashable application component",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudResourceDownload(c, ngsi, client)
			},
		},
		{
			Name:       "install",
			Usage:      "install mashable application component",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudFileRFlag,
				wireCloudPublicFlag,
				wireCloudOverwriteFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"file"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudResourceInstall(c, ngsi, client)
			},
		},
		{
			Name:       "uninstall",
			Usage:      "uninstall mashable application component",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
				ngsicli.RunFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudResourceUninstall(c, ngsi, client)
			},
		},
	},
}

var WorkspacesCmd = ngsicli.Command{
	Name:     "workspaces",
	Usage:    "manage workspaces for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list workspaces",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudWorkspacesList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get workspace",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudWorkspaceIdRFlag,
				wireCloudUsersFlag,
				wireCloudTabsFlag,
				wireCloudWidgetsFlag,
				wireCloudOperatorsFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"wid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudWorkspaceGet(c, ngsi, client)
			},
		},
	},
}

var TabsCmd = ngsicli.Command{
	Name:     "tabs",
	Usage:    "manage tabs for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list tabs",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudWorkspaceIdRFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"wid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudTabsList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get tab",
			ServerList: []string{"wirecloud"},
			Flags: []ngsicli.Flag{
				wireCloudWorkspaceIdRFlag,
				wireCloudTabIdRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"wid", "tid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return wireCloudTabGet(c, ngsi, client)
			},
		},
	},
}
