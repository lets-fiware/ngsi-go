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

package ngsicmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

// types respose
type typesResponseV2 struct {
	Attrs map[string]interface{} `json:"attrs"`
	Count int                    `json:"count"`
}

type subscriptionEntityV2 struct {
	ID          string `json:"id,omitempty"`
	IDPattern   string `json:"idPattern,omitempty"`
	Type        string `json:"type,omitempty"`
	TypePattern string `json:"typePattern,omitempty"`
}

type subscriptionHTTPV2 struct {
	URL string `json:"url,omitempty"`
}
type subscriptionHTTPCustomV2 struct {
	URL     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Qs      map[string]string `json:"qs,omitempty"`
	Method  string            `json:"method,omitempty"`
	Payload string            `json:"payload,omitempty"`
}

type subscriptionExpressionV2 struct {
	Q        string `json:"q,omitempty"`
	Mq       string `json:"mq,omitempty"`
	Georel   string `json:"georel,omitempty"`
	Geometry string `json:"geometry,omitempty"`
	Coords   string `json:"coords,omitempty"`
}

type subscriptionConditionV2 struct {
	Attrs      []string                  `json:"attrs,omitempty"`
	Expression *subscriptionExpressionV2 `json:"expression,omitempty"`
}

type subscriptionSubjectV2 struct {
	Entities  []subscriptionEntityV2   `json:"entities,omitempty"`
	Condition *subscriptionConditionV2 `json:"condition,omitempty"`
}

type subscriptionNotificationV2 struct {
	HTTP        *subscriptionHTTPV2       `json:"http,omitempty"`
	HTTPCustom  *subscriptionHTTPCustomV2 `json:"httpCustom,omitempty"`
	Attrs       []string                  `json:"attrs,omitempty"`
	Metadata    []string                  `json:"metadata,omitempty"`
	ExceptAttrs []string                  `json:"exceptAttrs,omitempty"`
	AttrsFormat string                    `json:"attrsFormat,omitempty"`
}

type subscriptionV2 struct {
	Description  string                      `json:"description,omitempty"`
	Subject      *subscriptionSubjectV2      `json:"subject,omitempty"`
	Notification *subscriptionNotificationV2 `json:"notification,omitempty"`
	Throttling   int64                       `json:"throttling,omitempty"`
	Expires      string                      `json:"expires,omitempty"`
	Status       string                      `json:"status,omitempty"`
}

type subscriptionResposeV2 struct {
	ID          string `json:"id"`
	Description string `json:"description,omitempty"`
	Subject     struct {
		Entities  []subscriptionEntityV2   `json:"entities"`
		Condition *subscriptionConditionV2 `json:"condition,omitempty"`
	} `json:"subject"`
	Notification struct {
		TimesSent         *int64                    `json:"timesSent,omitempty"`
		LastNotification  string                    `json:"lastNotification,omitempty"`
		LastSuccess       string                    `json:"lastSuccess,omitempty"`
		LastSuccessCode   *int                      `json:"lastSuccessCode,omitempty"`
		LastFailure       string                    `json:"lastFailure,omitempty"`
		LastFailureReason string                    `json:"lastFailureReason,omitempty"`
		OnlyChangedAttrs  *bool                     `json:"onlyChangedAttrs,omitempty"`
		HTTP              *subscriptionHTTPV2       `json:"http,omitempty"`
		HTTPCustom        *subscriptionHTTPCustomV2 `json:"httpCustom,omitempty"`
		Attrs             []string                  `json:"attrs,omitempty"`
		Metadata          []string                  `json:"metadata,omitempty"`
		ExceptAttrs       []string                  `json:"exceptAttrs,omitempty"`
		AttrsFormat       string                    `json:"attrsFormat,omitempty"`
	} `json:"notification"`
	Throttling int64  `json:"throttling,omitempty"`
	Expires    string `json:"expires,omitempty"`
	Status     string `json:"status,omitempty"`
}

func subscriptionsListV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsListV2"

	filters := []string{}

	if c.IsSet("status") {
		status := []string{"active", "inactive", "oneshot", "expired", "failed"}
		filters = strings.Split(c.String("status"), ",")
		for i, v := range filters {
			filters[i] = strings.ToLower(v)
			if !ngsilib.Contains(status, filters[i]) {
				return ngsierr.New(funcName, 1, "error: "+filters[i]+" (active, inactive, oneshot, expired, failed)", nil)
			}
		}
	}

	page := 0
	count := 0
	limit := 100
	total := 0

	var subscriptions []subscriptionResposeV2

	for {
		client.SetPath("/subscriptions")

		v := url.Values{}
		v.Set("options", "count")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		client.SetQuery(&v)

		res, body, err := client.HTTPGet()
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
		}
		count, err = client.ResultsCount(res)
		if err != nil {
			return ngsierr.New(funcName, 4, "ResultsCount error", err)
		}
		if count == 0 {
			break
		}
		var subs []subscriptionResposeV2
		if err := ngsilib.JSONUnmarshalDecode(body, &subs, client.IsSafeString()); err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
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
		var subs []subscriptionResposeV2
		for _, e := range subscriptions {
			if ngsilib.Contains(filters, e.Status) {
				subs = append(subs, e)
			}
		}
		subscriptions = subs
	}

	if c.IsSet("query") {
		reg := regexp.MustCompile(c.String("query"))

		var subs []subscriptionResposeV2
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
				return ngsierr.New(funcName, 6, err.Error(), err)
			}
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
				if err != nil {
					return ngsierr.New(funcName, 7, err.Error(), err)
				}
				fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			} else {
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		}
	} else if c.IsSet("verbose") {
		items := []string{"id", "status", "expires", "description"}
		var err error
		if c.IsSet("items") {
			items, err = checkItems(c)
			if err != nil {
				return ngsierr.New(funcName, 8, err.Error(), err)
			}
		}
		local := c.IsSet("localTime")
		for _, e := range subscriptions {
			if local {
				toLocaltime(&e)
			}
			fmt.Fprintln(ngsi.StdWriter, sprintItems(&e, items))
		}
	} else {
		for _, e := range subscriptions {
			fmt.Fprintln(ngsi.StdWriter, e.ID)
		}
	}

	return nil
}

func subscriptionGetV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionGetV2"

	id := c.String("id")
	client.SetPath("/subscriptions/" + id)

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s %s", res.Status, string(body), id), nil)
	}

	b := body

	if !c.Bool("raw") {
		var sub subscriptionResposeV2
		if err := ngsilib.JSONUnmarshalDecode(body, &sub, client.IsSafeString()); err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}

		if c.IsSet("localTime") {
			toLocaltime(&sub)
		}
		b, err = ngsilib.JSONMarshal(&sub)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(b))
	return nil
}

func subscriptionsCreateV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsCreateV2"

	var b []byte
	var err error

	client.SetPath("/subscriptions")

	opts := []string{"skipInitialNotification"}
	v := ngsicli.ParseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetHeader("Content-Type", "application/json")

	if c.Bool("raw") && c.IsSet("data") {
		b, err = ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
	} else {

		var t subscriptionV2

		if err := setSubscriptionValuesV2(c, ngsi, &t, false); err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}

		b, err = ngsilib.JSONMarshalEncode(&t, true)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated {
		return ngsierr.New(funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	location := res.Header.Get("Location")
	p := "/v2/subscriptions/"
	location = strings.TrimPrefix(location, p)

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is created, FIWARE-Service: %s, FIWARE-ServicePath: %s",
		res.Header.Get("Location"), c.String("service"), c.String("path")))

	fmt.Fprintln(ngsi.StdWriter, location)
	return nil
}

func subscriptionsUpdateV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsUpdateV2"

	var b []byte
	var err error

	id := c.String("id")

	client.SetPath("/subscriptions/" + id)

	opts := []string{"skipInitialNotification"}
	v := ngsicli.ParseOptions(c, nil, opts)
	client.SetQuery(v)

	client.SetHeader("Content-Type", "application/json")

	if c.Bool("raw") && c.IsSet("data") {
		b, err = ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
	} else {
		var t subscriptionV2

		if err := setSubscriptionValuesV2(c, ngsi, &t, true); err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}

		b, err = ngsilib.JSONMarshalEncode(&t, true)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 5, fmt.Sprintf("%s %s %s", res.Status, string(body), id), nil)
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is updated, FIWARE-Service: %s, FIWARE-ServicePath: %s",
		id, c.String("service"), c.String("path")))

	fmt.Fprintln(ngsi.StdWriter, id)

	return nil
}

func subscriptionsDeleteV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "subscriptionsDeleteV2"

	id := c.String("id")
	path := "/subscriptions/" + id
	client.SetPath(path)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s %s", res.Status, string(body), id), nil)
	}

	ngsi.Logging(ngsilib.LogInfo, fmt.Sprintf("%s is deleted, FIWARE-Service: %s, FIWARE-ServicePath: %s",
		path, c.String("service"), c.String("path")))

	return nil
}

func subscriptionsTemplateV2(c *ngsicli.Context, ngsi *ngsilib.NGSI) error {
	const funcName = "subscriptionsTemplateV2"

	var t subscriptionV2

	if err := setSubscriptionValuesV2(c, ngsi, &t, false); err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	b, err := ngsilib.JSONMarshal(t)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(b))
	return nil
}

func setSubscriptionValuesV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, t *subscriptionV2, checkskip bool) error {
	const funcName = "setSubscriptionValuesV2"

	if c.IsSet("data") {
		b, err := ngsi.ReadAll(c.String("data"))
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		err = ngsilib.JSONUnmarshal(b, t)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}
	err := getAttributesV2(c, ngsi, t)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	if c.IsSet("entityId") || c.IsSet("idPattern") {
		if t.Subject == nil {
			t.Subject = new(subscriptionSubjectV2)
		}
		if len(t.Subject.Entities) == 0 {
			t.Subject.Entities = append(t.Subject.Entities, *new(subscriptionEntityV2))
		}
		if c.IsSet("entityId") {
			t.Subject.Entities[0].ID = c.String("entityId")
			t.Subject.Entities[0].IDPattern = ""
		}
		if c.IsSet("idPattern") {
			t.Subject.Entities[0].IDPattern = c.String("idPattern")
			t.Subject.Entities[0].ID = ""
		}
	}

	if c.IsSet("type") && c.IsSet("typePattern") {
		return ngsierr.New(funcName, 4, "type or typePattern", nil)
	}

	if c.IsSet("type") || c.IsSet("typePattern") {
		if t.Subject == nil {
			t.Subject = new(subscriptionSubjectV2)
		}
		if len(t.Subject.Entities) == 0 {
			t.Subject.Entities = append(t.Subject.Entities, *new(subscriptionEntityV2))
		}
		if c.IsSet("type") {
			t.Subject.Entities[0].Type = c.String("type")
			t.Subject.Entities[0].TypePattern = ""
		}
		if c.IsSet("typePattern") {
			t.Subject.Entities[0].TypePattern = c.String("typePattern")
			t.Subject.Entities[0].Type = ""
		}
	}

	if c.IsSetOR([]string{"wAttrs", "query", "mq", "georel", "geometry", "coords"}) {
		if t.Subject == nil {
			t.Subject = new(subscriptionSubjectV2)
		}
		if t.Subject.Condition == nil {
			t.Subject.Condition = new(subscriptionConditionV2)
		}
		if c.IsSet("wAttrs") {
			t.Subject.Condition.Attrs = strings.Split(c.String("wAttrs"), ",")
		}

		if c.IsSetOR([]string{"query", "mq", "georel", "geometry", "coords"}) {
			if t.Subject.Condition.Expression == nil {
				t.Subject.Condition.Expression = new(subscriptionExpressionV2)
			}
			if c.IsSet("query") {
				t.Subject.Condition.Expression.Q = c.String("query")
			}
			if c.IsSet("mq") {
				t.Subject.Condition.Expression.Mq = c.String("mq")
			}
			if c.IsSet("georel") {
				t.Subject.Condition.Expression.Georel = c.String("georel")
			}
			if c.IsSet("geometry") {
				t.Subject.Condition.Expression.Geometry = c.String("geometry")
			}
			if c.IsSet("coords") {
				t.Subject.Condition.Expression.Coords = c.String("coords")
			}
		}
	}

	if c.IsSetOR([]string{"uri", "url", "headers", "qs", "method", "payload"}) {
		url := ""
		if c.IsSet("uri") {
			url = c.String("uri")
		} else if c.IsSet("url") {
			url = c.String("url")
		}
		if url != "" {
			if t.Notification == nil {
				t.Notification = new(subscriptionNotificationV2)
			}
			if t.Notification.HTTP == nil {
				t.Notification.HTTP = new(subscriptionHTTPV2)
			}
			t.Notification.HTTP.URL = url
		}
		if c.IsSetOR([]string{"headers", "qs", "method", "payload"}) {
			if t.Notification.HTTPCustom == nil {
				t.Notification.HTTPCustom = new(subscriptionHTTPCustomV2)
			}
			if t.Notification.HTTP != nil {
				t.Notification.HTTPCustom.URL = t.Notification.HTTP.URL
				t.Notification.HTTP = nil
			}
			if c.IsSet("headers") {
				var headers map[string]string
				err := ngsilib.JSONUnmarshal([]byte(c.String("headers")), &headers)
				if err != nil {
					return ngsierr.New(funcName, 5, "err"+c.String("headers"), err)
				}
				t.Notification.HTTPCustom.Headers = headers
			}
			if c.IsSet("qs") {
				var qs map[string]string
				err := ngsilib.JSONUnmarshal([]byte(c.String("qs")), &qs)
				if err != nil {
					return ngsierr.New(funcName, 6, "err"+c.String("qs"), err)
				}
				t.Notification.HTTPCustom.Qs = qs
			}
			if c.IsSet("method") {
				t.Notification.HTTPCustom.Method = c.String("method")
			}
			if c.IsSet("payload") {
				s := c.String("payload")
				s = strings.ReplaceAll(s, `\"`, `"`)
				s = strings.ReplaceAll(s, `\\`, `\`)
				s = ngsilib.SafeStringEncode(s)
				t.Notification.HTTPCustom.Payload = s
			}
		}
	}

	if c.IsSetOR([]string{"nAttrs", "metadata", "exceptAttrs", "attrsFormat"}) {
		if t.Notification == nil {
			t.Notification = new(subscriptionNotificationV2)
		}
		if c.IsSet("nAttrs") {
			t.Notification.Attrs = strings.Split(c.String("nAttrs"), ",")
		}
		if c.IsSet("metadata") {
			t.Notification.Metadata = strings.Split(c.String("metadata"), ",")
		}
		if c.IsSet("exceptAttrs") && c.IsSet("nAttrs") {
			return ngsierr.New(funcName, 7, "error exceptAttrs or nAttrs", nil)
		}
		if c.IsSet("exceptAttrs") {
			t.Notification.ExceptAttrs = strings.Split(c.String("exceptAttrs"), ",")
		}
		if c.IsSet("attrsFormat") {
			t.Notification.AttrsFormat = c.String("attrsFormat")
		}
	}

	if c.IsSet("throttling") {
		t.Throttling = c.Int64("throttling")
	}
	if c.IsSet("expires") {
		s := c.String("expires")
		if !ngsilib.IsOrionDateTime(s) {
			s, err = ngsilib.GetExpirationDate(s)
			if err != nil {
				return ngsierr.New(funcName, 8, err.Error(), err)
			}
		}
		t.Expires = s
	}
	if c.IsSet("status") {
		s := strings.ToLower(c.String("status"))
		if !ngsilib.Contains([]string{"active", "inactive", "oneshot"}, s) {
			return ngsierr.New(funcName, 9, "error: "+s+" (active, inactive, oneshot)", nil)
		}
		t.Status = s
	}
	return nil
}

func getAttributesV2(c *ngsicli.Context, ngsi *ngsilib.NGSI, t *subscriptionV2) error {
	const funcName = "getAttributesV2"

	entityType := c.String("type")

	if entityType != "" && c.IsSet("host") && c.IsSet("get") {
		ngsi.Host = c.String("host")
		client, err := ngsicli.NewClient(ngsi, c, false, []string{"brokerv2"})
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}

		client.SetPath("/types/" + entityType)

		res, body, err := client.HTTPGet()
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if res.StatusCode != http.StatusOK {
			return ngsierr.New(funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
		}

		var typeInfo typesResponseV2

		if err := ngsilib.JSONUnmarshal(body, &typeInfo); err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}

		attrs := make([]string, len(typeInfo.Attrs))
		var i = 0
		for key := range typeInfo.Attrs {
			attrs[i] = key
			i++
		}

		if t.Subject == nil {
			t.Subject = new(subscriptionSubjectV2)
		}
		if t.Subject.Condition == nil {
			t.Subject.Condition = new(subscriptionConditionV2)
		}
		t.Subject.Condition.Attrs = attrs

		if t.Notification == nil {
			t.Notification = new(subscriptionNotificationV2)
		}
		t.Notification.Attrs = attrs

		if len(t.Subject.Entities) == 0 {
			t.Subject.Entities = append(t.Subject.Entities, *new(subscriptionEntityV2))
		}
		t.Subject.Entities[0].Type = entityType
	}

	return nil
}

func toLocaltime(sub *subscriptionResposeV2) {
	sub.Expires = getLocalTime(sub.Expires)
	sub.Notification.LastNotification = getLocalTime(sub.Notification.LastNotification)
	sub.Notification.LastSuccess = getLocalTime(sub.Notification.LastSuccess)
	sub.Notification.LastFailure = getLocalTime(sub.Notification.LastFailure)
}

func getLocalTime(s string) string {
	if s == "" {
		return s
	}
	p := strings.Index(s, "Z") - strings.Index(s, ".")
	if p == 3 {
		layout := "2006-01-02T15:04:05.00Z"
		t, _ := time.Parse(layout, s)
		t = t.In(time.Local)
		layoutLocal := "2006-01-02T15:04:05.00"
		z := t.Format(time.RFC1123Z)
		return t.Format(layoutLocal) + z[len(z)-5:]
	} else if p == 4 {
		layout := "2006-01-02T15:04:05.000Z"
		t, _ := time.Parse(layout, s)
		t = t.In(time.Local)
		layoutLocal := "2006-01-02T15:04:05.000"
		z := t.Format(time.RFC1123Z)
		return t.Format(layoutLocal) + z[len(z)-5:]
	}

	return s
}

func checkItems(c *ngsicli.Context) ([]string, error) {
	const funcName = "checkItems"

	subscriptionItems := []string{"description", "timessent", "lastnotification", "lastsuccess",
		"lastsuccesscode", "url", "expires", "status"}

	items := []string{"id"}
	if c.IsSet("items") {
		list := strings.Split(c.String("items"), ",")
		for _, e := range list {
			e = strings.ToLower(e)
			if !ngsilib.Contains(subscriptionItems, e) {
				return nil, ngsierr.New(funcName, 1, fmt.Sprintf("error: %s in --items", e), nil)
			}
			items = append(items, e)
		}
	}
	return items, nil
}

func sprintItems(e *subscriptionResposeV2, items []string) string {
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
		case "lastsuccesscode":
			c := "-"
			if e.Notification.LastSuccessCode != nil {
				c = strconv.Itoa(*(e.Notification.LastSuccessCode))
			}
			s = append(s, c)
		case "url":
			url := "-"
			if e.Notification.HTTP != nil {
				url = e.Notification.HTTP.URL
			} else if e.Notification.HTTPCustom != nil {
				url = e.Notification.HTTPCustom.URL
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
