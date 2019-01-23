package app

import (
	"charityreports/projects/gitlab"
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

	fmt.Println(client.ListCommits("fernandomendes1/sd-kmeans-mpi"))

	// TODO
	//err = golang.GetCoverage("https://github.com/haya14busa/goverage.git", "578a76dd1c685cbe2df589c43f4259912bb28889")
	//if err != nil {
	//	fmt.Println(err)
	//}

}
