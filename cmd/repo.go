package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/91go/gh-alfredworkflow/utils"

	aw "github.com/deanishe/awgo"

	"github.com/spf13/cobra"
)

const updateReposJobName = "update-repos"

type Repository struct {
	LastUpdated time.Time
	URL         string
	Name        string
	User        string
	Description string
}

func (r Repository) FullName() string {
	return fmt.Sprintf("%s/%s", r.User, r.Name)
}

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Searching Starred Repositories And My Repositories",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := ListRepositories()
		if err != nil {
			wf.FatalError(err)
		}
		for _, repo := range repos {
			url := repo.URL
			items := wf.NewItem(repo.FullName()).Arg(url).Copytext(url).Quicklook(url).Largetype(repo.Description).Valid(true).Subtitle(repo.Description).Icon(&aw.Icon{Value: "icons/repo.png"}).Title(repo.FullName()).Autocomplete(repo.FullName())
			items.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}
		if len(args) > 0 {
			wf.Filter(args[0])
		}
		if !wf.IsRunning(updateReposJobName) {
			cmd := exec.Command("./exe", "actions", "update-repos")
			if err := wf.RunInBackground(updateReposJobName, cmd); err != nil {
				ErrorHandle(err)
			}
		}

		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}

// Search from sqlite
func ListRepositories() ([]Repository, error) {
	db, err := utils.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, url,description, name,user,updated_at FROM repository")
	if err != nil {
		return nil, err
	}

	var repos []Repository

	for rows.Next() {
		var id, url, descr, name, user string
		var updated time.Time
		err = rows.Scan(&id, &url, &descr, &name, &user, &updated)
		if err != nil {
			return nil, err
		}

		repos = append(repos, Repository{
			URL:         url,
			Name:        name,
			User:        user,
			Description: descr,
			LastUpdated: updated,
		})
	}

	return repos, nil
}
