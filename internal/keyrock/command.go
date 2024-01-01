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

package keyrock

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewNgsiApp() *ngsicli.App {
	return KeyrockApp
}

var KeyrockApp = &ngsicli.App{
	Copyright: ngsicli.Copyright,
	Version:   ngsicli.Version,
	Usage:     "keyrock command",
	Flags:     ngsicli.GlobalFlags,
	Commands: []*ngsicli.Command{
		&ApplicationsCmd,
		&UsersCmd,
		&OrganizationsCmd,
		&ProvidersCmd,
	},
}

var ApplicationsCmd = ngsicli.Command{
	Name:     "applications",
	Usage:    "manage applications for Keyrock",
	Category: "Keyrock",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list applications",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return applicationsList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockApplicationIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return applicationsGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
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
				keyrcokApplicationScopeFlag,
				keyrcokApplicationOpenIDFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			OptionFlags: &ngsicli.ValidationFlag{Mode: ngsicli.XnorCondition, Flags: []string{"data", "name"}},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return applicationsCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockApplicationIDRFlag,
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
				keyrcokApplicationScopeFlag,
				keyrcokApplicationOpenIDFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return applicationsUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockApplicationIDRFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return applicationsDelete(c, ngsi, client)
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

var appsOrgsCmd = ngsicli.Command{
	Name:     "organizations",
	Usage:    "manage authorized organizations in an application for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockApplicationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list organizations in an application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsOrgsRolesList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get roles of an organization in an application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockOrganizationIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "oid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsOrgsRolesGet(c, ngsi, client)
			},
		},
		{
			Name:       "assign",
			Usage:      "assign a role to an organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockOrganizationIDRFlag,
				keyrockRolesIDRFlag,
				keyrcokOrganizationRoleIDRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "oid", "rid", "orid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsOrgsRolesAssign(c, ngsi, client)
			},
		},
		{
			Name:       "unassign",
			Usage:      "delete a role assignment from an organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockOrganizationIDRFlag,
				keyrockRolesIDRFlag,
				keyrcokOrganizationRoleIDRFlag,
			},
			RequiredFlags: []string{"aid", "oid", "rid", "orid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsOrgsRolesUnassign(c, ngsi, client)
			},
		},
	},
}

var iotAgentCmd = ngsicli.Command{
	Name:     "iota",
	Usage:    "manage IoT Agents for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockApplicationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list iot agents",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return iotAgentsList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get iot agent",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockIotAgentsIDRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "iid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return iotAgentsGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create iot agent",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return iotAgentsCreate(c, ngsi, client)
			},
		},
		{
			Name:       "reset",
			Usage:      "reset iot agent",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockIotAgentsIDRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "iid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return iotAgentsReset(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete iot agent",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockIotAgentsIDRFlag,
			},
			RequiredFlags: []string{"aid", "iid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return iotAgentsDelete(c, ngsi, client)
			},
		},
	},
}

var OrganizationsCmd = ngsicli.Command{
	Name:     "organizations",
	Usage:    "manage organizations for Keyrock",
	Category: "Keyrock",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list organizations",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return organizationsList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockOrganizationIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"oid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return organizationsGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockOrganizationNameRFlag,
				keyrockDescriptionFlag,
				keyrockWebsiteFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"name"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return organizationsCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockOrganizationIDRFlag,
				keyrockOrganizationNameFlag,
				keyrockDescriptionFlag,
				keyrockWebsiteFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"oid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return organizationsUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockOrganizationIDRFlag,
			},
			RequiredFlags: []string{"oid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return organizationsDelete(c, ngsi, client)
			},
		},
		&orgsUsersCmd,
	},
}

var orgsUsersCmd = ngsicli.Command{
	Name:     "users",
	Usage:    "manage users of an organization for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockOrganizationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list users of an organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"oid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return orgUsersList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get info of user organization relationship",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"oid", "uid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return orgUsersGet(c, ngsi, client)
			},
		},
		{
			Name:       "add",
			Usage:      "add a user to an organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				keyrcokOrganizationRoleIDRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"oid", "uid", "orid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return orgUsersCreate(c, ngsi, client)
			},
		},
		{
			Name:       "remove",
			Usage:      "remove a user from an organization",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				keyrcokOrganizationRoleIDRFlag,
			},
			RequiredFlags: []string{"oid", "uid", "orid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return orgUsersDelete(c, ngsi, client)
			},
		},
	},
}
var UsersCmd = ngsicli.Command{
	Name:     "users",
	Usage:    "manage users for Keyrock",
	Category: "Keyrock",
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list users",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return usersList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get user",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"uid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return usersGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create user",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserNameRFlag,
				keyrockPasswordRFlag,
				keyrockEmailRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"username", "email", "password"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return usersCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update user",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				keyrockUserNameFlag,
				keyrockPasswordFlag,
				keyrockEmailFlag,
				keyrockGravatarFlag,
				keyrockDescriptionFlag,
				keyrockWebsiteFlag,
				keyrockExtraFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"uid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return usersUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete user",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
			},
			RequiredFlags: []string{"uid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return usersDelete(c, ngsi, client)
			},
		},
	},
}

var rolesCmd = ngsicli.Command{
	Name:     "roles",
	Usage:    "manage roles for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockApplicationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list roles",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return rolesList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get a role",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockRolesIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "rid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return rolesGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create a role",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockRoleDataFlag,
				keyrockRoleNameFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			OptionFlags:   &ngsicli.ValidationFlag{Mode: ngsicli.XnorCondition, Flags: []string{"data", "name"}},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return rolesCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update a role",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockRolesIDRFlag,
				keyrockRoleDataFlag,
				keyrockRoleNameFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "rid"},
			OptionFlags:   &ngsicli.ValidationFlag{Mode: ngsicli.XnorCondition, Flags: []string{"data", "name"}},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return rolesUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete a role",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockRolesIDRFlag,
			},
			RequiredFlags: []string{"aid", "rid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return rolesDelete(c, ngsi, client)
			},
		},
		{
			Name:       "permissions",
			Usage:      "list permissions associated to a role",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockRolesIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "rid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsRolePermList(c, ngsi, client)
			},
		},
		{
			Name:       "assign",
			Usage:      "Assign a permission to a role",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockRolesIDRFlag,
				keyrockPermissionIDRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "rid", "pid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsRolePermAssign(c, ngsi, client)
			},
		},
		{
			Name:       "unassign",
			Usage:      "delete a permission from a role",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockRolesIDRFlag,
				keyrockPermissionIDRFlag,
			},
			RequiredFlags: []string{"aid", "rid", "pid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsRolePermDelete(c, ngsi, client)
			},
		},
	},
}

var permissionsCmd = ngsicli.Command{
	Name:     "permissions",
	Usage:    "manage permissions for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockApplicationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list permissions",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockPermissionIDFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return permissionsList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get permission",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockPermissionIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "pid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return permissionsGet(c, ngsi, client)
			},
		},
		{
			Name:       "create",
			Usage:      "create permission",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrcokPermissionDataFlag,
				keyrockPermissionNameFlag,
				keyrockPermissionDescriptionFlag,
				keyrockPermissionActionFlag,
				keyrockPermissionResourceFlag,
				keyrockPermissionRegexFlag,
				keyrockPermissionServiceHeaderFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return permissionsCreate(c, ngsi, client)
			},
		},
		{
			Name:       "update",
			Usage:      "update permission",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockPermissionIDRFlag,
				keyrcokPermissionDataFlag,
				keyrockPermissionNameFlag,
				keyrockPermissionDescriptionFlag,
				keyrockPermissionActionFlag,
				keyrockPermissionResourceFlag,
				keyrockPermissionRegexFlag,
				keyrockPermissionServiceHeaderFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "pid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return permissionsUpdate(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete permission",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockPermissionIDRFlag,
			},
			RequiredFlags: []string{"aid", "pid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return permissionsDelete(c, ngsi, client)
			},
		},
	},
}

var pepProxiesCmd = ngsicli.Command{
	Name:     "pep",
	Usage:    "manage PEP Proxies for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockApplicationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list pep proxy",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.PrettyFlag,
			},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return pepProxiesList(c, ngsi, client)
			},
			RequiredFlags: []string{"aid"},
		},
		{
			Name:       "create",
			Usage:      "create pep proxy",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.RunFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return pepProxiesCreate(c, ngsi, client)
			},
		},
		{
			Name:       "reset",
			Usage:      "reset pep proxy",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.RunFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return pepProxiesReset(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete pep proxy",
			ServerList: []string{"keyrock"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return pepProxiesDelete(c, ngsi, client)
			},
			RequiredFlags: []string{"aid"},
			Flags: []ngsicli.Flag{
				ngsicli.RunFlag,
			},
		},
	},
}

var appsUsersCmd = ngsicli.Command{
	Name:     "users",
	Usage:    "manage authorized users in an application for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockApplicationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list users in an application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsUsersList(c, ngsi, client)
			},
		},
		{
			Name:       "get",
			Usage:      "get roles of a user in an application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "uid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsUsersGet(c, ngsi, client)
			},
		},
		{
			Name:       "assign",
			Usage:      "assign a role to a user",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				keyrockRolesIDRFlag,
				ngsicli.VerboseFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "uid", "rid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsUsersAssign(c, ngsi, client)
			},
		},
		{
			Name:       "unassign",
			Usage:      "delete a role assignment from a user",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockUserIDRFlag,
				keyrockRolesIDRFlag,
			},
			RequiredFlags: []string{"aid", "uid", "rid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return appsUsersUnassign(c, ngsi, client)
			},
		},
	},
}

var trustedAppsCmd = ngsicli.Command{
	Name:     "trusted",
	Usage:    "manage trusted applications for Keyrock",
	Category: "sub-command",
	Flags: []ngsicli.Flag{
		keyrockApplicationIDRFlag,
	},
	Subcommands: []*ngsicli.Command{
		{
			Name:       "list",
			Usage:      "list trusted applications associated to an application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return trustedAppList(c, ngsi, client)
			},
		},
		{
			Name:       "add",
			Usage:      "add trusted application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockTrustedIDRFlag,
				ngsicli.PrettyFlag,
			},
			RequiredFlags: []string{"aid", "tid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return trustedAppAdd(c, ngsi, client)
			},
		},
		{
			Name:       "delete",
			Usage:      "delete trusted application",
			ServerList: []string{"keyrock"},
			Flags: []ngsicli.Flag{
				keyrockTrustedIDRFlag,
			},
			RequiredFlags: []string{"aid", "tid"},
			Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
				return trustedAppDelete(c, ngsi, client)
			},
		},
	},
}

var ProvidersCmd = ngsicli.Command{
	Name:       "providers",
	Usage:      "print service providers for Keyrock",
	Category:   "Keyrock",
	ServerList: []string{"keyrock"},
	Flags: []ngsicli.Flag{
		ngsicli.HostRFlag,
		ngsicli.OAuthTokenFlag,
		ngsicli.PrettyFlag,
	},
	Action: func(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
		return providersGet(c, ngsi, client)
	},
}
