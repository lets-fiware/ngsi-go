package ngsilib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTenantString(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "fiware", expected: true},
		{arg: "open_iot", expected: true},
		{arg: "open@iot", expected: false},
		{arg: "FIWARE", expected: false},
	}

	for _, c := range cases {
		actual := isTenantString(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsScopeString(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "", expected: true},
		{arg: "/", expected: true},
		{arg: "/fiware", expected: true},
		{arg: "/FIWARE", expected: true},
		{arg: "/FIWARE_orion", expected: true},
		{arg: "/fiware/orion", expected: true},
		{arg: "/fiware/orion,/keyrock", expected: true},
		{arg: "/fiware/orion,/keyrock, /abc, /def, /xyz/abc", expected: true},
		{arg: "/#", expected: true},
		{arg: "/*", expected: true},
		{arg: "/FIWARE@orion", expected: false},
		{arg: "FIWARE", expected: false},
	}

	for _, c := range cases {
		actual := isScopeString(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsHTTP(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "http://orion", expected: true},
		{arg: "https://orion", expected: true},
		{arg: "http:/orion", expected: false},
		{arg: "https:/orion", expected: false},
		{arg: "orion", expected: false},
	}

	for _, c := range cases {
		actual := IsHTTP(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsIPAddress(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "192.168.1.1", expected: true},
		{arg: "192.168.1.1:1026", expected: true},
		{arg: "orion", expected: false},
		{arg: "orion:1026", expected: false},
	}

	for _, c := range cases {
		actual := isIPAddress(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsLocalHost(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "localhost", expected: true},
		{arg: "localhost:1026", expected: true},
		{arg: "192.168.1.1", expected: false},
		{arg: "192.168.1.1:1026", expected: false},
	}

	for _, c := range cases {
		actual := isLocalHost(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestContains(t *testing.T) {
	list := []string{"abc", "def", "xyz", "123"}

	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "abc", expected: true},
		{arg: "123", expected: true},
		{arg: "orion", expected: false},
		{arg: "576", expected: false},
		{arg: "", expected: false},
	}

	for _, c := range cases {
		actual := Contains(list, c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestIsExpirationDate(t *testing.T) {
	cases := []struct {
		arg      string
		expected bool
	}{
		{arg: "10years", expected: true},
		{arg: "1year", expected: true},
		{arg: "65months", expected: true},
		{arg: "1month", expected: true},
		{arg: "365days", expected: true},
		{arg: "1day", expected: true},
		{arg: "123hours", expected: true},
		{arg: "1hour", expected: true},
		{arg: "orion", expected: false},
		{arg: "576", expected: false},
		{arg: "", expected: false},
	}

	for _, c := range cases {
		actual := isExpirationDate(c.arg)
		assert.Equal(t, c.expected, actual)
	}
}

func TestGetExpirationDate1Year(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("1year")

	if assert.NoError(t, err) {
		assert.Equal(t, "2007-01-02T15:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}
func TestGetExpirationDate5Years(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("5years")

	if assert.NoError(t, err) {
		assert.Equal(t, "2011-01-02T15:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}

func TestGetExpirationDate1Month(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("1month")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-02-02T15:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}
func TestGetExpirationDate5Months(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("5months")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-06-02T15:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}

func TestGetExpirationDate1Day(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("1day")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-01-03T15:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}
func TestGetExpirationDate5Days(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("5days")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-01-07T15:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}

func TestGetExpirationDate1Hour(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("1hour")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-01-02T16:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}
func TestGetExpirationDate5Hours(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("5hours")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-01-02T20:04:05.000Z", date)
	} else {
		t.FailNow()
	}
}

func TestGetExpirationDate1Minute(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("1minute")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-01-02T15:05:05.000Z", date)
	} else {
		t.FailNow()
	}
}

func TestGetExpirationDate1MinusMinute(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("-1minute")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-01-02T15:03:05.000Z", date)
	} else {
		t.FailNow()
	}
}

func TestGetExpirationDate5Minutes(t *testing.T) {
	ngsi := NewNGSI()

	ngsi.TimeLib = &MockTimeLib{dateTime: "2006-01-02T15:04:05.000Z"}

	date, err := GetExpirationDate("5minutes")

	if assert.NoError(t, err) {
		assert.Equal(t, "2006-01-02T15:09:05.000Z", date)
	} else {
		t.FailNow()
	}
}

func TestGetExpirationDateError(t *testing.T) {
	_ = NewNGSI()

	_, err := GetExpirationDate("test")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error test", ngsiErr.Message)
	} else {
		t.FailNow()
	}
}

func TestIsOrionDateTime(t *testing.T) {
	actual := IsOrionDateTime("2022-09-24T12:07:54.035Z")
	expected := true

	assert.Equal(t, expected, actual)

	actual = IsOrionDateTime("2022-09-24T12:07:54.035+09:00")
	expected = true

	assert.Equal(t, expected, actual)

	actual = IsOrionDateTime("2022-09-24")
	expected = true

	assert.Equal(t, expected, actual)
}

func TestIsNameString(t *testing.T) {
	cases := []struct {
		name string
		rc   bool
	}{
		{name: "a", rc: true},
		{name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", rc: true},
		{name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", rc: false},
		{name: "a-________---0", rc: true},
		{name: "0123", rc: false},
		{name: "user@fware", rc: true},
		{name: "user@fware.org", rc: true},
		{name: "", rc: false},
		{name: "0_", rc: false},
		{name: "_", rc: false},
		{name: "-", rc: false},
		{name: "@", rc: false},
	}

	for _, c := range cases {
		if b := IsNameString(c.name); b != c.rc {
			t.Error(fmt.Printf("error \"%s\" is %v", c.name, b))
		}
	}
}

func TestIsNgsiV2(t *testing.T) {
	actual := IsNgsiV2("ngsi-v2")
	expected := true

	assert.Equal(t, expected, actual)

	actual = IsNgsiV2("ngsiv2")
	expected = true

	assert.Equal(t, expected, actual)

	actual = IsNgsiV2("v2")
	expected = true

	assert.Equal(t, expected, actual)

	actual = IsNgsiV2("ld")
	expected = false

	assert.Equal(t, expected, actual)
}

func TestIsNgsiLd(t *testing.T) {
	actual := IsNgsiLd("ngsi-ld")
	expected := true

	assert.Equal(t, expected, actual)

	actual = IsNgsiLd("ld")
	expected = true

	assert.Equal(t, expected, actual)

	actual = IsNgsiLd("v2")
	expected = false

	assert.Equal(t, expected, actual)
}
