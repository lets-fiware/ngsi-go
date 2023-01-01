/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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

package ngsicli

import (
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestArgsGet(t *testing.T) {

	cmdArgs := cmdArgs{"orion", "iota", "wirecloud", "cygnus", "comet"}

	assert.Equal(t, "orion", cmdArgs.Get(0))
	assert.Equal(t, "iota", cmdArgs.Get(1))
	assert.Equal(t, "wirecloud", cmdArgs.Get(2))
	assert.Equal(t, "cygnus", cmdArgs.Get(3))
	assert.Equal(t, "comet", cmdArgs.Get(4))
}

func TestArgsGetError(t *testing.T) {

	cmdArgs := cmdArgs{"orion", "iota", "wirecloud", "cygnus", "comet"}

	assert.Equal(t, "", cmdArgs.Get(-1))
	assert.Equal(t, "", cmdArgs.Get(5))
}

func TestArgsLen(t *testing.T) {

	cmdArgs := cmdArgs{"orion", "iota", "wirecloud", "cygnus", "comet"}

	assert.Equal(t, 5, cmdArgs.Len())
}
