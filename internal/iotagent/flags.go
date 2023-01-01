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

package iotagent

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

var (
	typeFlag = &ngsicli.StringFlag{
		Name:    "type",
		Aliases: []string{"t"},
		Usage:   "Entity Type",
	}
)

// IoT Agent
var (
	servicesLimitFlag = &ngsicli.Int64Flag{
		Name:  "limit",
		Usage: "maximum number of services",
	}
	servicesOffsetFlag = &ngsicli.Int64Flag{
		Name:  "offset",
		Usage: "offset to skip a given number of elements at the beginning",
	}
	resourceFlag = &ngsicli.StringFlag{
		Name:  "resource",
		Usage: "uri for the iotagent",
	}
	resourceRFlag = &ngsicli.StringFlag{
		Name:     "resource",
		Usage:    "uri for the iotagent",
		Required: true,
	}
	apikeyFlag = &ngsicli.StringFlag{
		Name:  "apikey",
		Usage: "a key used for devices belonging to this service",
	}
	cbrokerFlag = &ngsicli.StringFlag{
		Name:  "cbroker",
		Usage: "url of context broker or broker alias",
	}
	servicesDeviceFlag = &ngsicli.BoolFlag{
		Name:  "device",
		Usage: "remove devices in service/subservice",
		Value: false,
	}
	servicesDataFlag = &ngsicli.StringFlag{
		Name:  "data",
		Usage: "data body (payload)",
	}
	servicesTokenFlag = &ngsicli.StringFlag{
		Name:  "token",
		Usage: "token obtained from the authentication system",
	}
	devicesLimit = &ngsicli.Int64Flag{
		Name:  "limit",
		Usage: "maximum number of devices",
	}
	devicesOffset = &ngsicli.Int64Flag{
		Name:  "offset",
		Usage: "offset to skip a given number of elements at the beginning",
	}
	devicesDetailed = &ngsicli.StringFlag{
		Name:    "detailed",
		Usage:   "all device information (on) or only name (off)",
		Value:   "off",
		Choices: []string{"off", "on"},
	}
	devicesEntity = &ngsicli.StringFlag{
		Name:  "entity",
		Usage: "get a device from entity name",
	}
	devicesProtocol = &ngsicli.StringFlag{
		Name:  "protocol",
		Usage: "get devices with this protocol",
	}
	devicesDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Usage:    "data body (payload)",
		Required: true,
	}
	devicesIDRFlag = &ngsicli.StringFlag{
		Name:     "id",
		Usage:    "device id",
		Required: true,
	}
)
