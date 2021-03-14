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
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitHeader(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}

	err := client.InitHeader()

	assert.NoError(t, err)
}

func TestInitHeaderKeyrock(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}, Token: "1234"}
	client.Server = &Server{ServerType: "keyrock"}

	err := client.InitHeader()

	assert.NoError(t, err)
}

func TestInitHeaderXAuthToken(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.XAuthToken = true
	client.Token = "000000000000000000000"

	err := client.InitHeader()
	actual := client.Headers["X-Auth-Token"]
	expected := client.Token

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestInitHeaderAuthorization(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.XAuthToken = false
	client.Token = "000000000000000000000"

	err := client.InitHeader()
	actual := client.Headers["Authorization"]
	expected := "Bearer " + client.Token

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestInitHeaderTenant(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.XAuthToken = false
	client.Tenant = "fiware"

	err := client.InitHeader()
	actual := client.Headers["Fiware-Service"]
	expected := client.Tenant

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestInitHeaderTenantLD(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.XAuthToken = false
	client.Tenant = "fiware"
	client.NgsiType = ngsiLd

	err := client.InitHeader()
	actual := client.Headers["NGSILD-Tenant"]
	expected := client.Tenant

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestInitHeaderErrorTenant(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.XAuthToken = false
	client.Tenant = "FIWARE"

	err := client.InitHeader()

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error FIWARE Service: FIWARE", ngsiErr.Message)
	}
}

func TestInitHeaderNgsiV2(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	err := client.InitHeader()

	assert.NoError(t, err)
}

func TestInitHeaderNgsiV2Scope(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2
	client.Scope = "/iot"

	err := client.InitHeader()
	actual := client.Headers["Fiware-ServicePath"]
	expected := client.Scope

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestInitHeaderErrorNgsiV2Scope(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.NgsiType = ngsiV2
	client.Scope = "iot"

	err := client.InitHeader()

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error FIWARE ServicePath: iot", ngsiErr.Message)
	}
}

func TestInitHeaderNgsiLD(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiLd

	err := client.InitHeader()

	assert.NoError(t, err)
}

func TestInitHeaderNgsiLdLink(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiLd
	link := "http://context"
	client.Link = &link

	err := client.InitHeader()
	actual := client.Headers["link"]
	expected := "<http://context>; rel=\"http://www.w3.org/ns/json-ld#context\"; type=\"application/ld+json\""

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestSetHeadersNil(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	client.SetHeaders(nil)
}

func TestSetHeaders(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	headers := make(map[string]string)
	headers["Fiware-ServicePath"] = "/iot"
	client.SetHeaders(headers)

	assert.Equal(t, "/iot", client.Headers["Fiware-ServicePath"])
}

func TestSetHeader(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	client.SetHeader("Fiware-Service", "iot")

	assert.Equal(t, "iot", client.Headers["Fiware-Service"])
}

func TestRemoveHeader(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	client.SetHeader("Fiware-Service", "iot")
	client.RemoveHeader("Fiware-Service")

	_, actual := client.Headers["Fiware-Service"]
	expected := false

	assert.Equal(t, expected, actual)
}

func TestSetContentType(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	client.SetContentType()

	actual := client.Headers["Content-Type"]
	expected := "application/json"
	assert.Equal(t, expected, actual)
}

func TestSetContentTypeLD(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiLd
	client.Link = nil

	client.SetContentType()

	actual := client.Headers["Content-Type"]
	expected := "application/ld+json"
	assert.Equal(t, expected, actual)
}

func TestSetContentJSON(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}

	client.SetContentJSON()

	actual := client.Headers["Content-Type"]
	expected := "application/json"
	assert.Equal(t, expected, actual)
}

func TestSetContentLdJSON(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}

	client.SetContentLdJSON()

	actual := client.Headers["Content-Type"]
	expected := "application/ld+json"
	assert.Equal(t, expected, actual)
}

func TestSetAcceptJSON(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}

	client.SetAcceptJSON()

	actual := client.Headers["Accept"]
	expected := "application/json"
	assert.Equal(t, expected, actual)
}

func TestSetAcceptGeoJSON(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.Server = &Server{ServerType: "broker"}

	client.SetAcceptGeoJSON()

	actual := client.Headers["Accept"]
	expected := "application/geo+json"
	assert.Equal(t, expected, actual)
}

func TestSetPath(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}, Server: &Server{ServerType: "broker"}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	client.SetPath("/version")

	actual := client.URL.Path
	expected := "/version"
	assert.Equal(t, expected, actual)
}

func TestSetPathV2(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}, Server: &Server{ServerType: "broker"}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2

	client.SetPath("/entities")

	actual := client.URL.Path
	expected := "/v2/entities"
	assert.Equal(t, expected, actual)
}

func TestSetPathLD(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}, Server: &Server{ServerType: "broker"}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiLd

	client.SetPath("/entities")

	actual := client.URL.Path
	expected := "/ngsi-ld/v1/entities"
	assert.Equal(t, expected, actual)
}

func TestSetPathV2APIPath(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}, Server: &Server{ServerType: "broker"}}
	client.Server = &Server{ServerType: "broker"}
	client.NgsiType = ngsiV2
	client.APIPathBefore = "/"
	client.APIPathAfter = "/orion"

	client.SetPath("/entities")

	actual := client.URL.Path
	expected := "/orion/v2/entities"
	assert.Equal(t, expected, actual)
}

func TestHasPrefixTrue(t *testing.T) {
	actual := hasPrefix([]string{"/version", "/v2"}, "/v2/entities")
	expected := true
	assert.Equal(t, expected, actual)
}

func TestHasPrefixFalse(t *testing.T) {
	actual := hasPrefix([]string{"/version", "/v2"}, "/v1/entities")
	expected := false
	assert.Equal(t, expected, actual)
}

func TestIsSafeStringFalse(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.SafeString = false

	actual := client.IsSafeString()
	expected := false
	assert.Equal(t, expected, actual)
}

func TestIsSafeStringTrue(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.SafeString = true

	actual := client.IsSafeString()
	expected := true
	assert.Equal(t, expected, actual)
}

func TestIsNgsiV2False(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.NgsiType = ngsiLd

	actual := client.IsNgsiV2()
	expected := false
	assert.Equal(t, expected, actual)
}

func TestIsNgsiV2True(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.NgsiType = ngsiV2

	actual := client.IsNgsiV2()
	expected := true
	assert.Equal(t, expected, actual)
}

func TestIsNgsiLDFalse(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.NgsiType = ngsiV2

	actual := client.IsNgsiLd()
	expected := false
	assert.Equal(t, expected, actual)
}

func TestIsNgsiLDStringTrue(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.NgsiType = ngsiLd

	actual := client.IsNgsiLd()
	expected := true
	assert.Equal(t, expected, actual)
}

func TestResultsCountV2(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.NgsiType = ngsiV2

	res := http.Response{}
	res.Header = make(http.Header)
	res.Header["Fiware-Total-Count"] = []string{"10"}

	actual, err := client.ResultsCount(&res)
	expected := 10

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestResultsCountLd(t *testing.T) {
	client := &Client{URL: &url.URL{}, Headers: map[string]string{}}
	client.NgsiType = ngsiLd

	res := http.Response{}
	res.Header = make(http.Header)
	res.Header["Ngsild-Results-Count"] = []string{"20"}

	actual, err := client.ResultsCount(&res)
	expected := 20

	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}
}

func TestIdmURL(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}
	client.Server.IdmHost = "https://keyrock"

	actual := client.idmURL()
	expected := "https://keyrock"
	assert.Equal(t, expected, actual)
}

func TestIdmURLPath(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}
	client.Server.IdmHost = "/token"
	client.Server.ServerHost = "https://orion"

	actual := client.idmURL()
	expected := "https://orion/token"
	assert.Equal(t, expected, actual)
}

func TestStoreToken(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}

	client.storeToken("token")
	actual := client.Token
	expected := "token"
	assert.Equal(t, expected, actual)
}

func TestGetExpiresIn(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}

	actual := client.getExpiresIn()
	expected := int64(3600)
	assert.Equal(t, expected, actual)
}

func TestCheckTenant(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}

	err := client.CheckTenant("fiware")
	assert.NoError(t, err)
}

func TestCheckTenantError(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}

	err := client.CheckTenant("FIWARE")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error FIWARE Service: FIWARE", ngsiErr.Message)
	}
}

func TestCheckScope(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}

	err := client.CheckScope("/fiware")
	assert.NoError(t, err)
}

func TestCheckScopeError(t *testing.T) {
	client := &Client{URL: &url.URL{}, Server: &Server{}}

	err := client.CheckScope("fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error FIWARE ServicePath: fiware", ngsiErr.Message)
	}
}
