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
	"encoding/json"
	"fmt"
)

// IsJSON is ...
func IsJSON(b []byte) bool {
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' || b[i] == '\t' {
			continue
		}
		if b[i] == '{' || b[i] == '[' {
			return true
		}
		break
	}
	return false
}

// IsJSONArray is ...
func IsJSONArray(b []byte) bool {
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' || b[i] == '\t' {
			continue
		}
		if b[i] == '[' {
			return true
		}
		break
	}
	return false
}

// GetJSONArray is ...
func GetJSONArray(b []byte, v interface{}) error {
	const funcName = "GetJSONArray"

	for i := 0; i < len(b); i++ {
		if string(b[i]) == " " || string(b[i]) == "\t" {
			continue
		}
		if string(b[i]) == "[" {
			err := JSONUnmarshal(b, v)
			if err != nil {
				return &NgsiLibError{funcName, 1, err.Error(), err}
			}
			return nil
		}
		break
	}
	return &NgsiLibError{funcName, 2, "not JSON Array:" + string(b), nil}
}

// JSONMarshal is ...
func JSONMarshal(v interface{}) ([]byte, error) {
	return jsonMarshal(v, false, nil)
}

// JSONMarshalEncode is ...
func JSONMarshalEncode(v interface{}, safeString bool) ([]byte, error) {
	return jsonMarshal(v, safeString, SafeStringEncode)
}

// JSONMarshalDecode is ...
func JSONMarshalDecode(v interface{}, safeString bool) ([]byte, error) {
	return jsonMarshal(v, safeString, SafeStringDecode)
}

// JSONMarshal is ...
func jsonMarshal(v interface{}, safeString bool, f func(string) string) ([]byte, error) {
	const funcName = "jsonMarshal"

	buffer := &bytes.Buffer{}
	err := gNGSI.JSONConverter.Encode(buffer, v)
	if err != nil {
		return nil, &NgsiLibError{funcName, 1, err.Error(), err}
	}
	b := buffer.Bytes()
	if b[len(b)-1] == 0xa {
		b = b[:len(b)-1]
	}

	if safeString {
		b, _ = jsonParser(b, f)
	}

	return b, nil
}

// JSONUnmarshal is ...
func JSONUnmarshal(data []byte, v interface{}) error {
	return jsonUnmarshal(data, v, false, nil)
}

// JSONUnmarshalEncode is ...
func JSONUnmarshalEncode(data []byte, v interface{}, safeString bool) error {
	return jsonUnmarshal(data, v, safeString, SafeStringEncode)
}

// JSONUnmarshalDecode is ...
func JSONUnmarshalDecode(data []byte, v interface{}, safeString bool) error {

	return jsonUnmarshal(data, v, safeString, SafeStringDecode)
}

// JSONUnmarshalEncode is ...
func jsonUnmarshal(data []byte, v interface{}, safeString bool, f func(string) string) error {
	const funcName = "jsonUnmarshal"

	if safeString {
		var err error
		data, err = jsonParser(data, f)
		if err != nil {
			return &NgsiLibError{funcName, 1, err.Error(), err}
		}

	}

	err := gNGSI.JSONConverter.Decode(bytes.NewReader(data), &v)
	if err != nil {
		switch err := err.(type) {
		case *json.SyntaxError:
			s := err.Offset - 15
			if s < 0 {
				s = 0
			}
			e := err.Offset + 15
			if e > int64(len(data)) {
				e = int64(len(data))
			}
			return &NgsiLibError{funcName, 2, fmt.Sprintf("%s (%d) %s", err.Error(), err.Offset, string(data[s:e])), err}
		case *json.UnmarshalTypeError:
			s := err.Offset - 15
			if s < 0 {
				s = 0
			}
			e := err.Offset + 15
			if e > int64(len(data)) {
				e = int64(len(data))
			}
			return &NgsiLibError{funcName, 3, fmt.Sprintf("%s Field:%s (%d) %s", err.Error(), err.Field, err.Offset, string(data[s:e])), err}
		}
		return &NgsiLibError{funcName, 4, err.Error(), err}
	}

	return nil
}
