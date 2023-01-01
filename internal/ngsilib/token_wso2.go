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

// WSO2Token is ...
type WSO2Token struct {
	AccessToken  string      `json:"access_token"`
	ExpiresIn    int64       `json:"expires_in"`
	RefreshToken string      `json:"refresh_token"`
	Scope        interface{} `json:"scope"`
	TokenType    string      `json:"token_type"`
}

type idmWSO2 struct {
}

func (i *idmWSO2) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenWSO2"

	var res *http.Response
	var body []byte
	var payloads []string
	broker := client.Server

	if tokenInfo.RefreshToken != "" {
		payloads = append(payloads, fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", tokenInfo.RefreshToken))
	}

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}
	payloads = append(payloads,
		fmt.Sprintf("grant_type=password&username=%s&password=%s&client_id=%s&client_secret=%s",
			username, password, broker.ClientID, broker.ClientSecret))

	for _, payload := range payloads {

		headers := make(map[string]string)
		u, _ := url.Parse(client.idmURL())
		idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}
		broker := client.Server

		idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
		auth := fmt.Sprintf("%s:%s", broker.ClientID, broker.ClientSecret)
		idm.SetHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth))))

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

			var token WSO2Token
			err = JSONUnmarshal(body, &token)
			if err != nil {
				return nil, ngsierr.New(funcName, 3, err.Error(), err)
			}

			tokenInfo.Type = CWSO2
			tokenInfo.Token = token.AccessToken
			tokenInfo.Expires = time.Unix(utime+token.ExpiresIn, 0)
			tokenInfo.RefreshToken = token.RefreshToken
			tokenInfo.WSO2 = &token

			return tokenInfo, nil
		}
	}

	return nil, ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
}

func (i *idmWSO2) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	const funcName = "revokeTokenWSO2"

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

func (i *idmWSO2) getAuthHeader(token string) (string, string) {
	return "Authorization", "Bearer " + token
}

func (i *idmWSO2) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoWSO2"

	b, err := JSONMarshal(tokenInfo.WSO2)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}
	return b, nil
}

func (i *idmWSO2) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsWSO2"

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
