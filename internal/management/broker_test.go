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

package management

import (
	"errors"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func TestBrokersList(t *testing.T) {
	c := setupTest([]string{"broker", "list"})

	err := brokersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "orion orion-alias orion-ld scorpio\n"
		assert.Equal(t, expected, actual)
	}
}

func TestBrokersListHost(t *testing.T) {
	c := setupTest([]string{"broker", "list", "--host", "orion"})

	err := brokersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "brokerHost https://orion\nngsiType v2\n"
		assert.Equal(t, expected, actual)
	}
}

func TestBrokersListJSON(t *testing.T) {
	c := setupTest([]string{"broker", "list", "--json"})

	err := brokersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"orion\":{\"serverType\":\"broker\",\"serverHost\":\"https://orion\",\"ngsiType\":\"v2\"},\"orion-alias\":{\"serverType\":\"broker\",\"serverHost\":\"orion-ld\",\"brokerType\":\"orion-ld\",\"ngsiType\":\"ld\"},\"orion-ld\":{\"serverType\":\"broker\",\"serverHost\":\"https://orion-ld\",\"brokerType\":\"orion-ld\",\"ngsiType\":\"ld\"},\"scorpio\":{\"serverType\":\"broker\",\"serverHost\":\"https://scorpio:9090\",\"brokerType\":\"scorpio\",\"ngsiType\":\"ld\"}}"
		assert.Equal(t, expected, actual)
	}
}

func TestBrokersListJSONPretty(t *testing.T) {
	c := setupTest([]string{"broker", "list", "--json", "--pretty"})

	err := brokersList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"orion\": {\n    \"serverType\": \"broker\",\n    \"serverHost\": \"https://orion\",\n    \"ngsiType\": \"v2\"\n  },\n  \"orion-alias\": {\n    \"serverType\": \"broker\",\n    \"serverHost\": \"orion-ld\",\n    \"brokerType\": \"orion-ld\",\n    \"ngsiType\": \"ld\"\n  },\n  \"orion-ld\": {\n    \"serverType\": \"broker\",\n    \"serverHost\": \"https://orion-ld\",\n    \"brokerType\": \"orion-ld\",\n    \"ngsiType\": \"ld\"\n  },\n  \"scorpio\": {\n    \"serverType\": \"broker\",\n    \"serverHost\": \"https://scorpio:9090\",\n    \"brokerType\": \"scorpio\",\n    \"ngsiType\": \"ld\"\n  }\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestBrokersListErrorHost(t *testing.T) {
	c := setupTest([]string{"broker", "list", "--host", "orion-v2"})

	err := brokersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "orion-v2 not found", ngsiErr.Message)
	}
}

func TestBrokersListErrorJSON(t *testing.T) {
	c := setupTest([]string{"broker", "list", "--json"})

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := brokersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json.Marshl error", ngsiErr.Message)
	}
}

func TestBrokersListErrorJSONPretty(t *testing.T) {
	c := setupTest([]string{"broker", "list", "--json", "--pretty"})

	helper.SetJSONIndentError(c.Ngsi)

	err := brokersList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestBrokersGet(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion"})

	err := brokersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "brokerHost https://orion\nngsiType v2\n"
		assert.Equal(t, expected, actual)
	}
}

func TestBrokersGetJSON(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion", "--json"})

	err := brokersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"serverType\":\"broker\",\"serverHost\":\"https://orion\",\"ngsiType\":\"v2\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestBrokersGetJSONPretty(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion", "--json", "--pretty"})

	err := brokersGet(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"serverType\": \"broker\",\n  \"serverHost\": \"https://orion\",\n  \"ngsiType\": \"v2\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestBrokersGetErrorBrokerListErrorJSON(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion-v2", "--json"})

	err := brokersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "orion-v2 not found", ngsiErr.Message)
	}
}

func TestBrokersGetErrorJSONPretty(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion", "--json", "--pretty"})

	helper.SetJSONIndentError(c.Ngsi)

	err := brokersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestBrokersGetErrorBrokerList(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion-v2"})

	err := brokersGet(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion-v2 not found", ngsiErr.Message)
	}
}

func TestBrokersAdd(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orion-v2", "--brokerHost", "http://orion", "--ngsiType", "v2"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["orion-v2"]
		assert.Equal(t, "v2", v.NgsiType)
	}
}

func TestBrokersAddTenant(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orion-v2", "--brokerHost", "http://orion", "--ngsiType", "v2", "--service", "foo"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["orion-v2"]
		assert.Equal(t, "foo", v.Tenant)
	}
}

func TestBrokersAddLD(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orionld", "--brokerHost", "http://orion", "--ngsiType", "ld"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["orionld"]
		assert.Equal(t, "ld", v.NgsiType)
	}
}

func TestBrokersAddLDSafeString(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orionld", "--brokerHost", "http://orion", "--ngsiType", "ld", "--safeString", "on"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["orionld"]
		assert.Equal(t, "on", v.SafeString)
	}
}

func TestBrokersAddOverWrite(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orion", "--brokerHost", "http://overwrite", "--ngsiType", "v2", "--overWrite"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["orion"]
		assert.Equal(t, "http://overwrite", v.ServerHost)
	}
}

func TestBrokersAddErrorNameString(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "@orion", "--brokerHost", "http://orion", "--ngsiType", "v2"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "name error @orion", ngsiErr.Message)
	}
}

func TestBrokersAddErrorOverWrite(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orion", "--brokerHost", "http://orion", "--ngsiType", "v2", "--overWrite"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("open error"), Filename: helper.StrPtr("ngsi-go-config.json")}

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestBrokersAddErrorAlreadyExists(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orion", "--brokerHost", "http://orion", "--ngsiType", "v2"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "orion already exists", ngsiErr.Message)
	}
}

func TestBrokersAddErrorBrokerHost(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orion-v2", "--ngsiType", "v2"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "brokerHost is missing", ngsiErr.Message)
	}
}

func TestBrokersAddErrorNgsiType(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orion-v2", "--brokerHost", "http://orion"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "ngsiType is missing", ngsiErr.Message)
	}
}

func TestBrokersAddErrorNgsiType2(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orionv2", "--brokerHost", "orion2", "--ngsiType", "v2"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "can't specify ngsiType", ngsiErr.Message)
	}
}

func TestBrokersAddErrorCreateBroker(t *testing.T) {
	c := setupTest([]string{"broker", "add", "--host", "orionld", "--brokerHost", "http://orion-ld", "--ngsiType", "ld", "--safeString", "123"})

	err := brokersAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "unknown parameter: 123", ngsiErr.Message)
	}
}

func TestBrokersUpdate(t *testing.T) {
	c := setupTest([]string{"broker", "update", "--host", "orion", "--ngsiType", "ld"})

	err := brokersUpdate(c, c.Ngsi, c.Client)

	assert.NoError(t, err)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["orion"]
		assert.Equal(t, "ld", v.NgsiType)
	}
}

func TestBrokersUpdateService(t *testing.T) {
	c := setupTest([]string{"broker", "update", "--host", "orion", "--ngsiType", "ld", "--service", "foo"})

	err := brokersUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		list := c.Ngsi.AllServersList()
		v := (*list)["orion"]
		assert.Equal(t, "ld", v.NgsiType)
		assert.Equal(t, "foo", v.Tenant)
	}
}

func TestBrokersUpdateErrorAlreadyExists(t *testing.T) {
	c := setupTest([]string{"broker", "update", "--host", "orionld"})

	err := brokersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "orionld not found", ngsiErr.Message)
	}
}

func TestBrokersUpdateErrorCreateBroker(t *testing.T) {
	c := setupTest([]string{"broker", "update", "--host", "orion", "--ngsiType", "v1"})

	err := brokersUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "v1 not found", ngsiErr.Message)
	}
}

func TestBrokersDelete(t *testing.T) {
	c := setupTest([]string{"broker", "delete", "--host", "orion"})

	err := brokersDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestBrokersDeleteErrorNotFound(t *testing.T) {
	c := setupTest([]string{"broker", "delete", "--host", "orionld"})

	err := brokersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "orionld not found", ngsiErr.Message)
	}
}

func TestBrokersDeleteNoItem(t *testing.T) {
	c := setupTest([]string{"broker", "delete", "--host", "orion", "--items", "noitem"})

	err := brokersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "delete error - noitem", ngsiErr.Message)
	}
}

func TestBrokersDeleteErrorUpdateBroker(t *testing.T) {
	c := setupTest([]string{"broker", "delete", "--host", "orion", "--items", "ngsiType"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("open error"), Filename: helper.StrPtr("ngsi-go-config.json")}

	err := brokersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestBrokersDeleteErrorDeleteServer(t *testing.T) {
	c := setupTest([]string{"broker", "delete", "--host", "orion"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("open error"), Filename: helper.StrPtr("ngsi-go-config.json")}

	err := brokersDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestDeleteBrokerAlias(t *testing.T) {
	c := setupTest([]string{"broker", "delete", "--host", "orion"})

	err := deleteBrokerAlias(c.Ngsi, "orion")

	assert.NoError(t, err)
}

func TestDeleteBrokerAliasErrorReference(t *testing.T) {
	conf := `{
		"version": "1",
		"servers": {
			"orion": {
				  "serverHost": "https://orion",
				  "ngsiType": "v2"
			},
			"orion-ld": {
				  "serverHost": "orion"
			}
		}
	}`

	c := setupTestWithConfig([]string{"broker", "delete", "--host", "orion"}, conf)

	err := deleteBrokerAlias(c.Ngsi, "orion")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "orion is referenced", ngsiErr.Message)
	}
}

func TestDeleteBrokerAliasErrorDeleteServer(t *testing.T) {
	c := setupTest([]string{"broker", "delete", "--host", "orion"})

	c.Ngsi.ConfigFile = &helper.MockIoLib{OpenErr: errors.New("open error"), Filename: helper.StrPtr("ngsi-go-config.json")}

	err := deleteBrokerAlias(c.Ngsi, "orion")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "open error", ngsiErr.Message)
	}
}

func TestPrintBrokerInfoV2(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion"})

	broker := ngsilib.Server{
		ServerHost:   "http://orion",
		NgsiType:     "v2",
		Tenant:       "openiot",
		Scope:        "/iot",
		Context:      "http://context",
		SafeString:   "on",
		IdmType:      "keyrock",
		IdmHost:      "http://keyrock",
		Username:     "fiware",
		Password:     "1234",
		ClientID:     "clientid",
		ClientSecret: "clientsecret",
		XAuthToken:   "false",
		Token:        "token",
		APIPath:      "/path",
	}

	printBrokerInfo(c.Ngsi, &broker, false)

	actual := helper.GetStdoutString(c)
	expected := "brokerHost http://orion\nngsiType v2\nFIWARE-Service openiot\nFIWARE-ServicePath /iot\nContext http://context\nSafeString on\nIdmType keyrock\nIdmHost http://keyrock\nUsername fiware\nPassword ****\nClientID ********\nClientSecret ************\nXAuthToken false\nToken token\nAPIPath /path\n"
	assert.Equal(t, expected, actual)
}

func TestPrintBrokerInfoV2ClearText(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion"})

	broker := ngsilib.Server{
		ServerHost:   "http://orion",
		NgsiType:     "v2",
		Tenant:       "openiot",
		Scope:        "/iot",
		Context:      "http://context",
		SafeString:   "on",
		IdmType:      "keyrock",
		IdmHost:      "http://keyrock",
		Username:     "fiware",
		Password:     "1234",
		ClientID:     "clientid",
		ClientSecret: "clientsecret",
		XAuthToken:   "false",
		Token:        "token",
		APIPath:      "/path",
	}

	printBrokerInfo(c.Ngsi, &broker, true)

	actual := helper.GetStdoutString(c)
	expected := "brokerHost http://orion\nngsiType v2\nFIWARE-Service openiot\nFIWARE-ServicePath /iot\nContext http://context\nSafeString on\nIdmType keyrock\nIdmHost http://keyrock\nUsername fiware\nPassword 1234\nClientID clientid\nClientSecret clientsecret\nXAuthToken false\nToken token\nAPIPath /path\n"
	assert.Equal(t, expected, actual)
}

func TestPrintBrokerInfoLD(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "orion-ld"})

	broker := ngsilib.Server{
		ServerHost: "http://orion-ld",
		NgsiType:   "ld",
		Tenant:     "openiot",
		Scope:      "/iot",
	}

	printBrokerInfo(c.Ngsi, &broker, false)

	actual := helper.GetStdoutString(c)
	expected := "brokerHost http://orion-ld\nngsiType ld\nTenant openiot\nScope /iot\n"
	assert.Equal(t, expected, actual)
}

func TestPrintBrokerInfoScorpio(t *testing.T) {
	c := setupTest([]string{"broker", "get", "--host", "scorpio"})

	broker := ngsilib.Server{
		ServerHost: "http://scorpio",
		NgsiType:   "ld",
		BrokerType: "scorpio",
		Tenant:     "openiot",
		Scope:      "/iot",
	}

	printBrokerInfo(c.Ngsi, &broker, false)

	actual := helper.GetStdoutString(c)
	expected := "brokerHost http://scorpio\nngsiType ld\nbrokerType scorpio\nTenant openiot\nScope /iot\n"
	assert.Equal(t, expected, actual)
}

func TestObfuscateText(t *testing.T) {
	cases := []struct {
		text      string
		clearText bool
		expected  string
	}{
		{text: "fiware", clearText: false, expected: "******"},
		{text: "fiware", clearText: true, expected: "fiware"},
		{text: "", clearText: false, expected: ""},
		{text: "", clearText: true, expected: ""},
	}
	for _, c := range cases {
		actual := obfuscateText(c.text, c.clearText)
		assert.Equal(t, c.expected, actual)
	}
}
