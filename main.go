package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: analyzer <file.go or directory>")
		os.Exit(1)
	}

	target := os.Args[1]

	var filesToCheck []string
	info, err := os.Stat(target)
	if err != nil {
		fmt.Println("stat error:", err)
		os.Exit(1)
	}
	if info.IsDir() {
		filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() && strings.HasSuffix(path, ".go") {
				filesToCheck = append(filesToCheck, path)
			}
			return nil
		})
	} else {
		filesToCheck = []string{target}
	}

	for _, filename := range filesToCheck {

		tokens := token.NewFileSet()
		node, err := parser.ParseFile(tokens, filename, nil, parser.AllErrors)
		if err != nil {
			fmt.Println("parse error:", err)
			continue
		}

		checkUncheckedErrors(tokens, node)
		checkCyclomaticComplexity(tokens, node)
	}
}
