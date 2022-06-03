package main

import (
	"fmt"
	"log"
	"os/exec"
)

const gitBinary = "/usr/bin/git"

func main() {

	// TODO: use PWD
	hasMaster, err := repoHasBranch("master", "/Users/r622233/dev/paladin")
	hasMain, err := repoHasBranch("main", "/Users/r622233/dev/paladin")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(hasMaster)
	fmt.Println(hasMain)

}

func repoHasBranch(branch, repoPath string) (bool, error) {

	argList := []string{"-C", repoPath, "branch", "--list", branch}
	lsCmd := exec.Command(gitBinary, argList...)

	lsOut, err := lsCmd.Output()

	if err != nil {
		log.Println(err)
	}

	return len(lsOut) > 0, nil
}
