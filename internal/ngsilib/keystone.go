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

// https://thinking-cities.readthedocs.io/en/latest/authentication_api/index.html

type keyStoneDomain struct {
	Name string `json:"name"`
}

type keyStoneProject struct {
	Domain keyStoneDomain `json:"domain"`
	Name   string         `json:"name"`
}

type keyStoneUser struct {
	Domain   keyStoneDomain `json:"domain,omitempty"`
	Name     string         `json:"name"`
	Password string         `json:"password"`
}

type keyStoneRequest struct {
	Auth struct {
		Identity struct {
			Methods  []string `json:"methods"`
			Password struct {
				User keyStoneUser `json:"user"`
			} `json:"password"`
		} `json:"identity"`
		Scope struct {
			Project *keyStoneProject `json:"project,omitempty"`
			Domain  *keyStoneDomain  `json:"domain,omitempty"`
		} `json:"scope"`
	} `json:"auth"`
}

type keyStoneRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type keyStoneEndpoint struct {
	RegionID  string `json:"region_id"`
	URL       string `json:"url"`
	Region    string `json:"region"`
	Interface string `json:"interface"`
	ID        string `json:"id"`
}

type keyStoneCatalog struct {
	Endpoints []keyStoneEndpoint `json:"endpoints"`
	Type      string             `json:"type"`
	ID        string             `json:"id"`
	Name      string             `json:"name"`
}

// KeyStoneToken is ...
type KeyStoneToken struct {
	Token struct {
		Domain struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"domain"`
		Methods   []string          `json:"methods"`
		Roles     []keyStoneRole    `json:"roles"`
		ExpiresAt string            `json:"expires_at"`
		Catalog   []keyStoneCatalog `json:"catalog"`
		Extras    struct {
			PasswordCreationTime   string `json:"password_creation_time"`
			LastLoginAttemptTime   string `json:"last_login_attempt_time"`
			PwdUserInBlacklist     bool   `json:"pwd_user_in_blacklist"`
			PasswordExpirationTime string `json:"password_expiration_time"`
		} `json:"extras"`
		User struct {
			PasswordExpiresAt string `json:"password_expires_at"`
			Domain            struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"domain"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"user"`
		AuditIds []string `json:"audit_ids"`
		IssuedAt string   `json:"issued_at"`
	} `json:"token"`
}

func getKeyStoneTokenRequest(name, password, tenant, scorpe string) string {
	var r keyStoneRequest

	r.Auth.Identity.Methods = []string{"password"}
	r.Auth.Identity.Password.User.Name = name
	r.Auth.Identity.Password.User.Password = password

	if tenant != "" {
		r.Auth.Identity.Password.User.Domain.Name = tenant
		d := keyStoneDomain{Name: tenant}
		if scorpe != "" {
			p := keyStoneProject{Name: scorpe, Domain: d}
			r.Auth.Scope.Project = &p
		} else {
			r.Auth.Scope.Domain = &d
		}
	} else {
		if scorpe != "" {
			p := keyStoneProject{Name: scorpe}
			r.Auth.Scope.Project = &p
		} else {
			d := keyStoneDomain{}
			r.Auth.Scope.Domain = &d
		}
	}

	b, _ := JSONMarshal(&r)

	return string(b)
}
