package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting directory rename utility.")
	renameActiveDirectories()
	fmt.Println("Rename utility finished.")
}

func renameActiveDirectories() {
	cwd, _ := os.Getwd()
	cwdContents, _ := os.ReadDir(cwd)

	var isStarwindActive bool
	var isMWActive bool

	for _, entry := range cwdContents {
		if strings.Contains(entry.Name(), "starwind") {
			isMWActive = true
			break
		} else if strings.Contains(entry.Name(), "mw") {
			isStarwindActive = true
			break
		}
	}

	if isStarwindActive {
		fmt.Printf("Starwind currently active, switching to Morrowind")
		for _, entry := range cwdContents {
			if entry.Name() == "data" || entry.Name() == "config" {
				os.Rename(entry.Name(), "starwind"+entry.Name())
			}
			if entry.Name() == "mwdata" {
				os.Rename(entry.Name(), "data")
			}
			if entry.Name() == "mwconfig" {
				os.Rename(entry.Name(), "config")
			}
		}
	} else if isMWActive {
		fmt.Printf("Morrowind currently active, switching to Starwind")
		for _, entry := range cwdContents {
			if entry.Name() == "data" || entry.Name() == "config" {
				os.Rename(entry.Name(), "mw"+entry.Name())
			}
			if entry.Name() == "starwinddata" {
				os.Rename(entry.Name(), "data")
			}
			if entry.Name() == "starwindconfig" {
				os.Rename(entry.Name(), "config")
			}

		}
	} else {
		fmt.Printf("Doing nothing.")
	}
}
