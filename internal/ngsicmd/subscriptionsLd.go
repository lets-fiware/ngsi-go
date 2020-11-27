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

const subscriptionTemplateLD string = `{
	"description": "description",
	"type": "Subscription",
	"entities": [{"type": "Template"}],
	"watchedAttributes": ["watchedAttribute"],
	"notification": {
	  "attributes": ["attribute"],
	  "format": "normalized",
	  "endpoint": {
		"uri": "http://template",
		"accept": "application/ld+json"
	  }
	}
  }`

func subscriptionsListLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsListLd"

	page := 0
	count := 0
	limit := 100
	total := 0

	var subscriptions []map[string]interface{}

	for {
		client.SetPath("/subscriptions")

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
		subscriptions = append(subscriptions, subs...)

		total += len(subs)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	if c.IsSet("json") {
		b, err := ngsilib.JSONMarshal(subscriptions)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, string(b))
	} else if c.IsSet("verbose") {
		for _, e := range subscriptions {
			fmt.Fprintf(ngsi.StdWriter, "%s %s\n", e["id"].(string), e["description"].(string))
		}
	} else {
		for _, e := range subscriptions {
			fmt.Fprintln(ngsi.StdWriter, e["id"].(string))
		}
	}

	return nil
}

func subscriptionGetLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionGetLd"

	id := c.String("id")
	client.SetPath("/subscriptions/" + id)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s %s", res.Status, string(body), id), nil}
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

func subscriptionsCreateLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsCreateLd"

	client.SetPath("/subscriptions")

	client.SetContentType()

	b, err := readAll(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
	}

	location := res.Header.Get("Location")
	p := "/ngsi-ld/v1/subscriptions/"
	if strings.HasPrefix(location, p) {
		location = location[len(p):]
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is created", res.Header.Get("Location")))

	fmt.Fprintln(ngsi.StdWriter, location)

	return nil
}

func subscriptionsUpdateLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsUpdateLd"

	return &ngsiCmdError{funcName, 1, "not yet implemented", nil}
}

func subscriptionsDeleteLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsDeleteLd"

	id := c.String("id")
	path := "/subscriptions/" + id
	client.SetPath(path)

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s %s", res.Status, string(body), id), nil}
	}

	return nil
}

func subscriptionsTemplateLd(c *cli.Context, ngsi *ngsilib.NGSI) error {
	const funcName = "subscriptionsTemplateLd"

	var template = make(map[string]interface{})
	ngsilib.JSONUnmarshalEncode([]byte(subscriptionTemplateLD), &template, false)

	if c.IsSet("type") {
		template["entities"].([]interface{})[0].(map[string]interface{})["type"] = c.String("type")
	}
	if c.IsSet("uri") {
		s := c.String("uri")
		if ngsilib.IsHTTP(s) {
			template["notification"].(map[string]interface{})["endpoint"].(map[string]interface{})["uri"] = s
		} else {
			e := fmt.Sprintf("notification url error: %s", s)
			return &ngsiCmdError{funcName, 1, e, nil}
		}
	}

	if c.IsSet("query") {
		template["q"] = c.String("query")
	}

	if c.IsSet("keyValues") {
		template["notification"].(map[string]interface{})["format"] = "keyValues"
	}

	if c.IsSet("link") {
		link := c.String("link")
		if !ngsilib.IsHTTP(link) {
			value, err := ngsi.GetContext(link)
			if err != nil {
				return &ngsiCmdError{funcName, 2, err.Error(), err}
			}
			link = value
		}
		template["@Context"] = link
	}

	if c.IsSet("wAttrs") {
		attrs := []string{}
		for _, v := range strings.Split(c.String("wAttrs"), ",") {
			attrs = append(attrs, v)
		}
		template["watchedAttributes"] = attrs
	}

	if c.IsSet("nAttrs") {
		attrs := []string{}
		for _, v := range strings.Split(c.String("nAttrs"), ",") {
			attrs = append(attrs, v)
		}
		template["notification"].(map[string]interface{})["attributes"] = attrs
	}

	if c.IsSet("description") {
		s := c.String("description")
		template["description"] = s
	}

	b, err := ngsilib.JSONMarshal(template)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	fmt.Fprint(ngsi.StdWriter, string(b))

	return nil
}
