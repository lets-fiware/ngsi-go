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

// https://github.com/FIWARE-Ops/KeyrockTokenProvider

type idmKeyrockTokenProvider struct {
}

func (i *idmKeyrockTokenProvider) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenTokenProvider"

	headers := make(map[string]string)
	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}

	idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
	data := fmt.Sprintf("username=%s&password=%s", username, password)

	res, body, err := idm.HTTPPost(data)
	if err != nil {
		return nil, &LibError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, &LibError{funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	utime := ngsi.TimeLib.NowUnix()

	var token OauthToken
	token.AccessToken = string(body)
	token.ExpiresIn = 3600

	tokenInfo.Type = CKeyrocktokenprovider
	tokenInfo.Token = token.AccessToken
	tokenInfo.Expires = time.Unix(utime+token.ExpiresIn, 0)
	tokenInfo.RefreshToken = ""
	tokenInfo.Oauth = &token

	return tokenInfo, nil
}

func (i *idmKeyrockTokenProvider) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	return nil
}

func (i *idmKeyrockTokenProvider) getAuthHeader(token string) (string, string) {
	return "Authorization", "Bearer " + token
}
