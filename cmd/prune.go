/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pruneCmd represents the logout command
var pruneCmd = &cobra.Command{
	Use:   "clear-caches",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// 删除 my.json
		err := wf.ClearCache()
		if err != nil {
			wf.FatalError(err)
		}
	},
}

func init() {
	actionsCmd.AddCommand(pruneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pruneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pruneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
