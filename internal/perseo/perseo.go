/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

This file is part of NGSI Go

https://github.com/lets-fiware/ngsi-go

Permission is hereby granted, free of charge, to any perseon obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit perseons to whom the Software is
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

package perseo

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type perseoRules struct {
	Error interface{}  `json:"error"`
	Data  []perseoRule `json:"data"`
	Count int          `json:"count"`
}

type perseoRule struct {
	ID         string      `json:"_id"`
	Name       string      `json:"name"`
	Text       string      `json:"text"`
	Action     interface{} `json:"action"`
	Subservice string      `json:"subservice"`
	Service    string      `json:"service"`
}

func perseoRulesList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "perseoRulesList"

	v := ngsicli.ParseOptions(c, []string{"limit", "offset"}, nil)
	client.SetQuery(v)

	client.SetPath("/rules")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	if err = perseoPrintRespose(c, ngsi, body); err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}
	return nil
}

func perseoPrintRespose(c *ngsicli.Context, ngsi *ngsilib.NGSI, body []byte) error {
	const funcName = "perseoPrintRespose"

	if c.Bool("raw") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 1, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(body))
		}
	} else {
		var rules perseoRules
		err := ngsilib.JSONUnmarshal(body, &rules)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if c.Bool("count") {
			fmt.Fprintln(ngsi.StdWriter, rules.Count)
		} else {
			if c.Bool("verbose") || c.Bool("pretty") {
				b, err := ngsilib.JSONMarshal(rules.Data)
				if err != nil {
					return ngsierr.New(funcName, 3, err.Error(), err)
				}
				if c.Bool("pretty") {
					newBuf := new(bytes.Buffer)
					err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
					if err != nil {
						return ngsierr.New(funcName, 4, err.Error(), err)
					}
					fmt.Fprintln(ngsi.StdWriter, newBuf.String())
				} else {
					fmt.Fprintln(ngsi.StdWriter, string(b))
				}
			} else {
				for _, rule := range rules.Data {
					fmt.Fprintln(ngsi.StdWriter, rule.Name)
				}
			}
		}
	}

	return nil
}

func perseoRulesGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "perseoRulesGet"

	client.SetPath("/rules/" + c.String("name"))

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprintln(ngsi.StdWriter, string(body))

	return nil
}

func perseoRulesCreate(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "perseoRulesCreate"

	b, err := ngsi.ReadAll(c.String("data"))
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/rules")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 3, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 4, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(body))
		}
	}

	return nil
}

func perseoRulesDelete(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "perseoRulesDelete"

	client.SetPath("/rules/" + c.String("name"))

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("%s %s", res.Status, string(body)), nil)
	}

	return nil
}
