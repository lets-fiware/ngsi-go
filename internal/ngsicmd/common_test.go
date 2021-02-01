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

package ngsicmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceAddrs(t *testing.T) {
	n := &netLib{}
	_, err := n.InterfaceAddrs()

	assert.NoError(t, err)
}

func TestListenAndServe(t *testing.T) {
	n := &netLib{}
	err := n.ListenAndServe("", nil)

	assert.Error(t, err)
}
func TestListenAndServeTLS(t *testing.T) {
	n := &netLib{}
	err := n.ListenAndServeTLS("", "", "", nil)
	assert.Error(t, err)
}

func TestGetDateTimeISO8601(t *testing.T) {

	actual, err := getDateTime("2022-09-24T12:07:54.035Z")
	expected := "2022-09-24T12:07:54.035Z"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestGetDateTime1Day(t *testing.T) {
	setupTest()

	_, err := getDateTime("1day")

	assert.NoError(t, err)
}

func TestGetDateTimeError(t *testing.T) {
	setupTest()

	_, err := getDateTime("1")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}
