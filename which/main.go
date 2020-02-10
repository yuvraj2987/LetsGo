package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func fileExists(dir string, file string) bool {
	fullPath := filepath.Join(dir, file)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func search(paths []string, args []string) []string {
	pathLen := len(paths)
	fileFoundPaths := make([]string, pathLen)
	fileCount := 0
	for _, file := range args {
		for _, dir := range paths {
			if fileExists(dir, file) {
				fileFoundPaths[fileCount] = dir
				fileCount += 1
			}
		}
	}

	return fileFoundPaths[0:fileCount]
}

func main() {
	ptrFlagA := flag.Bool("a", false, "a")
	ptrFlagS := flag.Bool("s", false, "s")
	flag.Parse()

	if flag.NArg() < 1 {
		os.Exit(1)
	}

	args := flag.Args()

	ospaths := strings.Split(os.Getenv("PATH"), ":")
	pathDir := search(ospaths, args)

	for _, d := range pathDir {
		fmt.Println(d)
	}

	if *ptrFlagA {
		fmt.Println("Print all paths")
	}

	if *ptrFlagS {
		fmt.Println("Return true if executable found in any one path.")
	}

}
