package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func help() {
	fmt.Println("Bad Option")
	fmt.Println("Correct Usage")
	fmt.Printf("pwd\n pwd -P")
}

func main() {
	args := os.Args
	flagP := false
	if len(args) > 1 {
		opt := args[1]
		if len(args) > 2 && opt != "-P" {
			help()
			os.Exit(1)
		}
		flagP = true

	}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if !flagP {
		fmt.Println(cwd)
		os.Exit(0)
	}

	fileinfo, err := os.Lstat(cwd)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(cwd)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println(realpath)

	} else {

		fmt.Println(cwd)
	}

	os.Exit(0)

}
