package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var searchDir string
	if len(os.Args) > 1 {
		searchDir = os.Args[1]
	} else {
		searchDir = "."
	}
	var files []string
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if f.Name() == ".git" || f.Name() == ".." || f.Name() == "output.zip" {
			return filepath.SkipDir
		}
		if f.Name() != "." && !f.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	output := "output.zip"
	err := ZipDir(output, files)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Zipped Directory: " + output)
}
