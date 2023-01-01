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

package helper

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestMultiPartLibCreatePart(t *testing.T) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "application/octet-stream")
	mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+"file"+"\"")

	m := &MockMultiPartLib{Mw: w}

	actual, err := m.CreatePart(mh)

	if assert.NoError(t, err) {
		assert.NotEqual(t, (io.Writer)(nil), actual)
	}
}

func TestMultiPartLibCreatePartError(t *testing.T) {
	m := &MockMultiPartLib{CreatePartErr: errors.New("CreatePart error")}

	actual, err := m.CreatePart(nil)

	if assert.Error(t, err) {
		assert.Equal(t, (io.Writer)(nil), actual)
		assert.Equal(t, "CreatePart error", err.Error())
	}
}

func TestMultiPartLibFormDataContentType(t *testing.T) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "application/octet-stream")
	mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+"file"+"\"")

	m := &MockMultiPartLib{Mw: w}

	_, err := m.CreatePart(mh)

	assert.NoError(t, err)

	s := m.FormDataContentType()

	assert.NotEqual(t, "", s)

}

func TestMultiPartLibClose(t *testing.T) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "application/octet-stream")
	mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+"file"+"\"")

	m := &MockMultiPartLib{Mw: w}

	_, err := m.CreatePart(mh)

	assert.NoError(t, err)

	err = m.Close()

	assert.NoError(t, err)
}

func TestMultiPartLibCloseError(t *testing.T) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "application/octet-stream")
	mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+"file"+"\"")

	m := &MockMultiPartLib{Mw: w, CloseErr: errors.New("Close error")}

	_, err := m.CreatePart(mh)

	assert.NoError(t, err)

	err = m.Close()

	if assert.Error(t, err) {
		assert.Equal(t, "Close error", err.Error())
	}
}
