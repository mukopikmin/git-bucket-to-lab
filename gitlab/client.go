package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Client for GitLab
type Client struct {
	Endpoint   string
	apiversion string
	token      string
	maxPage    int
	perPage    int
	*http.Client
}

// APIEndpoint ...
func (c *Client) APIEndpoint() string {
	return c.Endpoint + "/api/" + c.apiversion
}

type errorsResult struct {
	Errors []string `json:"message"`
}

// NewClient is constructor for client
func NewClient(endpoint string, token string) *Client {
	return &Client{endpoint, "v4", token, 100, 50, http.DefaultClient}
}

func (c *Client) authGet(path string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", c.APIEndpoint()+path, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.Do(req)
	if err != nil {
		return nil, 0, err
	}
	if res.StatusCode != 200 {
		return nil, 0, fmt.Errorf("error with status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 1, err
	}

	total, err := strconv.Atoi(res.Header.Get("X-TOTAL-PAGES"))
	if err != nil {
		return body, 0, nil
	}

	return body, total, nil
}

func (c *Client) authPost(path string, jsonBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.APIEndpoint()+path, strings.NewReader(string(jsonBody)))
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

func (c *Client) authPut(path string, jsonBody []byte) ([]byte, error) {
	req, err := http.NewRequest("`PUT", c.APIEndpoint()+path, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 204 {
		return nil, fmt.Errorf("error with status: %d", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
