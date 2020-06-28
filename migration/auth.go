package migration

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
)

// AuthState ...
type AuthState struct {
	BucketUser   *gitbucket.User   `json:"gitbucket_user"`
	LabUser      *gitlab.User      `json:"gitlab_user"`
	BucketGroups []gitbucket.Group `json:"gitbucket_groups"`
	LabGroups    []gitlab.Group    `json:"gitlab_groups"`
}

// GetAuthState ...
func (c *Client) GetAuthState() (*AuthState, error) {
	bucketUser, err := c.bucket.GetAuthorizedUser()
	if err != nil {
		return nil, err
	}

	bucketGroups, err := c.bucket.GetAuthorizedGroups()
	if err != nil {
		return nil, err
	}

	labUser, err := c.lab.GetAuthorizedUser()
	if err != nil {
		return nil, err
	}

	labGroups, err := c.lab.GetAuthorizedGroups()
	if err != nil {
		return nil, err
	}

	return &AuthState{bucketUser, labUser, bucketGroups, labGroups}, nil
}
