package gitbucket

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client for GitBucket
type Client struct {
	Endpoint   string
	apiversion string
	apikey     string
	maxPage    int
	perPage    int
	*http.Client
}

// APIEndpoint ...
func (c *Client) APIEndpoint() string {
	return c.Endpoint + "/api/" + c.apiversion
}

// NewClient is constructor fot client
func NewClient(endpoint string, apikey string) *Client {
	return &Client{endpoint, "v3", apikey, 100, 50, http.DefaultClient}
}

func (c *Client) authGet(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.APIEndpoint()+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+c.apikey)

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
	req, err := http.NewRequest("POST", c.APIEndpoint()+path, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+c.apikey)

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

	return body, err
}
