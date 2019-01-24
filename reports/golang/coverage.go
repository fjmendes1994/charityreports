package golang

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"os/user"
)

func baseDir() string {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/go/src/"

}



func runSh(){
	cmd := exec.Command("sh", "coverage.sh","-c", "1>&2")
	cmd.Dir = baseDir() + "github.com/xanzy/go-gitlab"
	stdoutStderr, err := cmd.CombinedOutput()

	if err != nil {

		log.Fatal(err)

	}

	fmt.Printf("%s\n", stdoutStderr)
}





// go test -coverprofile=cov.out -race -v $(go list ./... | grep -v /vendor/) && go tool cover -func=cov.out | grep total: | awk ' {print $3} ';
func GetCoverage(projectUrl string, commitId string) error {
	fmt.Println("GetCoverage:")


	runSh()

	//err := get(projectUrl)
	//if err != nil {
	//	return err
	//}

	//err = checkout(commitId)
	//if err != nil {
	//	return err
	//}
	//
	//err = depInit()
	//if err != nil {
	//	return err
	//}
	//
	//err = depEnsure()
	//if err != nil {
	//	return err
	//}
	//
	//err = cover()
	//if err != nil {
	//	return err
	//}
	//
	//err = clean()
	//if err != nil {
	//	return err
	//}

	return nil

}


func get(projectUrl string) error {
	goGet := exec.Command("go", "get", projectUrl)
	goGet.Dir = baseDir()
	goGetOut, err := goGet.StdoutPipe()
	if err != nil {
		return err
	}

	err = goGet.Start()
	if err != nil {
		return err
	}
	fmt.Println("Get:")

	readGoGetOut, err := ioutil.ReadAll(goGetOut)
	if err != nil {
		return err
	}

	err = goGet.Wait()
	if err != nil {
		return err
	}

	fmt.Println(string(readGoGetOut))

	return nil
}


func clone(projectUrl string) error {
	gitClone := exec.Command("git", "clone", projectUrl)
	gitClone.Dir = baseDir()
	gitCloneOut, err := gitClone.StderrPipe()
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

func checkout(commitId string) error {
	gitCheckout := exec.Command("git", "checkout", commitId)
	gitCheckout.Dir = baseDir() + "github.com/xanzy/go-gitlab"
	gitCheckoutOut, err := gitCheckout.StderrPipe()
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


func depInit() error {
	depInit := exec.Command("dep", "init")
	depInit.Dir = baseDir() + "github.com/xanzy/go-gitlab"
	depInitOut, err := depInit.StderrPipe()
	if err != nil {
		return err
	}

	err = depInit.Start()
	if err != nil {
		return err
	}
	fmt.Println("Dep init:")

	readDepInitOut, err := ioutil.ReadAll(depInitOut)
	if err != nil {
		return err
	}

	err = depInit.Wait()
	if err != nil {
		return err
	}

	fmt.Println(string(readDepInitOut))

	return nil
}

func depEnsure() error {
	depEnsure := exec.Command("dep", "ensure")
	depEnsure.Dir = baseDir() + "github.com/xanzy/go-gitlab"
	depEnsureOut, err := depEnsure.StdoutPipe()
	if err != nil {
		return err
	}

	err = depEnsure.Start()
	if err != nil {
		return err
	}
	fmt.Println("Ensure:")

	readDepEnsureOut, err := ioutil.ReadAll(depEnsureOut)
	if err != nil {
		return err
	}

	err = depEnsure.Wait()
	if err != nil {
		return err
	}

	fmt.Println(string(readDepEnsureOut))

	return nil
}

func cover() error {
	coverage := exec.Command("go", "test", "./...", "-coverprofile=cov.out", "-race")
	coverage.Dir = baseDir() + "github.com/xanzy/go-gitlab"
	coverageOut, err := coverage.StdoutPipe()
	if err != nil {
		return err
	}

	err = coverage.Start()
	if err != nil {
		return err
	}
	fmt.Println("Cover:")

	readCoverageOut, err := ioutil.ReadAll(coverageOut)
	if err != nil {
		return err
	}

	err = coverage.Wait()
	if err != nil {
		return err
	}

	fmt.Println(string(readCoverageOut))

	return nil
}

func report() error {
	// TODO

	return nil
}



func clean() error {
	rm := exec.Command("rm", "-rf", "github.com/xanzy/go-gitlab")
	rm.Dir = baseDir()
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
