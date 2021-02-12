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

package ngsicmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type keyrockUserItems struct {
	ID                string      `json:"id,omitempty"`
	Username          string      `json:"username,omitempty"`
	Email             string      `json:"email,omitempty"`
	Password          string      `json:"password,omitempty"`
	Image             string      `json:"image,omitempty"`
	Enabled           bool        `json:"enabled,omitempty"`
	Admin             bool        `json:"admin,omitempty"`
	Gravatar          bool        `json:"gravatar,omitempty"`
	StartersTourEnded bool        `json:"starters_tour_ended,omitempty"`
	DatePassword      string      `json:"date_password,omitempty"`
	Description       string      `json:"description,omitempty"`
	Website           string      `json:"website,omitempty"`
	Extra             interface{} `json:"extra,omitempty"`
}

type keyrockUser struct {
	User keyrockUserItems `json:"user"`
}

type keyrockUsers struct {
	Users []keyrockUserItems `json:"users"`
}

func usersList(c *cli.Context) error {
	const funcName = "usersList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	client.SetPath("/v1/users")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 5, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var users keyrockUsers
		err := ngsilib.JSONUnmarshal(body, &users)
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		for _, user := range users.Users {
			fmt.Fprintln(ngsi.StdWriter, user.ID)
		}
	}

	return nil
}

func usersGet(c *cli.Context) error {
	const funcName = "usersGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 3, "specify user id", nil}
	}
	path := "/v1/users/" + c.String("uid")
	client.SetPath(path)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 6, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func usersCreate(c *cli.Context) error {
	const funcName = "usersCreate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !isSetAND(c, []string{"username", "email", "password"}) {
		return &ngsiCmdError{funcName, 3, "specify username, email and password", nil}
	}

	user := keyrockUser{}
	user.User.Username = c.String("username")
	user.User.Password = c.String("password")
	user.User.Email = c.String("email")

	b, err := ngsilib.JSONMarshal(user)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	client.SetHeader("Content-Type", "application/json")
	client.SetPath("/v1/users")

	res, body, err := client.HTTPPost(b)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusCreated {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("verbose") || c.Bool("pretty") {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 7, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprint(ngsi.StdWriter, string(body))
		}
	} else {
		var user keyrockUser
		err = ngsilib.JSONUnmarshal(body, &user)
		if err != nil {
			return &ngsiCmdError{funcName, 8, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, user.User.ID)
	}

	return nil
}

func usersUpdate(c *cli.Context) error {
	const funcName = "usersUpdate"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 3, "specify user id", nil}
	}
	client.SetPath("/v1/users/" + c.String("uid"))

	client.SetHeader("Content-Type", "application/json")

	b, err := setUsersParam(c)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}

	res, body, err := client.HTTPPatch(b)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 6, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		return nil
	}

	fmt.Fprint(ngsi.StdWriter, string(body))

	return nil
}

func usersDelete(c *cli.Context) error {
	const funcName = "usersDelete"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"keyrock"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if !c.IsSet("uid") {
		return &ngsiCmdError{funcName, 3, "specify user id", nil}
	}
	client.SetPath("/v1/users/" + c.String("uid"))

	res, body, err := client.HTTPDelete()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}

func setUsersParam(c *cli.Context) ([]byte, error) {
	const funcName = "setUsersParam"

	user := keyrockUser{}

	if c.IsSet("username") {
		user.User.Username = c.String("username")
	}
	if c.IsSet("password") {
		user.User.Password = c.String("password")
	}
	if c.IsSet("email") {
		user.User.Email = c.String("email")
	}
	if c.IsSet("description") {
		user.User.Description = c.String("description")
	}
	if c.IsSet("website") {
		user.User.Website = c.String("website")
	}
	if c.IsSet("gravatar") {
		b, err := getBool(c, "gravatar")
		if err != nil {
			return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		user.User.Gravatar = b
	}
	if c.IsSet("extra") {
		user.User.Extra = c.String("extra")
	}

	b, err := ngsilib.JSONMarshal(user)
	if err != nil {
		return nil, &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	return b, nil
}

func getBool(c *cli.Context, name string) (bool, error) {
	const funcName = "getBool"

	b := strings.ToLower(c.String(name))
	if ngsilib.Contains([]string{"true", "on"}, b) {
		return true, nil
	} else if ngsilib.Contains([]string{"false", "off"}, b) {
		return false, nil
	} else {
		return false, &ngsiCmdError{funcName, 1, "specify either true or false to --" + name, nil}
	}
}
