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
	"path/filepath"
	"strings"
	"time"
)

// OauthToken is ...
type OauthToken struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int64    `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

type TokenInfo struct {
	Type         string         `json:"type"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
	Expires      time.Time      `json:"expires"`
	Oauth        *OauthToken    `json:"Oauth,omitempty"`
	Keyrock      *KeyrockToken  `json:"keyrock,omitempty"`
	Keystone     *KeyStoneToken `json:"keystone,omitempty"`
}

type tokenInfoList map[string]TokenInfo

type tokens struct {
	Version string        `json:"version"`
	Tokens  tokenInfoList `json:"tokens"`
}

//
// TokenPlugin interface
//
type TokenPlugin interface {
	requestToken(ngsi *NGSI, client *Client, broker *Server, refresh string) (*TokenInfo, error)
	getAuthHeader(string) (string, string)
}

const (
	cContentType           = "Content-Type"
	cAppXWwwFormUrlencoded = "application/x-www-form-urlencoded"
	cAppJSON               = "application/json"
)

const cacheFileName = "ngsi-go-token-cache.json"

// InitTokenMgr is ..
func (ngsi *NGSI) InitTokenMgr(file *string) error {
	const funcName = "InitTokenMgr"

	ngsi.Logging(LogDebug, funcName+"\n")

	cacheFile := ngsi.CacheFile

	if file == nil {
		home, err := getConfigDir(cacheFile)
		if err != nil {
			return &LibError{funcName, 1, err.Error(), err}
		}

		s := filepath.Join(home, cacheFileName)
		ngsi.CacheFile.SetFileName(&s)
	} else {
		if *file == "" {
			ngsi.CacheFile.SetFileName(file)
		} else {
			s, err := cacheFile.FilePathAbs(*file)
			if err != nil {
				return &LibError{funcName, 2, err.Error() + " " + s, err}
			}
			cacheFile.SetFileName(&s)
		}
	}

	if err := initTokenList(cacheFile); err != nil {
		return &LibError{funcName, 3, err.Error() + " " + *cacheFile.FileName(), err}
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
			return &LibError{funcName, 1, err.Error(), err}
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
	return nil, &LibError{funcName, 1, "not found", nil}
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
	token, err := requestToken(ngsi, client, info.RefreshToken)
	if err != nil {
		return "", &LibError{funcName, 1, err.Error(), err}
	}
	return token, nil
}

// GetAuthHeader is ...
func (ngsi *NGSI) GetAuthHeader(client *Client) (string, string, error) {
	const funcName = "GetAuthHeader"

	token, err := ngsi.GetToken(client)
	if err != nil {
		return "", "", &LibError{funcName, 1, err.Error(), err}
	}

	broker := client.Server
	idmType := strings.ToLower(broker.IdmType)

	idm, ok := tokenPlugins[idmType]
	if !ok {
		return "", "", &LibError{funcName, 2, "unknown idm type: " + idmType, nil}
	}

	key, value := idm.getAuthHeader(token)

	return key, value, nil
}

var tokenPlugins = map[string]TokenPlugin{
	CKeyrock:              &idmKeyrock{},
	CPasswordCredentials:  &idmPasswordCredentials{},
	CKeyrocktokenprovider: &idmKeyrockTokenProvider{},
	CTokenproxy:           &idmTokenProxy{},
	CKeyrockIDM:           &idmKeyrockIDM{},
	CThinkingCities:       &idmThinkingCities{},
	CBasic:                &idmBasic{},
}

func requestToken(ngsi *NGSI, client *Client, refresh string) (string, error) {
	const funcName = "requestToken"

	ngsi.Logging(LogInfo, funcName+"\n")

	broker := client.Server
	idmType := strings.ToLower(broker.IdmType)

	idm, ok := tokenPlugins[idmType]
	if !ok {
		return "", &LibError{funcName, 1, "unknown idm type: " + idmType, nil}
	}

	var tokenInfo *TokenInfo
	tokenInfo, err := idm.requestToken(ngsi, client, broker, refresh)
	if err != nil {
		return "", &LibError{funcName, 2, err.Error(), err}
	}

	client.storeToken(tokenInfo.Token)

	newTokenList := make(tokenInfoList)
	hash := getHash(client)
	newTokenList[hash] = *tokenInfo

	utime := ngsi.TimeLib.NowUnix()

	for k, v := range ngsi.tokenList {
		if v.Expires.Unix() > utime+gNGSI.Margin {
			newTokenList[k] = v
		}
	}

	ngsi.tokenList = newTokenList

	tokens := make(map[string]interface{})
	tokens["tokens"] = ngsi.tokenList
	tokens["version"] = "1"

	err = saveToken(*ngsi.CacheFile.FileName(), tokens)
	if err != nil {
		return "", &LibError{funcName, 3, err.Error(), err}
	}

	return tokenInfo.Token, nil
}

func saveToken(file string, tokens map[string]interface{}) error {
	const funcName = "saveToken"

	gNGSI.Logging(LogInfo, funcName+"\n")

	if file == "" {
		return nil
	}

	cacheFile := gNGSI.CacheFile

	err := cacheFile.OpenFile(oWRONLY|oCREATE, 0600)
	if err != nil {
		return &LibError{funcName, 1, err.Error() + " " + file, err}
	}
	defer func() { _ = cacheFile.Close() }()

	if err := cacheFile.Truncate(0); err != nil {
		return &LibError{funcName, 2, err.Error(), err}
	}

	err = cacheFile.Encode(tokens)
	if err != nil {
		return &LibError{funcName, 3, err.Error(), err}
	}

	return nil
}

func getHash(client *Client) string {
	s := client.Server.ServerHost + client.Server.Username
	if client.Server.IdmType == CThinkingCities {
		s = s + client.Server.Tenant + client.Server.Scope
	}
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func getUserName(client *Client) (string, error) {
	const funcName = "getUserName"

	s := client.Server.Username
	if s == "" {
		return "", &LibError{funcName, 1, "username is required", nil}
	}
	return s, nil
}

func getPassword(client *Client) (string, error) {
	const funcName = "getPassword"

	s := client.Server.Password
	if s == "" {
		return "", &LibError{funcName, 1, "password is required", nil}
	}
	return s, nil
}

func getUserNamePassword(client *Client) (string, string, error) {
	const funcName = "getUserNamePassword"

	username, err := getUserName(client)
	if err != nil {
		return "", "", &LibError{funcName, 1, err.Error(), err}
	}
	password, err := getPassword(client)
	if err != nil {
		return "", "", &LibError{funcName, 2, err.Error(), err}
	}

	return username, password, nil
}
