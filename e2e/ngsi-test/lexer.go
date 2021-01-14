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
	"fmt"
	"io"
	"strings"
)

type lexer struct {
	reader    *bufio.Reader
	lineNo    int
	line      string
	startLine int
	token     []string
	s         string
}

func (l *lexer) scan() ([]string, error) {
	const funcName = "scan"

	l.line = ""
	l.lineNo++
	l.startLine = l.lineNo
	l.token = []string{}
	l.s = ""

	for {
		c, err := l.readChar()
		if err == io.EOF {
			return nil, err
		}
		if err != nil {
			return nil, &ngsiCmdError{funcName, 1, err.Error(), err}
		}

		switch c {
		case '\n':
			return l.token, nil
		case ' ', '\t':
			l.s = ""
			break
		case '\\':
			c, err := l.readChar()
			if err != nil {
				return nil, &ngsiCmdError{funcName, 2, err.Error(), err}
			}
			if c == '\n' {
				c, err = l.readChar()
				if err != nil {
					return nil, &ngsiCmdError{funcName, 3, err.Error(), err}
				}
				l.unreadChar()
				l.s = ""
				l.lineNo++
				break
			}
			return nil, &ngsiCmdError{funcName, 4, err.Error(), err}
		case '\'':
			err = l.readSingleQuotedWord()
			if err != nil {
				return nil, &ngsiCmdError{funcName, 5, err.Error(), err}
			}
		case '"':
			err = l.readDoubleQuotedWord()
			if err != nil {
				return nil, &ngsiCmdError{funcName, 6, err.Error(), err}
			}
		case '#':
			if len(l.token) == 0 {
				_, err = l.readLine()
				if err != nil {
					return nil, &ngsiCmdError{funcName, 7, err.Error(), err}
				}
				return l.token, nil
			}
			err = l.readWord()
			if err != nil {
				return nil, &ngsiCmdError{funcName, 8, err.Error(), err}
			}
			break
		case '$':
			l.readValiable()
		default:
			err = l.readWord()
			if err != nil {
				return nil, err
			}
			if len(l.token) == 1 {
				if strings.HasPrefix(l.token[0], "```") {
					err = l.readBackquotedWord()
					if err != nil {
						return nil, &ngsiCmdError{funcName, 9, err.Error(), err}
					}
					return l.token, nil
				}
			}
		}
	}
}

func (l *lexer) readChar() (rune, error) {
	const funcName = "readChar"

	c, _, err := l.reader.ReadRune()

	if err == io.EOF {
		return c, err
	}
	if err != nil {
		return c, &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	s := string(c)
	l.s += s
	l.line += s

	return c, err
}

func (l *lexer) unreadChar() error {
	l.s = l.s[:len(l.s)-1]
	l.line = l.line[:len(l.line)-1]
	return l.reader.UnreadRune()
}

func (l *lexer) readWord() error {
	const funcName = "readWord"

	for {
		c, err := l.readChar()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if c == ' ' || c == '\t' || c == '\n' {
			l.unreadChar()
			break
		}
	}
	l.token = append(l.token, l.s)
	l.s = ""

	return nil
}

func (l *lexer) readLine() (string, error) {
	const funcName = "readLine"

	s := l.s
	l.s = ""

	for {
		c, err := l.readChar()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if c == '\n' {
			l.s = l.s[:len(l.s)-1]
			break
		}
		s += string(c)
	}

	l.token = append(l.token, s)
	l.s = ""

	return s, nil
}

func (l *lexer) readValiable() error {
	const funcName = "readSingleQuotedWord"

	save := l.s[:len(l.s)-1]

	c, err := l.readChar()
	if err != nil {
		return &ngsiCmdError{funcName, 1, err.Error(), err}
	}
	if c != '{' {
		return &ngsiCmdError{funcName, 2, "{ not found", nil}
	}
	n := ""
	for {
		c, err = l.readChar()
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}
		if c == '}' {
			break
		}
		n += string(c)
	}

	if _, ok := val[n]; ok {
		l.s = save + val[n][0]
	} else {
		return &ngsiCmdError{funcName, 4, "${" + n + "} not found", nil}
	}
	return nil
}

func (l *lexer) readSingleQuotedWord() error {
	const funcName = "readSingleQuotedWord"

	escaped := false
	l.s = l.s[:len(l.s)-1]

	for {
		c, err := l.readChar()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if escaped {
			escaped = false
		} else {
			if c == '\\' {
				escaped = true
				l.s = l.s[:len(l.s)-1]
			} else if c == '\'' {
				l.s = l.s[:len(l.s)-1]
				break
			} else if c == '$' {
				err = l.readValiable()
				if err != nil {
					return &ngsiCmdError{funcName, 2, err.Error(), err}
				}
			}
		}
	}
	l.token = append(l.token, l.s)
	l.s = ""

	return nil
}

func (l *lexer) readDoubleQuotedWord() error {
	const funcName = "readDoubleQuotedWord"

	escaped := false
	l.s = l.s[:len(l.s)-1]

	for {
		c, err := l.readChar()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if escaped {
			escaped = false
		} else {
			if c == '\\' {
				escaped = true
				l.s = l.s[:len(l.s)-1]
			} else if c == '"' {
				l.s = l.s[:len(l.s)-1]
				break
			}
		}
	}
	l.token = append(l.token, l.s)
	l.s = ""

	return nil
}

func (l *lexer) readBackquotedWord() error {
	const funcName = "readBackquotedWord"

	for {
		c, err := l.readChar()
		if err != nil {
			return &ngsiCmdError{funcName, 1, err.Error(), err}
		}
		if c == ' ' || c == '\t' {
			continue
		} else if c == '\n' {
			l.s = ""
			break
		} else {
			return &ngsiCmdError{funcName, 2, fmt.Sprintf("Illegal char: %s", string(c)), nil}
		}
	}
	for {
		s, err := l.readLine()
		if err != nil {
			return &ngsiCmdError{funcName, 3, err.Error(), err}
		}
		l.lineNo++
		if s == "```" {
			return nil
		}
		l.line = ""
	}
}
