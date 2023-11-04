package cmd

import (
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	repo  = "91go/gh-alfredworkflow"
	wf    *aw.Workflow
	av    = aw.NewArgVars()
	gt    = "GITHUB_TOKEN"
	token string
)

// ErrorHandle handle error
func ErrorHandle(err error) {
	av.Var("error", err.Error())
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args to Alfred: %v", err)
	}
}

// TODO
func checkEnv(cmd *cobra.Command, args []string) {
	if token = wf.Config.GetString(gt); token == "" {
		wf.NewItem("Please set your github token first").Valid(false).Icon(&aw.Icon{Value: "icons/warning.png"})
		wf.SendFeedback()
		return
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "gh",
	Short:            "gh-alfredworkflow is a Alfred shortcut actions workflow for GitHub",
	PersistentPreRun: checkEnv,
	Run: func(cmd *cobra.Command, args []string) {
		wf.SendFeedback()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().StringVar(&token, gt, "", gt)
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
