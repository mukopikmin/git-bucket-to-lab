package gitlab

import (
	"encoding/json"
	"time"
)

// Group ...
type Group struct {
	ID                             int         `json:"id"`
	Name                           string      `json:"name"`
	Path                           string      `json:"path"`
	Description                    string      `json:"description"`
	Visibility                     string      `json:"visibility"`
	ShareWithGroupLock             bool        `json:"share_with_group_lock"`
	RequireTwoFactorAuthentication bool        `json:"require_two_factor_authentication"`
	TwoFactorGracePeriod           int         `json:"two_factor_grace_period"`
	ProjectCreationLevel           string      `json:"project_creation_level"`
	AutoDevopsEnabled              interface{} `json:"auto_devops_enabled"`
	SubgroupCreationLevel          string      `json:"subgroup_creation_level"`
	EmailsDisabled                 interface{} `json:"emails_disabled"`
	MentionsDisabled               interface{} `json:"mentions_disabled"`
	LfsEnabled                     bool        `json:"lfs_enabled"`
	DefaultBranchProtection        int         `json:"default_branch_protection"`
	AvatarURL                      string      `json:"avatar_url"`
	WebURL                         string      `json:"web_url"`
	RequestAccessEnabled           bool        `json:"request_access_enabled"`
	FullName                       string      `json:"full_name"`
	FullPath                       string      `json:"full_path"`
	FileTemplateProjectID          int         `json:"file_template_project_id"`
	ParentID                       interface{} `json:"parent_id"`
	CreatedAt                      time.Time   `json:"created_at"`
}

// GetAuthorizedGroups ...
func (c *Client) GetAuthorizedGroups() ([]Group, error) {
	groups := make([]Group, 0)
	for i := range make([]int, c.maxPage) {
		body, total, err := c.authGet("/groups")
		if err != nil {
			return nil, err
		}

		var _groups []Group
		if err = json.Unmarshal([]byte(body), &_groups); err != nil {
			return nil, err
		}

		groups = append(groups, _groups...)

		if total == i+1 {
			break
		}
	}

	return groups, nil
}

// GetGroup ...
func (c *Client) GetGroup(name string) (*Group, error) {
	body, _, err := c.authGet("/groups/" + name)
	if err != nil {
		return nil, err
	}

	var group Group
	if err = json.Unmarshal([]byte(body), &group); err != nil {
		return nil, err
	}

	return &group, nil
}

// IsPrivate ...
func (g *Group) IsPrivate() bool {
	return g.Visibility == "private"
}
