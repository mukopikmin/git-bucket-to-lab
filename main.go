package main

import (
	"flag"
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/handler"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/bxcodec/faker/v3"
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

	fixture := flag.Bool("fixture", false, "Generate fixture data")
	flag.Parse()

	if *fixture {
		err := generateFixtures()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		os.Exit(0)
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

func generateFixtures() error {
	fmt.Println("Generating fixture data in GitBucket ...")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), "855a9c623ef34a433f9118c0ddc52ec79b956d54")

	repo, err := b.CreateRepo(faker.Word(), faker.Sentence(), false)
	if err != nil {
		return err
	}

	fmt.Printf("Created repository : %s\n", repo.Name)

	for i := 0; i < 5; i++ {
		issue, err := b.CreateIssue(repo, faker.Sentence(), faker.Paragraph())
		if err != nil {
			return err
		}

		fmt.Printf("Created issue : #%d %s\n", issue.Number, issue.Title)

		for i := 0; i < 5; i++ {
			comment, err := b.CreateComment(repo, issue.Number, faker.Sentence())
			if err != nil {
				return err
			}

			fmt.Printf("Created issue comment : %d\n", comment.ID)
		}
	}

	return nil
}
