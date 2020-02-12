package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func exeExists(dir string, file string) bool {
	fullPath := filepath.Join(dir, file)
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		return false
	}

	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return false
	}

	// if mode&0111 != 0 {
	// 	return false // execute bit is not set for any of owner,group and other mode.
	// }

	return true

}

func searchExe(paths []string, file string) []string {
	pathLen := len(paths)
	fileFoundPaths := make([]string, pathLen)
	fileCount := 0
	for _, dir := range paths {
		if exeExists(dir, file) {
			fileFoundPaths[fileCount] = filepath.Join(dir, file)
			fileCount += 1
		}
	}

	return fileFoundPaths[0:fileCount]
}

func search(paths []string, exes []string) map[string][]string {
	var mapPaths = make(map[string][]string)
	for _, exe := range exes {
		mapPaths[exe] = searchExe(paths, exe)
	}
	return mapPaths

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
	exeMaps := search(ospaths, args)

	foundAllFiles := true

	for file, paths := range exeMaps {
		if *ptrFlagS {
			// Do not print anything
			if len(paths) < 1 {
				foundAllFiles = false
			}

		} else if *ptrFlagA {
			// Print all instances of all executables
			if len(paths) < 1 {
				fmt.Println(file, " not found")
				foundAllFiles = false
			} else {
				for _, path := range paths {
					fmt.Println(path)
				}
			}
		} else {
			// No flag mode, print only first path of each
			if len(paths) < 1 {
				foundAllFiles = false
				fmt.Println(file, " not found")
			} else {
				fmt.Println(paths[0])
			}
		}

	}

	if !foundAllFiles {
		os.Exit(1)
	}

	os.Exit(0)

}
