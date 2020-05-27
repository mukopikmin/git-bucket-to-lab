package gitlab

import "net/http"

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
