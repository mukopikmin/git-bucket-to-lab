package gitbucket

import (
	"encoding/json"
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
