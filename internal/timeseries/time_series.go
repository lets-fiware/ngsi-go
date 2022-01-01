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

package timeseries

import (
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func tsAttrRead(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "tsAttrRead"
	var err error

	if client.Server.ServerType == "comet" {
		err = cometAttrReadMain(c, ngsi, client)
	} else {
		err = qlAttrReadMain(c, ngsi, client)
	}

	if err != nil {
		err = ngsierr.New(funcName, 1, err.Error(), err)
	}

	return err
}

func tsEntitiesDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "tsEntitiesDelete"
	var err error

	if client.Server.ServerType == "comet" {
		err = cometEntitiesDeleteMain(c, ngsi, client)
	} else {
		err = qlEntitiesDeleteMain(c, ngsi, client)
	}

	if err != nil {
		err = ngsierr.New(funcName, 1, err.Error(), err)
	}

	return err
}

func tsEntityDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "tsEntityDelete"
	var err error

	if client.Server.ServerType == "comet" {
		err = cometEntityDeleteMain(c, ngsi, client)
	} else {
		err = qlEntityDeleteMain(c, ngsi, client)
	}

	if err != nil {
		err = ngsierr.New(funcName, 1, err.Error(), err)
	}

	return err
}
