package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client for GitLab
type Client struct {
	endpoint string
	token    string
	*http.Client
}

type errorsResult struct {
	Errors []string `json:"message"`
}

// NewClient is constructor for client
func NewClient(endpoint string, token string) *Client {
	return &Client{endpoint, token, http.DefaultClient}
}

func (c *Client) authGet(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.endpoint+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error with status: %d", res.StatusCode)
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 201 {
		return nil, fmt.Errorf("error with status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
