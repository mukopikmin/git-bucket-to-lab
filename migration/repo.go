package migration

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
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
		m.RepoMigratable = m.isRepoMigratable(lus)
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

	lus, err := c.lab.GetAuthorizedGroups()
	if err != nil {
		return nil, err
	}

	m := &Migration{r, p, false, false, false}
	m.RepoMigratable = m.isRepoMigratable(lus)
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

func (m *Migration) isRepoMigratable(labGroups []gitlab.Group) bool {
	if m.Repo.Owner.IsOrganization() {
		for _, g := range labGroups {
			if m.Repo.Owner.Login == g.Path {
				if !m.Repo.Private || g.IsPrivate() {
					return false
				}
			}
		}
	}

	return true
}
