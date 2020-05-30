package main

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

func showRepo(c echo.Context) error {
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), "855a9c623ef34a433f9118c0ddc52ec79b956d54")
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), "8vJG_YxuJ5K1xTt5xeM-")
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

		project.Issues = i
	}

	params := RepoParam{repo, project}

	return c.Render(http.StatusOK, "repo", params)
}

func migrateRepo(c echo.Context) error {
	owner := c.Param("owner")
	name := c.Param("name")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), "855a9c623ef34a433f9118c0ddc52ec79b956d54")
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), "8vJG_YxuJ5K1xTt5xeM-")

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = l.CreateProject(name, repo.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusFound, "/"+owner+"/"+name)
}

func migrateIssues(c echo.Context) error {
	owner := c.Param("owner")
	name := c.Param("name")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), "855a9c623ef34a433f9118c0ddc52ec79b956d54")
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), "8vJG_YxuJ5K1xTt5xeM-")

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	issues, err := b.GetIssues(repo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := l.GetProject(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, i := range issues {
		issue, err := l.CreateIssue(project, i.Title, i.Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		for _, comment := range i.Comments {
			_, err := l.CreateComment(project, issue, comment.Body, comment.CreatedAt)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}
	}

	return c.Redirect(http.StatusFound, "/"+owner+"/"+name)
}
