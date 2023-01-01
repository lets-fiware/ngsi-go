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

package ngsierr

import (
	"errors"
	"syscall"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestNgsiErrorNew(t *testing.T) {
	err := New("ngsierr", 1, "error message", nil)

	if assert.Error(t, err) {
		assert.Equal(t, "ngsierr", err.Function)
		assert.Equal(t, 1, err.ErrNo)
		assert.Equal(t, "error message", err.Message)
		assert.Equal(t, nil, err.Err)
	}
}

func TestErrorMessage(t *testing.T) {
	e := NgsiError{"test", 1, "message", nil}

	actual := e.Error()
	expected := "message"

	assert.Equal(t, expected, actual)
	assert.Error(t, &e)
}

func TestErrorString(t *testing.T) {
	e := NgsiError{"test", 1, "message", nil}

	actual := e.String()
	expected := "test001 message"

	assert.Equal(t, expected, actual)
	assert.Error(t, &e)
}

func TestErrorStringSysError(t *testing.T) {
	sys := syscall.Errno(1)
	e := NgsiError{"test", 1, "message", sys}

	actual := e.String()
	expected := "test001 message: operation not permitted"

	assert.Equal(t, expected, actual)
	assert.Error(t, &e)
}

func TestErrorUnwap(t *testing.T) {
	expected := syscall.Errno(1)
	e := NgsiError{"test", 1, "message", expected}

	actual := e.Unwrap()

	assert.Equal(t, expected, actual)
	assert.Error(t, &e)
}

func TestSprintMsg(t *testing.T) {
	expected := "NGSI-GO001 error"
	actual := SprintMsg("NGSI-GO", 1, "error")

	assert.Equal(t, expected, actual)
}

func TestMessage(t *testing.T) {
	e := errors.New("error message")
	s := Message(e)

	assert.Equal(t, "error message", s)
}

func TestMessageCmdError(t *testing.T) {
	e := New("Test", 1, "error message", nil)
	s := Message(e)

	assert.Equal(t, "Test001 error message", s)
}
