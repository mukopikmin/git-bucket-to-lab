package gitbucket

import (
	"encoding/json"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// Repo in GitBucket
type Repo struct {
	Name          string   `json:"name"`
	FullName      string   `json:"full_name"`
	Description   string   `json:"description"`
	Watchers      int      `json:"watchers"`
	Forks         int      `json:"forks"`
	Private       bool     `json:"private"`
	DefaultBranch string   `json:"default_branch"`
	Owner         User     `json:"owner"`
	ID            int      `json:"id"`
	ForksCount    int      `json:"forks_count"`
	WatchersCount int      `json:"watchers_count"`
	URL           string   `json:"url"`
	HTTPURL       string   `json:"http_url"`
	CloneURL      string   `json:"clone_url"`
	HTMLURL       string   `json:"html_url"`
	Issues        []Issue  `json:"issues"`
	Pulls         []Pull   `json:"pulls"`
	Branches      []Branch `json:"branches"`
}

// RepoRequest ...
type RepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Homepage    string `json:"homepage"`
	Private bool `json:"private"`
	// HasIssues bool `json:"has_issues"`
	// HasProjects bool   `json:"has_projects"`
	// HasWiki     bool   `json:"has_wiki"`
}

// GetRepos ...
func (c *Client) GetRepos() ([]Repo, error) {
	body, err := c.authGet("/user/repos")
	if err != nil {
		return nil, err
	}

	var repos []Repo
	if err = json.Unmarshal([]byte(body), &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

// GetRepo ...
func (c *Client) GetRepo(owner string, name string) (*Repo, error) {
	path := "/repos/" + owner + "/" + name
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	fmt.Println(path)
	fmt.Println(c.apikey)

	var repo Repo
	if err = json.Unmarshal([]byte(body), &repo); err != nil {
		return nil, err
	}

	branches, err := c.GetBranches(&repo)
	if err != nil {
		return nil, err
	}

	repo.Branches = branches

	return &repo, nil
}

// CreateRepo ...
func (c *Client) CreateRepo(name string, description string, private bool) (*Repo, error) {
	repoReq := RepoRequest{name, description, private}
	jsonBody, err := json.Marshal(repoReq)
	if err != nil {
		return nil, err
	}

	body, err := c.authPost("/user/repos", jsonBody)
	if err != nil {
		return nil, err
	}

	var repo Repo
	if err = json.Unmarshal([]byte(body), &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

// Clone ...
func (repo *Repo) Clone() error {
	r, err := git.PlainClone("tmp/"+repo.FullName, false, &git.CloneOptions{
		URL:          repo.CloneURL,
		SingleBranch: false,
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	for _, branch := range repo.Branches {
		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/remotes/origin/" + branch.Name),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
