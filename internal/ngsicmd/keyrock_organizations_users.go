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

type keyrockOrganizationUsersItems struct {
	UserID         string `json:"user_id"`
	OrganizationID string `json:"organization_id"`
	Role           string `json:"role"`
}

type keyrockOrganizationUsers struct {
	OrganizationUsers []keyrockOrganizationUsersItems `json:"organization_users"`
}

func orgUsersList(c *cli.Context) error {
	const funcName = "orgUsersList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("oid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	client.SetPath("/v1/organizations/" + c.String("oid") + "/users")

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
		var users keyrockOrganizationUsers
		err := ngsilib.JSONUnmarshal(body, &users)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		for _, user := range users.OrganizationUsers {
			fmt.Fprintln(ngsi.StdWriter, user.UserID)
		}
	}

	return nil
}

func orgUsersGet(c *cli.Context) error {
	const funcName = "orgUsersGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("oid") {
		return &ngsiCmdError{funcName, 3, "specify organization id", nil}
	}
	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 4, "specify user id", nil}
	}
	client.SetPath("/v1/organizations/" + c.String("oid") + "/users/" + c.String("uid") + "/organization_roles")

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

func orgUsersCreate(c *cli.Context) error {
	const funcName = "orgUsersCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("oid") {
		return &ngsiCmdError{funcName, 3, "specify organization id", nil}
	}
	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 4, "specify user id", nil}
	}
	if !c.IsSet("orid") {
		return &ngsiCmdError{funcName, 5, "specify organization role id", nil}
	}
	client.SetPath("/v1/organizations/" + c.String("oid") + "/users/" + c.String("uid") + "/organization_roles/" + c.String("orid"))

	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPost("")
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func orgUsersDelete(c *cli.Context) error {
	const funcName = "orgUsersDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("oid") {
		return &ngsiCmdError{funcName, 3, "specify organization id", nil}
	}
	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 4, "specify user id", nil}
	}
	if !c.IsSet("orid") {
		return &ngsiCmdError{funcName, 5, "specify organization role id", nil}
	}
	client.SetPath("/v1/organizations/" + c.String("oid") + "/users/" + c.String("uid") + "/organization_roles/" + c.String("orid"))

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}
