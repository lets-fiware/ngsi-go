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

package ngsilib

import (
	"fmt"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// TimeLib is ...
type TimeLib interface {
	Now() time.Time
	NowUnix() int64
	Unix(sec int64, nsec int64) time.Time
	Format(layout string) string
}

type timeLib struct {
	TTime time.Time
}

func (t *timeLib) Now() time.Time {
	return time.Now()
}

func (t *timeLib) NowUnix() int64 {
	return time.Now().Unix()
}

func (t *timeLib) Unix(sec int64, nsec int64) time.Time {
	t.TTime = time.Unix(sec, nsec)
	return t.TTime
}

func (t *timeLib) Format(layout string) string {
	return t.TTime.Format(layout)
}

// GetDateTime
func GetDateTime(dateTime string) (string, error) {
	const funcName = "getDateTime"

	var err error
	if !IsOrionDateTime(dateTime) {
		dateTime, err = GetExpirationDate(dateTime)
		if err != nil {
			return "", ngsierr.New(funcName, 1, err.Error(), err)
		}
	}
	return dateTime, nil
}

// GetTime
func GetTime(ngsi *NGSI, v int64) string {
	_ = ngsi.TimeLib.Unix(v/1000, 0)
	return ngsi.TimeLib.Format("2006/01/02 15:04:05")
}

// HumanizeUptime
func HumanizeUptime(t int64) string {
	return fmt.Sprintf("%d d, %d h, %d m, %d s", (t/3600)/24, (t/3600)%24, (t/60)%60, t%60)
}
