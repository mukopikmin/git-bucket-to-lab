package main

import (
	"git-bucket-to-lab/handler"
	"io"
	"log"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template is templates of views
type Template struct {
	templates *template.Template
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

	e.GET("/", handler.Index)
	e.GET("/:owner/:name", handler.ShowRepo)
	e.POST("/:owner/:name/repo", handler.MigrateRepo)
	e.POST("/:owner/:name/issues", handler.MigrateIssues)
	e.POST("/:owner/:name/pulls", handler.MigratePulls)

	e.Logger.Fatal(e.Start(":1323"))
}
