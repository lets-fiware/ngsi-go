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

package ngsilib

import (
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestGetAtContext(t *testing.T) {
	ngsi := testNgsiLibInit()

	_, err := ngsi.GetAtContext("{}")

	assert.NoError(t, err)
}

func TestGetAtContextHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	actual, err := ngsi.GetAtContext("etsi1.3")

	if assert.NoError(t, err) {
		assert.Equal(t, "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld", actual)
	}
}

func TestGetAtContextErrorNotFound(t *testing.T) {
	ngsi := testNgsiLibInit()

	_, err := ngsi.GetAtContext("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestGetAtContextErrorNotJSON(t *testing.T) {
	ngsi := testNgsiLibInit()

	_, err := ngsi.GetAtContext("1")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "data not json: 1", ngsiErr.Message)
	}
}

func TestGetAtContextErrorJSON(t *testing.T) {
	ngsi := testNgsiLibInit()

	SetJSONDecodeErr(ngsi, 0)

	_, err := ngsi.GetAtContext("{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContext(t *testing.T) {
	ngsi := testNgsiLibInit()

	payload := []byte(`{"id":"I"}`)

	cases := []struct {
		context  string
		expected string
	}{
		{
			context:  `https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld`,
			expected: "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"id\":\"I\"}",
		},
		{
			context:  "[\"http://example.org/ngsi-ld/latest/parking.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]",
			expected: "{\"@context\":[\"http://example.org/ngsi-ld/latest/parking.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"],\"id\":\"I\"}",
		},
		{
			context:  "{\"parking\":\"http://example.org/ngsi-ld/latest/parking.jsonld\"}",
			expected: "{\"@context\":{\"parking\":\"http://example.org/ngsi-ld/latest/parking.jsonld\"},\"id\":\"I\"}",
		},
		{
			context:  "etsi1.3",
			expected: "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"id\":\"I\"}",
		},
	}

	for _, ca := range cases {
		actual, err := ngsi.InsertAtContext(payload, ca.context)

		if assert.NoError(t, err) {
			assert.Equal(t, ca.expected, string(actual))
		}
	}
}

func TestInsertAtContextAtContext(t *testing.T) {
	ngsi := testNgsiLibInit()

	SetJSONEncodeErr(ngsi, 0)
	payload := []byte(`[{"@context":""}]`)

	_, err := ngsi.InsertAtContext(payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorPayload(t *testing.T) {
	ngsi := testNgsiLibInit()

	payload := []byte(`context`)

	_, err := ngsi.InsertAtContext(payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data not json", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorGetAtContext(t *testing.T) {
	ngsi := testNgsiLibInit()

	SetJSONDecodeErr(ngsi, 1)
	payload := []byte(`[]`)

	_, err := ngsi.InsertAtContext(payload, "{")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorArrayUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	SetJSONDecodeErr(ngsi, 1)
	payload := []byte(`[]`)

	_, err := ngsi.InsertAtContext(payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorArrayMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	SetJSONEncodeErr(ngsi, 0)
	payload := []byte(`[]`)

	_, err := ngsi.InsertAtContext(payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorObjectUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	SetJSONDecodeErr(ngsi, 1)
	payload := []byte(`{}`)

	_, err := ngsi.InsertAtContext(payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorObjectMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	SetJSONEncodeErr(ngsi, 0)

	payload := []byte(`{}`)

	_, err := ngsi.InsertAtContext(payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
