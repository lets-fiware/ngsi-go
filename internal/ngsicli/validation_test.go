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
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestValidationNandCondition(t *testing.T) {
	c := &Context{
		Ngsi: &ngsilib.NGSI{},
		Flags: []Flag{
			&StringFlag{Name: "flag1", Value: "", Set: true},
			&StringFlag{Name: "flag2", Value: "", Set: false},
		},
	}

	v := &ValidationFlag{Mode: NandCondition, Flags: []string{"flag1", "flag2"}}

	err := validation(v, c)

	assert.NoError(t, err)
}

func TestValidationXnorCondition(t *testing.T) {
	c := &Context{
		Ngsi: &ngsilib.NGSI{},
		Flags: []Flag{
			&StringFlag{Name: "flag1", Value: "", Set: true},
			&StringFlag{Name: "flag2", Value: "", Set: false},
		},
	}

	v := &ValidationFlag{Mode: XnorCondition, Flags: []string{"flag1", "flag2"}}

	err := validation(v, c)

	assert.NoError(t, err)
}

func TestValidationErrorNonCondition(t *testing.T) {
	c := &Context{
		Ngsi: &ngsilib.NGSI{},
		Flags: []Flag{
			&StringFlag{Name: "flag1", Value: "", Set: true},
			&StringFlag{Name: "flag2", Value: "", Set: true},
		},
	}

	v := &ValidationFlag{Mode: NonCondition, Flags: []string{"flag1", "flag2"}}

	err := validation(v, c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "validation mode error", err.Error())
	}
}

func TestValidationErrorNandCondition(t *testing.T) {
	c := &Context{
		Ngsi: &ngsilib.NGSI{},
		Flags: []Flag{
			&StringFlag{Name: "flag1", Value: "", Set: true},
			&StringFlag{Name: "flag2", Value: "", Set: true},
		},
	}

	v := &ValidationFlag{Mode: NandCondition, Flags: []string{"flag1", "flag2"}}

	err := validation(v, c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "specify either --flag1 or --flag2", err.Error())
	}
}

func TestValidationErrorXnorCondition(t *testing.T) {
	c := &Context{
		Ngsi: &ngsilib.NGSI{},
		Flags: []Flag{
			&StringFlag{Name: "flag1", Value: "", Set: true},
			&StringFlag{Name: "flag2", Value: "", Set: true},
		},
	}

	v := &ValidationFlag{Mode: XnorCondition, Flags: []string{"flag1", "flag2"}}

	err := validation(v, c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify either --flag1 or --flag2", err.Error())
	}
}
