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

package ngsicli

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
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
	  },
	  "wirecloud": {
		"serverHost": "https://wirecloud",
		"serverType": "wirecloud"
	  },
	  "scorpio": {
		"serverHost": "https://scorpio:9090",
		"ngsiType": "ld",
		"serverType": "broker",
		"brokerType": "scorpio"
	  },
	  "regproxy": {
		"serverHost": "https://regproxy",
		"serverType": "regproxy"
	  },
	  "tokenproxy": {
		"serverHost": "https://tokenproxy",
		"serverType": "tokenproxy"
	  },
	  "queryproxy": {
		"serverHost": "https://queryproxy",
		"serverType": "queryproxy"
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

var mockNGSI *ngsilib.NGSI

func setupTestInitNGSI() *ngsilib.NGSI {
	ngsilib.Reset()

	filename := ""
	ngsi := ngsilib.NewNGSI()

	mockNGSI = ngsi

	ngsi.ConfigFile = &MockIoLib{}
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(configData)}

	ngsi.CacheFile = &MockIoLib{}
	ngsi.CacheFile.SetFileName(&filename)

	ngsi.HTTP = NewMockHTTP()
	buffer := &bytes.Buffer{}

	ngsi.StdWriter = buffer
	ngsi.LogWriter = &bytes.Buffer{}

	ngsi.FilePath = &MockFilePathLib{}
	ngsi.Ioutil = &MockIoutilLib{}
	ngsi.ZipLib = &MockZipLib{}

	ngsi.ReadAll = mockReadAll
	ngsi.GetReader = mockGetReader

	ngsi.PreviousArgs = &ngsilib.Settings{}

	return ngsi
}

func setupTestInitCmd() *Context {
	ngsi := setupTestInitNGSI()

	c := &Context{Ngsi: ngsi}
	_, err := InitCmd(c)
	if err != nil {
		panic(err)
	}

	return c
}

func NewMockHTTP() *MockHTTP {
	m := MockHTTP{}
	return &m
}

// MockReader
func mockReadAll(s string) (bytes []byte, err error) {
	return ngsilib.ReadAll(s)
}

func mockGetReader(s string) (ngsilib.FileLib, error) {
	return ngsilib.GetReader(s)
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
	RawQuery   *string
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
			return nil, nil, ngsierr.New(funcName, 0, "Unsupported type", nil)
		}
	}
	if data != nil && r.ReqData != nil {
		if !reflect.DeepEqual(r.ReqData, data) {
			fmt.Printf("r.ReqData: %s\n", string(r.ReqData))
			fmt.Printf("Data:      %s\n", string(data))
			return nil, nil, ngsierr.New(funcName, 1, "body data error", nil)
		}
	}
	if r.Path != "" && r.Path != url.Path {
		return nil, nil, ngsierr.New(funcName, 3, "url error", nil)
	}
	if r.RawQuery != nil {
		if *r.RawQuery != url.RawQuery {
			return nil, nil, ngsierr.New(funcName, 4, "raw query error: "+url.RawQuery, nil)
		}
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
	OpenErr    error
	EncodeErr  error
	filename   *string
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
	io.filename = filename
}

func (io *MockIoLib) FileName() *string {
	return io.filename
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

//
// MockJSONLib
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
	format   *string
	tTime    *time.Time
}

func (t *MockTimeLib) Now() time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	tm, _ := time.Parse(layout, t.dateTime)
	return tm
}

func (t *MockTimeLib) NowUnix() int64 {
	return t.unixTime + time.Now().Unix()
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

//
//  MockIoutilLib
//
type MockIoutilLib struct {
	CopyErr      error
	ReadFullErr  error
	ReadFullData []byte
	WriteFileErr error
	ReadFileErr  error
	ReadFileData []byte
	WriteSkip    bool
}

func (i *MockIoutilLib) Copy(dst io.Writer, src io.Reader) (int64, error) {
	if i.CopyErr != nil {
		return 0, i.CopyErr
	}
	return io.Copy(dst, src)
}

func (i *MockIoutilLib) ReadFull(r io.Reader, buf []byte) (n int, err error) {
	if i.ReadFullErr != nil {
		return 0, i.ReadFullErr
	}
	if i.ReadFullData != nil {
		for i, v := range i.ReadFullData {
			buf[i] = v
		}
		return len(buf), nil
	}
	return io.ReadFull(r, buf)
}

func (i *MockIoutilLib) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if i.WriteSkip {
		return nil
	}
	if i.WriteFileErr != nil {
		return i.WriteFileErr
	}
	return ioutil.WriteFile(filename, data, perm)
}

func (i *MockIoutilLib) ReadFile(filename string) ([]byte, error) {
	if i.ReadFileErr != nil {
		return nil, i.ReadFileErr
	}
	if i.ReadFileData != nil {
		return i.ReadFileData, nil
	}
	return ioutil.ReadFile(filename)
}

//
// MockFilePathLib
//
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

//
// MockZipLib
//
type MockZipLib struct {
	Zip       error
	ZipReader *zip.Reader
}

func (z *MockZipLib) NewReader(r io.ReaderAt, size int64) (*zip.Reader, error) {
	if z.Zip != nil {
		return nil, z.Zip
	}
	if z.ZipReader != nil {
		return z.ZipReader, nil
	}
	return zip.NewReader(r, size)
}

type MockMultiPart struct {
	CreatePartErr error
	CloseErr      error
}

func (m *MockMultiPart) NewWriter(w io.Writer) ngsilib.MultiPartLib {
	return &MockMultiPartLib{Mw: multipart.NewWriter(w), CreatePartErr: m.CreatePartErr, CloseErr: m.CloseErr}
}

//
// MockMultiPartLib
//
type MockMultiPartLib struct {
	CreatePartErr error
	CloseErr      error
	Mw            *multipart.Writer
}

func (m MockMultiPartLib) CreatePart(header textproto.MIMEHeader) (io.Writer, error) {
	if m.CreatePartErr != nil {
		return nil, m.CreatePartErr
	}
	return m.Mw.CreatePart(header)
}

func (m MockMultiPartLib) FormDataContentType() string {
	return m.Mw.FormDataContentType()
}

func (m MockMultiPartLib) Close() error {
	if m.CloseErr != nil {
		return m.CloseErr
	}
	return m.Mw.Close()
}

//
// MockZipFile
//
type MockZipFile struct {
}

func (f *MockZipFile) DataOffset() (offset int64, err error) {
	return 0, nil

}

func (f *MockZipFile) Open() (io.ReadCloser, error) {
	return nil, nil
}

// Flag
var (
	syslogFlag = &StringFlag{
		Name:  "syslog",
		Usage: "specify logging `LEVEL` (off, err, info, debug)",
	}
	stderrFlag = &StringFlag{
		Name:  "stderr",
		Usage: "specify logging `LEVEL` (off, err, info, debug)",
	}
	configFlag = &StringFlag{
		Name:  "config",
		Usage: "specify configuration `FILE`",
	}
	cacheFlag = &StringFlag{
		Name:  "cache",
		Usage: "specify cache `FILE`",
	}
	marginFlag = &Int64Flag{
		Name:   "margin",
		Usage:  "I/O time out (second)",
		Hidden: true,
		Value:  180,
	}
	timeOutFlag = &Int64Flag{
		Name:   "timeout",
		Usage:  "I/O time out (second)",
		Hidden: true,
		Value:  60,
	}
	maxCountFlag = &Int64Flag{
		Name:   "maxCount",
		Usage:  "maxCount",
		Hidden: true,
		Value:  100,
	}
	batchFlag = &BoolFlag{
		Name:    "batch",
		Aliases: []string{"B"},
		Usage:   "don't use previous args (batch)",
	}
	insecureSkipVerifyFlag = &BoolFlag{
		Name:  "insecureSkipVerify",
		Usage: "TLS/SSL skip certificate verification",
	}
	hostFlag = &StringFlag{
		Name:    "host",
		Usage:   "broker or server host `VALUE`",
		Aliases: []string{"h"},
	}
	tokenFlag = &StringFlag{
		Name:  "oAuthToken",
		Usage: "oauth token `VALUE`",
	}
	tenantFlag = &StringFlag{
		Name:    "service",
		Aliases: []string{"s"},
		Usage:   "FIWARE Service `VALUE`",
	}
	scopeFlag = &StringFlag{
		Name:    "path",
		Aliases: []string{"p"},
		Usage:   "FIWARE ServicePath `VALUE`",
	}
	linkFlag = &StringFlag{
		Name:    "link",
		Aliases: []string{"L"},
		Usage:   "@context `VALUE` (LD)",
	}
	destinationFlag = &StringFlag{
		Name:     "host2",
		Aliases:  []string{"d"},
		Usage:    "host or alias",
		Value:    "",
		Required: true,
	}
	token2Flag = &StringFlag{
		Name:  "oAuthToken2",
		Usage: "oauth token for destination",
	}
	tenant2Flag = &StringFlag{
		Name:  "service2",
		Usage: "FIWARE Service for destination",
	}
	scope2Flag = &StringFlag{
		Name:  "path2",
		Usage: "FIWARE ServicePath for destination",
	}
	link2Flag = &StringFlag{
		Name:  "link2",
		Usage: "@context (LD)",
	}
	xAuthTokenFlag = &BoolFlag{
		Name:   "xAuthToken",
		Usage:  "use X-Auth-Token",
		Hidden: true,
	}
	safeStringFlag = &StringFlag{
		Name:  "safeString",
		Usage: "use safe string (`VALUE`: on/off)",
	}
	queryFlag = &StringFlag{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "filtering by attribute value",
	}
	servicesDeviceFlag = &BoolFlag{
		Name:  "device",
		Usage: "remove devices in service/subservice",
		Value: false,
	}
)
