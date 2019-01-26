package app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fjmendes1994/charityreports/reports/golang"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type App struct {
	providers struct{}
	config    struct{}
}

func Start() {
	var repository, username, password string

	flag.StringVar(&repository, "r", "", "Repository")
	flag.StringVar(&username, "u", "", "Gitlab username")
	flag.StringVar(&password, "p", "", "Gitlab password")
	flag.Parse()

	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	r, _ := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://" + repository,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
	})

	// ... retrieves the branch pointed by HEAD
	ref, _ := r.Head()

	// ... retrieves the commit history
	cIter, _ := r.Log(&git.LogOptions{From: ref.Hash()})

	// ... just iterates over the commits, printing it
	var i int
	coverages := make([][]string, 0)

	_ = cIter.ForEach(func(commit *object.Commit) error {
		cov, err := golang.GetCoverage(repository, commit.Hash.String())
		if err != nil {
			fmt.Println(err)
		}

		c := strings.Split(cov, "%")

		fmt.Println(c[0])

		coverages = append(coverages, []string{commit.Hash.String(), c[0]})
		i++

		return nil
	})
	fmt.Println(coverages)
	Write(coverages)

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
