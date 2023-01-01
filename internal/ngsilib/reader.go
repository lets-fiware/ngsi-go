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

package ngsilib

import (
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

type ReadAllFunc func(s string) (bytes []byte, err error)
type GetReaderFunc func(s string) (FileLib, error)

func ReadAll(s string) (bytes []byte, err error) {
	const funcName = "readAll"

	ngsi := gNGSI

	fileReader := ngsi.FileReader

	if s == "" {
		return nil, ngsierr.New(funcName, 1, "data is empty", nil)
	}
	if s == "stdin" || s == "@-" { // from stdin
		bytes, err = fileReader.ReadAll(ngsi.StdReader)
		if err != nil {
			return nil, ngsierr.New(funcName, 2, err.Error(), err)
		}
	} else if strings.HasPrefix(s, "@") { // from file
		if len(s) > 1 {
			path, err := fileReader.FilePathAbs(s[1:])
			if err != nil {
				return nil, ngsierr.New(funcName, 3, err.Error(), err)
			}
			bytes, err = fileReader.ReadFile(path)
			if err != nil {
				return nil, ngsierr.New(funcName, 4, err.Error(), err)
			}
			return bytes, nil
		}
		return nil, ngsierr.New(funcName, 5, "file name error", nil)
	} else {
		bytes = []byte(s)
	}

	return
}

func GetReader(s string) (FileLib, error) {
	const funcName = "getReader"

	ngsi := gNGSI

	fileReader := ngsi.FileReader

	if s == "" {
		return nil, ngsierr.New(funcName, 1, "data is empty", nil)
	}
	if s == "stdin" || s == "@-" { // from stdin
		fileReader.SetReader(ngsi.StdReader)
		return fileReader, nil
	} else if strings.HasPrefix(s, "@") { // from file
		if len(s) > 1 {
			path, err := fileReader.FilePathAbs(s[1:])
			if err != nil {
				return nil, ngsierr.New(funcName, 2, err.Error(), err)
			}
			err = fileReader.Open(path)
			if err != nil {
				return nil, ngsierr.New(funcName, 3, err.Error(), err)
			}
			return fileReader, nil
		}
		return nil, ngsierr.New(funcName, 4, "file name error", nil)
	}
	fileReader.SetReader(strings.NewReader(s))
	return fileReader, nil
}
