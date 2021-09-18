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
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestGetKeyStoneTokenRequest(t *testing.T) {
	testNgsiLibInit()

	cases := []struct {
		name     string
		password string
		tenant   string
		scorpe   string
		expected string
	}{
		{
			name:     "fiware",
			password: "1234",
			tenant:   "",
			scorpe:   "",
			expected: `{"auth":{"identity":{"methods":["password"],"password":{"user":{"domain":{"name":""},"name":"fiware","password":"1234"}}},"scope":{"domain":{"name":""}}}}`,
		},
		{
			name:     "fiware",
			password: "1234",
			tenant:   "smartcity",
			scorpe:   "",
			expected: `{"auth":{"identity":{"methods":["password"],"password":{"user":{"domain":{"name":"smartcity"},"name":"fiware","password":"1234"}}},"scope":{"domain":{"name":"smartcity"}}}}`,
		},
		{
			name:     "fiware",
			password: "1234",
			tenant:   "",
			scorpe:   "/madrid",
			expected: `{"auth":{"identity":{"methods":["password"],"password":{"user":{"domain":{"name":""},"name":"fiware","password":"1234"}}},"scope":{"project":{"domain":{"name":""},"name":"/madrid"}}}}`,
		},
		{
			name:     "fiware",
			password: "1234",
			tenant:   "smartcity",
			scorpe:   "/madrid",
			expected: `{"auth":{"identity":{"methods":["password"],"password":{"user":{"domain":{"name":"smartcity"},"name":"fiware","password":"1234"}}},"scope":{"project":{"domain":{"name":"smartcity"},"name":"/madrid"}}}}`,
		},
	}

	for _, c := range cases {
		actual := getKeyStoneTokenRequest(c.name, c.password, c.tenant, c.scorpe)
		assert.Equal(t, c.expected, actual)
	}
}
