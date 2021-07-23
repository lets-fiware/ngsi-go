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
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func serverList(c *cli.Context) error {
	const funcName = "serverList"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")

	if host != "" {
		info, err := ngsi.AllServersList().ServerInfo(host, gCmdMode)
		if err != nil {
			return &ngsiCmdError{funcName, 2, host + " not found", err}
		}
		clearText := c.Bool("clearText")
		printServerInfo(ngsi, info, clearText)
	} else {
		if c.IsSet("json") || c.Bool("pretty") {
			lists, err := ngsi.AllServersList().ServerInfoJSON("", gCmdMode)
			if err != nil {
				return &ngsiCmdError{funcName, 3, err.Error(), err}
			}
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, []byte(*lists), "", "  ")
				if err != nil {
					return &ngsiCmdError{funcName, 4, err.Error(), err}
				}
				fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			} else {
				fmt.Fprint(ngsi.StdWriter, *lists)
			}
		} else {
			info := ngsi.AllServersList().ServerList(gCmdMode, c.Bool("all"))
			list := info.List(c.Bool("singleLine"))
			fmt.Fprintln(ngsi.StdWriter, list)
		}
	}

	return nil
}

func serverGet(c *cli.Context) error {
	const funcName = "serverGet"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")

	if host == "" {
		return &ngsiCmdError{funcName, 2, "required host not found", nil}
	}

	if c.IsSet("json") || c.Bool("pretty") {
		lists, err := ngsi.AllServersList().ServerInfoJSON(host, gCmdMode)
		if err != nil {
			return &ngsiCmdError{funcName, 3, host + " not found", err}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, []byte(*lists), "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, *lists)
		}
	} else {
		info, err := ngsi.AllServersList().ServerInfo(host, gCmdMode)
		if err != nil {
			return &ngsiCmdError{funcName, 5, host + " not found", err}
		}
		clearText := c.Bool("clearText")
		printServerInfo(ngsi, info, clearText)
	}

	return nil
}

func serverAdd(c *cli.Context) error {
	const funcName = "serverAdd"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")
	if host == "" {
		return &ngsiCmdError{funcName, 2, "required host not found", nil}
	}

	if !ngsilib.IsNameString(host) {
		return &ngsiCmdError{funcName, 3, "name error " + host, err}
	}
	if ngsi.ExistsBrokerHost(host) {
		return &ngsiCmdError{funcName, 4, host + " already exists", err}
	}

	if !c.IsSet("serverHost") && gCmdMode == "" {
		return &ngsiCmdError{funcName, 5, "serverHost is missing", err}
	}

	if !c.IsSet("serverType") {
		return &ngsiCmdError{funcName, 6, "serverType is missing", err}
	}

	serverType := gCmdMode
	if c.IsSet("serverType") {
		serverType = strings.ToLower(c.String("serverType"))
	}
	if !ngsilib.Contains(ngsi.ServerTypeArgs(), serverType) {
		return &ngsiCmdError{funcName, 7, "serverType error: " + serverType + " (Comet, Cygnus, Iota, Keyrock, Perseo, QuantumLeap, WireCloud, Geoproxy, Regproxy, Tokenproxy)", err}
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

	err = ngsi.CreateServer(host, param)
	if err != nil {
		return &ngsiCmdError{funcName, 8, err.Error(), err}
	}
	return nil
}

func serverUpdate(c *cli.Context) error {
	const funcName = "serverUpdate"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")
	if host == "" {
		return &ngsiCmdError{funcName, 2, "required host not found", nil}
	}

	if !ngsi.ExistsBrokerHost(host) {
		return &ngsiCmdError{funcName, 3, host + " not found", err}
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

	err = ngsi.UpdateServer(host, param)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	return nil
}

func serverDelete(c *cli.Context) error {
	const funcName = "serverDelete"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")
	if host == "" {
		return &ngsiCmdError{funcName, 2, "required host not found", nil}
	}

	if !ngsi.ExistsBrokerHost(host) {
		return &ngsiCmdError{funcName, 3, host + " not found", err}
	}

	if c.IsSet("items") {
		items := c.String("items")
		for _, item := range strings.Split(items, ",") {
			if err := ngsi.DeleteItem(host, item); err != nil {
				return &ngsiCmdError{funcName, 4, "delete error - " + item, err}
			}
		}
		err = ngsi.UpdateServer(host, nil)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	} else {
		if err = ngsi.IsHostReferenced(host); err != nil {
			return &ngsiCmdError{funcName, 6, host + " is referenced", err}
		}
		if ngsi.PreviousArgs.Host == host {
			ngsi.PreviousArgs.Host = ""
			ngsi.PreviousArgs.Tenant = ""
			ngsi.PreviousArgs.Scope = ""
		}
		_ = ngsi.DeleteServer(host)
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
