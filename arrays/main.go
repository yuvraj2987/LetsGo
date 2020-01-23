package main

import (
	"fmt"
)

func main() {
	fmt.Println("==== Testing array package =====")
	q := [3]int{1, 2, 3}
	q = [3]int{2, 3, 4}

	for i, v := range q {
		fmt.Printf("{%v} -> %v\n", i, v)
	}

	// If array elements are comparable arrays are comparable

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}

	fmt.Println("a == b", a == b, "\ta == c", a == c, "\t	b == c", b == c)
}
