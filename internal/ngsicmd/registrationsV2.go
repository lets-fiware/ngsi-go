/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"

	"github.com/urfave/cli/v2"
)

type registrationQuery struct {
	Description  string `json:"description"`
	DataProvided struct {
		Entities []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"entities"`
		Attrs []string `json:"attrs"`
	} `json:"dataProvided"`
	Provider struct {
		HTTP struct {
			URL string `json:"url"`
		} `json:"http"`
	} `json:"provider"`
}

const registrationTemplateV2 string = `{
	"description": "Registration template",
	"dataProvided": {
	  "entities": [
		{
		  "id": "",
		  "type": "Room"
		}
	  ],
	  "attrs": [
		"attr"
	  ]
	},
	"provider": {
	  "http": {
		"url": "http://localhost:1234"
	  }
	}
}`

func registrationsListV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registratinsListV2"

	page := 0
	count := 0
	limit := 100
	total := 0

	var registrations []map[string]interface{}

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
		var subs []map[string]interface{}
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

	if c.IsSet("json") {
		b, err := ngsilib.JSONMarshal(registrations)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, string(b))
	} else if c.IsSet("verbose") {
		for _, e := range registrations {
			fmt.Fprintf(ngsi.StdWriter, "%s %s\n", e["id"].(string), e["description"].(string))
		}
	} else {
		for _, e := range registrations {
			fmt.Fprintln(ngsi.StdWriter, e["id"].(string))
		}
	}

	return nil
}

func registrationsGetV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registratinGetV2"

	id := c.String("id")
	client.SetPath("/registrations/" + id)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s %s", id, res.Status, string(body)), nil}
	}

	if client.IsSafeString() {
		body, err = ngsilib.JSONSafeStringDecode(body)
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}
	}
	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func registrationsCreateV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsCreateV2"

	client.SetPath("/registrations")

	client.SetHeader("Content-Type", "application/json")

	s, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	res, body, err := client.HTTPPost(s)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	location := res.Header.Get("Location")
	p := "/v2/registrations/"
	if strings.HasPrefix(location, p) {
		location = location[len(p):]
	}

	fmt.Fprintln(ngsi.StdWriter, location)

	return nil
}

func registrationsDeleteV2(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsDeleteV2"

	id := c.String("id")

	path := "/registrations/" + id
	client.SetPath(path)

	res, body, err := client.HTTPDelete()
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

	var template registrationQuery
	ngsilib.JSONUnmarshal([]byte(registrationTemplateV2), &template)

	if c.IsSet("description") {
		template.Description = c.String("description")
	}
	if c.IsSet("id") {
		template.DataProvided.Entities[0].ID = c.String("id")
	}
	if c.IsSet("type") {
		template.DataProvided.Entities[0].Type = c.String("type")
	}
	if c.IsSet("attrs") {
		s := c.String("attrs")
		template.DataProvided.Attrs = strings.Split(s, ",")
	}
	if c.IsSet("provider") {
		s := c.String("provider")
		if ngsilib.IsHTTP(s) {
			template.Provider.HTTP.URL = s
		} else {
			e := fmt.Sprintf("provider url error: %s", s)
			return &ngsiCmdError{funcName, 1, e, nil}
		}
	}

	b, err := ngsilib.JSONMarshalEncode(&template, true)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	fmt.Fprintln(ngsi.StdWriter, string(b))

	return nil
}
