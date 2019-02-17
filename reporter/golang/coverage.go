package golang

import (
	"fmt"
	"github.com/fjmendes1994/charityreports/domain"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func GetCoverage(report domain.ProjectReport) (map[string]domain.Coverage, error) {

	coverages := make(map[string]domain.Coverage)

	for _, commit := range report.Commits {

		cov, err := runCoverageScript(report.ProjectPath, commit.Id)
		if err != nil {
			return coverages, err
		}
		coverages[commit.Id] = cov
	}
	return coverages, nil
}

func runCoverageScript(projectPath string, commitID string) (domain.Coverage, error) {
	fmt.Print("GetCoverage (" + commitID + "): ")

	cmd := exec.Command("sh", "golang_coverage.sh", projectPath, commitID, "-c", "1>&2")

	cmd.Dir = "/charityreports/scripts/"

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return domain.Coverage{}, errors.Wrap(err, "fail go coverage script")
	}

	cov := fmt.Sprintf("%s\n", stdoutStderr)

	fmt.Println(strings.TrimSpace(cov))
	covOutput := strings.Split(cov, "%")

	coverage, err := strconv.ParseFloat(covOutput[0], 64)
	if err != nil {
		return domain.Coverage{}, err
	}

	return domain.Coverage{
		Percentage: coverage,
	}, nil

}
