package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/91go/gh-alfredworkflow/utils"
	aw "github.com/deanishe/awgo"

	"github.com/google/go-github/v56/github"
	"github.com/spf13/cobra"
)

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
	Short: "search github repo directly",
	Args:  cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := ListRepositories()
		if err != nil {
			wf.FatalError(err)
		}
		for _, repo := range repos {
			// alfred 内置 fuzzy filter, 不需要自己判断
			url := repo.URL
			// TODO 判断是否私有，不同的 icon
			// TODO 语言，star 数量
			items := wf.NewItem(repo.FullName()).Arg(url).Copytext(url).Quicklook(url).Largetype(repo.Description).Valid(true).Subtitle(repo.Description).Icon(&aw.Icon{Value: "icons/repo.png"}).Title(repo.FullName()).Autocomplete(repo.FullName())
			items.Cmd().Subtitle("Press Enter to copy this url to clipboard")
		}
		if len(args) > 0 {
			wf.Filter(args[0])
		}
		wf.SendFeedback()

		// else {
		// 	// 2.2 有参数，显示搜索结果
		// 	query := args[0]
		// 	repos, err := searchRepo(query)
		// 	if err != nil {
		// 		wf.FatalError(err)
		// 	}
		//
		// 	for _, repo := range repos {
		// 		// alfred 内置 fuzzy filter, 不需要自己判断
		// 		url := repo.url
		// 		items := wf.NewItem(repo.FullName()).Arg(url).Copytext(url).Quicklook(url).Largetype(repo.Description).Valid(true).subtitle(repo.Description).Icon(&aw.Icon{Value: "icons/repos.png"}).Title(repo.FullName()).Autocomplete(repo.FullName())
		// 		items.Cmd().subtitle("Press Enter to copy this url to clipboard")
		// 	}
		// 	wf.Filter(query)
		// }
		// 3. 搜索 repo
		// 4. 显示搜索结果
		// 5. 选择结果，打开浏览器
	},
}

func init() {
	myCmd.AddCommand(repoCmd)
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

// github API
func ListStarredRepositories(client *github.Client) ([]*github.Repository, error) {
	opt := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{PerPage: 45},
		Sort:        "pushed",
	}

	var repos []*github.Repository

	for {
		result, resp, err := client.Activity.ListStarred(context.Background(), "", opt)
		if err != nil {
			return repos, err
		}
		for _, starred := range result {
			repos = append(repos, starred.Repository)
		}
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}

	return repos, nil
}

func ListUserRepositories(client *github.Client) ([]*github.Repository, error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 45},
		Sort:        "pushed",
	}

	var repos []*github.Repository

	for {
		result, resp, err := client.Repositories.List(context.Background(), "", opt)
		if err != nil {
			return repos, err
		}
		repos = append(repos, result...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}

	return repos, nil
}

// [Search - GitHub Docs](https://docs.github.com/en/rest/search/search?apiVersion=2022-11-28#search-repositories)
func SearchRepositories(client *github.Client, title string) ([]*github.Repository, error) {
	opt := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 20},
		Sort:        "stars",
	}
	var repos []*github.Repository
	result, _, err := client.Search.Repositories(context.Background(), title, opt)
	if err != nil {
		return repos, err
	}
	repos = append(repos, result.Repositories...)
	// if resp.NextPage == 0 {
	// 	break
	// }
	// opt.ListOptions.Page = resp.NextPage
	return repos, nil
}
