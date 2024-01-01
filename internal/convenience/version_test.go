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

package convenience

import (
	"errors"
	"net/http"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestVersionV2(t *testing.T) {
	c := setupTest([]string{"version", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/version"
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestVersionLD(t *testing.T) {
	c := setupTest([]string{"version", "--host", "orion-ld"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/ngsi-ld/ex/v1/version"
	reqRes.ResBody = []byte(`{"Orion-LD version":"1.0.1-PRE-468","based on orion":"1.15.0-next","kbase version":"0.8","kalloc version":"0.8","khash version":"0.8","kjson version":"0.8","microhttpd version":"0.9.72-0","rapidjson version":"1.0.2","libcurl version":"7.61.1","libuuid version":"UNKNOWN","mongocpp version":"1.1.3","mongoc version":"1.17.5","mongodb server version":"4.4.11","boost version":"1_66","openssl version":"OpenSSL 1.1.1k  FIPS 25 Mar 2021","branch":"","cached subscriptions":0,"Next File Descriptor":20}`)
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"Orion-LD version\":\"1.0.1-PRE-468\",\"based on orion\":\"1.15.0-next\",\"kbase version\":\"0.8\",\"kalloc version\":\"0.8\",\"khash version\":\"0.8\",\"kjson version\":\"0.8\",\"microhttpd version\":\"0.9.72-0\",\"rapidjson version\":\"1.0.2\",\"libcurl version\":\"7.61.1\",\"libuuid version\":\"UNKNOWN\",\"mongocpp version\":\"1.1.3\",\"mongoc version\":\"1.17.5\",\"mongodb server version\":\"4.4.11\",\"boost version\":\"1_66\",\"openssl version\":\"OpenSSL 1.1.1k  FIPS 25 Mar 2021\",\"branch\":\"\",\"cached subscriptions\":0,\"Next File Descriptor\":20}"
		assert.Equal(t, expected, actual)
	}
}

func TestVersionIota(t *testing.T) {
	c := setupTest([]string{"version", "--host", "iota"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/about"
	reqRes.ResBody = []byte(`{"libVersion":"2.14.0","port":"4041","baseRoot":"/","version":"1.15.0"}`)
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\"libVersion\":\"2.14.0\",\"port\":\"4041\",\"baseRoot\":\"/\",\"version\":\"1.15.0\"}"
		assert.Equal(t, expected, actual)
	}
}

func TestVersionCygnus(t *testing.T) {
	c := setupTest([]string{"version", "--host", "cygnus"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/v1/version"
	reqRes.ResBody = []byte(`{"success":"true","version":"2.6.0.0f2695c2bb9a290854cce6243771c08abe07e281"}`)
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"success":"true","version":"2.6.0.0f2695c2bb9a290854cce6243771c08abe07e281"}`
		assert.Equal(t, expected, actual)
	}
}

func TestVersionWireCloud(t *testing.T) {
	c := setupTest([]string{"version", "--host", "wirecloud"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/api/features"
	reqRes.ResBody = []byte(`{"ApplicationMashup": "2.2", "ComponentManagement": "1.0", "DashboardManagement": "1.0", "FIWARE": "7.7.1", "FullscreenWidget": "0.5", "NGSI": "1.2.1", "OAuth2Provider": "0.5", "ObjectStorage": "0.5", "StyledElements": "0.10.0", "Wirecloud": "1.3.1"}`)
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"ApplicationMashup": "2.2", "ComponentManagement": "1.0", "DashboardManagement": "1.0", "FIWARE": "7.7.1", "FullscreenWidget": "0.5", "NGSI": "1.2.1", "OAuth2Provider": "0.5", "ObjectStorage": "0.5", "StyledElements": "0.10.0", "Wirecloud": "1.3.1"}`
		assert.Equal(t, expected, actual)
	}
}

func TestVersionIotaPretty(t *testing.T) {
	c := setupTest([]string{"version", "--host", "iota", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/about"
	reqRes.ResBody = []byte(`{"libVersion":"2.14.0","port":"4041","baseRoot":"/","version":"1.15.0"}`)
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{\n  \"libVersion\": \"2.14.0\",\n  \"port\": \"4041\",\n  \"baseRoot\": \"/\",\n  \"version\": \"1.15.0\"\n}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestVersionPerseoCore(t *testing.T) {
	c := setupTest([]string{"version", "--host", "perseo-core"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/perseo-core/version"
	reqRes.ResBody = []byte("1.5.0")
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "1.5.0"
		assert.Equal(t, expected, actual)
	}
}

func TestVersionKeyrock(t *testing.T) {
	c := setupTest([]string{"version", "--host", "keyrock"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/version"
	reqRes.ResBody = []byte(`{"keyrock":{"version":"7.0.1","release_date":"2018-06-25","uptime":"01:38:39.9","git_hash":"https://github.com/ging/fiware-idm/releases/tag/7.0.1","doc":"https://fiware-idm.readthedocs.io/en/7.0.1/","api":{"version":"v1","link":"https://keyrock.e-suda.info/v1"}}}`)
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := `{"keyrock":{"version":"7.0.1","release_date":"2018-06-25","uptime":"01:38:39.9","git_hash":"https://github.com/ging/fiware-idm/releases/tag/7.0.1","doc":"https://fiware-idm.readthedocs.io/en/7.0.1/","api":{"version":"v1","link":"https://keyrock.e-suda.info/v1"}}}`
		assert.Equal(t, expected, actual)
	}
}

func TestVersionErrorHTTP(t *testing.T) {
	c := setupTest([]string{"version", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/version"
	reqRes.Err = errors.New("error")
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "error", ngsiErr.Message)
	}
}

func TestVersionErrorStatusCode(t *testing.T) {
	c := setupTest([]string{"version", "--host", "orion"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.Path = "/version"
	helper.SetClientHTTP(c, reqRes)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  ", ngsiErr.Message)
	}
}

func TestVersionIotaErrorPretty(t *testing.T) {
	c := setupTest([]string{"version", "--host", "iota", "--pretty"})

	reqRes := helper.MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.Path = "/iot/about"
	reqRes.ResBody = []byte(`{"libVersion":"2.14.0","port":"4041","baseRoot":"/","version":"1.15.0"}`)
	helper.SetClientHTTP(c, reqRes)

	helper.SetJSONIndentError(c.Ngsi)

	err := cbVersion(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}
