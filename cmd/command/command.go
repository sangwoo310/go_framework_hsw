package command

import "github.com/spf13/cobra"

type Commander interface {
	SetStruct()

	ReadConfig(configPath string, configStruct interface{})

	AddCommand(commandName string)

	AddFlag(flagName string)
}

type Command struct {
	Command 			cobra.Command
}

func (c Command) ReadConfig() (*Commander, error) {

}

func (c Command) AddCommand() (*Commander, error) {

}

func (c Command) AddFlag() (*Commander, error) {

}

