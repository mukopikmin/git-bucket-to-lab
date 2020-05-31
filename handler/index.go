package handler

import (
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// RepoProject ...
type RepoProject struct {
	Repo    *gitbucket.Repo
	Project *gitlab.Project
}

// IndexParam ...
type IndexParam struct {
	GitBucketURL string
	GitLabURL    string
	RepoProject  []RepoProject
}

// Index ...
func Index(c echo.Context) error {
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), "855a9c623ef34a433f9118c0ddc52ec79b956d54")
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), "8vJG_YxuJ5K1xTt5xeM-")

	fmt.Println(c.Request().Header)

	repos, err := b.GetRepos()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	projects, err := l.GetProjects()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := IndexParam{b.Endpoint, l.Endpoint, []RepoProject{}}
	for i, r := range repos {
		var project *gitlab.Project
		for _, p := range projects {
			if r.FullName == p.PathWithNamespace {
				project = &p
				break
			}
		}

		params.RepoProject = append(params.RepoProject, RepoProject{&repos[i], project})
	}

	return c.Render(http.StatusOK, "index", params)
}
