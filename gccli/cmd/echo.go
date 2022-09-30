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

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(1),

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside echoCmd PersistentPreRun with args: %v\n", args)
	},

	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside echoCmd PreRun with args: %v\n", args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside echoCmd Run with args: %v\n", args)
		fmt.Println("Print: " + strings.Join(args, " "))
	},

	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside echoCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside echoCmd PersistentPostRun with args: %v\n", args)
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
