/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeStringEncode(t *testing.T) {
	testNgsiLibInit()

	actual := SafeStringEncode((`<>"'=;()%`))
	expected := "%3C%3E%22%27%3D%3B%28%29%25"

	assert.Equal(t, expected, actual)

	actual = SafeStringEncode(("%%%%%"))
	expected = "%25%25%25%25%25"

	assert.Equal(t, expected, actual)

	actual = SafeStringEncode((`123abc<>"'=;()%`))
	expected = "123abc%3C%3E%22%27%3D%3B%28%29%25"

	assert.Equal(t, expected, actual)
}

func TestSafeStringDecode(t *testing.T) {
	testNgsiLibInit()

	actual := SafeStringDecode(("%3C%3E%22%27%3D%3B%28%29%25"))
	expected := `<>"'=;()%`

	assert.Equal(t, expected, actual)

	actual = SafeStringDecode(("%3c%3e%22%27%3d%3b%28%29%25"))
	expected = `<>"'=;()%`

	assert.Equal(t, expected, actual)

	actual = SafeStringDecode(("%25%25%25%25%25"))
	expected = "%%%%%"

	assert.Equal(t, expected, actual)

	actual = SafeStringDecode(("123abc%3c%3e%22%27%3d%3b%28%29%25"))
	expected = `123abc<>"'=;()%`

	assert.Equal(t, expected, actual)

	actual = SafeStringDecode(("123abc%3C%3E%22%27%3D%3B%28%29%25"))
	expected = `123abc<>"'=;()%`

	assert.Equal(t, expected, actual)
}

func TestJSONSafeStringEncode(t *testing.T) {
	testNgsiLibInit()

	input := `{"type": "<>\"'=;()%"}`
	expected := `{"type":"%3C%3E%22%27%3D%3B%28%29%25"}`

	actual, err := JSONSafeStringEncode([]byte(input))

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(expected), actual)
	} else {
		t.FailNow()
	}

	input = `[{"type": "<>\"'=;()%"}]`
	expected = `[{"type":"%3C%3E%22%27%3D%3B%28%29%25"}]`

	actual, err = JSONSafeStringEncode([]byte(input))

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(expected), actual)
	} else {
		t.FailNow()
	}

	input = `{"id":"abc","type":"%<>\"'=;()","speed":{"type":"Text","value":"xyz()"},"array":{"type":"=struct%%","value":[1,"<abc>"]}}`
	expected = `{"array":{"type":"%3Dstruct%25%25","value":[1,"%3Cabc%3E"]},"id":"abc","speed":{"type":"Text","value":"xyz%28%29"},"type":"%25%3C%3E%22%27%3D%3B%28%29"}`

	actual, err = JSONSafeStringEncode([]byte(input))

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(expected), actual)
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringErrorEncode1(t *testing.T) {
	ngsi := testNgsiLibInit()

	j := gNGSI.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{Jsonlib: j, DecodeErr: errors.New("json error")}

	input := `{"id":"abc","type":"%<>\"'=;()","speed":{"type":"Text","value":"xyz()"},"array":{"type":"=struct%%","value":[1,"<abc>"]}}`
	_, err := JSONSafeStringEncode([]byte(input))

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringErrorEncode2(t *testing.T) {
	ngsi := testNgsiLibInit()

	j := gNGSI.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	input := `{"id":"abc","type":"%<>\"'=;()","speed":{"type":"Text","value":"xyz()"},"array":{"type":"=struct%%","value":[1,"<abc>"]}}`
	_, err := JSONSafeStringEncode([]byte(input))

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
func TestJSONSafeStringDecodeString(t *testing.T) {
	testNgsiLibInit()

	actual, err := JSONSafeStringDecode([]byte("%25"))
	expected := "%"

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(expected), actual)
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringDecodeJSON(t *testing.T) {
	testNgsiLibInit()

	actual, err := JSONSafeStringDecode([]byte(`{"name":"%25"}`))
	expected := `{"name":"%"}`

	if assert.NoError(t, err) {
		assert.Equal(t, []byte(expected), actual)
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringDecodeErrorMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}

	_, err := JSONSafeStringDecode([]byte(`{"name":"%25"}`))

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringDecodeErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	_, err := JSONSafeStringDecode([]byte(`{"name":"%25"}`))

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}
