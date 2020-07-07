package gitbucket

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
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

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
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

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}

	return body, err
}

// MigratedBody ...
func (c *Client) migratedBody(b string, r *Repo) (*string, error) {
	body := b
	if strings.Contains(b, "#") {
		re := regexp.MustCompile("#(\\d+)")
		result := re.FindAllStringSubmatch(b, -1)

		if len(result) > 0 {
			pulls, err := c.GetPulls(r)
			if err != nil {
				return nil, err
			}

			for _, r := range result {
				id, err := strconv.Atoi(r[1])
				if err != nil {
					return nil, err
				}

				for i, p := range pulls {
					if p.Number == id {
						body = strings.Replace(body, "#"+strconv.Itoa(id), "!"+strconv.Itoa(i+1), -1)
					}
				}
			}
		}
	}

	return &body, nil
}
