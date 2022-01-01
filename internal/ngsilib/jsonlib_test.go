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

package ngsilib

import (
	"bytes"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestJSONLibDecode(t *testing.T) {
	j := &jsonLib{}
	buf := &bytes.Buffer{}
	var i int
	err := j.Decode(buf, &i)

	if assert.Error(t, err) {
		assert.Equal(t, "EOF", err.Error())
	}
}

func TestJSONLibEncode(t *testing.T) {
	j := &jsonLib{}
	buf := &bytes.Buffer{}
	var i int
	err := j.Encode(buf, &i)

	assert.NoError(t, err)
}

func TestIndent(t *testing.T) {
	j := &jsonLib{}
	buf := &bytes.Buffer{}
	src := []byte("{}")
	err := j.Indent(buf, src, "", "  ")

	assert.NoError(t, err)
}

func TestValid(t *testing.T) {
	j := &jsonLib{}
	src := []byte("{}")
	actual := j.Valid(src)

	expected := true

	assert.Equal(t, expected, actual)
}
