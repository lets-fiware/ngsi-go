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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeStringEncode(t *testing.T) {
	testNgsiLibInit()

	cases := []struct {
		testdata string
		expected string
	}{
		{
			testdata: `<>"'=;()%`,
			expected: "%3C%3E%22%27%3D%3B%28%29%25",
		},
		{
			testdata: "%%%%%",
			expected: "%25%25%25%25%25",
		},
		{
			testdata: `123abc<>"'=;()%`,
			expected: "123abc%3C%3E%22%27%3D%3B%28%29%25",
		},
	}
	for _, c := range cases {
		actual := SafeStringEncode(c.testdata)
		assert.Equal(t, c.expected, actual)
	}
}

func TestSafeStringDecode(t *testing.T) {
	testNgsiLibInit()

	cases := []struct {
		testdata string
		expected string
	}{
		{
			testdata: "%3C%3E%22%27%3D%3B%28%29%25",
			expected: `<>\"'=;()%`,
		},
		{

			testdata: "%3c%3e%22%27%3d%3b%28%29%25",
			expected: `<>\"'=;()%`,
		},
		{
			testdata: "%25%25%25%25%25",
			expected: "%%%%%",
		},
		{
			testdata: "123abc%3c%3e%22%27%3d%3b%28%29%25",
			expected: `123abc<>\"'=;()%`,
		},
		{
			testdata: "123abc%3C%3E%22%27%3D%3B%28%29%25",
			expected: `123abc<>\"'=;()%`,
		},
		{
			testdata: "123abc%3C%3E%22%27%3D%3B%28%29%25%11",
			expected: `123abc<>\"'=;()%%11`,
		},
		{
			testdata: "123abc%3C%3E%22%27%3D%3B%28%29%25%a",
			expected: `123abc<>\"'=;()%%a`,
		},
		{
			testdata: "123abc%3C%3E%22%27%3D%3B%28%29%25%",
			expected: `123abc<>\"'=;()%%`,
		},
	}
	for _, c := range cases {
		actual := SafeStringDecode(c.testdata)
		assert.Equal(t, c.expected, actual)
	}
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
	expected = `{"id":"abc","type":"%25%3C%3E%22%27%3D%3B%28%29","speed":{"type":"Text","value":"xyz%28%29"},"array":{"type":"%3Dstruct%25%25","value":[1,"%3Cabc%3E"]}}`

	actual, err = JSONSafeStringEncode([]byte(input))

	if assert.NoError(t, err) {
		assert.Equal(t, string([]byte(expected)), string(actual))
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringErrorEncode1(t *testing.T) {
	testNgsiLibInit()

	input := `{"id":"abc","type:"%<>\"'=;()","speed":{"type":"Text","value":"xyz()"},"array":{"type":"=struct%%","value":[1,"<abc>"]}}`
	_, err := JSONSafeStringEncode([]byte(input))

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '%' after object key (19) \":\"abc\",\"type:\"%<>\\\"'=;()\",\"sp", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringErrorEncode2(t *testing.T) {
	testNgsiLibInit()

	input := `{"id":"abc","type":"%<>\"'=;()","speed:{"type":"Text","value":"xyz()"},"array":{"type":"=struct%%","value":[1,"<abc>"]}}`
	_, err := JSONSafeStringEncode([]byte(input))

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 't' after object key (41) =;()\",\"speed:{\"type\":\"Text\",\"v", ngsiErr.Message)
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
	testNgsiLibInit()

	_, err := JSONSafeStringDecode([]byte(`{"name":%25"}`))

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character '%' looking for beginning of value (7) {\"name\":%25\"}", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestJSONSafeStringDecodeErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	j := ngsi.JSONConverter
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), Jsonlib: j}

	_, err := JSONSafeStringDecode([]byte(`{"name":"%25}`))

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func testStringFunc(s string) string {
	return s
}
func TestJsonParser(t *testing.T) {
	cases := []struct {
		data     string
		expected string
	}{
		{data: `1`, expected: `1`},
		{data: `null`, expected: `null`},
		{data: `{}`, expected: `{}`},
		{data: `[]`, expected: `[]`},
		{data: `[1,2,3]`, expected: `[1,2,3]`},
		{data: `["a","b","c"]`, expected: `["a","b","c"]`},
		{data: `[{},{},{}]`, expected: `[{},{},{}]`},
		{data: `{"array":[{},{},{}]}`, expected: `{"array":[{},{},{}]}`},
		{
			data:     `{"a":12.01,"Null":null,"id":"5fd412e8ecb082767349b975","subject":{"entities":[{"idPattern":".*"},{}],"condition":{}},"notification":{"timesSent":17,"lastNotification":"2020-12-12T06:16:11.000Z","lastSuccess":"2020-12-12T06:16:11.000Z","lastSuccessCode":204,"onlyChangedAttrs":false,"http":{"url":"http://172.22.143.188:1028/"},"attrsFormat":"normalized"},"status":"active"}`,
			expected: `{"a":12.01,"Null":null,"id":"5fd412e8ecb082767349b975","subject":{"entities":[{"idPattern":".*"},{}],"condition":{}},"notification":{"timesSent":17,"lastNotification":"2020-12-12T06:16:11.000Z","lastSuccess":"2020-12-12T06:16:11.000Z","lastSuccessCode":204,"onlyChangedAttrs":false,"http":{"url":"http://172.22.143.188:1028/"},"attrsFormat":"normalized"},"status":"active"}`,
		},
	}

	for _, c := range cases {
		actual, err := jsonParser([]byte(c.data), testStringFunc)
		if assert.NoError(t, err) {
			assert.Equal(t, c.expected, string(actual))
		} else {
			t.FailNow()
		}
	}
}

func TestJsonParserError(t *testing.T) {
	data := `{"name":`
	_, err := jsonParser([]byte(data), testStringFunc)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error: {\"name\"", ngsiErr.Message)
	}
}

func TestJsonParserError2(t *testing.T) {
	data := `{"name":"abcdefghijklmn"`
	_, err := jsonParser([]byte(data), testStringFunc)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error: abcdefghijklmn\"", ngsiErr.Message)
	}
}
