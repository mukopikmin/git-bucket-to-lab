package handler

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Pair ...
type Pair struct {
	Repo    *gitbucket.Repo `json:"repo"`
	Project *gitlab.Project `json:"project"`
}

// IndexParam ...
type IndexParam struct {
	GitbucketUser *gitbucket.User `json:"gitbucket_user"`
	GitlabUser    *gitlab.User    `json:"gitlab_user"`
	Pairs         []Pair          `json:"pairs"`
}

// Index ...
func Index(c echo.Context) error {
	h := c.Request().Header
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), h.Get("X-GITBUCKET-TOKEN"))
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), h.Get("X-GITLAB-TOKEN"))

	buser, err := b.GetAuthorizedUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	luser, err := l.GetAuthorizedUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	repos, err := b.GetRepos()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	projects, err := l.GetProjects()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := IndexParam{buser, luser, []Pair{}}
	for i, r := range repos {
		var project *gitlab.Project
		for _, p := range projects {
			if r.FullName == p.PathWithNamespace {
				project = &p
				break
			}
		}

		params.Pairs = append(params.Pairs, Pair{&repos[i], project})
	}

	return c.JSON(http.StatusOK, params)
}
