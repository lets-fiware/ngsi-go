/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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
	"path"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

type KeycloakToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	IDToken          string `json:"id_token"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type idmKeycloak struct {
}

const (
	keyCloakToken  = "/protocol/openid-connect/token"
	keyCLoakRevoke = "/protocol/openid-connect/revoke"
)

func (i *idmKeycloak) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenKeycloak"

	var res *http.Response
	var body []byte
	var payloads []string
	broker := client.Server

	if tokenInfo.RefreshToken != "" {
		payloads = append(payloads, fmt.Sprintf("grant_type=refresh_token&refresh_token=%s&client_id=%s&client_secret=%s",
			tokenInfo.RefreshToken, broker.ClientID, broker.ClientSecret))
	}

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}

	openid := "openid"
	payloads = append(payloads, fmt.Sprintf("grant_type=password&client_id=%s&client_secret=%s&username=%s&password=%s&scope=%s",
		broker.ClientID, broker.ClientSecret, username, password, openid))

	for _, payload := range payloads {

		headers := make(map[string]string)
		u, _ := url.Parse(client.idmURL())
		u.Path = path.Join(u.Path, keyCloakToken)
		idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

		idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)

		res, body, err = idm.HTTPPost(payload)
		if err != nil {
			return nil, ngsierr.New(funcName, 2, err.Error(), err)
		}

		switch res.StatusCode {
		default:
			gNGSI.Logging(LogInfo, fmt.Sprintf("%s %d\n", funcName, res.StatusCode))
			continue
		case http.StatusUnauthorized:
			gNGSI.Logging(LogInfo, funcName+" Unauthorized\n")
			continue
		case http.StatusOK:
			utime := ngsi.TimeLib.NowUnix()

			var oToken KeycloakToken
			err = JSONUnmarshal(body, &oToken)
			if err != nil {
				return nil, ngsierr.New(funcName, 3, err.Error(), err)
			}

			tokenInfo := &TokenInfo{
				Type:         CKeycloak,
				Token:        oToken.AccessToken,
				RefreshToken: oToken.RefreshToken,
				Expires:      time.Unix(utime+int64(oToken.ExpiresIn), 0),
				Keycloak:     &oToken,
			}
			return tokenInfo, nil
		}
	}

	return nil, ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
}

func (i *idmKeycloak) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	const funcName = "revokeTokenKeycloak"

	headers := make(map[string]string)

	u, _ := url.Parse(client.idmURL())
	u.Path = path.Join(u.Path, keyCLoakRevoke)
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	broker := client.Server
	idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
	payload := fmt.Sprintf("token=%s&client_id=%s&client_secret=%s", tokenInfo.RefreshToken, broker.ClientID, broker.ClientSecret)

	res, body, err := idm.HTTPPost(payload)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}

func (i *idmKeycloak) getAuthHeader(token string) (string, string) {
	return "Authorization", "Bearer " + token
}

func (i *idmKeycloak) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoKeycloak"

	b, err := JSONMarshal(tokenInfo.Keycloak)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}
	return b, nil
}

func (i *idmKeycloak) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsKeycloak"

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
