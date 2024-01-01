/*
MIT License

Copyright (c) 2020-2024 Kazuhito Suda

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

package wirecloud

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsicli"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
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

type wireCloudWiringEndPoint struct {
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
	Authors         []wireCloudAuthor       `json:"authors"`
	Changelog       string                  `json:"changelog"`
	Contributors    []wireCloudAuthor       `json:"contributors"`
	DefaultLang     string                  `json:"default_lang"`
	Description     string                  `json:"description"`
	Doc             string                  `json:"doc"`
	Email           string                  `json:"email"`
	Homepage        string                  `json:"homepage"`
	Image           string                  `json:"image"`
	Issuetracker    string                  `json:"issuetracker"`
	JsFiles         []string                `json:"js_files"`
	License         string                  `json:"license"`
	Licenseurl      string                  `json:"licenseurl"`
	Longdescription string                  `json:"longdescription"`
	Name            string                  `json:"name"`
	Preferences     interface{}             `json:"preferences"`
	Properties      []wireCloudProperty     `json:"properties"`
	Requirements    []wireCloudRequirement  `json:"requirements"`
	Smartphoneimage string                  `json:"smartphoneimage"`
	Title           string                  `json:"title"`
	Type            string                  `json:"type"`
	Vendor          string                  `json:"vendor"`
	Version         string                  `json:"version"`
	Wiring          wireCloudWiringEndPoint `json:"wiring"`
}

type wireCloudResources map[string]wireCloudResource

func wireCloudResourcesList(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "resourceList"

	client.SetPath("/api/resources")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	noOption := !c.IsSetOR([]string{"widget", "operator", "mashup"})

	if noOption && c.IsSetOR([]string{"json", "pretty"}) {
		if c.IsSet("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 3, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(body))
		}
	} else {
		resource := make(wireCloudResources)
		err = ngsilib.JSONUnmarshal(body, &resource)
		if err != nil {
			return ngsierr.New(funcName, 4, err.Error(), err)
		}

		r := make(wireCloudResources)
		widget := c.Bool("widget")
		operator := c.Bool("operator")
		mashup := c.Bool("mashup")
		for key, value := range resource {
			if (widget && value.Type == "widget") ||
				(operator && value.Type == "operator") ||
				(mashup && value.Type == "mashup") ||
				(noOption) {
				if c.IsSet("vender") && value.Vendor != c.String("vender") {
					continue
				}
				if c.IsSet("name") && value.Name != c.String("name") {
					continue
				}
				if c.IsSet("version") && value.Version != c.String("version") {
					continue
				}
				r[key] = value
			}
		}

		if c.IsSetOR([]string{"json", "pretty"}) {
			b, err := ngsilib.JSONMarshal(r)
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
			if c.Bool("pretty") {
				newBuf := new(bytes.Buffer)
				err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
				if err != nil {
					return ngsierr.New(funcName, 6, err.Error(), err)
				}
				fmt.Fprintln(ngsi.StdWriter, newBuf.String())
			} else {
				fmt.Fprintln(ngsi.StdWriter, string(b))
			}
		} else {
			keys := make([]string, len(r))
			i := 0
			for key := range r {
				keys[i] = key
				i++
			}
			sort.Strings(keys)
			for _, key := range keys {
				fmt.Fprintln(ngsi.StdWriter, key)
			}
		}
	}

	return nil
}

func wireCloudResourceGet(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "resourceGet"

	var macName string

	if c.IsSetAND([]string{"vender", "name", "version"}) && c.Args().Len() == 0 {
		macName = c.String("vender") + "/" + c.String("name") + "/" + c.String("version")
	} else if c.Args().Len() == 1 {
		macName = c.Args().Get(0)
	} else {
		return ngsierr.New(funcName, 1, "argument error", nil)
	}

	client.SetPath("/api/resources")

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 3, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	resources := make(wireCloudResources)
	err = ngsilib.JSONUnmarshal(body, &resources)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}

	mac, ok := resources[macName]
	if !ok {
		return ngsierr.New(funcName, 5, macName+" not found", nil)
	}

	b, err := ngsilib.JSONMarshal(mac)
	if err != nil {
		return ngsierr.New(funcName, 6, err.Error(), err)
	}

	if c.Bool("pretty") {
		newBuf := new(bytes.Buffer)
		err := ngsi.JSONConverter.Indent(newBuf, b, "", "  ")
		if err != nil {
			return ngsierr.New(funcName, 7, err.Error(), err)
		}
		fmt.Fprintln(ngsi.StdWriter, newBuf.String())
	} else {
		fmt.Fprintln(ngsi.StdWriter, string(b))
	}

	return nil
}

func wireCloudResourceDownload(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "resourceDownload"

	var mashup string

	if c.IsSetAND([]string{"vender", "name", "version"}) && c.Args().Len() == 0 {
		mashup = fmt.Sprintf("%s/%s/%s", c.String("vender"), c.String("name"), c.String("version"))
	} else if c.Args().Len() == 1 {
		mashup = c.Args().Get(0)
	} else {
		return ngsierr.New(funcName, 1, "argument error", nil)
	}

	client.SetPath("/api/resource/" + mashup)

	res, body, err := client.HTTPGet()
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return ngsierr.New(funcName, 3, mashup+" not found", nil)
		}
		return ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	fileName := strings.Replace(mashup, "/", "_", -1) + ".wgt"

	err = ngsi.Ioutil.WriteFile(fileName, body, 0644)
	if err != nil {
		return ngsierr.New(funcName, 5, err.Error(), err)
	}

	return nil
}

func wireCloudResourceInstall(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "resourceInstall"

	file := c.String("file")

	file, err := ngsi.FilePath.FilePathAbs(file)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	fileName := ngsi.FilePath.FilePathBase(file)

	b, err := ngsi.Ioutil.ReadFile(file)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}

	_, name, err := getMacName(ngsi, b)
	if err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	exists, err := existsMac(ngsi, client, name)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}

	if exists {
		if c.Bool("overwrite") {
			err = uninstallMac(ngsi, client, name)
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
		} else {
			return ngsierr.New(funcName, 6, name+" already exists", nil)
		}
	}

	var body bytes.Buffer
	m := ngsi.MultiPart.NewWriter(&body)

	contentType, err := makeMultipart(ngsi, m, fileName, b)
	if err != nil {
		return ngsierr.New(funcName, 7, err.Error(), err)
	}

	if c.Bool("public") {
		v := url.Values{}
		v.Set("public", "true")
		client.SetQuery(&v)
	}

	url := "/api/resources"
	client.SetPath(url)
	client.SetHeader("Content-Type", contentType)

	res, resBody, err := client.HTTPPost(body.Bytes())
	if err != nil {
		return ngsierr.New(funcName, 8, err.Error(), err)
	}
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		return ngsierr.New(funcName, 9, fmt.Sprintf("error %s %s", res.Status, string(resBody)), nil)
	}

	if c.IsSetOR([]string{"json", "pretty"}) {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, resBody, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 10, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(resBody))
		}
	}

	return nil
}

func makeMultipart(ngsi *ngsilib.NGSI, mw ngsilib.MultiPartLib, fileName string, b []byte) (string, error) {
	const funcName = "makeMultipart"

	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "application/octet-stream")
	mh.Set("Content-Disposition", "form-data; name=\"file\"; filename=\""+fileName+"\"")
	pw, err := mw.CreatePart(mh)
	if err != nil {
		return "", ngsierr.New(funcName, 1, err.Error(), err)
	}
	_, err = ngsi.Ioutil.Copy(pw, bytes.NewReader(b))
	if err != nil {
		return "", ngsierr.New(funcName, 2, err.Error(), err)
	}
	contentType := mw.FormDataContentType()
	err = mw.Close()
	if err != nil {
		return "", ngsierr.New(funcName, 3, err.Error(), err)
	}

	return contentType, nil
}

func wireCloudResourceUninstall(c *ngsicli.Context, ngsi *ngsilib.NGSI, client *ngsilib.Client) error {
	const funcName = "resourceUninstall"

	var mashup string

	if c.IsSetAND([]string{"vender", "name"}) && c.Args().Len() == 0 {
		if c.IsSet("version") {
			mashup = fmt.Sprintf("%s/%s/%s", c.String("vender"), c.String("name"), c.String("version"))
		} else {
			mashup = fmt.Sprintf("%s/%s", c.String("vender"), c.String("name"))
		}
	} else if c.Args().Len() == 1 {
		mashup = c.Args().Get(0)
	} else {
		return ngsierr.New(funcName, 1, "argument error", nil)
	}

	if !c.IsSet("run") {
		fmt.Fprintf(ngsi.StdWriter, "%s will be uninstalled. run uninstall with --run option\n", mashup)
		return nil
	}

	affected := c.IsSetOR([]string{"json", "pretty"})
	if affected {
		v := url.Values{}
		v.Set("affected", "true")
		client.SetQuery(&v)
	}

	client.SetPath("/api/resource/" + mashup)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 2, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return ngsierr.New(funcName, 3, mashup+" not found", nil)
		}
		return ngsierr.New(funcName, 4, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	if affected {
		if c.Bool("pretty") {
			newBuf := new(bytes.Buffer)
			err := ngsi.JSONConverter.Indent(newBuf, body, "", "  ")
			if err != nil {
				return ngsierr.New(funcName, 5, err.Error(), err)
			}
			fmt.Fprintln(ngsi.StdWriter, newBuf.String())
		} else {
			fmt.Fprintln(ngsi.StdWriter, string(body))
		}
	}

	return nil
}

var mashupRegEx = regexp.MustCompile(`<(widget|operator|mashup)\s+xmlns="([^"]+)"\s+vendor="([^"]+)"\s+name="([^"]+)"\s+version="([^"]+)"\s*>`)

func uninstallMac(ngsi *ngsilib.NGSI, client *ngsilib.Client, name string) error {
	const funcName = "uninstallMac"

	client.SetPath("/api/resource/" + name)

	res, body, err := client.HTTPDelete(nil)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusNotFound {
		return ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return nil
}

func getMacName(ngsi *ngsilib.NGSI, b []byte) (string, string, error) {
	const funcName = "getMacName"

	r := bytes.NewReader(b)
	zr, err := ngsi.ZipLib.NewReader(r, r.Size())
	if err != nil {
		return "", "", ngsierr.New(funcName, 1, err.Error(), err)
	}

	for _, f := range zr.File {
		if f.Name == "config.xml" {
			rc, _ := f.Open()
			defer func() { _ = rc.Close() }()
			return getFromConfigXML(ngsi, rc, f.UncompressedSize)
		}
	}

	return "", "", ngsierr.New(funcName, 2, "config.xml not found", nil)
}

func getFromConfigXML(ngsi *ngsilib.NGSI, rc io.ReadCloser, size uint32) (string, string, error) {
	const funcName = "configXML"

	buf := make([]byte, size)
	_, err := ngsi.Ioutil.ReadFull(rc, buf)
	if err != nil {
		return "", "", ngsierr.New(funcName, 1, err.Error(), err)
	}

	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		s := strings.TrimSpace(line)
		if strings.HasPrefix(s, "<mashup") ||
			strings.HasPrefix(s, "<widget") ||
			strings.HasPrefix(s, "<operator") {
			mac := mashupRegEx.FindStringSubmatch(s)
			if len(mac) == 6 {
				mashup := mac[1]
				name := fmt.Sprintf("%s/%s/%s", mac[3], mac[4], mac[5])
				return mashup, name, nil
			}
			return "", "", ngsierr.New(funcName, 2, "config.xml error", nil)
		}
	}
	return "", "", ngsierr.New(funcName, 3, "config.xml error", nil)
}

func existsMac(ngsi *ngsilib.NGSI, client *ngsilib.Client, name string) (bool, error) {
	const funcName = "existMac"

	client.SetPath("/api/resource/" + name)

	res, body, err := client.HTTPGet()
	if err != nil {
		return false, ngsierr.New(funcName, 1, err.Error(), err)
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNotFound {
		return false, ngsierr.New(funcName, 2, fmt.Sprintf("error %s %s", res.Status, string(body)), nil)
	}

	return res.StatusCode != http.StatusNotFound, nil
}
