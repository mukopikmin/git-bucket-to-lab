package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Auth ...
func Auth(c echo.Context) error {
	return c.Render(http.StatusOK, "auth", nil)
}
