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
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// IsJSON is ...
func IsJSON(b []byte) bool {
	for i := 0; i < len(b); i++ {
		if string(b[i]) == " " || string(b[i]) == "\t" {
			continue
		}
		if string(b[i]) == "{" || string(b[i]) == "[" {
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

	if safeString {
		convert(v, f)
	}

	buffer := &bytes.Buffer{}
	err := gNGSI.JSONConverter.Encode(buffer, v)
	if err != nil {
		return nil, &NgsiLibError{funcName, 1, err.Error(), err}
	}
	b := buffer.Bytes()
	if b[len(b)-1] == 0xa {
		b = b[:len(b)-1]
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

	err := gNGSI.JSONConverter.Decode(bytes.NewReader(data), &v)
	if err != nil {
		if err, ok := err.(*json.SyntaxError); ok {
			s := err.Offset - 15
			if s < 0 {
				s = 0
			}
			e := err.Offset + 15
			if e > int64(len(data)) {
				e = int64(len(data))
			}
			return &NgsiLibError{funcName, 1, fmt.Sprintf("%s (%d) %s", err.Error(), err.Offset, string(data[s:e])), err}
		}
		return &NgsiLibError{funcName, 2, err.Error(), err}
	}
	if safeString {
		if reflect.ValueOf(v).Kind() != reflect.Ptr {
			return &NgsiLibError{funcName, 3, "non-pointer", nil}
		}
		convert(v, f)
	}
	return nil
}

// convert is ...
func convert(e interface{}, f func(string) string) {
	const funcName = "convert"

	if (e == nil) || reflect.ValueOf(e).IsNil() {
		return
	}
	v := reflect.ValueOf(e)
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(reflect.ValueOf(e))
	}

	t := v.Type()

	switch t.Kind() {
	case reflect.Interface:
		switch reflect.ValueOf(v.Interface()).Kind() {
		case reflect.Map:
			convert(v.Interface().(map[string]interface{}), f)
		case reflect.Slice:
			convert(v.Interface().([]interface{}), f)
		}
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			ft := t.Field(i)
			fv := v.FieldByName(ft.Name)
			switch ft.Type.Kind() {
			case reflect.Struct, reflect.Map, reflect.Slice:
				convert(fv.Addr().Interface(), f)
			case reflect.Ptr:
				switch ft.Type.Elem().Kind() {
				case reflect.Struct, reflect.Map, reflect.Slice:
					convert(fv.Interface(), f)
				}
			case reflect.String:
				fv.SetString(f(fv.String()))
			}
		}
	case reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Struct, reflect.Map, reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				e := v.Index(i)
				convert(e.Addr().Interface(), f)
			}
		case reflect.String:
			for i := 0; i < v.Len(); i++ {
				e := v.Index(i)
				s := f(e.String())
				v.Index(i).SetString(s)
			}
		case reflect.Interface:
			for i := 0; i < v.Len(); i++ {
				e := v.Index(i)
				switch reflect.ValueOf(e.Interface()).Kind() {
				case reflect.Map:
					convert(e.Interface().(map[string]interface{}), f)
				case reflect.String:
					s := f(e.Interface().(string))
					v.Index(i).Set(reflect.ValueOf(s))
				}
			}
		}
	case reflect.Map:
		kt := v.Type().Key()
		if kt.Kind() == reflect.String {
			iter := v.MapRange()
			for iter.Next() {
				key := iter.Key().String()
				mapv := iter.Value()
				switch mapv.Kind() {
				case reflect.Struct, reflect.Map, reflect.Slice:
					convert(mapv.Interface(), f)
				case reflect.String:
					s := f(mapv.String())
					v.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(s))
				case reflect.Interface:
					switch reflect.ValueOf(mapv.Interface()).Kind() {
					case reflect.String:
						s := f(mapv.Interface().(string))
						v.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(s))
					case reflect.Struct, reflect.Map, reflect.Slice:
						convert(mapv.Interface(), f)
					}
				}
			}
		}
	}
}
