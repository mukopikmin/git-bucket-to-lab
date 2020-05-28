package gitlab

import (
	"encoding/json"
	"time"
)

// User ...
type User struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	Username         string        `json:"username"`
	State            string        `json:"state"`
	AvatarURL        string        `json:"avatar_url"`
	WebURL           string        `json:"web_url"`
	CreatedAt        time.Time     `json:"created_at"`
	Bio              interface{}   `json:"bio"`
	Location         interface{}   `json:"location"`
	PublicEmail      string        `json:"public_email"`
	Skype            string        `json:"skype"`
	Linkedin         string        `json:"linkedin"`
	Twitter          string        `json:"twitter"`
	WebsiteURL       string        `json:"website_url"`
	Organization     interface{}   `json:"organization"`
	JobTitle         string        `json:"job_title"`
	WorkInformation  interface{}   `json:"work_information"`
	LastSignInAt     time.Time     `json:"last_sign_in_at"`
	ConfirmedAt      time.Time     `json:"confirmed_at"`
	LastActivityOn   string        `json:"last_activity_on"`
	Email            string        `json:"email"`
	ThemeID          int           `json:"theme_id"`
	ColorSchemeID    int           `json:"color_scheme_id"`
	ProjectsLimit    int           `json:"projects_limit"`
	CurrentSignInAt  time.Time     `json:"current_sign_in_at"`
	Identities       []interface{} `json:"identities"`
	CanCreateGroup   bool          `json:"can_create_group"`
	CanCreateProject bool          `json:"can_create_project"`
	TwoFactorEnabled bool          `json:"two_factor_enabled"`
	External         bool          `json:"external"`
	PrivateProfile   bool          `json:"private_profile"`
	IsAdmin          bool          `json:"is_admin"`
}

// GetSelf ...
func (c *Client) GetSelf() (*User, error) {
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
