package gitbucket

import (
	"encoding/json"
	"fmt"
	"time"
)

// Comment ...
type Comment struct {
	ID   int `json:"id"`
	User struct {
		Login     string    `json:"login"`
		Email     string    `json:"email"`
		Type      string    `json:"type"`
		SiteAdmin bool      `json:"site_admin"`
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
		URL       string    `json:"url"`
		HTMLURL   string    `json:"html_url"`
		AvatarURL string    `json:"avatar_url"`
	} `json:"user"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	HTMLURL   string    `json:"html_url"`
}

// GetComments ...
func (c *Client) GetComments(r *Repo, id int) ([]Comment, error) {
	path := fmt.Sprintf("/api/v3/repos/%s/issues/%d/comments", r.FullName, id)
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	var comments []Comment
	if err = json.Unmarshal([]byte(body), &comments); err != nil {
		return nil, err
	}

	return comments, nil
}
