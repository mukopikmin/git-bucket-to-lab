package gitlab

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage"
)

// Project ...
type Project struct {
	ID                int         `json:"id"`
	Description       interface{} `json:"description"`
	DefaultBranch     string      `json:"default_branch"`
	SSHURLToRepo      string      `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string      `json:"http_url_to_repo"`
	WebURL            string      `json:"web_url"`
	ReadmeURL         string      `json:"readme_url"`
	TagList           []string    `json:"tag_list"`
	Name              string      `json:"name"`
	NameWithNamespace string      `json:"name_with_namespace"`
	Path              string      `json:"path"`
	PathWithNamespace string      `json:"path_with_namespace"`
	CreatedAt         time.Time   `json:"created_at"`
	LastActivityAt    time.Time   `json:"last_activity_at"`
	ForksCount        int         `json:"forks_count"`
	AvatarURL         string      `json:"avatar_url"`
	StarCount         int         `json:"star_count"`
	Issues            []Issue     `json:"issues"`
	Merges            []Merge     `json:"merges"`
	Branches          []Branch    `json:"branches"`
}

// ProjectRequest ...
type ProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetProjects ...
func (c *Client) GetProjects() ([]Project, error) {
	var projects []Project
	for i := range make([]int, c.maxPage) {
		path := fmt.Sprintf("/projects?per_page=%d&page=%d", c.perPage, i+1)
		body, total, err := c.authGet(path)
		if err != nil {
			return nil, err
		}

		var _projects []Project
		if err = json.Unmarshal([]byte(body), &_projects); err != nil {
			return nil, err
		}

		projects = append(projects, _projects...)

		if total == i+1 {
			break
		}
	}

	return projects, nil
}

// GetProject ...
func (c *Client) GetProject(owner string, name string) (*Project, error) {
	path := "/projects/" + owner + "%2F" + name
	body, _, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	var project Project
	if err = json.Unmarshal([]byte(body), &project); err != nil {
		return nil, err
	}

	branches, err := c.GetBranches(&project)
	if err != nil {
		return nil, err
	}

	project.Branches = branches

	return &project, nil
}

// CreateProject ...
func (c *Client) CreateProject(name string, description string) (*Project, error) {
	proReq := ProjectRequest{name, description}
	jsonBody, err := json.Marshal(proReq)
	if err != nil {
		return nil, err
	}

	body, err := c.authPost("/projects", jsonBody)
	if err != nil {
		return nil, err
	}

	var project Project
	if err = json.Unmarshal([]byte(body), &project); err != nil {
		return nil, err
	}

	return &project, nil
}

// Push ...
func (p *Project) Push(storage storage.Storer, worktree billy.Filesystem, token string) error {
	remote := "lab"
	r, err := git.Open(storage, worktree)
	if err != nil {
		return err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: remote,
		URLs: []string{p.HTTPURLToRepo},
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
			Username: "oauth2",
			Password: token,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
