package gitbucket

import (
	"encoding/json"
	"fmt"
	"sort"
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

// CommentRequest ...
type CommentRequest struct {
	Body string `json:"body"`
}

// GetComments ...
func (c *Client) GetComments(r *Repo, id int) ([]Comment, error) {
	path := fmt.Sprintf("/repos/%s/issues/%d/comments", r.FullName, id)
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	comments := make([]Comment, 0)
	if err = json.Unmarshal([]byte(body), &comments); err != nil {
		return nil, err
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].ID < comments[j].ID
	})

	return comments, nil
}

// CreateComment ...
func (c *Client) CreateComment(r *Repo, id int, message string) (*Comment, error) {
	path := fmt.Sprintf("/repos/%s/issues/%d/comments", r.FullName, id)
	reqBody := CommentRequest{message}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resBody, err := c.authPost(path, jsonBody)
	if err != nil {
		return nil, err
	}

	var comment Comment
	if err = json.Unmarshal([]byte(resBody), &comment); err != nil {
		return nil, err
	}

	return &comment, nil
}

// MigratedBody ...
func (c *Comment) MigratedBody() string {
	format := "2006/1/2 15:04:05"
	prefix := fmt.Sprintf(`> This comment is migrated from [#%d](%s).
>
> Original author: %s  
> Original created date: %s UTC  
> Original updated date: %s UTC  

`, c.ID, c.HTMLURL, c.User.Login, c.CreatedAt.Format(format), c.UpdatedAt.Format(format))

	return prefix + c.Body
}
