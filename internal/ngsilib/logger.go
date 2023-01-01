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

package ngsilib

import (
	"fmt"
	"io"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

const (
	// LogOff is ...
	LogOff = iota
	// LogErr is ...
	LogErr
	// LogWarn is ...
	LogWarn
	// LogInfo is ...
	LogInfo
	// LogDebug is ...
	LogDebug
	// LogTrace is ...
	LogTrace
)

// LogWriter is ...
type LogWriter struct {
	Writer   io.Writer
	LogLevel int
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	if w.LogLevel >= logLevel {
		return w.Writer.Write(p)
	}
	return len(p), nil
}

// var logger io.Writer = os.Stderr
var logLevel int = LogErr

// Logging is ...
func (ngsi *NGSI) Logging(level int, s string) {
	save := logLevel
	logLevel = level
	_, _ = (*ngsi).LogWriter.Write([]byte(s))
	logLevel = save
}

// LogLevel is ...
func LogLevel(s string) (int, error) {
	const funcName = "LogLevel"

	s = strings.ToLower(s)
	if strings.HasPrefix(s, "log") && len(s) > 3 {
		s = s[3:]
	}
	switch s {
	case "off":
		return LogOff, nil
	case "err":
		return LogErr, nil
	case "info":
		return LogInfo, nil
	case "debug":
		return LogDebug, nil
	}
	return -1, ngsierr.New(funcName, 1, fmt.Sprintf("unknown LogLevel: %s", s), nil)
}
