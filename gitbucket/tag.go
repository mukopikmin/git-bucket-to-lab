package gitbucket

import (
	"strings"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Tag ...
type Tag struct {
	Name string `json:"name"`
	Sha  string `json:"sha"`
}

// GetTags ...
func (c *Client) GetTags(repo *Repo, user string) ([]Tag, error) {
	storage := memory.NewStorage()
	worktree := memfs.New()

	r, err := git.Clone(storage, worktree, &git.CloneOptions{
		URL: repo.CloneURL,
		Auth: &http.BasicAuth{
			Username: user,
			Password: c.apikey,
		},
	})
	if err != nil {
		return nil, err
	}

	tagrefs, err := r.Tags()
	tags := make([]Tag, 0)
	err = tagrefs.ForEach(func(ref *plumbing.Reference) error {
		n := strings.Split(ref.Name().String(), "/")
		tag := Tag{n[len(n)-1], ref.Hash().String()}
		tags = append(tags, tag)

		return nil
	})

	return tags, nil
}
