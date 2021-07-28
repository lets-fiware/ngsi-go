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
	"fmt"
	"sort"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// ContextsInfo is ...
type ContextsInfo map[string]interface{}

// AddContext is ...
func (ngsi *NGSI) AddContext(key string, value string) error {
	const funcName = "AddContext"

	if _, ok := ngsi.contextList[key]; ok {
		return ngsierr.New(funcName, 1, fmt.Sprintf("%s already exists", key), nil)
	}
	if IsHTTP(value) {
		ngsi.contextList[key] = value
	} else {
		var json interface{}
		err := JSONUnmarshal([]byte(value), &json)
		if err != nil {
			return ngsierr.New(funcName, 2, fmt.Sprintf("%s is neither url nor json", value), nil)
		}
		ngsi.contextList[key] = json
	}

	if err := ngsi.saveConfigFile(); err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	return nil
}

// UpdateContext is ...
func (ngsi *NGSI) UpdateContext(key string, value string) error {
	const funcName = "UpdateContext"

	if _, ok := ngsi.contextList[key]; ok {
		if !IsHTTP(value) {
			return ngsierr.New(funcName, 1, fmt.Sprintf("%s is not url", value), nil)
		}
		ngsi.contextList[key] = value
		if err := ngsi.saveConfigFile(); err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		return nil
	}
	return ngsierr.New(funcName, 3, fmt.Sprintf("%s not found", key), nil)
}

// DeleteContext is ...
func (ngsi *NGSI) DeleteContext(key string) error {
	const funcName = "DeleteContext"

	if err := ngsi.IsContextReferenced(key); err != nil {
		return ngsierr.New(funcName, 1, key+" is referenced", err)
	}
	if _, ok := ngsi.contextList[key]; ok {
		delete(ngsi.contextList, key)
		if err := ngsi.saveConfigFile(); err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		return nil
	}
	return ngsierr.New(funcName, 3, fmt.Sprintf("%s not found", key), nil)
}

// GetContext is ...
func (ngsi *NGSI) GetContext(key string) (string, error) {
	const funcName = "GetContext"

	if IsHTTP(key) {
		return key, nil
	}
	if _, ok := ngsi.contextList[key]; ok {
		value := ngsi.contextList[key]
		switch value := value.(type) {
		default:
			return "", ngsierr.New(funcName, 1, fmt.Sprintf("%s neither url nor json", key), nil)
		case string:
			if IsHTTP(value) {
				return value, nil
			}
			return "", ngsierr.New(funcName, 2, fmt.Sprintf("%s is not url", key), nil)
		case []interface{}, map[string]interface{}:
			b, err := JSONMarshal(value)
			if err != nil {
				return "", ngsierr.New(funcName, 3, err.Error(), err)
			}
			return string(b), nil
		}
	}
	return "", ngsierr.New(funcName, 4, fmt.Sprintf("%s not found", key), nil)
}

// GetContextHTTP is ...
func (ngsi *NGSI) GetContextHTTP(key string) (string, error) {
	const funcName = "GetContextHTTP"

	s, err := ngsi.GetContext(key)
	if err != nil {
		return "", ngsierr.New(funcName, 1, err.Error(), err)
	}
	if IsHTTP(s) {
		return s, nil
	}
	return "", ngsierr.New(funcName, 2, fmt.Sprintf("%s is not url", key), nil)
}

// GetContextList is ...
func (ngsi *NGSI) GetContextList() ContextsInfo {
	slice := make([]string, len(ngsi.contextList))
	index := 0
	for key := range ngsi.contextList {
		slice[index] = key
		index++
	}
	sort.Strings(slice)
	return ngsi.contextList
}
