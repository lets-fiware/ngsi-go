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
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitServerList(t *testing.T) {
	testNgsiLibInit()

	InitServerList()

	if gNGSI.serverList == nil {
		t.Error("brokerList is nil")
	}
}

func TestAllServersList(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	actual := ngsi.AllServersList()
	expected := &gNGSI.serverList

	assert.Equal(t, expected, actual)
}

func TestList(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual := gNGSI.serverList.List(false)
	expected := "orion orion-ld"

	assert.Equal(t, expected, actual)

}

func TestListSingleLine(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual := gNGSI.serverList.List(true)
	expected := "orion\norion-ld"

	assert.Equal(t, expected, actual)

}

func TestBrokerInfo(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	orion := Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion"] = &orion
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.serverList.BrokerInfo("orion")
	expected := &orion

	assert.Equal(t, expected, actual)
}

func TestBrokerInfoErrorNotBroker(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.serverList.BrokerInfo("comet")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "host found: comet, but type is comet", ngsiErr.Message)
	}
}

func TestBrokerInfoErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	_, err := gNGSI.serverList.BrokerInfo("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestServerInfo(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	comet := Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["comet"] = &comet
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.serverList.ServerInfo("comet", "")
	expected := &comet

	assert.Equal(t, expected, actual)
}

func TestServerInfoWithFilter(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	comet := Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["comet"] = &comet
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.serverList.ServerInfo("comet", "comet")
	expected := &comet

	assert.Equal(t, expected, actual)
}

func TestServerInfoErrorNotServer(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.serverList.ServerInfo("orion", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "host found: orion, but type is broker", ngsiErr.Message)
	}
}

func TestServerInfoErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	_, err := gNGSI.serverList.ServerInfo("fiware", "comet")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestBrokerInfoJSON(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.serverList.BrokerInfoJSON("orion")
	expected := `{"serverHost":"http://orion/", "serverType": "broker"}`

	assert.JSONEq(t, expected, *actual)
}

func TestBrokerInfoJSONAll(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.serverList.BrokerInfoJSON("")
	expected := `{"orion":{"serverHost":"http://orion/", "serverType": "broker"},"orion-ld":{"serverHost":"http://orion-ld/", "serverType": "broker"}}`

	assert.JSONEq(t, expected, *actual)
}

func TestBrokerInfoJSONErrorMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	_, err := gNGSI.serverList.BrokerInfoJSON("")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	_, err := gNGSI.serverList.BrokerInfoJSON("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorServerType(t *testing.T) {
	_ = testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.serverList.BrokerInfoJSON("comet")
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "host found: comet, but type is comet", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorMarshal2(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	_, err := gNGSI.serverList.BrokerInfoJSON("orion")
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestBrokerList(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual := gNGSI.serverList.BrokerList()
	expected := &gNGSI.serverList

	assert.Equal(t, *expected, actual)
}

func TestServerInfoJSON(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	actual, _ := gNGSI.serverList.ServerInfoJSON("comet", "comet")
	expected := `{"serverHost":"http://comet/", "serverType":"comet"}`

	assert.JSONEq(t, expected, *actual)
}

func TestServerInfoJSONAll(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	actual, _ := gNGSI.serverList.ServerInfoJSON("", "")
	expected := `{"comet":{"serverHost":"http://comet/", "serverType":"comet"}, "ql":{"serverHost":"http://quantumleap/", "serverType":"quantumleap"}}`

	assert.JSONEq(t, expected, *actual)
}

func TestServerInfoJSONErrorMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	_, err := gNGSI.serverList.ServerInfoJSON("", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestServerInfoJSONErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	_, err := gNGSI.serverList.ServerInfoJSON("fiware", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestServerInfoJSONErrorServerType(t *testing.T) {
	_ = testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.serverList.ServerInfoJSON("orion", "")
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "host found: orion, but type is broker", ngsiErr.Message)
	}
}

func TestServerInfoJSONErrorMarshal2(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	_, err := gNGSI.serverList.ServerInfoJSON("ql", "quantumleap")
	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestServerList(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	list := gNGSI.serverList.ServerList("", false)
	actual, _ := json.Marshal(list)
	expected := "{\"comet\":{\"serverType\":\"comet\",\"serverHost\":\"http://comet/\"},\"ql\":{\"serverType\":\"quantumleap\",\"serverHost\":\"http://quantumleap/\"}}"
	assert.JSONEq(t, expected, string(actual))
}

func TestServerListAll(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.serverList = nil
	InitServerList()

	gNGSI.serverList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.serverList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.serverList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.serverList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: errors.New("json error"), DecodeErr: errors.New("json error")}
	list := gNGSI.serverList.ServerList("", true)
	actual, _ := json.Marshal(list)
	expected := `{"comet":{"serverType":"comet","serverHost":"http://comet/"},"orion":{"serverType":"broker","serverHost":"http://orion/"},"orion-ld":{"serverType":"broker","serverHost":"http://orion-ld/"},"ql":{"serverType":"quantumleap","serverHost":"http://quantumleap/"}}`
	assert.JSONEq(t, expected, string(actual))
}
