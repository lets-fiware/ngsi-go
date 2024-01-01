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
	"encoding/json"
	"os"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type MockIoLib struct {
	OpenErr    error
	EncodeErr  error
	Filename   *string
	HomeDir    error
	PathAbs    error
	StatErr    error
	Tokens     *string
	TruncIndex int
	Trunc      []error
}

func (io *MockIoLib) Open() (err error) {
	return io.OpenErr
}

func (io *MockIoLib) OpenFile(flag int, perm os.FileMode) (err error) {
	return io.OpenErr
}

func (io *MockIoLib) Truncate(size int64) error {
	if io.Trunc != nil {
		e := io.Trunc[io.TruncIndex]
		io.TruncIndex++
		return e
	}
	return nil
}

func (io *MockIoLib) Close() error {
	return nil
}

func (io *MockIoLib) Decode(v interface{}) error {
	if io.Tokens != nil {
		return json.Unmarshal([]byte(*io.Tokens), v)
	}
	config, ok := v.(*ngsilib.NgsiConfig)
	if ok {
		config.DefaultValues = ngsilib.Settings{UsePreviousArgs: true}
	}
	return nil
}

func (io *MockIoLib) Encode(v interface{}) error {
	return io.EncodeErr
}

func (io *MockIoLib) UserHomeDir() (string, error) {
	return "", io.HomeDir
}

func (io *MockIoLib) UserConfigDir() (string, error) {
	return "", nil
}

func (io *MockIoLib) MkdirAll(path string, perm os.FileMode) error {
	return nil
}

func (io *MockIoLib) Stat(filename string) (os.FileInfo, error) {
	return nil, io.StatErr
}

func (io *MockIoLib) SetFileName(filename *string) {
	io.Filename = filename
}

func (io *MockIoLib) FileName() *string {
	return io.Filename
}

func (io *MockIoLib) Getenv(key string) string {
	return ""
}

func (io *MockIoLib) FilePathAbs(path string) (string, error) {
	return path, io.PathAbs
}

func (io *MockIoLib) FilePathJoin(elem ...string) string {
	return strings.Join(elem, "/")
}
