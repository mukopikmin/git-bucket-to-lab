package gitlab

import (
	"encoding/json"
	"fmt"
	"time"
)

type Tag struct {
	Commit struct {
		ID             string    `json:"id"`
		ShortID        string    `json:"short_id"`
		Title          string    `json:"title"`
		CreatedAt      time.Time `json:"created_at"`
		ParentIds      []string  `json:"parent_ids"`
		Message        string    `json:"message"`
		AuthorName     string    `json:"author_name"`
		AuthorEmail    string    `json:"author_email"`
		AuthoredDate   string    `json:"authored_date"`
		CommitterName  string    `json:"committer_name"`
		CommitterEmail string    `json:"committer_email"`
		CommittedDate  string    `json:"committed_date"`
	} `json:"commit"`
	Release struct {
		TagName     string `json:"tag_name"`
		Description string `json:"description"`
	} `json:"release"`
	Name      string      `json:"name"`
	Target    string      `json:"target"`
	Message   interface{} `json:"message"`
	Protected bool        `json:"protected"`
}

// GetTags ...
func (c *Client) GetTags(p *Project) ([]Tag, error) {
	tags := make([]Tag, 0)
	for i := range make([]int, c.maxPage) {
		path := fmt.Sprintf("/projects/%d/repository/tags?per_page=%d&page=%d", p.ID, c.perPage, i+1)
		body, total, err := c.authGet(path)
		if err != nil {
			return nil, err
		}

		var _tags []Tag
		if err = json.Unmarshal([]byte(body), &_tags); err != nil {
			return nil, err
		}

		tags = append(tags, _tags...)

		if total == i+1 {
			break
		}
	}

	return tags, nil
}
