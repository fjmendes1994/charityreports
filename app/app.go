package app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/fjmendes1994/charityreports/reports/golang"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type App struct {
	providers struct{}
	config    struct{}
}

func Start() {
	var repositoryUrl, repositoryPath, username, password, language string

	flag.StringVar(&repositoryUrl, "r", "", "Repository")
	flag.StringVar(&username, "u", "", "User")
	flag.StringVar(&password, "p", "", "Pass")
	flag.StringVar(&language, "l", "", "Language")

	flag.Parse()

	url, err := url.Parse(repositoryUrl)
	if err != nil {
		fmt.Println(err)
	}

	repositoryPath = url.Host + url.Path
	fmt.Println("Repository path: " + url.Host + url.Path)

	var repository *git.Repository
	switch url.Scheme {
	case "https":
		repository, err = getRepository(repositoryUrl, username, password)
	default:
		fmt.Println("Not suported: " + url.Scheme)
	}
	if err != nil {
		fmt.Println(err)
	}

	numberOfCommits, err := countCommits(repository)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of commits: %d \n", numberOfCommits)

	coverages := make([][]string, 0)
	switch language {
	case "golang":
		coverages, err = getGolangCoverages(repository, repositoryPath)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Not suported: " + language)

	}

	fmt.Println(coverages)
	Write(coverages)

}

func getGolangCoverages(repository *git.Repository, repositoryPath string) ([][]string, error) {
	head, err := repository.Head()
	if err != nil {
		fmt.Println(err)
	}

	coverages := make([][]string, 0)

	commitIter, err := repository.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		fmt.Println(err)
	}
	err = commitIter.ForEach(func(commit *object.Commit) error {
		cov, err := golang.GetCoverage(repositoryPath, commit.Hash.String())
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(strings.TrimSpace(cov))
		covOutput := strings.Split(cov, "%")

		coverages = append(coverages, []string{commit.Hash.String(), covOutput[0]})

		return nil
	})
	return coverages, err
}

func countCommits(repository *git.Repository) (int, error) {
	head, err := repository.Head()
	if err != nil {
		fmt.Println(err)
	}

	commitIter, err := repository.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		fmt.Println(err)
	}

	var numberOfCommits int

	err = commitIter.ForEach(func(commit *object.Commit) error {
		numberOfCommits++
		return nil
	})
	return numberOfCommits, err
}

func getRepository(repositoryUrl string, username string, password string) (*git.Repository, error) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repositoryUrl,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		return nil, err
	}
	return r, nil
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

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
