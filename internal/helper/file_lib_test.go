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

package helper

import (
	"bufio"
	"bytes"
	"errors"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestFileLibOpen(t *testing.T) {
	err := errors.New("open error")
	f := &MockFileLib{OpenError: err}

	actual := f.Open("path")

	if assert.Error(t, actual) {
		expected := "open error"
		assert.Equal(t, expected, actual.Error())
	}
}

func TestFileLibClose(t *testing.T) {
	err := errors.New("close error")
	f := &MockFileLib{CloseError: err}

	actual := f.Close()

	if assert.Error(t, actual) {
		expected := "close error"
		assert.Equal(t, expected, actual.Error())
	}
}

func TestFilePathAbs(t *testing.T) {
	f := &MockFileLib{FilePathAbsString: "/root"}

	s, err := f.FilePathAbs("root")

	if assert.NoError(t, err) {
		assert.Equal(t, "/root", s)
	}
}

func TestFilePathAbsError(t *testing.T) {
	f := &MockFileLib{FilePathAbsString: "/root", FilePathAbsError: [5]error{errors.New("FilePathAbs error")}}

	_, err := f.FilePathAbs("root")

	if assert.Error(t, err) {
		assert.Equal(t, "FilePathAbs error", err.Error())
	}
}

func TestSetFilePatAbsError(t *testing.T) {
	ngsi := &ngsilib.NGSI{FileReader: &MockFileLib{FilePathAbsError: [5]error{}}}

	SetFilePatAbsError(ngsi, 1)

	assert.Equal(t, "filepathabs error", (ngsi.FileReader).(*MockFileLib).FilePathAbsError[1].Error())

}
func TestReadAll(t *testing.T) {
	f := &MockFileLib{ReadallData: []byte("fiware")}

	actual, err := f.ReadAll(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", string(actual))
	}
}

func TestReadAllError(t *testing.T) {
	f := &MockFileLib{ReadallError: errors.New("ReadAll error")}

	_, err := f.ReadAll(nil)

	if assert.Error(t, err) {
		assert.Equal(t, "ReadAll error", err.Error())
	}
}

func TestReadFile(t *testing.T) {
	f := &MockFileLib{ReadFileData: []byte("orion")}

	actual, err := f.ReadFile("fiware")

	if assert.NoError(t, err) {
		assert.Equal(t, "orion", string(actual))
	}
}

func TestReadFileError(t *testing.T) {
	f := &MockFileLib{ReadFileError: [5]error{errors.New("ReadFile error")}}

	actual, err := f.ReadFile("fiware")

	if assert.Error(t, err) {
		assert.Equal(t, ([]byte)(nil), actual)
		assert.Equal(t, "ReadFile error", err.Error())
	}
}

func TestSetReadFileError(t *testing.T) {
	ngsi := &ngsilib.NGSI{FileReader: &MockFileLib{ReadFileError: [5]error{}}}

	SetReadFileError(ngsi, 0)

	assert.Equal(t, "readfile error", (ngsi.FileReader).(*MockFileLib).ReadFileError[0].Error())

}

func TestSetReader(t *testing.T) {
	f := &MockFileLib{}

	r := bytes.NewReader([]byte("fiware"))

	f.SetReader(r)

	assert.NotEqual(t, (*bufio.Reader)(nil), f.IoReader)

}

func TestFile(t *testing.T) {
	f := &MockFileLib{}
	r := bytes.NewReader([]byte("fiware"))
	f.SetReader(r)

	_ = f.File()
}

func TestFileNil(t *testing.T) {
	r := bytes.NewReader([]byte("fiware"))
	f := &MockFileLib{FileError: bufio.NewReader(r)}

	_ = f.File()
}
