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

type keyrockRoleUserAssignmentItems struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

type keyrockRoleUserAssignments struct {
	RoleUserAssignments []keyrockRoleUserAssignmentItems `json:"role_user_assignments"`
}

func appsUsersList(c *cli.Context) error {
	const funcName = "appsUsersList"

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
	path := fmt.Sprintf("/v1/applications/%s/users", c.String("aid"))
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
		var users keyrockRoleUserAssignments
		err := ngsilib.JSONUnmarshal(body, &users)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		for _, user := range users.RoleUserAssignments {
			fmt.Fprintln(ngsi.StdWriter, user.UserID)
		}
	}

	return nil
}

func appsUsersGet(c *cli.Context) error {
	const funcName = "appsUsersGet"

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
	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 4, "specify user id", nil}
	}
	path := fmt.Sprintf("/v1/applications/%s/users/%s/roles", c.String("aid"), c.String("uid"))
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 7, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var users keyrockRoleUserAssignments
		err := ngsilib.JSONUnmarshal(body, &users)
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		for _, user := range users.RoleUserAssignments {
			fmt.Fprintln(ngsi.StdWriter, user.RoleID)
		}
	}

	return nil
}

func appsUsersAssign(c *cli.Context) error {
	const funcName = "appsUsersAssign"

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
	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 4, "specify user id", nil}
	}
	if !c.IsSet("rid") {
		return &ngsiCmdError{funcName, 5, "specify role id", nil}
	}
	path := fmt.Sprintf("/v1/applications/%s/users/%s/roles/%s", c.String("aid"), c.String("uid"), c.String("rid"))
	client.SetPath(path)

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

func appsUsersUnassign(c *cli.Context) error {
	const funcName = "appsUsersUnassign"

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
	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 4, "specify user id", nil}
	}
	if !c.IsSet("rid") {
		return &ngsiCmdError{funcName, 5, "specify role id", nil}
	}
	path := fmt.Sprintf("/v1/applications/%s/users/%s/roles/%s", c.String("aid"), c.String("uid"), c.String("rid"))
	client.SetPath(path)

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}
