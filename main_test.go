package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramRenamesFolderGivenMWDataCurrentlyActive(t *testing.T) {
	setupTestDirectories("starwind")
	defer cleanupTestDirectories()

	renameActiveDirectories()
	cwd, _ := os.Getwd()
	cwdContents, _ := os.ReadDir(cwd)
	var containsDirectoriesNamedStarwind bool
	var directoryNames []string
	for _, entry := range cwdContents {
		directoryNames = append(directoryNames, entry.Name())
		if strings.Contains(entry.Name(), "starwind") {
			containsDirectoriesNamedStarwind = true
		}
	}
	fmt.Print(directoryNames)
	assert.False(t, containsDirectoriesNamedStarwind)
}

func TestProgramRenamesFolderGivenStarwindDataCurrentlyActive(t *testing.T) {
	setupTestDirectories("mw")
	defer cleanupTestDirectories()

	renameActiveDirectories()
	cwd, _ := os.Getwd()
	cwdContents, _ := os.ReadDir(cwd)
	var containsDirectoriesNamedMW bool
	var directoryNames []string
	for _, entry := range cwdContents {
		directoryNames = append(directoryNames, entry.Name())
		if strings.Contains(entry.Name(), "mw") {
			containsDirectoriesNamedMW = true
		}
	}
	fmt.Print(directoryNames)
	assert.False(t, containsDirectoriesNamedMW)
}

func setupTestDirectories(inactiveSet string) {
	os.Mkdir("config", 0777)
	os.Mkdir("data", 0777)
	os.Mkdir(inactiveSet+"config", 0777)
	os.Mkdir(inactiveSet+"data", 0777)
}

func cleanupTestDirectories() {
	os.RemoveAll("config")
	os.RemoveAll("data")
	os.RemoveAll("starwindconfig")
	os.RemoveAll("starwinddata")
	os.RemoveAll("mwconfig")
	os.RemoveAll("mwdata")
}

func contentToEntryHelper(contents []fs.DirEntry, entries []string) {
	for _, entry := range contents {
		fmt.Println(entry.Name())
		entries = append(entries, entry.Name())
	}
}
