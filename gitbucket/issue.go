package gitbucket

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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
	// Assignees []string `json:"assignees"`
	// Milestone int      `json:"milestone"`
	// Labels    []string `json:"labels"`
}

// GetIssues ...
func (c *Client) GetIssues(repo *Repo) ([]Issue, error) {
	url := c.endpoint + "/api/v3/repos/" + repo.FullName + "/issues"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var issues []Issue
	if err = json.Unmarshal([]byte(body), &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

// CreateIssue ...
func (c *Client) CreateIssue(repo *Repo, title string, body string) (*Issue, error) {
	url := c.endpoint + "/api/v3/repos/" + repo.FullName + "/issues"
	issueReq := IssueRequest{title, body}

	jsonBody, err := json.Marshal(issueReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var issue Issue
	if err = json.Unmarshal([]byte(resBody), &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
