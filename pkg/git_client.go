package checkr

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const gitBinary = "/usr/bin/git"

func PullBranch(branch, repoPath string) error {

	argList := []string{"-C", repoPath, "pull", "origin", branch}
	lsCmd := exec.Command(gitBinary, argList...)

	cmdOutput, err := lsCmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(cmdOutput))

	return nil
}

func ChangeToBranch(branch, repoPath string) error {
	argList := []string{"-C", repoPath, "checkout", branch}
	lsCmd := exec.Command(gitBinary, argList...)

	cmdOutput, err := lsCmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmdOutput))
	return nil
}

func RepoHasBranch(branch, repoPath string) (bool, error) {

	argList := []string{"-C", repoPath, "branch", "--list", branch}
	lsCmd := exec.Command(gitBinary, argList...)

	lsOut, err := lsCmd.Output()

	if err != nil {
		log.Println(err)
	}

	return len(lsOut) > 0, nil
}

func GetCurrentBranch() (string, error) {
	// pre Git 2.2.2: rev-parse --abbrev-ref HEAD
	argList := []string{"branch", "--show-current"}
	gbCmd := exec.Command(gitBinary, argList...)

	lsOut, err := gbCmd.Output()
	currentBranch := strings.Trim(string(lsOut), "\n")

	if err != nil {
		log.Println(err)
	}

	return currentBranch, nil

}
