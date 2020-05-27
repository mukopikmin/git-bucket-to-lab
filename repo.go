package main

import (
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RepoParam ...
type RepoParam struct {
	Repo    *gitbucket.Repo
	Project *gitlab.Project
	Issues  []gitbucket.Issue
	Pulls   []gitbucket.Pull
}

func showRepo(c echo.Context) error {
	b := gitbucket.NewClient("http://localhost:8080", "855a9c623ef34a433f9118c0ddc52ec79b956d54")
	l := gitlab.NewClient("http://localhost/api/v4", "8vJG_YxuJ5K1xTt5xeM-")
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

	project, err := l.GetProject(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := RepoParam{repo, project, issues, pulls}

	return c.Render(http.StatusOK, "repo", params)
}

func migrateRepo(c echo.Context) error {
	owner := c.Param("owner")
	name := c.Param("name")

	b := gitbucket.NewClient("http://localhost:8080", "855a9c623ef34a433f9118c0ddc52ec79b956d54")
	l := gitlab.NewClient("http://localhost/api/v4", "8vJG_YxuJ5K1xTt5xeM-")

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

	project, err := l.CreateProject(name, repo.Description)
	if err != nil {
		fmt.Println("aaaaaaaaaaaaa")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := RepoParam{repo, project, issues, pulls}

	return c.Render(http.StatusOK, "repo", params)
}
