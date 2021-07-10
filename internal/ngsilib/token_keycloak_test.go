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
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestTokenKeycloak(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA","expires_in":300,"refresh_expires_in":1800,"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJjOTUzMTYwYy0yMGEyLTRiNDQtYmI5NC01ZDgyNjQ2ODZmMzUifQ.eyJleHAiOjE2MjU5NTkzMzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiYzYzZjhhYmEtYWIzOC00OWVkLWExYjYtMjNhZjcyYzQ1MmI4IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXV0aC9yZWFsbXMvZml3YXJlX3NlcnZpY2UiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoibmdzaV9hcGkiLCJzZXNzaW9uX3N0YXRlIjoiMGZkZGZhZDctODA1Yi00MzcxLTkwMDktZmIxOTI3YWZkYjAyIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSJ9.rJ7sMlxKZ-IA0AoqKznZLZnNAK7xsKHoPIRc2owIsw0","token_type":"Bearer","id_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwiYXV0aF90aW1lIjowLCJqdGkiOiI4YzA4YWJiYy05YjI1LTRmMDgtOWMwYi1iMDRmNmEwYWJjNmMiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0L2F1dGgvcmVhbG1zL2Zpd2FyZV9zZXJ2aWNlIiwiYXVkIjoibmdzaV9hcGkiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJJRCIsImF6cCI6Im5nc2lfYXBpIiwic2Vzc2lvbl9zdGF0ZSI6IjBmZGRmYWQ3LTgwNWItNDM3MS05MDA5LWZiMTkyN2FmZGIwMiIsImF0X2hhc2giOiJfalh0eDQxbVlobWRrZGsyc1h5Y0FnIiwiYWNyIjoiMSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.JYjRzJckKsYhVChFw1Y2SGHJgY_PXhjTZRJMDzls7Xj5J7nJ5IiglHY1cSI-A1wzcJHttJYvGwKRK9QziVZfQh2LQoQSChnhHYY1Uq0VMrWfuOgofCQSqOmoXJiu7VwFRWVbg6RNVxlgT_z1SyncXQrCmjtfUBUKsBSSAZxOucVgDmHptT_JcNqKPeeo8-7PelDtx4PZDkf4Qf_77qHFPy0cwXio57UFAGtJAsztxd6nwZ6Q0QQY7XxCQLLALIsJeYJzfB2b58YwTdnpmSHG6oMrWp_Ie-P8cYkHgmNmI_Q1KIYuWYwA6NRqL26rC5CwN7irxn2sgEwShBNeqwhJlA","not-before-policy":0,"session_state":"0fddfad7-805b-4371-9009-fb1927afdb02","scope":"openid email profile"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CKeycloak, actual.Type)
		expected := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA"
		assert.Equal(t, expected, actual.Keycloak.AccessToken)
	}
}

func TestRequestTokenKeycloakRefresh(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA","expires_in":300,"refresh_expires_in":1800,"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJjOTUzMTYwYy0yMGEyLTRiNDQtYmI5NC01ZDgyNjQ2ODZmMzUifQ.eyJleHAiOjE2MjU5NTkzMzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiYzYzZjhhYmEtYWIzOC00OWVkLWExYjYtMjNhZjcyYzQ1MmI4IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXV0aC9yZWFsbXMvZml3YXJlX3NlcnZpY2UiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoibmdzaV9hcGkiLCJzZXNzaW9uX3N0YXRlIjoiMGZkZGZhZDctODA1Yi00MzcxLTkwMDktZmIxOTI3YWZkYjAyIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSJ9.rJ7sMlxKZ-IA0AoqKznZLZnNAK7xsKHoPIRc2owIsw0","token_type":"Bearer","id_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwiYXV0aF90aW1lIjowLCJqdGkiOiI4YzA4YWJiYy05YjI1LTRmMDgtOWMwYi1iMDRmNmEwYWJjNmMiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0L2F1dGgvcmVhbG1zL2Zpd2FyZV9zZXJ2aWNlIiwiYXVkIjoibmdzaV9hcGkiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJJRCIsImF6cCI6Im5nc2lfYXBpIiwic2Vzc2lvbl9zdGF0ZSI6IjBmZGRmYWQ3LTgwNWItNDM3MS05MDA5LWZiMTkyN2FmZGIwMiIsImF0X2hhc2giOiJfalh0eDQxbVlobWRrZGsyc1h5Y0FnIiwiYWNyIjoiMSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.JYjRzJckKsYhVChFw1Y2SGHJgY_PXhjTZRJMDzls7Xj5J7nJ5IiglHY1cSI-A1wzcJHttJYvGwKRK9QziVZfQh2LQoQSChnhHYY1Uq0VMrWfuOgofCQSqOmoXJiu7VwFRWVbg6RNVxlgT_z1SyncXQrCmjtfUBUKsBSSAZxOucVgDmHptT_JcNqKPeeo8-7PelDtx4PZDkf4Qf_77qHFPy0cwXio57UFAGtJAsztxd6nwZ6Q0QQY7XxCQLLALIsJeYJzfB2b58YwTdnpmSHG6oMrWp_Ie-P8cYkHgmNmI_Q1KIYuWYwA6NRqL26rC5CwN7irxn2sgEwShBNeqwhJlA","not-before-policy":0,"session_state":"0fddfad7-805b-4371-9009-fb1927afdb02","scope":"openid email profile"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{RefreshToken: "refresh"}

	actual, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.NoError(t, err) {
		assert.Equal(t, CKeycloak, actual.Type)
		expected := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA"
		assert.Equal(t, expected, actual.Keycloak.AccessToken)
	}
}

func TestRequestTokenKeycloakErrorUser(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA","expires_in":300,"refresh_expires_in":1800,"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJjOTUzMTYwYy0yMGEyLTRiNDQtYmI5NC01ZDgyNjQ2ODZmMzUifQ.eyJleHAiOjE2MjU5NTkzMzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiYzYzZjhhYmEtYWIzOC00OWVkLWExYjYtMjNhZjcyYzQ1MmI4IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXV0aC9yZWFsbXMvZml3YXJlX3NlcnZpY2UiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoibmdzaV9hcGkiLCJzZXNzaW9uX3N0YXRlIjoiMGZkZGZhZDctODA1Yi00MzcxLTkwMDktZmIxOTI3YWZkYjAyIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSJ9.rJ7sMlxKZ-IA0AoqKznZLZnNAK7xsKHoPIRc2owIsw0","token_type":"Bearer","id_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwiYXV0aF90aW1lIjowLCJqdGkiOiI4YzA4YWJiYy05YjI1LTRmMDgtOWMwYi1iMDRmNmEwYWJjNmMiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0L2F1dGgvcmVhbG1zL2Zpd2FyZV9zZXJ2aWNlIiwiYXVkIjoibmdzaV9hcGkiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJJRCIsImF6cCI6Im5nc2lfYXBpIiwic2Vzc2lvbl9zdGF0ZSI6IjBmZGRmYWQ3LTgwNWItNDM3MS05MDA5LWZiMTkyN2FmZGIwMiIsImF0X2hhc2giOiJfalh0eDQxbVlobWRrZGsyc1h5Y0FnIiwiYWNyIjoiMSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.JYjRzJckKsYhVChFw1Y2SGHJgY_PXhjTZRJMDzls7Xj5J7nJ5IiglHY1cSI-A1wzcJHttJYvGwKRK9QziVZfQh2LQoQSChnhHYY1Uq0VMrWfuOgofCQSqOmoXJiu7VwFRWVbg6RNVxlgT_z1SyncXQrCmjtfUBUKsBSSAZxOucVgDmHptT_JcNqKPeeo8-7PelDtx4PZDkf4Qf_77qHFPy0cwXio57UFAGtJAsztxd6nwZ6Q0QQY7XxCQLLALIsJeYJzfB2b58YwTdnpmSHG6oMrWp_Ie-P8cYkHgmNmI_Q1KIYuWYwA6NRqL26rC5CwN7irxn2sgEwShBNeqwhJlA","not-before-policy":0,"session_state":"0fddfad7-805b-4371-9009-fb1927afdb02","scope":"openid email profile"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "password is required", ngsiErr.Message)
	}
}

func TestRequestTokenKeycloakErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA","expires_in":300,"refresh_expires_in":1800,"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJjOTUzMTYwYy0yMGEyLTRiNDQtYmI5NC01ZDgyNjQ2ODZmMzUifQ.eyJleHAiOjE2MjU5NTkzMzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiYzYzZjhhYmEtYWIzOC00OWVkLWExYjYtMjNhZjcyYzQ1MmI4IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXV0aC9yZWFsbXMvZml3YXJlX3NlcnZpY2UiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoibmdzaV9hcGkiLCJzZXNzaW9uX3N0YXRlIjoiMGZkZGZhZDctODA1Yi00MzcxLTkwMDktZmIxOTI3YWZkYjAyIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSJ9.rJ7sMlxKZ-IA0AoqKznZLZnNAK7xsKHoPIRc2owIsw0","token_type":"Bearer","id_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwiYXV0aF90aW1lIjowLCJqdGkiOiI4YzA4YWJiYy05YjI1LTRmMDgtOWMwYi1iMDRmNmEwYWJjNmMiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0L2F1dGgvcmVhbG1zL2Zpd2FyZV9zZXJ2aWNlIiwiYXVkIjoibmdzaV9hcGkiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJJRCIsImF6cCI6Im5nc2lfYXBpIiwic2Vzc2lvbl9zdGF0ZSI6IjBmZGRmYWQ3LTgwNWItNDM3MS05MDA5LWZiMTkyN2FmZGIwMiIsImF0X2hhc2giOiJfalh0eDQxbVlobWRrZGsyc1h5Y0FnIiwiYWNyIjoiMSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.JYjRzJckKsYhVChFw1Y2SGHJgY_PXhjTZRJMDzls7Xj5J7nJ5IiglHY1cSI-A1wzcJHttJYvGwKRK9QziVZfQh2LQoQSChnhHYY1Uq0VMrWfuOgofCQSqOmoXJiu7VwFRWVbg6RNVxlgT_z1SyncXQrCmjtfUBUKsBSSAZxOucVgDmHptT_JcNqKPeeo8-7PelDtx4PZDkf4Qf_77qHFPy0cwXio57UFAGtJAsztxd6nwZ6Q0QQY7XxCQLLALIsJeYJzfB2b58YwTdnpmSHG6oMrWp_Ie-P8cYkHgmNmI_Q1KIYuWYwA6NRqL26rC5CwN7irxn2sgEwShBNeqwhJlA","not-before-policy":0,"session_state":"0fddfad7-805b-4371-9009-fb1927afdb02","scope":"openid email profile"}`)
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRequestTokenKeycloakErrorUnmarshal(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ResBody = []byte(`"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiNmZhMWNlMzEtZjkxNi00NTI2LWJlZDItYjk0NDg0MGFhMWUyIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJuZ3NpX2FwaSIsInNlc3Npb25fc3RhdGUiOiIwZmRkZmFkNy04MDViLTQzNzEtOTAwOS1mYjE5MjdhZmRiMDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWZpd2FyZV9zZXJ2aWNlIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.IbNIe7fAO7Q4ei4E7ICoNKMCEuEO1ebxP8zWM222ar22vF4Mx46UR4q9Qfc0Zrhdv3BG1bxwN8G6YLJyVx_ws3fRi0vFX_wZXRlVboKGo_4aBQBQb_rxRgMYDH3S5dQp2JPBwUPVznAz6M66zJM94G3ZUwlPB2mF-UfY_jlFxWUccN3OuFN91dEfjIxYwXL4T5ymdm2BwcZUuYDKDps15j7lcK-UC5tqpOzmYYlwxsrwMFVbWKpSYc-SJ3_Wz_Yj-m6TChXsDMRS9UkWuatmfq-i00b_AJgCo7B-bAUwt5YbW8KQGT-WN_as3TpfT6VuR8aLGwUg_00YcAAMdumUUA","expires_in":300,"refresh_expires_in":1800,"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJjOTUzMTYwYy0yMGEyLTRiNDQtYmI5NC01ZDgyNjQ2ODZmMzUifQ.eyJleHAiOjE2MjU5NTkzMzksImlhdCI6MTYyNTk1NzUzOSwianRpIjoiYzYzZjhhYmEtYWIzOC00OWVkLWExYjYtMjNhZjcyYzQ1MmI4IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC9hdXRoL3JlYWxtcy9maXdhcmVfc2VydmljZSIsImF1ZCI6Imh0dHA6Ly9sb2NhbGhvc3QvYXV0aC9yZWFsbXMvZml3YXJlX3NlcnZpY2UiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoibmdzaV9hcGkiLCJzZXNzaW9uX3N0YXRlIjoiMGZkZGZhZDctODA1Yi00MzcxLTkwMDktZmIxOTI3YWZkYjAyIiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSJ9.rJ7sMlxKZ-IA0AoqKznZLZnNAK7xsKHoPIRc2owIsw0","token_type":"Bearer","id_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfX0NGVHZXTHRCNHBVMWVZdk1OZFlQd1BfZHZmcXlYSDFOd1JHMFljZGhNIn0.eyJleHAiOjE2MjU5NTc4MzksImlhdCI6MTYyNTk1NzUzOSwiYXV0aF90aW1lIjowLCJqdGkiOiI4YzA4YWJiYy05YjI1LTRmMDgtOWMwYi1iMDRmNmEwYWJjNmMiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0L2F1dGgvcmVhbG1zL2Zpd2FyZV9zZXJ2aWNlIiwiYXVkIjoibmdzaV9hcGkiLCJzdWIiOiJhNDg4M2JiOS02MzBmLTQwYjktODY3Yi1iZmJlM2Q3OGU5NjMiLCJ0eXAiOiJJRCIsImF6cCI6Im5nc2lfYXBpIiwic2Vzc2lvbl9zdGF0ZSI6IjBmZGRmYWQ3LTgwNWItNDM3MS05MDA5LWZiMTkyN2FmZGIwMiIsImF0X2hhc2giOiJfalh0eDQxbVlobWRrZGsyc1h5Y0FnIiwiYWNyIjoiMSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiZml3YXJlIn0.JYjRzJckKsYhVChFw1Y2SGHJgY_PXhjTZRJMDzls7Xj5J7nJ5IiglHY1cSI-A1wzcJHttJYvGwKRK9QziVZfQh2LQoQSChnhHYY1Uq0VMrWfuOgofCQSqOmoXJiu7VwFRWVbg6RNVxlgT_z1SyncXQrCmjtfUBUKsBSSAZxOucVgDmHptT_JcNqKPeeo8-7PelDtx4PZDkf4Qf_77qHFPy0cwXio57UFAGtJAsztxd6nwZ6Q0QQY7XxCQLLALIsJeYJzfB2b58YwTdnpmSHG6oMrWp_Ie-P8cYkHgmNmI_Q1KIYuWYwA6NRqL26rC5CwN7irxn2sgEwShBNeqwhJlA","not-before-policy":0,"session_state":"0fddfad7-805b-4371-9009-fb1927afdb02","scope":"openid email profile"}`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		actual := "json: cannot unmarshal string into Go value of type ngsilib.KeycloakToken Field: (14) \"access_token\":\"eyJhbGciOiJSU"
		assert.Equal(t, actual, ngsiErr.Message)
	}
}

func TestRequestTokenKeycloakErrorHTTPStatus(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ResBody = []byte(`bad request`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestRequestTokenKeycloakErrorHTTPStatusUnauthorized(t *testing.T) {
	ngsi := testNgsiLibInit()

	ngsi.tokenList = tokenInfoList{}
	filename := ""
	ngsi.CacheFile = &MockIoLib{filename: &filename}
	ngsi.LogWriter = &bytes.Buffer{}
	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusUnauthorized
	reqRes.ResBody = []byte(`Unauthorized`)
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock
	ngsi.tokenList["token1"] = TokenInfo{}
	ngsi.tokenList["token2"] = TokenInfo{}

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CKeycloak, IdmHost: "http://idm", Username: "fiware", Password: "1234", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{}

	_, err := idm.requestToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "error  Unauthorized", ngsiErr.Message)
	}
}

func TestRevokeTokenKeycloak(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte("token=1a8346b8df2881c8b3407b0f39c80d1374204b93&client_id=0000&client_secret=1111")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{RefreshToken: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	assert.NoError(t, err)
}

func TestRevokeTokenKeycloakErrorHTTP(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusOK
	reqRes.ReqData = []byte("token=1a8346b8df2881c8b3407b0f39c80d1374204b93&client_id=0000&client_secret=1111")
	reqRes.Err = errors.New("http error")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{RefreshToken: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "http error", ngsiErr.Message)
	}
}

func TestRevokeTokenKeycloakErrorHTTPStatus(t *testing.T) {
	ngsi := testNgsiLibInit()

	reqRes := MockHTTPReqRes{}
	reqRes.Res.StatusCode = http.StatusBadRequest
	reqRes.ReqData = []byte("token=1a8346b8df2881c8b3407b0f39c80d1374204b93&client_id=0000&client_secret=1111")
	reqRes.ResBody = []byte("bad request")
	mock := NewMockHTTP()
	mock.ReqRes = append(mock.ReqRes, reqRes)
	ngsi.HTTP = mock

	client := &Client{Server: &Server{ServerHost: "http://orion/", IdmType: CBasic, IdmHost: "http://idm", Username: "fiware", ClientID: "0000", ClientSecret: "1111"}}
	idm := &idmKeycloak{}
	tokenInfo := &TokenInfo{RefreshToken: "1a8346b8df2881c8b3407b0f39c80d1374204b93"}

	err := idm.revokeToken(ngsi, client, tokenInfo)

	if assert.Error(t, err) {
		ngsiErr := err.(*LibError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error  bad request", ngsiErr.Message)
	}
}

func TestGetAuthHeaderKeycloak(t *testing.T) {
	idm := &idmKeycloak{}

	key, value := idm.getAuthHeader("9e7067026d0aac494e8fedf66b1f585e79f52935")

	assert.Equal(t, "Authorization", key)
	assert.Equal(t, "Bearer 9e7067026d0aac494e8fedf66b1f585e79f52935", value)
}
