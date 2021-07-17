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
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

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

	bufStdin := bufio.NewReader(stdin)
	bufStdout := bufio.NewWriter(stdout)
	ngsi.InitLog(bufStdin, bufStdout, stderr)

	Version = fmt.Sprintf("%s (git_hash:%s)", Version, Revision)

	gNetLib = &netLib{}
	cli.ErrWriter = stderr
	cli.HelpFlag = helpFlag

	app := getNgsiApp()
	err := app.Run(args)
	if err != nil {
		s := strings.TrimRight(message(err), "\n") + "\n"
		ngsi.Logging(ngsilib.LogErr, s)
		for err != nil {
			err = errors.Unwrap(err)
			if err == nil {
				break
			}
			ngsi.Logging(ngsilib.LogDebug, fmt.Sprintf("%T\n", err))
			s = strings.TrimRight(message(err), "\n") + "\n"
			ngsi.Logging(ngsilib.LogInfo, s)
		}
		ngsi.Logging(ngsilib.LogInfo, "abnormal termination\n")
		_ = bufStdout.Flush()
		return 1
	}
	ngsi.Logging(ngsilib.LogInfo, "normal termination\n")

	_ = bufStdout.Flush()

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
			insecureSkipVerifyFlag,
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
			&regProxyCmd,
			&replaceCmd,
			&perseoRulesCmd,
			&settingsCmd,
			&serverCmd,
			&servicesCmd,
			&templateCmd,
			&tokenCmd,
			&updateCmd,
			&upsertCmd,
			&versionCmd,
			&applicationsCmd,
			&usersCmd,
			&organizationsCmd,
			&providersCmd,
			&cygnusNamemappingsCmd,
			&cygnusGroupingrulesCmd,
			&wireCloudPreferencesCmd,
			&wireCloudResourcesCmd,
			&wireCloudWorkspacesCmd,
			&wireCloudTabsCmd,
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
		linkFlag,
		typeFlag,
		token2Flag,
		tenant2Flag,
		scope2Flag,
		context2Flag,
		ngsiV1Flag,
		skipForwardingFlag,
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
				skipForwardingFlag,
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
		skipForwardingFlag,
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
		typeFlag,
		linkFlag,
		ngsiV1Flag,
		skipForwardingFlag,
		runFlag,
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

var regProxyCmd = cli.Command{
	Name:     "regproxy",
	Category: "CONVENIENCE",
	Usage:    "registration proxy",
	Flags: []cli.Flag{
		regProxyHostFlag,
		regProxyRhostFlag,
		regProxyPortFlag,
		regProxyURLFlag,
		regProxyReplaceTenantFlag,
		regProxyReplaceScopeFlag,
		regProxyReplaceURLFlag,
		regProxyHTTPSFlag,
		regProxyKeyFlag,
		regProxyCertFlag,
		verboseFlag,
	},
	Action: func(c *cli.Context) error {
		return regProxy(c)
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
				clearTextFlag,
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
				clearTextFlag,
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
				brokerTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				tokenScopeFlag,
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
				brokerTypeFlag,
				idmTypeFlag,
				idmHostFlag,
				apiPathFlag,
				usernameFlag,
				passwordFlag,
				clientIDFlag,
				clientSecretFlag,
				tokenScopeFlag,
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
				clearTextFlag,
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
				clearTextFlag,
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
				tokenScopeFlag,
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
				tokenScopeFlag,
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
		{
			Name:  "previousArgs",
			Usage: "Set PreviousArgs mode",
			Flags: []cli.Flag{
				offFlag,
				onFlag,
			},
			Action: func(c *cli.Context) error {
				return settingsPreviousArgs(c)
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
		prettyFlag,
		expiresFlag,
		revokeFlag,
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
		{
			Name:  "tattrs",
			Usage: "append attribute instance of temporal entity",
			Flags: []cli.Flag{
				idFlag,
				dataFlag,
				linkFlag,
				contextFlag,
			},
			Action: func(c *cli.Context) error {
				return troeAttrsAppend(c)
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
			Name:  "tentity",
			Usage: "create temporal entity",
			Flags: []cli.Flag{
				dataFlag,
				linkFlag,
				contextFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return troeCreate(c)
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
			Name:  "tentity",
			Usage: "delete temporal entity",
			Flags: []cli.Flag{
				idFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return troeDelete(c)
			},
		},
		{
			Name:  "attr",
			Usage: "delete attr",
			Flags: []cli.Flag{
				idRFlag,
				typeFlag,
				attrRFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return attrDelete(c)
			},
		},
		{
			Name:  "tattr",
			Usage: "delete attr for temporal entity",
			Flags: []cli.Flag{
				idFlag,
				attrFlag,
				deleteAllFlag,
				deleteDatasetID,
				instanceIDFlag,
				linkFlag,
			},
			Action: func(c *cli.Context) error {
				return troeAttrDelete(c)
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
				acceptGeoJSONFlag,
				prettyFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return entityRead(c)
			},
		},
		{
			Name:  "tentity",
			Usage: "get temporal entity",
			Flags: []cli.Flag{
				idFlag,
				attrsFlag,
				timePropertyFlag,
				fromDateFlag,
				toDateFlag,
				lastNFlag,
				linkFlag,
				temporalValuesFlag,
				sysAttrsFlag,
				acceptJSONFlag,
				prettyFlag,
				safeStringFlag,
				etsi10Flag,
			},
			Action: func(c *cli.Context) error {
				return troeRead(c)
			},
		},
		{
			Name:  "attr",
			Usage: "get attr",
			Flags: []cli.Flag{
				idFlag,
				typeFlag,
				attrRFlag,
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
				skipForwardingFlag,
				linkFlag,
				acceptJSONFlag,
				acceptGeoJSONFlag,
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
			Name:  "tentities",
			Usage: "list temporal entities",
			Flags: []cli.Flag{
				idFlag,
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
				verboseFlag,
				linesFlag,
				prettyFlag,
				safeStringFlag,
				etsi10Flag,
			},
			Action: func(c *cli.Context) error {
				return troeList(c)
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
				attrRFlag,
				linkFlag,
				contextFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return attrUpdate(c)
			},
		},
		{
			Name:  "tattr",
			Usage: "update attr instance of temporal entity",
			Flags: []cli.Flag{
				idFlag,
				dataFlag,
				attrFlag,
				instanceIDFlag,
				linkFlag,
				contextFlag,
				safeStringFlag,
			},
			Action: func(c *cli.Context) error {
				return troeAttrUpdate(c)
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
	Usage:    "admin command for FIWARE Orion, Cygnus, Perseo, Scorpio",
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
		&cygnusLoggersCmd,
		&cygnusAppendersCmd,
		&scorpioCmd,
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
				attrFlag,
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
				attrFlag,
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
	Usage:    "manage services for IoT Agent",
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
	Usage:    "manage devices for IoT Agent",
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

// PERSEO FE Rules
var perseoRulesCmd = cli.Command{
	Name:     "rules",
	Usage:    "rules command for PERSEO",
	Category: "Context-Aware CEP",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		tenantFlag,
		scopeFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list all plain rules",
			Flags: []cli.Flag{
				perseoRulesLimitFlag,
				perseoRulesOffsetFlag,
				perseoRulesCount,
				perseoRulesRaw,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return perseoRulesList(c)
			},
		},
		{
			Name:  "create",
			Usage: "create a plain rule",
			Flags: []cli.Flag{
				perseoRulesDataFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return perseoRulesCreate(c)
			},
		},
		{
			Name:  "get",
			Usage: "get a plain rule",
			Flags: []cli.Flag{
				perseoRulesNameFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return perseoRulesGet(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete a plain rule",
			Flags: []cli.Flag{
				perseoRulesNameFlag,
			},
			Action: func(c *cli.Context) error {
				return perseoRulesDelete(c)
			},
		},
	},
}

// Keyrock
var usersCmd = cli.Command{
	Name:     "users",
	Usage:    "manage users for Keyrock",
	Category: "Keyrock",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list users",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return usersList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get user",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return usersGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create user",
			Flags: []cli.Flag{
				keyrockUserNameFlag,
				keyrockPasswordFlag,
				keyrockEmailFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return usersCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update user",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				keyrockUserNameFlag,
				keyrockPasswordFlag,
				keyrockEmailFlag,
				keyrockGravatarFlag,
				keyrockDescriptionFlag,
				keyrockWebsiteFlag,
				keyrockExtraFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return usersUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete user",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
			},
			Action: func(c *cli.Context) error {
				return usersDelete(c)
			},
		},
	},
}

var applicationsCmd = cli.Command{
	Name:     "applications",
	Usage:    "manage applications for Keyrock",
	Category: "Keyrock",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list applications",
			Flags: []cli.Flag{
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return applicationsList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get application",
			Flags: []cli.Flag{
				keyrockApplicationIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return applicationsGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create application",
			Flags: []cli.Flag{
				keyrcokApplicationDataFlag,
				keyrcokApplicationNameFlag,
				keyrcokApplicationDescriptionFlag,
				keyrcokApplicationURLFlag,
				keyrcokApplicationRedirectURIFlag,
				keyrcokApplicationRedirectSignOutURIFlag,
				keyrcokApplicationGrantTypeFlag,
				keyrcokApplicationTokenTypesFlag,
				keyrcokApplicationResponseTypeFlag,
				keyrcokApplicationClientTypeFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return applicationsCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update application",
			Flags: []cli.Flag{
				keyrockApplicationIDFlag,
				keyrcokApplicationDataFlag,
				keyrcokApplicationNameFlag,
				keyrcokApplicationDescriptionFlag,
				keyrcokApplicationURLFlag,
				keyrcokApplicationRedirectURIFlag,
				keyrcokApplicationRedirectSignOutURIFlag,
				keyrcokApplicationGrantTypeFlag,
				keyrcokApplicationTokenTypesFlag,
				keyrcokApplicationResponseTypeFlag,
				keyrcokApplicationClientTypeFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return applicationsUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete application",
			Flags: []cli.Flag{
				keyrockApplicationIDFlag,
			},
			Action: func(c *cli.Context) error {
				return applicationsDelete(c)
			},
		},
		&rolesCmd,
		&permissionsCmd,
		&pepProxiesCmd,
		&iotAgentCmd,
		&appsUsersCmd,
		&appsOrgsCmd,
		&trustedAppsCmd,
	},
}

var organizationsCmd = cli.Command{
	Name:     "organizations",
	Usage:    "manage organizations for Keyrock",
	Category: "Keyrock",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list organizations",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return organizationsList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get organization",
			Flags: []cli.Flag{
				keyrockOrganizationIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return organizationsGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create organization",
			Flags: []cli.Flag{
				keyrockOrganizationNameFlag,
				keyrockDescriptionFlag,
				keyrockWebsiteFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return organizationsCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update organization",
			Flags: []cli.Flag{
				keyrockOrganizationIDFlag,
				keyrockOrganizationNameFlag,
				keyrockDescriptionFlag,
				keyrockWebsiteFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return organizationsUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete organization",
			Flags: []cli.Flag{
				keyrockOrganizationIDFlag,
			},
			Action: func(c *cli.Context) error {
				return organizationsDelete(c)
			},
		},
		&orgsUsersCmd,
	},
}

var rolesCmd = cli.Command{
	Name:     "roles",
	Usage:    "manage roles for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockApplicationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list roles",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return rolesList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get a role",
			Flags: []cli.Flag{
				keyrockRolesIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return rolesGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create a role",
			Flags: []cli.Flag{
				keyrockRoleDataFlag,
				keyrockRoleNameFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return rolesCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update a role",
			Flags: []cli.Flag{
				keyrockRolesIDFlag,
				keyrockRoleDataFlag,
				keyrockRoleNameFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return rolesUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete a role",
			Flags: []cli.Flag{
				keyrockRolesIDFlag,
			},
			Action: func(c *cli.Context) error {
				return rolesDelete(c)
			},
		},
		{
			Name:  "permissions",
			Usage: "list permissions associated to a role",
			Flags: []cli.Flag{
				keyrockRolesIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsRolePermList(c)
			},
		},
		{
			Name:  "assign",
			Usage: "Assign a permission to a role",
			Flags: []cli.Flag{
				keyrockRolesIDFlag,
				keyrockPermissionIDFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsRolePermAssign(c)
			},
		},
		{
			Name:  "unassign",
			Usage: "delete a permission from a role",
			Flags: []cli.Flag{
				keyrockRolesIDFlag,
				keyrockPermissionIDFlag,
			},
			Action: func(c *cli.Context) error {
				return appsRolePermDelete(c)
			},
		},
	},
}

var permissionsCmd = cli.Command{
	Name:     "permissions",
	Usage:    "manage permissions for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockApplicationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list permissions",
			Flags: []cli.Flag{
				keyrockPermissionIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return permissionsList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get permission",
			Flags: []cli.Flag{
				keyrockPermissionIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return permissionsGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create permission",
			Flags: []cli.Flag{
				keyrcokPermissionDataFlag,
				keyrockPermissionNameFlag,
				keyrockPermissionDescriptionFlag,
				keyrockPermissionActionFlag,
				keyrockPermissionResourceFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return permissionsCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update permission",
			Flags: []cli.Flag{
				keyrockPermissionIDFlag,
				keyrcokPermissionDataFlag,
				keyrockPermissionNameFlag,
				keyrockPermissionDescriptionFlag,
				keyrockPermissionActionFlag,
				keyrockPermissionResourceFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return permissionsUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete permission",
			Flags: []cli.Flag{
				keyrockPermissionIDFlag,
			},
			Action: func(c *cli.Context) error {
				return permissionsDelete(c)
			},
		},
	},
}

var pepProxiesCmd = cli.Command{
	Name:     "pep",
	Usage:    "manage PEP Proxies for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockApplicationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list pep proxy",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return pepProxiesList(c)
			},
		},
		{
			Name:  "create",
			Usage: "create pep proxy",
			Flags: []cli.Flag{
				runFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return pepProxiesCreate(c)
			},
		},
		{
			Name:  "reset",
			Usage: "reset pep proxy",
			Flags: []cli.Flag{
				runFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return pepProxiesReset(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete pep proxy",
			Action: func(c *cli.Context) error {
				return pepProxiesDelete(c)
			},
			Flags: []cli.Flag{
				runFlag,
			},
		},
	},
}

var iotAgentCmd = cli.Command{
	Name:     "iota",
	Usage:    "manage IoT Agents for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockApplicationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list iot agents",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return iotAgentsList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get iot agent",
			Flags: []cli.Flag{
				keyrockIotAgentsIDFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return iotAgentsGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create iot agent",
			Flags: []cli.Flag{
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return iotAgentsCreate(c)
			},
		},
		{
			Name:  "reset",
			Usage: "reset iot agent",
			Flags: []cli.Flag{
				keyrockIotAgentsIDFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return iotAgentsReset(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete iot agent",
			Flags: []cli.Flag{
				keyrockIotAgentsIDFlag,
			},
			Action: func(c *cli.Context) error {
				return iotAgentsDelete(c)
			},
		},
	},
}

var appsUsersCmd = cli.Command{
	Name:     "users",
	Usage:    "manage authorized users in an application for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockApplicationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list users in an application",
			Flags: []cli.Flag{
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsUsersList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get roles of a user in an application",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsUsersGet(c)
			},
		},
		{
			Name:  "assign",
			Usage: "assign a role to a user",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				keyrockRolesIDFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsUsersAssign(c)
			},
		},
		{
			Name:  "unassign",
			Usage: "delete a role assignment from a user",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				keyrockRolesIDFlag,
			},
			Action: func(c *cli.Context) error {
				return appsUsersUnassign(c)
			},
		},
	},
}

var appsOrgsCmd = cli.Command{
	Name:     "organizations",
	Usage:    "manage authorized organizations in an application for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockApplicationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list organizations in an application",
			Flags: []cli.Flag{
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsOrgsRolesList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get roles of an organization in an application",
			Flags: []cli.Flag{
				keyrockOrganizationIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsOrgsRolesGet(c)
			},
		},
		{
			Name:  "assign",
			Usage: "assign a role to an organization",
			Flags: []cli.Flag{
				keyrockOrganizationIDFlag,
				keyrockRolesIDFlag,
				keyrcokOrganizationRoleIDFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appsOrgsRolesAssign(c)
			},
		},
		{
			Name:  "unassign",
			Usage: "delete a role assignment from an organization",
			Flags: []cli.Flag{
				keyrockOrganizationIDFlag,
				keyrockRolesIDFlag,
				keyrcokOrganizationRoleIDFlag,
			},
			Action: func(c *cli.Context) error {
				return appsOrgsRolesUnassign(c)
			},
		},
	},
}

var orgsUsersCmd = cli.Command{
	Name:     "users",
	Usage:    "manage users of an organization for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockOrganizationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list users of an organization",
			Flags: []cli.Flag{
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return orgUsersList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get info of user organization relationship",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return orgUsersGet(c)
			},
		},
		{
			Name:  "add",
			Usage: "add a user to an organization",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				keyrcokOrganizationRoleIDFlag,
				verboseFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return orgUsersCreate(c)
			},
		},
		{
			Name:  "remove",
			Usage: "remove a user from an organization",
			Flags: []cli.Flag{
				keyrockUserIDFlag,
				keyrcokOrganizationRoleIDFlag,
			},
			Action: func(c *cli.Context) error {
				return orgUsersDelete(c)
			},
		},
	},
}

var trustedAppsCmd = cli.Command{
	Name:     "trusted",
	Usage:    "manage trusted applications for Keyrock",
	Category: "sub-command",
	Flags: []cli.Flag{
		keyrockApplicationIDFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list trusted applications associated to an application",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return trustedAppList(c)
			},
		},
		{
			Name:  "add",
			Usage: "add trusted application",
			Flags: []cli.Flag{
				keyrockTrustedIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return trustedAppAdd(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete trusted application",
			Flags: []cli.Flag{
				keyrockTrustedIDFlag,
			},
			Action: func(c *cli.Context) error {
				return trustedAppDelete(c)
			},
		},
	},
}

var providersCmd = cli.Command{
	Name:     "providers",
	Usage:    "print service providers for Keyrock",
	Category: "Keyrock",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
		prettyFlag,
	},
	Action: func(c *cli.Context) error {
		return providersGet(c)
	},
}

var cygnusLoggersCmd = cli.Command{
	Name:     "loggers",
	Usage:    "manage loggers for Cygnus",
	Category: "sub-command",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list loggers",
			Flags: []cli.Flag{
				cygnusLoggersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return loggersList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get logger",
			Flags: []cli.Flag{
				cygnusLoggersNameFlag,
				cygnusLoggersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return loggersGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create logger",
			Flags: []cli.Flag{
				cygnusLoggersNameFlag,
				cygnusLoggersDataFlag,
				cygnusLoggersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return loggersCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update logger",
			Flags: []cli.Flag{
				cygnusLoggersNameFlag,
				cygnusLoggersDataFlag,
				cygnusLoggersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return loggersUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete logger",
			Flags: []cli.Flag{
				cygnusLoggersNameFlag,
				cygnusLoggersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return loggersDelete(c)
			},
		},
	},
}

var cygnusAppendersCmd = cli.Command{
	Name:     "appenders",
	Usage:    "manage appenders for Cygnus",
	Category: "sub-command",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list appenders",
			Flags: []cli.Flag{
				cygnusAppendersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appendersList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get appender",
			Flags: []cli.Flag{
				cygnusAppendersNameFlag,
				cygnusAppendersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appendersGet(c)
			},
		},
		{
			Name:  "create",
			Usage: "create appender",
			Flags: []cli.Flag{
				cygnusAppendersNameFlag,
				cygnusAppendersDataFlag,
				cygnusAppendersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appendersCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update appender",
			Flags: []cli.Flag{
				cygnusAppendersNameFlag,
				cygnusAppendersDataFlag,
				cygnusAppendersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appendersUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete appender",
			Flags: []cli.Flag{
				cygnusAppendersNameFlag,
				cygnusAppendersTransientFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return appendersDelete(c)
			},
		},
	},
}

var cygnusNamemappingsCmd = cli.Command{
	Name:     "namemappings",
	Usage:    "manage namemappings for Cygnus",
	Category: "PERSISTING CONTEXT DATA",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list namemappings",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return namemappingsList(c)
			},
		},
		{
			Name:  "create",
			Usage: "create namemapping",
			Flags: []cli.Flag{
				cygnusNamemappingsDataFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return namemappingsCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update namemapping",
			Flags: []cli.Flag{
				cygnusNamemappingsDataFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return namemappingsUpdate(c)
			},
			Hidden: true,
		},
		{
			Name:  "delete",
			Usage: "delete namemapping",
			Flags: []cli.Flag{
				cygnusNamemappingsDataFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return namemappingsDelete(c)
			},
		},
	},
}

var cygnusGroupingrulesCmd = cli.Command{
	Name:     "groupingrules",
	Usage:    "manage groupingrules for Cygnus",
	Category: "PERSISTING CONTEXT DATA",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list groupingrules",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return groupingrulesList(c)
			},
		},
		{
			Name:  "create",
			Usage: "create groupingrule",
			Flags: []cli.Flag{
				cygnusGroupingrulesDataFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return groupingrulesCreate(c)
			},
		},
		{
			Name:  "update",
			Usage: "update groupingrule",
			Flags: []cli.Flag{
				cygnusGroupingrulesIDFlag,
				cygnusGroupingrulesDataFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return groupingrulesUpdate(c)
			},
		},
		{
			Name:  "delete",
			Usage: "delete groupingrule",
			Flags: []cli.Flag{
				cygnusGroupingrulesIDFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return groupingrulesDelete(c)
			},
		},
	},
}

var scorpioCmd = cli.Command{
	Name:     "scorpio",
	Usage:    "information command for Scorpio broker",
	Category: "sub-command",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "List of information paths",
			Action: func(c *cli.Context) error {
				return scorpioCommand(c, "")
			},
		},
		{
			Name:  "types",
			Usage: "print types",
			Action: func(c *cli.Context) error {
				return scorpioCommand(c, "types")
			},
		},
		{
			Name:  "localtypes",
			Usage: "print local types",
			Action: func(c *cli.Context) error {
				return scorpioCommand(c, "localtypes")
			},
		},
		{
			Name:  "stats",
			Usage: "print stats",
			Action: func(c *cli.Context) error {
				return scorpioCommand(c, "stats")
			},
		},
		{
			Name:  "health",
			Usage: "print health",
			Action: func(c *cli.Context) error {
				return scorpioCommand(c, "health")
			},
		},
	},
}

var wireCloudWorkspacesCmd = cli.Command{
	Name:     "workspaces",
	Usage:    "manage workspaces for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list workspaces",
			Flags: []cli.Flag{
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudWorkspacesList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get workspace",
			Flags: []cli.Flag{
				wireCloudWorkspaceIdFlag,
				wireCloudUsersFlag,
				wireCloudTabsFlag,
				wireCloudWidgetsFlag,
				wireCloudOperatorsFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudWorkspaceGet(c)
			},
		},
	},
}

var wireCloudTabsCmd = cli.Command{
	Name:     "tabs",
	Usage:    "manage tabs for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list tabs",
			Flags: []cli.Flag{
				wireCloudWorkspaceIdFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudTabsList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get tab",
			Flags: []cli.Flag{
				wireCloudWorkspaceIdFlag,
				wireCloudTabIdFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudTabGet(c)
			},
		},
	},
}

var wireCloudResourcesCmd = cli.Command{
	Name:     "macs",
	Usage:    "manage mashable application components for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list mashable application components",
			Flags: []cli.Flag{
				wireCloudWidgetFlag,
				wireCloudOperatorFlag,
				wireCloudMashupFlag,
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudResourcesList(c)
			},
		},
		{
			Name:  "get",
			Usage: "get mashable application component",
			Flags: []cli.Flag{
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudResourceGet(c)
			},
		},
		{
			Name:  "download",
			Usage: "download mashable application component",
			Flags: []cli.Flag{
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudResourceDownload(c)
			},
		},
		{
			Name:  "install",
			Usage: "install mashable application component",
			Flags: []cli.Flag{
				wireCloudFileFlag,
				wireCloudPublicFlag,
				wireCloudOverwriteFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudResourceInstall(c)
			},
		},
		{
			Name:  "uninstall",
			Usage: "uninstall mashable application component",
			Flags: []cli.Flag{
				wireCloudVenderFlag,
				wireCloudNameFlag,
				wireCloudVersionFlag,
				runFlag,
				jsonFlag,
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudResourceUninstall(c)
			},
		},
	},
}

var wireCloudPreferencesCmd = cli.Command{
	Name:     "preferences",
	Usage:    "manage preferences for WireCloud",
	Category: "APPLICATION MASHUP",
	Flags: []cli.Flag{
		hostFlag,
		tokenFlag,
	},
	Subcommands: []*cli.Command{
		{
			Name:  "get",
			Usage: "get preferences",
			Flags: []cli.Flag{
				prettyFlag,
			},
			Action: func(c *cli.Context) error {
				return wireCloudPreferencesGet(c)
			},
		},
	},
}
