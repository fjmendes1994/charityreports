package golang

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

func GetCoverage(projectUrl string, commitId string) (string, error) {
	fmt.Print("GetCoverage (" + commitId + "): ")

	return runCoverageScript(projectUrl, commitId)

}

func runCoverageScript(projectPath string, commitId string) (string, error) {
	cmd := exec.Command("sh", "coverage.sh", projectPath, commitId, "-c", "1>&2")

	goPath := os.Getenv("GOPATH")

	cmd.Dir = goPath + "/src/" + projectPath + "/reports/golang/"

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, "fail go coverage script")
	}

	return fmt.Sprintf("%s\n", stdoutStderr), nil
}
