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
	"errors"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestReadAll(t *testing.T) {
	_ = testNgsiLibInit()

	b, err := ReadAll(`{"id":test"}`)

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(`{"id":test"}`), b)
	}
}

func TestReadAllStdReader(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.FileReader = &MockFileLib{ReadallData: []byte("test data")}

	b, err := ReadAll("stdin")

	if assert.NoError(t, err) {
		assert.Equal(t, []byte("test data"), b)
	}
}

func TestReadAllAt(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.FileReader = &MockFileLib{FilePathAbsString: "file", ReadFileData: []byte(`{"id":test"}`)}

	b, err := ReadAll("@file")

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(`{"id":test"}`), b)
	}
}

func TestReadAllErrorEmpty(t *testing.T) {
	_ = testNgsiLibInit()

	_, err := ReadAll("")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestReadAllErrorStdReader(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.FileReader = &MockFileLib{ReadallError: errors.New("ReadAll error")}

	_, err := ReadAll("stdin")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "ReadAll error", ngsiErr.Message)
	}
}

func TestReadAllAt3(t *testing.T) {
	ngsi := testNgsiLibInit()

	setFilePatAbsError(ngsi, 0)

	_, err := ReadAll("@file")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "filepathabs error", ngsiErr.Message)
	}
}

func TestReadAllAt4(t *testing.T) {
	ngsi := testNgsiLibInit()

	setReadFileError(ngsi, 0)

	_, err := ReadAll("@file")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "readfile error", ngsiErr.Message)
	}
}

func TestReadAllErrorAt5(t *testing.T) {
	testNgsiLibInit()

	_, err := ReadAll("@")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestGetReader(t *testing.T) {
	testNgsiLibInit()

	_, err := GetReader("{abc}")

	assert.NoError(t, err)
}

func TestGetReaderFIle(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.FileReader = &MockFileLib{FilePathAbsString: "file", ReadFileData: []byte(`{"id":test"}`)}

	_, err := GetReader("@file")

	assert.NoError(t, err)
}

func TestGetReaderStdin(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.FileReader = &MockFileLib{Name: "stdin test"}

	f, err := GetReader("stdin")

	if assert.NoError(t, err) {
		m := f.(*MockFileLib)
		assert.Equal(t, "stdin test", m.Name)
	}
}

func TestGetReaderErrorEmpty(t *testing.T) {
	testNgsiLibInit()

	_, err := GetReader("")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	}
}

func TestGetReaderAt3(t *testing.T) {
	ngsi := testNgsiLibInit()

	setFilePatAbsError(ngsi, 0)

	_, err := GetReader("@file")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "filepathabs error", ngsiErr.Message)
	}
}

func TestGetReaderAt4(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.FileReader = &MockFileLib{FilePathAbsString: "file", OpenError: errors.New("error @file")}

	_, err := GetReader("@file")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error @file", ngsiErr.Message)
	}
}

func TestGetReaderErrorAt5(t *testing.T) {
	testNgsiLibInit()

	_, err := GetReader("@")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}
