package gitbucket

import (
	"encoding/json"
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
	body, err := c.authGet("/admin/users")
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	if err = json.Unmarshal([]byte(body), &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetGroups ...
func (c *Client) GetGroups() ([]User, error) {
	body, err := c.authGet("/user/orgs")
	if err != nil {
		return nil, err
	}

	var orgs []User
	if err = json.Unmarshal([]byte(body), &orgs); err != nil {
		return nil, err
	}

	return orgs, nil
}

// GetAuthorizedUser ...
func (c *Client) GetAuthorizedUser() (*User, error) {
	body, err := c.authGet("/user")
	if err != nil {
		return nil, err
	}

	var user User
	if err = json.Unmarshal([]byte(body), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// func (c *Client) CreateUser(login string, email string) (*User, error) {
// 	url := c.Endpoint + "/admin/users"
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
