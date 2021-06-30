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
	"io/ioutil"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.Open()

	assert.Error(t, err)
}

func TestOpenFile(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)

	assert.Error(t, err)
}
func TestTruncate(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	err = io.Truncate(0)

	assert.Error(t, err)
}

func TestFileClose(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	err = io.Close()

	assert.Error(t, err)
}

func TestNewDecoder(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	var i int
	err = io.Decode(&i)

	assert.Error(t, err)
}

func TestNewEncoder(t *testing.T) {
	fileName := "???"
	io := ioLib{fileName: &fileName}
	err := io.OpenFile(oWRONLY, 0600)
	assert.Error(t, err)

	var i int
	err = io.Encode(&i)

	assert.Error(t, err)
}

func TestUserHomeDir(t *testing.T) {
	io := ioLib{}

	_, err := io.UserHomeDir()

	assert.NoError(t, err)
}

func TestUserConfigDir(t *testing.T) {
	io := ioLib{}

	_, err := io.UserConfigDir()

	assert.NoError(t, err)
}

func TestMkdirAll(t *testing.T) {
	io := ioLib{}

	err := io.MkdirAll("", 0700)

	if assert.Error(t, err) {
		assert.Equal(t, "mkdir : no such file or directory", err.Error())
	}
}

func TestStat(t *testing.T) {
	io := ioLib{}

	_, err := io.Stat("")

	if assert.Error(t, err) {
		assert.Equal(t, "stat : no such file or directory", err.Error())
	}
}

func TestSetFileName(t *testing.T) {
	io := ioLib{}

	s := "test"
	io.SetFileName(&s)

	assert.Equal(t, s, *io.fileName)
}

func TestFileName(t *testing.T) {
	io := ioLib{}

	s := "test"
	io.SetFileName(&s)

	f := io.FileName()

	assert.Equal(t, s, *f)
}

func TestGetenv(t *testing.T) {
	io := ioLib{}

	e := io.Getenv("")

	assert.Equal(t, "", e)
}

func TestFilePathAbs(t *testing.T) {
	io := ioLib{}

	_, err := io.FilePathAbs("")

	assert.NoError(t, err)
}

func TestFilePathJoin(t *testing.T) {
	io := ioLib{}

	s := io.FilePathJoin("/", "home")

	assert.Equal(t, "/home", s)
}

func TestFileLibOpen(t *testing.T) {
	f := &fileLib{}
	err := f.Open("???")

	assert.Error(t, err)
}

func TestFileLibClose(t *testing.T) {
	f := &fileLib{}
	err := f.Open(".")
	assert.NoError(t, err)

	err = f.Close()

	assert.Error(t, err)
}

func TestFileLibCloseNil(t *testing.T) {
	f := &fileLib{}
	f.file = nil
	err := f.Close()

	assert.NoError(t, err)
}

func TestFileLibCloseError(t *testing.T) {
	f := &fileLib{}
	err := f.Open("???")
	assert.Error(t, err)

	err = f.Close()

	assert.NoError(t, err)
}

func TestFileLibFilePathAbs(t *testing.T) {
	f := &fileLib{}
	_, err := f.FilePathAbs("")

	assert.NoError(t, err)
}

func TestFileLibReadAll(t *testing.T) {
	f := &fileLib{}
	buf := &bytes.Buffer{}
	_, err := f.ReadAll(buf)

	assert.NoError(t, err)
}

func TestFileLibReadFile(t *testing.T) {
	f := &fileLib{}
	_, err := f.ReadFile("???")

	assert.Error(t, err)
}

func TestFileLibSetReader(t *testing.T) {
	f := &fileLib{}
	buf := &bytes.Buffer{}
	f.SetReader(buf)

	reader := f.File()
	_, err := reader.Peek(1)

	if assert.Error(t, err) {
		assert.Equal(t, "EOF", err.Error())
	}
}

func TestFileLibFile(t *testing.T) {
	f := &fileLib{}
	buf := &bytes.Buffer{}
	f.SetReader(buf)

	file := f.File()
	_, err := file.Peek(1)

	if assert.Error(t, err) {
		assert.Equal(t, "EOF", err.Error())
	}
}

func TestJSONLibDecode(t *testing.T) {
	j := &jsonLib{}
	buf := &bytes.Buffer{}
	var i int
	err := j.Decode(buf, &i)

	if assert.Error(t, err) {
		assert.Equal(t, "EOF", err.Error())
	}
}

func TestJSONLibEncode(t *testing.T) {
	j := &jsonLib{}
	buf := &bytes.Buffer{}
	var i int
	err := j.Encode(buf, &i)

	assert.NoError(t, err)
}

func TestIndent(t *testing.T) {
	j := &jsonLib{}
	buf := &bytes.Buffer{}
	src := []byte("{}")
	err := j.Indent(buf, src, "", "  ")

	assert.NoError(t, err)
}

func TestValid(t *testing.T) {
	j := &jsonLib{}
	src := []byte("{}")
	actual := j.Valid(src)

	expected := true

	assert.Equal(t, expected, actual)
}

func TestTimeLibNow(t *testing.T) {
	time := &timeLib{}

	_ = time.Now()
}

func TestTimeLibNowUnix(t *testing.T) {
	time := &timeLib{}

	_ = time.NowUnix()
}

func TestTimeLibUnix(t *testing.T) {
	time := &timeLib{}

	_ = time.Unix(0, 0)
}

func TestTimeLibFormat(t *testing.T) {
	time := &timeLib{}

	_ = time.Unix(0, 0)
	_ = time.Format("2021/06/27 06:47:39")
}

func TestIoutilLibCopy(t *testing.T) {
	iolib := ioutilLib{}
	src := bytes.NewReader([]byte("FIWARE"))

	_, _ = iolib.Copy(ioutil.Discard, src)
}

func TestIoutilLibReadFull(t *testing.T) {
	iolib := ioutilLib{}
	src := bytes.NewReader([]byte("FIWARE"))
	buf := make([]byte, 10)

	_, _ = iolib.ReadFull(src, buf)
}

func TestIoutilLibWriteFile(t *testing.T) {
	iolib := ioutilLib{}
	buf := make([]byte, 10)

	_ = iolib.WriteFile("", buf, 0644)
}

func TestIoutilLibReadFile(t *testing.T) {
	iolib := ioutilLib{}

	_, _ = iolib.ReadFile("")
}

func TestFilePathLibFilePathAbs(t *testing.T) {
	f := filePathLib{}

	_, _ = f.FilePathAbs("")
}

func TestFilePathLibFilePathJoin(t *testing.T) {
	f := filePathLib{}

	_ = f.FilePathJoin("", "")
}

func TestFilePathLibFilePathBase(t *testing.T) {
	f := filePathLib{}

	_ = f.FilePathBase("")
}

func TestZipNewReader(t *testing.T) {
	z := zipLib{}
	src := bytes.NewReader([]byte(""))

	_, _ = z.NewReader(src, -1)
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
