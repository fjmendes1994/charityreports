package github

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
)

func (g *GitHub) ListCommits(projectUrl string) ([]*github.RepositoryCommit, error) {
	ctx := context.Background()
	url := strings.Split(projectUrl, "/")
	owner := url[1]
	repo := url[2]

	commits, _, err := g.client.Repositories.ListCommits(ctx, owner, repo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cant list github repo commits")
	}
	return commits, nil
}
