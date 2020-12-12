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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type receiverParam struct {
	ngsi   *ngsilib.NGSI
	pretty bool
}

var receiverGlobal *receiverParam

func receiver(c *cli.Context) error {
	const funcName = "receiver"

	ngsi, err := initCmd(c, funcName, false)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	host := c.String("host")
	port := c.String("port")
	addr := host + ":" + port

	path := c.String("url")
	url := addr + path

	pretty := c.IsSet("pretty")

	if c.Bool("https") {
		if !c.IsSet("key") {
			return &ngsiCmdError{funcName, 2, "no key file provided", nil}
		}
		if !c.IsSet("cert") {
			return &ngsiCmdError{funcName, 3, "no cert file provided", nil}
		}
		url = "https://" + url
	} else {
		url = "http://" + url
	}

	receiverGlobal = &receiverParam{ngsi: ngsi, pretty: pretty}

	http.HandleFunc(path, http.HandlerFunc(receiverHandler))

	if c.Bool("verbose") {
		fmt.Fprintf(ngsi.Stderr, "%s\n", url)
	}

	if c.Bool("https") {
		http.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), nil)
	} else {
		http.ListenAndServe(addr, nil)
	}

	return nil
}

func receiverHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusNoContent

	switch r.Method {
	default:
		fmt.Fprint(receiverGlobal.ngsi.Stderr, "Method not allowed.\n")
		status = http.StatusMethodNotAllowed
	case http.MethodPost:
		body := r.Body
		defer body.Close()
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		b := buf.Bytes()
		if receiverGlobal.pretty && ngsilib.IsJSON(b) {
			var j interface{}
			err := json.Unmarshal(b, &j)
			if err != nil {
				fmt.Fprintf(receiverGlobal.ngsi.Stderr, "json.Unmarshal error\n")
			}
			b, _ = json.MarshalIndent(j, "", "  ")
		}
		fmt.Fprint(receiverGlobal.ngsi.StdWriter, string(b)+"\n")
	}
	w.WriteHeader(status)
}
