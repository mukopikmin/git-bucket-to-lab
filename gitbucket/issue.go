package gitbucket

import (
	"encoding/json"
	"time"
)

// Issue ...
type Issue struct {
	Number      int           `json:"number"`
	Title       string        `json:"title"`
	User        User          `json:"user"`
	Labels      []interface{} `json:"labels"`
	State       string        `json:"state"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Body        string        `json:"body"`
	ID          int           `json:"id"`
	CommentsURL string        `json:"comments_url"`
	HTMLURL     string        `json:"html_url"`
	Comments    []Comment
}

// IssueRequest ...
type IssueRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// GetIssues ...
func (c *Client) GetIssues(repo *Repo) ([]Issue, error) {
	body, err := c.authGet("/repos/" + repo.FullName + "/issues")
	if err != nil {
		return nil, err
	}

	var issues []Issue
	if err = json.Unmarshal([]byte(body), &issues); err != nil {
		return nil, err
	}

	var commentIssues []Issue
	for _, issue := range issues {
		comments, err := c.GetComments(repo, issue.Number)
		if err != nil {
			return nil, err
		}
		issue.Comments = comments
		commentIssues = append(commentIssues, issue)
	}

	return commentIssues, nil
}

// CreateIssue ...
func (c *Client) CreateIssue(repo *Repo, title string, body string) (*Issue, error) {
	path := "/repos/" + repo.FullName + "/issues"
	issueReq := IssueRequest{title, body}

	jsonBody, err := json.Marshal(issueReq)
	if err != nil {
		return nil, err
	}

	resBody, err := c.authPost(path, jsonBody)
	if err != nil {
		return nil, err
	}

	var issue Issue
	if err = json.Unmarshal([]byte(resBody), &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
