/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"bytes"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
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

func rolesList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "rolesList"

	client.SetPath("/v1/applications/" + c.String("aid") + "/roles")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var roles keyrockRoles
		err := ngsilib.JSONUnmarshal(body, &roles)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		for _, role := range roles.Roles {
			fmt.Fprintln(ngsi.StdWriter, role.ID)
		}
	}

	return nil
}

func rolesGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "rolesGet"

	client.SetPath("/v1/applications/" + c.String("aid") + "/roles/" + c.String("rid"))

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func rolesCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "rolesCreate"

	client.SetPath("/v1/applications/" + c.String("aid") + "/roles")

	var b []byte
	var err error

	if c.IsSet("data") {
		b, err = ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
	}
	if c.IsSet("name") {
		r, err := getRoleByName(ngsi, client, c.String("aid"), c.String("name"))
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if r != (*keyrockRoleItmes)(nil) {
			kr := keyrockRole{Role: *r}
			b, err := ngsilib.JSONMarshal(kr)
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
			err = printRole(c, ngsi, b)
			if err != nil {
				return ngsierr.New(funcName, 4, err.Error(), err)
			}
			return nil
		}
		var role keyrockRole
		role.Role.Name = c.String("name")
		b, err = ngsilib.JSONMarshal(role)
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
	}
	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 6, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated {
		return ngsierr.New(funcName, 7, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}
	err = printRole(c, ngsi, body)
	if err != nil {
		return ngsierr.New(funcName, 8, err.Error(), err)
	}

	return nil
}

func printRole(c *ngsicli.Context, ngsi *ngsilib.NGSI, body []byte) error {
	const funcName = "printRole"

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 1, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var res keyrockRole
		err := ngsilib.JSONUnmarshal(body, &res)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, res.Role.ID)
	}
	return nil
}

func rolesUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "rolesUpdate"

	client.SetPath("/v1/applications/" + c.String("aid") + "/roles/" + c.String("rid"))

	var b []byte
	var err error

	if c.IsSet("data") {
		b, err = ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
	}
	if c.IsSet("name") {
		var role keyrockRole
		role.Role.Name = c.String("name")
		b, err = ngsilib.JSONMarshal(role)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}

	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func rolesDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "rolesDelete"

	client.SetPath("/v1/applications/" + c.String("aid") + "/roles/" + c.String("rid"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}

func getRoleByName(ngsi *ngsilib.NGSI, client *ngsilib.Client, aid string, name string) (*keyrockRoleItmes, error) {
	const funcName = "getRoleByName"

	client.SetPath("/v1/applications/" + aid + "/roles")

	res, body, err := client.HTTPGet()
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	var roles keyrockRoles
	err = ngsilib.JSONUnmarshal(body, &roles)
	if err != nil {
		return nil, ngsierr.New(funcName, 3, err.Error(), err)
	}
	for _, role := range roles.Roles {
		if role.Name == name {
			return &role, nil
		}
	}

	return nil, nil
}
