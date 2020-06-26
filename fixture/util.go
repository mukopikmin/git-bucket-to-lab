package fixture

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

func clone(storage *memory.Storage, worktree billy.Filesystem, url string) error {
	_, err := git.Clone(storage, worktree, &git.CloneOptions{
		URL:          url,
		SingleBranch: false,
	})
	if err != nil {
		return err
	}

	return nil
}

func push(storage *memory.Storage, worktree billy.Filesystem, url string) error {
	remote := "bucket"
	r, err := git.Open(storage, worktree)
	if err != nil {
		return err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: remote,
		URLs: []string{url},
	})
	if err != nil {
		return err
	}

	err = r.Push(&git.PushOptions{
		RemoteName: remote,
		RefSpecs: []config.RefSpec{
			config.RefSpec("+refs/remotes/origin/*:refs/heads/*"),
			config.RefSpec("+refs/tags/*:refs/tags/*"),
		},
		Auth: &http.BasicAuth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		return err
	}

	return nil
}
