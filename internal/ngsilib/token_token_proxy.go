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

// Token proxy is a proxy for Keyrock
// TokenProxy is ...
type TokenProxy struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int64    `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

type idmTokenProxy struct {
}

func (i *idmTokenProxy) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenTokenProxy"

	var res *http.Response
	var body []byte
	var payloads []string

	if tokenInfo.RefreshToken != "" {
		payloads = append(payloads, fmt.Sprintf("{\"refresh\": \"%s\"}", tokenInfo.RefreshToken))
	}
	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}
	payloads = append(payloads, fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\"}", username, password))

	for _, payload := range payloads {
		headers := make(map[string]string)
		u, _ := url.Parse(client.idmURL())
		idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

		idm.SetHeader(cContentType, cAppJSON)

		res, body, err = idm.HTTPPost(payload)
		if err != nil {
			return nil, &LibError{funcName, 2, err.Error(), err}
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

			var tokenProxy TokenProxy
			err = JSONUnmarshal(body, &tokenProxy)
			if err != nil {
				return nil, &LibError{funcName, 3, err.Error(), err}
			}

			tokenInfo := &TokenInfo{
				Type:         CTokenproxy,
				Token:        tokenProxy.AccessToken,
				Expires:      time.Unix(utime+tokenProxy.ExpiresIn, 0),
				RefreshToken: tokenProxy.RefreshToken,
				TokenProxy:   &tokenProxy,
			}
			return tokenInfo, nil
		}
	}

	return nil, &LibError{funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
}

func (i *idmTokenProxy) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	const funcName = "revokeTokenTokenProxy"

	headers := make(map[string]string)

	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	idm.SetHeader(cContentType, cAppJSON)
	payload := fmt.Sprintf(`{"token":"%s","token_type_hint":"refresh_token"}`, tokenInfo.RefreshToken)

	res, body, err := idm.HTTPPost(payload)
	if err != nil {
		return &LibError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &LibError{funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}

func (i *idmTokenProxy) getAuthHeader(token string) (string, string) {
	return "Authorization", "Bearer " + token
}

func (i *idmTokenProxy) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoTokenProxy"

	b, err := JSONMarshal(tokenInfo.TokenProxy)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}
	return b, nil
}

func (i *idmTokenProxy) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsTokenProxy"

	if idmParams.IdmHost != "" &&
		idmParams.Username != "" &&
		idmParams.Password != "" &&
		idmParams.ClientID == "" &&
		idmParams.ClientSecret == "" {
		return nil
	}
	return &LibError{funcName, 1, "idmHost, username and password are needed", nil}
}
