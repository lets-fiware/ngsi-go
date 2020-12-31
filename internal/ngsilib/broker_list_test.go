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

package ngsilib

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitBrokerList(t *testing.T) {
	testNgsiLibInit()

	InitBrokerList()

	if gNGSI.brokerList == nil {
		t.Error("brokerList is nil")
	}
}

func TestBrokerList(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	actual := ngsi.BrokerList()
	expected := &gNGSI.brokerList

	assert.Equal(t, expected, actual)
}

func TestList(t *testing.T) {
	testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	actual := gNGSI.brokerList.List()
	expected := "orion orion-ld"

	assert.Equal(t, expected, actual)

}

func TestBrokerInfo(t *testing.T) {
	testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	actual, _ := gNGSI.brokerList.BrokerInfo("orion")
	expected := &orion

	assert.Equal(t, expected, actual)
}

func TestBrokerInfoError(t *testing.T) {
	testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	_, err := gNGSI.brokerList.BrokerInfo("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestBrokerInfoJSON(t *testing.T) {
	testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	actual, _ := gNGSI.brokerList.BrokerInfoJSON("orion")
	expected := `{"brokerHost":"http://orion/"}`

	assert.JSONEq(t, expected, *actual)
}

func TestBrokerInfoJSONAll(t *testing.T) {
	testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	actual, _ := gNGSI.brokerList.BrokerInfoJSON("")
	expected := `{"orion":{"brokerHost":"http://orion/"},"orion-ld":{"brokerHost":"http://orion-ld/"}}`

	assert.JSONEq(t, expected, *actual)
}

func TestBrokerInfoJSONErrorMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	_, err := gNGSI.brokerList.BrokerInfoJSON("")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	_, err := gNGSI.brokerList.BrokerInfoJSON("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorMarshal2(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.brokerList = nil
	InitBrokerList()

	orion := Broker{BrokerHost: "http://orion/"}
	gNGSI.brokerList["orion"] = &orion
	orionld := Broker{BrokerHost: "http://orion-ld/"}
	gNGSI.brokerList["orion-ld"] = &orionld

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	_, err := gNGSI.brokerList.BrokerInfoJSON("orion")

	if assert.Error(t, err) {
		ngsiErr := err.(*NgsiLibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}
