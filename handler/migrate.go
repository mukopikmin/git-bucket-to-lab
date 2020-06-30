package handler

import (
	"git-bucket-to-lab/migration"
	"net/http"

	"github.com/labstack/echo/v4"
)

// MigrateRepo ...
func MigrateRepo(c echo.Context) error {
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

	m, err = client.ExecRepo(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m)
}

// MigrateIssues ...
func MigrateIssues(c echo.Context) error {
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

	m, err = client.ExecIssues(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m)
}

// MigratePulls ..
func MigratePulls(c echo.Context) error {
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

	m, err = client.ExecPulls(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, m)
}
