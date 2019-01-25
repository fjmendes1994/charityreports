package app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fjmendes1994/charityreports/reports/golang"

	"github.com/kr/pretty"

	"github.com/fjmendes1994/charityreports/projects/github"
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
	commits, err := ghclient.ListCommits("github.com/kr/pretty")
	if err != nil {
		fmt.Println(err)
	}

	for _, commit := range commits {
		fmt.Println(commit.GetSHA())
	}
	fmt.Println(len(commits), " commits.")

	coverages := make([][]string, len(commits))

	for i, commit := range commits {
		cov, err := golang.GetCoverage("github.com/kr/pretty", commit.GetSHA())
		if err != nil {
			fmt.Println(err)
		}

		c := strings.Split(cov, "%")

		fmt.Println(c[0])

		coverages[i] = []string{commit.GetSHA(), c[0]}

	}
	pretty.Println(coverages)
	Write(reverse(coverages))

}

func reverse(array [][]string) [][]string {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func Write(coverages [][]string) {
	records := [][]string{
		{"commit_id", "coverage"},
	}

	records = append(records, coverages...)

	file, err := os.Create("./reports.csv")
	if err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
