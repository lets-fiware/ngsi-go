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
	"errors"
	"io"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type MockFileLib struct {
	Name              string
	OpenError         error
	CloseError        error
	ReadallError      error
	ReadallData       []byte
	FilePathAbsString string
	FilePathAbsError  [5]error
	ab                int
	ReadFileData      []byte
	ReadFileError     [5]error
	rf                int
	FileError         *bufio.Reader
	FileError2        *bufio.Reader
	IoReader          *bufio.Reader
}

func (f *MockFileLib) Open(path string) (err error) {
	return f.OpenError
}

func (f *MockFileLib) Close() error {
	return f.CloseError
}

func (f *MockFileLib) FilePathAbs(path string) (string, error) {
	err := f.FilePathAbsError[f.ab]
	f.ab++
	if err == nil {
		return f.FilePathAbsString, nil
	}
	return "", err
}

func SetFilePatAbsError(ngsi *ngsilib.NGSI, p int) {
	f := ngsi.FileReader.(*MockFileLib)
	f.FilePathAbsError[p] = errors.New("filepathabs error")
	ngsi.FileReader = f
}

func (f *MockFileLib) ReadAll(r io.Reader) ([]byte, error) {
	if f.ReadallData == nil {
		return nil, f.ReadallError
	}
	return f.ReadallData, nil
}

func (f *MockFileLib) ReadFile(filename string) ([]byte, error) {
	err := f.ReadFileError[f.rf]
	f.rf++
	if err == nil {
		return f.ReadFileData, nil
	}
	return nil, err
}

func SetReadFileError(ngsi *ngsilib.NGSI, p int) {
	f := ngsi.FileReader.(*MockFileLib)
	f.ReadFileError[p] = errors.New("readfile error")
	ngsi.FileReader = f
}

func (f *MockFileLib) SetReader(r io.Reader) {
	f.IoReader = bufio.NewReader(r)
}

func (f *MockFileLib) File() bufio.Reader {
	err := f.FileError
	f.FileError = f.FileError2
	if err == (*bufio.Reader)(nil) {
		return *f.IoReader
	}
	return *err
}
