/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/91go/gh-alfredworkflow/utils"
	"github.com/google/go-github/v56/github"

	"github.com/spf13/cobra"
)

// syncCmd represents the updateRepos command
var syncCmd = &cobra.Command{
	Use:    "sync",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// sync repo
		if _, err := UpdateRepositories(token); err != nil {
			// wf.NewWarningItem("Sync Failed.", err.Error()).Valid(false).Title("Sync Failed.")
			// wf.SendFeedback()
			slog.Error("Sync Failed.", slog.Any("err", err))
		}

		// gh.yml
		url := wf.Config.GetString("url")
		if url != "" {
			resp, err := http.Get(url)
			if err != nil {
				slog.Error("request error", slog.Any("err", err))
				return
			}
			defer resp.Body.Close()

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}
			err = wf.Cache.Store(CustomRepo, data)
			if err != nil {
				return
			}
		}

		// wf.NewItem("Sync Repos Successfully.").Title("Sync Repos Successfully.").Valid(false)
		// wf.SendFeedback()
		slog.Info("Sync Repos Successfully.")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func UpdateRepositories(token string) (int64, error) {
	// my repos
	userRepos, err := utils.NewGithubClient(token).ListUserRepositories()
	if err != nil {
		return 0, err
	}

	// starred repos
	starredRepos, err := utils.NewGithubClient(token).ListStarredRepositories()
	if err != nil {
		return 0, err
	}

	db, err := utils.OpenDB(wf.CacheDir() + "/repo.db")
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
