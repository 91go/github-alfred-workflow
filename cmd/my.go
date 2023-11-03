package cmd

import (
	"fmt"
	"time"

	"github.com/91go/gh-alfredworkflow/utils"

	aw "github.com/deanishe/awgo"
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
		xxx := cacheUsername()

		my := []My{
			{item: "dashboard", url: "https://github.com/", subtitle: "View your dashboard", icon: &aw.Icon{Value: "icons/dashboard.png"}},
			{item: "notifications", url: "https://github.com/notifications", subtitle: "View your notifications", icon: &aw.Icon{Value: "icons/notifications.png"}},
			{item: "profile", url: fmt.Sprintf("https://github.com/%s", xxx), subtitle: "View your public user profile", icon: &aw.Icon{Value: "icons/profile.png"}},
			{item: "issues", url: "https://github.com/issues", subtitle: "View your issues", icon: &aw.Icon{Value: "icons/issue.png"}},
			{item: "PR", url: "https://github.com/pulls", subtitle: "View your pull requests", icon: &aw.Icon{Value: "icons/pull-request.png"}},
			// {item: "repos", url: fmt.Sprintf("https://github.com/%s?tab=repositories", xxx), subtitle: "View your repositories", icon: &aw.Icon{Value: "icons/repo.png"}},
			{item: "New", url: "https://github.com/new", subtitle: "", icon: &aw.Icon{Value: "icons/new.png"}},
			{item: "settings", url: "https://github.com/settings", subtitle: "View or edit your account settings", icon: &aw.Icon{Value: "icons/settings.png"}},
			{item: "stars", url: fmt.Sprintf("https://github.com/%s?tab=stars", xxx), subtitle: "View your starred repositories", icon: &aw.Icon{Value: "icons/stars.png"}},
			{item: "gist", url: fmt.Sprintf("https://gist.github.com/%s", xxx), subtitle: "View your gists", icon: &aw.Icon{Value: "icons/gists.png"}},
		}

		for _, m := range my {
			wf.NewItem(m.item).Icon(m.icon).Subtitle(m.subtitle).Arg(m.url).Valid(true).UID(m.item).Autocomplete(m.item).IsFile(true)
		}

		if len(args) > 0 {
			wf.Filter(args[0])
		}
		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(myCmd)
}

func cacheUsername() string {
	reload := func() ([]byte, error) {
		username := utils.NewGithubClient(token).GetUsername()
		return []byte(username), nil
	}
	store, _ := wf.Cache.LoadOrStore("username", time.Duration(0), reload)
	return string(store)
}
