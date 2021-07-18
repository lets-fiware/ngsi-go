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
	"time"
)

type idmBasic struct {
}

func (i *idmBasic) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	const funcName = "requestTokenBasic"

	username, password, err := getUserNamePassword(client)
	if err != nil {
		return nil, &LibError{funcName, 1, err.Error(), err}
	}

	token := base64.URLEncoding.EncodeToString([]byte(username + ":" + password))
	utime := ngsi.TimeLib.NowUnix()

	tokenInfo.Type = CBasic
	tokenInfo.Token = token
	tokenInfo.RefreshToken = ""
	tokenInfo.Expires = time.Unix(utime+3600, 0)

	return tokenInfo, nil
}

func (i *idmBasic) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	return nil
}

func (i *idmBasic) getAuthHeader(token string) (string, string) {
	return "Authorization", "Basic " + token
}

func (i *idmBasic) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoBasic"

	return nil, &LibError{funcName, 1, "no information available", nil}
}

func (i *idmBasic) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsBasic"

	if idmParams.IdmHost == "" &&
		idmParams.Username != "" &&
		idmParams.Password != "" &&
		idmParams.ClientID == "" &&
		idmParams.ClientSecret == "" {
		return nil
	}
	return &LibError{funcName, 1, "username and password are needed", nil}
}
