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
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func tokenCommand(c *cli.Context) error {
	const funcName = "tokenCommand"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if c.Bool("revoke") {
		return revokeTokenCommand(c, ngsi)
	}

	host := ngsi.Host

	client, err := newClient(ngsi, c, false, nil)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	token, err := ngsi.TokenInfo(client)
	if err != nil {
		return &ngsiCmdError{funcName, 3, host + " has no token", err}
	}

	time := token.Expires.Unix() - ngsi.TimeLib.NowUnix()
	if time < 0 {
		time = 0
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		var b []byte
		switch token.Type {
		default: // ngsilib.CKeyrock, ngsilib.CTokenproxy, ngsilib.CKeyrocktokenprovider
			token.Oauth.ExpiresIn = time
			b, err = ngsilib.JSONMarshal(token.Oauth)
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
		case ngsilib.CBasic:
			fmt.Fprintln(ngsi.StdWriter, "no information available")
			return nil
		case ngsilib.CThinkingCities:
			b, err = ngsilib.JSONMarshal(token.Keystone)
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
		case ngsilib.CKeycloak:
			b, err = ngsilib.JSONMarshal(token.Keycloak)
			if err != nil {
				return &ngsiCmdError{funcName, 6, err.Error(), err}
			}
		case ngsilib.CWSO2:
			b, err = ngsilib.JSONMarshal(token.WSO2)
			if err != nil {
				return &ngsiCmdError{funcName, 7, err.Error(), err}
			}
		case ngsilib.CKeyrockIDM:
			b, err = getKeyrockUserInfo(client, token.Token)
			if err != nil {
				return &ngsiCmdError{funcName, 8, err.Error(), err}
			}
		case ngsilib.CKong:
			b, err = ngsilib.JSONMarshal(token.Kong)
			if err != nil {
				return &ngsiCmdError{funcName, 9, err.Error(), err}
			}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 10, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(b))
		}
	} else if c.Bool("expires") {
		fmt.Fprintf(ngsi.StdWriter, "%d\n", time)
	} else {
		fmt.Fprintln(ngsi.StdWriter, client.Token)
	}

	return nil
}

func revokeTokenCommand(c *cli.Context, ngsi *ngsilib.NGSI) error {
	const funcName = "revokeTokenCommand"

	if isSetOR(c, []string{"verbose", "pretty", "expires"}) {
		return &ngsiCmdError{funcName, 1, "only --revoke can be specified", nil}
	}

	host := ngsi.Host

	client, err := newClientSkipGetToken(ngsi, c, false, nil)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	_, err = ngsi.TokenInfo(client)
	if err != nil {
		return &ngsiCmdError{funcName, 3, host + " has no token", err}
	}

	err = ngsi.RevokeToken(client)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	return nil
}

func getKeyrockUserInfo(client *ngsilib.Client, token string) ([]byte, error) {
	const funcName = "getKeyrockUserInfo"

	if token == "" {
		return nil, &ngsiCmdError{funcName, 1, "token is empty", nil}
	}

	client.SetPath("/v1/auth/tokens")
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("X-Subject-token", token)

	res, body, err := client.HTTPGet()
	if err != nil {
		return nil, &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return nil, &ngsiCmdError{funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}
	return body, nil
}
