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

package ngsicli

import (
	"net/url"
	"strconv"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func parseFlags(ngsi *ngsilib.NGSI, c *Context) (*ngsilib.CmdFlags, error) {
	const funcName = "parseFlags"

	cmdFlags := new(ngsilib.CmdFlags)

	for _, flag := range c.FlagNames() {
		switch flag {
		case "oAuthToken":
			if c.IsSet(flag) {
				token := c.String(flag)
				cmdFlags.Token = &token
			}
		case "service":
			if c.IsSet(flag) {
				tenant := c.String(flag)
				cmdFlags.Tenant = &tenant
			}
		case "path":
			if c.IsSet(flag) {
				scope := c.String(flag)
				cmdFlags.Scope = &scope
			}
		case "link":
			if c.IsSet(flag) {
				if link, err := ngsi.GetContextHTTP(c.String(flag)); err == nil {
					cmdFlags.Link = &link
				} else {
					return nil, ngsierr.New(funcName, 1, err.Error(), err)
				}
			}
		case "safeString":
			if c.IsSet(flag) {
				s := c.String(flag)
				cmdFlags.SafeString = &s
			}
		case "xAuthToken":
			cmdFlags.XAuthToken = c.Bool(flag)
		}
	}
	return cmdFlags, nil
}

func parseFlags2(ngsi *ngsilib.NGSI, c *Context) (*ngsilib.CmdFlags, error) {
	const funcName = "parseFlags2"

	cmdFlags := new(ngsilib.CmdFlags)

	for _, flag := range c.FlagNames() {
		switch flag {
		case "oAuthToken2":
			if c.IsSet(flag) {
				token := c.String(flag)
				cmdFlags.Token = &token
			}
		case "service2":
			if c.IsSet(flag) {
				tenant := c.String(flag)
				cmdFlags.Tenant = &tenant
			}
		case "path2":
			if c.IsSet(flag) {
				scope := c.String(flag)
				cmdFlags.Scope = &scope
			}
		case "link2":
			if c.IsSet(flag) {
				if link, err := ngsi.GetContextHTTP(c.String(flag)); err == nil {
					cmdFlags.Link = &link
				} else {
					return nil, ngsierr.New(funcName, 1, err.Error(), err)
				}
			}
		}
	}
	return cmdFlags, nil
}

func ParseOptions(c *Context, args []string, opts []string) *url.Values {
	v := url.Values{}
	for _, key := range args {
		if c.IsSet(key) {
			if key == "limit" || key == "offset" || key == "hLimit" || key == "hOffset" || key == "lastN" {
				value := c.Int64(key)
				v.Set(key, strconv.FormatInt(value, 10))
			} else {
				switch key {
				default:
					v.Set(key, c.String(key))
				case "query":
					v.Set("q", c.String(key))
				case "device":
					if b := c.Bool(("device")); b {
						v.Set(key, "true")
					}
				case "details":
					v.Set(key, c.String(key))
				}
			}
		}
	}

	if opts != nil {
		options := ""
		for _, key := range opts {
			if c.IsSet(key) {
				if values := c.Bool(key); values {
					options += key + ","
				}
			}
		}
		if len(options) > 0 {
			options = options[:len(options)-1]
			v.Set("options", options)
		}
	}

	return &v
}
