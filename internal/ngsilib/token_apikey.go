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
	"os"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

type idmApikey struct {
}

func (i *idmApikey) requestToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) (*TokenInfo, error) {
	return nil, nil
}

func (i *idmApikey) revokeToken(ngsi *NGSI, client *Client, tokenInfo *TokenInfo) error {
	return nil
}

func (i *idmApikey) getAuthHeader(token string) (string, string) {
	return "", ""
}

func (i *idmApikey) getTokenInfo(tokenInfo *TokenInfo) ([]byte, error) {
	const funcName = "getTokenInfoApikey"

	return nil, ngsierr.New(funcName, 1, "no information available", nil)
}

func (i *idmApikey) checkIdmParams(idmParams *IdmParams) error {
	const funcName = "checkIdmParamsApikey"

	if idmParams.IdmHost == "" &&
		idmParams.Username == "" &&
		idmParams.Password == "" &&
		idmParams.ClientID == "" &&
		idmParams.ClientSecret == "" &&
		idmParams.HeaderName != "" &&
		((idmParams.HeaderValue == "") != (idmParams.HeaderEnvValue == "")) {
		return nil
	}
	return ngsierr.New(funcName, 1, "headerName and either headerValue or headerEnvValue", nil)
}

func GetApikeyHeader(client *Client) (string, string) {
	n := client.Server.HeaderName
	v := client.Server.HeaderValue
	if v != "" {
		return n, v
	}
	v = client.Server.HeaderEnvValue
	if v != "" {
		return n, os.Getenv(v)
	}
	return n, ""
}
