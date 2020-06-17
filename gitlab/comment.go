package gitlab

import (
	"encoding/json"
	"fmt"
	"sort"
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
}

// CommentRequest ...
type CommentRequest struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// GetIssueComments ...
func (c *Client) GetIssueComments(p *Project, i *Issue) ([]Comment, error) {
	return c.getComments(p, "issues", i.Iid)
}

// GetMergeComments ...
func (c *Client) GetMergeComments(p *Project, m *Merge) ([]Comment, error) {
	return c.getComments(p, "merge_requests", m.Iid)
}

func (c *Client) getComments(p *Project, target string, id int) ([]Comment, error) {
	comments := make([]Comment, 0)
	for i := range make([]int, c.maxPage) {
		path := fmt.Sprintf("/projects/%d/%s/%d/notes?per_page=%d&page=%d", p.ID, target, id, c.perPage, i+1)
		body, total, err := c.authGet(path)
		if err != nil {
			return nil, err
		}

		var _comments []Comment
		if err = json.Unmarshal([]byte(body), &_comments); err != nil {
			return nil, err
		}

		comments = append(comments, _comments...)

		if total == i+1 {
			break
		}
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].ID < comments[j].ID
	})

	return comments, nil
}

// CreateIssueComment ...
func (c *Client) CreateIssueComment(p *Project, i *Issue, message string, timestamp time.Time) (*Comment, error) {
	return c.createComment(p, "issues", i.Iid, message, timestamp)
}

// CreateMergeComment ...
func (c *Client) CreateMergeComment(p *Project, m *Merge, message string, timestamp time.Time) (*Comment, error) {
	return c.createComment(p, "merge_requests", m.Iid, message, timestamp)
}

func (c *Client) createComment(p *Project, target string, id int, message string, timestamp time.Time) (*Comment, error) {
	reqBody := &CommentRequest{message, timestamp}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/projects/%d/%s/%d/notes", p.ID, target, id)
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
