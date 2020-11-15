/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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
	"net/url"
	"strconv"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"

	"github.com/urfave/cli/v2"
)

func parseFlags(ngsi *ngsilib.NGSI, c *cli.Context) (*ngsilib.CmdFlags, error) {
	const funcName = "parseFlags"

	cmdFlags := new(ngsilib.CmdFlags)

	for _, flag := range c.FlagNames() {
		switch flag {
		case "token":
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
				if link, err := ngsi.GetContext(c.String(flag)); err == nil {
					cmdFlags.Link = &link
				} else {
					return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
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

func parseFlags2(ngsi *ngsilib.NGSI, c *cli.Context) (*ngsilib.CmdFlags, error) {
	const funcName = "parseFlags2"

	cmdFlags := new(ngsilib.CmdFlags)

	for _, flag := range c.FlagNames() {
		switch flag {
		case "token2":
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
				if link, err := ngsi.GetContext(c.String(flag)); err == nil {
					cmdFlags.Link = &link
				} else {
					return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
				}
			}
		case "safeString2":
			if c.IsSet(flag) {
				s := c.String(flag)
				cmdFlags.SafeString = &s
			}
		case "xAuthToken2":
			cmdFlags.XAuthToken = c.Bool(flag)
		}
	}
	return cmdFlags, nil
}

func parseOptions(c *cli.Context, args []string, opts []string) *url.Values {
	v := url.Values{}
	if args != nil {
		for _, key := range args {
			if key == "limit" || key == "offset" {
				value := c.Int64(key)
				if value > 0 {
					v.Set(key, strconv.FormatInt(value, 10))
				}
			} else {
				value := c.String(key)
				if value != "" {
					if key == "query" {
						v.Set("q", c.String(key))
					} else {
						v.Set(key, c.String(key))
					}
				}
			}
		}
	}

	if opts != nil {
		options := ""
		for _, key := range opts {
			if values := c.Bool(key); values {
				options += key + ","
			}
		}
		if len(options) > 0 {
			options = options[:len(options)-1]
			v.Set("options", options)
		}
	}

	return &v
}
