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

package ngsicli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

func printVersion(c *Context, token *token) bool {
	arg := token.Next()

	if arg != nil && (*arg == "--version" || *arg == "-v") {
		fmt.Fprintf(c.Ngsi.StdWriter, "ngsi version %s\n", c.App.Version)
		return true
	} else if arg != nil && *arg == "--serial" {
		var x, y, z int
		if _, err := fmt.Sscanf(c.App.Version, "%d.%d.%d", &x, &y, &z); err == nil {
			fmt.Fprintf(c.Ngsi.StdWriter, "%d%02d%02d", x, y, z)
		}
		return true
	}
	_ = token.Prev()

	return false
}

type category map[string][]*Command

func printHelp(c *Context) {
	msg := ""
	msg += fmt.Sprintf("NAME:\n   ngsi - %s\n\n", c.App.Usage)
	msg += fmt.Sprintf("USAGE:\n   %s [global options] command [options] [arguments...]\n\n", c.CommandName)
	msg += fmt.Sprintf("VERSION:\n   ngsi version %s\n\n", c.App.Version)

	// Commands
	msg += commandList(c)

	// Global flags
	msg += commandFlags("GLOBAL OPTIONS", c.GlobalFlags)

	// Previous args
	msg += previousArgs(c)

	fmt.Fprint(c.Ngsi.StdWriter, msg)
}

func printCommandHelp(c *Context) {
	var cmd *Command
	cmdName := ""
	usage := ""

	for _, cc := range c.Commands {
		cmdName += " " + cc.Name
		cmd = cc
		usage = cc.Usage
	}
	cmdName = cmdName[1:]

	msg := ""
	msg += fmt.Sprintf("NAME:\n   ngsi %s - %s\n\n", cmdName, usage)
	if cmd.Subcommands == nil {
		msg += fmt.Sprintf("USAGE:\n   %s [global options] %s [options] [arguments...]\n\n", c.CommandName, cmdName)
	} else {
		msg += fmt.Sprintf("USAGE:\n   %s [global options] %s [options] [command] [arguments...]\n\n", c.CommandName, cmdName)
	}

	// Category
	msg += "CATEGORY:\n"
	msg += "   " + c.Commands[0].Category + "\n\n"

	// Commands
	msg += subCommandList(cmd)

	// Command flags
	msg += commandFlags("OPTIONS", c.Flags)

	// Global flags
	msg += commandFlags("GLOBAL OPTIONS", c.GlobalFlags)

	// Previous args
	msg += previousArgs(c)

	fmt.Fprint(c.Ngsi.StdWriter, msg)
}

func commandList(c *Context) string {
	msg := "COMMANDS:\n"
	msg += "   help, h  Shows a list of commands or help for one command\n"

	category := make(category)
	for _, cmd := range c.App.Commands {
		category[cmd.Category] = append(category[cmd.Category], cmd)
	}
	keys := make(categories, len(category))
	i := 0
	for key := range category {
		keys[i] = key
		i++
	}
	sort.Sort(keys)
	for _, name := range keys {
		cat := category[name]
		msg += fmt.Sprintf("   %s:\n", name)
		max := 0
		for _, cmd := range cat {
			max = maxInt(max, len(cmd.Name))
		}
		for _, cmd := range cat {
			if !cmd.Hidden {
				s := cmd.Name + "          "
				msg += fmt.Sprintf("     %s  %s\n", s[:max], cmd.Usage)
			}
		}
	}

	msg += "\n"

	return msg
}

func subCommandList(cmd *Command) string {
	msg := ""

	if cmd.Subcommands != nil {
		format := ""
		msg = "COMMANDS:\n"
		max := 0
		for _, subCmd := range cmd.Subcommands {
			if subCmd.Subcommands == nil {
				max = maxInt(max, len(subCmd.Name))
			}
		}
		if max < 7 {
			max = 7
		}
		for _, subCmd := range cmd.Subcommands {
			if subCmd.Subcommands == nil && !subCmd.Hidden {
				format = fmt.Sprintf("   %%-%ds  %%s\n", max)
				msg += fmt.Sprintf(format, subCmd.Name, subCmd.Usage)
			}
		}
		msg += fmt.Sprintf(format, "help, h", "Shows a list of commands or help for one command")

		max = 0
		for _, subCmd := range cmd.Subcommands {
			if subCmd.Subcommands != nil {
				max = maxInt(max, len(subCmd.Name))
			}
		}
		if max > 0 {
			category := ""
			for _, subCmd := range cmd.Subcommands {
				if subCmd.Subcommands != nil && !subCmd.Hidden {
					if category != subCmd.Category {
						category = subCmd.Category
						msg += fmt.Sprintf("   %s:\n", category)
					}
					format = fmt.Sprintf("     %%-%ds  %%s\n", max)
					msg += fmt.Sprintf(format, subCmd.Name, subCmd.Usage)
				}
			}

		}

		msg += "\n"
	}

	return msg
}

func commandFlags(title string, flags []Flag) string {
	msg := ""

	if flags != nil {
		msg = title + ":\n"
		max := 0
		for _, flag := range flags {
			if !flag.FlagHidden() {
				max = maxInt(max, len(flag.FlagNameList()))
			}
		}
		for _, flag := range flags {
			if !flag.FlagHidden() {
				format := fmt.Sprintf("   %%-%ds  %%s\n", max)
				msg += fmt.Sprintf(format, flag.FlagNameList(), flag.FlagUsage())
			}
		}
		msg += "\n"
	}

	return msg
}

func previousArgs(c *Context) string {
	msg := ""

	if c.Ngsi != nil && c.Ngsi.PreviousArgs != nil {
		m := previousArgsMaxLen(c.Ngsi.PreviousArgs)
		msg += "PREVIOUS ARGS:\n"
		if !c.Ngsi.PreviousArgs.UsePreviousArgs {
			msg += "   off\n"
			msg += "   (To enable it, run 'ngsi settings previousArgs --on')\n"
		} else {
			if m == 0 {
				msg += "   None\n"
			} else {
				msg += previousArg("Host", c.Ngsi.PreviousArgs.Host, m)
				msg += previousArg("FIWARE-Service", c.Ngsi.PreviousArgs.Tenant, m)
				msg += previousArg("FIWARE-ServicePath", c.Ngsi.PreviousArgs.Scope, m)
				msg += previousArg("Token", c.Ngsi.PreviousArgs.Token, m)
				msg += previousArg("Syslog", c.Ngsi.PreviousArgs.Syslog, m)
				msg += previousArg("Stderr", c.Ngsi.PreviousArgs.Stderr, m)
				msg += previousArg("LogFile", c.Ngsi.PreviousArgs.Logfile, m)
				msg += previousArg("LogLevel", c.Ngsi.PreviousArgs.Loglevel, m)
				msg += "   (To clear args, run 'ngsi settings clear')\n"
			}
		}
	}

	return msg
}

func previousArg(name, msg string, l int) string {
	if msg == "" {
		return ""
	}
	name += "                                                  "
	return fmt.Sprintf("   %s  %s\n", name[:l], msg)
}

func previousArgsMaxLen(p *ngsilib.Settings) int {
	max := 0

	max = maxInt(max, previousArgsLen("Host", p.Host))
	max = maxInt(max, previousArgsLen("FIWARE-Service", p.Tenant))
	max = maxInt(max, previousArgsLen("FIWARE-ServicePath", p.Scope))
	max = maxInt(max, previousArgsLen("Token", p.Token))
	max = maxInt(max, previousArgsLen("Syslog", p.Syslog))
	max = maxInt(max, previousArgsLen("Stderr", p.Stderr))
	max = maxInt(max, previousArgsLen("LogFile", p.Logfile))
	max = maxInt(max, previousArgsLen("LogLevel", p.Loglevel))

	return max
}

func previousArgsLen(name, value string) int {
	if value == "" {
		return 0
	}
	return len(name)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type categories []string

func (c categories) Len() int {
	return len(c)
}

func (c categories) Less(i, j int) bool {
	if c[i] == "MANAGEMENT" {
		return false
	}
	if c[j] == "MANAGEMENT" {
		return true
	}
	return strings.ToUpper(c[i]) < strings.ToUpper(c[j])
}

func (c categories) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
