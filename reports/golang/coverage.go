package golang

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
)

// go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/) && go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
func GetCoverage(projectUrl string, commitId string) (string, error) {
	fmt.Println("GetCoverage (" + commitId + "):")

	return runCoverageScript(projectUrl, commitId)

}

func runCoverageScript(projectPath string, commitId string) (string, error) {
	cmd := exec.Command("sh", "coverage.sh", projectPath, commitId, "-c", "1>&2")
	cmd.Dir = "/home/fjmendes1994/go/src/charityreports/reports/golang/"

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, "fail go coverage script")
	}

	return fmt.Sprintf("%s\n", stdoutStderr), nil
}
