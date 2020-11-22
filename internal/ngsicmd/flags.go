/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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

package ngsicmd

import (
	"github.com/urfave/cli/v2"
)

// Global Flags
var (
	syslogFlag = &cli.StringFlag{
		Name:  "syslog",
		Usage: "specify logging `LEVEL` (off, err, info, debug)",
	}
	stderrFlag = &cli.StringFlag{
		Name:  "stderr",
		Usage: "specify logging `LEVEL` (off, err, info, debug)",
	}
	logFileFlag = &cli.StringFlag{
		Name:  "logfile",
		Usage: "specify log `FILE`",
	}
	logLevelFlag = &cli.StringFlag{
		Name:  "loglevel",
		Usage: "specify logging level",
	}
	configFlag = &cli.StringFlag{
		Name:  "config",
		Usage: "specify configuration `FILE`",
	}
	cacheFlag = &cli.StringFlag{
		Name:  "cache",
		Usage: "specify cache `FILE`",
	}
	helpFlag = &cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	marginFlag = &cli.Int64Flag{
		Name:   "margin",
		Usage:  "I/O time out (second)",
		Hidden: true,
		Value:  180,
	}
	timeOutFlag = &cli.IntFlag{
		Name:   "timeout",
		Usage:  "I/O time out (second)",
		Hidden: true,
		Value:  60,
	}
	maxCountFlag = &cli.IntFlag{
		Name:   "maxCount",
		Usage:  "maxCount",
		Hidden: true,
		Value:  100,
	}
	batchFlag = &cli.BoolFlag{
		Name:    "batch",
		Aliases: []string{"B"},
		Usage:   "don't use previous args (batch)",
	}
)

// Common flags
var (
	hostFlag = &cli.StringFlag{
		Name:    "host",
		Usage:   "host or alias",
		Aliases: []string{"h"},
		Value:   "",
	}
	destinationFlag = &cli.StringFlag{
		Name:     "destination",
		Aliases:  []string{"d"},
		Required: true,
		Usage:    "host or alias",
		Value:    "",
	}
	tokenFlag = &cli.StringFlag{
		Name:  "token",
		Usage: "oauth token",
	}
	tenantFlag = &cli.StringFlag{
		Name:    "service",
		Aliases: []string{"s"},
		Usage:   "FIWARE Service",
	}
	scopeFlag = &cli.StringFlag{
		Name:    "path",
		Aliases: []string{"p"},
		Usage:   "FIWARE ServicePath",
	}
	linkFlag = &cli.StringFlag{
		Name:    "link",
		Aliases: []string{"L"},
		Usage:   "@context",
	}
	xAuthTokenFlag = &cli.BoolFlag{
		Name:   "x-auth-token",
		Usage:  "use X-Auth-Token",
		Hidden: true,
	}
	safeStringFlag = &cli.StringFlag{
		Name:  "safeString",
		Usage: "use safe string (value: on/off)",
	}
)

// Common flags for copy command
var (
	token2Flag = &cli.StringFlag{
		Name:  "token2",
		Usage: "oauth token for destination",
	}
	tenant2Flag = &cli.StringFlag{
		Name:  "service2",
		Usage: "FIWARE Service for destination",
	}
	scope2Flag = &cli.StringFlag{
		Name:  "path2",
		Usage: "FIWARE ServicePath for destination",
	}
	link2Flag = &cli.StringFlag{
		Name:  "link2",
		Usage: "@context",
		Value: "",
	}
	xAuthToken2Flag = &cli.BoolFlag{
		Name:   "x-auth-token2",
		Usage:  "use X-Auth-Token",
		Hidden: true,
	}
	safeString2Flag = &cli.StringFlag{
		Name:  "safeString2",
		Usage: "use safe string (value: on/off)",
	}
	runFlag = &cli.BoolFlag{
		Name:  "run",
		Usage: "run command",
		Value: false,
	}
)

// flags for NGSI API
var (
	typeFlag = &cli.StringFlag{
		Name:    "type",
		Aliases: []string{"t"},
		Usage:   "Entity Type",
	}
	typeRFlag = &cli.StringFlag{
		Name:     "type",
		Aliases:  []string{"t"},
		Required: true,
		Usage:    "Entity Type",
	}
	idPatternFlag = &cli.StringFlag{
		Name:  "idPattern",
		Usage: "idPattern",
	}
	typePatternFlag = &cli.StringFlag{
		Name:  "typePattern",
		Usage: "typePattern",
	}
	queryFlag = &cli.StringFlag{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "query",
	}
	mqFlag = &cli.StringFlag{
		Name:    "mq",
		Aliases: []string{"m"},
		Usage:   "mq",
	}
	georelFlag = &cli.StringFlag{
		Name:  "georel",
		Usage: "georel",
	}
	geometryFlag = &cli.StringFlag{
		Name:  "geometry",
		Usage: "geometry",
	}
	coordsFlag = &cli.StringFlag{
		Name:  "coords",
		Usage: "coords",
	}
	headersFlag = &cli.StringFlag{
		Name:  "headers",
		Usage: "headers",
	}
	qsFlag = &cli.StringFlag{
		Name:  "qs",
		Usage: "qs",
	}
	methodFlag = &cli.StringFlag{
		Name:  "method",
		Usage: "method",
	}
	payloadFlag = &cli.StringFlag{
		Name:  "payload",
		Usage: "payload",
	}
	exceptAttrsFlag = &cli.StringFlag{
		Name:  "exceptAttrs",
		Usage: "exceptAttrs",
	}
	attrsFormatFlag = &cli.StringFlag{
		Name:  "attrsFormat",
		Usage: "attrsFormat",
	}
	subjectIDFlag = &cli.StringFlag{
		Name:  "subjectId",
		Usage: "subjectId",
	}
	attrsFlag = &cli.StringFlag{
		Name:  "attrs",
		Usage: "attrs",
		Value: "",
	}
	metadataFlag = &cli.StringFlag{
		Name:  "metadata",
		Usage: "metadata",
		Value: "",
	}
	orderByFlag = &cli.StringFlag{
		Name:  "orderBy",
		Usage: "orderBy",
		Value: "",
	}
	actionTypeFlag = &cli.StringFlag{
		Name:     "actionType",
		Usage:    "actionType",
		Required: true,
	}
	attrNameRFlag = &cli.StringFlag{
		Name:     "attrName",
		Usage:    "attrName",
		Value:    "",
		Required: true,
	}
	providedIDFlag = &cli.StringFlag{
		Name:  "providedId",
		Usage: "providedId",
	}
	legacyFlag = &cli.BoolFlag{
		Name:  "legacy",
		Usage: "legacy forwarding mode",
	}
	forwardingModeFlag = &cli.StringFlag{
		Name:  "forwardingMode",
		Usage: "forwarding mode",
	}
)

// flags for options
var (
	countFlag = &cli.BoolFlag{
		Name:    "count",
		Aliases: []string{"C"},
		Usage:   "count",
	}
	keyValuesFlag = &cli.BoolFlag{
		Name:    "keyValues",
		Aliases: []string{"K"},
		Usage:   "keyValues",
	}
	valuesFlag = &cli.BoolFlag{
		Name:    "values",
		Aliases: []string{"V"},
		Usage:   "values",
	}
	uniqueFlag = &cli.BoolFlag{
		Name:    "unique",
		Aliases: []string{"U"},
		Usage:   "unique",
	}
	upsertFlag = &cli.BoolFlag{
		Name:  "upsert",
		Usage: "upsert",
	}
	appendFlag = &cli.BoolFlag{
		Name:    "append",
		Aliases: []string{"a"},
		Usage:   "append",
	}
	noOverwriteFlag = &cli.BoolFlag{
		Name:    "noOverwrite",
		Aliases: []string{"n"},
		Usage:   "noOverwrite",
	}
	replaceFlag = &cli.BoolFlag{
		Name:    "replace",
		Aliases: []string{"r"},
		Usage:   "replace",
	}
	updateFlag = &cli.BoolFlag{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "update",
	}
	sysAttrsFlag = &cli.BoolFlag{
		Name:    "sysAttrs",
		Aliases: []string{"S"},
		Usage:   "sysAttrs",
	}
	linesFlag = &cli.BoolFlag{
		Name:    "lines",
		Aliases: []string{"1"},
		Usage:   "lines",
	}
)

var (
	dataFlag = &cli.StringFlag{
		Name:    "data",
		Usage:   "data",
		Aliases: []string{"d"},
		Value:   "",
	}
	idFlag = &cli.StringFlag{
		Name:    "id",
		Aliases: []string{"i"},
		Usage:   "id",
		Value:   "",
	}
	idRFlag = &cli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "id",
		Required: true,
	}
	uriFlag = &cli.StringFlag{
		Name:  "uri",
		Usage: "url or uri",
	}
	notifyURLFlag = &cli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Usage:   "url to be invoked when a notification is generated",
	}
	expiresFlag = &cli.BoolFlag{
		Name:    "expires",
		Aliases: []string{"e"},
		Usage:   "expires",
	}
	expiresSFlag = &cli.StringFlag{
		Name:    "expires",
		Aliases: []string{"e"},
		Usage:   "expires",
	}
	throttlingFlag = &cli.IntFlag{
		Name:  "throttling",
		Usage: "throttling",
	}
	verboseFlag = &cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "verbose",
	}
	jsonFlag = &cli.BoolFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "JSON format",
	}
	skipInitialNotificationFlag = &cli.BoolFlag{
		Name:  "skipInitialNotification",
		Usage: "skipInitialNotification",
	}
	localTimeFlag = &cli.BoolFlag{
		Name:  "localTime",
		Usage: "localTime",
	}
	getFlag = &cli.BoolFlag{
		Name:  "get",
		Usage: "get",
	}
	statusFlag = &cli.StringFlag{
		Name:  "status",
		Usage: "status",
	}
	descriptionFlag = &cli.StringFlag{
		Name:  "description",
		Usage: "description",
	}
	providerFlag = &cli.StringFlag{
		Name:    "provider",
		Aliases: []string{"p"},
		Usage:   "Url of context provider/source",
	}
	wAttrsFlag = &cli.StringFlag{
		Name:  "wAttrs",
		Usage: "watched attributes",
	}
	nAttrsFlag = &cli.StringFlag{
		Name:  "nAttrs",
		Usage: "attributes to be notified",
	}
)

// flag for broker config
var (
	brokerHostFlag = &cli.StringFlag{
		Name:    "brokerHost",
		Aliases: []string{"b"},
		Usage:   "specify context broker host",
	}
	ngsiTypeFlag = &cli.StringFlag{
		Name:  "ngsiType",
		Usage: "specify NGSI type: v2 or ld",
	}
	idmTypeFlag = &cli.StringFlag{
		Name:    "idmType",
		Aliases: []string{"t"},
		Usage:   "specify token type",
	}
	idmHostFlag = &cli.StringFlag{
		Name:    "idmHost",
		Aliases: []string{"m"},
		Usage:   "specify identity manager host",
	}
	apiPathFlag = &cli.StringFlag{
		Name:    "apiPath",
		Aliases: []string{"a"},
		Usage:   "specify API path",
	}
	usernameFlag = &cli.StringFlag{
		Name:    "username",
		Aliases: []string{"U"},
		Usage:   "specify username",
	}
	passwordFlag = &cli.StringFlag{
		Name:    "password",
		Aliases: []string{"P"},
		Usage:   "specify password",
	}
	clientIDFlag = &cli.StringFlag{
		Name:    "clientId",
		Aliases: []string{"I"},
		Usage:   "specify client id",
	}
	clientSecretFlag = &cli.StringFlag{
		Name:    "clientSecret",
		Aliases: []string{"S"},
		Usage:   "specify client secret",
	}
	itemsFlag = &cli.StringFlag{
		Name:    "items",
		Aliases: []string{"i"},
		Usage:   "specify itmes",
	}
	allFlag = &cli.BoolFlag{
		Name:  "all",
		Usage: "ail itmes",
	}
)

// flag for context
var (
	nameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "@context name",
	}
	nameRFlag = &cli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "@context name",
		Required: true,
	}
	urlFlag = &cli.StringFlag{
		Name:     "url",
		Aliases:  []string{"u"},
		Usage:    "url for @context",
		Required: true,
	}
)
