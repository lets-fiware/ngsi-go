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

package convenience

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

// flags for Admin command
var (
	levelFlag = &ngsicli.StringFlag{
		Name:    "level",
		Aliases: []string{"l"},
		Usage:   "log level",
	}
	deleteFlag = &ngsicli.BoolFlag{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete",
	}
	setFlag = &ngsicli.BoolFlag{
		Name:    "set",
		Aliases: []string{"s"},
		Usage:   "set",
	}
	resetFlag = &ngsicli.BoolFlag{
		Name:    "reset",
		Aliases: []string{"r"},
		Usage:   "reset",
	}
	loggingFlag = &ngsicli.BoolFlag{
		Name:    "logging",
		Aliases: []string{"L"},
		Usage:   "logging",
	}
)

// flags for copy command
var (
	typeRFlag = &ngsicli.StringFlag{
		Name:     "type",
		Aliases:  []string{"t"},
		Usage:    "Entity Type",
		Required: true,
	}
	linkFlag = &ngsicli.StringFlag{
		Name:    "link",
		Aliases: []string{"L"},
		Usage:   "@context `VALUE` (LD)",
	}
	destinationFlag = &ngsicli.StringFlag{
		Name:     "host2",
		Aliases:  []string{"d"},
		Usage:    "host or alias",
		Value:    "",
		Required: true,
	}
	token2Flag = &ngsicli.StringFlag{
		Name:   "OAuthToken2",
		Usage:  "oauth token for destination",
		Hidden: true,
	}
	tenant2Flag = &ngsicli.StringFlag{
		Name:  "service2",
		Usage: "FIWARE Service for destination",
	}
	scope2Flag = &ngsicli.StringFlag{
		Name:  "path2",
		Usage: "FIWARE ServicePath for destination",
	}
	context2Flag = &ngsicli.StringFlag{
		Name:  "context2",
		Usage: "@context for destination",
	}
	ngsiV1Flag = &ngsicli.BoolFlag{
		Name:  "ngsiV1",
		Usage: "NGSI v1 mode",
	}
	skipForwardingFlag = &ngsicli.BoolFlag{
		Name:  "skipForwarding",
		Usage: "skip forwarding to CPrs (v2)",
	}
)

// flag for receiver
var (
	receiverHostFlag = &ngsicli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Value:   "0.0.0.0",
		Usage:   "host for receiver",
	}
	receiverPortFlag = &ngsicli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1028",
		Usage:   "port for receiver",
	}
	receiverURLFlag = &ngsicli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Value:   "/",
		Usage:   "url for receiver",
	}
	receiverHTTPSFlag = &ngsicli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	receiverKeyFlag = &ngsicli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	receiverCertFlag = &ngsicli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	headerFlag = &ngsicli.BoolFlag{
		Name:  "header",
		Usage: "print receive header",
	}
)

// flag for registration proxy
var (
	regProxyHostFlag = &ngsicli.StringFlag{
		Name:         "host",
		Aliases:      []string{"h"},
		Usage:        "context broker or csource host",
		Required:     true,
		InitClient:   true,
		SkipGetToken: true,
	}
	regProxyRhostFlag = &ngsicli.StringFlag{
		Name:  "rhost",
		Value: "0.0.0.0",
		Usage: "host for registration proxy",
	}
	regProxyPortFlag = &ngsicli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1028",
		Usage:   "port for registration proxy",
	}
	regProxyURLFlag = &ngsicli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Value:   "/",
		Usage:   "url for registration proxy",
	}
	regProxyHTTPSFlag = &ngsicli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	regProxyKeyFlag = &ngsicli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	regProxyCertFlag = &ngsicli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	regProxyReplaceTenantFlag = &ngsicli.StringFlag{
		Name:  "replaceService",
		Usage: "replace FIWARE-Serivce",
	}
	regProxyReplaceScopeFlag = &ngsicli.StringFlag{
		Name:  "replacePath",
		Usage: "replace FIWARE-SerivcePath",
	}
	regProxyAddScopeFlag = &ngsicli.StringFlag{
		Name:  "addPath",
		Usage: "add path to FIWARE-SerivcePath",
	}
	regProxyReplaceURLFlag = &ngsicli.StringFlag{
		Name:  "replaceURL",
		Usage: "replace URL of forwarding destination",
	}
	regProxyRegProxyHostFlag = &ngsicli.StringFlag{
		Name:       "host",
		Aliases:    []string{"h"},
		Usage:      "regproxy host",
		Required:   true,
		InitClient: true,
	}
	regProxyVerboseFlag = &ngsicli.StringFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "verbose log (on/off)",
	}
)

// flag for tokenproxy
var (
	tokenProxyHostFlag = &ngsicli.StringFlag{
		Name:  "host",
		Value: "0.0.0.0",
		Usage: "host for tokenproxy",
	}
	tokenProxyPortFlag = &ngsicli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1029",
		Usage:   "port for tokenproxy",
	}
	tokenProxyHTTPSFlag = &ngsicli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	tokenProxyKeyFlag = &ngsicli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	tokenProxyCertFlag = &ngsicli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	tokenProxyIdmHostTenantFlag = &ngsicli.StringFlag{
		Name:  "idmHost",
		Usage: "host for Keyrock",
	}
	tokenProxyClientIDFlag = &ngsicli.StringFlag{
		Name:    "clientId",
		Aliases: []string{"I"},
		Usage:   "client id for Keyrock",
	}
	tokenProxyClientSecretFlag = &ngsicli.StringFlag{
		Name:    "clientSecret",
		Aliases: []string{"S"},
		Usage:   "client secret for Keyrock",
	}
	tokenProxyTokenProxyHostFlag = &ngsicli.StringFlag{
		Name:       "host",
		Aliases:    []string{"h"},
		Usage:      "tokenproxy server host",
		Required:   true,
		InitClient: true,
	}
)

// flag for queryproxy
var (
	queryProxyHostFlag = &ngsicli.StringFlag{
		Name:       "host",
		Aliases:    []string{"h"},
		Usage:      "context broker",
		Required:   true,
		InitClient: true,
	}
	getPorxyReplaceURLFlag = &ngsicli.StringFlag{
		Name:    "replaceURL",
		Aliases: []string{"u"},
		Usage:   "replace URL",
		Value:   "/v2/ex/entities",
	}
	queryProxyGHostFlag = &ngsicli.StringFlag{
		Name:  "qhost",
		Value: "0.0.0.0",
		Usage: "host for queryproxy",
	}
	queryProxyPortFlag = &ngsicli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1030",
		Usage:   "port for queryproxy",
	}
	queryProxyHTTPSFlag = &ngsicli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	queryProxyKeyFlag = &ngsicli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	queryProxyCertFlag = &ngsicli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	queryProxyQueryProxyHostFlag = &ngsicli.StringFlag{
		Name:       "host",
		Aliases:    []string{"h"},
		Usage:      "queryproxy server host",
		Required:   true,
		InitClient: true,
	}
)
