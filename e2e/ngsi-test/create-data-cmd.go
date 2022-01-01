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

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type v1Metadata struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type v1Attribute struct {
	Name      string       `json:"name,omitempty"`
	Type      string       `json:"type,omitempty"`
	Value     interface{}  `json:"value,omitempty"`
	Metadatas []v1Metadata `json:"metadatas,omitempty"`
}

type contextResponse struct {
	ContextElement struct {
		Type       string        `json:"type,omitempty"`
		IsPattern  string        `json:"isPattern,omitempty"`
		ID         string        `json:"id,omitempty"`
		Attributes []v1Attribute `json:"attributes,omitempty"`
	} `json:"contextElement,omitempty"`
	StatusCode struct {
		Code         string `json:"code,omitempty"`
		ReasonPhrase string `json:"reasonPhrase,omitempty"`
	} `json:"statusCode,omitempty"`
}

type v1Notify struct {
	SubscriptionID   string            `json:"subscriptionId,omitempty"`
	Originator       string            `json:"originator,omitempty"`
	ContextResponses []contextResponse `json:"contextResponses,omitempty"`
}

type v2Metadata struct {
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type v2Attribute struct {
	Type     string                `json:"type,omitempty"`
	Value    int                   `json:"value,omitempty"`
	Metadata map[string]v2Metadata `json:"metadata,omitempty"`
}
type v2Entity map[string]interface{}

type v2Notify struct {
	SubscriptionID string     `json:"subscriptionId,omitempty"`
	Data           []v2Entity `json:"data,omitempty"`
}

func createdataCmd(line int, args []string) error {
	const funcName = "createdataCmd"

	if len(args) > 3 {
		switch args[1] {
		default:
			return &ngsiCmdError{funcName, 2, "command error: " + args[0], nil}
		case "v1notify":
			return createV1NotifyData(args[2:])
		case "v2notify":
			return createV2NotifyData(args[2:])
		}

	}
	return &ngsiCmdError{funcName, 1, "no args error", nil}
}

func createV1NotifyData(args []string) error {
	const funcName = "createV1NotifyData"

	if len(args) == 0 {
		return &ngsiCmdError{funcName, 1, "no args error", nil}
	}

	opts, err := getOpts(args, []string{"url", "service", "path", "id", "type", "datetime", "count", "subsId", "attrs", "values", "period"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), nil}
	}

	if err = checkRequiredOpt(opts, []string{"url", "id", "datetime", "attrs", "values", "count"}); err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), nil}
	}

	url := opts["url"]
	if !isHTTP(url) {
		return &ngsiCmdError{funcName, 4, url + " not url", nil}
	}

	header := map[string]string{}
	header["Content-type"] = "application/json"
	for _, k := range []string{"service", "path"} {
		if v, ok := opts[k]; ok {
			if k == "service" {
				k = "Fiware-service"
			} else if k == "path" {
				k = "Fiware-servicepath"
			}
			header[k] = v
		}
	}

	subsID, ok := opts["subsId"]
	if !ok {
		subsID = "000000000000000000000001"
	}
	entityType, ok := opts["type"]
	if !ok {
		entityType = "Thing"
	}
	period, ok := opts["period"]
	if !ok {
		period = "minute"
	}
	if !contains([]string{"month", "day", "hour", "minute"}, period) {
		return &ngsiCmdError{funcName, 5, "period (month, day, hour, minute): " + period, nil}
	}

	dateTime := opts["datetime"]
	layout := "2006-01-02T15:04:05.000Z"
	dt, err := time.Parse(layout, dateTime)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), nil}
	}

	count, err := strconv.Atoi(opts["count"])
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), nil}
	}

	keys := strings.Split(opts["attrs"], ",")
	values := strings.Split(opts["values"], ",")

	if len(keys) != len(values) {
		return &ngsiCmdError{funcName, 8, "attrs, values error", nil}
	}

	attrs := map[string]int{}
	for i, k := range keys {
		v, err := strconv.Atoi(values[i])
		if err != nil {
			return &ngsiCmdError{funcName, 9, err.Error(), err}
		}
		attrs[k] = v
	}

	progress := -1
	for i := 0; i < count; i++ {
		if progress != (i*100)/count {
			progress = (i * 100) / count
			fmt.Printf("[%3d / 100]\033[0G", progress)
		}
		cr := contextResponse{}
		cr.ContextElement.Type = entityType
		cr.ContextElement.IsPattern = "false"
		cr.ContextElement.ID = opts["id"]
		cr.ContextElement.Attributes = []v1Attribute{}
		cr.StatusCode.Code = "200"
		cr.StatusCode.ReasonPhrase = "OK"

		for k, v := range attrs {
			attr := v1Attribute{}
			attr.Name = k
			attr.Type = "Number"
			attr.Value = v
			attr.Metadatas = []v1Metadata{}
			metadata := v1Metadata{Name: "TimeInstant", Type: "DateTime", Value: dt.Format(layout)}
			attr.Metadatas = append(attr.Metadatas, metadata)
			cr.ContextElement.Attributes = append(cr.ContextElement.Attributes, attr)
			attrs[k] = v + 1
			dt = incrementTime(period, dt)
		}

		notify := &v1Notify{SubscriptionID: subsID, Originator: "localhost"}
		notify.ContextResponses = []contextResponse{}
		notify.ContextResponses = append(notify.ContextResponses, cr)

		b, err := json.Marshal(notify)
		if err != nil {

			return &ngsiCmdError{funcName, 10, err.Error(), nil}
		}

		err = httpRequest("POST", url, header, b)
		if err != nil {
			return &ngsiCmdError{funcName, 11, err.Error(), nil}
		}
	}
	fmt.Println("[100 / 100]")

	return nil
}

type testEntity struct {
	ID   string         `json:"id"`
	Attr map[string]int `json:"attr"`
}
type testEntities []testEntity
type testData map[string]testEntities

func createV2NotifyData(args []string) error {
	const funcName = "createV2NotifyData"

	if len(args) == 0 {
		return &ngsiCmdError{funcName, 1, "no args error", nil}
	}

	opts, err := getOpts(args, []string{"url", "service", "path", "datetime", "count", "subsId", "period", "data", "wait"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), nil}
	}

	if err = checkRequiredOpt(opts, []string{"url", "datetime", "count", "data"}); err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), nil}
	}

	data := testData{}
	if err = json.Unmarshal([]byte(opts["data"]), &data); err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), nil}
	}

	url := opts["url"]
	if !isHTTP(url) {
		return &ngsiCmdError{funcName, 3, url + " not url", nil}
	}

	header := map[string]string{}
	header["Content-type"] = "application/json"
	for _, k := range []string{"service", "path"} {
		if v, ok := opts[k]; ok {
			if k == "service" {
				k = "Fiware-service"
			} else if k == "path" {
				k = "Fiware-servicepath"
			}
			header[k] = v
		}
	}

	wait := 0
	if w, ok := opts["wait"]; ok {
		wait, err = strconv.Atoi(w)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), nil}
		}
	}

	subsID, ok := opts["subsId"]
	if !ok {
		subsID = "000000000000000000000001"
	}
	period, ok := opts["period"]
	if !ok {
		period = "minute"
	}
	if !contains([]string{"month", "day", "hour", "minute"}, period) {
		return &ngsiCmdError{funcName, 4, "period (month, day, hour, minute): " + period, nil}
	}

	dateTime := opts["datetime"]
	layout := "2006-01-02T15:04:05.000Z"
	dt, err := time.Parse(layout, dateTime)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), nil}
	}

	count, err := strconv.Atoi(opts["count"])
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), nil}
	}

	progress := -1
	for i := 0; i < count; i++ {
		if progress != (i*100)/count {
			progress = (i * 100) / count
			fmt.Printf("[%3d / 100]\033[0G", progress)
		}
		notify := &v2Notify{SubscriptionID: subsID}
		notify.Data = []v2Entity{}

		for entityType, entities := range data {
			for _, e := range entities {
				entity := v2Entity{}
				entity["type"] = entityType
				entity["id"] = e.ID
				for k, v := range e.Attr {
					attr := v2Attribute{}
					attr.Type = "Number"
					attr.Value = v
					attr.Metadata = map[string]v2Metadata{}
					attr.Metadata["dateCreated"] = v2Metadata{Type: "DateTime", Value: dateTime}
					attr.Metadata["dateModified"] = v2Metadata{Type: "DateTime", Value: dt.Format(layout)}
					entity[k] = attr
					e.Attr[k] = v + 1
				}
				notify.Data = append(notify.Data, entity)
			}
		}
		dt = incrementTime(period, dt)

		b, err := json.Marshal(notify)
		if err != nil {

			return &ngsiCmdError{funcName, 9, err.Error(), nil}
		}

		err = httpRequest("POST", url, header, b)
		if err != nil {
			return &ngsiCmdError{funcName, 10, err.Error(), nil}
		}

		if wait > 0 {
			time.Sleep(time.Millisecond * time.Duration(wait))
		}
	}
	fmt.Println("[100 / 100]")

	return nil
}

func incrementTime(period string, t time.Time) time.Time {
	switch period {
	default: // "minute":
		t = t.Add(time.Duration(1) * 60 * time.Second)
	case "month":
		t = t.AddDate(0, 1, 0)
	case "day":
		t = t.AddDate(0, 0, 1)
	case "hour":
		t = t.Add(time.Duration(1) * 3600 * time.Second)
	}
	return t
}
