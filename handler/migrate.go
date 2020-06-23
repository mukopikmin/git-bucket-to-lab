package handler

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"net/http"
	"os"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/labstack/echo/v4"
)

// MigrateUserRepo ...
func MigrateUserRepo(c echo.Context) error {
	h := c.Request().Header
	user := h.Get("X-GITBUCKET-USER")
	btoken := h.Get("X-GITBUCKET-TOKEN")
	ltoken := h.Get("X-GITLAB-TOKEN")
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), btoken)
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), ltoken)
	owner := c.Param("owner")
	name := c.Param("name")
	storage := memory.NewStorage()
	worktree := memfs.New()

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = repo.Clone(storage, worktree, user, btoken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	luser, err := l.GetAuthorizedUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := l.GetProject(owner, name)
	if project == nil {
		project, err = l.CreateProject(luser.ID, repo.Name, repo.Description, repo.Private)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	err = project.Push(storage, worktree, ltoken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ShowRepo(c)
}

// MigrateGroupRepo ...
func MigrateGroupRepo(c echo.Context) error {
	h := c.Request().Header
	user := h.Get("X-GITBUCKET-USER")
	btoken := h.Get("X-GITBUCKET-TOKEN")
	ltoken := h.Get("X-GITLAB-TOKEN")
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), btoken)
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), ltoken)
	owner := c.Param("owner")
	name := c.Param("name")
	storage := memory.NewStorage()
	worktree := memfs.New()

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = repo.Clone(storage, worktree, user, btoken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	group, err := l.GetGroup(owner)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := l.GetProject(owner, name)
	if project == nil {
		project, err = l.CreateProject(group.ID, name, repo.Description, repo.Private)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	err = project.Push(storage, worktree, ltoken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ShowRepo(c)
}

// MigrateIssues ...
func MigrateIssues(c echo.Context) error {
	h := c.Request().Header
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), h.Get("X-GITBUCKET-TOKEN"))
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), h.Get("X-GITLAB-TOKEN"))
	owner := c.Param("owner")
	name := c.Param("name")

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := l.GetProject(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, i := range repo.Issues {
		issue, err := l.CreateIssue(project, i.Title, i.Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		for _, comment := range i.Comments {
			_, err := l.CreateIssueComment(project, issue, comment.Body, comment.CreatedAt)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}
	}

	return ShowRepo(c)
}

// MigratePulls ..
func MigratePulls(c echo.Context) error {
	h := c.Request().Header
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), h.Get("X-GITBUCKET-TOKEN"))
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), h.Get("X-GITLAB-TOKEN"))
	owner := c.Param("owner")
	name := c.Param("name")

	repo, err := b.GetRepo(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	project, err := l.GetProject(owner, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, p := range repo.Pulls {
		m, err := l.CreateMerge(project, p.Title, p.Head.Ref, p.Base.Ref, p.Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		for _, comment := range p.Comments {
			_, err := l.CreateMergeComment(project, m, comment.Body, comment.CreatedAt)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}
	}

	return ShowRepo(c)
}
