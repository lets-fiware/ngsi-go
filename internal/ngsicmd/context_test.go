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

package ngsicmd

import (
	"bytes"
	"errors"
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestContextList(t *testing.T) {
	_, set, app, buf := setupTest()
	setupFlagString(set, "name")

	c := cli.NewContext(app, set, nil)
	err := contextList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "etsi https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\nld https://schema.lab.fiware.org/ld/context\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListJSON(t *testing.T) {
	_, set, app, buf := setupTest3()

	setupFlagString(set, "name,json")
	c := cli.NewContext(app, set, nil)
	set.Parse([]string{"--name=fiware", "--json={}"})
	err := contextAdd(c)
	assert.NoError(t, err)

	set = flag.NewFlagSet("test", 0)
	c = cli.NewContext(app, set, nil)
	err = contextList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "array [\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"]\ndata-model http://context-provider:3000/data-models/ngsi-context.jsonld\netsi https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\nld https://schema.lab.fiware.org/ld/context\nobject {\"ld\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld\"}\ntutorial http://context-provider:3000/data-models/ngsi-context.jsonld\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListName(t *testing.T) {
	_, set, app, buf := setupTest()
	setupFlagString(set, "name")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi"})
	err := contextList(c)

	if assert.NoError(t, err) {
		actual := buf.String()
		expected := "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\n"
		assert.Equal(t, expected, actual)
	}
}

func TestContextListErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextListErrorName(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestContextListErrorJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest3()

	setupFlagString(set, "name,json")
	c := cli.NewContext(app, set, nil)
	set.Parse([]string{"--name=fiware", "--json={}"})
	err := contextAdd(c)
	assert.NoError(t, err)

	JSONEncodeErr(ngsi, 2)

	set = flag.NewFlagSet("test", 0)
	c = cli.NewContext(app, set, nil)

	err = contextList(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestContextAdd(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware", "--url=http://fiware"})
	err := contextAdd(c)

	if assert.NoError(t, err) {
		_, set, app, _ := setupTest()
		setupFlagString(set, "name")
		c := cli.NewContext(app, set, nil)
		_ = set.Parse([]string{"--name=fiware"})
		_ = contextDelete(c)
	}
}

func TestContextAddErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextAddErrorName(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "name not found", ngsiErr.Message)
	}
}

func TestContextAddErrorNameString(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=@fiware"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "name error @fiware", ngsiErr.Message)
	}
}

func TestContextAddErrorUrl(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "url or json not provided", ngsiErr.Message)
	}
}

func TestContextAddErrorUrlJSON(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url,json")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware", "--url=http://context", "--json={}"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "specify either url or json", ngsiErr.Message)
	}
}

func TestContextAddErrorUrlError(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware", "--url=abc"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestContextAddErrorJSONError(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	setupFlagString(set, "name,json")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware", "--json=http://context"})
	err := contextAdd(c)
	b := false
	ngsi.JSONConverter = &MockJSONLib{ValidErr: &b}

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "url error", ngsiErr.Message)
	}
}

func TestContextAddError(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi", "--url=http://context"})
	err := contextAdd(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "etsi already exists", ngsiErr.Message)
	}
}

func TestContextUpdate(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi", "--url=http://fiware"})
	err := contextUpdate(c)

	assert.NoError(t, err)
}

func TestContextUpdateErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextUpdateErrorName(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "name not found", ngsiErr.Message)
	}
}

func TestContextUpdateErrorUrl(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "url not found", ngsiErr.Message)
	}
}

func TestContextUpdateErrorUrlError(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi", "--url=abc"})
	err := contextUpdate(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "abc is not url", ngsiErr.Message)
	}
}

func TestContextDelete(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=etsi", "--url=http://fiware"})
	err := contextDelete(c)

	assert.NoError(t, err)
}

func TestContextDeleteErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextDeleteErrorName(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	err := contextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "name not found", ngsiErr.Message)
	}
}

func TestContextDeleteErrorUrl(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,url")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--name=fiware"})
	err := contextDelete(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestGetAtContext(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	_, err := getAtContext(ngsi, "{}")

	assert.NoError(t, err)
}

func TestGetAtContextHTTP(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	actual, err := getAtContext(ngsi, "etsi")

	if assert.NoError(t, err) {
		assert.Equal(t, "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld", actual)
	}
}

func TestGetAtContextErrorNotFound(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	_, err := getAtContext(ngsi, "fiware")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestGetAtContextErrorNotJSON(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	_, err := getAtContext(ngsi, "1")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "data not json: 1", ngsiErr.Message)
	}
}

func TestGetAtContextErrorJSON(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	JSONDecodeErr(ngsi, 0)

	_, err := getAtContext(ngsi, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContext(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	payload := []byte(`{"id":"I"}`)

	cases := []struct {
		context  string
		expected string
	}{
		{
			context:  `https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld`,
			expected: "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"id\":\"I\"}",
		},
		{
			context:  "[\"http://example.org/ngsi-ld/latest/parking.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"]",
			expected: "{\"@context\":[\"http://example.org/ngsi-ld/latest/parking.jsonld\",\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\"],\"id\":\"I\"}",
		},
		{
			context:  "{\"parking\":\"http://example.org/ngsi-ld/latest/parking.jsonld\"}",
			expected: "{\"@context\":{\"parking\":\"http://example.org/ngsi-ld/latest/parking.jsonld\"},\"id\":\"I\"}",
		},
		{
			context:  "etsi",
			expected: "{\"@context\":\"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld\",\"id\":\"I\"}",
		},
	}

	for _, c := range cases {
		actual, err := insertAtContext(ngsi, payload, c.context)

		if assert.NoError(t, err) {
			assert.Equal(t, c.expected, string(actual))
		}
	}
}

func TestInsertAtContextErrorPayload(t *testing.T) {
	ngsi, _, _, _ := setupTest()

	payload := []byte(`context`)
	_, err := insertAtContext(ngsi, payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "data not json", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorArrayUnmarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	JSONDecodeErr(ngsi, 1)

	payload := []byte(`[]`)
	_, err := insertAtContext(ngsi, payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorArrayMarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	JSONEncodeErr(ngsi, 0)

	payload := []byte(`[]`)
	_, err := insertAtContext(ngsi, payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorObjectUnmarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	JSONDecodeErr(ngsi, 1)

	payload := []byte(`{}`)
	_, err := insertAtContext(ngsi, payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestInsertAtContextErrorObjectMarshal(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	JSONEncodeErr(ngsi, 0)

	payload := []byte(`{}`)
	_, err := insertAtContext(ngsi, payload, "{}")

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "json error", ngsiErr.Message)
	}
}

func TestContextServer(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,name")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--name=etsi"})
	err := contextServer(c)

	assert.NoError(t, err)
}

func TestContextServerHTTPS(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,key,cert,name")
	setupFlagBool(set, "https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--https", "--key=test.key", "--cert=test.cert", "--port=aaaa", "--url=/context", "--name=etsi"})
	err := contextServer(c)

	assert.NoError(t, err)
}

func TestContextServerJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest3()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	set = flag.NewFlagSet("test", 0)
	setupFlagString(set, "port,url,name")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--name=array"})
	err := contextServer(c)

	assert.NoError(t, err)
}

func TestContextServerDataJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--data={}"})
	err := contextServer(c)

	assert.NoError(t, err)
}

func TestContextServerDataHTTP(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--data=http://context"})
	err := contextServer(c)

	assert.NoError(t, err)
}

func TestContextServerDataFile(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,data")

	ngsi.FileReader = &MockFileLib{ReadFileData: []byte("{}")}
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--data=@file"})
	err := contextServer(c)

	assert.NoError(t, err)
}

func TestContextServerErrorInitCmd(t *testing.T) {
	_, set, app, _ := setupTest()
	setupFlagString(set, "name,syslog")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--syslog="})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "syslog logLevel error", ngsiErr.Message)
	}
}

func TestContextServerErrorNoArgs(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,name,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "name or data  not found", ngsiErr.Message)
	}
}

func TestContextServerErrorTooMuchArgs(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,name,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--name=etsi", "--data=@file"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "specify either name or data", ngsiErr.Message)
	}
}

func TestContextServerErrorNotFoundName(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,name,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--name=fiware"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestContextServerErrorFilePathAbs(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,data")

	ngsi.FileReader = &MockFileLib{FilePathAbsError: errors.New("filePathAbsError")}
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--data=@file"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "filePathAbsError", ngsiErr.Message)
	}
}

func TestContextServerErrorReadFileError(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,data")

	ngsi.FileReader = &MockFileLib{ReadFileError: errors.New("readFileError")}
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--data=@file"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 6, ngsiErr.ErrNo)
		assert.Equal(t, "readFileError", ngsiErr.Message)
	}
}

func TestContextServerErrorFileNameError(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,data")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--data=@"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 7, ngsiErr.ErrNo)
		assert.Equal(t, "file name error", ngsiErr.Message)
	}
}

func TestContextServerErrorNotJSON(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,data")

	ngsi.FileReader = &MockFileLib{ReadFileData: []byte("context")}
	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--port=aaaa", "--url=/context", "--data=@file"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 8, ngsiErr.ErrNo)
		assert.Equal(t, "data not json", ngsiErr.Message)
	}
}

func TestContextServerErrorKey(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,key,cert,name")
	setupFlagBool(set, "https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--https", "--port=aaaa", "--url=/", "--name=etsi"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 9, ngsiErr.ErrNo)
		assert.Equal(t, "no key file provided", ngsiErr.Message)
	}
}

func TestContextServerErrorCert(t *testing.T) {
	ngsi, set, app, _ := setupTest()
	buf := new(bytes.Buffer)
	ngsi.Stderr = buf

	setupFlagString(set, "port,url,key,cert,name")
	setupFlagBool(set, "https")

	c := cli.NewContext(app, set, nil)
	_ = set.Parse([]string{"--https", "--key=a", "--port=aaaa", "--url=/", "--name=etsi"})
	err := contextServer(c)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsiCmdError)
		assert.Equal(t, 10, ngsiErr.ErrNo)
		assert.Equal(t, "no cert file provided", ngsiErr.Message)
	}
}

func TestServerHander(t *testing.T) {
	serverGlobal = &serverParam{context: "{\"context\":\"http://context\"}"}

	req := httptest.NewRequest(http.MethodGet, "http://receiver/", nil)

	got := httptest.NewRecorder()

	serverHandler(got, req)

	expected := http.StatusOK

	assert.Equal(t, expected, got.Code)
}

func TestServerHanderErrorStatusMethodNotAllowed(t *testing.T) {
	ngsi, _, _, _ := setupTest()
	serverGlobal = &serverParam{ngsi: ngsi, context: "{\"context\":\"http://context\"}"}

	req := httptest.NewRequest(http.MethodPost, "http://receiver/", nil)

	got := httptest.NewRecorder()

	serverHandler(got, req)

	expected := http.StatusMethodNotAllowed

	assert.Equal(t, expected, got.Code)
}
