/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topics",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://github.com/topics/%s?o=desc&s=updated"

		if topics := wf.Config.GetString("Topics"); topics != "" {
			ts := strings.Split(topics, ";")
			for i := 0; i < len(ts); i++ {
				wf.NewItem(ts[i]).Title(ts[i]).Subtitle(fmt.Sprintf("Looking for %s trending", ts[i])).Arg(fmt.Sprintf(url, ts[i])).Valid(true)
			}
		}

		for _, lang := range langs {
			wf.NewItem(lang).Title(lang).Subtitle(fmt.Sprintf("Looking for %s trending", lang)).Arg(fmt.Sprintf(url, lang)).Valid(true)
		}

		if len(args) > 0 {
			wf.Filter(args[0])
		}

		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(topicCmd)
}
