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
	"encoding/json"
	"errors"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestInitServerList(t *testing.T) {
	testNgsiLibInit()

	InitServerList()

	if gNGSI.ServerList == nil {
		t.Error("brokerList is nil")
	}
}

func TestAllServersList(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	actual := ngsi.AllServersList()
	expected := &gNGSI.ServerList

	assert.Equal(t, expected, actual)
}

func TestList(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual := gNGSI.ServerList.List(false)
	expected := "orion orion-ld"

	assert.Equal(t, expected, actual)

}

func TestListSingleLine(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual := gNGSI.ServerList.List(true)
	expected := "orion\norion-ld"

	assert.Equal(t, expected, actual)

}

func TestBrokerInfo(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	orion := Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion"] = &orion
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.ServerList.BrokerInfo("orion")
	expected := &orion

	assert.Equal(t, expected, actual)
}

func TestBrokerInfoErrorNotBroker(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.ServerList.BrokerInfo("comet")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "host found: comet, but type is comet", ngsiErr.Message)
	}
}

func TestBrokerInfoErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	_, err := gNGSI.ServerList.BrokerInfo("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestServerInfo(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	comet := Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["comet"] = &comet
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.ServerList.ServerInfo("comet", "")
	expected := &comet

	assert.Equal(t, expected, actual)
}

func TestServerInfoWithFilter(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	comet := Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["comet"] = &comet
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.ServerList.ServerInfo("comet", "comet")
	expected := &comet

	assert.Equal(t, expected, actual)
}

func TestServerInfoErrorNotServer(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.ServerList.ServerInfo("orion", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "host found: orion, but type is broker", ngsiErr.Message)
	}
}

func TestServerInfoErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	_, err := gNGSI.ServerList.ServerInfo("fiware", "comet")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestBrokerInfoJSON(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.ServerList.BrokerInfoJSON("orion")
	expected := `{"serverHost":"http://orion/", "serverType": "broker"}`

	actualMap := make(map[string]string)
	expectedMap := make(map[string]string)

	_ = json.Unmarshal([]byte(*actual), &actualMap)
	_ = json.Unmarshal([]byte(expected), &expectedMap)

	assert.Equal(t, expectedMap, actualMap)
}

func TestBrokerInfoJSONAll(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual, _ := gNGSI.ServerList.BrokerInfoJSON("")
	expected := `{"orion":{"serverHost":"http://orion/", "serverType": "broker"},"orion-ld":{"serverHost":"http://orion-ld/", "serverType": "broker"}}`

	actualMap := make(map[string]map[string]string)
	expectedMap := make(map[string]map[string]string)

	_ = json.Unmarshal([]byte(*actual), &actualMap)
	_ = json.Unmarshal([]byte(expected), &expectedMap)

	assert.Equal(t, expectedMap, actualMap)
}

func TestBrokerInfoJSONErrorMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}
	_, err := gNGSI.ServerList.BrokerInfoJSON("")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	_, err := gNGSI.ServerList.BrokerInfoJSON("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorServerType(t *testing.T) {
	_ = testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.ServerList.BrokerInfoJSON("comet")
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "host found: comet, but type is comet", ngsiErr.Message)
	}
}

func TestBrokerInfoJSONErrorMarshal2(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}
	_, err := gNGSI.ServerList.BrokerInfoJSON("orion")
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestBrokerList(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	actual := gNGSI.ServerList.BrokerList()
	expected := &gNGSI.ServerList

	assert.Equal(t, *expected, actual)
}

func TestServerInfoJSON(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	actual, _ := gNGSI.ServerList.ServerInfoJSON("comet", "comet")
	expected := `{"serverHost":"http://comet/", "serverType":"comet"}`

	actualMap := make(map[string]string)
	expectedMap := make(map[string]string)

	_ = json.Unmarshal([]byte(*actual), &actualMap)
	_ = json.Unmarshal([]byte(expected), &expectedMap)

	assert.Equal(t, expectedMap, actualMap)
}

func TestServerInfoJSONAll(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	actual, _ := gNGSI.ServerList.ServerInfoJSON("", "")
	expected := `{"comet":{"serverHost":"http://comet/", "serverType":"comet"}, "ql":{"serverHost":"http://quantumleap/", "serverType":"quantumleap"}}`

	actualMap := make(map[string]map[string]string)
	expectedMap := make(map[string]map[string]string)

	_ = json.Unmarshal([]byte(*actual), &actualMap)
	_ = json.Unmarshal([]byte(expected), &expectedMap)

	assert.Equal(t, expectedMap, actualMap)
}

func TestServerInfoJSONErrorMarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}
	_, err := gNGSI.ServerList.ServerInfoJSON("", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestServerInfoJSONErrorHostNotFound(t *testing.T) {
	testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	_, err := gNGSI.ServerList.ServerInfoJSON("fiware", "")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host not found: fiware", ngsiErr.Message)
	}
}

func TestServerInfoJSONErrorServerType(t *testing.T) {
	_ = testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}

	_, err := gNGSI.ServerList.ServerInfoJSON("orion", "")
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "host found: orion, but type is broker", ngsiErr.Message)
	}
}

func TestServerInfoJSONErrorMarshal2(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}
	_, err := gNGSI.ServerList.ServerInfoJSON("ql", "quantumleap")
	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestServerList(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}

	actual := gNGSI.ServerList.ServerList("", false)

	var expected ServerList
	_ = json.Unmarshal([]byte("{\"comet\":{\"serverType\":\"comet\",\"serverHost\":\"http://comet/\"},\"ql\":{\"serverType\":\"quantumleap\",\"serverHost\":\"http://quantumleap/\"}}"), &expected)

	assert.Equal(t, expected, actual)

}

func TestServerListAll(t *testing.T) {
	ngsi := testNgsiLibInit()

	gNGSI.ServerList = nil
	InitServerList()

	gNGSI.ServerList["orion"] = &Server{ServerHost: "http://orion/", ServerType: "broker"}
	gNGSI.ServerList["orion-ld"] = &Server{ServerHost: "http://orion-ld/", ServerType: "broker"}
	gNGSI.ServerList["comet"] = &Server{ServerHost: "http://comet/", ServerType: "comet"}
	gNGSI.ServerList["ql"] = &Server{ServerHost: "http://quantumleap/", ServerType: "quantumleap"}

	ngsi.JSONConverter = &MockJSONLib{EncodeErr: [5]error{errors.New("json error")}, DecodeErr: [5]error{errors.New("json error")}}

	actual := gNGSI.ServerList.ServerList("", true)

	var expected ServerList
	_ = json.Unmarshal([]byte(`{"comet":{"serverType":"comet","serverHost":"http://comet/"},"orion":{"serverType":"broker","serverHost":"http://orion/"},"orion-ld":{"serverType":"broker","serverHost":"http://orion-ld/"},"ql":{"serverType":"quantumleap","serverHost":"http://quantumleap/"}}`), &expected)

	assert.Equal(t, expected, actual)
}

func TestGetServerInfo(t *testing.T) {
	ngsi := NGSI{
		ServerList: ServerList{
			"orion1": &Server{ServerHost: "http://localhost:1026", Tenant: "fiware", Scope: "/iot"},
			"orion2": &Server{ServerHost: "orion1", Tenant: "wirecloud", Scope: "/macs"},
			"orion3": &Server{ServerHost: "orion1"},
		},
	}

	tests := []struct {
		Host        string
		Expected    *Server
		SkipRefHost bool
	}{
		{Host: "orion3", Expected: &Server{ServerHost: "orion1", Tenant: "", Scope: ""}, SkipRefHost: true},
		{Host: "orion1", Expected: &Server{ServerHost: "http://localhost:1026", Tenant: "fiware", Scope: "/iot"}},
		{Host: "orion2", Expected: &Server{ServerHost: "http://localhost:1026", Tenant: "wirecloud", Scope: "/macs"}},
		{Host: "orion3", Expected: &Server{ServerHost: "http://localhost:1026", Tenant: "fiware", Scope: "/iot"}},
	}

	for _, test := range tests {
		actual, err := ngsi.GetServerInfo(test.Host, test.SkipRefHost)

		if assert.NoError(t, err) {
			assert.Equal(t, test.Expected.ServerHost, actual.ServerHost)
			assert.Equal(t, test.Expected.Tenant, actual.Tenant)
			assert.Equal(t, test.Expected.Scope, actual.Scope)
		}
	}
}

func TestGetServerInfoError1(t *testing.T) {
	ngsi := NGSI{
		ServerList: ServerList{
			"orion2": &Server{ServerHost: "orion1", Tenant: "wirecloud", Scope: "/macs"},
		},
	}

	actual, err := ngsi.GetServerInfo("orion2", false)

	if assert.Error(t, err) {
		assert.Equal(t, (*Server)(nil), actual)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "host error: orion2, orion1 not found", ngsiErr.Message)
	}
}

func TestGetServerInfoError2(t *testing.T) {
	ngsi := NGSI{
		ServerList: ServerList{
			"orion1": &Server{ServerHost: "http://localhost:1026", Tenant: "fiware", Scope: "/iot"},
			"orion2": &Server{ServerHost: "orion1", Tenant: "wirecloud", Scope: "/macs"},
			"orion3": &Server{ServerHost: "orion3"},
		},
	}

	actual, err := ngsi.GetServerInfo("orion3", false)

	if assert.Error(t, err) {
		assert.Equal(t, (*Server)(nil), actual)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "host error: orion3, orion3 not found", ngsiErr.Message)
	}
}

func TestGetServerInfoError3(t *testing.T) {
	ngsi := NGSI{
		ServerList: ServerList{},
	}

	actual, err := ngsi.GetServerInfo("orion", false)

	if assert.Error(t, err) {
		assert.Equal(t, (*Server)(nil), actual)
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion not found", ngsiErr.Message)
	}
}
