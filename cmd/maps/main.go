package main

import (
	"fmt"
)

func printMap(aMap map[int]int) {
	for key, val := range aMap {
		fmt.Printf("{%d} => %d\n", key, val)
	}

}

func toMap(slice []int) (aMap map[int]int) {
	aMap = make(map[int]int)

	for idx, value := range slice {
		aMap[idx] = value
	}
	return

}

func main() {
	fmt.Println(" Testing go maps")
	arr1 := [4]int{4, 3, 2, 1}
	for _, v := range arr1 {
		fmt.Printf("%d", v)
	}
	fmt.Println()
	map1 := toMap(arr1[:])
	printMap(map1)

}
