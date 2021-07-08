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

package ngsicmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sort"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type receiverParam struct {
	ngsi   *ngsilib.NGSI
	pretty bool
	header bool
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

	pretty := c.Bool("pretty")
	header := c.Bool("header")

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

	receiverGlobal = &receiverParam{ngsi: ngsi, pretty: pretty, header: header}

	mux := http.NewServeMux()
	mux.HandleFunc(path, http.HandlerFunc(receiverHandler))

	addrs, _ := gNetLib.InterfaceAddrs()
	ip := []string{}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = append(ip, ipnet.IP.String())
			}
		}
	}

	if c.Bool("verbose") {
		fmt.Fprintln(ngsi.Stderr, ip)
		fmt.Fprintf(ngsi.Stderr, "%s\n", url)
	}

	ngsi.Logging(ngsilib.LogInfo, url+"\n")

	if c.Bool("https") {
		err = gNetLib.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), mux)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
	} else {
		err = gNetLib.ListenAndServe(addr, mux)
		if err != nil {
			return &ngsiCmdError{funcName, 5, err.Error(), err}
		}
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
		defer func() { _ = body.Close() }()
		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, body)

		header := ""
		key := []string{}
		for k := range r.Header {
			key = append(key, k)
		}
		sort.Strings(key)

		for _, k := range key {
			header += fmt.Sprintf("%s:[%s] ", k, r.Header.Get(k))
		}
		header = "[" + strings.TrimSpace(header) + "]"
		receiverGlobal.ngsi.Logging(ngsilib.LogInfo, header)

		if receiverGlobal.header {
			for _, k := range key {
				fmt.Fprintf(receiverGlobal.ngsi.StdWriter, "%s: %s\n", k, r.Header.Get(k))
			}
			fmt.Fprintln(receiverGlobal.ngsi.StdWriter, "")
		}

		b := buf.Bytes()
		receiverGlobal.ngsi.Logging(ngsilib.LogInfo, string(b))

		if receiverGlobal.pretty && ngsilib.IsJSON(b) {
			newBuf := new(bytes.Buffer)
			err := json.Indent(newBuf, b, "", "  ")
			if err == nil {
				b = newBuf.Bytes()
			}
		}
		fmt.Fprintf(receiverGlobal.ngsi.StdWriter, "%s\n", string(b))
		receiverGlobal.ngsi.StdoutFlush()
	}
	w.WriteHeader(status)
}
