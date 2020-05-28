package gitbucket

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Repo in GitBucket
type Repo struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	Watchers      int    `json:"watchers"`
	Forks         int    `json:"forks"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`
	Owner         User   `json:"owner"`
	ID            int    `json:"id"`
	ForksCount    int    `json:"forks_count"`
	WatchersCount int    `json:"watchers_count"`
	URL           string `json:"url"`
	HTTPURL       string `json:"http_url"`
	CloneURL      string `json:"clone_url"`
	HTMLURL       string `json:"html_url"`
	Issues        []Issue
	Pulls         []Pull
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
	url := c.endpoint + "/api/v3/user/repos"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
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
	url := c.endpoint + "/api/v3/repos/" + owner + "/" + name
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var repo Repo
	if err = json.Unmarshal([]byte(body), &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

// CreateRepo ...
func (c *Client) CreateRepo(name string, description string, private bool) (*Repo, error) {
	url := c.endpoint + "/api/v3/user/repos"
	repoReq := RepoRequest{name, description, private}

	jsonBody, err := json.Marshal(repoReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var repo Repo
	if err = json.Unmarshal([]byte(body), &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}
