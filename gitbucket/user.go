package gitbucket

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// User in GitBucket
type User struct {
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Type      string    `json:"type"`
	SiteAdmin bool      `json:"site_admin"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	HTMLURL   string    `json:"html_url"`
	AvatarURL string    `json:"avatar_url"`
}

// UserRequest is request parameter to create user
type UserRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

// GetUsers ...
func (c *Client) GetUsers() ([]User, error) {
	url := c.endpoint + "/api/v3/admin/users"
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

	var repos []User
	if err = json.Unmarshal([]byte(body), &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

// func (c *Client) CreateUser(login string, email string) (*User, error) {
// 	url := c.endpoint + "/api/v3/admin/users"
// 	userReq := UserRequest{login, email}

// 	jsonBody, err := json.Marshal(userReq)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Authorization", "token "+c.apikey)

// 	res, err := c.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var user User
// 	if err = json.Unmarshal([]byte(body), &user); err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
