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
	"sort"

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
				fmt.Fprint(ngsi.StdWriter, fmt.Sprintf("%s %s\n", key, contexts[key]))
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
	if !c.IsSet("url") {
		return &ngsiCmdError{funcName, 4, "url not found", nil}
	}
	url := c.String("url")

	if err := ngsi.AddContext(name, url); err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
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
