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

/*
type wireCloudWorkspace struct {
	Description     string `json:"description"`
	ID              string `json:"id"`
	Lastmodified    int64  `json:"lastmodified"`
	Longdescription string `json:"longdescription"`
	Name            string `json:"name"`
	Owner           string `json:"owner"`
	Public          bool   `json:"public"`
	Removable       bool   `json:"removable"`
	Requireauth     bool   `json:"requireauth"`
	Shared          bool   `json:"shared"`
	Title           string `json:"title"`
}
*/

func wireCloudTabsList(c *cli.Context) error {
	const funcName = "tabsList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"wirecloud"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	var wid string
	if c.IsSet("wid") && c.Args().Len() == 0 {
		wid = c.String("wid")
	} else if c.Args().Len() == 1 {
		wid = c.Args().Get(0)
	} else {
		return &ngsiCmdError{funcName, 3, "workspace id required", nil}
	}

	client.SetPath("/api/workspace/" + wid)
	client.SetAcceptJSON()

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return &ngsiCmdError{funcName, 5, "workspace not found", nil}
		}
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	ws := wireCloudWorkspace{}
	err = ngsilib.JSONUnmarshal(body, &ws)
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}

	if isSetOR(c, []string{"json", "pretty"}) {
		b, err := ngsilib.JSONMarshal(ws.Tabs)
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}

		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 9, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			return nil
		} else {
			fmt.Fprint(ngsi.StdWriter, string(b))
		}
	} else {
		for _, tab := range ws.Tabs {
			fmt.Fprintf(ngsi.StdWriter, "%s %s %s\n", tab.ID, tab.Name, tab.Title)
		}
	}

	return nil
}

func wireCloudTabGet(c *cli.Context) error {
	const funcName = "tabGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"wirecloud"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	var wid, tid string
	if c.IsSet("wid") && c.IsSet("tid") && c.Args().Len() == 0 {
		wid = c.String("wid")
		tid = c.String("tid")
	} else if c.Args().Len() == 2 {
		wid = c.Args().Get(0)
		tid = c.Args().Get(1)
	} else {
		for _, opt := range []option{{Name: "wid", Description: "workspace id"}, {Name: "tid", Description: "tab id"}} {
			if !c.IsSet(opt.Name) {
				return &ngsiCmdError{funcName, 3, opt.Description + " required", nil}
			}
		}
	}

	client.SetPath("/api/workspace/" + wid + "/tab/" + tid)
	client.SetAcceptJSON()

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return &ngsiCmdError{funcName, 5, "workspace or tab not found", nil}
		}
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
