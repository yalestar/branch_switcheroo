package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	checkr "github.com/yalestar/branch_switcheroo/pkg"
	"log"
	"os"
)

func main() {

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
