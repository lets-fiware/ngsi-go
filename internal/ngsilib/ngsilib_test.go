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
	"bufio"
	"errors"
	"os"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestNGSI(t *testing.T) {
	ngsi := testNgsiLibInit()

	assert.Equal(t, ngsi, gNGSI)
}

func TestClose(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.Close()

	assert.Equal(t, (*NGSI)((nil)), gNGSI)
}

func TestReset(t *testing.T) {
	testNgsiLibInit()

	Reset()

	assert.Equal(t, (*NGSI)((nil)), gNGSI)
}

func TestInitLog(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.InitLog(os.Stdin, os.Stdout, os.Stderr)

	assert.Equal(t, os.Stdin, ngsi.StdReader)
	assert.Equal(t, os.Stdout, ngsi.StdWriter)
	assert.Equal(t, LogErr, ngsi.LogLevel)
}

func TestBoolFlag(t *testing.T) {
	ngsi := testNgsiLibInit()

	b, err := ngsi.BoolFlag("")
	if assert.NoError(t, err) {
		assert.Equal(t, false, b)
	}
}

func TestBoolFlagOff(t *testing.T) {
	ngsi := testNgsiLibInit()

	b, err := ngsi.BoolFlag("OFF")
	if assert.NoError(t, err) {
		assert.Equal(t, false, b)
	}
}

func TestBoolFalse(t *testing.T) {
	ngsi := testNgsiLibInit()

	b, err := ngsi.BoolFlag("False")
	if assert.NoError(t, err) {
		assert.Equal(t, false, b)
	}
}

func TestBoolON(t *testing.T) {
	ngsi := testNgsiLibInit()

	b, err := ngsi.BoolFlag("ON")
	if assert.NoError(t, err) {
		assert.Equal(t, true, b)
	}
}

func TestBoolTrue(t *testing.T) {
	ngsi := testNgsiLibInit()

	b, err := ngsi.BoolFlag("True")
	if assert.NoError(t, err) {
		assert.Equal(t, true, b)
	}
}

func TestBoolError(t *testing.T) {
	ngsi := testNgsiLibInit()

	b, err := ngsi.BoolFlag("fiware")
	if assert.Error(t, err) {
		assert.Equal(t, false, b)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: fiware", ngsiErr.Message)
	}
}

func TestStdoutFlush(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.StdWriter = bufio.NewWriter(os.Stdout)

	ngsi.StdoutFlush()

}

func TestGetConfigDir(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config"}

	dir, err := getConfigDir(io)

	if assert.NoError(t, err) {
		assert.Equal(t, "/home/ngsi/.config/fiware", dir)
	}
}

func TestGetConfigDirGetenv(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.OsType = "windows"
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config", Env: `C:\Users\`}

	dir, err := getConfigDir(io)

	if assert.NoError(t, err) {
		assert.Equal(t, "C:\\Users\\/fiware", dir)
	}
}

func TestGetConfigDirGetenvNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.OsType = "windows"
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config", Env: ""}

	dir, err := getConfigDir(io)

	if assert.NoError(t, err) {
		assert.Equal(t, "/fiware", dir)
	}
}

func TestGetConfigDirNotExists(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config", StatErr: errors.New("error")}

	dir, err := getConfigDir(io)

	if assert.NoError(t, err) {
		assert.Equal(t, "/home/ngsi/.config/fiware", dir)
	}
}

func TestGetConfigDirErrorUserConfigDir(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{HomeDir: "/home/ngsi", ConfigDir: "/home/ngsi/.config", ConfigDirErr: errors.New("error")}

	dir, err := getConfigDir(io)

	if assert.NoError(t, err) {
		assert.Equal(t, "/home/ngsi/.config/fiware", dir)
	}
}

func TestGetConfigDirErrorUserHomeDir(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config", HomeDirErr: errors.New("error homedir")}

	_, err := getConfigDir(io)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error homedir", ngsiErr.Message)
	}
}

func TestGetConfigDirErrorMkdir(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config", StatErr: errors.New("error stat"), MkdirErr: errors.New("error mkdir")}

	_, err := getConfigDir(io)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error mkdir", ngsiErr.Message)
	}
}

func TestExistsFileEmpty(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config", StatErr: errors.New("error")}

	b := existsFile(io, "")

	assert.Equal(t, false, b)
}

func TestExistsFileExists(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config"}

	b := existsFile(io, "fiware")

	assert.Equal(t, true, b)
}

func TestExistsFileNotExists(t *testing.T) {
	testNgsiLibInit()
	io := &MockIoLib{ConfigDir: "/home/ngsi/.config", StatErr: errors.New("error")}

	b := existsFile(io, "fiware")

	assert.Equal(t, false, b)
}
