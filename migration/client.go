package migration

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"os"
)

// Client ...
type Client struct {
	bucket *gitbucket.Client
	lab    *gitlab.Client
}

// NewClient ...
func NewClient(bucketToken string, labToken string) (*Client, error) {
	bucketURL := os.Getenv("GITBUCKET_URL")
	labURL := os.Getenv("GITLAB_URL")
	bucket := gitbucket.NewClient(bucketURL, bucketToken)
	lab := gitlab.NewClient(labURL, labToken)

	return &Client{bucket, lab}, nil
}
