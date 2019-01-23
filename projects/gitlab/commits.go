package gitlab

import (
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

// ListCommits list all commit of an application
func (g *Gitlab) ListCommits(appUrl string) ([]*gitlab.Commit, error) {

	pid, err := g.getProjectId(appUrl)
	if err != nil {
		return nil, errors.Wrap(err, "cant get project id")
	}

	commits, _, err := g.client.Commits.ListCommits(pid, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cant get project commits")
	}
	return commits, nil
}
