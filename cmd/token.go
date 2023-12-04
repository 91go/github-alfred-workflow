package cmd

import (
	"fmt"

	"github.com/91go/gh-alfredworkflow/utils/secret"

	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		store := secret.NewStore(wf)
		token := args[0]

		if len(token) == 0 {
			token, err := store.GetAPIToken()
			if err != nil {
				wf.FatalError(err)
				return
			}
			if len(token) == 0 {
				wf.WarnEmpty("No token found.", "")
				return
			}

			wf.NewItem(fmt.Sprintf("Found token:%s", token))
			return
		}

		err := store.Store(secret.KeyGithubAPIToken, token)
		if err != nil {
			wf.FatalError(err)
			return
		}
		// wf.Config.Set(secret.KeyGithubAPIToken, cmd.Token)

		wf.NewItem("Success.")
	},
}

func init() {
	actionsCmd.AddCommand(tokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
