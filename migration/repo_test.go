package migration

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
	"testing"
)

func TestIsRepoMigratable(t *testing.T) {
	labUser := gitlab.User{Username: "root"}
	labGroups := []gitlab.Group{
		gitlab.Group{Name: "public_group", Visibility: "public"},
		gitlab.Group{Name: "private_group", Visibility: "private"},
	}

	owner := gitbucket.User{Login: "public_group"}
	repo := gitbucket.Repo{Name: "test", Owner: owner, Private: false}
	m := Migration{Repo: &repo}

	if !m.isRepoMigratable(&labUser, labGroups) {
		t.Fatal("isRepoMigratable: public repo to public project for public group")
	}

	owner = gitbucket.User{Login: "private_group"}
	repo = gitbucket.Repo{Name: "test", Owner: owner, Private: false}
	m = Migration{Repo: &repo}

	if m.isRepoMigratable(&labUser, labGroups) {
		t.Fatal("isRepoMigratable: public repo to public project for private group")
	}
}
