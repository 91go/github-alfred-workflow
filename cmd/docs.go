package cmd

import (
	"fmt"

	aw "github.com/deanishe/awgo"

	"github.com/spf13/cobra"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "A brief description of your command",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		markets := []Metadata{
			{item: "docs", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://docs.github.com/%s"},
			{item: "actions", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://docs.github.com/%s/actions"},
			{item: "packages", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://docs.github.com/%s/packages"},
			{item: "copilot", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://docs.github.com/%s/copilot"},
			{item: "repositories", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://docs.github.com/%s/repositories"},
			{item: "code-security", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://docs.github.com/%s/code-security"},
		}
		lang := wf.Config.GetString("lang", "en")
		for _, m := range markets {
			wf.NewItem(m.item).Icon(m.icon).Subtitle(m.subtitle).Arg(fmt.Sprintf(m.url, lang)).Valid(true).UID(m.item).Autocomplete(m.item).IsFile(true)
		}
		if len(args) > 0 {
			wf.Filter(args[0])
		}
		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
