/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	myCmd.AddCommand(topicCmd)
}
