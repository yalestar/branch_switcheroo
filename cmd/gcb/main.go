package main

import (
	"fmt"
	"github.com/atotto/clipboard"
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

	currentBranch, err := checkr.GetCurrentBranch()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	err = clipboard.WriteAll(currentBranch)
	if err == nil {
		fmt.Printf("Current branch %s has been zapped to clipboard", currentBranch)
		os.Exit(0)
	}
}
