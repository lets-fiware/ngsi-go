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

package ngsicli

import (
	"fmt"
	"strings"
	"time"

	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
)

var (
	Version  string
	Revision string
)

var CmdMode = ""

type App struct {
	Name        string
	HelpName    string
	Usage       string
	UsageText   string
	ArgsUsage   string
	Version     string
	Description string
	Commands    []*Command
	Flags       []Flag
	Compiled    time.Time
	Copyright   string
}

type Command struct {
	Name          string
	Aliases       []string
	Usage         string
	UsageText     string
	Description   string
	ArgsUsage     string
	Category      string
	Subcommands   []*Command
	Flags         []Flag
	RequiredFlags []string
	OptionFlags   *ValidationFlag
	ServerList    []string
	Hidden        bool

	Action ActionFunc
}

var Copyright string

const bashCompletion = "--generate-bash-completion"

type ActionFunc func(*Context, *ngsilib.NGSI, *ngsilib.Client) error

func (r *App) Run(args []string) error {
	const funcName = "Run"

	command, c, err := ngsiRun(r, args)
	if err != nil {
		if len(args) > 0 && args[len(args)-1] == bashCompletion {
			return nil
		}
		return ngsierr.New(funcName, 1, err.Error(), err)
	}
	if command == nil {
		return nil
	}
	if command.Subcommands != nil {
		printCommandHelp(c)
		return nil
	}

	err = command.Action(c, c.Ngsi, c.Client)
	if c.Ngsi.Updated && c.Ngsi.GetPreviousArgs().UsePreviousArgs {
		e := c.Ngsi.SavePreviousArgs()
		if e != nil {
			fmt.Fprint(c.Ngsi.Stderr, ngsierr.SprintMsg(funcName, 2, e.Error()))
		}
	}
	return err
}

func (r *App) Parse(args []string) (*Command, *Context, error) {
	const funcName = "Parse"

	cmd, ctx, err := ngsiRun(r, args)
	if err != nil {
		err = ngsierr.New(funcName, 1, err.Error(), err)
	}

	return cmd, ctx, err
}

func ngsiRun(r *App, args []string) (*Command, *Context, error) {
	const funcName = "ngsiRun"

	c := NewContext(r)

	token := newToken(args)

	cmdName := token.Next()
	if cmdName == nil {
		return nil, nil, ngsierr.New(funcName, 1, "command name error", nil)
	}
	c.CommandName = *cmdName

	if s := token.Peek(); s != nil && *s == bashCompletion {
		fmt.Fprint(c.Ngsi.StdWriter, hintCmdList(r.Commands))
		return nil, nil, nil
	}

	if printVersion(c, token) {
		return nil, nil, nil
	}

	err := parseGlobalFlag(c, token)
	if err != nil {
		return nil, nil, ngsierr.New(funcName, 2, err.Error(), err)
	}
	if c.Bashcompletion {
		return nil, nil, nil
	}

	ngsi, err := InitCmd(c)
	if err != nil {
		return nil, nil, ngsierr.New(funcName, 3, err.Error(), err)
	}
	c.Ngsi = ngsi

	c.Flags = []Flag{}

	cmd := token.Next()
	if cmd == nil || *cmd == "help" || *cmd == "h" { // no commnand, help, h
		printHelp(c)
		return nil, nil, nil
	}

	if *cmd == bashCompletion {
		fmt.Fprint(c.Ngsi.StdWriter, hintCmdList(r.Commands))
		return nil, nil, nil
	}

	c.GlobalFlags = removeFlag(c.GlobalFlags, "help")
	c.GlobalFlags = removeFlag(c.GlobalFlags, "version")

	command, err := runCmd(c, token, *cmd)
	if err != nil {
		return nil, nil, ngsierr.New(funcName, 4, err.Error(), err)
	}

	if command != nil {
		c.ServerList = command.ServerList
		err = createNewClient(c)
		if err != nil {
			return nil, nil, ngsierr.New(funcName, 5, err.Error(), err)
		}
		return command, c, nil
	}

	return nil, nil, nil
}

func runCmd(c *Context, token *token, cmdName string) (*Command, error) {
	const funcName = "parseCmd"

	for _, cmd := range c.App.Commands {
		if cmd.Name == cmdName {
			cmd, err := parseCmdFlag(c, token, cmd)
			if err != nil {
				return nil, ngsierr.New(funcName, 1, err.Error(), err)
			}
			if c.HelpCommand {
				printCommandHelp(c)
				return nil, nil
			} else if c.Bashcompletion {
				return nil, nil
			}
			return cmd, nil
		}
	}

	return nil, ngsierr.New(funcName, 2, cmdName+" not found", nil)
}

func createNewClient(c *Context) error {
	const funcName = "newClient"

	ngsi := c.Ngsi

	if f := c.GetStringFlag("host"); f != nil {
		pArgs := ngsi.GetPreviousArgs()
		ngsi.Host = f.Value

		if ngsi.Host == "" {
			ngsi.Host = pArgs.Host
		}

		if pArgs.Host != ngsi.Host {
			pArgs.Host = ngsi.Host
			pArgs.Tenant = ""
			pArgs.Scope = ""
			ngsi.Updated = true
		}

		if f.IsInitClient() {
			var err error
			c.Client, err = newClient(ngsi, c, false, c.ServerList, f.SkipGetToken, false)
			if err != nil {
				return ngsierr.New(funcName, 1, err.Error(), err)
			}
		}
	}

	if c.HasFlag("host2") && c.IsSet("host2") {
		ngsi.Destination = c.String("host2")

		var err error
		c.Client2, err = newClient(ngsi, c, false, c.ServerList, false, true)
		if err != nil {
			return ngsierr.New(funcName, 2, err.Error(), err)
		}
	}
	return nil
}

func parseCmdFlag(c *Context, token *token, cmd *Command) (*Command, error) {
	const funcName = "parseCmdFlag"

	var flags []Flag

	for _, f := range cmd.Flags {
		flags = append(flags, f.Copy(true))
	}
	helpFlag := &BoolFlag{Name: "help", Usage: "show help"}
	flags = append(flags, helpFlag)
	c.Flags = flags
	c.RequiredFlags = append(c.RequiredFlags, cmd.RequiredFlags...)

	c.Commands = append(c.Commands, cmd)

	arg := token.Peek()
	if (arg == nil || *arg == "help" || *arg == "h") && cmd.Subcommands != nil {
		c.HelpCommand = true
		return nil, nil
	}

	for {
		arg := token.Next()
		if arg == nil {
			break
		}
		if *arg == bashCompletion {
			fmt.Fprint(c.Ngsi.StdWriter, hintCmdList(cmd.Subcommands))
			c.Bashcompletion = true
			return nil, nil
		}
		if s := token.Peek(); s != nil && *s == bashCompletion && strings.HasPrefix(*arg, "-") {
			fmt.Fprint(c.Ngsi.StdWriter, hintFlagList(c.Flags, *arg))
			c.Bashcompletion = true
			return nil, nil
		}
		name, alias, opt := isOption(*arg)
		if opt {
			err := parseOpt(flags, token, name, alias)
			if err != nil {
				return nil, ngsierr.New(funcName, 1, err.Error(), err)
			}
			if helpFlag.IsSet() {
				c.HelpCommand = true
				return nil, nil
			}
		} else {
			if cmd.Subcommands != nil {
				cmd = searchSubCommnad(cmd.Subcommands, name)
				if cmd == nil {
					return nil, ngsierr.New(funcName, 2, name+" not found", nil)
				}
				c.Commands = append(c.Commands, cmd)
				var cmdFlags []Flag
				for _, f := range cmd.Flags {
					cmdFlags = append(cmdFlags, f.Copy(true))
				}
				flags = removeFlag(flags, "help")
				flags = append(flags, cmdFlags...)
				flags = append(flags, helpFlag)
				c.Flags = flags
				c.RequiredFlags = append(c.RequiredFlags, cmd.RequiredFlags...)
			} else {
				*c.Arg = append(*c.Arg, name)
			}
		}
	}

	if required(flags, c) {
		return nil, ngsierr.New(funcName, 3, "missing required options", nil)
	}

	if checkEmpty(flags, c) {
		return nil, ngsierr.New(funcName, 4, "option values are empty", nil)
	}

	err := validation(cmd.OptionFlags, c)
	if err != nil {
		return nil, ngsierr.New(funcName, 5, err.Error(), err)
	}

	if checkChoices(flags, c) {
		return nil, ngsierr.New(funcName, 6, "option values error", nil)
	}

	return cmd, nil
}

func required(flags []Flag, c *Context) bool {
	const funcName = "required"

	setPrevArgs(c)

	missing := false

	if len(c.RequiredFlags) > 0 {
		if c.IsSetAND(c.RequiredFlags) && c.Args().Len() == 0 {
			missing = false
		} else if !c.IsSetAND(c.RequiredFlags) && c.Args().Len() == len(c.RequiredFlags) {
			for i, name := range c.RequiredFlags {
				flag := c.GetFlag(name)
				if flag != nil {
					err := flag.SetValue(c.Args().Get(i))
					if err != nil {
						fmt.Fprintf(c.Ngsi.Stderr, "--%s: %s\n", flag.FlagName(), err.Error())
						missing = true
					}
				} else {
					fmt.Fprint(c.Ngsi.Stderr, ngsierr.SprintMsg(funcName, 1, fmt.Sprintf("--%s: not found\n", name)))
					missing = true
				}
			}
			c.Arg = &cmdArgs{}
		} else {
			missing = true
		}
	}

	for _, flag := range flags {
		if flag.IsRequired() && !flag.IsSet() {
			fmt.Fprint(c.Ngsi.Stderr, ngsierr.SprintMsg(funcName, 2, fmt.Sprintf("--%s not found\n", flag.FlagName())))
			missing = true
		}
		if flag.FlagName() == "data" && flag.IsSet() {
			data := c.GetFlag("data").(*StringFlag)
			if data.Value == "" {
				if !data.ValueEmpty {
					fmt.Fprintln(c.Ngsi.Stderr, ngsierr.SprintMsg(funcName, 3, "data is empty"))
					missing = true
				}
			}
		}
	}

	return missing
}

func checkEmpty(flags []Flag, c *Context) bool {
	const funcName = "checkEmpty"

	err := false

	for _, flag := range flags {
		if f, ok := flag.(*StringFlag); ok {
			if f.IsSet() && f.Value == "" && !f.AllowEmpty() {
				fmt.Fprint(c.Ngsi.Stderr, ngsierr.SprintMsg(funcName, 1, fmt.Sprintf("--%s: value is empty\n", f.FlagName())))
				err = true
			}
		}
	}

	return err
}

func checkChoices(flags []Flag, c *Context) bool {
	const funcName = "checkChoices"

	err := false

	for _, f := range flags {
		if ff, ok := f.(*StringFlag); ok {
			if ff.IsSet() && ff.Choices != nil {
				found := false
				v := strings.ToLower(ff.Value)
				for _, c := range ff.Choices {
					if v == c {
						found = true
						break
					}
				}
				if !found {
					l := len(ff.Choices)
					msg := ""
					if l > 2 {
						msg = fmt.Sprintf("specify one of %s and %s to --%s\n", strings.Join(ff.Choices[0:l-1], ", "), ff.Choices[l-1], ff.FlagName())
					} else {
						msg = fmt.Sprintf("specify either %s or %s to --%s\n", ff.Choices[0], ff.Choices[1], ff.FlagName())
					}
					fmt.Fprint(c.Ngsi.Stderr, ngsierr.SprintMsg(funcName, 1, msg))
					err = true
				}
			}
		}
	}

	return err
}

func setPrevArgs(c *Context) {
	if !c.Ngsi.GetPreviousArgs().UsePreviousArgs {
		return
	}

	host := c.GetStringFlag("host")
	if host == nil {
		return
	}

	h, err := c.Ngsi.GetServerInfo(host.Value, host.SkipRefHost)
	if err == nil {
		f := c.GetStringFlag("service")
		if h.Tenant != "" && f != nil && !f.IsSet() {
			_ = f.SetValue(h.Tenant)
		}
		f = c.GetStringFlag("path")
		if h.Scope != "" && f != nil && !f.IsSet() {
			_ = f.SetValue(h.Scope)
		}
	}

	prevArgs := c.Ngsi.GetPreviousArgs()

	if host.PreviousArgs {
		if host.IsSet() && host.Value != prevArgs.Host {
			copyFlagsToPrevArgs(c)
		} else {
			copyPrevArgsToFlags(c)
			copyFlagsToPrevArgs(c)
		}
	} else {
		if host.IsSet() {
			copyFlagsToPrevArgs(c)
		}
	}
}

type prevFlags []struct {
	name  string
	value *string
}

func getPrevFlags(c *Context) *prevFlags {
	prevArgs := c.Ngsi.GetPreviousArgs()

	flags := prevFlags{
		{"host", &prevArgs.Host},
		{"service", &prevArgs.Tenant},
		{"path", &prevArgs.Scope},
		{"oAuthToken", &prevArgs.Token},
	}
	return &flags
}

func copyPrevArgsToFlags(c *Context) {
	flags := getPrevFlags(c)

	for _, flag := range *flags {
		f := c.GetStringFlag(flag.name)
		if f != nil && !c.IsSet(flag.name) {
			if !(flag.name == "host" && *flag.value == "") {
				_ = f.SetValue(*flag.value)
			}
		}
	}
}

func copyFlagsToPrevArgs(c *Context) {
	flags := getPrevFlags(c)

	for _, flag := range *flags {
		f := c.GetStringFlag(flag.name)
		if f != nil && c.IsSet(flag.name) {
			*flag.value = c.String(flag.name)
		} else {
			*flag.value = ""
		}
	}
	c.Ngsi.Updated = true
}

func searchSubCommnad(cmds []*Command, cmdName string) *Command {
	for _, cmd := range cmds {
		if cmd.Name == cmdName {
			return cmd
		}
	}
	return nil
}

func parseGlobalFlag(c *Context, token *token) error {
	const funcName = "parseGlobalFlag"

	r := c.App

	for _, f := range r.Flags {
		c.GlobalFlags = append(c.GlobalFlags, f.Copy(true))
	}

	c.GlobalFlags = append(c.GlobalFlags, &BoolFlag{Name: "help", Usage: "show help"})
	c.GlobalFlags = append(c.GlobalFlags, &BoolFlag{Name: "version", Aliases: []string{"v"}, Usage: "print the version"})

	for {
		arg := token.Next()
		if arg == nil {
			break
		}

		if *arg == bashCompletion {
			_ = token.Prev()
			break
		}

		if s := token.Peek(); s != nil && *s == bashCompletion && strings.HasPrefix(*arg, "-") {
			fmt.Fprint(c.Ngsi.StdWriter, hintFlagList(c.GlobalFlags, *arg))
			c.Bashcompletion = true
			return nil
		}

		name, alias, opt := isOption(*arg)
		if opt {
			err := parseOpt(c.GlobalFlags, token, name, alias)
			if err != nil {
				return ngsierr.New(funcName, 1, err.Error(), err)
			}
		} else {
			_ = token.Prev()
			break
		}
	}
	c.Flags = c.GlobalFlags

	return nil
}

func isOption(arg string) (string, string, bool) {
	if (strings.HasPrefix(arg, "--") && len(arg) > 2) ||
		(strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 1) {
		name := ""
		alias := ""
		if strings.HasPrefix(arg, "--") {
			name = (arg)[2:]
		} else {
			alias = string((arg)[1:])
		}
		return name, alias, true
	}

	return arg, "", false
}

func parseOpt(flags []Flag, token *token, name, alias string) error {
	const funcName = "parseOpt"

	for _, ff := range flags {
		switch f := ff.(type) {
		case *StringFlag:
			if f.Check(name, alias) {
				v := token.Next()
				if v == nil {
					return ngsierr.New(funcName, 1, "value missing", nil)
				}
				return f.SetValue(*v)
			}
		case *Int64Flag:
			if f.Check(name, alias) {
				v := token.Next()
				if v == nil {
					return ngsierr.New(funcName, 2, "value missing", nil)
				}
				err := f.SetValue(*v)
				if err != nil {
					return ngsierr.New(funcName, 3, err.Error(), err)
				}
				return nil
			}
		case *BoolFlag:
			if f.Check(name, alias) {
				if token.Peek() == nil {
					f.Value = true
				} else {
					v := token.Next()
					err := f.SetValue(*v)
					if err != nil {
						f.Value = true
						_ = token.Prev()
					}
				}
				f.Set = true
				return nil
			}
		}
	}

	if name != "" {
		name = "--" + name
	}
	if alias != "" {
		alias = "-" + alias
	}

	return ngsierr.New(funcName, 4, "unknown flag: "+name+alias, nil)
}

func hintCmdList(cmds []*Command) string {
	hint := ""
	for _, cmd := range cmds {
		hint += cmd.Name + "\n"
	}
	return hint
}

func hintFlagList(flags []Flag, arg string) string {
	hint := ""
	for _, f := range flags {
		if f.IsSet() || f.FlagHidden() {
			continue
		}
		name := "--" + f.FlagName()
		if name == arg {
			return ""
		}
		if strings.HasPrefix(name, arg) {
			hint += name + "\n"
		}
	}
	if !strings.HasPrefix(arg, "--") {
		for _, f := range flags {
			if f.IsSet() || f.FlagHidden() {
				continue
			}
			for _, f := range f.FlagAliases() {
				name := "-" + f
				if name == arg {
					return ""
				}
				if strings.HasPrefix(name, arg) {
					hint += name + "\n"
				}
			}
		}
	}
	return hint
}
