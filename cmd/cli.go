package cmd

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/fjmendes1994/charityreports/domain"
	git2 "github.com/fjmendes1994/charityreports/git"
	"github.com/kr/pretty"
	"log"
	"net/url"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/fjmendes1994/charityreports/reporter/golang"
)

func StartCli() {
	var repositoryUrl, repositoryPath, language, output string
	var auth bool

	flag.StringVar(&repositoryUrl, "r", "https://github.com/kr/pretty", "Repository")
	flag.StringVar(&language, "l", "golang", "Language")
	flag.StringVar(&output, "o", "reports.csv", "Output")
	flag.BoolVar(&auth, "a", false, "Auth")

	flag.Parse()

	url, err := url.Parse(repositoryUrl)
	if err != nil {
		fmt.Println(err)
	}

	repositoryPath = url.Host + url.Path
	fmt.Println("Repository path: " + url.Host + url.Path)

	switch url.Scheme {
	case "https":

		git := git2.Builder().Build()
		if err != nil {
			fmt.Println(err)
			return
		}

		commits, err := git.GetCommits(repositoryUrl)
		if err != nil {
			fmt.Println(err)
			return
		}
		pretty.Println("Number of commits: ", len(commits))

		projectReport := domain.ProjectReport{
			ProjectPath: repositoryPath,
			Branch:      "master",
			Commits:     commits,
		}

		coverages, err := golang.GetCoverage(projectReport)
		if err != nil {
			fmt.Println(err)
			return
		}
		pretty.Println(coverages)

	default:
		fmt.Println("Not suported: " + url.Scheme)
	}

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

func Write(coverages [][]string, output string) {
	records := [][]string{
		{"commit_id", "coverage"},
	}

	records = append(records, coverages...)

	file, err := os.Create("/charityreports/out/reports/" + output)
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
