package main

import (
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramDetectsFolderByName(t *testing.T) {
	err := os.Mkdir("testDir", 0755)
	if err != nil {
		t.Errorf("Failed to create testDir in TestProgramDetectsFolderByName: %v", err)
	}
	defer os.RemoveAll("testDir")
	assert.Equal(t, "testDir", detectFolderByName("testDir"))
}

func TestProgramCorrectlyCopiesDirectoryIntoTarget(t *testing.T) {
	os.Mkdir("testTargetDir", 0755)
	defer os.RemoveAll("testTargetDir")
	os.MkdirAll("testSourceDir/testSubdir", 0755)
	os.WriteFile("testSourceDir/testSubdir/testFile1", []byte(""), 0755)
	defer os.RemoveAll("testSourceDir/testSubdir/testFile1")

	err := copyDirectoryIntoTarget("testSourceDir", "testTargetDir")
	if err != nil {
		t.Errorf("Failed to copy: %v", err)
	}

	originalDirContents, _ := os.ReadDir("testSourceDir")
	targetDirContents, _ := os.ReadDir("testTargetDir")

	var originalDirEntries []string
	var targetDirEntries []string

	fmt.Println("originalDirContents:")
	contentToEntryHelper(originalDirContents, originalDirEntries)

	fmt.Println("targetDirContents:")
	contentToEntryHelper(targetDirContents, targetDirEntries)

	assert.Equal(t, originalDirEntries, targetDirEntries)
}

func contentToEntryHelper(contents []fs.DirEntry, entries []string) {
	for _, entry := range contents {
		fmt.Println(entry.Name())
		entries = append(entries, entry.Name())
	}
}
