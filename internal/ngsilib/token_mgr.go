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
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
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

// KeyrockToken is ...
type KeyrockToken struct {
	Token struct {
		Methods   []string `json:"methods"`
		ExpiresAt string   `json:"expires_at"`
	} `json:"token"`
	IdmAuthorizationConfig struct {
		Level      string `json:"level"`
		Authzforce bool   `json:"authzforce"`
	} `json:"idm_authorization_config"`
}

// TokenInfo is ...
type TokenInfo struct {
	Expires       int64          `json:"expires,omitempty"`
	OauthToken    *OauthToken    `json:"token,omitempty"`
	Keyrock       *KeyrockToken  `json:"keyrock,omitempty"`
	KeyrockToken  *string        `json:"keyrock_token,omitempty"`
	Keystone      *KeyStoneToken `json:"keystone,omitempty"`
	KeystoneToken *string        `json:"keystone_token,omitempty"`
}

type tokenInfoList map[string]TokenInfo
type tokens struct {
	Tokens tokenInfoList `json:"tokens"`
}

const (
	cContentType           = "Content-Type"
	cAppXWwwFormUrlencoded = "application/x-www-form-urlencoded"
	cAppJSON               = "application/json"
)

// var cacheFile string

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

	if existsFile(io, *io.FileName()) {
		err = io.Open()
		if err != nil {
			return &LibError{funcName, 1, err.Error(), err}
		}
		defer func() { _ = io.Close() }()

		tokens := tokens{}
		err = io.Decode(&tokens)
		if err != nil {
			return &LibError{funcName, 3, err.Error(), err}
		}

		gNGSI.tokenList = tokens.Tokens
	} else {
		gNGSI.tokenList = make(tokenInfoList)
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
		accessToken := ""
		expires := info.Expires
		if info.OauthToken != nil {
			accessToken = info.OauthToken.AccessToken
		} else if info.KeyrockToken != nil {
			accessToken = *info.KeyrockToken
		} else if info.KeystoneToken != nil {
			accessToken = *info.KeystoneToken
		} else {
			return "", &LibError{funcName, 1, "token list error", nil}
		}

		utime := ngsi.TimeLib.NowUnix()

		if expires > utime+gNGSI.Margin {
			gNGSI.Logging(LogInfo, "Cached token is used\n")
			gNGSI.Logging(LogDebug, accessToken+"\n")
			return accessToken, nil
		}
	}
	token, err := getToken(ngsi, client)
	if err != nil {
		return "", &LibError{funcName, 2, err.Error(), err}
	}
	return token, nil
}

func getToken(ngsi *NGSI, client *Client) (string, error) {
	const funcName = "getToken"

	ngsi.Logging(LogInfo, funcName+"\n")

	var data string
	headers := make(map[string]string)
	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	username, err := getUserName(client)
	if err != nil {
		return "", &LibError{funcName, 1, err.Error(), err}
	}
	password, err := getPassword(client)
	if err != nil {
		return "", &LibError{funcName, 2, err.Error(), err}
	}

	broker := client.Server
	idmType := strings.ToLower(broker.IdmType)

	switch idmType {
	default:
		return "", &LibError{funcName, 3, "unknown idm type: " + idmType, nil}
	case cKeyrock:
		idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
		auth := fmt.Sprintf("%s:%s", broker.ClientID, broker.ClientSecret)
		idm.SetHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth))))
		data = fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)
	case cPasswordCredentials:
		idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
		data = fmt.Sprintf("grant_type=password&username=%s&password=%s&client_id=%s&client_secret=%s", username, password, broker.ClientID, broker.ClientSecret)
	case cKeyrocktokenprovider:
		idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
		data = fmt.Sprintf("username=%s&password=%s", username, password)
	case cTokenproxy:
		idm.SetHeader(cContentType, cAppJSON)
		data = fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\"}", username, password)
	case cKeyrockIDM: // Keyrock
		idm.SetHeader(cContentType, cAppJSON)
		data = fmt.Sprintf("{\"name\": \"%s\", \"password\": \"%s\"}", username, password)
	case cThinkingCities:
		idm.SetHeader(cContentType, cAppJSON)
		data = getKeyStoneTokenRequest(username, password, client.Server.Tenant, client.Server.Scope)
	}

	res, body, err := idm.HTTPPost(data)
	if err != nil {
		return "", &LibError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return "", &LibError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	accessToken := ""

	var tokenInfo TokenInfo
	utime := ngsi.TimeLib.NowUnix()

	switch idmType {
	default:
		var token OauthToken
		err := JSONUnmarshal(body, &token)
		if err != nil {
			return "", &LibError{funcName, 6, err.Error(), err}
		}
		tokenInfo.Expires = utime + token.ExpiresIn
		tokenInfo.OauthToken = &token
		accessToken = token.AccessToken
	case cKeyrocktokenprovider:
		var token OauthToken
		r := fmt.Sprintf(`{"access_token":"%s", "expires_in":%d}`, string(body), client.getExpiresIn())
		err := JSONUnmarshal([]byte(r), &token)
		if err != nil {
			return "", &LibError{funcName, 7, err.Error(), err}
		}
		tokenInfo.Expires = utime + token.ExpiresIn
		tokenInfo.OauthToken = &token
		accessToken = token.AccessToken
	case cKeyrockIDM:
		var token KeyrockToken
		err := JSONUnmarshal(body, &token)
		if err != nil {
			return "", &LibError{funcName, 8, err.Error(), err}
		}
		layout := "2006-01-02T15:04:05.000Z"
		t, _ := time.Parse(layout, token.Token.ExpiresAt)
		accessToken = res.Header.Get("X-Subject-Token")
		tokenInfo.Expires = t.Unix()
		tokenInfo.Keyrock = &token
		tokenInfo.KeyrockToken = &accessToken
	case cThinkingCities:
		var token KeyStoneToken
		err := JSONUnmarshal(body, &token)
		if err != nil {
			return "", &LibError{funcName, 9, err.Error(), err}
		}
		layout := "2006-01-02T15:04:05.000000Z"
		t, _ := time.Parse(layout, token.Token.ExpiresAt)
		accessToken = res.Header.Get("X-Subject-Token")
		tokenInfo.Expires = t.Unix()
		tokenInfo.Keystone = &token
		tokenInfo.KeystoneToken = &accessToken
	}

	client.storeToken(accessToken)

	newTokenList := make(tokenInfoList)
	hash := getHash(client)
	newTokenList[hash] = tokenInfo

	for k, v := range ngsi.tokenList {
		if v.Expires > utime+gNGSI.Margin {
			newTokenList[k] = v
		}
	}

	ngsi.tokenList = newTokenList

	tokens := make(map[string]interface{})
	tokens["tokens"] = ngsi.tokenList

	err = saveToken(*ngsi.CacheFile.FileName(), tokens)
	if err != nil {
		return "", &LibError{funcName, 10, err.Error(), err}
	}
	return accessToken, nil
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
	if client.Server.IdmType == cThinkingCities {
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
