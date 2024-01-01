/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

package iotagent

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return IotagentkApp
}

var IotagentkApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "Iot Agent command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&ServicesCmd,
		&DevicesCmd,
	},
}

var ServicesCmd = ngsicli.Command{
	Name:     "services",
	Usage:    "manage services for IoT Agent",
	Category: "IoT Agent",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list configuration group",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				servicesLimitFlag,
				servicesOffsetFlag,
				resourceFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasServicesList(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create a configuration group",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				servicesDataFlag,
				apikeyFlag,
				servicesTokenFlag,
				cbrokerFlag,
				typeFlag,
				resourceFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasServicesCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update a configuration group",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				resourceRFlag,
				servicesDataFlag,
				apikeyFlag,
				servicesTokenFlag,
				cbrokerFlag,
				typeFlag,
			},
			RequiredFlags: []string{"resource"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasServicesUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "remove a configuration group",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				resourceRFlag,
				apikeyFlag,
				servicesDeviceFlag,
			},
			RequiredFlags: []string{"resource"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasServicesDelete(c, ngsi, client)
			},
		},
	},
}

var DevicesCmd = ngsicli.Command{
	Name:     "devices",
	Usage:    "manage devices for IoT Agent",
	Category: "IoT Agent",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list all devices",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				devicesLimit,
				devicesOffset,
				devicesDetailed,
				devicesEntity,
				devicesProtocol,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasDevicesList(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create a device",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				devicesDataRFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasDevicesCreate(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get a device",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				devicesIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasDevicesGet(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update a device",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				devicesIDRFlag,
				devicesDataRFlag,
			},
			RequiredFlags: []string{"id", "data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasDevicesUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete a device",
			ServerList: []string{"iota"},
			Flags: []ngsicli.Flag{
				devicesIDRFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return idasDevicesDelete(c, ngsi, client)
			},
		},
	},
}
