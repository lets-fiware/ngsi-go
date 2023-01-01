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

package ngsilib

import "github.com/lets-fiware/ngsi-go/internal/ngsierr"

// CreateServer is ...
func (ngsi *NGSI) CreateServer(name string, brokerParam map[string]string) error {
	const funcName = "CreateServer"

	broker := new(Server)
	if err := setServerParam(broker, brokerParam); err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}

	if err := ngsi.checkAllParams(broker); err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	ngsi.ServerList[name] = broker

	if err := ngsi.saveConfigFile(); err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	return nil
}

// UpdateServer is ...
func (ngsi *NGSI) UpdateServer(host string, brokerParam map[string]string) error {
	const funcName = "UpdateServer"

	if broker, ok := ngsi.ServerList[host]; ok {
		if err := setServerParam(broker, brokerParam); err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		if err := ngsi.checkAllParams(broker); err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
		if err := ngsi.saveConfigFile(); err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	} else {
		return ngsierr.New(funcName, 4, host+" not found", nil)
	}
	return nil
}

// DeleteServer is ...
func (ngsi *NGSI) DeleteServer(host string) error {
	const funcName = "DeleteServer"

	if _, ok := ngsi.ServerList[host]; ok {
		delete(ngsi.ServerList, host)
		if err := ngsi.saveConfigFile(); err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
	} else {
		return ngsierr.New(funcName, 2, host+" not found", nil)
	}
	return nil
}
