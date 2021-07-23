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
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func healthCheck(c *cli.Context) error {
	const funcName = "healthCheck"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap", "broker", "regproxy", "tokenproxy", "geoproxy"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	path := "/health"
	if client.Server.ServerType == "broker" {
		if client.Server.NgsiType == "ld" && client.Server.BrokerType == "scorpio" {
			path = "/scorpio/v1/info/health"
		} else {
			return &ngsiCmdError{funcName, 3, "brokerType error", err}
		}

	}
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}
