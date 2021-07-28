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

package management

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func serverList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "serverList"

	host := c.String("host")

	if host != "" {
		info, err := ngsi.AllServersList().ServerInfo(host, ngsicli.CmdMode)
		if err != nil {
			return ngsierr.New(funcName, 1, host+" not found", err)
		}
		clearText := c.Bool("clearText")
		printServerInfo(ngsi, info, clearText)
	} else {
		if c.IsSet("json") || c.Bool("pretty") {
			lists, err := ngsi.AllServersList().ServerInfoJSON("", ngsicli.CmdMode)
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error(), err)
			}
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, []byte(*lists), "", "  ")
				if err != nil {
					return ngsierr.New(funcName, 3, err.Error(), err)
				}
				fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			} else {
				fmt.Fprint(ngsi.StdWriter, *lists)
			}
		} else {
			info := ngsi.AllServersList().ServerList(ngsicli.CmdMode, c.Bool("all"))
			list := info.List(c.Bool("singleLine"))
			fmt.Fprintln(ngsi.StdWriter, list)
		}
	}

	setPreviousArgs(ngsi, host)

	return nil
}

func serverGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "serverGet"

	host := c.String("host")

	if c.IsSet("json") || c.Bool("pretty") {
		lists, err := ngsi.AllServersList().ServerInfoJSON(host, ngsicli.CmdMode)
		if err != nil {
			return ngsierr.New(funcName, 1, host+" not found", err)
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, []byte(*lists), "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, *lists)
		}
	} else {
		info, err := ngsi.AllServersList().ServerInfo(host, ngsicli.CmdMode)
		if err != nil {
			return ngsierr.New(funcName, 3, host+" not found", err)
		}
		clearText := c.Bool("clearText")
		printServerInfo(ngsi, info, clearText)
	}

	setPreviousArgs(ngsi, host)

	return nil
}

func serverAdd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "serverAdd"

	host := c.String("host")

	if !ngsilib.IsNameString(host) {
		return ngsierr.New(funcName, 1, "name error "+host, nil)
	}
	if ngsi.ExistsBrokerHost(host) {
		return ngsierr.New(funcName, 2, host+" already exists", nil)
	}

	if !c.IsSet("serverHost") && ngsicli.CmdMode == "" {
		return ngsierr.New(funcName, 3, "serverHost is missing", nil)
	}

	if !c.IsSet("serverType") {
		return ngsierr.New(funcName, 4, "serverType is missing", nil)
	}

	serverType := ngsicli.CmdMode
	if c.IsSet("serverType") {
		serverType = strings.ToLower(c.String("serverType"))
	}
	if !ngsilib.Contains(ngsi.ServerTypeArgs(), serverType) {
		return ngsierr.New(funcName, 5, "serverType error: "+serverType+" (Comet, Cygnus, Iota, Keyrock, Perseo, QuantumLeap, WireCloud, Queryproxy, Regproxy, Tokenproxy)", nil)
	}

	param := make(map[string]string)
	args := ngsi.ServerInfoArgs()
	for i := 0; i < len(args); i++ {
		key := args[i]
		if c.IsSet(key) {
			value := c.String(key)
			if key == "service" || key == "serverType" {
				value = strings.ToLower(value)
			}
			if value != "" {
				param[key] = value
			}
		}
	}

	if serverType == "keyrock" {
		param["idmType"] = "idm"
		param["idmHost"] = strings.TrimRight(param["serverHost"], "/") + "/v1/auth/tokens"
	}

	setPreviousArgs(ngsi, host)

	err := ngsi.CreateServer(host, param)
	if err != nil {
		ngsi.Updated = false
		return ngsierr.New(funcName, 6, err.Error(), err)
	}

	return nil
}

func serverUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "serverUpdate"

	host := c.String("host")

	if !ngsi.ExistsBrokerHost(host) {
		return ngsierr.New(funcName, 1, host+" not found", nil)
	}

	param := make(map[string]string)
	args := ngsi.ServerInfoArgs()
	for i := 0; i < len(args); i++ {
		key := args[i]
		if c.IsSet(key) {
			value := c.String(key)
			if key == "service" {
				value = strings.ToLower(value)
			}
			if value != "" {
				param[key] = value
			}
		}
	}

	setPreviousArgs(ngsi, host)

	err := ngsi.UpdateServer(host, param)
	if err != nil {
		ngsi.Updated = false
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	return nil
}

func serverDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "serverDelete"

	host := c.String("host")

	if !ngsi.ExistsBrokerHost(host) {
		return ngsierr.New(funcName, 1, host+" not found", nil)
	}

	if c.IsSet("items") {
		items := c.String("items")
		for _, item := range strings.Split(items, ",") {
			if err := ngsi.DeleteItem(host, item); err != nil {
				return ngsierr.New(funcName, 2, "delete error - "+item, err)
			}
		}
		err := ngsi.UpdateServer(host, nil)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	} else {
		err := ngsi.IsHostReferenced(host)
		if err != nil {
			return ngsierr.New(funcName, 4, host+" is referenced", err)
		}

		setPreviousArgs(ngsi, host)

		err = ngsi.DeleteServer(host)
		if err != nil {
			ngsi.Updated = false
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
	}
	return nil
}

func printServerInfo(ngsi *ngsilib.NGSI, info *ngsilib.Server, clearText bool) {
	if info.ServerType == "broker" {
		fmt.Fprintln(ngsi.StdWriter, "server type error")
		return
	}
	fmt.Fprintln(ngsi.StdWriter, "serverType "+info.ServerType)
	fmt.Fprintln(ngsi.StdWriter, "serverHost "+info.ServerHost)
	if info.Tenant != "" {
		fmt.Fprintln(ngsi.StdWriter, "FIWARE-Service "+info.Tenant)
	}
	if info.Scope != "" {
		fmt.Fprintln(ngsi.StdWriter, "FIWARE-ServicePath "+info.Scope)
	}

	items := []struct {
		key   string
		value string
	}{
		{"Context", info.Context},
		{"SafeString", info.SafeString},
		{"IdmType", info.IdmType},
		{"IdmHost", info.IdmHost},
		{"Username", info.Username},
		{"Password", obfuscateText(info.Password, clearText)},
		{"ClientID", obfuscateText(info.ClientID, clearText)},
		{"ClientSecret", obfuscateText(info.ClientSecret, clearText)},
		{"HeaderName", info.HeaderName},
		{"HeaderValue", obfuscateText(info.HeaderValue, clearText)},
		{"HeaderEnvValue", info.HeaderEnvValue},
		{"TokenScope", info.TokenScope},
		{"XAuthToken", info.XAuthToken},
		{"Token", info.Token},
		{"APIPath", info.APIPath},
	}

	for _, item := range items {
		if item.value != "" {
			fmt.Fprintf(ngsi.StdWriter, "%s %s\n", item.key, item.value)
		}
	}
}
