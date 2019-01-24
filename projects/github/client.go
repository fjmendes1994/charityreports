package github

import (
	"github.com/google/go-github/github"
)

type GitHub struct {
	client *github.Client
}

func New() *GitHub {
	client := github.NewClient(nil)
	return &GitHub{client: client}
}
