package gitlab

import "github.com/pkg/errors"

func (g *Gitlab) getProjectId(url string) (int, error) {

	proj, _, err := g.client.Projects.GetProject(url)
	if err != nil {
		return -1, errors.Wrap(err, "cant get project id")
	}
	return proj.ID, nil
}
