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
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileLib is ...
type FileLib interface {
	Open(path string) error
	Close() error
	FilePathAbs(path string) (string, error)
	ReadAll(r io.Reader) ([]byte, error)
	ReadFile(filename string) ([]byte, error)
	SetReader(r io.Reader)
	File() bufio.Reader
}

type fileLib struct {
	file   *bufio.Reader
	osFile *os.File
}

func (f *fileLib) Open(path string) error {
	osFile, err := os.Open(path)
	if err != nil {
		f.file = nil
		return err
	}
	f.file = bufio.NewReader(osFile)
	return nil
}

func (f *fileLib) Close() error {
	if f.file == nil {
		return nil
	}
	err := f.osFile.Close()
	f.file = nil
	return err
}

func (f *fileLib) FilePathAbs(path string) (string, error) {
	return filepath.Abs(path)
}

func (f *fileLib) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

func (f *fileLib) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (f *fileLib) SetReader(r io.Reader) {
	f.file = bufio.NewReader(r)
}

func (f *fileLib) File() bufio.Reader {
	return *f.file
}
