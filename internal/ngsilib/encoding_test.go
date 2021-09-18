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

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestIsJSON(t *testing.T) {
	testNgsiLibInit()

	cases := []struct {
		value    []byte
		expected bool
	}{
		{value: []byte("["), expected: true},
		{value: []byte("{"), expected: true},
		{value: []byte(" ["), expected: true},
		{value: []byte(" {"), expected: true},
		{value: []byte("\t{"), expected: true},
		{value: []byte("\t["), expected: true},
		{value: []byte("	{"), expected: true},
		{value: []byte("	["), expected: true},
		{value: []byte("abc"), expected: false},
		{value: []byte("123"), expected: false},
		{value: []byte(" abc"), expected: false},
		{value: []byte(" 123"), expected: false},
		{value: []byte("\t123"), expected: false},
		{value: []byte("\tabc"), expected: false},
		{value: []byte("	123"), expected: false},
		{value: []byte("	abc"), expected: false},
		{value: []byte("a[bc"), expected: false},
		{value: []byte("1{23"), expected: false},
		{value: []byte(" a[bc"), expected: false},
		{value: []byte(" 1{23"), expected: false},
		{value: []byte("\t1{23"), expected: false},
		{value: []byte("\ta[bc"), expected: false},
		{value: []byte("	1{23"), expected: false},
		{value: []byte("	a[bc"), expected: false},
	}

	for _, c := range cases {
		actual := IsJSON(c.value)
		assert.Equal(t, c.expected, actual)
	}

}

func TestIsJSONArray(t *testing.T) {
	testNgsiLibInit()

	cases := []struct {
		value    []byte
		expected bool
	}{
		{value: []byte("["), expected: true},
		{value: []byte("{"), expected: false},
		{value: []byte(" ["), expected: true},
		{value: []byte(" {"), expected: false},
		{value: []byte("\t{"), expected: false},
		{value: []byte("\t["), expected: true},
		{value: []byte("	{"), expected: false},
		{value: []byte("	["), expected: true},
		{value: []byte("abc"), expected: false},
		{value: []byte("123"), expected: false},
		{value: []byte(" abc"), expected: false},
		{value: []byte(" 123"), expected: false},
		{value: []byte("\t123"), expected: false},
		{value: []byte("\tabc"), expected: false},
		{value: []byte("	123"), expected: false},
		{value: []byte("	abc"), expected: false},
		{value: []byte("a[bc"), expected: false},
		{value: []byte("1{23"), expected: false},
		{value: []byte(" a[bc"), expected: false},
		{value: []byte(" 1{23"), expected: false},
		{value: []byte("\t1{23"), expected: false},
		{value: []byte("\ta[bc"), expected: false},
		{value: []byte("	1{23"), expected: false},
		{value: []byte("	a[bc"), expected: false},
	}

	for _, c := range cases {
		actual := IsJSONArray(c.value)
		assert.Equal(t, c.expected, actual)
	}

}

func TestGetJSONArray(t *testing.T) {
	testNgsiLibInit()

	var v interface{}

	var a, b float64
	a = -1
	b = 100
	err := GetJSONArray([]byte("[-1, 100]"), &v)
	expected := []interface{}([]interface{}{a, b})
	if assert.NoError(t, err) {
		assert.Equal(t, expected, v)
	}

	err = GetJSONArray([]byte(" [-1, 100"), v)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}

	err = GetJSONArray([]byte(" {}"), v)
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "not JSON Array: {}", ngsiErr.Message)
	}
}

func TestJSONMarshalEncode(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery
	err := JSONUnmarshalEncode([]byte(subscriptionTemplate), &template, true)
	assert.NoError(t, err)

	actual, err := JSONMarshalEncode(template, false)
	expected := "{\"description\":\"%3C%3ESubscription template\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"\"}],\"condition\":{\"attrs\":[]}},\"notification\":{\"http\":{\"url\":\"http://localhost:1028/accumulate\"},\"attrs\":[]},\"expires\":\"2099-12-31T14:00:00.00Z\",\"throttling\":0}"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, string(actual))
	}

}

func TestJSONUnmarshal(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery
	err := JSONUnmarshal([]byte(subscriptionTemplate), &template)
	assert.NoError(t, err)

	actual, err := JSONMarshal(template)
	expected := "{\"description\":\"<>Subscription template\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"\"}],\"condition\":{\"attrs\":[]}},\"notification\":{\"http\":{\"url\":\"http://localhost:1028/accumulate\"},\"attrs\":[]},\"expires\":\"2099-12-31T14:00:00.00Z\",\"throttling\":0}"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, string(actual))
	}

}

func TestJSONMarshalDecode(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery
	err := JSONUnmarshalEncode([]byte(subscriptionTemplate), &template, true)
	assert.NoError(t, err)

	actual, err := JSONMarshalDecode(template, false)
	expected := "{\"description\":\"%3C%3ESubscription template\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"\"}],\"condition\":{\"attrs\":[]}},\"notification\":{\"http\":{\"url\":\"http://localhost:1028/accumulate\"},\"attrs\":[]},\"expires\":\"2099-12-31T14:00:00.00Z\",\"throttling\":0}"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, string(actual))
	}

	actual, err = JSONMarshalDecode(&template, true)
	expected = "{\"description\":\"<>Subscription template\",\"subject\":{\"entities\":[{\"idPattern\":\".*\",\"type\":\"\"}],\"condition\":{\"attrs\":[]}},\"notification\":{\"http\":{\"url\":\"http://localhost:1028/accumulate\"},\"attrs\":[]},\"expires\":\"2099-12-31T14:00:00.00Z\",\"throttling\":0}"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, string(actual))
	}
}

func TestJSONUnmarshalEncode(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery

	err := JSONUnmarshalEncode([]byte(subscriptionTemplate), &template, true)

	if assert.NoError(t, err) {
		assert.Equal(t, "%3C%3ESubscription template", template.Description)
		assert.Equal(t, "http://localhost:1028/accumulate", template.Notification.HTTP.URL)
		assert.Equal(t, 0, template.Throttling)
	}

	err = JSONUnmarshalEncode([]byte(subscriptionTemplate), &template, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "<>Subscription template", template.Description)
		assert.Equal(t, "http://localhost:1028/accumulate", template.Notification.HTTP.URL)
		assert.Equal(t, 0, template.Throttling)
	}
}

func TestJSONUnmarshalDecode(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery

	err := JSONUnmarshalDecode([]byte(test), &template, true)

	if assert.NoError(t, err) {
		assert.Equal(t, "<>Subscription template", template.Description)
	}

	err = JSONUnmarshalDecode([]byte(test), &template, false)

	if assert.NoError(t, err) {
		assert.Equal(t, "%3C%3ESubscription template", template.Description)
	}
}

func TestJsonUnmarshalErrorJSON(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery
	err := jsonUnmarshal([]byte(`{"id": aa`), template, true, SafeStringEncode)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'a' looking for beginning of value (5) {\"id\": aa", ngsiErr.Message)
	}
}

func TestJsonUnmarshalErrorJSONEof(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery
	err := jsonUnmarshal([]byte("{"), template, true, SafeStringEncode)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json error: {", ngsiErr.Message)
	}
}

func TestJsonUnmarshalErrorJSONNoSafeString(t *testing.T) {
	testNgsiLibInit()

	var template subscriptionQuery
	err := jsonUnmarshal([]byte(`{"name":aa`), &template, false, SafeStringEncode)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "invalid character 'a' looking for beginning of value (9) {\"name\":aa", ngsiErr.Message)
	}
}

func TestJsonUnmarshalErrorUnmarshalTypeError(t *testing.T) {
	testNgsiLibInit()

	var template []interface{}
	err := jsonUnmarshal([]byte(`{}`), &template, false, SafeStringEncode)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json: cannot unmarshal object into Go value of type []interface {} Field: (1) {}", ngsiErr.Message)
	}
}

func TestJsonUnmarshalErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()
	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}

	var template interface{}
	err := jsonUnmarshal([]byte(`{}`), &template, false, SafeStringEncode)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

// subscription query
type subscriptionQuery struct {
	Description string `json:"description"`
	Subject     struct {
		Entities []struct {
			IDPattern string `json:"idPattern"`
			Type      string `json:"type"`
		} `json:"entities"`
		Condition struct {
			Attrs []string `json:"attrs"`
		} `json:"condition"`
	} `json:"subject"`
	Notification struct {
		HTTP struct {
			URL string `json:"url"`
		} `json:"http"`
		Attrs []string `json:"attrs"`
	} `json:"notification"`
	Expires    string `json:"expires"`
	Throttling int    `json:"throttling"`
}

const test = `{
	"description": "%3C%3ESubscription template",
	"subject": {
		"entities": [
			{
				"idPattern": ".*",
				"type": ""
			}
		 ],
		 "condition": {
				"attrs": []
			 }
		},
		"notification": {
			 "http": {
				"url": "http://localhost:1028/accumulate"
			},
			"attrs": []
		},
		"expires": "2099-12-31T14:00:00.00Z",
		"throttling": 0
}`

const subscriptionTemplate string = `{
	"description": "<>Subscription template",
	"subject": {
	"entities": [
		{
			"idPattern": ".*",
			"type": ""
		}
 	],
 	"condition": {
			"attrs": []
 		}
	},
	"notification": {
 		"http": {
			"url": "http://localhost:1028/accumulate"
		},
		"attrs": []
	},
	"expires": "2099-12-31T14:00:00.00Z",
	"throttling": 0
}`
