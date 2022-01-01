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
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestOpen(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.Open()

	assert.Error(t, err)
}

func TestOpenFile(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)

	assert.Error(t, err)
}
func TestTruncate(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	err = io.Truncate(0)

	assert.Error(t, err)
}

func TestFileClose(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	err = io.Close()

	assert.Error(t, err)
}

func TestNewDecoder(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	var i int
	err = io.Decode(&i)

	assert.Error(t, err)
}

func TestNewEncoder(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	var i int
	err = io.Encode(&i)

	assert.Error(t, err)
}

func TestUserHomeDir(t *testing.T) {
	io := ioLib{}

	_, err := io.UserHomeDir()

	assert.NoError(t, err)
}

func TestUserConfigDir(t *testing.T) {
	io := ioLib{}

	_, err := io.UserConfigDir()

	assert.NoError(t, err)
}

func TestMkdirAll(t *testing.T) {
	io := ioLib{}

	err := io.MkdirAll("", 0700)

	if assert.Error(t, err) {
		assert.Equal(t, "mkdir : no such file or directory", err.Error())
	}
}

func TestStat(t *testing.T) {
	io := ioLib{}

	_, err := io.Stat("")

	if assert.Error(t, err) {
		assert.Equal(t, "stat : no such file or directory", err.Error())
	}
}

func TestSetFileName(t *testing.T) {
	io := ioLib{}

	s := "test"
	io.SetFileName(&s)

	assert.Equal(t, s, *io.fileName)
}

func TestFileName(t *testing.T) {
	io := ioLib{}

	s := "test"
	io.SetFileName(&s)

	f := io.FileName()

	assert.Equal(t, s, *f)
}

func TestGetenv(t *testing.T) {
	io := ioLib{}

	e := io.Getenv("")

	assert.Equal(t, "", e)
}

func TestFilePathAbs(t *testing.T) {
	io := ioLib{}

	_, err := io.FilePathAbs("")

	assert.NoError(t, err)
}

func TestFilePathJoin(t *testing.T) {
	io := ioLib{}

	s := io.FilePathJoin("/", "home")

	assert.Equal(t, "/home", s)
}
