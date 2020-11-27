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

package ngsilib

import (
	"strings"
)

// https://fiware-orion.readthedocs.io/en/master/user/forbidden_characters/index.html

var forbiddenCharsEncode = map[string]string{
	`"`: "%22",
	"'": "%27",
	"(": "%28",
	")": "%29",
	";": "%3B",
	"<": "%3C",
	"=": "%3D",
	">": "%3E",
}

var forbiddenCharsDecode = map[string]string{
	"%22": `"`,
	"%25": `%`,
	"%27": "'",
	"%28": "(",
	"%29": ")",
	"%3b": ";",
	"%3B": ";",
	"%3c": "<",
	"%3C": "<",
	"%3d": "=",
	"%3D": "=",
	"%3e": ">",
	"%3E": ">",
}

// SafeStringEncode is ...
func SafeStringEncode(s string) string {
	s = strings.ReplaceAll(s, "%", "%25")
	for k, v := range forbiddenCharsEncode {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

// SafeStringDecode is ...
func SafeStringDecode(s string) string {
	for k, v := range forbiddenCharsDecode {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

// JSONSafeStringEncode is ...
func JSONSafeStringEncode(data []byte) ([]byte, error) {
	const funcName = "JSONSafeStringEncode"

	var e interface{}

	if err := JSONUnmarshalEncode(data, &e, true); err != nil {
		return nil, &NgsiLibError{funcName, 1, err.Error(), err}
	}

	bytes, err := JSONMarshalDecode(&e, false)
	if err != nil {
		return nil, &NgsiLibError{funcName, 2, err.Error(), err}
	}
	return bytes, nil
}

// JSONSafeStringDecode is ...
func JSONSafeStringDecode(data []byte) ([]byte, error) {
	const funcName = "JSONSafeStringDecode"

	if IsJSON(data) {
		var e interface{}

		if err := JSONUnmarshalEncode(data, &e, false); err != nil {
			return nil, &NgsiLibError{funcName, 1, err.Error(), err}
		}

		bytes, err := JSONMarshalDecode(&e, true)
		if err != nil {
			return nil, &NgsiLibError{funcName, 2, err.Error(), err}
		}
		return bytes, nil
	}
	s := SafeStringDecode(string(data))
	return []byte(s), nil
}
