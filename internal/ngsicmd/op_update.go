/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func opUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client, actionType string) (err error) {
	const funcName = "opUpdate"

	keyValues := c.Bool("keyValues")
	safeStirng := client.IsSafeString()
	lines := false

	client.SetHeader("Content-Type", "application/json")

	fileReader, err := ngsi.GetReader(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	defer func() { _ = fileReader.Close() }()

	reader := fileReader.File()
	b, err := reader.Peek(1)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	switch string(b) {
	default:
		return ngsierr.New(funcName, 3, "data not JSON", err)
	case "{":
		lines = true
	case "[":
		lines = false
	}

	dec := json.NewDecoder(&reader)

	if !lines {
		_, _ = dec.Token()
	}

	var entities []interface{}

	for dec.More() {
		entity := make(map[string]interface{})
		err := dec.Decode(&entity)
		if err != nil {
			if err, ok := err.(*json.SyntaxError); ok {
				return ngsierr.New(funcName, 4, fmt.Sprintf("%s (%d)", err.Error(), err.Offset), err)
			}
			return ngsierr.New(funcName, 5, err.Error(), err)
		}
		entities = append(entities, entity)

		if len(entities) >= 100 {
			res, body, err := client.OpUpdate(entities, actionType, keyValues, safeStirng)
			if err != nil {
				return ngsierr.New(funcName, 6, err.Error(), err)
			}
			if res.StatusCode != http.StatusNoContent {
				return ngsierr.New(funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
			}
			entities = nil
		}
	}

	if len(entities) > 0 {
		res, body, err := client.OpUpdate(entities, actionType, keyValues, safeStirng)
		if err != nil {
			return ngsierr.New(funcName, 8, err.Error(), err)
		}
		if res.StatusCode != http.StatusNoContent {
			return ngsierr.New(funcName, 9, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
		}
	}

	if !lines {
		_, err = dec.Token()
		if err != nil {
			return ngsierr.New(funcName, 10, err.Error(), err)
		}
	}

	return nil
}
