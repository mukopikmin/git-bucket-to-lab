package fixture

import (
	"fmt"
	"git-bucket-to-lab/gitbucket"
	"os"

	"github.com/bxcodec/faker/v3"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Generate ...
func Generate() error {
	fmt.Println("Generating fixture data in GitBucket ...")

	b := gitbucket.NewClient(os.Getenv("GITBUCKET_URL"), os.Getenv("GITBUCKET_TOKEN"))
	storage := memory.NewStorage()
	worktree := memfs.New()

	name := faker.Word() + "-" + faker.Word()
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

		_, err = b.CreateComment(repo, issue.Number, "See #3 and #7")
		if err != nil {
			return err
		}
	}

	_, err = b.CreateIssue(repo, faker.Sentence(), "See #7")
	if err != nil {
		return err
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

		fmt.Printf("Created pull request : #%d\n", p.Number)

		for i := 0; i < 5; i++ {
			comment, err := b.CreateComment(repo, p.Number, faker.Sentence())
			if err != nil {
				return err
			}

			fmt.Printf("Created pull request comment : %d\n", comment.ID)
		}

		_, err = b.CreateComment(repo, p.Number, "See #3 and #7")
		if err != nil {
			return err
		}
	}

	_, err = b.CreatePull(repo, faker.Sentence(), repo.Branches[0].Name, "master", "Fix #2, Ref #7")
	if err != nil {
		return err
	}

	return nil
}
