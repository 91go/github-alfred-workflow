package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/91go/gh-alfredworkflow/utils"
	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/v56/github"
	"github.com/spf13/cobra"
)

type Action struct {
	item     string
	icon     *aw.Icon
	subtitle string
}

var actions = []Action{
	{item: "Update Workflow", subtitle: "Enter to check update", icon: &aw.Icon{Value: "icons/update.png"}},
	{item: "Flush Repositories", subtitle: "Enter to flush repositories", icon: &aw.Icon{Value: "icons/flush.svg"}},
}

// actionsCmd represents the actions command
var actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		for _, m := range actions {
			items := wf.NewItem(m.item).Largetype(m.subtitle).Valid(true).Subtitle(m.subtitle).Icon(m.icon).Title(m.item).Autocomplete(m.item)
			items.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}
		// switch args[0] {
		// case "update":
		// 	_, err := UpdateRepositories(token)
		// 	if err != nil {
		// 		return
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(actionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
