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

type keyrockRoleItmes struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	IsInternal    *bool  `json:"is_internal,omitempty"`
	OauthClientID string `json:"oauth_client_id,omitempty"`
}

type keyrockRole struct {
	Role keyrockRoleItmes `json:"role"`
}

type keyrockRoles struct {
	Roles []keyrockRoleItmes `json:"roles"`
}

func rolesList(c *cli.Context) error {
	const funcName = "rolesList"

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
	client.SetPath("/v1/applications/" + c.String("aid") + "/roles")

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
		var roles keyrockRoles
		err := ngsilib.JSONUnmarshal(body, &roles)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		for _, role := range roles.Roles {
			fmt.Fprintln(ngsi.StdWriter, role.ID)
		}
	}

	return nil
}

func rolesGet(c *cli.Context) error {
	const funcName = "rolesGet"

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
	if !c.IsSet("rid") {
		return &ngsiCmdError{funcName, 4, "specify role id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/roles/" + c.String("rid"))

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

func rolesCreate(c *cli.Context) error {
	const funcName = "rolesCreate"

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
	client.SetPath("/v1/applications/" + c.String("aid") + "/roles")

	var b []byte
	if c.IsSet("data") {
		b, err = readAll(c, ngsi)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
	} else if c.IsSet("name") {
		var role keyrockRole
		role.Role.Name = c.String("name")
		b, err = ngsilib.JSONMarshal(role)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	} else {
		return &ngsiCmdError{funcName, 6, "specify either name or data", err}
	}
	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 8, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var res keyrockRole
		err = ngsilib.JSONUnmarshal(body, &res)
		if err != nil {
			return &ngsiCmdError{funcName, 9, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, res.Role.ID)
	}
	return nil
}

func rolesUpdate(c *cli.Context) error {
	const funcName = "rolesUpdate"

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
	if !c.IsSet("rid") {
		return &ngsiCmdError{funcName, 4, "specify role id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/roles/" + c.String("rid"))

	var b []byte
	if c.IsSet("data") {
		b, err = readAll(c, ngsi)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	} else if c.IsSet("name") {
		var role keyrockRole
		role.Role.Name = c.String("name")
		b, err = ngsilib.JSONMarshal(role)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	} else {
		return &ngsiCmdError{funcName, 7, "specify either name or data", err}
	}

	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return &ngsiCmdError{funcName, 8, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 9, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 10, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func rolesDelete(c *cli.Context) error {
	const funcName = "rolesDelete"

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
	if !c.IsSet("rid") {
		return &ngsiCmdError{funcName, 4, "specify role id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/roles/" + c.String("rid"))

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}
