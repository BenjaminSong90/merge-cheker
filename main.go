package main

import (
	"fmt"
	"git_tools/commands"
	"os"

	"github.com/jessevdk/go-flags"
)

type Option struct {
	Value       bool                        `short:"v"`
	Init        commands.InitCommand        `command:"init"`
	Branch      commands.BranchCommand      `command:"branch"`
	Verify      commands.VerifyCommand      `command:"verify"`
	BranchGroup commands.BranchGroupCommand `command:"group"`
}

func main() {
	var opt Option
	parser := flags.NewParser(&opt, flags.HelpFlag)
	var err error
	if len(os.Args) == 1 {
		_, err = parser.ParseArgs([]string{"--help"})
	} else {
		_, err = parser.Parse()
	}

	if err != nil {
		fmt.Println(err.Error())
	}

}
