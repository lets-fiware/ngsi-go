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
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/urfave/cli/v2"
)

type wireCloudAuthor struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type wireCloudEndPoint struct {
	Actionlabel string `json:"actionlabel"`
	Description string `json:"description"`
	Friendcode  string `json:"friendcode"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type wireCloudWiring struct {
	Inputs  []wireCloudEndPoint `json:"inputs"`
	Outputs []wireCloudEndPoint `json:"outputs"`
}

/*
type wireCloudPreference struct {
	Default     string      `json:"default"`
	Description string      `json:"description"`
	Label       string      `json:"label"`
	Multiuser   bool        `json:"multiuser"`
	Name        string      `json:"name"`
	Readonly    bool        `json:"readonly"`
	Required    bool        `json:"required"`
	Secure      bool        `json:"secure"`
	Type        string      `json:"type"`
	Value       interface{} `json:"value"`
}

type wireCloudMashupPreference struct {
	Baselayout string `json:"baselayout"`
}
*/

type wireCloudProperty struct {
	Default     string `json:"default"`
	Description string `json:"description"`
	Label       string `json:"label"`
	Multiuser   bool   `json:"multiuser"`
	Name        string `json:"name"`
	Secure      bool   `json:"secure"`
	Type        string `json:"type"`
}

type wireCloudRequirement struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type wireCloudResource struct {
	Authors         []wireCloudAuthor      `json:"authors"`
	Changelog       string                 `json:"changelog"`
	Contributors    []wireCloudAuthor      `json:"contributors"`
	DefaultLang     string                 `json:"default_lang"`
	Description     string                 `json:"description"`
	Doc             string                 `json:"doc"`
	Email           string                 `json:"email"`
	Homepage        string                 `json:"homepage"`
	Image           string                 `json:"image"`
	Issuetracker    string                 `json:"issuetracker"`
	JsFiles         []string               `json:"js_files"`
	License         string                 `json:"license"`
	Licenseurl      string                 `json:"licenseurl"`
	Longdescription string                 `json:"longdescription"`
	Name            string                 `json:"name"`
	Preferences     interface{}            `json:"preferences"`
	Properties      []wireCloudProperty    `json:"properties"`
	Requirements    []wireCloudRequirement `json:"requirements"`
	Smartphoneimage string                 `json:"smartphoneimage"`
	Title           string                 `json:"title"`
	Type            string                 `json:"type"`
	Vendor          string                 `json:"vendor"`
	Version         string                 `json:"version"`
	Wiring          wireCloudWiring        `json:"wiring"`
}

type wireCloudResources map[string]wireCloudResource

func wireCloudResourcesList(c *cli.Context) error {
	const funcName = "resourceList"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"wirecloud"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	client.SetPath("/api/resources")

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	resource := make(wireCloudResources)
	err = ngsilib.JSONUnmarshal(body, &resource)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}

	r := make(wireCloudResources)
	widget := c.Bool("widget")
	operator := c.Bool("operator")
	mashup := c.Bool("mashup")
	noOption := !isSetOR(c, []string{"widget", "operator", "mashup"})
	for key, value := range resource {
		if (widget && value.Type == "widget") ||
			(operator && value.Type == "operator") ||
			(mashup && value.Type == "mashup") ||
			(noOption) {
			if c.IsSet("vender") && value.Vendor != c.String("vender") {
				continue
			}
			if c.IsSet("name") && value.Vendor != c.String("name") {
				continue
			}
			if c.IsSet("version") && value.Vendor != c.String("version") {
				continue
			}
			r[key] = value
		}
	}

	if isSetOR(c, []string{"json", "pretty"}) {
		b, err := ngsilib.JSONMarshal(r)
		if err != nil {
			return &ngsiCmdError{funcName, 7, err.Error(), err}
		}
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
			if err != nil {
				return &ngsiCmdError{funcName, 8, err.Error(), err}
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(b))
		}
	} else {
		for key := range r {
			fmt.Fprintln(ngsi.StdWriter, key)
		}
	}

	return nil
}

func wireCloudResourceGet(c *cli.Context) error {
	const funcName = "resourceGet"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"wirecloud"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	var mashup string

	if isSetAND(c, []string{"vender", "name", "version"}) && c.Args().Len() == 0 {
		mashup = fmt.Sprintf("%s/%s/%s", c.String("vender"), c.String("name"), c.String("version"))
	} else if c.Args().Len() == 1 {
		mashup = c.Args().Get(0)
	} else {
		return &ngsiCmdError{funcName, 3, "argument error", nil}
	}

	client.SetPath("/api/resource/" + mashup)

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
	} else {
		fmt.Fprintln(ngsi.StdWriter, string(body))
	}

	return nil
}

func wireCloudResourceDownload(c *cli.Context) error {
	const funcName = "resourceDownload"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"wirecloud"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	var mashup string

	if isSetAND(c, []string{"vender", "name", "version"}) && c.Args().Len() == 0 {
		mashup = fmt.Sprintf("%s/%s/%s", c.String("vender"), c.String("name"), c.String("version"))
	} else if c.Args().Len() == 1 {
		mashup = c.Args().Get(0)
	} else {
		return &ngsiCmdError{funcName, 3, "argument error", nil}
	}

	client.SetPath("/api/resource/" + mashup)

	res, body, err := client.HTTPGet()
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusOK {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	fileName := strings.Replace(mashup, "/", "_", -1) + ".wgt"

	err = ioutil.WriteFile(fileName, body, 0644)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}

	return nil
}

func wireCloudResourceInstall(c *cli.Context) error {
	const funcName = "resourceInstall"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"wirecloud"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	var file string
	if c.IsSet("file") && c.Args().Len() == 0 {
		file = c.String("file")
	} else if c.Args().Len() == 1 {
		file = c.Args().Get(0)
	} else {
		return &ngsiCmdError{funcName, 3, "argument error", nil}
	}
	file, err = filepath.Abs(file)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), nil}
	}
	fileName := filepath.Base(file)

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return &ngsiCmdError{funcName, 5, err.Error(), err}
	}

	var body bytes.Buffer

	contentType, err := makeMultipart(&body, fileName, b)
	if err != nil {
		return &ngsiCmdError{funcName, 6, err.Error(), err}
	}

	url := "/api/resource"
	if c.Bool("public") {
		url += "?public=true"
	}

	client.SetPath(url)
	client.SetHeader("Content-Type", contentType)

	res, resbody, err := client.HTTPPost(body.Bytes())
	if err != nil {
		return &ngsiCmdError{funcName, 7, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 8, fmt.Sprintf("error %s %s", res.Status, string(resbody)), nil}
	}

	return nil
}

func makeMultipart(buf *bytes.Buffer, fileName string, b []byte) (string, error) {
	const funcName = "makeMultipart"

	mw := multipart.NewWriter(buf)
	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "application/octet-stream")
	mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+fileName+"\"")
	pw, err := mw.CreatePart(mh)
	if err != nil {
		return "", &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	_, err = io.Copy(pw, bytes.NewReader(b))
	if err != nil {
		return "", &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	contentType := mw.FormDataContentType()
	err = mw.Close()
	if err != nil {
		return "", &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	return contentType, nil
}

func wireCloudResourceUninstall(c *cli.Context) error {
	const funcName = "resourceUninstall"

	ngsi, err := initCmd(c, funcName, true)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client, err := newClient(ngsi, c, false, []string{"wirecloud"})
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	var mashup string

	if isSetAND(c, []string{"vender", "name"}) && c.Args().Len() == 0 {
		if c.IsSet("version") {
			mashup = fmt.Sprintf("%s/%s/%s", c.String("vender"), c.String("name"), c.String("version"))
		} else {
			mashup = fmt.Sprintf("%s/%s", c.String("vender"), c.String("name"))
		}
	} else if c.Args().Len() == 1 {
		mashup = c.Args().Get(0)
	} else {
		return &ngsiCmdError{funcName, 3, "argument error", nil}
	}

	if !c.IsSet("run") {
		fmt.Fprintf(ngsi.StdWriter, "%s will be uninstalled. run uninstall with --run option\n", mashup)
		return nil
	}

	client.SetPath("/api/resource/" + mashup + "?affected=true")

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return &ngsiCmdError{funcName, 4, err.Error(), err}
	}
	if res.StatusCode != http.StatusNoContent {
		return &ngsiCmdError{funcName, 5, fmt.Sprintf("error %s %s", res.Status, string(body)), nil}
	}

	return nil
}
