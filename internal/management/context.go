/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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

package management

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func contextList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "contextList"

	if c.IsSet("name") {
		name := c.String("name")
		value, err := ngsi.GetContext(name)
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
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
						return ngsierr.New(funcName, 2, err.Error(), err)
					}
					s = string(b)
				}
				fmt.Fprintf(ngsi.StdWriter, "%s %s\n", key, s)
			}
		}
	}

	return nil
}

func contextAdd(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "contextAdd"

	name := c.String("name")

	if !ngsilib.IsNameString(name) {
		return ngsierr.New(funcName, 1, "name error "+name, nil)
	}

	var value string

	if c.IsSet("url") {
		value = c.String("url")
		if !ngsilib.IsHTTP(value) {
			return ngsierr.New(funcName, 2, "url error", nil)
		}
	} else if c.IsSet("json") {
		value = c.String("json")
		if !ngsi.JSONConverter.Valid([]byte(value)) {
			return ngsierr.New(funcName, 3, "json error", nil)
		}
	}

	if err := ngsi.AddContext(name, value); err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}

	return nil
}

func contextUpdate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "contextUpdate"

	name := c.String("name")
	url := c.String("url")

	if err := ngsi.UpdateContext(name, url); err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	return nil
}

func contextDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "contextDelete"

	name := c.String("name")

	if err := ngsi.DeleteContext(name); err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	return nil
}

func contextServer(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "contextServer"

	atContext := ""

	if c.IsSet("name") {
		name := c.String("name")
		name, err := ngsi.GetContext(name)
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
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
					return ngsierr.New(funcName, 2, err.Error(), err)
				}
				b, err := fileReader.ReadFile(path)
				if err != nil {
					return ngsierr.New(funcName, 3, err.Error(), err)
				}
				data = string(b)
			} else {
				return ngsierr.New(funcName, 4, "file name error", nil)
			}
		}
		if ngsilib.IsHTTP(data) {
			atContext = fmt.Sprintf(`["%s"]`, data)
		} else if ngsilib.IsJSON([]byte(data)) && ngsi.JSONConverter.Valid([]byte(data)) {
			atContext = data
		} else {
			return ngsierr.New(funcName, 5, "data not json", nil)
		}
	}

	host := c.String("host")
	port := c.String("port")
	addr := host + ":" + port

	path := c.String("url")

	if c.Bool("https") {
		if !c.IsSet("key") {
			return ngsierr.New(funcName, 6, "no key file provided", nil)
		}
		if !c.IsSet("cert") {
			return ngsierr.New(funcName, 7, "no cert file provided", nil)
		}
	}

	atContext = strings.TrimRight(atContext, "\n") + "\n"

	mux := http.NewServeMux()

	mux.Handle(path, &atContextServerHandler{ngsi: ngsi, context: atContext})

	if c.Bool("https") {
		_ = http.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), mux)
	} else {
		_ = http.ListenAndServe(addr, mux)
	}

	return nil
}

type atContextServerHandler struct {
	ngsi    *ngsilib.NGSI
	context string
}

func (h *atContextServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/ld+json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(h.context))
	}
}
