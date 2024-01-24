package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/91go/gh-alfredworkflow/utils"

	aw "github.com/deanishe/awgo"

	"github.com/spf13/cobra"
)

const syncJob = "sync"

const (
	GistSearch = "https://gist.github.com/search?q=%s"
	RepoSearch = "https://github.com/search?q=%s&type=repositories"
)

type Repository struct {
	LastUpdated time.Time
	URL         string
	Name        string
	User        string
	Description string
}

type Repo struct {
	Feat string `yaml:"feat"`
	Name string `yaml:"name"`
}

func (r Repository) FullName() string {
	return fmt.Sprintf("%s/%s", r.User, r.Name)
}

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Searching from starred repositories and my repositories",
	// Args:    cobra.RangeArgs(0, 2),
	Example: "icons/repo.svg",
	PostRun: func(cmd *cobra.Command, args []string) {
		if !wf.IsRunning(syncJob) {
			cmd := exec.Command("./exe", "list", "actions", syncJob)
			if err := wf.RunInBackground(syncJob, cmd); err != nil {
				ErrorHandle(err)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := ListRepositories()
		if err != nil {
			wf.FatalError(err)
		}

		var ghs []Repo
		fy := wf.Cache.Dir + "/gh.yml"
		if wf.Cache.Exists(fy) {

			f, err := os.Open(fy)
			if err != nil {
				return
			}
			d := yaml.NewDecoder(f)

			for {
				err := d.Decode(&ghs)
				if ghs == nil {
					continue
				}
				if errors.Is(err, io.EOF) {
					break
				}
				if err != nil {
					panic(err)
				}
			}
		}

		for _, repo := range repos {
			url := repo.URL

			item := wf.NewItem(repo.FullName()).
				Arg(url).
				Copytext(url).
				Valid(true).
				Title(repo.FullName()).
				Autocomplete(repo.FullName())

			for _, gh := range ghs {
				if gh.Name == url {
					item.Icon(&aw.Icon{Value: "icons/arrow.svg"}).Subtitle(gh.Feat)
				} else {
					item.Icon(&aw.Icon{Value: "icons/repo.svg"}).Subtitle(repo.Description)
				}
			}

			item.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}

		if len(args) > 0 {
			wf.Filter(args[0])
		}

		wf.NewItem("Search On Github").
			Arg(fmt.Sprintf(RepoSearch, strings.Join(args, "+"))).
			Valid(true).
			Icon(&aw.Icon{Value: "icons/search.svg"}).
			Title("Searching On Github").
			Subtitle(fmt.Sprintf("%s %s", "searching...", strings.Join(args, " ")))
		wf.NewItem("Search On Github Gist").
			Arg(fmt.Sprintf(GistSearch, strings.Join(args, "+"))).
			Valid(true).
			Icon(&aw.Icon{Value: "icons/gist.svg"}).
			Title("Searching On Github Gist").
			Subtitle(fmt.Sprintf("%s %s", "searching...", strings.Join(args, " ")))

		wf.SendFeedback()
	},
}

func init() {
	listCmd.AddCommand(repoCmd)
}

// Search from sqlite
func ListRepositories() ([]Repository, error) {
	db, err := utils.OpenDB(wf.CacheDir() + "/repo.db")
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
