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
	"net/http"
	"net/url"
	"time"
)

type idmThinkingCities struct {
}

func (i *idmThinkingCities) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenThinkingCities"

	headers := make(map[string]string)
	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}

	idm.SetHeader(cContentType, cAppJSON)
	data := getKeyStoneTokenRequest(username, password, client.Server.Tenant, client.Server.Scope)

	res, body, err := idm.HTTPPost(data)
	if err != nil {
		return nil, &LibError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, &LibError{funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	var token KeyStoneToken
	err = JSONUnmarshal(body, &token)
	if err != nil {
		return nil, &LibError{funcName, 4, err.Error(), err}
	}
	layout := "2006-01-02T15:04:05.000000Z"
	t, _ := time.Parse(layout, token.Token.ExpiresAt)

	tokenInfo.Type = CThinkingCities
	tokenInfo.Token = res.Header.Get("X-Subject-Token")
	tokenInfo.RefreshToken = ""
	tokenInfo.Expires = t
	tokenInfo.Keystone = &token

	return tokenInfo, nil
}

func (i *idmThinkingCities) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	return nil
}

func (i *idmThinkingCities) getAuthHeader(token string) (string, string) {
	return "X-Auth-Token", token
}

func (i *idmThinkingCities) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoThinkingCities"

	b, err := JSONMarshal(tokenInfo.Keystone)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}
	return b, nil
}

func (i *idmThinkingCities) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsThinkingCities"

	if idmParams.IdmHost != "" &&
		idmParams.Username != "" &&
		idmParams.Password != "" &&
		idmParams.ClientID == "" &&
		idmParams.ClientSecret == "" {
		return nil
	}
	return &LibError{funcName, 1, "idmHost, username and password are needed", nil}
}
