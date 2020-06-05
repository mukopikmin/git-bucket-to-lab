package handler

import (
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// MigrateRepo ...
func MigrateRepo(c echo.Context) error {
	owner := c.Param("owner")
	name := c.Param("name")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), os.Getenv("GITBUCKE_TOKEN"))
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), os.Getenv("GITLAB_TOKEN"))

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = repo.Clone()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := l.CreateProject(name, repo.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = project.Push()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Redirect(http.StatusFound, "/"+owner+"/"+name)
}

// MigrateIssues ...
func MigrateIssues(c echo.Context) error {
	owner := c.Param("owner")
	name := c.Param("name")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), os.Getenv("GITBUCKE_TOKEN"))
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), os.Getenv("GITLAB_TOKEN"))

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

// MigratePulls ..
func MigratePulls(c echo.Context) error {
	owner := c.Param("owner")
	name := c.Param("name")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), os.Getenv("GITBUCKE_TOKEN"))
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), os.Getenv("GITLAB_TOKEN"))

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println(repo)

	pulls, err := b.GetPulls(repo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := l.GetProject(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, p := range pulls {
		fmt.Println(p)
		_, err := l.CreateMerge(project, p.Title, p.Head.Ref, p.Base.Ref, p.Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// for _, comment := range p.Comments {
		// 	_, err := l.CreateComment(project, p, comment.Body, comment.CreatedAt)
		// 	if err != nil {
		// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		// 	}
		// }
	}

	return c.Redirect(http.StatusFound, "/"+owner+"/"+name)
}
