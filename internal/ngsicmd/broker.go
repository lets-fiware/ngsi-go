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

func brokersList(c *cli.Context) error {
	const funcName = "brokersList"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")

	if host != "" {
		info, err := ngsi.AllServersList().BrokerInfo(host)
		if err != nil {
			return &ngsiCmdError{funcName, 2, host + " not found", err}
		}
		printBrokerInfo(ngsi, info)
		return nil
	}

	if c.IsSet("json") || c.Bool("pretty") {
		lists, err := ngsi.AllServersList().BrokerInfoJSON("")
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
		info := ngsi.AllServersList().BrokerList()
		list := info.List()
		fmt.Fprintln(ngsi.StdWriter, list)
	}

	return nil
}

func brokersGet(c *cli.Context) error {
	const funcName = "brokersGet"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")

	if host == "" {
		return &ngsiCmdError{funcName, 2, "required host not found", nil}
	}

	if c.IsSet("json") || c.Bool("pretty") {
		lists, err := ngsi.AllServersList().BrokerInfoJSON(host)
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
		info, err := ngsi.AllServersList().BrokerInfo(host)
		if err != nil {
			return &ngsiCmdError{funcName, 5, host + " not found", err}
		}
		printBrokerInfo(ngsi, info)
	}

	return nil
}

func brokersAdd(c *cli.Context) error {
	const funcName = "brokersAdd"

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

	if !c.IsSet("brokerHost") {
		return &ngsiCmdError{funcName, 5, "brokerHost is missing", err}
	}

	if !c.IsSet("ngsiType") && ngsilib.IsHTTP(c.String("brokerHost")) {
		return &ngsiCmdError{funcName, 6, "ngsiType is missing", err}
	}

	if c.IsSet("ngsiType") && !ngsilib.IsHTTP(c.String("brokerHost")) {
		return &ngsiCmdError{funcName, 7, "can't specify ngsiType", err}
	}

	param := map[string]string{"serverType": "broker"}
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

	err = ngsi.CreateServer(host, param)
	if err != nil {
		return &ngsiCmdError{funcName, 8, err.Error(), err}
	}
	return nil
}

func brokersUpdate(c *cli.Context) error {
	const funcName = "brokersUpdate"

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

func brokersDelete(c *cli.Context) error {
	const funcName = "brokersDelete"

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

func printBrokerInfo(ngsi *ngsilib.NGSI, info *ngsilib.Server) {

	fmt.Fprintln(ngsi.StdWriter, "brokerHost "+info.ServerHost)
	if info.NgsiType != "" {
		fmt.Fprintln(ngsi.StdWriter, "ngsiType "+info.NgsiType)
	}
	if strings.ToLower(info.NgsiType) == "v2" {
		if info.Tenant != "" {
			fmt.Fprintln(ngsi.StdWriter, "FIWARE-Service "+info.Tenant)
		}
		if info.Scope != "" {
			fmt.Fprintln(ngsi.StdWriter, "FIWARE-ServicePath "+info.Scope)
		}
	} else {
		if info.BrokerType != "" {
			fmt.Fprintln(ngsi.StdWriter, "brokerType "+info.BrokerType)
		}
		if info.Tenant != "" {
			fmt.Fprintln(ngsi.StdWriter, "Tenant "+info.Tenant)
		}
		if info.Scope != "" {
			fmt.Fprintln(ngsi.StdWriter, "Scope "+info.Scope)
		}
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
		{"Password", info.Password},
		{"ClientID", info.ClientID},
		{"ClientSecret", info.ClientSecret},
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
