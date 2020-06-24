package gitlab

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// Issue ...
type Issue struct {
	ID          int           `json:"id"`
	Iid         int           `json:"iid"`
	ProjectID   int           `json:"project_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	State       string        `json:"state"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	ClosedAt    interface{}   `json:"closed_at"`
	ClosedBy    interface{}   `json:"closed_by"`
	Labels      []interface{} `json:"labels"`
	Milestone   interface{}   `json:"milestone"`
	Assignees   []interface{} `json:"assignees"`
	Author      struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	Assignee           interface{} `json:"assignee"`
	UserNotesCount     int         `json:"user_notes_count"`
	MergeRequestsCount int         `json:"merge_requests_count"`
	Upvotes            int         `json:"upvotes"`
	Downvotes          int         `json:"downvotes"`
	DueDate            interface{} `json:"due_date"`
	Confidential       bool        `json:"confidential"`
	DiscussionLocked   interface{} `json:"discussion_locked"`
	WebURL             string      `json:"web_url"`
	TimeStats          struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasTasks bool `json:"has_tasks"`
	Links    struct {
		Self       string `json:"self"`
		Notes      string `json:"notes"`
		AwardEmoji string `json:"award_emoji"`
		Project    string `json:"project"`
	} `json:"_links"`
	References struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	MovedToID interface{} `json:"moved_to_id"`
	Comments  []Comment   `json:"comments"`
}

// IssueRequest ...
type IssueRequest struct {
	IID         int    `json:"iid"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// GetIssues ...
func (c *Client) GetIssues(p *Project) ([]Issue, error) {
	issues := make([]Issue, 0)
	for i := range make([]int, c.maxPage) {
		path := fmt.Sprintf("/projects/%d/issues?per_page=%d&page=%d", p.ID, c.perPage, i+1)
		body, total, err := c.authGet(path)
		if err != nil {
			return nil, err
		}

		var _issues []Issue
		if err = json.Unmarshal([]byte(body), &_issues); err != nil {
			return nil, err
		}

		for _, issue := range _issues {
			comments, err := c.GetIssueComments(p, &issue)
			if err != nil {
				return nil, err
			}
			issue.Comments = comments
			issues = append(issues, issue)
		}

		if total == i+1 {
			break
		}
	}

	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Iid < issues[j].Iid
	})

	return issues, nil
}

// CreateIssue ...
func (c *Client) CreateIssue(p *Project, iid int, title string, description string) (*Issue, error) {
	path := fmt.Sprintf("/projects/%d/issues", p.ID)
	issueReq := IssueRequest{iid, title, description}

	jsonBody, err := json.Marshal(issueReq)
	if err != nil {
		return nil, err
	}

	body, err := c.authPost(path, jsonBody)
	if err != nil {
		return nil, err
	}

	var issue Issue
	if err = json.Unmarshal([]byte(body), &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
