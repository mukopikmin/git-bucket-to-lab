package main

import (
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template is templates of views
type Template struct {
	templates *template.Template
}

// RepoProject ...
type RepoProject struct {
	Repo    *gitbucket.Repo
	Project *gitlab.Project
}

// IndexParam ...
type IndexParam struct {
	GitBucketURL string
	GitLabURL    string
	RepoProject  []RepoProject
}

// Render acts as renderer of templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", index)
	e.GET("/:owner/:name", showRepo)
	e.POST("/:owner/:name/repo", migrateRepo)
	e.POST("/:owner/:name/issues", migrateIssues)

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), "855a9c623ef34a433f9118c0ddc52ec79b956d54")
	l := gitlab.NewClient(os.Getenv("GITLAB_URL"), "8vJG_YxuJ5K1xTt5xeM-")

	fmt.Println(c.Request().Header)

	repos, err := b.GetRepos()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	projects, err := l.GetProjects()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := IndexParam{b.Endpoint, l.Endpoint, []RepoProject{}}
	for i, r := range repos {
		var project *gitlab.Project
		for _, p := range projects {
			if r.FullName == p.PathWithNamespace {
				project = &p
				break
			}
		}

		params.RepoProject = append(params.RepoProject, RepoProject{&repos[i], project})
	}

	return c.Render(http.StatusOK, "index", params)
}
