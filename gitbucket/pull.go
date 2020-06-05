package gitbucket

import (
	"encoding/json"
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
	Comments          []Comment
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
	path := "/repos/" + repo.FullName + "/pulls"
	body, err := c.authGet(path)
	if err != nil {
		return nil, err
	}

	var _pulls []Pull
	if err = json.Unmarshal([]byte(body), &_pulls); err != nil {
		return nil, err
	}

	var pulls []Pull
	for _, pull := range _pulls {
		comments, err := c.GetComments(repo, pull.Number)
		if err != nil {
			return nil, err
		}

		pull.Comments = comments
		pulls = append(pulls, pull)
	}

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

	var pull Pull
	if err = json.Unmarshal([]byte(resBody), &pull); err != nil {
		return nil, err
	}

	return &pull, nil
}
