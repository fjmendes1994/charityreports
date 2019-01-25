package app

import (
	"flag"
	"fmt"

	"github.com/fjmendes1994/charityreports/projects/github"
	"github.com/fjmendes1994/charityreports/reports/golang"
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

	//glclient, err := gitlab.New().
	//	AddEndpoint(domain).
	//	AddUsername(username).
	//	AddPassword(password).
	//	Build()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//commits, err := glclient.ListCommits("fernandomendes1/sd-kmeans-mpi")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(commits)

	ghclient := github.New()
	commits, err := ghclient.ListCommits("github.com/xanzy/go-gitlab")
	if err != nil {
		fmt.Println(err)
	}

	for _, commit := range commits {
		fmt.Println(commit.GetSHA())
	}
	fmt.Println(len(commits), " commits.")

	coverages := make([]string, len(commits))

	for i, commit := range commits {
		cov, err := golang.GetCoverage("github.com/xanzy/go-gitlab", commit.GetSHA())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cov)
		coverages[i] = cov

	}

}
