package main

import (
	"fmt"
	checkr "github.com/yalestar/branch_switcheroo/pkg"
	"log"
	"os"
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
		fmt.Println("I'm bout to check out MASTER")
		err := checkr.ChangeToBranch("master", daRepo)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(-1)
		}
		os.Exit(0)
	}

	hasMain, err := checkr.RepoHasBranch("main", daRepo)

	if hasMain {
		fmt.Println("I'm bout to check out MAIN")
		err := checkr.ChangeToBranch("main", daRepo)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(-1)
		}
		os.Exit(0)
	}

}
