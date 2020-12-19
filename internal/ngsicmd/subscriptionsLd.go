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
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

// lib/orionld/kjTree/kjTreeFromSubscription.cpp

// 5.2.8 EntityInfo
type entityInfoLd struct {
	ID        string `json:"id,omitempty"`
	IDPattern string `json:"idPattern,omitempty"`
	Type      string `json:"type,omitempty"`
}

// 5.2.13 GeoQuery
type geoQueryLd struct {
	Geometry    string      `json:"geometry,omitempty"`
	Coordinates interface{} `json:"coordinates,omitempty"` // JSON Array or string
	Georel      string      `json:"georel,omitempty"`
	Geoproperty string      `json:"geoproperty,omitempty"`
}

// 5.2.14 NotificationParams
type notificationParamsLd struct {
	Attributes []string    `json:"attributes,omitempty"`
	Format     string      `json:"format,omitempty"`
	Endpoint   *endpointLd `json:"endpoint,omitempty"`
	Status     string      `json:"status,omitempty"`

	TimesSent        *int64 `json:"timesSent,omitempty"`
	LastNotification string `json:"lastNotification,omitempty"`
	LastFailure      string `json:"lastFailure,omitempty"`
	LastSuccess      string `json:"lastSuccess,omitempty"`
}

// 5.2.15 Endpoint
type endpointLd struct {
	URI          string                 `json:"uri"`
	Accept       string                 `json:"accept,omitempty"`
	ReceiverInfo map[string]interface{} `json:"receiverInfo,omitempty"` // KeyValue Array
	NotifierInfo map[string]interface{} `json:"notifierInfo,omitempty"` // KeyValue Array
}

// 5.2.21 TemporalQuery
type temporalQueryLd struct {
	Timerel      string `json:"timerel"`
	TimeAt       string `json:"timeAt"`
	EndTimeAt    string `json:"endTimeAt,omitempty"`
	Timeproperty string `json:"timeproperty,omitempty"`
}

// 5.2.12 Subscription
type subscriptionLd struct {
	ID                string                `json:"id,omitempty"`
	Type              string                `json:"type,omitempty"`
	Name              string                `json:"name,omitempty"` // subscriptionName???
	Description       string                `json:"description,omitempty"`
	CreatedAt         string                `json:"createdAt,omitempty"`
	ModifiedAt        string                `json:"modifiedAt,omitempty"`
	Entities          []entityInfoLd        `json:"entities,omitempty"`
	WatchedAttributes []string              `json:"watchedAttributes,omitempty"`
	TimeInterval      *int64                `json:"timeInterval,omitempty"`
	Q                 string                `json:"q,omitempty"`
	GeoQ              *geoQueryLd           `json:"geoQ,omitempty"`
	Csf               string                `json:"csf,omitempty"`
	IsActive          *bool                 `json:"isActive,omitempty"`
	Notification      *notificationParamsLd `json:"notification,omitempty"`
	Expires           string                `json:"expires,omitempty"` // expiresAt???
	Throttling        *int64                `json:"throttling,omitempty"`
	TemporalQ         *temporalQueryLd      `json:"temporalQ,omitempty"`
	AtContext         interface{}           `json:"@context,omitempty"`
	Status            string                `json:"status,omitempty"`
}

func subscriptionsListLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsListLd"

	filters := []string{}

	if c.IsSet("status") {
		status := []string{"active", "paused", "expired"}
		filters = strings.Split(c.String("status"), ",")
		for i, v := range filters {
			filters[i] = strings.ToLower(v)
			if !ngsilib.Contains(status, filters[i]) {
				return &ngsiCmdError{funcName, 1, "error: " + filters[i] + " (active, paused, expired)", nil}
			}
		}
	}

	page := 0
	count := 0
	limit := 100
	total := 0

	var subscriptions []subscriptionLd

	for {
		client.SetPath("/subscriptions/")

		v := url.Values{}
		v.Set("count", "true")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		client.SetQuery(&v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}
		count, err = client.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 4, "ResultsCount error", err}
		}
		if count == 0 {
			break
		}
		var subs []subscriptionLd
		if err := ngsilib.JSONUnmarshalDecode(body, &subs, client.IsSafeString()); err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		subscriptions = append(subscriptions, subs...)

		total += len(subs)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	if c.IsSet("status") {
		var subs []subscriptionLd
		for _, e := range subscriptions {
			if ngsilib.Contains(filters, e.Status) {
				subs = append(subs, e)
			}
		}
		subscriptions = subs
	}

	if c.IsSet("query") {
		reg := regexp.MustCompile(c.String("query"))

		var subs []subscriptionLd
		for _, e := range subscriptions {
			if reg.MatchString(e.Description) {
				subs = append(subs, e)
			}
		}
		subscriptions = subs
	}

	if c.IsSet("count") {
		fmt.Fprintf(ngsi.StdWriter, "%d\n", len(subscriptions))
	} else if c.IsSet("json") || c.Bool("pretty") {
		if len(subscriptions) > 0 {
			b, err := ngsilib.JSONMarshal(subscriptions)
			if err != nil {
				return &ngsiCmdError{funcName, 6, err.Error(), err}
			}
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
				if err != nil {
					return &ngsiCmdError{funcName, 7, err.Error(), err}
				}
				fmt.Fprintln(ngsi.StdWriter, string(newBuf.Bytes()))
			} else {
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		}
	} else if c.IsSet("verbose") {
		items := []string{"id", "status", "expires", "description"}
		var err error
		if c.IsSet("items") {
			items, err = checkItemsLd(c)
			if err != nil {
				return &ngsiCmdError{funcName, 8, err.Error(), err}
			}
		}
		local := c.IsSet("localTime")
		for _, e := range subscriptions {
			if local {
				toLocaltimeLd(&e)
			}
			fmt.Fprintln(ngsi.StdWriter, sprintItemsLd(&e, items))
		}
	} else {
		for _, e := range subscriptions {
			fmt.Fprintln(ngsi.StdWriter, e.ID)
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

	var sub subscriptionLd
	if err := ngsilib.JSONUnmarshalDecode(body, &sub, client.IsSafeString()); err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	if c.IsSet("localTime") {
		toLocaltimeLd(&sub)
	}
	b, err := ngsilib.JSONMarshal(&sub)
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

	fmt.Fprintln(ngsi.StdWriter, string(b))

	return nil
}

func subscriptionsCreateLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsCreateLd"

	client.SetPath("/subscriptions")

	client.SetContentType()

	var t subscriptionLd

	err := setSubscriptionValuesLd(c, ngsi, &t, false)

	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	b, err := ngsilib.JSONMarshalEncode(&t, true)
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
	p := "/ngsi-ld/v1/subscriptions/"
	if strings.HasPrefix(location, p) {
		location = location[len(p):]
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is created, FIWARE-Service: %s", res.Header.Get("Location"), c.String("service")))

	fmt.Fprintln(ngsi.StdWriter, location)

	return nil
}

func subscriptionsUpdateLd(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsUpdateLd"

	id := c.String("id")

	client.SetPath("/subscriptions/" + id)

	client.SetContentType()

	var t subscriptionLd

	if err := setSubscriptionValuesLd(c, ngsi, &t, true); err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	b, err := ngsilib.JSONMarshalEncode(&t, true)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s %s %s", res.Status, string(body), id), nil}
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is updated, FIWARE-Service: %s", id, c.String("service")))

	fmt.Fprintln(ngsi.StdWriter, id)

	return nil
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

	var t subscriptionLd

	if err := setSubscriptionValuesLd(c, ngsi, &t, false); err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	b, err := ngsilib.JSONMarshal(t)
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

func setSubscriptionValuesLd(c *cli.Context, ngsi *ngsilib.NGSI, t *subscriptionLd, checkskip bool) error {
	const funcName = "setSubscriptionValuesLd"

	if c.IsSet("data") {
		b, err := readAll(c, ngsi)
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		err = ngsilib.JSONUnmarshal(b, t)
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
	}

	if c.IsSet("subscriptionId") {
		t.ID = c.String("subscriptionId")
	}

	t.Type = "Subscription"

	if c.IsSet("name") {
		s := c.String("name")
		t.Name = s
	}

	if c.IsSet("description") {
		s := c.String("description")
		t.Description = s
	}

	if isSetOR(c, []string{"entityId", "idPattern", "type"}) {
		if len(t.Expires) == 0 {
			t.Entities = append(t.Entities, *new(entityInfoLd))
		}
		if c.IsSet("entityId") {
			t.Entities[0].ID = c.String("entityId")
		}
		if c.IsSet("idPattern") {
			t.Entities[0].IDPattern = c.String("idPattern")
		}
		if c.IsSet("type") {
			t.Entities[0].Type = c.String("type")
		}
	}

	if c.IsSet("wAttrs") {
		attrs := []string{}
		for _, v := range strings.Split(c.String("wAttrs"), ",") {
			attrs = append(attrs, v)
		}
		t.WatchedAttributes = attrs
	}

	if c.IsSet("timeInterval") {
		v := c.Int64("timeInterval")
		t.TimeInterval = &v
	}

	if c.IsSet("query") {
		t.Q = c.String("query")
	}

	//geoQ
	if isSetOR(c, []string{"geometry", "coords", "georel", "geoproperty"}) {
		if t.GeoQ == nil {
			t.GeoQ = new(geoQueryLd)
		}
		if c.IsSet("geometry") {
			t.GeoQ.Geometry = c.String("geometry")
		}
		if c.IsSet("coords") {
			coords := c.String("coords")
			err := ngsilib.GetJSONArray([]byte(coords), &t.GeoQ.Coordinates)
			if err != nil {
				return &ngsiCmdError{funcName, 3, "coords: " + err.Error(), nil}
			}
		}
		if c.IsSet("georel") {
			t.GeoQ.Georel = c.String("georel")
		}
		if c.IsSet("geoproperty") {
			t.GeoQ.Geoproperty = c.String("geoproperty")
		}
	}

	if c.IsSet("csf") {
		t.Csf = c.String("csf")
	}

	// temporarily pause the subscription
	if c.IsSet("active") && c.IsSet("inactive") {
		return &ngsiCmdError{funcName, 4, "cannot specify both active and inactive options", nil}
	}
	if c.IsSet("active") {
		b := true
		t.IsActive = &b
	}
	if c.IsSet("inactive") {
		b := false
		t.IsActive = &b
	}

	// Notification
	if isSetOR(c, []string{"nAttrs", "keyValues", "uri", "accept"}) {
		if t.Notification == nil {
			t.Notification = new(notificationParamsLd)
		}
		if c.IsSet("nAttrs") {
			attrs := []string{}
			for _, v := range strings.Split(c.String("nAttrs"), ",") {
				attrs = append(attrs, v)
			}
			t.Notification.Attributes = attrs
		}
		if c.IsSet("keyValues") {
			t.Notification.Format = "keyValues"
		}

		if isSetOR(c, []string{"uri", "accept"}) {
			if t.Notification.Endpoint == nil {
				t.Notification.Endpoint = new(endpointLd)
			}
			if c.IsSet("uri") {
				s := c.String("uri")
				if ngsilib.IsHTTP(s) {
					t.Notification.Endpoint.URI = s
				} else {
					e := fmt.Sprintf("notification url error: %s", s)
					return &ngsiCmdError{funcName, 5, e, nil}
				}
			}
			if c.IsSet("accept") {
				a := strings.ToLower(c.String("accept"))
				if a == "json" {
					t.Notification.Endpoint.Accept = "application/json"
				} else if ngsilib.Contains([]string{"ld+json", "ld"}, a) {
					t.Notification.Endpoint.Accept = "application/ld+json"
				} else {
					return &ngsiCmdError{funcName, 6, "unknown param: " + a, nil}
				}
			}
		}
	}

	if c.IsSet("expires") {
		s := c.String("expires")
		if !ngsilib.IsOrionDateTime(s) {
			var err error
			s, err = ngsilib.GetExpirationDate(s)
			if err != nil {
				return &ngsiCmdError{funcName, 7, err.Error(), nil}
			}
		}
		t.Expires = s
	}

	if c.IsSet("throttling") {
		throttling := c.Int64("throttling")
		t.Throttling = &throttling
	}

	// TemporalQ
	if isSetOR(c, []string{"timeRel", "timeAt", "endTimeAt", "timeProperty"}) {
		if t.TemporalQ == nil {
			t.TemporalQ = new(temporalQueryLd)
		}
		if c.IsSet("timeRel") {
			tr := strings.ToLower(c.String("timeRel"))
			if !ngsilib.Contains([]string{"before", "after", "bwtween"}, tr) {
				return &ngsiCmdError{funcName, 8, "unknown param: " + tr, nil}
			}
			t.TemporalQ.Timerel = tr
		}
		if c.IsSet("timeAt") {
			t.TemporalQ.TimeAt = c.String("timeAt")
		}
		if c.IsSet("endTimeAt") {
			t.TemporalQ.EndTimeAt = c.String("endTimeAt")
		}
		if c.IsSet("timeProperty") {
			t.TemporalQ.Timeproperty = c.String("timeProperty")
		}
	}

	// @context
	if c.IsSet("link") {
		link := c.String("link")
		if !ngsilib.IsHTTP(link) {
			value, err := ngsi.GetContext(link)
			if err != nil {
				return &ngsiCmdError{funcName, 9, err.Error(), err}
			}
			link = value
		}
		t.AtContext = link
	}

	return nil
}

func toLocaltimeLd(sub *subscriptionLd) {
	sub.Expires = getLocalTime(sub.Expires)
	if sub.Notification != nil {
		sub.Notification.LastNotification = getLocalTime(sub.Notification.LastNotification)
		sub.Notification.LastSuccess = getLocalTime(sub.Notification.LastSuccess)
		sub.Notification.LastFailure = getLocalTime(sub.Notification.LastFailure)
	}
}

func checkItemsLd(c *cli.Context) ([]string, error) {
	const funcName = "checkItems"

	subscriptionItems := []string{"description", "timessent", "lastnotification", "lastsuccess",
		"notificationstatus", "uri", "expires", "status"}

	items := []string{"id"}
	if c.IsSet("items") {
		list := strings.Split(c.String("items"), ",")
		for _, e := range list {
			e = strings.ToLower(e)
			if !ngsilib.Contains(subscriptionItems, e) {
				return nil, &ngsiCmdError{funcName, 1, fmt.Sprintf("error: %s in --items", e), nil}
			}
			items = append(items, e)
		}
	}
	return items, nil
}

func sprintItemsLd(e *subscriptionLd, items []string) string {
	var s []string

	for _, item := range items {
		switch item {
		case "id":
			s = append(s, e.ID)
		case "description":
			s = append(s, e.Description)
		case "timessent":
			t := "-"
			if e.Notification.TimesSent != nil {
				t = strconv.FormatInt(*(e.Notification.TimesSent), 10)
			}
			s = append(s, t)
		case "lastnotification":
			s = append(s, e.Notification.LastNotification)
		case "lastsuccess":
			s = append(s, e.Notification.LastSuccess)
		case "notificationstatus":
			c := "-"
			if e.Notification.Status != "" {
				c = e.Notification.Status
			}
			s = append(s, c)
		case "uri":
			url := "-"
			if e.Notification.Endpoint.URI != "" {
				url = e.Notification.Endpoint.URI
			}
			s = append(s, url)
		case "expires":
			s = append(s, e.Expires)
		case "status":
			s = append(s, e.Status)
		}
	}

	return strings.Join(s, " ")
}
