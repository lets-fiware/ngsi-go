/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"net/http"
	"net/url"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// KeyrockToken is ...
type KeyrockIDMToken struct {
	Token struct {
		Methods   []string `json:"methods"`
		ExpiresAt string   `json:"expires_at"`
	} `json:"token"`
	IdmAuthorizationConfig struct {
		Level      string `json:"level"`
		Authzforce bool   `json:"authzforce"`
	} `json:"idm_authorization_config"`
}

type idmKeyrockIDM struct {
}

func (i *idmKeyrockIDM) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenKeyrockIDM"

	headers := make(map[string]string)
	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}

	idm.SetHeader(cContentType, cAppJSON)
	data := fmt.Sprintf("{\"name\": \"%s\", \"password\": \"%s\"}", username, password)

	res, body, err := idm.HTTPPost(data)
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	var token KeyrockIDMToken
	err = JSONUnmarshal(body, &token)
	if err != nil {
		return nil, ngsierr.New(funcName, 4, err.Error(), err)
	}
	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, token.Token.ExpiresAt)

	tokenInfo.Type = CKeyrockIDM
	tokenInfo.Token = res.Header.Get("X-Subject-Token")
	tokenInfo.Expires = t
	tokenInfo.RefreshToken = ""
	tokenInfo.KeyrockIDM = &token

	return tokenInfo, nil
}

func (i *idmKeyrockIDM) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	const funcName = "revokeTokenKeyrockIDM"

	headers := make(map[string]string)

	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	idm.SetHeader("X-Auth-token", tokenInfo.Token)
	idm.SetHeader("X-Subject-token", tokenInfo.Token)

	res, body, err := idm.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}

func (i *idmKeyrockIDM) getAuthHeader(token string) (string, string) {
	return "X-Auth-Token", token
}

func (i *idmKeyrockIDM) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	// unused: ngsicmd/token.go
	return nil, nil
}

func (i *idmKeyrockIDM) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsKeyIDM"

	if idmParams.IdmHost != "" &&
		idmParams.Username != "" &&
		idmParams.Password != "" &&
		idmParams.ClientID == "" &&
		idmParams.ClientSecret == "" &&
		idmParams.HeaderName == "" &&
		idmParams.HeaderValue == "" &&
		idmParams.HeaderEnvValue == "" {
		return nil
	}
	return ngsierr.New(funcName, 1, "username and password are needed", nil)
}
