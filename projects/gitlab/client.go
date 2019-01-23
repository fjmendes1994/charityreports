package gitlab

import (
	"github.com/pkg/errors"
	"github.com/xanzy/go-gitlab"
)

type Gitlab struct {
	client   *gitlab.Client
	endpoint string
	username string
	password string
}

func New() *Gitlab {
	return &Gitlab{}
}

func (g *Gitlab) AddEndpoint(endpoint string) *Gitlab {
	g.endpoint = endpoint
	return g
}

func (g *Gitlab) AddUsername(username string) *Gitlab {
	g.username = username
	return g
}

func (g *Gitlab) AddPassword(password string) *Gitlab {
	g.password = password
	return g
}

func (g *Gitlab) addClient() error {
	client, err := gitlab.NewBasicAuthClient(nil, g.endpoint, g.username, g.password)
	if err != nil {
		return errors.Wrap(err, "cannot set gilab client")
	}
	g.client = client
	return nil
}

func (g *Gitlab) Build() (*Gitlab, error) {
	err := g.addClient()
	if err != nil {
		return nil, errors.Wrap(err, "cant set gilab client")
	}
	return g, err
}
