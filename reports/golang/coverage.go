package golang

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

const BASEDIR = "./temp/"

// go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/) && go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
func GetCoverage(projectUrl string, commitId string) error {
	fmt.Println("GetCoverage:")

	err := clone(projectUrl)
	if err != nil {
		return err
	}

	err = checkout(commitId)
	if err != nil {
		return err
	}

	err = clean(commitId)
	if err != nil {
		return err
	}

	return nil

}

func checkout(commitId string) error {
	gitCheckout := exec.Command("git", "checkout", commitId)
	gitCheckout.Dir = BASEDIR + "goverage"
	gitCheckoutOut, err := gitCheckout.StdoutPipe()
	if err != nil {
		return err
	}

	err = gitCheckout.Start()
	if err != nil {
		return err
	}
	fmt.Println("Checkout:")

	readGitCloneOut, err := ioutil.ReadAll(gitCheckoutOut)
	if err != nil {
		return err
	}

	err = gitCheckout.Wait()
	if err != nil {
		return err
	}

	fmt.Println(string(readGitCloneOut))

	return nil
}

func clone(projectUrl string) error {
	gitClone := exec.Command("git", "clone", projectUrl)
	gitClone.Dir = BASEDIR
	gitCloneOut, err := gitClone.StdoutPipe()
	if err != nil {
		return err
	}

	err = gitClone.Start()
	if err != nil {
		return err
	}
	fmt.Println("Clone:")

	readGitCloneOut, err := ioutil.ReadAll(gitCloneOut)
	if err != nil {
		return err
	}

	err = gitClone.Wait()
	if err != nil {
		return err
	}

	fmt.Println(string(readGitCloneOut))

	return nil
}

func clean(projectUrl string) error {
	rm := exec.Command("rm", "-r", BASEDIR+"goverage")
	rm.Dir = BASEDIR
	rmOut, err := rm.StdoutPipe()
	if err != nil {
		return err
	}

	err = rm.Start()
	if err != nil {
		return err
	}
	fmt.Println("Rm:")

	readGitCloneOut, err := ioutil.ReadAll(rmOut)
	if err != nil {
		return err
	}

	err = rm.Wait()
	if err != nil {
		return err
	}

	fmt.Println(string(readGitCloneOut))

	return nil
}
