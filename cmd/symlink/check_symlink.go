package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("Invalid number of arguments!!!")
		os.Exit(1)
	}

	file := arg[1]

	fileinfo, err := os.Lstat(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filemode := fileinfo.Mode()

	if filemode&os.ModeSymlink != 0 {
		fmt.Println(file, " is a symlink")
		realpath, err := filepath.EvalSymlinks(file)
		if err != nil {
			fmt.Println("Failed to evaluate symlink: ", err.Error())
			os.Exit(1)
		} else {
			fmt.Println("Real path: ", realpath)
		}
	} else {
		fmt.Println(file, " is not a symlink")
	}

	os.Exit(0)

}
