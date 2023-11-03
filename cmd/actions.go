package cmd

import (
	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

type Action struct {
	item     string
	icon     *aw.Icon
	subtitle string
	args     string
}

var actions = []Action{
	{item: "actions update-repos", subtitle: "Enter to flush repositories local database", icon: &aw.Icon{Value: "icons/actions-update.svg"}, args: "update-repos"},
	{item: "actions update-workflow", subtitle: "Enter to check workflow's update", icon: &aw.Icon{Value: "icons/actions-download.svg"}, args: "update"},
	{item: "actions clear-caches", subtitle: "Enter to clear caches", icon: &aw.Icon{Value: "icons/actions-cache.svg"}, args: "flush"},
}

// actionsCmd represents the actions command
var actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		for _, m := range actions {
			items := wf.NewItem(m.item).Largetype(m.subtitle).Valid(true).Subtitle(m.subtitle).Icon(m.icon).Title(m.item).Autocomplete(m.item)
			items.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}
		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(actionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
