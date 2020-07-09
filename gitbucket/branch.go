package gitbucket

import (
	"encoding/json"
	"fmt"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage"
)

// Branch ...
type Branch struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	Protected  bool `json:"protected"`
	Protection struct {
		Enabled              bool `json:"enabled"`
		RequiredStatusChecks struct {
			EnforcementLevel string   `json:"enforcement_level"`
			Contexts         []string `json:"contexts"`
		} `json:"required_status_checks"`
	} `json:"protection"`
	ProtectionURL string `json:"protection_url"`
}

// GetBranches ...
func (c *Client) GetBranches(r *Repo) ([]Branch, error) {
	path := "/repos/" + r.FullName + "/branches"
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	branches := make([]Branch, 0)
	if err = json.Unmarshal([]byte(body), &branches); err != nil {
		return nil, err
	}

	return branches, nil
}

// CreateBranch ...
func (c *Client) CreateBranch(repo *git.Repository, name string, hash string, storage storage.Storer, worktree billy.Filesystem) error {
	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Create: true,
		Hash:   plumbing.NewHash(hash),
		Branch: plumbing.ReferenceName("refs/heads/" + name),
	})
	if err != nil {
		// Ignore error
		// Assume most checkout error is "already exists", and it is not a problem
		fmt.Println(err)
	}

	return nil
}
