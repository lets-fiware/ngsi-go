/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

type token struct {
	args []string
	cur  int
	len  int
}

func newToken(args []string) *token {
	return &token{args: args, len: len(args)}
}

func (t *token) Next() *string {
	if t.cur >= t.len {
		t.cur += 1
		if t.cur > t.len+1 {
			t.cur = t.len + 1
		}
		return nil
	}
	arg := t.args[t.cur]
	t.cur += 1

	return &arg
}

func (t *token) Peek() *string {
	if t.cur >= t.len {
		return nil
	}
	arg := t.args[t.cur]

	return &arg
}

func (t *token) Prev() *string {
	t.cur -= 1
	if t.cur < 0 {
		t.cur = 0
		return nil
	}
	if t.cur >= t.len {
		return nil
	}
	arg := t.args[t.cur]

	return &arg
}
