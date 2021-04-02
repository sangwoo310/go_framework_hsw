package cmd

import "github.com/urfave/cli/v2"

type CmdStruct struct {
	Name			string
	Usage			string
	UsageText		string
	Description		string
	Category		string
	SubCommands		[]*cli.Command
	Flags			[]*cli.Flag
}

var cmd []*cli.Command



func Build()  {

}