package gitlab

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

// Merge ...
type Merge struct {
	ID             int         `json:"id"`
	Iid            int         `json:"iid"`
	ProjectID      int         `json:"project_id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	State          string      `json:"state"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	MergedBy       interface{} `json:"merged_by"`
	MergedAt       interface{} `json:"merged_at"`
	ClosedBy       interface{} `json:"closed_by"`
	ClosedAt       interface{} `json:"closed_at"`
	TargetBranch   string      `json:"target_branch"`
	SourceBranch   string      `json:"source_branch"`
	UserNotesCount int         `json:"user_notes_count"`
	Upvotes        int         `json:"upvotes"`
	Downvotes      int         `json:"downvotes"`
	Assignee       interface{} `json:"assignee"`
	Author         struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	Assignees                 []interface{} `json:"assignees"`
	SourceProjectID           int           `json:"source_project_id"`
	TargetProjectID           int           `json:"target_project_id"`
	Labels                    []interface{} `json:"labels"`
	WorkInProgress            bool          `json:"work_in_progress"`
	Milestone                 interface{}   `json:"milestone"`
	MergeWhenPipelineSucceeds bool          `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string        `json:"merge_status"`
	Sha                       string        `json:"sha"`
	MergeCommitSha            interface{}   `json:"merge_commit_sha"`
	SquashCommitSha           interface{}   `json:"squash_commit_sha"`
	DiscussionLocked          interface{}   `json:"discussion_locked"`
	ShouldRemoveSourceBranch  interface{}   `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool          `json:"force_remove_source_branch"`
	Reference                 string        `json:"reference"`
	References                struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	WebURL    string `json:"web_url"`
	TimeStats struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Squash               bool `json:"squash"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasConflicts                bool      `json:"has_conflicts"`
	BlockingDiscussionsResolved bool      `json:"blocking_discussions_resolved"`
	Comments                    []Comment `json:"comments"`
}

// MergeRequest ...
type MergeRequest struct {
	SourceBranch string `json:"source_branch"`
	TargetBrach  string `json:"target_branch"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	StateEvent   string `json:"state_event"`
}

// GetMerges ...
func (c *Client) GetMerges(p *Project) ([]Merge, error) {
	path := fmt.Sprintf("/projects/%d/merge_requests", p.ID)
	body, err := c.authGet(path)

	var _merges []Merge
	if err = json.Unmarshal([]byte(body), &_merges); err != nil {
		return nil, err
	}

	var merges []Merge
	for _, m := range _merges {
		comments, err := c.GetMergeComments(p, &m)
		if err != nil {
			return nil, err
		}

		m.Comments = comments
		merges = append(merges, m)
	}

	sort.Slice(merges, func(i, j int) bool {
		return merges[i].Iid < merges[j].Iid
	})

	return merges, nil
}

// CreateMerge ...
func (c *Client) CreateMerge(p *Project, title string, sb string, tb string, description string) (*Merge, error) {
	path := fmt.Sprintf("/projects/%d/merge_requests", p.ID)
	mergeReq := &MergeRequest{sb, tb, title, description, "close"}
	jsonBody, err := json.Marshal(mergeReq)
	if err != nil {
		return nil, err
	}

	body, err := c.authPost(path, jsonBody)
	if err != nil {
		return nil, err
	}

	var merge Merge
	if err = json.Unmarshal([]byte(body), &merge); err != nil {
		return nil, err
	}

	return &merge, nil
}
