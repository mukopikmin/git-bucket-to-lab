package gitbucket

import (
	"io/ioutil"
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

func (c *Client) authGet(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.endpoint+path, nil)
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

	return body, nil
}

func (c *Client) authPost(path string, jsonBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.endpoint+path, strings.NewReader(string(jsonBody)))
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

	return body, err
}
