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

package management

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

// flag for broker config
var (
	hostBrokerFlag = &ngsicli.StringFlag{
		Name:        "host",
		Aliases:     []string{"h"},
		Usage:       "context broker host alias",
		SkipRefHost: true,
	}
	hostBrokerRFlag = &ngsicli.StringFlag{
		Name:        "host",
		Aliases:     []string{"h"},
		Usage:       "context broker host alias",
		Required:    true,
		SkipRefHost: true,
	}
	hostBrokerRPFlag = &ngsicli.StringFlag{
		Name:         "host",
		Aliases:      []string{"h"},
		Usage:        "context broker host alias",
		Required:     true,
		PreviousArgs: true,
		SkipRefHost:  true,
	}
	brokerHostFlag = &ngsicli.StringFlag{
		Name:    "brokerHost",
		Aliases: []string{"b"},
		Usage:   "context broker host address or alias",
	}
	ngsiTypeFlag = &ngsicli.StringFlag{
		Name:  "ngsiType",
		Usage: "NGSI type: v2 or ld",
	}
	brokerTypeFlag = &ngsicli.StringFlag{
		Name:  "brokerType",
		Usage: "NGSI-LD broker type: orion-ld, scorpio or stellio",
	}
	idmTypeFlag = &ngsicli.StringFlag{
		Name:    "idmType",
		Aliases: []string{"t"},
		Usage:   "token type",
	}
	idmHostFlag = &ngsicli.StringFlag{
		Name:    "idmHost",
		Aliases: []string{"m"},
		Usage:   "identity manager host",
	}
	apiPathFlag = &ngsicli.StringFlag{
		Name:    "apiPath",
		Aliases: []string{"a"},
		Usage:   "API path",
	}
	usernameFlag = &ngsicli.StringFlag{
		Name:    "username",
		Aliases: []string{"U"},
		Usage:   "username",
	}
	passwordFlag = &ngsicli.StringFlag{
		Name:    "password",
		Aliases: []string{"P"},
		Usage:   "password",
	}
	clientIDFlag = &ngsicli.StringFlag{
		Name:    "clientId",
		Aliases: []string{"I"},
		Usage:   "client id",
	}
	clientSecretFlag = &ngsicli.StringFlag{
		Name:    "clientSecret",
		Aliases: []string{"S"},
		Usage:   "client secret",
	}
	tokenScopeFlag = &ngsicli.StringFlag{
		Name:  "tokenScope",
		Usage: "scope for token",
	}
	itemsFlag = &ngsicli.StringFlag{
		Name:    "items",
		Aliases: []string{"i"},
		Usage:   "itmes",
	}
	itemsHFlag = &ngsicli.StringFlag{
		Name:    "items",
		Aliases: []string{"i"},
		Usage:   "itmes",
		Hidden:  true,
	}
	allFlag = &ngsicli.BoolFlag{
		Name:  "all",
		Usage: "ail itmes",
	}
	headerNameFlag = &ngsicli.StringFlag{
		Name:  "headerName",
		Usage: "header name for apikey",
	}
	headerValueFlag = &ngsicli.StringFlag{
		Name:  "headerValue",
		Usage: "header value for apikey",
	}
	headerEnvValueFlag = &ngsicli.StringFlag{
		Name:  "headerEnvValue",
		Usage: "name of environment variable for apikey",
	}
	clearTextFlag = &ngsicli.BoolFlag{
		Name:  "clearText",
		Usage: "show obfuscated items as clear text",
	}
	singleLineFlag = &ngsicli.BoolFlag{
		Name:    "singleLine",
		Aliases: []string{"1"},
		Usage:   "list one file per line",
		Value:   false,
	}
	tokenFlag = &ngsicli.StringFlag{
		Name:  "token",
		Usage: "token `VALUE`",
	}
	brokerOverWrite = &ngsicli.BoolFlag{
		Name:    "overWrite",
		Aliases: []string{"O"},
		Usage:   "overwrite broker alias",
	}
)

// flag for server config
var (
	hostServerFlag = &ngsicli.StringFlag{
		Name:        "host",
		Aliases:     []string{"h"},
		Usage:       "server host alias",
		SkipRefHost: true,
	}
	hostServerRFlag = &ngsicli.StringFlag{
		Name:        "host",
		Aliases:     []string{"h"},
		Usage:       "server host alias",
		Required:    true,
		SkipRefHost: true,
	}
	hostServerRPFlag = &ngsicli.StringFlag{
		Name:         "host",
		Aliases:      []string{"h"},
		Usage:        "server host alias",
		Required:     true,
		PreviousArgs: true,
		SkipRefHost:  true,
	}
	serverHostFlag = &ngsicli.StringFlag{
		Name:  "serverHost",
		Usage: "server host address or alias",
	}
	serverTypeFlag = &ngsicli.StringFlag{
		Name:  "serverType",
		Usage: "serverType (comet, ql)",
	}
	allServersFlag = &ngsicli.BoolFlag{
		Name:   "all",
		Usage:  "print all servers",
		Hidden: true,
	}
	serverOverWrite = &ngsicli.BoolFlag{
		Name:    "overWrite",
		Aliases: []string{"O"},
		Usage:   "overwrite server alias",
	}
)

// flag for context
var (
	contextNameFlag = &ngsicli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "@context name",
	}
	contextNameRFlag = &ngsicli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "@context name",
		Required: true,
	}
	contextDataFlag = &ngsicli.StringFlag{
		Name:       "data",
		Usage:      "@context data",
		Aliases:    []string{"d"},
		Value:      "",
		ValueEmpty: false,
	}
	contextUrlFlag = &ngsicli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Usage:   "url for @context",
	}
	contextUrlRFlag = &ngsicli.StringFlag{
		Name:     "url",
		Aliases:  []string{"u"},
		Usage:    "url for @context",
		Required: true,
	}
	contextJsonFlag = &ngsicli.StringFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "url for @context",
	}
)

// flag for context server
var (
	contextServerHostFlag = &ngsicli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Value:   "0.0.0.0",
		Usage:   "host for server",
	}
	contextServerPortFlag = &ngsicli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "3000",
		Usage:   "port for server",
	}
	contextServerURLFlag = &ngsicli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Value:   "/",
		Usage:   "url for server",
	}
	contextServerHTTPSFlag = &ngsicli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	contextServerKeyFlag = &ngsicli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	contextServerCertFlag = &ngsicli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
)

var (
	onFlag = &ngsicli.BoolFlag{
		Name:    "on",
		Aliases: []string{"e"},
		Usage:   "on (enable)",
	}
	offFlag = &ngsicli.BoolFlag{
		Name:    "off",
		Aliases: []string{"d"},
		Usage:   "off (disable)",
	}
)

var (
	expiresFlag = &ngsicli.BoolFlag{
		Name:    "expires",
		Aliases: []string{"e"},
		Usage:   "expires",
	}
	revokeFlag = &ngsicli.BoolFlag{
		Name:    "revoke",
		Aliases: []string{"r"},
		Usage:   "revoke token",
	}
)
