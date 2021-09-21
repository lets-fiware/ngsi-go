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

package helper

import (
	"errors"
	"net"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestInterfaceAddrs(t *testing.T) {
	n := &MockNetLib{}

	actual, err := n.InterfaceAddrs()

	if assert.NoError(t, err) {
		assert.NotEqual(t, ([]net.Addr)(nil), actual)
	}
}

func TestInterfaceAddrsError(t *testing.T) {
	n := &MockNetLib{AddrErr: errors.New("InterfaceAddrs error")}

	actual, err := n.InterfaceAddrs()

	if assert.Error(t, err) {
		assert.Equal(t, ([]net.Addr)(nil), actual)
		assert.Equal(t, "InterfaceAddrs error", err.Error())
	}
}

func TestListenAndServe(t *testing.T) {
	n := &MockNetLib{ListenAndServeErr: errors.New("ListenAndServe error")}

	err := n.ListenAndServe("", nil)

	if assert.Error(t, err) {
		assert.Equal(t, "ListenAndServe error", err.Error())
	}
}

func TestListenAndServeTLS(t *testing.T) {
	n := &MockNetLib{ListenAndServeTLSErr: errors.New("ListenAndServeTLS error")}

	err := n.ListenAndServeTLS("", "", "", nil)

	if assert.Error(t, err) {
		assert.Equal(t, "ListenAndServeTLS error", err.Error())
	}
}
