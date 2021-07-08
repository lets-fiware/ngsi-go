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

package ngsilib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// NGSI is ...
type NGSI struct {
	configVresion string
	serverList    ServerList
	tokenList     tokenInfoList
	contextList   ContextsInfo

	LogLevel      int
	ConfigFile    IoLib
	CacheFile     IoLib
	StdReader     io.Reader
	StdWriter     io.Writer
	LogWriter     io.Writer
	FileReader    FileLib
	JSONConverter JSONLib
	FilePath      FilePathLib
	Ioutil        IoutilLib
	ZipLib        ZipLib
	MultiPart     MultiPart
	Host          string
	Destination   string
	Margin        int64
	Maxsize       int
	Timeout       time.Duration
	PreviousArgs  *Settings
	Updated       bool
	HTTP          HTTPRequest
	Stderr        io.Writer
	OsType        string
	SyslogLib     SyslogLib
	TimeLib       TimeLib
	BatchFlag     *bool
}

// CmdFlags is ...
type CmdFlags struct {
	Token      *string
	Tenant     *string
	Scope      *string
	SafeString *string
	XAuthToken bool
	Link       *string
}

var gNGSI *NGSI

// NewNGSI is ...
func NewNGSI() *NGSI {
	if gNGSI == nil {
		gNGSI = &NGSI{}
		gNGSI.configVresion = "1"
		gNGSI.InitLog(os.Stdin, os.Stdout, os.Stderr)
		gNGSI.HTTP = &httpRequest{}
		gNGSI.Margin = 180
		gNGSI.Timeout = 60
		gNGSI.Maxsize = 100
		gNGSI.ConfigFile = &ioLib{}
		gNGSI.CacheFile = &ioLib{}
		gNGSI.JSONConverter = &jsonLib{}
		gNGSI.FileReader = &fileLib{}
		gNGSI.FilePath = &filePathLib{}
		gNGSI.Ioutil = &ioutilLib{}
		gNGSI.ZipLib = &zipLib{}
		gNGSI.MultiPart = &multiPart{}
		gNGSI.Stderr = os.Stderr
		gNGSI.OsType = runtime.GOOS
		gNGSI.SyslogLib = &syslogLib{}
		gNGSI.PreviousArgs = &Settings{UsePreviousArgs: true}
		gNGSI.TimeLib = &timeLib{}
		gNGSI.serverList = make(ServerList)
		gNGSI.contextList = make(ContextsInfo)
		gNGSI.contextList["etsi1.0"] = "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"
		gNGSI.contextList["etsi1.3"] = "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"
		gNGSI.contextList["etsi1.4"] = "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.4.jsonld"
		gNGSI.contextList["ld"] = "https://schema.lab.fiware.org/ld/context"
	}
	return gNGSI
}

// Close is ...
func (ngsi *NGSI) Close() {
	gNGSI = nil
}

// Reset is ...
func Reset() {
	gNGSI = nil
}

// InitLog is ...
func (ngsi *NGSI) InitLog(stdin io.Reader, stdout, stderr io.Writer) *NGSI {
	ngsi.StdReader = stdin
	ngsi.StdWriter = stdout
	ngsi.Stderr = stderr
	ngsi.LogWriter = &LogWriter{stderr, LogErr}
	ngsi.LogLevel = LogErr
	return ngsi
}

// BoolFlag is ...
func (ngsi *NGSI) BoolFlag(s string) (bool, error) {
	const funcName = "BoolFlag"

	switch strings.ToLower(s) {
	case "", "off", "false":
		return false, nil
	case "on", "true":
		return true, nil
	}
	return false, &LibError{funcName, 1, fmt.Sprintf("unknown parameter: %s", s), nil}
}

// StdoutFlush is ...
func (ngsi *NGSI) StdoutFlush() {
	out, ok := ngsi.StdWriter.(*bufio.Writer)
	if ok {
		_ = out.Flush()
	}
}

func getConfigDir(io IoLib) (string, error) {
	const funcName = "getConfigDir"

	var path string
	home, err := io.UserHomeDir()
	if err != nil {
		return "", &LibError{funcName, 1, err.Error(), err}
	}
	if gNGSI.OsType == "windows" {
		path = io.Getenv("APPDATA")
		if path == "" {
			path = home
		}
	} else {
		path, err = io.UserConfigDir()
		if err != nil {
			path = io.FilePathJoin(home, ".config")
		}
	}
	home = io.FilePathJoin(path, "fiware")
	if !existsFile(io, home) {
		err := io.MkdirAll(home, 0700)
		if err != nil {
			return "", &LibError{funcName, 2, err.Error(), err}
		}
	}
	return home, nil
}

func existsFile(io IoLib, filename string) bool {
	if filename == "" {
		return false
	}
	_, err := io.Stat(filename)
	return err == nil
}
