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

package ngsicli

var GlobalFlags = []Flag{
	SyslogFlag,
	StderrFlag,
	ConfigFlag,
	CacheFlag,
	MarginFlag,
	TimeOutFlag,
	MaxCountFlag,
	BatchFlag,
	InsecureSkipVerifyFlag,
	CmdNameFlag,
}

// Global Flags
var (
	SyslogFlag = &StringFlag{
		Name:    "syslog",
		Usage:   "syslog logging `LEVEL` (off, err, info, debug)",
		Choices: []string{"off", "err", "info", "debug"},
	}
	StderrFlag = &StringFlag{
		Name:    "stderr",
		Usage:   "stderr logging `LEVEL` (err, info, debug)",
		Choices: []string{"off", "err", "info", "debug"},
	}
	ConfigFlag = &StringFlag{
		Name:  "config",
		Usage: "configuration `FILE` name",
	}
	CacheFlag = &StringFlag{
		Name:  "cache",
		Usage: "cache `FILE` name",
	}
	HelpFlag = &BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	MarginFlag = &Int64Flag{
		Name:   "margin",
		Usage:  "I/O time out (second)",
		Hidden: true,
		Value:  180,
	}
	TimeOutFlag = &Int64Flag{
		Name:   "timeout",
		Usage:  "I/O time out (second)",
		Hidden: true,
		Value:  60,
	}
	MaxCountFlag = &Int64Flag{
		Name:   "maxCount",
		Usage:  "maxCount",
		Hidden: true,
		Value:  100,
	}
	BatchFlag = &BoolFlag{
		Name:    "batch",
		Aliases: []string{"B"},
		Usage:   "don't use previous args (batch)",
	}
	CmdNameFlag = &StringFlag{
		Name:   "cmdName",
		Hidden: true,
	}
	InsecureSkipVerifyFlag = &BoolFlag{
		Name:  "insecureSkipVerify",
		Usage: "TLS/SSL skip certificate verification",
	}
)

var (
	HostRFlag = &StringFlag{
		Name:         "host",
		Usage:        "broker or server host `VALUE`",
		Aliases:      []string{"h"},
		Required:     true,
		InitClient:   true,
		PreviousArgs: true,
	}
	HostFlag = &StringFlag{
		Name:    "host",
		Usage:   "broker or server host `VALUE`",
		Aliases: []string{"h"},
	}
	TenantFlag = &StringFlag{
		Name:       "service",
		Aliases:    []string{"s"},
		Usage:      "FIWARE Service `VALUE`",
		ValueEmpty: true,
	}
	ScopeFlag = &StringFlag{
		Name:       "path",
		Aliases:    []string{"p"},
		Usage:      "FIWARE ServicePath `VALUE`",
		ValueEmpty: true,
	}
	OAuthTokenFlag = &StringFlag{
		Name:       "oAuthToken",
		Usage:      "OAuth token `VALUE`",
		Hidden:     true,
		ValueEmpty: true,
	}
	XAuthTokenFlag = &BoolFlag{
		Name:   "xAuthToken",
		Usage:  "use X-Auth-Token",
		Hidden: true,
	}
)

var (
	SafeStringFlag = &StringFlag{
		Name:  "safeString",
		Usage: "use safe string (`VALUE`: on/off)",
	}
	NgsiTypeFlag = &StringFlag{
		Name:  "ngsiType",
		Usage: "NGSI type: v2 or ld",
	}
)

var (
	VerboseFlag = &BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Usage:   "verbose",
	}
	JsonFlag = &BoolFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "JSON format",
	}
	PrettyFlag = &BoolFlag{
		Name:    "pretty",
		Aliases: []string{"P"},
		Value:   false,
		Usage:   "pretty format",
	}
	RunFlag = &BoolFlag{
		Name:  "run",
		Usage: "run command",
		Value: false,
	}
)
