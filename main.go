package main

import (
	"flag"
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/handler"
	"os"
	"strings"

	"github.com/bxcodec/faker/v3"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
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
	e.Use(middleware.CORS())

	e.GET("/api", handler.Index)
	e.GET("/api/auth", handler.Auth)
	e.GET("/api/:owner/:name", handler.ShowRepo)
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

func generateFixtures() error {
	fmt.Println("Generating fixture data in GitBucket ...")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), os.Getenv("GITBUCKET_TOKEN"))
	storage := memory.NewStorage()
	worktree := memfs.New()

	name := faker.Word()
	repo, err := b.CreateRepo(name, faker.Sentence(), false)
	if err != nil {
		return err
	}

	err = clone(storage, worktree, "https://github.com/mukopikmin/git-bucket-to-lab.git")
	if err != nil {
		return err
	}

	err = push(storage, worktree, repo.CloneURL)
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

func clone(storage *memory.Storage, worktree billy.Filesystem, url string) error {
	r, err := git.Clone(storage, worktree, &git.CloneOptions{
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

func push(storage *memory.Storage, worktree billy.Filesystem, url string) error {
	remote := "bucket"
	r, err := git.Open(storage, worktree)
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
			config.RefSpec("+refs/tags/*:refs/tags/*"),
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
