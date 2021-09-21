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

package ngsicli

import (
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func NewClient(ngsi *ngsilib.NGSI, c *Context, isHTTPVerb bool, serverList []string) (*ngsilib.Client, error) {
	return newClient(ngsi, c, isHTTPVerb, serverList, false, false)
}

func NewClientSkipGetToken(ngsi *ngsilib.NGSI, c *Context, isHTTPVerb bool, serverList []string) (*ngsilib.Client, error) {
	return newClient(ngsi, c, isHTTPVerb, serverList, true, false)
}

// newClient is a wrapper function for ngsi.NewClient function

func newClient(ngsi *ngsilib.NGSI, c *Context, isHTTPVerb bool, serverList []string, skipGetToken, dest bool) (*ngsilib.Client, error) {
	const funcName = "newClient"

	var flags *ngsilib.CmdFlags
	var err error
	var host string

	if !dest {
		flags, err = parseFlags(ngsi, c)
		host = ngsi.Host
	} else {
		flags, err = parseFlags2(ngsi, c)
		host = ngsi.Destination
	}
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}

	client, err := ngsi.NewClient(host, flags, isHTTPVerb, skipGetToken)
	if err != nil {
		msg := err.Error()
		if dest {
			msg += " (destination)"
		}
		return nil, ngsierr.New(funcName, 2, msg, err)
	}

	if serverList != nil {
		serverType := client.Server.ServerType
		if serverType == "broker" {
			serverType += client.Server.NgsiType
		}
		if !ngsilib.Contains(serverList, serverType) {
			if strings.HasPrefix(serverType, "broker") {
				if serverType == "brokerld" && ngsilib.Contains(serverList, "brokerv2") {
					return nil, ngsierr.New(funcName, 3, "only available on NGSIv2", nil)
				}
				if serverType == "brokerv2" && ngsilib.Contains(serverList, "brokerld") {
					return nil, ngsierr.New(funcName, 4, "only available on NGSI-LD", nil)
				}
				serverType = client.Server.ServerType
			}
			return nil, ngsierr.New(funcName, 5, "not supported by "+serverType, nil)
		}
	}

	return client, nil
}
