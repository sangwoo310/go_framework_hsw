package main

import (
	"fmt"
	"github.com/spf13/cobra"
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
	Commands := &[]cobra.Command{
		{
			Use:   "flag",
			Short: "flag1",
			Long: `example for flag option`,
			Args: cobra.MinimumNArgs(0),
			Run: qq,
		},

	}
}