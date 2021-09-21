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

package wirecloud

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type wireCloudWorkspaceInfos []wireCloudWorkspaceInfo

type wireCloudWorkspaceInfo struct {
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

type wireCloudIPreference struct {
	Name     string      `json:"name,omitempty"`
	Secure   bool        `json:"secure,omitempty"`
	Readonly bool        `json:"readonly,omitempty"`
	Hidden   bool        `json:"hidden,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

type wireCloudWiringOperator struct {
	ID          string                          `json:"id"`
	Name        string                          `json:"name"`
	Preferences map[string]wireCloudIPreference `json:"preferences"`
	Properties  interface{}                     `json:"properties"`
}

type wireCloudWiringVisualdescription struct {
	Behaviours  []interface{}             `json:"behaviours"`
	Components  wireCloudWiringComponents `json:"components"`
	Connections []wireCloudConnection     `json:"connections"`
}

type wireCloudWiring struct {
	Version           string                             `json:"version"`
	Connections       []wireCloudWiringConnection        `json:"connections"`
	Operators         map[string]wireCloudWiringOperator `json:"operators"`
	Visualdescription wireCloudWiringVisualdescription   `json:"visualdescription"`
}

type wireCloudWiringComponents struct {
	Operator map[string]wireCloudWidget `json:"operator"`
	Widget   map[string]wireCloudWidget `json:"widget"`
}

type wireCloudWiringConnection struct {
	Readonly bool                              `json:"readonly"`
	Source   wireCloudWiringConnectionEndpoint `json:"source"`
	Target   wireCloudWiringConnectionEndpoint `json:"target"`
}

type wireCloudWiringConnectionEndpoint struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	Endpoint string `json:"endpoint"`
}

type wireCloudInherit struct {
	Inherit bool   `json:"inherit"`
	Value   string `json:"value"`
}

type wireCloudUser struct {
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	Organization bool   `json:"organization"`
	Accesslevel  string `json:"accesslevel"`
}

type wireCloudPreferences struct {
	Public        wireCloudInherit `json:"public,omitempty"`
	Requireauth   wireCloudInherit `json:"requireauth,omitempty"`
	Sharelist     wireCloudInherit `json:"sharelist,omitempty"`
	Initiallayout wireCloudInherit `json:"initiallayout,omitempty"`
	Baselayout    wireCloudInherit `json:"baselayout,omitempty"`
}

type wireCloudIwidget struct {
	ID            string                          `json:"id"`
	Title         string                          `json:"title"`
	Tab           int                             `json:"tab"`
	Layout        int                             `json:"layout"`
	Widget        string                          `json:"widget"`
	Top           int                             `json:"top"`
	Left          int                             `json:"left"`
	Anchor        string                          `json:"anchor"`
	Relx          bool                            `json:"relx"`
	Rely          bool                            `json:"rely"`
	Relheight     bool                            `json:"relheight"`
	Relwidth      bool                            `json:"relwidth"`
	ZIndex        int                             `json:"zIndex"`
	Width         int                             `json:"width"`
	Height        int                             `json:"height"`
	Fulldragboard bool                            `json:"fulldragboard"`
	Minimized     bool                            `json:"minimized"`
	Readonly      bool                            `json:"readonly"`
	Permissions   interface{}                     `json:"permissions"`
	Preferences   map[string]wireCloudIPreference `json:"preferences,omitempty"`
	Properties    interface{}                     `json:"properties"`
	Titlevisible  bool                            `json:"titlevisible"`
}

type wireCloudTabs []wireCloudTab

type wireCloudTab struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Title       string               `json:"title"`
	Visible     bool                 `json:"visible"`
	Preferences wireCloudPreferences `json:"preferences"`
	Iwidgets    []wireCloudIwidget   `json:"iwidgets"`
}

type wireCloudConnection struct {
	Sourcename   string `json:"sourcename"`
	Sourcehandle string `json:"sourcehandle"`
	Targetname   string `json:"targetname"`
	Targethandle string `json:"targethandle"`
}

type wireCloudWidgetPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type wireCloudEndpoints struct {
	Source []interface{} `json:"source"`
	Target []interface{} `json:"target"`
}

type wireCloudWidget struct {
	Name      string                  `json:"name"`
	Position  wireCloudWidgetPosition `json:"position"`
	Collapsed bool                    `json:"collapsed"`
	Endpoints wireCloudEndpoints      `json:"endpoints"`
}

type wireCloudWorkspace struct {
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	Title           string               `json:"title"`
	Public          bool                 `json:"public"`
	Shared          bool                 `json:"shared"`
	Requireauth     bool                 `json:"requireauth"`
	Owner           string               `json:"owner"`
	Removable       bool                 `json:"removable"`
	Lastmodified    int64                `json:"lastmodified"`
	Description     string               `json:"description"`
	Longdescription string               `json:"longdescription"`
	Preferences     wireCloudPreferences `json:"preferences"`
	Users           []wireCloudUser      `json:"users"`
	Groups          []interface{}        `json:"groups"`
	EmptyParams     []interface{}        `json:"empty_params"`
	ExtraPrefs      []interface{}        `json:"extra_prefs"`
	Tabs            wireCloudTabs        `json:"tabs"`
	Wiring          wireCloudWiring      `json:"wiring"`
}

func wireCloudWorkspacesList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "workspacesList"

	client.SetPath("/api/workspaces")
	client.SetAcceptJSON()

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if c.IsSetOR([]string{"json", "pretty"}) {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			return nil
		}
		fmt.Fprint(ngsi.StdWriter, string(body))
	} else {
		wss := wireCloudWorkspaceInfos{}
		err = ngsilib.JSONUnmarshal(body, &wss)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}

		sort.Sort(wireCloudWorkspaceInfos(wss))

		for _, ws := range wss {
			fmt.Fprintf(ngsi.StdWriter, "%s %s %s %s\n", ws.ID, ws.Name, ws.Title, ngsilib.GetTime(ngsi, ws.Lastmodified))
		}
	}

	return nil
}

func wireCloudWorkspaceGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "workspaceGet"

	client.SetPath("/api/workspace/" + c.String("wid"))
	client.SetAcceptJSON()

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return ngsierr.New(funcName, 2, "workspace not found", nil)
		}
		return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if !c.IsSetOR([]string{"json", "pretty"}) {
		var ws wireCloudWorkspace
		err := ngsilib.JSONUnmarshal(body, &ws)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
		if c.Bool("tabs") {
			sort.Sort(wireCloudTabs(ws.Tabs))
			for _, v := range ws.Tabs {
				fmt.Fprintf(ngsi.StdWriter, "%s %s %s\n", v.ID, v.Name, v.Title)
			}
		} else if c.Bool("widgets") {
			for _, v := range ws.Wiring.Visualdescription.Components.Widget {
				fmt.Fprintf(ngsi.StdWriter, "%s\n", v.Name)
			}
		} else if c.Bool("operators") {
			for _, v := range ws.Wiring.Visualdescription.Components.Operator {
				fmt.Fprintf(ngsi.StdWriter, "%s\n", v.Name)
			}
		} else if c.Bool("users") {
			for _, v := range ws.Users {
				fmt.Fprintf(ngsi.StdWriter, "%s %s\n", v.Username, v.Accesslevel)
			}
		} else {
			fmt.Fprintf(ngsi.StdWriter, "%s %s %s %s\n", ws.ID, ws.Name, ws.Title, ngsilib.GetTime(ngsi, ws.Lastmodified))
		}
	} else {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			return nil
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	}

	return nil
}

func (s wireCloudWorkspaceInfos) Len() int {
	return len(s)
}

func (s wireCloudWorkspaceInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s wireCloudWorkspaceInfos) Less(i, j int) bool {
	ii := "000000" + s[i].ID
	jj := "000000" + s[j].ID
	return ii[len(ii)-6:] < jj[len(jj)-6:]
}

func (s wireCloudTabs) Len() int {
	return len(s)
}

func (s wireCloudTabs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s wireCloudTabs) Less(i, j int) bool {
	ii := "000000" + s[i].ID
	jj := "000000" + s[j].ID
	return ii[len(ii)-6:] < jj[len(jj)-6:]
}
