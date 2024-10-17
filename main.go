package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
}

func detectFolderByName(name string) string {
	var output string
	cwdContents, _ := os.ReadDir(".")
	for _, entry := range cwdContents {
		if entry.Name() == name {
			output = entry.Name()
		}
	}

	return output
}

func copyDirectoryIntoTarget(sourceDirectory string, targetDirectory string) error {
	sourceFilesystem := os.DirFS(sourceDirectory)
	return os.CopyFS(targetDirectory, sourceFilesystem)
}
