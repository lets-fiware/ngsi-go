/*
MIT License

Copyright (c) 2020-2021 Kazuhito Suda

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
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

var configData = `{
	"version": "1",
	"servers": {
	  "orion": {
		"serverHost": "https://orion",
		"ngsiType": "v2",
		"serverType": "broker"
	  },
	  "orion-ld": {
		"serverHost": "https://orion-ld",
		"ngsiType": "ld",
		"serverType": "broker"
	  },
	  "orion-alias": {
		"serverHost": "orion-ld",
		"ngsiType": "ld",
		"serverType": "broker"
	  },
	  "comet": {
		"serverHost": "https://comet",
		"serverType": "comet"
	  },
	  "cygnus": {
		"serverHost": "https://cygnus",
		"serverType": "cygnus"
	  },
	  "ql": {
		"serverHost": "https://quantumleap",
		"serverType": "quantumleap"
	  },
	  "iota": {
		"serverHost": "https://iota",
		"serverType": "iota"
	  },
	  "perseo": {
		"serverHost": "https://perseo",
		"serverType": "perseo"
	  },
	  "perseo-core": {
		"serverHost": "https://perseo-core",
		"serverType": "perseo-core"
	  },
	  "keyrock": {
		"serverHost": "https://keyrock",
		"serverType": "keyrock"
	  }
	},
	"contexts": {
	  "data-model": "http://context-provider:3000/data-models/ngsi-context.jsonld",
	  "etsi": "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld",
	  "ld": "https://schema.lab.fiware.org/ld/context",
	  "tutorial": "http://context-provider:3000/data-models/ngsi-context.jsonld",
	  "array": [
		"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"
	  ],
	  "object": {
		"ld": "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"
	  }
	},
	"settings": {
		"usePreviousArgs": true
	}
  }`

func setupTest() (*ngsilib.NGSI, *flag.FlagSet, *cli.App, *bytes.Buffer) {
	ngsilib.Reset()

	filename := ""
	ngsi := ngsilib.NewNGSI()
	ngsi.ConfigFile = &MockIoLib{}
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.CacheFile = &MockIoLib{}
	ngsi.CacheFile.SetFileName(&filename)
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(configData)}
	ngsi.HTTP = NewMockHTTP()
	buffer := &bytes.Buffer{}
	ngsi.StdWriter = buffer
	ngsi.LogWriter = &bytes.Buffer{}

	set := flag.NewFlagSet("test", 0)
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
		switch body := body.(type) {
		case []byte:
			data = body
		case string:
			data = []byte(body)
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
	StatErr   error
	Tokens    *string
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
	FileError         io.Reader
	FileError2        io.Reader
	IoReader          io.Reader
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

func setFilePatAbsError(ngsi *ngsilib.NGSI, p int) {
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

func setReadFileError(ngsi *ngsilib.NGSI, p int) {
	f := ngsi.FileReader.(*MockFileLib)
	f.ReadFileError[p] = errors.New("readfile error")
	ngsi.FileReader = f
}

func (f *MockFileLib) SetReader(r io.Reader) {
	f.IoReader = r
}

func (f *MockFileLib) File() io.Reader {
	err := f.FileError
	f.FileError = f.FileError2
	if err == nil {
		return f.IoReader
	}
	return err
}

//
// MockJSONLIB
//
type MockJSONLib struct {
	IndentErr error
	ValidErr  *bool
	Jsonlib   ngsilib.JSONLib
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

func setJSONDecodeErr(ngsi *ngsilib.NGSI, p int) {
	j := ngsi.JSONConverter
	mockj := &MockJSONLib{Jsonlib: j}
	mockj.DecodeErr[p] = errors.New("json error")
	ngsi.JSONConverter = mockj
}

func setJSONEncodeErr(ngsi *ngsilib.NGSI, p int) {
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

func setJSONIndentError(ngsi *ngsilib.NGSI) {
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{IndentErr: errors.New("json error"), Jsonlib: j}
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

type MockNetLib struct {
	AddrErr              error
	ListenAndServeErr    error
	ListenAndServeTLSErr error
}

func (n *MockNetLib) InterfaceAddrs() ([]net.Addr, error) {
	if n.AddrErr != nil {
		return nil, n.AddrErr
	}
	return net.InterfaceAddrs()
}

func (n *MockNetLib) ListenAndServe(addr string, handler http.Handler) error {
	return n.ListenAndServeErr
}
func (n *MockNetLib) ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	return n.ListenAndServeTLSErr
}
