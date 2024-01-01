/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

package helper

import (
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestMockReadAll(t *testing.T) {
	_ = ngsilib.NewNGSI()

	actual, err := MockReadAll("fiware")

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", string(actual))
	}
}

func TestMockReadAllError(t *testing.T) {
	_ = ngsilib.NewNGSI()

	actual, err := MockReadAllError("fiware")

	if assert.Error(t, err) {
		assert.Equal(t, ([]byte)(nil), actual)
		assert.Equal(t, "readall error", err.Error())
	}
}

func TestGetMockReader(t *testing.T) {
	_ = ngsilib.NewNGSI()

	actual, err := MockGetReader("fiware")

	if assert.NoError(t, err) {
		assert.NotEqual(t, (ngsilib.FileLib)(nil), actual)
	}
}

func TestGetMockReaderError(t *testing.T) {
	_ = ngsilib.NewNGSI()

	actual, err := MockGetReaderError("fiware")

	if assert.Error(t, err) {
		assert.Equal(t, (ngsilib.FileLib)(nil), actual)
		assert.Equal(t, "getreader error", err.Error())
	}
}
