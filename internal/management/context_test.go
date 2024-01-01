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

package management

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/helper"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestContextList(t *testing.T) {
	c := setupTest([]string{"context", "list"})

	err := contextList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "array [\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]\ndata-model http://context-provider:3000/data-models/ngsi-context.jsonld\netsi https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\nld https://schema.lab.fiware.org/ld/context\nobject {\"ld\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"}\ntutorial http://context-provider:3000/data-models/ngsi-context.jsonld\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListJSON(t *testing.T) {
	c := setupTest([]string{"context", "list", "--name", "fiware"})

	_ = c.Ngsi.AddContext("fiware", "{}")

	err := contextList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "{}\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListName(t *testing.T) {
	c := setupTest([]string{"context", "list", "--name", "etsi"})

	err := contextList(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual := helper.GetStdoutString(c)
		expected := "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListErrorName(t *testing.T) {
	c := setupTest([]string{"context", "list", "--name", "fiware"})

	err := contextList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestContextListErrorJSON(t *testing.T) {
	c := setupTest([]string{"context", "list"})

	_ = c.Ngsi.AddContext("fiware", "{}")

	helper.SetJSONEncodeErr(c.Ngsi, 0)

	err := contextList(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestContextAddUrl(t *testing.T) {
	c := setupTest([]string{"context", "add", "--name", "fiware", "--url", "http://fiware"})

	err := contextAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual, _ := c.Ngsi.GetContext("fiware")
		assert.Equal(t, "http://fiware", actual)
	}
}

func TestContextAddJSON(t *testing.T) {
	c := setupTest([]string{"context", "add", "--name", "fiware", "--json", "{}"})

	err := contextAdd(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual, _ := c.Ngsi.GetContext("fiware")
		assert.Equal(t, "{}", actual)
	}
}

func TestContextAddErrorNameString(t *testing.T) {
	c := setupTest([]string{"context", "add", "--name", "@fiware", "--json", "http://fiware"})

	err := contextAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "name error @fiware", ngsiErr.Message)
	}
}

func TestContextAddErrorUrlError(t *testing.T) {
	c := setupTest([]string{"context", "add", "--name", "fiware", "--url", "abc"})

	err := contextAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestContextAddErrorJSONError(t *testing.T) {
	c := setupTest([]string{"context", "add", "--name", "fiware", "--json", "http://context"})

	err := contextAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestContextAddError(t *testing.T) {
	c := setupTest([]string{"context", "add", "--name", "etsi", "--url", "http://context"})

	err := contextAdd(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "etsi already exists", ngsiErr.Message)
	}
}

func TestContextUpdate(t *testing.T) {
	c := setupTest([]string{"context", "update", "--name", "etsi", "--url", "http://context"})

	err := contextUpdate(c, c.Ngsi, c.Client)

	if assert.NoError(t, err) {
		actual, _ := c.Ngsi.GetContext("etsi")
		assert.Equal(t, "http://context", actual)
	}
}

func TestContextUpdateErrorUrlError(t *testing.T) {
	c := setupTest([]string{"context", "update", "--name", "etsi", "--url", "abc"})

	err := contextUpdate(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "abc is not url", ngsiErr.Message)
	}
}

func TestContextDelete(t *testing.T) {
	c := setupTest([]string{"context", "delete", "--name", "etsi"})

	err := contextDelete(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestContextDeleteErrorUrl(t *testing.T) {
	c := setupTest([]string{"context", "delete", "--name", "fiware"})

	err := contextDelete(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestContextServer(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--name", "etsi"})

	err := contextServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestContextServerHTTPS(t *testing.T) {
	c := setupTest([]string{"context", "server", "--https", "--key", "test.key", "--cert", "test.cert", "--port", "aaaa", "--url", "/context", "--name", "etsi"})

	err := contextServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestContextServerJSON(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--name", "array"})

	err := contextServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestContextServerDataJSON(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--data", "{}"})

	err := contextServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestContextServerDataHTTP(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--data", "http://context"})

	err := contextServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestContextServerDataFile(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--data", "@file"})

	c.Ngsi.FileReader = &helper.MockFileLib{ReadFileData: []byte("{}")}

	err := contextServer(c, c.Ngsi, c.Client)

	assert.NoError(t, err)
}

func TestContextServerErrorNotFoundName(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--name", "fiware"})

	err := contextServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestContextServerErrorFilePathAbs(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--data", "@file"})

	helper.SetFilePatAbsError(c.Ngsi, 0)

	err := contextServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "filepathabs error", ngsiErr.Message)
	}
}

func TestContextServerErrorReadFileError(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--data", "@file"})

	helper.SetReadFileError(c.Ngsi, 1)

	err := contextServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "readfile error", ngsiErr.Message)
	}
}

func TestContextServerErrorFileNameError(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--data", "@"})

	err := contextServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestContextServerErrorNotJSON(t *testing.T) {
	c := setupTest([]string{"context", "server", "--port", "aaaa", "--url", "/context", "--data", "file"})

	helper.SetReadFileError(c.Ngsi, 0)

	err := contextServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "data not json", ngsiErr.Message)
	}
}

func TestContextServerErrorKey(t *testing.T) {
	c := setupTest([]string{"context", "server", "--https", "--port", "aaaa", "--url", "/", "--name", "etsi"})

	err := contextServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestContextServerErrorCert(t *testing.T) {
	c := setupTest([]string{"context", "server", "--https", "--key", "a", "--port", "aaaa", "--url", "/", "--name", "etsi"})

	err := contextServer(c, c.Ngsi, c.Client)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestServerHander(t *testing.T) {
	h := &atContextServerHandler{context: "{\"context\":\"http://context\"}"}

	req := httptest.NewRequest(http.MethodGet, "http://receiver/", nil)

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestServerHanderErrorStatusMethodNotAllowed(t *testing.T) {
	c := setupTest([]string{"context", "server", "--name", "etsi"})

	h := &atContextServerHandler{ngsi: c.Ngsi, context: "{\"context\":\"http://context\"}"}

	req := httptest.NewRequest(http.MethodPost, "http://receiver/", nil)

	got := httptest.NewRecorder()

	h.ServeHTTP(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}
