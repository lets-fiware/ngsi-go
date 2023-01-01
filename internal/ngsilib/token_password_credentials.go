/*
MIT License

Copyright (c) 2020-2023 Kazuhito Suda

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
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// OauthToken is ...
type OAuthToken struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int64    `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

type idmPasswordCredentials struct {
}

func (i *idmPasswordCredentials) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenPasswordCredentials"

	broker := client.Server

	headers := make(map[string]string)
	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}

	idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
	data := fmt.Sprintf("grant_type=password&username=%s&password=%s&client_id=%s&client_secret=%s",
		username, password, broker.ClientID, broker.ClientSecret)

	res, body, err := idm.HTTPPost(data)
	if err != nil {
		return nil, ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	utime := ngsi.TimeLib.NowUnix()

	var token OAuthToken
	err = JSONUnmarshal(body, &token)
	if err != nil {
		return nil, ngsierr.New(funcName, 4, err.Error(), err)
	}

	tokenInfo.Type = CPasswordCredentials
	tokenInfo.Token = token.AccessToken
	tokenInfo.Expires = time.Unix(utime+token.ExpiresIn, 0)
	tokenInfo.RefreshToken = token.RefreshToken
	tokenInfo.OAuth = &token

	return tokenInfo, nil
}

func (i *idmPasswordCredentials) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	const funcName = "revokeTokenPasswordCredentials"

	headers := make(map[string]string)

	s := strings.Replace(client.idmURL(), "/token", "/revoke", 1)
	if !strings.HasSuffix(s, "/revoke") {
		s += "/revoke"
	}
	u, _ := url.Parse(s)
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}
	broker := client.Server

	idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
	auth := fmt.Sprintf("%s:%s", broker.ClientID, broker.ClientSecret)
	idm.SetHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth))))
	payload := fmt.Sprintf("token=%s&token_type_hint=refresh_token", tokenInfo.RefreshToken)

	res, body, err := idm.HTTPPost(payload)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}

func (i *idmPasswordCredentials) getAuthHeader(token string) (string, string) {
	return "Authorization", "Bearer " + token
}

func (i *idmPasswordCredentials) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoPasswordCredentials"

	b, err := JSONMarshal(tokenInfo.OAuth)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}

	return b, nil
}

func (i *idmPasswordCredentials) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsPasswordCredentials"

	if idmParams.IdmHost != "" &&
		idmParams.Username != "" &&
		idmParams.Password != "" &&
		idmParams.ClientID != "" &&
		idmParams.ClientSecret != "" &&
		idmParams.HeaderName == "" &&
		idmParams.HeaderValue == "" &&
		idmParams.HeaderEnvValue == "" {
		return nil
	}
	return ngsierr.New(funcName, 1, "idmHost, username, password, clientID and clientSecret are needed", nil)
}
