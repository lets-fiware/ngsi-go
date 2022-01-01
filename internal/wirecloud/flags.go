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

package wirecloud

import "github.com/lets-fiware/ngsi-go/internal/ngsicli"

// WireCloud
var (
	wireCloudWorkspaceIdRFlag = &ngsicli.StringFlag{
		Name:     "wid",
		Aliases:  []string{"w"},
		Usage:    "workspace id",
		Required: true,
	}
	wireCloudTabIdRFlag = &ngsicli.StringFlag{
		Name:     "tid",
		Aliases:  []string{"t"},
		Usage:    "tab id",
		Required: true,
	}
	wireCloudWidgetFlag = &ngsicli.BoolFlag{
		Name:  "widget",
		Usage: "filtering widget",
	}
	wireCloudOperatorFlag = &ngsicli.BoolFlag{
		Name:  "operator",
		Usage: "filtering operator",
	}
	wireCloudMashupFlag = &ngsicli.BoolFlag{
		Name:  "mashup",
		Usage: "filtering mashup",
	}
	wireCloudVenderFlag = &ngsicli.StringFlag{
		Name:    "vender",
		Aliases: []string{"v"},
		Usage:   "vender name of mashable application component",
	}
	wireCloudNameFlag = &ngsicli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   "name of mashable application component",
	}
	wireCloudVersionFlag = &ngsicli.StringFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "version of mashable application component",
	}
	wireCloudFileRFlag = &ngsicli.StringFlag{
		Name:     "file",
		Aliases:  []string{"f"},
		Usage:    "mashable application component file",
		Required: true,
	}
	wireCloudPublicFlag = &ngsicli.BoolFlag{
		Name:    "public",
		Aliases: []string{"p"},
		Usage:   "install mashable application component as public",
	}
	wireCloudOverwriteFlag = &ngsicli.BoolFlag{
		Name:    "overwrite",
		Aliases: []string{"o"},
		Usage:   "overwrite mashable application component",
	}
	wireCloudWidgetsFlag = &ngsicli.BoolFlag{
		Name:    "widgets",
		Aliases: []string{"W"},
		Usage:   "list widgets",
	}
	wireCloudOperatorsFlag = &ngsicli.BoolFlag{
		Name:    "operators",
		Aliases: []string{"o"},
		Usage:   "list operators",
	}
	wireCloudTabsFlag = &ngsicli.BoolFlag{
		Name:    "tabs",
		Aliases: []string{"t"},
		Usage:   "list tabs",
	}
	wireCloudUsersFlag = &ngsicli.BoolFlag{
		Name:    "users",
		Aliases: []string{"u"},
		Usage:   "list users",
	}
)
