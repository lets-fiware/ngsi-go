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

package ngsicmd

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
)

// id

var (
	idEntityRFlag = &ngsicli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "entity id",
		Required: true,
	}
	idEntityFlag = &ngsicli.StringFlag{
		Name:    "id",
		Aliases: []string{"i"},
		Usage:   "entity id",
	}
	idSubscriptionRFlag = &ngsicli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "subscription id",
		Required: true,
	}
	idRegistrationRFlag = &ngsicli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "registration id",
		Required: true,
	}
	idTemporalEntityRFlag = &ngsicli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "temporal entity id",
		Required: true,
	}
	idTemporalEntityFlag = &ngsicli.StringFlag{
		Name:    "id",
		Aliases: []string{"i"},
		Usage:   "temporal entity id",
	}
)

// flags for NGSI-LD
var (
	linkFlag = &ngsicli.StringFlag{
		Name:    "link",
		Aliases: []string{"L"},
		Usage:   "@context `VALUE` (LD)",
	}
	contextFlag = &ngsicli.StringFlag{
		Name:    "context",
		Aliases: []string{"C"},
		Usage:   "@context `VLAUE` (LD)",
	}
	acceptJSONFlag = &ngsicli.BoolFlag{
		Name:  "acceptJson",
		Usage: "set accecpt header to application/json (LD)",
		Value: false,
	}
	acceptGeoJSONFlag = &ngsicli.BoolFlag{
		Name:  "acceptGeoJson",
		Usage: "set accecpt header to application/geo+json (LD)",
		Value: false,
	}
	attrsDetailFlag = &ngsicli.BoolFlag{
		Name:    "details",
		Aliases: []string{"d"},
		Usage:   "detailed attribute information (LD)",
	}
	typeDetailFlag = &ngsicli.BoolFlag{
		Name:    "details",
		Aliases: []string{"d"},
		Usage:   "detailed entity type information (LD)",
	}
	ldContextsIDRFlag = &ngsicli.StringFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "jsonldContexts id (LD)",
		Required: true,
	}
	ldContextsDetailsFlag = &ngsicli.BoolFlag{
		Name:    "details",
		Aliases: []string{"d"},
		Usage:   "detailed jsonldContexts information (LD)",
	}
	ldContextsDataFlag = &ngsicli.StringFlag{
		Name:     "data",
		Aliases:  []string{"d"},
		Usage:    "jsonldContexts data (LD)",
		Required: true,
	}
)

// flags for NGSI API
var (
	typeFlag = &ngsicli.StringFlag{
		Name:    "type",
		Aliases: []string{"t"},
		Usage:   "entity type",
	}
	idPatternFlag = &ngsicli.StringFlag{
		Name:  "idPattern",
		Usage: "idPattern",
	}
	typePatternFlag = &ngsicli.StringFlag{
		Name:  "typePattern",
		Usage: "typePattern (v2)",
	}
	queryFlag = &ngsicli.StringFlag{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "filtering by attribute value",
	}
	mqFlag = &ngsicli.StringFlag{
		Name:    "mq",
		Aliases: []string{"m"},
		Usage:   "filtering by metadata (v2)",
	}
	georelFlag = &ngsicli.StringFlag{
		Name:  "georel",
		Usage: "georel",
	}
	geometryFlag = &ngsicli.StringFlag{
		Name:  "geometry",
		Usage: "geometry",
	}
	coordsFlag = &ngsicli.StringFlag{
		Name:  "coords",
		Usage: "coords",
	}
	geopropertyFlag = &ngsicli.StringFlag{
		Name:  "geoproperty",
		Usage: "geoproperty (LD)",
	}
	headersFlag = &ngsicli.StringFlag{
		Name:  "headers",
		Usage: "headers (v2)",
	}
	qsFlag = &ngsicli.StringFlag{
		Name:  "qs",
		Usage: "qs (v2)",
	}
	methodFlag = &ngsicli.StringFlag{
		Name:  "method",
		Usage: "method (v2)",
	}
	payloadFlag = &ngsicli.StringFlag{
		Name:  "payload",
		Usage: "payload (v2)",
	}
	exceptAttrsFlag = &ngsicli.StringFlag{
		Name:  "exceptAttrs",
		Usage: "exceptAttrs (v2)",
	}
	attrsFormatFlag = &ngsicli.StringFlag{
		Name:  "attrsFormat",
		Usage: "attrsFormat (v2)",
	}
	subscriptionIDFlag = &ngsicli.StringFlag{
		Name:  "subscriptionId",
		Usage: "subscription id (LD)",
	}
	subscriptionNameFlag = &ngsicli.StringFlag{
		Name:  "name",
		Usage: "subscription name (LD)",
	}
	entityIDFlag = &ngsicli.StringFlag{
		Name:  "entityId",
		Usage: "entity id",
	}
	attrsFlag = &ngsicli.StringFlag{
		Name:  "attrs",
		Usage: "attributes",
	}
	metadataFlag = &ngsicli.StringFlag{
		Name:  "metadata",
		Usage: "metadata (v2)",
	}
	orderByFlag = &ngsicli.StringFlag{
		Name:  "orderBy",
		Usage: "orderBy",
	}
	attrFlag = &ngsicli.StringFlag{
		Name:  "attr",
		Usage: "attribute name",
	}
	attrRFlag = &ngsicli.StringFlag{
		Name:     "attr",
		Usage:    "attribute name",
		Required: true,
	}
	itemsFlag = &ngsicli.StringFlag{
		Name:    "items",
		Aliases: []string{"i"},
		Usage:   "itmes",
	}
)

// Attribute data
var (
	attrDataFlag = &ngsicli.StringFlag{
		Name:       "data",
		Usage:      "attribute data",
		Aliases:    []string{"d"},
		ValueEmpty: true,
	}
)

// Attributes data
var (
	attrsDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Usage:   "attributes data",
		Aliases: []string{"d"},
	}
)

// Temporal attribute data
var (
	tattrDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Usage:   "attribute instance of temporal entity",
		Aliases: []string{"d"},
	}
)

// Temporal attributes data
var (
	tattrsDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Usage:   "attributes instance of temporal entity",
		Aliases: []string{"d"},
	}
)

// Entity data
var (
	/*
		entityDataFlag = &ngsicli.StringFlag{
			Name:    "data",
			Usage:   "entity data",
			Aliases: []string{"d"},
		}
	*/
	entityDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Usage:    "entity data",
		Aliases:  []string{"d"},
		Required: true,
	}
)

// Entities
var (
	entitiesDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Usage:   "entities data",
		Aliases: []string{"d"},
	}
	entitiesDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Usage:    "entities data",
		Aliases:  []string{"d"},
		Required: true,
	}
)

// Temporal entity

var (
	tentityDataRFlag = &ngsicli.StringFlag{
		Name:     "data",
		Usage:    "temporal entity data",
		Aliases:  []string{"d"},
		Required: true,
	}
)

// Subscription
var (
	subscriptionDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Usage:   "subscription data",
		Aliases: []string{"d"},
	}
)

// Registration
var (
	registrationDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Usage:   "registration data",
		Aliases: []string{"d"},
	}
	providedIDFlag = &ngsicli.StringFlag{
		Name:  "providedId",
		Usage: "providedId",
	}
	propertiesFlag = &ngsicli.StringFlag{
		Name:  "properties",
		Usage: "properties (LD)",
	}
	relationshipsFlag = &ngsicli.StringFlag{
		Name:  "relationships",
		Usage: "relationships (LD)",
	}
	providerFlag = &ngsicli.StringFlag{
		Name:    "provider",
		Aliases: []string{"p"},
		Usage:   "Url of context provider/source",
	}
	legacyFlag = &ngsicli.BoolFlag{
		Name:  "legacy",
		Usage: "legacy forwarding mode (V2)",
	}
	forwardingModeFlag = &ngsicli.StringFlag{
		Name:  "forwardingMode",
		Usage: "forwarding mode (V2)",
	}
)

// Temporal
var (
	timeRelFlag = &ngsicli.StringFlag{
		Name:  "timeRel",
		Usage: "temporal relationship (LD)",
	}
	timeAtFlag = &ngsicli.StringFlag{
		Name:  "timeAt",
		Usage: "timeAt (LD)",
	}
	endTimeAtFlag = &ngsicli.StringFlag{
		Name:  "endTimeAt",
		Usage: "endTimeAt (LD)",
	}
	timePropertyFlag = &ngsicli.StringFlag{
		Name:  "timeProperty",
		Usage: "timeProperty (LD)",
	}
	geoPropertyFlag = &ngsicli.StringFlag{
		Name:  "geoProperty",
		Usage: "geo property (LD)",
	}
	instanceIDFlag = &ngsicli.StringFlag{
		Name:  "instanceId",
		Usage: "attribute instance id (LD)",
	}
	instanceIDRFlag = &ngsicli.StringFlag{
		Name:     "instanceId",
		Usage:    "attribute instance id (LD)",
		Required: true,
	}
	temporalValuesFlag = &ngsicli.BoolFlag{
		Name:  "temporalValues",
		Usage: "temporal simplified representation of entity",
	}
	etsi10Flag = &ngsicli.BoolFlag{
		Name:  "etsi10",
		Usage: "ETSI CIM 009 V1.0",
	}
	deleteAllFlag = &ngsicli.BoolFlag{
		Name:  "deleteAll",
		Usage: "all attribute instances are deleted",
	}
	deleteDatasetID = &ngsicli.StringFlag{
		Name:  "datasetId",
		Usage: "datasetId of the dataset to be deleted",
	}
)

// flags for Temporal
var (
	fromDateFlag = &ngsicli.StringFlag{
		Name:  "fromDate",
		Usage: "starting date from which data should be retrieved",
	}
	toDateFlag = &ngsicli.StringFlag{
		Name:  "toDate",
		Usage: "final date until which data should be retrieved",
	}
	lastNFlag = &ngsicli.Int64Flag{
		Name:  "lastN",
		Usage: "number of data entries to retrieve since the final date backwards",
	}
)

// flags for options
var (
	countFlag = &ngsicli.BoolFlag{
		Name:    "count",
		Aliases: []string{"C"},
		Usage:   "count",
	}
	keyValuesFlag = &ngsicli.BoolFlag{
		Name:    "keyValues",
		Aliases: []string{"K"},
		Usage:   "keyValues",
	}
	valuesFlag = &ngsicli.BoolFlag{
		Name:    "values",
		Aliases: []string{"V"},
		Usage:   "values",
	}
	uniqueFlag = &ngsicli.BoolFlag{
		Name:    "unique",
		Aliases: []string{"U"},
		Usage:   "unique",
	}
	upsertFlag = &ngsicli.BoolFlag{
		Name:  "upsert",
		Usage: "upsert",
	}
	appendFlag = &ngsicli.BoolFlag{
		Name:    "append",
		Aliases: []string{"a"},
		Usage:   "append",
	}
	noOverwriteFlag = &ngsicli.BoolFlag{
		Name:    "noOverwrite",
		Aliases: []string{"n"},
		Usage:   "noOverwrite",
	}
	replaceFlag = &ngsicli.BoolFlag{
		Name:    "replace",
		Aliases: []string{"r"},
		Usage:   "replace",
	}
	updateFlag = &ngsicli.BoolFlag{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "update",
	}
	sysAttrsFlag = &ngsicli.BoolFlag{
		Name:    "sysAttrs",
		Aliases: []string{"S"},
		Usage:   "sysAttrs",
	}
	linesFlag = &ngsicli.BoolFlag{
		Name:    "lines",
		Aliases: []string{"1"},
		Usage:   "lines",
	}
	skipForwardingFlag = &ngsicli.BoolFlag{
		Name:  "skipForwarding",
		Usage: "skip forwarding to CPrs (v2)",
	}
	rawFlag = &ngsicli.BoolFlag{
		Name:  "raw",
		Usage: "handle raw data",
	}
)

var (
	uriFlag = &ngsicli.StringFlag{
		Name:    "uri",
		Aliases: []string{"u"},
		Usage:   "uri/url to be invoked when a notification is generated",
	}
	acceptFlag = &ngsicli.StringFlag{
		Name:  "accept",
		Usage: "accept header (json or ld+json)",
	}

	expiresSFlag = &ngsicli.StringFlag{
		Name:    "expires",
		Aliases: []string{"e"},
		Usage:   "expires",
	}
	throttlingFlag = &ngsicli.Int64Flag{
		Name:  "throttling",
		Usage: "throttling",
	}
	skipInitialNotificationFlag = &ngsicli.BoolFlag{
		Name:  "skipInitialNotification",
		Usage: "skipInitialNotification",
	}
	localTimeFlag = &ngsicli.BoolFlag{
		Name:  "localTime",
		Usage: "localTime",
	}
	statusFlag = &ngsicli.StringFlag{
		Name:  "status",
		Usage: "status",
	}
	timeIntervalFlag = &ngsicli.Int64Flag{
		Name:  "timeInterval",
		Usage: "time interval (LD)",
	}
	csfFlag = &ngsicli.StringFlag{
		Name:  "csf",
		Usage: "context source filter (LD)",
	}
	activeFlag = &ngsicli.BoolFlag{
		Name:  "active",
		Usage: "active (LD)",
	}
	inActiveFlag = &ngsicli.BoolFlag{
		Name:  "inactive",
		Usage: "inactive (LD)",
	}
	descriptionFlag = &ngsicli.StringFlag{
		Name:  "description",
		Usage: "description",
	}
	wAttrsFlag = &ngsicli.StringFlag{
		Name:  "wAttrs",
		Usage: "watched attributes",
	}
	nAttrsFlag = &ngsicli.StringFlag{
		Name:  "nAttrs",
		Usage: "attributes to be notified",
	}
	getFlag = &ngsicli.BoolFlag{
		Name:   "get",
		Usage:  "get (v2)",
		Hidden: true,
	}
	notifyURLFlag = &ngsicli.StringFlag{
		Name:   "url",
		Usage:  "url to be invoked when a notification is generated (v2)",
		Hidden: true,
	}
)
