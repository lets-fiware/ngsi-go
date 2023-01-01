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

func TestNewToken(t *testing.T) {
	token := newToken([]string{"ngsi", "create", "entity", "--host", "orion"})

	assert.Equal(t, 5, token.len)
	assert.Equal(t, []string{"ngsi", "create", "entity", "--host", "orion"}, token.args)
}

func TestTokenNext(t *testing.T) {
	token := newToken([]string{"ngsi", "create", "entity", "--host", "orion"})

	expected := []string{"ngsi", "create", "entity", "--host", "orion"}

	for _, e := range expected {
		s := token.Next()
		assert.Equal(t, e, *s)
	}

	s := token.Next()
	assert.Equal(t, (*string)(nil), s)

	s = token.Next()
	assert.Equal(t, (*string)(nil), s)
}

func TestTokenPeek(t *testing.T) {
	token := newToken([]string{"ngsi"})

	assert.Equal(t, "ngsi", *token.Peek())

	_ = token.Next()
	assert.Equal(t, (*string)(nil), token.Peek())

	_ = token.Next()
	assert.Equal(t, (*string)(nil), token.Peek())
}

func TestTokenPrev(t *testing.T) {
	token := newToken([]string{"ngsi", "create", "entity", "--host", "orion"})

	s := token.Prev()
	assert.Equal(t, (*string)(nil), s)

	for i := 0; i < 10; i++ {
		_ = token.Next()
	}

	s = token.Prev()
	assert.Equal(t, (*string)(nil), s)

	expected := []string{"orion", "--host", "entity", "create", "ngsi"}

	for _, e := range expected {
		s := token.Prev()
		assert.Equal(t, e, *s)
	}
}
