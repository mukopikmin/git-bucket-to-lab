package handler

import (
	"git-bucket-to-lab/migration"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// AuthURLs ...
type AuthURLs struct {
	GitBucketURL string `json:"gitbucket_url"`
	GitLabURL    string `json:"gitlab_url"`
}

// Auth ...
func Auth(c echo.Context) error {
	params := &AuthURLs{os.Getenv("GITBUCKET_URL"), os.Getenv("GITLAB_URL")}

	return c.JSON(http.StatusOK, params)
}

// AuthState ...
func AuthState(c echo.Context) error {
	h := c.Request().Header
	client, err := migration.NewClient(h.Get("X-GITBUCKET-TOKEN"), h.Get("X-GITLAB-TOKEN"))
	if err != nil {
		return err
	}

	params, err := client.GetAuthState()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, params)
}
