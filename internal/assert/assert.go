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

package assert

import (
	"reflect"
	"strconv"
	"strings"
)

type TestingT interface {
	Errorf(format string, args ...interface{})
}

func NoError(t TestingT, err error, msgAndArgs ...interface{}) bool {
	if err == nil {
		return true
	}

	t.Errorf("\nReceived unexpected error:\n%s", err.Error())
	return false
}

func Error(t TestingT, err error, msgAndArgs ...interface{}) bool {
	if err != nil {
		return true
	}

	t.Errorf("\nAn error is expected but got nil.")
	return false
}

func Equal(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	if reflect.DeepEqual(expected, actual) {
		return true
	}

	if exp, ok := expected.(string); ok {
		if act, ok := actual.(string); ok {
			t.Errorf("\nNot equal")
			t.Errorf("\nexpected: \"%s\"", escapeString(exp))
			t.Errorf("\nactual  : \"%s\"", escapeString(act))
		}
		return false
	}

	t.Errorf("\nShould not be: %s", getString(expected))
	return false
}

func NotEqual(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	if !reflect.DeepEqual(expected, actual) {
		return true
	}

	t.Errorf("\nShould be: %s", getString(expected))

	return false
}

func escapeString(s string) string {
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

func getString(value interface{}) string {
	switch v := value.(type) {
	default:
		return "error"
	case nil:
		return "nil"
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case string:
		return v
	}
}
