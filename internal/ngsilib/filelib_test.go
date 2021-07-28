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

package ngsilib

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileLibOpen(t *testing.T) {
	f := &fileLib{}
	err := f.Open("???")

	assert.Error(t, err)
}

func TestFileLibClose(t *testing.T) {
	f := &fileLib{}
	err := f.Open(".")
	assert.NoError(t, err)

	err = f.Close()

	assert.Error(t, err)
}

func TestFileLibCloseNil(t *testing.T) {
	f := &fileLib{}
	f.file = nil
	err := f.Close()

	assert.NoError(t, err)
}

func TestFileLibCloseError(t *testing.T) {
	f := &fileLib{}
	err := f.Open("???")
	assert.Error(t, err)

	err = f.Close()

	assert.NoError(t, err)
}

func TestFileLibFilePathAbs(t *testing.T) {
	f := &fileLib{}
	_, err := f.FilePathAbs("")

	assert.NoError(t, err)
}

func TestFileLibReadAll(t *testing.T) {
	f := &fileLib{}
	buf := &bytes.Buffer{}
	_, err := f.ReadAll(buf)

	assert.NoError(t, err)
}

func TestFileLibReadFile(t *testing.T) {
	f := &fileLib{}
	_, err := f.ReadFile("???")

	assert.Error(t, err)
}

func TestFileLibSetReader(t *testing.T) {
	f := &fileLib{}
	buf := &bytes.Buffer{}
	f.SetReader(buf)

	reader := f.File()
	_, err := reader.Peek(1)

	if assert.Error(t, err) {
		assert.Equal(t, "EOF", err.Error())
	}
}

func TestFileLibFile(t *testing.T) {
	f := &fileLib{}
	buf := &bytes.Buffer{}
	f.SetReader(buf)

	file := f.File()
	_, err := file.Peek(1)

	if assert.Error(t, err) {
		assert.Equal(t, "EOF", err.Error())
	}
}
