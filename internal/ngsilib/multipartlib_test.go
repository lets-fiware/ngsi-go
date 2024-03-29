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

package ngsilib

import (
	"bytes"
	"net/textproto"
	"os"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestMultiPart(t *testing.T) {
	mp := &multiPart{}

	w := mp.NewWriter(os.Stdout)

	assert.NotEqual(t, nil, w)
}

func TestMultiPartLibNewWriter(t *testing.T) {
	var body bytes.Buffer

	mp := &multiPart{}
	m := mp.NewWriter(&body)
	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "application/octet-stream")
	mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\"test\"")
	_, err := m.CreatePart(mh)
	assert.NoError(t, err)
	_ = m.FormDataContentType()
	err = m.Close()
	assert.NoError(t, err)
}
