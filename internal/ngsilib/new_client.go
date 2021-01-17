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
	"net/url"
	"strings"
)

// NewClient is ...
func (ngsi *NGSI) NewClient(name string, cmdFlags *CmdFlags, isHTTPVerb bool) (client *Client, err error) {
	const funcName = "NewClient"

	client = &Client{}
	client.Broker = &Broker{}
	client.HTTP = ngsi.HTTP

	if IsHTTP(name) {
		client.URL, err = url.Parse(name)
		if err != nil {
			return nil, &NgsiLibError{funcName, 1, fmt.Sprintf("illegal url: %s", name), nil}
		}
	} else {
		host, path, query := parseURL(name)

		broker, ok := ngsi.brokerList[host]
		if ok {
			client.Broker = broker
			host = client.Broker.BrokerHost
			if host == "" {
				return nil, &NgsiLibError{funcName, 2, "host not found", nil}
			}
			if !IsHTTP(host) {
				broker1, ok := ngsi.brokerList[host]
				if !ok {
					return nil, &NgsiLibError{funcName, 3, host + " not found", nil}
				}
				copyBrokerInfo(broker1, client.Broker)
				host = client.Broker.BrokerHost
				if !IsHTTP(host) {
					return nil, &NgsiLibError{funcName, 4, "url error: " + host, nil}
				}
			}
			if strings.HasSuffix(host, "/") {
				host = host[:len(host)-1]
			}
		} else {
			if !isIPAddress(host) && !isLocalHost(host) {
				return nil, &NgsiLibError{funcName, 5, "error host: " + host, nil}
			}
			host = "http://" + host
		}
		if !isHTTPVerb {
			path = ""
		}
		host += path
		if query != "" {
			host += query
		}
		client.URL, err = url.Parse(host)
		if err != nil {
			return nil, &NgsiLibError{funcName, 6, "illegal url: " + name + ", " + host, nil}
		}
	}

	var tenant *string
	var scope *string

	d := ngsi.GetPreviousArgs()

	if d.Tenant != "" {
		tenant = &d.Tenant
	}
	if d.Scope != "" {
		scope = &d.Scope
	}

	if cmdFlags.Tenant == nil {
		cmdFlags.Tenant = tenant
	}
	if cmdFlags.Scope == nil {
		cmdFlags.Scope = scope
	}
	setTenantAndScope(client, cmdFlags.Tenant, cmdFlags.Scope)

	if d.Tenant != client.Tenant {
		d.Tenant = client.Tenant
		ngsi.Updated = true
	}

	if d.Scope != client.Scope {
		d.Scope = client.Scope
		ngsi.Updated = true
	}

	if client.Broker != nil {
		if apiPath := client.Broker.APIPath; apiPath != "" {
			client.APIPathBefore, client.APIPathAfter, err = getAPIPath(apiPath)
			if err != nil {
				return nil, &NgsiLibError{funcName, 7, err.Error(), err}
			}
		}
		client.NgsiType = ngsiV2
		if ngsiType := client.Broker.NgsiType; ngsiType != "" {
			if Contains(ngsiLdTypes, strings.ToLower(ngsiType)) {
				client.NgsiType = ngsiLd
			}
		}
	}

	token := ""
	if d.Token != "" {
		token = d.Token
	}
	if cmdFlags.Token != nil && *cmdFlags.Token != token {
		token = *cmdFlags.Token
		d.Token = token
		ngsi.Updated = true
	}
	if token != "" {
		client.Token = token
	} else if client.Broker.IdmType != "" {
		token, err := ngsi.GetToken(client)
		if err != nil {
			return nil, &NgsiLibError{funcName, 8, err.Error(), err}
		}
		client.Token = token
	}

	b, err := client.Broker.safeString()
	if err != nil {
		return nil, &NgsiLibError{funcName, 9, err.Error(), err}
	}
	client.SafeString = b
	if cmdFlags.SafeString != nil {
		b, err := ngsi.BoolFlag(*cmdFlags.SafeString)
		if err != nil {
			return nil, &NgsiLibError{funcName, 10, err.Error(), err}
		}
		client.SafeString = b
	}

	client.XAuthToken = cmdFlags.XAuthToken
	client.Link = cmdFlags.Link

	if err = client.InitHeader(); err != nil {
		return nil, &NgsiLibError{funcName, 11, err.Error(), err}
	}

	if ngsi.Updated {
		if IsHTTP(ngsi.PreviousArgs.Host) {
			ngsi.PreviousArgs.Host = ""
			ngsi.PreviousArgs.Tenant = ""
			ngsi.PreviousArgs.Scope = ""
		}
		if err = ngsi.saveConfigFile(); err != nil {
			return nil, &NgsiLibError{funcName, 12, err.Error(), err}
		}
	}
	return client, nil
}

func setTenantAndScope(client *Client, tenant *string, scope *string) {

	client.Tenant = client.Broker.Tenant
	client.Scope = client.Broker.Scope

	if tenant != nil {
		client.Tenant = *tenant
	}
	if scope != nil {
		client.Scope = *scope
	}
}

func parseURL(url string) (string, string, string) {
	var host, path, query string

	pos := strings.Index(url, "?")
	if pos == -1 {
		host = url
		query = ""
	} else {
		host = url[:pos]
		query = url[pos+1:]
	}
	pos = strings.Index(host, "/")
	if pos != -1 {
		path = host[pos:]
		host = host[:pos]
	}
	return host, path, query
}
