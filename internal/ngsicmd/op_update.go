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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

var actionTypes = []string{"append", "append_strict", "update", "delete", "replace"}

func opUpdate(c *cli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client, actionType string) error {
	const funcName = "opUpdate"

	keyValues := c.Bool("keyValues")
	safeStirng := client.IsSafeString()
	lines := false

	client.SetHeader("Content-Type", "application/json")

	fileReader, err := getReader(c, ngsi)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	defer fileReader.Close()
	reader := fileReader.File()
	dec := json.NewDecoder(reader)

	t, err := dec.Token()
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	switch t.(type) {
	default:
		return &ngsiCmdError{funcName, 3, "data is not JSON", nil}
	case json.Delim:
		if t == json.Delim('{') {
			fileReader.Close()
			fileReader, err = getReader(c, ngsi)
			reader = fileReader.File()
			dec = json.NewDecoder(reader)
			lines = true
		}
	}
	var entities []interface{}

	for dec.More() {
		entity := make(map[string]interface{})
		err := dec.Decode(&entity)
		if err != nil {
			if err, ok := err.(*json.SyntaxError); ok {
				return &ngsiCmdError{funcName, 4, fmt.Sprintf("%s (%d)", err.Error(), err.Offset), err}
			}
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		entities = append(entities, entity)

		if len(entities) >= 100 {
			res, body, err := client.OpUpdate(entities, actionType, keyValues, safeStirng)
			if err != nil {
				return &ngsiCmdError{funcName, 6, err.Error(), err}
			}
			if res.StatusCode != http.StatusNoContent {
				return &ngsiCmdError{funcName, 7, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
			}
			entities = nil
		}
	}

	if len(entities) > 0 {
		res, body, err := client.OpUpdate(entities, actionType, keyValues, safeStirng)
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		if res.StatusCode != http.StatusNoContent {
			return &ngsiCmdError{funcName, 9, fmt.Sprintf("%s %s", res.Status, string(body)), nil}
		}
	}

	if !lines {
		_, err = dec.Token()
		if err != nil {
			return &ngsiCmdError{funcName, 10, err.Error(), err}
		}
	}

	return nil
}
