package cmd

import (
	"fmt"
	"github.com/samber/lo"
	"strings"

	"github.com/spf13/cobra"
)

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:     "topics",
	Short:   "List all starred topics",
	Example: "icons/topics.svg",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://github.com/topics/%s?o=desc&s=updated"

		if topics := wf.Config.GetString("topics"); topics != "" {
			ts := strings.Split(topics, ";")
			// sanitize topics
			tsN := lo.WithoutEmpty(ts)
			for i := 0; i < len(tsN); i++ {
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
	listCmd.AddCommand(topicCmd)
}
