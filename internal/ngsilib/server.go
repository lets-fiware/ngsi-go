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

// Server is
type Server struct {
	ServerType           string `json:"serverType,omitempty"`
	DeprecatedBrokerHost string `json:"brokerHost,omitempty"`
	ServerHost           string `json:"serverHost,omitempty"`
	BrokerType           string `json:"brokerType,omitempty"`
	NgsiType             string `json:"ngsiType,omitempty"`
	APIPath              string `json:"apiPath,omitempty"`
	IdmType              string `json:"idmType,omitempty"`
	IdmHost              string `json:"idmHost,omitempty"`
	Token                string `json:"token,omitempty"`
	Username             string `json:"username,omitempty"`
	Password             string `json:"password,omitempty"`
	ClientID             string `json:"clientId,omitempty"`
	ClientSecret         string `json:"clientSecret,omitempty"`
	HeaderName           string `json:"headerName,omitempty"`
	HeaderValue          string `json:"headerValue,omitempty"`
	HeaderEnvValue       string `json:"headerEnvValue,omitempty"`
	TokenScope           string `json:"tokenScope,omitempty"`
	Context              string `json:"context,omitempty"`
	Tenant               string `json:"tenant,omitempty"`
	Scope                string `json:"scope,omitempty"`
	SafeString           string `json:"safeString,omitempty"`
	XAuthToken           string `json:"xAuthToken,omitempty"`
}

const (
	cServerType        = "serverType"
	cServerHost        = "serverHost"
	cBrokerHost        = "brokerHost"
	cBrokerType        = "brokerType"
	cNgsiType          = "ngsiType"
	cAPIPath           = "apiPath"
	cIdmType           = "idmType"
	cIdmHost           = "idmHost"
	cToken             = "token"
	cUsername          = "username"
	cPassword          = "password"
	cClientID          = "clientId"
	cClientSecret      = "clientSecret"
	cHeaderName        = "headerName"
	cHeaderValue       = "headerValue"
	cHeaderEnvValue    = "headerEnvValue"
	cTokenScope        = "tokenScope"
	cContext           = "context"
	cFiwareService     = "service"
	cFiwareServicePath = "path"
	cSafeString        = "safeString"
	cXAuthToken        = "xAuthToken"
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

const (
	cComet         = "comet"
	cCygnus        = "cygnus"
	cQuantumLeap   = "quantumleap"
	cIota          = "iota"
	cfiwareKeyrock = "keyrock"
	cPerseo        = "perseo"
	cPerseoCore    = "perseo-core"
	cWireCloud     = "wirecloud"
	cOrionLD       = "orion-ld"
	cScorpio       = "scorpio"
	cStellio       = "stellio"
	cRegProxy      = "regproxy"
	cTokenProxy    = "tokenproxy"
	cQueryProxy    = "queryproxy"
)

var (
	brokerArgs = []string{cServerType, cServerHost, cBrokerHost, cBrokerType, cNgsiType, cAPIPath,
		cIdmType, cIdmHost, cToken, cUsername, cPassword, cClientID, cClientSecret, cHeaderName, cHeaderValue, cHeaderEnvValue, cTokenScope,
		cContext, cFiwareService, cFiwareServicePath, cSafeString, cXAuthToken}
	brokerTypeArgs = []string{cOrionLD, cScorpio, cStellio}
	serverTypeArgs = []string{cComet, cCygnus, cQuantumLeap, cIota, cfiwareKeyrock, cPerseo, cPerseoCore, cWireCloud, cRegProxy, cTokenProxy, cQueryProxy}

	ngsiV2Types = []string{cNgsiV2, cNgsiv2, cV2}
	ngsiLdTypes = []string{cNgsiLd, cLd}
	apiPaths    = []string{cPathRoot, cPathV2, cPathNgsiLd}
)

func (ngsi *NGSI) checkAllParams(host *Server) error {
	const funcName = "checkAllParams"

	serverHost := host.ServerHost
	if serverHost == "" {
		return &LibError{funcName, 1, "host not found", nil}
	}
	if !IsHTTP(serverHost) {
		if _, ok := ngsi.serverList[serverHost]; !ok {
			return &LibError{funcName, 2, fmt.Sprintf("host error: %s", serverHost), nil}
		}
	}

	if host.ServerType == "" {
		host.ServerType = "broker"
	}

	if ngsiType := host.NgsiType; ngsiType != "" {
		ngsiType = strings.ToLower(ngsiType)
		if !(Contains(ngsiV2Types, ngsiType) || Contains(ngsiLdTypes, ngsiType)) {
			return &LibError{funcName, 3, fmt.Sprintf("%s not found", ngsiType), nil}
		}
		if Contains(ngsiV2Types, ngsiType) {
			if host.BrokerType != "" {
				return &LibError{funcName, 4, "can'n specify broker Type", nil}
			}
			host.NgsiType = cV2
			host.BrokerType = ""
		} else {
			host.NgsiType = cLd
			if host.BrokerType == "" {
				host.BrokerType = cOrionLD
			} else {
				host.BrokerType = strings.ToLower(host.BrokerType)
				if !Contains(brokerTypeArgs, host.BrokerType) {
					return &LibError{funcName, 5, fmt.Sprintf("brokerType Error: %s", host.BrokerType), nil}
				}
			}
		}
	}

	if apiPath := host.APIPath; apiPath != "" {
		if _, _, err := getAPIPath(apiPath); err != nil {
			return &LibError{funcName, 6, err.Error(), err}
		}
	}

	idmParams := &IdmParams{
		IdmType:        host.IdmType,
		IdmHost:        host.IdmHost,
		Username:       host.Username,
		Password:       host.Password,
		ClientID:       host.ClientID,
		ClientSecret:   host.ClientSecret,
		HeaderName:     host.HeaderName,
		HeaderValue:    host.HeaderValue,
		HeaderEnvValue: host.HeaderEnvValue,
		TokenScope:     host.TokenScope,
	}
	err := checkIdmParams(idmParams)
	if err != nil {
		return &LibError{funcName, 7, err.Error(), err}
	}

	var client *Client
	if tenant := host.Tenant; tenant != "" {
		err = client.CheckTenant(tenant)
		if err != nil {
			return &LibError{funcName, 8, err.Error(), err}
		}
	}

	if scope := host.Scope; scope != "" {
		err = client.CheckScope(scope)
		if err != nil {
			return &LibError{funcName, 9, err.Error(), err}
		}
	}

	if _, err := host.safeString(); err != nil {
		return &LibError{funcName, 10, err.Error(), err}
	}

	return nil
}

func getAPIPath(apiPath string) (string, string, error) {
	const funcName = "getAPIPath"

	pos := strings.Index(apiPath, ",")
	if pos == -1 {
		return "", "", &LibError{funcName, 1, fmt.Sprintf("apiPath error: %s", apiPath), nil}
	}
	pathBefore := apiPath[:pos]
	if !Contains(apiPaths, pathBefore) {
		return "", "", &LibError{funcName, 2, fmt.Sprintf("apiPath error: %s", pathBefore), nil}
	}
	pathAfter := apiPath[pos+1:]
	if !strings.HasPrefix(pathAfter, "/") {
		return "", "", &LibError{funcName, 3, fmt.Sprintf("must start with '/': %s", pathAfter), nil}
	}
	if strings.HasSuffix(pathAfter, "/") {
		return "", "", &LibError{funcName, 4, fmt.Sprintf("trailing '/' is not required: %s", pathAfter), nil}
	}
	return pathBefore, pathAfter, nil
}

// ExistsBrokerHost is ...
func (ngsi *NGSI) ExistsBrokerHost(host string) bool {
	_, ok := ngsi.serverList[host]
	return ok
}

// ServerInfoArgs is ...
func (ngsi *NGSI) ServerInfoArgs() []string {
	return brokerArgs
}

// ServerTypeArgs is ...
func (ngsi *NGSI) ServerTypeArgs() []string {
	return serverTypeArgs
}

// BrokerTypeArgs is ...
func (ngsi *NGSI) BrokerTypeArgs() []string {
	return brokerTypeArgs
}

func copyServerInfo(from *Server, to *Server) {
	if from.ServerType != "" {
		to.ServerType = from.ServerType
	}
	if from.ServerHost != "" {
		to.ServerHost = from.ServerHost
	}
	if from.NgsiType != "" && to.NgsiType == "" {
		to.NgsiType = from.NgsiType
	}
	if from.BrokerType != "" {
		to.BrokerType = from.BrokerType
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
	if from.HeaderName != "" && to.HeaderName == "" {
		to.HeaderName = from.HeaderName
	}
	if from.HeaderValue != "" && to.HeaderValue == "" {
		to.HeaderValue = from.HeaderValue
	}
	if from.HeaderEnvValue != "" && to.HeaderEnvValue == "" {
		to.HeaderEnvValue = from.HeaderEnvValue
	}
	if from.TokenScope != "" && to.TokenScope == "" {
		to.TokenScope = from.TokenScope
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
func setServerParam(broker *Server, param map[string]string) error {
	const funcName = "setServerParam"

	for key, value := range param {
		switch key {
		default:
			return &LibError{funcName, 1, fmt.Sprintf("%s not found", key), nil}
		case cServerType:
			broker.ServerType = value
		case cServerHost:
			broker.ServerHost = value
		case cBrokerHost:
			broker.ServerHost = value
		case cNgsiType:
			broker.NgsiType = value
		case cBrokerType:
			broker.BrokerType = value
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
		case cHeaderName:
			broker.HeaderName = value
		case cHeaderValue:
			broker.HeaderValue = value
		case cHeaderEnvValue:
			broker.HeaderEnvValue = value
		case cTokenScope:
			broker.TokenScope = value
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

	broker, ok := ngsi.serverList[host]
	if !ok {
		return &LibError{funcName, 1, fmt.Sprintf("%s not found", host), nil}
	}
	param := map[string]string{item: ""}

	err := setServerParam(broker, param)

	if err != nil {
		return &LibError{funcName, 2, err.Error(), err}
	}
	return nil
}

// IsHostReferenced is ...
func (ngsi *NGSI) IsHostReferenced(host string) error {
	const funcName = "IsHostReferenced"

	for k, v := range ngsi.serverList {
		value := v.ServerHost
		if host == value {
			return &LibError{funcName, 1, fmt.Sprintf("%s is referenced in %s", host, k), nil}
		}
	}
	return nil
}

// IsContextReferenced is ...
func (ngsi *NGSI) IsContextReferenced(context string) error {
	const funcName = "IsContextReferenced"

	for k, v := range ngsi.serverList {
		value := v.Context
		if context == value {
			return &LibError{funcName, 1, fmt.Sprintf("%s is referenced in %s", context, k), nil}
		}
	}
	return nil
}

func isIdmType(name string) bool {
	return Contains(idmTypes, strings.ToLower(name))
}

func (info *Server) safeString() (bool, error) {
	const funcName = "safeString"

	value := info.SafeString
	b, err := gNGSI.BoolFlag(value)
	if err != nil {
		return false, &LibError{funcName, 1, err.Error(), err}
	}
	return b, nil
}

func (info *Server) xAuthToken() (bool, error) {
	const funcName = "xAuthToken"

	value := info.XAuthToken
	b, err := gNGSI.BoolFlag(value)
	if err != nil {
		return false, &LibError{funcName, 1, err.Error(), err}
	}
	return b, nil
}
