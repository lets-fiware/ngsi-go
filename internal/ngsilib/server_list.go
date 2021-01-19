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

// ServerList is ...
type ServerList map[string]*Server

// AllServersList is ...
func (ngsi *NGSI) AllServersList() *ServerList {
	return &ngsi.serverList
}

// InitServerList is ...
func InitServerList() {
	gNGSI.serverList = make(ServerList)
}

// List is ...
func (info *ServerList) List() string {
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
func (info *ServerList) BrokerInfo(name string) (*Server, error) {
	const funcName = "BrokerInfo"

	client, ok := gNGSI.serverList[name]
	if ok {
		if client.ServerType == "broker" {
			return client, nil
		}
		return nil, &LibError{funcName, 1, fmt.Sprintf("host found: %s, but type is %s", name, client.ServerType), nil}
	}
	return nil, &LibError{funcName, 2, fmt.Sprintf("host not found: %s", name), nil}
}

// ServerInfo is ...
func (info *ServerList) ServerInfo(name, filter string) (*Server, error) {
	const funcName = "ServerInfo"

	client, ok := gNGSI.serverList[name]
	if ok {
		if filter == "" {
			if client.ServerType != "broker" {
				return client, nil
			}
			return nil, &LibError{funcName, 1, fmt.Sprintf("host found: %s, but type is %s", name, client.ServerType), nil}
		}
		if client.ServerType == filter {
			return client, nil
		}
	}
	return nil, &LibError{funcName, 2, fmt.Sprintf("host not found: %s", name), nil}
}

// BrokerInfoJSON is ...
func (info *ServerList) BrokerInfoJSON(name string) (*string, error) {
	const funcName = "BrokerInfoJSON"

	var s string
	if name == "" {
		json, err := JSONMarshal(info.BrokerList())
		if err != nil {
			return nil, &LibError{funcName, 1, "json.Marshl error", err}
		}
		s = string(json)
	} else {
		info, ok := gNGSI.serverList[name]
		if !ok {
			return nil, &LibError{funcName, 2, fmt.Sprintf("host not found: %s", name), nil}
		}
		if info.ServerType != "broker" {
			return nil, &LibError{funcName, 3, fmt.Sprintf("host found: %s, but type is %s", name, info.ServerType), nil}
		}
		json, err := JSONMarshal(info)
		if err != nil {
			return nil, &LibError{funcName, 4, "json.Marshl error", err}
		}
		s = string(json)
	}
	return &s, nil

}

// BrokerList is ...
func (info *ServerList) BrokerList() ServerList {
	list := ServerList{}

	for k, v := range gNGSI.serverList {
		if v.ServerType == "broker" {
			list[k] = v
		}
	}
	return list
}

// ServerInfoJSON is ...
func (info *ServerList) ServerInfoJSON(name, filter string) (*string, error) {
	const funcName = "ServerInfoJSON"

	var s string
	if name == "" {
		json, err := JSONMarshal(info.ServerList(filter, false))
		if err != nil {
			return nil, &LibError{funcName, 1, "json.Marshl error", err}
		}
		s = string(json)
	} else {
		info, ok := gNGSI.serverList[name]
		if !ok {
			return nil, &LibError{funcName, 2, fmt.Sprintf("host not found: %s", name), nil}
		}
		if info.ServerType == "broker" {
			return nil, &LibError{funcName, 3, fmt.Sprintf("host found: %s, but type is %s", name, info.ServerType), nil}
		}
		json, err := JSONMarshal(info)
		if err != nil {
			return nil, &LibError{funcName, 4, "json.Marshl error", err}
		}
		s = string(json)
	}
	return &s, nil

}

// ServerList is ...
func (info *ServerList) ServerList(filter string, all bool) ServerList {
	list := ServerList{}

	for k, v := range gNGSI.serverList {
		if v.ServerType != "broker" || all {
			if filter == "" || v.ServerType == filter {
				list[k] = v
			}
		}
	}
	return list
}
