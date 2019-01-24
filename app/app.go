package app

import (
	"charityreports/reports/golang"
	"fmt"
)

type App struct {
	providers struct{}
	config    struct{}
}

func Start() {
	//var domain, username, password string
	//
	//flag.StringVar(&domain, "d", "", "Gitlab doman")
	//flag.StringVar(&username, "u", "", "Gitlab username")
	//flag.StringVar(&password, "p", "", "Gitlab password")
	//flag.Parse()

	//client, err := gitlab.New().
	//	AddEndpoint(domain).
	//	AddUsername(username).
	//	AddPassword(password).
	//	Build()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(client.ListCommits("fernandomendes1/sd-kmeans-mpi"))

	// TODO
	err := golang.GetCoverage("github.com/xanzy/go-gitlab", "14423f46413f55f9447c41a9ec20c032d1ff53f4")
	if err != nil {
		fmt.Println(err)
	}

}
