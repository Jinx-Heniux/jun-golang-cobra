/*
Copyright © 2022 jinx-heniux

*/

// Cobra 中文文档 - 掘金

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var echoTimes int

// timesCmd represents the times command
var timesCmd = &cobra.Command{
	Use:   "times [# times] [string to echo]",
	Short: "Echo anyting to the screen more times",
	Long: `echo things multiple times back to the user y providing
	a count and a string.`,
	Args: cobra.MinimumNArgs(1),

	// 如果子命令中没有PersistentPreRun，将会调用父命令中的PersistentPreRun
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside timesCmd PersistentPreRun with args: %v\n", args)
	},

	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside timesCmd PreRun with args: %v\n", args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside timesCmd Run with args: %v\n", args)
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},

	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside timesCmd PostRun with args: %v\n", args)
	},

	// 如果子命令中没有PersistentPreRun，将会调用父命令中的PersistentPreRun
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside timesCmd PersistentPostRun with args: %v\n", args)
	},
}

func init() {
	// rootCmd.AddCommand(timesCmd)
	echoCmd.AddCommand(timesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	timesCmd.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

}
