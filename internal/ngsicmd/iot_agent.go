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

func idasServicesList(c *cli.Context) error {
	const funcName = "idasServicesList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	v := parseOptions(c, []string{"limit", "offset", "resource"}, nil)
	client.SetQuery(v)

	client.SetPath("/iot/services")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func idasServicesCreate(c *cli.Context) error {
	const funcName = "idasServicesCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	var b []byte

	if c.IsSet("data") && !isSetOR(c, []string{"apikey", "token", "type", "resource", "cbroker"}) {
		b, err = readAll(c, ngsi)
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}
	} else if !c.IsSet("data") && isSetAND(c, []string{"apikey", "type", "resource"}) {
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
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
		}
		config := confGroup{}
		config.Services = []iotaService{service}

		b, err = ngsilib.JSONMarshal(&config)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
	} else {
		return &ngsiCmdError{funcName, 6, "apikey, type and resource are needed", nil}
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/iot/services")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func idasServicesUpdate(c *cli.Context) error {
	const funcName = "idasServicesUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("resource") {
		return &ngsiCmdError{funcName, 3, "resource not fuond", nil}
	}

	v := parseOptions(c, []string{"resource", "apikey"}, nil)
	client.SetQuery(v)

	var b []byte

	flag := isSetOR(c, []string{"token", "type", "cbroker"})
	if c.IsSet("data") && !flag {
		b, err = readAll(c, ngsi)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
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
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
		}

		b, err = ngsilib.JSONMarshal(&service)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	} else {
		return &ngsiCmdError{funcName, 7, "configuration group field not found", nil}
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/iot/services")

	res, body, err := client.HTTPPut(b)
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func idasServicesDelete(c *cli.Context) error {
	const funcName = "idasServicesDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("resource") {
		return &ngsiCmdError{funcName, 3, "resource not fuond", nil}
	}

	v := parseOptions(c, []string{"resource", "apikey", "device"}, nil)
	if !c.IsSet("apikey") {
		v.Set("apikey", "")
	}
	client.SetQuery(v)

	client.SetPath("/iot/services")

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil

}

func idasDevicesList(c *cli.Context) error {
	const funcName = "idasDevicesList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if c.IsSet("detailed") {
		value := c.String("detailed")
		if value != "on" && value != "off" {
			return &ngsiCmdError{funcName, 3, "specify either on or off to --detailed", err}
		}
	}

	v := parseOptions(c, []string{"limit", "offset", "detailed", "entity", "protocol"}, nil)
	client.SetQuery(v)

	client.SetPath("/iot/devices")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func idasDevicesGet(c *cli.Context) error {
	const funcName = "idasDevicesGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("id") {
		return &ngsiCmdError{funcName, 3, "device id not found", err}
	}

	path := "/iot/devices/" + c.String("id")
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func idasDevicesCreate(c *cli.Context) error {
	const funcName = "idasDevicesCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("data") {
		return &ngsiCmdError{funcName, 3, "--data not found", nil}
	}

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/iot/devices")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}
	return nil
}

func idasDevicesUpdate(c *cli.Context) error {
	const funcName = "idasDevicesUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("data") {
		return &ngsiCmdError{funcName, 3, "--data not found", nil}
	}

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	client.SetHeader("Content-Type", "application/json")

	if !c.IsSet("id") {
		return &ngsiCmdError{funcName, 5, "device id not found", err}
	}

	path := "/iot/devices/" + c.String("id")
	client.SetPath(path)

	res, body, err := client.HTTPPut(b)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func idasDevicesDelete(c *cli.Context) error {
	const funcName = "idasDevicesDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	client, err := newClient(ngsi, c, false, []string{"iota"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("id") {
		return &ngsiCmdError{funcName, 3, "device id not found", err}
	}

	path := "/iot/devices/" + c.String("id")
	client.SetPath(path)

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	return nil
}

func getCbroker(ngsi *ngsilib.NGSI, host string) (string, error) {
	const funcName = "getCbroker"

	if ngsilib.IsHTTP(host) {
		return host, nil
	}

	info, err := ngsi.AllServersList().BrokerInfo(host)
	if err == nil {
		if ngsilib.IsHTTP(info.ServerHost) {
			return info.ServerHost, nil
		}
		info, err = ngsi.AllServersList().BrokerInfo(info.ServerHost)
		if err == nil {
			return info.ServerHost, nil
		}
	}

	return "", &ngsiCmdError{funcName, 1, "specify url or broker alias to --cbroker", nil}
}
