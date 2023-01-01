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
	"bytes"
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestIoutilLibCopy(t *testing.T) {
	i := &MockIoutilLib{}

	src := bytes.NewReader([]byte("fiware"))
	dst := &bytes.Buffer{}

	actual, err := i.Copy(dst, src)

	if assert.NoError(t, err) {
		assert.Equal(t, int64(6), actual)
		assert.Equal(t, "fiware", dst.String())
	}
}

func TestIoutilLibCopyErr(t *testing.T) {
	i := &MockIoutilLib{CopyErr: errors.New("Copy error")}

	src := bytes.NewReader([]byte("fiware"))
	dst := &bytes.Buffer{}

	_, err := i.Copy(dst, src)

	if assert.Error(t, err) {
		assert.Equal(t, "Copy error", err.Error())
	}
}

func TestIoutilLibReadFull(t *testing.T) {
	i := &MockIoutilLib{}

	src := bytes.NewReader([]byte("fiware"))
	dst := []byte("      ")

	actual, err := i.ReadFull(src, dst)

	if assert.NoError(t, err) {
		assert.Equal(t, 6, actual)
		assert.Equal(t, "fiware", string(dst))
	}
}

func TestIoutilLibReadFullData(t *testing.T) {
	i := &MockIoutilLib{ReadFullData: []byte("fiware")}

	dst := []byte("      ")

	actual, err := i.ReadFull(nil, dst)

	if assert.NoError(t, err) {
		assert.Equal(t, 6, actual)
		assert.Equal(t, "fiware", string(dst))
	}
}

func TestIoutilLibReadFullError(t *testing.T) {
	i := &MockIoutilLib{ReadFullErr: errors.New("ReadFull error")}

	actual, err := i.ReadFull(nil, nil)

	if assert.Error(t, err) {
		assert.Equal(t, 0, actual)
		assert.Equal(t, "ReadFull error", err.Error())
	}
}

func TestIoutilLibWriteFile(t *testing.T) {
	i := &MockIoutilLib{}

	err := i.WriteFile("", nil, fs.FileMode(os.O_RDONLY))

	if assert.Error(t, err) {
		assert.Equal(t, "open : no such file or directory", err.Error())
	}
}

func TestIoutilLibWriteFileWriteSkip(t *testing.T) {
	i := &MockIoutilLib{WriteSkip: true}

	err := i.WriteFile("", nil, fs.FileMode(os.O_RDONLY))

	assert.NoError(t, err)
}

func TestIoutilLibWriteFileError(t *testing.T) {
	i := &MockIoutilLib{WriteFileErr: errors.New("WriteFile error")}

	err := i.WriteFile("", nil, fs.FileMode(os.O_RDONLY))

	if assert.Error(t, err) {
		assert.Equal(t, "WriteFile error", err.Error())
	}
}

func TestIoutilLibReadFile(t *testing.T) {
	i := &MockIoutilLib{}

	_, err := i.ReadFile("")

	if assert.Error(t, err) {
		assert.Equal(t, "open : no such file or directory", err.Error())
	}
}

func TestIoutilLibReadFileError(t *testing.T) {
	i := &MockIoutilLib{ReadFileErr: errors.New("ReadFile error")}

	_, err := i.ReadFile("")

	if assert.Error(t, err) {
		assert.Equal(t, "ReadFile error", err.Error())
	}
}

func TestIoutilLibReadFileData(t *testing.T) {
	i := &MockIoutilLib{ReadFileData: []byte("fiware")}

	actual, err := i.ReadFile("")

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", string(actual))
	}
}
