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

package helper

import (
	"bytes"
	"errors"
	"io"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

type MockJSONLib struct {
	IndentErr error
	ValidErr  *bool
	Jsonlib   ngsilib.JSONLib
	DecodeErr [5]error
	EncodeErr [5]error
	dp        int
	ep        int
}

func (j *MockJSONLib) Decode(r io.Reader, v interface{}) error {
	err := j.DecodeErr[j.dp]
	j.dp++
	if err == nil {
		return j.Jsonlib.Decode(r, v)
	}
	return err
}

func (j *MockJSONLib) Encode(w io.Writer, v interface{}) error {
	err := j.EncodeErr[j.ep]
	j.ep++
	if err == nil {
		return j.Jsonlib.Encode(w, v)
	}
	return err
}

func SetJSONDecodeErr(ngsi *ngsilib.NGSI, p int) {
	j := ngsi.JSONConverter
	mockj := &MockJSONLib{Jsonlib: j}
	mockj.DecodeErr[p] = errors.New("json error")
	ngsi.JSONConverter = mockj
}

func SetJSONEncodeErr(ngsi *ngsilib.NGSI, p int) {
	j := ngsi.JSONConverter
	mockj := &MockJSONLib{Jsonlib: j}
	mockj.EncodeErr[p] = errors.New("json error")
	ngsi.JSONConverter = mockj
}

func (j *MockJSONLib) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	if j.IndentErr == nil {
		return j.Jsonlib.Indent(dst, src, prefix, indent)
	}
	return j.IndentErr
}

func SetJSONIndentError(ngsi *ngsilib.NGSI) {
	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{IndentErr: errors.New("json error"), Jsonlib: j}
}

func (j *MockJSONLib) Valid(data []byte) bool {
	if j.ValidErr != nil {
		return *j.ValidErr
	}
	return j.Jsonlib.Valid(data)
}
