package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
}

// ProjectRequest ...
type ProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetProjects ...
func (c *Client) GetProjects() ([]Project, error) {
	url := c.endpoint + "/projects"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error with status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var projects []Project
	if err = json.Unmarshal([]byte(body), &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// CreateProject ...
func (c *Client) CreateProject(name string, description string) (*Project, error) {
	url := c.endpoint + "/projects"
	proReq := ProjectRequest{name, description}

	jsonBody, err := json.Marshal(proReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 204 {
		return nil, fmt.Errorf("error with status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var project Project
	if err = json.Unmarshal([]byte(body), &project); err != nil {
		return nil, err
	}

	return &project, nil
}
