package gitbucket

import "time"

// Comment ...
type Comment struct {
	ID   int `json:"id"`
	User struct {
		Login     string    `json:"login"`
		Email     string    `json:"email"`
		Type      string    `json:"type"`
		SiteAdmin bool      `json:"site_admin"`
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
		URL       string    `json:"url"`
		HTMLURL   string    `json:"html_url"`
		AvatarURL string    `json:"avatar_url"`
	} `json:"user"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	HTMLURL   string    `json:"html_url"`
}
