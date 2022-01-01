/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"errors"
	"fmt"
	"testing"
)

type helperTesting struct {
	err string
}

func (h *helperTesting) Errorf(format string, args ...interface{}) {
	h.err += fmt.Sprintf(format, args...)
}

func TestNoErrorTrue(t *testing.T) {
	h := &helperTesting{}

	actual := NoError(h, nil)

	if !actual {
		t.Error("True is expected but false")
	}
}

func TestNoErrorFalse(t *testing.T) {
	h := &helperTesting{}

	actual := NoError(h, errors.New("error"))

	if actual {
		t.Error("False is expected but true")
	}

	if h.err != "\nReceived unexpected error:\nerror" {
		t.Error(h.err)
	}
}

func TestErrorTrue(t *testing.T) {
	h := &helperTesting{}

	actual := Error(h, errors.New("error"))

	if !actual {
		t.Error("True is expected but false")
	}
}

func TestErrorFalse(t *testing.T) {
	h := &helperTesting{}

	actual := Error(h, nil)

	if actual {
		t.Error("False is expected but true")
	}

	if h.err != "\nAn error is expected but got nil." {
		t.Error(h.err)
	}
}

func TestEqualTrue(t *testing.T) {
	tests := []struct {
		expected interface{}
		actual   interface{}
	}{
		{nil, nil},
		{1, 1},
		{int64(64), int64(64)},
		{"test", "test"},
	}

	for _, test := range tests {
		actual := Equal(t, test.expected, test.actual)

		if !actual {
			t.Error("True is expected but false")
		}
	}
}

func TestEqualFalse(t *testing.T) {
	tests := []struct {
		expected interface{}
		actual   interface{}
		msg      string
	}{
		{nil, 1, "\nShould not be: nil"},
		{1, 2, "\nShould not be: 1"},
		{int64(64), int64(32), "\nShould not be: 64"},
	}

	for _, test := range tests {
		h := &helperTesting{}

		actual := Equal(h, test.expected, test.actual)

		if actual {
			t.Error("False is expected but true")
		}

		if h.err != test.msg {
			t.Error(h.err)
		}
	}
}

func TestEqualFalseString(t *testing.T) {
	tests := []struct {
		expected interface{}
		actual   interface{}
		msg      string
	}{
		{"test1", "test2", "\nNot equal\nexpected: \"test1\"\nactual  : \"test2\""},
		{"\"test3", "test4", "\nNot equal\nexpected: \"\\\"test3\"\nactual  : \"test4\""},
		{"\ntest5", "test6", "\nNot equal\nexpected: \"\\ntest5\"\nactual  : \"test6\""},
		{"\ttest7", "test8", "\nNot equal\nexpected: \"\\ttest7\"\nactual  : \"test8\""},
	}

	for _, test := range tests {
		h := &helperTesting{}

		actual := Equal(h, test.expected, test.actual)

		if actual {
			t.Error("False is expected but true")
		}

		if h.err != test.msg {
			t.Error(h.err)
		}
	}
}

func TestNotEqualTrue(t *testing.T) {
	tests := []struct {
		expected interface{}
		actual   interface{}
	}{
		{nil, 1},
		{1, 2},
		{int64(64), int64(32)},
		{"test", "test2"},
	}

	for _, test := range tests {
		actual := NotEqual(t, test.expected, test.actual)

		if !actual {
			t.Error("True is expected but false")
		}
	}
}

func TestNotEqualFalse(t *testing.T) {
	tests := []struct {
		expected interface{}
		actual   interface{}
		msg      string
	}{
		{nil, nil, "\nShould be: nil"},
		{1, 1, "\nShould be: 1"},
		{int64(64), int64(64), "\nShould be: 64"},
		{"test", "test", "\nShould be: test"},
	}

	for _, test := range tests {
		h := &helperTesting{}

		actual := NotEqual(h, test.expected, test.actual)

		if actual {
			t.Error("False is expected but true")
		}

		if h.err != test.msg {
			t.Error(h.err)
		}
	}
}

func TestEscapeString(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"test", "test"},
		{"\n\ntest", "\\n\\ntest"},
		{"\t\ttest", "\\t\\ttest"},
		{"\"\"test", "\\\"\\\"test"},
	}

	for _, test := range tests {
		actual := escapeString(test.value)

		if test.expected != actual {
			t.Error("Should be " + test.expected)
		}
	}
}

func TestGetString(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{nil, "nil"},
		{1, "1"},
		{int64(64), "64"},
		{"test", "test"},
		{1.1, "error"},
	}

	for _, test := range tests {
		actual := getString(test.value)

		if test.expected != actual {
			t.Error("Should be " + test.expected)
		}
	}
}
