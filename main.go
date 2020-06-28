package main

import (
	"flag"
	"fmt"
	"git-bucket-to-lab/fixture"
	"git-bucket-to-lab/handler"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// APIError Status:...
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	f := flag.Bool("fixture", false, "Generate fixture data")
	flag.Parse()

	if *f {
		err := fixture.Generate()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/auth", handler.Auth)
	e.GET("/api/auth/state", handler.AuthState)
	e.GET("/api/repos", handler.Repos)
	e.GET("/api/:owner/:name", handler.Repo)
	e.POST("/api/:owner/:name/repo", handler.MigrateUserRepo)
	e.POST("/api/:owner/:name/repo/group", handler.MigrateGroupRepo)
	e.POST("/api/:owner/:name/issues", handler.MigrateIssues)
	e.POST("/api/:owner/:name/pulls", handler.MigratePulls)
	e.Static("*", "view/dist")

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		path := c.Request().URL.Path

		if strings.HasPrefix(path, "/api") {
			if he, ok := err.(*echo.HTTPError); ok {
				code := he.Code
				message := he.Message
				c.JSON(code, APIError{
					Status:  code,
					Message: message.(string),
				})
			}
		} else {
			if err := c.File("view/dist/index.html"); err != nil {
				c.Logger().Error(err)
			}
		}
	}

	e.Logger.Fatal(e.Start(":1323"))
}
