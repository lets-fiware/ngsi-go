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
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type keyrockApplicationItems struct {
	ID                 string      `json:"id,omitempty"`
	Name               string      `json:"name,omitempty"`
	Description        string      `json:"description,omitempty"`
	Image              string      `json:"image,omitempty"`
	URL                string      `json:"url,omitempty"`
	RedirectURI        string      `json:"redirect_uri,omitempty"`
	RedirectSignOutUri string      `json:"redirect_sign_out_uri"`
	GrantType          interface{} `json:"grant_type,omitempty"`
	ResponseType       interface{} `json:"response_type,omitempty"`
	TokenTypes         interface{} `json:"token_types,omitempty"`
	ClientType         interface{} `json:"client_type,omitempty"`
	Scope              interface{} `json:"scope,omitempty"`
}

type keyrockApplication struct {
	Application keyrockApplicationItems `json:"application"`
}

type keyrockApplications struct {
	Applications []keyrockApplicationItems `json:"applications"`
}

func applicationsList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "applicationsList"

	client.SetPath("/v1/applications")

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
		var apps keyrockApplications
		err := ngsilib.JSONUnmarshal(body, &apps)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		for _, app := range apps.Applications {
			fmt.Fprintln(ngsi.StdWriter, app.ID)
		}
	}

	return nil
}

func applicationsGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "applicationsGet"

	client.SetPath("/v1/applications/" + c.String("aid"))

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

func applicationsCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "applicationsCreate"

	b, err := makeAppBody(c, ngsi, false)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/v1/applications")

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
		var res keyrockApplication
		err = ngsilib.JSONUnmarshal(body, &res)
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, res.Application.ID)
	}

	return nil
}

func applicationsUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "applicationsUpdate"

	client.SetPath("/v1/applications/" + c.String("aid"))

	b, err := makeAppBody(c, ngsi, true)
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

func applicationsDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "applicationsDelete"

	client.SetPath("/v1/applications/" + c.String("aid"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}

func makeAppBody(c *ngsicli.Context, ngsi *ngsilib.NGSI, update bool) ([]byte, error) {
	const funcName = "makeAppBody"

	if c.IsSet("data") {
		b, err := ngsi.ReadAll(c.String("data"))
		if err != nil {
			return nil, ngsierr.New(funcName, 1, err.Error(), err)
		}
		return b, nil
	}

	var app keyrockApplication

	app.Application.Name = c.String("name")
	app.Application.Description = c.String("description")
	if app.Application.Description == "" && !update {
		app.Application.Description = app.Application.Name
	}
	app.Application.URL = c.String("url")
	if app.Application.URL == "" && !update {
		app.Application.URL = "http://localhost"
	}
	app.Application.RedirectURI = c.String("redirectUri")
	if app.Application.RedirectURI == "" && !update {
		app.Application.RedirectURI = "http://localhost"
	}
	app.Application.RedirectSignOutUri = c.String("redirectSignOutUri")
	s := c.String("grantType")
	if s != "" {
		app.Application.GrantType = strings.Split(s, ",")
	} else {
		if !update {
			app.Application.GrantType = []string{"client_credentials", "password", "implicit", "authorization_code", "refresh_token"}
		}
	}
	s = c.String("tokenTypes")
	if s != "" {
		app.Application.TokenTypes = strings.Split(s, ",")
	}
	s = c.String("responseType")
	if s != "" {
		app.Application.ResponseType = strings.Split(s, ",")
	}
	s = c.String("clientType")
	if s != "" {
		app.Application.ClientType = strings.Split(s, ",")
	}
	s = c.String("scope")
	if s != "" {
		app.Application.Scope = strings.Split(s, ",")
	}
	if c.IsSet("openid") {
		app.Application.Scope = "openid"
		app.Application.GrantType = []string{"client_credentials", "password", "implicit", "authorization_code", "refresh_token"}
		app.Application.TokenTypes = []string{"jwt"}
	}

	b, err := ngsilib.JSONMarshal(app)
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}

	return b, nil
}
