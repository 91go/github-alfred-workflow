package cmd

import (
	"fmt"
	"strings"

	"github.com/91go/gh-alfredworkflow/utils"
	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

// repoSearchCmd represents the repoSearch command
var repoSearchCmd = &cobra.Command{
	Use:     "repos",
	Short:   "Searching repositories from github",
	Args:    cobra.RangeArgs(1, 4),
	Example: "icons/repos.svg",
	Run: func(cmd *cobra.Command, args []string) {
		// priority list
		xs := fmt.Sprintf("https://github.com/search?q=%s&type=repositories", strings.Join(args[0:], "+"))
		wf.NewItem("Search on github").Arg(xs).Valid(true).Icon(&aw.Icon{Value: "icons/repo.png"}).Title("Searching on github").Subtitle("searching...")
		// wf.Rerun(0.1)

		if len(args) == 1 {
			repos, err := utils.NewGithubClient(token).SearchRepositories(args[0])
			if err != nil {
				wf.FatalError(err)
			}

			for _, repo := range repos {
				url := repo.GetHTMLURL()
				fullName := repo.GetFullName()
				des := repo.GetDescription()
				// TODO 判断是否私有，不同的 icon
				// TODO 语言，star 数量
				wf.NewItem(fullName).Arg(url).Valid(true).Subtitle(des).Icon(&aw.Icon{Value: "icons/repo.png"}).Title(fullName).Autocomplete(fullName)
			}
			// if len(args) > 1 {
			// 	wf.Filter(args[1])
			// }
		}

		wf.SendFeedback()
	},
}

func init() {
	execCmd.AddCommand(repoSearchCmd)
}
