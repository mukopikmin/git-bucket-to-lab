package gitbucket

import (
	"encoding/json"
	"time"
)

// Group ...
type Group struct {
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	AvatarURL string    `json:"avatar_url"`
}

// GetAuthenticatedGroups ...
func (c *Client) GetAuthenticatedGroups() ([]Group, error) {
	path := "/user/orgs"
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	var groups []Group
	if err = json.Unmarshal([]byte(body), &groups); err != nil {
		return nil, err
	}

	return groups, nil
}
