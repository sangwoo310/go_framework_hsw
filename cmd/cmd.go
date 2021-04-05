package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type SCmd struct {
	Name			string
	Usage			string
	UsageText		string
	Description		string
	Category		string
	SubCommands		[]*cli.Command
	// Flag 구조체 생성 및 변경 필요?
	Flags			[]cli.Flag
}

var app cli.App

func init() {
	app = *cli.NewApp()
}

func (c SCmd) TestMethod(s string) (string, error) {
	fmt.Println("Message is : ", s)
	str := "TestMethod Cmd Function Start : " + s

	return str, nil
}

func (c SCmd) AddCommand() (*cli.App, error) {
	castCmd := cli.Command{
		Name:        	c.Name,
		Usage:       	c.Usage,
		UsageText:   	c.UsageText,
		Description: 	c.Description,
		Category:    	c.Category,
		Subcommands:	c.SubCommands,
		Flags:			c.Flags,
	}

	_ = append(app.Commands, &castCmd)

	return &app, nil
}

func (c SCmd) AddCommands(sc ...SCmd) (*cli.App, error) {
	for _, rsc := range sc {
		castCmd := cli.Command{
			Name:        	rsc.Name,
			Usage:       	rsc.Usage,
			UsageText:   	rsc.UsageText,
			Description: 	rsc.Description,
			Category:    	rsc.Category,
			Subcommands:	rsc.SubCommands,
			Flags:			rsc.Flags,
		}

		_ = append(app.Commands, &castCmd)
	}

	return &app, nil
}

func (c SCmd) Build() (string, error) {
	str := "Build Cmd Function Start"

	return str, nil

}