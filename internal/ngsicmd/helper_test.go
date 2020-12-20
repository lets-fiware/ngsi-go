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

package ngsicmd

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func setupTest() (*ngsilib.NGSI, *flag.FlagSet, *cli.App, *bytes.Buffer) {
	ngsilib.Reset()
	filename := ""
	ngsi := ngsilib.NewNGSI()
	ngsi.ConfigFile = &MockIoLib{}
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.CacheFile = &MockIoLib{}
	ngsi.CacheFile.SetFileName(&filename)
	ngsi.HTTP = NewMockHTTP()
	buffer := &bytes.Buffer{}
	ngsi.StdWriter = buffer
	ngsi.LogWriter = &bytes.Buffer{}

	set := flag.NewFlagSet("test", 0)
	setupFlagString(set, "config,cacheFile")
	app := cli.NewApp()

	_ = set.Parse([]string{"--config=", "--cacheFile="})

	return ngsi, set, app, buffer
}

func setupTest2() (*ngsilib.NGSI, *flag.FlagSet, *cli.App, *bytes.Buffer) {
	buffer := &bytes.Buffer{}
	ngsi := ngsilib.NewNGSI()
	ngsi.StdWriter = buffer

	set := flag.NewFlagSet("test", 0)
	// setupFlagString(set, "config,cacheFile")
	app := cli.NewApp()

	return ngsi, set, app, buffer
}

func setupFlagString(set *flag.FlagSet, s string) {
	for _, flag := range strings.Split(s, ",") {
		set.String(flag, "", "doc")
	}
}

func setupFlagBool(set *flag.FlagSet, s string) {
	for _, flag := range strings.Split(s, ",") {
		set.Bool(flag, false, "doc")
	}
}

func setupFlagInt64(set *flag.FlagSet, s string) {
	for _, flag := range strings.Split(s, ",") {
		set.Int64(flag, 0, "doc")
	}
}

func setupAddBroker(t *testing.T, ngsi *ngsilib.NGSI, host string, brokerHost string, ngsiType string) {
	broker := ngsilib.Broker{BrokerHost: brokerHost, NgsiType: ngsiType}

	list := ngsi.BrokerList()
	(*list)[host] = &broker
	ngsi.Host = host
}

func setupAddBroker2(t *testing.T, ngsi *ngsilib.NGSI, host, brokerHost, ngsiType, idmType, idmHost, username, password string) {
	broker := ngsilib.Broker{BrokerHost: brokerHost, NgsiType: ngsiType, IdmType: idmType, IdmHost: idmHost, Username: username, Password: password}

	list := ngsi.BrokerList()
	(*list)[host] = &broker
	ngsi.Host = host
}

func setupDeleteBroker(t *testing.T, host string) {
	_, set, app, _ := setupTest()

	setupFlagString(set, "host")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--host=" + host})
	err := brokersDelete(c)
	assert.NoError(t, err)
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

func AddReqRes(ngsi *ngsilib.NGSI, r MockHTTPReqRes) {
	h, _ := ngsi.HTTP.(*MockHTTP)
	h.ReqRes = append(h.ReqRes, r)
}

// Request is ...
func (h *MockHTTP) Request(method string, url *url.URL, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	const funcName = "Request"

	if len(h.ReqRes) == 0 {
		panic(errors.New("ReqRes length is 0"))
	}
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
			return nil, nil, &ngsiCmdError{funcName, 0, "Unsupported type", nil}
		}
	}
	if data != nil && r.ReqData != nil {
		if !reflect.DeepEqual(r.ReqData, data) {
			fmt.Printf("r.ReqData: %s\n", string(r.ReqData))
			fmt.Printf("Data: %s\n", string(data))
			return nil, nil, &ngsiCmdError{funcName, 1, "body data error", nil}
		}
	}
	if r.Path != "" && r.Path != url.Path {
		return nil, nil, &ngsiCmdError{funcName, 3, "url error", nil}
	}
	if r.ResHeader != nil {
		r.Res.Header = r.ResHeader
	}
	return &r.Res, r.ResBody, r.Err
}

//
// MockIoLib
//

type MockIoLib struct {
	OpenErr   error
	EncodeErr error
	filename  *string
	HomeDir   error
	PathAbs   error
}

func (io *MockIoLib) Open() (err error) {
	return io.OpenErr
}

func (io *MockIoLib) OpenFile(flag int, perm os.FileMode) (err error) {
	return io.OpenErr
}

func (io *MockIoLib) Truncate(size int64) error {
	return nil
}

func (io *MockIoLib) Close() error {
	return nil
}

func (io *MockIoLib) Decode(v interface{}) error {
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
	return nil, nil
}

func (io *MockIoLib) SetFileName(filename *string) {
	io.filename = filename
}

func (io *MockIoLib) FileName() *string {
	return io.filename
}

func (io *MockIoLib) Getenv(key string) string {
	return ""
}

func (io *MockIoLib) FilePathAbs(path string) (string, error) {
	return "", io.PathAbs
}

func (io *MockIoLib) FilePathJoin(elem ...string) string {
	return strings.Join(elem, "/")
}

//
// MockFileLib
//
type MockFileLib struct {
	Name             string
	openError        error
	readallError     error
	readall          []byte
	filePathAbs      string
	filePathAbsError error
	readFile         []byte
	readFileError    error
	fileError        io.Reader
	fileError2       io.Reader
}

func (f *MockFileLib) Open(path string) (err error) {
	return f.openError
}

func (f *MockFileLib) Close() error {
	return nil
}

func (f *MockFileLib) FilePathAbs(path string) (string, error) {
	if f.filePathAbsError == nil {
		return f.filePathAbs, nil
	}
	return "", f.filePathAbsError
}

func (f *MockFileLib) ReadAll(r io.Reader) ([]byte, error) {
	if f.readall == nil {
		return nil, f.readallError
	}
	return f.readall, nil
}

func (f *MockFileLib) ReadFile(filename string) ([]byte, error) {
	if f.readFileError == nil {
		return f.readFile, nil
	}
	return nil, f.readFileError
}

func (f *MockFileLib) SetReader(r io.Reader) {
}

func (f *MockFileLib) File() io.Reader {
	r := f.fileError
	f.fileError = f.fileError2
	return r
}

//
// MockJSONLIB
//
type MockJSONLib struct {
	DecodeErr  error
	DecodeErr2 error
	EncodeErr  error
	IndentErr  error
	ValidErr   *bool
	Jsonlib    ngsilib.JSONLib
}

func (j *MockJSONLib) Decode(r io.Reader, v interface{}) error {
	if j.DecodeErr == nil {
		j.DecodeErr = j.DecodeErr2
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

func (j *MockJSONLib) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	if j.IndentErr == nil {
		return j.Jsonlib.Indent(dst, src, prefix, indent)
	}
	return j.IndentErr
}

func (j *MockJSONLib) Valid(data []byte) bool {
	if j.ValidErr != nil {
		return *j.ValidErr
	}
	return j.Jsonlib.Valid(data)
}

//
// MockSyslogLIb
//
type MockSyslogLib struct {
	Err error
	Buf *bytes.Buffer
}

func (s *MockSyslogLib) New() (io.Writer, error) {
	if s.Err == nil {
		s.Buf = new(bytes.Buffer)
		return s.Buf, nil
	}
	return nil, s.Err
}

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
	return t.unixTime + time.Now().Unix()
}
