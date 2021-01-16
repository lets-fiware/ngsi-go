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
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	gConfig     = flag.String("config", "./.ngsi-test-config.json", "config file for ngsi-test")
	gNgsiConfig = flag.String("ngsi-config", "", "config file for NGSI Go")
	gNgsiCache  = flag.String("ngsi-cache", "", "cahce file for NGSI Go")
	gVerbose    = flag.Bool("verbose", false, "verobose")
	gArgs       = flag.Bool("args", false, "")
)

var val map[string][]string

var rComment = regexp.MustCompile(`^# *[0-9]{1,5}`)

func ngsiTest() int {
	flag.Parse()

	code := 0
	if len(flag.Args()) == 1 {
		if err := runTestCases(flag.Args()[0]); err != nil {
			printError(err)
			code = 1
		}
	} else {
		fmt.Printf("testcase file not found\n")
		code = 1
	}
	return code
}

type configDef struct {
	Valiables map[string]string `json:"valiables"`
}

func readConfigFile() error {
	const funcName = "readConfigFile"

	val = make(map[string][]string)

	fileName := *gConfig

	_, err := os.Stat(fileName)
	if err != nil {
		home, err := os.UserHomeDir()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}

		fileName = filepath.Join(home, ".ngsi-test-config.json")
	}

	fmt.Printf("config: %s\n", fileName)

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	var config configDef

	err = json.Unmarshal(b, &config)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	for k, v := range config.Valiables {
		val[k] = []string{v}
	}

	if *gVerbose {
		for k, v := range val {
			fmt.Printf("${%s}: %s\n", k, v[0])
		}
	}
	return nil
}

func runTestCases(fileName string) error {
	const funcName = "runTestCases"

	initCmdTable()

	err := readConfigFile()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}

	fInfo, err := os.Stat(fileName)
	if err != nil {
		return &ngsiCmdError{funcName, 2, err.Error(), err}
	}

	if fInfo.IsDir() {
		dirs := 0
		files := 0
		cases := 0
		err := runTestCaseDir(fileName, &dirs, &files, &cases)
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}
		fmt.Printf("Results:\n   %d cases, %d files, %d directories\n", cases, files, dirs)
	} else {
		cases, err := runTestCaseFile(fileName)
		if err != nil {
			return &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		fmt.Printf("   cases: %d\n", cases)
	}

	return nil
}

func runTestCaseDir(dirName string, dirs, files, cases *int) error {
	const funcName = "runTestCaseFile"

	fmt.Println("directory: " + dirName)
	*dirs++

	dirList := []string{}
	fileList := []string{}

	ls, err := ioutil.ReadDir(dirName)
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	for _, f := range ls {
		name := f.Name()
		if '0' > name[0] || name[0] > '9' {
			continue
		}
		if f.IsDir() {
			dirList = append(dirList, filepath.Join(dirName, name))
		} else {
			fileList = append(fileList, filepath.Join(dirName, name))
		}
	}
	sort.Strings(dirList)
	sort.Strings(fileList)

	for _, f := range dirList {
		err := runTestCaseDir(f, dirs, files, cases)
		if err != nil {
			return &ngsiCmdError{funcName, 2, err.Error(), err}
		}
	}
	for _, f := range fileList {
		c, err := runTestCaseFile(f)
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}
		*files++
		*cases += c
	}
	return nil
}

func runTestCaseFile(fileName string) (cases int, err error) {
	const funcName = "runTestCaseFile"

	cases = 0

	fmt.Println("file: " + fileName)

	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		return cases, &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	defer func() { setNewError(funcName, 2, f.Close(), &err) }()

	l := lexer{reader: bufio.NewReader(f)}

	for {
		token, err := l.scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			return cases, &ngsiCmdError{funcName, 3, err.Error(), err}
		}

		c, err := parser(l.lineNo, token)
		if err != nil {
			return cases, &ngsiCmdError{funcName, 4, err.Error(), err}
		}
		cases += c
	}

	return cases, nil
}

func parser(line int, token []string) (int, error) {
	const funcName = "execCmd"

	cases := 0

	if len(token) == 0 {
		return cases, nil
	}

	t := token[0]
	if strings.HasPrefix(t, "#") {
		if rComment.Match([]byte(t)) {
			fmt.Println(t)
			cases++
		}
		return cases, nil
	}
	if strings.HasPrefix(t, "$") {
		v := strings.Split(t, "=")
		if len(v) == 2 {
			if strings.HasPrefix(v[1], "$") {
				s, ok := val[v[1][1:]]
				if ok {
					val[v[0][1:]] = make([]string, len(s))
					copy(val[v[0][1:]], val[v[1][1:]])
					return cases, nil
				}
				return cases, &ngsiCmdError{funcName, 1, fmt.Sprintf("%s not found: L%04d %s", s, line, strings.Join(token, " ")), nil}
			}
			s := []string{v[1]}
			val[v[0][1:]] = make([]string, len(s))
			copy(val[v[0][1:]], s)
			return cases, nil
		}
		return cases, &ngsiCmdError{funcName, 2, fmt.Sprintf("%04d %s", line, strings.Join(token, " ")), nil}
	}

	for i, s := range token {
		if strings.HasPrefix(s, "$") {
			if v, ok := val[s[1:]]; ok {
				token[i] = strings.Join(v, "\n")
			} else {
				return cases, &ngsiCmdError{funcName, 3, fmt.Sprintf("%s not found: L%04d %s", s, line, strings.Join(token, " ")), nil}
			}
		}
	}

	if t == "break" {
		return cases, &ngsiCmdError{funcName, 4, "stopped running test cases by break command.", nil}
	}

	if strings.HasPrefix(t, "```") {
		t = "```"
	}

	if f, ok := cmdTable[t]; ok {
		err := f(line, token)
		if err != nil {
			return cases, &ngsiCmdError{funcName, 5, err.Error(), err}
		}
		return cases, nil
	}
	return cases, &ngsiCmdError{funcName, 6, fmt.Sprintf("%s not found: L%04d %s", t, line, strings.Join(token, " ")), nil}
}
