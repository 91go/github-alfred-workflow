package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
)

// myCmd represents the my command
var myCmd = &cobra.Command{
	Use:   "my",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("my called")
	},
}

var (
	mySubCommand = []string{"dashboard", "notifications", "profile", "issues", "pulls", "repos", "settings", "stars", "gists"}
)

func init() {
	rootCmd.AddCommand(myCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// myCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// myCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	args := wf.Args()
	if !slices.Contains(mySubCommand, args[0]) {
		ErrorHandle(fmt.Errorf("invalid My subcommand: %s", args[0]))
	}
	// 需要单独处理 pulls 和 issues
	if slices.Contains([]string{"pulls", "issues"}, args[0]) {

		subs := map[string]string{
			"created":   "Created",
			"assigned":  "Assigned",
			"mentioned": "Mentioned",
		}

		if len(args) == 1 {
			wf.NewItem("Pulls").Arg("pulls").Valid(true).Subtitle(subs[]).Icon("icons/pulls.png")
		}
		wf.NewItem("Created").Arg("created").Valid(true)
	}
}

// func addRepos(repos[], comparatorPrefix string) {
//
// }
