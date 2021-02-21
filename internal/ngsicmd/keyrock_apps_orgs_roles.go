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
	"bytes"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type keyrockRoleOrganizationAssignmentsItems struct {
	OrganizationID   string `json:"organization_id"`
	RoleOrganization string `json:"role_organization"`
	RoleID           string `json:"role_id"`
}

type keyrockRoleOrganizationAssignments struct {
	RoleOrganizationAssignments []keyrockRoleOrganizationAssignmentsItems `json:"role_organization_assignments"`
}

func appsOrgsRolesList(c *cli.Context) error {
	const funcName = "appsOrgsRolesList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	path := fmt.Sprintf("/v1/applications/%s/organizations", c.String("aid"))
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 6, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var orgs keyrockRoleOrganizationAssignments
		err := ngsilib.JSONUnmarshal(body, &orgs)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		for _, org := range orgs.RoleOrganizationAssignments {
			fmt.Fprintln(ngsi.StdWriter, org.OrganizationID)
		}
	}

	return nil
}

func appsOrgsRolesGet(c *cli.Context) error {
	const funcName = "appsOrgsRolesGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	if !c.IsSet("oid") {
		return &ngsiCmdError{funcName, 4, "specify organization id", nil}
	}
	path := fmt.Sprintf("/v1/applications/%s/organizations/%s/roles", c.String("aid"), c.String("oid"))
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func appsOrgsRolesAssign(c *cli.Context) error {
	const funcName = "appsOrgsRolesAssign"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	if !c.IsSet("oid") {
		return &ngsiCmdError{funcName, 4, "specify organization id", nil}
	}
	if !c.IsSet("rid") {
		return &ngsiCmdError{funcName, 5, "specify role id", nil}
	}
	if !c.IsSet("orid") {
		return &ngsiCmdError{funcName, 6, "specify organization role id", nil}
	}
	path := fmt.Sprintf("/v1/applications/%s/organizations/%s/roles/%s/organization_roles/%s", c.String("aid"), c.String("oid"), c.String("rid"), c.String("orid"))
	client.SetPath(path)

	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPost("")
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 9, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func appsOrgsRolesUnassign(c *cli.Context) error {
	const funcName = "appsOrgsRolesUnassign"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	if !c.IsSet("oid") {
		return &ngsiCmdError{funcName, 4, "specify organization id", nil}
	}
	if !c.IsSet("rid") {
		return &ngsiCmdError{funcName, 5, "specify role id", nil}
	}
	if !c.IsSet("orid") {
		return &ngsiCmdError{funcName, 6, "specify organization role id", nil}
	}
	path := fmt.Sprintf("/v1/applications/%s/organizations/%s/roles/%s/organization_roles/%s", c.String("aid"), c.String("oid"), c.String("rid"), c.String("orid"))
	client.SetPath(path)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}
