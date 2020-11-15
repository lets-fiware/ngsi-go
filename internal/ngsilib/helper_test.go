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
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

type MockTimeLib struct {
	dateTime string
	unixTime int64
}

func (t *MockTimeLib) Now() time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	tm, _ := time.Parse(layout, t.dateTime)
	return tm
}

func (t *MockTimeLib) NowUnix() int64 {
	return t.unixTime
}

//
// MockJSONLIB
//
type MockJSONLib struct {
	DecodeErr error
	EncodeErr error
	Jsonlib   JSONLib
}

func (j *MockJSONLib) Decode(r io.Reader, v interface{}) error {
	if j.DecodeErr == nil {
		return j.Jsonlib.Decode(r, v)
	}
	return j.DecodeErr
}

func (j *MockJSONLib) Encode(w io.Writer, v interface{}) error {
	if j.EncodeErr == nil {
		return j.Jsonlib.Encode(w, v)
	}
	return j.EncodeErr
}

func testNgsiLibInit() *NGSI {
	gNGSI = nil
	return NewNGSI()
}

//
// MockIoLib
//

type MockIoLib struct {
	OpenErr      error
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
	return nil
}

func (io *MockIoLib) Decode(v interface{}) error {
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

// MockHTTPRequest is ...
type MockHTTPRequest interface {
	Request(method string, url *url.URL, headers map[string]string, body interface{}) (*http.Response, []byte, error)
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
		switch body.(type) {
		case []byte:
			data = body.([]byte)
		case string:
			data = []byte(body.(string))
		default:
			return nil, nil, &NgsiLibError{funcName, 0, "Unsupported type", nil}
		}
	}
	if data != nil && r.ReqData != nil {
		if !reflect.DeepEqual(r.ReqData, data) {
			return nil, nil, &NgsiLibError{funcName, 1, "body data error", nil}
		}
	}
	if r.Path != "" && r.Path != url.Path {
		return nil, nil, &NgsiLibError{funcName, 3, "url error", nil}
	}
	if r.ResHeader != nil {
		r.Res.Header = r.ResHeader
	}
	return &r.Res, r.ResBody, r.Err
}
