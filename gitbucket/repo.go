package gitbucket

import (
	"encoding/json"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage"
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
	Tags          []Tag    `json:"tags"`
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
	groups, err := c.GetAuthenticatedGroups()
	if err != nil {
		return nil, err
	}

	repos := make([]Repo, 0)
	for _, g := range groups {
		body, err := c.authGet("/orgs/" + g.Login + "/repos")
		if err != nil {
			return nil, err
		}

		_repos := make([]Repo, 0)
		if err = json.Unmarshal([]byte(body), &_repos); err != nil {
			return nil, err
		}

		repos = append(repos, _repos...)
	}

	body, err := c.authGet("/user/repos")
	if err != nil {
		return nil, err
	}

	_repos := make([]Repo, 0)
	if err = json.Unmarshal([]byte(body), &_repos); err != nil {
		return nil, err
	}

	repos = append(repos, _repos...)

	return repos, nil
}

// GetRepo ...
func (c *Client) GetRepo(owner string, name string) (*Repo, error) {
	path := "/repos/" + owner + "/" + name
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	var repo Repo
	if err = json.Unmarshal([]byte(body), &repo); err != nil {
		return nil, err
	}

	branches, err := c.GetBranches(&repo)
	if err != nil {
		return nil, err
	}

	tags, err := c.GetTags(&repo, owner)
	if err != nil {
		return nil, err
	}

	issues, err := c.GetIssues(&repo)
	if err != nil {
		return nil, err
	}

	pulls, err := c.GetPulls(&repo)
	if err != nil {
		return nil, err
	}

	repo.Branches = branches
	repo.Tags = tags
	repo.Issues = issues
	repo.Pulls = pulls

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
func (repo *Repo) Clone(storage storage.Storer, worktree billy.Filesystem, user string, token string) error {
	r, err := git.Clone(storage, worktree, &git.CloneOptions{
		URL: repo.CloneURL,
		Auth: &http.BasicAuth{
			Username: user,
			Password: token,
		},
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
