package golang

import (
	"fmt"
	"log"
	"os/exec"
)

// go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/) && go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
func GetCoverage(projectUrl string, commitId string) error {
	fmt.Println("GetCoverage (" + commitId + "):")

	runCoverageScript(projectUrl, commitId)

	return nil

}

func runCoverageScript(projectPath string, commitId string) {
	cmd := exec.Command("sh", "coverage.sh", projectPath, commitId, "-c", "1>&2")
	cmd.Dir = "/home/fjmendes1994/go/src/charityreports/reports/golang/"

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
