package handler

import (
	"git-bucket-to-lab/migration"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Repos ...
func Repos(c echo.Context) error {
	h := c.Request().Header
	client, err := migration.NewClient(h.Get("X-GITBUCKET-TOKEN"), h.Get("X-GITLAB-TOKEN"))
	if err != nil {
		return err
	}

	params, err := client.GetMigrations()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, params)
}

// Repo ...
func Repo(c echo.Context) error {
	h := c.Request().Header
	owner := c.Param("owner")
	name := c.Param("name")

	client, err := migration.NewClient(h.Get("X-GITBUCKET-TOKEN"), h.Get("X-GITLAB-TOKEN"))
	if err != nil {
		return err
	}

	m, err := client.GetMigration(owner, name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m)
}
