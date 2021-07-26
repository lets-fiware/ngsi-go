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
	cmdNameFlag = &cli.StringFlag{
		Name:   "cmdName",
		Hidden: true,
	}
	insecureSkipVerifyFlag = &cli.BoolFlag{
		Name:  "insecureSkipVerify",
		Usage: "TLS/SSL skip certificate verification",
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
		Usage:   "@context (LD)",
	}
	contextFlag = &cli.StringFlag{
		Name:    "context",
		Aliases: []string{"C"},
		Usage:   "@context (LD)",
	}
	acceptJSONFlag = &cli.BoolFlag{
		Name:  "acceptJson",
		Usage: "set accecpt header to application/json (LD)",
		Value: false,
	}
	acceptGeoJSONFlag = &cli.BoolFlag{
		Name:  "acceptGeoJson",
		Usage: "set accecpt header to application/geo+json (LD)",
		Value: false,
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
	ngsiV1Flag = &cli.BoolFlag{
		Name:  "ngsiV1",
		Usage: "NGSI v1 mode",
	}
	clearTextFlag = &cli.BoolFlag{
		Name:  "clearText",
		Usage: "show obfuscated items as clear text",
	}
	onFlag = &cli.BoolFlag{
		Name:    "on",
		Aliases: []string{"e"},
		Usage:   "on (enable)",
	}
	offFlag = &cli.BoolFlag{
		Name:    "off",
		Aliases: []string{"d"},
		Usage:   "off (disable)",
	}
)

// flags for NGSI-LD
var (
	attrsDetailFlag = &cli.BoolFlag{
		Name:    "details",
		Aliases: []string{"d"},
		Usage:   "detailed attribute information",
	}
	typeDetailFlag = &cli.BoolFlag{
		Name:    "details",
		Aliases: []string{"d"},
		Usage:   "detailed entity type information (LD)",
	}
	ldContextsIDFlag = &cli.StringFlag{
		Name:    "id",
		Aliases: []string{"i"},
		Usage:   "jsonldContexts id",
	}
	ldContextsDetailsFlag = &cli.BoolFlag{
		Name:    "details",
		Aliases: []string{"d"},
		Usage:   "detailed jsonldContexts information",
	}
	ldContextsDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "jsonldContexts data",
	}
)

// Common flags for copy command
var (
	destinationFlag = &cli.StringFlag{
		Name:     "host2",
		Aliases:  []string{"d"},
		Required: true,
		Usage:    "host or alias",
		Value:    "",
	}
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
	context2Flag = &cli.StringFlag{
		Name:  "context2",
		Usage: "@context for destination",
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
	idPatternFlag = &cli.StringFlag{
		Name:  "idPattern",
		Usage: "idPattern",
	}
	typePatternFlag = &cli.StringFlag{
		Name:  "typePattern",
		Usage: "typePattern (v2)",
	}
	queryFlag = &cli.StringFlag{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "filtering by attribute value",
	}
	mqFlag = &cli.StringFlag{
		Name:    "mq",
		Aliases: []string{"m"},
		Usage:   "filtering by metadata (v2)",
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
	geopropertyFlag = &cli.StringFlag{
		Name:  "geoproperty",
		Usage: "geoproperty (LD)",
	}
	headersFlag = &cli.StringFlag{
		Name:  "headers",
		Usage: "headers (v2)",
	}
	qsFlag = &cli.StringFlag{
		Name:  "qs",
		Usage: "qs (v2)",
	}
	methodFlag = &cli.StringFlag{
		Name:  "method",
		Usage: "method (v2)",
	}
	payloadFlag = &cli.StringFlag{
		Name:  "payload",
		Usage: "payload (v2)",
	}
	exceptAttrsFlag = &cli.StringFlag{
		Name:  "exceptAttrs",
		Usage: "exceptAttrs (v2)",
	}
	attrsFormatFlag = &cli.StringFlag{
		Name:  "attrsFormat",
		Usage: "attrsFormat (v2)",
	}
	subscriptionIDFlag = &cli.StringFlag{
		Name:  "subscriptionId",
		Usage: "subscription id (LD)",
	}
	subscriptionNameFlag = &cli.StringFlag{
		Name:  "name",
		Usage: "subscription name (LD)",
	}
	entityIDFlag = &cli.StringFlag{
		Name:  "entityId",
		Usage: "entity id",
	}
	attrsFlag = &cli.StringFlag{
		Name:  "attrs",
		Usage: "attributes",
	}
	metadataFlag = &cli.StringFlag{
		Name:  "metadata",
		Usage: "metadata (v2)",
	}
	orderByFlag = &cli.StringFlag{
		Name:  "orderBy",
		Usage: "orderBy",
	}
	attrFlag = &cli.StringFlag{
		Name:  "attr",
		Usage: "attribute name",
	}
	attrRFlag = &cli.StringFlag{
		Name:     "attr",
		Usage:    "attribute name",
		Value:    "",
		Required: true,
	}
)

// Registration
var (
	providedIDFlag = &cli.StringFlag{
		Name:  "providedId",
		Usage: "providedId",
	}
	propertiesFlag = &cli.StringFlag{
		Name:  "properties",
		Usage: "properties (LD)",
	}
	relationshipsFlag = &cli.StringFlag{
		Name:  "relationships",
		Usage: "relationships (LD)",
	}
	providerFlag = &cli.StringFlag{
		Name:    "provider",
		Aliases: []string{"p"},
		Usage:   "Url of context provider/source",
	}
	legacyFlag = &cli.BoolFlag{
		Name:  "legacy",
		Usage: "legacy forwarding mode (V2)",
	}
	forwardingModeFlag = &cli.StringFlag{
		Name:  "forwardingMode",
		Usage: "forwarding mode (V2)",
	}
)

// Temporal
var (
	timeRelFlag = &cli.StringFlag{
		Name:  "timeRel",
		Usage: "temporal relationship (LD)",
	}
	timeAtFlag = &cli.StringFlag{
		Name:  "timeAt",
		Usage: "timeAt (LD)",
	}
	endTimeAtFlag = &cli.StringFlag{
		Name:  "endTimeAt",
		Usage: "endTimeAt (LD)",
	}
	timePropertyFlag = &cli.StringFlag{
		Name:  "timeProperty",
		Usage: "timeProperty (LD)",
	}
	geoPropertyFlag = &cli.StringFlag{
		Name:  "geoProperty",
		Usage: "geo property (LD)",
	}
	instanceIDFlag = &cli.StringFlag{
		Name:  "instanceId",
		Usage: "attribute instance id (LD)",
	}
	temporalValuesFlag = &cli.BoolFlag{
		Name:  "temporalValues",
		Usage: "temporal simplified representation of entity",
	}
	etsi10Flag = &cli.BoolFlag{
		Name:  "etsi10",
		Usage: "ETSI CIM 009 V1.0",
	}
	deleteAllFlag = &cli.BoolFlag{
		Name:  "deleteAll",
		Usage: "all attribute instances are deleted",
	}
	deleteDatasetID = &cli.StringFlag{
		Name:  "datasetId",
		Usage: "datasetId of the dataset to be deleted",
	}
)

// IoT Agent
var (
	servicesLimitFlag = &cli.Int64Flag{
		Name:  "limit",
		Usage: "maximum number of services",
	}
	servicesOffsetFlag = &cli.Int64Flag{
		Name:  "offset",
		Usage: "offset to skip a given number of elements at the beginning",
	}
	resourceFlag = &cli.StringFlag{
		Name:  "resource",
		Usage: "uri for the iotagent",
	}
	apikeyFlag = &cli.StringFlag{
		Name:  "apikey",
		Usage: "a key used for devices belonging to this service",
	}
	cbrokerFlag = &cli.StringFlag{
		Name:  "cbroker",
		Usage: "url of context broker or broker alias",
	}
	servicesDeviceFlag = &cli.BoolFlag{
		Name:  "device",
		Usage: "remove devices in service/subservice",
		Value: false,
	}
	servicesDataFlag = &cli.StringFlag{
		Name:  "data",
		Usage: "data body (payload)",
	}
	servicesTokenFlag = &cli.StringFlag{
		Name:  "token",
		Usage: "token obtained from the authentication system",
	}
	devicesLimit = &cli.Int64Flag{
		Name:  "limit",
		Usage: "maximum number of devices",
	}
	devicesOffset = &cli.Int64Flag{
		Name:  "offset",
		Usage: "offset to skip a given number of elements at the beginning",
	}
	devicesDetailed = &cli.StringFlag{
		Name:  "detailed",
		Usage: "all device information (on) or only name (off)",
		Value: "off",
	}
	devicesEntity = &cli.StringFlag{
		Name:  "entity",
		Usage: "get a device from entity name",
	}
	devicesProtocol = &cli.StringFlag{
		Name:  "protocol",
		Usage: "get devices with this protocol",
	}
	devicesDataFlag = &cli.StringFlag{
		Name:  "data",
		Usage: "data body (payload)",
	}
	devicesIDFlag = &cli.StringFlag{
		Name:  "id",
		Usage: "device id",
	}
)

// TIME SERIES
var (
	hLimitFlag = &cli.Int64Flag{
		Name:  "hLimit",
		Usage: "maximum number of data entries to retrieve",
	}
	hOffsetFlag = &cli.Int64Flag{
		Name:  "hOffset",
		Usage: "offset to be applied to data entries to be retrieved",
	}
	lastNFlag = &cli.Int64Flag{
		Name:  "lastN",
		Usage: "number of data entries to retrieve since the final date backwards",
	}
	aggrMethodFlag = &cli.StringFlag{
		Name:  "aggrMethod",
		Usage: "aggregation method (max, min, sum, sum, occur)",
	}
	aggrPeriodFlag = &cli.StringFlag{
		Name:  "aggrPeriod",
		Usage: "aggregation period or resolution of the aggregated data to be retrieved",
	}
	fromDateFlag = &cli.StringFlag{
		Name:  "fromDate",
		Usage: "starting date from which data should be retrieved",
	}
	toDateFlag = &cli.StringFlag{
		Name:  "toDate",
		Usage: "final date until which data should be retrieved",
	}
)

// TIME SERIES (quantumleap)
var (
	dropTableFlag = &cli.BoolFlag{
		Name:  "dropTable",
		Usage: "drop the table storing an entity type",
	}
	sameTypeFlag = &cli.BoolFlag{
		Name:  "sameType",
		Usage: "same type",
	}
	nTypesFlag = &cli.BoolFlag{
		Name:  "nTypes",
		Usage: "nTypes",
	}
	valueFlag = &cli.BoolFlag{
		Name:  "value",
		Usage: "values only",
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
	skipForwardingFlag = &cli.BoolFlag{
		Name:  "skipForwarding",
		Usage: "skip forwarding to CPrs (v2)",
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
		Name:    "uri",
		Aliases: []string{"u"},
		Usage:   "uri/url to be invoked when a notification is generated",
	}
	acceptFlag = &cli.StringFlag{
		Name:  "accept",
		Usage: "accept header (json or ld+json)",
	}
	expiresFlag = &cli.BoolFlag{
		Name:    "expires",
		Aliases: []string{"e"},
		Usage:   "expires",
	}
	revokeFlag = &cli.BoolFlag{
		Name:    "revoke",
		Aliases: []string{"r"},
		Usage:   "revoke token",
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
	prettyFlag = &cli.BoolFlag{
		Name:    "pretty",
		Aliases: []string{"P"},
		Value:   false,
		Usage:   "pretty format",
	}
	skipInitialNotificationFlag = &cli.BoolFlag{
		Name:  "skipInitialNotification",
		Usage: "skipInitialNotification",
	}
	localTimeFlag = &cli.BoolFlag{
		Name:  "localTime",
		Usage: "localTime",
	}
	statusFlag = &cli.StringFlag{
		Name:  "status",
		Usage: "status",
	}
	timeIntervalFlag = &cli.Int64Flag{
		Name:  "timeInterval",
		Usage: "time interval (LD)",
	}
	csfFlag = &cli.StringFlag{
		Name:  "csf",
		Usage: "context source filter (LD)",
	}
	activeFlag = &cli.BoolFlag{
		Name:  "active",
		Usage: "active (LD)",
	}
	inActiveFlag = &cli.BoolFlag{
		Name:  "inactive",
		Usage: "inactive (LD)",
	}
	descriptionFlag = &cli.StringFlag{
		Name:  "description",
		Usage: "description",
	}
	wAttrsFlag = &cli.StringFlag{
		Name:  "wAttrs",
		Usage: "watched attributes",
	}
	nAttrsFlag = &cli.StringFlag{
		Name:  "nAttrs",
		Usage: "attributes to be notified",
	}
	getFlag = &cli.BoolFlag{
		Name:   "get",
		Usage:  "get (v2)",
		Hidden: true,
	}
	notifyURLFlag = &cli.StringFlag{
		Name:   "url",
		Usage:  "url to be invoked when a notification is generated (v2)",
		Hidden: true,
	}
	singleLineFlag = &cli.BoolFlag{
		Name:    "singleLine",
		Aliases: []string{"1"},
		Usage:   "list one file per line",
		Value:   false,
	}
)

// flag for server config
var (
	serverHost2Flag = &cli.StringFlag{
		Name:  "serverHost",
		Usage: "specify server host",
	}
	serverTypeFlag = &cli.StringFlag{
		Name:  "serverType",
		Usage: "serverType (comet, ql)",
	}
	allServersFlag = &cli.BoolFlag{
		Name:   "all",
		Usage:  "print all servers",
		Hidden: true,
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
	brokerTypeFlag = &cli.StringFlag{
		Name:  "brokerType",
		Usage: "specify NGSI-LD broker type: orion-ld, scorpio or stellio",
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
	tokenScopeFlag = &cli.StringFlag{
		Name:  "tokenScope",
		Usage: "specify scope for token",
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
	headerNameFlag = &cli.StringFlag{
		Name:  "headerName",
		Usage: "specify header name for apikey",
	}
	headerValueFlag = &cli.StringFlag{
		Name:  "headerValue",
		Usage: "specify header value for apikey",
	}
	headerEnvValueFlag = &cli.StringFlag{
		Name:  "headerEnvValue",
		Usage: "specify name of environment variable for apikey",
	}
)

// flag for server
var (
	serverHostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Value:   "0.0.0.0",
		Usage:   "host for server",
	}
	serverPortFlag = &cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "3000",
		Usage:   "port for server",
	}
	serverURLFlag = &cli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Value:   "/",
		Usage:   "url for server",
	}
	serverHTTPSFlag = &cli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	serverKeyFlag = &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	serverCertFlag = &cli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
)

// flag for receiver
var (
	receiverHostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Value:   "0.0.0.0",
		Usage:   "host for receiver",
	}
	receiverPortFlag = &cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1028",
		Usage:   "port for receiver",
	}
	receiverURLFlag = &cli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Value:   "/",
		Usage:   "url for receiver",
	}
	receiverHTTPSFlag = &cli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	receiverKeyFlag = &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	receiverCertFlag = &cli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	headerFlag = &cli.BoolFlag{
		Name:  "header",
		Usage: "print receive header",
	}
)

// flag for registration proxy
var (
	regProxyHostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Usage:   "context broker or csource host",
	}
	regProxyRhostFlag = &cli.StringFlag{
		Name:  "rhost",
		Value: "0.0.0.0",
		Usage: "host for registration proxy",
	}
	regProxyPortFlag = &cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1028",
		Usage:   "port for registration proxy",
	}
	regProxyURLFlag = &cli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Value:   "/",
		Usage:   "url for registration proxy",
	}
	regProxyHTTPSFlag = &cli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	regProxyKeyFlag = &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	regProxyCertFlag = &cli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	regProxyReplaceTenantFlag = &cli.StringFlag{
		Name:  "replaceService",
		Usage: "replace FIWARE-Serivce",
	}
	regProxyReplaceScopeFlag = &cli.StringFlag{
		Name:  "replacePath",
		Usage: "replace FIWARE-SerivcePath",
	}
	regProxyAddScopeFlag = &cli.StringFlag{
		Name:  "addPath",
		Usage: "add path to FIWARE-SerivcePath",
	}
	regProxyReplaceURLFlag = &cli.StringFlag{
		Name:  "replaceURL",
		Usage: "replace URL of forwarding destination",
	}
	regProxyRegProxyHostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Usage:   "regproxy host",
	}
	regProxyVerboseFlag = &cli.StringFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "verbose log (on/off)",
	}
)

// flag for tokenproxy
var (
	tokenProxyHostFlag = &cli.StringFlag{
		Name:  "host",
		Value: "0.0.0.0",
		Usage: "host for tokenproxy",
	}
	tokenProxyPortFlag = &cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1029",
		Usage:   "port for tokenproxy",
	}
	tokenProxyHTTPSFlag = &cli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	tokenProxyKeyFlag = &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	tokenProxyCertFlag = &cli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	tokenProxyIdmHostTenantFlag = &cli.StringFlag{
		Name:  "idmHost",
		Usage: "host for Keyrock",
	}
	tokenProxyClientIDFlag = &cli.StringFlag{
		Name:    "clientId",
		Aliases: []string{"I"},
		Usage:   "specify client id for Keyrock",
	}
	tokeProxyClientSecretFlag = &cli.StringFlag{
		Name:    "clientSecret",
		Aliases: []string{"S"},
		Usage:   "specify client secret for Keyrock",
	}
	tokeProxyTokenProxyHostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Usage:   "specify tokenproxy server",
	}
)

// flag for queryproxy
var (
	queryProxyHostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Usage:   "context broker",
	}
	getPorxyReplaceURLFlag = &cli.StringFlag{
		Name:    "replaceURL",
		Aliases: []string{"u"},
		Usage:   "replace URL",
		Value:   "/v2/ex/entities",
	}
	queryProxyGHostFlag = &cli.StringFlag{
		Name:  "qhost",
		Value: "0.0.0.0",
		Usage: "host for queryproxy",
	}
	queryProxyPortFlag = &cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "1030",
		Usage:   "port for queryproxy",
	}
	queryProxyHTTPSFlag = &cli.BoolFlag{
		Name:    "https",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "start in https",
	}
	queryProxyKeyFlag = &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "key file (only needed if https is enabled)",
	}
	queryProxyCertFlag = &cli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "cert file (only needed if https is enabled)",
	}
	queryProxyQueryProxyHostFlag = &cli.StringFlag{
		Name:    "host",
		Aliases: []string{"h"},
		Usage:   "specify queryproxy server",
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

var (
	levelFlag = &cli.StringFlag{
		Name:    "level",
		Aliases: []string{"l"},
		Usage:   "specify log level",
	}
	deleteFlag = &cli.BoolFlag{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete",
	}
	setFlag = &cli.BoolFlag{
		Name:    "set",
		Aliases: []string{"s"},
		Usage:   "set",
	}
	resetFlag = &cli.BoolFlag{
		Name:    "reset",
		Aliases: []string{"r"},
		Usage:   "reset",
	}
	loggingFlag = &cli.BoolFlag{
		Name:    "logging",
		Aliases: []string{"L"},
		Usage:   "logging",
	}
)

// PERSEO FE
var (
	perseoRulesNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "rule name",
	}
	perseoRulesDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "rule data",
	}
	perseoRulesLimitFlag = &cli.Int64Flag{
		Name:  "limit",
		Usage: "maximum number of rules",
	}
	perseoRulesOffsetFlag = &cli.Int64Flag{
		Name:  "offset",
		Usage: "offset to skip a given number of rules at the beginning",
	}
	perseoRulesRaw = &cli.BoolFlag{
		Name:  "raw",
		Usage: "print raw data",
	}
	perseoRulesCount = &cli.BoolFlag{
		Name:  "count",
		Usage: "print number of rules",
	}
)

// Keyrock Common
var (
	keyrockDescriptionFlag = &cli.StringFlag{
		Name:    "description",
		Aliases: []string{"d"},
		Usage:   "description",
	}
	keyrockWebsiteFlag = &cli.StringFlag{
		Name:    "website",
		Aliases: []string{"w"},
		Usage:   "website",
	}
	keyrcokOrganizationRoleIDFlag = &cli.StringFlag{
		Name:    "orid",
		Aliases: []string{"c"},
		Usage:   "organization role id",
	}
)

// Keyrock users
var (
	keyrockUserIDFlag = &cli.StringFlag{
		Name:    "uid",
		Aliases: []string{"i"},
		Usage:   "user id",
	}
	keyrockUserNameFlag = &cli.StringFlag{
		Name:    "username",
		Aliases: []string{"u"},
		Usage:   "user name",
	}
	keyrockPasswordFlag = &cli.StringFlag{
		Name:    "password",
		Aliases: []string{"p"},
		Usage:   "password",
	}
	keyrockEmailFlag = &cli.StringFlag{
		Name:    "email",
		Aliases: []string{"e"},
		Usage:   "email",
	}
	keyrockGravatarFlag = &cli.StringFlag{
		Name:    "gravatar",
		Aliases: []string{"g"},
		Usage:   "gravatar (true or false)",
	}
	keyrockExtraFlag = &cli.StringFlag{
		Name:    "extra",
		Aliases: []string{"E"},
		Usage:   "extra information",
	}
)

// Keyrock applications
var (
	keyrockApplicationIDFlag = &cli.StringFlag{
		Name:    "aid",
		Aliases: []string{"i"},
		Usage:   "application id",
	}
	keyrcokApplicationDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "application data",
	}
	keyrcokApplicationNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "application name",
	}
	keyrcokApplicationDescriptionFlag = &cli.StringFlag{
		Name:    "description",
		Aliases: []string{"D"},
		Usage:   "description",
	}
	keyrcokApplicationRedirectURIFlag = &cli.StringFlag{
		Name:    "redirectUri",
		Aliases: []string{"R"},
		Usage:   "redirect uri",
	}
	keyrcokApplicationRedirectSignOutURIFlag = &cli.StringFlag{
		Name:    "redirectSignOutUri",
		Aliases: []string{"S"},
		Usage:   "redirect redirect sign out uri",
	}
	keyrcokApplicationURLFlag = &cli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Usage:   "url",
	}
	keyrcokApplicationGrantTypeFlag = &cli.StringFlag{
		Name:    "grantType",
		Aliases: []string{"g"},
		Usage:   "grant type",
	}
	keyrcokApplicationTokenTypesFlag = &cli.StringFlag{
		Name:    "tokenTypes",
		Aliases: []string{"t"},
		Usage:   "token types",
	}
	keyrcokApplicationResponseTypeFlag = &cli.StringFlag{
		Name:    "responseType",
		Aliases: []string{"r"},
		Usage:   "response type",
	}
	keyrcokApplicationClientTypeFlag = &cli.StringFlag{
		Name:    "clientType",
		Aliases: []string{"c"},
		Usage:   "client type",
	}
)

// Keyrock organizations
var (
	keyrockOrganizationIDFlag = &cli.StringFlag{
		Name:    "oid",
		Aliases: []string{"o"},
		Usage:   "organization id",
	}
	keyrockOrganizationNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "organization name",
	}
)

// Keyrock roles
var (
	keyrockRolesIDFlag = &cli.StringFlag{
		Name:    "rid",
		Aliases: []string{"r"},
		Usage:   "role id",
	}
	keyrockRoleDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "role data",
	}
	keyrockRoleNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "role name",
	}
)

// Keyrock permissions
var (
	keyrockPermissionIDFlag = &cli.StringFlag{
		Name:    "pid",
		Aliases: []string{"p"},
		Usage:   "permission id",
	}
	keyrcokPermissionDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "permissionrole data",
	}
	keyrockPermissionNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "permission name",
	}
	keyrockPermissionDescriptionFlag = &cli.StringFlag{
		Name:    "description",
		Aliases: []string{"D"},
		Usage:   "description",
	}
	keyrockPermissionActionFlag = &cli.StringFlag{
		Name:    "action",
		Aliases: []string{"a"},
		Usage:   "action",
	}
	keyrockPermissionResourceFlag = &cli.StringFlag{
		Name:    "resource",
		Aliases: []string{"r"},
		Usage:   "resoruce",
	}
)

// Keyrock Iot Agents
var (
	keyrockIotAgentsIDFlag = &cli.StringFlag{
		Name:    "iid",
		Aliases: []string{"i"},
		Usage:   "IoT Agent id",
	}
)

// Keyrock Trusted application
var (
	keyrockTrustedIDFlag = &cli.StringFlag{
		Name:    "tid",
		Aliases: []string{"t"},
		Usage:   "trusted application id",
	}
)

// Cygnus namemappings

var (
	cygnusNamemappingsDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "name mapping data",
	}
)

// Cygnus groupingrules

var (
	cygnusGroupingrulesIDFlag = &cli.StringFlag{
		Name:    "id",
		Aliases: []string{"i"},
		Usage:   "grouping rule id",
	}
	cygnusGroupingrulesDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "grouping rule data",
	}
)

// Cygnus loggers
var (
	cygnusLoggersNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "logger name",
	}
	cygnusLoggersDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "logger information",
	}
	cygnusLoggersTransientFlag = &cli.BoolFlag{
		Name:    "transient",
		Aliases: []string{"t"},
		Usage:   "true, retrieving from memory, or false, retrieving from file",
		Value:   false,
	}
)

// Cygnus appenders
var (
	cygnusAppendersNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "appender name",
	}
	cygnusAppendersDataFlag = &cli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "appender information",
	}
	cygnusAppendersTransientFlag = &cli.BoolFlag{
		Name:    "transient",
		Aliases: []string{"t"},
		Usage:   "true, retrieving from memory, or false, retrieving from file",
		Value:   false,
	}
)

// WireCloud
var (
	wireCloudWorkspaceIdFlag = &cli.StringFlag{
		Name:    "wid",
		Aliases: []string{"w"},
		Usage:   "workspace id",
	}
	wireCloudTabIdFlag = &cli.StringFlag{
		Name:    "tid",
		Aliases: []string{"t"},
		Usage:   "tab id",
	}
	wireCloudWidgetFlag = &cli.BoolFlag{
		Name:  "widget",
		Usage: "filtering widget",
	}
	wireCloudOperatorFlag = &cli.BoolFlag{
		Name:  "operator",
		Usage: "filtering operator",
	}
	wireCloudMashupFlag = &cli.BoolFlag{
		Name:  "mashup",
		Usage: "filtering mashup",
	}
	wireCloudVenderFlag = &cli.StringFlag{
		Name:    "vender",
		Aliases: []string{"v"},
		Usage:   "vender name of mashable application component",
	}
	wireCloudNameFlag = &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "name of mashable application component",
	}
	wireCloudVersionFlag = &cli.StringFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "version of mashable application component",
	}
	wireCloudFileFlag = &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Usage:   "mashable application component file",
	}
	wireCloudPublicFlag = &cli.BoolFlag{
		Name:    "public",
		Aliases: []string{"p"},
		Usage:   "install mashable application component as public",
	}
	wireCloudOverwriteFlag = &cli.BoolFlag{
		Name:    "overwrite",
		Aliases: []string{"o"},
		Usage:   "overwrite mashable application component",
	}
	wireCloudWidgetsFlag = &cli.BoolFlag{
		Name:    "widgets",
		Aliases: []string{"W"},
		Usage:   "list widgets",
	}
	wireCloudOperatorsFlag = &cli.BoolFlag{
		Name:    "operators",
		Aliases: []string{"o"},
		Usage:   "list operators",
	}
	wireCloudTabsFlag = &cli.BoolFlag{
		Name:    "tabs",
		Aliases: []string{"t"},
		Usage:   "list tabs",
	}
	wireCloudUsersFlag = &cli.BoolFlag{
		Name:    "users",
		Aliases: []string{"u"},
		Usage:   "list users",
	}
)
