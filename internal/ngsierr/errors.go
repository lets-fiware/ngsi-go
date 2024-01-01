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

package ngsierr

import (
	"errors"
	"fmt"
	"syscall"
)

type NgsiError struct {
	Function string
	ErrNo    int
	Message  string
	Err      error
}

func New(function string, errNo int, message string, err error) *NgsiError {
	return &NgsiError{Function: function, ErrNo: errNo, Message: message, Err: err}
}

func (e *NgsiError) String() string {
	var errno syscall.Errno
	var s string
	if errors.As(e, &errno) {
		s = fmt.Sprintf(": %s", errno)
	}
	return fmt.Sprintf("%s%03d %s%s", e.Function, e.ErrNo, e.Message, s)
}

func (e *NgsiError) Error() string {
	return e.Message
}

func (e *NgsiError) Unwrap() error { return e.Err }

func SprintMsg(funcName string, no int, msg string) string {
	return fmt.Sprintf("%s%03d %s", funcName, no, msg)
}

func Message(err error) (s string) {
	switch e := err.(type) {
	case *NgsiError:
		s = e.String()
	default:
		s = e.Error()
	}
	return
}
