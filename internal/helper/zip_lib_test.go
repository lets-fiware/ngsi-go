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

package helper

import (
	"archive/zip"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockIolib struct {
}

func (i *mockIolib) Write(p []byte) (n int, err error) {
	return 0, nil
}

type mockReaderAt struct {
}

func (r *mockReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}

func TestZipLibNewReader(t *testing.T) {
	r := &mockReaderAt{}
	z := &MockZipLib{}

	actual, err := z.NewReader(r, 0)

	assert.Equal(t, (*zip.Reader)(nil), actual)
	assert.Equal(t, "zip: not a valid zip file", err.Error())
}

func TestZipLibNewReaderZipReader(t *testing.T) {
	r := &mockReaderAt{}
	z := &MockZipLib{ZipReader: &zip.Reader{}}

	actual, err := z.NewReader(r, 0)

	assert.NotEqual(t, (*zip.Reader)(nil), actual)
	assert.Equal(t, nil, err)
}

func TestZipLibNewReaderErrorZip(t *testing.T) {
	r := &mockReaderAt{}
	z := &MockZipLib{Zip: errors.New("zip error")}

	actual, err := z.NewReader(r, 0)

	assert.Equal(t, (*zip.Reader)(nil), actual)
	assert.Equal(t, "zip error", err.Error())
}

func TestZipLibNewWriter(t *testing.T) {
	w := &mockIolib{}
	m := &MockMultiPart{}

	actual := m.NewWriter(w)

	assert.NotEqual(t, nil, actual)
}
