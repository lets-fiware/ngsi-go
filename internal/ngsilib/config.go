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
	"path/filepath"
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
	DefaultValues Settings     `json:"settings"`
	Brokers       BrokerList   `json:"brokers"`
	Contexts      ContextsInfo `json:"contexts"`
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

	if io.FileName() == nil {
		home, err := getConfigDir(io)
		if err != nil {
			return &NgsiLibError{funcName, 1, err.Error(), err}
		}
		s := filepath.Join(home, configFileName)
		io.SetFileName(&s)
	} else {
		if *io.FileName() != "" {
			s, err := io.FilePathAbs(*io.FileName())
			if err != nil {
				return &NgsiLibError{funcName, 2, err.Error() + " " + *io.FileName(), err}
			}
			io.SetFileName(&s)
		}
	}

	if existsFile(io, *io.FileName()) {
		b, err := ngsi.FileReader.ReadFile(*io.FileName())
		if err != nil {
			return &NgsiLibError{funcName, 3, err.Error(), err}
		}

		ngsiConfig := NgsiConfig{}
		err = JSONUnmarshal(b, &ngsiConfig)
		if err != nil {
			return &NgsiLibError{funcName, 4, err.Error(), err}
		}

		ngsi.PreviousArgs = &ngsiConfig.DefaultValues
		if ngsi.BatchFlag != nil && *ngsi.BatchFlag {
			ngsi.PreviousArgs = &Settings{UsePreviousArgs: false}
		}
		ngsi.brokerList = ngsiConfig.Brokers
		ngsi.contextList = ngsiConfig.Contexts
	}

	if ngsi.brokerList == nil {
		ngsi.brokerList = make(BrokerList)
	}
	if ngsi.contextList == nil {
		ngsi.contextList = make(ContextsInfo)
		ngsi.contextList["etsi"] = "https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context-v1.3.jsonld"
		ngsi.contextList["ld"] = "https://schema.lab.fiware.org/ld/context"
	}

	errflag := false
	for k, v := range ngsi.brokerList {
		if err := gNGSI.checkAllParams(v); err != nil {
			fmt.Fprintf(gNGSI.LogWriter, "%s in %s\n", err, k)
			errflag = true
		}
	}
	for k, v := range ngsi.contextList {
		switch v.(type) {
		default:
			fmt.Fprintf(gNGSI.LogWriter, "%s is neither url nor json\n", k)
			errflag = true
		case string:
			s := v.(string)
			if !IsHTTP(s) {
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
		return &NgsiLibError{funcName, 5, "error in config file", nil}
	}
	return nil
}

func (ngsi *NGSI) saveConfigFile() error {
	const funcName = "saveConfigFile"

	io := ngsi.ConfigFile

	if *io.FileName() == "" {
		return nil
	}

	config := make(map[string]interface{})

	config["settings"] = *ngsi.PreviousArgs
	config["brokers"] = ngsi.brokerList
	config["contexts"] = ngsi.contextList

	err := io.OpenFile(oWRONLY|oCREATE, 0600)
	if err != nil {
		return &NgsiLibError{funcName, 1, err.Error(), err}
	}
	defer io.Close()

	if err := io.Truncate(0); err != nil {
		return &NgsiLibError{funcName, 2, err.Error(), err}
	}

	err = io.Encode(&config)
	if err != nil {
		return &NgsiLibError{funcName, 3, err.Error(), err}
	}

	return nil
}

// GetPreviousArgs is ...
func (ngsi *NGSI) GetPreviousArgs() *Settings {
	return ngsi.PreviousArgs
}

// SavePreviousArgs is ...
func (ngsi *NGSI) SavePreviousArgs() error {
	if ngsi.PreviousArgs.UsePreviousArgs {
		return ngsi.saveConfigFile()
	}
	return nil
}
