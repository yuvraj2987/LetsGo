package main

import "fmt"

func isUnique(str string) bool {
	if len(str) == 0 {
		return true
	}
	aMap := make(map[rune]bool)
	for _, word := range str {
		_, exist := aMap[word]
		if exist {
			return false
		} else {
			aMap[word] = true
		}

	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
