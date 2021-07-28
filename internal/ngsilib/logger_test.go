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
	"bytes"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/stretchr/testify/assert"
)

func TestLogWrite(t *testing.T) {
	l := &LogWriter{Writer: &bytes.Buffer{}, LogLevel: LogErr}

	logLevel = LogDebug
	n, err := l.Write([]byte("test"))

	if assert.NoError(t, err) {
		assert.Equal(t, 4, n)
	}
}

func TestLogWriteLogOff(t *testing.T) {
	l := &LogWriter{Writer: &bytes.Buffer{}, LogLevel: LogErr}

	logLevel = LogOff
	n, err := l.Write([]byte("test"))

	if assert.NoError(t, err) {
		assert.Equal(t, 4, n)
	}
}

func TestLogging(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.LogWriter = &bytes.Buffer{}

	ngsi.Logging(LogDebug, "test")
}

func TestLogLevelOff(t *testing.T) {

	actual, err := LogLevel("off")

	if assert.NoError(t, err) {
		assert.Equal(t, LogOff, actual)
	}
}

func TestLogLevelErr(t *testing.T) {

	actual, err := LogLevel("ERR")

	if assert.NoError(t, err) {
		assert.Equal(t, LogErr, actual)
	}
}

func TestLogLevelInfo(t *testing.T) {

	actual, err := LogLevel("info")

	if assert.NoError(t, err) {
		assert.Equal(t, LogInfo, actual)
	}
}

func TestLogLevelDebug(t *testing.T) {

	actual, err := LogLevel("debug")

	if assert.NoError(t, err) {
		assert.Equal(t, LogDebug, actual)
	}
}

func TestLogLevelLogDebug(t *testing.T) {

	actual, err := LogLevel("LOGDEBUG")

	if assert.NoError(t, err) {
		assert.Equal(t, LogDebug, actual)
	}
}

func TestLogLevelError(t *testing.T) {

	actual, err := LogLevel("fiware")

	if assert.Error(t, err) {
		assert.Equal(t, -1, actual)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown LogLevel: fiware", ngsiErr.Message)
	}
}
