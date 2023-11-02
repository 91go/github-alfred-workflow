package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/v56/github"

	aw "github.com/deanishe/awgo"
	"github.com/gregjones/httpcache"
	"github.com/spf13/cobra"
)

type My struct {
	icon     *aw.Icon
	item     string
	url      string
	subtitle string
}

// myCmd represents the my command
var myCmd = &cobra.Command{
	Use:   "my",
	Short: "list all my github shortcut actions",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		xxx := GetUsername(github.NewClient(httpcache.NewMemoryCacheTransport().Client()).WithAuthToken(token))

		my := []My{
			{item: "dashboard", url: "https://github.com/", subtitle: "View your dashboard", icon: &aw.Icon{Value: "icons/dashboard.png"}},
			{item: "notifications", url: "https://github.com/notifications", subtitle: "View your notifications", icon: &aw.Icon{Value: "icons/notifications.png"}},
			{item: "profile", url: fmt.Sprintf("https://github.com/%s", xxx), subtitle: "View your public user profile", icon: &aw.Icon{Value: "icons/profile.png"}},
			{item: "issues", url: "https://github.com/issues", subtitle: "View your issues", icon: &aw.Icon{Value: "icons/issue.png"}},
			{item: "PR", url: "https://github.com/pulls", subtitle: "View your pull requests", icon: &aw.Icon{Value: "icons/pull-request.png"}},
			{item: "repos", url: fmt.Sprintf("https://github.com/%s?tab=repositories", xxx), subtitle: "View your repositories", icon: &aw.Icon{Value: "icons/repo.png"}},
			{item: "New", url: "https://github.com/new", subtitle: "", icon: &aw.Icon{Value: "icons/new.png"}},
			{item: "settings", url: "https://github.com/settings", subtitle: "View or edit your account settings", icon: &aw.Icon{Value: "icons/settings.png"}},
			{item: "stars", url: fmt.Sprintf("https://github.com/%s?tab=stars", xxx), subtitle: "View your starred repositories", icon: &aw.Icon{Value: "icons/stars.png"}},
			{item: "gist", url: fmt.Sprintf("https://gist.github.com/%s", xxx), subtitle: "View your gists", icon: &aw.Icon{Value: "icons/gists.png"}},
			{item: "topic", url: fmt.Sprintf("https://github.com/%s?tab=stars", xxx), subtitle: "View your starred topics", icon: &aw.Icon{Value: "icons/topics.png"}},
		}

		// list all sub
		for _, m := range my {
			items := wf.NewItem(m.item).Arg(m.url).Copytext(m.url).Quicklook(m.url).Largetype(m.subtitle).Valid(true).Subtitle(m.subtitle).Icon(m.icon).Title(m.item).Autocomplete(m.item)
			items.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}
		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(myCmd)
}

func GetUsername(client *github.Client) string {
	user := &github.User{}
	response, _, err := client.Users.Get(context.Background(), user.GetName())
	if err != nil {
		return ""
	}
	return response.GetLogin()
}
