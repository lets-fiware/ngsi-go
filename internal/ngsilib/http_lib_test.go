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
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGet(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	client := &Client{URL: u, Headers: map[string]string{}}
	client.HTTP = NewHTTPRequet()

	res, _, err := client.HTTPGet()
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

func TestHTTPPost(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	client := &Client{URL: u, Headers: map[string]string{}}
	client.HTTP = NewHTTPRequet()

	res, _, err := client.HTTPPost("")
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

func TestHTTPPut(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	client := &Client{URL: u, Headers: map[string]string{}}
	client.HTTP = NewHTTPRequet()

	res, _, err := client.HTTPPut("")
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

func TestHTTPPatch(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	client := &Client{URL: u, Headers: map[string]string{}}
	client.HTTP = NewHTTPRequet()

	res, _, err := client.HTTPPatch("")
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

func TestHTTPDelete(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	client := &Client{URL: u, Headers: map[string]string{}}
	client.HTTP = NewHTTPRequet()

	res, _, err := client.HTTPDelete()
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

func TestNewHTTPRequest(t *testing.T) {
	r := NewHTTPRequet()

	actual := reflect.TypeOf(r).String()
	expected := "*ngsilib.httpRequest"
	assert.Equal(t, expected, actual)
}

func TestRequest(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	r := NewHTTPRequet()
	u, _ := url.Parse(ts.URL)
	res, _, err := r.Request("GET", u, nil, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

func TestRequestHeaders(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	r := NewHTTPRequet()
	u, _ := url.Parse(ts.URL)
	headers := map[string]string{"Fiware-Service": "fiware"}
	res, _, err := r.Request("GET", u, headers, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, res.StatusCode)
	}
}

func TestRequestErrorNewReader(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	r := NewHTTPRequet()
	u, _ := url.Parse(ts.URL)
	_, _, err := r.Request("POST", u, nil, 1)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unsupported type", ngsiErr.Message)
	}
}

func TestRequestErrorNewRequest(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	r := NewHTTPRequet()
	u, _ := url.Parse(ts.URL)
	u.Host = ":\n"
	_, _, err := r.Request("GET", u, nil, nil)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "parse \"http://:%0A\": invalid port \":%0A\" after host", ngsiErr.Message)
	}
}

func TestRequestErrorDo(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	r := NewHTTPRequet()
	u, _ := url.Parse(ts.URL)
	u.Host = ""
	_, _, err := r.Request("GET", u, nil, nil)
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "Get \"http:\": http: no Host in request URL", ngsiErr.Message)
	}
}

func TestRequestErrorReadAll(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	r := NewHTTPRequet()
	u, _ := url.Parse(ts.URL + "/error")

	_, _, err := r.Request("POST", u, nil, "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "unexpected EOF", ngsiErr.Message)
	}
}

func TestNewReaderString(t *testing.T) {
	s := "tset data"

	_, err := newReader(s)

	assert.NoError(t, err)
}

func TestNewReaderByte(t *testing.T) {
	b := []byte("tset data")

	_, err := newReader(b)

	assert.NoError(t, err)
}

func TestNewReaderError(t *testing.T) {
	i := 123

	_, err := newReader(i)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "unsupported type", ngsiErr.Message)
	}
}

func Route() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})
	m.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	})
	return m
}
