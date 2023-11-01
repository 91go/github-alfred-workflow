package cmd

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

type My struct {
	icon     *aw.Icon
	item     string
	url      string
	subtitle string
}

var my = []My{
	{item: "Dashboard", url: "https://github.com/", subtitle: "View your dashboard", icon: &aw.Icon{Value: "icons/dashboard.png"}},
	{item: "Notification", url: "https://github.com/notifications", subtitle: "View your notifications", icon: &aw.Icon{Value: "icons/notifications.png"}},
	{item: "Profile", url: "https://github.com/%s", subtitle: "View your public user profile", icon: &aw.Icon{Value: "icons/profile.png"}},
	{item: "Issue", url: "https://github.com/issues", subtitle: "View your issues", icon: &aw.Icon{Value: "icons/issues.png"}},
	{item: "PR", url: "https://github.com/pulls", subtitle: "View your pull requests", icon: &aw.Icon{Value: "icons/pull-request.png"}},
	{item: "repo", url: "https://github.com/%s?tab=repositories", subtitle: "View your repositories", icon: &aw.Icon{Value: "icons/repo.png"}},
	{item: "New", url: "https://github.com/new", subtitle: "", icon: &aw.Icon{Value: "icons/new.png"}},
	{item: "setting", url: "https://github.com/settings", subtitle: "View or edit your account settings", icon: &aw.Icon{Value: "icons/settings.png"}},
	{item: "star", url: "https://github.com/%s?tab=stars", subtitle: "View your starred repositories", icon: &aw.Icon{Value: "icons/stars.png"}},
	{item: "gist", url: "https://gist.github.com/%s", subtitle: "View your gists", icon: &aw.Icon{Value: "icons/gists.png"}},
	{item: "topic", url: "https://github.com/%s?tab=stars", subtitle: "View your starred topics", icon: &aw.Icon{Value: "icons/topics.png"}},
}

// myCmd represents the my command
var myCmd = &cobra.Command{
	Use:   "my",
	Short: "list all my github shortcut actions",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		// list all sub
		for _, m := range my {
			url := fmt.Sprintf(m.url, "hapihacking")
			items := wf.NewItem(m.item).Arg(url).Copytext(url).Quicklook(url).Largetype(m.subtitle).Valid(true).Subtitle(m.subtitle).Icon(m.icon).Title(m.item).Autocomplete(m.item)
			items.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}
		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(myCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// myCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// myCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func addRepos(repos[], comparatorPrefix string) {
//
// }
