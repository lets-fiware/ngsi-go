/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

// https://fiware-orion.readthedocs.io/en/master/user/forbidden_characters/index.html

var forbiddenCharsEncode = map[rune]string{
	'%':  "%25",
	'"':  "%22",
	'\'': "%27",
	'(':  "%28",
	')':  "%29",
	';':  "%3B",
	'<':  "%3C",
	'=':  "%3D",
	'>':  "%3E",
}

var forbiddenCharsDecode = map[string]string{
	"%22": `\"`,
	"%25": `%`,
	"%27": "'",
	"%28": "(",
	"%29": ")",
	"%3b": ";",
	"%3B": ";",
	"%3c": "<",
	"%3C": "<",
	"%3d": "=",
	"%3D": "=",
	"%3e": ">",
	"%3E": ">",
}

// SafeStringEncode is ...
func SafeStringEncode(s string) string {
	ss := ""
	for _, c := range s {
		if cc, ok := forbiddenCharsEncode[c]; ok {
			ss += cc
		} else {
			ss += string(c)
		}
	}
	return ss
}

// SafeStringDecode is ...
func SafeStringDecode(s string) string {
	ss := ""
	escape := 0
	escapeStr := ""

	for _, c := range s {
		switch escape {
		case 0:
			if c == '%' {
				escapeStr = "%"
				escape = 1
			} else {
				ss += string(c)
			}
		case 1:
			escapeStr += string(c)
			escape = 2
		case 2:
			escapeStr += string(c)
			if cc, ok := forbiddenCharsDecode[escapeStr]; ok {
				ss += cc
			} else {
				ss += escapeStr
			}
			escape = 0
		}
	}
	if escape > 0 {
		ss += escapeStr
	}
	return ss
}

// JSONSafeStringEncode is ...
func JSONSafeStringEncode(data []byte) ([]byte, error) {
	const funcName = "JSONSafeStringEncode"

	bytes, err := jsonParser(data, SafeStringEncode)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}
	return bytes, nil
}

// JSONSafeStringDecode is ...
func JSONSafeStringDecode(data []byte) ([]byte, error) {
	const funcName = "JSONSafeStringDecode"

	bytes, err := jsonParser(data, SafeStringDecode)
	if err != nil {
		return nil, ngsierr.New(funcName, 1, err.Error(), err)
	}
	return bytes, nil
}

const (
	tokenNone          = iota
	tokenBracesOpen    // {
	tokenBracesClose   // }
	tokenBracketsOpen  // [
	tokenBracketsClose // ]
	tokenJSONKey
	tokenJSONValue
)

func jsonParser(jsonStream []byte, f func(string) string) ([]byte, error) {
	const funcName = "jsonParser"

	if !IsJSON(jsonStream) {
		s := f(string(jsonStream))
		return []byte(s), nil
	}

	var err error
	var stack [128]int
	tokenTable := map[byte]int{
		'{': tokenBracesOpen,
		'}': tokenBracesClose,
		'[': tokenBracketsOpen,
		']': tokenBracketsClose,
	}

	p := -1
	prevToken := tokenNone
	mode := tokenNone

	dst := new(bytes.Buffer)
	dec := json.NewDecoder(bytes.NewReader(jsonStream))
	for {
		var t json.Token
		t, err = dec.Token()
		if err == io.EOF {
			if p == -1 && (prevToken == tokenBracesClose || prevToken == tokenBracketsClose) {
				err = nil
				break
			}
			s := dst.String()
			l := len(s)
			if l > 15 {
				l = 15
			}
			return nil, ngsierr.New(funcName, 1, "json error: "+s[len(s)-l:], err)
		}
		if err != nil {
			break
		}
		switch tt := t.(type) {
		case json.Delim:
			c := byte(tt)
			switch t {
			case json.Delim('{'), json.Delim('['):
				p++
				stack[p] = mode
				switch prevToken {
				case tokenJSONKey:
					dst.WriteByte(':')
				case tokenJSONValue, tokenBracesClose, tokenBracketsClose:
					dst.WriteByte(',')
				}
				dst.WriteByte(c)
				prevToken = tokenTable[c]
				mode = prevToken
			case json.Delim('}'), json.Delim(']'):
				dst.WriteByte(c)
				prevToken = tokenTable[c]
				mode = stack[p]
				p--
			}
		case string:
			s := `"` + f(t.(string)) + `"`
			switch mode {
			case tokenBracketsOpen: // [
				if prevToken != tokenBracketsOpen {
					dst.WriteByte(',')
				}
				dst.WriteString(s)
				prevToken = tokenJSONValue
			case tokenBracesOpen: // {
				switch prevToken {
				case tokenBracesOpen:
					dst.WriteString(s)
					prevToken = tokenJSONKey
				case tokenJSONKey:
					dst.WriteByte(':')
					dst.WriteString(s)
					prevToken = tokenJSONValue
				case tokenJSONValue, tokenBracesClose, tokenBracketsClose:
					dst.WriteByte(',')
					dst.WriteString(s)
					prevToken = tokenJSONKey
				}
			}
		case float64:
			s := strconv.FormatFloat(t.(float64), 'f', -1, 64)
			writeTokenValue(dst, mode, s, &prevToken)
		case bool:
			s := strconv.FormatBool(t.(bool))
			writeTokenValue(dst, mode, s, &prevToken)
		case nil:
			s := "null"
			writeTokenValue(dst, mode, s, &prevToken)
		}
	}
	if err != nil {
		if err, ok := err.(*json.SyntaxError); ok {
			s := err.Offset - 15
			if s < 0 {
				s = 0
			}
			e := err.Offset + 15
			if e > int64(len(jsonStream)) {
				e = int64(len(jsonStream))
			}

			return nil, ngsierr.New(funcName, 2, fmt.Sprintf("%s (%d) %s", err.Error(), err.Offset, string(jsonStream[s:e])), err)
		}
		return nil, ngsierr.New(funcName, 3, err.Error(), err)
	}
	return dst.Bytes(), nil
}

func writeTokenValue(dst *bytes.Buffer, mode int, s string, prevToken *int) {
	switch mode {
	case tokenBracketsOpen: // [
		if *prevToken != tokenBracketsOpen {
			dst.WriteByte(',')
		}
		dst.WriteString(s)
		*prevToken = tokenJSONValue
	case tokenBracesOpen: // {
		dst.WriteByte(':')
		dst.WriteString(s)
		*prevToken = tokenJSONValue
	}
}
