package gitlab

import (
	"encoding/json"
	"fmt"
)

// Branch ...
type Branch struct {
	Name               string `json:"name"`
	Merged             bool   `json:"merged"`
	Protected          bool   `json:"protected"`
	Default            bool   `json:"default"`
	DevelopersCanPush  bool   `json:"developers_can_push"`
	DevelopersCanMerge bool   `json:"developers_can_merge"`
	CanPush            bool   `json:"can_push"`
	WebURL             string `json:"web_url"`
	Commit             struct {
		AuthorEmail    string   `json:"author_email"`
		AuthorName     string   `json:"author_name"`
		AuthoredDate   string   `json:"authored_date"`
		CommittedDate  string   `json:"committed_date"`
		CommitterEmail string   `json:"committer_email"`
		CommitterName  string   `json:"committer_name"`
		ID             string   `json:"id"`
		ShortID        string   `json:"short_id"`
		Title          string   `json:"title"`
		Message        string   `json:"message"`
		ParentIds      []string `json:"parent_ids"`
	} `json:"commit"`
}

// GetBranches ...
func (c *Client) GetBranches(p *Project) ([]Branch, error) {
	path := fmt.Sprintf("/projects/%d/repository/branches", p.ID)
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	var branches []Branch
	if err = json.Unmarshal([]byte(body), &branches); err != nil {
		return nil, err
	}

	return branches, nil
}
