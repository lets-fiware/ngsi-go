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
	"strings"
)

// Broker is
type Broker struct {
	BrokerHost   string `json:"brokerHost,omitempty"`
	NgsiType     string `json:"ngsiType,omitempty"`
	APIPath      string `json:"apiPath,omitempty"`
	IdmType      string `json:"idmType,omitempty"`
	IdmHost      string `json:"idmHost,omitempty"`
	Token        string `json:"token,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	ClientID     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Context      string `json:"context,omitempty"`
	Tenant       string `json:"tenant,omitempty"`
	Scope        string `json:"scope,omitempty"`
	SafeString   string `json:"safeString,omitempty"`
	XAuthToken   string `json:"xAuthToken,omitempty"`
}

const (
	cBrokerHost        = "brokerHost"
	cNgsiType          = "ngsiType"
	cAPIPath           = "apiPath"
	cIdmType           = "idmType"
	cIdmHost           = "idmHost"
	cToken             = "token"
	cUsername          = "username"
	cPassword          = "password"
	cClientID          = "clientId"
	cClientSecret      = "clientSecret"
	cContext           = "context"
	cFiwareService     = "service"
	cFiwareServicePath = "path"
	cSafeString        = "safeString"
	cXAuthToken        = "xAuthToken"
)

const (
	cPasswordCredentials  = "password"
	cKeyrock              = "keyrock"
	cKeyrocktokenprovider = "keyrocktokenprovider"
	cTokenproxy           = "tokenproxy"
)

const (
	cNgsiV2     = "ngsi-v2"
	cNgsiv2     = "ngsiv2"
	cV2         = "v2"
	cNgsiLd     = "ngsi-ld"
	cLd         = "ld"
	cPathRoot   = "/"
	cPathV2     = "/v2"
	cPathNgsiLd = "/ngsi-ld"
)

var (
	brokerArgs = []string{cBrokerHost, cNgsiType, cAPIPath,
		cIdmType, cIdmHost, cToken, cUsername, cPassword, cClientID, cClientSecret,
		cContext, cFiwareService, cFiwareServicePath, cSafeString, cXAuthToken}
	idmTypes    = []string{cPasswordCredentials, cKeyrock, cKeyrocktokenprovider, cTokenproxy}
	ngsiV2Types = []string{cNgsiV2, cNgsiv2, cV2}
	ngsiLdTypes = []string{cNgsiLd, cLd}
	apiPaths    = []string{cPathRoot, cPathV2, cPathNgsiLd}
)

func (ngsi *NGSI) checkAllParams(host *Broker) error {
	const funcName = "checkAllParams"

	brokerHost := host.BrokerHost
	if brokerHost == "" {
		return &NgsiLibError{funcName, 1, "brokerHost not found", nil}
	}
	if !IsHTTP(brokerHost) {
		if _, ok := ngsi.brokerList[brokerHost]; !ok {
			return &NgsiLibError{funcName, 2, fmt.Sprintf("brokerHost error: %s", brokerHost), nil}
		}
	}

	if ngsiType := host.NgsiType; ngsiType != "" {
		ngsiType = strings.ToLower(ngsiType)
		if !(Contains(ngsiV2Types, ngsiType) || Contains(ngsiLdTypes, ngsiType)) {
			return &NgsiLibError{funcName, 3, fmt.Sprintf("%s not found", ngsiType), nil}
		}
	}

	if apiPath := host.APIPath; apiPath != "" {
		if _, _, err := getAPIPath(apiPath); err != nil {
			return &NgsiLibError{funcName, 4, err.Error(), err}
		}
	}

	err := checkIdmParams(host.IdmType, host.IdmHost, host.Username, host.Password,
		host.ClientID, host.ClientSecret)
	if err != nil {
		return &NgsiLibError{funcName, 5, err.Error(), err}
	}

	var client *Client
	if tenant := host.Tenant; tenant != "" {
		err = client.CheckTenant(tenant)
		if err != nil {
			return &NgsiLibError{funcName, 6, err.Error(), err}
		}
	}

	if scope := host.Scope; scope != "" {
		err = client.CheckScope(scope)
		if err != nil {
			return &NgsiLibError{funcName, 7, err.Error(), err}
		}
	}

	if _, err := host.safeString(); err != nil {
		return &NgsiLibError{funcName, 8, err.Error(), err}
	}

	return nil
}

func getAPIPath(apiPath string) (string, string, error) {
	const funcName = "getAPIPath"

	pos := strings.Index(apiPath, ",")
	if pos == -1 {
		return "", "", &NgsiLibError{funcName, 1, fmt.Sprintf("apiPath error: %s", apiPath), nil}
	}
	pathBefore := apiPath[:pos]
	if !Contains(apiPaths, pathBefore) {
		return "", "", &NgsiLibError{funcName, 2, fmt.Sprintf("apiPath error: %s", pathBefore), nil}
	}
	pathAfter := apiPath[pos+1:]
	if !strings.HasPrefix(pathAfter, "/") {
		return "", "", &NgsiLibError{funcName, 3, fmt.Sprintf("must start with '/': %s", pathAfter), nil}
	}
	if strings.HasSuffix(pathAfter, "/") {
		return "", "", &NgsiLibError{funcName, 4, fmt.Sprintf("trailing '/' is not required: %s", pathAfter), nil}
	}
	return pathBefore, pathAfter, nil
}

func checkIdmParams(idmType string, idmHost string, username string, password string,
	clientID string, clientSecret string) error {
	const funcName = "checkIdmParams"

	if idmType == "" {
		if !(idmHost == "" && username == "" && password == "" && clientID == "" && clientSecret == "") {
			return &NgsiLibError{funcName, 1, "required idmType not found", nil}
		}
		return nil
	}
	if !isIdmType(idmType) {
		return &NgsiLibError{funcName, 2, fmt.Sprintf("idmType error: %s", idmType), nil}
	}

	if idmHost == "" {
		return &NgsiLibError{funcName, 3, "required idmHost not found", nil}
	}

	if !(IsHTTP(idmHost) || strings.HasPrefix(idmHost, "/")) {
		return &NgsiLibError{funcName, 4, fmt.Sprintf("idmHost error: %s", idmHost), nil}
	}

	switch strings.ToLower(idmType) {
	case cKeyrock, cPasswordCredentials:
		if clientID == "" || clientSecret == "" {
			return &NgsiLibError{funcName, 5, fmt.Sprintf("clientID and clientSecret are needed"), nil}
		}
		fallthrough
	case cKeyrocktokenprovider, cTokenproxy:
		if username == "" && password != "" {
			return &NgsiLibError{funcName, 6, "username is needed", nil}
		}
	}
	return nil
}

// ExistsBrokerHost is ...
func (ngsi *NGSI) ExistsBrokerHost(host string) bool {
	const funcName = "ExistsBrokerHost"
	_, ok := ngsi.brokerList[host]
	return ok
}

// ServerInfoArgs is ...
func (ngsi *NGSI) ServerInfoArgs() []string {
	return brokerArgs
}

func copyBrokerInfo(from *Broker, to *Broker) {
	if from.BrokerHost != "" {
		to.BrokerHost = from.BrokerHost
	}
	if from.NgsiType != "" && to.NgsiType == "" {
		to.NgsiType = from.NgsiType
	}
	if from.APIPath != "" && to.APIPath == "" {
		to.APIPath = from.APIPath
	}
	if from.IdmType != "" && to.IdmType == "" {
		to.IdmType = from.IdmType
	}
	if from.IdmHost != "" && to.IdmHost == "" {
		to.IdmHost = from.IdmHost
	}
	if from.Token != "" && to.Token == "" {
		to.Token = from.Token
	}
	if from.Username != "" && to.Username == "" {
		to.Username = from.Username
	}
	if from.Password != "" && to.Password == "" {
		to.Password = from.Password
	}
	if from.ClientID != "" && to.ClientID == "" {
		to.ClientID = from.ClientID
	}
	if from.ClientSecret != "" && to.ClientSecret == "" {
		to.ClientSecret = from.ClientSecret
	}
	if from.Context != "" && to.Context == "" {
		to.Context = from.Context
	}
	if from.Tenant != "" && to.Tenant == "" {
		to.Tenant = from.Tenant
	}
	if from.Scope != "" && to.Scope == "" {
		to.Scope = from.Scope
	}
	if from.SafeString != "" && to.SafeString == "" {
		to.SafeString = from.SafeString
	}
	if from.XAuthToken != "" && to.XAuthToken == "" {
		to.XAuthToken = from.XAuthToken
	}
}
func setBrokerParam(broker *Broker, param map[string]string) error {
	const funcName = "setBrokerParam"

	for key, value := range param {
		switch key {
		default:
			return &NgsiLibError{funcName, 1, fmt.Sprintf("%s not found", key), nil}
		case cBrokerHost:
			broker.BrokerHost = value
		case cNgsiType:
			broker.NgsiType = value
		case cAPIPath:
			broker.APIPath = value
		case cIdmType:
			broker.IdmType = strings.ToLower(value)
		case cIdmHost:
			broker.IdmHost = value
		case cToken:
			broker.Token = value
		case cUsername:
			broker.Username = value
		case cPassword:
			broker.Password = value
		case cClientID:
			broker.ClientID = value
		case cClientSecret:
			broker.ClientSecret = value
		case cContext:
			broker.Context = value
		case cFiwareService:
			broker.Tenant = value
		case cFiwareServicePath:
			broker.Scope = value
		case cSafeString:
			broker.SafeString = value
		case cXAuthToken:
			broker.XAuthToken = value
		}
	}
	return nil
}

// DeleteItem is ...
func (ngsi *NGSI) DeleteItem(host string, item string) error {
	const funcName = "DeleteItem"

	broker, ok := ngsi.brokerList[host]
	if !ok {
		return &NgsiLibError{funcName, 1, fmt.Sprintf("%s not found", host), nil}
	}
	param := map[string]string{item: ""}

	err := setBrokerParam(broker, param)

	if err != nil {
		return &NgsiLibError{funcName, 2, err.Error(), nil}
	}
	return nil
}

// IsHostReferenced is ...
func (ngsi *NGSI) IsHostReferenced(host string) error {
	const funcName = "IsHostReferenced"

	for k, v := range ngsi.brokerList {
		value := v.BrokerHost
		if host == value {
			return &NgsiLibError{funcName, 1, fmt.Sprintf("%s is referenced in %s", host, k), nil}
		}
	}
	return nil
}

// IsContextReferenced is ...
func (ngsi *NGSI) IsContextReferenced(context string) error {
	const funcName = "IsContextReferenced"

	for k, v := range ngsi.brokerList {
		value := v.Context
		if context == value {
			return &NgsiLibError{funcName, 1, fmt.Sprintf("%s is referenced in %s", context, k), nil}
		}
	}
	return nil
}

func isIdmType(name string) bool {
	return Contains(idmTypes, strings.ToLower(name))
}

func (info *Broker) safeString() (bool, error) {
	const funcName = "safeString"

	value := info.SafeString
	b, err := gNGSI.BoolFlag(value)
	if err != nil {
		return false, &NgsiLibError{funcName, 1, err.Error(), err}
	}
	return b, nil
}

func (info *Broker) xAuthToken() (bool, error) {
	const funcName = "xAuthToken"

	value := info.XAuthToken
	b, err := gNGSI.BoolFlag(value)
	if err != nil {
		return false, &NgsiLibError{funcName, 1, err.Error(), err}
	}
	return b, nil
}
