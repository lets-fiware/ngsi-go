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

package ngsicmd

import (
	"fmt"
	"io"
)

type buffer interface {
	bufferOpen(w io.Writer)
	bufferWrite(interface{})
	bufferClose()
}

type jsonBuffer struct {
	writer    io.Writer
	buf       []byte
	delimiter string
}

func (j *jsonBuffer) bufferOpen(w io.Writer) {
	j.writer = w
	j.delimiter = "["
}

func (j *jsonBuffer) bufferWrite(b []byte) {
	if len(j.buf) > 0 {
		fmt.Fprint(j.writer, j.delimiter)
		fmt.Fprint(j.writer, string(j.buf))
		j.delimiter = ","
	}
	if len(b) > 0 {
		j.buf = b[1 : len(b)-1]
	} else {
		j.buf = nil
	}
}

func (j *jsonBuffer) bufferClose() {
	if len(j.buf) > 0 {
		j.bufferWrite(nil)
	}
	if j.delimiter != "[" {
		fmt.Fprint(j.writer, "]")
	}
}
