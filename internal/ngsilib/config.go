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

import (
	"fmt"
	"path/filepath"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// Settings is ...
type Settings struct {
	UsePreviousArgs bool   `json:"usePreviousArgs"`
	Syslog          string `json:"syslog"`
	Stderr          string `json:"stderr"`
	Logfile         string `json:"logfile"`
	Loglevel        string `json:"loglevel"`
	CacheFile       string `json:"cachefile"`
	Host            string `json:"host"`
	Tenant          string `json:"tenant"`
	Scope           string `json:"scope"`
	Token           string `json:"token"`
}

// NgsiConfig is ...
type NgsiConfig struct {
	Version           string       `json:"version"`
	DefaultValues     Settings     `json:"settings"`
	DeprecatedBrokers ServerList   `json:"brokers,omitempty"`
	Servers           ServerList   `json:"servers"`
	Contexts          ContextsInfo `json:"contexts"`
}

// var configFile string

// var defaultValues *Default
var configFileName = "ngsi-go-config.json"

// InitConfig is ...
func (ngsi *NGSI) InitConfig(file *string) error {
	ngsi.ConfigFile.SetFileName(file)
	return initConfig(ngsi, ngsi.ConfigFile)
}

func initConfig(ngsi *NGSI, io IoLib) error {
	const funcName = "initConfig"

	saveFlag := false

	if io.FileName() == nil {
		home, err := getConfigDir(ngsi.ConfigDir, io)
		if err != nil {
			return ngsierr.New(funcName, 1, err.Error(), err)
		}
		s := filepath.Join(home, configFileName)
		io.SetFileName(&s)
	} else {
		if *io.FileName() != "" {
			s, err := io.FilePathAbs(*io.FileName())
			if err != nil {
				return ngsierr.New(funcName, 2, err.Error()+" "+*io.FileName(), err)
			}
			io.SetFileName(&s)
		}
	}

	if existsFile(io, *io.FileName()) {
		b, err := ngsi.FileReader.ReadFile(*io.FileName())
		if err != nil {
			return ngsierr.New(funcName, 3, err.Error(), err)
		}

		ngsiConfig := NgsiConfig{}
		if len(b) > 0 {
			err = JSONUnmarshal(b, &ngsiConfig)
			if err != nil {
				return ngsierr.New(funcName, 4, err.Error(), err)
			}
		}

		ngsi.configVresion = ngsiConfig.Version
		if ngsi.configVresion == "" {
			migration(&ngsiConfig)
			ngsi.configVresion = "1"
			saveFlag = true
		}

		ngsi.PreviousArgs = &ngsiConfig.DefaultValues
		if ngsi.BatchFlag != nil && *ngsi.BatchFlag {
			ngsi.PreviousArgs = &Settings{UsePreviousArgs: false}
		}
		ngsi.ServerList = ngsiConfig.Servers
		ngsi.contextList = ngsiConfig.Contexts
	}

	if ngsi.configVresion != "1" {
		return ngsierr.New(funcName, 5, "error: config file version", nil)
	}

	if ngsi.ServerList == nil {
		ngsi.ServerList = make(ServerList)
	}
	if ngsi.contextList == nil {
		ngsi.contextList = make(ContextsInfo)
		ngsi.contextList["etsi1.0"] = "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"
		ngsi.contextList["etsi1.3"] = "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"
		ngsi.contextList["etsi1.4"] = "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.4.jsonld"
		ngsi.contextList["ld"] = "https://schema.lab.fiware.org/ld/context"
	}

	errflag := false
	for k, v := range ngsi.ServerList {
		if err := gNGSI.checkAllParams(v); err != nil {
			fmt.Fprintf(gNGSI.LogWriter, "%s in %s\n", err, k)
			errflag = true
		}
	}
	for k, v := range ngsi.contextList {
		switch v := v.(type) {
		default:
			fmt.Fprintf(gNGSI.LogWriter, "%s is neither url nor json\n", k)
			errflag = true
		case string:
			if !IsHTTP(v) {
				fmt.Fprintf(gNGSI.LogWriter, "%s is not url\n", k)
				errflag = true
			}
		case []interface{}, map[string]interface{}:
			_, err := JSONMarshal(v)
			if err != nil {
				fmt.Fprintf(gNGSI.LogWriter, "%s is not json\n", k)
				errflag = true
			}
		}
	}
	if errflag {
		return ngsierr.New(funcName, 6, "error in config file", nil)
	}

	if saveFlag {
		if err := ngsi.saveConfigFile(); err != nil {
			return ngsierr.New(funcName, 7, err.Error(), err)
		}
	}
	return nil
}

func (ngsi *NGSI) saveConfigFile() (err error) {
	const funcName = "saveConfigFile"

	ngsi.Updated = false
	io := ngsi.ConfigFile

	if *io.FileName() == "" {
		return nil
	}

	config := make(map[string]interface{})

	config["version"] = ngsi.configVresion
	config["settings"] = *ngsi.PreviousArgs
	config["servers"] = ngsi.ServerList
	config["contexts"] = ngsi.contextList

	err = io.OpenFile(oWRONLY|oCREATE, 0600)
	if err != nil {
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	defer func() { _ = io.Close() }()

	if err := io.Truncate(0); err != nil {
		return ngsierr.New(funcName, 3, err.Error(), err)
	}

	err = io.Encode(&config)
	if err != nil {
		return ngsierr.New(funcName, 4, err.Error(), err)
	}

	return nil
}

// GetPreviousArgs is ...
func (ngsi *NGSI) GetPreviousArgs() *Settings {
	return ngsi.PreviousArgs
}

// SavePreviousArgs is ...
func (ngsi *NGSI) SavePreviousArgs() error {
	return ngsi.saveConfigFile()
}

func migration(config *NgsiConfig) {
	config.Servers = ServerList{}
	for k, v := range config.DeprecatedBrokers {
		v.ServerHost = v.DeprecatedBrokerHost
		v.DeprecatedBrokerHost = ""
		config.Servers[k] = v
	}
	config.DeprecatedBrokers = ServerList{}
}
