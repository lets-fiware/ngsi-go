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
	"sort"
)

// BrokerList is ...
type BrokerList map[string]*Broker

// BrokerList is ...
func (ngsi *NGSI) BrokerList() *BrokerList {
	return &ngsi.brokerList
}

// InitBrokerList is ...
func InitBrokerList() {
	gNGSI.brokerList = make(BrokerList)
}

// List is ...
func (info *BrokerList) List() string {
	list := ""

	keys := make([]string, len(*info))
	i := 0
	for key := range *info {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	for _, key := range keys {
		list += key + " "
	}

	if len(list) != 0 {
		list = list[:len(list)-1]
	}
	return list
}

// BrokerInfo is ...
func (info *BrokerList) BrokerInfo(name string) (*Broker, error) {
	const funcName = "BrokerInfo"

	client, ok := gNGSI.brokerList[name]
	if ok {
		return client, nil
	}
	return nil, &NgsiLibError{funcName, 1, fmt.Sprintf("host not found: %s", name), nil}
}

// BrokerInfoJSON is ...
func (info *BrokerList) BrokerInfoJSON(name string) (*string, error) {
	const funcName = "BrokerInfoJSON"

	var s string
	if name == "" {
		json, err := JSONMarshal(gNGSI.brokerList)
		if err != nil {
			return nil, &NgsiLibError{funcName, 1, "json.Marshl error", err}
		}
		s = string(json)
	} else {
		info, ok := gNGSI.brokerList[name]
		if !ok {
			return nil, &NgsiLibError{funcName, 2, fmt.Sprintf("host not found: %s", name), nil}
		}
		json, err := JSONMarshal(info)
		if err != nil {
			return nil, &NgsiLibError{funcName, 3, "json.Marshl error", err}
		}
		s = string(json)
	}
	return &s, nil

}
