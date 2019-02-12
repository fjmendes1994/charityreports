package app

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

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
	var repositoryUrl, repositoryPath, language string
	var auth bool

	flag.StringVar(&repositoryUrl, "r", "https://github.com/olivere/elastic", "Repository")
	flag.StringVar(&language, "l", "golang", "Language")
	flag.BoolVar(&auth, "auth", false, "Auth")

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
		var username, password string
		if auth {
			username, password, err = getCredentials()
			if err != nil {
				fmt.Println(err)
				break
			}
		}

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

	if username == "" && password == "" {
		r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: repositoryUrl,
		})
		if err != nil {
			return nil, err
		}
		return r, nil

	}

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

func getCredentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", "", err
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password), nil
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
