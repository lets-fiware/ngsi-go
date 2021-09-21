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

import "github.com/lets-fiware/ngsi-go/internal/ngsierr"

func (ngsi *NGSI) InsertAtContext(payload []byte, context string) (b []byte, err error) {
	const funcName = "insertAtContext"

	if !IsJSON(payload) {
		return nil, ngsierr.New(funcName, 1, "data not json", nil)
	}

	var atContext interface{}
	atContext, err = ngsi.GetAtContext(context)
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}

	if IsJSONArray(payload) { // JSON Array
		var entities NgsiEntities
		err = JSONUnmarshal(payload, &entities)
		if err != nil {
			return nil, ngsierr.New(funcName, 3, err.Error(), err)
		}
		for _, e := range entities {
			e["@context"] = atContext
		}
		b, err = JSONMarshal(entities)
		if err != nil {
			return nil, ngsierr.New(funcName, 4, err.Error(), err)
		}
	} else { // JSON Object
		var e NgsiEntity
		err = JSONUnmarshal(payload, &e)
		if err != nil {
			return nil, ngsierr.New(funcName, 5, err.Error(), err)
		}
		e["@context"] = atContext
		b, err = JSONMarshal(e)
		if err != nil {
			return nil, ngsierr.New(funcName, 6, err.Error(), err)
		}
	}

	return b, nil
}

func (ngsi *NGSI) GetAtContext(context string) (interface{}, error) {
	const funcName = "getAtContext"

	var atContext interface{}

	if IsNameString(context) {
		value, err := ngsi.GetContext(context)
		if err != nil {
			return nil, ngsierr.New(funcName, 1, err.Error(), err)
		}
		context = value
	}

	if !IsJSON([]byte(context)) {
		if IsHTTP(context) {
			context = `"` + context + `"`
		} else {
			return nil, ngsierr.New(funcName, 2, "data not json: "+context, nil)
		}
	}

	err := JSONUnmarshal([]byte(context), &atContext)
	if err != nil {
		return nil, ngsierr.New(funcName, 3, err.Error(), err)
	}

	return atContext, nil
}
