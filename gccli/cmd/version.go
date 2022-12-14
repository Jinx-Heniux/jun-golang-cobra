/*
Copyright © 2022 jinx-heniux

*/

// 视频 Golang+Cobra打造简单易用的命令行工具

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "short version information",
	Long:  `long and longer version information`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")

		author, err := cmd.Flags().GetString("author")
		if err != nil {
			fmt.Println("Author Error!")
			return
		}
		fmt.Println("Author: ", author)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	versionCmd.Flags().StringP("author", "a", "jinx heniux", "author of this argument")
}
