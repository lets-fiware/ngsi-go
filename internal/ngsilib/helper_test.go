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
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// testNgsiLibInit
func testNgsiLibInit() *NGSI {
	gNGSI = nil

	ngsi := NewNGSI()
	ngsi.FileReader = &MockFileLib{}

	return ngsi
}

// MockTimeLib
type MockTimeLib struct {
	dateTime string
	unixTime int64
	format   *string
	tTime    *time.Time
}

func (t *MockTimeLib) Now() time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	tm, _ := time.Parse(layout, t.dateTime)
	return tm
}

func (t *MockTimeLib) NowUnix() int64 {
	return t.unixTime
}

func (t *MockTimeLib) Unix(sec int64, nsec int64) time.Time {
	if t.tTime != nil {
		return *t.tTime
	}
	tt := time.Unix(sec, nsec)
	t.tTime = &tt
	return tt
}

func (t *MockTimeLib) Format(layout string) string {
	if t.format != nil {
		return *t.format
	}
	return t.tTime.Format(layout)
}

// MockJSONLib
type MockJSONLib struct {
	IndentErr error
	ValidErr  *bool
	Jsonlib   JSONLib
	DecodeErr [5]error
	EncodeErr [5]error
	dp        int
	ep        int
}

func (j *MockJSONLib) Decode(r io.Reader, v interface{}) error {
	err := j.DecodeErr[j.dp]
	j.dp++
	if err == nil {
		return j.Jsonlib.Decode(r, v)
	}
	return err
}

func (j *MockJSONLib) Encode(w io.Writer, v interface{}) error {
	err := j.EncodeErr[j.ep]
	j.ep++
	if err == nil {
		return j.Jsonlib.Encode(w, v)
	}
	return err
}

func SetJSONDecodeErr(ngsi *NGSI, p int) {
	j := ngsi.JSONConverter
	mockj := &MockJSONLib{Jsonlib: j}
	mockj.DecodeErr[p] = errors.New("json error")
	ngsi.JSONConverter = mockj
}

func SetJSONEncodeErr(ngsi *NGSI, p int) {
	j := ngsi.JSONConverter
	mockj := &MockJSONLib{Jsonlib: j}
	mockj.EncodeErr[p] = errors.New("json error")
	ngsi.JSONConverter = mockj
}

func (j *MockJSONLib) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	if j.IndentErr == nil {
		return j.Jsonlib.Indent(dst, src, prefix, indent)
	}
	return j.IndentErr
}

func SetJSONIndentError(ngsi *NGSI) {
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{IndentErr: errors.New("json error"), Jsonlib: j}
}

func (j *MockJSONLib) Valid(data []byte) bool {
	if j.ValidErr != nil {
		return *j.ValidErr
	}
	return j.Jsonlib.Valid(data)
}

// MockIoLib
//

type MockIoLib struct {
	OpenErr      error
	CloseErr     error
	TruncateErr  error
	EncodeErr    error
	filename     *string
	HomeDir      string
	HomeDirErr   error
	PathAbsErr   error
	ConfigDir    string
	ConfigDirErr error
	StatErr      error
	MkdirErr     error
	DecodeErr    error
	Env          string
	Data         *string
}

func (io *MockIoLib) Open() (err error) {
	return io.OpenErr
}

func (io *MockIoLib) OpenFile(flag int, perm os.FileMode) (err error) {
	return io.OpenErr
}

func (io *MockIoLib) Truncate(size int64) error {
	return io.TruncateErr
}

func (io *MockIoLib) Close() error {
	return io.CloseErr
}

func (io *MockIoLib) Decode(v interface{}) error {
	if io.Data != nil {
		return json.Unmarshal([]byte(*io.Data), v)
	}
	return io.DecodeErr
}

func (io *MockIoLib) Encode(v interface{}) error {
	return io.EncodeErr
}

func (io *MockIoLib) UserHomeDir() (string, error) {
	return io.HomeDir, io.HomeDirErr
}

func (io *MockIoLib) UserConfigDir() (string, error) {
	return io.ConfigDir, io.ConfigDirErr
}

func (io *MockIoLib) MkdirAll(path string, perm os.FileMode) error {
	return io.MkdirErr
}

func (io *MockIoLib) Stat(filename string) (os.FileInfo, error) {
	return nil, io.StatErr
}

func (io *MockIoLib) SetFileName(filename *string) {
	io.filename = filename
}

func (io *MockIoLib) FileName() *string {
	return io.filename
}

func (io *MockIoLib) Getenv(key string) string {
	return io.Env
}

func (io *MockIoLib) FilePathAbs(path string) (string, error) {
	return path, io.PathAbsErr
}

func (io *MockIoLib) FilePathJoin(elem ...string) string {
	return strings.Join(elem, "/")
}

func NewMockHTTP() *MockHTTP {
	m := MockHTTP{}
	return &m
}

// MockHTTPReqRes is ...
type MockHTTP struct {
	index  int
	ReqRes []MockHTTPReqRes
}

type MockHTTPReqRes struct {
	Res        http.Response
	ResBody    []byte
	ResHeader  http.Header
	Err        error
	StatusCode int
	ReqData    []byte
	Path       string
}

// Request is ...
func (h *MockHTTP) Request(method string, url *url.URL, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	const funcName = "Request"

	r := h.ReqRes[h.index]
	h.index++

	if r.Err != nil {
		return nil, nil, r.Err
	}
	var data []byte
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		switch body := body.(type) {
		case []byte:
			data = body
		case string:
			data = []byte(body)
		default:
			return nil, nil, ngsierr.New(funcName, 0, "Unsupported type", nil)
		}
	}
	if data != nil && r.ReqData != nil {
		if !reflect.DeepEqual(r.ReqData, data) {
			return nil, nil, ngsierr.New(funcName, 1, "body data error", nil)
		}
	}
	if r.Path != "" && r.Path != url.Path {
		return nil, nil, ngsierr.New(funcName, 3, "url error", nil)
	}
	if r.ResHeader != nil {
		r.Res.Header = r.ResHeader
	}
	return &r.Res, r.ResBody, r.Err
}

// MockFileLib
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

func setFilePatAbsError(ngsi *NGSI, p int) {
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

func setReadFileError(ngsi *NGSI, p int) {
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
	if err == nil {
		return *f.IoReader
	}
	return *err
}

// MockIoutilLib
type MockIoutilLib struct {
	CopyErr      error
	WriteFileErr error
	ReadFileErr  error
}

func (i *MockIoutilLib) Copy(dst io.Writer, src io.Reader) (int64, error) {
	if i.CopyErr != nil {
		return 0, i.CopyErr
	}
	return io.Copy(dst, src)
}

func (i *MockIoutilLib) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if i.WriteFileErr != nil {
		return i.WriteFileErr
	}
	return os.WriteFile(filename, data, perm)
}

func (i *MockIoutilLib) ReadFile(filename string) ([]byte, error) {
	if i.ReadFileErr != nil {
		return nil, i.ReadFileErr
	}
	return os.ReadFile(filename)
}

// MockFilePathLib
type MockFilePathLib struct {
	PathAbsErr error
}

func (i *MockFilePathLib) FilePathAbs(path string) (string, error) {
	return path, i.PathAbsErr
}

func (i *MockFilePathLib) FilePathJoin(elem ...string) string {
	return strings.Join(elem, "/")
}

func (i *MockFilePathLib) FilePathBase(path string) string {
	return filepath.Base(path)
}
