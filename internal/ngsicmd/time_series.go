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
	"github.com/urfave/cli/v2"
)

func tsAttrRead(c *cli.Context) error {
	const funcName = "tsAttrRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"comet", "quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.ServerType == "comet" {
		return cometAttrReadMain(c, ngsi, client)
	}
	return qlAttrReadMain(c, ngsi, client)
}

func tsAttrsRead(c *cli.Context) error {
	const funcName = "tsAttrsRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	return qlAttrsReadMain(c, ngsi, client)
}

func tsEntitiesRead(c *cli.Context) error {
	const funcName = "tsEntitiesRead"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return qlEntitiesReadMain(c, ngsi, client)
}

func tsEntitiesDelete(c *cli.Context) error {
	const funcName = "tsEntitiesDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"comet", "quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.ServerType == "comet" {
		return cometEntitiesDeleteMain(c, ngsi, client)
	}
	return qlEntitiesDeleteMain(c, ngsi, client)
}

func tsEntityDelete(c *cli.Context) error {
	const funcName = "tsEntityDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"comet", "quantumleap"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if client.Server.ServerType == "comet" {
		return cometEntityDeleteMain(c, ngsi, client)
	}
	return qlEntityDeleteMain(c, ngsi, client)
}
