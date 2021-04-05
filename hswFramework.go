package hswFramework

import "github.com/sangwoo310/go_framework_hsw/cmd"

type IFrame interface {
	Cmd() *cmd.SCmd
}

type SFrame struct {
	Name string
}

func NewApp() IFrame {
	var framework IFrame = &SFrame{}

	return framework
}

func (f SFrame) Cmd() *cmd.SCmd {
	return &cmd.SCmd{}
}


//func (f SFrame) Cmd() *cmd.CmdInterface {
//	var s cmd.CmdInterface = &cmd.CmdStruct{}
//
//	return &s
//}

