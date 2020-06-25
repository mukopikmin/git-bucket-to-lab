package fixture

import (
	"fmt"
	"git-bucket-to-lab/gitbucket"
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
)

// Generate ...
func Generate() error {
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

	repo, err = b.GetRepo(repo.Owner.Login, repo.Name)
	if err != nil {
		return err
	}

	for _, branch := range repo.Branches {
		if branch.Name == "master" {
			continue
		}

		p, err := b.CreatePull(repo, faker.Sentence(), branch.Name, "master", faker.Paragraph())
		if err != nil {
			return err
		}

		fmt.Printf("Created pull request : #%d\n", p.ID)
	}

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
