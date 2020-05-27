package gitbucket

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
}

// GetPulls ...
func (c *Client) GetPulls(repo *Repo) ([]Pull, error) {
	url := c.endpoint + "/api/v3/repos/" + repo.FullName + "/pulls"
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

	var pulls []Pull
	if err = json.Unmarshal([]byte(body), &pulls); err != nil {
		return nil, err
	}

	return pulls, nil
}
