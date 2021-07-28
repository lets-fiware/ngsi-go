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
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func tokenCommand(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "tokenCommand"

	if c.Bool("revoke") {
		return revokeTokenCommand(c, ngsi)
	}

	host := ngsi.Host

	tokenClient, err := ngsicli.NewClient(ngsi, c, false, nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	token, err := ngsi.TokenInfo(tokenClient)
	if err != nil {
		return ngsierr.New(funcName, 2, host+" has no token", err)
	}

	time := token.Expires.Unix() - ngsi.TimeLib.NowUnix()
	if time < 0 {
		time = 0
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		var b []byte
		if token.Type == ngsilib.CKeyrockIDM {
			b, err = getKeyrockUserInfo(tokenClient, token.Token)
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
		} else {
			b, err = ngsi.GetTokenInfo(token)
			if err != nil {
				return ngsierr.New(funcName, 4, err.Error(), err)
			}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(b))
		}
	} else if c.Bool("expires") {
		fmt.Fprintf(ngsi.StdWriter, "%d\n", time)
	} else {
		fmt.Fprintln(ngsi.StdWriter, tokenClient.Token)
	}

	return nil
}

func revokeTokenCommand(c *ngsicli.Context, ngsi *ngsilib.NGSI) error {
	const funcName = "revokeTokenCommand"

	if c.IsSetOR([]string{"verbose", "pretty", "expires"}) {
		return ngsierr.New(funcName, 1, "only --revoke can be specified", nil)
	}

	host := ngsi.Host

	client, err := ngsicli.NewClientSkipGetToken(ngsi, c, false, nil)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	_, err = ngsi.TokenInfo(client)
	if err != nil {
		return ngsierr.New(funcName, 3, host+" has no token", err)
	}

	err = ngsi.RevokeToken(client)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}

	return nil
}

func getKeyrockUserInfo(client *ngsilib.Client, token string) ([]byte, error) {
	const funcName = "getKeyrockUserInfo"

	if token == "" {
		return nil, ngsierr.New(funcName, 1, "token is empty", nil)
	}

	client.SetPath("/v1/auth/tokens")
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("X-Subject-token", token)

	res, body, err := client.HTTPGet()
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}
	return body, nil
}
