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
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return NgsiApp
}

var NgsiApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "ngsi command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&CountCmd,
		&LsCmd,
		&AppendCmd,
		&CreateCmd,
		&DeleteCmd,
		&GetCmd,
		&ListCmd,
		&ReplaceCmd,
		&UpdateCmd,
		&UpsertCmd,
		&TemplateCmd,
	},
}

var CountCmd = ngsicli.Command{
	Name:     "wc",
	Usage:    "print number of entities, subscriptions, registrations or types",
	Category: "CONVENIENCE",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "entities",
			Usage:      "print number of entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				typeFlag,
				linkFlag,
				skipForwardingFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return entitiesCount(c, ngsi, client)
			},
		},
		{
			Name:       "subscriptions",
			Usage:      "print number of subscriptions",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return subscriptionsCount(c, ngsi, client)
			},
		},
		{
			Name:       "registrations",
			Usage:      "print number of registrations",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return registrationsCount(c, ngsi, client)
			},
		},
		{
			Name:       "types",
			Usage:      "print number of types",
			ServerList: []string{"brokerv2", "brokerld"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return typesCount(c, ngsi, client)
			},
		},
	},
}

var LsCmd = ngsicli.Command{
	Name:       "ls",
	Usage:      "list entities",
	Category:   "CONVENIENCE",
	ServerList: []string{"brokerv2", "brokerld"},
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
		idEntityFlag,
		typeFlag,
		idPatternFlag,
		typePatternFlag,
		queryFlag,
		mqFlag,
		georelFlag,
		geometryFlag,
		coordsFlag,
		attrsFlag,
		metadataFlag,
		orderByFlag,
		countFlag,
		keyValuesFlag,
		valuesFlag,
		uniqueFlag,
		skipForwardingFlag,
		linkFlag,
		ngsicli.VerboseFlag,
		linesFlag,
		ngsicli.PrettyFlag,
		ngsicli.SafeStringFlag,
	},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return entitiesList(c, ngsi, client)
	},
}

var AppendCmd = ngsicli.Command{
	Name:     "append",
	Usage:    "append attributes",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attrs",
			Usage:      "append ttributes",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				typeFlag,
				keyValuesFlag,
				appendFlag,
				attrsDataFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attrsAppend(c, ngsi, client)
			},
		},
		{
			Name:       "tattrs",
			Usage:      "append attribute instance of temporal entity",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				idTemporalEntityRFlag,
				tattrsDataFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return troeAttrsAppend(c, ngsi, client)
			},
		},
	},
}

var CreateCmd = ngsicli.Command{
	Name:     "create",
	Usage:    "create entity(ies), subscription, registration or ldContext",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "entity",
			Usage:      "create entity",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				entityDataRFlag,
				keyValuesFlag,
				upsertFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return entityCreate(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "create entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				entitiesDataRFlag,
				keyValuesFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return batch(c, c.Ngsi, c.Client, "create")
			},
		},
		{
			Name:       "tentity",
			Usage:      "create temporal entity",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				tentityDataRFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return troeCreate(c, ngsi, client)
			},
		},
		{
			Name:       "subscription",
			Usage:      "create subscription",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				subscriptionDataFlag,
				skipInitialNotificationFlag,
				subscriptionIDFlag,
				subscriptionNameFlag,
				descriptionFlag,
				entityIDFlag,
				idPatternFlag,
				typeFlag,
				typePatternFlag,
				wAttrsFlag,
				timeIntervalFlag,
				queryFlag,
				mqFlag,
				geometryFlag,
				coordsFlag,
				georelFlag,
				geopropertyFlag,
				csfFlag,
				activeFlag,
				inActiveFlag,
				nAttrsFlag,
				keyValuesFlag,
				uriFlag,
				acceptFlag,
				expiresSFlag,
				throttlingFlag,
				timeRelFlag,
				timeAtFlag,
				endTimeAtFlag,
				timePropertyFlag,
				linkFlag,
				contextFlag,
				statusFlag,
				headersFlag,
				qsFlag,
				methodFlag,
				payloadFlag,
				metadataFlag,
				exceptAttrsFlag,
				attrsFormatFlag,
				ngsicli.SafeStringFlag,
				rawFlag,
				notifyURLFlag,
				getFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return subscriptionsCreate(c, ngsi, client)
			},
		},
		{
			Name:       "registration",
			Usage:      "create registration",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				registrationDataFlag,
				linkFlag,
				contextFlag,
				providedIDFlag,
				idPatternFlag,
				typeFlag,
				attrsFlag,
				providerFlag,
				descriptionFlag,
				legacyFlag,
				forwardingModeFlag,
				expiresSFlag,
				statusFlag,
				ngsicli.SafeStringFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return registrationsCreate(c, ngsi, client)
			},
		},
		{
			Name:       "ldContext",
			Usage:      "create jsonldContext",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				ldContextsDataFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return jsonldContextCreate(c, ngsi, client)
			},
		},
	},
}

var DeleteCmd = ngsicli.Command{
	Name:     "delete",
	Usage:    "delete entity(ies), attribute, subscription, registration or ldContext",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attr",
			Usage:      "delete attr",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				attrRFlag,
				typeFlag,
				linkFlag,
			},
			RequiredFlags: []string{"id", "attr"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attrDelete(c, ngsi, client)
			},
		},
		{
			Name:       "tattr",
			Usage:      "delete attr for temporal entity",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				idTemporalEntityRFlag,
				attrRFlag,
				deleteAllFlag,
				deleteDatasetID,
				instanceIDFlag,
				linkFlag,
			},
			RequiredFlags: []string{"id", "attr"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return troeAttrDelete(c, ngsi, client)
			},
		},
		{
			Name:       "entity",
			Usage:      "delete entity",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				typeFlag,
				linkFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return entityDelete(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "delete entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				entitiesDataRFlag,
				keyValuesFlag,
				linkFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return batch(c, c.Ngsi, c.Client, "delete")
			},
		},
		{
			Name:       "tentity",
			Usage:      "delete temporal entity",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				idTemporalEntityRFlag,
				linkFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return troeDelete(c, ngsi, client)
			},
		},
		{
			Name:       "subscription",
			Usage:      "delete subscription",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idSubscriptionRFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return subscriptionsDelete(c, ngsi, client)
			},
		},
		{
			Name:       "registration",
			Usage:      "delete registration",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idRegistrationRFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return registrationsDelete(c, ngsi, client)
			},
		},
		{
			Name:       "ldContext",
			Usage:      "delete jsonldContext",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				ldContextsIDRFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return jsonldContextDelete(c, ngsi, client)
			},
		},
	},
}

var GetCmd = ngsicli.Command{
	Name:     "get",
	Usage:    "get entity(ies), attribute(s), subscription, registration type or ldContext",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attr",
			Usage:      "get attr",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				attrRFlag,
				typeFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id", "attr"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attrRead(c, ngsi, client)
			},
		},
		{
			Name:       "attrs",
			Usage:      "get attrs",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				typeFlag,
				attrsFlag,
				metadataFlag,
				keyValuesFlag,
				valuesFlag,
				uniqueFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attrsRead(c, ngsi, client)
			},
		},
		{
			Name:       "entity",
			Usage:      "get entity",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				typeFlag,
				attrsFlag,
				keyValuesFlag,
				valuesFlag,
				uniqueFlag,
				sysAttrsFlag,
				linkFlag,
				acceptJSONFlag,
				acceptGeoJSONFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return entityRead(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "get entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				orderByFlag,
				countFlag,
				keyValuesFlag,
				valuesFlag,
				uniqueFlag,
				ngsicli.VerboseFlag,
				linesFlag,
				entitiesDataFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return opQuery(c, ngsi, client)
			},
		},
		{
			Name:       "tentity",
			Usage:      "get temporal entity",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				idTemporalEntityRFlag,
				attrsFlag,
				timePropertyFlag,
				fromDateFlag,
				toDateFlag,
				lastNFlag,
				linkFlag,
				temporalValuesFlag,
				sysAttrsFlag,
				acceptJSONFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
				etsi10Flag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return troeRead(c, ngsi, client)
			},
		},
		{
			Name:       "subscription",
			Usage:      "get subscription",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idSubscriptionRFlag,
				localTimeFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
				rawFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return subscriptionGet(c, ngsi, client)
			},
		},
		{
			Name:       "registration",
			Usage:      "get registration",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idRegistrationRFlag,
				localTimeFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return registrationsGet(c, ngsi, client)
			},
		},
		{
			Name:       "type",
			Usage:      "get type",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				typeFlag,
				ngsicli.PrettyFlag,
				linkFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return typeGet(c, ngsi, client)
			},
		},
		{
			Name:       "ldContext",
			Usage:      "get jsonldContext",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				ldContextsIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return jsonldContextGet(c, ngsi, client)
			},
		},
	},
}

var ListCmd = ngsicli.Command{
	Name:     "list",
	Usage:    "list types, attributes, entities, tentities, subscriptions or registrations",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attributes",
			Usage:      "list attributes",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				attrFlag,
				attrsDetailFlag,
				ngsicli.PrettyFlag,
				linkFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attributesList(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "list entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityFlag,
				typeFlag,
				idPatternFlag,
				typePatternFlag,
				queryFlag,
				mqFlag,
				georelFlag,
				geometryFlag,
				coordsFlag,
				attrsFlag,
				metadataFlag,
				orderByFlag,
				countFlag,
				keyValuesFlag,
				valuesFlag,
				uniqueFlag,
				skipForwardingFlag,
				linkFlag,
				acceptJSONFlag,
				acceptGeoJSONFlag,
				ngsicli.VerboseFlag,
				linesFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return entitiesList(c, ngsi, client)
			},
		},
		{
			Name:       "tentities",
			Usage:      "list temporal entities",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				idTemporalEntityFlag,
				typeFlag,
				idPatternFlag,
				attrsFlag,
				queryFlag,
				csfFlag,
				georelFlag,
				geometryFlag,
				coordsFlag,
				geoPropertyFlag,
				timePropertyFlag,
				fromDateFlag,
				toDateFlag,
				lastNFlag,
				temporalValuesFlag,
				sysAttrsFlag,
				linkFlag,
				acceptJSONFlag,
				ngsicli.VerboseFlag,
				linesFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
				etsi10Flag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return troeList(c, ngsi, client)
			},
		},
		{
			Name:       "types",
			Usage:      "list types",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				typeDetailFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
				linkFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return typesList(c, ngsi, client)
			},
		},
		{
			Name:       "subscriptions",
			Usage:      "list subscriptions",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.JsonFlag,
				statusFlag,
				localTimeFlag,
				queryFlag,
				itemsFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
				countFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return subscriptionsList(c, ngsi, client)
			},
		},
		{
			Name:       "registrations",
			Usage:      "list registrations",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.JsonFlag,
				localTimeFlag,
				ngsicli.PrettyFlag,
				ngsicli.SafeStringFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return registrationsList(c, ngsi, client)
			},
		},
		{
			Name:       "ldContexts",
			Usage:      "list jsonldContexts",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				ldContextsDetailsFlag,
				ngsicli.JsonFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return jsonldContextsList(c, ngsi, client)
			},
		},
	},
}

var ReplaceCmd = ngsicli.Command{
	Name:     "replace",
	Usage:    "replace entities or attributes",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
		linkFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attrs",
			Usage:      "replace attrs",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				typeFlag,
				attrsDataFlag,
				keyValuesFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attrsReplace(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "replace entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				entitiesDataRFlag,
				keyValuesFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return batch(c, c.Ngsi, c.Client, "replace")
			},
		},
	},
}

var UpdateCmd = ngsicli.Command{
	Name:     "update",
	Usage:    "update entities, attribute(s) or subscription",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "attr",
			Usage:      "update attr",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				attrRFlag,
				attrDataFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id", "attr"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attrUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "attrs",
			Usage:      "update attrs",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idEntityRFlag,
				typeFlag,
				keyValuesFlag,
				attrsDataFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return attrsUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "tattr",
			Usage:      "update attr instance of temporal entity",
			ServerList: []string{"brokerld"},
			Flags: []ngsicli.Flag{
				idTemporalEntityRFlag,
				attrRFlag,
				instanceIDRFlag,
				tattrDataFlag,
				linkFlag,
				contextFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"id", "attr", "instanceId"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return troeAttrUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "update entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				entitiesDataRFlag,
				keyValuesFlag,
				noOverwriteFlag,
				replaceFlag,
				linkFlag,
				contextFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return batch(c, c.Ngsi, c.Client, "update")
			},
		},
		{
			Name:       "subscription",
			Usage:      "update subscription",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				idSubscriptionRFlag,
				subscriptionDataFlag,
				skipInitialNotificationFlag,
				subscriptionIDFlag,
				subscriptionNameFlag,
				descriptionFlag,
				entityIDFlag,
				idPatternFlag,
				typeFlag,
				typePatternFlag,
				wAttrsFlag,
				timeIntervalFlag,
				queryFlag,
				mqFlag,
				geometryFlag,
				coordsFlag,
				georelFlag,
				geopropertyFlag,
				csfFlag,
				activeFlag,
				inActiveFlag,
				nAttrsFlag,
				keyValuesFlag,
				uriFlag,
				acceptFlag,
				expiresSFlag,
				throttlingFlag,
				timeRelFlag,
				timeAtFlag,
				endTimeAtFlag,
				timePropertyFlag,
				linkFlag,
				contextFlag,
				statusFlag,
				headersFlag,
				qsFlag,
				methodFlag,
				payloadFlag,
				metadataFlag,
				exceptAttrsFlag,
				attrsFormatFlag,
				ngsicli.SafeStringFlag,
				rawFlag,
				notifyURLFlag,
				getFlag,
			},
			RequiredFlags: []string{"id"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return subscriptionsUpdate(c, ngsi, client)
			},
		},
	},
}

var UpsertCmd = ngsicli.Command{
	Name:     "upsert",
	Usage:    "upsert entity or entities",
	Category: "NGSI",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "entity",
			Usage:      "upsert entity",
			ServerList: []string{"brokerv2"},
			Flags: []ngsicli.Flag{
				entityDataRFlag,
				keyValuesFlag,
				ngsicli.SafeStringFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return entityUpsert(c, ngsi, client)
			},
		},
		{
			Name:       "entities",
			Usage:      "upsert entities",
			ServerList: []string{"brokerv2", "brokerld"},
			Flags: []ngsicli.Flag{
				entitiesDataRFlag,
				replaceFlag,
				updateFlag,
				linkFlag,
				contextFlag,
			},
			RequiredFlags: []string{"data"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return batch(c, c.Ngsi, c.Client, "upsert")
			},
		},
	},
}

var TemplateCmd = ngsicli.Command{
	Name:     "template",
	Usage:    "create template of subscription or registration",
	Category: "CONVENIENCE",
	Flags: []ngsicli.Flag{
		ngsicli.HostFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.TenantFlag,
		ngsicli.ScopeFlag,
		linkFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:  "subscription",
			Usage: "create template of subscription",
			Flags: []ngsicli.Flag{
				ngsicli.NgsiTypeFlag,
				subscriptionDataFlag,
				subscriptionIDFlag,
				subscriptionNameFlag,
				descriptionFlag,
				entityIDFlag,
				idPatternFlag,
				typeFlag,
				typePatternFlag,
				wAttrsFlag,
				timeIntervalFlag,
				queryFlag,
				mqFlag,
				geometryFlag,
				coordsFlag,
				georelFlag,
				geopropertyFlag,
				csfFlag,
				activeFlag,
				inActiveFlag,
				nAttrsFlag,
				keyValuesFlag,
				uriFlag,
				acceptFlag,
				expiresSFlag,
				throttlingFlag,
				timeRelFlag,
				timeAtFlag,
				endTimeAtFlag,
				timePropertyFlag,
				contextFlag,
				statusFlag,
				headersFlag,
				qsFlag,
				methodFlag,
				payloadFlag,
				metadataFlag,
				exceptAttrsFlag,
				attrsFormatFlag,
				notifyURLFlag,
				getFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return subscriptionsTemplate(c, ngsi, client)
			},
		},
		{
			Name:  "registration",
			Usage: "create template of registration",
			Flags: []ngsicli.Flag{
				ngsicli.NgsiTypeFlag,
				registrationDataFlag,
				descriptionFlag,
				typeFlag,
				providedIDFlag,
				idPatternFlag,
				propertiesFlag,
				relationshipsFlag,
				expiresSFlag,
				providerFlag,
				attrsFlag,
				legacyFlag,
				forwardingModeFlag,
				statusFlag,
				contextFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return registrationsTemplate(c, ngsi, client)
			},
		},
	},
}
