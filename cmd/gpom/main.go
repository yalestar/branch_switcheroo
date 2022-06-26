package main

import (
	"fmt"
	"log"
	"os"

	checkr "github.com/yalestar/branch_switcheroo/pkg"
)

func main() {

	daRepo, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	hasMaster, err := checkr.RepoHasBranch("master", daRepo)
	if err != nil {
		log.Println(err)
	}

	if hasMaster {
		fmt.Println("I'm bout to pull from MASTER")
		err := checkr.PullBranch("master", daRepo)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(-1)
		}
		os.Exit(0)
	}

	hasMain, err := checkr.RepoHasBranch("main", daRepo)

	if hasMain {
		fmt.Println("I'm bout to pull from MAIN")
		err := checkr.PullBranch("main", daRepo)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(-1)
		}
		os.Exit(0)
	}

}
