package app

import (
	"charityreports/projects/gitlab"
	"charityreports/reports/golang"
	"flag"
	"fmt"
)

type App struct {
	providers struct{}
	config    struct{}
}

func Start() {
	var domain, username, password string

	flag.StringVar(&domain, "d", "", "Gitlab doman")
	flag.StringVar(&username, "u", "", "Gitlab username")
	flag.StringVar(&password, "p", "", "Gitlab password")
	flag.Parse()

	client, err := gitlab.New().
		AddEndpoint(domain).
		AddUsername(username).
		AddPassword(password).
		Build()

	if err != nil {
		fmt.Println(err)
	}

	commits, err := client.ListCommits("fernandomendes1/sd-kmeans-mpi")
	fmt.Println(commits)
	if err != nil {
		fmt.Println(err)
	}

	for _, commit := range commits {
		err = golang.GetCoverage("github.com/xanzy/go-gitlab", commit.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

}
