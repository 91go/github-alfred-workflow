package cmd

import (
	"github.com/91go/gh-alfredworkflow/utils"
	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

// repoSearchCmd represents the repoSearch command
var repoSearchCmd = &cobra.Command{
	Use:   "repos",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := utils.NewGithubClient(token).SearchRepositories(args[0])
		if err != nil {
			wf.FatalError(err)
		}
		for _, repo := range repos {
			// alfred 内置 fuzzy filter, 不需要自己判断
			url := *repo.URL
			fullName := *repo.FullName
			des := *repo.Description
			// TODO 判断是否私有，不同的 icon
			// TODO 语言，star 数量
			wf.NewItem(fullName).Arg(url).Valid(true).Subtitle(des).Icon(&aw.Icon{Value: "icons/repo.png"}).Title(fullName).Autocomplete(fullName)
		}
		if len(args) > 0 {
			wf.Filter(args[0])
		}
		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(repoSearchCmd)
}
