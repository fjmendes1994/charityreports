package git

import (
	"fmt"
	"github.com/fjmendes1994/charityreports/domain"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type Client struct {
	httpAuth transport.AuthMethod
}

func Builder() *Client {
	return &Client{}

}

func (c *Client) Auth(method transport.AuthMethod) {
	c.httpAuth = method
}

func (c *Client) Build() *Client {
	return c
}

func (c *Client) GetCommits(repositoryUrl string) ([]domain.Commit, error) {

	repository, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:  repositoryUrl,
		Auth: c.httpAuth,
	})
	if err != nil {
		fmt.Println(err)
	}

	head, err := repository.Head()
	if err != nil {
		fmt.Println(err)
	}

	commitIter, err := repository.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		fmt.Println(err)
	}

	var index int
	commits := make([]domain.Commit, 0)

	err = commitIter.ForEach(func(commit *object.Commit) error {
		commits = append(commits, domain.Commit{
			Id:    commit.Hash.String(),
			Index: index,
		})
		index++
		return nil
	})
	return commits, err
}
