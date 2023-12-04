package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/91go/gh-alfredworkflow/utils"
	"github.com/91go/gh-alfredworkflow/utils/secret"
	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/v56/github"
	"github.com/spf13/cobra"
)

var actions = []Metadata{
	{item: "actions token", subtitle: "Enter to set github token", icon: &aw.Icon{Value: "icons/actions-token.svg"}},
	{item: "actions sync", subtitle: "Enter to flush repositories local database", icon: &aw.Icon{Value: "icons/actions-sync.svg"}},
	{item: "actions update", subtitle: "Enter to check workflow's update", icon: &aw.Icon{Value: "icons/actions-update.svg"}},
	{item: "actions clean", subtitle: "Enter to clear caches", icon: &aw.Icon{Value: "icons/actions-clean.svg"}},
}

// actionsCmd represents the actions command
var actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "LIST ALL ACTIONS",
	Run: func(cmd *cobra.Command, args []string) {
		for _, m := range actions {
			items := wf.NewItem(m.item).Valid(true).Subtitle(m.subtitle).Icon(m.icon)
			items.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}

		wf.SendFeedback()
	},
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "UPDATE WORKFLOW",
	Run: func(cmd *cobra.Command, args []string) {
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		wf.NewItem("Success.")
	},
}

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

// pruneCmd represents the logout command
var pruneCmd = &cobra.Command{
	Use:   "clean",
	Short: "CLEAR CACHES",
	Run: func(cmd *cobra.Command, args []string) {
		// 删除 my.json
		err := wf.ClearCache()
		if err != nil {
			wf.FatalError(err)
		}
		wf.NewItem("Success.")
	},
}

// syncCmd represents the updateRepos command
var syncCmd = &cobra.Command{
	Use:    "sync",
	Short:  "A brief description of your command",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := UpdateRepositories(token)
		if err != nil {
			wf.FatalError(err)
		}
		wf.NewItem("Success.")
	},
}

func init() {
	rootCmd.AddCommand(actionsCmd)
	actionsCmd.AddCommand(updateCmd)
	actionsCmd.AddCommand(tokenCmd)
	actionsCmd.AddCommand(pruneCmd)
	actionsCmd.AddCommand(syncCmd)
}

func UpdateRepositories(token string) (int64, error) {
	userRepos, err := utils.NewGithubClient(token).ListUserRepositories()
	if err != nil {
		return 0, err
	}

	starredRepos, err := utils.NewGithubClient(token).ListStarredRepositories()
	if err != nil {
		return 0, err
	}

	db, err := utils.OpenDB()
	if err != nil {
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	found := map[string]struct{}{}
	counter := int64(0)

	for _, repo := range append(userRepos, starredRepos...) {
		log.Printf("Updating %s/%s", *repo.Owner.Login, *repo.Name)

		name := fmt.Sprintf("%s/%s", *repo.Owner.Login, *repo.Name)
		res, err := db.Exec(
			`INSERT OR REPLACE INTO repository (
					id,
					url,
					description,
					name, user,
					pushed_at,
					updated_at,
					created_at
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			name,
			nilableString(repo.HTMLURL),
			nilableString(repo.Description),
			*repo.Name,
			*repo.Owner.Login,
			githubTime(repo.PushedAt),
			githubTime(repo.UpdatedAt),
			githubTime(repo.CreatedAt),
		)
		if err != nil {
			return counter, err
		}
		found[name] = struct{}{}
		rows, _ := res.RowsAffected()
		counter += rows
	}

	existing, err := ListRepositories()
	if err != nil {
		return 0, err
	}

	// purge repos that don't exit any more
	for _, repo := range existing {
		if _, exists := found[repo.FullName()]; !exists {
			log.Printf("Repo %s doesn't exist, deleting", repo.FullName())

			_, err := db.Exec(
				`DELETE FROM repository WHERE id=?`,
				repo.FullName(),
			)
			if err != nil {
				return 0, err
			}

		}
	}

	return counter, tx.Commit()
}

func nilableString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func githubTime(t *github.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	return &t.Time
}
