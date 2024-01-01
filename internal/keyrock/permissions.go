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
	"bytes"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type keyrockPermissionItems struct {
	ID                            string      `json:"id,omitempty"`
	Name                          string      `json:"name,omitempty"`
	Description                   string      `json:"description,omitempty"`
	IsInternal                    *bool       `json:"is_internal,omitempty"`
	Action                        string      `json:"action,omitempty"`
	Resource                      string      `json:"resource,omitempty"`
	AuthorizationServiceHeader    string      `json:"authorization_service_header,omitempty"`
	UseAuthorizationServiceHeader interface{} `json:"use_authorization_service_header,omitempty"`
	IsRegex                       interface{} `json:"is_regex,omitempty"`
	XML                           string      `json:"xml,omitempty"`
	OauthClientID                 string      `json:"oauth_client_id,omitempty"`
}

type keyrockPermission struct {
	Permission keyrockPermissionItems `json:"permission,omitempty"`
}

type keyrockPermissions struct {
	Permissions []keyrockPermissionItems `json:"permissions,omitempty"`
}

func permissionsList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "permissionsList"

	client.SetPath("/v1/applications/" + c.String("aid") + "/permissions")

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
		var permissions keyrockPermissions
		err := ngsilib.JSONUnmarshal(body, &permissions)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		for _, app := range permissions.Permissions {
			fmt.Fprintln(ngsi.StdWriter, app.ID)
		}
	}

	return nil
}

func permissionsGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "permissionsGet"

	client.SetPath("/v1/applications/" + c.String("aid") + "/permissions/" + c.String("pid"))

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

func permissionsCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "permissionsCreate"

	client.SetPath("/v1/applications/" + c.String("aid") + "/permissions")

	b, err := makePermissionBody(c, ngsi)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated {
		return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 4, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var res keyrockPermission
		err = ngsilib.JSONUnmarshal(body, &res)
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, res.Permission.ID)
	}

	return nil
}

func permissionsUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "permissionsUpdate"

	client.SetPath("/v1/applications/" + c.String("aid") + "/permissions/" + c.String("pid"))

	b, err := makePermissionBody(c, ngsi)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func permissionsDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "permissionsDelete"

	client.SetPath("/v1/applications/" + c.String("aid") + "/permissions/" + c.String("pid"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}

func makePermissionBody(c *ngsicli.Context, ngsi *ngsilib.NGSI) ([]byte, error) {
	const funcName = "makeAppBody"

	if c.IsSet("data") {
		b, err := ngsi.ReadAll(c.String("data"))
		if err != nil {
			return nil, ngsierr.New(funcName, 1, err.Error(), err)
		}
		return b, nil
	}

	var perm keyrockPermission

	perm.Permission.Name = c.String("name")
	perm.Permission.Description = c.String("description")
	perm.Permission.Action = c.String("action")
	perm.Permission.Resource = c.String("resource")

	if c.IsSet("regex") {
		perm.Permission.IsRegex = c.Bool(("regex"))
	}

	if c.IsSet("serviceHeader") {
		if c.String("serviceHeader") == "" {
			perm.Permission.UseAuthorizationServiceHeader = false
		} else {
			perm.Permission.UseAuthorizationServiceHeader = true
			perm.Permission.AuthorizationServiceHeader = c.String(("serviceHeader"))
		}

	}

	b, err := ngsilib.JSONMarshal(perm)
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}

	return b, nil
}
