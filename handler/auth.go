package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// AuthParam ...
type AuthParam struct {
	GitBucketURL string `json:"gitbucket_url"`
	GitLabURL    string `json:"gitlab_url"`
}

// Auth ...
func Auth(c echo.Context) error {
	params := &AuthParam{os.Getenv("GITBUCKET_URL"), os.Getenv("GITLAB_URL")}

	return c.JSON(http.StatusOK, params)
}
