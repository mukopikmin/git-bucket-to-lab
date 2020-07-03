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

	repo := gitbucket.Repo{
		Name:    "test",
		Owner:   gitbucket.User{Login: "root"},
		Private: false,
	}
	m := Migration{Repo: &repo}

	if !m.isRepoMigratable(&labUser, labGroups) {
		t.Fatal("test failed")
	}

	repo = gitbucket.Repo{
		Name:    "test",
		Owner:   gitbucket.User{Login: "admin"},
		Private: true,
	}
	m = Migration{Repo: &repo}

	if m.isRepoMigratable(&labUser, labGroups) {
		t.Fatal("test failed")
	}

	repo = gitbucket.Repo{
		Name:    "test",
		Owner:   gitbucket.User{Login: "public_group"},
		Private: true,
	}
	m = Migration{Repo: &repo}

	if m.isRepoMigratable(&labUser, labGroups) {
		t.Fatal("test failed")
	}

	repo = gitbucket.Repo{
		Name:    "test",
		Owner:   gitbucket.User{Login: "public_group"},
		Private: false,
	}
	m = Migration{Repo: &repo}

	if m.isRepoMigratable(&labUser, labGroups) {
		t.Fatal("test failed")
	}

	repo = gitbucket.Repo{
		Name:    "test",
		Owner:   gitbucket.User{Login: "private_group"},
		Private: true,
	}
	m = Migration{Repo: &repo}

	if m.isRepoMigratable(&labUser, labGroups) {
		t.Fatal("test failed")
	}
}
