/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

package ngsimain

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/convenience"
	"github.com/lets-fiware/ngsi-go/internal/cygnus"
	"github.com/lets-fiware/ngsi-go/internal/iotagent"
	"github.com/lets-fiware/ngsi-go/internal/keyrock"
	"github.com/lets-fiware/ngsi-go/internal/management"
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsicmd"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/lets-fiware/ngsi-go/internal/perseo"
	"github.com/lets-fiware/ngsi-go/internal/timeseries"
	"github.com/lets-fiware/ngsi-go/internal/wirecloud"
)

// Version  has a version number of NGSI Go
var Version = ""

// Revision has a git hash value
var Revision = ""

const copyright = "(c) 2020-2024 Kazuhito Suda"

var usage = "command-line tool for FIWARE Open APIs"

// Run is a main rouitne of NGSI Go
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	ngsi := ngsilib.NewNGSI()
	defer ngsi.Close()

	bufStdin := bufio.NewReader(stdin)
	bufStdout := bufio.NewWriter(stdout)
	ngsi.InitLog(bufStdin, bufStdout, stderr)

	ngsicli.Version = Version
	ngsicli.Revision = Revision

	Version = fmt.Sprintf("%s (git_hash:%s)", Version, Revision)

	app := NewNgsiApp()

	err := app.Run(args)
	if err != nil {
		s := strings.TrimRight(ngsierr.Message(err), "\n") + "\n"
		ngsi.Logging(ngsilib.LogErr, s)
		for err != nil {
			err = errors.Unwrap(err)
			if err == nil {
				break
			}
			ngsi.Logging(ngsilib.LogDebug, fmt.Sprintf("%T\n", err))
			s = strings.TrimRight(ngsierr.Message(err), "\n") + "\n"
			ngsi.Logging(ngsilib.LogInfo, s)
		}
		ngsi.Logging(ngsilib.LogInfo, "abnormal termination\n")
		_ = bufStdout.Flush()
		return 1
	}
	ngsi.Logging(ngsilib.LogInfo, "normal termination\n")

	_ = bufStdout.Flush()

	return 0
}

func NewNgsiApp() *ngsicli.App {
	return &ngsicli.App{
		Copyright: copyright,
		Version:   Version,
		Usage:     usage,
		Flags:     ngsicli.GlobalFlags,
		Commands: []*ngsicli.Command{
			&convenience.AdminCmd,
			&convenience.ApisCmd,
			&ngsicmd.AppendCmd,
			&management.BrokersCmd,
			&management.ContextCmd,
			&convenience.CopyCmd,
			&ngsicmd.CountCmd,
			&ngsicmd.CreateCmd,
			&convenience.DebugCmd,
			&ngsicmd.DeleteCmd,
			&iotagent.DevicesCmd,
			&convenience.DocumentsCmd,
			&ngsicmd.GetCmd,
			&timeseries.HDeleteCmd,
			&timeseries.HGetCmd,
			&convenience.HealthCmd,
			&ngsicmd.ListCmd,
			&ngsicmd.LsCmd,
			&convenience.QueryProxyCmd,
			&convenience.RemoveCmd,
			&convenience.ReceiverCmd,
			&convenience.RegProxyCmd,
			&ngsicmd.ReplaceCmd,
			&perseo.RulesCmd,
			&management.SettingsCmd,
			&management.ServerCmd,
			&iotagent.ServicesCmd,
			&ngsicmd.TemplateCmd,
			&management.TokenCmd,
			&management.LicenseCmd,
			&convenience.TokenProxyCmd,
			&ngsicmd.UpdateCmd,
			&ngsicmd.UpsertCmd,
			&convenience.VersionCmd,
			&keyrock.ApplicationsCmd,
			&keyrock.UsersCmd,
			&keyrock.OrganizationsCmd,
			&keyrock.ProvidersCmd,
			&cygnus.NamemappingsCmd,
			&cygnus.GroupingrulesCmd,
			&wirecloud.PreferencesCmd,
			&wirecloud.ResourcesCmd,
			&wirecloud.WorkspacesCmd,
			&wirecloud.TabsCmd,
		},
	}
}
