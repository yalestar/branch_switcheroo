package checkr

import (
	"log"
	"os/exec"
)

const gitBinary = "/usr/bin/git"

func ChangeToBranch(branch, repoPath string) error {
	argList := []string{"-C", repoPath, "checkout", branch}
	lsCmd := exec.Command(gitBinary, argList...)

	_, err := lsCmd.Output()

	if err != nil {
		log.Fatal(err)
	}

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
