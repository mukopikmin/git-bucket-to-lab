package gitbucket

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Pull ...
type Pull struct {
	Number    int       `json:"number"`
	State     string    `json:"state"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Head      struct {
		Sha   string `json:"sha"`
		Ref   string `json:"ref"`
		Repo  Repo   `json:"repo"`
		Label string `json:"label"`
		User  User   `json:"user"`
	} `json:"head"`
	Base struct {
		Sha   string `json:"sha"`
		Ref   string `json:"ref"`
		Repo  Repo   `json:"repo"`
		Label string `json:"label"`
		User  User   `json:"user"`
	} `json:"base"`
	Merged            bool          `json:"merged"`
	Title             string        `json:"title"`
	Body              string        `json:"body"`
	User              User          `json:"user"`
	Labels            []interface{} `json:"labels"`
	ID                int           `json:"id"`
	HTMLURL           string        `json:"html_url"`
	URL               string        `json:"url"`
	CommitsURL        string        `json:"commits_url"`
	ReviewCommentsURL string        `json:"review_comments_url"`
	ReviewCommentURL  string        `json:"review_comment_url"`
	CommentsURL       string        `json:"comments_url"`
	StatusesURL       string        `json:"statuses_url"`
	Comments          []Comment     `json:"comments"`
}

// PullRequest ...
type PullRequest struct {
	Title string `json:"title"`
	Head  string `json:"head"`
	Base  string `json:"base"`
	Body  string `json:"body"`
}

// GetPulls ...
func (c *Client) GetPulls(repo *Repo) ([]Pull, error) {
	pulls := make([]Pull, 0)
	for _, s := range []string{"open", "closed"} {
		for i := range make([]int, c.maxPage) {
			path := fmt.Sprintf("/repos/%s/pulls?state=%s&per_page=%d&page=%d", repo.FullName, s, c.perPage, i+1)
			body, err := c.authGet(path)
			if err != nil {
				return nil, err
			}

			var _pulls []Pull
			if err = json.Unmarshal([]byte(body), &_pulls); err != nil {
				return nil, err
			}

			if len(_pulls) == 0 {
				break
			}

			for _, pull := range _pulls {
				comments, err := c.GetComments(repo, pull.Number)
				if err != nil {
					return nil, err
				}

				pull.Comments = comments
				pulls = append(pulls, pull)
			}
		}
	}

	sort.Slice(pulls, func(i, j int) bool {
		return pulls[i].Number < pulls[j].Number
	})

	return pulls, nil
}

// CreatePull ...
func (c *Client) CreatePull(repo *Repo, title string, head string, base string, body string) (*Pull, error) {
	path := "/repos/" + repo.FullName + "/pulls"

	reqBody := PullRequest{title, head, base, body}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resBody, err := c.authPost(path, jsonBody)
	if err != nil {
		return nil, err
	}

	// Patch to avoid wrong response of GitBucket API
	// GitBucket returns response as escaped JSON string
	// Need to remove escaped character and double quotes at edge of string
	// Ref. https://github.com/gitbucket/gitbucket/issues/2306
	fixedJSON := strings.ReplaceAll(fmt.Sprintf("%s", resBody), "\\", "")
	fixedJSON = strings.TrimLeft(fixedJSON, "\"")
	fixedJSON = strings.TrimRight(fixedJSON, "\"")

	var pull Pull
	if err = json.Unmarshal([]byte([]byte(fixedJSON)), &pull); err != nil {
		return nil, err
	}

	return &pull, nil
}

// MigratedPullBody ...
func (c *Client) MigratedPullBody(p *Pull, r *Repo) (*string, error) {
	prefix := fmt.Sprintf(`> This merge request is migrated from [#%d@GitBucket](%s).    
> Original author: %s  

`, p.Number, p.HTMLURL, p.User.Login)

	if p.Merged {
		prefix += "> This merge request has been merged.\n\n"
	} else {
		prefix += "> This merge request is not merged.\n\n"
	}

	b, err := c.migratedBody(p.Body, r)
	if err != nil {
		return nil, err
	}

	body := prefix + *b

	return &body, nil
}
