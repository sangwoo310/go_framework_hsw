package hswFramework

import (
	"fmt"
	"github.com/sangwoo310/go_framework_hsw/cmd"
	"github.com/spf13/cobra"
	"github.com/urfave/cli/v2"
	"strings"
	"testing"
)

func Test_command(t *testing.T) {

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdWow = &cobra.Command{
		Use:   "wow",
		Short: "wow1",
		Long: `wow2`,
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("wow start")
		},
	}

	var cmdWow2 = &cobra.Command{
		Use:   "wow3",
		Short: "wow4",
		Long: `wow5`,
		Args: cobra.MaximumNArgs(5),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("wow start")
		},
	}

	var cmdFlag = &cobra.Command{
		Use:   "flag",
		Short: "flag1",
		Long: `example for flag option`,
		Args: cobra.MinimumNArgs(0),
		Run: qq,
	}
	var text string
	cmdFlag.PersistentFlags().StringVar(&text, "text2", "123", "text option")

	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(cmdPrint, cmdWow, cmdFlag, cmdWow2)
	rootCmd.Execute()
}

func qq(cmd *cobra.Command, args []string) {
	fmt.Println("qq start")
}

func init() {
}

func TestInterface(t *testing.T)  {
	//// 인터페이스 함수로 강제함
	//var ci cmd.CmdInterface = &cmd.CmdStruct{}
	//ci.SetCommand()
	//
	//// 클래스 함수를 사용가능
	//ci2 := &cmd.CmdStruct{}
	//ci2.InterTest()
}

func TestCmdAddCommand(t *testing.T)  {
	q := cmd.SCmd{
		Name:        "",
		Usage:       "",
		UsageText:   "",
		Description: "",
		Category:    "",
		Action:      scTe,
		SubCommands: nil,
		Flags:       nil,
	}

	fmt.Println(q)
}

func scTe(ctx *cli.Context) error {
	return nil
}