package migration

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Migration ...
type Migration struct {
	Repo             *gitbucket.Repo `json:"repo"`
	Project          *gitlab.Project `json:"project"`
	IssuesMigratable bool            `json:"issues_migratable"`
	PullsMigratable  bool            `json:"pulls_migratable"`
	RepoMigratable   bool            `json:"repo_migratable"`
}

// GetMigrations ...
func (c *Client) GetMigrations() ([]Migration, error) {
	repos, err := c.bucket.GetRepos()
	if err != nil {
		return nil, err
	}

	projects, err := c.lab.GetProjects()
	if err != nil {
		return nil, err
	}

	lu, err := c.lab.GetAuthorizedUser()
	if err != nil {
		return nil, err
	}

	lus, err := c.lab.GetAuthorizedGroups()
	if err != nil {
		return nil, err
	}

	migrations := []Migration{}
	for _, r := range repos {
		repo := r

		var project *gitlab.Project
		for _, p := range projects {
			if repo.FullName == p.PathWithNamespace {
				project = &p
				break
			}
		}

		m := Migration{&repo, project, false, false, false}
		m.RepoMigratable = m.isRepoMigratable(lu, lus)
		m.IssuesMigratable = m.isIssuesMigratable()
		m.PullsMigratable = m.isPullsMigratable()
		migrations = append(migrations, m)
	}

	return migrations, nil
}

// GetMigration ...
func (c *Client) GetMigration(owner string, name string) (*Migration, error) {
	r, err := c.bucket.GetRepo(owner, name)
	if err != nil {
		return nil, err
	}

	p, err := c.lab.GetProject(owner, name)
	if err != nil {
		// Ignore error
	}

	lu, err := c.lab.GetAuthorizedUser()
	if err != nil {
		return nil, err
	}

	lus, err := c.lab.GetAuthorizedGroups()
	if err != nil {
		return nil, err
	}

	m := &Migration{r, p, false, false, false}
	m.RepoMigratable = m.isRepoMigratable(lu, lus)
	m.IssuesMigratable = m.isIssuesMigratable()
	m.PullsMigratable = m.isPullsMigratable()

	return m, nil
}

func (m *Migration) isPullsMigratable() bool {
	return m.Project != nil && len(m.Repo.Pulls) > len(m.Project.Merges)
}

func (m *Migration) isIssuesMigratable() bool {
	return m.Project != nil && len(m.Repo.Issues) > len(m.Project.Issues)
}

func (m *Migration) isRepoMigratable(labUser *gitlab.User, labGroups []gitlab.Group) bool {
	if m.Repo.Owner.IsOrganization() {
		for _, g := range labGroups {
			if m.Repo.Owner.Login == g.Path {
				return m.Repo.Private || !g.IsPrivate()
			}
		}

		return false
	}

	return m.Repo.Owner.Login == labUser.Username
}

// ExecRepo ...
func (c *Client) ExecRepo(m *Migration) (*Migration, error) {
	if m.Project == nil {
		nsID := 0
		if m.Repo.Owner.IsOrganization() {
			group, err := c.lab.GetGroup(m.Repo.Owner.Login)
			if err != nil {
				return nil, err
			}
			nsID = group.ID
		} else {
			user, err := c.lab.GetAuthorizedUser()
			if err != nil {
				return nil, err
			}
			nsID = user.ID
		}

		project, err := c.lab.CreateProject(nsID, m.Repo.Name, m.Repo.Description, m.Repo.Private)
		if err != nil {
			return nil, err
		}

		m.Project = project
	}

	storage := memory.NewStorage()
	worktree := memfs.New()

	err := c.bucket.Clone(m.Repo, storage, worktree)
	if err != nil {
		return nil, err
	}

	err = c.lab.Push(m.Project, storage, worktree)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// ExecIssues ...
func (c *Client) ExecIssues(m *Migration) (*Migration, error) {
	for _, i := range m.Repo.Issues {
		issue, err := c.lab.CreateIssue(m.Project, i.Number, i.Title, i.Body)
		if err != nil {
			return nil, err
		}

		for _, comment := range i.Comments {
			_, err := c.lab.CreateIssueComment(m.Project, issue, comment.Body, comment.CreatedAt)
			if err != nil {
				return nil, err
			}
		}

		if i.State == "closed" {
			err = c.lab.CloseIssue(issue)
			if err != nil {
				return nil, err
			}
		}
	}

	return c.GetMigration(m.Repo.Owner.Login, m.Repo.Name)
}

// ExecPulls ...
func (c *Client) ExecPulls(m *Migration) (*Migration, error) {
	for _, p := range m.Repo.Pulls {
		merge, err := c.lab.CreateMerge(m.Project, p.Title, p.Head.Ref, p.Base.Ref, p.Body)
		if err != nil {
			return nil, err
		}

		for _, comment := range p.Comments {
			_, err := c.lab.CreateMergeComment(m.Project, merge, comment.Body, comment.CreatedAt)
			if err != nil {
				return nil, err
			}
		}

		if p.State == "closed" {
			err = c.lab.CloseMerge(merge)
			if err != nil {
				return nil, err
			}
		}
	}

	return c.GetMigration(m.Repo.Owner.Login, m.Repo.Name)
}
