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
	"net/url"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"

	"github.com/urfave/cli/v2"
)

// lib/apiTypesV2/Registration.h in fiware-orion

type registrationEntitiesV2 struct {
	ID          string `json:"id,omitempty"`
	IDPattern   string `json:"idPattern,omitempty"`
	Type        string `json:"type,omitempty"`
	TypePattern string `json:"typePattern,omitempty"`
}

type registrationDataProvidedV2 struct {
	Entities []registrationEntitiesV2 `json:"entities,omitempty"`
	Atts     []string                 `json:"attrs,omitempty"`
}

type registrationHTTPV2 struct {
	URL string `json:"url,omitempty"`
}

type registrationProviderV2 struct {
	HTTP                    *registrationHTTPV2 `json:"http,omitempty"`
	SupportedForwardingMode string              `json:"supportedForwardingMode,omitempty"`
	LegacyForwarding        *bool               `json:"legacyForwarding,omitempty"`
}

type registrationForwardingInformationV2 struct {
	TimesSent      *int64 `json:"timesSent,omitempty"`
	LastForwarding string `json:"lastForwarding,omitempty"`
	LastSuccess    string `json:"lastSuccess,omitempty"`
	LastFailure    string `json:"lastFailure,omitempty"`
}

type registrationQueryV2 struct {
	Description  string                      `json:"description,omitempty"`
	DataProvided *registrationDataProvidedV2 `json:"dataProvided,omitempty"`
	Provider     *registrationProviderV2     `json:"provider,omitempty"`
	Expires      string                      `json:"expires,omitempty"`
	Status       string                      `json:"status,omitempty"`
}

type registrationResposeV2 struct {
	ID                    string                               `json:"id"`
	Description           string                               `json:"description,omitempty"`
	DataProvided          *registrationDataProvidedV2          `json:"dataProvided,omitempty"`
	Provider              *registrationProviderV2              `json:"provider,omitempty"`
	Expires               string                               `json:"expires,omitempty"`
	Status                string                               `json:"status,omitempty"`
	ForwardingInformation *registrationForwardingInformationV2 `json:"forwardingInformation,omitempty"`
}

func registrationsListV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registratinsListV2"

	page := 0
	count := 0
	limit := 100
	total := 0

	var registrations []registrationResposeV2

	for {
		client.SetPath("/registrations")

		v := url.Values{}
		v.Set("options", "count")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		client.SetQuery(&v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}
		count, err = client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 3, "ResultsCount error", err}
		}
		if count == 0 {
			break
		}
		var regs []registrationResposeV2
		if err := ngsilib.JSONUnmarshalDecode(body, &regs, client.IsSafeString()); err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		registrations = append(registrations, regs...)

		total += len(regs)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	if c.IsSet("json") || c.Bool("pretty") {
		if len(registrations) > 0 {
			b, err := ngsilib.JSONMarshal(registrations)
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
				if err != nil {
					return &ngsiCmdError{funcName, 6, err.Error(), err}
				}
				fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			} else {
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		}
	} else if c.IsSet("verbose") {
		local := c.IsSet("localTime")
		for _, e := range registrations {
			if local {
				toLocaltimeRegistration(&e)
			}
			fmt.Fprintf(ngsi.StdWriter, "%s %s %s\n", e.ID, e.Description, e.Expires)
		}
	} else {
		for _, e := range registrations {
			fmt.Fprintln(ngsi.StdWriter, e.ID)
		}
	}

	return nil
}

func registrationsGetV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsGetV2"

	id := c.String("id")
	client.SetPath("/registrations/" + id)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s %s", id, res.Status, string(body)), nil}
	}

	var r registrationResposeV2
	if err := ngsilib.JSONUnmarshalDecode(body, &r, client.IsSafeString()); err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	if c.IsSet("localTime") {
		toLocaltimeRegistration(&r)
	}
	b, err := ngsilib.JSONMarshal(&r)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}
	fmt.Fprint(ngsi.StdWriter, string(b))

	return nil
}

func registrationsCreateV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsCreateV2"

	client.SetPath("/registrations")

	client.SetHeader("Content-Type", "application/json")

	var r registrationQueryV2

	if err := setRegistrationsValuleV2(c, ngsi, &r); err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	b, err := ngsilib.JSONMarshalEncode(&r, true)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	location := res.Header.Get("Location")
	p := "/v2/registrations/"
	location = strings.TrimPrefix(location, p)

	fmt.Fprintln(ngsi.StdWriter, location)

	return nil
}

func registrationsDeleteV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsDeleteV2"

	id := c.String("id")

	path := "/registrations/" + id
	client.SetPath(path)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s %s", id, res.Status, string(body)), nil}
	}

	return nil
}

func registrationsTemplateV2(c *cli.Context, ngsi *ngsilib.NGSI) error {
	const funcName = "registrationsTemplateV2"

	var r registrationQueryV2

	if err := setRegistrationsValuleV2(c, ngsi, &r); err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	b, err := ngsilib.JSONMarshal(r)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(b))
	return nil
}

func setRegistrationsValuleV2(c *cli.Context, ngsi *ngsilib.NGSI, r *registrationQueryV2) error {
	const funcName = "setRegistrationsValuleV2"

	if c.IsSet("data") {
		b, err := readAll(c, ngsi)
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		err = ngsilib.JSONUnmarshal(b, r)
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
	}

	if c.IsSet("description") {
		r.Description = c.String("description")
	}

	if c.IsSet("providedId") || c.IsSet("idPattern") || c.IsSet("type") || c.IsSet("attrs") {
		if r.DataProvided == nil {
			r.DataProvided = new(registrationDataProvidedV2)
		}
		if len(r.DataProvided.Entities) == 0 {
			r.DataProvided.Entities = append(r.DataProvided.Entities, *new(registrationEntitiesV2))
		}
	}

	if c.IsSet("providedId") {
		r.DataProvided.Entities[0].ID = c.String("providedId")
	}
	if c.IsSet("idPattern") {
		r.DataProvided.Entities[0].IDPattern = c.String("idPattern")
	}
	if c.IsSet("type") {
		r.DataProvided.Entities[0].Type = c.String("type")
	}
	if c.IsSet("attrs") {
		s := c.String("attrs")
		r.DataProvided.Atts = strings.Split(s, ",")
	}

	if c.IsSet("provider") || c.IsSet("legacy") || c.IsSet("forwardingModeFlag") {
		if r.Provider == nil {
			r.Provider = new(registrationProviderV2)
		}
		if r.Provider.HTTP == nil {
			r.Provider.HTTP = new(registrationHTTPV2)
		}
		s := c.String("provider")
		if ngsilib.IsHTTP(s) {
			r.Provider.HTTP.URL = s
		} else {
			e := fmt.Sprintf("provider url error: %s", s)
			return &ngsiCmdError{funcName, 1, e, nil}
		}

		if c.IsSet("legacy") {
			legacy := c.Bool("legacy")
			if legacy {
				r.Provider.LegacyForwarding = &legacy
			}
		}

		if c.IsSet("forwardingModeFlag") {
			mode := c.String("forwardingModeFlag")
			if !ngsilib.Contains([]string{"all", "none", "query", "update"}, mode) {
				return &ngsiCmdError{funcName, 3, "unknown mode: " + mode, nil}
			}
			r.Provider.SupportedForwardingMode = mode
		}
	}

	if c.IsSet("expires") {
		s := c.String("expires")
		if !ngsilib.IsOrionDateTime(s) {
			var err error
			s, err = ngsilib.GetExpirationDate(s)
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), nil}
			}
		}
		r.Expires = s
	}

	if c.IsSet("status") {
		status := c.String("status")
		if !ngsilib.Contains([]string{"active", "inactive"}, status) {
			return &ngsiCmdError{funcName, 5, "unknown status: " + status, nil}
		}
		r.Status = status
	}

	return nil
}

func toLocaltimeRegistration(reg *registrationResposeV2) {
	reg.Expires = getLocalTime(reg.Expires)
	if reg.ForwardingInformation != nil {
		reg.ForwardingInformation.LastForwarding = getLocalTime(reg.ForwardingInformation.LastForwarding)
		reg.ForwardingInformation.LastSuccess = getLocalTime(reg.ForwardingInformation.LastSuccess)
		reg.ForwardingInformation.LastFailure = getLocalTime(reg.ForwardingInformation.LastFailure)
	}
	reg.Expires = getLocalTime(reg.Expires)
}
