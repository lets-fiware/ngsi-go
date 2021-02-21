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

type keyrockIotAgentItems struct {
	ID            string `json:"id,omitempty"`
	OauthClientID string `json:"oauth_client_id,omitempty"`
}

type keyrockIotAgents struct {
	IotAgents []keyrockIotAgentItems `json:"iot_agents,omitempty"`
}

func iotAgentsList(c *cli.Context) error {
	const funcName = "iotAgentsList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/iot_agents")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			fmt.Fprintln(ngsi.StdWriter, "iot agents not found")
			return nil
		}
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 6, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var idas keyrockIotAgents
		err := ngsilib.JSONUnmarshal(body, &idas)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		for _, iota := range idas.IotAgents {
			fmt.Fprintln(ngsi.StdWriter, iota.ID)
		}
	}

	return nil
}

func iotAgentsGet(c *cli.Context) error {
	const funcName = "iotAgentsGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	if !c.IsSet("iid") {
		return &ngsiCmdError{funcName, 4, "specify iot agent id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/iot_agents/" + c.String("iid"))

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func iotAgentsCreate(c *cli.Context) error {
	const funcName = "iotAgentsCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/iot_agents")

	client.SetHeader("Content-Type", "application/json")

	res, body, err := client.HTTPPost([]byte(""))
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
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

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func iotAgentsReset(c *cli.Context) error {
	const funcName = "iotAgentsReset"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	if !c.IsSet("iid") {
		return &ngsiCmdError{funcName, 4, "specify iot agent id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/iot_agents/" + c.String("iid"))

	client.SetContentJSON()

	res, body, err := client.HTTPPatch([]byte(""))
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func iotAgentsDelete(c *cli.Context) error {
	const funcName = "iotAgentsDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("aid") {
		return &ngsiCmdError{funcName, 3, "specify application id", nil}
	}
	if !c.IsSet("iid") {
		return &ngsiCmdError{funcName, 4, "specify iot agent id", nil}
	}
	client.SetPath("/v1/applications/" + c.String("aid") + "/iot_agents/" + c.String("iid"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}
