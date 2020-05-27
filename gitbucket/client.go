package gitbucket

import (
	"net/http"
	"strings"
)

// Client for GitBucket
type Client struct {
	endpoint string
	apikey   string
	*http.Client
}

// NewClient is constructor fot client
func NewClient(endpoint string, apikey string) *Client {
	return &Client{endpoint, apikey, http.DefaultClient}
}

func (c *Client) adminGet(path string) (*http.Request, error) {
	url := c.endpoint + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

	return req, nil
}

func (c *Client) adminPost(path string, reader *strings.Reader) (*http.Request, error) {
	url := c.endpoint + path
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

	return req, nil
}
