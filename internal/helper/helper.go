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

package helper

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

var configData = `{
	"version": "1",
	"servers": {
	  "orion": {
		"serverHost": "https://orion",
		"ngsiType": "v2",
		"serverType": "broker"
	  },
	  "orion-ld": {
		"serverHost": "https://orion-ld",
		"ngsiType": "ld",
		"serverType": "broker"
	  },
	  "orion-alias": {
		"serverHost": "orion-ld",
		"ngsiType": "ld",
		"serverType": "broker"
	  },
	  "comet": {
		"serverHost": "https://comet",
		"serverType": "comet"
	  },
	  "cygnus": {
		"serverHost": "https://cygnus",
		"serverType": "cygnus"
	  },
	  "ql": {
		"serverHost": "https://quantumleap",
		"serverType": "quantumleap"
	  },
	  "iota": {
		"serverHost": "https://iota",
		"serverType": "iota"
	  },
	  "perseo": {
		"serverHost": "https://perseo",
		"serverType": "perseo"
	  },
	  "perseo-core": {
		"serverHost": "https://perseo-core",
		"serverType": "perseo-core"
	  },
	  "keyrock": {
		"serverHost": "https://keyrock",
		"serverType": "keyrock"
	  },
	  "wirecloud": {
		"serverHost": "https://wirecloud",
		"serverType": "wirecloud"
	  },
	  "scorpio": {
		"serverHost": "https://scorpio:9090",
		"ngsiType": "ld",
		"serverType": "broker",
		"brokerType": "scorpio"
	  },
	  "regproxy": {
		"serverHost": "https://regproxy",
		"serverType": "regproxy"
	  },
	  "tokenproxy": {
		"serverHost": "https://tokenproxy",
		"serverType": "tokenproxy"
	  },
	  "queryproxy": {
		"serverHost": "https://queryproxy",
		"serverType": "queryproxy"
	  }
	},
	"contexts": {
	  "data-model": "http://context-provider:3000/data-models/ngsi-context.jsonld",
	  "etsi": "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld",
	  "ld": "https://schema.lab.fiware.org/ld/context",
	  "tutorial": "http://context-provider:3000/data-models/ngsi-context.jsonld",
	  "array": [
		"https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"
	  ],
	  "object": {
		"ld": "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"
	  }
	},
	"settings": {
		"usePreviousArgs": true
	}
  }`

func SetupTestInitNGSI() *ngsilib.NGSI {
	c := SetupTestWithConfigAndCache(nil, nil, "", "")
	return c.Ngsi
}

func SetupTestInitCmd(callback func(*ngsilib.NGSI)) *ngsicli.Context {
	ngsi := SetupTestInitNGSI()

	if callback != nil {
		callback(ngsi)
	}
	c := &ngsicli.Context{Ngsi: ngsi}
	_, err := ngsicli.InitCmd(c)
	if err != nil {
		return nil // panic(err)
	}

	return c
}

func SetupTest(app *ngsicli.App, args []string) *ngsicli.Context {
	return SetupTestWithConfig(app, args, "")
}

func SetupTestWithConfig(app *ngsicli.App, args []string, config string) *ngsicli.Context {
	return SetupTestWithConfigAndCache(app, args, config, "")
}

func SetupTestWithConfigAndCache(app *ngsicli.App, args []string, config, cache string) *ngsicli.Context {
	ngsilib.Reset()

	if config == "" {
		config = configData
	}
	filename := ""
	ngsi := ngsilib.NewNGSI()

	ngsi.ConfigFile = &MockIoLib{}
	ngsi.ConfigFile.SetFileName(&filename)
	ngsi.FileReader = &MockFileLib{ReadFileData: []byte(config)}

	if cache == "" {
		ngsi.CacheFile = &MockIoLib{}
		ngsi.CacheFile.SetFileName(&filename)
	} else {
		filename := "ngsi-cache-token.json"
		ngsi.CacheFile = &MockIoLib{Tokens: &cache, Filename: &filename}
	}

	ngsi.HTTP = NewMockHTTP()
	ngsi.NetLib = &MockNetLib{}

	buffer := &bytes.Buffer{}
	stderrBuffer := &bytes.Buffer{}
	logBuffer := &bytes.Buffer{}

	ngsi.StdWriter = buffer
	ngsi.Stderr = stderrBuffer
	ngsi.LogWriter = logBuffer

	ngsi.FilePath = &MockFilePathLib{}
	ngsi.Ioutil = &MockIoutilLib{}
	ngsi.ZipLib = &MockZipLib{}

	ngsi.ReadAll = MockReadAll
	ngsi.GetReader = MockGetReader

	ngsi.PreviousArgs = &ngsilib.Settings{}

	if app == nil && args == nil {
		return &ngsicli.Context{Ngsi: ngsi}
	}

	args = append([]string{"ngsi"}, args...)

	_, c, err := app.Parse(args)
	if err != nil {
		s := strings.TrimRight(ngsierr.Message(err), "\n") + "\n"
		fmt.Fprintln(os.Stderr, s)
		for err != nil {
			err = errors.Unwrap(err)
			if err == nil {
				break
			}
			s = strings.TrimRight(ngsierr.Message(err), "\n")
			fmt.Fprintf(os.Stderr, "%T %s\n", err, s)
		}
		fmt.Fprintln(os.Stderr, stderrBuffer.String())
		fmt.Fprintf(os.Stderr, "args: %s\n\n", strings.Join(args, " "))
		return nil
	}

	ngsi.LogWriter = logBuffer

	if c.Client != nil {
		c.Client.HTTP = NewMockHTTP()
	}
	if c.Client2 != nil {
		c.Client2.HTTP = NewMockHTTP()
	}

	return c
}

func GetStdoutString(c *ngsicli.Context) string {
	return c.Ngsi.StdWriter.(*bytes.Buffer).String()
}

func GetStderrString(c *ngsicli.Context) string {
	return c.Ngsi.Stderr.(*bytes.Buffer).String()
}

func StrPtr(s string) *string {
	return &s
}

func UrlParse(s string) *url.URL {
	u, _ := url.Parse(s)
	return u
}
