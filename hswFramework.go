package hswFramework

import "github.com/sangwoo310/go_framework_hsw/cmd"

type IHswFrame interface {
	Cmd()
}

type SHswFrame struct {
	Name string
}

func NewFrame(*IHswFrame) *SHswFrame {

	s := SHswFrame{}

	return &s
}

func (h SHswFrame) Cmd() *cmd.CmdStruct {
	s := cmd.NewApp()

	return s
}