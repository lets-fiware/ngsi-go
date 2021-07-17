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
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func copy(c *cli.Context) error {
	const funcName = "copy"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	d := ngsi.GetPreviousArgs()
	d.Host = ""
	d.Tenant = ""
	d.Scope = ""

	err = ngsi.SavePreviousArgs()
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	source, err := newClient(ngsi, c, false, []string{"broker"})
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	flags, err := parseFlags2(ngsi, c)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	var f func(c *cli.Context, ngsi *ngsilib.NGSI, source, destination *ngsilib.Client, entityType string) error

	destination, err := ngsi.NewClient(ngsi.Destination, flags, false, false)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error() + " (destination)", err}
	}

	if destination.Server.ServerType != "broker" {
		return &ngsiCmdError{funcName, 6, "destination not broker", nil}
	}

	if source.Server.ServerHost == destination.Server.ServerHost &&
		source.Tenant == destination.Tenant &&
		source.Scope == destination.Scope {
		return &ngsiCmdError{funcName, 7, "source and destination are same", nil}
	}

	if source.IsNgsiV2() && source.Scope == "" {
		source.Scope = "/"
		source.Headers["Fiware-ServicePath"] = "/"
	}
	if destination.IsNgsiV2() && destination.Scope == "" {
		destination.Scope = "/"
		destination.Headers["Fiware-ServicePath"] = "/"
	}

	if source.IsNgsiV2() && destination.IsNgsiV2() {
		if c.Bool("ngsiV1") {
			f = copyV1V1
		} else {
			f = copyV2V2
		}
	} else if source.IsNgsiV2() && destination.IsNgsiLd() {
		f = copyV2LD
	} else if source.IsNgsiLd() && destination.IsNgsiLd() {
		f = copyLDLD
	} else {
		return &ngsiCmdError{funcName, 8, "cannot copy entities from NGSI-LD to NGSI v2", err}
	}

	if !c.IsSet("type") {
		return &ngsiCmdError{funcName, 9, "specify entity type", nil}
	}

	entityType := c.String("type")

	entities := strings.Split(entityType, ",")

	for _, e := range entities {
		err = f(c, ngsi, source, destination, e)
		if err != nil {
			return &ngsiCmdError{funcName, 10, err.Error(), err}
		}
		ngsi.StdoutFlush()
	}

	return nil
}

func copyV2V2(c *cli.Context, ngsi *ngsilib.NGSI, source, destination *ngsilib.Client, entityType string) error {
	const funcName = "copyV2V2"

	page := 0
	count := 0
	limit := 100
	total := 0
	for {
		source.SetPath("/entities")

		v := url.Values{}
		v.Set("type", entityType)
		if c.Bool("skipForwarding") {
			v.Set("options", "count,skipForwarding")
		} else {
			v.Set("options", "count")
		}
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		source.SetQuery(&v)

		res, body, err := source.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}
		count, err = source.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}

		if !c.IsSet("run") {
			fmt.Fprintf(ngsi.StdWriter, "%d entities will be copied. run copy with --run option\n", count)
			return nil
		}

		if count == 0 {
			break
		}

		var entities entitiesRespose
		err = ngsilib.JSONUnmarshal(body, &entities)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}

		res, body, err = destination.OpUpdate(&entities, "append", false, false)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		if res.StatusCode != http.StatusNoContent {
			return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		total += len(entities)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	fmt.Fprintln(ngsi.StdWriter, total)

	return nil
}

func copyLDLD(c *cli.Context, ngsi *ngsilib.NGSI, source, destination *ngsilib.Client, entityType string) error {
	const funcName = "copyLDLD"

	page := 0
	limit := 100
	total := 0
	for {
		// get count
		source.SetPath("/entities")

		v := url.Values{}
		v.Set("type", entityType)
		v.Set("count", "true")
		v.Set("limit", fmt.Sprintf("%d", limit))
		source.SetQuery(&v)

		res, body, err := source.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		count, err := source.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 3, "results count error", nil}
		}

		if !c.IsSet("run") {
			fmt.Fprintf(ngsi.StdWriter, "%d entities will be copied. run copy with --run option\n", count)
			return nil
		}

		if count == 0 {
			break
		}

		destination.SetPath("/entityOperations/create")
		destination.SetContentLdJSON()

		if c.IsSet("context2") {
			body, err = insertAtContext(ngsi, body, c.String("context2"))
			if err != nil {
				return &ngsiCmdError{funcName, 4, err.Error(), err}
			}
		}
		res, body, err = destination.HTTPPost(body)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		if res.StatusCode != http.StatusCreated {
			return &ngsiCmdError{funcName, 6, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		total += count

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	fmt.Fprintln(ngsi.StdWriter, total)

	return nil
}

func copyV1V1(c *cli.Context, ngsi *ngsilib.NGSI, source, destination *ngsilib.Client, entityType string) error {
	const funcName = "copyV1V1"

	page := 0
	count := 0
	limit := 100
	total := 0
	for {

		source.SetPath("/v1/queryContext")
		payload := fmt.Sprintf("{\"entities\":[{\"type\":\"%s\",\"isPattern\":\"true\",\"id\":\".*\"}]}", entityType)

		v := url.Values{}
		v.Set("details", "on")
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		source.SetQuery(&v)
		source.SetContentJSON()

		res, body, err := source.HTTPPost([]byte(payload))
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		body, count, err = makeV1Entities(body, "APPEND")
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}

		if !c.IsSet("run") {
			fmt.Fprintf(ngsi.StdWriter, "%d entities will be copied. run copy with --run option\n", count)
			return nil
		}

		if count == 0 {
			break
		}

		destination.SetPath("/v1/updateContext")
		v = url.Values{}
		v.Set("details", "on")
		source.SetContentJSON()
		destination.SetContentJSON()

		res, body, err = destination.HTTPPost(body)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 5, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		var resBody v1Response
		err = ngsilib.JSONUnmarshal(body, &resBody)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}

		for _, e := range resBody.ContextResponses {
			if e.StatusCode.Code != "200" {
				return &ngsiCmdError{funcName, 7, fmt.Sprintf("error %s %s", e.StatusCode.Code, e.StatusCode.ReasonPhrase), err}
			}
		}
		total += len(resBody.ContextResponses)

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	fmt.Fprintln(ngsi.StdWriter, total)

	return nil
}

func makeV1Entities(body []byte, actionType string) ([]byte, int, error) {
	const funcName = "makeV1Entities"

	var res v1Response
	err := ngsilib.JSONUnmarshal(body, &res)
	if err != nil {
		return nil, 0, &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if res.ErrorCode.Code != "200" {
		return nil, 0, &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.ErrorCode, res.ErrorCode.ReasonPhrase), nil}
	}

	if !strings.HasPrefix(res.ErrorCode.Details, "Count: ") {
		return nil, 0, &ngsiCmdError{funcName, 3, "count error", nil}
	}

	count, err := strconv.Atoi(res.ErrorCode.Details[7:])
	if err != nil {
		return nil, 0, &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	var req v1Request
	for _, e := range res.ContextResponses {
		req.ContextElements = append(req.ContextElements, e.ContextElement)
	}
	req.UpdateAction = actionType

	b, err := ngsilib.JSONMarshal(req)
	if err != nil {
		return nil, 0, &ngsiCmdError{funcName, 5, err.Error(), err}
	}

	return b, count, nil
}

func copyV2LD(c *cli.Context, ngsi *ngsilib.NGSI, source, destination *ngsilib.Client, entityType string) error {
	const funcName = "copyV2LD"

	page := 0
	limit := 100
	total := 0
	for {
		// get count
		source.SetPath("/entities")

		v := url.Values{}
		v.Set("type", entityType)
		if c.Bool("skipForwarding") {
			v.Set("options", "count,skipForwarding")
		} else {
			v.Set("options", "count")
		}
		v.Set("limit", fmt.Sprintf("%d", limit))
		v.Set("offset", fmt.Sprintf("%d", page*limit))
		source.SetQuery(&v)

		res, body, err := source.HTTPGet()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if res.StatusCode != http.StatusOK {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}

		count, err := source.ResultsCount(res)
		if err != nil {
			return &ngsiCmdError{funcName, 3, "results count error", nil}
		}

		if !c.IsSet("run") {
			fmt.Fprintf(ngsi.StdWriter, "%d entities will be copied. run copy with --run option\n", count)
			return nil
		}

		if count == 0 {
			break
		}

		body, err = normalized2LD(body)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}

		destination.SetPath("/entityOperations/create")
		destination.SetContentLdJSON()

		if c.IsSet("context2") {
			body, err = insertAtContext(ngsi, body, c.String("context2"))
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
		}
		res, resBody, err := destination.HTTPPost(body)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		if res.StatusCode != http.StatusCreated {
			fmt.Fprintln(ngsi.Stderr, string(body))
			return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(resBody)), nil}
		}

		total += count

		if (page+1)*limit < count {
			page = page + 1
		} else {
			break
		}
	}

	fmt.Fprintln(ngsi.StdWriter, total)
	return nil
}

// Porting of https://github.com/FIWARE/dataModels/blob/master/tools/normalized2LD.py

type ldEntity map[string]interface{}
type ldEntities []ldEntity

func normalized2LD(body []byte) ([]byte, error) {
	const funcName = "normalized2LD"

	ldEntities := ldEntities{}
	v2Entities := entitiesRespose{}

	err := ngsilib.JSONUnmarshal(body, &v2Entities)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	for _, v2 := range v2Entities {
		ld, err := normalized2LDEntity(v2)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 2, err.Error(), err}
		}
		ldEntities = append(ldEntities, ld)
	}

	b, err := ngsilib.JSONMarshal(ldEntities)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	return b, nil
}

func normalized2LDEntity(v2 ngsiEntity) (ldEntity, error) {
	const funcName = "normalized2LDEntity"

	ld := make(ldEntity)
	if value, ok := v2["id"]; ok {
		v, ok := value.(string)
		if !ok {
			return nil, &ngsiCmdError{funcName, 1, "id not string", nil}
		}
		t, ok := v2["type"].(string)
		if !ok {
			return nil, &ngsiCmdError{funcName, 2, "type not string", nil}
		}
		ld["id"] = ldId(v, t)
	}
	for key, value := range v2 {
		switch key {
		case "id":
			continue
		case "type":
			v, ok := value.(string)
			if !ok {
				return nil, &ngsiCmdError{funcName, 3, "type not string", nil}
			}
			ld[key] = v
		default:
			attr, ok := value.(map[string]interface{})
			if !ok {
				return nil, &ngsiCmdError{funcName, 4, key + ": attribute error", nil}
			}
			v, err := ldAttribute(key, attr)
			if err != nil {
				return nil, &ngsiCmdError{funcName, 5, err.Error(), err}
			}
			ld[key] = v
		}
	}
	return ld, nil
}

type ngsiTypeValue struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func ldAttribute(key string, attr map[string]interface{}) (map[string]interface{}, error) {
	const funcName = "ldAttribute"

	ldAttr := map[string]interface{}{}

	t := ""
	if v, ok := attr["type"]; ok {
		t, ok = v.(string)
		if !ok {
			return nil, &ngsiCmdError{funcName, 1, "type not string", nil}
		}
	}

	switch t {
	case "Relationship":
		ldAttr["type"] = "Relationship"
		entityId, ok := attr["value"].(string)
		if !ok {
			return nil, &ngsiCmdError{funcName, 2, t + ": value not string", nil}
		}
		ldAttr["object"] = ldRelationship(key, entityId)
	case "DateTime":
		ldAttr["type"] = "Property"
		v, ok := attr["value"].(string)
		if !ok {
			return nil, &ngsiCmdError{funcName, 3, t + ": value not string", nil}
		}
		ldAttr["value"] = normalizeDate(v)
	case "geo:point":
		ldAttr["type"] = "GeoProperty"
		var coord string
		s, err := ngsilib.JSONMarshalEncode(attr["value"], false)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		err = ngsilib.JSONUnmarshalDecode(s, &coord, false)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		coords := []string{coord}
		v, _ := toGeoJSON(t, coords)
		ldAttr["value"] = *v
	case "geo:line", "geo:box", "geo:polygon":
		ldAttr["type"] = "GeoProperty"
		var coords []string
		s, err := ngsilib.JSONMarshalEncode(attr["value"], false)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		err = ngsilib.JSONUnmarshalDecode(s, &coords, false)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		v, _ := toGeoJSON(t, coords)
		ldAttr["value"] = *v
	case "geo:json":
		ldAttr["type"] = "GeoProperty"
		ldAttr["value"] = attr["value"]
	default:
		ldAttr["type"] = "Property"
		ldAttr["value"] = attr["value"]
	}

	if v, ok := attr["metadata"]; ok {
		var metadata map[string]ngsiTypeValue
		s, err := ngsilib.JSONMarshalEncode(v, false)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		err = ngsilib.JSONUnmarshalDecode(s, &metadata, false)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 9, err.Error(), err}
		}
		for key, value := range metadata {
			switch key {
			case "timestamp":
				ldAttr["observedAt"] = value.Value
			case "unitCode":
				ldAttr[key] = value.Value
			default:
				ldAttr[key] = ngsiTypeValue{Type: "Property", Value: value.Value}
			}
		}
	}

	return ldAttr, nil
}

type geoJSONValue struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

// transform from simple location format to GeoJSON
func toGeoJSON(t string, location []string) (*geoJSONValue, error) {
	const funcName = "toGeoJSON"

	var err error
	var coords [][2]float64

	for _, v := range location {
		points := strings.Split(strings.TrimSpace(v), ",")
		var coord [2]float64
		coord[1], err = strconv.ParseFloat(strings.TrimSpace(points[0]), 64)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		coord[0], err = strconv.ParseFloat(strings.TrimSpace(points[1]), 64)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 2, err.Error(), err}
		}
		coords = append(coords, coord)
	}

	var value geoJSONValue

	switch t {
	case "geo:point":
		value.Type = "Point"
		value.Coordinates = coords[0]
	case "geo:line":
		value.Type = "LineString"
		value.Coordinates = coords
	case "geo:polygon":
		value.Type = "Polygon"
		value.Coordinates = [...][][2]float64{coords}
	case "geo:box":
		var box [][2]float64
		box = append(box, coords[0])
		box = append(box, [2]float64{coords[1][0], coords[0][1]})
		box = append(box, coords[1])
		box = append(box, [2]float64{coords[0][0], coords[1][1]})
		box = append(box, coords[0])
		value.Type = "Polygon"
		value.Coordinates = [...][][2]float64{box}
	default:
		return nil, &ngsiCmdError{funcName, 3, "unknown type: " + t, nil}
	}
	return &value, nil
}

var (
	uriParse = regexp.MustCompile(`(?i)^(?:[^:/?#]+)`)
)

func ldRelationship(attrName string, entityId string) string {
	uri := entityId

	scheme := uriParse.FindAllStringSubmatch(entityId, -1)
	if !ngsilib.Contains([]string{"urn", "http", "https"}, scheme[0][0]) {
		entityType := attrName
		if strings.HasPrefix(attrName, "ref") {
			entityType = attrName[3:]
		}
		return ngsildUri(entityType, entityId)
	}
	return uri
}

func ngsildUri(typePart, idPart string) string {
	entityType := uriParse.FindAllStringSubmatch(idPart, -1)
	if strings.EqualFold(entityType[0][0], typePart) {
		return fmt.Sprintf("urn:ngsi-ld:%s", idPart)
	}
	return fmt.Sprintf("urn:ngsi-ld:%s:%s", typePart, idPart)
}

func ldId(entityId, entityType string) string {
	scheme := uriParse.FindAllStringSubmatch(entityId, -1)

	if !ngsilib.Contains([]string{"urn", "http", "https"}, scheme[0][0]) {
		return ngsildUri(entityType, entityId)
	}

	return entityId
}

type ngsiAtTypeValue struct {
	Type  string      `json:"@type"`
	Value interface{} `json:"@value"`
}

func normalizeDate(dateStr string) ngsiAtTypeValue {
	if !strings.HasSuffix(dateStr, "Z") {
		dateStr = dateStr + "Z"
	}
	return ngsiAtTypeValue{Type: "DateTime", Value: dateStr}
}
