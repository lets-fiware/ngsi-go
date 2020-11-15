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

const registrationTemplateLd string = `{
	"type": "ContextSourceRegistration",
	"description": "registration template",
    "information": [
        {
            "entities": [
                {
                    "type": "Registration",
                    "id": "urn:ngsi-ld:Registration:001"
                }
            ],
            "properties": [
                "attr"
            ]
        }
    ],
    "endpoint": "http://registration"
}`

func registrationsListLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registratinsListLd"

	page := 0
	count := 0
	limit := 100
	total := 0

	var registrations []map[string]interface{}

	for {
		client.SetPath("/csourceRegistrations")

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

func registrationsGetLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registratinGetLd"

	id := c.String("id")
	client.SetPath("/csourceRegistrations/" + id)

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

func registrationsCreateLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "registrationsCreateLd"

	client.SetPath("/csourceRegistrations")

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
	p := "/ngsi-ld/v1/csourceRegistrations/"
	if strings.HasPrefix(location, p) {
		location = location[len(p):]
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is created, FIWARE-Service: %s, FIWARE-ServicePath: %s",
		res.Header.Get("Location"), c.String("service"), c.String("path")))

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

	var template = make(map[string]interface{})
	ngsilib.JSONUnmarshal([]byte(registrationTemplateLd), &template)

	if c.IsSet("description") {
		template["description"] = c.String("description")
	}
	if c.IsSet("type") {
		template["information"].([]interface{})[0].(map[string]interface{})["entities"].([]interface{})[0].(map[string]interface{})["type"] = c.String("type")
	}
	if c.IsSet("id") {
		template["information"].([]interface{})[0].(map[string]interface{})["entities"].([]interface{})[0].(map[string]interface{})["id"] = c.String("id")
	}
	if c.IsSet("attrs") {
		s := c.String("attrs")
		template["information"].([]interface{})[0].(map[string]interface{})["properties"] = strings.Split(s, ",")
	}
	if c.IsSet("provider") {
		s := c.String("provider")
		if ngsilib.IsHTTP(s) {
			template["endpoint"] = s
		} else {
			e := fmt.Sprintf("provider url error: %s", s)
			return &ngsiCmdError{funcName, 1, e, nil}
		}
	}

	b, err := ngsilib.JSONMarshal(template)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	fmt.Fprint(ngsi.StdWriter, string(b))

	return nil
}
