package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	searchDir := "."
	var files []string
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if f.Name() == ".git" || f.Name() == ".." || f.Name() == "output.zip" {
			return filepath.SkipDir
		}
		if f.Name() != "." {
			files = append(files, f.Name())
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
