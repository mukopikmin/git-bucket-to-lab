package handler

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// RepoParam ...
type RepoParam struct {
	Repo    *gitbucket.Repo
	Project *gitlab.Project
}

// ShowRepo ...
func ShowRepo(c echo.Context) error {
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), os.Getenv("GITBUCKE_TOKEN"))
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), os.Getenv("GITLAB_TOKEN"))
	owner := c.Param("owner")
	name := c.Param("name")

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	issues, err := b.GetIssues(repo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pulls, err := b.GetPulls(repo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	repo.Issues = issues
	repo.Pulls = pulls

	project, err := l.GetProject(owner, name)
	if err == nil {
		i, err := l.GetIssues(project)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		m, err := l.GetMerges(project)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		project.Issues = i
		project.Merges = m
	}

	params := RepoParam{repo, project}

	return c.Render(http.StatusOK, "repo", params)
}
