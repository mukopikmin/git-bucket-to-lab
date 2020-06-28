package migration

import (
	"git-bucket-to-lab/gitbucket"
	"git-bucket-to-lab/gitlab"
)

// Migration ...
type Migration struct {
	Repo           *gitbucket.Repo `json:"repo"`
	Project        *gitlab.Project `json:"project"`
	RepoMigrated   bool            `json:"repo_migrated"`
	IssuesMigrated bool            `json:"issues_migrated"`
	PullsMigrated  bool            `json:"pulls_migrated"`
	Migratable     bool            `json:"migratable"`
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

		rm := project != nil
		im := rm && len(repo.Issues) == len(project.Issues)
		pm := rm && len(repo.Pulls) == len(project.Merges)
		m := Migration{&repo, project, rm, im, pm, false}
		m.Migratable = m.isMigratable(lus)
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

	rm := p != nil
	im := rm && len(r.Issues) == len(p.Issues)
	pm := rm && len(r.Pulls) == len(p.Merges)
	m := &Migration{r, p, rm, im, pm, false}
	m.Migratable = m.isMigratable(lus)

	return m, nil
}

func (m *Migration) isMigratable(labGroups []gitlab.Group) bool {
	if m.Project == nil {
		return true
	}

	for _, g := range labGroups {
		if m.Repo.Owner.Login == g.Path {
			return true
		}
	}

	return false
}
