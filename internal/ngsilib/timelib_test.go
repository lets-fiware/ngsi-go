/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestTimeLibNow(t *testing.T) {
	time := &timeLib{}

	_ = time.Now()
}

func TestTimeLibNowUnix(t *testing.T) {
	time := &timeLib{}

	_ = time.NowUnix()
}

func TestTimeLibUnix(t *testing.T) {
	time := &timeLib{}

	_ = time.Unix(0, 0)
}

func TestTimeLibFormat(t *testing.T) {
	time := &timeLib{}

	_ = time.Unix(0, 0)
	_ = time.Format("2021/06/27 06:47:39")
}

func TestGetDateTimeISO8601(t *testing.T) {

	actual, err := GetDateTime("2022-09-24T12:07:54.035Z")
	expected := "2022-09-24T12:07:54.035Z"

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestGetDateTime1Day(t *testing.T) {
	_ = testNgsiLibInit()

	_, err := GetDateTime("1day")

	assert.NoError(t, err)
}

func TestGetDateTimeError(t *testing.T) {
	_ = testNgsiLibInit()

	_, err := GetDateTime("1")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error 1", ngsiErr.Message)
	}
}

func TestGetTime(t *testing.T) {
	ngsi := testNgsiLibInit()

	expected := "2021/06/27 06:47:39"
	ngsi.TimeLib = &MockTimeLib{format: &expected}
	actual := GetTime(ngsi, 0)
	assert.Equal(t, expected, actual)
}

func TestHumanizeUptime(t *testing.T) {
	cases := []struct {
		t        int64
		expected string
	}{
		{t: 0, expected: "0 d, 0 h, 0 m, 0 s"},
		{t: 10, expected: "0 d, 0 h, 0 m, 10 s"},
		{t: 61, expected: "0 d, 0 h, 1 m, 1 s"},
		{t: 3662, expected: "0 d, 1 h, 1 m, 2 s"},
		{t: 86401, expected: "1 d, 0 h, 0 m, 1 s"},
		{t: 86401 + 3662, expected: "1 d, 1 h, 1 m, 3 s"},
		{t: 86400 - 1, expected: "0 d, 23 h, 59 m, 59 s"},
		{t: 86401*10 + 3662, expected: "10 d, 1 h, 1 m, 12 s"},
	}

	for _, c := range cases {
		actual := HumanizeUptime(c.t)
		assert.Equal(t, c.expected, actual)
	}

}
