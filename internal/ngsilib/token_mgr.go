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
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// TokenInfo is ...
type TokenInfo struct {
	Type          string           `json:"type"`
	Token         string           `json:"token"`
	RefreshToken  string           `json:"refresh_token"`
	Expires       time.Time        `json:"expires"`
	OAuth         *OAuthToken      `json:"oauth,omitempty"`
	KeyrockIDM    *KeyrockIDMToken `json:"keyrock_idm,omitempty"`
	Keyrock       *KeyrockToken    `json:"keyrock,omitempty"`
	TokenProxy    *TokenProxy      `json:"tokenproxy,omitempty"`
	TokenProvider *TokenProvider   `json:"tokenprovider,omitempty"`
	Keystone      *KeyStoneToken   `json:"keystone,omitempty"`
	Keycloak      *KeycloakToken   `json:"keycloak,omitempty"`
	WSO2          *WSO2Token       `json:"wso2,omitempty"`
	Kong          *KongToken       `json:"kong,omitempty"`
}

type tokenInfoList map[string]TokenInfo

type tokens struct {
	Version string        `json:"version"`
	Tokens  tokenInfoList `json:"tokens"`
}

// IdmParam is ...
type IdmParams struct {
	IdmType        string
	IdmHost        string
	Username       string
	Password       string
	ClientID       string
	ClientSecret   string
	HeaderName     string
	HeaderValue    string
	HeaderEnvValue string
	TokenScope     string
}

//
// TokenPlugin interface
//
// TokenPlugin is ...
type TokenPlugin interface {
	requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error)
	revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error
	getAuthHeader(string) (string, string)
	getTokenInfo(tokenInfo *TokenInfo) ([]byte, error)
	checkIdmParams(idmParams *IdmParams) error
}

const (
	cContentType           = "Content-Type"
	cAppXWwwFormUrlencoded = "application/x-www-form-urlencoded"
	cAppJSON               = "application/json"
)

const (
	CPasswordCredentials  = "password"
	CKeyrock              = "keyrock"
	CKeyrocktokenprovider = "keyrocktokenprovider"
	CTokenproxy           = "tokenproxy"
	CKeyrockIDM           = "idm"
	CThinkingCities       = "thinkingcities"
	CBasic                = "basic"
	CKeycloak             = "keycloak"
	CWSO2                 = "wso2"
	CKong                 = "kong"
	CApikey               = "apikey"
)

var idmTypes = []string{
	CPasswordCredentials, CKeyrock, CKeyrocktokenprovider, CTokenproxy, CKeyrockIDM,
	CThinkingCities, CBasic, CKeycloak, CWSO2, CKong, CApikey,
}

var tokenPlugins = map[string]TokenPlugin{
	CKeyrock:              &idmKeyrock{},
	CPasswordCredentials:  &idmPasswordCredentials{},
	CKeyrocktokenprovider: &idmKeyrockTokenProvider{},
	CTokenproxy:           &idmTokenProxy{},
	CKeyrockIDM:           &idmKeyrockIDM{},
	CThinkingCities:       &idmThinkingCities{},
	CBasic:                &idmBasic{},
	CKeycloak:             &idmKeycloak{},
	CWSO2:                 &idmWSO2{},
	CKong:                 &idmKong{},
	CApikey:               &idmApikey{},
}

const cacheFileName = "ngsi-go-token-cache.json"

// InitTokenMgr is ..
func (ngsi *NGSI) InitTokenMgr(file *string) error {
	const funcName = "InitTokenMgr"

	ngsi.Logging(LogDebug, funcName+"\n")

	cacheFile := ngsi.CacheFile

	if file == nil {
		home, err := getConfigDir(cacheFile)
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}

		s := filepath.Join(home, cacheFileName)
		ngsi.CacheFile.SetFileName(&s)
	} else {
		if *file == "" {
			ngsi.CacheFile.SetFileName(file)
		} else {
			s, err := cacheFile.FilePathAbs(*file)
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error()+" "+s, err)
			}
			cacheFile.SetFileName(&s)
		}
	}

	if err := initTokenList(cacheFile); err != nil {
		return ngsierr.New(funcName, 3, err.Error()+" "+*cacheFile.FileName(), err)
	}

	return nil
}

func initTokenList(io IoLib) (err error) {
	const funcName = "initTokenList"

	if *io.FileName() == "" {
		return nil
	}

	gNGSI.tokenList = make(tokenInfoList)

	if existsFile(io, *io.FileName()) {
		err = io.Open()
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		defer func() { _ = io.Close() }()

		tokens := tokens{}
		err = io.Decode(&tokens)
		if err == nil {
			if tokens.Version == "1" {
				gNGSI.tokenList = tokens.Tokens
			}
		}
	}

	return nil
}

// TokenList is ...
func (ngsi *NGSI) TokenList() string {
	list := ""

	for key := range ngsi.tokenList {
		list += key + " "
	}
	if len(list) != 0 {
		list = list[:len(list)-1]
	}
	return list
}

// TokenInfo is ...
func (ngsi *NGSI) TokenInfo(client *Client) (*TokenInfo, error) {
	const funcName = "TokenInfo"

	hash := getHash(client)
	if v, ok := ngsi.tokenList[hash]; ok {
		return &v, nil
	}
	return nil, ngsierr.New(funcName, 1, "not found", nil)
}

// GetToken is ...
func (ngsi *NGSI) GetToken(client *Client) (string, error) {
	const funcName = "GetToken"

	hash := getHash(client)
	info, ok := ngsi.tokenList[hash]
	if ok {
		accessToken := info.Token
		expires := info.Expires
		utime := ngsi.TimeLib.NowUnix()

		if accessToken != "" && expires.Unix() > utime+gNGSI.Margin {
			gNGSI.Logging(LogInfo, "Cached token is used\n")
			gNGSI.Logging(LogDebug, accessToken+"\n")
			return accessToken, nil
		}
	}
	token, err := requestToken(ngsi, client, &info)
	if err != nil {
		return "", ngsierr.New(funcName, 1, err.Error(), err)
	}
	return token, nil
}

// GetAuthHeader is ...
func (ngsi *NGSI) GetAuthHeader(client *Client) (string, string, error) {
	const funcName = "GetAuthHeader"

	token, err := ngsi.GetToken(client)
	if err != nil {
		return "", "", ngsierr.New(funcName, 1, err.Error(), err)
	}

	idmType := strings.ToLower(client.Server.IdmType)

	idm, ok := tokenPlugins[idmType]
	if !ok {
		return "", "", ngsierr.New(funcName, 2, "unknown idm type: "+idmType, nil)
	}

	key, value := idm.getAuthHeader(token)

	return key, value, nil
}

// RevokeToken is ...
func (ngsi *NGSI) RevokeToken(client *Client) error {
	const funcName = "RevokeToken"

	hash := getHash(client)

	if tokenInfo, ok := ngsi.tokenList[hash]; ok {
		idmType := strings.ToLower(client.Server.IdmType)
		idm, ok := tokenPlugins[idmType]
		if !ok {
			return ngsierr.New(funcName, 1, "unknown idm type: "+idmType, nil)
		}

		err := idm.revokeToken(ngsi, client, &tokenInfo)
		if err != nil {
			fmt.Fprint(ngsi.Stderr, ngsierr.SprintMsg(funcName, 2, err.Error()))
		}

		delete(ngsi.tokenList, hash)

		err = saveToken(*ngsi.CacheFile.FileName(), &ngsi.tokenList)
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}
	}

	return nil
}

func (ngsi *NGSI) GetTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "GetTokenInfo"

	idmType := tokenInfo.Type
	idm, ok := tokenPlugins[idmType]
	if !ok {
		return nil, ngsierr.New(funcName, 1, "unknown idm type: "+idmType, nil)
	}

	b, err := idm.getTokenInfo(tokenInfo)
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}

	return b, nil
}

func requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (string, error) {
	const funcName = "requestToken"

	ngsi.Logging(LogInfo, funcName+"\n")

	idmType := strings.ToLower(client.Server.IdmType)

	idm, ok := tokenPlugins[idmType]
	if !ok {
		return "", ngsierr.New(funcName, 1, "unknown idm type: "+idmType, nil)
	}

	tokenInfo, err := idm.requestToken(ngsi, client, tokenInfo)
	if err != nil {
		return "", ngsierr.New(funcName, 2, err.Error(), err)
	}

	client.storeToken(tokenInfo.Token)

	hash := getHash(client)
	newList := appendToken(ngsi, hash, tokenInfo)
	ngsi.tokenList = *newList

	err = saveToken(*ngsi.CacheFile.FileName(), newList)
	if err != nil {
		return "", ngsierr.New(funcName, 3, err.Error(), err)
	}

	return tokenInfo.Token, nil
}

func appendToken(ngsi *NGSI, hash string, tokenInfo *TokenInfo) *tokenInfoList {
	newTokenList := make(tokenInfoList)
	newTokenList[hash] = *tokenInfo

	utime := ngsi.TimeLib.NowUnix()

	for k, v := range ngsi.tokenList {
		if v.Expires.Unix() > utime+gNGSI.Margin {
			newTokenList[k] = v
		}
	}

	return &newTokenList
}

func saveToken(file string, tokenList *tokenInfoList) error {
	const funcName = "saveToken"

	gNGSI.Logging(LogInfo, funcName+"\n")

	if file == "" {
		return nil
	}

	tokens := &tokens{
		Version: "1",
		Tokens:  *tokenList,
	}

	cacheFile := gNGSI.CacheFile

	err := cacheFile.OpenFile(oWRONLY|oCREATE, 0600)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error()+" "+file, err)
	}
	defer func() { _ = cacheFile.Close() }()

	if err := cacheFile.Truncate(0); err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	err = cacheFile.Encode(tokens)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	return nil
}

func getHash(client *Client) string {
	s := client.Server.ServerHost + client.Server.Username + client.Server.Password
	if client.Server.IdmType == CThinkingCities {
		s = s + client.Server.Tenant + client.Server.Scope
	}
	if client.Server.IdmType == CKeycloak {
		s = s + client.Server.ClientID + client.Server.ClientSecret
	}
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func getUserName(client *Client) (string, error) {
	const funcName = "getUserName"

	s := client.Server.Username
	if s == "" {
		return "", ngsierr.New(funcName, 1, "username is required", nil)
	}
	return s, nil
}

func getPassword(client *Client) (string, error) {
	const funcName = "getPassword"

	s := client.Server.Password
	if s == "" {
		return "", ngsierr.New(funcName, 1, "password is required", nil)
	}
	return s, nil
}

func getUserNamePassword(client *Client) (string, string, error) {
	const funcName = "getUserNamePassword"

	username, err := getUserName(client)
	if err != nil {
		return "", "", ngsierr.New(funcName, 1, err.Error(), err)
	}
	password, err := getPassword(client)
	if err != nil {
		return "", "", ngsierr.New(funcName, 2, err.Error(), err)
	}

	return username, password, nil
}

func checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParams"

	idmType := strings.ToLower(idmParams.IdmType)

	if idmType == "" {
		if !(idmParams.IdmHost == "" &&
			idmParams.Username == "" &&
			idmParams.Password == "" &&
			idmParams.ClientID == "" &&
			idmParams.ClientSecret == "" &&
			idmParams.HeaderName == "" &&
			idmParams.HeaderValue == "" &&
			idmParams.HeaderEnvValue == "" &&
			idmParams.TokenScope == "") {
			return ngsierr.New(funcName, 1, "required idmType not found", nil)
		}
		return nil
	}

	idm, ok := tokenPlugins[idmType]
	if !ok {
		return ngsierr.New(funcName, 2, "unknown idm type: "+idmType, nil)
	}

	err := idm.checkIdmParams(idmParams)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	return nil
}
