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

package iotagent

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type confGroup struct {
	Services []iotaService `json:"services,omitempty"`
}

type iotaService struct {
	Apikey     string `json:"apikey,omitempty"`
	Token      string `json:"token,omitempty"`
	Cbroker    string `json:"cbroker,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	Resource   string `json:"resource,omitempty"`
}

func idasServicesList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "idasServicesList"

	v := ngsicli.ParseOptions(c, []string{"limit", "offset", "resource"}, nil)
	client.SetQuery(v)

	client.SetPath("/iot/services")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func idasServicesCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "idasServicesCreate"

	var b []byte
	var err error

	if c.IsSet("data") && !c.IsSetOR([]string{"apikey", "token", "type", "resource", "cbroker"}) {
		b, err = ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
	} else if !c.IsSet("data") && c.IsSetAND([]string{"apikey", "type", "resource"}) {
		service := iotaService{}
		service.Apikey = c.String("apikey")
		service.EntityType = c.String("type")
		service.Resource = c.String("resource")
		if c.IsSet("token") {
			service.Token = c.String("token")
		}
		if c.IsSet("cbroker") {
			service.Cbroker, err = getCbroker(ngsi, c.String("cbroker"))
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error(), err)
			}
		}
		config := confGroup{}
		config.Services = []iotaService{service}

		b, err = ngsilib.JSONMarshal(&config)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	} else {
		return ngsierr.New(funcName, 4, "apikey, type and resource are needed", nil)
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/iot/services")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 5, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated {
		return ngsierr.New(funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func idasServicesUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "idasServicesUpdate"

	v := ngsicli.ParseOptions(c, []string{"resource", "apikey"}, nil)
	client.SetQuery(v)

	var b []byte
	var err error

	flag := c.IsSetOR([]string{"token", "type", "cbroker"})
	if c.IsSet("data") && !flag {
		b, err = ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
	} else if !c.IsSet("data") && flag {
		service := iotaService{}
		if c.IsSet("token") {
			service.Token = c.String("token")
		}
		if c.IsSet("type") {
			service.EntityType = c.String("type")
		}
		if c.IsSet("cbroker") {
			service.Cbroker, err = getCbroker(ngsi, c.String("cbroker"))
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error(), err)
			}
		}

		b, err = ngsilib.JSONMarshal(&service)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	} else {
		return ngsierr.New(funcName, 4, "configuration group field not found", nil)
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/iot/services")

	res, body, err := client.HTTPPut(b)
	if err != nil {
		return ngsierr.New(funcName, 5, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}

func idasServicesDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "idasServicesDelete"

	v := ngsicli.ParseOptions(c, []string{"resource", "apikey", "device"}, nil)
	if !c.IsSet("apikey") {
		v.Set("apikey", "")
	}
	client.SetQuery(v)

	client.SetPath("/iot/services")

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil

}
