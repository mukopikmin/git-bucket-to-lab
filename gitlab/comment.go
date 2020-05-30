package gitlab

import (
	"encoding/json"
	"fmt"
	"time"
)

// Comment ...
type Comment struct {
	ID         int         `json:"id"`
	Type       interface{} `json:"type"`
	Body       string      `json:"body"`
	Attachment interface{} `json:"attachment"`
	Author     struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	System          bool      `json:"system"`
	NoteableID      int       `json:"noteable_id"`
	NoteableType    string    `json:"noteable_type"`
	Resolvable      bool      `json:"resolvable"`
	Confidential    bool      `json:"confidential"`
	NoteableIid     int       `json:"noteable_iid"`
	CommandsChanges struct {
	} `json:"commands_changes"`
	HTMLURL string
}

// CommentRequest ...
type CommentRequest struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// GetComments ...
func (c *Client) GetComments(p *Project, i *Issue) ([]Comment, error) {
	path := fmt.Sprintf("/projects/%d/issues/%d/notes", p.ID, i.Iid)
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	var _comments []Comment
	if err = json.Unmarshal([]byte(body), &_comments); err != nil {
		return nil, err
	}

	var comments []Comment
	for _, comment := range _comments {
		comment.HTMLURL = fmt.Sprintf("%s/%s/-/issues/%d#note_%d", c.Endpoint, p.PathWithNamespace, i.Iid, comment.ID)
		comments = append(comments, comment)
	}

	return comments, nil
}

// CreateComment ...
func (c *Client) CreateComment(p *Project, i *Issue, message string, timestamp time.Time) (*Comment, error) {
	reqBody := &CommentRequest{message, timestamp}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/projects/%d/issues/%d/notes", p.ID, i.Iid)
	body, err := c.authPost(path, jsonBody)
	if err != nil {
		return nil, err
	}

	var comment Comment
	if err = json.Unmarshal([]byte(body), &comment); err != nil {
		return nil, err
	}

	return &comment, nil
}
