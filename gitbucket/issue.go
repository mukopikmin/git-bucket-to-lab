package gitbucket

import (
	"encoding/json"
	"fmt"
	"sort"
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
	Comments    []Comment     `json:"comments"`
}

// IssueRequest ...
type IssueRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// GetIssues ...
func (c *Client) GetIssues(repo *Repo) ([]Issue, error) {
	issues := make([]Issue, 0)
	for _, s := range []string{"open", "closed"} {
		for i := range make([]int, c.maxPage) {
			path := fmt.Sprintf("/repos/%s/issues?state=%s&per_page=%d&page=%d", repo.FullName, s, c.perPage, i+1)
			body, err := c.authGet(path)
			if err != nil {
				return nil, err
			}

			var _issues []Issue
			if err = json.Unmarshal([]byte(body), &_issues); err != nil {
				return nil, err
			}

			if len(_issues) == 0 {
				break
			}

			for _, issue := range _issues {
				comments, err := c.GetComments(repo, issue.Number)
				if err != nil {
					return nil, err
				}
				issue.Comments = comments
				issues = append(issues, issue)
			}
		}
	}

	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Number < issues[j].Number
	})

	return issues, nil
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

// MigratedBody ...
func (i *Issue) MigratedBody() string {
	format := "2006/1/2 15:04:05"
	prefix := fmt.Sprintf(`> This issue is migrated from [#%d](%s).  
> Original author: %s  

`, i.Number, i.HTMLURL, i.User.Login, i.CreatedAt.Format(format), i.UpdatedAt.Format(format))

	return prefix + i.Body
}
