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
	oneD := [4]int{1, 2, 3, 4}
	fmt.Println("oneD array len ", len(oneD))

	threeD := [3][3][3]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}

	for _, v1 := range threeD {
		for _, v2 := range v1 {
			for l3, v3 := range v2 {
				fmt.Printf("{%v}=>%d\t", l3, v3)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
