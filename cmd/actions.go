package cmd

// actionsCmd represents the actions command
// var actionsCmd = &cobra.Command{
// 	Use:     "actions",
// 	Short:   "Common Operations",
// 	Example: "icons/settings.png",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		actions := []Metadata{
// 			{item: "actions token", subtitle: "Enter to set github token", icon: &aw.Icon{Value: "icons/actions-token.svg"}},
// 			// {item: "actions sync", subtitle: "Enter to flush repositories local database", icon: &aw.Icon{Value: "icons/actions-sync.svg"}},
// 			// {item: "actions update", subtitle: "Enter to check workflow's update", icon: &aw.Icon{Value: "icons/actions-update.svg"}},
// 			{item: "actions clean", subtitle: "Enter to clear caches", icon: &aw.Icon{Value: "icons/actions-clean.svg"}},
// 		}
// 		for _, m := range actions {
// 			wf.NewItem(m.item).Valid(false).Subtitle(m.subtitle).Icon(m.icon).Autocomplete(m.item).Title(m.item)
// 		}
// 		wf.SendFeedback()
// 	},
// }

// updateCmd represents the update command
// var updateCmd = &cobra.Command{
// 	Use:   "update",
// 	Short: "UPDATE WORKFLOW",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if err := wf.CheckForUpdate(); err != nil {
// 			// wf.FatalError(err)
// 			wf.NewWarningItem("No Available Releases Found.", "Please check later.").Valid(false).Title("No Available Releases Found.")
// 			wf.SendFeedback()
// 		}
// 		wf.NewItem("Workflow Release is Available.").Valid(false).Title("Workflow Release is Available.")
// 		wf.SendFeedback()
// 	},
// }

// func CheckForUpdate() {
// 	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
// 		logrus.Println("Running update check in background...")
// 		cmd := exec.Command(os.Args[0], "update")
// 		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
// 			logrus.Printf("Error starting update check: %s", err)
// 		}
// 	}
//
// 	if wf.UpdateAvailable() {
// 		wf.Configure(aw.SuppressUIDs(true))
// 		wf.NewItem("An update is available!").
// 			Subtitle("⇥ or ↩ to install update").
// 			Valid(false).
// 			Autocomplete("workflow:update").
// 			Icon(&aw.Icon{Value: "update-available.png"})
// 	}
// }

// tokenCmd represents the token command
// var tokenCmd = &cobra.Command{
// 	Use:   "token",
// 	Short: "A brief description of your command",
// 	Args:  cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		store := secret.NewStore(wf)
// 		token := args[0]
//
// 		if token, err := store.GetAPIToken(); err != nil || len(token) == 0 {
// 			wf.NewWarningItem("No Token Found.", "").Valid(false).Title("No Token Found.")
// 			wf.SendFeedback()
// 		}
//
// 		if err := store.Store(secret.KeyGithubAPIToken, token); err != nil {
// 			wf.NewWarningItem("Store Token Failed.", err.Error()).Valid(false).Title("Store Token Failed.")
// 			wf.SendFeedback()
// 		}
// 		wf.NewItem("Set Token Successfully.").Title("Set Token Successfully.")
// 		wf.SendFeedback()
// 	},
// }

// pruneCmd represents the logout command
// var pruneCmd = &cobra.Command{
// 	Use:   "clean",
// 	Short: "CLEAR CACHES",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// 删除 my.json
// 		// if err := wf.ClearCache(); err != nil {
// 		// 	wf.NewWarningItem("Clear Caches Failed.", err.Error()).Valid(false).Title("Clear Caches Failed.")
// 		// 	wf.SendFeedback()
// 		// }
//
// 		if err := wf.Config.Set("username", utils.NewGithubClient(token).GetUsername(), false).Do(); err != nil {
// 			wf.NewWarningItem("Clear Caches Failed.", err.Error()).Valid(false).Title("Clear Caches Failed.")
// 			wf.SendFeedback()
// 		}
// 		wf.NewItem("Clear Caches Successfully.").Title("Clear Caches Successfully.").Valid(false)
// 		wf.SendFeedback()
// 	},
// }

func init() {
	// rootCmd.AddCommand(actionsCmd)
	// // actionsCmd.AddCommand(updateCmd)
	// actionsCmd.AddCommand(tokenCmd)
	// actionsCmd.AddCommand(pruneCmd)
	// actionsCmd.AddCommand(syncCmd)
}
