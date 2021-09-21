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

package keyrock

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestNewNgsiApp(t *testing.T) {
	actual := NewNgsiApp()

	assert.NotEqual(t, nil, actual)
}

func TestNGSICommand(t *testing.T) {
	cases := []struct {
		args  []string
		flags []ngsicli.Flag
		rc    int
	}{
		{args: []string{"applications", "get", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "create", "--host", "keyrock"}, rc: 1},
		{args: []string{"applications", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "update", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "delete", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "roles", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "roles", "list", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "roles", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "roles", "get", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "roles", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", "@"}, rc: 1},
		{args: []string{"applications", "roles", "create", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--data", "@"}, rc: 1},
		{args: []string{"applications", "roles", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", "@"}, rc: 1},
		{args: []string{"applications", "roles", "update", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--data", "@"}, rc: 1},
		{args: []string{"applications", "roles", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "roles", "delete", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "roles", "permissions", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "roles", "permissions", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "roles", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"applications", "roles", "assign", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"applications", "roles", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--rid", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "--pid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"applications", "roles", "unassign", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33fd15c0-e919-47b0-9e05-5f47999f6d91", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"applications", "permissions", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "permissions", "list", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "permissions", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "permissions", "get", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33fd15c0-e919-47b0-9e05-5f47999f6d91"}, rc: 1},
		{args: []string{"applications", "permissions", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "permissions", "create", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "permissions", "update", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"}, rc: 1},
		{args: []string{"applications", "permissions", "update", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"}, rc: 1},
		{args: []string{"applications", "permissions", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--pid", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"}, rc: 1},
		{args: []string{"applications", "permissions", "delete", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "15ca810b-27d1-44a1-8491-a3fb4b6bc6f3"}, rc: 1},
		{args: []string{"applications", "pep", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "pep", "list", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "pep", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "pep", "create", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "pep", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "pep", "reset", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "pep", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "pep", "delete", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "iota", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "iota", "list", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "iota", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"}, rc: 1},
		{args: []string{"applications", "iota", "get", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "iot_sensor_47886fa2-883e-4550-bb56-0138ae9862b7"}, rc: 1},
		{args: []string{"applications", "iota", "create", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "iota", "create", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "iota", "reset", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}, rc: 1},
		{args: []string{"applications", "iota", "reset", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}, rc: 1},
		{args: []string{"applications", "iota", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--iid", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}, rc: 1},
		{args: []string{"applications", "iota", "delete", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "iot_sensor_9d209e51-bff1-4a20-a0f1-d75706f95b04"}, rc: 1},
		{args: []string{"applications", "users", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "users", "list", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "users", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"applications", "users", "get", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"applications", "users", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"}, rc: 1},
		{args: []string{"applications", "users", "assign", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "purchaser"}, rc: 1},
		{args: []string{"applications", "users", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "--rid", "purchaser"}, rc: 1},
		{args: []string{"applications", "users", "unassign", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "2d6f5391-6130-48d8-a9d0-01f20699a7eb", "purchaser"}, rc: 1},
		{args: []string{"applications", "organizations", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "organizations", "list", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "organizations", "get", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec"}, rc: 1},
		{args: []string{"applications", "organizations", "get", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec"}, rc: 1},
		{args: []string{"applications", "organizations", "assign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"}, rc: 1},
		{args: []string{"applications", "organizations", "assign", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "provider", "owner"}, rc: 1},
		{args: []string{"applications", "organizations", "unassign", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--oid", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "--rid", "provider", "--orid", "owner"}, rc: 1},
		{args: []string{"applications", "organizations", "unassign", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "33cf4d3c-8dfb-4bed-bf37-7647f45528ec", "provider", "owner"}, rc: 1},
		{args: []string{"applications", "trusted", "list", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "trusted", "list", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c"}, rc: 1},
		{args: []string{"applications", "trusted", "add", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"applications", "trusted", "add", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"applications", "trusted", "delete", "--host", "keyrock", "--aid", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "--tid", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"applications", "trusted", "delete", "--host", "keyrock", "0fbfa58c-e5b6-41c3-b748-ab29f1567a9c", "0118ccb7-756e-42f9-8a19-5b4e83ca8c46"}, rc: 1},
		{args: []string{"organizations", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"organizations", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "get", "--host", "keyrock", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "create", "--host", "keyrock", "--name", "Testorganization"}, rc: 1},
		{args: []string{"organizations", "create", "--host", "keyrock", "Testorganization"}, rc: 1},
		{args: []string{"organizations", "update", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "update", "--host", "keyrock", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "delete", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "delete", "--host", "keyrock", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "users", "list", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "users", "list", "--host", "keyrock", "3e20722f-d420-422d-89ba-3ae87bc1c0cd"}, rc: 1},
		{args: []string{"organizations", "users", "get", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin"}, rc: 1},
		{args: []string{"organizations", "users", "get", "--host", "keyrock", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "admin"}, rc: 1},
		{args: []string{"organizations", "users", "add", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"}, rc: 1},
		{args: []string{"organizations", "users", "add", "--host", "keyrock", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "admin", "owner"}, rc: 1},
		{args: []string{"organizations", "users", "remove", "--host", "keyrock", "--oid", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "--uid", "admin", "--orid", "owner"}, rc: 1},
		{args: []string{"organizations", "users", "remove", "--host", "keyrock", "3e20722f-d420-422d-89ba-3ae87bc1c0cd", "admin", "owner"}, rc: 1},
		{args: []string{"users", "list", "--host", "keyrock"}, rc: 1},
		{args: []string{"users", "get", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"users", "get", "--host", "keyrock", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"users", "create", "--host", "keyrock", "--username", "alice", "--email", "alice@test.com", "--password", "passw0rd"}, rc: 1},
		{args: []string{"users", "create", "--host", "keyrock", "alice", "alice@test.com", "passw0rd"}, rc: 1},
		{args: []string{"users", "update", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"users", "update", "--host", "keyrock", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"users", "delete", "--host", "keyrock", "--uid", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"users", "delete", "--host", "keyrock", "2d6f5391-6130-48d8-a9d0-01f20699a7eb"}, rc: 1},
		{args: []string{"providers", "--host", "keyrock"}, rc: 1},
	}

	for _, c := range cases {
		ngsi := helper.SetupTestInitNGSI()

		ngsi.HTTP = &helper.MockHTTP{ReqRes: []helper.MockHTTPReqRes{{StatusCode: http.StatusBadRequest}}}
		syslog := []string{"ngsi", "--stderr", "off"}
		args := append(syslog, c.args...)

		app := NewNgsiApp()

		err := app.Run(args)
		rc := 0
		if err != nil {
			if err.(*ngsierr.NgsiError).Message == "missing required flags" {
				t.Error(strings.Join(args, "\", \""))
			}
			rc = 1
		}

		if rc != c.rc {
			fmt.Printf("*** %s *** rc expected:%d, actual:%d)\n", strings.Join(c.args, " "), c.rc, rc)
		}
	}
}
