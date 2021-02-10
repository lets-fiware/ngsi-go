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
	"errors"
	"fmt"
	"io"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"

	"github.com/urfave/cli/v2"
)

// Version  has a version number of NGSI Go
var Version = ""

// Revision has a git hash value
var Revision = ""

const copyright = "(c) 2020-2021 Kazuhito Suda"

var usage = "command-line tool for FIWARE NGSI and NGSI-LD"

var gCmdMode = ""

var gNetLib NetLib

// Run is a main rouitne of NGSI Go
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	ngsi := ngsilib.NewNGSI()
	defer ngsi.Close()

	ngsi.InitLog(stdin, stdout, stderr)
	Version = fmt.Sprintf("%s (git_hash:%s)", Version, Revision)

	gNetLib = &netLib{}
	cli.ErrWriter = stderr
	cli.HelpFlag = helpFlag

	app := getNgsiApp()
	err := app.Run(args)
	if err != nil {
		ngsi.Logging(ngsilib.LogErr, message(err)+"\n")
		for err != nil {
			err = errors.Unwrap(err)
			if err == nil {
				break
			}
			ngsi.Logging(ngsilib.LogDebug, fmt.Sprintf("%T\n", err))
			ngsi.Logging(ngsilib.LogInfo, message(err)+"\n")
		}
		ngsi.Logging(ngsilib.LogInfo, "abnormal termination\n")
		return 1
	}
	ngsi.Logging(ngsilib.LogInfo, "normal termination\n")
	return 0
}

func getNgsiApp() *cli.App {
	return &cli.App{
		EnableBashCompletion: true,
		Copyright:            copyright,
		Version:              Version,
		Usage:                usage,
		HideVersion:          false,
		Flags: []cli.Flag{
			syslogFlag,
			stderrFlag,
			configFlag,
			cacheFlag,
			marginFlag,
			timeOutFlag,
			maxCountFlag,
			batchFlag,
			cmdNameFlag,
		},
		Commands: []*cli.Command{
			&adminCmd,
			&apisCmd,
			&appendCmd,
			&brokersCmd,
			&contextCmd,
			&copyCmd,
			&countCmd,
			&createCmd,
			&debugCmd,
			&deleteCmd,
			&devicesCmd,
			&documentsCmd,
			&getCmd,
			&hDeleteCmd,
			&hGetCmd,
			&healthCmd,
			&listCmd,
			&lsCmd,
			&removeCmd,
			&receiverCmd,
			&replaceCmd,
			&settingsCmd,
			&serverCmd,
			&servicesCmd,
			&templateCmd,
			&tokenCmd,
			&updateCmd,
			&upsertCmd,
			&versionCmd,
		},
	}
}

func message(err error) (s string) {
	switch e := err.(type) {
	case *ngsilib.LibError:
		s = e.String()
	case *ngsiCmdError:
		s = e.String()
	default:
		s = e.Error()
	}
	return
}

func isSetOR(c *cli.Context, params []string) bool {
	for _, param := range params {
		if c.IsSet(param) {
			return true
		}
	}
	return false
}

func isSetAND(c *cli.Context, params []string) bool {
	for _, param := range params {
		if !c.IsSet(param) {
			return false
		}
	}
	return true
}

var copyCmd = cli.Command{
	Name:     "cp",
	Usage:    "copy entities",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		destinationFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
		typeRFlag,
		token2Flag,
		tenant2Flag,
		scope2Flag,
		runFlag,
	},
	Action: func(c *cli.Context) error {
		return copy(c)
	},
}

var countCmd = cli.Command{
	Name:     "wc",
	Usage:    "print number of entities, subscriptions, registrations or types",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "entities",
			Usage: "print number of entities",
			Flags: []cli.Flag{
				typeFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return entitiesCount(c)
			},
		},
		{
			Name:  "subscriptions",
			Usage: "print number of subscriptions",
			Action: func(c *cli.Context) error {
				return subscriptionsCount(c)
			},
		},
		{
			Name:  "registrations",
			Usage: "print number of registrations",
			Action: func(c *cli.Context) error {
				return registrationsCount(c)
			},
		},
		{
			Name:  "types",
			Usage: "print number of types",
			Action: func(c *cli.Context) error {
				return typesCount(c)
			},
		},
	},
}

var documentsCmd = cli.Command{
	Name:     "man",
	Usage:    "print urls of document",
	Category: "CONVENIENCE",
	Action: func(c *cli.Context) error {
		return documents(c)
	},
}

var lsCmd = cli.Command{
	Name:     "ls",
	Usage:    "list entities",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
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
		idFlag,
		linkFlag,
		verboseFlag,
		linesFlag,
		prettyFlag,
		safeStringFlag,
	},
	Action: func(c *cli.Context) error {
		return entitiesList(c)
	},
}

var removeCmd = cli.Command{
	Name:     "rm",
	Usage:    "remove entities",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
		typeRFlag,
		runFlag,
		linkFlag,
	},
	Action: func(c *cli.Context) error {
		return remove(c)
	},
}

var templateCmd = cli.Command{
	Name:     "template",
	Usage:    "create template of subscription or registration",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
		linkFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "subscription",
			Usage: "create template of subscription",
			Flags: []cli.Flag{
				ngsiTypeFlag,
				dataFlag,
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
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return subscriptionsTemplate(c)
			},
		},
		{
			Name:  "registration",
			Usage: "create template of registration",
			Flags: []cli.Flag{
				ngsiTypeFlag,
				dataFlag,
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
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return registrationsTemplate(c)
			},
		},
	},
}

var receiverCmd = cli.Command{
	Name:     "receiver",
	Category: "CONVENIENCE",
	Usage:    "notification receiver",
	Flags: []cli.Flag{
		receiverHostFlag,
		receiverPortFlag,
		receiverURLFlag,
		prettyFlag,
		receiverHTTPSFlag,
		receiverKeyFlag,
		receiverCertFlag,
		verboseFlag,
		headerFlag,
	},
	Action: func(c *cli.Context) error {
		return receiver(c)
	},
}

var versionCmd = cli.Command{
	Name:     "version",
	Category: "CONVENIENCE",
	Usage:    "print the version",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		prettyFlag,
	},
	Action: func(c *cli.Context) error {
		return cbVersion(c)
	},
}

var brokersCmd = cli.Command{
	Name:     "broker",
	Usage:    "manage config for broker",
	Category: "MANAGEMENT",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list brokers",
			Flags: []cli.Flag{
				hostFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return brokersList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get broker",
			Flags: []cli.Flag{
				hostFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return brokersGet(c)
			},
		},
		{
			Name:  "add",
			Usage: "add broker",
			Flags: []cli.Flag{
				hostFlag,
				brokerHostFlag,
				ngsiTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				tokenFlag,
				tenantFlag,
				scopeFlag,
				safeStringFlag,
				xAuthTokenFlag,
			},
			Action: func(c *cli.Context) error {
				return brokersAdd(c)
			},
		},
		{
			Name:  "update",
			Usage: "update broker",
			Flags: []cli.Flag{
				hostFlag,
				brokerHostFlag,
				ngsiTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				tokenFlag,
				tenantFlag,
				scopeFlag,
				safeStringFlag,
				xAuthTokenFlag,
			},
			Action: func(c *cli.Context) error {
				return brokersUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete broker",
			Flags: []cli.Flag{
				hostFlag,
			},
			Action: func(c *cli.Context) error {
				return brokersDelete(c)
			},
		},
	},
}

var serverCmd = cli.Command{
	Name:     "server",
	Usage:    "manage config for server",
	Category: "MANAGEMENT",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list servers",
			Flags: []cli.Flag{
				hostFlag,
				jsonFlag,
				prettyFlag,
				allServersFlag,
			},
			Action: func(c *cli.Context) error {
				return serverList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get server",
			Flags: []cli.Flag{
				hostFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return serverGet(c)
			},
		},
		{
			Name:  "add",
			Usage: "add server",
			Flags: []cli.Flag{
				hostFlag,
				serverHost2Flag,
				serverTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				tokenFlag,
				tenantFlag,
				scopeFlag,
				safeStringFlag,
				xAuthTokenFlag,
			},
			Action: func(c *cli.Context) error {
				return serverAdd(c)
			},
		},
		{
			Name:  "update",
			Usage: "update server",
			Flags: []cli.Flag{
				hostFlag,
				serverHost2Flag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				tokenFlag,
				tenantFlag,
				scopeFlag,
				safeStringFlag,
				xAuthTokenFlag,
			},
			Action: func(c *cli.Context) error {
				return serverUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete server",
			Flags: []cli.Flag{
				hostFlag,
			},
			Action: func(c *cli.Context) error {
				return serverDelete(c)
			},
		},
	},
}

var contextCmd = cli.Command{
	Name:     "context",
	Usage:    "manage @context",
	Category: "MANAGEMENT",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "List @context",
			Flags: []cli.Flag{
				nameFlag,
			},
			Action: func(c *cli.Context) error {
				return contextList(c)
			},
		},
		{
			Name:  "add",
			Usage: "Add @context",
			Flags: []cli.Flag{
				nameRFlag,
				urlFlag,
				jsonFlag,
			},
			Action: func(c *cli.Context) error {
				return contextAdd(c)
			},
		},
		{
			Name:  "update",
			Usage: "Update @context",
			Flags: []cli.Flag{
				nameRFlag,
				urlFlag,
			},
			Action: func(c *cli.Context) error {
				return contextUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "Delete @context",
			Flags: []cli.Flag{
				nameRFlag,
			},
			Action: func(c *cli.Context) error {
				return contextDelete(c)
			},
		},
		{
			Name:  "server",
			Usage: "serve @context",
			Flags: []cli.Flag{
				nameFlag,
				dataFlag,
				serverHostFlag,
				serverPortFlag,
				serverURLFlag,
				serverHTTPSFlag,
				serverKeyFlag,
				serverCertFlag,
			},
			Action: func(c *cli.Context) error {
				return contextServer(c)
			},
		},
	},
}

var settingsCmd = cli.Command{
	Name:     "settings",
	Category: "MANAGEMENT",
	Usage:    "manage settings",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "List settings",
			Flags: []cli.Flag{
				allFlag,
			},
			Action: func(c *cli.Context) error {
				return settingsList(c)
			},
		},
		{
			Name:  "delete",
			Usage: "Delete setting",
			Flags: []cli.Flag{
				itemsFlag,
			},
			Action: func(c *cli.Context) error {
				return settingsDelete(c)
			},
		},
		{
			Name:  "clear",
			Usage: "Clear settings",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) error {
				return settingsClear(c)
			},
		},
	},
}

var tokenCmd = cli.Command{
	Name:  "token",
	Usage: "manage token",
	Flags: []cli.Flag{
		hostFlag,
		verboseFlag,
		expiresFlag,
		prettyFlag,
	},
	Category: "MANAGEMENT",
	Action: func(c *cli.Context) error {
		return tokenCommand(c)
	},
}

var appendCmd = cli.Command{
	Name:     "append",
	Usage:    "append attributes",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "attrs",
			Usage: "append attrs",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				keyValuesFlag,
				appendFlag,
				dataFlag,
				linkFlag,
				contextFlag,
			},
			Action: func(c *cli.Context) error {
				return attrsAppend(c)
			},
		},
	},
}

var createCmd = cli.Command{
	Name:     "create",
	Usage:    "create entity(ies), subscription or registration",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "entities",
			Usage: "create entities",
			Flags: []cli.Flag{
				keyValuesFlag,
				dataFlag,
				linkFlag,
				contextFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return batch(c, "create")
			},
		},
		{
			Name:  "entity",
			Usage: "create entity",
			Flags: []cli.Flag{
				dataFlag,
				keyValuesFlag,
				upsertFlag,
				linkFlag,
				contextFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return entityCreate(c)
			},
		},
		{
			Name:  "subscription",
			Usage: "create subscription",
			Flags: []cli.Flag{
				dataFlag,
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
				safeStringFlag,
				notifyURLFlag,
				getFlag,
			},
			Action: func(c *cli.Context) error {
				return subscriptionsCreate(c)
			},
		},
		{
			Name:  "registration",
			Usage: "create registration",
			Flags: []cli.Flag{
				dataFlag,
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
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return registrationsCreate(c)
			},
		},
	},
}

var deleteCmd = cli.Command{
	Name:     "delete",
	Usage:    "delete entity(ies), attribute, subscription or registration",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "entities",
			Usage: "delete entities",
			Flags: []cli.Flag{
				keyValuesFlag,
				dataFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return batch(c, "delete")
			},
		},
		{
			Name:  "entity",
			Usage: "delete entity",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return entityDelete(c)
			},
		},
		{
			Name:  "attr",
			Usage: "delete attr",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				attrNameRFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return attrDelete(c)
			},
		},
		{
			Name:  "subscription",
			Usage: "delete subscription",
			Flags: []cli.Flag{
				idRFlag,
			},
			Action: func(c *cli.Context) error {
				return subscriptionsDelete(c)
			},
		},
		{
			Name:  "registration",
			Usage: "delete registration",
			Flags: []cli.Flag{
				idRFlag,
			},
			Action: func(c *cli.Context) error {
				return registrationsDelete(c)
			},
		},
	},
}

var getCmd = cli.Command{
	Name:     "get",
	Usage:    "get entity(ies), attribute(s), subscription, registration or type",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "entities",
			Usage: "get entities",
			Flags: []cli.Flag{
				orderByFlag,
				countFlag,
				keyValuesFlag,
				valuesFlag,
				uniqueFlag,
				verboseFlag,
				linesFlag,
				dataFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return opQuery(c)
			},
		},
		{
			Name:  "entity",
			Usage: "get entity",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				attrsFlag,
				keyValuesFlag,
				valuesFlag,
				uniqueFlag,
				sysAttrsFlag,
				linkFlag,
				acceptJSONFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return entityRead(c)
			},
		},
		{
			Name:  "attr",
			Usage: "get attr",
			Flags: []cli.Flag{
				idFlag,
				typeFlag,
				attrNameRFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return attrRead(c)
			},
		},
		{
			Name:  "attrs",
			Usage: "get attrs",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				attrsFlag,
				metadataFlag,
				keyValuesFlag,
				valuesFlag,
				uniqueFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return attrsRead(c)
			},
		},
		{
			Name:  "subscription",
			Usage: "get subscription",
			Flags: []cli.Flag{
				idRFlag,
				localTimeFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return subscriptionGet(c)
			},
		},
		{
			Name:  "registration",
			Usage: "get registration",
			Flags: []cli.Flag{
				idRFlag,
				localTimeFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return registrationsGet(c)
			},
		},
		{
			Name:  "type",
			Usage: "get type",
			Flags: []cli.Flag{
				typeRFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return typeGet(c)
			},
		},
	},
}

var listCmd = cli.Command{
	Name:     "list",
	Usage:    "list types, entities, subscriptions or registrations",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "types",
			Usage: "list types",
			Flags: []cli.Flag{
				jsonFlag,
				prettyFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return typesList(c)
			},
		},
		{
			Name:  "entities",
			Usage: "list entities",
			Flags: []cli.Flag{
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
				idFlag,
				linkFlag,
				acceptJSONFlag,
				verboseFlag,
				linesFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return entitiesList(c)
			},
		},
		{
			Name:  "subscriptions",
			Usage: "list subscriptions",
			Flags: []cli.Flag{
				verboseFlag,
				jsonFlag,
				statusFlag,
				localTimeFlag,
				queryFlag,
				itemsFlag,
				prettyFlag,
				safeStringFlag,
				countFlag,
			},
			Action: func(c *cli.Context) error {
				return subscriptionsList(c)
			},
		},
		{
			Name:  "registrations",
			Usage: "list registrations",
			Flags: []cli.Flag{
				verboseFlag,
				jsonFlag,
				localTimeFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return registrationsList(c)
			},
		},
	},
}

var replaceCmd = cli.Command{
	Name:     "replace",
	Usage:    "replace entities or attributes",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
		linkFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "entities",
			Usage: "replace entities",
			Flags: []cli.Flag{
				keyValuesFlag,
				dataFlag,
			},
			Action: func(c *cli.Context) error {
				return batch(c, "replace")
			},
		},
		{
			Name:  "attrs",
			Usage: "replace attrs",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				dataFlag,
				keyValuesFlag,
			},
			Action: func(c *cli.Context) error {
				return attrsReplace(c)
			},
		},
	},
}

var updateCmd = cli.Command{
	Name:     "update",
	Usage:    "update entities, attribute(s) or subscription",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "attr",
			Usage: "update attr",
			Flags: []cli.Flag{
				idRFlag,
				dataFlag,
				attrNameRFlag,
				linkFlag,
				contextFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return attrUpdate(c)
			},
		},
		{
			Name:  "attrs",
			Usage: "update attrs",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				keyValuesFlag,
				dataFlag,
				linkFlag,
				contextFlag,
			},
			Action: func(c *cli.Context) error {
				return attrsUpdate(c)
			},
		},
		{
			Name:  "subscription",
			Usage: "update subscription",
			Flags: []cli.Flag{
				idRFlag,
				dataFlag,
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
				safeStringFlag,
				notifyURLFlag,
				getFlag,
			},
			Action: func(c *cli.Context) error {
				return subscriptionsUpdate(c)
			},
		},
		{
			Name:  "entities",
			Usage: "update entities",
			Flags: []cli.Flag{
				keyValuesFlag,
				dataFlag,
				noOverwriteFlag,
				replaceFlag,
				linkFlag,
				contextFlag,
			},
			Action: func(c *cli.Context) error {
				return batch(c, "update")
			},
		},
	},
}

var upsertCmd = cli.Command{
	Name:     "upsert",
	Usage:    "upsert entity or entities",
	Category: "NGSI",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "entity",
			Usage: "upsert entity",
			Flags: []cli.Flag{
				dataFlag,
				keyValuesFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return entityUpsert(c)
			},
		},
		{
			Name:  "entities",
			Usage: "upsert entities",
			Flags: []cli.Flag{
				dataFlag,
				replaceFlag,
				updateFlag,
				linkFlag,
				contextFlag,
			},
			Action: func(c *cli.Context) error {
				return batch(c, "upsert")
			},
		},
	},
}

var adminCmd = cli.Command{
	Name:     "admin",
	Usage:    "admin command for FIWARE Orion",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "log",
			Usage: "admin log",
			Flags: []cli.Flag{
				levelFlag,
				loggingFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return adminLog(c)
			},
		},
		{
			Name:  "trace",
			Usage: "admin trace",
			Flags: []cli.Flag{
				levelFlag,
				setFlag,
				deleteFlag,
				loggingFlag,
			},
			Action: func(c *cli.Context) error {
				return adminTrace(c)
			},
		},
		{
			Name:  "semaphore",
			Usage: "print semaphore",
			Flags: []cli.Flag{
				loggingFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return adminSemaphore(c)
			},
		},
		{
			Name:  "metrics",
			Usage: "manage metrics",
			Flags: []cli.Flag{
				deleteFlag,
				resetFlag,
				loggingFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return adminMetrics(c)
			},
		},
		{
			Name:  "statistics",
			Usage: "print statistics",
			Flags: []cli.Flag{
				deleteFlag,
				loggingFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return adminStatistics(c)
			},
		},
		{
			Name:  "cacheStatistics",
			Usage: "print cache statistics",
			Flags: []cli.Flag{
				deleteFlag,
				loggingFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return adminCacheStatistics(c)
			},
		},
	},
}

var apisCmd = cli.Command{
	Name:     "apis",
	Usage:    "print endpoints of API",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		prettyFlag,
	},
	Action: func(c *cli.Context) error {
		return apis(c)
	},
}

var healthCmd = cli.Command{
	Name:     "health",
	Usage:    "print health status",
	Category: "CONVENIENCE",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Action: func(c *cli.Context) error {
		return healthCheck(c)
	},
}

var hGetCmd = cli.Command{
	Name:     "hget",
	Usage:    "get historical raw and aggregated time series context information",
	Category: "TIME SERIES",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "attr",
			Usage: "history of an attribute",
			Flags: []cli.Flag{
				typeFlag,
				idFlag,
				attrNameFlag,
				sameTypeFlag,
				nTypesFlag,
				aggrMethodFlag,
				aggrPeriodFlag,
				fromDateFlag,
				toDateFlag,
				lastNFlag,
				hLimitFlag,
				hOffsetFlag,
				georelFlag,
				geometryFlag,
				coordsFlag,
				valueFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return tsAttrRead(c)
			},
		},
		{
			Name:  "attrs",
			Usage: "history of attributes",
			Flags: []cli.Flag{
				typeFlag,
				idFlag,
				attrsFlag,
				sameTypeFlag,
				nTypesFlag,
				aggrMethodFlag,
				aggrPeriodFlag,
				fromDateFlag,
				toDateFlag,
				lastNFlag,
				hLimitFlag,
				hOffsetFlag,
				georelFlag,
				geometryFlag,
				coordsFlag,
				valueFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return tsAttrsRead(c)
			},
		},
		{
			Name:  "entities",
			Usage: "list of all the entity id",
			Flags: []cli.Flag{
				typeFlag,
				fromDateFlag,
				toDateFlag,
				hLimitFlag,
				hOffsetFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return tsEntitiesRead(c)
			},
		},
	},
}

var hDeleteCmd = cli.Command{
	Name:     "hdelete",
	Usage:    "delete historical raw and aggregated time series context information",
	Category: "TIME SERIES",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "entities",
			Usage: "delete historical data of all entities of a certain type",
			Flags: []cli.Flag{
				idFlag,
				typeFlag,
				dropTableFlag,
				fromDateFlag,
				toDateFlag,
				runFlag,
			},
			Action: func(c *cli.Context) error {
				return tsEntitiesDelete(c)
			},
		},
		{
			Name:  "entity",
			Usage: "delete historical data of a certain entity",
			Flags: []cli.Flag{
				idFlag,
				typeFlag,
				fromDateFlag,
				toDateFlag,
				runFlag,
			},
			Action: func(c *cli.Context) error {
				return tsEntityDelete(c)
			},
		},
		{
			Name:  "attr",
			Usage: "delete all the data associated to certain attribute of certain entity",
			Flags: []cli.Flag{
				idFlag,
				typeFlag,
				attrNameFlag,
				runFlag,
			},
			Action: func(c *cli.Context) error {
				return cometAttrDelete(c)
			},
		},
	},
}

var servicesCmd = cli.Command{
	Name:     "services",
	Usage:    "services command for IoT Agent",
	Category: "IoT Agent",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list configuration group",
			Flags: []cli.Flag{
				servicesLimitFlag,
				servicesOffsetFlag,
				resourceFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return idasServicesList(c)
			},
		},
		{
			Name:  "create",
			Usage: "create a configuration group",
			Flags: []cli.Flag{
				servicesDataFlag,
				apikeyFlag,
				servicesTokenFlag,
				cbrokerFlag,
				typeFlag,
				resourceFlag,
			},
			Action: func(c *cli.Context) error {
				return idasServicesCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update a configuration group",
			Flags: []cli.Flag{
				servicesDataFlag,
				apikeyFlag,
				servicesTokenFlag,
				cbrokerFlag,
				typeFlag,
				resourceFlag,
			},
			Action: func(c *cli.Context) error {
				return idasServicesUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "remove a configuration group",
			Flags: []cli.Flag{
				apikeyFlag,
				resourceFlag,
				servicesDeviceFlag,
			},
			Action: func(c *cli.Context) error {
				return idasServicesDelete(c)
			},
		},
	},
}

var devicesCmd = cli.Command{
	Name:     "devices",
	Usage:    "devices command for IoT Agent",
	Category: "IoT Agent",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list all devices",
			Flags: []cli.Flag{
				devicesLimit,
				devicesOffset,
				devicesDetailed,
				devicesEntity,
				devicesProtocol,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return idasDevicesList(c)
			},
		},
		{
			Name:  "create",
			Usage: "create a device",
			Flags: []cli.Flag{
				devicesDataFlag,
			},
			Action: func(c *cli.Context) error {
				return idasDevicesCreate(c)
			},
		},
		{
			Name:  "get",
			Usage: "get a device",
			Flags: []cli.Flag{
				devicesIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return idasDevicesGet(c)
			},
		},
		{
			Name:  "update",
			Usage: "update a device",
			Flags: []cli.Flag{
				devicesIDFlag,
				devicesDataFlag,
			},
			Action: func(c *cli.Context) error {
				return idasDevicesUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete a device",
			Flags: []cli.Flag{
				devicesIDFlag,
			},
			Action: func(c *cli.Context) error {
				return idasDevicesDelete(c)
			},
		},
	},
}
