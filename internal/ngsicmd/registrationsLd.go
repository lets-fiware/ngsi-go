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

// 4.7 GeoJSON geometry
type geometry struct {
	Type        string        `json:"entities,omitempty"`
	Coordinates []interface{} `json:"coordinates,omitempty"`
}

// 5.2.9 CsourceRegistration
type cSourceRegistration struct {
	ID                  string             `json:"id,omitempty"`
	Type                string             `json:"type,omitempty"`
	Name                string             `json:"name,omitempty"` // registrationName???
	Description         string             `json:"description,omitempty"`
	Information         []registrationInfo `json:"registrationInfo,omitempty"`
	Tenant              string             `json:"tenant,omitempty"`
	ObservationInterval *timeInterval      `json:"observationInterval,omitempty"`
	ManagementInterval  *timeInterval      `json:"ManagementInterval,omitempty"`
	Location            *geometry          `json:"location,omitempty"`
	ObservationSpace    *geometry          `json:"observationSpace,omitempty"`
	OperationSpace      *geometry          `json:"operationSpace,omitempty"`
	Expires             string             `json:"expires,omitempty"` // expiresAt???
	Endpoint            string             `json:"endpoint,omitempty"`
	AtContext           interface{}        `json:"@context,omitempty"`
}

// 5.2.10 RegistrationInfo
type registrationInfo struct {
	Entities      []entityInfoLd `json:"entities,omitempty"`
	Properties    []string       `json:"properties,omitempty"`    // PropertyNames???
	Relationships []string       `json:"relationships,omitempty"` // RelationshipNames???
}

// 5.2.11 TimeInterval

type timeInterval struct {
	StartAt string `json:"startAt,omitempty"`
	EndAt   string `json:"endAt,omitempty"`
}

func registrationsListLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registratinsListLd"

	page := 0
	count := 0
	limit := 100
	total := 0

	var registrations []cSourceRegistration

	for {
		client.SetPath("/csourceRegistrations/")

		v := url.Values{}
		v.Set("count", "true")
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
		var subs []cSourceRegistration
		if err := ngsilib.JSONUnmarshalDecode(body, &subs, client.IsSafeString()); err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		registrations = append(registrations, subs...)

		total += len(subs)

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
				fmt.Fprintln(ngsi.StdWriter, string(newBuf.Bytes()))
			} else {
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		}
	} else if c.IsSet("verbose") {
		local := c.IsSet("localTime")
		for _, e := range registrations {
			if local {
				e.Expires = getLocalTime(e.Expires)
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

func registrationsGetLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsGetLd"

	id := c.String("id")
	client.SetPath("/csourceRegistrations/" + id)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s %s", id, res.Status, string(body)), nil}
	}

	var r cSourceRegistration
	if err := ngsilib.JSONUnmarshalDecode(body, &r, client.IsSafeString()); err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	if c.IsSet("localTime") {
		r.Expires = getLocalTime(r.Expires)
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
		fmt.Fprintln(ngsi.StdWriter, string(newBuf.Bytes()))
		return nil
	}
	fmt.Fprint(ngsi.StdWriter, string(b))

	return nil
}

func registrationsCreateLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsCreateLd"

	client.SetPath("/csourceRegistrations")

	client.SetHeader("Content-Type", "application/json")

	var r cSourceRegistration

	if err := setRegistrationsValuleLd(c, ngsi, &r); err != nil {
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
	p := "/ngsi-ld/v1/csourceRegistrations/"
	if strings.HasPrefix(location, p) {
		location = location[len(p):]
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is created, FIWARE-ServicePath: %s", res.Header.Get("Location"), c.String("service")))

	fmt.Fprintln(ngsi.StdWriter, location)

	return nil
}

func registrationsDeleteLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsDeleteLd"

	id := c.String("id")

	path := "/csourceRegistrations/" + id
	client.SetPath(path)

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s %s", id, res.Status, string(body)), nil}
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is deleted, FIWARE-Service: %s, FIWARE-ServicePath: %s",
		path, c.String("service"), c.String("path")))

	return nil
}

func registrationsTemplateLd(c *cli.Context, ngsi *ngsilib.NGSI) error {
	const funcName = "registrationsTemplateLd"

	var r cSourceRegistration

	err := setRegistrationsValuleLd(c, ngsi, &r)
	if err != nil {
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
		fmt.Fprintln(ngsi.StdWriter, string(newBuf.Bytes()))
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(b))

	return nil
}

func setRegistrationsValuleLd(c *cli.Context, ngsi *ngsilib.NGSI, r *cSourceRegistration) error {
	const funcName = "setRegistrationsValuleLd"

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

	if isSetOR(c, []string{"type", "providedId", "idPattern", "properties", "relationships"}) {
		if len(r.Information) == 0 {
			r.Information = append(r.Information, *new(registrationInfo))
		}
		if len(r.Information[0].Entities) == 0 {
			r.Information[0].Entities = append(r.Information[0].Entities, *new(entityInfoLd))
		}
		if c.IsSet("type") {
			r.Information[0].Entities[0].Type = c.String("type")
		}
		if c.IsSet("providedId") {
			r.Information[0].Entities[0].ID = c.String("providedId")
		}
		if c.IsSet("idPattern") {
			r.Information[0].Entities[0].IDPattern = c.String("idPattern")
		}
		if c.IsSet("properties") {
			s := c.String("properties")
			r.Information[0].Properties = strings.Split(s, ",")
		}
		if c.IsSet("relationships") {
			s := c.String("relationships")
			r.Information[0].Relationships = strings.Split(s, ",")
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

	if c.IsSet("provider") {
		s := c.String("provider")
		if ngsilib.IsHTTP(s) {
			r.Endpoint = s
		} else {
			e := fmt.Sprintf("provider url error: %s", s)
			return &ngsiCmdError{funcName, 5, e, nil}
		}
	}

	if c.IsSet("context") {
		context := c.String("context")
		var atContext interface{}
		atContext, err := getAtContext(ngsi, context)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), nil}
		}
		r.AtContext = atContext
	}

	return nil
}
