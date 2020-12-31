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
	"sort"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

func contextList(c *cli.Context) error {
	const funcName = "contextList"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if c.IsSet("name") {
		name := c.String("name")
		value, err := ngsi.GetContext(name)
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
		fmt.Fprint(ngsi.StdWriter, value+"\n")
	} else {
		if contexts := ngsi.GetContextList(); contexts != nil {
			keys := make([]string, len(contexts))
			i := 0
			for key := range contexts {
				keys[i] = key
				i++
			}
			sort.Strings(keys)
			for _, key := range keys {
				var s string
				switch contexts[key].(type) {
				case string:
					s = contexts[key].(string)
				case []interface{}, map[string]interface{}:
					b, err := ngsilib.JSONMarshal(contexts[key])
					if err != nil {
						return &ngsiCmdError{funcName, 3, err.Error(), err}
					}
					s = string(b)
				}
				fmt.Fprint(ngsi.StdWriter, fmt.Sprintf("%s %s\n", key, s))
			}
		}
	}

	return nil
}

func contextAdd(c *cli.Context) error {
	const funcName = "contextAdd"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if !c.IsSet("name") {
		return &ngsiCmdError{funcName, 2, "name not found", nil}
	}

	name := c.String("name")

	if ngsilib.IsNameString(name) == false {
		return &ngsiCmdError{funcName, 3, "name error " + name, nil}
	}

	if !c.IsSet("url") && !c.IsSet("json") {
		return &ngsiCmdError{funcName, 4, "url or json not provided", nil}
	}
	if c.IsSet("url") && c.IsSet("json") {
		return &ngsiCmdError{funcName, 5, "specify either url or json", nil}
	}

	var value string

	if c.IsSet("url") {
		value = c.String("url")
		if !ngsilib.IsHTTP(value) {
			return &ngsiCmdError{funcName, 6, "url error", nil}
		}
	} else if c.IsSet("json") {
		value = c.String("json")
		if !ngsi.JSONConverter.Valid([]byte(value)) {
			return &ngsiCmdError{funcName, 7, "url error", nil}
		}
	}

	if err := ngsi.AddContext(name, value); err != nil {
		return &ngsiCmdError{funcName, 8, err.Error(), err}
	}

	return nil
}

func contextUpdate(c *cli.Context) error {
	const funcName = "contextUpdate"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if !c.IsSet("name") {
		return &ngsiCmdError{funcName, 2, "name not found", nil}
	}
	name := c.String("name")

	if !c.IsSet("url") {
		return &ngsiCmdError{funcName, 3, "url not found", nil}
	}
	url := c.String("url")

	if err := ngsi.UpdateContext(name, url); err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	return nil
}

func contextDelete(c *cli.Context) error {
	const funcName = "contextDelete"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if !c.IsSet("name") {
		return &ngsiCmdError{funcName, 2, "name not found", nil}
	}
	name := c.String("name")

	if err := ngsi.DeleteContext(name); err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	return nil
}

func getAtContext(ngsi *ngsilib.NGSI, context string) (interface{}, error) {
	const funcName = "getAtContext"

	var atContext interface{}

	if ngsilib.IsNameString(context) {
		value, err := ngsi.GetContext(context)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		context = value
	}

	if ngsilib.IsJSON([]byte(context)) == false {
		if ngsilib.IsHTTP(context) {
			context = `"` + context + `"`
		} else {
			return nil, &ngsiCmdError{funcName, 2, "data not json: " + context, nil}
		}
	}

	err := ngsilib.JSONUnmarshal([]byte(context), &atContext)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	return atContext, nil
}

func insertAtContext(ngsi *ngsilib.NGSI, payload []byte, context string) (b []byte, err error) {
	const funcName = "insertAtContext"

	if !ngsilib.IsJSON(payload) {
		return nil, &ngsiCmdError{funcName, 1, "data not json", nil}
	}

	var atContext interface{}
	atContext, err = getAtContext(ngsi, context)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 2, err.Error(), nil}
	}

	if ngsilib.IsJSONArray(payload) { // JSON Array
		var entities ngsiEntities
		err = ngsilib.JSONUnmarshal(payload, &entities)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 3, err.Error(), err}
		}
		for _, e := range entities {
			e["@context"] = atContext
		}
		b, err = ngsilib.JSONMarshal(entities)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 4, err.Error(), err}
		}
	} else { // JSON Object
		var e ngsiEntity
		err = ngsilib.JSONUnmarshal(payload, &e)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		e["@context"] = atContext
		b, err = ngsilib.JSONMarshal(e)
		if err != nil {
			return nil, &ngsiCmdError{funcName, 6, err.Error(), err}
		}
	}

	return b, nil
}

type serverParam struct {
	ngsi    *ngsilib.NGSI
	context string
}

var serverGlobal *serverParam

func contextServer(c *cli.Context) error {
	const funcName = "contextServer"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	if !c.IsSet("name") && !c.IsSet("data") {
		return &ngsiCmdError{funcName, 2, "name or data  not found", nil}
	}

	if c.IsSet("name") && c.IsSet("data") {
		return &ngsiCmdError{funcName, 3, "specify either name or data", nil}
	}

	atContext := ""

	if c.IsSet("name") {
		name := c.String("name")
		name, err := ngsi.GetContext(name)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		if ngsilib.IsHTTP(name) {
			atContext = fmt.Sprintf(`["%s"]`, name)
		} else {
			atContext = name
		}
	}

	if c.IsSet("data") {
		data := c.String("data")
		if strings.HasPrefix(data, "@") {
			if len(data) > 1 {
				fileReader := ngsi.FileReader
				path, err := fileReader.FilePathAbs(data[1:])
				if err != nil {
					return &ngsiCmdError{funcName, 5, err.Error(), err}
				}
				b, err := fileReader.ReadFile(path)
				if err != nil {
					return &ngsiCmdError{funcName, 6, err.Error(), err}
				}
				data = string(b)
			} else {
				return &ngsiCmdError{funcName, 7, "file name error", nil}
			}
		}
		if ngsilib.IsHTTP(data) {
			atContext = fmt.Sprintf(`["%s"]`, data)
		} else if ngsilib.IsJSON([]byte(data)) && ngsi.JSONConverter.Valid([]byte(data)) {
			atContext = data
		} else {
			return &ngsiCmdError{funcName, 8, "data not json", nil}
		}
	}

	host := c.String("host")
	port := c.String("port")
	addr := host + ":" + port

	path := c.String("url")
	url := addr + path

	if c.Bool("https") {
		if !c.IsSet("key") {
			return &ngsiCmdError{funcName, 9, "no key file provided", nil}
		}
		if !c.IsSet("cert") {
			return &ngsiCmdError{funcName, 10, "no cert file provided", nil}
		}
		url = "https://" + url
	} else {
		url = "http://" + url
	}

	atContext = strings.TrimRight(atContext, "\n") + "\n"
	serverGlobal = &serverParam{context: atContext}

	mux := http.NewServeMux()

	mux.HandleFunc(path, http.HandlerFunc(serverHandler))

	if c.Bool("https") {
		http.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), mux)
	} else {
		http.ListenAndServe(addr, mux)
	}

	return nil
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/ld+json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(serverGlobal.context))
	}
}
