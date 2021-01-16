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
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client is
type Client struct {
	Broker        *Broker
	URL           *url.URL
	Headers       map[string]string
	Token         string
	Tenant        string
	Scope         string
	APIPathBefore string
	APIPathAfter  string
	NgsiType      int
	SafeString    bool
	XAuthToken    bool
	Link          *string
	HTTP          HTTPRequest
}

const (
	ngsiV2 = iota
	ngsiLd
)

// InitHeader is ...
func (client *Client) InitHeader() error {
	const funcName = "InitHeader"
	client.Headers = make(map[string]string)

	if client.Token != "" {
		if client.XAuthToken {
			client.Headers["X-Auth-Token"] = client.Token
		} else {
			client.Headers["Authorization"] = "Bearer " + client.Token
		}
	}
	if client.Tenant != "" {
		if err := client.CheckTenant(client.Tenant); err != nil {
			return &NgsiLibError{funcName, 1, err.Error(), err}
		}
		client.Headers["Fiware-Service"] = client.Tenant
	}
	if client.NgsiType == ngsiV2 {
		if client.Scope != "" {
			if err := client.CheckScope(client.Scope); err != nil {
				return &NgsiLibError{funcName, 2, err.Error(), err}
			}
			client.Headers["Fiware-ServicePath"] = client.Scope
		}
	}
	if client.NgsiType == ngsiLd {
		if client.Link != nil {
			client.Headers["link"] =
				fmt.Sprintf(`<%s>; rel="http://www.w3.org/ns/json-ld#context"; type="application/ld+json"`, *(client.Link))
		}
	}

	client.Headers["Accept"] = "*/*"

	return nil
}

// SetHeaders is ...
func (client *Client) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		client.Headers[key] = value
	}
}

// SetHeader is ...
func (client *Client) SetHeader(key string, value string) {
	client.Headers[key] = value
}

// RemoveHeader is ...
func (client *Client) RemoveHeader(key string) {
	_, ok := client.Headers[key]
	if ok {
		delete(client.Headers, key)
	}
}

// SetContentType is ...
func (client *Client) SetContentType() {
	client.Headers["Content-Type"] = "application/json"
	if client.NgsiType == ngsiLd && client.Link == nil {
		client.Headers["Content-Type"] = "application/ld+json"
	}
}

// SetAcceptJSON is ...
func (client *Client) SetAcceptJSON() {
	client.Headers["Accept"] = "application/json"
}

// SetPath is ...
func (client *Client) SetPath(path string) {
	if !hasPrefix([]string{"/version", "/admin", "/log", "/statistics", "/cache"}, path) {
		if client.NgsiType == ngsiLd {
			path = "/ngsi-ld/v1" + path
		} else {
			path = "/v2" + path
		}
	}
	if client.APIPathBefore != "" {
		if strings.HasPrefix(path, client.APIPathBefore) {
			path = client.APIPathAfter + "/" + path[len(client.APIPathBefore):]
		}
	}
	client.URL.Path = path
}

func hasPrefix(prefixes []string, path string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}

// SetQuery is ...
func (client *Client) SetQuery(values *url.Values) {
	client.URL.RawQuery = (*values).Encode()
}

// IsSafeString is ...
func (client *Client) IsSafeString() bool {
	return client.SafeString
}

// IsNgsiV2 is
func (client *Client) IsNgsiV2() bool {
	return client.NgsiType == ngsiV2
}

// IsNgsiLd is
func (client *Client) IsNgsiLd() bool {
	return client.NgsiType == ngsiLd
}

// ResultsCount is ...
func (client *Client) ResultsCount(res *http.Response) (int, error) {
	if client.IsNgsiLd() {
		return strconv.Atoi(res.Header.Get("Ngsild-Results-Count"))
	}
	return strconv.Atoi(res.Header.Get("Fiware-Total-Count"))
}

func (client *Client) idmURL() string {
	tokenURL := client.Broker.IdmHost
	if strings.HasPrefix(tokenURL, "http") {
		return tokenURL
	}
	baseURL, _ := url.Parse(client.Broker.BrokerHost)
	baseURL.Path = client.Broker.IdmHost
	return baseURL.String()
}

func (client *Client) storeToken(token string) {
	client.Token = token
}

func (client *Client) getExpiresIn() int64 {
	return 3600
}

// CheckTenant is ...
func (client *Client) CheckTenant(tenant string) error {
	const funcName = "CheckTenant"

	if isTenantString(tenant) {
		return nil
	}
	return &NgsiLibError{funcName, 1, fmt.Sprintf("error FIWARE Service: %s", tenant), nil}
}

// CheckScope is ...
func (client *Client) CheckScope(scope string) error {
	const funcName = "CheckScope"

	if isScopeString(scope) {
		return nil
	}
	return &NgsiLibError{funcName, 1, fmt.Sprintf("error FIWARE ServicePath: %s", scope), nil}
}
