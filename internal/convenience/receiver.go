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

package convenience

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sort"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func receiver(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "receiver"

	host := c.String("host")
	port := c.String("port")
	addr := host + ":" + port

	path := c.String("url")
	url := addr + path

	pretty := c.Bool("pretty")
	header := c.Bool("header")

	if c.Bool("https") {
		if !c.IsSet("key") {
			return ngsierr.New(funcName, 1, "no key file provided", nil)
		}
		if !c.IsSet("cert") {
			return ngsierr.New(funcName, 2, "no cert file provided", nil)
		}
		url = "https://" + url
	} else {
		url = "http://" + url
	}

	mux := http.NewServeMux()
	mux.Handle(path, &receiverHandler{ngsi: ngsi, pretty: pretty, header: header})

	addrs, _ := ngsi.NetLib.InterfaceAddrs()
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
		err := ngsi.NetLib.ListenAndServeTLS(addr, c.String("cert"), c.String("key"), mux)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	} else {
		err := ngsi.NetLib.ListenAndServe(addr, mux)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}
	}

	return nil
}

type receiverHandler struct {
	ngsi   *ngsilib.NGSI
	pretty bool
	header bool
}

func (h *receiverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := http.StatusNoContent

	switch r.Method {
	default:
		fmt.Fprint(h.ngsi.Stderr, "Method not allowed.\n")
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
		h.ngsi.Logging(ngsilib.LogInfo, header)

		if h.header {
			for _, k := range key {
				fmt.Fprintf(h.ngsi.StdWriter, "%s: %s\n", k, r.Header.Get(k))
			}
			fmt.Fprintln(h.ngsi.StdWriter, "")
		}

		b := buf.Bytes()
		h.ngsi.Logging(ngsilib.LogInfo, string(b))

		if h.pretty && ngsilib.IsJSON(b) {
			newBuf := new(bytes.Buffer)
			err := json.Indent(newBuf, b, "", "  ")
			if err == nil {
				b = newBuf.Bytes()
			}
		}
		fmt.Fprintf(h.ngsi.StdWriter, "%s\n", string(b))
		h.ngsi.StdoutFlush()
	}
	w.WriteHeader(status)
}
