/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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

// FileLib is ...
type FileLib interface {
	Open(path string) error
	Close() error
	FilePathAbs(path string) (string, error)
	ReadAll(r io.Reader) ([]byte, error)
	ReadFile(filename string) ([]byte, error)
	SetReader(r io.Reader)
	File() io.Reader
}

type fileLib struct {
	file io.Reader
}

func (f *fileLib) Open(path string) (err error) {
	f.file, err = os.Open(path)
	return
}

func (f *fileLib) Close() error {
	if f.file == nil {
		return nil
	}
	file, _ := f.file.(*os.File)
	err := file.Close()
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
	f.file = r
}

func (f *fileLib) File() io.Reader {
	return f.file
}

// JSONLib is
type JSONLib interface {
	Decode(r io.Reader, v interface{}) error
	Encode(w io.Writer, v interface{}) error
	Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
}

type jsonLib struct {
}

func (j *jsonLib) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (j *jsonLib) Encode(w io.Writer, v interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(v)
}

func (j *jsonLib) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}

// TimeLib is ...
type TimeLib interface {
	Now() time.Time
	NowUnix() int64
}

type timeLib struct{}

func (t *timeLib) Now() time.Time {
	return time.Now()
}

func (t *timeLib) NowUnix() int64 {
	return time.Now().Unix()
}
