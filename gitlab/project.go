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
	ID            int         `json:"id"`
	Description   interface{} `json:"description"`
	DefaultBranch string      `json:"default_branch"`
	Visibility    string      `json:"visibility"`
	SSHURLToRepo  string      `json:"ssh_url_to_repo"`
	HTTPURLToRepo string      `json:"http_url_to_repo"`
	WebURL        string      `json:"web_url"`
	ReadmeURL     string      `json:"readme_url"`
	TagList       []string    `json:"tag_list"`
	Owner         struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"owner"`
	Name                           string    `json:"name"`
	NameWithNamespace              string    `json:"name_with_namespace"`
	Path                           string    `json:"path"`
	PathWithNamespace              string    `json:"path_with_namespace"`
	IssuesEnabled                  bool      `json:"issues_enabled"`
	OpenIssuesCount                int       `json:"open_issues_count"`
	MergeRequestsEnabled           bool      `json:"merge_requests_enabled"`
	JobsEnabled                    bool      `json:"jobs_enabled"`
	WikiEnabled                    bool      `json:"wiki_enabled"`
	SnippetsEnabled                bool      `json:"snippets_enabled"`
	CanCreateMergeRequestIn        bool      `json:"can_create_merge_request_in"`
	ResolveOutdatedDiffDiscussions bool      `json:"resolve_outdated_diff_discussions"`
	ContainerRegistryEnabled       bool      `json:"container_registry_enabled"`
	CreatedAt                      time.Time `json:"created_at"`
	LastActivityAt                 time.Time `json:"last_activity_at"`
	CreatorID                      int       `json:"creator_id"`
	Namespace                      struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Path     string `json:"path"`
		Kind     string `json:"kind"`
		FullPath string `json:"full_path"`
	} `json:"namespace"`
	ImportStatus                              string        `json:"import_status"`
	Archived                                  bool          `json:"archived"`
	AvatarURL                                 string        `json:"avatar_url"`
	SharedRunnersEnabled                      bool          `json:"shared_runners_enabled"`
	ForksCount                                int           `json:"forks_count"`
	StarCount                                 int           `json:"star_count"`
	RunnersToken                              string        `json:"runners_token"`
	CiDefaultGitDepth                         int           `json:"ci_default_git_depth"`
	PublicJobs                                bool          `json:"public_jobs"`
	SharedWithGroups                          []interface{} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool          `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               bool          `json:"allow_merge_on_skipped_pipeline"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool          `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool          `json:"remove_source_branch_after_merge"`
	RequestAccessEnabled                      bool          `json:"request_access_enabled"`
	MergeMethod                               string        `json:"merge_method"`
	AutocloseReferencedIssues                 bool          `json:"autoclose_referenced_issues"`
	SuggestionCommitMessage                   interface{}   `json:"suggestion_commit_message"`
	MarkedForDeletionAt                       string        `json:"marked_for_deletion_at"`
	MarkedForDeletionOn                       string        `json:"marked_for_deletion_on"`
	Statistics                                struct {
		CommitCount      int `json:"commit_count"`
		StorageSize      int `json:"storage_size"`
		RepositorySize   int `json:"repository_size"`
		WikiSize         int `json:"wiki_size"`
		LfsObjectsSize   int `json:"lfs_objects_size"`
		JobArtifactsSize int `json:"job_artifacts_size"`
		PackagesSize     int `json:"packages_size"`
	} `json:"statistics"`
	Links struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
	} `json:"_links"`
	Issues   []Issue  `json:"issues"`
	Merges   []Merge  `json:"merges"`
	Branches []Branch `json:"branches"`
	Tags     []Tag    `json:"tags"`
}

// ProjectRequest ...
type ProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	NamespaceID int    `json:"namespace_id"`
}

// GetProjects ...
func (c *Client) GetProjects() ([]Project, error) {
	projects := make([]Project, 0)
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

// CreateGroupProject ...
func (c *Client) CreateGroupProject(group *Group, name string, description string) (*Project, error) {
	proReq := ProjectRequest{name, description, group.ID}
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

// CreateProject ...
func (c *Client) CreateProject(name string, description string) (*Project, error) {
	proReq := ProjectRequest{}
	proReq.Name = name
	proReq.Description = description

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
			config.RefSpec("+refs/tags/*:refs/tags/*"),
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
