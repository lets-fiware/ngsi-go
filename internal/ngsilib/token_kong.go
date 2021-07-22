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
	"path"
	"strings"
	"time"
)

// KongToken is ...
type KongToken struct {
	ExpiresIn   int64  `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type idmKong struct {
}

const (
	cKongService = 0
	cKongIdm     = 1
)

func (i *idmKong) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenKong"

	var err error
	var res *http.Response
	var body []byte
	var payloads []string
	broker := client.Server

	payloads = append(payloads, fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", broker.ClientID, broker.ClientSecret))

	for _, payload := range payloads {

		headers := make(map[string]string)
		u, _ := url.Parse(getKongHost(client.idmURL(), cKongService))
		idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

		idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)

		res, body, err = idm.HTTPPost(payload)
		if err != nil {
			return nil, &LibError{funcName, 1, err.Error(), err}
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

			var token KongToken
			err = JSONUnmarshal(body, &token)
			if err != nil {
				return nil, &LibError{funcName, 2, err.Error(), err}
			}

			tokenInfo.Type = CKong
			tokenInfo.Token = token.AccessToken
			tokenInfo.Expires = time.Unix(utime+token.ExpiresIn, 0)
			tokenInfo.RefreshToken = ""
			tokenInfo.Kong = &token

			return tokenInfo, nil
		}
	}

	return nil, &LibError{funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
}

func (i *idmKong) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	const funcName = "revokeTokenKong"

	headers := make(map[string]string)

	u, _ := url.Parse(getKongHost(client.idmURL(), cKongIdm))
	u.Path = path.Join(u.Path, "/oauth2_tokens/", tokenInfo.Token)
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	res, body, err := idm.HTTPDelete(nil)
	if err != nil {
		return &LibError{funcName, 1, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &LibError{funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}

func (i *idmKong) getAuthHeader(token string) (string, string) {
	return "Authorization", "Bearer " + token
}

func (i *idmKong) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoKong"

	b, err := JSONMarshal(tokenInfo.Kong)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}
	return b, nil
}

func (i *idmKong) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsKeyrock"

	if idmParams.IdmHost != "" &&
		idmParams.Username == "" &&
		idmParams.Password == "" &&
		idmParams.ClientID != "" &&
		idmParams.ClientSecret != "" &&
		idmParams.HeaderName == "" &&
		idmParams.HeaderValue == "" &&
		idmParams.HeaderEnvValue == "" {
		return nil
	}
	return &LibError{funcName, 1, "idmHost, clientID and clientSecret are needed", nil}
}

func getKongHost(idm string, i int) string {
	hosts := strings.Split(idm, ",")
	if len(hosts) <= i {
		return "http://kong-error/"
	}
	return hosts[i]
}
