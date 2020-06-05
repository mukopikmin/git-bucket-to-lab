package main

// This is unused comment for creating test branch 2

import (
	"flag"
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/handler"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/bxcodec/faker/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), os.Getenv("GITBUCKE_TOKEN"))

	name := faker.Word()
	repo, err := b.CreateRepo(name, faker.Sentence(), false)
	if err != nil {
		return err
	}

	err = clone(repo.FullName, "https://github.com/mukopikmin/git-bucket-to-lab.git")
	if err != nil {
		return err
	}

	err = push(repo.FullName, repo.CloneURL)
	if err != nil {
		return err
	}

	fmt.Printf("Created repository : %s\n", repo.FullName)

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

	// repo, err = b.GetRepo(repo.Owner.Login, repo.Name)
	// if err != nil {
	// 	return err
	// }

	// for _, branch := range repo.Branches {
	// 	if branch.Name == "master" {
	// 		continue
	// 	}

	// 	p, err := b.CreatePull(repo, faker.Sentence(), "root:"+branch.Name, "master", faker.Paragraph())
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Printf("Created pull request : #%d\n", p.ID)
	// }

	return nil
}

func clone(name string, url string) error {
	r, err := git.PlainClone("tmp/"+name, false, &git.CloneOptions{
		URL:          url,
		SingleBranch: false,
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	refs, err := r.References()
	if err != nil {
		return err
	}

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		s := strings.Split(ref.Name().String(), "/")
		branch := s[len(s)-1]

		if !(len(s) > 1 && s[1] == "remotes") {
			return nil
		}

		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/remotes/origin/" + branch),
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func push(name string, url string) error {
	remote := "bucket"
	r, err := git.PlainOpen("tmp/" + name)
	if err != nil {
		return err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: remote,
		URLs: []string{url},
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	refs, err := r.References()
	if err != nil {
		return err
	}

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		s := strings.Split(ref.Name().String(), "/")
		branch := s[len(s)-1]

		if !(len(s) > 1 && s[1] == "remotes") {
			return nil
		}

		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/remotes/origin/" + branch),
		})
		if err != nil {
			return err
		}

		head, err := r.Head()
		if err != nil {
			return err
		}

		href := plumbing.NewHashReference(plumbing.ReferenceName("refs/heads/"+branch), head.Hash())
		err = r.Storer.SetReference(href)

		return nil
	})
	if err != nil {
		return err
	}

	err = r.Push(&git.PushOptions{
		RemoteName: remote,
		RefSpecs: []config.RefSpec{
			config.RefSpec("+refs/heads/*:refs/heads/*"),
		},
		Auth: &http.BasicAuth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		return err
	}

	return nil
}
