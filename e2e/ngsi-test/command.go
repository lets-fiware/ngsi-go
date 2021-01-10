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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type cmdDef func([]string) error

var cmdTable map[string]cmdDef

func initCmdTable() {
	cmdTable = map[string]cmdDef{
		"```":   compareCmd,
		"ngsi":  ngsiCmd,
		"halt":  haltCmd,
		"http":  httpCmd,
		"print": printCmd,
		"sleep": sleepCmd,
		"wait":  waitCmd,
	}
}

func printCmd(args []string) error {
	if len(args) == 2 {
		fmt.Println(args[1])
	}
	return nil
}

func sleepCmd(args []string) error {
	const funcName = "sleepCmd"

	if len(args) == 2 {
		v := strings.Split(args[1], ".")
		if len(v) > 2 {
			return &ngsiCmdError{funcName, 1, "value error: " + args[1], nil}
		}
		i, err := strconv.Atoi(v[0])
		if err != nil {
			return &ngsiCmdError{funcName, 2, "value error: " + v[0], nil}
		}
		t := time.Second * time.Duration(i)
		if len(v) == 2 && len(v[1]) == 1 {
			i, err = strconv.Atoi(v[1])
			if err != nil {
				return &ngsiCmdError{funcName, 3, "value error: " + v[1], nil}
			}
			t += time.Millisecond * time.Duration(i*100)
		}
		time.Sleep(t)
		return nil
	}

	return &ngsiCmdError{funcName, 4, "param error" + args[1], nil}
}

func haltCmd(args []string) error {
	const funcName = "haltCmd"

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	fmt.Println("halt")

	<-sig

	fmt.Println("resume")

	return nil
}

func waitCmd(args []string) error {
	const funcName = "sleepCmd"

	retry := 60
	if len(args) == 2 {
		if !isHTTP(args[1]) {
			return &ngsiCmdError{funcName, 1, "url error: " + args[1], nil}
		}
		fmt.Printf("Waiting for response from %s\n", args[1])
		for {
			res, err := http.Get(args[1])
			if err != nil {
				retry--
				if retry == 0 {
					return &ngsiCmdError{funcName, 2, "no response from " + args[1], nil}
				}
				time.Sleep(time.Second * time.Duration(1))
				continue
			}
			defer res.Body.Close()
			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				return &ngsiCmdError{funcName, 3, err.Error(), err}
			}
			if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
				return nil
			}
		}
	}
	return &ngsiCmdError{funcName, 4, "param error" + args[1], nil}
}

func ngsiCmd(args []string) error {
	const funcName = "ngsiCmd"

	if *gArgs {
		for i, s := range args {
			fmt.Printf("%d: %s\n", i, s)
		}
	}
	param := []string{}
	if *gNgsiConfig != "" {
		param = append(param, "--config", *gNgsiConfig)
	}
	if *gNgsiCache != "" {
		param = append(param, "--cache", *gNgsiCache)
	}
	param = append(param, args[1:]...)

	cmd := exec.Command(args[0], param...)
	cmd.Stderr = nil
	rc := "0"

	result, err := cmd.Output()

	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			result = e.Stderr
		}
		rc = strconv.Itoa(cmd.ProcessState.ExitCode())
	}
	s := strings.TrimRight(string(result), "\n")
	val["$"] = strings.Split(s, "\n")
	val["?"] = []string{rc}
	return nil
}

func compareCmd(args []string) error {
	const funcName = "compareCmd"

	var err error

	expectedCode := args[0][3:]
	actualCode := val["?"][0]

	if expectedCode != actualCode {
		fmt.Printf("Exit code error, expected:%s, actual:%s\n", expectedCode, actualCode)
		err = &ngsiCmdError{funcName, 1, fmt.Sprintf("Exit code error, expected:%s, actual:%s", expectedCode, actualCode), nil}
	}

	expected := args[1 : len(args)-1]
	if len(expected) == 0 {
		expected = []string{""}
	}
	actual := val["$"]

	if err := diffLines(expected, actual); err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	return err
}

func httpCmd(args []string) error {
	const funcName = "httpCmd"

	if len(args) < 2 {
		return &ngsiCmdError{funcName, 1, "http error", nil}
	}

	url := args[2]
	if !isHTTP(url) {
		return &ngsiCmdError{funcName, 2, "url error: " + url, nil}
	}

	switch args[1] {
	default:
		return &ngsiCmdError{funcName, 3, "http verb error", nil}
	case "get":
		return httpGet(args)
	case "post":
		if len(args) < 4 {
			return &ngsiCmdError{funcName, 4, "http post url --data \"{\"data\":\"post data\"}", nil}
		}
		if args[3] != "--data" {
			return &ngsiCmdError{funcName, 4, "http post url --data \"{\"data\":\"post data\"}", nil}
		}
		return httpPost(args)
	case "delete":
		return httpDelete(args)
	}

}

func httpGet(args []string) error {
	const funcName = "httpGet"

	res, err := http.Get(args[2])
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	status := "0"
	if res.StatusCode != http.StatusOK {
		status = "1"
	}
	s := strings.TrimRight(string(b), "\n")
	val["$"] = strings.Split(s, "\n")
	val["?"] = []string{status}

	return nil
}

func httpPost(args []string) error {
	const funcName = "httpPost"

	req, err := http.NewRequest("POST", args[2], bytes.NewBuffer([]byte(args[4])))
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	status := "0"
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		status = "1"
	}
	s := strings.TrimRight(string(b), "\n")
	val["$"] = strings.Split(s, "\n")
	val["?"] = []string{status}

	return nil
}

func httpDelete(args []string) error {
	const funcName = "httpDelete"

	req, err := http.NewRequest("DELETE", args[2], nil)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &ngsiCmdError{funcName, 3, err.Error(), err}
	}

	status := "0"
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		status = "1"
	}
	s := strings.TrimRight(string(b), "\n")
	val["$"] = strings.Split(s, "\n")
	val["?"] = []string{status}

	return nil
}

func isHTTP(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
