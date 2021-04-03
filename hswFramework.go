package hswFramework

import "github.com/sangwoo310/go_framework_hsw/cmd"

type HswFrame interface {
	Cmd()
}

type SHswFrame struct {
	Frame *HswFrame
}

//func NewFrame() *SHswFrame {
//
//	return &SHswFrame{}
//}

func NewFrame() *SHswFrame {

	return &SHswFrame{}
}

func (h SHswFrame) Cmd() *cmd.CmdStruct {
	s := cmd.NewApp()

	return s
}