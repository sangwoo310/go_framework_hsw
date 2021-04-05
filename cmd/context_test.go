package cmd

import (
	"github.com/urfave/cli/v2"
	"testing"
)

func TestNewContext(t *testing.T) {
	//app := cli.NewApp()
	//flags := flag.NewFlagSet("flagName", 1)
	//ctx := cli.NewContext(*app, flags)

	ctx := cli.Context{
		Context: nil,
		App:     nil,
		Command: nil,
	}
}