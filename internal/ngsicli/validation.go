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

package ngsicli

import (
	"fmt"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

const (
	NonCondition = iota
	NandCondition
	XnorCondition
)

type ValidationFlag struct {
	Mode  int
	Flags []string
}

func validation(f *ValidationFlag, c *Context) error {
	const funcName = "validation"

	if f == nil {
		return nil
	}

	switch f.Mode {
	default:
		return ngsierr.New(funcName, 1, "validation mode error", nil)
	case NandCondition:
		b := c.IsSet(f.Flags[0]) && c.IsSet(f.Flags[1])
		if b {
			return ngsierr.New(funcName, 2, fmt.Sprintf("specify either --%s or --%s", f.Flags[0], f.Flags[1]), nil)
		}
	case XnorCondition:
		b := c.IsSet(f.Flags[0]) == c.IsSet(f.Flags[1])
		if b {
			return ngsierr.New(funcName, 3, fmt.Sprintf("specify either --%s or --%s", f.Flags[0], f.Flags[1]), nil)
		}
	}

	return nil
}
