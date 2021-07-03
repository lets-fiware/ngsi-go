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
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type idmKeyrock struct {
}

func (i *idmKeyrock) requestToken(ngsi *NGSI, client *Client, broker *Server, refresh string) (*TokenInfo, error) {
	const funcName = "requestTokenKeyrock"

	headers := make(map[string]string)
	u, _ := url.Parse(client.idmURL())
	idm := Client{URL: u, Headers: headers, HTTP: ngsi.HTTP}

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}
	idm.SetHeader(cContentType, cAppXWwwFormUrlencoded)
	auth := fmt.Sprintf("%s:%s", broker.ClientID, broker.ClientSecret)
	idm.SetHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(auth))))
	data := fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password)

	res, body, err := idm.HTTPPost(data)
	if err != nil {
		return nil, &LibError{funcName, 2, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, &LibError{funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	var tokenInfo TokenInfo
	utime := ngsi.TimeLib.NowUnix()

	var token OauthToken
	err = JSONUnmarshal(body, &token)
	if err != nil {
		return nil, &LibError{funcName, 4, err.Error(), err}
	}

	tokenInfo.Type = CKeyrock
	tokenInfo.Token = token.AccessToken
	tokenInfo.RefreshToken = token.RefreshToken
	tokenInfo.Expires = time.Unix(utime+token.ExpiresIn, 0)
	tokenInfo.Oauth = &token

	return &tokenInfo, nil
}

func (i *idmKeyrock) getAuthHeader(token string) (string, string) {
	return "Authorization", "Bearer " + token
}
