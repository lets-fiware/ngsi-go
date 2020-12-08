/*
MIT License

Copyright (c) 2020 Kazuhito Suda

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
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	tenantRegexp            = regexp.MustCompile(`^[-0-9a-z_]{1,50}$`)
	scopeRegexp             = regexp.MustCompile(`^(/[-0-9A-Za-z_]{1,50}(/[0-9A-Za-z_]{1,50}){0,9},[ ]*)*/[0-9A-Za-z_]{1,50}(/[0-9A-Za-z_]{1,50}){0,9}$`)
	ipRegexp                = regexp.MustCompile(`^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(:[1-9][0-9]*)*$`)
	localhostRegexp         = regexp.MustCompile(`^localhost(:[1-9][0-9]{0,3})*$`)
	expirationDateRegexp    = regexp.MustCompile(`^[1-9][0-9]{0,2}(year|month|day|hour|minute)s?$`)
	orionDateTimeZoneRegexp = regexp.MustCompile(`^2[0-9]{3}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])(T([0-1][0-9]|2[0-4]):[0-5][0-9]:[0-5][0-9].[0-9]{2,3}(Z|(\+|-)([01][0-9]|2[0-4]):?([0-5][0-9])?))?$`)
	nameRegexp              = regexp.MustCompile(`^[A-Za-z][-@0-9A-Za-z_]{0,31}$`)
)

func isTenantString(s string) bool {
	return tenantRegexp.MatchString(s)
}

func isScopeString(s string) bool {
	if s == "" || s == "/" {
		return true
	}
	return scopeRegexp.MatchString(s)
}

// IsHTTP is ...
func IsHTTP(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func isIPAddress(s string) bool {
	return ipRegexp.MatchString(s)
}

func isLocalHost(s string) bool {
	return localhostRegexp.MatchString(s)
}

// Contains is ...
func Contains(a []string, e string) bool {
	e = strings.ToLower(e)
	for _, v := range a {
		if e == strings.ToLower(v) {
			return true
		}
	}
	return false
}

// GetExpirationDate is ...
func GetExpirationDate(s string) (string, error) {
	const funcName = "GetExpirationDate"

	if !isExpirationDate(s) {
		return s, &NgsiLibError{funcName, 1, "error " + s, nil}
	}
	if strings.HasSuffix(s, "s") {
		s = strings.TrimSuffix(s, "s")
	}

	saveTime := time.Local

	time.Local = time.FixedZone("Local", 0)
	t := gNGSI.TimeLib.Now()

	for _, v := range []string{"year", "month", "day", "hour", "minute"} {
		if strings.HasSuffix(s, v) {
			i, _ := strconv.Atoi(strings.TrimSuffix(s, v))
			switch v {
			case "year":
				t = t.AddDate(i, 0, 0)
			case "month":
				t = t.AddDate(0, i, 0)
			case "day":
				t = t.AddDate(0, 0, i)
			case "hour":
				t = t.Add(time.Duration(i) * 3600 * time.Second)
			case "minute":
				t = t.Add(time.Duration(i) * 60 * time.Second)
			}
			break
		}
	}

	s = t.Format("2006-01-02T15:04:05.000Z")
	time.Local = saveTime

	return s, nil
}

// isExpirationDate is ...
func isExpirationDate(s string) bool {
	return expirationDateRegexp.MatchString(s)
}

// IsOrionDateTime is ...
func IsOrionDateTime(s string) bool {
	return orionDateTimeZoneRegexp.MatchString(s)
}

// IsNameString is ...
func IsNameString(s string) bool {
	return nameRegexp.MatchString(s)
}

// IsNgsiV2 is ...
func IsNgsiV2(s string) bool {
	return Contains(ngsiV2Types, strings.ToLower(s))
}

// IsNgsiLd is ...
func IsNgsiLd(s string) bool {
	return Contains(ngsiLdTypes, strings.ToLower(s))
}
