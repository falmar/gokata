package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// clean up other languages

	keep := []string{
		".github",
		"texttests",
		"go",
		".git",
		".idea", // intellij
	}

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip files
		if !d.IsDir() || d.Name() == "." || d.Name() == ".." {
			return nil
		}

		// skip needed dirs
		for _, k := range keep {
			if k == d.Name() {
				return fs.SkipDir
			}
		}

		fmt.Println("removing", d.Name())
		err = os.RemoveAll(d.Name())

		return fs.SkipDir
	})

	if err != nil {
		log.Fatalf("walk dir err: %v", err)
	}
}
