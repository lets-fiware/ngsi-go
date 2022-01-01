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

package helper

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestIoLibOpen(t *testing.T) {
	i := &MockIoLib{OpenErr: errors.New("Open error")}

	err := i.Open()

	if assert.Error(t, err) {
		assert.Equal(t, "Open error", err.Error())
	}
}

func TestIoLibOpenFile(t *testing.T) {
	i := &MockIoLib{OpenErr: errors.New("OpenFile error")}

	err := i.OpenFile(0, fs.FileMode(os.O_RDONLY))

	if assert.Error(t, err) {
		assert.Equal(t, "OpenFile error", err.Error())
	}
}

func TestIoLibTruncate(t *testing.T) {
	i := &MockIoLib{}

	err := i.Truncate(0)

	assert.NoError(t, err)
}

func TestIoLibTruncateError(t *testing.T) {
	i := &MockIoLib{Trunc: []error{errors.New("Truncate error")}}

	err := i.Truncate(0)

	if assert.Error(t, err) {
		assert.Equal(t, "Truncate error", err.Error())
	}
}

func TestIoLibClose(t *testing.T) {
	i := &MockIoLib{}

	err := i.Close()

	assert.NoError(t, err)
}

func TestIoLibDecode(t *testing.T) {
	i := &MockIoLib{}

	err := i.Decode(nil)

	assert.NoError(t, err)
}

func TestIoLibDecodeToken(t *testing.T) {
	s := `{"name":"fiware"}`
	i := &MockIoLib{Tokens: &s}

	v := make(map[string]interface{})

	err := i.Decode(&v)

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", (v["name"]).(string))
	}
}

func TestIoLibDecodeConfig(t *testing.T) {
	i := &MockIoLib{}

	config := &ngsilib.NgsiConfig{}

	err := i.Decode(config)

	if assert.NoError(t, err) {
		assert.Equal(t, true, config.DefaultValues.UsePreviousArgs)
	}
}

func TestIoLibEncode(t *testing.T) {
	i := &MockIoLib{EncodeErr: errors.New("Encode error")}

	err := i.Encode(nil)

	if assert.Error(t, err) {
		assert.Equal(t, "Encode error", err.Error())
	}
}

func TestIoLibUserHomeDir(t *testing.T) {
	i := &MockIoLib{HomeDir: errors.New("UserHomeDir error")}

	actual, err := i.UserHomeDir()

	if assert.Error(t, err) {
		assert.Equal(t, "", actual)
		assert.Equal(t, "UserHomeDir error", err.Error())
	}
}

func TestIoLibUserConfigDir(t *testing.T) {
	i := &MockIoLib{}

	actual, err := i.UserConfigDir()

	if assert.NoError(t, err) {
		assert.Equal(t, "", actual)
	}
}

func TestIoLibMkdirAll(t *testing.T) {
	i := &MockIoLib{}

	err := i.MkdirAll("", fs.FileMode(os.O_RDONLY))

	assert.NoError(t, err)
}

func TestIoLibStat(t *testing.T) {
	i := &MockIoLib{StatErr: errors.New("StatErr error")}

	actual, err := i.Stat("file")

	if assert.Error(t, err) {
		assert.Equal(t, (os.FileInfo)(nil), actual)
		assert.Equal(t, "StatErr error", err.Error())
	}
}

func TestIoLibSetFileName(t *testing.T) {
	i := &MockIoLib{}
	s := "fiware"

	i.SetFileName(&s)

	assert.Equal(t, "fiware", *i.Filename)
}

func TestIoLibFileName(t *testing.T) {
	s := "fiware"
	i := &MockIoLib{Filename: &s}

	file := i.FileName()

	assert.Equal(t, "fiware", *file)
}

func TestIoLibGetenv(t *testing.T) {
	i := &MockIoLib{}

	e := i.Getenv("fiware")

	assert.Equal(t, "", e)
}

func TestIoLibFilePathAbs(t *testing.T) {
	i := &MockIoLib{}

	s, err := i.FilePathAbs("/ngsi-ld")

	if assert.NoError(t, err) {
		assert.Equal(t, "/ngsi-ld", s)
	}
}

func TestIoLibFilePathAbsError(t *testing.T) {
	i := &MockIoLib{PathAbs: errors.New("FilePathAbs error")}

	s, err := i.FilePathAbs("/ngsi-ld")

	if assert.Error(t, err) {
		assert.Equal(t, "/ngsi-ld", s)
		assert.Equal(t, "FilePathAbs error", err.Error())
	}
}

func TestIoLibFilePathJoin(t *testing.T) {
	i := &MockIoLib{}

	s := i.FilePathJoin("/ngsi-ld", "v1")

	assert.Equal(t, "/ngsi-ld/v1", s)
}
