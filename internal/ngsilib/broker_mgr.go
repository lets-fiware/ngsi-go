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

package ngsilib

// CreateBroker is ...
func (ngsi *NGSI) CreateBroker(name string, brokerParam map[string]string) error {
	const funcName = "CreateBroker"

	broker := new(Broker)
	setBrokerParam(broker, brokerParam)

	if err := ngsi.checkAllParams(broker); err != nil {
		return &NgsiLibError{funcName, 1, err.Error(), err}
	}

	ngsi.brokerList[name] = broker

	if err := ngsi.saveConfigFile(); err != nil {
		return &NgsiLibError{funcName, 2, err.Error(), err}
	}

	return nil
}

// UpdateBroker is ...
func (ngsi *NGSI) UpdateBroker(host string, brokerParam map[string]string) error {
	const funcName = "UpdateBroker"

	if broker, ok := ngsi.brokerList[host]; ok {
		if err := setBrokerParam(broker, brokerParam); err != nil {
			return &NgsiLibError{funcName, 1, err.Error(), err}
		}
		if err := ngsi.checkAllParams(broker); err != nil {
			return &NgsiLibError{funcName, 2, err.Error(), err}
		}
		if err := ngsi.saveConfigFile(); err != nil {
			return &NgsiLibError{funcName, 3, err.Error(), err}
		}
	} else {
		return &NgsiLibError{funcName, 4, host + " not found", nil}
	}
	return nil
}

// DeleteBroker is ...
func (ngsi *NGSI) DeleteBroker(host string) error {
	const funcName = "DeleteBroker"

	if _, ok := ngsi.brokerList[host]; ok {
		delete(ngsi.brokerList, host)
		if err := ngsi.saveConfigFile(); err != nil {
			return &NgsiLibError{funcName, 1, err.Error(), err}
		}
	} else {
		return &NgsiLibError{funcName, 2, host + " not found", nil}
	}
	return nil
}
