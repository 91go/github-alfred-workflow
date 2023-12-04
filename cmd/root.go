package cmd

import (
	"log"
	"os"

	"github.com/91go/gh-alfredworkflow/utils/secret"
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/spf13/cobra"
)

var (
	repo = "91go/gh-alfredworkflow"
	wf   *aw.Workflow
	av   = aw.NewArgVars()
	// gt    = "GITHUB_TOKEN"
	token string
)

// ErrorHandle handle error
func ErrorHandle(err error) {
	av.Var("error", err.Error())
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args to Alfred: %v", err)
	}
}

// checkEnv Get github-token from keychain directly
func checkEnv(cmd *cobra.Command, args []string) {
	// if token = wf.Config.GetString(gt); token == "" {
	// 	wf.NewItem("Please set your github token first").Valid(false).Icon(&aw.Icon{Value: "icons/warning.png"})
	// 	wf.SendFeedback()
	// 	return
	// }
	if cmd.Use == "token" {
		return
	}
	store := secret.NewStore(wf)
	if token, _ = store.Get(secret.KeyGithubAPIToken); token == "" {
		wf.NewItem("Please set your github token first").Valid(false).Icon(&aw.Icon{Value: "icons/warning.png"})
		wf.SendFeedback()
		return
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "gh",
	Short:            "list all subcommands",
	PersistentPreRun: checkEnv,
	Run: func(cmd *cobra.Command, args []string) {
		// get all subcommands
		for _, c := range cmd.Commands() {
			wf.NewItem(c.Name()).Arg(c.Name()).Valid(true).Icon(&aw.Icon{Value: "icons/repo.png"}).Title(c.Name()).Subtitle(c.Short).Autocomplete(c.Name())
		}
		if len(args) > 0 {
			wf.Filter(args[0])
		}
		// wf.WarnEmpty("No matching commands", "Try a different query?")
		wf.SendFeedback()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd.PersistentFlags().StringVar(&token, gt, "", gt)
	wf.Run(func() {
		if err := rootCmd.Execute(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	})
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
	wf.Args() // magic for "workflow:update"
}
