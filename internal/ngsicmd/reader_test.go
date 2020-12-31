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

package ngsicmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestReadAll(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{`--data={"id":test"}`})

	b, err := readAll(c, ngsi)

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(`{"id":test"}`), b)
	} else {
		t.FailNow()
	}
}

func TestReadAllStdReader(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=stdin"})
	ngsi.FileReader = &MockFileLib{ReadallData: []byte("test data")}

	b, err := readAll(c, ngsi)

	if assert.NoError(t, err) {
		assert.Equal(t, []byte("test data"), b)
	} else {
		t.FailNow()
	}
}

func TestReadAllAt(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@file"})
	ngsi.FileReader = &MockFileLib{FilePathAbsString: "file", ReadFileData: []byte(`{"id":test"}`)}

	b, err := readAll(c, ngsi)

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(`{"id":test"}`), b)
	} else {
		t.FailNow()
	}
}

func TestReadAllErrorEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	c := cli.NewContext(app, set, nil)

	_, err := readAll(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestReadAllErrorStdReader(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=stdin"})
	ngsi.FileReader = &MockFileLib{ReadallError: errors.New("ReadAll error")}

	_, err := readAll(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "ReadAll error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestReadAllAt3(t *testing.T) {
	ngsi, set, app, _ := setupTest()

	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@file"})
	setFilePatAbsError(ngsi, 0)

	_, err := readAll(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "filepathabs error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestReadAllAt4(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@file"})
	setReadFileError(ngsi, 0)

	_, err := readAll(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "readfile error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestReadAllErrorAt5(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@"})

	_, err := readAll(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetReader(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data={abc}"})

	_, err := getReader(c, ngsi)

	assert.NoError(t, err)
}

func TestGetReaderFIle(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@file"})
	ngsi.FileReader = &MockFileLib{FilePathAbsString: "file", ReadFileData: []byte(`{"id":test"}`)}

	_, err := getReader(c, ngsi)

	assert.NoError(t, err)
}

func TestGetReaderStdin(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=stdin"})
	ngsi.FileReader = &MockFileLib{Name: "stdin test"}

	f, err := getReader(c, ngsi)

	if assert.NoError(t, err) {
		m := f.(*MockFileLib)
		assert.Equal(t, "stdin test", m.Name)
	} else {
		t.FailNow()
	}
}

func TestGetReaderErrorEmpty(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	c := cli.NewContext(app, set, nil)

	_, err := getReader(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data is empty", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetReaderAt3(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@file"})
	setFilePatAbsError(ngsi, 0)

	_, err := getReader(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "filepathabs error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetReaderAt4(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@file"})
	ngsi.FileReader = &MockFileLib{FilePathAbsString: "file", OpenError: errors.New("error @file")}

	_, err := getReader(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "error @file", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestGetReaderErrorAt5(t *testing.T) {
	ngsi, set, app, _ := setupTest()


	setupFlagString(set, "data")
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--data=@"})

	_, err := getReader(c, ngsi)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
