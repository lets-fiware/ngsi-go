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
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPGet is ...
func (client *Client) HTTPGet() (*http.Response, []byte, error) {
	return client.HTTP.Request(http.MethodGet, client.URL, client.Headers, nil)
}

// HTTPPost is ...
func (client *Client) HTTPPost(body interface{}) (*http.Response, []byte, error) {
	return client.HTTP.Request(http.MethodPost, client.URL, client.Headers, body)
}

// HTTPPut is ...
func (client *Client) HTTPPut(body interface{}) (*http.Response, []byte, error) {
	return client.HTTP.Request(http.MethodPut, client.URL, client.Headers, body)
}

// HTTPPatch is ...
func (client *Client) HTTPPatch(body interface{}) (*http.Response, []byte, error) {
	return client.HTTP.Request(http.MethodPatch, client.URL, client.Headers, body)
}

// HTTPDelete is
func (client *Client) HTTPDelete(body interface{}) (*http.Response, []byte, error) {
	return client.HTTP.Request(http.MethodDelete, client.URL, client.Headers, body)
}

// HTTPRequest is ...
type HTTPRequest interface {
	Request(method string, url *url.URL, headers map[string]string, body interface{}) (*http.Response, []byte, error)
}
type httpRequest struct{}

// NewHTTPRequet is ...
func NewHTTPRequet() HTTPRequest {
	return &httpRequest{}
}

func (r *httpRequest) Request(method string, url *url.URL, headers map[string]string, body interface{}) (res *http.Response, b []byte, err error) {
	const funcName = "Request"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: gNGSI.InsecureSkipVerify},
	}

	client := &http.Client{Timeout: time.Duration(60 * time.Second), Transport: tr}

	var reader io.Reader

	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		if body != nil {
			reader, err = newReader(body)
			if err != nil {
				return nil, nil, &LibError{funcName, 1, err.Error(), err}
			}
		}
	}

	u := url.String()
	if p := strings.Index(u, "/ngsi-ld/v1/attributes/"); p > 0 {
		u = u[:p] + url.Path
	} else if p := strings.Index(u, "/ngsi-ld/v1/types/"); p > 0 {
		u = u[:p] + url.Path
	}

	var req *http.Request
	req, err = http.NewRequest(method, u, reader)
	if err != nil {
		return nil, nil, &LibError{funcName, 2, err.Error(), err}
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return nil, nil, &LibError{funcName, 3, err.Error(), err}
	}
	defer func() { _ = resp.Body.Close() }()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, &LibError{funcName, 5, err.Error(), err}
	}
	return resp, b, nil
}

func newReader(v interface{}) (io.Reader, error) {
	const funcName = "newReader"

	switch v := v.(type) {
	case []byte:
		return bytes.NewReader(v), nil
	case string:
		return strings.NewReader(v), nil
	}
	return nil, &LibError{funcName, 1, "unsupported type", nil}
}
