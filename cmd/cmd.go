package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type CmdInterface interface {
	SetCommand() error

	Build() error

}

type CmdStruct struct {
	Name			string
	Usage			string
	UsageText		string
	Description		string
	Category		string
	SubCommands		[]*cli.Command
	Flags			[]*cli.Flag
}

//var cmd []*cli.Command

type App2 struct {
	app2 *cli.App
}

func NewApp() *CmdStruct {
	return &CmdStruct{}
}

func (c CmdStruct)SetCommand(cmd *CmdInterface) error {
	//cmd :=
	return nil
}

func (c CmdStruct)Build() error {
	fmt.Print("this is build function")

	//app := &cli.App{
	//	Commands: ,
	//}
	//
	//if err := app.Run(os.Args); err != nil {
	//	return err
	//}


	return nil
}