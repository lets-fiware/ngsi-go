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

package keyrock

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

// Keyrock Common
var (
	keyrockDescriptionFlag = &ngsicli.StringFlag{
		Name:    "description",
		Aliases: []string{"d"},
		Usage:   "description",
	}
	keyrockWebsiteFlag = &ngsicli.StringFlag{
		Name:    "website",
		Aliases: []string{"w"},
		Usage:   "website",
	}
)

// Keyrock users
var (
	keyrockUserIDRFlag = &ngsicli.StringFlag{
		Name:     "uid",
		Aliases:  []string{"i"},
		Usage:    "user id",
		Required: true,
	}
	keyrockUserNameFlag = &ngsicli.StringFlag{
		Name:    "username",
		Aliases: []string{"u"},
		Usage:   "user name",
	}
	keyrockUserNameRFlag = &ngsicli.StringFlag{
		Name:     "username",
		Aliases:  []string{"u"},
		Usage:    "user name",
		Required: true,
	}
	keyrockPasswordFlag = &ngsicli.StringFlag{
		Name:    "password",
		Aliases: []string{"p"},
		Usage:   "password",
	}
	keyrockPasswordRFlag = &ngsicli.StringFlag{
		Name:     "password",
		Aliases:  []string{"p"},
		Usage:    "password",
		Required: true,
	}
	keyrockEmailFlag = &ngsicli.StringFlag{
		Name:    "email",
		Aliases: []string{"e"},
		Usage:   "email",
	}
	keyrockEmailRFlag = &ngsicli.StringFlag{
		Name:     "email",
		Aliases:  []string{"e"},
		Usage:    "email",
		Required: true,
	}
	keyrockGravatarFlag = &ngsicli.BoolFlag{
		Name:    "gravatar",
		Aliases: []string{"g"},
		Usage:   "gravatar",
	}
	keyrockExtraFlag = &ngsicli.StringFlag{
		Name:    "extra",
		Aliases: []string{"E"},
		Usage:   "extra information",
	}
)

// Keyrock applications
var (
	keyrockApplicationIDRFlag = &ngsicli.StringFlag{
		Name:     "aid",
		Aliases:  []string{"i"},
		Usage:    "application id",
		Required: true,
	}
	keyrcokApplicationDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "application data",
	}
	keyrcokApplicationNameFlag = &ngsicli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "application name",
	}
	keyrcokApplicationDescriptionFlag = &ngsicli.StringFlag{
		Name:    "description",
		Aliases: []string{"D"},
		Usage:   "description",
	}
	keyrcokApplicationRedirectURIFlag = &ngsicli.StringFlag{
		Name:    "redirectUri",
		Aliases: []string{"R"},
		Usage:   "redirect uri",
	}
	keyrcokApplicationRedirectSignOutURIFlag = &ngsicli.StringFlag{
		Name:    "redirectSignOutUri",
		Aliases: []string{"S"},
		Usage:   "redirect redirect sign out uri",
	}
	keyrcokApplicationURLFlag = &ngsicli.StringFlag{
		Name:    "url",
		Aliases: []string{"u"},
		Usage:   "url",
	}
	keyrcokApplicationGrantTypeFlag = &ngsicli.StringFlag{
		Name:    "grantType",
		Aliases: []string{"g"},
		Usage:   "grant type",
	}
	keyrcokApplicationTokenTypesFlag = &ngsicli.StringFlag{
		Name:    "tokenTypes",
		Aliases: []string{"t"},
		Usage:   "token types",
	}
	keyrcokApplicationResponseTypeFlag = &ngsicli.StringFlag{
		Name:    "responseType",
		Aliases: []string{"r"},
		Usage:   "response type",
	}
	keyrcokApplicationClientTypeFlag = &ngsicli.StringFlag{
		Name:    "clientType",
		Aliases: []string{"c"},
		Usage:   "client type",
	}
)

// Keyrock organizations
var (
	keyrockOrganizationIDRFlag = &ngsicli.StringFlag{
		Name:     "oid",
		Aliases:  []string{"o"},
		Usage:    "organization id",
		Required: true,
	}
	keyrockOrganizationNameFlag = &ngsicli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "organization name",
	}
	keyrockOrganizationNameRFlag = &ngsicli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "organization name",
		Required: true,
	}
	keyrcokOrganizationRoleIDRFlag = &ngsicli.StringFlag{
		Name:     "orid",
		Aliases:  []string{"c"},
		Usage:    "organization role id",
		Required: true,
	}
)

// Keyrock roles
var (
	keyrockRolesIDRFlag = &ngsicli.StringFlag{
		Name:     "rid",
		Aliases:  []string{"r"},
		Usage:    "role id",
		Required: true,
	}
	keyrockRoleDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "role data",
	}
	keyrockRoleNameFlag = &ngsicli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "role name",
	}
)

// Keyrock permissions
var (
	keyrockPermissionIDFlag = &ngsicli.StringFlag{
		Name:    "pid",
		Aliases: []string{"p"},
		Usage:   "permission id",
	}
	keyrockPermissionIDRFlag = &ngsicli.StringFlag{
		Name:     "pid",
		Aliases:  []string{"p"},
		Usage:    "permission id",
		Required: true,
	}
	keyrcokPermissionDataFlag = &ngsicli.StringFlag{
		Name:    "data",
		Aliases: []string{"d"},
		Usage:   "permissionrole data",
	}
	keyrockPermissionNameFlag = &ngsicli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "permission name",
	}
	keyrockPermissionDescriptionFlag = &ngsicli.StringFlag{
		Name:    "description",
		Aliases: []string{"D"},
		Usage:   "description",
	}
	keyrockPermissionActionFlag = &ngsicli.StringFlag{
		Name:    "action",
		Aliases: []string{"a"},
		Usage:   "action",
	}
	keyrockPermissionResourceFlag = &ngsicli.StringFlag{
		Name:    "resource",
		Aliases: []string{"r"},
		Usage:   "resoruce",
	}
)

// Keyrock Iot Agents
var (
	keyrockIotAgentsIDRFlag = &ngsicli.StringFlag{
		Name:     "iid",
		Aliases:  []string{"i"},
		Usage:    "IoT Agent id",
		Required: true,
	}
)

// Keyrock Trusted application
var (
	keyrockTrustedIDRFlag = &ngsicli.StringFlag{
		Name:     "tid",
		Aliases:  []string{"t"},
		Usage:    "trusted application id",
		Required: true,
	}
)
