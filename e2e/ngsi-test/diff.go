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
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func diffLines(expected, actual []string) error {
	const funcName = "diffLines"

	r := regexp.MustCompile(`(.*)REGEX\((.*)\)(.*)`)

	if len(expected) != len(actual) {
		printDiff(expected, actual)
		return &ngsiCmdError{funcName, 1, fmt.Sprintf("Number of lines error: expected %d, actual %d", len(expected), len(actual)), nil}
	}

	diff := false
	for i, s := range expected {
		result := r.FindAllStringSubmatch(s, -1)
		if result != nil {
			rs := strings.Join(result[0][1:], "")
			r2 := regexp.MustCompile(rs)
			if r2.Match([]byte(actual[i])) == false {
				fmt.Printf("diff error line %d\nExpected:\n%s\nActual:\n%s\n", i, s, actual[i])
				diff = true
			}
		} else {
			if s != actual[i] {
				fmt.Printf("diff error line %d\nExpected:\n%s\nActual:\n%s\n", i, s, actual[i])
				diff = true
			}
		}
	}

	if diff {
		printDiff(expected, actual)
		return &ngsiCmdError{funcName, 2, "diff error", nil}
	}

	return nil
}

func printDiff(expected []string, actual []string) {
	fmt.Println("--------------------------------------- Expected ---------------------------------------")
	for _, s := range expected {
		fmt.Println(s)
	}
	fmt.Println("---------------------------------------- Actual ----------------------------------------")
	for _, s := range actual {
		fmt.Println(s)
	}
	fmt.Println("----------------------------------------------------------------------------------------")
}

func printJSONIndent(s string, b []byte) {
	if s != "" {
		fmt.Println(s)
	}
	buf := new(bytes.Buffer)
	if err := json.Indent(buf, b, "", "  "); err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println(buf.String())
}
