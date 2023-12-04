/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/91go/gh-alfredworkflow/utils"
	"github.com/google/go-github/v56/github"

	"github.com/spf13/cobra"
)

// updateReposCmd represents the updateRepos command
var updateReposCmd = &cobra.Command{
	Use:    "update-repos",
	Short:  "A brief description of your command",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := UpdateRepositories(token)
		if err != nil {
			wf.FatalError(err)
		}
	},
}

func init() {
	actionsCmd.AddCommand(updateReposCmd)
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
