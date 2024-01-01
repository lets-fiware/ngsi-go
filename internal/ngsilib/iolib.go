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

package ngsilib

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	oWRONLY int = os.O_WRONLY
	oCREATE int = os.O_CREATE
)

// IoLib is ...
type IoLib interface {
	Open() error
	OpenFile(flag int, perm os.FileMode) error
	Truncate(size int64) error
	Close() error
	Decode(v interface{}) error
	Encode(v interface{}) error
	MkdirAll(path string, perm os.FileMode) error
	Stat(name string) (os.FileInfo, error)
	UserConfigDir() (string, error)
	UserHomeDir() (string, error)
	SetFileName(filename *string)
	FileName() *string
	Getenv(key string) string
	FilePathAbs(path string) (string, error)
	FilePathJoin(elem ...string) string
}

type ioLib struct {
	file     *os.File
	fileName *string
}

func (io *ioLib) Open() (err error) {
	io.file, err = os.Open(*io.fileName)
	return
}

func (io *ioLib) OpenFile(flag int, perm os.FileMode) (err error) {
	io.file, err = os.OpenFile(*io.fileName, flag, perm)
	return
}

func (io *ioLib) Truncate(size int64) error {
	return io.file.Truncate(size)
}

func (io *ioLib) Close() error {
	return io.file.Close()
}

func (io *ioLib) Decode(v interface{}) error {
	return json.NewDecoder(io.file).Decode(v)
}

func (io *ioLib) Encode(v interface{}) error {
	encoder := json.NewEncoder(io.file)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(v)
}

func (io *ioLib) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (io *ioLib) UserConfigDir() (string, error) {
	return os.UserConfigDir()
}

func (io *ioLib) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (io *ioLib) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}

func (io *ioLib) SetFileName(filename *string) {
	io.fileName = filename
}

func (io *ioLib) FileName() *string {
	return io.fileName
}

func (io *ioLib) Getenv(key string) string {
	return os.Getenv(key)
}

func (io *ioLib) FilePathAbs(path string) (string, error) {
	return filepath.Abs(path)
}

func (io *ioLib) FilePathJoin(elem ...string) string {
	return filepath.Join(elem...)
}
