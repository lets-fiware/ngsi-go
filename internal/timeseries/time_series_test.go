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

package timeseries

import (
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestTsAttrReadComet(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "comet"})

	err := tsAttrRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestTsAttrReadQuantumleap(t *testing.T) {
	c := setupTest([]string{"hget", "attr", "--host", "ql"})

	err := tsAttrRead(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing attr", ngsiErr.Message)
	}
}

func TestTsEntitiesDeleteComet(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "comet"})

	err := tsEntitiesDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestTsEntitiesDeleteQuantumleap(t *testing.T) {
	c := setupTest([]string{"hdelete", "entities", "--host", "ql"})

	err := tsEntitiesDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestTsEntityDeleteComet(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "comet"})

	err := tsEntityDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing type", ngsiErr.Message)
	}
}

func TestTsEntityDeleteQuantumleap(t *testing.T) {
	c := setupTest([]string{"hdelete", "entity", "--host", "ql"})

	err := tsEntityDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "missing id", ngsiErr.Message)
	}
}
