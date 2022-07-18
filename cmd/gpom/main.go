package main

import (
	"fmt"
	"log"
	"os"

	checkr "github.com/yalestar/branch_switcheroo/pkg"
)

var version = ""

func main() {

	versionArg := os.Args[1:]

	if len(versionArg) == 1 && versionArg[0] == "version" {
		fmt.Println(version)
		os.Exit(0)
	}

	daRepo, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	hasMaster, err := checkr.RepoHasBranch("master", daRepo)
	if err != nil {
		log.Println(err)
	}

	if hasMaster {
		cb, err := checkr.GetCurrentBranch()
		if cb != "master" {
			fmt.Println("Cowardlyly refusing to pull master into a non-master branch. " +
				"Do it yourself.")
			fmt.Println("Exiting")
			os.Exit(-1)
		}
		fmt.Println("I'm bout to pull from MASTER")
		err = checkr.PullBranch("master", daRepo)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(-1)
		}
		os.Exit(0)
	}

	hasMain, err := checkr.RepoHasBranch("main", daRepo)

	if hasMain {
		cb, err := checkr.GetCurrentBranch()
		if cb != "main" {
			fmt.Println("Cowardlyly refusing to pull main into a non-main branch. " +
				"Do it yourself.")
			fmt.Println("Exiting")
			os.Exit(-1)
		}
		fmt.Println("I'm bout to pull from MAIN")

		err = checkr.PullBranch("main", daRepo)

		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(-1)
		}
		os.Exit(0)
	}

}
