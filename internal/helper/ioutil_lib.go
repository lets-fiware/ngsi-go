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
	"io"
	"io/ioutil"
	"os"
)

type MockIoutilLib struct {
	CopyErr      error
	ReadFullErr  error
	ReadFullData []byte
	WriteFileErr error
	ReadFileErr  error
	ReadFileData []byte
	WriteSkip    bool
}

func (i *MockIoutilLib) Copy(dst io.Writer, src io.Reader) (int64, error) {
	if i.CopyErr != nil {
		return 0, i.CopyErr
	}
	return io.Copy(dst, src)
}

func (i *MockIoutilLib) ReadFull(r io.Reader, buf []byte) (n int, err error) {
	if i.ReadFullErr != nil {
		return 0, i.ReadFullErr
	}
	if i.ReadFullData != nil {
		for i, v := range i.ReadFullData {
			buf[i] = v
		}
		return len(buf), nil
	}
	return io.ReadFull(r, buf)
}

func (i *MockIoutilLib) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if i.WriteSkip {
		return nil
	}
	if i.WriteFileErr != nil {
		return i.WriteFileErr
	}
	return ioutil.WriteFile(filename, data, perm)
}

func (i *MockIoutilLib) ReadFile(filename string) ([]byte, error) {
	if i.ReadFileErr != nil {
		return nil, i.ReadFileErr
	}
	if i.ReadFileData != nil {
		return i.ReadFileData, nil
	}
	return ioutil.ReadFile(filename)
}
