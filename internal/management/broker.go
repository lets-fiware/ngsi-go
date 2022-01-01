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

package management

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func brokersList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "brokersList"

	host := c.String("host")

	if host != "" {
		info, err := ngsi.AllServersList().BrokerInfo(host)
		if err != nil {
			return ngsierr.New(funcName, 1, host+" not found", err)
		}
		clearText := c.Bool("clearText")
		printBrokerInfo(ngsi, info, clearText)
		return nil
	}

	if c.IsSet("json") || c.Bool("pretty") {
		lists, err := ngsi.AllServersList().BrokerInfoJSON("")
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
		info := ngsi.AllServersList().BrokerList()
		list := info.List(c.Bool("singleLine"))
		fmt.Fprintln(ngsi.StdWriter, list)
	}

	setPreviousArgs(ngsi, host)

	return nil
}

func brokersGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "brokersGet"

	host := c.String("host")

	if c.IsSet("json") || c.Bool("pretty") {
		lists, err := ngsi.AllServersList().BrokerInfoJSON(host)
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
		info, err := ngsi.AllServersList().BrokerInfo(host)
		if err != nil {
			return ngsierr.New(funcName, 3, host+" not found", err)
		}
		clearText := c.Bool("clearText")
		printBrokerInfo(ngsi, info, clearText)
	}

	setPreviousArgs(ngsi, host)

	return nil
}

func brokersAdd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "brokersAdd"

	host := c.String("host")

	if !ngsilib.IsNameString(host) {
		return ngsierr.New(funcName, 1, "name error "+host, nil)
	}
	if ngsi.ExistsBrokerHost(host) {
		if c.Bool("overWrite") {
			if err := deleteBrokerAlias(ngsi, host); err != nil {
				return ngsierr.New(funcName, 2, err.Error(), err)
			}
		} else {
			return ngsierr.New(funcName, 3, host+" already exists", nil)
		}
	}

	if !c.IsSet("brokerHost") {
		return ngsierr.New(funcName, 4, "brokerHost is missing", nil)
	}

	if !c.IsSet("ngsiType") && ngsilib.IsHTTP(c.String("brokerHost")) {
		return ngsierr.New(funcName, 5, "ngsiType is missing", nil)
	}

	if c.IsSet("ngsiType") && !ngsilib.IsHTTP(c.String("brokerHost")) {
		return ngsierr.New(funcName, 6, "can't specify ngsiType", nil)
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

	setPreviousArgs(ngsi, host)

	err := ngsi.CreateServer(host, param)
	if err != nil {
		ngsi.Updated = false
		return ngsierr.New(funcName, 7, err.Error(), err)
	}

	return nil
}

func brokersUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "brokersUpdate"

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
		return ngsierr.New(funcName, 2, err.Error(), nil)
	}
	return nil
}

func brokersDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "brokersDelete"

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
			return ngsierr.New(funcName, 3, err.Error(), nil)
		}
	} else {
		err := deleteBrokerAlias(ngsi, host)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
	}
	return nil
}

func deleteBrokerAlias(ngsi *ngsilib.NGSI, host string) error {
	const funcName = "deleteBrokerAlias"

	err := ngsi.IsHostReferenced(host)
	if err != nil {
		return ngsierr.New(funcName, 1, host+" is referenced", err)
	}

	setPreviousArgs(ngsi, "")

	err = ngsi.DeleteServer(host)
	if err != nil {
		ngsi.Updated = false
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	return nil
}

func printBrokerInfo(ngsi *ngsilib.NGSI, info *ngsilib.Server, clearText bool) {

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

func obfuscateText(text string, clearText bool) string {
	if !clearText {
		s := ""
		for range text {
			s += "*"
		}
		return s
	}
	return text
}

func setPreviousArgs(ngsi *ngsilib.NGSI, host string) {
	if ngsi.GetPreviousArgs().UsePreviousArgs {
		ngsi.PreviousArgs.Host = host
		ngsi.PreviousArgs.Tenant = ""
		ngsi.PreviousArgs.Scope = ""
		ngsi.Updated = true
	}
}
